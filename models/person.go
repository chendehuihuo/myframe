package models

import (
	db "myframe/database"
	"strconv"
)

type Name struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Time int    `json:"time" form:"time"`
}

func (t *Name) TableName() string {
	return "name"
}

func (p *Name) AddPerson() (id int64, err error) {
	var name = Name{Name: p.Name, Time:p.Time}
	err=db.SqlDB.Create(&name).Error
	if err != nil {
		return
	}
	id,err= strconv.ParseInt(strconv.Itoa(name.Id), 10, 64)
	if err != nil {
		return
	}
	return
}

func (p *Name) GetPersons() (persons []Name, err error) {
	persons = make([]Name, 0)

	//rows,err:=db.SqlDB.Where("name = ?", "李钢").First(&user);
	err=db.SqlDB.Find(&persons).Error
	//rows, err := db.SqlDB.Query("SELECT `name`,`time` FROM ceshi limit 10000")

	if err != nil {
		return
	}

	//for name {
	//	var person Name
	//	rows.Scan(&person.Name, &person.Time)
	//	persons = append(persons, person)
	//}
	//if err = rows.Err(); err != nil {
	//	return
	//}
	return
}

func (p *Name) GetPersonsone() (persons []Name) {
	persons = make([]Name, 0)
	err:=db.SqlDB.Where("id=?",p.Id).First(&persons).Error
	if err !=nil{
		return
	}


	//var person Person
	//persons = make([]Name, 0)
	//errs := db.SqlDB.QueryRow("SELECT  `name`,`time` FROM `name` where  name=? limit 1", p.Name).Scan(
	//	&person.Name, &person.Time,
	//)
	//if errs != nil {
	//	return
	//}
	//persons = append(persons, person)
	return
}

func (p *Name) PutPersonsone() (id int64, err error) {
	var name Name
	name.Name = p.Name
	name.Time = p.Time
	err= db.SqlDB.Model(&name).Where("id = ?", p.Id).Updates(name).Error
	if err != nil {
		return
	}

	id,err= strconv.ParseInt(strconv.Itoa(p.Id), 10, 64)

	if err != nil {
		return
	}

	//stmt, err := db.SqlDB.Prepare("UPDATE name SET  name=? WHERE time=?")
	//defer stmt.Close()
	//if err != nil {
	//	return
	//}
	//rs, err := stmt.Exec(p.Name, p.Time)
	//if err != nil {
	//	return
	//}
	//ra, err = rs.RowsAffected()
	//if err != nil {
	//	return
	//}
	return
}

func (p *Name) DeletePersonone() (id int64,err error) {
	var name Name
	err=db.SqlDB.Where("id = ?", p.Id).Delete(&name).Error
	if err != nil {
			return
	}
	id,err= strconv.ParseInt(strconv.Itoa(p.Id), 10, 64)

	if err != nil {
		return
	}
	//rs, err := db.SqlDB.Exec("DELETE FROM name WHERE time=?", p.Time)
	//if err != nil {
	//	return
	//}
	//ra, err = rs.RowsAffected()
	//if err != nil {
	//	return
	//}
	return
}
