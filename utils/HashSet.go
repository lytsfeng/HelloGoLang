package utils

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}] bool
}

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}


func NewHashSet() *HashSet {
	return &HashSet{m:make(map[interface{}] bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e]{
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{})  {
	delete(set.m,e)
}

func (set *HashSet) Clear()  {
	set.m = make(map[interface{}] bool)
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Len() int{
	return len(set.m)
}

func (set *HashSet) Same(other Set) bool  {
	if other == nil{
		return false
	}
	if set.Len() != other.Len(){
		return false
	}

	for  key := range set.m{
		if !other.Contains(key){
			return false
		}
	}
	return true
}

func (set *HashSet) Elements() ([] interface{}){
	initialLen := len(set.m)
	snapshot := make([]interface{},initialLen)
	actualLen := 0

	for key := range set.m{
		if actualLen < initialLen{
			snapshot[actualLen] = key
		}else {
			snapshot = append(snapshot,key)
		}
		actualLen++
	}
	if actualLen < initialLen{
		snapshot = snapshot[: actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string  {
	var buf bytes.Buffer
	buf.WriteString("set{")
	first := true
	for key := range set.m {
		if first{
			first = false
		}else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v",key))
	}
	buf.WriteString("}")
	return buf.String()
}
//超集
func IsSuperset(one Set,other Set) bool {

	if one == nil || other == nil{
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()

	if oneLen == 0 || oneLen == otherLen{
		return false
	}
	if oneLen > 0 && otherLen == 0{
		return true
	}

	for _,v := range other.Elements(){
		if !one.Contains(v){
			return false
		}
	}
	return true
}



















