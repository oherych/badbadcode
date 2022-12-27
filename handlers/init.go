package handlers

import "github.com/oherych/badbadcode/pkg"

var tm *pkg.TaskManager

func init() {
	tm := new(pkg.TaskManager)
	tm.List = make(map[pkg.Task]string, 1000)
}
