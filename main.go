package main

import (
	"log"
	"main/conn"
	"main/export"
	"main/object"
	"sync"
)

func main() {
	conn.InitDB()
	defer conn.DisconnectAll()
	poolArr := conn.GetPools()
	dbcp := poolArr.Pools[0]
	trgArr, err := object.GetAllProcedures(*dbcp)
	if err != nil {
		log.Println(err)
	}

	var wg sync.WaitGroup

	for _, trg := range trgArr {
		trg := trg
		wg.Add(1)
		go func() {
			defer wg.Done()
			export.SaveFile("./export/"+trg.Db+"/", trg.Name, trg.Def)
		}()
	}

	wg.Wait()

	log.Println("jobs done")

}
