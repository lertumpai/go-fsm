package app

import "go-fsm/state/manager"

func Start() {
	ConfigInit()
	m := manager.CreateManager()
	m.Init()
	m.FireProcessing()
}
