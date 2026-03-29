package notes

import (
	"context"

	repo "github.com/Davidmuthee12/notes-rest-API/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListNotes(ctx context.Context) ([]repo.Note, error)
	CreateNote(ctx context.Context, content string) (repo.Note, error)
}

type svc struct {
	// Repository
	repo repo.Querier
}

// CreateNote implements [Service].
func (s *svc) CreateNote(ctx context.Context, content string) (repo.Note, error) {
	return s.repo.CreateNote(ctx, content)
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListNotes(ctx context.Context) ([]repo.Note, error) {
	return s.repo.ListNotes(ctx)
}
