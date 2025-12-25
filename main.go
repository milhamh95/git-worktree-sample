package main

import "fmt"

func main() {
	fmt.Println("Hello from git-worktree-sample!")
	fmt.Println(getText())
}

func getText() string {
	return "This is a sample function."
}

func getText2() string {
	return "this is utility branch"
}
