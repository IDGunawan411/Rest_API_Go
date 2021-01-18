package student_delete
import (
    "context"
    "fmt"
    // "encoding/json"
	"belajar_golang/rest_api_student4/student_delete/student_del"
	"belajar_golang/rest_api_student4/utils"
    "belajar_golang/rest_api_student4/student_models"
	// "log"
	// "strconv"
	"net/http"
)
// DeleteMahasiswa  
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var student student_models.Student
		id := r.URL.Query().Get("id")
		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong bos", http.StatusBadRequest)
			return
		}
		student.ID_Student = id
		if err := student_del.Delete(ctx, student); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}
		res := map[string]string{
			"status": "Succesfully",
		}
		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
