package main

import (
	"gotour/ch12/bpool"
	"io"
	"os"
	"sync"
)

func main() {
	var bp = bpool.NewBytePoolCap(500, 1024, 1024)
	var sp = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024, 1024)
		},
	}

	opBytePool(bp)
	opSyncPool(sp)
}

func opBytePool(bp *bpool.BytePoolCap) {
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func(bp *bpool.BytePoolCap) {
			buffer := bp.Get()
			defer bp.Put(buffer)
			mockReadFile(buffer)
			wg.Done()
		}(bp)
	}
	wg.Wait()
}

func opSyncPool(sp *sync.Pool) {
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func(sp *sync.Pool) {
			buffer := sp.Get().([]byte)
			defer sp.Put(buffer)
			mockReadFile(buffer)
			wg.Done()
		}(sp)
	}
	wg.Wait()
}

func mockReadFile(b []byte) {
	f, _ := os.Open("water")
	for {
		n, err := io.ReadFull(f, b)
		if n == 0 || err == io.EOF {
			break
		}
	}
}
