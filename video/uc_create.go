package video

import (
	"go.wdy.de/nago/pkg/blob"
	"go.wdy.de/nago/pkg/data"
	"go.wdy.de/nago/presentation/core"
)

func NewCreate(repo Repository, files blob.Store) Create {
	return func(file core.File) (ID, error) {
		blobKey := data.RandIdent[string]()
		reader, err := file.Open()
		if err != nil {
			return "", err
		}

		defer reader.Close()

		n, err := blob.Write(files, blobKey, reader)
		if err != nil {
			return "", err
		}

		vid := Video{
			ID:      data.RandIdent[ID](),
			BlobKey: blobKey,
			Name:    file.Name(),
			Size:    n,
			Title:   file.Name(),
		}

		return vid.ID, repo.Save(vid)
	}
}
