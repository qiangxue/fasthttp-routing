package routing

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

type Header struct {
	c *fasthttp.RequestCtx
}

func (p Header) get(key string) string {
	return string(p.c.Request.Header.Peek(key))
}

func (p Header) Set(key string, value string) {
	p.c.Request.Header.Set(key, value)
}

func (p Header) SetInt(key string, value int) {
	vStr := strconv.Itoa(value)
	p.c.Request.Header.Set(key, vStr)
}

func (p Header) SetInt32(key string, value int32) {
	vStr := strconv.FormatInt(int64(value), 10)
	p.c.Request.Header.Set(key, vStr)
}

func (p Header) SetInt64(key string, value int64) {
	vStr := strconv.FormatInt(value, 10)
	p.c.Request.Header.Set(key, vStr)
}

func (p Header) SetBool(key string, value bool) {
	vStr := strconv.FormatBool(value)
	p.c.Request.Header.Set(key, vStr)
}

func (p Header) Int(key string) (int, error) {
	vStr := p.get(key)
	return strconv.Atoi(vStr)
}

func (p Header) IntDefault(key string, def int) int {
	vStr := p.get(key)
	value, err := strconv.Atoi(vStr)
	if err != nil{
		return def
	}
	return value
}

func (p Header) Int32(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil{
		return 0, err
	}
	return int32(value), nil
}

func (p Header) Int32Default(key string, def int32) int32 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil{
		return def
	}
	return int32(value)
}

func (p Header) Int64(key string) (int32, error) {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil{
		return 0, err
	}
	return int32(value), nil
}

func (p Header) Int64Default(key string, def int64) int64 {
	vStr := p.get(key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil{
		return def
	}
	return value
}

func (p Header) Value(key string) string {
	return p.get(key)
}

func (p Header) ValueDefault(key string, def string) string {
	if p.get(key) == "" {
		return def
	}
	return p.get(key)
}

func (p Header) Bool(key string) (bool, error) {
	vStr := p.get(key)
	return strconv.ParseBool(vStr)
}

func (p Header) BoolDefault(key string, def bool) bool {
	vStr := p.get(key)
	value, err := strconv.ParseBool(vStr)
	if err != nil{
		return def
	}
	return value
}
