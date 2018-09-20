package diff

import (
	"gopkg.in/bblfsh/sdk.v2/uast/nodes"
	"reflect"
)

// pointerOf returns a Go pointer for Node that is a reference type (Arrays and Objects).
func pointerOf(n nodes.Node) uintptr {
	if n == nil {
		return 0
	}
	v := reflect.ValueOf(n)
	if v.IsNil() {
		return 0
	}
	return v.Pointer()
}

type arrayPtr uintptr
type mapPtr uintptr

type NodeID interface{}

// UniqueKey returns a unique key of the node in the current tree. The key can be used in maps.
func UniqueKey(n nodes.Node) NodeID {
	switch n := n.(type) {
	case nil:
		return nil
	case nodes.Value:
		return n
	default:
		ptr := pointerOf(n)
		// distinguish nil arrays and maps
		switch n.(type) {
		case nodes.Object:
			return mapPtr(ptr)
		case nodes.Array:
			return arrayPtr(ptr)
		}
		return ptr
	}
}
