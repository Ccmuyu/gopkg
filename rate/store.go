package rate

import (
	"sync"
	"time"
)

// CountResult 单次计数/查询的结果
type CountResult struct {
	Current    int64 // 当前窗口内累计请求数
	Limited    bool  // 是否触发限流（Peek 时表示「下一次 Incr 会被限」）
	ResetAfter int64 // 建议重试等待秒数
	Remaining  int64 // 窗口内剩余可请求次数
}

// Store 限流计数存储接口（预留 Redis 扩展点）
type Store interface {
	// Incr 在未达到 limit 时对 key 在当前窗口内计数 +1。
	// 已达到 limit 时拒绝且不增加计数，避免超限流量延长封禁时间。
	Incr(key string, limit, window int64) (CountResult, error)
	// Peek 只读查询当前计数，不增加计数
	Peek(key string, limit, window int64) (CountResult, error)
	// Cleanup 清理过期数据
	Cleanup() error
}

// slot 单个时间槽
type slot struct {
	timestamp int64 // 槽起始时间戳（Unix 秒，按 slotWidth 对齐）
	count     int64 // 槽内累计请求数
}

// keyState 单个 key 的计数状态
type keyState struct {
	slots      []slot
	lastAccess int64 // 最近一次 Incr 的 Unix 秒
	lastWindow int64 // 最近一次 Incr 使用的 window
}

// storeKey 将 window 作为规则身份的一部分。同一个业务 key 使用不同窗口时，
// 必须维护独立状态，否则短窗口检查会不可逆地删除长窗口仍需使用的历史。
type storeKey struct {
	key    string
	window int64
}

// slidingWindowStore 基于滑动窗口计数器的内存实现。
// 每次 Incr/Peek 按入参 window 计算槽宽，避免构造时 window 与调用 window 不一致。
type slidingWindowStore struct {
	mu      sync.Mutex
	data    map[storeKey]*keyState
	slots   int   // 子槽数量（用于粒度；不因槽满丢弃未过期计数）
	idleTTL int64 // Cleanup 空闲阈值下限（秒）
	nowFunc func() time.Time
}

// NewSlidingWindowStore 创建内存滑动窗口存储。
// window: 默认空闲清理参考窗口（秒）；实际计数窗口以 Incr/Peek 的 window 为准。
// slots:  子槽数量，影响时间粒度。
func NewSlidingWindowStore(window int64, slots int) Store {
	if window <= 0 {
		window = 60
	}
	if slots <= 0 {
		slots = 10
	}
	idleTTL := window * 2
	if idleTTL < 120 {
		idleTTL = 120
	}
	return &slidingWindowStore{
		data:    make(map[storeKey]*keyState),
		slots:   slots,
		idleTTL: idleTTL,
		nowFunc: time.Now,
	}
}

func (s *slidingWindowStore) slotWidth(window int64) int64 {
	w := window / int64(s.slots)
	if w <= 0 {
		return 1
	}
	return w
}

func remaining(limit, current int64) int64 {
	r := limit - current
	if r < 0 {
		return 0
	}
	return r
}

func resetAfter(now, window, slotWidth int64, slots []slot, limited bool) int64 {
	if !limited || len(slots) == 0 {
		return 0
	}
	// 槽内没有保存单次请求时间，因此按整个最老槽过期计算。
	// 这会略偏保守，但不会返回一个过早、等待后仍被限流的时间。
	ra := slots[0].timestamp + slotWidth + window - now
	if ra < 1 {
		return 1
	}
	return ra
}

// collectValid 收集仍在窗口内的槽并求和（不修改原数据）
func collectValid(slots []slot, windowStart, slotWidth int64) ([]slot, int64) {
	valid := make([]slot, 0, len(slots)+1)
	var current int64
	for _, sl := range slots {
		// 只要槽尾仍与窗口相交，就保守地保留整个槽。按槽起点判断会
		// 提前丢弃槽后半段仍在窗口内的请求，造成限流漏放。
		if sl.timestamp+slotWidth > windowStart {
			valid = append(valid, sl)
			current += sl.count
		}
	}
	return valid, current
}

// Incr 对 key 计数 +1
func (s *slidingWindowStore) Incr(key string, limit, window int64) (CountResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limit <= 0 || window <= 0 {
		return CountResult{Limited: true}, nil
	}

	now := s.nowFunc().Unix()
	sw := s.slotWidth(window)
	currentSlotTs := (now / sw) * sw
	windowStart := now - window

	sk := storeKey{key: key, window: window}
	ks := s.data[sk]
	if ks == nil {
		ks = &keyState{}
		s.data[sk] = ks
	}

	validSlots, current := collectValid(ks.slots, windowStart, sw)
	ks.slots = validSlots
	ks.lastAccess = now
	ks.lastWindow = window

	// 拒绝的请求不进入窗口。否则持续的超限流量会不断抬高 current，
	// 即使按 ResetAfter 等待后也仍然无法恢复。
	if current >= limit {
		return CountResult{
			Current:    current,
			Limited:    true,
			ResetAfter: resetAfter(now, window, sw, validSlots, true),
			Remaining:  0,
		}, nil
	}

	found := false
	for i, sl := range validSlots {
		if sl.timestamp == currentSlotTs {
			validSlots[i].count++
			current++
			found = true
			break
		}
	}
	if !found {
		validSlots = append(validSlots, slot{timestamp: currentSlotTs, count: 1})
		current++
	}

	ks.slots = validSlots

	return CountResult{
		Current:    current,
		Limited:    false,
		ResetAfter: 0,
		Remaining:  remaining(limit, current),
	}, nil
}

// Peek 只读查询，不增加计数。Limited 表示下一次 Incr 会被限流（current >= limit）。
func (s *slidingWindowStore) Peek(key string, limit, window int64) (CountResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limit <= 0 || window <= 0 {
		return CountResult{Limited: true}, nil
	}

	now := s.nowFunc().Unix()
	windowStart := now - window
	sw := s.slotWidth(window)

	var slots []slot
	if ks := s.data[storeKey{key: key, window: window}]; ks != nil {
		slots = ks.slots
	}
	validSlots, current := collectValid(slots, windowStart, sw)
	limited := current >= limit

	return CountResult{
		Current:    current,
		Limited:    limited,
		ResetAfter: resetAfter(now, window, sw, validSlots, limited),
		Remaining:  remaining(limit, current),
	}, nil
}

// Cleanup 清理空闲 key：空闲超过 max(idleTTL, 2*lastWindow) 则删除
func (s *slidingWindowStore) Cleanup() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := s.nowFunc().Unix()
	for key, ks := range s.data {
		if ks == nil || len(ks.slots) == 0 {
			delete(s.data, key)
			continue
		}
		ttl := ks.lastWindow * 2
		if ttl < s.idleTTL {
			ttl = s.idleTTL
		}
		if ks.lastAccess < now-ttl {
			delete(s.data, key)
		}
	}
	return nil
}

// lenKeys 测试用：当前 key 数量
func (s *slidingWindowStore) lenKeys() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}
