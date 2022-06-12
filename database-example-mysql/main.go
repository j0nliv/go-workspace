package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql","<USERNAME>:<PASSWORD>@/<DATABASENAME>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	createStatement := "`users`(`ID` int(11) NOT NULL AUTO_INCREMENT, `Username` varchar(45) NOT NULL, `Email` varchar(45) NOT NULL, `Password` varchar(45) NOT NULL, `FirstName` varchar(45) NOT NULL, `LastName` varchar(45) NOT NULL, `BirthDate` varchar(45) DEFAULT NULL, PRIMARY KEY (`ID`), UNIQUE KEY `ID_UNIQUE` (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;"

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS "+ createStatement)
	if err != nil {
		log.Fatal(err)
	}

	var (
		ID 			int
		Username 	string
		Email 		string
		Password 	string
		FirstName 	string
		LastName 	string
		BirthDate 	string
	)

	/*
	res, err := db.Exec("INSERT INTO users(Username, Email, Password, FirstName, LastName, BirthDate) VALUES('samet','7654321!','example@example.com','Samet','Emiroğlu','14.04.1997')")
	if err != nil {
		log.Fatal(err)
	}
	

	rowC, err := res.RowsAffected()
	if err != nil{
		log.Fatal(err)
	}

	log.Printf("Inserted %d rows", rowC)

	// ### MULTIPLE ROW ###

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Datas; %q", strconv.Itoa(ID) + " " +Username+ " " +Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate)
	}
	


	// ### SINGLE ROW ###

	rows, errQ := db.Query("SELECT * FROM users WHERE FirstName = ?", "Samet")
	if errQ != nil {
		log.Fatal(errQ)
	}

	defer rows.Close()

	for rows.Next(){
		errQ = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate)
		if errQ != nil {
			log.Fatal(errQ)
		}
		log.Printf("Datas; %q", strconv.Itoa(ID) + " " +Username+ " " +Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate)
	}
	err = rows.Err()
	if errQ != nil {
		log.Fatal(errQ)
	}

	err = db.QueryRow("SELECT * FROM users limit 1").Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate)
	if err != nil {
		if err == sql.ErrNoRows {
			// empty data process
		}else {
			log.Fatal(err)
		}
	}
	log.Println("A row was found: " , FirstName)
	
	*/

	
	/*
	// ### MULTIPLE ACTIVE RESULT SET (MARS) ###

	_, err := db.Exec("DELETE FROM table1, DELETE FROM table2")
	*/

	/*

	// ### PREPARING QUERIES ###

	statement, errQ := db.Prepare("SELECT * FROM users WHERE ID = ?")
	if errQ != nil {
		log.Fatal(errQ)
	}
	defer statement.Close()
	rows, errX := statement.Query(5)
	if errX != nil {
		log.Fatal(errX)
	}

	defer rows.Close()
	for rows.Next() {
		scanErr := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		log.Printf("Datas; %q", strconv.Itoa(ID) + " " +Username+ " " +Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate)
	}

	*/


}