package shell

var commands = make(map[string]Command)

type Command interface {
	Name() string
	Description() string
	Call(shell *Shell, args []string)
}
