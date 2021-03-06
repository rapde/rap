package website

import (
	"archive/zip"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
)

//go:generate go run generator/makestatic.go

const warning = `// Code generated by "makestatic"; DO NOT EDIT.`

const license = `// Copyright Rapde/Rap. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause License
// license that can be found in the LICENSE file.`

// Generate reads a set of files and returns a file buffer that declares
// a map of string constants containing contents of the input files.
func Generate() ([]byte, error) {

	buf := new(bytes.Buffer)

	assets, err := ZipWriter("build")
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(buf, "%v\n\n%v\n\npackage website\n\n", license, warning)
	fmt.Fprintf(buf, "func init(){\n\n")
	fmt.Fprintf(buf, "RAR = []byte{\n")

	for i, b := range assets {

		buf.WriteString(fmt.Sprintf("%d,", b))

		if (i+1)%30 == 0 {
			buf.WriteString("\n")
		}
	}

	fmt.Fprintln(buf, "\n}\n\n}")

	return format.Source(buf.Bytes())
}

// ZipWriter write website zip
func ZipWriter(root string) ([]byte, error) {

	buf := new(bytes.Buffer)
	compresser := zip.NewWriter(buf)

	err := filepath.Walk(root, func(path string, fileInfo os.FileInfo, err error) error {

		if err != nil || fileInfo.IsDir() {
			return err
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		f, err := compresser.Create(path)
		if err != nil {
			return err
		}

		_, err = f.Write(data)
		return err
	})

	if err != nil {
		return nil, err
	}

	if err = compresser.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
