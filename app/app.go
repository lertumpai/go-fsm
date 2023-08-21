package app

import (
	"fmt"
	"go-fsm/state/backup"
	"go-fsm/state/manager"
)

func Start() {
	ConfigInit()
	b := backup.Create().Init()
	m := manager.Create(manager.Config{
		Backup: b,
	}).Init()

	fmt.Println("task 1")
	m.FireEventStartBackup()
	fmt.Println("task 2")
	m.FireEventStartBackup()
}
