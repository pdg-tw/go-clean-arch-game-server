package persistence

import (
	"go-clean-arch-game-server/internal/domain/entities/member"
)

type MemberMemRepository struct {
	members map[string]member.Member
}

func NewMemberMemRepository() member.Repository {
	members := make(map[string]member.Member)
	return &MemberMemRepository{members}
}

// Add the provided member
func (r *MemberMemRepository) Add(member member.Member) error {
	r.members[member.ID.String()] = member
	return nil
}
