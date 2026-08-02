package singleton

import (
	"sync"
	"testing"
)

const parallelCount = 100

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Fail()
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parallelCount)
	instances := [parallelCount]*singleton{}
	for i := 0; i < parallelCount; i++ {
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
		}(i)
	}
	wg.Wait()

	for one := 0; one < parallelCount-1; one++ {
		for two := one + 1; two < parallelCount; two++ {
			if instances[one] != instances[two] {
				t.Fail()
			}
		}
	}
}
