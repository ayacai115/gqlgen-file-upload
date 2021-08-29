package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/ayacai115/gqlgen-file-upload/graph/generated"
	"github.com/ayacai115/gqlgen-file-upload/graph/model"
)

func (r *mutationResolver) UploadCsv(ctx context.Context, input model.UploadCsvInput) (*model.UploadCsvPayload, error) {
	// ShiftJISで読み込むコンバータを用意
	//converter := transform.NewReader(input.CsvData.File, japanese.ShiftJIS.NewDecoder())
	// csvのreaderを作成
	csvReader := csv.NewReader(input.CsvData.File) // UTF-8なら converterの代わりにinput.CsvData.Fileを直で入れる

	// 1行目はヘッダなのでスキップ
	_, err := csvReader.Read()
	if err != nil {
		panic("failed to read first csv")
	}

	for index := 0; ; index++ {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("failed to read csv")
		}

		fmt.Println(fmt.Sprintf("%s: %s", row[0], row[1]))
	}

	return &model.UploadCsvPayload{
		Result: model.UploadResultSucceeded,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
