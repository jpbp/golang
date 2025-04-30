package main
import "fmt"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggested   := []string{}
	users := make(map[string]int)

	for _,friends := range friendships[username]{
		users[friends]=1
	}

	for _,friends := range friendships[username]{
		for _,friendDoMeuFriend := range friendships[friends]{
			if _,ok := users[friendDoMeuFriend]; !ok && username != friendDoMeuFriend{
				users[friendDoMeuFriend]=1
				suggested = append(suggested,friendDoMeuFriend)
			}

		}
	}
	if (len(suggested)>0){
		return suggested
	}
	return nil
}

func main (){

	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}
	suggestedFriends := findSuggestedFriends("E", friendships)
	fmt.Println(suggestedFriends)
}