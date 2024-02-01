package models

type Environment string

const (
	EnvironmentProduction  Environment = "production"
	EnvironmentDevelopment Environment = "development"
	EnvironmentTest        Environment = "test"
)

func (e Environment) IsProduction() bool {
	return e == EnvironmentProduction
}

func (e Environment) IsDevelopment() bool {
	return e == EnvironmentDevelopment
}

func (e Environment) IsTest() bool {
	return e == EnvironmentTest
}
