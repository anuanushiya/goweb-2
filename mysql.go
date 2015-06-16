package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
    db, err := sql.Open("mysql", "qiwihui:qiwihui@/sqltest?charset=utf8")
    checkErr(err)

    // Insert data
    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
    checkErr(err)

    res, err := stmt.Exec("qiwihui", "devops", "2015-06-16")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

    // update data
    stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
    checkErr(err)

    res, err = stmt.Exec("qiwihuiupdate", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    // query
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
	var username string
	var departname string
	var created string
	err := rows.Scan(&uid, &username, &departname, &created)
	checkErr(err)
	fmt.Println(uid)
	fmt.Println(username)
	fmt.Println(departname)
	fmt.Println(created)
    }

    // delete data
    stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
