package conn

import (
	"database/sql"
	"fmt"
	"log"
	"main/config"

	_ "github.com/go-sql-driver/mysql"
)

type DBcp struct {
	Dsn  string
	Conn *sql.DB
}

type pools struct {
	Pools []*DBcp
}

var cPools pools

// TODO CONFIG BASED
func InitDB() {
	conf := config.GetConfig()

	for _, dbConf := range conf.DB {
		connect(dbConf)
	}
}

func connect(dbConf config.DBConfig) DBcp {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//연결 확인
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	newDBcp := DBcp{Conn: db, Dsn: dsn}
	cPools.Pools = append(cPools.Pools, &newDBcp)

	log.Println("new DB connection")
	return newDBcp
}

func (d *DBcp) disconenct() {
	d.Conn.Close()
}

func GetPools() pools {
	return cPools
}

func DisconnectAll() {
	for _, dbcp := range cPools.Pools {
		dbcp.disconenct()
		log.Println(dbcp.Dsn, " disconencted ")
	}
}
