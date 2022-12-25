package models

type JwtDecodedModel struct {
	Data Data  `json:"data"`
	Exp  int64 `json:"exp"`
}

type Data struct {
	CreatedAt string      `json:"CreatedAt"`
	DeletedAt interface{} `json:"DeletedAt"`
	ID        int64       `json:"ID"`
	UpdatedAt string      `json:"UpdatedAt"`
	Role      string      `json:"role"`
	Category  interface{} `json:"categorCategory"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Task      interface{} `json:"task"`
	Full_name string      `json:"full_name"`
}
