package faktoryandroute

import (
	uh "master/internal/handler/userhandler"
	ru "master/internal/repo/repouser"
	us "master/internal/service/userservice"
	"master/middlewares"

	lh "master/internal/handler/loginhandler"
	rl "master/internal/repo/repologin"
	ls "master/internal/service/loginservice"

	eh "master/internal/handler/employeehandler"
	re "master/internal/repo/repoemployee"
	es "master/internal/service/employservice"

	sh "master/internal/handler/salaryhandler"
	rs "master/internal/repo/reposalary"
	ss "master/internal/service/Salaryservice"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FaktoryAndRoute(e *echo.Echo, db *gorm.DB) {
	rpu := ru.NewRepoUser(db)
	userservice := us.NewServiceUser(rpu)
	userhandle := uh.NewHandleUser(userservice)
	usergrup := e.Group("/user")
	usergrup.POST("/register", userhandle.RegisterUser)
	usergrup.GET("", userhandle.AllUser)
	usergrup.GET("/profile", userhandle.Profile, middlewares.JWTMiddleware())
	// usergrup.GET("/filterlist", hndlmhs.Filter)

	rpl := rl.NewRepoLogin(db)
	servicelogin := ls.NewServiceLogin(rpl, rpu)
	handlelogin := lh.NewHandlLogin(servicelogin)
	logingrup := e.Group("/login")

	logingrup.POST("/user", handlelogin.LoginUser)
	logingrup.POST("/logout", handlelogin.LogoutUser)
	logingrup.POST("/employee", handlelogin.LoginEmp)

	rpe := re.NewRepoEmployee(db)
	empservice := es.NewServiceemployee(rpe, rpu)
	emphandle := eh.NewHandleemployee(empservice)
	empgrup := e.Group("/employee")
	empgrup.POST("/add", emphandle.AddEmployee)
	empgrup.GET("/ceksalary", emphandle.GetSalary, middlewares.JWTMiddleware())

	rps := rs.NewRepoSalary(db)
	salservice := ss.NewServiceSalary(rps, rpe)
	salhandle := sh.NewHandleSalary(salservice)
	salgrup := e.Group("/salary")
	salgrup.POST("/add", salhandle.AddSalary, middlewares.JWTMiddleware())

}
