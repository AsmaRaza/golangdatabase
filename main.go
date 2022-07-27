package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "princes13"
	dbname   = "caldb"
)

type DBConnect struct {
	conn *sql.DB
}

type Numbers struct {
	Num1 int `json:"number1"`
	Num2 int `json:"number2"`
}
type Response struct {
	result float64 `json:"result"`
}
type GetAllData struct {
	ID         int     `json:"id" db:"id"`
	Num1       int     `json:"number1" db:"number1"`
	Num2       int     `json:"number2" db:"number2"`
	result     float64 `json:"result" db:"result"`
	operations string  `json:"operation" db:"operation"`
}

func (db DBConnect) Addition(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 + u.Num2
	fmt.Println("The Addition of 2 numbers is=", result)
	sqlInsert := `INSERT INTO "myschema".arthimetic(Num1, Num2, result, operations)VALUES( $1, $2, $3, $4)`
	stmt, err := db.conn.Prepare(sqlInsert)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Num1, u.Num2, result, "+")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(result)

	return c.JSON(http.StatusOK, result)
}
func (db DBConnect) Subtraction(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 - u.Num2
	fmt.Println(result)
	sqlInsert := `INSERT INTO "myschema".arthimetic( Num1, Num2, result, operations)VALUES($1, $2, $3, $4)`
	stmt, err := db.conn.Prepare(sqlInsert)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Num1, u.Num2, result, "-")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}
func (db DBConnect) Multiplication(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 * u.Num2
	fmt.Println("The Multiplication of 2 numbers is=", result)
	sqlInsert := `INSERT INTO "myschema".arthimetic( Num1, Num2, result, operations)VALUES($1, $2, $3, $4)`
	stmt, err := db.conn.Prepare(sqlInsert)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Num1, u.Num2, result, "*")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}
func (db DBConnect) Division(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := u.Num1 / u.Num2

	fmt.Println(result)
	sqlInsert := `INSERT INTO "myschema".arthimetic( Num1, Num2, result, operations)VALUES($1, $2, $3, $4)`
	stmt, err := db.conn.Prepare(sqlInsert)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Num1, u.Num2, result, "/")
	if err != nil {
		fmt.Println(err)
		panic(
			err)
	}
	return c.JSON(http.StatusOK, result)

}

func (db DBConnect) SquareRoot(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}
	result := math.Sqrt(float64(u.Num1))
	fmt.Println(result)
	sqlInsert := `INSERT INTO "myschema".arthimetic( Num1, Num2, result, operations)VALUES($1, $2, $3, $4)`
	stmt, err := db.conn.Prepare(sqlInsert)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Num1, u.Num2, result, "√ ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}
func (db DBConnect) Modulus(c echo.Context) error {
	u := new(Numbers)

	if err := c.Bind(u); err != nil {
		return err
	}

	result := u.Num1 % u.Num2
	fmt.Println("The Modulus  of 2 numbers is=", result)
	sqlInsert := `INSERT INTO "myschema".arthimetic( Num1, Num2, result, operations)VALUES( $1, $2, $3, $4)`
	stmt, err := db.conn.Prepare(sqlInsert)
	if err != nil {
		fmt.Println("HERE", err)
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Num1, u.Num2, result, "%")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}
func database() *DBConnect {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return &DBConnect{
		conn: db,
	}
}
func (db DBConnect) GetAllData(c echo.Context) error {

	rows, err := db.conn.Query(`SELECT * FROM "myschema".arthimetic `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ID int
		var Num1 int
		var Num2 int
		var result float64
		var operations string

		err = rows.Scan(&ID, &Num1, &Num2, &result, &operations)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(ID, Num1, Num2, result, operations)
	}

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, "success")
}

func (db DBConnect) DeleteData(c echo.Context) error {

	id := c.Param("id")
	fmt.Println(id)

	_, err := db.conn.Query(`DELETE FROM "myschema".arthimetic where Id=$1`, id)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)

	}
	return c.String(http.StatusOK, "deleted")
}

func main() {
	db := database()
	e := echo.New()
	e.POST("/addition", db.Addition)
	e.POST("/addition", db.Addition)
	e.POST("/subtraction", db.Subtraction)
	e.POST("/multipication", db.Multiplication)
	e.POST("/division", db.Division)
	e.POST("/squareroot", db.SquareRoot)
	e.POST("/modulus", db.Modulus)
	e.GET("/getalldata", db.GetAllData)
	e.DELETE("/deleterecord/:id", db.DeleteData)
	e.Start(":6001")

}
