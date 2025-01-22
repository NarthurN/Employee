package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int `json:"age"`
	Salary int `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
		counter: 1,
	}
}