package get_student_detail
 
import (
    "context"
    "fmt"
    // "errors"
    "belajar_golang/rest_api_student4/config"
    "belajar_golang/rest_api_student4/student_models"
    "log"
)
const (
    table          = "student"
    // layoutDateTime = "2006-01-02 15:04:05"
)
// GetAll
func GetAll(ctx context.Context, student student_models.Student) ([]student_models.Student, error) {
    var students []student_models.Student
    db, err := config.MySQL()
    
    if err != nil {
        log.Fatal("Cant connect to MySQL", err)
    }
    
    queryText := fmt.Sprintf("SELECT * FROM %v WHERE ID_Student='%s'", table, student.ID_Student)
    rowQuery, err := db.QueryContext(ctx, queryText)
    if err != nil {
        log.Fatal(err)
    }
    for rowQuery.Next() {
        var student student_models.Student
        // var createdAt, updatedAt string
        if err = rowQuery.Scan(
            &student.ID_Student,
            &student.First_Name,
            &student.Last_Name,
            &student.EnrollmentDate);
            // &createdAt,
            // &updatedAt); 
            
            err != nil {
                return nil, err
            }
        //  Change format string to datetime for created_at and updated_at
        // mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
        // if err != nil {
        //     log.Fatal(err)
        // }
        // mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
        // if err != nil {
        //     log.Fatal(err)
        // }
        students = append(students, student)
    }
    return students, nil
} 
