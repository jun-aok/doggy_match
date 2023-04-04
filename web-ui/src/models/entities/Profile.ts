import { Gender, getGenderView } from "../enums/gender";
import moment from "moment"
import { UserEntity } from "./UserEntity";
import { DogEntity } from "./DogEntity";

export class Profile {
  constructor(email: string, token: string, user: UserEntity | null, dogs: DogEntity[] ) {
    this._email = email;
    this._token = token;
    this._user = user;
    this._dogs = dogs;
  }
  private _email: string;
  get email(): string{
    return this._email;
  }
  private _token: string | null;
  get token(): string | null {
    return this._token;
  }
  private _user: UserEntity | null;
  get user(): UserEntity | null {
    return this._user;
  }
  set user(user: UserEntity | null) {
    this._user = user;
  }
  private _dogs: DogEntity[];
  get dogs(): DogEntity[] {
    return this._dogs;
  }
  set dogs(dogs: DogEntity[]) {
    this._dogs = dogs;
  }
}