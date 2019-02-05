package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
)

type User struct {
	uid  int
	name string
}

//const (
//	host     = "172.16.10.133"
//	port     = 5432
//	user     = "sneaky"
//	password = "77WN88wwc"
//	dbname   = "sneaky"
//)
//
//func main() {
//	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Successfully connected")
//	query := "select uid,name from pw_user where uid=991"
//	//rows, err := db.Query(query)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//user := User{}
//	//for rows.Next() {
//	//	err := rows.Scan(&user.uid, &user.name)
//	//	if err != nil {
//	//		log.Fatal(err)
//	//	}
//	//}
//	var user User
//	row := db.QueryRow(query)
//	err = row.Scan(&user.uid, &user.name)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(user)
//}

var db *sql.DB

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func connPG() {
	var err error
	db, err = sql.Open("postgres", "postgres://sneaky:77WN88wwc@172.16.10.133:5432/sneaky?sslmode=verify-full")
	logErr(err)
}

func insert(uid int, name string) {
	stmt, err := db.Prepare("INSERT INTO pw_user (uid,name) VALUES ($1,$2)")
	logErr(err)
	res, err := stmt.Exec(uid, name)
	logErr(err)
	affect, err := res.RowsAffected()
	logErr(err)
	fmt.Println("rows affect:", affect)
}

func delete(uid int) {
	stmt, err := db.Prepare("DELETE FROM pw_user where uid=$1")
	logErr(err)
	res, err := stmt.Exec(uid)
	logErr(err)
	affect, err := res.RowsAffected()
	logErr(err)
	fmt.Println("rows affect:", affect)
}

