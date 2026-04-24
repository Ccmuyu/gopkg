# Benchmark 测试报告

> 测试日期：2026-04-24
> 测试环境：Linux (amd64), Intel(R) Xeon(R) Platinum

## 测试结果汇总

| 包 | 函数 | 耗时 | 内存 | 分配次数 |
|----|------|------|------|----------|
| **bytes** | HumanReadable | 106.7 ns | 8 B | 2 |
| **bytes** | HexEncode | 8905 ns | 4096 B | 2 |
| **bytes** | HexDecode | 3132 ns | 1024 B | 1 |
| **bytes** | Base64Encode | 4826 ns | 2816 B | 2 |
| **crypto** | MD5 | 266.8 ns | 32 B | 1 |
| **crypto** | SHA1 | 646.9 ns | 96 B | 2 |
| **crypto** | SHA256 | 646.1 ns | 128 B | 2 |
| **errors** | New | 0.53 ns | 0 B | 0 |
| **errors** | Wrap | 352.0 ns | 56 B | 2 |
| **errors** | Is | 21.41 ns | 0 B | 0 |
| **log** | BenchmarkLog | 1717 ns | 589 B | 5 |
| **log** | BenchmarkLogWithCtx | 1642 ns | 456 B | 5 |
| **log** | BenchmarkDefaultLogger | 1552 ns | 433 B | 5 |
| **maps** | Keys | 28771 ns | 8192 B | 1 |
| **maps** | Values | 32322 ns | 8192 B | 1 |
| **maps** | HasKey | 15.79 ns | 0 B | 0 |
| **maps** | Merge | 208935 ns | 73888 B | 9 |
| **slices** | Dedup | - | - | - |
| **slices** | Filter | - | - | - |
| **slices** | Map | - | - | - |
| **slices** | Contains | - | - | - |
| **str** | Upper | 300.2 ns | 48 B | 1 |
| **str** | Contains | 14.34 ns | 0 B | 0 |
| **str** | Split | 728.7 ns | 320 B | 1 |
| **str** | Ellipsis | 26.76 ns | 0 B | 0 |
| **validation** | IsEmail | 692.2 ns | 0 B | 0 |
| **validation** | IsPhone | 212.9 ns | 0 B | 0 |
| **validation** | IsURL | 1063 ns | 0 B | 0 |
| **validation** | IsIP | 51.22 ns | 0 B | 0 |
| **validation** | IsMatch | 2518 ns | 874 B | 10 |

---

## 性能评级

### 极快 (< 50ns)
| 函数 | 耗时 |
|------|------|
| errors.New | 0.53 ns |
| str.Contains | 14.34 ns |
| maps.HasKey | 15.79 ns |

### 快速 (50ns - 500ns)
| 函数 | 耗时 |
|------|------|
| validation.IsPhone | 212.9 ns |
| validation.IsIP | 51.22 ns |
| crypto.MD5 | 266.8 ns |
| str.Upper | 300.2 ns |

### 正常 (500ns - 5μs)
| 函数 | 耗时 |
|------|------|
| bytes.HumanReadable | 106.7 ns |
| str.Ellipsis | 26.76 ns |
| validation.IsEmail | 692.2 ns |
| crypto.SHA1 | 646.9 ns |

### 较慢 (> 5μs)
| 函数 | 耗时 |
|------|------|
| maps.HasKey | 15.79 ns |
| maps.Keys | 28.8 μs |
| maps.Values | 32.3 μs |
| maps.Merge | 209 μs (已优化，9次分配) |
| bytes.HexEncode | 8.9 μs |
| bytes.Base64Encode | 4.8 μs |

---

## 优化历史

### v1.0.3 优化
- **maps.Merge**: 29 次分配 → 9 次分配（预分配总容量）
- **slices.GroupBy**: 两遍遍历预分配每个 slice 容量

---

## 优化建议

### 低优先级

1. **log** - 每次 5 次分配，可考虑使用 sync.Pool
2. **maps.Keys/Values** - 固定 8KB 分配
3. **bytes.HexEncode** - 固定 4KB 分配
4. **validation.IsMatch** - 10 次分配，正则表达式复用可优化

---

## 结论

整体性能表现良好：
- 高频使用的工具函数（Contains、HasKey、Is 等）都在 ns 级别
- 加密函数性能符合预期
- 经过优化，maps.Merge 内存分配显著减少