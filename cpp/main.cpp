#include <cstdio>
#include <iostream>
#include <chrono>
#include <string>
#include <fstream>
#include <nlohmann/json.hpp>

using namespace std;
using namespace nlohmann;

const string INPUT_JSON_LOCATION = "../go/aos_soa_test_input.json";

#define BENCH_FUNC(_func_name, _block) \
    auto start = chrono::high_resolution_clock::now();\
    _block \
    auto end = chrono::high_resolution_clock::now(); \
    auto duration = chrono::duration_cast<chrono::microseconds>(end - start).count();\
    cout << _func_name << " execution Time: " << duration << " microsec." << endl

struct Results {
	int32_t *arr;
	int count;
	int cap;

    Results(int cap);
    void clear();
    void add(int32_t v);
};

Results::Results(int cap) {
    this->cap = cap;
    this->arr = new int32_t[cap];
}

void Results::clear() {
	this->count = 0;
}

void Results::add(int32_t v) {
	if (this->count == this->cap) {
		this->arr[0] = v;
		this->count = 1;
		return;
	}

	this->arr[this->count++] = v;
}

struct Vector {
    int32_t X;
    int32_t Y;
    int32_t Z;
};

struct  InnerTrash{
    Vector V1;
    Vector V2;
    string S1;
    Vector V3;
};

struct MetaStruct {
    string Id;
    string Hash;
    uint32_t Time;
    InnerTrash Trash;
};

struct Entity {
    Vector Position;
    MetaStruct Meta;
    Vector Velocity;
};

bool shouldSkip(int32_t velX, int32_t velZ) {
	// val := velX > 0 || velZ > 0
	// if !val {
	// 	panic("error")
	// }
	// return val
	return false;
}

void aosBench(
	int entitiesCount, vector<Vector> positions, vector<Vector> velocities,
	int32_t randomVelX, int32_t randomVelY, int32_t randomVelZ
) {
    char buffer[128];
    auto ENTITY_COUNT = entitiesCount;
    Entity *entitiesArray = new Entity[ENTITY_COUNT];

    for (auto ei = 0; ei < ENTITY_COUNT; ei++) {
        entitiesArray[ei].Position = positions[ei];
        entitiesArray[ei].Velocity = velocities[ei];
        snprintf(buffer, sizeof(buffer), "id %d", positions[ei].Y);
        entitiesArray[ei].Meta.Id = string(buffer);
        snprintf(buffer, sizeof(buffer), "hash %d", positions[ei].Y);
        entitiesArray[ei].Meta.Hash = string(buffer);
        // TODO (delphifeel): random here
        entitiesArray[ei].Meta.Time = ei % 333;

        auto Trash = &entitiesArray[ei].Meta.Trash;
        Trash->V1 = positions[ei];
        Trash->V2 = positions[ei];
        snprintf(buffer, sizeof(buffer), "S1 %d", positions[ei].Z);
        Trash->S1 = string(buffer);
        Trash->V3 = positions[ei];
    }

    Results results(ENTITY_COUNT / 2);
    int32_t sum = 0;
    BENCH_FUNC("aosBench", {
        for (auto i = 0; i < 1000; i++) {
            // results.clear();
            for (auto eai = 0; eai < ENTITY_COUNT; eai++) {
                if (entitiesArray[eai].Position.X%2 == 0) {
                    continue;
                }
                if (shouldSkip(entitiesArray[eai].Velocity.X, entitiesArray[eai].Velocity.Z)) {
                    continue;
                }
                sum += entitiesArray[eai].Position.Y;
                if (entitiesArray[eai].Velocity.Z > randomVelZ) {
                    sum -= entitiesArray[eai].Velocity.X;
                    continue;
                }
                if (entitiesArray[eai].Velocity.Y > randomVelY) {
                    if (eai != ENTITY_COUNT) {
                        entitiesArray[eai+1].Velocity.Y = entitiesArray[eai].Velocity.Y;
                    }
                    sum *= 2;
                    continue;
                }
                if (entitiesArray[eai].Velocity.X < randomVelX) {
                    if (eai != ENTITY_COUNT) {
                        entitiesArray[eai+1].Velocity.X = entitiesArray[eai].Velocity.X;
                    }
                    sum /= 2;
                    continue;
                }

                // results.add(sum);
            }
        }
    });

    cout << "sum: " << sum << endl;

		// results.print()
}



int main(void) {
    ifstream f(INPUT_JSON_LOCATION);
    json input = json::parse(f);

    
    vector<Vector> positions;
    for (auto &i : input["Positions"]) {
        positions.push_back({i["X"], i["Y"], i["Z"]});
    }

    vector<Vector> velocities;
    for (auto &i : input["Velocities"]) {
        velocities.push_back({i["X"], i["Y"], i["Z"]});
    }
	auto randomVelX = input["RandomVelX"];
	auto randomVelY = input["RandomVelY"];
	auto randomVelZ = input["RandomVelZ"];

    aosBench(1000000, positions, velocities, randomVelX, randomVelY, randomVelZ);
}