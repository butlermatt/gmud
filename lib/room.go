package lib

import "fmt"

// Room represents an environment for the player.
type Room struct {
	short     string
	long      string
	inventory []Objecter
}

// Name returns the short name of the room.
func (r *Room) Name() string {
	return r.short
}

// SetName sets the short name of the room.
func (r *Room) SetName(name string) {
	r.short = name
}

// Description returns the short and long description of the room.
func (r *Room) Description() string {
	return fmt.Sprintf("%s\r\n%s", r.short, r.long)
}

// SetDescription sets the long description of the room.
func (r *Room) SetDescription(desc string) {
	r.long = desc
}

// Add adds the specified Objecter to the room inventory. Rooms can accept anything
// except rooms.
func (r *Room) Add(obj Objecter) bool {
	r.inventory = append(r.inventory, obj)
	return true
}

// Remove removes the specified object from the room. Returns false if the room
// does not have the object. Returns true if it was removed from the room inventory.
func (r *Room) Remove(obj Objecter) bool {
	var i int
	for i = 0; i < len(r.inventory); i++ {
		if obj == r.inventory[i] {
			break
		}
	}

	if i == len(r.inventory) {
		return false
	}

	copy(r.inventory[i:], r.inventory[i+1:])
	r.inventory[len(r.inventory)-1] = nil
	r.inventory = r.inventory[:len(r.inventory)-1]
	return true
}

// Inventory returns a slice of the Room's inventory.
func (r *Room) Inventory() []Objecter {
	return r.inventory
}

// DefaultRoom is the standard room when all else fails.
var DefaultRoom = &Room{
	short: "Empty Void",
	long:  "You are floating weightlessly in an empty void",
}
