package main
import ("fmt")

func euclidean_algorithm(integer1 int, integer2 int) (int) {
	for integer2 != 0 {
		integer1, integer2 = integer2, integer1 % integer2;
	}
	return integer1
}

func update_abxy_myarray(a int, b int, x int, y int, my_array []int) (int, int, int, int, []int) {
	return my_array[0], a, y, - (my_array[0] / a) * y + x, my_array[1 : len(my_array)]
}

func extended_euclidean_algorithm(input1 int, input2 int) (x int, y int, gcd int) {
	var my_array = []int{};
	var integer1, integer2 = input1, input2;
	var a, b int;

	for integer2 != 0 {
		my_array = append([]int{integer1}, my_array...);
		integer1, integer2 = integer2, integer1 % integer2;
	}
	gcd = integer1;

	a, b, x, y = gcd, 0, 1, 0; //x*a + y*b = gcd, and this will remain true

	for (a != input1) || (b != input2) {
		a, b, x, y, my_array = update_abxy_myarray(a, b, x, y, my_array);
	}
	return

}

func main() {
	var n = 41;
	var m = 8;

	fmt.Println("The gcd of", n, "and", m, "is:", euclidean_algorithm(n, m));
	var x, y, gcd = extended_euclidean_algorithm(n, m);
	fmt.Println(n, "*", x, "+", m, "*", y, "=", gcd);
}
