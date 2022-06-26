
package main

import "fmt"

func main() {


	fmt.Print("how many times do you want your table?: ")


	var number int


	fmt.Scanln(&number)




	for i := 1; i <= number; i++{


		for j := 1; j <= number; j++{


			fmt.Printf("%d %d = %d\n", i, j, i*j)
		}


		fmt.Println("##########")
	}
}