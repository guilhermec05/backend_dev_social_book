package application

func (u *Publish) UnLikePublish(id_publish, id_author uint64) error {

	return u.repo.UnLikePublish(id_publish, id_author)
}
