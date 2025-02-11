package validator

import (
	"go-clean-arch-game-server/pkg/utils/structure"
	"sync"
)

var lock = &sync.Mutex{}

var validatorInstance structure.Validator

func GetValidator() structure.Validator {
	if validatorInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if validatorInstance == nil {
			validatorInstance = structure.NewValidator()
		}
	}

	return validatorInstance
}
