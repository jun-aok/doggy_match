package model

import (
	"database/sql/driver"

	"github.com/shopspring/decimal"
)

type Position struct {
	Val decimal.Decimal
}

func NewPosition(value decimal.Decimal) Position {
	return Position{Val: value}
}

// json変換する際に呼ばれる
func (p Position) MarshalJSON() ([]byte, error) {
	str := p.Val.String()
	return []byte(str), nil
}

func (p *Position) GetValue() decimal.Decimal {
	return p.Val
}

// echoでformにbindする際に利用
func (p *Position) UnmarshalParam(src string) error {
	d, error := decimal.NewFromString(src)
	if error != nil {
		return error
	}
	*p = NewPosition(d)
	return error
}

// SQLの結果をbindする時に利用
func (p Position) Value() (driver.Value, error) {
	return p.Val.Value()
}

// SQLの結果をbindする時に利用
func (p *Position) Scan(value interface{}) error {
	return p.Val.Scan(value)
}
