package main

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestForCatImages(t *testing.T) {

	_ = os.Setenv("ALLURE_OUTPUT_PATH", "../tests")
	runner.Run(t, "Check cat image", func(t provider.T) {
		t.Title("Cat Image")
		t.Description("This test will check if any image is returned from the cats folder")

		_, err := RandomFileFromDir("cats")
		if err != nil {
			t.Error("Error getting image from the cat directory")
		}
	})
}

func TestForDogImages(t *testing.T) {

	_ = os.Setenv("ALLURE_OUTPUT_PATH", "../tests")
	runner.Run(t, "Check dog image", func(t provider.T) {
		t.Title("Dog Image")
		t.Description("This test will check if any image is returned from the dogs folder")

		_, err := RandomFileFromDir("dogs")
		if err != nil {
			t.Error("Error getting image from the cat directory")
		}
	})
}

func TestForMetrics(t *testing.T) {

	_ = os.Setenv("ALLURE_OUTPUT_PATH", "../tests")
	runner.Run(t, "Check for prometheus metricss", func(t provider.T) {
		t.Title("Prometheus metrics")
		t.Description("This test will test the prometheus metrics")
		req := httptest.NewRequest("GET", "/metrics", nil)

		rr := httptest.NewRecorder()

		handler := promhttp.Handler()

		handler.ServeHTTP(rr, req)

		response := rr.Result()

		assert.Equal(t, http.StatusOK, response.StatusCode, "Expected status code 200")

	})
}

func TestCatImageHandler(t *testing.T) {
	_ = os.Setenv("ALLURE_OUTPUT_PATH", "../tests")

	runner.Run(t, "Cat handler check", func(t provider.T) {

		t.Title("Checking the cat handler")
		t.Description("This test checks the cat handler")
		req, err := http.NewRequest("GET", "/cat", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		handler := CatImageHandler("templates/cat_template.html", "cats")
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

}

func TestDogImageHandler(t *testing.T) {
	_ = os.Setenv("ALLURE_OUTPUT_PATH", "../tests")
	runner.Run(t, "Dog handler check", func(t provider.T) {
		t.Title("Checking the dog handler")
		t.Description("This test checks the dog handler")
		req, err := http.NewRequest("GET", "/dog", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		handler := DogImageHandler("templates/dog_template.html", "dogs")
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})

}
