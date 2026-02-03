package dto

// EmployeeDTO untuk request/response
// Sesuaikan dengan field Employee

type EmployeeDTO struct {
	EmployeeID   int32   `json:"employee_id"`
	FirstName    *string `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	HireDate     string  `json:"hire_date"`
	JobID        int32   `json:"job_id"`
	Salary       float64 `json:"salary"`
	ManagerID    *int32  `json:"manager_id"`
	DepartmentID *int32  `json:"department_id"`
	PhotoProfile *string `json:"photo_profile"`
}
