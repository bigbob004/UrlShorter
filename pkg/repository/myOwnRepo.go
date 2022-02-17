package repository

import (
	gen "UrlShorter/pkg/generate"
	"errors"
)

type MyRepo struct {
	ShortToLong map[string]string
	LongToShort map[string]string
}

func (rep MyRepo) Get(hash string) (string, error) {
	if findedLongURL, ok := rep.ShortToLong[hash]; !ok {
		return "", errors.New("There isn't key")
	} else {
		return findedLongURL, nil
	}
}

func (rep *MyRepo) Save(longURL string) string {
	if finded_hash, ok := rep.LongToShort[longURL]; !ok {
		hash := gen.RandSeq(10)
		rep.ShortToLong[hash] = longURL
		rep.LongToShort[longURL] = hash
		return hash
	} else {
		return finded_hash
	}
}
