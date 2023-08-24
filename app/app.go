package app

import (
	"go-fsm/app/state/backup"
	"go-fsm/app/state/manager"
	"go-fsm/config"
)

func Start() {
	config.Load()
	b := backup.Create().Init()
	m := manager.Create(manager.Config{
		Backup: b,
	}).Init()

	m.FireEventStartBackup()
}
