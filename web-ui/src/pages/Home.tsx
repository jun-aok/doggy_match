// import { auth } from '../firebase';
import { useEffect } from 'react';
import { useHistory, Link } from 'react-router-dom';
import CheckIn from '../components/CheckIn';
import CheckOut from '../components/CheckOut';
import { useAuthContext } from '../context/AuthContext';
import './style/Home.css'


const Home = () => {
  const history = useHistory();
  const { userInfo } = useAuthContext();

  useEffect(() => {
    //console.log(auther);
  }, [])

  return (
    <div className="home">
      <main>
        <h1>
          {
            userInfo!.checkIn ? "チェックイン済み" : "チェックイン" 
          }
        </h1>
        <div>
          {
            userInfo!.checkIn ? <CheckOut /> : <CheckIn />
          }
          </div>
      </main>
      <div id="footer" className="fixed-bottom footer">
        <p>
          { 
            userInfo!.checkIn ? "チェックアウト" : "チェックイン" 
          }
        </p>
        <p>友達検索</p>
        <p>友達管理</p>
        <p><Link to="/mypage">マイページ</Link></p>
      </div>
    </div>
    
  );
};

export default Home;