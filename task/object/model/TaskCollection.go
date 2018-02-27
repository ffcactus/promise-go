package model

import (
	"time"
)

// TaskMember Task member object.
type TaskMember struct {
	URI             string
	Name            string
	Description     string
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
	Start       int
	Count       int
	Total       int
	Members     []TaskMember
	NextPageURI string
	PrevPageURI string
}
