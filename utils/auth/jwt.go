// JWT标准中注册的声明 (建议但不强制使用) ：
//
// iss: jwt签发者
// sub: jwt所面向的用户
// aud: 接收jwt的一方
// exp: jwt的过期时间，这个过期时间必须要大于签发时间
// nbf: 定义在什么时间之前，该jwt都是不可用的.
// iat: jwt的签发时间
// jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
//
package auth

import (
	"github.com/cbwfree/micro-game/utils/dtype"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Jwt struct {
	SigningMethod jwt.SigningMethod
	SecretKey     []byte
	Expire        time.Duration
	Refresh       time.Duration
}

func (j *Jwt) Encrypt(data map[string]interface{}) (string, error) {
	t := time.Now()
	// JWT声明
	claims := jwt.MapClaims{"nbf": t.Unix(), "iat": t.Unix()}
	if j.Expire > 0 {
		claims["exp"] = t.Add(j.Expire).Unix()
	}
	// 写入数据
	for k, v := range data {
		claims[k] = v
	}

	token := jwt.NewWithClaims(j.SigningMethod, claims)
	ts, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}
	return ts, nil
}

func (j *Jwt) Verify(str string) (*JwtClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(str, &claims, func(t *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		c := &JwtClaims{claims: claims}
		// 检查是否需要刷新
		iat := c.GetInt64("iat")
		if iat > 0 && j.Refresh > 0 && time.Now().Add(-j.Refresh).Unix() > iat {
			c.refresh = true
		}
		return c, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func NewJwt(opts ...JwtOption) *Jwt {
	j := &Jwt{}

	for _, o := range opts {
		o(j)
	}

	return j
}

type JwtClaims struct {
	refresh bool // 是否需要刷新
	claims  map[string]interface{}
}

func (c *JwtClaims) HasRefresh() bool {
	return c.refresh
}

func (c *JwtClaims) Get(key string) interface{} {
	if val, ok := c.claims[key]; ok {
		return val
	}
	return nil
}

func (c *JwtClaims) GetStr(key string) string {
	val := c.Get(key)
	if val != nil {
		return dtype.ParseStr(val)
	}
	return ""
}

func (c *JwtClaims) GetInt(key string) int {
	val := c.Get(key)
	if val != nil {
		return dtype.ParseInt(val)
	}
	return 0
}

func (c *JwtClaims) GetInt64(key string) int64 {
	val := c.Get(key)
	if val != nil {
		return dtype.ParseInt64(val)
	}
	return 0
}

func (c *JwtClaims) GetFloat32(key string) float32 {
	val := c.Get(key)
	if val != nil {
		return dtype.ParseFloat32(val)
	}
	return 0
}

func (c *JwtClaims) GetFloat64(key string) float64 {
	val := c.Get(key)
	if val != nil {
		return dtype.ParseFloat64(val)
	}
	return 0
}
