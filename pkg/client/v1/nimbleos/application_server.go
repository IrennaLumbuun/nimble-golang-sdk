// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ApplicationServer - An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.
// Export ApplicationServerFields for advance operations like search filter etc.
var ApplicationServerFields *ApplicationServerStringFields

func init() {
	IDfield := "id"
	Namefield := "name"
	Hostnamefield := "hostname"
	Portfield := "port"
	Usernamefield := "username"
	Descriptionfield := "description"
	Passwordfield := "password"
	ServerTypefield := "server_type"
	Metadatafield := "metadata"
	CreationTimefield := "creation_time"
	LastModifiedfield := "last_modified"

	ApplicationServerFields = &ApplicationServerStringFields{
		ID:           &IDfield,
		Name:         &Namefield,
		Hostname:     &Hostnamefield,
		Port:         &Portfield,
		Username:     &Usernamefield,
		Description:  &Descriptionfield,
		Password:     &Passwordfield,
		ServerType:   &ServerTypefield,
		Metadata:     &Metadatafield,
		CreationTime: &CreationTimefield,
		LastModified: &LastModifiedfield,
	}
}

type ApplicationServer struct {
	// ID - Identifier for the application server.
	ID *string `json:"id,omitempty"`
	// Name - Name for the application server.
	Name *string `json:"name,omitempty"`
	// Hostname - Application server hostname.
	Hostname *string `json:"hostname,omitempty"`
	// Port - Application server port number.
	Port *int64 `json:"port,omitempty"`
	// Username - Application server username.
	Username *string `json:"username,omitempty"`
	// Description - Text description of application server.
	Description *string `json:"description,omitempty"`
	// Password - Application server password.
	Password *string `json:"password,omitempty"`
	// ServerType - Application server type ({invalid|vss|vmware|cisco|stack_vision|container_node}).
	ServerType *NsAppServerType `json:"server_type,omitempty"`
	// Metadata - Key-value pairs that augment an application server's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// CreationTime - Time when this application server was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this application server was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
}

// Struct for ApplicationServerFields
type ApplicationServerStringFields struct {
	ID           *string
	Name         *string
	Hostname     *string
	Port         *string
	Username     *string
	Description  *string
	Password     *string
	ServerType   *string
	Metadata     *string
	CreationTime *string
	LastModified *string
}
