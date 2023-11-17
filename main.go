package main

import (
	"database/sql"

	"github.com/manifoldco/promptui"
	_ "github.com/mattn/go-sqlite3"
)

func createPrompt(Lable string, Items []string) (string, error) {

	prompt := promptui.Select{
		Label: Lable,
		Items: Items,
	}

	_, result, err := prompt.Run()

	return result, err
}

//"github.com\HusseinMansourIQ\go_playground\src"
//	"github.com/manifoldco/promptui"

func main() {

	database, err := sql.Open("sqlite3", "pmg.db")
	if err != nil {
		panic(err)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS Vault (id INTEGER PRIMARY KEY, vaultName TEXT , userName TEXT, password TEXT)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	//statement, err = database.Prepare("INSERT INTO Vault (vaultName , userName, password) VALUES (?,?, ?)")
	//statement.Exec("google.com", "hussein", "deez")

	//if err != nil {
	//	panic(err)
	//}
	for {

		homePormpt("")

		result, err := vaultListprompt()

		if err != nil {
			panic(err)
		}
		inVaultPrompt(result, " ")

	}
	//genereatePassword()

}
