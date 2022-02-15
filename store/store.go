package store

import (
	"fmt"
)

type StorageService struct {
}

func InitializeStore() *StorageService {
	return &StorageService{}
}

// if user id was not provided generate one on the fly : case for not logged in users

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/
func SaveUrlMapping(shortUrl string, originalUrl string) {
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/
func RetrieveInitialUrl(shortUrl string) string {
	return "www.google.com"
}
