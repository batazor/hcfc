package generate

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/helm/pkg/engine"
	"os"
	"strconv"
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
	cfg.Template.Filename = "Chart.yaml"
	if err := createFile(cfg, "Chart.yaml", cfg.Template.Values.Chart); err != nil {
		return err
	}

	cfg.Template.Filename = "values.yaml"
	if err := createValuesFile(cfg, "values.yaml", cfg.Template.Values); err != nil {
		return err
	}

	cfg.Template.Filename = "templates/NOTES.txt"
	if err := createFile(cfg, "templates/NOTES.txt", cfg.Template.Values.Chart); err != nil {
		return err
	}

	cfg.Template.Filename = ".helmignore"
	if err := createFile(cfg, ".helmignore", cfg.Template.Values.Chart); err != nil {
		return err
	}

	for index, secretConfig := range cfg.Template.Values.Secret {
		cnf := struct {
			Chart
			Secret
			Index int
		}{
			cfg.Template.Values.Chart,
			secretConfig,
			index,
		}
		templateName := "templates/secret.yaml"

		if len(cfg.Template.Values.Secret) > 1 {
			cnf.Chart.Name = strings.Join([]string{cnf.Chart.Name, "-", strconv.Itoa(index)}, "")
			cfg.Template.Filename = strings.Join([]string{"templates/secret-", strconv.Itoa(index), ".yaml"}, "")
		} else {
			cfg.Template.Filename = templateName
		}

		if err := createFile(cfg, templateName, cnf); err != nil {
			return err
		}
	}

	for index, deploymentConfig := range cfg.Template.Values.Deployment {
		cnf := struct {
			Chart
			Deployment
			Index int
		}{
			cfg.Template.Values.Chart,
			deploymentConfig,
			index,
		}
		templateName := "templates/deployment.yaml"

		if len(cfg.Template.Values.Deployment) > 1 {
			cnf.Chart.Name = strings.Join([]string{cnf.Chart.Name, "-", strconv.Itoa(index)}, "")
			cfg.Template.Filename = strings.Join([]string{"templates/deployment-", strconv.Itoa(index), ".yaml"}, "")
		} else {
			cfg.Template.Filename = templateName
		}

		if err := createFile(cfg, templateName, cnf); err != nil {
			return err
		}
	}

	for index, serviceConfig := range cfg.Template.Values.Service {
		cnf := struct {
			Chart
			Service
			Index int
		}{
			cfg.Template.Values.Chart,
			serviceConfig,
			index,
		}
		templateName := "templates/service.yaml"

		if len(cfg.Template.Values.Service) > 1 {
			cnf.Chart.Name = strings.Join([]string{cnf.Chart.Name, "-", strconv.Itoa(index)}, "")
			cfg.Template.Filename = strings.Join([]string{"templates/service-", strconv.Itoa(index), ".yaml"}, "")
		} else {
			cfg.Template.Filename = templateName
		}

		if err := createFile(cfg, templateName, cnf); err != nil {
			return err
		}
	}

	for index, ingressConfig := range cfg.Template.Values.Ingress {
		cnf := struct {
			Chart
			Ingress
			Index int
		}{
			cfg.Template.Values.Chart,
			ingressConfig,
			index,
		}
		templateName := "templates/ingress.yaml"

		if len(cfg.Template.Values.Ingress) > 1 {
			cnf.Chart.Name = strings.Join([]string{cnf.Chart.Name, "-", strconv.Itoa(index)}, "")
			cfg.Template.Filename = strings.Join([]string{"templates/ingress-", strconv.Itoa(index), ".yaml"}, "")
		} else {
			cfg.Template.Filename = templateName
		}

		if err := createFile(cfg, templateName, cnf); err != nil {
			return err
		}
	}

	return nil
}

func createValuesFile(cfg BuildConfig, templateName string, values interface{}) error {
	response, err := yaml.Marshal(&values)

	// Create path to file
	s := []string{cfg.Template.Output, cfg.Template.Filename}
	pathOutput := strings.Join(s, "/")

	// Check directory
	s = []string{}
	path := strings.Split(pathOutput, "/")
	for _, v := range path[:len(path)-1] {
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

func createFile(cfg BuildConfig, templateName string, values interface{}) error {
	// Get template file
	err, template := cfg.Templates.Get(templateName)
	if err != nil {
		return err
	}

	err, response := ExecuteTemplates(template, values)
	if err != nil {
		return err
	}

	// Create path to file
	s := []string{cfg.Template.Output, cfg.Template.Filename}
	pathOutput := strings.Join(s, "/")

	// Check directory
	s = []string{}
	path := strings.Split(pathOutput, "/")
	for _, v := range path[:len(path)-1] {
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
func ExecuteTemplates(file string, values interface{}) (error, string) {
	newTemplate := "hcfc"
	FuncMap := engine.FuncMap()

	FuncMap["ignore"] = Ignore
	FuncMap["text"] = Text

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

func Text(v interface{}) string {
	return strings.Join([]string{v.(string)}, " ")
}
