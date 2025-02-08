package user

import (
	context "context"
	"errors"
)

func (*UserFilter) Validate(context.Context) error {
	return errors.New("Validate failed")
}
