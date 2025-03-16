package main

import "fmt"

type Vertex struct {
	Num        int
	Neighbours []*Vertex
}

func (v *Vertex) member(set []*Vertex) bool {
	for _, i := range set {
		if v == i {
			return true
		}
	}
	return false
}

func (v *Vertex) remove(set []*Vertex) []*Vertex {
	s := make([]*Vertex, len(set))
	copy(s, set)
	for a, b := range s {
		if v == b {
			return append(s[:a], s[a+1:]...)
		}
	}
	return s
}

func path(visited, destinations []*Vertex, current *Vertex) bool {
	newvisited := append(visited, current)
	newdestination := current.remove(destinations)
	if len(newdestination) == 0 {
		fmt.Println("A solution is: ")
		for _, ver := range newvisited {
			fmt.Println(ver.Num)
		}
		return true
	}
	var exists bool
	for _, i := range current.Neighbours {
		if i.member(newdestination) {
			exists = exists || path(newvisited, newdestination, i)
		}
	}
	return exists
}