package core

//go:generate mockgen -package core -self_package github.com/nigoroll/mock/mockgen/internal/tests/self_package -destination mock.go github.com/nigoroll/mock/mockgen/internal/tests/self_package Methods

type Info struct{}

type Methods interface {
	getInfo() Info
}
