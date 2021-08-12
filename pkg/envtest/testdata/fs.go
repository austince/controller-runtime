package testdata

import (
	"embed"
	"io/fs"
)

//go:embed *.yaml crds/*.yaml crdv1_original/*.yaml crdv1_updated/*.yaml webhooks/*.yaml
var testData embed.FS

func FS() fs.FS {
	return testData
}

func SubFS(dir string) fs.FS {
	fsys, err := fs.Sub(testData, dir)
	if err != nil {
		panic(err)
	}
	return fsys
}
