package plugin


// Result is the response type of plugin execution.
type Result struct {
	Val interface{} 	`json:"res"`
	Err error			`json:"err,omitempty"`
}

// Plugin is the core type for extending go-monitor
type Plugin interface {
	Name() string
	Exec() Result
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

func (p defaultPlugin) Exec() Result {
	val, err := p.execFn()
	return Result{val, err}
}
