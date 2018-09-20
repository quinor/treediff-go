package diff

import (
	"gopkg.in/bblfsh/sdk.v2/uast/nodes"
)

type Changelist []Change

type Change interface {
	change()
}

type changeBase struct {
	txID uint64
}

func (_ *changeBase) change() {}

// TODO: proper ID of a node somehow
type ID int64

// key in a node, string for nodes.Object and int for nodes.Array
type Key interface{ key() }

type StringKey string
type IntKey int

func (_ IntKey) key()    {}
func (_ StringKey) key() {}

// four change types

// create a node
type Create struct {
	changeBase
	node nodes.Node
}

// delete a node by ID
type Delete struct {
	changeBase
	nodeID ID
}

// attach a node as a child of another node with a given key
type Attach struct {
	changeBase
	parent ID
	key    Key
	child  ID
}

// deattach a child from a node
type Deattach struct {
	changeBase
	parent ID
	key    Key //or string, how to do alternative?
}
