package main

import (
	"fmt"
	"github.com/justtaldevelops/hcaptcha-solver-go"
	"github.com/jviguy/libvote"
	"time"
)

func main() {
	// Create the solver with the default options.
	// I wouldn't rely on this solver however, I haven't updated it for a while.
	// Use something like 2captcha, AntiCaptcha, etc. Those are more reliable.
	s, err := hcaptcha.NewSolver("minecraftpocket-servers.com", hcaptcha.SolverOptions{
		SiteKey: "e6b7bb01-42ff-4114-9245-3d2b7842ed92",
	})
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// We provide a deadline that the solver must have the solution done by.
	// If the deadline is not reached, an error is sent instead of the solution.
	solution, err := s.Solve(time.Now().Add(3 * time.Minute))
	if err != nil {
		panic(err)
	}

	// The HyperLands server key. This can be found in the server's URL.
	// Example: https://minecraftpocket-servers.com/server/80103/ (HyperLands)
	serverKey := 80103
	// The username you want to vote with.
	username := "JustTalDevelops"

	ok, err := libvote.NewClient().Vote(serverKey, username, solution)
	if !ok {
		panic(err)
	}
	fmt.Println("Successfully voted for HyperLands!")
}
