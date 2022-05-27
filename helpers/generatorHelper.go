package helpers

import (
	"github.com/segmentio/ksuid"
)

func GenerateId() (string, error) {
	id, err := ksuid.NewRandom()

	if err != nil {
		return "", err
	}

	return id.String(), nil
}
