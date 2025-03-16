package main

import (
	"fmt"
)

type Vertex struct {
	Num        int
	Neighbours []*Vertex
}

func (v *Vertex) add(n *Vertex) {
	v.Neighbours = append(v.Neighbours, n)
}

func (v *Vertex) member(set []*Vertex) bool {
	for _, i := range set {
		if v.Num == i.Num {
			return true
		}
	}
	return false
}

func (v *Vertex) remove(set []*Vertex) []*Vertex {
	s := make([]*Vertex, len(set))
	copy(s, set)
	for a, b := range s {
		if v.Num == b.Num {
			s = append(s[:a], s[a+1:]...)
			return s
		}
	}
	return s
}

func path(visited []*Vertex, destinations []*Vertex, current *Vertex) bool {
	newvisited := append(visited, current)
	newdestination := current.remove(destinations)
	if len(newdestination) == 0 {
		fmt.Println("A solution is: ")
		for _, ver := range newvisited {
			fmt.Println(ver.Num)
		}
		return true
	}
	exists := false
	for _, i := range current.Neighbours {
		if i.member(newdestination) {
			if path(newvisited, newdestination, i) {
				exists = true
			}
		}
	}
	return exists
}

func main() {
	v1 := &Vertex{Num: 1}
	v2 := &Vertex{Num: 2}
	v3 := &Vertex{Num: 3}
	v4 := &Vertex{Num: 4}

	v1.add(v2)
	v2.add(v3)
	v3.add(v4)
	v2.add(v4)
	v4.add(v3)

	if !path([]*Vertex{}, []*Vertex{v1, v2, v3, v4}, v1) {
		fmt.Println("No such path found.")
	}
}
