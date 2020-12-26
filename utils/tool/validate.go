package tool

import (
	"regexp"
	"sync"
)

var (
	validateInstance *validate
	validateOnce     sync.Once
)

func Valid() *validate {
	validateOnce.Do(func() {
		validateInstance = &validate{}
	})
	return validateInstance
}

var (
	validateRules = map[string]string{
		"alpha":            `^[A-Za-z]+$`,
		"alphaNum":         `^[A-Za-z0-9]+$`,
		"alphaDash":        `^[A-Za-z0-9\-\_]+$`,
		"chs":              `^[\x{4e00}-\x{9fa5}]+$`,
		"chsAlpha":         `^[\x{4e00}-\x{9fa5}a-zA-Z]+$`,
		"chsAlphaNum":      `^[\x{4e00}-\x{9fa5}a-zA-Z0-9]+$`,
		"chsDash":          `^[\x{4e00}-\x{9fa5}a-zA-Z0-9\_\-]+$`,
		"mobile":           `^1[3-9]\d{9}$`,
		"idCard":           `(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`,
		"zip":              `\d{6}`,
		"email":            `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`,
		"number":           `^\d*$`,
		"integer":          `^(\-|\+)?\d+$`,
		"float":            `^(-?\d+)(\.\d+)?$`,
		"positive_integer": `^[1-9]\d*$`,
		"negative_integer": `^-[1-9]\d*$`,
	}
)

type validate struct {
	rules []string
}

// 检查是否为大小写字母
func (v *validate) Alpha(val string) bool {
	return regexp.MustCompile(validateRules["alpha"]).MatchString(val)
}

// 检查是否为大小写字母和数字
func (v *validate) AlphaNum(val string) bool {
	return regexp.MustCompile(validateRules["alphaNum"]).MatchString(val)
}

// 检查是否为大小写字母和数字及_-
func (v *validate) AlphaDash(val string) bool {
	return regexp.MustCompile(validateRules["alphaDash"]).MatchString(val)
}

// 检查是否为中文
func (v *validate) Chs(val string) bool {
	return regexp.MustCompile(validateRules["chs"]).MatchString(val)
}

// 检查是否为中文、字母
func (v *validate) ChsAlpha(val string) bool {
	return regexp.MustCompile(validateRules["chsAlpha"]).MatchString(val)
}

// 检查是否为中文、字母、数字
func (v *validate) ChsAlphaNum(val string) bool {
	return regexp.MustCompile(validateRules["chsAlphaNum"]).MatchString(val)
}

// 检查是否为中文、字母、数字、_、-
func (v *validate) ChsDash(val string) bool {
	return regexp.MustCompile(validateRules["chsDash"]).MatchString(val)
}

// 检查手机号格式是否正确
func (v *validate) Mobile(val string) bool {
	return regexp.MustCompile(validateRules["mobile"]).MatchString(val)
}

// 检查身份证号码是否正确
func (v *validate) IdCard(val string) bool {
	return regexp.MustCompile(validateRules["idcard"]).MatchString(val)
}

// 检查邮编是否正确
func (v *validate) Zip(val string) bool {
	return regexp.MustCompile(validateRules["zip"]).MatchString(val)
}

// 检查邮箱地址是否正确
func (v *validate) Email(val string) bool {
	return regexp.MustCompile(validateRules["email"]).MatchString(val)
}

// 检查是否为数字
func (v *validate) Number(val string) bool {
	return regexp.MustCompile(validateRules["number"]).MatchString(val)
}

// 整数
func (v *validate) Integer(val string) bool {
	return regexp.MustCompile(validateRules["integer"]).MatchString(val)
}

// 小数
func (v *validate) Float(val string) bool {
	return regexp.MustCompile(validateRules["float"]).MatchString(val)
}

// 非零正整数
func (v *validate) PositiveInteger(val string) bool {
	return regexp.MustCompile(validateRules["positive_integer"]).MatchString(val)
}

// 非零负整数
func (v *validate) NegativeInteger(val string) bool {
	return regexp.MustCompile(validateRules["negative_integer"]).MatchString(val)
}
