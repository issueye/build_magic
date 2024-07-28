package logic

import (
	"embed"
	"fmt"

	"github.com/issueye/build_magic/backend/pkg/utils"
)

//go:embed docs/*
var docData embed.FS

type Docs struct{}

func (docs *Docs) GetDocsTree() string {
	return docs.GetDocsContent("docs/docs.json")
}

func (docs *Docs) GetDocsContent(path string) string {
	d, err := docData.ReadFile(path)
	if err != nil {
		return ""
	}

	fmt.Println("path: ", path, "data", string(d))

	return utils.Bytes2String(d)
}
