import React, { createContext, useState, useContext, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { Authenticator } from '../Authenticator'
import { Profile } from '../models/entities/Profile';
import { UserInfo } from '../models/entities/UserInfo';
import { UserEntity } from "../models/entities/UserEntity";
import { DogEntity } from "../models/entities/DogEntity";
import { useErrorContext } from './ErrorContext';
import moment from 'moment';
import axios from 'axios';
import { CheckIn } from '../models/entities/CheckIn';

type AuthContextType = {
  userInfo: UserInfo | null;
  setUserInfo: (auther: UserInfo) => void;
}

const AuthContext = createContext<AuthContextType>({
  userInfo: null,
  setUserInfo: (_) => {}
})

export function useAuthContext() {
  return useContext(AuthContext);
}

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [userInfo, setUserInfo] = useState<UserInfo | null>(null);
  const [loading, setLoading] = useState(true);
  const history = useHistory();
  const error = useErrorContext();

  useEffect(() => {
    // ログイン状態を監視
    const unsubscribed = Authenticator.getAuth().onAuthStateChanged(async (user) => {
      let token = Authenticator.getToken()
      // ログインしていない
      if(!user || token == null) {
        setLoading(false);
        return
      }
      axios.get('http://localhost:1323/api/user_info', { timeout: 3000 })
        .then(async res => {
          let userInfo : UserInfo = new UserInfo();
          //var profile : Profile | null = null
          if(res.data.userInfo) {

            const u = res.data.userInfo.profile.user
            const ds = res.data.userInfo.profile.dogs

            userInfo.profile = new Profile(
              user.email!,
              await user.getIdToken(),
              u ? new UserEntity(u.name, u.birthDate, u.gender) : null,
              ds ? ds.map((d: any) => new DogEntity(d.dogId, d.name, d.birthDate, d.gender)) : [],
            )
            if (res.data.userInfo.checkIn) {
              const cds = res.data.userInfo.checkIn.dogs;
              userInfo.checkIn = new CheckIn(
                res.data.userInfo.checkIn.checkInTime,
                cds ? cds.map((d: any) => new DogEntity(d.dogId, d.name, d.birthDate, d.gender)) : [],
                res.data.userInfo.checkIn.latitude,
                res.data.userInfo.checkIn.longitude
              );
            }
          } else {
            userInfo.profile = new Profile(
              user.email!,
              await user.getIdToken(),
              null,
              []
            )
          }
          setUserInfo(userInfo);
          setLoading(false);
        }).catch(err => {
          if (err.response?.status == 400) {
            // トークンの期限切れ
            Authenticator.signOut()
          } 
          if (!err.response) {
            // サーバーダウン？
            // alert("サーバーダウン")
            error.setErrorType(500)
            setLoading(false);
          }
          // ここ、エラー条件によってはサーバーダウンとか知らせる必要があるかも
          setLoading(false);
          console.log(err);
        });
    });
    return () => {
      unsubscribed();
    };
  }, []);

  if (loading) {
    return <p>loading...</p>;
  } else {
    return (
      <AuthContext.Provider value={{userInfo, setUserInfo}}>
        {!loading && children}
      </AuthContext.Provider>
    );
  }
}