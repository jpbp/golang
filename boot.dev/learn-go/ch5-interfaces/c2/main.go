package main

import (
	"fmt"

)

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() (calcImportance int){
	if dm.isUrgent{
		return 50
	}
	return dm.priorityLevel
}

func (gm groupMessage) importance() (calcImportance int){
	return gm.priorityLevel
}
func (sa systemAlert) importance() (calcImportance int){
	return 100
}

func processNotification(n notification) (string, int) {
	switch newNotification := n.(type){
		case directMessage: 
			return newNotification.senderUsername , n.importance()
		case groupMessage:
			return newNotification.groupName , n.importance()
		case systemAlert:
			return newNotification.alertCode, n.importance()
		default:
			return "",0
	}
}

func main(){
	dm := directMessage{senderUsername: "jp",messageContent: "teste123",priorityLevel: 12,isUrgent: false}
	gm := groupMessage{groupName: "jp-group",messageContent: "teste456",priorityLevel: 45}
	sa := systemAlert{alertCode: "alerta!!",messageContent: "teste789"}

	fmt.Println(processNotification(dm))
	fmt.Println(processNotification(gm))
	fmt.Println(processNotification(sa))
}