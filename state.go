package main

import (
	"github.com/DorothyDanger/BlogAggregator/internal/config"
	"github.com/DorothyDanger/BlogAggregator/internal/database"
)

// state holds a pointer to a config
type state struct {
	db     *database.Queries
	config *config.Config
}
