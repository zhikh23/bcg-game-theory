package game

type Game interface {
	Start() error
	Play() error
	Results() map[Name]Score
}
