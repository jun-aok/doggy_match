import { Profile } from "./Profile";
import { CheckIn } from "./CheckIn";

export class UserInfo {
  private _profile: Profile | null;
  get profile(): Profile | null {
    return this._profile
  }
  set profile(profile: Profile | null) {
    this._profile = profile;
  }
  private _checkIn: CheckIn | null;
  get checkIn(): CheckIn | null {
    return this._checkIn;
  }
  set checkIn(checkIn: CheckIn | null) {
    this._checkIn = checkIn;
  }
}