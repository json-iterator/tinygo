package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/json-iterator/tinygo/generator"
)

var cwd string
var fset = token.NewFileSet()
var allImports = map[string]string{}

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		reportError(err)
		return
	}
}

func main() {
	typeSpec := locateTypeSpec()
	outputPath := filepath.Join(cwd, fmt.Sprintf("%s_json.go", typeSpec.Name.Name))
	os.WriteFile(outputPath, generator.Generate(fset, allImports, typeSpec), 0644)
}

func reportError(err error) {
	panic(err)
}

func locateTypeSpec() *ast.TypeSpec {
	goline, err := strconv.Atoi(os.Getenv("GOLINE"))
	if err != nil {
		reportError(err)
		return nil
	}
	f, err := parser.ParseFile(fset, os.Getenv("GOFILE"), nil, parser.ParseComments)
	if err != nil {
		reportError(err)
		return nil
	}
	for _, importSpec := range f.Imports {
		path, _ := strconv.Unquote(importSpec.Path.Value)
		if importSpec.Name == nil {
			parts := strings.Split(path, "/")
			allImports[parts[len(parts)-1]] = path
		} else {
			allImports[importSpec.Name.Name] = path
		}
	}
	var located *ast.TypeSpec
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if goline+1 == fset.Position(x.Pos()).Line {
				located = x
			}
			return false
		}
		return true
	})
	if located == nil {
		reportError(fmt.Errorf("%s:%s go generate should be marked just before type definition",
			os.Getenv("GOFILE"), os.Getenv("GOLINE")))
		return nil
	}
	return located
}
