package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

func main(){
	//m := sync.Mutex{}
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		//m.Lock()
		number++
		//m.Unlock()
		str := fmt.Sprintf("Você teve acesso a página %d vezes\n",number)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(str))
	})
	http.ListenAndServe(":3000",nil)
}