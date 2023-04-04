import React, { createContext, useState, useContext, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { Authenticator } from '../Authenticator'
import { Profile } from '../models/entities/Profile';
import { UserEntity } from "../models/entities/UserEntity";
import { DogEntity } from "../models/entities/DogEntity";
import moment from 'moment';
import axios from 'axios';
import RegisterDog from '../pages/register/RegisterDog';
import Error500 from '../pages/Error500';

type ErrorContextType = {
  errorType: number;
  setErrorType: (errorType: number) => void;
}

const ErrorContext = createContext<ErrorContextType>({
  errorType: 0,
  setErrorType: (_) => {}
})

export function useErrorContext() {
  return useContext(ErrorContext);
}

export const ErrorProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [errorType, setErrorType] = useState<number>(0);
  return (

    <ErrorContext.Provider value={{ errorType, setErrorType }}>
      { errorType == 0 && children }
      { errorType == 500 && <Error500 /> }
    </ErrorContext.Provider>
  );
}