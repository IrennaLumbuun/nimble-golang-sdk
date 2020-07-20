// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Support - View and alter support-based parameters.
// Export SupportFields for advance operations like search filter etc.
var SupportFields *Support

func init() {

	SupportFields = &Support{
		ID: "id",
	}
}

type Support struct {
	// ID - Object identifier for the group.
	ID string `json:"id,omitempty"`
	// PasswordMode - Mode for support password.
	PasswordMode *NsSupportPasswordMode `json:"password_mode,omitempty"`
	// ArrayCount - Count of arrays for support password.
	ArrayCount int64 `json:"array_count,omitempty"`
	// ArrayList - Details of support passwords for arrays.
	ArrayList []*NsSupportPasswordArray `json:"array_list,omitempty"`
}

// sdk internal struct
type support struct {
	ID           *string                   `json:"id,omitempty"`
	PasswordMode *NsSupportPasswordMode    `json:"password_mode,omitempty"`
	ArrayCount   *int64                    `json:"array_count,omitempty"`
	ArrayList    []*NsSupportPasswordArray `json:"array_list,omitempty"`
}

// EncodeSupport - Transform Support to support type
func EncodeSupport(request interface{}) (*support, error) {
	reqSupport := request.(*Support)
	byte, err := json.Marshal(reqSupport)
	if err != nil {
		return nil, err
	}
	respSupportPtr := &support{}
	err = json.Unmarshal(byte, respSupportPtr)
	return respSupportPtr, err
}

// DecodeSupport Transform support to Support type
func DecodeSupport(request interface{}) (*Support, error) {
	reqSupport := request.(*support)
	byte, err := json.Marshal(reqSupport)
	if err != nil {
		return nil, err
	}
	respSupport := &Support{}
	err = json.Unmarshal(byte, respSupport)
	return respSupport, err
}
