package models

import (
	"math/rand"
)

type HumanCode struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	TimeTrackedModel
}

func (m *HumanCode) GenerateCode() {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	m.Code = ""
	for i := 0; i < 5; i++ {
		m.Code += string(chars[rand.Intn(len(chars))])
	}
}
