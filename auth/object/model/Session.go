package model

// Session The Session object
type Session struct {
	ID        string
	AccountID string
	Token     string
	Expire    int64
}
