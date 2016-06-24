package main

import "os"
import "io/ioutil"
import "path/filepath"
import "text/template"
import "strings"
import "fmt"

type Catalog struct {
	CatalogPath   string
	AccessControl string

	// generated properties
	EnumName     string
	EnumInitName string
	Assets       map[string]string
}

// Load the catalog from .CatalogPath and set generated properties accordingly.
func NewCatalog(catalogPath string, accessControl string) *Catalog {
	var assets = make(map[string]string)

	files, _ := ioutil.ReadDir(catalogPath)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".imageset" {
			filename := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			name := cammelCase(filename)

			assets[name] = filename
		}
	}

	enumName := cammelCase(strings.TrimSuffix(filepath.Base(catalogPath), filepath.Ext(catalogPath))) + "Asset"
	enumInitName := strings.ToLower(enumName[0:1]) + enumName[1:]

	return &Catalog{CatalogPath: catalogPath, AccessControl: accessControl, EnumName: enumName, EnumInitName: enumInitName, Assets: assets}
}

// Write the Swift source code to a file.
func (c *Catalog) writeEnum() {
	tmpl, err := template.New("swift").Parse(defaultSwiftTemplate)
	if err != nil {
		panic(err)
	}

	outPath := filepath.Join(filepath.Dir(c.CatalogPath), c.EnumName+".swift")
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, c)
	if err != nil {
		panic(err)
	}
}

// Convert a string with varying word delimination to camel case.
// Examples: more: More, trash-activity: TrashActivity, welcome_placeholder: WelcomePlaceholder
func cammelCase(input string) string {
	capatalizeNext := true
	output := ""

	var allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for _, char := range input {
		charString := fmt.Sprintf("%c", char)

		if strings.Contains(allowed, charString) {
			if capatalizeNext {
				output = output + strings.ToUpper(charString)
				capatalizeNext = false
			} else {
				output = output + charString
			}
		} else {
			capatalizeNext = true
		}
	}
	return output
}
