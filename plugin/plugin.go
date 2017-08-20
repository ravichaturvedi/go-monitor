package plugin


// Plugin is the core type for extending go-monitor
type Plugin interface {
	Name() string
	Exec() (interface{}, error)
}


// New creates a new plugin providing the plugin name and the function to be executed.
func New(name string, execFn func() (interface{}, error)) Plugin {
	return defaultPlugin{name, execFn}
}


type defaultPlugin struct {
	name string
	execFn func() (interface{}, error)
}

func (p defaultPlugin) Name() string {
	return p.name
}

func (p defaultPlugin) Exec() (interface{}, error) {
	return p.execFn()
}
