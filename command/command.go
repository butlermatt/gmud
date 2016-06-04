package command

import (
	"fmt"
	"github.com/butlermatt/gmud/lib"
)

// Player is an in game character player.
type Player interface {
	lib.Objecter
	// Write sends a message to the Player in a non-blocking way.
	Write(string)
	// Send sends a message to the Player in a potentially blocking way.
	Send(string)
	// Room returns the room of the current player.
	Room() lib.Holder
	Quit()
}

// Commands should implement this interface.
type Commander interface {
	// Exec executes the command.
	Exec()
	// Name returns the name of the command (and the keyword to invoke it)
	Name() string
	// Help returns a formatted help string for the command.
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

// Help returns a formatted help string for the command.
func (c Command) Help() string {
	return fmt.Sprintf("%s\n\t%s", c.name, c.help)
}

var commands = make(map[string]Command)

// addCommand will add the specified Command to the command parser.
func addCommand(cmd Command) {
	commands[cmd.name] = cmd
}

// GetCommand takes the player and their input and creates a new Command instance
// If the command is invalid, it returns an error.
func GetCommand(user Player, cmdArgs []string) (Commander, error) {
	cmd := commands[cmdArgs[0]]

	if cmd.exec == nil {
		return nil, fmt.Errorf("Unknown command %s", cmdArgs[0])
	}

	cmd.User = user
	cmd.args = cmdArgs[1:]
	return cmd, nil
}
