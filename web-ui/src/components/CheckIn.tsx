import { Link, Redirect } from 'react-router-dom';
import { useState, useEffect } from 'react';
import axios from 'axios';
import { UserInfo } from '../models/entities/UserInfo';
import { CheckIn as CheckInEntity } from '../models/entities/CheckIn';
import { useAuthContext } from '../context/AuthContext';

// type Props = {
//   profile: Profile
// }

const CheckIn = () => {
  const [error, setError] = useState<{[key: string]:string}>({});
  //const [checkedDog, setCheckedDog] = useState<number[]>([]);

  const { userInfo, setUserInfo } = useAuthContext();
  

  const checkIn = async (event: any) => {
    event.preventDefault();
    const { dogIds } = event.target.elements;
    navigator.geolocation.getCurrentPosition((p: GeolocationPosition) => {
      // 緯度
      const latitude: number = p.coords.latitude;
      // 経度
      const longitude: number = p.coords.longitude;

      const params = new URLSearchParams();
      params.append("latitude", latitude.toString());
      params.append("longitude", longitude.toString());
      params.append("dog_ids", dogIds.value);

      axios.post('http://localhost:1323/api/check_in', params).then(res => {
        const u : UserInfo = new UserInfo()
        u.profile = userInfo!.profile!
        u.checkIn = new CheckInEntity(
          res.data.checkIn.checkInTime,
          res.data.checkIn.dogs,
          res.data.checkIn.latitude,
          res.data.checkIn.longitude,
        )
        setUserInfo(u)
      }).catch(res => {
        if (res.response.status == 400) {
          setError(res.response.data)
        }
      })
    }, (e: GeolocationPositionError) => {
      alert("現在地が取得できません")
      console.log(e);
    })
  };

  return (
    <div>
      <form onSubmit={checkIn}>
        <select name="dogIds">
          <option value="0">犬なし</option>
          {
            userInfo && userInfo.profile!.dogs.map(d => 
              <option value={d.dogId} key={d.dogId}>{d.name}</option>
            )
          }
        </select>
        <div>
          <button>チェックイン</button>
        </div>
      </form>
    </div>
  );
};

export default CheckIn;