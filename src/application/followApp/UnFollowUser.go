package application

import "fmt"

func (u *Follow) UnFollowUser(user_id uint64, follow_id uint64) (int, error) {
	if user_id == follow_id {
		return 0, fmt.Errorf("não pode seguir você mesmo")
	}

	return u.repo.UnFollowUser(user_id, follow_id)
}
