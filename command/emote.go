package command

import (
	"fmt"
	"strings"

	"github.com/butlermatt/gmud/lib"
)

func init() {
	addCommand(Command{
		name: "emote",
		help: "allows you to send an action to the room.",
		exec: emoteCmd,
	})
}

func emoteCmd(user lib.Player, args []string) (ok bool) {
	if len(args) == 0 {
		return false
	}

	str := strings.Join(args, " ")

	inv := user.Room().Inventory()
	var tmp []lib.Player
	for _, i := range inv {
		p, ok := i.(lib.Player)
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
