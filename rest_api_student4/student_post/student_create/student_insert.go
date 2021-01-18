package student_insert
 
import (
    "context"
    "fmt"
    // "errors"
    "belajar_golang/rest_api_student4/config"
    "belajar_golang/rest_api_student4/student_models"
    "log"
	"time"
)
 
const (
    table          = "student"
    // layoutDateTime = "2006-01-02 15:04:05"
)

// Insert
func Insert(ctx context.Context, std student_models.Student) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (ID_Student, First_Name, Last_Name, EnrollmentDate) values('%v','%v','%v','%v')", 
		table,
		std.ID_Student,
		std.First_Name,
		std.Last_Name,
		time.Now().Format("01/02/2006 15:04:05"),
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
