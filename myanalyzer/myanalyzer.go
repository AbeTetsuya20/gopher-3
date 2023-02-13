package myanalyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "myanalyzer checks the number of functions and outputs a warning if there are more than 5 functions."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "myanalyzer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			// 引数があるかどうかチェック
			if len(n.Type.Params.List) != 0 {
				// 引数の数をチェック
				params := len(n.Type.Params.List[0].Names)
				if params > 5 {
					// 引数が5以上だったら警告を出す
					pass.Reportf(n.Pos(), "too many arguments: func name: %s", n.Name.Name)
				}
			}
		}
	})

	return nil, nil
}
