package model

type ModelI interface {
	SetController(ControllerI)
	InjectAPI()
	SetID(id string) error
	GetID() string
	GetName() string
	CreateValidation() (bool, map[string]string)
	UpdateValidation() (bool, map[string]string)
	BeforeValidation()
	AfterValidation()
	OnCreate() (status int, err error)
	OnUpdate() (status int, err error)
	OnDelete() (status int, err error)
}
