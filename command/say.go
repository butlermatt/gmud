package command

import (
	"strings"
	"fmt"
)

func init() {
	addCommand(Command{
		name: "say",
		help: "allows you to speak to the room.",
		exec: sayCmd,
	})
}

func sayCmd(user Player, args []string) (ok bool) {
	if len(args) == 0 {
		return false
	}

	str := strings.Join(args, " ")

	inv := user.Room().Inventory()
	var tmp []Player
	for _, i := range inv {
		p, ok := i.(Player)
		if ok {
			tmp = append(tmp, p)
		}
	}

	for _, p := range tmp {
		if p == user {
			user.Write(fmt.Sprintf("you say: %s", str))
		} else {
			p.Write(fmt.Sprintf("%s says: %s", user.Name(), str))
		}
	}

	return true
}