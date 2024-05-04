package main

import (
	"fmt"
	"geetest/internal/solver"
)

func main() {
	token, err := solver.New("7a82ed93bc5ef1f522b3cb36093e9bde").Solve()

	if err != nil {
		panic(err)
	}

	fmt.Printf("token: %v\n", token)
}
