package schemas

import "github.com/attapon-th/null"

type FooRequest struct {
	FooName  null.String     `json:"foo_name"`
	FooDate  null.DateString `json:"foo_date"`
	FooInt   null.Int        `json:"foo_int"`
	FooFloat null.Float      `json:"foo_float"`
	FooTime  null.Time       `json:"foo_time"`
}

type FooResponse struct {
}
