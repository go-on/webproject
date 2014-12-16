package typeconverter

import (
	"time"
)

type DefaultType int

func Default(b int) DefaultType { return DefaultType(b) }

func (ø DefaultType) String() string              { return "" }
func (ø DefaultType) Int() int                    { return 0 }
func (ø DefaultType) Float() float64              { return float64(0) }
func (ø DefaultType) Time() time.Time             { return time.Time{} }
func (ø DefaultType) Json() string                { return "{}" }
func (ø DefaultType) Array() []interface{}        { return []interface{}{} }
func (ø DefaultType) Bool() bool                  { return false }
func (ø DefaultType) Map() map[string]interface{} { return map[string]interface{}{} }
