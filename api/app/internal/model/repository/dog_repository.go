package repository

import (
	"app/internal/model/entity/orm"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type DogRepository struct {
	dbConnection *sql.DB
}

func NewDogRepository(db *sql.DB) *DogRepository {
	return &DogRepository{
		dbConnection: db,
	}
}
func (r *DogRepository) Regist(dog *orm.Dog) (int, error) {
	if r.dbConnection == nil {
		return 0, errors.New("正しく初期化されていません")
	}
	stmt, e1 := r.dbConnection.Prepare(`
INSERT INTO dog(user_id, name, birth_date, gender, personality, created_at, updated_at) 
VALUES(?,?,?,?,?,?,?)`)
	if e1 != nil {
		// エラー
		return 0, fmt.Errorf("DBのPrepareに失敗しました %s", e1.Error())
	}
	defer stmt.Close()
	res, e2 := stmt.Exec(dog.UserId, string(dog.Name), time.Time(dog.BirthDate), dog.Gender, dog.Personality, dog.CreatedAt, dog.UpdatedAt)
	if e2 != nil {
		return 0, fmt.Errorf("insertに失敗しました %s", e2.Error())
	}
	insertId, _ := res.LastInsertId()
	return int(insertId), nil
}

func (r *DogRepository) QueryByUserId(userId int) ([]*orm.Dog, error) {
	if r.dbConnection == nil {
		return nil, errors.New("正しく初期化されていません")
	}
	row, error := r.dbConnection.Query(`
SELECT 
	dog_id DogId,
	name Name,
	birth_date BirthDate,
	gender Gender,
	personality Personality
FROM dog 
WHERE user_id = ?`, userId)
	if error != nil {
		return nil, error
	}
	defer row.Close()

	array := make([]*orm.Dog, 0, 5)
	for row.Next() {
		d := orm.Dog{}
		e := row.Scan(&d.DogId, &d.Name, &d.BirthDate, &d.Gender, &d.Personality)
		if e != nil {
			return nil, e
		}
		array = append(array, &d)
	}
	if row.Err() != nil {
		return nil, row.Err()
	}

	return array, nil
}
