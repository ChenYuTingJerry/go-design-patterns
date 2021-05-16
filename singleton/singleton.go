package singleton

import "fmt"

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	fmt.Println("instacne: ", instance)
	if instance == nil {
		instance = &singleton{}
	}

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

