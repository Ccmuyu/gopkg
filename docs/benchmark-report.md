# Benchmark 测试报告

> 测试日期：2026-05-08
> 测试环境：Linux (amd64), Intel(R) Xeon(R) Platinum

## 测试结果汇总

| 包 | 函数 | 耗时 | 内存 | 分配次数 |
|----|------|------|------|----------|
| **bytes** | HumanReadable | 92.89 ns | 8 B | 2 |
| **bytes** | HexEncode | 6048 ns | 4096 B | 2 |
| **bytes** | HexDecode | 2719 ns | 1024 B | 1 |
| **bytes** | Base64Encode | 4202 ns | 2816 B | 2 |
| **bytes** | Base64EncodeString | 232.8 ns | 128 B | 2 |
| **bytes** | Base64Decode | 167.1 ns | 48 B | 1 |
| **bytes** | Base64URLEncodeString | 116.2 ns | 32 B | 1 |
| **bytes** | Base64URLDecode | 111.4 ns | 24 B | 1 |
| **crypto** | MD5 | 238.7 ns | 32 B | 1 |
| **crypto** | SHA1 | 581.1 ns | 96 B | 2 |
| **crypto** | SHA256 | 564.8 ns | 128 B | 2 |
| **errors** | New | 0.47 ns | 0 B | 0 |
| **errors** | Wrap | 301.2 ns | 56 B | 2 |
| **errors** | Is | 19.22 ns | 0 B | 0 |
| **logs** | BenchmarkLog | 1476 ns | 542 B | 5 |
| **logs** | BenchmarkLogWithCtx | 1315 ns | 401 B | 5 |
| **logs** | BenchmarkDefaultLogger | 1334 ns | 426 B | 5 |
| **maps** | Keys | 27668 ns | 8192 B | 1 |
| **maps** | Values | 27488 ns | 8192 B | 1 |
| **maps** | HasKey | 12.91 ns | 0 B | 0 |
| **maps** | Merge | 198703 ns | 73888 B | 9 |
| **slices** | Dedup | 976.6 ns | 440 B | 4 |
| **slices** | Filter | 9716 ns | 8192 B | 1 |
| **slices** | Map | 8938 ns | 8192 B | 1 |
| **slices** | Contains | 481.1 ns | 0 B | 0 |
| **slices** | Split | 3293 ns | 2688 B | 1 |
| **slices** | GroupBy | 46035 ns | 9952 B | 17 |
| **str** | Upper | 269.9 ns | 48 B | 1 |
| **str** | Contains | 12.83 ns | 0 B | 0 |
| **str** | Split | 638.1 ns | 320 B | 1 |
| **str** | Ellipsis | 23.99 ns | 0 B | 0 |
| **validation** | IsEmail | 616.7 ns | 0 B | 0 |
| **validation** | IsPhone | 192.5 ns | 0 B | 0 |
| **validation** | IsURL | 941.0 ns | 0 B | 0 |
| **validation** | IsIP | 46.49 ns | 0 B | 0 |
| **validation** | IsMatch | 303.1 ns | 0 B | 0 |

---

## 性能评级

### 极快 (< 50ns)
| 函数 | 耗时 |
|------|------|
| errors.New | 0.47 ns |
| maps.HasKey | 12.91 ns |
| str.Contains | 12.83 ns |
| errors.Is | 19.22 ns |
| str.Ellipsis | 23.99 ns |
| validation.IsIP | 46.49 ns |

### 快速 (50ns - 500ns)
| 函数 | 耗时 |
|------|------|
| bytes.HumanReadable | 92.89 ns |
| bytes.Base64URLDecode | 111.4 ns |
| bytes.Base64URLEncodeString | 116.2 ns |
| bytes.Base64Decode | 167.1 ns |
| validation.IsPhone | 192.5 ns |
| bytes.Base64EncodeString | 232.8 ns |
| crypto.MD5 | 238.7 ns |
| str.Upper | 269.9 ns |
| errors.Wrap | 301.2 ns |
| validation.IsMatch | 303.1 ns |
| slices.Contains | 481.1 ns |

### 正常 (500ns - 5μs)
| 函数 | 耗时 |
|------|------|
| crypto.SHA1 | 581.1 ns |
| crypto.SHA256 | 564.8 ns |
| validation.IsEmail | 616.7 ns |
| str.Split | 638.1 ns |
| validation.IsURL | 941.0 ns |
| slices.Dedup | 976.6 ns |

### 较慢 (> 5μs)
| 函数 | 耗时 |
|------|------|
| slices.Split | 3.3 μs |
| bytes.HexDecode | 2.7 μs |
| bytes.Base64Encode | 4.2 μs |
| bytes.HexEncode | 6.0 μs |
| slices.Map | 8.9 μs |
| slices.Filter | 9.7 μs |
| maps.Keys | 27.7 μs |
| maps.Values | 27.5 μs |
| slices.GroupBy | 46.0 μs |
| maps.Merge | 199 μs (9 次分配) |

---

## 优化历史

### v1.0.3 优化
- **maps.Merge**: 29 次分配 → 9 次分配（预分配总容量）
- **slices.GroupBy**: 两遍遍历预分配每个 slice 容量

### v1.0.4 优化
- **validation.IsMatch**: 2518 ns / 10 次分配 → 303 ns / 0 次分配（引入 sync.Map 正则缓存，8x 提速）
- **slices.Split**: 预分配结果切片容量，减少扩容分配
- **slices**: 新增 Benchmark 覆盖

---

## 优化建议

### 低优先级

1. **logs** - 每次 5 次分配，可考虑使用 sync.Pool
2. **maps.Keys/Values** - 固定 8KB 分配
3. **bytes.HexEncode** - 固定 4KB 分配

---

## 结论

整体性能表现良好：
- 高频使用的工具函数（Contains、HasKey、Is 等）都在 ns 级别
- 加密函数性能符合预期
- 经过优化，maps.Merge、validation.IsMatch 内存分配显著减少
- slices 包新增完整 benchmark 覆盖