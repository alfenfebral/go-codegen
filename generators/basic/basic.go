package basic

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type BasicTemplateData struct {
	Name string
}

func Generate(name, dir string) error {
	// Ensure directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Read template from file
	tmplContent, err := os.ReadFile("templates/basic.tmpl")
	if err != nil {
		return fmt.Errorf("error reading template: %w", err)
	}

	data := struct {
		Package    string
		StructName string
	}{
		Package:    filepath.Base(dir),
		StructName: name,
	}

	filename := filepath.Join(dir, strings.ToLower(name)+".go")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	tmpl, err := template.New("basic").Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	fmt.Printf("Generated file: %s\n", filename)
	return nil
}
