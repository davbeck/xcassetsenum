package main

import "os"
import "io/ioutil"
import "path/filepath"
import "text/template"
import "github.com/mkideal/cli"
import "strings"
import "fmt"

type Catalog struct {
	cli.Helper
	CatalogPath   string `cli:"*c,catalog" usage:"The xcassets catalog to process into a Swift enum. Required."`
	AccessControl string `cli:"a,access_control" usage:"The access to grant for the generated enum, such as public, private, internal. Defaults to nothing, which causes Swift to use internal." dft:"internal"`
	SwiftVersion  string `cli:"s,swift_version" usage:"The version of swift to generate." dft:"2.3"`

	// generated properties
	EnumName     string
	EnumInitName string
	Assets       map[string]string
}

// Load the catalog from .CatalogPath and set generated properties accordingly.
func (c *Catalog) loadAssets() {
	var assets = make(map[string]string)

	files, _ := ioutil.ReadDir(c.CatalogPath)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".imageset" {
			filename := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			name := cammelCase(filename)

			assets[name] = filename
		}
	}

	c.Assets = assets
	c.EnumName = cammelCase(strings.TrimSuffix(filepath.Base(c.CatalogPath), filepath.Ext(c.CatalogPath))) + "Asset"
	c.EnumInitName = strings.ToLower(c.EnumName[0:1]) + c.EnumName[1:]
}

// Write the Swift source code to a file.
func (c *Catalog) writeEnum() {
	tmpl, err := template.New("swift3").Parse(defaultSwift3Template)
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
