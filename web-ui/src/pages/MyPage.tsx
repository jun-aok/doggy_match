import { useEffect, useState } from 'react';
import axios from 'axios';
import { useAuthContext } from '../context/AuthContext';
import { Authenticator } from '../Authenticator'

const MyPage = () => {
  const auther = useAuthContext();
  // useEffect(() => {
  //   const getData = async () => {
  //     const res: any = await axios.get('https://localhost:5001/CurrentUser')
  //     setRes(new CurrentUserResponse(res.data.email, res.data.name, res.data.completed));
  //   }
  //   getData();
  // }, [])
  const handleLogout = () => {
    Authenticator.signOut();
    window.location.href = '/login'
  };

  return (
    <div>
      <h1>マイページ</h1>
      {
        auther && 
          <>
            <h2>メールアドレス</h2>
            <p>{auther.userInfo?.profile?.email}</p>
            <h2>名前</h2>
            {/* <p>{auther.name}</p>
            <h2>性別</h2>
            <p>{auther.gender}</p>
            <h2>誕生日</h2>
            <p>{auther.birthDateView}</p> */}
          </>
      }
      <button onClick={handleLogout}>ログアウト</button>
    </div>
  );
};

export default MyPage;