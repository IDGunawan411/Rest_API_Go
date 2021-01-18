package student_put

import (
    "context"
    // "log"
    "encoding/json"
	"belajar_golang/rest_api_student4/student_put/student_update"
	"belajar_golang/rest_api_student4/utils"
    "belajar_golang/rest_api_student4/student_models"
	// "strconv"
	"net/http"
	"fmt"
)
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type aplication / json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var std student_models.Student
		if err := json.NewDecoder(r.Body).Decode(&std); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		fmt.Println(std)
		if err := student_update.Update(ctx, std); err != nil {
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