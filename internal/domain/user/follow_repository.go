package domain

type FollowRepository interface {
	FollowUser(user_id uint64, follow_id uint64) (int, error)
	UnFollowUser(user_id uint64, follow_id uint64) (int, error)
	GetFollows(user_id uint64) ([]User, error)
	GetFollowMe(user_id uint64) ([]User, error)
}
