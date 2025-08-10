package core

import "github.com/arian-press2015/uniac/internal/validators"

type Database struct {
	Name    string
	Storage string
	Region  string
	Engine  string
	Version float64
	Tags    map[string]string
}

func NewDatabase(configDatabase *validators.Database) *Database {
	return &Database{
		Name:    configDatabase.Name,
		Storage: configDatabase.Storage,
		Region:  configDatabase.Region,
		Engine:  configDatabase.Engine,
		Version: configDatabase.Version,
		Tags:    configDatabase.Tags,
	}
}
