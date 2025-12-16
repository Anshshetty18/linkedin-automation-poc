package stealth

type Behavior interface {
	BeforeAction() error
	AfterAction() error
}
