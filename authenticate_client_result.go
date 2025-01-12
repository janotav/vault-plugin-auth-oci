// Copyright © 2019, Oracle and/or its affiliates.
package ociauth

import "github.com/oracle/oci-go-sdk/v46/common"

// Do not edit this file. This is based on standard OCI GO SDK format

// Contains the result of the Authenticate Client request.
type AuthenticateClientResult struct {
	Principal    *Principal `json:"principal"`
	ErrorMessage *string    `json:"errorMessage"`
	IsSuccess    *bool      `json:"success"`
}

// Prints the values of pointers in AuthenticateClientResult,
// producing a human friendly string for an struct with pointers. Useful when debugging the values of a struct.
func (m AuthenticateClientResult) String() string {
	return common.PointerString(m)
}
