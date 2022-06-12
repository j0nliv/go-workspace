package models

import (
	"database/sql"
	"fmt"
	"log"
)



const (
	host 		= 		"localhost"
	port 		= 		5432
	user 		= 		"USERNAME"
	password 	= 		"PASSWORD"
	dbname 		= 		"exampledb"
)

var db *sql.DB

func init() {
	var err error
	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	
	db,err = sql.Open("sametemiroglu", connectionStr)
	// db.SetConnMaxIdleTime(5)  -- Max deactive connection
	// db.SetMaxOpenConns(7) -- Max parallel connection
	// db.SetConnMaxIdleTime(1 * time.Second) -- Connection Idle Time
	// db.SetConnMaxLifetime(30 * time.Second) -- Connection Max Life Time 

	if err != nil {
		log.Fatal(err)
	}
}

type Product struct {
	ID					int
	Title,Description	string
	Price				float32	
}


func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title,description,price) VALUES($1, $2 , $3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowAffected, err := result.RowsAffected()
	fmt.Printf("Affected Row Count(%d", rowAffected)
}

func UpdateProduct(data Product){
	result, err := db.Exec("UPDATE products SET title=$2 ,description=$3 ,price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowAffected, err := result.RowsAffected()
	fmt.Printf("Affected Row Count(%d", rowAffected)	
}

func GetProduct() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found!")
		}
		log.Fatal(err)
	}
	defer rows.Close()


	var products []*Product
	for rows.Next() {
		prod := &Product{}
		err := rows.Scan(&prod.ID,&prod.Title,&prod.Description,&prod.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prod)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, pr := range products {
		fmt.Printf("%d - %s, %s, $%.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
}

func GetProdByID(id int) {
	var product string
	err := db.QueryRow("SELECT * FROM products WHERE id=$1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %s\n", product)
	}
}