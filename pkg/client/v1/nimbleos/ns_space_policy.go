// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsSpacePolicy Enum.

type NsSpacePolicy string

const (
 cNsSpacePolicyOffline NsSpacePolicy = "offline"
 cNsSpacePolicyLoginOnly NsSpacePolicy = "login_only"
 cNsSpacePolicyNonWritable NsSpacePolicy = "non_writable"
 cNsSpacePolicyReadOnly NsSpacePolicy = "read_only"
 cNsSpacePolicyInvalid NsSpacePolicy = "invalid"
)

var pNsSpacePolicyOffline NsSpacePolicy
var pNsSpacePolicyLoginOnly NsSpacePolicy
var pNsSpacePolicyNonWritable NsSpacePolicy
var pNsSpacePolicyReadOnly NsSpacePolicy
var pNsSpacePolicyInvalid NsSpacePolicy

// NsSpacePolicyOffline enum exports
var NsSpacePolicyOffline *NsSpacePolicy

// NsSpacePolicyLoginOnly enum exports
var NsSpacePolicyLoginOnly *NsSpacePolicy

// NsSpacePolicyNonWritable enum exports
var NsSpacePolicyNonWritable *NsSpacePolicy

// NsSpacePolicyReadOnly enum exports
var NsSpacePolicyReadOnly *NsSpacePolicy

// NsSpacePolicyInvalid enum exports
var NsSpacePolicyInvalid *NsSpacePolicy

func init() {
 pNsSpacePolicyOffline = cNsSpacePolicyOffline
 NsSpacePolicyOffline = &pNsSpacePolicyOffline

 pNsSpacePolicyLoginOnly = cNsSpacePolicyLoginOnly
 NsSpacePolicyLoginOnly = &pNsSpacePolicyLoginOnly

 pNsSpacePolicyNonWritable = cNsSpacePolicyNonWritable
 NsSpacePolicyNonWritable = &pNsSpacePolicyNonWritable

 pNsSpacePolicyReadOnly = cNsSpacePolicyReadOnly
 NsSpacePolicyReadOnly = &pNsSpacePolicyReadOnly

 pNsSpacePolicyInvalid = cNsSpacePolicyInvalid
 NsSpacePolicyInvalid = &pNsSpacePolicyInvalid

}

