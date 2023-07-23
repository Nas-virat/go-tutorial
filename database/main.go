package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Test struct {
	Id   int
	Name string
}


var db *sql.DB
var dbx *sqlx.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<databasename>")
	if err != nil {
		panic(err)
	}

	// Get All
	test, err := GetTests()
	if err != nil{
		fmt.Println(err)
		return
	}

	for _,v := range test{
		fmt.Println(v)
	}

	// Get by id
	test1,err := GetTest(1)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d,%s",test1.Id,test1.Name)

	// insert item
	data := "insert data"
	err = AddTest(data)

	//update
	err = UpdateTest("test update2",3)
	if err != nil{
		panic(err)
	}

}

func GetTestX()([]Test,error){
	query := "select id, name from test"
	tests := []Test{}
	err := dbx.Select(&tests, query)
	if err != nil{
		return nil,err
	}
	return tests,nil
}

func GetTests() ([]Test,error){
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, name from test"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tests := []Test{}

	for rows.Next() {
		test := Test{}
		err = rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	return tests , nil
}

// use point because can not set Test as 
func GetTest(id int)(*Test,error){
	err := db.Ping()
	if err != nil{
		return nil,err
	}

	query :="select id,name from test where id=?"
	// QueryRow for One output row
	row := db.QueryRow(query,id)

	test := Test{}
	err = row.Scan(&test.Id,&test.Name)
	if err != nil{
		return nil,err
	}

	return &test,nil
}

func AddTest(name string)error{
	query:= "insert into test (create_time,name) values (CURRENT_TIMESTAMP,?)"
	result, err := db.Exec(query,name)
	if err != nil{
		return err
	}
	affected,err :=result.RowsAffected()

	if err != nil{
		return err
	}

	if affected <=0{
		return errors.New("cannot insert")
	}

	return nil
}

func UpdateTest(name string,id int) error{
	query:= "update test set name=? where id=?"
	result, err := db.Exec(query,name,id)
	if err != nil{
		return err
	}
	affected,err :=result.RowsAffected()

	if err != nil{
		return err
	}

	if affected <=0{
		return errors.New("cannot update")
	}

	return nil
}

func DeleteTest(id int) error{
	query := "delete from test where id=?"
	result, err := db.Exec(query,id)
	if err != nil{
		return err
	}
	affected,err :=result.RowsAffected()

	if err != nil{
		return err
	}

	if affected <=0{
		return errors.New("cannot update")
	}

	return nil
}
