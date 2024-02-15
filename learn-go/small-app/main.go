package main

func main() {
	setupSlog()
	err := startApp()
	if err != nil {
		panic(err)
	}
}
