package student_short

import (
    "context"
    "fmt"
	"belajar_golang/rest_api_student4/student_short/get_student_short"
	"belajar_golang/rest_api_student4/utils"
	"net/http"
)
// GetMahasiswa
func GetStudent(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ctx, cancel := context.WithCancel(context.Background())
		defer cancel() 
        students, err := get_student_short.GetAll(ctx, r)
        if err != nil {
            fmt.Println(err)
        }
        utils.ResponseJSON(w, students, http.StatusOK)
        return
    }
    http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
    return
}