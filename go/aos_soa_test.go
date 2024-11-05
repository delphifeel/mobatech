package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Benchmark_AOS_SOA(b *testing.B) {
	const ENTITY_COUNT = 20000

	type Vector struct {
		X int
		Y int
		Z int
	}

	positions := make([]Vector, ENTITY_COUNT)
	for pi := range positions {
		pos := &positions[pi]
		pos.X = rand.Intn(3000)
		pos.Y = rand.Intn(3000)
		pos.Z = rand.Intn(3000)
	}

	velocities := make([]Vector, ENTITY_COUNT)
	for vi := range velocities {
		vel := &velocities[vi]
		vel.X = rand.Intn(3000)
		vel.Y = rand.Intn(3000)
		vel.Z = rand.Intn(3000)
	}

	// TEST 1 PREPARE

	type Entity struct {
		Position Vector
		Velocity Vector
		Meta     string
	}
	entitiesArray := make([]Entity, ENTITY_COUNT)
	for ei := range entitiesArray {
		entitiesArray[ei].Position = positions[ei]
		entitiesArray[ei].Velocity = velocities[ei]
		entitiesArray[ei].Meta = fmt.Sprintf("pos vel %v", ei)
	}

	// fmt.Printf("entitiesArray: %#v\n", entitiesArray)

	sum1 := 0
	b.Run("Array Of Structs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sum := 0
			for eai := range entitiesArray {
				sum += entitiesArray[eai].Position.Y
				if entitiesArray[eai].Velocity.Z > 1000 {
					sum -= entitiesArray[eai].Velocity.Z
				}
			}
			sum1 = sum
		}
	})
	fmt.Println(sum1)

	// TEST 2 PREPARE

	type Entities struct {
		PositionY []int
		VelocityZ []int
	}

	internalArray := make([]int, ENTITY_COUNT*2)

	entitiesStruct := Entities{
		PositionY: internalArray[:ENTITY_COUNT],
		VelocityZ: internalArray[ENTITY_COUNT:],
	}

	for i := 0; i < ENTITY_COUNT; i++ {
		entitiesStruct.PositionY[i] = positions[i].Y
		entitiesStruct.VelocityZ[i] = velocities[i].Z
	}
	// fmt.Printf("entitiesStruct: %#v\n", entitiesStruct)

	sum2 := 0
	b.Run("Struct Of Arrays", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sum := 0
			for esi := 0; esi < ENTITY_COUNT; esi++ {
				sum += entitiesStruct.PositionY[esi]
				if entitiesStruct.VelocityZ[esi] > 1000 {
					sum -= entitiesStruct.VelocityZ[esi]
				}
			}

			sum2 = sum
		}
	})

	fmt.Println(sum2)
}

// func Test_AOS_SOA(t *testing.T) {
// 	type AbilityPickRates struct {
// 		ability string
// 		rates   []uint
// 	}

// 	type Build struct {
// 		buildID          string
// 		matches          int
// 		wins             int
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
// 			buildID:          fmt.Sprintf("%v ID", rand.Intn(200)),
// 			matches:          rand.Intn(400),
// 			wins:             rand.Intn(300),
// 			earlyGameItems:   fmt.Sprintf("%v items", rand.Intn(200)),
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
// 				max = maxInt(max, int(rate))
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
// 					max = maxInt(max, int(rate))
// 				}

// 				build = append(build, uint(max))
// 			}

// 			result = append(result, build)
// 		}

// 		assert.Equal(t, expected, result)
// 	})

// }

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
}
