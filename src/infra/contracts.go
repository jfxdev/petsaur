package infra

type Migration interface {
	Run() error
}
