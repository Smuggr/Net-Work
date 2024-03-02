package main

import (
	"log"
	"testing"
	"reflect"

	"network/data/messages"
	"network/data/errors"
)

func TestErrors(t *testing.T) {
	log.Println("\n\nTesting errors:")

	log.Println("Formatted wrapper:", "\"", errors.ErrTest.Format("string", 10, 1.0, true), "\"", "=>", reflect.TypeOf(errors.ErrTest.Format("string", 10, 1.0, true)))
	log.Println("Formatted msg:", "\"", errors.ErrTest.FormatError("string", 10, 1.0, true), "\"", "=>", reflect.TypeOf(errors.ErrTest.FormatError("string", 10, 1.0, true)))
	log.Println("Error:", "\"", errors.ErrTest.Err, "\"", "=>", reflect.TypeOf(errors.ErrTest.Error()))

	formattedErrMsg := errors.ErrTest.FormatError()
	log.Println("Blank msg format w/o args:", "\"", formattedErrMsg, "\"", "=>", reflect.TypeOf(formattedErrMsg))

	if formattedErrMsg != errors.ErrTest.Error() {
		t.FailNow()
	}

	log.Println("Blank msg:", "\"", formattedErrMsg, "\"", "=>", reflect.TypeOf(formattedErrMsg))
}

func TestMessages(t *testing.T) {
	log.Println("\n\nTesting messages:")

	log.Println("Formatted wrapper:", "\"", messages.MsgTest.Format("string", 10, 1.0, true), "\"", "=>", reflect.TypeOf(messages.MsgTest.Format("string", 10, 1.0, true)))
	log.Println("Formatted:", "\"", messages.MsgTest.FormatMsg("string", 10, 1.0, true), "\"", "=>", reflect.TypeOf(messages.MsgTest.FormatMsg("string", 10, 1.0, true)))

	formattedMsg := messages.MsgTest.FormatMsg()
	log.Println("Blank msg format w/o args:", "\"", formattedMsg, "\"", "=>", reflect.TypeOf(formattedMsg))

	if formattedMsg != messages.MsgTest.Msg() {
		t.FailNow()
	}

	log.Println("Blank msg:", "\"", formattedMsg, "\"", "=>", reflect.TypeOf(formattedMsg))
}