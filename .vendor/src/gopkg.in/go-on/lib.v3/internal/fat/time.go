package fat

import (
	"fmt"
	//"gopkg.in/metakeule/fmtdate.v1"
	"time"
)

var (
	zeroTime, _ = time.Parse(time.UnixDate, time.UnixDate)
	//TimeFormat  = "YYYY-MM-DD hh:mm:ss"
)

type time_ time.Time

func Time(t time.Time) *time_    { t_ := time_(t); return &t_ }
func (t time_) Typ() string      { return "time" }
func (t time_) Get() interface{} { return time.Time(t) }

//func (t time_) String() string   { return fmt.Sprintf(fmtdate.Format(TimeFormat, time.Time(t))) }
func (t time_) String() string { return time.Time(t).Format(time.RFC3339) }

func (øtime *time_) MarshalJSON() ([]byte, error) {
	return []byte(`"` + øtime.String() + `"`), nil
}

func (øtime *time_) Set(i interface{}) error {
	t, ok := i.(time.Time)
	if !ok {
		return fmt.Errorf("can't convert %T to time", i)
	}
	*øtime = time_(t)
	return nil
}

func (øtime *time_) Scan(s string) error {
	t, err := time.Parse(time.RFC3339, s)
	// t, err := fmtdate.Parse(TimeFormat, s)
	if err != nil {
		return err
	}
	*øtime = time_(t)
	return nil
}
