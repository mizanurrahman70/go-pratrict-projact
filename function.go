package main
import "fmt"
func getNumber(num1 int, num2 int)(int, int){
 sum := num1 + num2;
 mul := num1 * num2;
 return sum, mul
}
// func getNumber(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2
// 	mul := num1 * num2
// 	return sum, mul
// }
func main (){
	a , b :=getNumber(20,20)
	fmt.Println(a)
	fmt.Println(b)
}