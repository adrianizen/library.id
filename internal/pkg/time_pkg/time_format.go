package time_pkg

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const TIME_ATOM = "2006-01-02T15:04:05-07:00"
const TIME_UTC = "2006-01-02T15:04:05Z"

const TIME_DB = "2006-01-02 15:04:05"

func TransformToUTCTime(timeString string) (time.Time, error) {
	var parsedTime time.Time
	parsedTime, err := time.Parse(TIME_ATOM, timeString)

	if err != nil {
		parsedTime, err = time.Parse(TIME_UTC, timeString)
		if err != nil {
			return parsedTime, errors.New("time format not valid")
		}

	}
	parsedTime, err = TimeIn(parsedTime, "UTC")
	if err != nil {
		return parsedTime, err
	}

	return parsedTime, nil
}

func TransformDBFormatToTime(timeString string) (time.Time, error) {
	var parsedTime time.Time
	parsedTime, err := time.Parse(TIME_DB, timeString)

	if err != nil {
		return parsedTime, errors.New("time format not valid")
	}

	return parsedTime, nil
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

type JsonWaresixTimeFormat time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonWaresixTimeFormat) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), "\"")
	parsedTime, err := time.Parse(TIME_ATOM, timeString)

	if err != nil {
		parsedTime, err = time.Parse(TIME_UTC, timeString)
		if err != nil {
			return errors.New("time format not valid")
		}

	}
	parsedTime, err = TimeIn(parsedTime, "UTC")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	*j = JsonWaresixTimeFormat(parsedTime)
	return nil
}

func (j JsonWaresixTimeFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}
