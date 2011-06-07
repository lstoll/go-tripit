// This file was generated by a tool. Do not edit.
package tripit

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for WeatherObject objects
type WeatherObjectVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *WeatherObjectVector) AppendVector(x *WeatherObjectVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *WeatherObjectVector) At(i int) WeatherObject {
	return p.Vector.At(i).(WeatherObject)
}

// Copy makes a copy of the vector and returns it.
func (p *WeatherObjectVector) Copy() WeatherObjectVector {
	return WeatherObjectVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *WeatherObjectVector) Do(f func(elem WeatherObject)) {
	p.Vector.Do(func(e interface{}) { f(e.(WeatherObject)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *WeatherObjectVector) Insert(i int, x WeatherObject) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *WeatherObjectVector) InsertVector(i int, x *WeatherObjectVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *WeatherObjectVector) Last() WeatherObject {
	return p.Vector.Last().(WeatherObject)
}

// Pop deletes the last element of the vector.
func (p *WeatherObjectVector) Pop() WeatherObject {
	return p.Vector.Pop().(WeatherObject)
}

// Push appends x to the end of the vector.
func (p *WeatherObjectVector) Push(x WeatherObject) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *WeatherObjectVector) Resize(length, capacity int) *WeatherObjectVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *WeatherObjectVector) Set(i int, x WeatherObject) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *WeatherObjectVector) Slice(i, j int) *WeatherObjectVector {
	v := p.Vector.Slice(i, j)
	return &WeatherObjectVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *WeatherObjectVector) UnmarshalJSON(b []byte) os.Error {
	var arr []WeatherObject
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]WeatherObject, 1)
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
func (p *WeatherObjectVector) MarshalJSON() ([]byte, os.Error) {
	var a []WeatherObject
	if p == nil {
		a = make([]WeatherObject, 0)
	} else {
		a = make([]WeatherObject, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *WeatherObjectVector) Data() []WeatherObject {
	arr := make([]WeatherObject, p.Len())
	var i int
	i = 0
	p.Do(func(v WeatherObject) {
		arr[i] = v
		i++
	})
	return arr
}
