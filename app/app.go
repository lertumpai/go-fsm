package app

import "go-fsm/state/manager"

func Start() {
	ConfigInit()
	m := manager.CreateManager().Init()
	m.FireProcessing()
}
