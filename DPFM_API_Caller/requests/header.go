package requests

type Header struct {
	Site		int		`json:"Site"`
	IsReleased	*bool	`json:"IsReleased"`
}
