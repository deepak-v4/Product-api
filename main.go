package main

import(
	"net/http"
	"log"
	"os"
	"os/signal"
	"context"
	"./handlers"
	"time"
	"syscall"
)

func main()  {
	

lg := log.New(os.Stdout,"product-api",5)
ph := handlers.NewProduct(lg)
sm := http.NewServeMux()
sm.Handle("/products",ph)
sm.Handle("/products/",ph)
	


s:= http.Server{
	Addr : ":9090",
	Handler: sm,
	ErrorLog:     lg,                 // set the logger for the server
	ReadTimeout:  5 * time.Second,   // max time to read request from the client
	WriteTimeout: 10 * time.Second,  // max time to write response to the client
	IdleTimeout:  120 * time.Second,

}

go func() {

	lg.Println("Starting server on port :9090")
	err := s.ListenAndServe()	
	if err !=nil{
	lg.Printf("Error starting server %s",err)
	os.Exit(1)
}

}()


ch :=make(chan os.Signal,1)


signal.Notify(ch,os.Interrupt)
signal.Notify(ch, syscall.SIGTERM)

sig :=<-ch
log.Println("Got signal",sig)

ctx,_:=context.WithTimeout(context.Background(),30*time.Second)
s.Shutdown(ctx)

}