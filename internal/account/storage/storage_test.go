package storage

import (
	"sqliteMoney/pkg/sqlite3"
	"testing"
	"fmt"
)


func TestGetAccounts(t *testing.T){
	config := sqlite3.NewConfig()
	s := sqlite3.NewStore(config)
	rep := NewRepository(s)

	accs := rep.GetAccounts()

	for _, acc := range accs{
		fmt.Println(acc.ID, acc.Name, acc.Amount)
	}
}