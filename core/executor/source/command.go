package source

import "github.com/baojiweicn/Surge/util/parser"

type Command struct {
	cmd []string
}

func NewCommand(args ...string) Command {
	c := Command{cmd: make([]string, 0)}
	for _, i := range args {
		c.cmd = append(c.cmd, i)
	}
	return c
}

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
