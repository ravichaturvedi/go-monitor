package plugin

import "encoding/json"

// Result is the response type of plugin execution.
type Result struct {
	Val interface{}
	Err error
}

func (r Result) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return json.Marshal(map[string]interface{}{"msg": r.Val})
	}

	return json.Marshal(map[string]interface{}{"msg": r.Val, "err": r.Err.Error()})
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
