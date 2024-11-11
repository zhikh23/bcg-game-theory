package actor

type Factory interface {
	New() (Actor, error)
	MustNew() Actor
}
