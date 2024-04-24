// Package ast holds the AST types
package ast

// Node the base for all AST node types
type Node interface {
	TokenLiteral() string
	IsNil() bool
	String() string
}

// Root stylesheet root
type Root interface {
	Node
	rootNode()
}

// Comment comment node
type Comment interface {
	Node
	commentNode()
}

// Rule a CSS rule node
type Rule interface {
	Node
	ruleNode()
}

// AtRule a CSS media query rule node
type AtRule interface {
	Node
	atRuleNode()
}

// Declaration a CSS key value pair
type Declaration interface {
	Node
	declNode()
}
