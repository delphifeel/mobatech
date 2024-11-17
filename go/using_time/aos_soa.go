package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

const MAX_ENTITY_COUNT = 3_200_000

type Vector struct {
	X int32
	Y int32
	Z int32
}

type Results struct {
	arr   []int32
	count int
	cap   int
}

func resultsInit(cap int) Results {
	return Results{
		arr:   make([]int32, cap),
		count: 0,
		cap:   cap,
	}
}

func (r *Results) clear() {
	r.count = 0
}

func (r *Results) add(v int32) {
	if r.count == r.cap {
		r.arr[0] = v
		r.count = 1
		return
	}

	r.arr[r.count] = v
	r.count++
}

// func (r *Results) print() {
// 	fmt.Printf("results: %#v\n", r.arr)
// }

func makeVector(src Vector) *Vector {
	return &Vector{
		X: src.X,
		Y: src.Y,
		Z: src.Z,
	}
}

func shouldSkip(velX, velZ int32) bool {
	// val := velX > 0 || velZ > 0
	// if !val {
	// 	panic("error")
	// }
	// return val
	return false
}

func aosPtrsBench(
	entitiesCount int, b *testing.B, positions []Vector, velocities []Vector,
	randomVelX int32, randomVelY int32, randomVelZ int32) {

	b.Run(fmt.Sprintf("[%v] Array Of Structs Pointers", entitiesCount), func(b *testing.B) {
		ENTITY_COUNT := entitiesCount

		type InnerTrash struct {
			V1 *Vector
			V2 *Vector
			S1 string
			V3 *Vector
		}

		type Meta struct {
			Id    string
			Hash  string
			Time  uint
			Trash *InnerTrash
		}

		type Entity struct {
			Position *Vector
			Meta     *Meta
			Velocity *Vector
		}
		entitiesArray := make([]Entity, ENTITY_COUNT)
		for ei := range entitiesArray {
			entitiesArray[ei].Position = makeVector(positions[ei])
			entitiesArray[ei].Meta = &Meta{
				Id:   fmt.Sprintf("id %v", positions[ei].Y),
				Hash: fmt.Sprintf("hash %v", velocities[ei].Z),
				Time: uint(rand.Uint32()),
				Trash: &InnerTrash{
					V1: makeVector(positions[ei]),
					V2: makeVector(positions[ei]),
					S1: fmt.Sprintf("S1 %v", positions[ei].Z),
					V3: makeVector(positions[ei]),
				},
			}
			entitiesArray[ei].Velocity = makeVector(velocities[ei])
		}

		results := resultsInit(ENTITY_COUNT / 2)
		// --RUN--
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			results.clear()
			sum := int32(0)
			for eai := range entitiesArray {
				if entitiesArray[eai].Position.X%2 == 0 {
					continue
				}
				if shouldSkip(entitiesArray[eai].Velocity.X, entitiesArray[eai].Velocity.Z) {
					continue
				}
				sum += entitiesArray[eai].Position.Y
				if entitiesArray[eai].Velocity.Z > randomVelZ {
					sum -= entitiesArray[eai].Velocity.X
					continue
				}
				if entitiesArray[eai].Velocity.Y > randomVelY {
					if eai != len(entitiesArray) {
						entitiesArray[eai+1].Velocity.Y = entitiesArray[eai].Velocity.Y
					}
					sum *= 2
					continue
				}
				if entitiesArray[eai].Velocity.X < randomVelX {
					if eai != len(entitiesArray) {
						entitiesArray[eai+1].Velocity.X = entitiesArray[eai].Velocity.X
					}
					sum /= 2
					continue
				}
				results.add(sum)
			}
		}

		// results.print()
	})
}

func aosBench(
	entitiesCount int, positions []Vector, velocities []Vector,
	randomVelX int32, randomVelY int32, randomVelZ int32) {

	ENTITY_COUNT := entitiesCount
	type InnerTrash struct {
		V1 Vector
		V2 Vector
		S1 string
		V3 Vector
	}

	type Meta struct {
		Id    string
		Hash  string
		Time  uint
		Trash InnerTrash
	}

	type Entity struct {
		Position Vector
		Meta     Meta
		Velocity Vector
	}
	entitiesArray := make([]Entity, ENTITY_COUNT)
	for ei := range entitiesArray {
		entitiesArray[ei].Position = positions[ei]
		entitiesArray[ei].Velocity = velocities[ei]
		entitiesArray[ei].Meta = Meta{
			Id:   fmt.Sprintf("id %v", positions[ei].Y),
			Hash: fmt.Sprintf("hash %v", velocities[ei].Z),
			Time: uint(rand.Uint32()),
			Trash: InnerTrash{
				V1: positions[ei],
				V2: positions[ei],
				S1: fmt.Sprintf("S1 %v", positions[ei].Z),
				V3: positions[ei],
			},
		}
	}

	results := resultsInit(ENTITY_COUNT / 2)
	// --RUN--
	start := time.Now()
	sum := int32(0)
	for i := 0; i < 1000; i++ {
		results.clear()
		for eai := range entitiesArray {
			if entitiesArray[eai].Position.X%2 == 0 {
				continue
			}
			if shouldSkip(entitiesArray[eai].Velocity.X, entitiesArray[eai].Velocity.Z) {
				continue
			}
			sum += entitiesArray[eai].Position.Y
			if entitiesArray[eai].Velocity.Z > randomVelZ {
				sum -= entitiesArray[eai].Velocity.X
				continue
			}
			if entitiesArray[eai].Velocity.Y > randomVelY {
				if eai != len(entitiesArray) {
					entitiesArray[eai+1].Velocity.Y = entitiesArray[eai].Velocity.Y
				}
				sum *= 2
				continue
			}
			if entitiesArray[eai].Velocity.X < randomVelX {
				if eai != len(entitiesArray) {
					entitiesArray[eai+1].Velocity.X = entitiesArray[eai].Velocity.X
				}
				sum /= 2
				continue
			}

			results.add(sum)
		}
	}

	// Calculate the elapsed time in microseconds
	elapsed := time.Since(start).Microseconds()

	// Print the elapsed time in microseconds
	fmt.Printf("Function execution time: %d microseconds\n", elapsed)
	fmt.Printf("sum: %v\n", sum)

	// results.print()
}

func soaBench(
	entitiesCount int, b *testing.B, positions []Vector, velocities []Vector,
	randomVelX int32, randomVelY int32, randomVelZ int32) {

	b.Run(fmt.Sprintf("[%v] Struct Of Arrays", entitiesCount), func(b *testing.B) {
		ENTITY_COUNT := entitiesCount

		type InnerTrash struct {
			V1 Vector
			V2 Vector
			S1 string
			V3 Vector
		}

		type Meta struct {
			Id    string
			Hash  string
			Time  uint
			Trash InnerTrash
		}

		type AllEntities struct {
			// PositionX []int32
			VelocityX []int32
			VelocityZ []int32
			PositionY []int32
			VelocityY []int32
			Meta      []Meta
		}

		PositionX := make([]int32, ENTITY_COUNT)
		entities := AllEntities{
			VelocityX: make([]int32, ENTITY_COUNT),
			VelocityZ: make([]int32, ENTITY_COUNT),
			PositionY: make([]int32, ENTITY_COUNT),
			VelocityY: make([]int32, ENTITY_COUNT),
			Meta:      make([]Meta, ENTITY_COUNT),
		}

		for i := 0; i < ENTITY_COUNT; i++ {
			PositionX[i] = positions[i].X
			entities.VelocityX[i] = velocities[i].X
			entities.VelocityZ[i] = velocities[i].Z
			entities.PositionY[i] = positions[i].Y
			entities.VelocityY[i] = velocities[i].Y

			entities.Meta[i] = Meta{
				Id:   fmt.Sprintf("id %v", positions[i].Y),
				Hash: fmt.Sprintf("hash %v", velocities[i].Z),
				Time: uint(rand.Uint32()),
				Trash: InnerTrash{
					V1: positions[i],
					V2: positions[i],
					S1: fmt.Sprintf("S1 %v", positions[i].Z),
					V3: positions[i],
				},
			}
		}

		results := resultsInit(ENTITY_COUNT / 2)
		// --RUN--
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			results.clear()
			sum := int32(0)
			for esi := 0; esi < ENTITY_COUNT; esi++ {
				if PositionX[esi]%2 == 0 {
					continue
				}
				if shouldSkip(entities.VelocityX[esi], entities.VelocityZ[esi]) {
					continue
				}

				sum += entities.PositionY[esi]
				if entities.VelocityZ[esi] > randomVelZ {
					sum -= entities.VelocityX[esi]
					continue
				}
				if entities.VelocityY[esi] > randomVelY {
					if esi != ENTITY_COUNT {
						entities.VelocityY[esi+1] = entities.VelocityY[esi]
					}
					sum *= 2
					continue
				}
				if entities.VelocityX[esi] < randomVelX {
					if esi != ENTITY_COUNT {
						entities.VelocityX[esi+1] = entities.VelocityX[esi]
					}
					sum /= 2
					continue
				}
				results.add(sum)
			}
		}

		// results.print()
	})
}

func writeJsonToFile(fileName string, value any) {
	jsonData, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}

	// Open or create the file
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func readJsonFromFile(outPtr any, fileName string) {
	// Open the JSON file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode JSON data into the struct
	err = json.NewDecoder(file).Decode(outPtr)
	if err != nil {
		panic(err)
	}
}

type Input struct {
	Positions  []Vector
	Velocities []Vector
	RandomVelX int32
	RandomVelY int32
	RandomVelZ int32
}

func main() {
	// createNewInput()
	// return

	var input Input
	readJsonFromFile(&input, "aos_soa_test_input.json")

	positions := input.Positions
	velocities := input.Velocities
	randomVelX := input.RandomVelX
	randomVelY := input.RandomVelY
	randomVelZ := input.RandomVelZ

	// fmt.Println(randomVelX)
	// fmt.Println(randomVelY)
	// fmt.Println(randomVelZ)

	// aosPtrsBench(1000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// aosBench(1000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// soaBench(1000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)

	// aosPtrsBench(10000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// aosBench(10000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// soaBench(10000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)

	// aosPtrsBench(100_000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// aosBench(100_000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// soaBench(100_000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)

	// aosPtrsBench(1_000_000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)
	aosBench(1_000_000, positions, velocities, randomVelX, randomVelY, randomVelZ)
	// soaBench(1_000_000, b, positions, velocities, randomVelX, randomVelY, randomVelZ)

	// fmt.Println(sum2)

}
