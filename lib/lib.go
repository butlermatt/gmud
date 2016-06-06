package lib

// Objecter is an interface that all MUD objects implement (but not commands).
type Objecter interface {
	Name() string
	SetName(string)
	Description() string
	SetDescription(string)
}

// Holder is an object which can hold other objects. It has an Inventory
type Holder interface {
	Add(obj Objecter) bool
	Inventory() []Objecter
	Remove(obj Objecter) bool
}
