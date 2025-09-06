package channel

import "go.wdy.de/nago/pkg/data"

func NewCreate(repo Repository) Create {
	return func(title, desc string) (ID, error) {
		c := Channel{
			ID:          data.RandIdent[ID](),
			Title:       title,
			Description: desc,
		}

		return c.ID, repo.Save(c)
	}
}
