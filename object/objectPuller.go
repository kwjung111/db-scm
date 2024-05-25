package object

import (
	"database/sql"
	"fmt"
	"main/conn"
)

type objectPuller struct {
}

func GetAllTriggers(dbcp conn.DBcp) ([]trigger, error) {
	query := getAllTriggersQuery()
	return getTriggers(dbcp.Conn, query)
}

func GetChangedTriggers(dbcp conn.DBcp) ([]trigger, error) {
	query := getChangedTriggersQuery()
	return getTriggers(dbcp.Conn, query)
}

func GetAllProcedures(dbcp conn.DBcp) ([]procedure, error) {
	query := getAllProceduresQuery()
	return getProcedures(dbcp.Conn, query)
}

func GetChangedProcedures(dbcp conn.DBcp) ([]procedure, error) {
	query := getChangedProceduresQuery()
	return getProcedures(dbcp.Conn, query)
}

func getTriggers(db *sql.DB, query string) ([]trigger, error) {

	var trgArr []trigger

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("sql error : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var trg trigger
		if err := rows.Scan(&trg.Db, &trg.Name, &trg.Def); err != nil {
			return nil, fmt.Errorf("parse error : %w", err)
		}
		trgArr = append(trgArr, trg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error : %w", err)
	}
	return trgArr, nil
}

func getProcedures(db *sql.DB, query string) ([]procedure, error) {

	var prcArr []procedure

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("sql error : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var prc procedure
		if err := rows.Scan(&prc.Db, &prc.Name, &prc.Def); err != nil {
			return nil, fmt.Errorf("parse error : %w", err)
		}
		prcArr = append(prcArr, prc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error : %w", err)
	}
	return prcArr, nil
}
