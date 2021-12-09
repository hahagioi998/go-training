package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Add 函数的单元测试
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"case1", 1, 2, 3},
		{"case2", 2, 2, 4},
		{"case3", 2, 3, 5},
	}

	for _, tt := range tests {
		assert.Equal(t, Add(tt.a, tt.b), tt.want, tt.name+"they should be equal")
	}

	assert.NotEqual(t, Add(2, 3), 456, "they should not be equal")
}

//作者
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//testify的Json assert使用
func TestJson(t *testing.T) {
	a := &Author{
		ID:   1,
		Name: "Yanfei Zhang",
	}

	j, _ := json.Marshal(a)
	assert.JSONEq(
		t,
		`{"id":1, "name":"Yanfei Zhang"}`,
		string(j),
	)
}
