package command

import (
	"os/exec"

	"github.com/baojiweicn/Surge/util/parser"
)

type Command struct {
	name string
	cmd  []string
}

// NewCommand : Create a command with executor and template
func NewCommand(name string, args ...string) Command {
	c := Command{cmd: make([]string, 0), name: name}
	for _, i := range args {
		c.cmd = append(c.cmd, i)
	}
	return c
}

// NewCommandTemplate : Create a command template without executor
func NewCommandTemplate(args ...string) Command {
	c := Command{cmd: make([]string, 0)}
	for _, i := range args {
		c.cmd = append(c.cmd, i)
	}
	return c
}

// Generate : create execable command for exec.Command
func (c Command) Generate(fields []parser.Field) *exec.Cmd {
	return exec.Command(c.name, c.Render(fields)...)
}

// SetExecutor : set executor path
func (c Command) SetExecutor(path string) Command {
	c.name = path
	return c
}

// Render : fulfill args into template
func (c Command) Render(fields []parser.Field) []string {
	res := make([]string, 0)
	for _, i := range c.cmd {
		match := false
		for _, field := range fields {
			if i == "{{"+field.Key()+"}}" {
				res = append(res, field.Value())
				match = true
				break
			}
		}
		if !match {
			res = append(res, i)
		}
	}
	return res
}

// Executor : get executor path
func (c Command) Executor() string {
	return c.name
}
