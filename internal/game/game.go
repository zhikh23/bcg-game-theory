package game

type BinaryGame interface {
	Play(rounds int, a, b *Participant) error
}
