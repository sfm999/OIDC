package main

import (
	"fmt"
	re "project/oidc/resource_owner"
	"project/oidc/utils"
)

func main() {
	resOwner := re.NewResourceOwner("1", "Stu", "stu@hotmail.com")
	fmt.Println(resOwner)
	fmt.Println(resOwner.CsvString())
	postData, err := resOwner.PrepareForPost()
	utils.Check(err, "Error creating Post data for "+resOwner.Username)
	fmt.Println(string(postData))
}
