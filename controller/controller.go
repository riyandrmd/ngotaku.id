package controller

import (
	"animebase/connection"
	"animebase/models"
	"encoding/json"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDB()
}

func PostStudio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data []models.Studio

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		DB.Create(&data)
		w.Write([]byte("Suscces Post Data"))
		w.WriteHeader(200)
		return
	}
}

func PostGenre(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data []models.Genre

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Suscces Post Data"))
		w.WriteHeader(200)
		return
	}
}

func PostAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		var data []models.Anime

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		DB.Create(&data)
		w.Write([]byte("Suscces Post Data"))
		w.WriteHeader(200)
		return
	}
}

func GetAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Anime

		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.Write(datajson)
		w.WriteHeader(200)
		return
	}

	http.Error(w, "Error Not Found", 404)
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Genre

		DB.Model(&models.Genre{}).Preload("Anime").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetStudio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Studio

		DB.Model(&models.Studio{}).Preload("Anime").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetAnimeFull(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var result []models.AnimeFull

		DB.Table("Animes").Select("animes.anime_id,animes.title,genres.genre,animes.status,animes.rating,studios.studio_name, animes.poster_url, animes.trailer").Joins("LEFT JOIN studios ON studios.studio_id = animes.studio_id").Joins("LEFT JOIN genres ON genres.genre_id = animes.genre_id").Scan(&result)

		datajson, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetAnimeFullbyId(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var result []models.AnimeFull

		DB.Table("Animes").Select("animes.anime_id,animes.title,genres.genre,animes.status,animes.rating,studios.studio_name, animes.poster_url, animes.trailer").Joins("LEFT JOIN studios ON studios.studio_id = animes.studio_id").Joins("LEFT JOIN genres ON genres.genre_id = animes.genre_id").Where("animes.anime_id = ?", id[2]).Scan(&result)

		datajson, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func DeleteAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Anime{}, "anime_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Anime{}, "anime_id = ?", id[2])

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Genre{}, "genre_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Genre{}, "genre_id = ?", id[2])

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func DeleteStudio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Studio{}, "studio_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Studio{}, "studio_id = ?", id[2])

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func UpdateAnime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var datarequest models.Anime
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		err := DB.First(&models.Anime{}, "anime_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.Anime{}).Where("anime_id = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
}

func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var datarequest models.Genre
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		err := DB.First(&models.Genre{}, "genre_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.Genre{}).Where("genre_id = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
}

func UpdateStudio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var datarequest models.Studio
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		err := DB.First(&models.Studio{}, "studio_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.Studio{}).Where("studio_id = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
}
