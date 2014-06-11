package MySQLdb

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySqlDB struct {
	Address, DbName, User, Password string

	DbConn *sql.DB
	DbErr error
}

func (v *MySqlDB) Connect(){
	v.DbConn,v.DbErr = OpenDB(fmt.Sprintf("tcp(%s)",v.Address),v.DbName,v.User,v.Password)
	if v.DbErr != nil {
		panic(v.DbErr)
	}
}

func (v *MySqlDB) Close() error {
	err := v.DbConn.Close()
	return err
}

func OpenDB(DB_HOST, DB_NAME,DB_USER, DB_PASS string) (*sql.DB,error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8", DB_USER, DB_PASS, DB_HOST,DB_NAME))
	return db,err
}

func NullInt64ToInt64(v sql.NullInt64) int64{
	if v.Valid{
		return v.Int64
	}
	return 0
}

func NullInt64ToInt32(v sql.NullInt64) uint32{
	if v.Valid{
		return (uint32)(v.Int64)
	}
	return 0
}

func NullStringToString(v sql.NullString) string{
	if v.Valid{
		return v.String
	}
	return ""
}

func (v *MySqlDB) GetOneField(query, field string) (interface {},error) {
	data,err := v.GetArray(query)
	if err != nil {
		return nil,err
	}
	if len(data) == 0{
		return nil, sql.ErrNoRows
	}
	d,ok := data[0][field]
	if !ok {
		return nil,sql.ErrNoRows
	}
	return d,nil
}

func (v * MySqlDB) GetOne(query string) (map[string]interface {},error){
	data,err := v.GetArray(query)
	if err != nil {
		return nil, err
	}
	if len(data) == 0{
		return nil, sql.ErrNoRows
	}
	return data[0], nil
}

func (v * MySqlDB) GetArray(query string) ([]map[string]interface {},error) {
	//	return v.DbConn.Query(query)
	records := make([]map[string]interface {},0)
	rows, err := v.DbConn.Query(query)
	if err != nil {
		return nil,err
	}
	columns, err := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				default:
				case bool:
					record[columns[i]] = col.(bool)
				case int:
					record[columns[i]] = col.(int)
				case int64:
					record[columns[i]] = col.(int64)
				case float64:
					record[columns[i]] = col.(float64)
				case string:
					record[columns[i]] = col.(string)
				case []byte:   // -- all cases go HERE!
					record[columns[i]] = string(col.([]byte))
				case time.Time:
					record[columns[i]] = col.(string)
				}
			}
		}
		records = append(records,record)
	}
	return records,nil
}

func (v * MySqlDB) Execute(query string) (sql.Result, error){
	return v.DbConn.Exec(query)
}
