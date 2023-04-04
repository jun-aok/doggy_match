import moment from "moment";
import { Gender, getGenderViewForDog } from "../enums/gender"

export class DogEntity {
  constructor(dogId: number, name: string, birthDate: moment.Moment, gender: number) {
    this._dogId = dogId;
    this._name = name;
    this._birthDate = birthDate;
    this._gender = gender;
  }
  private _dogId: number;
  get dogId(): number {
    return this._dogId;
  }
  private _name: string;
  get name(): string  {
    return this._name
  }
  private _gender: Gender | null;
  get gender(): Gender | null {
    return this._gender;
  }
  get genderView(): string {
    return getGenderViewForDog(this.gender)
  }
  private _birthDate: moment.Moment | null;
  get birthDate(): moment.Moment | null {
    return this._birthDate;
  }
  get birthDateView(): string {
    return this.birthDate?.format("YYYY-MM-DD") ?? ""
  }
}