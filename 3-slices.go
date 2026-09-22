package main

func countAdults(ages []int) int {
	adultsCount := 0
	for _, age := range ages {
		if age >= 18 {
			adultsCount++
		}
	}
	return adultsCount
}

func getUserName(users map[int]string, id int) string {
	name, exists := users[id]
	if !exists {
		return "Пользователь не найден"
	}
	return name
}

// func main() {
// 	//================ SLICE (срез) - динамический список элементов одного типа
// 	ages := []int{16, 20, -1, 67, 25, 30, 31, 17}
// 	fmt.Println(countAdults(ages))

// 	users := map[int]string{
// 		1: "Step",
// 		2: "Sam",
// 		3: "Ilya",
// 	}
// 	fmt.Println(getUserName(users, 3))
// Если индекс не требуется то вместо аргумента - "_"
// for _, age := range ages {
// 	fmt.Println(age)
// }
//================ MAP - словарь, хранит пары ключ-значение
// fmt.Println(users[1])

// users[3] = "Anna"
// fmt.Println(users[4])

// name, exists := users[3]

// if !exists {
// 	fmt.Println("Пользователь не найден")
// } else {
// 	fmt.Println(name)
// }
//}
