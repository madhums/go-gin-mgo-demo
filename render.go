package main

import (
	"path/filepath"

	"github.com/gin-gonic/contrib/renders/multitemplate"
)

func render() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("400.html", "templates/layout.html", "templates/400.html")

	tpls, err := filepath.Glob("templates/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, tpl := range tpls {
		r.AddFromFiles(filepath.Base(tpl), "templates/layout.html", tpl)
	}

	return r
}
