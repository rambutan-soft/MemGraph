package main

import (
	"errors"
	"reflect"
	"sync"
)

//Node ..
type Node struct {
	TypeID int
	Key    string
	Level  int
	Value  interface{}
	Heads  map[*Node]int //int----is id from edges
	Tails  map[*Node]int
	*sync.RWMutex
}

//MemGraph ...
type MemGraph struct {
	Nodes    map[int]map[int]*Node //typeid,id
	Index    map[int]map[string]int
	Types    map[int]interface{}
	Edges    map[int]interface{}
	MaxLevel int
	*sync.RWMutex
}

//NewMemGraph ...Initializes a new graph.
func NewMemGraph() *MemGraph {
	return &MemGraph{make(map[int]map[int]*Node), make(map[int]map[string]int), make(map[int]interface{}), make(map[int]interface{}), 0, new(sync.RWMutex)}
}

//AddType ..
func (g *MemGraph) AddType(typeID int, data interface{}) {
	g.Types[typeID] = data //update later
}

//Add ..
func (g *MemGraph) Add(typeID int, key string, data interface{}) {
	g.Lock()
	defer g.Unlock()

	_, ok := g.Index[typeID]
	if !ok {
		g.Index[typeID] = make(map[string]int)
	}
	id := len(g.Index[typeID]) + 1
	g.Index[typeID][key] = id

	_, ok = g.Nodes[typeID]
	if !ok {
		g.Nodes[typeID] = make(map[int]*Node)
	}
	g.Nodes[typeID][id] = &Node{typeID, key, 0, data, make(map[*Node]int), make(map[*Node]int), new(sync.RWMutex)}
}

//Connect ..
func (g *MemGraph) Connect(parentTypeID, childTypeID int, parentKey, childKey string, edge interface{}) bool {
	if parentKey == childKey && parentTypeID == childTypeID {
		return false
	}
	g.RLock()
	defer g.RUnlock()

	pid, ok1 := g.Index[parentTypeID][parentKey]
	cid, ok2 := g.Index[childTypeID][childKey]

	if !ok1 || !ok2 {
		return false
	}

	pNode, ok1 := g.Nodes[parentTypeID][pid]
	cNode, ok2 := g.Nodes[childTypeID][cid]

	if !ok1 || !ok2 {
		return false
	}

	edgeID := 0
	for k, v := range g.Edges {
		if reflect.TypeOf(v) == reflect.TypeOf(edge) {
			edgeID = k
			break
		}
	}
	if edgeID == 0 {
		edgeID = len(g.Edges) + 1
	}

	pNode.Lock()
	cNode.Lock()
	defer cNode.Unlock()
	defer pNode.Unlock()
	pNode.Tails[cNode] = edgeID
	cNode.Heads[pNode] = edgeID

	return true
}

//Get ... get the Account
func (g *MemGraph) Get(typeID int, key string) (*Node, error) {
	g.RLock()
	defer g.RUnlock()
	id, ok := g.Index[typeID][key]
	if ok {
		v, ok := g.Nodes[typeID][id]
		if ok {
			return v, nil
		}
	}
	return nil, errors.New("invalid key")
}
