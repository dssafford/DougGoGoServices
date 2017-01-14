package main

import (
<<<<<<< HEAD
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	service "github.com/cloudnativego/gogo-service/service"
=======
"fmt"
"os"

"github.com/cloudfoundry-community/go-cfenv"
"github.com/cloudnativego/gogo-service/service"
>>>>>>> afdb79a1d0ece6b42548f322c0bc546e2551821a
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("CF Environment not detected.")
	}

	server := service.NewServer(appEnv)
	server.Run(":" + port)
}
<<<<<<< HEAD
=======

>>>>>>> afdb79a1d0ece6b42548f322c0bc546e2551821a
