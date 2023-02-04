package myanalyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "myanalyzer is ..."

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
	fmt.Println("myanalyzer is running...")
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	//nodeFilter := []ast.Node{
	//	(*ast.Ident)(nil),
	//}

	var nodeFilter []ast.Node

	inspect.Preorder(nodeFilter, func(n ast.Node) {

		switch n := n.(type) {
		//case *ast.File:
		//	fmt.Printf("file: %s \n", n.Name.Name)
		//
		//case *ast.Ident:
		//	// 引数の数をチェック
		//	fmt.Printf("name: %s \n", n.Name)
		//	fmt.Printf("obj: %+v \n", n.Obj)
		//

		case *ast.FuncDecl:
			fmt.Printf("func: %s \n", n.Name.Name)

			// 引数があるかどうかチェック
			if len(n.Type.Params.List) != 0 {
				fmt.Printf("funcParams len: %d \n", len(n.Type.Params.List[0].Names))

				// 引数の数をチェック
				params := len(n.Type.Params.List[0].Names)
				if params > 3 {
					// 3以上だったら警告を出す
					pass.Reportf(n.Pos(), "too many arguments: func name: %s", n.Name.Name)
				}
			}

			// 関数の行数をチェック
			if n.Body != nil {
				fmt.Printf("funcBody len: %d \n", len(n.Body.List))
			}
		}
	})

	return nil, nil
}
