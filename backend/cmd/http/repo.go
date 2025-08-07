package main

import (
	"booklib/internal/domain/book"
	"booklib/internal/infra"
	repobook "booklib/internal/repo/book"
)

type Repo struct {
	Book book.Repository
}

func newRepo(res *infra.Resources) *Repo {
	return &Repo{
		Book: repobook.New(res.Database),
	}
}
