package internal

import "fmt"

/*
   for [Initial Statemtent]; [ Condition ]; [Post Statement] {
     [Action...]
   }
*/

/*
        [I]    [C]   [Post-Statement]
	for i := 0; i < 5; i++ {
		fmt.Println("Hello GoLang", i) // Action...
	}
*/

func ForLoop() {
	// Simple
	for i := 0; i < 5; i++ {
		fmt.Println("Hello GoLang", i)
	}

	// Casting Text into Array/List
	word := "hello"
	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Looping in Array of Numbers
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Looping in Arary/Map/List of Strings
	fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Cherry"}
	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	//Loopings with Conditionals
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("Chegou é 5: %d\n", i)
		}
		fmt.Println(i)
	}
}
