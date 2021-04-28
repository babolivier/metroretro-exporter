package main

import (
	"fmt"
	"regexp"
	"strings"
)

type APIResp map[string][]Note

type Note struct {
	Author   Author `json:"author"`
	Content  string `json:"content"`
	Comments []Note `json:"comments"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a *Author) Initials() string {
	r := regexp.MustCompile("[A-Z]")
	res := r.FindAllString(a.Name, 500)

	if len(res) > 0 {
		return strings.Join(res, "")
	}

	return strings.ToUpper(a.Name[:1])
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
			md += fmt.Sprintf("    * %s: %s\n", note.Author.Initials(), note.Content)
			for _, comment := range note.Comments {
				md += fmt.Sprintf("        * %s: %s\n", comment.Author.Initials(), comment.Content)
			}
		}
	}

	return md
}
