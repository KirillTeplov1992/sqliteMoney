package sqlite3

import "testing"


func TestStore(t *testing.T){
	config := NewConfig()
	s := NewStore(config)
	if err := s.Open(); err != nil{
		t.Fatal(err)
	}
}