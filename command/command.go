package command

import (
	"fmt"
)

type Player interface {
	Write(string)
	Quit()
}

type Commander interface {
	Exec()
	Name() string
	Help() string
}

type execFunc func(user Player, args []string) (ok bool)

// Command holds the command name, help string, User executing command, as well as
// function to execute and a string slice passed to the command.
type Command struct {
	name string
	help string
	User Player
	exec execFunc
	args []string
}

// Name returns the command name
func (c Command) Name() string {
	return c.name
}

// Exec executes the method associated with this command.
func (c Command) Exec() {
	if ok := c.exec(c.User, c.args); !ok {
		c.User.Write("Usage: " + c.Help())
	}
}

func (c Command) Help() string {
	return fmt.Sprintf("%s\n\t%s", c.name, c.help)
}

var commands = make(map[string]Command)

// addCommand will add the specified Command to the command parser.
func addCommand(cmd Command) {
	commands[cmd.name] = cmd;
}

func GetCommand(user Player, cmdArgs []string) (Commander, error) {
	cmd := commands[cmdArgs[0]]

	if cmd.exec == nil {
		return nil, fmt.Errorf("Unknown command %s", cmdArgs[0])
	}

	cmd.User = user
	cmd.args = cmdArgs[1:]
	return cmd, nil
}
