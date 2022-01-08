package Mssql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

type Company struct {
	CompanyName string
	CompanyCode string
}

var dbContext = context.Background()

func OpenConnection() *sql.DB{
	db, err := sql.Open("sqlserver", "sqlserver://sa:12QWaszx!!@localhost?database=GoTest&connection+timeout=30")
	if err != nil{
		log.Fatal(err)
	}
	log.Print("Sql con open")
	return db
}

func CreateCompany(db *sql.DB, company Company) error {

	query, err := db.Prepare("INSERT INTO Company ( CompanyName, CompanyCode ) VALUES ( @p1, @p2 );"); if err != nil {
		log.Fatal("Could not insert SqlDB - " , err)
	}

	var ctx context.Context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	defer query.Close()
	rows := query.QueryRowContext(ctx, company.CompanyName, company.CompanyCode)
	if rows.Err() != nil {
		log.Fatal("Could not insert SqlDB2")
	}

	return nil
}

func DeleteCompany(db *sql.DB, companyId int) error {
	queryStatement := fmt.Sprintf("DELETE FROM Company WHERE CompanyId='%v';", companyId)
	query, err := db.Prepare(queryStatement); if err != nil {
		log.Fatal("Could not delete SqlDB - " , err)
	}

	_, queryErr := query.Query(); if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
	}
	defer query.Close()
	return nil
}

func GetAllCompany(db *sql.DB) ([]string, []string, error) {
	var (
		CompanyName []string
		CompanyCode []string
		ctx         context.Context
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT CompanyName,CompanyCode FROM [dbo].[Company];")
	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var _companyName string
		var _companyCode string

		err := rows.Scan(&_companyName, &_companyCode)
		if err != nil{
			return CompanyName,CompanyCode, err
		} else{
			CompanyName = append(CompanyName, _companyName)
			CompanyCode = append(CompanyCode, _companyCode)
		}
	}
	return CompanyName,CompanyCode, nil
}
