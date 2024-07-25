// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"time"
)

// Defines values for SourceStatus.
const (
	SourceStatusError                     SourceStatus = "error"
	SourceStatusGatheringInitialInventory SourceStatus = "gathering-initial-inventory"
	SourceStatusUpToDate                  SourceStatus = "up-to-date"
	SourceStatusWaitingForCredentials     SourceStatus = "waiting-for-credentials"
)

// Error defines model for Error.
type Error struct {
	// Message Error message
	Message string `json:"message"`
}

// Source defines model for Source.
type Source struct {
	CreatedAt  time.Time    `json:"createdAt"`
	Id         string       `json:"id"`
	Inventory  string       `json:"inventory"`
	Name       string       `json:"name"`
	Status     SourceStatus `json:"status"`
	StatusInfo string       `json:"status_info"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}

// SourceStatus defines model for Source.Status.
type SourceStatus string

// SourceCreate defines model for SourceCreate.
type SourceCreate struct {
	Name string `json:"name"`
}

// SourceList defines model for SourceList.
type SourceList = []Source

// Status Status is a return value for calls that don't return other objects.
type Status struct {
	// Message A human-readable description of the status of this operation.
	Message *string `json:"message,omitempty"`

	// Reason A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	Reason *string `json:"reason,omitempty"`

	// Status Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status *string `json:"status,omitempty"`
}

// CreateSourceJSONRequestBody defines body for CreateSource for application/json ContentType.
type CreateSourceJSONRequestBody = SourceCreate
