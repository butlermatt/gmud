package command

import (
	"fmt"
	"strings"
)

func init() {
	addCommand(Command{
		name: "emote",
		help: "allows you to send an action to the room.",
		exec: emoteCmd,
	})
}

func emoteCmd(user Player, args []string) (ok bool) {
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
			user.Write(fmt.Sprintf("you emote: %s %s", user.Name(), str))
		} else {
			p.Write(fmt.Sprintf("%s %s", user.Name(), str))
		}
	}

	return true
}
