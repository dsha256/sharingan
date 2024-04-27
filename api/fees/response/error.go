package response

import (
	"errors"

	"encore.dev/beta/errs"
	"go.temporal.io/api/serviceerror"
)

var (
	ErrBillNotFound = errors.New("bill not found")
)

func Error(err error) error {
	var notFound *serviceerror.NotFound
	if errors.As(err, &notFound) {
		return errs.WrapCode(errors.Join(ErrBillNotFound, err), errs.NotFound, "")
	}

	return err
}
