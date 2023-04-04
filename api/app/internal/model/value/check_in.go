package value

import (
	. "app/internal/model"
	"app/internal/model/entity/orm"
	"app/pkg/calc"
	"app/pkg/util"
	"fmt"
	"reflect"
	"time"
)

type CheckIn struct {
	// ormじゃなくて専用のentityの方が良さそう
	Dogs        []*orm.Dog `json:"dogs"`
	Latitude    Position   `json:"latitude"`
	Longitude   Position   `json:"longitude"`
	CheckInTime time.Time  `json:"checkInTime"`
	isValid     bool
}

func (c *CheckIn) IsValid() bool {
	return true
}

// 犬ありの場合のチェックイン
// チェックイン済みだった場合のチェックアウト処理
func NewCheckIn(userId int, checkInDogIds []int, userDogs []*orm.Dog, latitude Position, longitude Position) (*CheckIn, error) {
	// todo latitude等々の妥当性
	var resultDogs []*orm.Dog = make([]*orm.Dog, 0)
	if len(checkInDogIds) > 0 {
		// 自分の犬かの確認
		var ids []int
		for _, d := range userDogs {
			ids = append(ids, d.DogId)
		}
		// dogIdsを自分のDogIdのものに絞り込んだ結果がdogIdsと同じなら全て自分の犬だということ
		if !reflect.DeepEqual(calc.Intersect(checkInDogIds, ids), checkInDogIds) {
			return &CheckIn{
				isValid: false,
			}, nil
		}
		dogDic, error := util.ToDictionary(userDogs, func(dog *orm.Dog) int {
			return dog.DogId
		})
		if error != nil {
			return nil, fmt.Errorf("不明なエラー")
		}
		resultDogs = util.Select(checkInDogIds, func(dogId int) *orm.Dog {
			return dogDic[dogId]
		})
	}
	return &CheckIn{
		Dogs:        resultDogs,
		Latitude:    latitude,
		Longitude:   longitude,
		CheckInTime: time.Now(),
		isValid:     true,
	}, nil

}
