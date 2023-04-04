import { Route, Redirect, useLocation } from 'react-router-dom';
import { useAuthContext } from '../context/AuthContext';
import { UserInfo } from '../models/entities/UserInfo';
const PrivateRoute: React.FC<any> = ({ component, exact, path }) => {
  const auther = useAuthContext();
  const location = useLocation();
  // 未ログイン
  // userInfoがnull
  if(!auther.userInfo) {
    return <Redirect to="/login" />
  }
  // ログインはしているがプロファイル未登録
  // userInfoにインスタンスがある
  if(!auther.userInfo.profile?.user && location.pathname != "/register_user") {
    return <Redirect to="/register_user" />
  }
  // プロファイル登録済みで/register_profileにアクセスしようとした
  if(auther.userInfo.profile?.user && location.pathname == "/register_user") {
    return <Redirect to="/" />
  }
  return ['/login', '/signup'].includes(location.pathname) ? <Redirect to="/" /> : <Route exact={exact} path={path} component={component} />
};

export default PrivateRoute;