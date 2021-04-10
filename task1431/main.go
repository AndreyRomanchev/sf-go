package main

import (
	"fmt"
	"math"
)

type HashMap struct {
	buckets [][][]string
	size    int
}

func (h *HashMap) Set(key string, value string) string {
	index := h.Hashstr([]byte(key))
	insertedKeyVal := []string{key, value}
	h.buckets[index] = append(h.buckets[index], insertedKeyVal)
	return key
}

func (h *HashMap) Get(key string) (string, bool) {
	index := h.Hashstr([]byte(key))
	if len(h.buckets[index]) == 0 {
		return "", false
	}

	for _, kv := range h.buckets[index] {
		if kv[0] == key {
			return kv[1], true
		}
	}

	return "", false
}

func (h *HashMap) Delete(key string) bool {
	index := h.Hashstr([]byte(key))

	if len(h.buckets[index]) == 0 {
		return false
	}

	for i, kv := range h.buckets[index] {
		if kv[0] == key {
			h.buckets[index] = append(h.buckets[index][:i], h.buckets[index][i+1:]...)
			return true
		}
	}

	return false
}

func (h *HashMap) Hashstr(val []byte) uint64 {
	var total uint64
	for k, v := range val {
		total += uint64(v) * uint64(math.Pow(10, float64(len(val)-k-1)))
	}
	return total % uint64(h.size)
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([][][]string, size),
		size:    size,
	}
}

func main() {
	fmt.Println("vim-go")
	h := NewHashMap(10)
	h.Set("first", "1")
	h.Set("second", "22")
	h.Set("third", "333")
	val, ok := h.Get("first")
	fmt.Println(val, ok)
	fmt.Println(h.Get("second"))

	fmt.Println(h.Get("third"))
	h.Delete("third")
	fmt.Println(h.Get("third"))
}
