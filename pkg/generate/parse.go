package generate

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
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
	err := filepath.Walk(cfg.Template, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() == false {
			err, response := ExecuteTemplates(path, config)
			if err != nil {
				return err
			}

			// Create path to file
			s := []string{cfg.Output}
			for _, v := range strings.Split(path, "/")[1:] {
				s = append(s, v)
			}
			pathOutput := strings.Join(s, "/")

			// Check directory
			s = []string{cfg.Output}
			for _, v := range strings.Split(path, "/")[:1] {
				s = append(s, v)
			}
			dirOutput := strings.Join(s, "/")
			if _, err := os.Stat(dirOutput); os.IsNotExist(err) {
				os.MkdirAll(dirOutput, os.ModePerm)
			}

			err = ioutil.WriteFile(pathOutput, []byte(response), 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
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
