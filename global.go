package container

var (
	// Global is the global container
	Global Container
)

// Register service by id
func Register(id string, service any) {
	Global.Register(id, service)
}

// Get gets a service by id
func Get(id string) (any, bool) {
	return Global.Get(id)
}

// MustGet calls Get underneath
// will panic if serviceect not found within container
func MustGet(id string) any {
	return Global.MustGet(id)
}

// Invoke gets a service safely typed by passing it to a closure
// will panic if callback is not a function
func Invoke(id string, fn any) {
	Global.Invoke(id, fn)
}

// MustInvoke calls MustGet underneath
// will panic if service not found within container
func MustInvoke(id string, fn any) {
	Global.MustInvoke(id, fn)
}

// Has checks if a service exists within the container
func Has(id string) bool {
	return Global.Has(id)
}
