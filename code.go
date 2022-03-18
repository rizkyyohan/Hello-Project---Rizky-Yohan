package main
import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
	var i, bil, total int
	i = 0
	bil = 0
	total = 0
	for i < n {
		fmt.Scanln(&bil)
		for bil < 0 || bil > 9 {
			fmt.Scanln(&bil)
		}
		total += bil
		i ++
	}
	fmt.Println(total)
}