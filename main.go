package main

import (
	"fmt"
	"github.com/guitarpawat/portscan/api/cache"
	"github.com/guitarpawat/portscan/api/model"
	"time"
)

func main() {
	defer cache.CloseDB()
	s := model.GetOutput{
		Results: []model.Result{
			{
				IP:       "127.0.0.1",
				Finished: false,
			},
		},
		LastUpdate: time.Now(),
	}

	b, _ := s.Marshal()
	fmt.Println(b)
	fmt.Println(model.UnMarshalGetOutput(b))
	fmt.Println(cache.PutNewToken("test", "127.0.0.1"))
	fmt.Println(cache.UpdateTokenInfo("test", model.MakeResult("127.0.0.1", 80)))
	fmt.Println(cache.GetTokenInfo("test"))
	fmt.Println(cache.DeleteToken("test"))
}
