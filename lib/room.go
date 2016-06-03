package lib

type Room struct {
	short     string
	long      string
	inventory []Objecter
}

func (r *Room) Name() string {
	return r.short
}

func (r *Room) Add(obj Objecter) bool {
	r.inventory = append(r.inventory, obj)
	return true
}

func (r *Room) Remove(obj Objecter) bool {
	var i int
	for i = 0; i < len(r.inventory); i++ {
		if obj == r.inventory[i] {
			break;
		}
	}

	if i == len(r.inventory) {
		return false
	}

	copy(r.inventory[i:], r.inventory[i+1:])
	r.inventory[len(r.inventory) - 1] = nil
	r.inventory = r.inventory[:len(r.inventory) - 1]
	return true
}
