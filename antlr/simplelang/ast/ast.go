package ast

type Expression interface{}

type IntLiteral struct {
	Value int
}

type VarRef struct {
	Name string
}

type BinOp struct {
	Left  Expression
	Op    string
	Right Expression
}

type Assign struct {
	Name string
	Expr Expression
}

type Print struct {
	Expr Expression
}

type Statement interface{}
