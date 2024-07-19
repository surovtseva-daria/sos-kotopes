package post

import (
	"gitflic.ru/spbu-se/sos-kotopes/internal/core"
)
func (p *Post) ToCorePost() *core.Post {
	if p == nil {
		return nil
	}

	return &core.Post{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		AnimalID:  p.AnimalID,
	}
}