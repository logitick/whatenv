package environments

type Environment struct {
	name string
}

type EnvironmentService interface {
	Read(name string) *Environment
	List() []*Environment
}