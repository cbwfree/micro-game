package formula

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Top level function
// Analytical expression and execution
// err is not nil if an error occurs (including arithmetic runtime errors)
func ParseAndExec(s string) (r float64, err error) {
	toks, err := Parse(s)
	if err != nil {
		return 0, err
	}
	ast := NewAST(toks, s)
	if ast.Err != nil {
		return 0, ast.Err
	}
	ar := ast.ParseExpression()
	if ast.Err != nil {
		return 0, ast.Err
	}
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	return ExprASTResult(ar), err
}

func ErrPos(s string, pos int) string {
	r := strings.Repeat("-", len(s)) + "\n"
	s += "\n"
	for i := 0; i < pos; i++ {
		s += " "
	}
	s += "^\n"
	return r + s + r
}

// the integer power of a number
func Pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	r := calPow(x, n)
	if n < 0 {
		r = 1 / r
	}
	return r
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	r := calPow(x, n>>1) // move right 1 byte
	r *= r
	if n&1 == 1 {
		r *= x
	}
	return r
}

func expr2Radian(expr ExprAST) float64 {
	r := ExprASTResult(expr)
	if TrigonometricMode == AngleMode {
		r = r / 180 * math.Pi
	}
	return r
}

// Float64ToStr float64 -> string
func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// RegFunction is Top level function
// register a new function to use in expressions
func RegFunction(name string, argc int, fun func(...ExprAST) float64) error {
	if len(name) == 0 {
		return errors.New("RegFunction name is not empty.")
	}
	if argc < 1 {
		return errors.New("RegFunction argc is must has one arg at least.")
	}
	if _, ok := defFunc[name]; ok {
		return errors.New("RegFunction name is already exist.")
	}
	defFunc[name] = defS{argc, fun}
	return nil
}

// ExprASTResult is a Top level function
// AST traversal
// if an arithmetic runtime error occurs, a panic exception is thrown
func ExprASTResult(expr ExprAST) float64 {
	var l, r float64
	switch expr.(type) {
	case BinaryExprAST:
		ast := expr.(BinaryExprAST)
		l = ExprASTResult(ast.Lhs)
		r = ExprASTResult(ast.Rhs)
		switch ast.Op {
		case "+":
			lh, _ := new(big.Float).SetString(Float64ToStr(l))
			rh, _ := new(big.Float).SetString(Float64ToStr(r))
			f, _ := new(big.Float).Add(lh, rh).Float64()
			return f
		case "-":
			lh, _ := new(big.Float).SetString(Float64ToStr(l))
			rh, _ := new(big.Float).SetString(Float64ToStr(r))
			f, _ := new(big.Float).Sub(lh, rh).Float64()
			return f
		case "*":
			f, _ := new(big.Float).Mul(new(big.Float).SetFloat64(l), new(big.Float).SetFloat64(r)).Float64()
			return f
		case "/":
			if r == 0 {
				panic(errors.New(
					fmt.Sprintf("violation of arithmetic specification: a division by zero in ExprASTResult: [%g/%g]",
						l,
						r)))
			}
			f, _ := new(big.Float).Quo(new(big.Float).SetFloat64(l), new(big.Float).SetFloat64(r)).Float64()
			return f
		case "%":
			return float64(int(l) % int(r))
		case "^":
			return Pow(l, int(r))
		default:

		}
	case NumberExprAST:
		return expr.(NumberExprAST).Val
	case FunCallerExprAST:
		f := expr.(FunCallerExprAST)
		def := defFunc[f.Name]
		return def.fun(f.Arg...)
	}

	return 0.0
}
