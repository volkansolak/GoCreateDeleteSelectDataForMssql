package main

import (
	"GoImportDataMssql/Mssql"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	var db = Mssql.OpenConnection()
	defer db.Close()

	fmt.Println("Welcome to Golang and Microsoft SQL Server Company Creation and Listing Applications.")
	fmt.Println("Select a numeric option \n [1] Create Company \n [2] Delete Company \n [3] Company List")

	consoleReader := bufio.NewScanner(os.Stdin)
	consoleReader.Scan()

	userChoice := consoleReader.Text()

	switch userChoice {
	case "1":
		var (
			companyName,
			companyCode string
		)

		fmt.Println("Please Enter Company Name: ")
		consoleReader.Scan()
		companyName = consoleReader.Text()

		fmt.Println("Please Enter Company Code: ")
		consoleReader.Scan()
		companyCode = consoleReader.Text()

		var comp Mssql.Company = Mssql.Company{CompanyName: companyName, CompanyCode: companyCode}
		err := Mssql.CreateCompany(db, comp); if err != nil{
			fmt.Println("Firma Eklenirken Hata Oluştu: #{err}")
			return
		}
		fmt.Println("Firma Başarıyla Eklendi.")
		break
	case "2":
		companyName, companyCode, err := Mssql.GetAllCompany(db)
		if err != nil{
			fmt.Println(err)
			return
		}

		for i := 0; i < len(companyName); i++{
			fmt.Println(i + 1, companyName[i] + " - " + companyCode[i])
		}

		fmt.Println("Please Choose a Company")
		consoleReader := bufio.NewScanner(os.Stdin)
		consoleReader.Scan()
		companyChoice := consoleReader.Text()
		i, err := strconv.Atoi(companyChoice)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}

		err1 := Mssql.DeleteCompany(db, i)
		if err1 != nil{
			fmt.Println(err1)
			return
		}

		companyName2, companyCode2, err2 := Mssql.GetAllCompany(db)
		if err2 != nil{
			fmt.Println(err2)
			return
		}

		for i := 0; i < len(companyName2); i++{
			fmt.Println(i + 1, companyName2[i] + " - " + companyCode2[i])
		}
		break
	case "3":
		companyName, companyCode, err := Mssql.GetAllCompany(db)
		if err != nil{
			fmt.Println(err)
			return
		}

		for i := 0; i < len(companyName); i++{
			fmt.Println(companyName[i] + " - " + companyCode[i])
		}
		break
	}


}
