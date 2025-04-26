package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	costPayDay := []float64{}
	for i:=0;i<len(costs);i++{
		if(day == costs[i].day){
			costPayDay = append(costPayDay,costs[i].value )
		}
	}
	return costPayDay
}

func main(){
	c := [7]cost{{0,1},{1,2},{1,3.1},{5,2.5},{2,3.6},{1,2.7},{1,3.3}}
	fmt.Println(getDayCosts(c[:],1))
}