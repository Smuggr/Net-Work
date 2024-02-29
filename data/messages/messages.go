package messages

type MessageWrapper struct {
	Key     string `json:"key"`
	Message	string `json:"message"`
}

func NewMessageWrapper(key string, message string) *MessageWrapper {
	return &MessageWrapper{
		Key:     key,
		Message: message,
	}
}

func (e *MessageWrapper) Msg() string {
	return e.Message
}

var (
	MsgUserRegisterSuccess   = NewMessageWrapper("MsgRegisterSuccess", "user successfully registered")
	MsgUserUpdateSuccess     = NewMessageWrapper("MsgUserUpdateSuccess", "user successfully updated")
	MsgUsersFetchSuccess     = NewMessageWrapper("MsgUsersFetchSuccess", "users successfully fetched")
	MsgUserFetchSuccess      = NewMessageWrapper("MsgUserFetchSuccess", "user successfully fetched")

	MsgDeviceRegisterSuccess = NewMessageWrapper("MsgDeviceRegisterSuccess", "device successfully registered")
	MsgDeviceUpdateSuccess   = NewMessageWrapper("MsgDeviceUpdateSuccess", "device successfully updated")
	MsgDevicesFetchSuccess   = NewMessageWrapper("MsgDevicesFetchSuccess", "devices successfully fetched")
	MsgDeviceFetchSuccess    = NewMessageWrapper("MsgDeviceFetchSuccess", "device successfully fetched")
)