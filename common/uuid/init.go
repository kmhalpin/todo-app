package uuid

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct {
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (*UUIDGenerator) Generate() (string, error) {
	uid, err := uuid.NewRandom()
	return uid.String(), err
}
