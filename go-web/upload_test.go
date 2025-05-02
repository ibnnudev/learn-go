package goweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := templateEmbed.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	if fileHeader.Size > 10*1024*1024 {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// var allowedTypes = map[string]bool{
	// 	"image/jpeg": true,
	// 	"image/png":  true,
	// }

	// if !allowedTypes[fileHeader.Header.Get("Content-Type")] {
	// 	http.Error(w, "Invalid file type", http.StatusBadRequest)
	// 	return
	// }

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)

	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	name := r.PostFormValue("name")
	err = templateEmbed.ExecuteTemplate(w, "upload.success.gohtml", map[string]any{
		"Name":     name,
		"FilePath": "/resources/" + fileHeader.Filename,
	})
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		return
	}
}

func TestUpload(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", UploadHandler)
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatalf("Server failed to start: %v", err)
	}
}

//go:embed resources/logo.png
var uploadFileTest []byte

func TestUploadHandler(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	err := writer.WriteField("name", "test")
	if err != nil {
		panic(err)
	}
	file, _ := writer.CreateFormFile("file", "uploaded_file.png")
	file.Write(uploadFileTest)
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadHandler(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(bodyResponse))

}
