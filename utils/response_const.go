package utils

const (
	Success               = 200
	ErrInternalServer int = iota + 1000
	ErrDataNotFound
	ErrInvalidParams
	ErrNoAuth
	ErrInvalidSign
)
