package main

import (
	"fmt"
)

var (
	build Version

	commit string
	branch string
	date   string
)

func init() {
	build = Version{
		Commit: commit,
		Branch: branch,
		Date:   date,
	}
}

type Version struct {
	Commit string
	Branch string
	Date   string
}

func (v Version) String() string {
	return fmt.Sprintf("%s (branch: %s, build date: %s)",
		v.Commit, v.Branch, v.Date)
}
