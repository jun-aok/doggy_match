//import { auth } from '../firebase';
import { Link, Redirect } from 'react-router-dom';


// ユーザー情報未登録時に表示するページ
const RegisterDog = () => {
  return (
    <div>
      <p>
        ようこそ
      </p>
      <p>
        <Link to={'/'}>ホームへ</Link>
      </p>
    </div>
  );
};

export default RegisterDog;