package command

func init() {
	addCommand(Command{
		name: "shutdown",
		help: "shuts down the MUD server.",
		exec: shutdownCmd,
		log:  true,
	})
}

func shutdownCmd(user Player, args []string) (ok bool) {
	if len(args) != 0 {
		return false
	}

	return true
}
