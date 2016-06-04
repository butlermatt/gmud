package lib

// Objecter is an interface that all MUD objects implement (but not commands).
type Objecter interface {
	Name() string
	Description() string
}

// Holder is an object which can hold other objects. It has an Inventory
type Holder interface {
	Objecter
	Add(obj Objecter) bool
	Inventory() []Objecter
	Remove(obj Objecter) bool
}
