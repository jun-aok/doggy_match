import Home from './pages/Home';
import MyPage from './pages/MyPage';
import SignUp from './pages/SignUp';
import Login from './pages/Login';
import RegisterUser from './pages/register/RegisterUser';
import RegisterDog from './pages/register/RegisterDog';
import Welcome from './pages/register/Welcome';
import { AuthProvider } from './context/AuthContext';
import { BrowserRouter, Route } from 'react-router-dom';
import PrivateRoute from './components/PrivateRoute';
import {} from './axios-middleware';
import {} from './Authenticator';
import { ErrorProvider } from './context/ErrorContext';

function App() {
  return (
    <ErrorProvider>
      <AuthProvider>
        <div style={{ margin: '2em' }}>
          <BrowserRouter>
            {/* サインインサインアップ */}
            <Route path="/signup" component={SignUp} />
            <Route path="/login" component={Login} />
            {/* 初回登録 */}
            <PrivateRoute exact path="/register_user" component={RegisterUser} />
            <PrivateRoute exact path="/register_dog" component={RegisterDog} />
            <PrivateRoute exact path="/welcome" component={Welcome} />

            <PrivateRoute exact path="/" component={Home} />
            <PrivateRoute exact path="/mypage" component={MyPage} />
          </BrowserRouter>
        </div>
      </AuthProvider>
    </ErrorProvider>
  );
}

export default App;