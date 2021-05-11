package main

import (
	"fmt"
	"strings"
)

type APIResp map[string][]Note

type Note struct {
	Author   Author `json:"author"`
	Content  string `json:"content"`
	Comments []Note `json:"comments"`
}

func (n *Note) Initials() string {
	splits := strings.Split(n.Content, ":")
	if len(splits) > 1 {
		if splits[0] == strings.ToUpper(splits[0]) {
			n.Content = strings.Trim(strings.Join(splits[1:], ":"), " ")
			return splits[0]
		}
	}

	return n.Author.Initials()
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a *Author) Initials() string {
	splits := strings.Split(a.Name, " ")
	initials := ""
	for _, split := range splits[:2] {
		initials += string(split[0])
	}

	return strings.ToUpper(initials)
}

func (r APIResp) ToMarkdown(sections []string) string {
	var md string
	for _, section := range sections {
		notes := r[section]

		if len(notes) == 0 {
			continue
		}

		md += fmt.Sprintf("\n%s\n", section)

		for _, note := range notes {
			md += fmt.Sprintf("  * %s: %s\n", note.Author.Initials(), note.Content)
			for _, comment := range note.Comments {
				md += fmt.Sprintf("    * %s: %s\n", comment.Initials(), comment.Content)
			}
		}
	}

	return md
}
