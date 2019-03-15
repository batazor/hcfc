package generate

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"text/template"
)

func Build(filePath, output string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var helmChart interface{}
	err = yaml.Unmarshal([]byte(data), &helmChart)
	if err != nil {
		return err
	}

	err, response := ExecuteTemplates("./templates/Chart.yaml", &helmChart)
	if err != nil {
		return err
	}

	path := strings.Join([]string{output, "/Chart.yaml"}, "")
	err = ioutil.WriteFile(path, []byte(response), 0644)
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
