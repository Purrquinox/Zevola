package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func isExported(name string) bool {
	return ast.IsExported(name)
}

func isGenericFunc(fn *ast.FuncDecl) bool {
	return fn.Type.TypeParams != nil && len(fn.Type.TypeParams.List) > 0
}

func main() {
	coreDir := "./core"
	outFile := "zevola.go"

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, coreDir, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	var exports []string
	var skipped []string

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				switch d := decl.(type) {
				case *ast.FuncDecl:
					if d.Recv == nil && isExported(d.Name.Name) {
						if isGenericFunc(d) {
							skipped = append(skipped, d.Name.Name)
							continue
						}
						exports = append(exports, fmt.Sprintf("var %s = core.%s", d.Name.Name, d.Name.Name))
					}
				case *ast.GenDecl:
					for _, spec := range d.Specs {
						switch s := spec.(type) {
						case *ast.TypeSpec:
							if isExported(s.Name.Name) {
								exports = append(exports, fmt.Sprintf("type %s = core.%s", s.Name.Name, s.Name.Name))
							}
						case *ast.ValueSpec:
							for _, name := range s.Names {
								if isExported(name.Name) {
									exports = append(exports, fmt.Sprintf("var %s = core.%s", name.Name, name.Name))
								}
							}
						}
					}
				}
			}
		}
	}

	var sb strings.Builder
	sb.WriteString("package main\n\n")
	sb.WriteString(`import "github.com/purrquinox/zevola/core"` + "\n\n")

	for _, line := range exports {
		sb.WriteString(line + "\n")
	}

	err = os.WriteFile(outFile, []byte(sb.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("✅ Generated %s with %d exports\n", outFile, len(exports))
	if len(skipped) > 0 {
		fmt.Println("⚠️  Skipped generic functions:")
		for _, name := range skipped {
			fmt.Printf(" - %s\n", name)
		}
	}
}
