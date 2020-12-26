/*
formula

--------

数学表达式计算引擎

使用 Go 实现的数学表达式解析计算引擎，完成了数学表达式解析执行，包括词法分析、语法分析、构建AST、运行。

能够处理的表达式样例：

1+127-21+(3-4)*6/2.5

(88+(1+8)*6)/2+99

123_345_456 * 1.5 - 2 ^ 4

-4 * 6 + 2e2 - 1.6e-3

sin(pi/2)+cos(45-45*1)+tan(pi/4)

99+abs(-1)-ceil(88.8)+floor(88.8)

支持的运算符和函数:
symbol	explanation	e.g.
+	加，plus	1+2 = 3
-	减，sub	8-3.5 = 4.5
*	乘，multiply	2*3 = 6
/	除，division	5/2 = 2.5
%	取余，remainder	5%2 = 1
^	整数次方，integer power	2^3 = 8, 3^2 = 9
e	科学计数法，E-notation	1.2e3 = 1.2e+3 = 1200，1.2e-2 = 0.012
()	括号，brackets	(2+3)*4 = 20
_	数字分隔符，number separator	123_456_789 = 123456789
pi	π	pi = 3.141592653589793
sin(x)	正弦函数，sine	sin(pi/2) = 1
cos(x)	余弦函数，cosine	cos(0) = 1
tan(x)	正切函数，tangent	tan(pi/4) = 1
cot(x)	余切函数，cotangent	cot(pi/4) = 1
sec(x)	正割函数，secant	sec(0) = 1
csc(x)	余割函数，cosecant	csc(pi/2) = 1
abs(x)	绝对值，absolute value	abs(-6) = 6
ceil(x)	向上取整	ceil(4.2) = 5
floor(x)	向下取整	floor(4.8) = 4
round(x)	四舍五入取整	round(4.4) = 4, round(4.5) = 5
sqrt(x)	平方根，square root	sqrt(4) = 2
cbrt(x)	立方根，cube root	cbrt(27) = 3
max(x, y)	x, y 中的较大值	max(2, 3) = 3
min(x, y)	x, y 中的较小值	min(2, 3) = 2
noerr(x)	计算 x 出错时返回 0	noerr(1 / 1) = 1, noerr( 1/ 0 ) = 0
double(x)	返回 x 的双倍值，这是一个自定义的函数示例，你可以注册任意的自定义函数到引擎中	double(6) = 12



*/
package formula
