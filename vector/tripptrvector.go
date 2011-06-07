// This file was generated by a tool. Do not edit.
package tripit

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for *Trip objects
type TripPtrVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *TripPtrVector) AppendVector(x *TripPtrVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *TripPtrVector) At(i int) *Trip {
	return p.Vector.At(i).(*Trip)
}

// Copy makes a copy of the vector and returns it.
func (p *TripPtrVector) Copy() TripPtrVector {
	return TripPtrVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *TripPtrVector) Do(f func(elem *Trip)) {
	p.Vector.Do(func(e interface{}) { f(e.(*Trip)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *TripPtrVector) Insert(i int, x *Trip) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *TripPtrVector) InsertVector(i int, x *TripPtrVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *TripPtrVector) Last() *Trip {
	return p.Vector.Last().(*Trip)
}

// Pop deletes the last element of the vector.
func (p *TripPtrVector) Pop() *Trip {
	return p.Vector.Pop().(*Trip)
}

// Push appends x to the end of the vector.
func (p *TripPtrVector) Push(x *Trip) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *TripPtrVector) Resize(length, capacity int) *TripPtrVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *TripPtrVector) Set(i int, x *Trip) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *TripPtrVector) Slice(i, j int) *TripPtrVector {
	v := p.Vector.Slice(i, j)
	return &TripPtrVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *TripPtrVector) UnmarshalJSON(b []byte) os.Error {
	var arr []*Trip
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]*Trip, 1)
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
func (p *TripPtrVector) MarshalJSON() ([]byte, os.Error) {
	var a []*Trip
	if p == nil {
		a = make([]*Trip, 0)
	} else {
		a = make([]*Trip, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *TripPtrVector) Data() []*Trip {
	arr := make([]*Trip, p.Len())
	var i int
	i = 0
	p.Do(func(v *Trip) {
		arr[i] = v
		i++
	})
	return arr
}
