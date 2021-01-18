package get_student_src
import (
    "context"
    "fmt"
    "belajar_golang/rest_api_student4/config"
    "belajar_golang/rest_api_student4/student_models"
    "log"
    "net/http"
)
const (
    table          = "student"
)
// GetAll
func GetAll(ctx context.Context, r *http.Request) ([]student_models.Student, error) {
    var students []student_models.Student
    db, err := config.MySQL()
    First_Name := r.URL.Query().Get("First_Name")
    Last_Name  := r.URL.Query().Get("Last_Name")
    Enrollment := r.URL.Query().Get("Enrollment")
    if err != nil {
        log.Fatal("Cant connect to MySQL", err)
    }
    queryText := fmt.Sprintf("SELECT * FROM %v WHERE First_Name LIKE'%%%s%%' AND Last_Name LIKE'%%%s%%' AND EnrollmentDate LIKE'%%%s%%'", 
    table, First_Name, Last_Name, Enrollment)
    rowQuery, err := db.QueryContext(ctx, queryText)
    if err != nil {
        log.Fatal(err)
    }
    for rowQuery.Next() {
        var student student_models.Student
        if err = rowQuery.Scan(
            &student.ID_Student,
            &student.First_Name,
            &student.Last_Name,
            &student.EnrollmentDate);
            err != nil {
                return nil, err
            }
        students = append(students, student)
    }
    return students, nil
} 
