package set

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkStoreLoad(b *testing.B) {
	s := String{}
	for i := rand.Intn(10000); i > 0; i-- {
		s.Store(strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		s.Load(strconv.Itoa(i))
	}
}

func TestStoreLoad(T *testing.T) {
	s := String{}
	s.Store("aa")

	if s.Load("aa") {
		T.Logf("aa is in set")
	} else {
		T.Errorf("aa is not in set")
	}
}

func TestStoreKeys(T *testing.T) {
	s := String{}
	s.Store("aa")
	s.Store("bb")

	keys := s.Keys()
	if len(keys) != 2 || keys[0] != "aa" || keys[1] != "bb" {
		T.Errorf("failed with keys: %+v", keys)
	}
}
