package main

import (
	"testing"
	"fmt"
	"os/exec"
)

func TestYamlConversion(t *testing.T){
	not_exit1:="./testfiles/testfiles/test9.json"
	arg31:="./testfiles/testfiles/test.json"
	arg41:="./testfiles/testfiles/test3.yaml"
	cmd1 := exec.Command("go","run","main.go", "-y", arg31,arg41)
	cmd3 := exec.Command("go","run","main.go", "-y", not_exit1,arg41)
	st, err := cmd1.Output()
	std2,err2:=cmd3.Output()
    if err != nil ||  err2!=nil  {
        t.Fatalf("failed")
    }
	if string(st) != "conversion sucessfull" && string(std2)!=not_exit1+" is not exit" {
		t.Errorf("failed")
	}
  
}
func TestJsonConversion(t *testing.T){
	// if file exit
	arg3:="./testfiles/testfiles/test.yaml"
	not_exit:="./testfiles/testfiles/test9.yaml"
	arg4:="./testfiles/testfiles/test4.json"
	cmd := exec.Command("go","run", "main.go","-j", arg3 ,arg4)
	st, err := cmd.Output()
	// if file not exit
	cmd2 := exec.Command("go","run","main.go", "-j", not_exit,arg4)
	std3,err3:=cmd2.Output()
	if err != nil ||  err3!=nil  {
        t.Fatalf("failed")
    }
	if string(st)!= "conversion sucessfull" && string(std3)!=not_exit+" is not exit"{
		t.Errorf("failed")
		fmt.Printf(string(std3))
	}
	fmt.Println(string(st))
	fmt.Println(string(std3))
}