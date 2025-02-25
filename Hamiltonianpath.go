package main

import (
	"fmt"
	"os"
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
	return nil
}

func path(visited []*Vertex, destinations []*Vertex, current *Vertex) {
	fmt.Println("Currently at ", current.Num)
	fmt.Println("Removing ", current.Num, " from ", destinations)
	newvisited := append(visited, current)
	newdestination := current.remove(destinations)
	if len(newdestination) == 0 {
		fmt.Println("Success!")
		fmt.Println("A solution is: ")
		for _, ver := range newvisited {
			fmt.Println(ver.Num)
		}
		os.Exit(1)
	}
	for _, i := range current.Neighbours {
		if i.member(newdestination) {
			path(newvisited, newdestination, i)
		}
	}
}

func main() {
	v1 := &Vertex{Num: 1}
	v2 := &Vertex{Num: 2}
	v3 := &Vertex{Num: 3}

	v1.add(v2)
	v1.add(v3)
	v3.add(v2)

	path([]*Vertex{}, []*Vertex{v1, v2, v3}, v1)
}
