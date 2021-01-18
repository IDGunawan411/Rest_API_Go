package get_student_short
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
    ODB        := r.URL.Query().Get("ODB")
    Data       := r.URL.Query().Get("Data")
    Short      := r.URL.Query().Get("Short")
    if err != nil {
        log.Fatal("Cant connect to MySQL", err)
    }
    queryText := fmt.Sprintf("SELECT * FROM %v WHERE First_Name LIKE'%%%s%%' AND Last_Name LIKE'%%%s%%' AND EnrollmentDate LIKE'%%%s%%' %s %s %s", 
    table, First_Name, Last_Name, Enrollment, ODB, Data, Short)
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
