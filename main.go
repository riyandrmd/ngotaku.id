package main

import (
	"animebase/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/poststudio", controller.PostStudio)
	http.HandleFunc("/postgenre", controller.PostGenre)
	http.HandleFunc("/postanime", controller.PostAnime)

	http.HandleFunc("/getstudio", controller.GetStudio)
	http.HandleFunc("/getanime", controller.GetAnime)
	http.HandleFunc("/getgenre", controller.GetGenre)
	http.HandleFunc("/getanimefull", controller.GetAnimeFull)
	http.HandleFunc("/getanimefull/", controller.GetAnimeFullbyId)

	http.HandleFunc("/deleteanime/", controller.DeleteAnime)
	http.HandleFunc("/d`eletegenre/", controller.DeleteGenre)
	http.HandleFunc("/deletestudio/", controller.DeleteStudio)

	http.HandleFunc("/updateanime/", controller.UpdateAnime)
	http.HandleFunc("/updategenre/", controller.UpdateGenre)
	http.HandleFunc("/updatestudio/", controller.UpdateStudio)

	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Printf("Error Starting Service")
	}
}
