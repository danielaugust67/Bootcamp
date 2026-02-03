package routes

import (
	"net/http"

	"github.com/danielaugust67/hr-api/internal/handlers"
	"github.com/danielaugust67/hr-api/internal/repositories"
	"github.com/danielaugust67/hr-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	//5. Initialize repositories
	regionRepo := repositories.NewRegionRepository(db)
	departmentRepo := repositories.NewDepartmentRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)

	//6. Initialize services
	regionService := services.NewRegionService(regionRepo)
	departmentService := services.NewDepartmentService(departmentRepo)
	employeeService := services.NewEmployeeService(employeeRepo)

	//7. Initialize handlers
	regionHandler := handlers.NewRegionHandler(regionService)
	depatmentHandler := handlers.NewDepartmentHandler(departmentService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)

	//3.1 call handler
	router.GET("/", welcomeHandler)

	//9.call basepath
	basePath := viper.GetString("SERVER.BASE_PATH")

	//8. grouping subroute with prefix /api
	api := router.Group(basePath)
	{
		// region routes endpoints
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.GET("/countries", regionHandler.GetRegionsWithCountry)
			regions.GET("/:id/countries", regionHandler.GetRegionByIdWithCountry)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
		}

		department := api.Group("/departments")
		{
			department.POST("", depatmentHandler.CreateDepartment)
			department.GET("/:id", depatmentHandler.GetDepartmentByID)
			department.GET("", depatmentHandler.GetAllDepartments)
			department.PUT("/:id", depatmentHandler.UpdateDepartment)
			department.DELETE("/:id", depatmentHandler.DeleteDepartment)
			department.GET("/search", depatmentHandler.SearchDepartments) // filter search
		}

		employees := api.Group("/employees")
		{
			employees.GET("", employeeHandler.GetAll)
			employees.GET("/:id", employeeHandler.GetByID)
			employees.POST("", employeeHandler.Create)
			employees.PUT("/:id", employeeHandler.Update)
			employees.DELETE("/:id", employeeHandler.Delete)
		}
	}

}

// 2. create first handler
func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to gin framework",
		"status":  "running",
	})
}
