package main

func main() {

	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }

	// go func() {
	// 	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("HTTP server error: %v", err)
	// 	}
	// 	log.Println("Stopped serving new connections.")
	// }()

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan

	// shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	// defer shutdownRelease()

	// if err := server.Shutdown(shutdownCtx); err != nil {
	// 	log.Fatalf("HTTP shutdown error: %v", err)
	// }
	// log.Println("Graceful shutdown complete.")
}
