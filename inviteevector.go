package tripit

// This file was generated by a tool. Do not edit.

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for Invitee objects
type InviteeVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *InviteeVector) AppendVector(x *InviteeVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *InviteeVector) At(i int) Invitee {
	return p.Vector.At(i).(Invitee)
}

// Copy makes a copy of the vector and returns it.
func (p *InviteeVector) Copy() InviteeVector {
	return InviteeVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *InviteeVector) Do(f func(elem Invitee)) {
	p.Vector.Do(func(e interface{}) { f(e.(Invitee)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *InviteeVector) Insert(i int, x Invitee) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *InviteeVector) InsertVector(i int, x *InviteeVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *InviteeVector) Last() Invitee {
	return p.Vector.Last().(Invitee)
}

// Pop deletes the last element of the vector.
func (p *InviteeVector) Pop() Invitee {
	return p.Vector.Pop().(Invitee)
}

// Push appends x to the end of the vector.
func (p *InviteeVector) Push(x Invitee) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *InviteeVector) Resize(length, capacity int) *InviteeVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *InviteeVector) Set(i int, x Invitee) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *InviteeVector) Slice(i, j int) *InviteeVector {
	v := p.Vector.Slice(i, j)
	return &InviteeVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *InviteeVector) UnmarshalJSON(b []byte) os.Error {
	var arr []Invitee
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]Invitee, 1)
		err := json.Unmarshal(b, &arr[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				arr = arr[0:0]
			} else {
				return err
			}
		}
	}
	p.Cut(0, p.Len())
	for _, v := range arr {
		p.Push(v)
	}
	return nil
}

// MarshalJSON customizes the JSON output for Vectors.
func (p *InviteeVector) MarshalJSON() ([]byte, os.Error) {
	var a []Invitee
	if p == nil {
		a = make([]Invitee, 0)
	} else {
		a = make([]Invitee, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *InviteeVector) Data() []Invitee {
	arr := make([]Invitee, p.Len())
	var i int
	i = 0
	p.Do(func(v Invitee) {
		arr[i] = v
		i++
	})
	return arr
}