package object

func getAllTriggersQuery() string {
	return `SELECT 
	ROUTINE_SCHEMA AS db
	,ROUTINE_NAME AS name
	,ROUTINE_DEFINITION AS def
	FROM information_schema.routines
	WHERE routine_type = 'procedure'
	AND routine_schema != 'sys'
	AND routine_schema != 'mysql';`
}

func getChangedTriggersQuery() string {
	return `SELECT 
	ROUTINE_SCHEMA AS db
	,ROUTINE_NAME as name
	,ROUTINE_DEFINITION AS def
	FROM information_schema.routines
	WHERE routine_type = 'procedure'
	AND routine_schema != 'sys'
	AND routine_schema != 'mysql'
	AND LAST_ALTERED >= DATE_SUB(CURDATE(), INTERVAL 1 DAY) `
}

func getAllProceduresQuery() string {
	return `
	select 
	ROUTINE_SCHEMA AS db
	,ROUTINE_NAME AS name
	,ROUTINE_DEFINITION AS def
	FROM information_schema.routines
	WHERE routine_type = 'procedure'
	AND routine_schema != 'sys'
	AND routine_schema != 'mysql';`
}

func getChangedProceduresQuery() string {
	return `SELECT 
	ROUTINE_SCHEMA AS db
	,ROUTINE_NAME as name
	,ROUTINE_DEFINITION AS def
	FROM information_schema.routines
	WHERE routine_type = 'procedure'
	AND routine_schema != 'sys'
	AND routine_schema != 'mysql'
		AND LAST_ALTERED >= DATE_SUB(CURDATE(), INTERVAL 1 DAY) `
}
