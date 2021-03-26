package main

// Are there any maintenance windows available?
// If no windows are available return false.
import (
	"fmt"
	"time"
)

// Wtime is WindowTime
type Wtime struct {
	Hour   int
	Minute int
}

// Window is maintenance window with day of week, start hour/minute and end hour/minute.
type Window struct {
	Day   time.Weekday
	Start Wtime
	End   Wtime
}

// Windows is a collection of window.
type Windows struct {
	List []Window
	Time time.Time
}

// String implements the stringer interface.
func (w Windows) String() string {
	return fmt.Sprint("maintenance windows available? ", check(w.Time, w.List))
}

func main() {
	w := Windows{
		List: []Window{
			{time.Tuesday, Wtime{9, 30}, Wtime{9, 50}},
			{time.Wednesday, Wtime{12, 30}, Wtime{12, 50}},
			{time.Wednesday, Wtime{9, 30}, Wtime{9, 50}},
			{time.Friday, Wtime{13, 30}, Wtime{19, 50}},
			{time.Saturday, Wtime{15, 30}, Wtime{15, 50}},
		},
		Time: time.Now().AddDate(0, 0, -20),
	}

	fmt.Println(w)
}

func check(proposed time.Time, windows []Window) bool {
	if windows == nil || len(windows) == 0 {
		return false
	}
	d := proposed.Weekday()
	for _, v := range windows {
		if d != v.Day {
			continue
		}
		start, end := createDates(proposed, &v)
		if (proposed.After(start) || proposed.Equal(start)) &&
			(proposed.Before(end) || proposed.Equal(end)) {
			return true
		}
	}
	return false
}

func createDates(t time.Time, v *Window) (start, end time.Time) {
	year, month, day := t.Date()

	start = time.Date(year, month, day, v.Start.Hour, v.Start.Minute, 0, 0, t.Location())
	start = start.AddDate(0, 0, int(v.Day-t.Weekday()))

	end = time.Date(year, month, day, v.End.Hour, v.End.Minute, 0, 0, t.Location())
	end = end.AddDate(0, 0, int(v.Day-t.Weekday()))
	return
}
