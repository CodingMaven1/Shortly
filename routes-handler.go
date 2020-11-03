package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/ventu-io/go-shortid"
)

// ErrorResponse to send in request.
type ErrorResponse struct {
	Code    int
	Message string
}

// RenderHome route to render the home page
func RenderHome(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./views/index.html")
}

// MakeShortURL for generating the short URL.
func MakeShortURL(res http.ResponseWriter, req *http.Request) {
	type URLObject struct {
		URL string `json:"url"`
	}

	type URLCollection struct {
		RedirectURL string
		ShortURL    string
	}

	type SuccessResponse struct {
		Code     int
		Message  string
		Response URLCollection
	}

	var urlObj URLObject
	errorobj := ErrorResponse{
		Code: http.StatusInternalServerError, Message: "There was some internal error.",
	}
	ctx := context.Background()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&urlObj)

	if err != nil {
		errorobj.Message = "The URL cannot be empty."
		respondBack(res, req, errorobj)
	} else if !validURL(urlObj.URL) {
		errorobj.Message = "Please enter a valid URL!"
		respondBack(res, req, errorobj)
	} else {
		id, error := shortid.Generate()
		if error != nil {
			respondBack(res, req, errorobj)
		} else {
			rediserr := Client.Set(ctx, id, urlObj.URL, 0).Err()
			if rediserr != nil {
				log.Println(rediserr)
				respondBack(res, req, errorobj)
			} else {
				successres := SuccessResponse{
					Code:    http.StatusOK,
					Message: "Successfully converted!",
					Response: URLCollection{
						RedirectURL: urlObj.URL,
						ShortURL:    req.Host + "/" + id,
					},
				}
				jsonResponse, err := json.Marshal(successres)
				if err != nil {
					panic(err)
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(successres.Code)
				res.Write(jsonResponse)
			}
		}
	}
}

func validURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func respondBack(res http.ResponseWriter, req *http.Request, errormsg ErrorResponse) {
	// httpError := &ErrorResponse{Code: errormsg.Code, Message: errormsg.Message}
	jsonResponse, err := json.Marshal(errormsg)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(errormsg.Code)
	res.Write(jsonResponse)
}
