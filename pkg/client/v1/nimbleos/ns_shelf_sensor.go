// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfSensor - A shelf sensor data.
// Export NsShelfSensorFields for advance operations like search filter etc.
var NsShelfSensorFields *NsShelfSensorStringFields

func init() {
	Typefield := "type"
	Namefield := "name"
	DisplayNamefield := "display_name"
	Locationfield := "location"
	Cidfield := "cid"
	Valuefield := "value"
	Statusfield := "status"

	NsShelfSensorFields = &NsShelfSensorStringFields{
		Type:        &Typefield,
		Name:        &Namefield,
		DisplayName: &DisplayNamefield,
		Location:    &Locationfield,
		Cid:         &Cidfield,
		Value:       &Valuefield,
		Status:      &Statusfield,
	}
}

type NsShelfSensor struct {
	// Type - Type of the sensor.
	Type *NsShelfSensorType `json:"type,omitempty"`
	// Name - Internal name of the sensor.
	Name *string `json:"name,omitempty"`
	// DisplayName - Name for display purpose.
	DisplayName *string `json:"display_name,omitempty"`
	// Location - Location of the sensor.
	Location *string `json:"location,omitempty"`
	// Cid - Which controller this sensor applies to.
	Cid *NsShelfCtrlrSide `json:"cid,omitempty"`
	// Value - Value of the sensor reading.
	Value *int64 `json:"value,omitempty"`
	// Status - Sensor status.
	Status *NsShelfSensorState `json:"status,omitempty"`
}

// Struct for NsShelfSensorFields
type NsShelfSensorStringFields struct {
	Type        *string
	Name        *string
	DisplayName *string
	Location    *string
	Cid         *string
	Value       *string
	Status      *string
}
