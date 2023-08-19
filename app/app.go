package app

import "go-fsm/app/state/manager"

func Start() {
	ConfigInit()
	m := manager.CreateManager()
	m.Init()
	m.FireProcessing()
}
