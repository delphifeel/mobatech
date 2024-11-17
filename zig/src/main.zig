const std = @import("std");
const debug = std.debug;
const fmt = std.fmt;
const time = std.time;
const Allocator = std.mem.Allocator;

const MAX_FILE_SIZE = 999_999_999;

const Vector = struct {
    X: i32,
    Y: i32,
    Z: i32,
};

const InnerTrash = struct {
    V1: Vector,
    V2: Vector,
    S1: []const u8,
    V3: Vector,
};

const Meta = struct {
    Id: []const u8,
    Hash: []const u8,
    Time: u32,
    Trash: InnerTrash,
};

const Entity = struct {
    Position: Vector,
    Meta: Meta,
    Velocity: Vector,
};

const Input = struct {
    Positions: []Vector,
	Velocities: []Vector,
	RandomVelX: i32,
	RandomVelY: i32,
	RandomVelZ: i32,
};

inline fn shouldSkip(_: i32, _: i32) bool {
    return false;
}

fn aosBench(
    allocator: Allocator,
	entitiesCount: i32, positions: []const Vector, velocities: []const Vector,
	randomVelX: i32, randomVelY: i32, randomVelZ: i32,
) !void {
	var ENTITY_COUNT: usize = @intCast(entitiesCount);
	
    var entitiesArray = std.MultiArrayList(Entity){};
    try entitiesArray.ensureTotalCapacity(allocator, ENTITY_COUNT);
    for (0..ENTITY_COUNT) |ei| {
        var ei_u32: u32 = @intCast(ei);
        try entitiesArray.append(allocator, .{
            .Position = positions[ei],
            .Velocity = velocities[ei],
            .Meta = .{
                .Id =   try fmt.allocPrint(allocator, "id {}", .{positions[ei].Y}),
                .Hash = try fmt.allocPrint(allocator, "hash {}", .{velocities[ei].Z}),
                .Time = ei_u32 % 333,
                .Trash = .{
                    .V1 = positions[ei],
                    .V2 = positions[ei],
                    .S1 = try fmt.allocPrint(allocator, "S1 {}", .{positions[ei].Z}),
                    .V3 = positions[ei],
                },
            },
        });
    }

	var results = try Results.init(allocator, @divExact(entitiesCount, 2));
	// --RUN--
    var timer = try time.Timer.start();
	var sum: i32 = 0;
    var entitiesArraySlice = entitiesArray.slice();
	for (0..1000) |_| {
		results.clear();
        for (
            entitiesArraySlice.items(.Position), 
            entitiesArraySlice.items(.Velocity), 
            0..
        ) | *position, *velocity, eai | {
			if (@mod(position.X, 2) == 0) {
				continue;
			}
			if (shouldSkip(velocity.X, velocity.Z)) {
				continue;
			}
			sum += position.Y;
			if (velocity.Z > randomVelZ) {
				sum -= velocity.X;
				continue;
			}
			if (velocity.Y > randomVelY) {
				if (eai < ENTITY_COUNT - 1) {
                    var newEntity = entitiesArraySlice.get(eai + 1);
                    newEntity.Velocity.Y = velocity.Y;
					entitiesArraySlice.set(eai + 1, newEntity);
				}
				sum *= 2;
				continue;
			}
			if (velocity.X < randomVelX) {
				if (eai < ENTITY_COUNT - 1) {
					var newEntity = entitiesArraySlice.get(eai + 1);
                    newEntity.Velocity.X = velocity.X;
					entitiesArraySlice.set(eai + 1, newEntity);
				}
				sum = @divExact(sum, 2);
				continue;
			}

			results.add(sum);
		}
	}

    var timer_v = timer.lap();
    debug.print("Passed: {} ns\n", .{timer_v});
    debug.print("sum: {}\n", .{sum});

	// results.print()
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    debug.assert(gpa.detectLeaks() == false);
    var allocator = gpa.allocator();

    var fileContent = try std.fs.cwd().readFileAlloc(allocator, "../go/aos_soa_test_input.json", MAX_FILE_SIZE);
    var parseOptions = std.json.ParseOptions{
        .ignore_unknown_fields = true,
        .allocate = .alloc_if_needed,
    };
    var parsedInput = try std.json.parseFromSlice(Input, allocator, fileContent, parseOptions);
    defer parsedInput.deinit();

    var input = &parsedInput.value;

    try aosBench(allocator, 1_000_000, input.Positions, input.Velocities, input.RandomVelX, input.RandomVelY, input.RandomVelZ);
}

const Results = struct {
	arr: []i32,
	count: i32,
	cap: i32,

    const Self = @This();

    pub fn init(allocator: Allocator, cap: i32) !Self {
        return .{
            .arr =  try allocator.alloc(i32, @intCast(cap)),
            .count = 0,
            .cap = cap,
        };
	}

    pub inline fn clear(self: *Self) void {
        self.count = 0;
    }

    pub inline fn add(self: *Self, v: i32) void {
        if (self.count == self.cap) {
            self.arr[0] = v;
            self.count = 1;
            return;
        }

        self.arr[@intCast(self.count)] = v;
        self.count += 1;
    }
};

// inline fn OOM() noreturn {
//     debug.panic("OOM", .{});
// }

test "simple test" {
}
