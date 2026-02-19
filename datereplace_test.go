package timereplace_test

import (
	"testing"
	"time"

	"github.com/koron-go/timereplace"
)

func testReplace(t *testing.T, origTime time.Time, v timereplace.Values, want string) {
	gotTime := timereplace.Replace(origTime, v)
	got := gotTime.Format(time.RFC3339Nano)
	if got != want {
		t.Helper()
		t.Errorf("replace failure:\nwant=%s\n got=%s", want, got)
	}
}

func TestReplace(t *testing.T) {
	baseTime := time.Date(2006, 1, 2, 3, 4, 5, 999999999, time.UTC)

	testReplace(t, baseTime,
		timereplace.Values{},
		"2006-01-02T03:04:05.999999999Z")

	testReplace(t, baseTime,
		timereplace.Values{
			Year: new(2026),
		},
		"2026-01-02T03:04:05.999999999Z")
	testReplace(t, baseTime,
		timereplace.Values{
			Month: new(time.Month(2)),
		},
		"2006-02-02T03:04:05.999999999Z")
	testReplace(t, baseTime,
		timereplace.Values{
			Day: new(19),
		},
		"2006-01-19T03:04:05.999999999Z")

	testReplace(t, baseTime,
		timereplace.Values{
			Hour: new(15),
		},
		"2006-01-02T15:04:05.999999999Z")
	testReplace(t, baseTime,
		timereplace.Values{
			Min: new(17),
		},
		"2006-01-02T03:17:05.999999999Z")
	testReplace(t, baseTime,
		timereplace.Values{
			Sec: new(30),
		},
		"2006-01-02T03:04:30.999999999Z")

	testReplace(t, baseTime,
		timereplace.Values{
			Nsec: new(123456789),
		},
		"2006-01-02T03:04:05.123456789Z")

	locTokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("not found Asia/Tokyo location: %s", err)
	}
	testReplace(t, baseTime,
		timereplace.Values{
			Loc: locTokyo,
		},
		"2006-01-02T03:04:05.999999999+09:00")
}
