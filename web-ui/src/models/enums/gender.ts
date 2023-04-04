//export class Enum {

export enum Gender {
  Male = 1,
  Female = 2,
  NoAnswer = 3,
}

export function getGenderView(gender: Gender | null): string {
  if (gender == null) {
    return ""
  }
  const map: { [key: number]: string }  = {}
  // この書き方じゃないとGenderをkeyにできない
  map[Gender.Male] = "男性"
  map[Gender.Female] = "女性"
  map[Gender.NoAnswer] = "無回答"
  return map[gender] ?? ""
}

export function getGenderViewForDog(gender: Gender | null): string {
  if (gender == null) {
    return ""
  }
  const map: { [key: number]: string }  = {}
  // この書き方じゃないとGenderをkeyにできない
  map[Gender.Male] = "♂"
  map[Gender.Female] = "♀"
  map[Gender.NoAnswer] = "不明"
  return map[gender] ?? ""
}