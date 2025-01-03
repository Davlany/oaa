package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// words := []string{
	// 	"CREATE",
	// 	"INSERT",
	// 	"PRINT_TREE",
	// 	"CONTAINS",
	// 	"SEARCH",
	// 	"INTERSECTS",
	// 	"CONTAINED_BY",
	// 	"WHERE",
	// }

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		//create_regex := regexp.MustCompile(`(?i)CREATE\s+(\S+)`)
		insert_regex := regexp.MustCompile(`(?i)INSERT\s+(\S+)\s+\{([^}]*)\};`)
		matches := insert_regex.FindStringSubmatch(input)
		if len(matches) > 1 {
			fmt.Println(matches)
			collection_name := matches[1]
			valueList := strings.Split(matches[2], ",")
			fmt.Printf("Sets %s %s %s has been added to %s\n", valueList[0], valueList[1], valueList[2], collection_name)
		}

	}
}
