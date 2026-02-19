// Package timereplace provides the Replace() function equivalent to Python's
// date.replace.
package timereplace

import "time"

// Values specifies the time.Time fields and its values to be rewritten with
// Replace.
type Values struct {
	Year  *int
	Month *time.Month
	Day   *int
	Hour  *int
	Min   *int
	Sec   *int
	Nsec  *int
	Loc   *time.Location
}

// Replace returns a new time.Time with some of the values in the given
// time.Time argument overwritten. The value to be rewritten is an element of
// the argument Values that is not nil.
func Replace(orig time.Time, v Values) time.Time {
	if v.Year == nil {
		v.Year = new(orig.Year())
	}
	if v.Month == nil {
		v.Month = new(orig.Month())
	}
	if v.Day == nil {
		v.Day = new(orig.Day())
	}
	if v.Hour == nil {
		v.Hour = new(orig.Hour())
	}
	if v.Min == nil {
		v.Min = new(orig.Minute())
	}
	if v.Sec == nil {
		v.Sec = new(orig.Second())
	}
	if v.Nsec == nil {
		v.Nsec = new(orig.Nanosecond())
	}
	if v.Loc == nil {
		v.Loc = orig.Location()
	}
	return time.Date(*v.Year, *v.Month, *v.Day, *v.Hour, *v.Min, *v.Sec, *v.Nsec, v.Loc)
}
