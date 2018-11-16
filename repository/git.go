package repository

import "net/url"

type Commit struct {
	Id string
	Name string
	Message string
}

type Tag struct {
	Name string
	Commit Commit
}

type Git struct {
	remote url.URL
}

func (g *Git) Clone()  {

}

func (g *Git) Log(limit int)  []*Commit {
	return nil
}
