package model

type Chef struct {
	ID   int64 `json:"-"`
	Name string `json:"name"`
}

type Menu struct {
	ID   int64 `json:"id"`
	Name string `json:"name"`
}
