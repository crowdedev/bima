package bima

import (
	services "github.com/crowdeco/bima/services"
)

const VERSION_MAYOR = 10000
const VERSION_MINOR = 3000
const VERSION_BUGFIX = 0
const VERSION_STRING = "v1.3.0"

type Repository = *services.Repository
// list disini apa yang perlu di-export untuk minimalisir import e.g import github.com/crowdeco/bima
