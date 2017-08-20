package plugin


// Result is the response type of plugin execution.
type Result struct {
	Val interface{} 	`json:"res"`
	Err error			`json:"err,omitempty"`
}

// Plugin is the core type for extending go-monitor
type Plugin interface {
	Exec() Result
}


// New creates a plugin from the provided function.
func New(execFn func() (interface{}, error)) Plugin {
	return defaultPlugin{execFn}
}


type defaultPlugin struct {
	execFn func() (interface{}, error)
}

func (p defaultPlugin) Exec() Result {
	val, err := p.execFn()
	return Result{val, err}
}
