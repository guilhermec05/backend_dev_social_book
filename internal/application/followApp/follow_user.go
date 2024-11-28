package application

import "fmt"

func (u *Follow) FollowUser(user_id uint64, follow_id uint64) (int, error) {
	if user_id == follow_id {
		return 0, fmt.Errorf("não pode seguir você mesmo")
	}

	return u.repo.FollowUser(user_id, follow_id)
}
