package storage

import (
	"fmt"
	"my-diet-server/config"

	storage "github.com/ramadani/go-filestorage"
)

var host = config.Config("SERVER_HOST")

var storageConfig = &storage.Config{
	Root: "storage",
	URL:  fmt.Sprintf("%s/public", host),
}

var LocalStorage = storage.NewStorage(storageConfig)
