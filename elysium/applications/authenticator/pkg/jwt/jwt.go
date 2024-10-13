package jwt

import (
	jwtPkg "github.com/golang-jwt/jwt"
	"golang.org/x/xerrors"
	"time"
)

type Config struct {
	SecretKey          string `env:"JWT_SECRET_KEY"`
	TokenDurationInSec int64  `env:"TOKEN_DURATION_IN_SEC" default:"108000"` // 30 days
}

type Info struct {
	Id   string
	User string
}

type Claim struct {
	jwtPkg.StandardClaims
	info Info
}

func (c *Claim) GetInfo() Info {
	return c.info
}

type Resolver interface {
	GenerateToken(info Info) (string, error)
	VerifyToken(token string) (*Claim, error)
}

type resolverImpl struct {
	cfg *Config
}

func NewResolver(cfg *Config) Resolver {
	return &resolverImpl{cfg: cfg}
}

func (r *resolverImpl) GenerateToken(info Info) (string, error) {

	claim := Claim{
		StandardClaims: jwtPkg.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(r.cfg.TokenDurationInSec) * time.Second).UnixNano(),
		},
		info: info,
	}

	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claim)
	return token.SignedString([]byte(r.cfg.SecretKey))
}

func (r *resolverImpl) VerifyToken(token string) (*Claim, error) {
	data, err := jwtPkg.ParseWithClaims(
		token, &Claim{}, func(token *jwtPkg.Token) (interface{}, error) {
			_, ok := token.Method.(*jwtPkg.SigningMethodHMAC)
			if !ok {
				return nil, xerrors.Errorf("unexpected token signing method")
			}
			return []byte(r.cfg.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claim, ok := data.Claims.(*Claim)
	if !ok {
		return nil, xerrors.Errorf("invalid token")
	}

	return claim, nil
}
