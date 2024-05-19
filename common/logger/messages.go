package logger

var (
	// User messages
	MsgUserRegisterSuccess             = NewMessageWrapper("MsgRegisterSuccess", "user %s successfully registered", InfoLevel)
	MsgUserUpdateSuccess               = NewMessageWrapper("MsgUserUpdateSuccess", "user %s successfully updated", InfoLevel)
	MsgUsersFetchSuccess               = NewMessageWrapper("MsgUsersFetchSuccess", "users (%d) successfully fetched", InfoLevel)
	MsgUserFetchSuccess                = NewMessageWrapper("MsgUserFetchSuccess", "user %s successfully fetched", InfoLevel)
	MsgUserAuthenticateSuccess         = NewMessageWrapper("MsgUserAuthenticateSuccess", "user %s successfully authenticated", InfoLevel)

	// Device messages
	MsgDeviceConnectSuccess            = NewMessageWrapper("MsgDeviceConnectSuccess", "device %s successfully connected", InfoLevel)
	MsgDeviceRegisterSuccess           = NewMessageWrapper("MsgDeviceRegisterSuccess", "device %s successfully registered", InfoLevel)
	MsgDeviceAuthenticateSuccess       = NewMessageWrapper("MsgDeviceAuthenticateSuccess", "device %s successfully authenticated", InfoLevel)
	MsgDeviceUpdateSuccess             = NewMessageWrapper("MsgDeviceUpdateSuccess", "device %s successfully updated", InfoLevel)
	MsgDevicesFetchSuccess             = NewMessageWrapper("MsgDevicesFetchSuccess", "devices (%d) successfully fetched", InfoLevel)
	MsgDeviceFetchSuccess              = NewMessageWrapper("MsgDeviceFetchSuccess", "device %s successfully fetched", InfoLevel)
	MsgPluginProvidersInfoFetchSuccess = NewMessageWrapper("MsgPluginProvidersInfoFetchSuccess", "plugin providers (%d) info successfully fetched", InfoLevel)
	MsgPluginProviderInfoFetchSuccess  = NewMessageWrapper("MsgPluginProviderInfoFetchSuccess", "plugin provider %s info successfully fetched", InfoLevel)

	// Removal messages
	MsgDeviceRemoveSuccess             = NewMessageWrapper("MsgDeviceRemoveSuccess", "device %s successfully removed", InfoLevel)
	MsgUserRemoveSuccess               = NewMessageWrapper("MsgUserRemoveSuccess", "user %s successfully removed", InfoLevel)

	// Uncategorizated messages
	MsgTest                            = NewMessageWrapper("MsgTest", "message test %s, %d, %f, %t", InfoLevel)
	MsgEnvFileLoaded				   = NewMessageWrapper("MsgEnvFileLoaded", "env file loaded", InfoLevel)
	MsgConfigFileLoaded				   = NewMessageWrapper("MsgConfigFileLoaded", "config file loaded %s", InfoLevel)
	MsgInitialized					   = NewMessageWrapper("MsgInitialized", "initialized", InfoLevel)
	MsgCleanedUp					   = NewMessageWrapper("MsgCleanedUp", "cleaned up", InfoLevel)
)

var (
	// Authentication errors
	ErrInvalidCredentials           = NewMessageWrapper("ErrInvalidCredentials", "invalid credentials", ErrorLevel)
	ErrUnauthorized                 = NewMessageWrapper("ErrUnauthorized", "unauthorized", ErrorLevel)

	// Token errors
	ErrCreatingToken                = NewMessageWrapper("ErrCreatingToken", "error creating token", ErrorLevel)
	ErrSigningToken                 = NewMessageWrapper("ErrSigningToken", "error signing token", ErrorLevel)
	ErrInvalidToken                 = NewMessageWrapper("ErrInvalidToken", "invalid token", ErrorLevel)
	ErrInvalidTokenFormat           = NewMessageWrapper("ErrInvalidTokenFormat", "invalid token format", ErrorLevel)

	// User errors
	ErrDefaultAdminNotFound         = NewMessageWrapper("ErrDefaultAdminNotFound", "default admin not found", ErrorLevel)
	ErrUserAlreadyExists            = NewMessageWrapper("ErrUserAlreadyExists", "user %s already exists", ErrorLevel)
	ErrUserNotFound                 = NewMessageWrapper("ErrUserNotFound", "user %s not found", ErrorLevel)
	ErrFetchingUsersFromDB          = NewMessageWrapper("ErrFetchingUsersFromDB", "error fetching users from database", ErrorLevel)

	// Device errors
	ErrDeviceAlreadyExists          = NewMessageWrapper("ErrDeviceAlreadyExists", "device %s already exists", ErrorLevel)
	ErrRegisteringDeviceInDB        = NewMessageWrapper("ErrRegisteringDeviceInDB", "error registering device %s in database", ErrorLevel)
	ErrUpdatingDeviceInDB           = NewMessageWrapper("ErrUpdatingDeviceInDB", "error updating device %s in database", ErrorLevel)
	ErrRemovingDeviceFromDB         = NewMessageWrapper("ErrRemovingDeviceFromDB", "error removing device %s from database", ErrorLevel)
	ErrFetchingDevicesFromDB        = NewMessageWrapper("ErrFetchingDevicesFromDB", "error fetching devices from database", ErrorLevel)
	ErrDeviceNotFound               = NewMessageWrapper("ErrDeviceNotFound", "device %s not found", ErrorLevel)
	ErrInvalidIPAddress             = NewMessageWrapper("ErrInvalidIPAddress", "invalid ip address", ErrorLevel)
	ErrPluginProviderAlreadyLoaded  = NewMessageWrapper("ErrPluginProviderAlreadyLoaded", "error plugin provider %s already loaded", ErrorLevel)
	ErrPluginProviderConflict       = NewMessageWrapper("ErrPluginProviderConflict", "plugin provider conflict, %s already exists", ErrorLevel)
	ErrAPIVersionMismatch           = NewMessageWrapper("ErrAPIVersionMismatch", "api version mismatch, plugin api version %s, expected %s", ErrorLevel)
	ErrRemovingDevicePlugin         = NewMessageWrapper("ErrRemovingDevicePlugin", "error removing device %s plugin %s", ErrorLevel)
	ErrCreatingDevicePlugin         = NewMessageWrapper("ErrCreatingDevicePlugin", "error creating device %s plugin %s", ErrorLevel)
	ErrClientNotFound               = NewMessageWrapper("ErrClientNotFound", "client %s not found", ErrorLevel)
	ErrDevicePluginNotFound         = NewMessageWrapper("ErrDevicePluginNotFound", "device %s plugin not found", ErrorLevel)

	// Database errors
	ErrConnectingToDB               = NewMessageWrapper("ErrConnectingToDB", "error connecting to database", ErrorLevel)
	ErrGettingDBConnection          = NewMessageWrapper("ErrGettingDBConnection", "error getting database connection", ErrorLevel)
	ErrClosingDBConnection          = NewMessageWrapper("ErrClosingDBConnection", "error closing database connection", ErrorLevel)
	ErrRegisteringUserInDB          = NewMessageWrapper("ErrRegisteringUserInDB", "error registering user %s in database", ErrorLevel)
	ErrUpdatingUserInDB             = NewMessageWrapper("ErrUpdatingUserInDB", "error updating user %s in database", ErrorLevel)
	ErrRemovingUserFromDB           = NewMessageWrapper("ErrRemovingUserFromDB", "error removing user %s from database", ErrorLevel)

	// Validation errors
	ErrInsufficientCharacters       = NewMessageWrapper("ErrInsufficientCharacters", "insufficient characters", ErrorLevel)
	ErrLengthNotInRange             = NewMessageWrapper("ErrLengthNotInRange", "length not in range <%d, %d>", ErrorLevel)
	ErrForbiddenCharacter           = NewMessageWrapper("ErrForbiddenCharacter", "forbidden character(s)", ErrorLevel)

	// File errors
	ErrReadingConfigFile            = NewMessageWrapper("ErrReadingConfigFile", "error reading config file", FatalLevel)
	ErrFormattingConfigFile         = NewMessageWrapper("ErrFormattingConfigFile", "error formatting config file", FatalLevel)
	ErrReadingEnvFile               = NewMessageWrapper("ErrReadingEnvFile", "error reading env file", FatalLevel)
	ErrLookingUpPluginSymbol        = NewMessageWrapper("ErrLookingUpPluginSymbol", "error looking up plugin symbol from file %s", ErrorLevel)
	ErrFetchingPluginProvidersInfo  = NewMessageWrapper("ErrFetchingPluginProvidersInfo", "error fetching plugin providers info", ErrorLevel)
	ErrPluginProviderNotFound       = NewMessageWrapper("ErrPluginProviderNotFound", "plugin provider %s not found", ErrorLevel)
	
	// General errors
	ErrHashingPassword              = NewMessageWrapper("ErrHashingPassword", "error hashing password", ErrorLevel)
	ErrInvalidRequestPayload        = NewMessageWrapper("ErrInvalidRequestPayload", "invalid request payload", ErrorLevel)
	ErrOperationNotPermitted        = NewMessageWrapper("ErrOperationNotPermitted", "operation not permitted", ErrorLevel)

	// Uncategorized errors
	ErrInvalidHookConfig            = NewMessageWrapper("ErrInvalidHookConfig", "invalid hook config", ErrorLevel)
	ErrTest                         = NewMessageWrapper("ErrTest", "error test %s, %d, %f, %t", ErrorLevel)
	ErrInitializing 			    = NewMessageWrapper("ErrInitializing", "error initializing: %s", ErrorLevel)
	ErrCleaningUp				    = NewMessageWrapper("ErrCleaningUp", "error cleaning up: %s", ErrorLevel)
)
