package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

type PageData struct {
	ImageUrl string
}

var catApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "Cat_API_Ping_Counter",
	Help: "Number of pings made to the endpoint",
})

var DogApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "Dog_API_Ping_Counter",
	Help: "Number of pings made to the endpoint",
})

func RandomFileFromDir(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", fmt.Errorf("no files in directory %s", dir)
	}
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	randomIndex := random.Intn(len(files))
	return files[randomIndex].Name(), nil
}

func CatImageHandler(templatePath, imagesDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		catApiPingCounter.Inc()
		imageName, err := RandomFileFromDir(imagesDir)
		if err != nil {
			http.Error(w, "Failed to get a random image", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		data := PageData{
			ImageUrl: "/cats/" + imageName,
		}

		tmpl.Execute(w, data)
	}
}
func DogImageHandler(templatePath, imageDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		DogApiPingCounter.Inc()
		imageName, errorGettingImage := RandomFileFromDir(imageDir)
		if errorGettingImage != nil {
			http.Error(w, "Error getting the file", http.StatusInternalServerError)
			return
		}
		Dogtemplate := template.Must(template.ParseFiles(templatePath))
		data := PageData{
			ImageUrl: "/dogs/" + imageName,
		}
		Dogtemplate.Execute(w, data)
	}
}
func main() {
	prometheus.MustRegister(catApiPingCounter)
	prometheus.MustRegister(DogApiPingCounter)
	http.Handle("/cats/", http.StripPrefix("/cats/", http.FileServer(http.Dir("cats"))))
	http.Handle("/dogs/", http.StripPrefix("/dogs/", http.FileServer(http.Dir("dogs"))))
	http.HandleFunc("/cat", CatImageHandler("templates/cat_template.html", "cats"))
	http.HandleFunc("/dog", DogImageHandler("templates/dog_template.html", "dogs"))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
}
