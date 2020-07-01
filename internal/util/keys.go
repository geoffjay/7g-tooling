package util

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ContextKeys holds the context keys through the project
type ContextKeys struct {
	GinContextKey ContextKey
}

var (
	// MasterContextKeys the project's context keys
	ServiceContextKeys = ContextKeys{
		GinContextKey: "gin",
	}
)
