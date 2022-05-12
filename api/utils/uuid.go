package utils

import (
	uuid "github.com/satori/go.uuid"
)

func NewUuid() string {
	newUuid := uuid.NewV4()
	uuidString := newUuid.String()
	return uuidString
}
