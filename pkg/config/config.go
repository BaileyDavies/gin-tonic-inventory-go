package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbUser     string
	dbPass     string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
	apiPort    string
	migrate    string
	redisPort  string
	redisPass  string
}

// Setup a new config instance from the config struct
func Init() *Config {
	// New object of the config struct
	conf := &Config{}

	// Reads the dbX flag from command line into the conf object variable, or uses env var as default if no flag option
	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "Username for the production DB")
	flag.StringVar(&conf.dbPass, "dbpswd", os.Getenv("POSTGRES_PASSWORD"), "Password for the production DB")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "Port for the production DB")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "Host for the production DB")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "Name of the production DB")

	flag.StringVar(&conf.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "Host of the test DB")
	flag.StringVar(&conf.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "Name of the test DB")
	flag.StringVar(&conf.apiPort, "apiPort", os.Getenv("API_PORT"), "API Port")
	flag.StringVar(&conf.migrate, "migrate", "up", "specify if we should be migrating DB 'up' or 'down'")
	flag.StringVar(&conf.redisPort, "redisport", os.Getenv("REDIS_PORT"), "Redis Cache Port")
	flag.StringVar(&conf.redisPass, "redispass", os.Getenv("REDIS_PASS"), "Redis DB Pass")

	flag.Parse()
	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBHost, c.testDBName)
}

// Private function that assmebles the correctly formatted connection string for postgres
func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPass,
		dbhost,
		c.dbPort,
		dbname)
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetMigration() string {
	return c.migrate
}

func (c *Config) GetRedisPort() string {
	return ":" + c.redisPort
}

func (c *Config) GetRedisPass() string {
	return c.redisPass
}
