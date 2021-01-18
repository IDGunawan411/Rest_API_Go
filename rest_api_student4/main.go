package main
import (
	"fmt"
	"net/http"
	"log"
	"os"
	// Student
	"belajar_golang/rest_api_student4/student_get"
	"belajar_golang/rest_api_student4/student_post"
	"belajar_golang/rest_api_student4/student_delete"
	"belajar_golang/rest_api_student4/student_put"
	"belajar_golang/rest_api_student4/student_detail"
	"belajar_golang/rest_api_student4/student_src"
	"belajar_golang/rest_api_student4/student_short"
	
	// Course
	// "github.com/belajar_golang/rest_api_student/course_api/course_get"

)

func main() {
	// Student
	http.HandleFunc("/student_get", student_get.GetStudent)
	http.HandleFunc("/student_post", student_post.PostStudent)
	http.HandleFunc("/student_delete", student_delete.DeleteStudent)
	http.HandleFunc("/student_put", student_put.UpdateStudent)
	http.HandleFunc("/student_detail", student_detail.GetStudent)
	http.HandleFunc("/student_src", student_src.GetStudent)
	http.HandleFunc("/student_short", student_short.GetStudent)
	http.HandleFunc("/config", 
    func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Database User   : " + os.Getenv("db_user") + "\n");
        fmt.Fprintf(w, "Database Pass   : " + os.Getenv("db_pass") + "\n");
        fmt.Fprintf(w, "Database Server : " + os.Getenv("db_server") + "\n");
    }) 
	// Course
	// http.HandleFunc("/course_get", course_get.GetCourse)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
