package main

import "fmt"

func maps() {
	var idMap1 map[string]int // this line can be skipped
	idMap1 = make(map[string]int)
	idMap1["a"] = 1
	fmt.Println("idMap1:", idMap1)

	idMap2 := map[string]int{"joe": 1, "jane": 2}

	fmt.Println("idMap2 joe:", idMap2["joe"])
	fmt.Println("idMap2 not_present:", idMap2["not_present"])

	delete(idMap2, "joe")
	fmt.Println("idMap2 joe deleted:", idMap2["joe"])

	id1, p1 := idMap2["jane"] // p is a boolean. if p is true, key is present. Else not present
	id2, p2 := idMap2["joe"]  // joe is deleted now

	fmt.Println("id1:", id1, "p1:", p1)
	fmt.Println("id2 (not present):", id2, "p2:", p2)

	// len() returns number of elements in map
	fmt.Println("len(idMap2):", len(idMap2))

	fmt.Println("----------------------------------------------------")

	idMap2["joe"] = 1

	// Iterate over map
	fmt.Println("idMap2: Iteration")
	for key, value := range idMap2 {
		fmt.Println("key:", key, "value:", value)
	}
}
