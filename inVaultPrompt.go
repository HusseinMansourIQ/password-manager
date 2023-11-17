package main

import (
	"database/sql"
	"fmt"

	"github.com/atotto/clipboard"
)

func editUserName(vault string) {

	var newUserName string
	fmt.Println("enter your new user name")
	fmt.Scan(&newUserName)

	database, err := sql.Open("sqlite3", "pmg.db")
	if err != nil {
		panic(err)
	}
	statement, err := database.Prepare("UPDATE Vault set userName = ? WHERE vaultName = ?")
	if err != nil {
		panic(err)
	}
	statement.Exec(newUserName, vault)

	inVaultPrompt(vault, "your user name has been changed")

}

func editPassword(vault string) {

	var newPassword string
	fmt.Println("enter your new password")
	fmt.Scan(&newPassword)

	database, err := sql.Open("sqlite3", "pmg.db")
	if err != nil {
		panic(err)
	}
	statement, err := database.Prepare("UPDATE Vault set password = ? WHERE vaultName = ?")
	if err != nil {
		panic(err)
	}
	statement.Exec(newPassword, vault)

	inVaultPrompt(vault, "your user name has been changed")

}

func inVaultPrompt(vault string, status string) {

	database, err := sql.Open("sqlite3", "pmg.db")
	if err != nil {
		panic(err)
	}

	Lable := "Select Operation"
	Items := []string{"Copy user name", "Copy Password", "Edit user name", "Edit password", "exit"}

	result, err := createPrompt(Lable, Items)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Copy user name" {
		rows, err := database.Query("SELECT userName FROM Vault WHERE vaultName = ? ", vault)
		if err != nil {
			panic(err)

		}

		var userName string
		for rows.Next() {
			rows.Scan(&userName)
			userName = userName
		}

		clipboard.WriteAll(userName)
		fmt.Println(status)

		inVaultPrompt(vault, "The user name was copied to your clipboard ")

	} else if result == "Copy Password" {

		rows, err := database.Query("SELECT password FROM Vault WHERE vaultName = ? ", vault)
		if err != nil {
			panic(err)
			return
		}

		var password string
		for rows.Next() {
			rows.Scan(&password)
			password = password
		}

		clipboard.WriteAll(password)
		fmt.Println(status)

		inVaultPrompt(vault, "The password was copied to your clipboard ")

	} else if result == "Edit user name" {

		editUserName(vault)

	} else if result == "Edit password" {

		editPassword(vault)

	}

	fmt.Printf("You choose %q\n", result)
}
