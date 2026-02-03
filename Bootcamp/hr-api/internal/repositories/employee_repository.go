package repositories

import (
	"context"

	"github.com/danielaugust67/hr-api/internal/domain/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) FindAll(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	if err := r.DB.WithContext(ctx).Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) FindByID(ctx context.Context, id int32) (*models.Employee, error) {
	var employee models.Employee
	if err := r.DB.WithContext(ctx).First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) Create(ctx context.Context, employee *models.Employee) error {
	return r.DB.WithContext(ctx).Create(employee).Error
}

func (r *EmployeeRepository) Update(ctx context.Context, employee *models.Employee) error {
	return r.DB.WithContext(ctx).Save(employee).Error
}

func (r *EmployeeRepository) Delete(ctx context.Context, id int32) error {
	return r.DB.WithContext(ctx).Delete(&models.Employee{}, id).Error
}
