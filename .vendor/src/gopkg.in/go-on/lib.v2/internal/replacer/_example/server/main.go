package main

import (
	"net/http"
	"time"

	"github.com/go-on/replacer"
	"gopkg.in/metakeule/fmtdate.v1"
)

var (
	hour   = replacer.Placeholder("hour")
	minute = replacer.Placeholder("minute")
	second = replacer.Placeholder("second")
	day    = replacer.Placeholder("day")
	month  = replacer.Placeholder("month")
	year   = replacer.Placeholder("year")

	currentTime = "Hour: " + hour.String() + " Minute: " + minute.String() + " Second: " + second.String()
	currentDate = "Day: " + day.String() + " Month: " + month.String() + " Year: " + year.String()
	dateAndTime = currentDate + " " + currentTime
)

type datetime struct{ *replacer.Template }

func (d *datetime) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	m := map[replacer.Placeholder]string{
		hour:   fmtdate.Format("hh", t),
		minute: fmtdate.Format("mm", t),
		second: fmtdate.Format("ss", t),
		day:    fmtdate.Format("DDDD", t),
		month:  fmtdate.Format("MMMM", t),
		year:   fmtdate.Format("YYYY", t),
	}
	w.Write(d.ReplaceStrings(m))
}

func main() {
	http.Handle("/time", &datetime{replacer.NewTemplateString(currentTime)})
	http.Handle("/date", &datetime{replacer.NewTemplateString(currentDate)})
	http.Handle("/datetime", &datetime{replacer.NewTemplateString(dateAndTime)})

	http.ListenAndServe(":7979", nil)
}
