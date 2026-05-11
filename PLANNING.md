# gopkg 功能规划

## 项目现状

已完成全部规划功能：

| 包 | 功能 | 状态 |
|---|---|---|
| `maths` | Min、Max、Clamp、Abs、Sum、Average、Pow、IsPrime、GCD、LCM、Fibonacci | ✅ 已实现 |
| `conv` | 指针转换、str子包、json子包 | ✅ 已实现 |
| `slices` | Split、Dedup、Filter、Map、FlatMap、Merge、Contains、Index、Reverse、GroupBy | ✅ 已实现 |
| `maps` | Keys、Values、Filter、Map、Merge、Clone、Get、Set、HasKey | ✅ 已实现 |
| `str` | Trim、Upper、Lower、Split、Join、Contains、Ellipsis等 | ✅ 已实现 |
| `errors` | New、Wrap、Wrapf、Is、As、Join | ✅ 已实现 |
| `validation` | IsEmail、IsPhone、IsURL、IsIP、IsAlpha等 | ✅ 已实现 |
| `net/ip` | IsValid、IsPrivate、ToInt、IntToIP | ✅ 已实现 |
| `net/url` | Parse、ParseQuery、Build、GetHost、Encode等 | ✅ 已实现 |
| `bytes` | Hex、Base64、HumanReadable | ✅ 已实现 |
| `crypto` | MD5、SHA1、SHA256 | ✅ 已实现 |
| `logs` | ctx支持、多级别、Formatter | ✅ 已实现 |
| `sorts` | Sort、SortFunc、SortFuncStable | ✅ 已实现 |
| `test` | AssertEqual、AssertTrue、AssertSliceEqual、ContainsStr | ✅ 已实现 |

---

## 代码风格规范

1. 使用 Go 1.18+ 泛型
2. 函数简短、直接
3. 包名简短（如 `conv`, `str` 而非 `conversion`, `string`）
4. 错误返回原始错误或使用 sentinel error
5. 避免过度封装，只做必要的简化
6. 每个功能配对应测试和 benchmark
7. 测试辅助函数统一放在 `test` 包
