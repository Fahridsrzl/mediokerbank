package manager

import "medioker-bank/repository"

type RepoManager interface {
}

type repoManager struct {
	infra InfraManager
}
