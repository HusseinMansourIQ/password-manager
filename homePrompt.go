package main

import (
	"database/sql"
	"fmt"

	"github.com/atotto/clipboard"
)

func createVault() {
	database, err := sql.Open("sqlite3", "pmg.db")
	if err != nil {
		panic(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS Vault (id INTEGER PRIMARY KEY, vaultName TEXT , userName TEXT, password TEXT)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	var vaultName string
	var userName string
	var password string

	fmt.Println("Enter Vault Name")
	fmt.Scan(&vaultName)

	fmt.Println("Enter User Name")
	fmt.Scan(&userName)

	fmt.Println("Password : Typer (r)If you want it to be auto generated")
	fmt.Scan(&password)

	rows, err := database.Query("SELECT vaultName FROM Vault WHERE vaultName = ?", vaultName)
	if err != nil {
		panic(err)
	}
	foundVault := false
	var vaultNameRow string

	for rows.Next() {
		rows.Scan(&vaultNameRow)
		foundVault = true
	}

	if foundVault == true {
		homePormpt("This vault is already there ")
	} else {
		if password == "r" {
			randomString, err := genereatePassword()
			if err != nil {
				panic(err)
			}

			password = randomString
		}

		statement, err = database.Prepare("INSERT INTO Vault (vaultName , userName, password) VALUES (?,?, ?)")

		if err != nil {
			panic(err)

		}
		statement.Exec(vaultName, userName, password)

	}
}

func homePormpt(status string) {

	fmt.Println(status)
	Lable := " "
	Items := []string{"Your vaults", "Create vault", "Genereate random passowrd"}
	result, err := createPrompt(Lable, Items)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Your vaults" {
		return
	} else if result == "Genereate random passowrd" {
		randomString, err := genereatePassword()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		clipboard.WriteAll(randomString)
		homePormpt("The password is " + randomString + " and it was copied to your clipboard")
	} else if result == "Create vault" {

		createVault()
		homePormpt("vautl created and the password was copied to your clipboard")

	}
}
