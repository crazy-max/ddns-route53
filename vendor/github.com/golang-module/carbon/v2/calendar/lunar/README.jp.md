# 中国の旧暦

日本語 | [English](README.md) | [简体中文](README.cn.md)

#### 使い方の例

> 現在は西暦` 1900 `年から` 2100 `年までの` 200 `年の時間スパンのみをサポートしている

##### `西暦`を`旧暦`に変換する

```go
// 旧暦の干支を手に入れる
carbon.Parse("2020-08-05 13:14:15").Lunar().Animal() // 鼠
// 旧暦の祝日を取得する
carbon.Parse("2021-02-12 13:14:15").Lunar().Festival() // 春节

// 旧正月の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Year() // 2020
// 旧暦月の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Month() // 6
// 旧暦うるう月の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().LeapMonth() // 4
// 旧暦日の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Day() // 16
// 旧暦時刻の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Hour() // 13
// 旧暦分の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Minute() // 14
// 旧暦の取得秒数
carbon.Parse("2020-08-05 13:14:15").Lunar().Second() // 15

// 旧暦日時文字列の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().String() // 2020-06-16 13:14:15
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 2020-06-16 13:14:15
// 旧正月文字列の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().ToYearString() // 二零二零
// 旧暦月文字列の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().ToMonthString() // 六月
// 旧暦うるう月文字列の取得
carbon.Parse("2020-04-23 13:14:15").Lunar().ToMonthString() // 闰四月
// 旧暦週文字列の取得
carbon.Parse("2020-04-23 13:14:15").Lunar().ToWeekString() // 周四
// 旧暦日文字列の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDayString() // 十六
// 旧暦日付文字列の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().ToDateString() // 二零二零年六月十六

// ゼロ値の時間ですか
carbon.Parse("0000-00-00 00:00:00").Lunar().IsZero() // true
carbon.Parse("2020-08-05 13:14:15").Lunar().IsZero() // false

// 旧暦うるう年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapYear() // true
// 旧暦うるう月かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapMonth() // false

// ねずみ年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRatYear() // true
// 牛年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsOxYear() // false
// 寅年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsTigerYear() // false
// うさぎ年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRabbitYear() // false
// 龍年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDragonYear() // false
// 蛇の年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsSnakeYear() // false
// 馬年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsHorseYear() // false
// 羊年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsGoatYear() // false
// 申年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsMonkeyYear() // false
// 鶏の年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsRoosterYear() // false
// 犬年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsDogYear() // false
// 豚年かどうか
carbon.Parse("2020-08-05 13:14:15").Lunar().IsPigYear() // false
```

##### `旧暦`を`西暦`に変換する

```go
// 2023 年の旧暦 12 月 11 日をグレゴリオ暦に変換します
carbon.CreateFromLunar(2023, 12, 11, 0, 0, 0, false).ToDateTimeString() // 2024-01-21 00:00:00
// 旧暦の 2023 年 2 月 11 日をグレゴリオ暦に変換します
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, false).ToDateTimeString() // 2023-03-02 00:00:00
// 旧暦 2023 年、閏 2 月 11 日をグレゴリオ暦に変換します
carbon.CreateFromLunar(2023, 2, 11, 0, 0, 0, true).ToDateTimeString() // 2023-04-01 00:00:00
```