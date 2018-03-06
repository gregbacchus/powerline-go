package main

import (
	"os"
)

func segmentUser(p *powerline) {
	var userPrompt string
	if *p.args.Shell == "bash" {
		userPrompt = "\\u"
	} else if *p.args.Shell == "zsh" {
		userPrompt = "%n"
	} else {
		user, _ := os.LookupEnv("USER")
		userPrompt = user
	}

	var background uint8
	if os.Getuid() == 0 {
		background = p.theme.UsernameRootBg
		userPrompt = p.symbolTemplates.RootUser
	} else {
		background = p.theme.UsernameBg
		// r, _ := regexp.Compile("([\\w.-_]+\\\\)?(.+)(@[\\w.-_]+)?")
		// userPrompt = r.FindAllStringSubmatch(userPrompt, -1)[0][2]
	}

	p.appendSegment("user", segment{
		content:    userPrompt,
		foreground: p.theme.UsernameFg,
		background: background,
	})
}
