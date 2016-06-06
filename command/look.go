package command

import (
	"fmt"
	"strings"

	"github.com/butlermatt/gmud/lib"
)

func init() {
	addCommand(Command{
		name: "look",
		help: "look around. If an object is specified it will look closely at that object.",
		exec: lookCmd,
	})
}

func lookCmd(user lib.Player, args []string) (ok bool) {
	if len(args) == 0 {
		room := user.Room()
		str := room.Description()
		var inv []lib.Objecter
		for _, obj := range room.Inventory() {
			if o, ok := obj.(lib.Player); ok && o == user {
				continue
			}

			inv = append(inv, obj)
		}

		if len(inv) > 0 {
			str += "\r\nYou also see here:\r\n"
			for _, obj := range inv {
				str += obj.Name() + "\r\n"
			}
		}

		user.Write(str)
		return true
	}

	var objName string
	if args[0] == "at" {
		objName = strings.Join(args[1:], " ")
	} else {
		objName = strings.TrimSpace(strings.Join(args[0:], " "))
	}

	for _, obj := range user.Inventory() {
		if objName == obj.Name() {
			user.Write(obj.Description())
			return true
		}
	}

	for _, obj := range user.Room().Inventory() {
		if objName == obj.Name() {
			user.Write(obj.Description())
			return true
		}
	}

	str := fmt.Sprintf("Can't see any %s here", objName)
	user.Write(str)
	return false
}
