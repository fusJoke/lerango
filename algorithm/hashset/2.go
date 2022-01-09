package main

import "container/list"

type myHashSet struct {
	data []list.List
}

func New() *myHashSet {
	return &myHashSet{data: make([]list.List, base)}
}

func (s *myHashSet) Hash(key int) int {
	return key % base
}

func (s *myHashSet) Add(key int) {
	if !s.Contains(key) {
		h := s.Hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *myHashSet) Remove(key int) {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e.Next() {
		if e.Value.(int) == key {
			s.data[h].Remove(e)
		}
	}
}

func (s *myHashSet) Contains(key int ) bool {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e.Next(){
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}