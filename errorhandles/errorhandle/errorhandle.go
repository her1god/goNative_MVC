package errorhandle

import (
	"go_native/setuproutes/setuproute"
	"log"
	"net/http"
)

func ErrorHandle() {
	r := setuproute.SetupRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server Running At Port 8080")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Tunggu sampai program berhenti
	<-make(chan struct{})
}
