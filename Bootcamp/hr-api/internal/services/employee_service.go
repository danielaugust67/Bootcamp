package services

import (
	"context"
	"time"

	"github.com/danielaugust67/hr-api/internal/domain/models"
	"github.com/danielaugust67/hr-api/internal/dto"
	"github.com/danielaugust67/hr-api/internal/repositories"
)

type EmployeeService struct {
	Repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) GetAll(ctx context.Context) ([]dto.EmployeeDTO, error) {
	employees, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var result []dto.EmployeeDTO
	for _, e := range employees {
		result = append(result, toEmployeeDTO(&e))
	}
	return result, nil
}

func (s *EmployeeService) GetByID(ctx context.Context, id int32) (*dto.EmployeeDTO, error) {
	e, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	dto := toEmployeeDTO(e)
	return &dto, nil
}

func (s *EmployeeService) Create(ctx context.Context, req *dto.EmployeeDTO) error {
	employee := fromEmployeeDTO(req)
	return s.Repo.Create(ctx, &employee)
}

func (s *EmployeeService) Update(ctx context.Context, req *dto.EmployeeDTO) error {
	employee := fromEmployeeDTO(req)
	return s.Repo.Update(ctx, &employee)
}

func (s *EmployeeService) Delete(ctx context.Context, id int32) error {
	return s.Repo.Delete(ctx, id)
}

func toEmployeeDTO(e *models.Employee) dto.EmployeeDTO {
	return dto.EmployeeDTO{
		EmployeeID:   e.EmployeeID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		Email:        e.Email,
		PhoneNumber:  e.PhoneNumber,
		HireDate:     e.HireDate.Format("2006-01-02"),
		JobID:        e.JobID,
		Salary:       e.Salary,
		ManagerID:    e.ManagerID,
		DepartmentID: e.DepartmentID,
		PhotoProfile: e.PhotoProfile,
	}
}

func fromEmployeeDTO(dto *dto.EmployeeDTO) models.Employee {
	return models.Employee{
		EmployeeID:   dto.EmployeeID,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Email:        dto.Email,
		PhoneNumber:  dto.PhoneNumber,
		HireDate:     parseDate(dto.HireDate),
		JobID:        dto.JobID,
		Salary:       dto.Salary,
		ManagerID:    dto.ManagerID,
		DepartmentID: dto.DepartmentID,
		PhotoProfile: dto.PhotoProfile,
	}
}

func parseDate(dateStr string) (t time.Time) {
	t, _ = time.Parse("2006-01-02", dateStr)
	return
}
