package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

// func genKeys() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)

// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func init() {
	utils.LoadTemplates()
	config.Load()
	cookies.Configure()
}

func main() {
	fmt.Printf("Rodando WebApp na porta %d\n", config.Port)
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
