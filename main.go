// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// func main() {
// 	// words := []string{
// 	// 	"CREATE",
// 	// 	"INSERT",
// 	// 	"PRINT_TREE",
// 	// 	"CONTAINS",
// 	// 	"SEARCH",
// 	// 	"INTERSECTS",
// 	// 	"CONTAINED_BY",
// 	// 	"WHERE",
// 	// }

// 	reader := bufio.NewReader(os.Stdin)
// 	full_str := ""
// 	for {
// 		input, _ := reader.ReadString('\n')
// 		input = strings.ReplaceAll(input, " ", "")
// 		input = strings.TrimSpace(input)
// 		if input == "RUN" {
// 			ReadString(full_str)
// 		} else {
// 			full_str += input
// 		}

// 	}

// }

// func ReadString(input string) {
// 	input = strings.TrimSpace(input)
// 	create_regex := regexp.MustCompile(`(?i)CREATE(\S+?);`)
// 	insert_regex := regexp.MustCompile(`(?i)INSERT(\S+)\{([^}]*)\};`)
// 	insert_matches := insert_regex.FindStringSubmatch(input)
// 	create_matches := create_regex.FindStringSubmatch(input)

//		if len(insert_matches) > 1 {
//			collection_name := insert_matches[1]
//			valueList := strings.Split(insert_matches[2], ",")
//			fmt.Print("Sets ")
//			for _, v := range valueList {
//				fmt.Printf("%s ", v)
//			}
//			fmt.Printf("has been added to %s\n", collection_name)
//		}
//		if len(create_matches) > 1 {
//			collection_name := create_matches[1]
//			fmt.Printf("Collection %s has been created\n", collection_name)
//		}
//	}
package main

import (
	"fmt"
	"sort"
)

// Node представляет узел дерева
type Node struct {
	Set      []int   // Множество, связанное с узлом
	Children []*Node // Дочерние узлы
}

// NewNode создает новый узел
func NewNode(set []int) *Node {
	return &Node{
		Set:      set,
		Children: []*Node{},
	}
}

// Insert вставляет новое множество в дерево
func (n *Node) Insert(newSet []int) {
	// Сортируем множество для согласованности
	sort.Ints(newSet)

	// Ищем, есть ли дочерний узел, который может включать новое множество
	for _, child := range n.Children {
		if isSubset(newSet, child.Set) {
			child.Insert(newSet)
			return
		}
	}

	// Добавляем новое множество как дочерний узел
	newNode := NewNode(newSet)

	// Проверяем, есть ли дочерние узлы, которые являются подмножествами нового множества
	var remainingChildren []*Node
	for _, child := range n.Children {
		if isSubset(child.Set, newSet) {
			newNode.Children = append(newNode.Children, child)
		} else {
			remainingChildren = append(remainingChildren, child)
		}
	}
	n.Children = remainingChildren
	n.Children = append(n.Children, newNode)
}

// isSubset проверяет, является ли одно множество подмножеством другого
func isSubset(subset, superset []int) bool {
	j := 0
	for i := 0; i < len(superset) && j < len(subset); i++ {
		if superset[i] == subset[j] {
			j++
		}
	}
	return j == len(subset)
}

// PrintTree выводит дерево с правильными отступами
func (n *Node) PrintTree(level int, prefix string) {
	// Формируем префикс для отступов
	fmt.Printf("%s[%v]\n", prefix, n.Set)

	// Для каждого дочернего узла вызываем PrintTree
	for _, child := range n.Children {
		child.PrintTree(level+1, prefix+"│   ")
	}
}

func main() {
	// Начальное множество (универсальное множество)
	root := NewNode([]int{1, 2, 3, 4, 5, 8, 9})

	// Генерация больших данных (множества)
	sets := [][]int{
		{1, 2, 3},
		{2, 5},
		{1, 4, 8},
		{1, 4, 9},
		{0, 1, 2, 3, 4},
		{1, 2, 4, 5},
		{2, 3, 5, 7},
		{1, 3, 4, 6, 7},
		{0, 1, 3, 4, 6, 8},
		{3, 5, 6, 7, 9},
		{0, 2, 4, 7, 8},
		{1, 4, 7, 9},
		{0, 3, 5, 8},
		{1, 2, 4, 6, 8},
		{2, 4, 6, 9},
		{1, 3, 4, 5, 6},
		{0, 2, 6, 7, 9},
		{2, 4, 5, 8},
		{1, 4, 5, 6, 9},
		{1, 2, 3, 5, 7},
		{0, 3, 4, 5, 9},
		{2, 4, 6, 8, 9},
	}

	// Вставляем множества в дерево
	for _, set := range sets {
		root.Insert(set)
	}

	// Выводим дерево
	root.PrintTree(1, "")
}
