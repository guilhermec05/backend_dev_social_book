package application

func (u *Publish) LikePublish(id_publish, id_author uint64) error {

	return u.repo.LikePublish(id_publish, id_author)
}
