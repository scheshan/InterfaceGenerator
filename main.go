package main

import "InterfaceGenerator/generate"
import "fmt"

func main() {
	str := `
# 用户实体
type User{
	int id
	string name
}

# 产品实体
type Product{
	int id
	string name
	User user
}`

	parser := generate.NewProtoParser()
	types, _ := parser.Parse(str)

	err := recover()

	if err != nil {
		fmt.Println(err)
	}

	for _, t := range types {
		fmt.Println(t)
	}

}
