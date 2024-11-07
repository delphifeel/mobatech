package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Trash struct {
	a1  string
	a2  float64
	a3  float64
	a4  float64
	aa1 int64
}

type HttpConnection struct {
	WantToLiveForMS int64
	AliveForMS      int64
	Content         string
	Trash           Trash
}

type TodoConnection struct {
	WantToLiveForMS int64
	Content         string
	Id              string
}

const CONNECTIONS_QUEUE = 10_000_000
const CONN_PER_TICK = 1
const MIN_MS = 300

type ConnectionsPool struct {
	array        []HttpConnection
	freeList     []*HttpConnection
	freeListSize int
}

const POOL_SIZE = CONN_PER_TICK * MIN_MS * 100

func NewConnectionsPool() ConnectionsPool {
	p := ConnectionsPool{
		array:    make([]HttpConnection, POOL_SIZE),
		freeList: make([]*HttpConnection, POOL_SIZE),
	}
	p.freeListSize = POOL_SIZE
	for i := 0; i < p.freeListSize; i++ {
		p.freeList[i] = &p.array[i]
	}
	return p
}

func (pool *ConnectionsPool) Free(v *HttpConnection) {
	// fmt.Println("Free")
	pool.freeList[pool.freeListSize] = v
	pool.freeListSize++
}

func (pool *ConnectionsPool) Allocate() *HttpConnection {
	// fmt.Println("allocate")
	if pool.freeListSize == 0 {
		fmt.Println("WARNING: no free. doubling")
		panic("")
	}

	existing := pool.freeList[pool.freeListSize-1]
	pool.freeListSize--

	return existing
}

// func (pool *ConnectionsPool) PutBack(v *HttpConnection) {
// 	freeListSize := len(pool.freeIndexes)
// 	if freeListSize == 0 {
// 		panic("OOM")
// 	}

// 	freeIndex := pool.freeIndexes[freeListSize-1]
// 	pool.freeIndexes = pool.freeIndexes[:freeListSize]

// 	existing := pool.array[freeIndex]
// 	existing.AliveForMS = v.AliveForMS
// 	existing.Content = v.Content
// 	existing.WantToLiveForMS = v.WantToLiveForMS

// 	pool.aliveIndexes = append(pool.aliveIndexes, freeIndex)
// }

func Benchmark_Pooling(b *testing.B) {

	connectionsTodoSrc := make([]TodoConnection, CONNECTIONS_QUEUE)
	for i := 0; i < CONNECTIONS_QUEUE; i++ {
		connectionsTodoSrc[i].WantToLiveForMS = MIN_MS + rand.Int63n(100)
		connectionsTodoSrc[i].Id = fmt.Sprintf("%v", i+666)
		connectionsTodoSrc[i].Content = fmt.Sprintf("%v", rand.Int63())
	}

	allocTestResults := make([]int, CONNECTIONS_QUEUE)
	poolTestResults := make([]int, CONNECTIONS_QUEUE)

	b.Run("allocations", func(b *testing.B) {
		aliveConnections := map[string]*HttpConnection{}
		aliveConnections["DUMMY"] = &HttpConnection{
			WantToLiveForMS: 1000,
			AliveForMS:      0,
			Content:         "DUMMY CONTENT",
		}
		timePassedMS := time.Now().UnixMilli()
		connectionsTodoPos := 0

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			timeDiffMS := time.Now().UnixMilli() - timePassedMS
			// if timeDiffMS < 13 {
			// 	continue
			// }
			timePassedMS = time.Now().UnixMilli()

			for nci := 0; nci < CONN_PER_TICK; nci++ {
				if connectionsTodoPos == len(connectionsTodoSrc) {
					panic("connections exhaust")
				}
				todoConnection := connectionsTodoSrc[connectionsTodoPos]
				connectionsTodoPos++

				_, exist := aliveConnections[todoConnection.Id]
				if exist {
					panic("exist")
				}

				aliveConnections[todoConnection.Id] = &HttpConnection{
					WantToLiveForMS: todoConnection.WantToLiveForMS,
					AliveForMS:      0,
					Content:         todoConnection.Content,
				}
			}

			idsToRemove := make([]string, 0)
			// fmt.Println(len(aliveConnections))
			allocTestResults = append(allocTestResults, len(aliveConnections))
			for id, iterConn := range aliveConnections {
				if iterConn.AliveForMS >= iterConn.WantToLiveForMS {
					idsToRemove = append(idsToRemove, id)
					continue
				}

				iterConn.AliveForMS += timeDiffMS
			}

			for _, idToRemove := range idsToRemove {
				delete(aliveConnections, idToRemove)
			}

			// fmt.Printf("%#v\n\n", aliveConnections)
		}
	})

	b.Run("pool", func(b *testing.B) {
		aliveConnections := make([]*HttpConnection, POOL_SIZE)
		aliveConnectionsSize := 0
		aliveConnectionsSwapBuffer := make([]*HttpConnection, POOL_SIZE)
		aliveConnectionsSwapBufferSize := 0

		connectionsPool := NewConnectionsPool()
		conn := connectionsPool.Allocate()

		conn.WantToLiveForMS = 1000
		conn.AliveForMS = 0
		conn.Content = "DUMMY CONTENT"
		aliveConnections[aliveConnectionsSize] = conn
		aliveConnectionsSize += 1

		timePassedMS := time.Now().UnixMilli()
		connectionsTodoPos := 0

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// fmt.Printf("START START\n\n")
			timeDiffMS := time.Now().UnixMilli() - timePassedMS
			// if timeDiffMS < 13 {
			// 	continue
			// }
			timePassedMS = time.Now().UnixMilli()

			// fmt.Println(aliveConnectionsCount)

			for nci := 0; nci < CONN_PER_TICK; nci++ {
				if connectionsTodoPos == len(connectionsTodoSrc) {
					panic("connections exhaust")
				}
				todoConnection := connectionsTodoSrc[connectionsTodoPos]
				connectionsTodoPos++

				conn := connectionsPool.Allocate()
				conn.AliveForMS = 0
				conn.WantToLiveForMS = todoConnection.WantToLiveForMS
				conn.Content = todoConnection.Content
				aliveConnections[aliveConnectionsSize] = conn
				aliveConnectionsSize += 1
			}

			// fmt.Println(aliveConnectionsSize)
			poolTestResults = append(poolTestResults, aliveConnectionsSize)
			for bi := 0; bi < aliveConnectionsSize; bi++ {
				iterConn := aliveConnections[bi]
				// fmt.Printf("%#v\n", iterConn)
				if iterConn.AliveForMS >= iterConn.WantToLiveForMS {
					connectionsPool.Free(iterConn)
				} else {
					iterConn.AliveForMS += timeDiffMS
					aliveConnectionsSwapBuffer[aliveConnectionsSwapBufferSize] = iterConn
					aliveConnectionsSwapBufferSize += 1
				}
			}

			aliveConnectionsSize = aliveConnectionsSwapBufferSize
			tmp := aliveConnectionsSwapBuffer
			aliveConnectionsSwapBuffer = aliveConnections
			aliveConnections = tmp
			aliveConnectionsSwapBufferSize = 0

			// fmt.Printf("%#v\n\n", aliveConnections)
		}
	})

}
