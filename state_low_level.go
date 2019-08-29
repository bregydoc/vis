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

func (state *RenvestgyState) getRGYXByID(id string) (*RGYx, error) {
	for _, r := range state.RGYs {
		if r.ID == id {
			return r, nil
		}
	}

	return nil, errors.New("rgyx id not found")
}

func (state *RenvestgyState) getInvestorByID(id string) (*Investor, error) {
	for _, i := range state.Investors {
		if i.ID == id {
			return i, nil
		}
	}
	return nil, errors.New("investor id not found")
}
