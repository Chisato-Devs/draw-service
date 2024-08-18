package context

import (
	"chisato-draw-service/server/controllers"
	"chisato-draw-service/server/v1/handlers/structs"
	"encoding/json"
	"errors"
	"net/http"
)

func CreateErrorResponse(w http.ResponseWriter, response string, status int) {
	logger := controllers.Logger()

	jsonResponse, err := json.Marshal(
		structs.ErrorResponse{Message: response},
	)
	if err != nil {
		http.Error(w, "An error occurred while serializing an object to json", http.StatusInternalServerError)
		logger.Warnf("An error occurred while serializing an object to json (500 Internal Server Error)")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsonResponse)

	if err != nil {
		return
	}
}

func CreateSuccessResponse(w http.ResponseWriter, response interface{}) {
	logger := controllers.Logger()

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "An error occurred while serializing an object to json", http.StatusInternalServerError)
		logger.Warnf("An error occurred while serializing an object to json (500 Internal Server Error)")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	logger.Infof("Successfully rendered an image")

	if err != nil {
		return
	}
}

func GetParameterFromURL(r *http.Request, paramName string, errorMessage string, w http.ResponseWriter) (string, error) {
	paramValue := r.URL.Query().Get(paramName)
	if paramValue == "" {
		logger := controllers.Logger()
		CreateErrorResponse(w, errorMessage, http.StatusBadRequest)

		logger.Warnf("Missing required parameter `%s`", paramName)
		return "", errors.New("missing required parameter")
	}
	return paramValue, nil
}
