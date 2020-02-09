package jsonapi

// The EntityNamer interface can be optionally implemented to directly return the
// name of resource used for the "type" field.
//
// Note: By default the name is guessed from the struct name.
type EntityNamer interface {
	GetName() string
}

// The EntityUrler interface can be optionally implemented to directly return the
// url of resource used for the "type" field.
//
// Note: By default the url is taken from the name
type EntityUrler interface {
	GetUrl() string
}
