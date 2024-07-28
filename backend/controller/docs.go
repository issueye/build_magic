package controller

import (
	"context"

	"github.com/issueye/build_magic/backend/logic"
)

var docs *Docs

type Docs struct {
	Ctx context.Context
}

func GetDocs() *Docs {
	if docs == nil {
		docs = &Docs{}
	}

	return docs
}

func (docs *Docs) GetDocsTree() string {
	return new(logic.Docs).GetDocsTree()
}

func (docs *Docs) GetDocsContent(path string) string {
	return new(logic.Docs).GetDocsContent(path)
}
