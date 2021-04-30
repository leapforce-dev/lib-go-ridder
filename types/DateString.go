package ridder

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	layout string = "2006-01-02"
)

type DateString civil.Date

func (d *DateString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(layout, unquoted)
	if err != nil {
		return err
	}

	*d = DateString(civil.DateOf(_t))
	return nil
}

func (d *DateString) MarshalJSON() ([]byte, error) {
	if d == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(utilities.DateToTime(civil.Date(*d)).Format(layout))
}

func (d *DateString) ValuePtr() *civil.Date {
	if d == nil {
		return nil
	}

	_d := civil.Date(*d)
	return &_d
}

func (d DateString) Value() civil.Date {
	return civil.Date(d)
}
