package db

import (
	"UrlShorter/pkg/generate"
	"database/sql"
	"fmt"
	"log"
)

const (
	Table = "urls"
)

//DB

//ПО факту, Search и Get - одинаковые, просто в 1-ом случае мы ищем
//по столбцу hashID, а во 2-ом по столбцу longURL

type DBRepo struct {
	DB *sql.DB
}

func (d *DBRepo) Search(longURL string) (string, error) {
	checkQuery := fmt.Sprintf("SELECT hashID FROM %s WHERE longURL = ($1)", Table)
	row := d.DB.QueryRow(checkQuery, longURL)
	var foundLongUrl string
	if err := row.Scan(&foundLongUrl); err != nil {
		return "", err
	} else {
		return foundLongUrl, nil
	}
}

func (d *DBRepo) Save(longURL string) string {
	if foundHashId, err := d.Search(longURL); err == nil {
		return foundHashId
	} else {
		hashID := generate.RandSeq(10)
		query := fmt.Sprintf("INSERT INTO %s (hashID, longURL) values($1, $2)", Table)
		_, err := d.DB.Exec(query, hashID, longURL)
		log.Print(err)
		return hashID
	}
	//Нужно дополнить код для обработки дргуих ошибок от БД
}

func (d *DBRepo) Get(hashID string) (string, error) {
	checkQuery := fmt.Sprintf("SELECT longURL FROM %s  WHERE hashID = ($1)", Table)
	row := d.DB.QueryRow(checkQuery, hashID)
	var foundLongUrl string
	if err := row.Scan(&foundLongUrl); err != nil {
		return "", err
	} else {
		return foundLongUrl, nil
	}
}
