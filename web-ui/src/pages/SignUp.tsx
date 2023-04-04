import { useState, useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import { Authenticator } from '../Authenticator'
import { Profile } from '../models/entities/Profile';
import { useAuthContext } from '../context/AuthContext';

const SignUp = () => {
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    const { email, password } = event.target.elements;
    var error = await Authenticator.signUp(email.value, password.value);
    if(error == null) {
      window.location.href = '/'
    } else {
      setError(error);
    }
  };

  return (
    <div className="login text-center">
      <p>Doggy Friends</p>
      <h1 className='h3 mb-3 font-weight-normal'>ユーザ登録</h1>
      <form onSubmit={handleSubmit} className="form-signin">
        <div>
          <input name="email" type="email" placeholder="メールアドレス" className='form-control' />
        </div>
        <div>
          <input name="password" type="password" placeholder="パスワード"  className='form-control' />
        </div>
        <div>
          <button className="btn btn-primary btn-block">登録</button>
        </div>
        {
          error &&
            <div>
              { error }
            </div>
        }
      </form>
    </div>
  );
};

export default SignUp;