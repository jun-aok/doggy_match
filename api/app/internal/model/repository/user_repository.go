package repository

import (
	"app/internal/model/entity/orm"
	"app/pkg/util"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserRepository struct {
	dbConnection *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		dbConnection: db,
	}
}
func (r UserRepository) Regist(user *orm.User) (int, error) {
	if r.dbConnection == nil {
		return 0, errors.New("正しく初期化されていません")
	}
	stmt, e1 := r.dbConnection.Prepare(`
INSERT INTO user(user_code, email, name, birth_date, gender, created_at, updated_at) 
VALUES(?,?,?,?,?,?,?)`)
	if e1 != nil {
		// エラー
		return 0, fmt.Errorf("DBのPrepareに失敗しました %s", e1.Error())
	}
	defer stmt.Close()
	res, e2 := stmt.Exec(user.UserCode, user.Email, string(user.Name), time.Time(user.BirthDate), user.Gender, user.CreatedAt, user.UpdatedAt)
	if e2 != nil {
		return 0, fmt.Errorf("insertに失敗しました %s", e2.Error())
	}
	insertId, _ := res.LastInsertId()
	return int(insertId), nil
}
func (r UserRepository) Query(userId int) (*orm.User, error) {
	if r.dbConnection == nil {
		// 実行時エラーになるのでdbConnectionを必ず渡す制約をつける方法はないのか、、
		return nil, errors.New("正しく初期化されていません")
	}
	row := r.dbConnection.QueryRow(`
SELECT 
	user_id UserId,
	user_code UserCode,
	email Email,
	name Name,
	birth_date BirthDate,
	gender Gender,
	created_at CreatedAt,
	updated_at UpdatedAt,
	deleted_at DeletedAt
FROM user 
WHERE user_id = ?`, userId)
	user, error := r.fetchUser(row)
	if error != nil {
		return nil, error
	}
	return user, nil
}
func (r UserRepository) QueryByUserCode(userCode string) (*orm.User, error) {
	if r.dbConnection == nil {
		// 実行時エラーになるのでdbConnectionを必ず渡す制約をつける方法はないのか、、
		return nil, errors.New("正しく初期化されていません")
	}
	row := r.dbConnection.QueryRow(`
SELECT 
	user_id UserId,
	user_code UserCode,
	email Email,
	name Name,
	birth_date BirthDate,
	gender Gender,
	created_at CreatedAt,
	updated_at UpdatedAt,
	deleted_at DeletedAt
FROM user 
WHERE user_code = ?`, userCode)
	user, error := r.fetchUser(row)
	if error != nil {
		return nil, error
	}
	return user, nil
}
func (u UserRepository) Update(userId int, user *orm.User) error {
	return nil
}

func (r UserRepository) fetchUser(row *sql.Row) (*orm.User, error) {
	var user orm.User
	var tmpCreatedAt string
	var tmpUpdatedAt string
	//var tmpDeletedAt string
	var nullTmpDeletedAt sql.NullString
	e := row.Scan(&user.UserId, &user.UserCode, &user.Email, &user.Name, &user.BirthDate, &user.Gender, &tmpCreatedAt, &tmpUpdatedAt, &nullTmpDeletedAt)
	if e != nil {
		if e.Error() == "sql: no rows in result set" {
			// データがなかった
			return nil, nil
		} else {
			// 予期しないエラー
			return nil, e
		}
	}
	createdAt, _ := util.JstStringToTime(tmpCreatedAt)
	user.CreatedAt = createdAt
	updatedAt, _ := util.JstStringToTime(tmpUpdatedAt)
	user.UpdatedAt = updatedAt
	var t *time.Time
	if nullTmpDeletedAt.Valid {
		tmpDeletedAt, _ := util.JstStringToTime(nullTmpDeletedAt.String)
		t = &tmpDeletedAt
	}
	user.DeletedAt = t

	return &user, nil
}
