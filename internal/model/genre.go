package model

import (
	"github.com/juicyluv/advanced-rest-server/internal/validator"
)

type Genre struct {
	Id    int64  `json:"id"`
	Genre string `json:"genre"`
}

// Validate checks whether track instance is valid. If it's not, it
// will fill up given validator Error map with appropriate error messages.
func (g *Genre) Validate(v *validator.Validator) {
	v.NotEmpty(g.Genre, "genre")
}
