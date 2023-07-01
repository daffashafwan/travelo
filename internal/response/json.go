package response

import (
	"encoding/json"
	"net/http"
	"travelo/internal/dto"
	"strconv"
)

func JSON(w http.ResponseWriter, status int, data any) error {
	dataJson := dto.BaseResponse{
		Data: data,
	}
	return JSONWithHeaders(w, status, dataJson, nil)
}

func JSONCustom(w http.ResponseWriter, data any, err error, status int) error {	
	
	message := "Successful"
	if err != nil {
		message = err.Error()
	}

	dataJson := dto.BaseResponse{
		Status:  strconv.Itoa(status) + " " + http.StatusText(status),
		Data:    data,
		Message: message,
	}
	return JSONWithHeaders(w, status, dataJson, nil)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
