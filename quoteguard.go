package quoteguard

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "quoteguard is a static analysis tool for Go source code that helps developers choose the optimal quoting style for string literals."

// Analyzer is a static analysis tool that checks for proper string quoting.
var Analyzer = &analysis.Analyzer{
	Name: "quoteguard",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
		(*ast.BasicLit)(nil),
	}

	pos := make(map[token.Pos]bool)

	inspect.Preorder(nodeFilter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.Field:
			if n.Tag != nil {
				pos[n.Tag.Pos()] = true
			}
		case *ast.BasicLit:
			if pos[n.Pos()] {
				return
			}

			if n.Kind != token.STRING {
				return
			}

			if isDoubleQuote(n.Value) {
				content := n.Value[1 : len(n.Value)-1]
				unquoted, err := strconv.Unquote(n.Value)
				if content == unquoted {
					return
				}
				if hasEscape(unquoted) {
					return
				}
				if err == nil {
					pass.Reportf(n.Pos(), "enclosed with back quotes `%s`", unquoted)
				}
			}

			if isBackQuote(n.Value) {
				content := n.Value[1 : len(n.Value)-1]
				unquoted, err := strconv.Unquote(n.Value)
				if err != nil {
					return
				}
				quote := strconv.Quote(unquoted)
				if quote == `"`+content+`"` {
					pass.Reportf(n.Pos(), `enclosed with double quotes "%s"`, content)
				}
			}
		}
	})

	return nil, nil
}

func isDoubleQuote(s string) bool {
	return strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`)
}

func isBackQuote(s string) bool {
	return strings.HasPrefix(s, "`") && strings.HasSuffix(s, "`")
}

func hasEscape(s string) bool {
	return strings.Contains(s, "\n") ||
		strings.Contains(s, "\t") ||
		strings.Contains(s, "\r") ||
		strings.Contains(s, "\b") ||
		strings.Contains(s, "\f") ||
		strings.Contains(s, "\v") ||
		strings.Contains(s, "\a")
}
