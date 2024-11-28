package application

func (u *Publish) DeletePublish(id_publish, id_author uint64) error {

	return u.repo.DeletePublish(id_publish, id_author)
}
