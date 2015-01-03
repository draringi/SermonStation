package shell

import (
	"io"
)

type Shell struct{
	stdin *io.Reader
	stdout *io.Writer
	running bool
}

func CreateShell(stdin *io.Reader, stdout *io.Writer) (shell *Shell){
	shell = new(Shell)
	shell.stdin = stdin
	shell.stdout = stdout
	shell.running = true
	go shell.run()
	return
}

func (shell *Shell) run(){
	while(running){

	}
}
