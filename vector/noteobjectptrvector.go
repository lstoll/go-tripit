// This file was generated by a tool. Do not edit.
package tripit

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for *NoteObject objects
type NoteObjectPtrVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *NoteObjectPtrVector) AppendVector(x *NoteObjectPtrVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *NoteObjectPtrVector) At(i int) *NoteObject {
	return p.Vector.At(i).(*NoteObject)
}

// Copy makes a copy of the vector and returns it.
func (p *NoteObjectPtrVector) Copy() NoteObjectPtrVector {
	return NoteObjectPtrVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *NoteObjectPtrVector) Do(f func(elem *NoteObject)) {
	p.Vector.Do(func(e interface{}) { f(e.(*NoteObject)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *NoteObjectPtrVector) Insert(i int, x *NoteObject) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *NoteObjectPtrVector) InsertVector(i int, x *NoteObjectPtrVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *NoteObjectPtrVector) Last() *NoteObject {
	return p.Vector.Last().(*NoteObject)
}

// Pop deletes the last element of the vector.
func (p *NoteObjectPtrVector) Pop() *NoteObject {
	return p.Vector.Pop().(*NoteObject)
}

// Push appends x to the end of the vector.
func (p *NoteObjectPtrVector) Push(x *NoteObject) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *NoteObjectPtrVector) Resize(length, capacity int) *NoteObjectPtrVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *NoteObjectPtrVector) Set(i int, x *NoteObject) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *NoteObjectPtrVector) Slice(i, j int) *NoteObjectPtrVector {
	v := p.Vector.Slice(i, j)
	return &NoteObjectPtrVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *NoteObjectPtrVector) UnmarshalJSON(b []byte) os.Error {
	var arr []*NoteObject
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]*NoteObject, 1)
		err := json.Unmarshal(b, &arr[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				arr = arr[0:0]
			} else {
				return err
			}
		}
		if arr[0] == nil {
			arr = arr[0:0]
		}
	}
	p.Cut(0, p.Len())
	for _, v := range arr {
		p.Push(v)
	}
	return nil
}

// MarshalJSON customizes the JSON output for Vectors.
func (p *NoteObjectPtrVector) MarshalJSON() ([]byte, os.Error) {
	var a []*NoteObject
	if p == nil {
		a = make([]*NoteObject, 0)
	} else {
		a = make([]*NoteObject, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *NoteObjectPtrVector) Data() []*NoteObject {
	arr := make([]*NoteObject, p.Len())
	var i int
	i = 0
	p.Do(func(v *NoteObject) {
		arr[i] = v
		i++
	})
	return arr
}
