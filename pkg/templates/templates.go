package templates

import (
	"github.com/gobuffalo/packr"
)

type Template struct {
	Box packr.Box
}

func (t *Template) Init(name, path string) {
	// set up a new box by giving it a name and an optional (relative) path to a folder on disk:
	t.Box = packr.NewBox("../../templates")
}

func (t *Template) Get(name string) (error, string) {
	// Get the string representation of a file, or an error if it doesn't exist:
	file, err := t.Box.FindString(name)
	return err, file
}
