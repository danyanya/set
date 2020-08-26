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
