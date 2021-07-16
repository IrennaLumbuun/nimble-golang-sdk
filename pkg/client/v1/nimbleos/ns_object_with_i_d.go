// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos


// NsObjectWithID - An object with an ID.
// Export NsObjectWithIDFields for advance operations like search filter etc.
var NsObjectWithIDFields *NsObjectWithID

func init(){
 IDfield:= "id"

 NsObjectWithIDFields= &NsObjectWithID{
  ID:&IDfield,
 }
}


type NsObjectWithID struct {
 // ID - ID of object.
  ID *string `json:"id,omitempty"`
}
