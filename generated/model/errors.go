package model

// This is a generated file
// Manual changes will be overwritten

// InvalidDataError No description provided
type InvalidDataError string

func (e InvalidDataError) Error() string {
	return string(e)
}

// NotFoundError No description provided
type NotFoundError string

func (e NotFoundError) Error() string {
	return string(e)
}
