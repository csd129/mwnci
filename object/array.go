package object

import (
	"bytes"
	"sort"
	"strings"
)

// Array wraps Object array and implements Object interface.
type Array struct {
	// Elements holds the individual members of the array we're wrapping.
	Elements []Object

	// offset holds our iteration-offset.
	offset int
}

// Type returns the type of this object.
func (ao *Array) Type() Type {
	return ARRAY_OBJ
}

func (ao *Array) String() string {
	return ao.Inspect()
}

// Inspect returns a string-representation of the given object.
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (ao *Array) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "methods" {
		static := []string{"len", "methods"}
		dynamic := env.Names("array.")

		var names []string
		names = append(names, static...)
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)

		result := make([]Object, len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

// Reset implements the Iterable interface, and allows the contents
// of the array to be reset to allow re-iteration.
func (ao *Array) Reset() {
	ao.offset = 0
}

// Next implements the Iterable interface, and allows the contents
// of our array to be iterated over.
func (ao *Array) Next() (Object, Object, bool) {
	if ao.offset < len(ao.Elements) {
		ao.offset++

		element := ao.Elements[ao.offset-1]
		return element, &Integer{Value: int64(ao.offset - 1)}, true
	}

	return nil, &Integer{Value: 0}, false
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (ao *Array) ToInterface() interface{} {
	return "<ARRAY>"
}

func (a *Array) Copy() *Array {
	elements := make([]Object, len(a.Elements))
	for i, e := range a.Elements {
		elements[i] = e
	}
	return &Array{Elements: elements}
}

func (a *Array) PopLeft() Object {
	if len(a.Elements) > 0 {
		e := a.Elements[0]
		a.Elements = a.Elements[1:]
		return e
	}
	return &Null{}

}

func (a *Array) Len() int {
	return len(a.Elements)
}

func (a *Array) Less(i, j int) bool {
	if cmp, ok := a.Elements[i].(Comparable); ok {
		return cmp.Compare(a.Elements[j]) == -1
	}
	return false
}

func (a *Array) Swap(i, j int) {
	a.Elements[i], a.Elements[j] = a.Elements[j], a.Elements[i]
}

func (a *Array) Insert(i int, j Object) {
	a.Elements = append(a.Elements, j)
	copy(a.Elements[i+1:], a.Elements[i:])
	a.Elements[i] = j
}

func (a *Array) Aset(i int, j Object) {
	a.Elements[i] = j
}

func (a *Array) Bool() bool {
	return len(a.Elements) > 0
}

func (a *Array) PopRight() Object {

	if len(a.Elements) > 0 {
		e := a.Elements[(len(a.Elements) - 1)]
		a.Elements = a.Elements[:(len(a.Elements) - 1)]
		return e
	}
	return &Null{}
}

func (a *Array) Reverse() {
	for i, j := 0, len(a.Elements)-1; i < j; i, j = i+1, j-1 {
		a.Elements[i], a.Elements[j] = a.Elements[j], a.Elements[i]
	}
}

func (ao *Array) Compare(other Object) int {
	if obj, ok := other.(*Array); ok {
		if len(ao.Elements) != len(obj.Elements) {
			return -1
		}
		for i, el := range ao.Elements {
			cmp, ok := el.(Comparable)
			if !ok {
				return -1
			}
			if cmp.Compare(obj.Elements[i]) != 0 {
				return cmp.Compare(obj.Elements[i])
			}
		}

		return 0
	}
	return -1
}

func (a *Array) Prepend(obj Object) {
	a.Elements = append([]Object{obj}, a.Elements...)
}

func (a *Array) Append(obj Object) {
	a.Elements = append(a.Elements, obj)
}

func (a *Array) SameType(obj Object) bool {
	First := a.Elements[0].Type()
	for _, el := range a.Elements {
		ThisType := el.Type()
		if First != ThisType {
			return false
		}
	}
	return true
}
