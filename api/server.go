package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yadavsushil07/user-management/auto"

	"github.com/yadavsushil07/user-management/config"

	"github.com/yadavsushil07/user-management/api/router"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
