package go_LRU_algorithm

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)           //2作废
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           //1作废
	fmt.Println(lru.Get(1)) //-1
	fmt.Println(lru.Get(3)) //3
	fmt.Println(lru.Get(4)) //4
}
