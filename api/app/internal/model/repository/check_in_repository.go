package repository

import (
	"app/internal/model/entity/orm"
	"app/internal/model/value"
	"app/pkg/util"
	"database/sql"
	"errors"
	"fmt"
)

type CheckInRepository struct {
	transaction  *sql.Tx
	dbConnection *sql.DB
}

func NewCheckInRepositoryWithTransaction(tx *sql.Tx) *CheckInRepository {
	return &CheckInRepository{
		transaction: tx,
	}
}

func (r *CheckInRepository) Query(userId int) (*value.CheckIn, error) {
	if r.transaction == nil {
		return nil, errors.New("正しく初期化されていません")
	}
	row := r.transaction.QueryRow(`
	SELECT 
		latitude Latitude,
		longitude Longitude,
		check_in_time CheckInTime
	FROM check_in 
	WHERE user_id = ? LIMIT 1`, userId)

	var checkIn value.CheckIn
	var tmpCheckInTime string
	e := row.Scan(&checkIn.Latitude, &checkIn.Longitude, &tmpCheckInTime)
	if e != nil {
		if e.Error() == "sql: no rows in result set" {
			// データがなかった
			return nil, nil
		} else {
			// 予期しないエラー
			return nil, e
		}
	}
	checkInTime, _ := util.JstStringToTime(tmpCheckInTime)
	checkIn.CheckInTime = checkInTime

	row2, error := r.transaction.Query(`
	SELECT 
		d.dog_id DogId,
		name Name,
		birth_date BirthDate,
		gender Gender,
		personality Personality
	FROM check_in_dog cid
	INNER JOIN dog d ON cid.dog_id = d.dog_id
	WHERE d.user_id = ?`, userId)
	if error != nil {
		return nil, error
	}
	defer row2.Close()

	dogs := make([]*orm.Dog, 0, 5)
	for row2.Next() {
		d := orm.Dog{}
		e := row.Scan(&d.DogId, &d.Name, &d.BirthDate, &d.Gender, &d.Personality)
		if e != nil {
			// データなし
			if e == sql.ErrNoRows {
				checkIn.Dogs = dogs
				return &checkIn, nil
			}
		}
		dogs = append(dogs, &d)
	}
	if row.Err() != nil {
		return nil, row.Err()
	}
	checkIn.Dogs = dogs
	return &checkIn, nil
}

func (r *CheckInRepository) Regist(userId int, checkIn *value.CheckIn) error {
	if r.transaction == nil {
		return errors.New("正しく初期化されていません")
	}
	stmt, e1 := r.transaction.Prepare(`INSERT INTO check_in(user_id, latitude, longitude, check_in_time) VALUES(?,?,?,?)`)
	if e1 != nil {
		// エラー
		return fmt.Errorf("DBのPrepareに失敗しました check_in %s", e1.Error())
	}
	defer stmt.Close()
	_, e2 := stmt.Exec(userId, checkIn.Latitude.GetValue().String(), checkIn.Longitude.GetValue().String(), checkIn.CheckInTime)
	if e2 != nil {
		return fmt.Errorf("insertに失敗しました check_in %s", e2.Error())
	}

	if len(checkIn.Dogs) < 1 {
		return nil
	}

	stmt2, e3 := r.transaction.Prepare(`INSERT INTO check_in_dog(user_id, dog_id) VALUE(?,?)`)
	if e3 != nil {
		// エラー
		return fmt.Errorf("DBのPrepareに失敗しました check_in_dog %s", e3.Error())
	}
	defer stmt2.Close()
	for _, dog := range checkIn.Dogs {
		_, e4 := stmt2.Exec(userId, dog.DogId)
		if e4 != nil {
			return fmt.Errorf("insertに失敗しました check_in_dog %s", e2.Error())
		}
	}
	return nil
}

func (r *CheckInRepository) Delete(userId int) error {
	if r.transaction == nil {
		return errors.New("正しく初期化されていません")
	}
	stmt, e1 := r.transaction.Prepare(`DELETE FROM check_in WHERE user_id = ?`)
	if e1 != nil {
		// エラー
		return fmt.Errorf("DBのPrepareに失敗しました check_in %s", e1.Error())
	}
	defer stmt.Close()
	_, e2 := stmt.Exec(userId)
	if e2 != nil {
		return fmt.Errorf("check_inのdeleteに失敗しました check_in %s", e2.Error())
	}

	stmt2, e3 := r.transaction.Prepare(`DELETE FROM check_in_dog WHERE user_id = ?`)
	if e3 != nil {
		// エラー
		return fmt.Errorf("DBのPrepareに失敗しました check_in_dog %s", e3.Error())
	}
	defer stmt2.Close()
	_, e4 := stmt2.Exec(userId)
	if e4 != nil {
		return fmt.Errorf("check_in_dogのdeleteに失敗しました check_in %s", e2.Error())
	}
	return nil
}

func (r *CheckInRepository) RegistLog(userId int, log *value.CheckIn) error {
	//ログの登録
	return nil
}

func (r *CheckInRepository) CompleteLog(userId int) error {
	return nil
}

//	RegistLog(userId int, log value.CheckInLog)
//}
