package main

import (
	"fmt"
	"github.com/guitarpawat/portscan/api/model"
	"time"
)

func main() {
	s := model.GetOutput{
		Results: []model.Result {
			{
				IP: "127.0.0.1",
				Finished: false,
			},
		},
		LastUpdate: time.Now(),
	}

	b, _ := s.Marshal()
	fmt.Println(b)
	fmt.Println(model.UnMarshalGetOutput(b))
}
