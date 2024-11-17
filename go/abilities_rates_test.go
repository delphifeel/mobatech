package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// const BUILDS_COUNT = 4
// const ABILITIES_COUNT = 4
// const RATES_COUNT = 16

type AbilityPickRates struct {
	ability string
	rates   []uint
}

type Build struct {
	buildID          string
	matches          int
	wins             int
	earlyGameItems   string
	abilityPickRates []AbilityPickRates
}

func noAllocFastBench(b *testing.B, input []Build, BUILDS_COUNT int, RATES_COUNT int, ABILITIES_COUNT int) {
	b.Run("noalloc fast", func(b *testing.B) {
		buildToRateToAbility := make([]uint, BUILDS_COUNT*RATES_COUNT*ABILITIES_COUNT)
		for i := 0; i < BUILDS_COUNT; i++ {
			inputBuild := input[i]
			for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
				for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
					buildToRateToAbility[i*RATES_COUNT*ABILITIES_COUNT+lvl_i*ABILITIES_COUNT+ab_i] =
						inputBuild.abilityPickRates[ab_i].rates[lvl_i]
				}
			}
		}

		b.ResetTimer()
		result := make([][]uint, BUILDS_COUNT)
		for b_i := 0; b_i < BUILDS_COUNT; b_i++ {
			result[b_i] = make([]uint, RATES_COUNT)
		}

		for i := 0; i < b.N; i++ {
			for i := 0; i < BUILDS_COUNT; i++ {
				for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
					max := 0
					for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
						rate := buildToRateToAbility[i*RATES_COUNT*ABILITIES_COUNT+lvl_i*ABILITIES_COUNT+ab_i]
						max = maxInt(max, int(rate))
					}

					result[i][lvl_i] = uint(max)
				}
			}
		}

		_ = result
	})
}

func noAllocSlowBench(b *testing.B, input []Build, BUILDS_COUNT int, RATES_COUNT int, ABILITIES_COUNT int) {
	b.Run("noalloc slow", func(b *testing.B) {

		result := make([][]uint, BUILDS_COUNT)
		for b_i := 0; b_i < BUILDS_COUNT; b_i++ {
			result[b_i] = make([]uint, RATES_COUNT)
		}

		for i := 0; i < b.N; i++ {
			for i := 0; i < BUILDS_COUNT; i++ {
				inputBuild := input[i]
				for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
					max := 0
					for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
						rate := inputBuild.abilityPickRates[ab_i].rates[lvl_i]
						max = maxInt(max, int(rate))
					}

					// build = append(build, uint(max))
					result[i][lvl_i] = uint(max)
				}

			}
		}

		_ = result
	})
}

func slowBench(b *testing.B, input []Build, BUILDS_COUNT int, RATES_COUNT int, ABILITIES_COUNT int) {
	b.Run("slow    ", func(b *testing.B) {
		result := [][]uint{}

		for i := 0; i < b.N; i++ {
			for i := 0; i < BUILDS_COUNT; i++ {
				inputBuild := input[i]
				build := []uint{}
				for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
					max := 0
					for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
						rate := inputBuild.abilityPickRates[ab_i].rates[lvl_i]
						max = maxInt(max, int(rate))
					}

					build = append(build, uint(max))
				}

				result = append(result, build)
			}
		}

		_ = result
	})
}

func fastBench(b *testing.B, input []Build, BUILDS_COUNT int, RATES_COUNT int, ABILITIES_COUNT int) {
	b.Run("fast    ", func(b *testing.B) {

		buildToRateToAbility := make([]uint, BUILDS_COUNT*RATES_COUNT*ABILITIES_COUNT)
		for i := 0; i < BUILDS_COUNT; i++ {
			inputBuild := input[i]
			for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
				for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
					buildToRateToAbility[i*RATES_COUNT*ABILITIES_COUNT+lvl_i*ABILITIES_COUNT+ab_i] =
						inputBuild.abilityPickRates[ab_i].rates[lvl_i]
				}
			}
		}

		b.ResetTimer()
		result := [][]uint{}
		for i := 0; i < b.N; i++ {
			for i := 0; i < BUILDS_COUNT; i++ {
				build := []uint{}
				for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
					max := 0
					for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
						rate := buildToRateToAbility[i*RATES_COUNT*ABILITIES_COUNT+lvl_i*ABILITIES_COUNT+ab_i]
						max = maxInt(max, int(rate))
					}

					build = append(build, uint(max))
				}

				result = append(result, build)
			}
		}

		_ = result
	})
}

func Benchmark_AbilitiesRates(b *testing.B) {

	// type BenchConfig struct {
	// 	input
	// }

	prepare := func(BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT int) []Build {
		input := []Build{}
		expected := [][]uint{}

		for i := 0; i < BUILDS_COUNT; i++ {
			abilityPickRates := []AbilityPickRates{}
			for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
				randomRates := arrayOfRandomNumbers1To100(RATES_COUNT)
				abilityPickRates = append(abilityPickRates, AbilityPickRates{
					ability: fmt.Sprintf("#%v ability", ab_i+1),
					rates:   randomRates,
				})
			}

			input = append(input, Build{
				buildID:          fmt.Sprintf("%v ID", rand.Intn(200)),
				matches:          rand.Intn(400),
				wins:             rand.Intn(300),
				earlyGameItems:   fmt.Sprintf("%v items", rand.Intn(200)),
				abilityPickRates: abilityPickRates,
			})
		}

		for i := 0; i < BUILDS_COUNT; i++ {
			inputBuild := input[i]
			expectedBuild := []uint{}
			for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
				max := 0
				for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
					rate := inputBuild.abilityPickRates[ab_i].rates[lvl_i]
					max = maxInt(max, int(rate))
				}

				expectedBuild = append(expectedBuild, uint(max))
			}

			expected = append(expected, expectedBuild)
		}

		_ = expected

		return input
	}

	b.Run("x1", func(b *testing.B) {
		BUILDS_COUNT := 4
		ABILITIES_COUNT := 4
		RATES_COUNT := 16
		input := prepare(BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT)

		b.ResetTimer()
		slowBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		fastBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		noAllocSlowBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		noAllocFastBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
	})
	b.Run("x5", func(b *testing.B) {
		BUILDS_COUNT := 4 * 5
		ABILITIES_COUNT := 4 * 5
		RATES_COUNT := 16 * 5
		input := prepare(BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT)

		b.ResetTimer()
		slowBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		fastBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		noAllocSlowBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		noAllocFastBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
	})
	b.Run("x25", func(b *testing.B) {
		BUILDS_COUNT := 4 * 25
		ABILITIES_COUNT := 4 * 25
		RATES_COUNT := 16 * 25
		input := prepare(BUILDS_COUNT, ABILITIES_COUNT, RATES_COUNT)

		b.ResetTimer()
		slowBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		fastBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		noAllocSlowBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
		noAllocFastBench(b, input, BUILDS_COUNT, RATES_COUNT, ABILITIES_COUNT)
	})
}

func arrayOfRandomNumbers1To100(size int) []uint {
	result := []uint{}
	for i := 0; i < size; i++ {
		result = append(result, uint(rand.Intn(100)+1))
	}
	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
