package shell

var commands = make(map[string]Command)

type Command interface {
	Name() string
	Description() string
	ArgumentCount() int
	Call(args []string)
}
