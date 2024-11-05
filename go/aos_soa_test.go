package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const ENTITY_COUNT = 100000
const RESULTS_COUNT = ENTITY_COUNT / 10

type Vector struct {
	X int64
	Y int64
	Z int64
}

func Benchmark_AOS_SOA(b *testing.B) {
	positions := make([]Vector, ENTITY_COUNT)
	for pi := range positions {
		pos := &positions[pi]
		pos.X = rand.Int63n(3000)
		pos.Y = rand.Int63n(3000)
		pos.Z = rand.Int63n(3000)
	}

	velocities := make([]Vector, ENTITY_COUNT)
	for vi := range velocities {
		vel := &velocities[vi]
		vel.X = rand.Int63n(3000)
		vel.Y = rand.Int63n(3000)
		vel.Z = rand.Int63n(3000)
	}

	b.Run("Array Of Structs", func(b *testing.B) {
		type Meta struct {
			Id   string
			Hash string
			Time uint
		}

		type Entity struct {
			Position Vector
			Velocity Vector
			Meta     *Meta
		}
		entitiesArray := make([]Entity, ENTITY_COUNT)
		for ei := range entitiesArray {
			entitiesArray[ei].Position = positions[ei]
			entitiesArray[ei].Velocity = velocities[ei]
			entitiesArray[ei].Meta = &Meta{
				Id:   fmt.Sprintf("id %v", positions[ei].Y),
				Hash: fmt.Sprintf("hash %v", velocities[ei].Z),
				Time: uint(rand.Uint32()),
			}
		}

		results := make([]int64, RESULTS_COUNT)

		// --RUN--
		for i := 0; i < b.N; i++ {
			sum := int64(0)
			for eai := range entitiesArray {
				sum += entitiesArray[eai].Position.Y
				if entitiesArray[eai].Velocity.Z > 1000 {
					sum -= entitiesArray[eai].Velocity.Z
				}
			}
			_ = sum
		}
	})

	b.Run("Struct Of Arrays", func(b *testing.B) {
		type Entity struct {
			PositionY int64
			VelocityZ int64
		}

		entities := make([]Entity, ENTITY_COUNT)

		for i := 0; i < ENTITY_COUNT; i++ {
			entities[i].PositionY = positions[i].Y
			entities[i].VelocityZ = velocities[i].Z
		}

		// --RUN--
		for i := 0; i < b.N; i++ {
			sum := int64(0)
			for esi := 0; esi < ENTITY_COUNT; esi++ {
				sum += entities[esi].PositionY
				if entities[esi].VelocityZ > 1000 {
					sum -= entities[esi].VelocityZ
				}
			}

			_ = sum
		}
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
// 		matches          int64
// 		wins             int64
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
// 				max = maxInt(max, int64(rate))
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
// 					max = maxInt(max, int64(rate))
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
