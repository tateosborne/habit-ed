package main

import (
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq" // the _ tells go we need this package even though we never rf it
)

const (
	host   = "localhost"
	port   = 5432
	user   = "tateposborne"
	dbname = "habit-ed"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s", host, port, user, dbname)
	fmt.Println(psqlInfo)
}
