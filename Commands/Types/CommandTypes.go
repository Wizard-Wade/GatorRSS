package CommandTypes

import (
	"GatorRss/internal/config"
	"GatorRss/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}
