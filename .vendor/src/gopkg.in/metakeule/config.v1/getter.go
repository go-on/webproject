package config

import (
	"time"
)

type BoolGetter struct {
	opt *Option
	cfg *Config
}

func (b *BoolGetter) Get() bool {
	return b.cfg.GetBool(b.opt.Name)
}

func (b *BoolGetter) IsSet() bool {
	return b.cfg.IsSet(b.opt.Name)
}

type Int32Getter struct {
	opt *Option
	cfg *Config
}

func (b *Int32Getter) IsSet() bool {
	return b.cfg.IsSet(b.opt.Name)
}

func (b *Int32Getter) Get() int32 {
	return b.cfg.GetInt32(b.opt.Name)
}

type Float32Getter struct {
	opt *Option
	cfg *Config
}

func (b *Float32Getter) IsSet() bool {
	return b.cfg.IsSet(b.opt.Name)
}

func (b *Float32Getter) Get() float32 {
	return b.cfg.GetFloat32(b.opt.Name)
}

type StringGetter struct {
	opt *Option
	cfg *Config
}

func (b *StringGetter) IsSet() bool {
	return b.cfg.IsSet(b.opt.Name)
}

func (b *StringGetter) Get() string {
	return b.cfg.GetString(b.opt.Name)
}

type DateTimeGetter struct {
	opt *Option
	cfg *Config
}

func (b *DateTimeGetter) IsSet() bool {
	return b.cfg.IsSet(b.opt.Name)
}

func (b *DateTimeGetter) Get() time.Time {
	return b.cfg.GetTime(b.opt.Name)
}

type JSONGetter struct {
	opt *Option
	cfg *Config
}

func (b *JSONGetter) IsSet() bool {
	return b.cfg.IsSet(b.opt.Name)
}

func (b *JSONGetter) Get(val interface{}) error {
	return b.cfg.GetJSON(b.opt.Name, val)
}
