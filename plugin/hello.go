package plugin


// NewHelloWorld return the plugin which says Hello World!
func NewHelloWorld() Plugin {
	return New(func() (interface{}, error) {
		return "Hello World!", nil
	})
}
