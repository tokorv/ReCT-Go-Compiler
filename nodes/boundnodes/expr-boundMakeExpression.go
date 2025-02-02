package boundnodes

import (
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"fmt"
)

type BoundMakeExpressionNode struct {
	BoundExpressionNode

	BaseType  symbols.ClassSymbol
	Arguments []BoundExpressionNode
}

func (BoundMakeExpressionNode) NodeType() BoundType { return BoundMakeExpression }

func (node BoundMakeExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ BoundMakeExpressionNode")
	fmt.Println(indent + "  └ Type: ")
	node.BaseType.Print(indent + "    ")
	fmt.Println(indent + "  └ Arguments: ")
	for _, v := range node.Arguments {
		v.Print(indent + "   ")
	}
}

func (BoundMakeExpressionNode) IsPersistent() bool { return false }

// implement the expression node interface
func (node BoundMakeExpressionNode) Type() symbols.TypeSymbol {
	return node.BaseType.Type
}

func CreateBoundMakeExpressionNode(baseType symbols.ClassSymbol, args []BoundExpressionNode) BoundMakeExpressionNode {
	return BoundMakeExpressionNode{
		BaseType:  baseType,
		Arguments: args,
	}
}
