package main

import (
	"fmt"
	//"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	)

var db *sql.DB

func connectDB(numberNode string) {
	dbLocal, err := sql.Open("sqlite3", "./databases/sqlite3_"+numberNode+".db")
	checkErr(err)
	//fmt.Println(db)
	db = dbLocal
}


func CreateBlockTable(){
	stmt, err := db.Exec("CREATE TABLE IF NOT EXISTS `blocks` ( `block_id` INTEGER PRIMARY KEY, `timestamp` DATE NULL, `hash` VARCHAR(64), `prevhash` VARCHAR(64) );")
	checkErr(err)
	fmt.Println(stmt)
}

func CreateTxTable(){
	stmt, err := db.Exec("PRAGMA foreign_keys = ON; CREATE TABLE IF NOT EXISTS `transactions` ( `hash` VARCHAR(64) PRIMARY KEY, `timestamp` DATE NULL, `tx_from` VARCHAR(64) NULL, `tx_to` VARCHAR(64) NULL, `amount` VARCHAR(64) NULL, `block_id` INTEGER NOT NULL, FOREIGN KEY (block_id) REFERENCES blocks(block_id) );")
	checkErr(err)
	fmt.Println(stmt)
}

func SaveBlock(block Block) {
	stmt, err := db.Prepare("INSERT INTO blocks(block_id, timestamp, hash, prevhash) values(?,?,?,?)")
	checkErr(err)
	//fmt.Println(stmt)

	res, err := stmt.Exec(block.Index, block.Timestamp, block.Hash, block.PrevHash)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}





func SaveTx(tx Transaction, index_block int) {
	fmt.Println(tx.Hash, tx.Timestamp, tx.From, tx.To, tx.Amount, index_block)
	stmt, err := db.Prepare("INSERT INTO transactions(hash, timestamp, tx_from, tx_to, amount, block_id) values(?,?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(tx.Hash, tx.Timestamp, tx.From, tx.To, tx.Amount, index_block)
	checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
} 




func CreateUserTable(){
	stmt, err := db.Exec("CREATE TABLE `userinfo` ( `uid` INTEGER PRIMARY KEY AUTOINCREMENT, `username` VARCHAR(64) NULL, `departname` VARCHAR(64) NULL, `created` DATE NULL );")
	checkErr(err)
	fmt.Println(stmt)
}

func InsertUser(username string, departname string, created string) {
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	//fmt.Println(stmt)

	res, err := stmt.Exec(username, departname, created)
	checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

}


func GetUsers(){
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	/*var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}*/

	rows.Close() //good habit to close
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
