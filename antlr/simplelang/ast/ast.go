package ast

type Expr interface{}

type IntLiteral struct {
	Value int
}

type VarRef struct {
	Name string
}

type BinOp struct {
	Left  Expr
	Op    string
	Right Expr
}

type Assign struct {
	Name string
	Expr Expr
}

type Print struct {
	Expr Expr
}

type Stat interface{}
