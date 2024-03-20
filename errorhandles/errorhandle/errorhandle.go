package errorhandle

import (
	"go_native/setuproutes/setuproute"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ErrorHandle() {

	godotenv.Load()
	PortStr := os.Getenv("PORT_SERV")
	Port, err := strconv.Atoi(PortStr)
	if err != nil {
		log.Fatalf("Invalid PORT: %v", err)
	}

	r := setuproute.SetupRoutes()
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(Port),
		Handler: r,
	}

	log.Println("Server Running At Port", Port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Tunggu sampai program berhenti
	<-make(chan struct{})
}
