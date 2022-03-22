package main

import (
	"fmt"
	"math"
	"sync"
	"github.com/OneOfOne/xxhash"
)

const (
	expandFactor = 0.75
)


type HashMap struct {
	array []*keyPairs
	capacity int
	len int
	capacityMask int
	lock sync.Mutex
}

type keyPairs struct {
	key string
	value interface{}
	next *keyPairs
}

func NewHashMap(capacity int) *HashMap {
	defaultCapacity := 1 << 4
	if capacity <= defaultCapacity {
		capacity = defaultCapacity
	} else {
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

func (m *HashMap) Len() int {
	return m.len
}


var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

func (m *HashMap) hashIndex(key string, mask int) int {
	hash := hashAlgorithm([]byte(key))
	index := hash & uint64(mask)
	return int(index)
}


func (m *HashMap) Put (key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)
	element := m.array[index]

	if element == nil {
		m.array[index] = &keyPairs{
			key: key,
			value: value,
		}
	} else {
		var lastPairs *keyPairs

		for element != nil {

			if element.key == key {
				element.value = value
				return
			}
			lastPairs = element
			element = element.next
		}

		lastPairs.next = &keyPairs{
			key:key,
			value:value,
		}

		newLen := m.len + 1

		if float64(newLen)/float64(m.capacity) >= expandFactor {
			newM := new(HashMap)
			newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
			newM.capacity = 2*m.capacity
			newM.capacityMask = 2*m.capacity - 1

			for _, pairs := range m.array {
				for pairs != nil {
					newM.Put(pairs.key, pairs.value)
					pairs = pairs.next
				}
			}

			m.array = newM.array
			m.capacity = newM.capacity
			m.capacityMask = newM.capacityMask
		}
		m.len = newLen
	}
}

func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Lock()

	index := m.hashIndex(key, m.capacityMask)
	element := m.array[index]

	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return
}


func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Lock()

	index := m.hashIndex(key, m.capacity)
	element := m.array[index]

	if element == nil {
		return
	}

	if element.key == key {
		m.array[index] = element.next
		m.len = m.len - 1
		return
	}

	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			m.len = m.len - 1
		}
		element = nextElement
		nextElement = nextElement.next
	}
}

func (m *HashMap) Range() {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("'%v'= '%v,'",pairs.key, pairs.value)
			pairs = pairs.next
		}
	}

	fmt.Println()
}
