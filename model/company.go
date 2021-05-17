package model

import (
	"TwoProject/common"
	"fmt"
	"log"
)

type Company struct {
	ID       string  `description:公司ID`
	Name     string  `description:公司全称`
	NickName string  `description:公司简称`
}

func GetAllCompany() (companies []Company, err error) {
	sql := "select id,name,nickname from company"
	rows, err := common.Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next(){
		c := Company{}
		err = rows.Scan(&c.ID, &c.Name, &c.NickName)
		if err != nil {
			return
		}
		companies = append(companies, c)
	}
	return
}


func GetCompanyById(id string) (company Company, err error) {
	sql := "SELECT id, name, nickname FROM company WHERE id=?"
	err = common.Db.QueryRow(sql, id).Scan(&company.ID, &company.Name, &company.NickName)
	return
}


// Insert ...
func (company *Company) Insert2() (err error) {
	sql := "INSERT INTO company (id, name, nickname) VALUES (?, ?, ?)"
	//sql := "insert into company (id, name, nickname) VALUES ($1, $2, $3)" // 准备好sql语句
	prepare, err := common.Db.Prepare(sql) // 预处理
	if err != nil {
		return
	}
	_, err = prepare.Exec(company.ID, company.Name, company.NickName)
	if err != nil {
		return
	}
	return err
}


// Insert ...
func (company *Company) Insert() (err error) {
	sql := "INSERT INTO company (id, name, nickname) VALUES (?, ?, ?)"
	stmt, err := common.Db.Prepare(sql)
	log.Println(err)
	if err != nil {
		return
	}

	log.Println("++++" + company.ID + "+++++")

	_, err = stmt.Exec(company.ID, company.Name, company.NickName)
	if err != nil {
		return
	}
	return
}


// update

func (company *Company) Update() (err error) {
	sql := "update  company set name=?, nickname=? where id=?"
	//prepare, err := common.Db.Prepare(sql)
	//if err != nil {
	//	return
	//}
	_, err = common.Db.Exec(sql,company.Name, company.NickName, company.ID)
	//if err != nil {
	//	return
	//}
	//return
	return
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}


// delete
func DeleteCompany(id string) (err error) {
	//sql := "DELETE FROM company WHERE id=?"
	//_, err = common.Db.Exec(sql, id)
	log.Println("开始删除")
	stmt, err := common.Db.Prepare("DELETE FROM company WHERE id=?")
	check(err)
	res, err := stmt.Exec(id)
	check(err)
	num, err := res.RowsAffected()
	check(err)
	log.Println(num)

	return
}

