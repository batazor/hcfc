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

	err = yaml.Unmarshal([]byte(data), &config.Template.Values)
	if err != nil {
		return err
	}

	if err = BuildFile(config); err != nil {
		return err
	}

	return nil
}

func BuildFile(cfg BuildConfig) error {
	if err := createFile(cfg, "Chart.yaml", cfg.Template.Values.Chart); err != nil {
		return err
	}

	if err := createFile(cfg, "values.yaml", cfg.Template.Values.Chart); err != nil {
		return err
	}

	if err := createFile(cfg, "templates/NOTES.txt", cfg.Template.Values.Chart); err != nil {
		return err
	}

	if err := createFile(cfg, "templates/.helmignore", cfg.Template.Values.Chart); err != nil {
		return err
	}

	for _, deploymentConfig := range cfg.Template.Values.Deployment {
		cnf := struct {
			Chart
			Deployment
		}{
			cfg.Template.Values.Chart,
			deploymentConfig,
		}

		if err := createFile(cfg, "templates/deployment.yaml", cnf); err != nil {
			return err
		}
	}

	for _, serviceConfig := range cfg.Template.Values.Service {
		cnf := struct {
			Chart
			Service
		}{
			cfg.Template.Values.Chart,
			serviceConfig,
		}

		if err := createFile(cfg, "templates/service.yaml", cnf); err != nil {
			return err
		}
	}

	for _, ingressConfig := range cfg.Template.Values.Ingress {
		cnf := struct {
			Chart
			Ingress
		}{
			cfg.Template.Values.Chart,
			ingressConfig,
		}

		if err := createFile(cfg, "templates/ingress.yaml", cnf); err != nil {
			return err
		}
	}

	return nil
}

func createFile(cfg BuildConfig, nameFile string, values interface{}) error {
	// Get template file
	err, template := cfg.Templates.Get(nameFile)
	if err != nil {
		return err
	}

	err, response := ExecuteTemplates(template, values)
	if err != nil {
		return err
	}

	// Create path to file
	s := []string{cfg.Template.Output, nameFile}
	pathOutput := strings.Join(s, "/")
	//fmt.Println("pathOutput", pathOutput)

	// Check directory
	s = []string{}
	path := strings.Split(pathOutput, "/")
	for _, v := range path[:len(path)-1] {
		s = append(s, v)
	}
	dirOutput := strings.Join(s, "/")
	fmt.Println("dirOutput", dirOutput)
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
func ExecuteTemplates(file string, values interface{}) (error, string) {
	newTemplate := "hcfc"
	FuncMap := engine.FuncMap()

	FuncMap["ignore"] = Ignore

	tpl, err := template.New(newTemplate).Funcs(FuncMap).Parse(file)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err), ""
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err), ""
	}

	return nil, buf.String()
}

func Ignore(v interface{}) string {
	return strings.Join([]string{"{{", v.(string), "}}"}, " ")
}
