package types

import (
	"encoding/json"

	"gorm.io/gorm"
)

type StringListInSQL []string

func (slis *StringListInSQL) Scan(value interface{}) error {
	if value == nil {
		*slis = nil
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return gorm.ErrInvalidValue
	}

	return json.Unmarshal(data, slis)
}

func (slis *StringListInSQL) Value() (interface{}, error) {
	return json.Marshal(slis)
}
