package main

import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "tateposborne"
    dbname = "habit-ed"
)

func main() {
	fmt.Println("hello")
}
