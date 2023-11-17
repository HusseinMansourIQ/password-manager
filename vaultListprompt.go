package main

import (
	"database/sql"
	"fmt"
)

func vaultListprompt() (string, error) {
	database, err := sql.Open("sqlite3", "pmg.db")
	if err != nil {
		panic(err)
	}
	rows, err := database.Query("SELECT vaultName FROM Vault")

	if err != nil {
		panic(err)
		return "", err
	}
	var vaults []string
	var vaultName string
	for rows.Next() {
		rows.Scan(&vaultName)
		vaults = append(vaults, vaultName)
	}

	//fmt.Println(rows)

	Lable := "Select Vault"

	result, err := createPrompt(Lable, vaults)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}
	return result, err
}
