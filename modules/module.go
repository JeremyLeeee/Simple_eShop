package modules

type IModule interface {
	GetModuleName() string
	GetController() interface{}
	GetRepository() interface{}
	GetService() interface{}
}
