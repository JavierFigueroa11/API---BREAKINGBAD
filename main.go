package main
import(
	"awesomeProject/router"
	"time"
)

func main () {
	go router.MapRoutes()
	time.Sleep(1)
	router.Run()
}
