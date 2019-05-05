package mmgo

import "time"

const (
	mgoHost         = "mongodb://localhost:27017"
	mgoTimeout      = 60 * time.Second
	mgoAuthdb       = "admin"
	mgoAuthUser     = ""
	mgoAuthPassword = ""
	PoolLimit       = 4096
)
