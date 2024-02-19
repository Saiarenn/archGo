package group

import (
	"time"

	"architecture_go/services/contact/internal/domain/group/name"
	"github.com/google/uuid"
)

type Group struct {
	id           uuid.UUID
	createdAt    time.Time
	modifiedAt   time.Time
	name         name.Name
	contactCount uint64
}

func NewWithID(id uuid.UUID, name name.Name, contactCount uint64) *Group {
	return &Group{
		id:           id,
		name:         name,
		contactCount: contactCount,
	}
}

func New(name name.Name) *Group {
	return &Group{
		id:   uuid.New(),
		name: name,
	}
}

func (g Group) ContactCount() uint64 {
	return g.contactCount
}

func (g Group) ID() uuid.UUID {
	return g.id
}

func (g Group) CreatedAt() time.Time {
	return g.createdAt
}

func (g Group) ModifiedAt() time.Time {
	return g.modifiedAt
}

func (g Group) Name() name.Name {
	return g.name
}
