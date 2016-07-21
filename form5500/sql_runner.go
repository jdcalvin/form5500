package main

import (
  "database/sql"
  "log"
  "fmt"
  "os/exec"
)

var db *sql.DB
var dbConnection string
var err error

// SetDBConnection set connection string to be used by sql.DB Open()
func SetDBConnection(connection string) {
  dbConnection = connection
}

// OpenDBConnection opens db connection - call at top of function
func OpenDBConnection() {
  if dbConnection == "" {
    log.Fatal("SQLRunner.Open() Must set connection with SetConnection(connection string)")
  }
  db, err = sql.Open("postgres", dbConnection)

  if err != nil {
    log.Fatal(err)
  }
}

// CloseDBConnection closes opened db connection - defer at top of function
func CloseDBConnection() {
  db.Close()
}

// SQLRunner struct to assign and call sql statements throughout the form5500 package
type SQLRunner struct {
  sql  string
  description string
}

// Exec runs #sql statement and prints description to command line
func (s SQLRunner) Exec() {
  s.Print()
  _, err := db.Exec(s.sql)
  if err != nil {
    fmt.Println(s)
    log.Fatal(err)
  }
}
// ExecCopy uses psql command line tool to copy data from a csv file
// Cannot use Exec due to permissions error on aws box
func (s SQLRunner) ExecCopy() {
  s.Print()
	cmd := exec.Command("psql", dbConnection, "-c", s.sql)
	_, err := cmd.Output()
	if err != nil {
    fmt.Println("psql \"" + dbConnection + "\" -c \"" + s.sql + "\"")
	  log.Fatal(err)
  }
}

// Print print formatted message to console
func (s SQLRunner) Print() {
  fmt.Println(fmt.Sprintf(" - %s", s.description))
}
