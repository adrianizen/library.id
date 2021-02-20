package word_repo

type Word struct {
	ID        uint   `json:"id"`
	Word      string `json:"word"`
	CreatedAt string `json:"created_at"`
}
