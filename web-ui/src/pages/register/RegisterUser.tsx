import { useState, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import axios, { AxiosResponse } from 'axios';
import { useAuthContext } from '../../context/AuthContext';
import { Profile } from '../../models/entities/Profile';
import { UserInfo } from '../../models/entities/UserInfo';
import { UserEntity } from '../../models/entities/UserEntity';


// ユーザー情報未登録時に表示するページ
const RegisterUser = () => {
  const [error, setError] = useState<{[key: string]:string}>({});
  
  const { userInfo, setUserInfo } = useAuthContext();
  const history = useHistory();

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    const { name, birthDate, gender } = event.target.elements;  
    // サインインすればAuthContextのオブサーバーが動く
    const params = new URLSearchParams();
    params.append("name",name.value);
    params.append("birthDate",birthDate.value);
    params.append("gender",gender.value);
    axios.post('http://localhost:1323/api/user', params).then((res: AxiosResponse<any>) => {
      userInfo!.profile = new Profile(
        userInfo!.profile!.email,
        userInfo!.profile!.token!,
        new UserEntity(
          res.data.name, 
          res.data.birthDate, 
          res.data.gender
        ),
        []
      )
      setUserInfo(userInfo!)
      history.push('/register_dog');
    }).catch(res => {
      if (res.response.status == 400) {
        setError(res.response.data)
      }
    })
  };

  return (
    <div>
      <p>
        あなたのことを教えてください
      </p>
      <form onSubmit={handleSubmit}>
        <div>
          <label>名前</label>
          <input name="name" type="text" placeholder="名前" />
          <span>{ error["name"] ?? "" }</span>
        </div>
        <div>
          <label>誕生日</label>
          <input name="birthDate" type="text" placeholder="2022-01-01" />
          <span>{ error["birthDate"] ?? "" }</span>
        </div>
        <div>
          <label>性別</label>
          <select name="gender">
            <option value="1">男</option>
            <option value="2">女</option>
            <option value="3">無回答</option>
          </select>
          <span>{ error["gender"] ?? "" }</span>
        </div>
        <div>
          <button>登録する</button>
        </div>
      </form>
    </div>
  );
};

export default RegisterUser;