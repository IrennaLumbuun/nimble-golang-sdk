// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArraySoftwareUpdateError - Software update error for specific controller.
// Export NsArraySoftwareUpdateErrorFields for advance operations like search filter etc.
var NsArraySoftwareUpdateErrorFields *NsArraySoftwareUpdateError

func init() {

	NsArraySoftwareUpdateErrorFields = &NsArraySoftwareUpdateError{
		Error: "error",
	}
}

type NsArraySoftwareUpdateError struct {
	// Error - Error code from software update.
	Error string `json:"error,omitempty"`
}

// sdk internal struct
type nsArraySoftwareUpdateError struct {
	Error *string `json:"error,omitempty"`
}

// EncodeNsArraySoftwareUpdateError - Transform NsArraySoftwareUpdateError to nsArraySoftwareUpdateError type
func EncodeNsArraySoftwareUpdateError(request interface{}) (*nsArraySoftwareUpdateError, error) {
	reqNsArraySoftwareUpdateError := request.(*NsArraySoftwareUpdateError)
	byte, err := json.Marshal(reqNsArraySoftwareUpdateError)
	if err != nil {
		return nil, err
	}
	respNsArraySoftwareUpdateErrorPtr := &nsArraySoftwareUpdateError{}
	err = json.Unmarshal(byte, respNsArraySoftwareUpdateErrorPtr)
	return respNsArraySoftwareUpdateErrorPtr, err
}

// DecodeNsArraySoftwareUpdateError Transform nsArraySoftwareUpdateError to NsArraySoftwareUpdateError type
func DecodeNsArraySoftwareUpdateError(request interface{}) (*NsArraySoftwareUpdateError, error) {
	reqNsArraySoftwareUpdateError := request.(*nsArraySoftwareUpdateError)
	byte, err := json.Marshal(reqNsArraySoftwareUpdateError)
	if err != nil {
		return nil, err
	}
	respNsArraySoftwareUpdateError := &NsArraySoftwareUpdateError{}
	err = json.Unmarshal(byte, respNsArraySoftwareUpdateError)
	return respNsArraySoftwareUpdateError, err
}
