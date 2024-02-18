package data

import (
	"errors"
	"strconv"
	"time"
)

var ErrInvalidDateFormat = errors.New("invalid date format")

type Date time.Time

// Implement json.Marshaler and json.Unmarshaler interfaces

// Because MarshalJSON() needs to return a byte slice and not
// modify the receiver, we can use a value receiver for this method.
func (d Date) MarshalJSON() ([]byte, error) {
	// format the time.Time value using the desired layout and
	// wrap it in double quotes before returning it
	return []byte(strconv.Quote(time.Time(d).Format("2006-01-02"))), nil
}

// IMPORTANT: Because UnmarshalJSON() needs to modify the
// receiver, we must use a pointer receiver for this to work correctly.
func (d *Date) UnmarshalJSON(jsonValue []byte) error {
	// try to unmarshal the data into a time.Time
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidDateFormat
	}

	t, err := time.Parse("2006-01-02", unquotedJSONValue)
	if err != nil {
		return ErrInvalidDateFormat
	}

	// then assign the time.Time value to the Date
	*d = Date(t)

	return nil
}
