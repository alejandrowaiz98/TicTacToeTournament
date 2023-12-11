package entitys

type Tournament struct {
}

type ITournament interface {
	New() ITournament
}

func (t *Tournament) New() ITournament {

	return &Tournament{}

}
