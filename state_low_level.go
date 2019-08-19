package main

import (
	"errors"
)

func (state *RenvestgyState) getDevByID(id string) (*Developer, error) {
	for _, d := range state.Developers {
		if d.ID == id {
			return d, nil
		}
	}

	return nil, errors.New("developer id not registered")
}
