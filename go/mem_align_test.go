package main

import (
	"math/rand"
	"testing"
)

type Unoptimized struct {
	A int64
	B int32
	C int64
	D int32
}

type Optimized struct {
	A int64
	C int64
	B int32
	D int32
}

const ARRAY_SIZE = 3_000_00

const MAX_VALUE = 44_444_444

func Benchmark_MemAlign(b *testing.B) {
	randGen := rand.New(rand.NewSource(777))
	arrayUnopt := make([]Unoptimized, ARRAY_SIZE)
	for i := 0; i < ARRAY_SIZE; i++ {
		arrayUnopt[i].A = randGen.Int63n(MAX_VALUE)
		arrayUnopt[i].B = randGen.Int31n(MAX_VALUE)
		arrayUnopt[i].C = randGen.Int63n(MAX_VALUE)
		arrayUnopt[i].D = randGen.Int31n(MAX_VALUE)
	}

	// Fill an array with instances of Unoptimized structs and perform simple operations
	b.Run("Unoptimized", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := int64(0)
			for j := 0; j < len(arrayUnopt); j++ {
				// if arrayUnopt[j].C > MAX_VALUE/2 {
				// 	continue
				// }
				if arrayUnopt[j].D < MAX_VALUE/2 {
					sum = 0
				}
				sum += arrayUnopt[j].A + int64(arrayUnopt[j].B)
			}
		}
	})

	// Fill an array with instances of Optimized structs and perform simple operations
	b.Run("Optimized", func(b *testing.B) {
		arrayOpt := make([]Optimized, ARRAY_SIZE)
		for i := 0; i < ARRAY_SIZE; i++ {
			arrayOpt[i].A = arrayUnopt[i].A
			arrayOpt[i].B = arrayUnopt[i].B
			arrayOpt[i].C = arrayUnopt[i].C
			arrayOpt[i].D = arrayUnopt[i].D
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := int64(0)
			for j := 0; j < len(arrayOpt); j++ {
				// if arrayOpt[j].C > MAX_VALUE/2 {
				// 	continue
				// }
				if arrayOpt[j].D < MAX_VALUE/2 {
					sum = 0
				}
				sum += arrayOpt[j].A + int64(arrayOpt[j].B)
			}
		}
	})

}
