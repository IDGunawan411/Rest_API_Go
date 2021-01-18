package student_del

import (
    "context"
    "fmt"
    "errors"
    "belajar_golang/rest_api_student4/config"
    "belajar_golang/rest_api_student4/student_models"
    "log"
    // "time"
)
const (
    table          = "student"
    layoutDateTime = "2006-01-02 15:04:05"
)
// Delete
func Delete(ctx context.Context, student student_models.Student) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}
	queryText := fmt.Sprintf("DELETE FROM %v where ID_Student='%s'", table, student.ID_Student)
	s, err := db.ExecContext(ctx, queryText)
    if err != nil {
		return err
	}
	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}
	return nil
}