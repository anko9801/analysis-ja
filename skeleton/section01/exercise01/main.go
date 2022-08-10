package main

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strconv"

	// "testing/fstest"
	// "github.com/klauspost/compress/fse"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("source code must be specified")
	}

	fname := args[0]
	fset := token.NewFileSet()
	// fnameのファイルをパースする
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)

	if err != nil {
		return err
	}

	for _, spec := range f.Imports {
		// パスの文字列リテラルから文字列を取得
		path, err := strconv.Unquote(spec.Path.Value)

		if err != nil {
			return err
		}

		if path == "unsafe" {
			// token.Pos型からtoken.Position型の値を取得
			pos := fset.Position(spec.Path.Pos())

			fmt.Fprintf(os.Stderr, "%s: import unsafe\n", pos)
		}
	}

	return nil
}
