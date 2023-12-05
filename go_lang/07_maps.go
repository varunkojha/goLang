package main
 
import "fmt"
 
func main() {
 
	// A Map is a collection of key value pairs
	// Created with varName := make(map[keyType] valueType)
	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 42
	fmt.Println(presAge["TheodoreRoosevelt"])
	
	// Get the number of items in the Map
	fmt.Println(len(presAge))
	
	// The size changes when a new item is added
	presAge["John F. Kennedy"] = 43
	fmt.Println(len(presAge))
	
	// We can delete by passing the key to delete
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge))
	
	fmt.Println(presAge)

	// We can store multiple items in a map as well
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"realname":"Clark Kent",
			"city":"Metropolis",
		},

		"Batman": map[string]string{
			"realname":"Bruce Wayne",
			"city":"Gotham City",
		},
	}

	// We can output data where the key matches Superman
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}

	for key, value:= range superhero {

		fmt.Println(value, key);
		for k, v:= range value {
			fmt.Println(v, "is", k, "of", key);
		}

	}

 
}
