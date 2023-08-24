package app

import (
	"fmt"
	"go-fsm/app/state/backup"
	"go-fsm/app/state/manager"
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
