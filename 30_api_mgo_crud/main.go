package main

import (
	"../30_api_mgo_crud/config"
	"../30_api_mgo_crud/dao"
	"../30_api_mgo_crud/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var (
	confi  config.Configuration // Database Configuration
	dataAO dao.MoviesDAO        // Data Access Object
)

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	confi.Read()
	dataAO.Server = confi.Server
	dataAO.Database = confi.Database
	dataAO.Connect()
}

// Define HTTP request routes
func main() {
	fmt.Println("HTTP port :3000")
	r := httprouter.New() // methods GET, POST, PUT, PATCH and DELETE

	r.GET("/", redirect)
	r.GET("/movies", AllMoviesEndPoint)
	r.POST("/movies", CreateMovieEndPoint)
	r.PUT("/movies", UpdateMovieEndPoint)
	r.DELETE("/movies", DeleteMovieEndPoint)
	r.GET("/movies/:id", FindMovieEndpoint)

	http.ListenAndServe(":3000", r)
}

// redirect
func redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/movies", http.StatusSeeOther)
}

// AllMoviesEndPoint - GET list of movies
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := dataAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, movies)
}

// FindMovieEndpoint - GET a movie by its ID
func FindMovieEndpoint(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	movie, err := dataAO.FindById(ps.ByName("id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJSON(w, http.StatusOK, movie)
}

// CreateMovieEndPoint - POST a new movie
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dataAO.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, movie)
}

// UpdateMovieEndPoint - PUT update an existing movie
func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dataAO.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteMovieEndPoint an existing movie
func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dataAO.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// curl -sSX POST -d '{"name":"dunkirk","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"world war 2 movie"}' http://localhost:3000/movies | jq '.'
// curl -i -X POST -d "isbn=001-8484314701&title=How to run Windows&author=Bill Gates&price=1.90" 192.168.0.2:8080/items/create/process
// curl -sSX POST -d '{"name":"dunkirk","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"world war 2 movie"}' http://192.168.1.2:3000/movies | jq '.'
// curl -sSX POST -d '{"name":"dunkirk","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"world war 2 movie"}' http://localhost:3000/movies
// curl -sSX DELETE -d '{"name":"dunkirk","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"world war 2 movie"}' http://localhost:3000/movies | jq '.'
/*
// UPDATE AND DELETE
curl -sSX PUT -d '{"id":"5a33deba263add78f980fcab", "name":"UPDATED name","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"world war 2 movie"}' http://localhost:3000/movies
curl -sSX DELETE -d '{"id":"5a33deba263add78f980fcab", "name":"UPDATED","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"world war 2 movie"}' http://localhost:3000/movies

// UPDATE AND DELETE
curl -sSX PUT -d '{"id":"5a33df1d263add78f980fcad", "name":"UPDATED 2","cover_image":"zHCcN.jpg", "description":"movie"}' http://localhost:3000/movies
curl -sSX DELETE -d '{"id":"5a33df1d263add78f980fcad", "name":"UPDATED 2","cover_image":"zHCcN.jpg", "description":"movie"}' http://localhost:3000/movies

// VIEW ONE
curl -sSX GET http://localhost:3000/movies/5a33df2d263add78f980fcae
http://192.168.1.2:3000/movies/5a33df2d263add78f980fcae

// VIEW ALL
curl -sSX GET http://localhost:3000/movies
http://192.168.1.2:3000/movies

// CREATE ONE
curl -i -X POST -d '{"name":"zelenko movie","cover_image":"https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg", "description":"the new family movie"}' http://localhost:3000/movies


// delete
curl -sSX DELETE -d '{"id":"5a33df2d263add78f980fcae"}' http://localhost:3000/movies

// Docker, Angular, Go, AWS
*/
