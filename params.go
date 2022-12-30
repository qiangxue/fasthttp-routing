package routing

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

type URLParams struct {
	c *fasthttp.RequestCtx
}

func (p URLParams) get(key string) string {
	return string(p.c.QueryArgs().Peek(key))
}

func (p URLParams) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p URLParams) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil{
		return def
	}
	return value
}

func (p URLParams) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil{
		return 0, err
	}
	return int32(value), nil
}

func (p URLParams) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil{
		return def
	}
	return int32(value)
}

func (p URLParams) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil{
		return 0, err
	}
	return int32(value), nil
}

func (p URLParams) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil{
		return def
	}
	return value
}

func (p URLParams) Value(key string) string {
	return p.get(key)
}

func (p URLParams) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p URLParams) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p URLParams) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil{
		return def
	}
	return value
}
