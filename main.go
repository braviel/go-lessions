package main

import (
	"fmt"
	"net/http"
	"os"
	"routine/boring"
	"routine/di"
	"routine/mock"
	"time"
)

//MyGreeterHandler Greet to response
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	Joe := boring.Boring("Joe")
	Ann := boring.Boring("Ann")
	Fan := boring.FanIn(Joe, Ann)
	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-Fan)
	}
	fmt.Println("You're boring, I'm leaving")
	//
	di.Greet(os.Stdout, "Elodie")
	//
	//http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	// learn mocking
	fmt.Println("\n#########")
	sleeper := &mock.ConfigurableSleeper{Duration: 1 * time.Second, SleepImpl: time.Sleep}
	mock.Countdown(os.Stdout, sleeper)
}
