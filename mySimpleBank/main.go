package main

import "os"

func main() {
	for i := 0; i < 10; i++ {
		os.CreateTemp("C:/Users/Michal/goproject/goProject", "somshit")
	}
}