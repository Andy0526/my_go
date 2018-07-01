package eval

type Var string

type Env map[Var]float64

type literal float64

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args [] Expr
}
