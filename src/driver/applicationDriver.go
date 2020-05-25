package driver

import (
	"fmt"
	"golangtuts/interfacez"
	"golangtuts/model"
	"golangtuts/utils"
)

/*
 * CHANNEL 1
 */
func CacheExamples(isDone chan string) {
	var cache = utils.MemCache{
		Cache: make(map[string]string),
	}
	cache.AddToCache("Anmol", "Deep")
	cache.AddToCache("Nitika", "Qazi")
	cache.AddToCache("Anila", "Koul")
	fmt.Println(cache.Get("Anila"))
	cache.AddToCache("Another", "Value")
	cache.Remove("Nitika")
	fmt.Println(cache.KeyExists("JUNK"))
	isDone <- "Finished CacheExample"

}

/*
 * CHANNEL 2
 */
func GoCompositions(isDone chan string) {
	var emp interfacez.Employee = utils.CreateEmployee("Operations")
	var achv interfacez.Achievements = model.WorkAchievements{
		Name:     "Employee Of The Month",
		Place:    "Bangalore",
		Year:     2019,
		Employee: emp,
	}
	achv.RecordAnAchievement()
	achv.FetchAllAchievements()
	achv.EmployeeDetails()
	isDone <- "Finished GoCompositionsExample"
}

/*
 * CHANNEL 3
 */
func DealWithEmployees(isDone chan string) {
	var employees []interfacez.Employee = utils.CreateEmployees()
	utils.CreateAndDescribeEmployee("HR")
	fmt.Println(employees)
	isDone <- "Finished DealWithEmployeesExample"
}
