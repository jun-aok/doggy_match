import moment from "moment";
import { DogEntity } from "./DogEntity";

export class CheckIn {
  constructor(checkInTime: moment.Moment, dogs: DogEntity[], latitude: string, longitude: string) {
    this._checkInTime = checkInTime;
    this._dogs = dogs;
    this._latitude = latitude;
    this._longitude = longitude;
  }
  private _checkInTime: moment.Moment;
  get checkInTime(): moment.Moment {
    return this._checkInTime;
  }

  private _dogs: DogEntity[];
  get dogs(): DogEntity[] {
    return this._dogs;
  }
  private _latitude: string
  get latitude(): string {
    return this._latitude;
  }
  private _longitude: string
  get longitude(): string {
    return this._longitude;
  }
}