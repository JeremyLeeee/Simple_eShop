package repositories

import (
	"database/sql"
	"errors"
	"eshop/common"
	"eshop/datamodels"
	"log"
	"strconv"
)

type IUser interface {
	Conn() error
	Select(userName string) (*datamodels.User, error)
	Insert(user *datamodels.User) (int64, error)
}

type UserManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewUserRepository(table string, db *sql.DB) IUser {
	return &UserManager{table: table, mysqlConn: db}
}

func (u *UserManager) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "order"
	}
	return
}

func (u *UserManager) Select(userName string) (user *datamodels.User, err error) {
	if userName == "" {
		return &datamodels.User{}, errors.New("user name is null")
	}
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}

	sql := "SELECT * FROM " + u.table + " WHERE userName=?"
	rows, err := u.mysqlConn.Query(sql, userName)

	if err != nil {
		return &datamodels.User{}, err
	}
	result := common.GetResultRow(rows)

	if len(result) == 0 {
		log.Println("user not exist")
		return &datamodels.User{}, errors.New("user not exist")
	}

	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return

}

func (u *UserManager) Insert(user *datamodels.User) (id int64, err error) {
	if err = u.Conn(); err != nil {
		return 0, err
	}

	sql := "INSERT " + u.table + " SET nickName=?, userName=?, passWord=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (u *UserManager) SelectById(id int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	sql := "SELECT * FROM " + u.table + " WHERE ID=" + strconv.FormatInt(id, 10)
	rows, err := u.mysqlConn.Query(sql)
	if err != nil {
		return &datamodels.User{}, err
	}
	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("user not exist")
	}
	user = &datamodels.User{}
	common.DataToStructByTagSql(result, user)
	return
}
