package main

import (
	"os"
)

func segmentShellVar(p *powerline) {
	shellVarName := *p.args.ShellVar
	varContent, varExists := os.LookupEnv(shellVarName)

	if !varExists || varContent == "" {
		return
	}

	p.appendSegment("shell-var", segment{
		content:    varContent,
		foreground: p.theme.ShellVarFg,
		background: p.theme.ShellVarBg,
	})
}
