package main

package main

/*
create a folder interface-prob , create a go module in root of the project
create main.go

	create a package named as data
	Inside the data package create stores/postgres , stores/mysql
		postgres -> // create a file named as postgres.go
			// struct -> conn to db (*sql.Db)
			// Create, Update, Delete
		mysql -> // repeat the same thing
	Inside  data package create an interface which could be implemented by postgres and mysql
	Create a global variable of the type interface , Export the global variable

In main.go call the Create method of mysql and postgres using the interface
*/
