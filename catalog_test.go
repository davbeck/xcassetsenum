package main

import "testing"
import "io/ioutil"
import "os"

func TestCatalog(t *testing.T) {
	_ = os.Remove("testdata/MediaAsset.swift")

	c := NewCatalog("testdata/Media.xcassets", "internal")
	c.writeEnum()

	outFile, err := ioutil.ReadFile(c.outputPath())
	if err != nil {
		t.Error("Could not read output file.")
	}

	expectedFile, err := ioutil.ReadFile("testdata/ExpectedMediaAsset.swift")
	if err != nil {
		t.Error("Could not read testdata/ExpectedMediaAsset.swift file.")
	}

	if string(outFile) != string(expectedFile) {
		t.Error("Output file does not match testdata/ExpectedMediaAsset.swift.")
	}

	_ = os.Remove("testdata/MediaAsset.swift")
}

func TestCatalogOutputPath(t *testing.T) {
	c := NewCatalog("testdata/Media.xcassets", "internal")

	if c.outputPath() != "testdata/MediaAsset.swift" {
		t.Error("Output path should be in the same directory as the asset catalog.")
	}
}
