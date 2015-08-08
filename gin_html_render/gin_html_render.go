package gin_html_render

// This work is based on gin contribs multitemplate render
// https://github.com/gin-gonic/contrib/blob/master/renders/multitemplate/multitemplate.go

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Render struct {
	Files     map[string][]string
	Templates map[string]*template.Template
	Mode      string
}

const (
	// TemplateDir holds the location of the templates
	TemplateDir = "templates"
)

var (
	// Mode is gin's env mode name
	Mode = os.Getenv(gin.ENV_GIN_MODE)
)

// Add sets the template
func (r Render) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	r.Templates[name] = tmpl
}

// AddFromFiles parses the files and returns the result
func (r Render) AddFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.ParseFiles(files...))
	if r.isDebug() {
		r.Files[name] = files
	}
	r.Add(name, tmpl)
	return tmpl
}

// Instance implements gin's HTML render interface
func (r Render) Instance(name string, data interface{}) render.Render {
	var tpl *template.Template

	if r.isDebug() {
		tpl = r.loadTemplate(name)
	} else {
		tpl = r.Templates[name]
	}

	return render.HTML{
		Template: tpl,
		Data:     data,
	}
}

// loadTemplate parses the specified template and returns it
func (r Render) loadTemplate(name string) *template.Template {
	return template.Must(template.ParseFiles(r.Files[name]...))
}

// isDebug checks if debug mode is active
func (r Render) isDebug() bool {
	return r.Mode == gin.DebugMode
}

// New returns a fresh instance of Render
func New() Render {
	return Render{
		Files:     make(map[string][]string),
		Templates: make(map[string]*template.Template),
		Mode:      Mode,
	}
}

// Create goes through the `TemplateDir` creating the template structure
// for rendering. Returns the Render instance.
// TODO: provide a way to customize template dir, layout dir
func (r Render) Create() Render {

	r.AddFromFiles("400.html", TemplateDir+"/layout.html", TemplateDir+"/400.html")

	tpls, err := filepath.Glob(TemplateDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, tpl := range tpls {
		r.AddFromFiles(filepath.Base(tpl), TemplateDir+"/layout.html", tpl)
	}

	return r
}
