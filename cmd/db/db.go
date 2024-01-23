package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var (
	fileDB       *sql.DB
	reverseIndex *sql.DB
	lemmatizedDB *sql.DB
	stopListDB   *sql.DB
	dbInitMutex  sync.Mutex // Mutex for synchronizing DB initialization
)

func GetFileDB() (*sql.DB, error) {
	dbInitMutex.Lock()
	defer dbInitMutex.Unlock()

	if fileDB != nil {
		return fileDB, nil
	}
	var err error
	fileDB, err = sql.Open("sqlite3", "file:filedb.sqlite?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("file not found error %v", err)
	}
	return fileDB, nil
}

func GetReverseIndexDB() (*sql.DB, error) {
	dbInitMutex.Lock()
	defer dbInitMutex.Unlock()

	if reverseIndex != nil {
		return reverseIndex, nil
	}
	var err error
	reverseIndex, err = sql.Open("sqlite3", "file:reverseindex.sqlite?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("reverse index file not found error %v", err)
	}
	return reverseIndex, nil
}

func GetLemmatizedDB() (*sql.DB, error) {
	dbInitMutex.Lock()
	defer dbInitMutex.Unlock()

	if lemmatizedDB != nil {
		return lemmatizedDB, nil
	}
	var err error
	lemmatizedDB, err = sql.Open("sqlite3", "file:lemmatizeddb.sqlite?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("lemmatized db file not found error %v", err)
	}
	return lemmatizedDB, nil
}

func GetStopListDB() (*sql.DB, error) {
	dbInitMutex.Lock()
	defer dbInitMutex.Unlock()

	if stopListDB != nil {
		return stopListDB, nil
	}
	var err error
	stopListDB, err = sql.Open("sqlite3", "file:stoplistdb.sqlite?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("stop list db file not found error %v", err)
	}
	return stopListDB, nil
}
