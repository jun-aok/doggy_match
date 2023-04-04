//import { auth } from '../firebase';
import { Link, Redirect } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { Authenticator } from '../../Authenticator'
import { useAuthContext } from '../../context/AuthContext';
import { Profile } from '../../models/entities/Profile';
import axios from 'axios';


// ユーザー情報未登録時に表示するページ
const RegisterDog = () => {
  const [error, setError] = useState<{[key: string]:string}>({});
  

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    const { name, birthDate, gender, personality } = event.target.elements;  
    // サインインすればAuthContextのオブサーバーが動く
    const params = new URLSearchParams();
    params.append("name",name.value);
    params.append("birthDate",birthDate.value);
    params.append("gender",gender.value);
    params.append("personality",personality.value);
    axios.post('http://localhost:1323/api/dog', params).then(_ => {
      window.location.href = '/'
    }).catch(res => {
      if (res.response.status == 400) {
        setError(res.response.data)
      }
    })
  };

  return (
    <div>
      <p>
        あなたの愛犬のことを教えてください
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
          <label>性格</label>
          <select name="personality">
            <option value="1">活発</option>
            <option value="2">おとなしい</option>
          </select>
          <span>{ error["personality"] ?? "" }</span>
        </div>
        <div>
          <button>登録する</button>
        </div>
      </form>
    </div>
  );
};

export default RegisterDog;