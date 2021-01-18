package student_detail
import (
    "context"
    "fmt"
    // "encoding/json"
	"belajar_golang/rest_api_student4/student_detail/get_student_detail"
	"belajar_golang/rest_api_student4/student_models"
	"belajar_golang/rest_api_student4/utils"
    // "github.com/belajar_golang/rest_api/models"
	// "strconv"
	"net/http"
)
// GetMahasiswa
func GetStudent(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ctx, cancel := context.WithCancel(context.Background())
		defer cancel() 
		var student student_models.Student
		id := r.URL.Query().Get("id")
		student.ID_Student = id
        students, err := get_student_detail.GetAll(ctx, student)
        if err != nil {
            fmt.Println(err)
        }
        utils.ResponseJSON(w, students, http.StatusOK)
        return
    }
    http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
    return
}
