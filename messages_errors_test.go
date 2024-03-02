package main

import (
	"testing"

	//"network/data/messages"
	"network/data/errors"
)

func TestErrors(t *testing.T) {
	t.Log("Formatted error wrapper:", errors.ErrTest.Format("string", 10, 1.0, true))
	t.Log("Formatted error message:", errors.ErrTest.FormatError("string", 10, 1.0, true))

	formattedErrMsg := errors.ErrTest.FormatError()
	t.Log("Blank error message using format without arguments:", formattedErrMsg)

	if formattedErrMsg != errors.ErrTest.Error() {
		t.FailNow()
	}

	t.Log("Blank error message:", formattedErrMsg)
}
