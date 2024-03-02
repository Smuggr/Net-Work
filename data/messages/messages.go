package messages

import "fmt"

type MessageWrapper struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func NewMessageWrapper(key string, message string) *MessageWrapper {
	return &MessageWrapper{
		Key:     key,
		Message: message,
	}
}

func (e *MessageWrapper) String() string {
	return e.Message
}

func (e *MessageWrapper) FormatMsg(vars ...interface{}) string {
	if vars == nil {
		return e.Message
    }

	return fmt.Sprintf(e.Message, vars...)
}

func (e *MessageWrapper) Format(vars ...interface{}) *MessageWrapper {
	e.Message = fmt.Sprintf(e.Message, vars...)
	return e
}

var (
	// User messages
	MsgUserRegisterSuccess   = NewMessageWrapper("MsgRegisterSuccess", "user %s successfully registered")
	MsgUserUpdateSuccess     = NewMessageWrapper("MsgUserUpdateSuccess", "user %s successfully updated")
	MsgUsersFetchSuccess     = NewMessageWrapper("MsgUsersFetchSuccess", "users successfully fetched")
	MsgUserFetchSuccess      = NewMessageWrapper("MsgUserFetchSuccess", "user %s successfully fetched")

	// Device messages
	MsgDeviceConnectSuccess  = NewMessageWrapper("MsgDeviceConnectSuccess", "device %s successfully connected")
	MsgDeviceRegisterSuccess = NewMessageWrapper("MsgDeviceRegisterSuccess", "device %s successfully registered")
	MsgDeviceUpdateSuccess   = NewMessageWrapper("MsgDeviceUpdateSuccess", "device %s successfully updated")
	MsgDevicesFetchSuccess   = NewMessageWrapper("MsgDevicesFetchSuccess", "devices successfully fetched")
	MsgDeviceFetchSuccess    = NewMessageWrapper("MsgDeviceFetchSuccess", "device %s successfully fetched")

	// Removal messages
	MsgDeviceRemoveSuccess   = NewMessageWrapper("MsgDeviceRemoveSuccess", "device %s successfully removed")
	MsgUserRemoveSuccess     = NewMessageWrapper("MsgUserRemoveSuccess", "user %s successfully removed")
)