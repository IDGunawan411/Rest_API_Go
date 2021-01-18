package student_update
 
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
)

// Update
func Update(ctx context.Context, std student_models.Student) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set First_Name = '%v', Last_Name ='%v', EnrollmentDate = '%v' where ID_Student = '%v'",
	table,
	std.First_Name,
	std.Last_Name,
	time.Now().Format("01/02/2006 15:04:05"),
	std.ID_Student)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}
