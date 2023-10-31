package model

type ModelI interface {
	SetController(ControllerI)
	InjectAPI()
	SetID(id string) error
	GetID() string
	CreateValidation() (bool, map[string]string)
	UpdateValidation() (bool, map[string]string)
	BeforeValidation()
	AfterValidation()
	OnCreate()
	OnUpdate()
	OnDelete()
}
