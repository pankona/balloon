package scene

type actioner interface {
	doAction()
	finalize()
}
