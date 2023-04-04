package repository

import (
	"app/internal/model/entity"
	"database/sql"
	"errors"
)

type ProfileRepository struct {
	dbConnection *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{
		dbConnection: db,
	}
}
func (r *ProfileRepository) Query(userId int) (*entity.Profile, error) {
	if r.dbConnection == nil {
		// 実行時エラーになるのでdbConnectionを必ず渡す制約をつける方法はないのか、、
		return nil, errors.New("正しく初期化されていません")
	}

	row := r.dbConnection.QueryRow(`
SELECT 
	email Email,
	user.user_id UserId,
	user.name UserName,
	user.birth_date UserBirthDate,
	user.gender UserGender
FROM user 
WHERE user_id = ? LIMIT 1`, userId) // 1件しかない

	p := entity.Profile{}
	u := entity.ProfileUser{}
	e := row.Scan(&p.Email, &u.UserId, &u.Name, &u.BirthDate, &u.Gender)
	if e != nil {
		// データなし
		if e == sql.ErrNoRows {
			return nil, nil
		}
	}
	p.User = &u

	dogRows, e := r.dbConnection.Query(`
SELECT
  dog.dog_id DogId, 
	dog.name Name,
	dog.birth_date BirthDate,
	dog.gender Gender,
	dog.personality Personality
FROM dog 
INNER JOIN user ON dog.user_id = user.user_id
WHERE user.user_id = ?`, userId)
	if e != nil {
		return nil, e
	}
	defer dogRows.Close()

	array := make([]*entity.ProfileDog, 0, 5)
	for dogRows.Next() {
		d := entity.ProfileDog{}
		e := dogRows.Scan(&d.DogId, &d.Name, &d.BirthDate, &d.Gender, &d.Personality)
		if e != nil {
			return nil, e
		}
		array = append(array, &d)
	}
	if row.Err() != nil {
		return nil, row.Err()
	}
	p.Dogs = array
	return &p, nil
}

// func (r UserRepository) fetchUser(row *sql.Row) (*orm.User, error) {
// 	var user orm.User
// 	var tmpCreatedAt string
// 	var tmpUpdatedAt string
// 	//var tmpDeletedAt string
// 	var nullTmpDeletedAt sql.NullString
// 	e := row.Scan(&user.UserId, &user.UserCode, &user.Email, &user.Name, &user.BirthDate, &user.Gender, &tmpCreatedAt, &tmpUpdatedAt, &nullTmpDeletedAt)
// 	if e != nil {
// 		if e.Error() == "sql: no rows in result set" {
// 			// データがなかった
// 			return nil, nil
// 		} else {
// 			// 予期しないエラー
// 			return nil, e
// 		}
// 	}
// 	createdAt, _ := util.JstStringToTime(tmpCreatedAt)
// 	user.CreatedAt = createdAt
// 	updatedAt, _ := util.JstStringToTime(tmpUpdatedAt)
// 	user.UpdatedAt = updatedAt
// 	var t *time.Time
// 	if nullTmpDeletedAt.Valid {
// 		tmpDeletedAt, _ := util.JstStringToTime(nullTmpDeletedAt.String)
// 		t = &tmpDeletedAt
// 	}
// 	user.DeletedAt = t

// 	return &user, nil
// }
