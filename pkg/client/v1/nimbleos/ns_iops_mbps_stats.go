// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsIopsMbpsStats - Average and maximum iops and mbps statistics in last 24 hours.
// Export NsIopsMbpsStatsFields for advance operations like search filter etc.
var NsIopsMbpsStatsFields *NsIopsMbpsStatsStringFields

func init() {
	AvgIopsfield := "avg_iops"
	MaxIopsfield := "max_iops"
	AvgMbpsfield := "avg_mbps"
	MaxMbpsfield := "max_mbps"

	NsIopsMbpsStatsFields = &NsIopsMbpsStatsStringFields{
		AvgIops: &AvgIopsfield,
		MaxIops: &MaxIopsfield,
		AvgMbps: &AvgMbpsfield,
		MaxMbps: &MaxMbpsfield,
	}
}

type NsIopsMbpsStats struct {
	// AvgIops - Average combined read and write iops.
	AvgIops *int64 `json:"avg_iops,omitempty"`
	// MaxIops - Maximum combined read and write iops.
	MaxIops *int64 `json:"max_iops,omitempty"`
	// AvgMbps - Average combined read and write throughput.
	AvgMbps *int64 `json:"avg_mbps,omitempty"`
	// MaxMbps - Maximum combined read and write throughput.
	MaxMbps *int64 `json:"max_mbps,omitempty"`
}

// Struct for NsIopsMbpsStatsFields
type NsIopsMbpsStatsStringFields struct {
	AvgIops *string
	MaxIops *string
	AvgMbps *string
	MaxMbps *string
}
