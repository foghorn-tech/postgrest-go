// This is basic example for postgrest-go library usage.
// For now this example is represent wanted syntax and bindings for library.
// After core development this test files will be used for CI tests.

package main

import (
	"fmt"

	"github.com/muratmirgun/postgrest-go"
)

var (
	REST_URL = `http://localhost:3000`
	headers  = map[string]string{}
	schema   = "public"
)

type DataModel struct {
	ID   int
	Task string
	Done bool
}

func main() {
	client := postgrest.NewClient(REST_URL, schema, headers)

	res, err := client.From("todos").Select("id,task,done", "", false).Eq("task", "that created from postgrest-go").Execute()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
