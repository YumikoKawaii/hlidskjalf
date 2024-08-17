package processor

import (
	"elysium.com/applications/yggdrasil/pkg/utilities"
	"elysium.com/shared/types"
	"strconv"
	"time"
)

type Generator interface {
	Generate(subject string) types.Prototype
}

func NewGenerator() Generator {
	return &generatorImpl{}
}

type generatorImpl struct {
}

func (g *generatorImpl) Generate(subject string) types.Prototype {

	switch subject {
	case types.SyntheticStudent{}.GetSubject():
		return g.generateSyntheticStudent()
	}

	return nil
}

func (g *generatorImpl) generateSyntheticStudent() types.SyntheticStudent {

	actionTime := time.Now()

	return types.SyntheticStudent{
		Id:    strconv.Itoa(int(actionTime.Unix())),
		Name:  utilities.RandomString(utilities.RandomNumber(10, 20)),
		Age:   uint32(utilities.RandomNumber(20, 30)),
		Sex:   utilities.RandomSex(),
		Major: utilities.RandomString(utilities.RandomNumber(10, 20)),
		Level: utilities.RandomString(utilities.RandomNumber(5, 10)),
		GPA:   float32(utilities.RandomNumber(0, 4)),
		// temporary ignore
		Hobbies:    nil,
		Country:    utilities.RandomString(utilities.RandomNumber(10, 20)),
		Province:   utilities.RandomString(utilities.RandomNumber(10, 20)),
		ActionTime: actionTime,
		Subject:    types.SyntheticStudent{}.Subject,
	}
}
