package models

import (
	"database/sql"
	"encoding/json"
	"reflect"
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

func (ns *NullString) marshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// func (ns *NullString) UnmarshalJSON(b []type) error {
// 	err := json.Unmarshal(b, &ns.String)
// 	ns.Valid = (err == nil)
// 	return err
// }
