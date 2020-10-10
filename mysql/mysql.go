package mysql

import (
	"database/sql"
	"fmt"
	mySQLDriver "github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"

)

const (
	maxAllowedPacketSize 	= 16777216		// 16MB
	configFilePath = "/config.toml"
)

type MySQLConnector struct {
	MySQLUrl string		`toml:"mySQLUrl"`
	Database string		`toml:"mySQLDatabase"`
	Username string		`toml:"mySQLUsername"`
	Password string		`toml:"mySQLPassword"`
	Network  string		`toml:"mySQLNetwork"`
}

type mySQLConfigParsed struct {
	MySQLConnector	MySQLConnector		`toml:"MySQLConnector"`
}

// mySQLStatus is mySQL active status flag
var mySQLStatus bool

// mDB is mySQL object pointer
var mySQLObject *sql.DB


// GetMySQL initializes MySQL object and return its pointer
func GetMySQL() *sql.DB {

	if mySQLStatus == false {
		mySQLObject = Init()
		if mySQLObject == nil {
			fmt.Println("Unable to intialize mySQL: mySQL: GetMySQL")
		} else {
			fmt.Println("Successfully connected to mySQL")
			mySQLStatus = true
		}
	}
	return mySQLObject
}

func Init() *sql.DB {
	
	var conf mySQLConfigParsed 
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		fmt.Println("Error in mySQL: Init: toml.DecodeFile: ", err)
		return nil
	}

	// Setting up the mySQL connection
	var mySQLConfig mySQLDriver.Config
	mySQLConfig.User = conf.MySQLConnector.Username
	mySQLConfig.Passwd = conf.MySQLConnector.Password
	mySQLConfig.Net = conf.MySQLConnector.Network
	mySQLConfig.Addr = conf.MySQLConnector.MySQLUrl
	mySQLConfig.DBName = conf.MySQLConnector.Database
	mySQLConfig.MaxAllowedPacket = maxAllowedPacketSize
	mySQLConfig.AllowNativePasswords=true

	db, err := sql.Open("mysql", mySQLConfig.FormatDSN())
	if err != nil {
		fmt.Println("Failed to connect to mySQL server :",err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Unable to ping the mySQL sever :",err)
		return nil
	}

	return db
}

