import { Link, Redirect } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { useAuthContext } from '../context/AuthContext';
import axios from 'axios';
import { UserInfo } from '../models/entities/UserInfo';

const CheckOut = () => {
  const [error, setError] = useState<{[key: string]:string}>({});
  const { userInfo, setUserInfo } = useAuthContext();
  

  const checkOut = async () => {
    axios.post('http://localhost:1323/api/check_out', {}).then(res => {
      const u : UserInfo = new UserInfo()
      u.profile = userInfo!.profile!
      // u.checkInは何もしない
      setUserInfo(u);
    }).catch(res => {
      if (res.response.status == 400) {
        setError(res.response.data)
      }
    })
  };

  return (
    <div>
      <button onClick={checkOut}>チェックアウト</button>
    </div>
  );
};

export default CheckOut;