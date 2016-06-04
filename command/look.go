package command

func init() {
	addCommand(Command{
		name: "look",
		help: "look around. If an object is specified it will look closely at that object.",
		exec: lookCmd,
	})
}

func lookCmd(user Player, args []string) (ok bool) {
	if len(args) == 0 {
		str := user.Room().Description()
		user.Write(str)
		return true
	}

	return true
}
