package lib

type PlayerImpl struct {
	// Name is the username of the player.
	name string
	// Long description of the player.
	long string
	// Room the user is currently located in.
	room *Room
	// inventory the user is holding.
	inventory []Objecter
}

// Holdable returns false as players cannot be picked up.
func (p *PlayerImpl) Holdable() bool { return false }

// Name returns the name of the client. Fulfills the Objecter interface.
func (p *PlayerImpl) Name() string {
	return p.name
}

// SetName sets the player's name.
func (p *PlayerImpl) SetName(name string) {
	p.name = name
}

// Description returns the description of the player.
func (p *PlayerImpl) Description() string {
	return p.long
}

// SetDescription sets the player's description.
func (p *PlayerImpl) SetDescription(desc string) {
	p.long = desc
}

// Room returns a pointer to the current room occupied by the user.
func (p *PlayerImpl) Room() *Room {
	return p.room
}

// SetRoom sets the players current room.
func (p *PlayerImpl) SetRoom(room *Room) {
	p.room = room
}

// Add adds the specified Objecter to the room inventory. Rooms can accept anything
// except rooms.
func (p *PlayerImpl) Add(obj Objecter) bool {
	p.inventory = append(p.inventory, obj)
	return true
}

// Remove removes the specified object from the room. Returns false if the room
// does not have the object. Returns true if it was removed from the room inventory.
func (p *PlayerImpl) Remove(obj Objecter) bool {
	var i int
	for i = 0; i < len(p.inventory); i++ {
		if obj == p.inventory[i] {
			break
		}
	}

	if i == len(p.inventory) {
		return false
	}

	copy(p.inventory[i:], p.inventory[i+1:])
	p.inventory[len(p.inventory)-1] = nil
	p.inventory = p.inventory[:len(p.inventory)-1]
	return true
}

// Inventory returns the list objects the user is holding.
func (p *PlayerImpl) Inventory() []Objecter {
	return p.inventory
}
