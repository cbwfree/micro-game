package formula

import "testing"

func TestParseAndExecSimple(t *testing.T) {
	type U struct {
		Expr string
		R    float64
	}
	exprs := []U{
		{"1", 1},
		{"1+", 1},
		{"1+2", 3},
		{"-1+2", 1},
		{"-(1+2)", -3},
		{"-(1+2)*5", -15},
		{"-(1+2)*5/3", -5},
		{"1+(-(1+2)*5/3)", -4},
		{"3^4", 81},
		{"3^4.5", 81},
		{"3.5^4.5", 150.0625},
		{"8%2", 0},
		{"8%3", 2},
		{"8%3.5", 2},
		{"1e2", 100},
		{"1e+2", 100},
		{"1e-2", 0.01},
		{"1e-2+1e2", 100.01},
		{"1e-2+1e2*6/3", 200.01},
		{"(1e-2+1e2)*6/3", 200.02},
		{"(88*8)+(1+1+1+1)+(6/1.5)-(99%9*(2^4))", 712},
		{"1/3*3", 1},
		{"123_456_789", 123456789},
		{"123_456_789___", 123456789},
		{"pi", 3.141592653589793},
		{"abs(1)", 1},
		{"abs(-1)", 1},
		{"ceil(90.2)", 91},
		{"ceil(90.8)", 91},
		{"ceil(90.0)", 90},
		{"floor(90.2)", 90},
		{"floor(90.8)", 90},
		{"floor(90.0)", 90},
		{"round(90.0)", 90},
		{"round(90.4)", 90},
		{"round(90.5)", 91},
		{"round(90.9)", 91},
		{"sqrt(4)", 2},
		{"cbrt(27)", 3},
		{"sqrt(4) + cbrt(27)", 5},
		{"sqrt(2^2) + cbrt(3^3)", 5},
		{"127^2+5/2-sqrt(2^2) + cbrt(3^3)", 16132.5},
		{"max(2,3.5)", 3.5},
		{"max(2^3,3+abs(-1)*6)", 9},
		{"min(2,3.5)", 2},
		{"min(2^3,3+abs(-1)*6)", 8},
		{"max(2^3,3^2)", 9},
		{"min(2^3,3^2)", 8},
		{"noerr(1/0)", 0},
		{"noerr(1/(1-1))", 0},
		{"0.1+0.2", 0.3},
		{"0.3-0.1", 0.2},
	}
	for _, e := range exprs {
		r, _ := ParseAndExec(e.Expr)
		if r != e.R {
			t.Error(e, " ParseAndExec:", r)
		}
	}
}

func TestParseAndExecTrigonometric(t *testing.T) {
	type U struct {
		Expr       string
		RadianMode float64
		AngleMode  float64
	}
	exprs := []U{
		{"sin(pi/2)", 1, 0.027412133592044294},
		{"csc(pi/2)", 1, 36.48019577324057},
		{"cos(0)", 1, 1},
		{"sec(0)", 1, 1},
		{"tan(pi/4)", 1, 0.013708642534394057},
		{"cot(pi/4)", 1, 72.94668290394674},

		{"sin(90)", 0.893996663600558, 1},
		{"csc(90)", 1.1185724071637082, 1},
		{"cos(0)", 1, 1},
		{"sec(0)", 1, 1},
		{"tan(45)", 1.6197751905438615, 1},
		{"cot(45)", 0.6173696237835551, 1},
	}
	for _, e := range exprs {
		TrigonometricMode = RadianMode
		r, _ := ParseAndExec(e.Expr)
		if r != e.RadianMode {
			t.Error(e, " ParseAndExec RadianMode:", r)
		}
		TrigonometricMode = AngleMode
		r, _ = ParseAndExec(e.Expr)
		if r != e.AngleMode {
			t.Error(e, " ParseAndExec AngleMode:", r)
		}
	}
}

func TestRegFunction(t *testing.T) {
	funs := []struct {
		Name string
		Argc int
		Fun  func(expr ...ExprAST) float64
		Exp  string
		R    float64
	}{
		{
			"double",
			1,
			func(expr ...ExprAST) float64 {
				return ExprASTResult(expr[0]) * 2
			},
			"double(6)",
			12,
		},
		{
			"percentage50",
			1,
			func(expr ...ExprAST) float64 {
				return ExprASTResult(expr[0]) / 2
			},
			"percentage50(6)",
			3,
		},
	}
	for _, f := range funs {
		_ = RegFunction(f.Name, f.Argc, f.Fun)
		r, err := ParseAndExec(f.Exp)
		if r != f.R {
			t.Error(err, "RegFunction errors when register new function: ", f.Name)
		}
	}

}

func TestParseAndExecError(t *testing.T) {
	exprs := []string{
		"(",
		"((((((",
		"(1",
		"(1+",
		"1#1",
		"_123_456_789___",
		"1ee3+3",
		"sin()",
		"sin",
		"pi(",
		"sin(1, 50)",
		"max(1,)",
		"min(1,)",
		"min(1,3, 099)",
		"1/0",
		"99.9 / (2-1-1)",
		"(1+2)3",
		"1+1 111",
		"1+1 111+2",
		"1 3",
		"1 3-",
	}
	for _, e := range exprs {
		_, err := ParseAndExec(e)
		if err == nil {
			t.Error(e, " this is error expr!")
		}
	}
}
