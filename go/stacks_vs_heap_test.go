package main

import (
	"math/rand"
	"testing"
)

type MoreInnerTrash struct {
	InnerValue   int32
	InnerAsFloat float32
	SomeStr      string
}

type InnerTrash struct {
	Value         int32
	AsFloat       float32
	EvenMoreTrash MoreInnerTrash
}

type TrashData struct {
	A          int64
	B          int32
	InnerTrash InnerTrash
	C          int64
	D          int32
}

const RAND_MAX = 4000

func newTrashData() *TrashData {
	return &TrashData{
		A: rand.Int63n(RAND_MAX),
		B: rand.Int31n(RAND_MAX),
		C: rand.Int63n(RAND_MAX),
		D: rand.Int31n(RAND_MAX),
		InnerTrash: InnerTrash{
			Value:   rand.Int31n(RAND_MAX),
			AsFloat: float32(rand.Int31n(RAND_MAX)),
		},
	}
}

const TRASH_COUNT = 100

func Benchmark_StackVsHeap(b *testing.B) {

	b.Run("x1 Stack", func(b *testing.B) {
		b.ResetTimer()
		sum := 0
		for bi := 0; bi < b.N; bi++ {
			for i := 0; i < TRASH_COUNT; i++ {
				trashItem := TrashData{
					A: rand.Int63n(RAND_MAX),
					B: rand.Int31n(RAND_MAX),
					C: rand.Int63n(RAND_MAX),
					D: rand.Int31n(RAND_MAX),
					InnerTrash: InnerTrash{
						Value:   rand.Int31n(RAND_MAX),
						AsFloat: float32(rand.Int31n(RAND_MAX)),
					},
				}
				if trashItem.InnerTrash.Value > 100 {
					sum++
				}
			}
		}
	})

	b.Run("x1 Heap", func(b *testing.B) {
		// newTrashDataEx := func() *TrashData {
		// 	return &TrashData{
		// 		A: rand.Int63n(RAND_MAX),
		// 		B: rand.Int31n(RAND_MAX),
		// 		C: rand.Int63n(RAND_MAX),
		// 		D: rand.Int31n(RAND_MAX),
		// 		InnerTrash: InnerTrash{
		// 			Value:   rand.Int31n(RAND_MAX),
		// 			AsFloat: float32(rand.Int31n(RAND_MAX)),
		// 		},
		// 	}
		// }

		b.ResetTimer()
		sum := 0
		for bi := 0; bi < b.N; bi++ {
			for i := 0; i < TRASH_COUNT; i++ {
				trashItem := newTrashData()
				if trashItem.InnerTrash.Value > 100 {
					sum++
				}
			}
		}
	})

	b.Run("x10 Stack", func(b *testing.B) {
		b.ResetTimer()
		sum := 0
		for bi := 0; bi < b.N; bi++ {
			for i := 0; i < TRASH_COUNT*10; i++ {
				trashItem := TrashData{
					A: rand.Int63n(RAND_MAX),
					B: rand.Int31n(RAND_MAX),
					C: rand.Int63n(RAND_MAX),
					D: rand.Int31n(RAND_MAX),
					InnerTrash: InnerTrash{
						Value:   rand.Int31n(RAND_MAX),
						AsFloat: float32(rand.Int31n(RAND_MAX)),
					},
				}
				if trashItem.InnerTrash.Value > 100 {
					sum++
				}
			}
		}
	})

	b.Run("x10 Heap", func(b *testing.B) {
		b.ResetTimer()
		sum := 0
		for bi := 0; bi < b.N; bi++ {
			for i := 0; i < TRASH_COUNT*10; i++ {
				trashItem := newTrashData()
				if trashItem.InnerTrash.Value > 100 {
					sum++
				}
			}
		}
	})

	b.Run("x100 Stack", func(b *testing.B) {
		b.ResetTimer()
		sum := 0
		for bi := 0; bi < b.N; bi++ {
			for i := 0; i < TRASH_COUNT*100; i++ {
				trashItem := TrashData{
					A: rand.Int63n(RAND_MAX),
					B: rand.Int31n(RAND_MAX),
					C: rand.Int63n(RAND_MAX),
					D: rand.Int31n(RAND_MAX),
					InnerTrash: InnerTrash{
						Value:   rand.Int31n(RAND_MAX),
						AsFloat: float32(rand.Int31n(RAND_MAX)),
					},
				}
				if trashItem.InnerTrash.Value > 100 {
					sum++
				}
			}
		}
	})

	b.Run("x100 Heap", func(b *testing.B) {
		b.ResetTimer()
		sum := 0
		for bi := 0; bi < b.N; bi++ {
			for i := 0; i < TRASH_COUNT*100; i++ {
				trashItem := newTrashData()
				if trashItem.InnerTrash.Value > 100 {
					sum++
				}
			}
		}
	})
}
