package models

type Config struct {
	BookDir     string `envconfig:"BOOK_DIR" required:"false"`
	TempDir     string `envconfig:"TEMP_DIR" required:"false"`
	Address     string `envconfig:"ADDRESS" required:"false"`
	IndexCovers bool   `envconfig:"INDEX_COVERS" default:"false"`
}
