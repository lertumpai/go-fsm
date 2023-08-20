package app

import (
	"go-fsm/state/backup"
)

func Start() {
	ConfigInit()
	b := backup.CreateBackup().Init()
	b.FireEventStartBackup()
	b.FireEventFinishExtract()
	b.FireEventFinishUpload()
	b.PrintGraph()
}
