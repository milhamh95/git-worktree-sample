package main

import "fmt"

func main() {
	fmt.Println("Hello from git-worktree-sample!")
	fmt.Println(getText())
	fmt.Println("----")
}

func getText() string {
	return "This is a sample function."
}
