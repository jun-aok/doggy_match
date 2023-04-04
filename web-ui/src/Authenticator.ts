import axios from 'axios';
import firebase from 'firebase/compat/app'
import 'firebase/compat/auth';
import { Profile } from './models/entities/Profile';
import { getFirebaseAuthError } from './responses/firebase_response'

export class Authenticator {
  public static getAuth = (): firebase.auth.Auth => {
    return firebase.auth();
  }
  public static signInGoogle = async (): Promise<string | null> => {
    const method = "signin"
    return firebase
      .auth()
      .signInWithPopup(new firebase.auth.GoogleAuthProvider())
        .then(async (userCredential: firebase.auth.UserCredential) => {
          if(!userCredential.user) {
            // 認証エラー
            return null;
          }
          const a = await userCredential.user.getIdToken()
          Authenticator.saveToken(await userCredential.user.getIdToken());
          return null;
        }).catch((error: firebase.auth.Error) => {
          return getFirebaseAuthError(error, method)
        })
  }
  public static signInFacebook = async (): Promise<string | null>  => {
    const method = "signin"
    return firebase
      .auth()
      .signInWithPopup(new firebase.auth.FacebookAuthProvider())
        .then(async (userCredential: firebase.auth.UserCredential) => {
          if(!userCredential.user) {
            // ここには来ないと思う
            return 'エラーが発生しました。しばらく時間をおいてお試しください'
          }
          Authenticator.saveToken(await userCredential.user.getIdToken());
          return null;
        }).catch((error: firebase.auth.Error) => {
          return getFirebaseAuthError(error, method)
        })
  }
  public static getToken = (): string | null => localStorage.getItem('jwt');
  public static refreshToken = async () : Promise<string | null> =>  {
    const tokenResult = await Authenticator.getAuth().currentUser?.getIdTokenResult()
    if(tokenResult != null) {
      Authenticator.saveToken(tokenResult.token)
      return tokenResult.token;
    }
    return null;
  }
  public static signUp = (email: string, password: string): Promise<string | null> => {
    const method = "signup"
    return Authenticator.getAuth().createUserWithEmailAndPassword(email, password)
      .then(async (userCredential: firebase.auth.UserCredential) => {
        if(!userCredential.user) {
          // ここには来ないと思う
          return 'エラーが発生しました。しばらく時間をおいてお試しください'

        }
        Authenticator.saveToken(await userCredential.user.getIdToken());
        return null;
      }).catch((error: firebase.auth.Error) => {
        return getFirebaseAuthError(error, method)
      })
  }
  public static signIn = (email: string, password: string): Promise<string | null> => {
    const method = "signin"
    return Authenticator.getAuth().signInWithEmailAndPassword(email, password)
    .then(async (userCredential: firebase.auth.UserCredential) => {
      if(!userCredential.user) {
        // ここには来ないと思う
        return 'エラーが発生しました。しばらく時間をおいてお試しください'
      }
      // この段階でonAuthStateChangedの処理が走っているがそちらではuserにアクセスしないようにしている
      Authenticator.saveToken(await userCredential.user.getIdToken());
      return null;
    }).catch((error: firebase.auth.Error) => {
      return getFirebaseAuthError(error, method)
    });
  }
    
  static signOut = () => {
    localStorage.removeItem('jwt');
    Authenticator.getAuth().signOut();
  }

  static saveToken = (token: string): void => localStorage.setItem('jwt', token);
}

firebase.initializeApp({
  apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
  authDomain: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
  projectId: process.env.REACT_APP_FIREBASE_PROJECT_ID,
  storageBucket: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.REACT_APP_FIREBASE_MESSAGE_SENDER_ID,
  appId: process.env.REACT_APP_FIREBASE_SENDER_ID,
});
// export default firebase;
// export const auth = firebase.auth();

// 現在ログインしているユーザーを取得する
  //public static getCurrentUser = async () : Promise<firebase.User | null> => {
  //  return new Promise((resolve, reject) => {
  //    Auther.getAuth().onAuthStateChanged(async(user) => {
  //      if (user) {
  //        resolve(user);
  //      } else {
  //        reject();
  //      }
  //    });
  //  })
  //}