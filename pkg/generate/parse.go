package generate

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"text/template"
)

func Build(config BuildConfig) error {
	data, err := ioutil.ReadFile(config.Filename)
	if err != nil {
		return err
	}

	var helmChart interface{}
	err = yaml.Unmarshal([]byte(data), &helmChart)
	if err != nil {
		return err
	}

	if err = RecursiveBuildFile(config, &helmChart); err != nil {
		return err
	}

	return nil
}

func RecursiveBuildFile(cfg BuildConfig, config interface{}) error {
	files, err := ioutil.ReadDir(cfg.Template)
	if err != nil {
		return err
	}

	for _, f := range files {
		pathTemplate := strings.Join([]string{cfg.Template, f.Name()}, "/")
		err, response := ExecuteTemplates(pathTemplate, config)
		if err != nil {
			return err
		}

		pathOutput := strings.Join([]string{cfg.Output, f.Name()}, "/")
		err = ioutil.WriteFile(pathOutput, []byte(response), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Reads a YAML document from the values_in stream, uses it as values
// for the tpl_files templates and writes the executed templates to
// the out stream.
func ExecuteTemplates(file string, config interface{}) (error, string) {
	tpl, err := template.ParseFiles(file)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err), ""
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, config)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err), ""
	}

	return nil, buf.String()
}
