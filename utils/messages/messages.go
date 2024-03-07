package messages

import (
	"fmt"
	"regexp"
)

type MessageWrapper struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func removeFormattingCharacters(message string) string {
	re := regexp.MustCompile(`%[a-zA-Z]`)
	return re.ReplaceAllString(message, "")
}

func NewMessageWrapper(key string, message string) *MessageWrapper {
	return &MessageWrapper{
		Key:     key,
		Message: message,
	}
}

func (e *MessageWrapper) String() string {
	return removeFormattingCharacters(e.Message)
}

func (e *MessageWrapper) Msg() string {
	return e.String()
}

func (e *MessageWrapper) FormatMsg(vars ...interface{}) string {
	if vars == nil {
		return e.Msg()
    }

	return fmt.Sprintf(e.Message, vars...)
}

func (e *MessageWrapper) Format(vars ...interface{}) *MessageWrapper {
	copy := *e
	copy.Message = e.FormatMsg(vars...)
	return &copy
}

var (
	// User messages
	MsgUserRegisterSuccess   = NewMessageWrapper("MsgRegisterSuccess", "user %s successfully registered")
	MsgUserUpdateSuccess     = NewMessageWrapper("MsgUserUpdateSuccess", "user %s successfully updated")
	MsgUsersFetchSuccess     = NewMessageWrapper("MsgUsersFetchSuccess", "users (%d) successfully fetched")
	MsgUserFetchSuccess      = NewMessageWrapper("MsgUserFetchSuccess", "user %s successfully fetched")

	// Device messages
	MsgDeviceConnectSuccess  = NewMessageWrapper("MsgDeviceConnectSuccess", "device %s successfully connected")
	MsgDeviceRegisterSuccess = NewMessageWrapper("MsgDeviceRegisterSuccess", "device %s successfully registered")
	MsgDeviceUpdateSuccess   = NewMessageWrapper("MsgDeviceUpdateSuccess", "device %s successfully updated")
	MsgDevicesFetchSuccess   = NewMessageWrapper("MsgDevicesFetchSuccess", "devices (%d) successfully fetched")
	MsgDeviceFetchSuccess    = NewMessageWrapper("MsgDeviceFetchSuccess", "device %s successfully fetched")

	// Removal messages
	MsgDeviceRemoveSuccess   = NewMessageWrapper("MsgDeviceRemoveSuccess", "device %s successfully removed")
	MsgUserRemoveSuccess     = NewMessageWrapper("MsgUserRemoveSuccess", "user %s successfully removed")

	// Uncategorizated messages
	MsgTest                  = NewMessageWrapper("MsgTest", "message test %s, %d, %f, %t")
)