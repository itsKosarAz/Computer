  
package main

import (
	"fmt"
)

type student struct {
	ID   int
	name string
	GPA  float64
}

type stack struct {
	array   [2]*student
	counter int
}

func main() {
	var s1 stack
	s1.input()
	s1.print()
}

func (s *stack) input() {
	for i := 0; i < len(s.array); i++ {
		fmt.Println("Enter student ID")
		var tmpID int
		fmt.Scanln(&tmpID)
		fmt.Println("Enter student name")
		var tmpName string
		fmt.Scanln(&tmpName)
		fmt.Println("Enter student GPA")
		var tmpGPA float64
		fmt.Scanln(&tmpGPA)
		s.push(tmpID, tmpName, tmpGPA)
	}
}

func (s *stack) push(tmpID int, tmpName string, tmpGPA float64) {
	tmpStudent := &student{tmpID, tmpName, tmpGPA}
	s.array[s.counter] = tmpStudent
	s.counter++
}

func (s *stack) pop() *student {
	s.counter--
	tmp := s.array[s.counter]
	s.array[s.counter] = nil
	return tmp
}

func (s *stack) print() {
	for i := 0; i < s.counter; i++ {
		fmt.Println(s.array[i].ID, s.array[i].name, s.array[i].GPA)
	}
	//fmt.Println(s.array)
}

func (s *stack) popAll() {
	tmp := s.counter
	for i := 0; i < tmp; i++ {
		s.pop()
	}
}
