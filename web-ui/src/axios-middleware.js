import axios from 'axios';
import { Service } from 'axios-middleware';
import { Authenticator } from './Authenticator';
const service = new Service(axios);

 
service.register({
  onRequest(config) {
    const token = Authenticator.getToken();
    if(token != null) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  }
});

service.register({
  onResponseError(err) {
    if (err.response && err.response.status === 401 && err.config && !err.config.hasRetriedRequest) {
      const http = axios.create();
      return Authenticator.refreshToken()
        .then((token) => {
          console.log(token);
          return http({
            ...err.config,
            hasRetriedRequest: true,
            headers: {
              ...err.config.headers,
              Authorization: `Bearer ${token}`
            }
          })
        })
        .catch((error) => {
          console.log('Refresh login error: ', error);
          throw error;
        });
    }
    throw err;
  }
})
