# Carbon  #

[![Carbon Release](https://img.shields.io/github/release/golang-module/carbon.svg)](https://github.com/golang-module/carbon/releases)
[![Go Test](https://github.com/golang-module/carbon/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/carbon/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/carbon/v2)](https://goreportcard.com/report/github.com/golang-module/carbon/v2)
[![Go Coverage](https://codecov.io/gh/golang-module/carbon/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/carbon)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/golang-module/carbon/badges/download-count.svg)](https://goproxy.cn)
[![Carbon Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/carbon/v2)
[![License](https://img.shields.io/github/license/golang-module/carbon)](https://github.com/golang-module/carbon/blob/master/LICENSE)

日本語 | [English](README.md) | [简体中文](README.cn.md)

軽量でセマンティックで開発者に優しい golang 時間処理ライブラリ

Carbon は [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") に収録されています, よかったら, スターをください

[github.com/golang-module/carbon](https://github.com/golang-module/carbon "github.com/golang-module/carbon")

[gitee.com/golang-module/carbon](https://gitee.com/golang-module/carbon "gitee.com/golang-module/carbon")

#### インストール使用

##### Golangバージョンは1.16以上です (推奨)

```go
// github倉庫を使う
go get -u github.com/golang-module/carbon/v2

import "github.com/golang-module/carbon/v2"


// gitee倉庫を使う
go get -u gitee.com/golang-module/carbon/v2

import "gitee.com/golang-module/carbon/v2"
```

##### Golangバージョンは1.16より小さいです (必要)

```go
// github倉庫を使う
go get -u github.com/golang-module/carbon

import "github.com/golang-module/carbon"

// gitee倉庫を使う
go get -u gitee.com/golang-module/carbon

import "gitee.com/golang-module/carbon"
```

> v1とv2の違いについては、<a href="#人気のある問題">人気のある問題</a>をご覧ください，v1 は更新されなくなり、維持されるだけです。

人気のある問題
#### 使い方の例

> 現在時刻が 2020-08-05 13:14:15.999999999 +0800 CST であると仮定します。

##### グローバル・デフォルトの設定

```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
  Timezone: carbon.PRC,
  WeekStartsAt: carbon.Sunday,
  Locale: "jp",
})
```

> 設定されていない場合，デフォルトのレイアウト テンプレートは `2006-01-02 15:04:05`，デフォルトのタイムゾーンは `Local`，デフォルトの週の開始日は `日曜日`，デフォルトの言語は `en` です。

##### Carbon と time.Time 交換

```go
// 標準の time.Time を Carbon に変換します
carbon.CreateFromStdTime(time.Now())
// Carbon を標準の time.Time に変換します
carbon.Now().StdTime()
```

##### 昨日、今日、明日

```go
// 今日の瞬間
fmt.Printf("%s", carbon.Now()) // 2020-08-05 13:14:15
carbon.Now().String() // 2020-08-05 13:14:15
carbon.Now().ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今日の日付
carbon.Now().ToDateString() // 2020-08-05
// 今日の時間
carbon.Now().ToTimeString() // 13:14:15
// タイムゾーン指定の今日
carbon.Now(Carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
// 今日は秒タイムスタンプ
carbon.Now().Timestamp() // 1596604455
// 今日のミリ秒タイムスタンプ
carbon.Now().TimestampMilli() // 1596604455999
// 今日のマイクロ秒タイムスタンプ
carbon.Now().TimestampMicro() // 1596604455999999
// 今日のナノ秒タイムスタンプ
carbon.Now().TimestampNano() // 1596604455999999999

// 昨日の今は
fmt.Printf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
carbon.Yesterday().String() // 2020-08-04 13:14:15
carbon.Yesterday().ToString() // 2020-08-04 13:14:15 +0800 CST
carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨日の日付
carbon.Yesterday().ToDateString() // 2020-08-04
// 昨日の時間
carbon.Yesterday().ToTimeString() // 13:14:15
// 日付指定の昨日
carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// タイムゾーン指定の昨日
carbon.Yesterday(Carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
// 昨日の秒タイムスタンプ
carbon.Yesterday().Timestamp() // 1596518055
// 昨日のミリ秒タイムスタンプ
carbon.Yesterday().TimestampMilli() // 1596518055999
// 昨日のマイクロ秒タイムスタンプ
carbon.Yesterday().TimestampMicro() // 1596518055999999
// 昨日のナノ秒タイムスタンプ
carbon.Yesterday().TimestampNano() // 1596518055999999999

// 明日の今は
fmt.Printf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
carbon.Tomorrow().String() // 2020-08-06 13:14:15
carbon.Tomorrow().ToString() // 2020-08-06 13:14:15 +0800 CST
carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明日の日付
carbon.Tomorrow().ToDateString() // 2020-08-06
// 明日の時間
carbon.Tomorrow().ToTimeString() // 13:14:15
// 日付指定の明日
carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// タイムゾーン指定の明日
carbon.Tomorrow(Carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
// 明日の秒タイムスタンプ
carbon.Tomorrow().Timestamp() // 1596690855
// 明日のミリ秒タイムスタンプ
carbon.Tomorrow().TimestampMilli() // 1596690855999
// 明日のマイクロ秒タイムスタンプ
carbon.Tomorrow().TimestampMicro() // 1596690855999999
// 明日のナノ秒タイムスタンプ
carbon.Tomorrow().TimestampNano() // 1596690855999999999
```

##### Carbon オブジェクトを作成する

```go
// 秒タイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestamp(-1).ToString() // 1970-01-01 07:59:59 +0800 CST
carbon.CreateFromTimestamp(0).ToString() // 1970-01-01 08:00:00 +0800 CST
carbon.CreateFromTimestamp(1).ToString() // 1970-01-01 08:00:01 +0800 CST
carbon.CreateFromTimestamp(1649735755).ToString() // 2022-04-12 11:55:55 +0800 CST
// ミリ秒のタイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestampMilli(1649735755981).ToString() // 2022-04-12 11:55:55.981 +0800 CST
// マイクロ秒タイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestampMicro(1649735755981566).ToString() // 2022-04-12 11:55:55.981566 +0800 CST
// ナノタイムスタンプから Carbon オブジェクトを作成します
carbon.CreateFromTimestampNano(1649735755981566000).ToString() // 2022-04-12 11:55:55.981566 +0800 CST

// 年月日から分秒で Carbon オブジェクトを作成します
carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
// 年月日から分秒で Carbon オブジェクトを作成します，ミリ秒を含む
carbon.CreateFromDateTimeMilli(2020, 1, 1, 13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0800 CST
// 年月日から分秒で Carbon オブジェクトを作成します，マイクロ秒を含む
carbon.CreateFromDateTimeMicro(2020, 1, 1, 13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0800 CST
// 年月日から分秒で Carbon オブジェクトを作成します，ナノ秒を含む
carbon.CreateFromDateTimeNano(2020, 1, 1, 13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0800 CST

// 年月日から Carbon オブジェクトを作成します
carbon.CreateFromDate(2020, 8, 5).ToString() // 2020-08-05 00:00:00 +0800 CST
// 年月日から Carbon オブジェクトを作成します，ミリ秒を含む
carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 00:00:00.999 +0800 CST
// 年月日から Carbon オブジェクトを作成します，マイクロ秒を含む
carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 00:00:00.999999 +0800 CST
// 年月日から Carbon オブジェクトを作成します，ナノ秒を含む
carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

// 時分秒から Carbon オブジェクトを作成します(年月日のデフォルトは現在の年月日です)
carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
// 時分秒から Carbon オブジェクトを作成します(年月日のデフォルトは現在の年月日です)，ミリ秒を含む
carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
// 時分秒から Carbon オブジェクトを作成します(年月日のデフォルトは現在の年月日です)，マイクロ秒を含む
carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
// 時分秒から Carbon オブジェクトを作成します(年月日のデフォルトは現在の年月日です)，ナノ秒を含む
carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 時間文字列を Carbon オブジェクトに解析します

```go
carbon.Parse("").ToDateTimeString() // 空の文字列
carbon.Parse("0").ToDateTimeString() // 空の文字列
carbon.Parse("00:00:00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00").ToDateTimeString() // 空の文字列
carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空の文字列

carbon.Parse("now").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("yesterday").ToString() // 2020-08-04 13:14:15 +0800 CST
carbon.Parse("tomorrow").ToString() // 2020-08-06 13:14:15 +0800 CST

carbon.Parse("2020").ToString() // 2020-01-01 00:00:00 +0800 CST
carbon.Parse("2020-8").ToString() // 2020-08-01 00:00:00 +0800 CST
carbon.Parse("2020-08").ToString() // 2020-08-01 00:00:00 +0800 CST
carbon.Parse("2020-8-5").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-8-05").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-08-05").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("2020-08-05.999").ToString() // 2020-08-05 00:00:00.999 +0800 CST
carbon.Parse("2020-08-05.999999").ToString() // 2020-08-05 00:00:00.999999 +0800 CST
carbon.Parse("2020-08-05.999999999").ToString() // 2020-08-05 00:00:00.999999999 +0800 CST

carbon.Parse("2020-8-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-8-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-5 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

carbon.Parse("2020-8-5T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-8-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

carbon.Parse("20200805").ToString() // 2020-08-05 00:00:00 +0800 CST
carbon.Parse("20200805131415").ToString() // 2020-08-05 13:14:15 +0800 CST
carbon.Parse("20200805131415.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("20200805131415.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
carbon.Parse("20200805131415.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
carbon.Parse("20200805131415.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
carbon.Parse("20200805131415.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 文字をフォーマットして文字列を Carbon オブジェクトに解析します

```go
carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
```

##### レイアウト文字を使用して文字列を Carbon オブジェクトに解析します

```go
carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 始まりと終わり

```go
// 世紀の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// 世紀の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// 十年の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// 十年の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// 今年の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 今年の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 季度の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// 季度の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// 本月の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 本月の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 本周の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// 本周の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

// 本日の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// 本日の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// 時間の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// 時間の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// 分钟の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// 分钟の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// 本秒の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfSecond().ToString() // 2020-08-05 13:14:15 +0800 CST
// 本秒の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfSecond().ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
```

##### 追加と減らす

```go
// 三ヶ世紀を追加
carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 三ヶ世紀を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// 一ヶ世紀を追加
carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 一ヶ世紀を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// 三ヶ世紀を減らす
carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 三ヶ世紀を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// 一ヶ世紀を減らす
carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 一ヶ世紀を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

// 三ヶ年代を追加
carbon.Parse("2020-02-29 13:14:15").AddDecades(3).ToDateTimeString() // 2050-03-01 13:14:15
// 三ヶ年代を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// 一ヶ年代を追加
carbon.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// 一ヶ年代を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// 三ヶ年代を減らす
carbon.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// 三ヶ年代を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// 一ヶ年代を減らす
carbon.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// 一ヶ年代を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

// 三か年を追加
carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 三か年を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// 一か年を追加
carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 一か年を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// 三か年を減らす
carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 三か年を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// 一か年を減らす
carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 一か年を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// 三ヶ四半期を追加
carbon.Parse("2019-05-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2020-03-02 13:14:15
// 三ヶ四半期を追加(オーバーフローなし)
carbon.Parse("2019-05-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2020-02-29 13:14:15
// 一ヶ四半期を追加
carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 一ヶ四半期を追加(オーバーフローなし)
carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三ヶ四半期を減らす
carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 三ヶ四半期を減らす(オーバーフローなし)
carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// 一ヶ四半期を減らす
carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 一ヶ四半期を減らす(オーバーフローなし)
carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三ヶ月を追加
carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 三ヶ月を追加(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// 一ヶ月を追加
carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一ヶ月を追加(オーバーフローなし)
carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三ヶ月を減らす
carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 三ヶ月を減らす(オーバーフローなし)
carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// 一ヶ月を減らす
carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一か月を減らす(オーバーフローなし)
carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三か週間を追加
carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 一か週間を追加
carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// 三か週間を減らす
carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 一か週間を減らす
carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// 三か日間を追加
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 一か日間を追加
carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// 三か日間を減らす
carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 一か日間を減らす
carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// 三か時間を追加
carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 二か時間半を追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 一か時間を追加
carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// 三か時間を減らす
carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 二か時間半を減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 一か時間を減らす
carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 三か分钟を追加
carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 二か分钟半を追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 一か分钟を追加
carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// 三か分钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 二か分钟半を減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 一か分钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// 三か秒钟を追加
carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 二か秒钟半を追加
carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 一か秒钟を追加
carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// 三か秒钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 二か秒钟半を減らす
carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 一か秒钟を減らす
carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14

// 三かミリ秒を追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(3).ToString() // 2020-08-05 13:14:15.225222222 +0800 CST
// 一かミリ秒を追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMillisecond().ToString() // 2020-08-05 13:14:15.223222222 +0800 CST
// 三かミリ秒を減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(3).ToString() // 2020-08-05 13:14:15.219222222 +0800 CST
// 一かミリ秒を減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMillisecond().ToString() // 2020-08-05 13:14:15.221222222 +0800 CST

// 三か微妙を追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(3).ToString() // 2020-08-05 13:14:15.222225222 +0800 CST
// 一か微妙を追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddMicrosecond().ToString() // 2020-08-05 13:14:15.222223222 +0800 CST
// 三か微妙を減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(3).ToString() // 2020-08-05 13:14:15.222219222 +0800 CST
// 一か微妙を減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubMicrosecond().ToString() // 2020-08-05 13:14:15.222221222 +0800 CST

// 三かナノ秒を追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222225 +0800 CST
// 一かナノ秒を追加
carbon.Parse("2020-08-05 13:14:15.222222222").AddNanossecond().ToString() // 2020-08-05 13:14:15.222222223 +0800 CST
// 三かナノ秒を減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanosseconds(3).ToString() // 2020-08-05 13:14:15.222222219 +0800 CST
// 一かナノ秒を減らす
carbon.Parse("2020-08-05 13:14:15.222222222").SubNanossecond().ToString() // 2020-08-05 13:14:15.222222221 +0800 CST
```

##### 时间差異

```go
// 何年違いますか
carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
// 何年違いますか（絶対値）
carbon.Parse("2021-08-05 13:14:15").DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")) // 1

// 何ヶ月違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
// 何ヶ月違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMonths(carbon.Parse("2020-07-05 13:14:15")) // 1

// 何週間違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
// 何週間違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInWeeks(carbon.Parse("2020-07-28 13:14:15")) // 1

// 何日間違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
// 何日間違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInDays(carbon.Parse("2020-08-04 13:14:15")) // 1

// 何時間違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
// 何時間違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(carbon.Parse("2020-08-05 12:14:15")) // 1

// 何分違いますか
carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
// 何分違いますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(carbon.Parse("2020-08-05 13:13:15")) // 1

// 何秒の差がありますか
carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
// 何秒の差がありますか（絶対値）
carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(carbon.Parse("2020-08-05 13:14:14")) // 1

// 時間差を文字列で表す
carbon.Now().DiffInString() // just now
carbon.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
carbon.Now().SubYearsNoOverflow(1).DiffInString() // 1 year
// 時間差を文字列で表す（絶対値）
carbon.Now().DiffAbsInString(carbon.Now()) // just now
carbon.Now().AddYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year
carbon.Now().SubYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year

// 相差時間長
now := carbon.Now()
now.DiffInDuration(now).String() // 0s
now.AddHour().DiffInDuration(now).String() // 1h0m0s
now.SubHour().DiffInDuration(now).String() // -1h0m0s
// 相差時間長（絶対値）
now.DiffAbsInDuration(now).String() // 0s
now.AddHour().DiffAbsInDuration(now).String() // 1h0m0s
now.SubHour().DiffAbsInDuration(now).String() // 1h0m0s

// 人間に優しい読み取り可能なフォーマットの時間差を取得します
carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after
```

##### 时间极值

```go
c0 := carbon.Parse("2023-04-01")
c1 := carbon.Parse("2023-03-28")
c2 := carbon.Parse("2023-04-16")
// 最近のCarbonインスタンスを返す
c0.Closest(c1, c2) // c1
// 最も遠いCarbonインスタンスを返す
c0.Farthest(c1, c2) // c2

yesterday := carbon.Yesterday()
today     := carbon.Now()
tomorrow  := carbon.Tomorrow()
// 最大の Carbon インスタンスを返します
carbon.Max(yesterday, today, tomorrow) // tomorrow
// 最小の Carbon インスタンスを返します
carbon.Min(yesterday, today, tomorrow) // yesterday
```

##### 时间比較

```go
// 夏時間ですか
carbon.Parse("").IsDST() // false
carbon.Parse("0").IsDST() // false
carbon.Parse("0000-00-00 00:00:00").IsDST() // false
carbon.Parse("0000-00-00").IsDST() // false
carbon.Parse("00:00:00").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Brisbane").IsDST() // false
carbon.Parse("2023-01-01", "Australia/Sydney").IsDST() // true

// ゼロ値の時間ですか
carbon.Parse("").IsZero() // true
carbon.Parse("0").IsZero() // true
carbon.Parse("0000-00-00 00:00:00").IsZero() // true
carbon.Parse("0000-00-00").IsZero() // true
carbon.Parse("00:00:00").IsZero() // true
carbon.Parse("2020-08-05 00:00:00").IsZero() // false
carbon.Parse("2020-08-05").IsZero() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// 有効時間かどうか
carbon.Parse("").IsValid() // false
carbon.Parse("0").IsValid() // false
carbon.Parse("0000-00-00 00:00:00").IsValid() // false
carbon.Parse("0000-00-00").IsValid() // false
carbon.Parse("00:00:00").IsValid() // false
carbon.Parse("2020-08-05 00:00:00").IsValid() // true
carbon.Parse("2020-08-05").IsValid() // true
carbon.Parse("2020-08-05").SetTimezone("xxx").IsValid() // false

// 無効な時間ですか
carbon.Parse("").IsInvalid() // true
carbon.Parse("0").IsInvalid() // true
carbon.Parse("0000-00-00 00:00:00").IsInvalid() // true
carbon.Parse("0000-00-00").IsInvalid() // true
carbon.Parse("00:00:00").IsInvalid() // true
carbon.Parse("2020-08-05 00:00:00").IsInvalid() // false
carbon.Parse("2020-08-05").IsInvalid() // false
carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// 朝かどうかを判断する
carbon.Parse("2020-08-05 00:00:00").IsAM() // true
carbon.Parse("2020-08-05 08:00:00").IsAM() // true
carbon.Parse("2020-08-05 12:00:00").IsAM() // false
carbon.Parse("2020-08-05 13:00:00").IsAM() // false
// 午後かどうかを判断します
carbon.Parse("2020-08-05 00:00:00").IsPM() // false
carbon.Parse("2020-08-05 08:00:00").IsPM() // false
carbon.Parse("2020-08-05 12:00:00").IsPM() // true
carbon.Parse("2020-08-05 13:00:00").IsPM() // true

// 現在かどうか
carbon.Now().IsNow() // true
// 未来かどうか
carbon.Tomorrow().IsFuture() // true
// 過去かどうか
carbon.Yesterday().IsPast() // true

// 閏年かどうか
carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// 長年ですか
carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 一月ですか
carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 二月ですか
carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 三月ですか
carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
// 四月ですか
carbon.Parse("2020-08-05 13:14:15").IsApril()  // false
// 五月ですか
carbon.Parse("2020-08-05 13:14:15").IsMay() // false
// 六月ですか
carbon.Parse("2020-08-05 13:14:15").IsJune() // false
// 七月ですか
carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
// 八月ですか
carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
// 八月ですか
carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 十月ですか
carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
// 十一月ですか
carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
// 十二月ですか
carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

// 月曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
// 火曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 水曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 木曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
// 金曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
// 土曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 日曜日ですか
carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

// 営業日ですか
carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 週末ですか
carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 昨日ですか
carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
carbon.Parse("2020-08-04").IsYesterday() // true
// 今日ですか
carbon.Parse("2020-08-05 13:14:15").IsToday() // true
carbon.Parse("2020-08-05 00:00:00").IsToday() // true
carbon.Parse("2020-08-05").IsToday() // true
// 明日ですか
carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
carbon.Parse("2020-08-06").IsTomorrow() // true

// 同じ世紀ですか
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("3020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameCentury(carbon.Parse("2099-08-05 13:14:15")) // true
// 同じ時代ですか
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2030-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameDecade(carbon.Parse("2120-08-05 13:14:15")) // true
// 同じ年ですか
carbon.Parse("2020-08-05 00:00:00").IsSameYear(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameYear(carbon.Parse("2020-12-31 13:14:15")) // true
// 同じ季節ですか
carbon.Parse("2020-08-05 00:00:00").IsSameQuarter(carbon.Parse("2020-09-05 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameQuarter(carbon.Parse("2021-01-31 13:14:15")) // true
// 同じ月ですか
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2021-01-31 13:14:15")) // false
carbon.Parse("2020-01-01 00:00:00").IsSameMonth(carbon.Parse("2020-01-31 13:14:15")) // true
// 同じ日ですか
carbon.Parse("2020-08-05 13:14:15").IsSameDay(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 00:00:00").IsSameDay(carbon.Parse("2020-08-05 13:14:15")) // true
// 同じ時間ですか
carbon.Parse("2020-08-05 13:14:15").IsSameHour(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:00:00").IsSameHour(carbon.Parse("2020-08-05 13:14:15")) // true
// 同じ分ですか
carbon.Parse("2020-08-05 13:14:15").IsSameMinute(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:00").IsSameMinute(carbon.Parse("2020-08-05 13:14:15")) // true
// 同じ秒ですか
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2021-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").IsSameSecond(carbon.Parse("2020-08-05 13:14:15")) // true

// 大きいかどうか
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

// 小さいかどうか
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

// 等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

// と等しくない
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

// 大きいか等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

// 小きいか等しいかどうか
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

// 二つの時間の間に(この二つの時間は含まれていません)
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(開始時間も含めて)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(終了時間も含めて)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(この二つの時間を含めて)
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
```

> 長年の定義については、読んでください https://en.wikipedia.org/wiki/ISO_8601#Week_dates

##### 时间設定

```go
// タイムゾーンを設定
carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
carbon.SetTimezone(carbon.Tokyo).Now().SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// リージョンを設定
utc, _ := time.LoadLocation(carbon.UTC)
carbon.SetLocation(utc).Now().ToDateTimeString() // 2022-06-28 09:25:38
tokyo, _ := time.LoadLocation(carbon.Tokyo)
carbon.SetLocation(tokyo).Now().ToDateTimeString() // 2022-06-28 18:25:38

// ロケール設定
carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()) // 1 month ago
carbon.Parse("2020-07-05 13:14:15").SetLocale("jp").DiffForHumans() // 1 ヶ月前

// 年月日時分秒を設定する
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 2, 13, 14, 15).ToString() // 2019-02-02 13:14:15 +0800 CST
carbon.Parse("2020-01-01").SetDateTime(2019, 2, 31, 13, 14, 15).ToString() // 2019-03-03 13:14:15 +0800 CST
// 年月日時分秒ミリ秒を設定する
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 2, 13, 14, 15, 999).ToString() // 2019-02-02 13:14:15.999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 31, 13, 14, 15, 999).ToString() // 2019-03-03 13:14:15.999 +0800 CST
// 年月日時分秒微秒を設定する
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 2, 13, 14, 15, 999999).ToString() // 2019-02-02 13:14:15.999999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 31, 13, 14, 15, 999999).ToString() // 2019-03-03 13:14:15.999999 +0800 CST
// 年月日時分秒ナノ秒を設定する
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 2, 13, 14, 15, 999999999).ToString() // 2019-02-02 13:14:15.999999999 +0800 CST
carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 31, 13, 14, 15, 999999999).ToString() // 2019-03-03 13:14:15.999999999 +0800 CST

// 年月日を設定する
carbon.Parse("2020-01-01").SetDate(2019, 2, 2).ToString() // 2019-02-02 00:00:00 +0800 CST
carbon.Parse("2020-01-01").SetDate(2019, 2, 31).ToString() // 2019-03-03 00:00:00 +0800 CST
// 年月日ミリ秒を設定する
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 2, 999).ToString() // 2019-02-02 00:00:00.999 +0800 CST
carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 31, 999).ToString() // 2019-03-03 00:00:00.999 +0800 CST
// 年月日微秒を設定する
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 2, 999999).ToString() // 2019-02-02 00:00:00.999999 +0800 CST
carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 31, 999999).ToString() // 2019-03-03 00:00:00.999999 +0800 CST
// 年月日ナノ秒を設定する
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 2, 999999999).ToString() // 2019-02-02 00:00:00.999999999 +0800 CST
carbon.Parse("2020-01-01").SetDateNano(2019, 2, 31, 999999999).ToString() // 2019-03-03 00:00:00.999999999 +0800 CST

// 時分秒を設定する
carbon.Parse("2020-01-01").SetTime(13, 14, 15).ToString() // 2020-01-01 13:14:15 +0800 CST
carbon.Parse("2020-01-01").SetTime(13, 14, 90).ToString() // 2020-01-01 13:15:30 +0800 CST
// 時分秒ミリ秒を設定する
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0800 CST
carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 90, 999).ToString() // 2020-01-01 13:15:30.999 +0800 CST
// 時分秒微秒を設定する
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0800 CST
carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 90, 999999).ToString() // 2020-01-01 13:15:30.999999 +0800 CST
// 時分秒ナノ秒を設定する
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0800 CST
carbon.Parse("2020-01-01").SetTimeNano(13, 14, 90, 999999999).ToString() // 2020-01-01 13:15:30.999999999 +0800 CST

// 年を設定する
carbon.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// 年を設定する(オーバーフローなし)
carbon.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// 月を設定する
carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// 月を設定する(オーバーフローなし)
carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// 週の開始日を設定する
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6

// 日数を設定する
carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 時間を設定する
carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 分を設定する
carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 秒を設定する
carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

// ミリ秒を設定
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
carbon.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

// 微妙に設定
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

// ナノ秒を設定する
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
carbon.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
```

##### 时间取得

```go
// 本年の総日数を取得
carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 今月の総日数を取得
carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 本年の第数日を取得
carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 本年の第数週を取得
carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 今月の何日目（1から）を取得
carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 今月の何週目（1から）を取得
carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 今月の何週目（1から）を取得
carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 現在の年月日時分秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020,8,5,13,14,15
// 現在の年月日時分秒ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020,8,5,13,14,15,999
// 現在の年月日時分秒マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020,8,5,13,14,15,999999
// 現在の年月日時分秒ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020,8,5,13,14,15,999999999

// 現在の年月日を取得
carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020,8,5
// 現在の年月日ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020,8,5,999
// 現在の年月日マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020,8,5,999999
// 現在の年月日ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020,8,5,999999999

// 現在の時分秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13,14,15
// 現在の時分秒ミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13,14,15,999
// 現在の時分秒マイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13,14,15,999999
// 現在の時分秒ナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13,14,15,999999999

// 現在の世紀を取得
carbon.Parse("2020-08-05 13:14:15").Century() // 21
// 現在の年代を取得
carbon.Parse("2019-08-05 13:14:15").Decade() // 10
carbon.Parse("2021-08-05 13:14:15").Decade() // 20
// 現在の年を取得
carbon.Parse("2020-08-05 13:14:15").Year() // 2020
// 現在の四半期を取得
carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
// 現在の月を取得
carbon.Parse("2020-08-05 13:14:15").Month() // 8
// 現在の週を取得(0から開始)
carbon.Parse("2020-08-02 13:14:15").Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
// 現在の日数を取得
carbon.Parse("2020-08-05 13:14:15").Day() // 5
// 現在の時間を取得
carbon.Parse("2020-08-05 13:14:15").Hour() // 13
// 現在の分を取得
carbon.Parse("2020-08-05 13:14:15").Minute() // 14
// 現在の秒を取得
carbon.Parse("2020-08-05 13:14:15").Second() // 15
// 現在のミリ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
// 現在のマイクロ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
// 現在のナノ秒を取得
carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

// 秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
// ミリ秒のタイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").TimestampMilli() // 1596604455000
// マイクロ秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").TimestampMicro() // 1596604455000000
// ナノ秒タイムスタンプを取得
carbon.Parse("2020-08-05 13:14:15").TimestampNano() // 1596604455000000000

// タイムゾーン名を取得
carbon.SetTimezone(carbon.PRC).Timezone() // CST
carbon.SetTimezone(carbon.Tokyo).Timezone() // JST

// ロケーション名を取得
carbon.SetTimezone(carbon.PRC).Location() // PRC
carbon.SetTimezone(carbon.Tokyo).Location() // Asia/Tokyo

// UTCタイムゾーンからのオフセットを取得，単位秒
carbon.SetTimezone(carbon.PRC).Offset() // 28800
carbon.SetTimezone(carbon.Tokyo).Offset() // 32400

// ロケール名を取得
carbon.Now().Locale() // en
carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

// 星座を取得
carbon.Now().Constellation() // Leo
carbon.Now().SetLocale("en").Constellation() // Leo
carbon.Now().SetLocale("jp").Constellation() // 獅子座

// 季節を取得
carbon.Now().Season() // Summer
carbon.Now().SetLocale("en").Season() // Summer
carbon.Now().SetLocale("jp").Season() // 夏か

// 年齢を取得
carbon.Parse("2002-01-01 13:14:15").Age() // 17
carbon.Parse("2002-12-31 13:14:15").Age() // 18
```

##### 时间出力

```go
// 日期时间文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
// 日期时间文字列を出力，ミリ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
// 日期时间文字列を出力，マイクロ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
// 日期时间文字列を出力，ナノ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

// 略語日期时间文字列を出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
// 略語日期时间文字列を出力，ミリ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
// 略語日期时间文字列を出力，マイクロ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
// 略語日期时间文字列を出力，ナノ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

// 日期文字列を出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
// 日期文字列を出力，ミリ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
// 日期文字列を出力，マイクロ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
// 日期文字列を出力，ナノ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

// 略語日期文字列を出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
// 输出简写日期字符串，ミリ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
// 输出简写日期字符串，マイクロ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
// 输出简写日期字符串，ナノ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

// 時間文字列を出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
// 時間文字列を出力，ミリ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
// 時間文字列を出力，マイクロ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
// 時間文字列を出力，ナノ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

// 略語時間文字列を出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
// 略語時間文字列を出力，ミリ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
// 略語時間文字列を出力，マイクロ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
// 略語時間文字列を出力，ナノ秒を含む
carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

// Ansic フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
// Atom フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
// UnixDate フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
// RubyDate フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
// Kitchen フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
// Cookie フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
// DayDateTime フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
// RSS フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
// W3C フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00

// ISO8601 フォーマット文字列を出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+08:00
// ISO8601Milli フォーマット文字列を出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+08:00
// ISO8601Micro フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+08:00
// ISO8601Nano フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+08:00
// ISO8601Zulu フォーマット文字列を出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluString() // 2020-08-05T13:14:15Z
// ISO8601ZuluMilli フォーマット文字列を出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMilliString() // 2020-08-05T13:14:15.999Z
// ISO8601ZuluMicro フォーマット文字列を出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluMicroString() // 2020-08-05T13:14:15.999999Z
// ISO8601ZuluNano フォーマット文字列を出力 
carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601ZuluNanoString() // 2020-08-05T13:14:15.999999999Z

// RFC822 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
// RFC822Z フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
// RFC850 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
// RFC1036 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
// RFC1123 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
// RFC1123Z フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
// RFC2822 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
// RFC7231 フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT

// RFC3339 フォーマット文字列を出力
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+08:00
// RFC3339 フォーマット文字列を出力，ミリ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+08:00
// RFC3339 フォーマット文字列を出力，マイクロ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+08:00
// RFC3339 フォーマット文字列を出力，ナノ秒を含む
carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+08:00

// 日付時間文字列を出力
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

// "2006-01-02 15:04:05.999999999 -0700 MST" フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15.999999 +0800 CST

// "Jan 2, 2006" フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToFormattedDateString() // Aug 5, 2020
// "Mon, Jan 2, 2006" フォーマット文字列を出力
carbon.Parse("2020-08-05 13:14:15").ToFormattedDayDateString() // Wed, Aug 5, 2020

// レイアウトを指定する文字列を出力
carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout) // 2020-08-05T13:14:15+08:00
carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

// 指定されたフォーマットの文字列を出力
carbon.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
```

> もっとフォーマットした出力記号は付録を見てください <a href="#format-sign-table">書式設定記号表</a>

##### 星座

```go
// 星座を取得
carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

// 牡羊座ですか
carbon.Parse("2020-08-05 13:14:15").IsAries() // false
// おうし座ですか
carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 双子座ですか
carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
// かに座ですか
carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
// 獅子座ですか
carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
// おとめ座ですか
carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 天秤座ですか
carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
// さそり座ですか
carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 射手座ですか
carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 山羊座ですか
carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 水瓶座ですか
carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 魚座ですか
carbon.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### 季節

> 気象区分によると、3-5月は春で、6-8月は夏で、9-11月は秋で、12-2月は冬です

```go
// シーズンを取得
carbon.Parse("2020-08-05 13:14:15").Season() // Summer

// この季節の始まり
carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// この季節の終わり
carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// 春かどうか
carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
// 夏かどうか
carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
// 秋かどうか
carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 冬かどうか
carbon.Parse("2020-08-05 13:14:15").IsWinter() // false
```

##### JSON

###### シナリオ 1: すべての時刻フィールドが同じ形式である
```go
carbon.SetDefault(carbon.Default{
  Layout: carbon.DateTimeLayout,
})

type Person struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
  
  Field1 carbon.Carbon `json:"field1"`
  Field2 carbon.Carbon `json:"field2"`
  Field3 carbon.Carbon `json:"field3"`
  Field4 carbon.Carbon `json:"field4"`
  
  Field5 carbon.Carbon `json:"field5"`
  Field6 carbon.Carbon `json:"field6"`
  Field7 carbon.Carbon `json:"field7"`
  Field8 carbon.Carbon `json:"field8"`
}

now := carbon.Parse("2020-08-05 13:14:15", carbon.PRC)
person := Person {
  Name:   "gouguoyin",
  Age:    18,
  
  Field1: now,
  Field2: now,
  Field3: now,
  Field4: now,
  Field5: now,
  Field6: now,
  Field7: now,
  Field8: now,
}

data, marshalErr := json.Marshal(person)
if marshalErr != nil {
  // エラー処理...
  log.Fatal(marshalErr)
}
fmt.Printf("%s", data)
// 出力
{
  "name": "gouguoyin",
  "age": 18,
  "field1": "2020-08-05 13:14:15",
  "field2": "2020-08-05 13:14:15",
  "field3": "2020-08-05 13:14:15",
  "field4": "2020-08-05 13:14:15",
  "field5": "2020-08-05 13:14:15",
  "field6": "2020-08-05 13:14:15",
  "field7": "2020-08-05 13:14:15",
  "field8": "2020-08-05 13:14:15"
}

unmarshalErr := json.Unmarshal(data, &person)
if unmarshalErr != nil {
  // エラー処理...
  log.Fatal(unmarshalErr)
}

fmt.Printf("%s", person.Field1) // 2020-08-05 13:14:15
fmt.Printf("%s", person.Field2) // 2020-08-05 13:14:15
fmt.Printf("%s", person.Field3) // 2020-08-05 13:14:15
fmt.Printf("%s", person.Field4) // 2020-08-05 13:14:15

fmt.Printf("%s", person.Field5) // 2020-08-05 13:14:15
fmt.Printf("%s", person.Field6) // 2020-08-05 13:14:15
fmt.Printf("%s", person.Field7) // 2020-08-05 13:14:15
fmt.Printf("%s", person.Field8) // 2020-08-05 13:14:15
```

###### シナリオ 2: 異なる時刻フィールドは異なる形式を持つ

> サポートされているすべての type 値については、<a href="https://github.com/golang-module/carbon/blob/master/tag.go#L24">ここ</a>をクリックしてご覧ください。「carbon」タグが設定されていない場合、デフォルトは「layout:2006-01-02 15:04:05」です。「tz」タグが設定されていない場合、デフォルトは「Local」です。

```go
type Person struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
  
  Field1 carbon.Carbon `json:"field1"`
  
  Field2 carbon.Carbon `json:"field2" carbon:"type:date" tz:"PRC"`
  Field3 carbon.Carbon `json:"field3" carbon:"type:time" tz:"PRC"`
  Field4 carbon.Carbon `json:"field4" carbon:"type:dateTime" tz:"PRC"`
  // または
  Field2 carbon.Carbon `json:"field2" carbon:"layout:2006-01-02" tz:"PRC"`
  Field3 carbon.Carbon `json:"field3" carbon:"layout:15:04:05" tz:"PRC"`
  Field4 carbon.Carbon `json:"field4" carbon:"layout:2006-01-02 15:04:05" tz:"PRC"`
  // または
  Field2 carbon.Carbon `json:"field2" carbon:"format:Y-m-d" tz:"PRC"`
  Field3 carbon.Carbon `json:"field3" carbon:"format:H:i:s" tz:"PRC"`
  Field4 carbon.Carbon `json:"field4" carbon:"format:Y-m-d H:i:s" tz:"PRC"`
  
  Field5 carbon.Carbon `json:"field5" carbon:"type:timestamp" tz:"PRC"`
  Field6 carbon.Carbon `json:"field6" carbon:"type:timestampMilli" tz:"PRC"`
  Field7 carbon.Carbon `json:"field7" carbon:"type:timestampMicro" tz:"PRC"`
  Field8 carbon.Carbon `json:"field8" carbon:"type:timestampNano" tz:"PRC"`
}

now := Parse("2020-08-05 13:14:15", carbon.PRC)
person := Person {
  Name:   "gouguoyin",
  Age:    18,
  
  Field1: now,
  Field2: now,
  Field3: now,
  Field4: now,
  Field5: now,
  Field6: now,
  Field7: now,
  Field8: now,
}

loadErr := carbon.LoadTag(&person)
if loadErr != nil {
  // エラー処理...
  log.Fatal(loadErr)
}
data, marshalErr := json.Marshal(person)
if marshalErr != nil {
  // エラー処理...
  log.Fatal(marshalErr)
}
fmt.Printf("%s", data)
// 出力
{
  "name": "gouguoyin",
  "age": 18,
  "field1": "2020-08-05 13:14:15",
  "field2": "2020-08-05",
  "field3": "13:14:15",
  "field4": "2020-08-05 13:14:15",
  "field5": 1596604455,
  "field6": 1596604455999,
  "field7": 1596604455999999,
  "field8": 1596604455999999999
}

var person Person

loadErr := carbon.LoadTag(&person)
if loadErr != nil {
  // エラー処理...
  log.Fatal(loadErr)
}

unmarshalErr := json.Unmarshal(data, &person)
if unmarshalErr != nil {
  // エラー処理...
  log.Fatal(unmarshalErr)
}

fmt.Printf("%s", person.Field1) // 2002-08-05 13:14:15
fmt.Printf("%s", person.Field2) // 2020-08-05
fmt.Printf("%s", person.Field3) // 13:14:15
fmt.Printf("%s", person.Field4) // 2002-08-05 13:14:15

fmt.Printf("%d", person.Field5) // 1596604455
fmt.Printf("%d", person.Field6) // 1596604455999
fmt.Printf("%d", person.Field7) // 1596604455999999
fmt.Printf("%d", person.Field8) // 1596604455999999999

```

##### カレンダ＃カレンダ＃

現在サポートされているカレンダーは

* [儒略の日/簡略化儒略の日](./calendar/julian/README.jp.md "儒略日/简化儒略日")
* [中国の旧暦](./calendar/lunar/README.jp.md "中国の旧暦")
* [ペルシャ暦/イラン暦](./calendar/persian/README.jp.md "ペルシャ暦/イラン暦")

##### 国際化

現在サポートされている言語

* [简体中国語(zh-CN)](./lang/zh-CN.json "简体中国語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") から翻訳されます
* [繁体中国語(zh-TW)](./lang/zh-TW.json "繁体中国語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") から翻訳されます
* [英語(en)](./lang/en.json "英語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") から翻訳されます
* [日本語(jp)](./lang/jp.json "日本語")：[gouguoyin](https://github.com/gouguoyin "gouguoyin") から翻訳されます
* [韓国語(kr)](./lang/kr.json "韓国語")：[benzammour](https://github.com/benzammour "benzammour") から翻訳されます
* [ドイツ語(de)](./lang/de.json "ドイツ語")：[benzammour](https://github.com/benzammour "benzammour") から翻訳されます
* [スペイン語(es)](./lang/es.json "スペイン語")：[hgisinger](https://github.com/hgisinger "hgisinger") から翻訳されます
* [トルコ語(tr)](./lang/tr.json "トルコ語")：[emresenyuva](https://github.com/emresenyuva "emresenyuva") から翻訳されます
* [ポルトガル語(pt)](./lang/pt.json "ポルトガル語")：[felipear89](https://github.com/felipear89 "felipear89") から翻訳されます
* [ロシア語(ru)](./lang/ru.json "ロシア語")：[zemlyak](https://github.com/zemlyak "zemlyak") から翻訳されます
* [ウクライナ語(uk)](./lang/uk.json "ウクライナ語")：[open-git](https://github.com/open-git "open-git") から翻訳されます
* [ルーマニア語(ro)](./lang/ro.json "ルーマニア語"): [DrOctavius](https://github.com/DrOctavius "DrOctavius") から翻訳されます
* [インドネシア語(id)](./lang/id.json "インドネシア語"): [justpoypoy](https://github.com/justpoypoy "justpoypoy") から翻訳されます
* [イタリア語(it)](./lang/it.json "イタリア語"): [nicoloHevelop](https://github.com/hollowaykeanho "nicoloHevelop") から翻訳されます
* [マレーシアバハマ語(ms-MY)](./lang/ms-MY.json "マレーシアバハマ語"): [hollowaykeanho](https://github.com/hollowaykeanho "hollowaykeanho") から翻訳されます
* [フランス語(fr)](./lang/fr.json "フランス語"): [hollowaykeanho](https://github.com/hollowaykeanho "hollowaykeanho") から翻訳されます
* [タイ語(th)](./lang/th.json "タイ語"): [izcream](https://github.com/izcream "izcream") から翻訳されます
* [スウェーデン語(se)](./lang/se.json "スウェーデン語"):  [jwanglof](https://github.com/jwanglof "jwanglof") から翻訳されます
* [ペルシア語(fa)](./lang/fa.json "ペルシア語"):  [Iranian](https://github.com/Iranian "Iranian") から翻訳されます
* [ポーランド語(nl)](./lang/nl.json "ポーランド語"): [RemcoE33](https://github.com/RemcoE33 "RemcoE33") から翻訳されます
* [ベトナム語(vi)](./lang/vi.json "ベトナム語"): [culy247](https://github.com/culy247 "culy247") から翻訳されます
* [ヒンディー語(hi)](./lang/hi.json "ヒンディー語"): [chauhan17nitin](https://github.com/chauhan17nitin "chauhan17nitin") から翻訳されます
* [ポーランド語(pl)](./lang/pl.json "ポーランド語"): [gouguoyin](https://github.com/gouguoyin "gouguoyin") から翻訳されます

現在サポートされている方法

* `Constellation()`：星座を取得，例えば `おひつじ座`
* `Season()`：シーズンを取得，例えば `春`
* `DiffForHumans()`：人間に優しい読み取り可能なフォーマットの時間差を取得します，例えば `一時間前`
* `ToMonthString()`：月文字列を出力，例えば `いちがつ`
* `ToShortMonthString()`：略語月文字列を出力，例えば `1がつ`
* `ToWeekString()`：週文字列を出力，例えば `日曜日`
* `ToShortWeekString()`：略語週文字列を出力，例えば `日`

###### エリアの設定

```go
lang := carbon.NewLanguage()
lang.SetLocale("jp")

c := carbon.SetLanguage(lang)
if c.Error != nil {
  // エラー処理
  log.Fatal(c.Error)
}

c.Now().AddHours(1).DiffForHumans() // 1 時間後
c.Now().AddHours(1).ToMonthString() // はちがつ
c.Now().AddHours(1).ToShortMonthString() // 8がつ
c.Now().AddHours(1).ToWeekString() // もくようび
c.Now().AddHours(1).ToShortWeekString() // 週三
c.Now().AddHours(1).Constellation() // ししざ
c.Now().AddHours(1).Season() // なつ
```

###### 翻訳リソースの一部を書き換える(残りはまだ指定された `locale` ファイルの内容によって翻訳されます)

```go
lang := carbon.NewLanguage()

resources := map[string]string {
  "hour": "%dh",
}
lang.SetLocale("en").SetResources(resources)

c := carbon.SetLanguage(lang)
if c.Error != nil {
  // エラー処理
  log.Fatal(c.Error)
}

c.Now().AddYears(1).DiffForHumans() // 1 year from now
c.Now().AddHours(1).DiffForHumans() // 1h from now
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

###### すべての翻訳リソースを書き換えます

```go
lang := carbon.NewLanguage()
resources := map[string]string {
  "months": "january|february|march|april|may|june|july|august|september|october|november|december",
  "short_months": "jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec",
  "weeks": "sunday|monday|tuesday|wednesday|thursday|friday|saturday",
  "short_weeks": "sun|mon|tue|wed|thu|fri|sat",
  "seasons": "spring|summer|autumn|winter",
  "constellations": "aries|taurus|gemini|cancer|leo|virgo|libra|scorpio|sagittarius|capricornus|aquarius|pisce",
  "year": "1 yr|%d yrs",
  "month": "1 mo|%d mos",
  "week": "%dw",
  "day": "%dd",
  "hour": "%dh",
  "minute": "%dm",
  "second": "%ds",
  "now": "just now",
  "ago": "%s ago",
  "from_now": "in %s",
  "before": "%s before",
  "after": "%s after",
}
lang.SetResources(resources)

c := carbon.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // in 1 yr
c.Now().AddHours(1).DiffForHumans() // in 1h
c.Now().ToMonthString() // august
c.Now().ToShortMonthString() // aug
c.Now().ToWeekString() // tuesday
c.Now().ToShortWeekString() // tue
c.Now().Constellation() // leo
c.Now().Season() // summer
```

##### 模擬テスト

```go
c := carbon.SetTimezone(carbon.UTC)

c.Now().ToDateString() // 2023-12-27
c.Now().IsSetTestNow() // false

c.SetTestNow(carbon.Parse("2020-08-05"))
c.Now().ToDateString() // 2020-08-05
c.Now().IsSetTestNow() // true

c.UnSetTestNow()
c.Now().ToDateString() // 2023-12-27
c.Now().IsSetTestNow() // false
```

##### エラー処理

> 複数のエラーが発生した場合、最初のエラーだけを返します。前のエラーは削除された後、次のエラーに戻ります

```go
c := carbon.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
  // エラー処理...
  log.Fatal(c.Error)
}
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

#### 付録

##### <a id="format-sign-table">書式設定記号表</a>

|     記号            | 説明            | 長さ | 範囲            |  例            |
| :------------: | :------------: |:------:|:----------------:| :------------: |
| d | 月の中の何日目ですか | 2 | 01-31 | 02 |
| D | 略語は何曜日を表しますか | 3 | Mon-Sun | Mon |
| j | 月の中の何日目ですか | - | 1-31 | 2 |
| S | 何日目の英語の略語の接尾語，普通はjと協力して使います | 2 | st/nd/rd/th | th |
| l | 完全な単語は何曜日を表しますか | - | Monday-Sunday | Monday |
| F | 完全な単語は月を表しますか | - | January-December | January |
| m | 数字が示す月は | 2 | 01-12 | 01 |
| M | 略語の月 | 3 | Jan-Dec | Jan |
| n | 数字が示す月 | - | 1-12 | 1 |
| Y | 数字が示す年 | 4 | 0000-9999 | 2006 |
| y | 数字が示す年 | 2 | 00-99 | 06 |
| a | 小文字の午前と午後の標識 | 2 | am/pm | pm |
| A | 大文字の午前と午後の表示 | 2 | AM/PM | PM |
| g | 時間, 12時間のフォーマット | - | 1-12 | 3 |
| G | 時間, 24時間のフォーマット | - | 0-23 | 15 |
| h | 時間, 12時間のフォーマット | 2 | 00-11 | 03 |
| H | 時間, 24時間のフォーマット | 2 | 00-23 | 15 |
| i | 分 | 2 | 01-59 | 04 |
| s | 秒 | 2 | 01-59 | 05 |
| O | グリニッジとの時間差の時間数 | - | - | -0700 |
| P | グリニッジと時間の差の時間数, 時間と分の間にコロンがあります | - | - | -07:00 |
| T | タイムゾーンの略語 | - | - | MST |
| W | ISO8601 フォーマットの数字は年の中の第数週を表します | 2 | 1-52 | 01 |
| N | ISO8601 フォーマットの数字は曜日の中の何日目を表しますか | 2 | 01-07 | 02 |
| L | うるう年かどうか, うるう年が1であれば, 0です | 1 | 0-1 | 0 |
| U | 秒タイムスタンプを | - | - | 1596604455 |
| V | ミリ秒のタイムスタンプを | - | - | 1596604455666 |
| X | マイクロ秒タイムスタンプを | - | - | 1596604455666666 |
| Z | ナノ秒タイムスタンプを | - | - | 1596604455666666666 |
| v | ミリ秒 | - | 1-999 | 999 |
| u | マイクロ秒| - | 1-999999 | 999999 |
| x | ナノ秒 | - | 1-999999999 | 999999999 |
| w | 数字の表示の曜日 | 1 | 0-6 | 1 |
| t | 月の総日数 | 2 | 28-31 | 31 |
| z | 年の中の何日目 | - | 1-365 | 2 |
| e | 位置 | - | - | America/New_York |
| Q | 季節 | 1 | 1-4 | 1 |
| C | 世紀 | - | 0-99 | 21 |

#### 人気のある問題

1、v1とv2のバージョンの違いは何ですか？
> APIのv1バージョンとv2バージョンに違いはありませんが `language.go`
> での翻訳リソースファイルの実装は異なります。v1は、サードパーティの拡張ライブラリ [packr](https://github.com/gobuffalo/packr) によって実装されています，v2は、 `golang1.16`
> の後に組み込みの標準ライブラリ [embed](https://pkg.go.dev/embed) によって実装されています。goバージョンが1.16より大きい場合は、v2バージョンを使用することをお勧めします。それ以外の場合は、v1バージョンを使用する必要があります。

#### 参考文献

* [briannesbitt/carbon](https://github.com/briannesbitt/Carbon)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [jodaOrg/joda-time](https://github.com/jodaOrg/joda-time)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)

#### 寄稿者
`Carbon` に貢献してくれた以下のすべてに感謝します：

<a href="https://github.com/golang-module/carbon/graphs/contributors"><img src="https://contrib.rocks/image?repo=golang-module/carbon&max=100&columns=16"/></a>

#### スポンサー

`Carbon` は非営利のオープンソースプロジェクトです，`Carbon` をサポートしたい場合は、開発者のために [コーヒーを1杯購入](https://opencollective.com/go-carbon) できます

#### ありがとう

`Carbon` は無料の JetBrains オープンソースライセンスを取得しました，これに感謝します

<a href="https://www.jetbrains.com"><img src="https://foruda.gitee.com/images/1704325523163241662/1bf21f86_544375.png" height="100" alt="JetBrains"/></a>
