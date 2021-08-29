package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ayacai115/gqlgen-file-upload/graph/generated"
	"github.com/ayacai115/gqlgen-file-upload/graph/model"
)

func (r *mutationResolver) UploadCsv(ctx context.Context, input model.UploadCsvInput) (*model.UploadCsvPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
