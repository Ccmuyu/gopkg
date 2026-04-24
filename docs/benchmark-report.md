# Benchmark 测试报告

> 测试日期：2026-04-24  
> 测试环境：Linux (amd64), Intel(R) Xeon(R) Platinum

## 测试结果汇总

| 包 | 函数 | 耗时 | 内存 | 分配次数 |
|----|------|------|------|----------|
| **bytes** | HumanReadable | 77.94 ns | 8 B | 2 |
| **bytes** | HexEncode | 4302 ns | 4096 B | 2 |
| **bytes** | HexDecode | 2078 ns | 1024 B | 1 |
| **bytes** | Base64Encode | 2931 ns | 2816 B | 2 |
| **crypto** | MD5 | 208.9 ns | 32 B | 1 |
| **crypto** | SHA1 | 450.1 ns | 96 B | 2 |
| **crypto** | SHA256 | 463.4 ns | 128 B | 2 |
| **errors** | New | 0.45 ns | 0 B | 0 |
| **errors** | Wrap | 247.9 ns | 56 B | 2 |
| **errors** | Is | 16.86 ns | 0 B | 0 |
| **log** | BenchmarkLog | 1313 ns | 531 B | 5 |
| **log** | BenchmarkLogWithCtx | 1051 ns | 382 B | 5 |
| **log** | BenchmarkDefaultLogger | 1048 ns | 382 B | 5 |
| **maps** | Keys | 20566 ns | 8192 B | 1 |
| **maps** | Values | 23988 ns | 8192 B | 1 |
| **maps** | HasKey | 11.66 ns | 0 B | 0 |
| **maps** | Merge | 277258 ns | 148152 B | 29 |
| **slices** | Dedup | 688357 ns | 377474 B | 34 |
| **slices** | Filter | 123309 ns | 81920 B | 1 |
| **slices** | Map | 155796 ns | 81920 B | 1 |
| **slices** | Contains | 5461 ns | 0 B | 0 |
| **str** | Upper | 217.3 ns | 48 B | 1 |
| **str** | Contains | 10.76 ns | 0 B | 0 |
| **str** | Split | 432.9 ns | 320 B | 1 |
| **str** | Ellipsis | 21.18 ns | 0 B | 0 |
| **validation** | IsEmail | 546.3 ns | 0 B | 0 |
| **validation** | IsPhone | 159.7 ns | 0 B | 0 |
| **validation** | IsURL | 811.6 ns | 0 B | 0 |
| **validation** | IsIP | 45.52 ns | 0 B | 0 |
| **validation** | IsMatch | 2357 ns | 874 B | 10 |

---

## 性能评级

### 极快 (< 50ns)
| 函数 | 耗时 |
|------|------|
| errors.New | 0.45 ns |
| str.Contains | 10.76 ns |
| maps.HasKey | 11.66 ns |

### 快速 (50ns - 500ns)
| 函数 | 耗时 |
|------|------|
| validation.IsPhone | 159.7 ns |
| crypto.MD5 | 208.9 ns |
| str.Upper | 217.3 ns |

### 正常 (500ns - 5μs)
| 函数 | 耗时 |
|------|------|
| bytes.HumanReadable | 77.94 ns |
| validation.IsIP | 45.52 ns |
| str.Ellipsis | 21.18 ns |
| validation.IsEmail | 546.3 ns |

### 较慢 (> 5μs)
| 函数 | 耗时 |
|------|------|
| slices.Contains | 5.46 μs |
| slices.Filter | 123.3 μs |
| slices.Map | 155.8 μs |
| maps.Keys | 20.6 μs |
| maps.Values | 24.0 μs |
| slices.Dedup | 688 μs |
| maps.Merge | 277 μs |

---

## 优化建议

### 高优先级

1. **maps.Merge** - 29次内存分配，需要优化合并逻辑
2. **slices.Dedup** - 34次分配，可考虑预分配 map 容量

### 中优先级

3. **log** - 每次5次分配，可考虑使用 sync.Pool
4. **maps.Keys/Values** - 固定8KB分配

### 低优先级

5. **bytes.HexEncode** - 固定4KB分配，对于大数据块影响不大

---

## 结论

整体性能表现良好：
- 高频使用的工具函数（Contains、HasKey、Is 等）都在 ns 级别
- 加密函数性能符合预期
- 切片/map 操作性能与算法复杂度匹配