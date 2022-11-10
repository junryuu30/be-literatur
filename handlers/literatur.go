package handlers

import (
	"encoding/json"
	literaturdto "literature/dto/literatur"
	dto "literature/dto/result"
	"literature/models"
	"literature/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerLiteratur struct {
	LiteraturRepository repositories.LiteraturRepository
}

func HandlerLiteratur(LiteraturRepository repositories.LiteraturRepository) *handlerLiteratur {
	return &handlerLiteratur{LiteraturRepository}
}

// func (h *handlerLiteratur) CreateLiteratur(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	UserId := int(userInfo["user_id"].(float64))

// 	dataContex := r.Context().Value("dataFile")
// 	filename := dataContex.(string)

// 	request := literaturdto.CreateLiteratureRequest{
// 		Title:           r.FormValue("title"),
// 		PublicationDate: r.FormValue("publicationdate"),
// 		Pages:           r.FormValue("Pages"),
// 		ISBN:            r.FormValue("isbn"),
// 		Author:          r.FormValue("author"),
// 		Attache:         r.FormValue("attache"),
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	literatur := models.Literatur{
// 		UserID:          UserId,
// 		Title:           request.Title,
// 		PublicationDate: request.PublicationDate,
// 		Pages:           request.Pages,
// 		ISBN:            request.ISBN,
// 		Author:          request.Author,
// 		Attache:         filename,
// 	}

// 	data, err := h.LiteraturRepository.CreateLiteratur(literatur)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerLiteratur) CreateLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pdfContext := r.Context().Value("dataPDF")
	filename := pdfContext.(string)

	request := literaturdto.CreateLiteratureRequest{
		Title:           r.FormValue("title"),
		PublicationDate: r.FormValue("publicationdate"),
		Pages:           r.FormValue("Pages"),
		ISBN:            r.FormValue("isbn"),
		Author:          r.FormValue("author"),
		Attache:         r.FormValue("attache"),
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// resp, err := cld.Upload.Upload(ctx, uploader.UploadParams{Folder: "goLiteratur"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	validation := validator.New()

	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	literatur := models.Literatur{
		// UserID: UserID,
		Title:           request.Title,
		PublicationDate: request.PublicationDate,
		Pages:           request.Pages,
		ISBN:            request.ISBN,
		Author:          request.Author,
		Attache:         filename,
	}

	// literatur, err = h.LiteraturRepository.CreateLiteratur(literatur)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// literatur, err = h.LiteraturRepository.GetLiteratur(literatur.ID)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }
	// literatur, _ = h.LiteraturRepository.GetLiteratur(literatur.ID)
	// literatur.Attache = os.Getenv("PATH_FILE") + literatur.Attache

	// w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: http.StatusOK, Data: literatur}
	// json.NewEncoder(w).Encode(response)

	literatur, _ = h.LiteraturRepository.CreateLiteratur(literatur)

	literatur, _ = h.LiteraturRepository.GetLiteratur(literatur.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literatur}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerLiteratur) FindLiteraturs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	literaturs, err := h.LiteraturRepository.FindLiteraturs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range literaturs {
		literaturs[i].Attache = os.Getenv("PATH_FILE") + p.Attache
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literaturs}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerLiteratur) GetLiteratur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	literatur, err := h.LiteraturRepository.GetLiteratur(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	literatur.Attache = os.Getenv("PATH_FILE") + literatur.Attache

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: literatur}
	json.NewEncoder(w).Encode(response)

}
