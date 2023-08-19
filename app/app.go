package app

func Start() {
	ConfigInit()
	manager := CreateManager()
	manager.Init()
	manager.FireProcessing()
}
