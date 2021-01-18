package student_models
 
import (
    // "time"
)
 
type (
    Student struct {
        ID_Student  string   `json:"ID_Student"`
        First_Name  string   `json:"First_Name"`
        Last_Name   string   `json:"Last_Name"`
        EnrollmentDate string  `json:"EnrollmentDate"`
    }
)
 