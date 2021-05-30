package main

import "fmt"

type student struct {
	ID   int
	name string
	GPA  float64
}

type stack struct {
	slice []*student
}

var n int

func main() {
	var s1 stack
	s1.input()
	s1.popAll()
	s1.print()
}

func (s *stack) input() {
	fmt.Println("How many students?")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
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
	tmpStudent := student{tmpID, tmpName, tmpGPA}
	s.slice = append(s.slice, &tmpStudent)
}

func (s *stack) pop() *student {
	lastIndex := len(s.slice) - 1
	tmp := s.slice[lastIndex]
	s.slice = s.slice[:lastIndex]
	return tmp
}

func (s *stack) print() {
	for i := 0; i < len(s.slice); i++ {
		fmt.Println(s.slice[i].ID, s.slice[i].name, s.slice[i].GPA)
	}
	//fmt.Println(s.slice)
}

func (s *stack) popAll() {
	for i := 0; i < n; i++ {
		s.pop()
	}
}
