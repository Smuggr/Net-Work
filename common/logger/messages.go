package logger

type Resource string

const (
	ResourceDevice  Resource = "DEVICE"
	ResourcePlugin  Resource = "PLUGIN"
	ResourceUser    Resource = "USER"
	ResourceConfig  Resource = "CONFIG"
	ResourceEnv     Resource = "ENV"
)

// Resource messages (User, Device, etc.)
var (
	MsgResourceRegisterSuccess     = NewMessageWrapper("MsgResourceRegisterSuccess", "resource %s (%s) successfully registered", InfoLevel)
	MsgResourceFetchSuccess        = NewMessageWrapper("MsgResourceFetchSuccess", "resource %s (%s) successfully fetched", InfoLevel)
	MsgResourceUpdateSuccess       = NewMessageWrapper("MsgResourceUpdateSuccess", "resource %s (%s) successfully updated", InfoLevel)
	MsgResourceRemoveSuccess       = NewMessageWrapper("MsgResourceRemoveSuccess", "resource %s (%s) successfully removed", InfoLevel)
	MsgResourceAuthenticateSuccess = NewMessageWrapper("MsgResourceAuthenticateSuccess", "resource %s (%s) successfully authenticated", InfoLevel)
)

// Device messages
var (
	MsgDeviceConnectSuccess = NewMessageWrapper("MsgDeviceConnectSuccess", "device %s successfully connected", InfoLevel)
)

// Uncategorizated messages
var (
	MsgResourceLoaded = NewMessageWrapper("MsgResourceLoaded", "resource %s (%s) loaded", InfoLevel)
	MsgInitialized    = NewMessageWrapper("MsgInitialized", "initialized", InfoLevel)
	MsgCleanedUp      = NewMessageWrapper("MsgCleanedUp", "cleaned up", InfoLevel)
)

// Authentication errors
var (
	ErrInvalidCredentials = NewMessageWrapper("ErrInvalidCredentials", "invalid credentials", ErrorLevel)
	ErrUnauthorized       = NewMessageWrapper("ErrUnauthorized", "unauthorized", ErrorLevel)
)

// Token errors
var (
	ErrCreatingToken      = NewMessageWrapper("ErrCreatingToken", "error creating token", ErrorLevel)
	ErrSigningToken       = NewMessageWrapper("ErrSigningToken", "error signing token", ErrorLevel)
	ErrInvalidToken       = NewMessageWrapper("ErrInvalidToken", "invalid token", ErrorLevel)
	ErrInvalidTokenFormat = NewMessageWrapper("ErrInvalidTokenFormat", "invalid token format", ErrorLevel)
)

// User errors
var (
	ErrDefaultAdminNotFound = NewMessageWrapper("ErrDefaultAdminNotFound", "default admin not found", ErrorLevel)
)

// Database errors
var (
	ErrConnectingToDB      = NewMessageWrapper("ErrConnectingToDB", "error connecting to database", ErrorLevel)
	ErrGettingDBConnection = NewMessageWrapper("ErrGettingDBConnection", "error getting database connection", ErrorLevel)
	ErrClosingDBConnection = NewMessageWrapper("ErrClosingDBConnection", "error closing database connection", ErrorLevel)

	ErrResourceAlreadyExists = NewMessageWrapper("ErrResourceAlreadyExists", "resource %s (%s) already exists", ErrorLevel)
	ErrResourceNotFound      = NewMessageWrapper("ErrResourceNotFound", "resource %s (%s) not found", ErrorLevel)

	ErrRegisteringResourceInDB = NewMessageWrapper("ErrRegisteringResourceInDB", "error registering resource %s (%s) in database", ErrorLevel)
	ErrFetchingResourceFromDB  = NewMessageWrapper("ErrFetchingResourceFromDB", "error fetching resource %s (%s) from database", ErrorLevel)
	ErrUpdatingResourceInDB    = NewMessageWrapper("ErrUpdatingResourceInDB", "error updating resource %s (%s) in database", ErrorLevel)
	ErrRemovingResourceFromDB  = NewMessageWrapper("ErrRemovingResourceFromDB", "error removing resource %s (%s) from database", ErrorLevel)
)

// Validation errors
var (
	ErrInvalidIPAddress       = NewMessageWrapper("ErrInvalidIPAddress", "invalid IP address", ErrorLevel)
	ErrInsufficientCharacters = NewMessageWrapper("ErrInsufficientCharacters", "insufficient characters", ErrorLevel)
	ErrLengthNotInRange       = NewMessageWrapper("ErrLengthNotInRange", "length not in range <%d, %d>", ErrorLevel)
	ErrForbiddenCharacter     = NewMessageWrapper("ErrForbiddenCharacter", "forbidden character(s)", ErrorLevel)
)

// Resource errors
var (
	ErrReadingResource    = NewMessageWrapper("ErrReadingResource", "error reading %s (%s)", FatalLevel)
	ErrFormattingResource = NewMessageWrapper("ErrFormattingResource", "error formatting %s (%s)", FatalLevel)
)

// Plugin errors
var (
	ErrLookingUpPluginSymbol       = NewMessageWrapper("ErrLookingUpPluginSymbol", "error looking up plugin symbol from file %s", ErrorLevel)
)

// General errors
var (
	ErrHashingPassword       = NewMessageWrapper("ErrHashingPassword", "error hashing password", ErrorLevel)
	ErrInvalidRequestPayload = NewMessageWrapper("ErrInvalidRequestPayload", "invalid request payload", ErrorLevel)
	ErrOperationNotPermitted = NewMessageWrapper("ErrOperationNotPermitted", "operation not permitted", ErrorLevel)
)

// Uncategorized errors
var (
	ErrInvalidHookConfig = NewMessageWrapper("ErrInvalidHookConfig", "invalid hook config", ErrorLevel)
	ErrInitializing      = NewMessageWrapper("ErrInitializing", "error initializing: %s", ErrorLevel)
	ErrCleaningUp        = NewMessageWrapper("ErrCleaningUp", "error cleaning up: %s", ErrorLevel)
)
