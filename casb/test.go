package casb

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Enforcers *casbin.Enforcer
)

func C()  {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, _ := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.
	Enforcers, _ = casbin.NewEnforcer("casb/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	_ = Enforcers.LoadPolicy()

	// Check the permission.
	//_, _ = Enforcers.Enforce("alice", "data1", "read")

	// Modify the policy.
	//Enforcers.AddPolicy("test1","root")
	//Enforcers.RemovePolicy("test1","root")

	// Save the policy back to DB.
	//_ = Enforcers.SavePolicy()
}


