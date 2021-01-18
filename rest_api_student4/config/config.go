package config
 
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "os"
)
 
// HubToMySQL
func MySQL() (*sql.DB, error) {
    db, err := sql.Open("mysql", ""+ os.Getenv("db_user") +":"+ os.Getenv("db_pass") +"@tcp("+ os.Getenv("db_server") +")/rest_student")
    
    if err != nil {
        return nil, err
    }
 
    return db, nil
}