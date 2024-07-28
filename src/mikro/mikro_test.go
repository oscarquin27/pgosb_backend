package mikro

import (
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestChaining(t *testing.T) {
	
	pool:= &pgxpool.Pool{}

	m := NewMkModel(pool)
	
	type User struct {
		Id  string `mk:"id_user"`
		Name pgtype.Text `mk:"name"`
		Username string `mk:"user_name"`
	}
	
	r, err := m.Model(&User{Id: "12345"}).Insert("tabla.user_123456")

	fmt.Println(r)

    r, err = m.Model(&User{Id: "12345"}).Omit("id_user").Where("id_user", "=", 1).Update("tabla.user_123456")

	fmt.Println(r)

	if err != nil {
		t.Fail()
	}
}