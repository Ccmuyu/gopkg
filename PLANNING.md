# gopkg 功能规划

## 项目现状

已完成全部规划功能：

| 包 | 功能 | 状态 |
|---|---|---|
| `maths` | Min、Max、Clamp、Abs、Sum、Average、Pow、IsPrime、GCD、LCM、Fibonacci | ✅ 已实现 |
| `maths` | Round、Floor、Ceil、MinSlice、MaxSlice、Factorial、Median | ✅ 已实现 |
| `conv` | 指针转换、str子包、json子包 | ✅ 已实现 |
| `slices` | Split、Dedup、Filter、Map、FlatMap、Merge、Contains、Index、Reverse、GroupBy | ✅ 已实现 |
| `slices` | Reduce、Some、Every、None、Without、Intersection、Union、Difference、Chunk | ✅ 已实现 |
| `slices` | Sort、Fill、Take、Drop | ✅ 已实现 |
| `maps` | Keys、Values、Filter、Map、Merge、Clone、Get、Set、HasKey | ✅ 已实现 |
| `str` | Trim、Upper、Lower、Split、Join、Contains、Ellipsis等 | ✅ 已实现 |
| `str` | After、Before、Between、Reverse、PadLeft、PadRight、Truncate、Capitalize、ToCamel、ToSnake、ToKebab、Mask | ✅ 已实现 |
| `errors` | New、Wrap、Wrapf、Is、As、Join | ✅ 已实现 |
| `errors` | MultiError、PanicToError | ✅ 已实现 |
| `validation` | IsEmail、IsPhone、IsURL、IsIP、IsAlpha等 | ✅ 已实现 |
| `validation` | IsUUID、IsMAC、IsJSON、IsHexColor、IsPostalCode、IsPort、IsLatitude、IsLongitude、IsCreditCard | ✅ 已实现 |
| `net/ip` | IsValid、IsPrivate、ToInt、IntToIP | ✅ 已实现 |
| `net/url` | Parse、ParseQuery、Build、GetHost、Encode等 | ✅ 已实现 |
| `bytes` | Hex、Base64、HumanReadable | ✅ 已实现 |
| `crypto` | MD5、SHA1、SHA256 | ✅ 已实现 |
| `crypto` | SHA512、SHA512Bytes、HMAC、CRC32、CRC32String、RandomToken | ✅ 已实现 |
| `logs` | ctx支持、多级别、Formatter | ✅ 已实现 |
| `sorts` | Sort、SortFunc、SortFuncStable | ✅ 已实现 |
| `times` | Format、ParseUTCTime、BeginOfDay、EndOfDay、BeginOfWeek、EndOfWeek、BeginOfMonth、EndOfMonth、BeginOfYear、EndOfYear、IsWeekend、IsLeapYear、DaysBetween等 | ✅ 已实现 |
| `set` | New、Add、Remove、Contains、ToSlice、Union、Intersection、Difference、ForEach | ✅ 已实现 |
| `coalesce` | Coalesce、CoalesceSlice | ✅ 已实现 |
| `rand` | String、Int、Bytes、Choice、Shuffle | ✅ 已实现 |
| `fn` | And、Or、Not、Tuple2、Tuple3 | ✅ 已实现 |
| `file` | Exists、IsDir、IsFile、Read、Write、Copy、Move、Remove、ListDir、TempDir、TempFile、Ext、Basename、Dir | ✅ 已实现 |
| `env` | Get、MustGet、Set、Unset、Has、EnvironMap | ✅ 已实现 |
| `retry` | Retry、WithMaxAttempts、WithDelay、WithBackoff、WithJitter | ✅ 已实现 |
| `rate` | 滑动窗口限流、三级规则、Peek/Remaining、FailClosed、HTTP Middleware、Redis Store | ✅ 已实现 |
| `test` | AssertEqual、AssertTrue、AssertSliceEqual、ContainsStr、SliceContains | ✅ 已实现 |
| `test` | AssertNil、AssertNotNil、AssertError、AssertNoError、AssertPanic、AssertMatch | ✅ 已实现 |

---

## 代码风格规范

1. 使用 Go 1.18+ 泛型
2. 函数简短、直接
3. 包名简短（如 `conv`, `str` 而非 `conversion`, `string`）
4. 错误返回原始错误或使用 sentinel error
5. 避免过度封装，只做必要的简化
6. 每个功能配对应测试和 benchmark
7. 测试辅助函数统一放在 `test` 包
