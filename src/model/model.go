package model

type ModelI interface {
	SetController(ControllerI)
	InjectAPI()
	GetID() uint
	SetID(id uint)
	CreateValidation() (bool, map[string]string)
	UpdateValidation() (bool, map[string]string)
	BeforeValidation()
	AfterValidation()
}
