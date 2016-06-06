package command

import "github.com/butlermatt/gmud/lib"

func init() {
	addCommand(Command{
		name: "shutdown",
		help: "shuts down the MUD server.",
		exec: shutdownCmd,
		log:  true,
	})
}

func shutdownCmd(user lib.Player, args []string) (ok bool) {
	if len(args) != 0 {
		return false
	}

	return true
}
