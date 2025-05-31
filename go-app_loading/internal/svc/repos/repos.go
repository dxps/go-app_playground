package repos

type Repos struct {
	FileRepo *FileRepo
}

func NewRepos() *Repos {
	return &Repos{
		FileRepo: NewFileRepo(),
	}
}
