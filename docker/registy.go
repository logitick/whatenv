package docker

type Tag string;

type Image struct {
	Id string
	Tags []Tag
}
type Registry interface {

}


