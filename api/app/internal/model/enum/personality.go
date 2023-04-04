package model

import (
	"fmt"
	"strconv"
)

type Personality int

const (
	PersonalityEmpty = 0 + iota
	Sociable
	Mild
)

func ItoPersonality(i int) (Personality, error) {
	m := map[int]Personality{
		1: Sociable,
		2: Mild,
	}
	if m[i] == 0 {
		return PersonalityEmpty, fmt.Errorf("Personalityではありません")
	}
	return m[i], nil
}

func (p *Personality) UnmarshalParam(src string) error {
	i, error := strconv.Atoi(src)
	if error != nil {
		return error
	}
	personality, error := ItoPersonality(i)
	if error != nil {
		return error
	}
	*p = personality
	return error
}
