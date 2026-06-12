# Benchmark 测试报告

> 测试日期：2026-06-12
> 测试环境：Linux (amd64), Intel(R) Xeon(R) Platinum

## 测试结果汇总

| 包 | 函数 | 耗时 | 内存 | 分配次数 |
|----|------|------|------|----------|
| **bytes** | HumanReadable | 91.69 ns | 8 B | 2 |
| **bytes** | HexEncode | 6094 ns | 4096 B | 2 |
| **bytes** | HexDecode | 2701 ns | 1024 B | 1 |
| **bytes** | Base64Encode | 4176 ns | 2816 B | 2 |
| **bytes** | Base64EncodeString | 231.1 ns | 128 B | 2 |
| **bytes** | Base64Decode | 176.8 ns | 48 B | 1 |
| **bytes** | Base64URLEncodeString | 96.25 ns | 32 B | 1 |
| **bytes** | Base64URLDecode | 115.3 ns | 24 B | 1 |
| **coalesce** | Coalesce | 4.38 ns | 0 B | 0 |
| **coalesce** | CoalesceString | 4.23 ns | 0 B | 0 |
| **crypto** | MD5 | 241.8 ns | 32 B | 1 |
| **crypto** | SHA1 | 555.4 ns | 96 B | 2 |
| **crypto** | SHA256 | 543.8 ns | 128 B | 2 |
| **env** | Get | 53.52 ns | 0 B | 0 |
| **env** | Has | 50.49 ns | 0 B | 0 |
| **env** | EnvironMap | 12732 ns | 6520 B | 52 |
| **errors** | New | 0.47 ns | 0 B | 0 |
| **errors** | Wrap | 307.0 ns | 56 B | 2 |
| **errors** | Is | 19.43 ns | 0 B | 0 |
| **file** | Exists | 1407 ns | 232 B | 2 |
| **file** | Read | 11288 ns | 840 B | 5 |
| **file** | Write | 858790 ns | 120 B | 3 |
| **file** | Copy | 861989 ns | 514 B | 12 |
| **fn** | And | 5.16 ns | 0 B | 0 |
| **fn** | Or | 5.33 ns | 0 B | 0 |
| **fn** | Not | 0.47 ns | 0 B | 0 |
| **logs** | Log | 1508 ns | 561 B | 5 |
| **logs** | LogWithCtx | 1323 ns | 402 B | 5 |
| **logs** | DefaultLogger | 1312 ns | 417 B | 5 |
| **maps** | Keys | 26924 ns | 8192 B | 1 |
| **maps** | Values | 26470 ns | 8192 B | 1 |
| **maps** | HasKey | 12.76 ns | 0 B | 0 |
| **maps** | Merge | 177715 ns | 73888 B | 9 |
| **maths** | Min | 0.47 ns | 0 B | 0 |
| **maths** | Max | 0.47 ns | 0 B | 0 |
| **maths** | Clamp | 0.47 ns | 0 B | 0 |
| **maths** | Abs | 0.46 ns | 0 B | 0 |
| **maths** | Sum | 6.52 ns | 0 B | 0 |
| **maths** | Average | 4.04 ns | 0 B | 0 |
| **maths** | Pow | 5.97 ns | 0 B | 0 |
| **maths** | IsPrime | 346.5 ns | 0 B | 0 |
| **maths** | GCD | 168.7 ns | 0 B | 0 |
| **maths** | Fibonacci | 39.72 ns | 0 B | 0 |
| **rand** | String | 298.3 ns | 16 B | 1 |
| **rand** | Int | 14.07 ns | 0 B | 0 |
| **rand** | Bytes | 946.4 ns | 0 B | 0 |
| **rand** | Choice | 15.22 ns | 0 B | 0 |
| **rand** | Shuffle | 14676 ns | 0 B | 0 |
| **set** | Add | 373.7 ns | 53 B | 0 |
| **set** | Contains | 12.95 ns | 0 B | 0 |
| **set** | Union | 82844 ns | 36944 B | 5 |
| **set** | Intersection | 125369 ns | 37512 B | 17 |
| **set** | ToSlice | 27077 ns | 8192 B | 1 |
| **slices** | Dedup | 973.8 ns | 440 B | 4 |
| **slices** | Filter | 9647 ns | 8192 B | 1 |
| **slices** | Map | 8464 ns | 8192 B | 1 |
| **slices** | Contains | 476.7 ns | 0 B | 0 |
| **slices** | Split | 3476 ns | 2688 B | 1 |
| **slices** | GroupBy | 46790 ns | 9952 B | 17 |
| **str** | Upper | 263.7 ns | 48 B | 1 |
| **str** | Contains | 11.09 ns | 0 B | 0 |
| **str** | Split | 563.1 ns | 320 B | 1 |
| **str** | Ellipsis | 23.76 ns | 0 B | 0 |
| **times** | Format | 269.2 ns | 24 B | 1 |
| **times** | BeginOfDay | 38.90 ns | 0 B | 0 |
| **times** | EndOfDay | 38.90 ns | 0 B | 0 |
| **times** | BeginOfMonth | 40.20 ns | 0 B | 0 |
| **times** | DaysBetween | 25.45 ns | 0 B | 0 |
| **times** | IsWeekend | 9.53 ns | 0 B | 0 |
| **validation** | IsEmail | 618.6 ns | 0 B | 0 |
| **validation** | IsPhone | 187.2 ns | 0 B | 0 |
| **validation** | IsURL | 926.6 ns | 0 B | 0 |
| **validation** | IsIP | 46.26 ns | 0 B | 0 |
| **validation** | IsMatch | 308.1 ns | 0 B | 0 |
| **validation** | IsUUID | 568.8 ns | 0 B | 0 |
| **validation** | IsJSON | 1606 ns | 576 B | 12 |
| **validation** | IsCreditCard | 337.6 ns | 0 B | 0 |

---

## 性能评级

### 极快 (< 50ns)
| 函数 | 耗时 |
|------|------|
| errors.New | 0.47 ns |
| fn.Not | 0.47 ns |
| maths.Min | 0.47 ns |
| maths.Max | 0.47 ns |
| maths.Clamp | 0.47 ns |
| maths.Abs | 0.46 ns |
| coalesce.Coalesce | 4.38 ns |
| fn.And | 5.16 ns |
| fn.Or | 5.33 ns |
| maths.Sum | 6.52 ns |
| times.IsWeekend | 9.53 ns |
| str.Contains | 11.09 ns |
| maps.HasKey | 12.76 ns |
| set.Contains | 12.95 ns |
| rand.Int | 14.07 ns |
| rand.Choice | 15.22 ns |
| errors.Is | 19.43 ns |
| str.Ellipsis | 23.76 ns |
| times.DaysBetween | 25.45 ns |
| times.BeginOfDay | 38.90 ns |
| times.EndOfDay | 38.90 ns |
| times.BeginOfMonth | 40.20 ns |
| validation.IsIP | 46.26 ns |
| env.Has | 50.49 ns |
| env.Get | 53.52 ns |

### 快速 (50ns - 500ns)
| 函数 | 耗时 |
|------|------|
| bytes.Base64URLEncodeString | 96.25 ns |
| bytes.Base64URLDecode | 115.3 ns |
| maths.GCD | 168.7 ns |
| bytes.Base64Decode | 176.8 ns |
| validation.IsPhone | 187.2 ns |
| bytes.Base64EncodeString | 231.1 ns |
| crypto.MD5 | 241.8 ns |
| str.Upper | 263.7 ns |
| times.Format | 269.2 ns |
| rand.String | 298.3 ns |
| errors.Wrap | 307.0 ns |
| validation.IsMatch | 308.1 ns |
| validation.IsCreditCard | 337.6 ns |
| maths.IsPrime | 346.5 ns |
| set.Add | 373.7 ns |
| slices.Contains | 476.7 ns |

### 正常 (500ns - 5μs)
| 函数 | 耗时 |
|------|------|
| crypto.SHA1 | 555.4 ns |
| crypto.SHA256 | 543.8 ns |
| validation.IsUUID | 568.8 ns |
| str.Split | 563.1 ns |
| validation.IsEmail | 618.6 ns |
| validation.IsURL | 926.6 ns |
| rand.Bytes | 946.4 ns |
| slices.Dedup | 973.8 ns |

### 较慢 (> 5μs)
| 函数 | 耗时 |
|------|------|
| validation.IsJSON | 1.6 μs |
| file.Exists | 1.4 μs |
| bytes.HexDecode | 2.7 μs |
| slices.Split | 3.5 μs |
| bytes.Base64Encode | 4.2 μs |
| bytes.HexEncode | 6.1 μs |
| slices.Map | 8.5 μs |
| slices.Filter | 9.6 μs |
| file.Read | 11.3 μs |
| env.EnvironMap | 12.7 μs |
| rand.Shuffle | 14.7 μs |
| maps.Keys | 26.9 μs |
| maps.Values | 26.5 μs |
| set.ToSlice | 27.1 μs |
| slices.GroupBy | 46.8 μs |
| set.Union | 82.8 μs |
| set.Intersection | 125.4 μs |
| maps.Merge | 177.7 μs |
| file.Write | 858.8 μs |
| file.Copy | 862.0 μs |

---

## 优化历史

### v1.0.3 优化
- **maps.Merge**: 29 次分配 → 9 次分配（预分配总容量）
- **slices.GroupBy**: 两遍遍历预分配每个 slice 容量

### v1.0.4 优化
- **validation.IsMatch**: 2518 ns / 10 次分配 → 303 ns / 0 次分配（引入 sync.Map 正则缓存，8x 提速）
- **slices.Split**: 预分配结果切片容量，减少扩容分配
- **slices**: 新增 Benchmark 覆盖

### v1.2.0 新增
- **set** 包：泛型集合实现，包含 Union / Intersection / Difference 等
- **coalesce** 包：取首非零值，4.38 ns，零分配
- **rand** 包：随机工具，String 298.3 ns / 1 次分配
- **fn** 包：函数式辅助，零分配
- **file** 包：文件操作（I/O 密集型，Write/Copy 约 860 μs）
- **env** 包：环境变量工具，Get/Has 约 50 ns，EnvironMap 约 12.7 μs / 52 次分配
- **validation** 新增：IsUUID、IsCreditCard、IsJSON 等
- **times** 大幅扩充：BeginOfDay 38.9 ns、IsWeekend 9.53 ns 等

---

## 优化建议

### 低优先级

1. **logs** - 每次 5 次分配，可考虑使用 sync.Pool
2. **maps.Keys/Values** - 固定 8KB 分配
3. **bytes.HexEncode** - 固定 4KB 分配
4. **env.EnvironMap** - 52 次分配，可考虑预分配

---

## 结论

整体性能表现良好：
- 高频使用的工具函数（Contains、HasKey、Is 等）都在 ns 级别
- 新增的 coalesce、fn 包接近零成本抽象
- 加密函数性能符合预期
- set 包集合运算在大数据量下表现合理（+ 预分配优化空间）
- I/O 操作（file.Write/Copy）受磁盘性能主导
