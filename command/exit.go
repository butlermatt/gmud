package command

import "github.com/butlermatt/gmud/lib"

func init() {
	addCommand(Command{
		name: "exit",
		help: "disconnects you from the mud.",
		exec: exitCmd,
	})
}

func exitCmd(user lib.Player, args []string) (ok bool) {
	if len(args) != 0 {
		return false
	}

	user.Quit()

	return true
}
