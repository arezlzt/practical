package practical

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	Username     string
	Id           int
	Password     string
	PasswordHash string
}

func TestStructColumn(t *testing.T) {
	params := make([]testStruct, 0)
	for i := 0; i < 10; i++ {
		params = append(params, testStruct{
			Username:     fmt.Sprintf("username%v", i),
			Id:           i,
			Password:     fmt.Sprintf("password%v", i),
			PasswordHash: fmt.Sprintf("password_pash%v", i),
		})
	}
	expect := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	actual := StructColumn[testStruct, int](params, "Id")
	assert.Equal(t, expect, actual)
}

func TestStructToMap(t *testing.T) {
	expect := map[string]interface{}{
		"username":      "practical",
		"age":           12,
		"password_salt": "",
	}
	actual := StructToMap(struct {
		Username     string
		Age          int
		PasswordSalt string
	}{
		"practical",
		12,
		"",
	})
	assert.Equal(t, expect, actual)
}
