package lib

// Objecter is an interface that all MUD objects implement (but not commands).
type Objecter interface {
	Name() string
	SetName(string)
	Description() string
	SetDescription(string)
	Holdable() bool
}

// Holder is an object which can hold other objects. It has an Inventory
type Holder interface {
	Add(obj Objecter) bool
	Inventory() []Objecter
	Remove(obj Objecter) bool
}

type Player interface {
	Objecter
	Holder
	// Write sends a message to the Player in a non-blocking way.
	Write(string)
	// Send sends a message to the Player in a potentially blocking way.
	Send(string)
	// Room returns the room of the current player.
	Room() *Room
	// SetRoom sets the player's room.
	SetRoom(*Room)
	// Sends a quit
	Quit()
}
