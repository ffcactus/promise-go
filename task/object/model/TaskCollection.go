package model

import (
	"time"
)

// TaskMember Task member object.
type TaskMember struct {
	URI             string
	Name            string
	Description     *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedByName   string
	CreatedByURI    string
	TargetName      string
	TargetURI       string
	ExecutionState  ExecutionState
	Percentage      int
	CurrentStep     string
	ExecutionResult ExecutionResult
}

// TaskCollection Task collection.
type TaskCollection struct {
	Start       int64
	Count       int64
	Total       int64
	Members     []TaskMember
	NextPageURI string
	PrevPageURI string
}
