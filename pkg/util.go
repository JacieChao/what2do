package pkg

import (
	"math/rand"
)

var (
	todoPlace = []string{"Japan", "Shenyang", "Beijing"}
	todoList  = map[string][]string{
		"Japan": []string{
			"大阪USJ",
			"奈良看小鹿",
			"京都稻禾大神",
		},
		"Shenyang": []string{
			"鸟岛",
			"莫子山公园",
			"故宫大帅府",
			"世博园",
		},
		"Beijing": []string{
			"环球影城",
			"故宫",
			"颐和园",
			"圆明园",
		},
	}
)

func Random() map[string][]string {
	index := rand.Intn(len(todoPlace))
	randomPlace := todoPlace[index]
	randomList := todoList[randomPlace]
	rTodo := map[string][]string{}
	rTodoList := []string{}
	for i := 0; i < 2; i++ {
		rIndex := rand.Intn(len(randomList))
		rTodoList = append(rTodoList, randomList[rIndex])
	}
	rTodo[randomPlace] = rTodoList
	return rTodo
}
