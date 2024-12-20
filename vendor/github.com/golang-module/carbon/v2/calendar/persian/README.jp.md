# ペルシア暦（イラン暦）

日本語 | [English](README.md) | [简体中文](README.cn.md)

#### 使い方の例

##### `西暦` を `ペルシャ暦` に変換

```go
// ペルシャ暦の取得
carbon.Parse("2020-08-05 13:14:15").Persian().Year() // 1399
// ペルシャ暦月の取得
carbon.Parse("2020-08-05 13:14:15").Persian().Month() // 5
// ペルシャ暦の取得日
carbon.Parse("2020-08-05 13:14:15").Persian().Day() // 15
// ペルシャ暦時間の取得
carbon.Parse("2020-08-05 13:14:15").Persian().Hour() // 13
// ペルシャ暦分の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Minute() // 14
// ペルシャ暦秒の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().Second() // 15

// ペルシャ暦日時文字列の取得
carbon.Parse("2020-08-05 13:14:15").Lunar().String() // 1399-05-15 13:14:15
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 1399-05-15 13:14:15
// ペルシア暦月文字列の取得
carbon.Parse("2020-08-05 13:14:15").Persian().ToMonthString() // مرداد
// ペルシャ暦週文字列の取得
carbon.Parse("2020-08-05 13:14:15").Persian().ToWeekString() // چهارشنبه

// ペルシャ暦ゼロ時間かどうか
carbon.Parse("0000-00-00 00:00:00").Persian().IsZero() // true
carbon.Parse("2020-08-05 13:14:15").Persian().IsZero() // false

// ペルシア暦閏年かどうか
carbon.Parse("2016-03-20 00:00:00").Persian().IsLeapYear() // true
carbon.Parse("2020-08-05 13:14:15").Persian().IsLeapYear() // false
```

##### ペルシャ暦を西暦に変換する

```go
carbon.CreateFromPersian(1395, 1, 1, 0, 0, 0).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(1399, 5, 15, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
```