package global

import (
	"by291.fun/scaffold/config"
	"gorm.io/gorm"
)

// some thread-safe global variables

var Config config.Config

var DB *gorm.DB
