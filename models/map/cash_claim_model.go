package models


import (
	"database/sql"
	"reflect"
	"time"
)

type NullString sql.NullString

func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

func (ns *NullString) marshalJSON() ([]byte, error){
	if !ns.Valid{
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(b []type) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// type time.Time mysql.time.Time

// func (nt *time.Time) Scan(value interface{}) error {
// 	var t mysql.time.Time
// 	if err := t.Scan(value); err != nil {
// 		return err
// 	}

// 	if reflect.TypeOf(value) == nil {
// 		*nt = time.Time{t.time.Time, false}
// 	} else {
// 		*nt = time.Time{t.time.Time, true}
// 	}

// 	return nil
// }

// func (nt *time.Time) marshalJSON() ([]byte, error) {
// 	if !nt.Valid {
// 		return []byte("Null"), nil
// 	}

// 	var := fmt.Sprintf("\"%s\"", nt.Time.Format(time.Stamp))
// 	return []byte(val), nil
// }

// func (nt *time.Time) UnmarshalJSON(b []byte) error {
// 	s := string(b)

// 	x, err := time.Parse(time.Stamp, s)
// 	if err != nil {
// 		nt.Valid = false
// 		return err
// 	}

// 	nt.Time = x
// 	nt.Valid = true
// 	return nil
// }

type NullInt64 sql.NullInt64

func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return nil
	}

	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
}

func (ni *NullInt64) marshalJSON() ([]byte, error) {
	if !ni.valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil )
	return err
}