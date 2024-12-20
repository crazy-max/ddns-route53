# 波斯历(伊朗历)

简体中文 | [English](README.md) | [日本語](README.jp.md)

#### 用法示例

##### 将 `公历` 转换成 `波斯历`

```go
// 获取波斯历年份
carbon.Parse("2020-08-05 13:14:15").Persian().Year() // 1399
// 获取波斯历月份
carbon.Parse("2020-08-05 13:14:15").Persian().Month() // 5
// 获取波斯历日期
carbon.Parse("2020-08-05 13:14:15").Persian().Day() // 15
// 获取波斯历小时
carbon.Parse("2020-08-05 13:14:15").Persian().Hour() // 13
// 获取波斯历分钟
carbon.Parse("2020-08-05 13:14:15").Lunar().Minute() // 14
// 获取波斯历秒数
carbon.Parse("2020-08-05 13:14:15").Lunar().Second() // 15

// 获取波斯历日期时间字符串
carbon.Parse("2020-08-05 13:14:15").Lunar().String() // 1399-05-15 13:14:15
fmt.Printf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 1399-05-15 13:14:15
// 获取波斯历月字符串
carbon.Parse("2020-08-05 13:14:15").Persian().ToMonthString() // مرداد
// 获取波斯历周字符串
carbon.Parse("2020-08-05 13:14:15").Persian().ToWeekString() // چهارشنبه

// 是否是波斯历零值时间
carbon.Parse("0000-00-00 00:00:00").Persian().IsZero() // true
carbon.Parse("2020-08-05 13:14:15").Persian().IsZero() // false

// 是否是波斯历闰年
carbon.Parse("2016-03-20 00:00:00").Persian().IsLeapYear() // true
carbon.Parse("2020-08-05 13:14:15").Persian().IsLeapYear() // false
```

##### 将 `波斯历` 转化成 `公历`

```go
carbon.CreateFromPersian(1395, 1, 1, 0, 0, 0).ToDateTimeString() // 2016-03-20 00:00:00
carbon.CreateFromPersian(1399, 5, 15, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
```