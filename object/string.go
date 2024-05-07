// The implementation of our string-object.

package object

import (
	"hash/fnv"
	"sort"
	"strings"
	"unicode/utf8"
)

// String wraps string and implements Object and Hashable interfaces.
type String struct {
	// Value holds the string value this object wraps.
	Value string

	// Offset holds our iteration-offset
	offset int
}

// Type returns the type of this object.
func (s *String) Type() Type {
	return STRING_OBJ
}

// Inspect returns a string-representation of the given object.
func (s *String) Inspect() string {
	return s.Value
}

func (s *String) String() string {
	return s.Value
}

// HashKey returns a hash key for the given object.
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// InvokeMethodinvokes a method against the object.
// (Built-in methods only.)
func (s *String) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "methods" {
		static := []string{"methods"}
		dynamic := env.Names("string.")

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
// of the string to be reset to allow re-iteration.
func (s *String) Reset() {
	s.offset = 0
}

// Next implements the Iterable interface, and allows the contents
// of our string to be iterated over.
func (s *String) Next() (Object, Object, bool) {

	if s.offset < utf8.RuneCountInString(s.Value) {
		s.offset++

		// Get the characters as an array of runes
		chars := []rune(s.Value)

		// Now index
		val := &String{Value: string(chars[s.offset-1])}

		return val, &Integer{Value: int64(s.offset - 1)}, true
	}

	return nil, &Integer{Value: 0}, false
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (s *String) ToInterface() interface{} {
	return s.Value
}

func (s *String) Len() int {
	return utf8.RuneCountInString(s.Value)
}

func (s *String) Compare(other Object) int {
        if obj, ok := other.(*String); ok {
                switch {
                case s.Value < obj.Value:
                        return -1
                case s.Value > obj.Value:
                        return 1
                default:
                        return 0
                }
        }
        return 1
}
