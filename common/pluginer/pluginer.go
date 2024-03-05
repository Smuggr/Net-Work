package pluginer

type Plugin interface {
	Initialize()
	Execute()
	Cleanup()
}