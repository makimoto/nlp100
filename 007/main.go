package main

import "fmt"

func main() {
	fmt.Println(templete("12", "気温", "22.4"))
}

func templete(x string, y string, z string) string {
	return fmt.Sprintf("%s時の%sは%s", x, y, z)
}
