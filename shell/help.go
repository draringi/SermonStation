package shell

type help struct {
}

func (cmd *help) Name() string {
	return "help"
}

func (cmd *help) Description() string {
	return "help can be called with either 0 or 1 argument.\nWith no arguments, in lists all available commands.\nWith 1 command as an argument, it provides a description of the argument"
}

func (cmd *help) Call(shell *Shell, args []string) {
	switch args.Size {
	case 0:
		cmds := commands.Keys()
		shell.PrintList(cmds)
		break
	case 1:
		arg := args[0]
		command, err := commands[arg]
		if err != nil {
			shell.Println("Error: '" + arg + "' does not exist")
			break
		}
		shell.Println(command.Name())
		shell.Println(command.Description())
		break
	default:
		shell.Println("Error: Too many arguments")
	}
}
