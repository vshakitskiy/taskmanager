package main

func main() {
	api := NewServer(Config{Addr: ":8080"})
	_ = api.Start()
}
