package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const ENTITY_COUNT = 64000
const RESULTS_COUNT = ENTITY_COUNT / 100

type Vector struct {
	X int32
	Y int32
	Z int32
}

type Results struct {
	arr   []int32
	count int
}

func resultsInit() Results {
	return Results{
		arr:   make([]int32, RESULTS_COUNT),
		count: 0,
	}
}

func (r *Results) clear() {
	r.count = 0
}

func (r *Results) add(v int32) {
	if r.count == RESULTS_COUNT {
		r.arr[0] = v
		r.count = 1
		return
	}

	r.arr[r.count] = v
	r.count++
}

func (r *Results) print() {
	fmt.Printf("results: %#v\n", r.arr)
}

func makeVector(src Vector) *Vector {
	return &Vector{
		X: src.X,
		Y: src.Y,
		Z: src.Z,
	}
}

func shouldSkip(velX, velZ int32) bool {
	return velX > 300000/10 || velZ < 300000/10
}

func Benchmark_AOS_SOA(b *testing.B) {
	positions := make([]Vector, ENTITY_COUNT)
	for pi := range positions {
		pos := &positions[pi]
		pos.X = rand.Int31n(300000)
		pos.Y = rand.Int31n(300000)
		pos.Z = rand.Int31n(300000)
	}

	velocities := make([]Vector, ENTITY_COUNT)
	for vi := range velocities {
		vel := &velocities[vi]
		vel.X = rand.Int31n(300000)
		vel.Y = rand.Int31n(300000)
		vel.Z = rand.Int31n(300000)
	}

	randomVelZ := int32(300000 / 10)
	randomVelX := int32(300000 / 10)
	randomVelY := int32(300000 / 10)

	b.Run("Array Of Structs Pointers", func(b *testing.B) {
		type Meta struct {
			Id   string
			Hash string
			Time uint
		}

		type Entity struct {
			Position *Vector
			Meta     *Meta
			Velocity *Vector
		}
		entitiesArray := make([]Entity, ENTITY_COUNT)
		for ei := range entitiesArray {
			entitiesArray[ei].Position = makeVector(positions[ei])
			entitiesArray[ei].Velocity = makeVector(velocities[ei])
			entitiesArray[ei].Meta = &Meta{
				Id:   fmt.Sprintf("id %v", positions[ei].Y),
				Hash: fmt.Sprintf("hash %v", velocities[ei].Z),
				Time: uint(rand.Uint32()),
			}
		}

		results := resultsInit()
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

	b.Run("Array Of Structs", func(b *testing.B) {
		type Meta struct {
			Id   string
			Hash string
			Time uint
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
			}
		}

		results := resultsInit()
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

	b.Run("Struct Of Arrays", func(b *testing.B) {
		// type Entity struct {
		// 	PositionY int32
		// 	VelocityZ int32
		// 	VelocityY int32
		// 	VelocityX int32
		// }

		// entities := make([]Entity, ENTITY_COUNT)

		type AllEntities struct {
			PositionY []int32
			VelocityZ []int32
			VelocityY []int32
			VelocityX []int32
		}

		xPositions := make([]int32, ENTITY_COUNT)
		entities := AllEntities{
			PositionY: make([]int32, ENTITY_COUNT),
			VelocityZ: make([]int32, ENTITY_COUNT),
			VelocityY: make([]int32, ENTITY_COUNT),
			VelocityX: make([]int32, ENTITY_COUNT),
		}

		for i := 0; i < ENTITY_COUNT; i++ {
			entities.PositionY[i] = positions[i].Y
			entities.VelocityZ[i] = velocities[i].Z
			entities.VelocityY[i] = velocities[i].Y
			entities.VelocityX[i] = velocities[i].X

			xPositions[i] = positions[i].X
		}

		results := resultsInit()
		// --RUN--
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			results.clear()
			sum := int32(0)
			for esi := 0; esi < ENTITY_COUNT; esi++ {
				if xPositions[esi]%2 == 0 {
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

	// fmt.Println(sum2)
}

// func Test_AOS_SOA(t *testing.T) {
// 	type AbilityPickRates struct {
// 		ability string
// 		rates   []uint
// 	}

// 	type Build struct {
// 		buildID          string
// 		matches          int32
// 		wins             int32
// 		earlyGameItems   string
// 		abilityPickRates []AbilityPickRates
// 	}

// 	BUILDS_COUNT := 1
// 	ABILITIES_COUNT := 2
// 	RATES_COUNT := 3

// 	input := []Build{}
// 	expected := [][]uint{}

// 	for i := 0; i < BUILDS_COUNT; i++ {
// 		abilityPickRates := []AbilityPickRates{}
// 		for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
// 			randomRates := arrayOfRandomNumbers1To100(RATES_COUNT)
// 			abilityPickRates = append(abilityPickRates, AbilityPickRates{
// 				ability: fmt.Sprintf("#%v ability", ab_i+1),
// 				rates:   randomRates,
// 			})
// 		}

// 		input = append(input, Build{
// 			buildID:          fmt.Sprintf("%v ID", rand.Int63n(200)),
// 			matches:          rand.Int63n(400),
// 			wins:             rand.Int63n(300),
// 			earlyGameItems:   fmt.Sprintf("%v items", rand.Int63n(200)),
// 			abilityPickRates: abilityPickRates,
// 		})
// 	}

// 	for i := 0; i < BUILDS_COUNT; i++ {
// 		inputBuild := input[i]
// 		expectedBuild := []uint{}
// 		for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
// 			max := 0
// 			for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
// 				rate := inputBuild.abilityPickRates[ab_i].rates[lvl_i]
// 				max = maxInt(max, int32(rate))
// 			}

// 			expectedBuild = append(expectedBuild, uint(max))
// 		}

// 		expected = append(expected, expectedBuild)
// 	}

// 	// TESTS
// 	t.Run("#1", func(t *testing.T) {
// 		result := [][]uint{}
// 		for i := 0; i < BUILDS_COUNT; i++ {
// 			inputBuild := input[i]
// 			build := []uint{}
// 			for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
// 				max := 0
// 				for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
// 					rate := inputBuild.abilityPickRates[ab_i].rates[lvl_i]
// 					max = maxInt(max, int32(rate))
// 				}

// 				build = append(build, uint(max))
// 			}

// 			result = append(result, build)
// 		}

// 		assert.Equal(t, expected, result)
// 	})

// }

func init() {
	// rand.Seed(time.Now().UnixNano()) // Seed the random number generator
}
