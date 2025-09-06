package video

import (
	"iter"
	"time"

	"github.com/worldiety/option"
	"go.wdy.de/nago/pkg/blob"
	"go.wdy.de/nago/pkg/data"
	"go.wdy.de/nago/presentation/core"
)

type ID string

type FindAll func() iter.Seq2[Video, error]

type FindAllIdent func() iter.Seq2[ID, error]

type FindByID func(id ID) (option.Opt[Video], error)

type Create func(file core.File) (ID, error)

type Update func(vid Video) error

type Delete func(id ID) error
type Video struct {
	ID          ID     `visible:"false"`
	BlobKey     string `visible:"false"`
	Name        string `visible:"false"`
	Size        int64  `visible:"false"`
	Title       string
	Description string    `lines:"4"`
	CreatedAt   time.Time `visible:"false"`
}

func (v Video) String() string {
	return v.Title
}

func (v Video) WithIdentity(id ID) Video {
	v.ID = id
	return v
}

func (v Video) Identity() ID {
	return v.ID
}

type Repository data.Repository[Video, ID]

type UseCases struct {
	FindAll      FindAll
	Create       Create
	FindAllIdent FindAllIdent
	FindByID     FindByID
	Delete       Delete
	Update       Update
}

func NewUseCases(repo Repository, videos blob.Store) UseCases {
	return UseCases{
		FindAll:      repo.All,
		FindAllIdent: repo.Identifiers,
		Create:       NewCreate(repo, videos),
		FindByID:     repo.FindByID,
		Update:       repo.Save,
		Delete:       repo.DeleteByID,
	}
}
