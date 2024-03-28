package manager

import "medioker-bank/usecase"

type UseCaseManager interface {
}

type useCaseManager struct {
	repo RepoManager
}

