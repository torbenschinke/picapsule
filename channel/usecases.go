package channel

import (
	"iter"

	"github.com/torbenschinke/picapsule/video"
	"github.com/worldiety/option"
	"go.wdy.de/nago/pkg/data"
)

type ID string
type Create func(title, desc string) (ID, error)

type Delete func(id ID) error
type Update func(c Channel) error

type FindByID func(id ID) (option.Opt[Channel], error)
type FindAll func() iter.Seq2[ID, error]

type Channel struct {
	ID          ID `visible:"false"`
	Title       string
	Description string
	Videos      []video.ID `source:"pic.videos"`
}

func (c Channel) Identity() ID {
	return c.ID
}

type Repository data.Repository[Channel, ID]

type UseCases struct {
	Create   Create
	Delete   Delete
	Update   Update
	FindByID FindByID
	FindAll  FindAll
}

func NewUseCases(repo Repository) UseCases {
	return UseCases{
		Create:   NewCreate(repo),
		Delete:   repo.DeleteByID,
		Update:   repo.Save,
		FindByID: repo.FindByID,
		FindAll:  repo.Identifiers,
	}
}
