package types

import (
	"encoding/json"
	"golang.org/x/xerrors"
)

func Analyze(bytes []byte) (Prototype, error) {

	primo := Primordial{}
	if err := json.Unmarshal(bytes, &primo); err != nil {
		return nil, err
	}

	switch primo.Subject {
	case SyntheticStudent{}.GetSubject():
		student := SyntheticStudent{}
		if err := json.Unmarshal(bytes, &student); err != nil {
			return nil, err
		}
		return student, nil
	default:
		return nil, xerrors.Errorf("error unknown prototype: %s", primo.Subject)
	}
}
