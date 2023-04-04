//import { auth } from '../firebase';
import { Link, Redirect } from 'react-router-dom';
import { useState, useEffect } from 'react';
import { Authenticator } from '../Authenticator'
//import { useDark } from '../context/AuthContext';
import { Profile } from '../models/entities/Profile';
import './style/Login.css'


const Login = () => {
  const [error, setError] = useState<string | null>(null);
  const [oauthError, setOauthError] = useState<string | null>(null);
  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    const { email, password } = event.target.elements;
    // サインインすればAuthContextのオブサーバーが動く
    // トークンがなくなっていたら強制的にサインアウトする
    // これをしないとIndexedDBを消さないと動かなくなる
    Authenticator.signOut()
    var error = await Authenticator.signIn(email.value, password.value);
    if(error) {
      setError(error);
    } else {
      // スマートじゃないけど<Redirect>だとAuthContextがそのままで正しく動かせない
      // ログイン時にContextの値を更新できれば良いがやり方がいまいちわからず
      window.location.href = '/'
    }
  };

  const handleGoogleLogin = async (event: any) => {
    event.preventDefault();
    Authenticator.signOut()
    const error = await Authenticator.signInGoogle();
    if(error) {
      setOauthError(error);
    } else {
      window.location.href = '/'
    }
  }

  const handleFacebookLogin = async (event: any) => {
    event.preventDefault();
    Authenticator.signOut()
    const error = await Authenticator.signInFacebook();
    if(error) {
      setOauthError(error);
    } else {
      window.location.href = '/'
    }
  }

  return (
    <div className="login text-center">
      <p>Doggy Friends</p>
      <h1 className='h3 mb-3 font-weight-normal'>ログイン</h1>
      
      <form onSubmit={handleSubmit} className="form-signin">
        <div>
          <input name="email" type="email" placeholder="email" className='form-control' />
        </div>
        <div>
          <input name="password" type="password" placeholder="password"  className='form-control' />
        </div>
        <div>
          <button className="btn btn-primary btn-block">ログイン</button>
        </div>
        {
          error && <p>{ error }</p>
        }
        <div>
          他の方法でログイン
          <a onClick={handleGoogleLogin} className="login_other">
            <img src="image/logo/google_button_logo.svg" />Googleアカウントでログイン
          </a>
          <a onClick={handleFacebookLogin} className="login_other">
            <img src="image/logo/f_logo_RGB-Blue_250.png" />Facebookアカウントでログイン
          </a>
          {
            oauthError && <p>{oauthError}</p>
          }
        </div>
        <div>
          ユーザ登録は<Link to={'/signup'}>こちら</Link>から
        </div>
      </form>
      
    </div>
  );
};

export default Login;