// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSensorData - A list of samples (in order of sample time) for a sensor.
// Export NsSensorDataFields for advance operations like search filter etc.
var NsSensorDataFields *NsSensorData

func init() {

	NsSensorDataFields = &NsSensorData{
		Sensor: "sensor",
	}
}

type NsSensorData struct {
	// Sensor - Sensor name.
	Sensor string `json:"sensor,omitempty"`
	// Samples - A list of samples for the sensor.
	Samples []*float64 `json:"samples,omitempty"`
}

// sdk internal struct
type nsSensorData struct {
	Sensor  *string    `json:"sensor,omitempty"`
	Samples []*float64 `json:"samples,omitempty"`
}

// EncodeNsSensorData - Transform NsSensorData to nsSensorData type
func EncodeNsSensorData(request interface{}) (*nsSensorData, error) {
	reqNsSensorData := request.(*NsSensorData)
	byte, err := json.Marshal(reqNsSensorData)
	if err != nil {
		return nil, err
	}
	respNsSensorDataPtr := &nsSensorData{}
	err = json.Unmarshal(byte, respNsSensorDataPtr)
	return respNsSensorDataPtr, err
}

// DecodeNsSensorData Transform nsSensorData to NsSensorData type
func DecodeNsSensorData(request interface{}) (*NsSensorData, error) {
	reqNsSensorData := request.(*nsSensorData)
	byte, err := json.Marshal(reqNsSensorData)
	if err != nil {
		return nil, err
	}
	respNsSensorData := &NsSensorData{}
	err = json.Unmarshal(byte, respNsSensorData)
	return respNsSensorData, err
}
