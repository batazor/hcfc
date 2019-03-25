package generate

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/helm/pkg/engine"
	"os"
	"strings"
	"text/template"
)

func Build(config BuildConfig) error {
	data, err := ioutil.ReadFile(config.Filename)
	if err != nil {
		return err
	}

	var helmChart Project
	err = yaml.Unmarshal([]byte(data), &helmChart)
	if err != nil {
		return err
	}

	if err = BuildFile(config, helmChart); err != nil {
		return err
	}

	return nil
}

func BuildFile(cfg BuildConfig, config Project) error {
	if err := createFile("templates/Chart.yaml", cfg.Output, config.Chart); err != nil {
		return err
	}

	if err := createFile("templates/templates/NOTES.txt", cfg.Output, config.Chart); err != nil {
		return err
	}

	if err := createFile("templates/values.yaml", cfg.Output, config.Chart); err != nil {
		return err
	}

	if err := createFile("templates/.helmignore", cfg.Output, config.Chart); err != nil {
		return err
	}

	for _, deploymentConfig := range config.Deployment {
		cnf := struct {
			Chart
			Deployment
		}{
			config.Chart,
			deploymentConfig,
		}

		if err := createFile("templates/templates/deployment.yaml", cfg.Output, cnf); err != nil {
			return err
		}
	}

	for _, serviceConfig := range config.Service {
		cnf := struct {
			Chart
			Service
		}{
			config.Chart,
			serviceConfig,
		}

		if err := createFile("templates/templates/service.yaml", cfg.Output, cnf); err != nil {
			return err
		}
	}

	for _, ingressConfig := range config.Ingress {
		cnf := struct {
			Chart
			Ingress
		}{
			config.Chart,
			ingressConfig,
		}

		if err := createFile("templates/templates/ingress.yaml", cfg.Output, cnf); err != nil {
			return err
		}
	}

	return nil
}

func createFile(path, output string, config interface{}) error {
	err, response := ExecuteTemplates(path, config)
	if err != nil {
		return err
	}

	// Create path to file
	s := []string{output}
	for _, v := range strings.Split(path, "/")[1:] {
		s = append(s, v)
	}
	pathOutput := strings.Join(s, "/")

	// Check directory
	s = []string{output}
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

	return nil
}

// Reads a YAML document from the values_in stream, uses it as values
// for the tpl_files templates and writes the executed templates to
// the out stream.
func ExecuteTemplates(file string, config interface{}) (error, string) {
	newTemplate := strings.Split(file, "/")
	FuncMap := engine.FuncMap()

	FuncMap["ignore"] = Ignore

	tpl, err := template.New(newTemplate[len(newTemplate)-1]).Funcs(FuncMap).ParseFiles(file)
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

func Ignore(v interface{}) string {
	return strings.Join([]string{"{{", v.(string), "}}"}, " ")
}
