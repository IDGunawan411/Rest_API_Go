package student_get
import (
    "context"
    "fmt"
    // "encoding/json"
	"belajar_golang/rest_api_student4/student_get/student"
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
        students, err := student.GetAll(ctx)
        if err != nil {
            fmt.Println(err)
        }
        utils.ResponseJSON(w, students, http.StatusOK)
        return
    }
    http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
    return
}
