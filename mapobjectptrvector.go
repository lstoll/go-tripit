package tripit

// This file was generated by a tool. Do not edit.

import (
	"encoding/json"
)

// MapObjectPtrVector is a specialization of Vector for *MapObject objects.
type MapObjectPtrVector []*MapObject

// UnmarshalJSON builds the vector from the JSON in b.
func (p *MapObjectPtrVector) UnmarshalJSON(b []byte) error {
	var arr *[]*MapObject
	arr = (*[]*MapObject)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]*MapObject, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}
		
		if (*arr)[0] == nil {
			*arr = (*arr)[0:0]
		}
		
	}
	return nil
}
