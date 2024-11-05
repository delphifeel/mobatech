package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const BUILDS_COUNT = 4
const ABILITIES_COUNT = 4
const RATES_COUNT = 16

func Benchmark_AbilitiesRates(b *testing.B) {
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

	// prepare := func() {

	// }

	// TESTS
	b.Run("#1    ", func(b *testing.B) {
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

	b.Run("#2    ", func(b *testing.B) {
		// #2 prep
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

	b.Run("#1 no allocs", func(b *testing.B) {
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

	b.Run("#2 no allocs", func(b *testing.B) {
		// #2 prep
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

func Test_AbilitiesRates(t *testing.T) {
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

	BUILDS_COUNT := 1
	ABILITIES_COUNT := 2
	RATES_COUNT := 3

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

	// TESTS
	t.Run("#1", func(t *testing.T) {
		result := [][]uint{}
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

		assert.Equal(t, expected, result)
	})

	// #2 prep
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
	ind := func(i, lvl_i, ab_i int) int {
		return i*RATES_COUNT*ABILITIES_COUNT + lvl_i*ABILITIES_COUNT + ab_i
	}

	t.Run("#2", func(t *testing.T) {
		result := [][]uint{}
		for i := 0; i < BUILDS_COUNT; i++ {
			build := []uint{}
			for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
				max := 0
				for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
					rate := buildToRateToAbility[ind(i, lvl_i, ab_i)]
					max = maxInt(max, int(rate))
				}

				build = append(build, uint(max))
			}

			result = append(result, build)
		}

		assert.Equal(t, expected, result)
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

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
}

// b.Run("#2 small", func(b *testing.B) {
// 	// #2 prep
// 	buildToRateToAbility := make([]uint8, BUILDS_COUNT*RATES_COUNT*ABILITIES_COUNT)
// 	for i := 0; i < BUILDS_COUNT; i++ {
// 		inputBuild := input[i]
// 		for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
// 			for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
// 				buildToRateToAbility[i*RATES_COUNT*ABILITIES_COUNT+lvl_i*ABILITIES_COUNT+ab_i] =
// 					uint8(inputBuild.abilityPickRates[ab_i].rates[lvl_i])
// 			}
// 		}
// 	}

// 	result := [][]uint{}
// 	for i := 0; i < b.N; i++ {
// 		for i := 0; i < BUILDS_COUNT; i++ {
// 			build := []uint{}
// 			for lvl_i := 0; lvl_i < RATES_COUNT; lvl_i++ {
// 				var max uint = 0
// 				for ab_i := 0; ab_i < ABILITIES_COUNT; ab_i++ {
// 					rate := buildToRateToAbility[i*RATES_COUNT*ABILITIES_COUNT+lvl_i*ABILITIES_COUNT+ab_i]
// 					if uint(rate) > max {
// 						max = uint(rate)
// 					}
// 				}

// 				build = append(build, max)
// 			}

// 			result = append(result, build)
// 		}
// 	}

// 	_ = result
// })
