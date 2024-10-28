package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
)

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}