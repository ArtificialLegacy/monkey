package evaluator

import (
	"github.com/ArtificialLegacy/monkey/pkg/ast"
	"github.com/ArtificialLegacy/monkey/pkg/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
