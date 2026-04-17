package main

import (
	"fmt"
	"math"
	"net/http"
	"testing"
)

type Expr interface {
	Eval(env Env) float64
}

type Var string
type literal float64

type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

type Env map[Var]float64

func testVal(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 8888, "pi": math.Pi}, "3993"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		//Print expr only when it changes,
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := t.Parse(test.expr)
		if err != nil {
			t.Errorf("Parse(%q): %v", test.expr, err)
			continue
		}
		got := fmt.Sprintf("%.0f", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("Eval(%q, %v): got %s, want %s", test.expr, test.env, got, test.want)
		}
	}

}

func plot(w http.ResponseWriter, r *http.Request) {
	//some code in here;
	fmt.Fprintf(w, "some string in here")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/plot", plot)

	http.ListenAndServe("localhost:8888", mux)
}
