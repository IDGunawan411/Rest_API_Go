package student_post

import (
    "context"
    // "log"
    "encoding/json"
	"belajar_golang/rest_api_student4/student_post/student_create"
	"belajar_golang/rest_api_student4/utils"
    "belajar_golang/rest_api_student4/student_models"
	// "strconv"
	"net/http"
)
//post Mahasiswa
func PostStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var std student_models.Student

		if err := json.NewDecoder(r.Body).Decode(&std); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		if err := student_insert.Insert(ctx, std); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
		res := map[string]string{
			"status": "Succesfully",
		}
		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
