package rate

// RateLimitItem 单一限流规则（limit + window）
type RateLimitItem struct {
	Limit  int64 // 窗口内最大请求数
	Window int64 // 窗口大小（秒）
}

// RateLimitRouteItem 路由级限流规则
type RateLimitRouteItem struct {
	Route  string // 路由路径（如 "/api/creator/sync"）
	Limit  int64  // 窗口内最大请求数
	Window int64  // 窗口大小（秒）
}

// RateLimitConfig 完整的限流配置
type RateLimitConfig struct {
	Enabled  bool                 // 是否启用
	Global   *RateLimitItem       // 全局规则；nil 表示不启用该级
	PerIP    *RateLimitItem       // IP 级规则；nil 表示不启用该级
	PerRoute []RateLimitRouteItem // 路由级规则列表
}

// DefaultConfig 返回默认配置（默认禁用，避免误伤）
func DefaultConfig() *RateLimitConfig {
	return &RateLimitConfig{
		Enabled:  false,
		Global:   &RateLimitItem{Limit: 1000, Window: 60},
		PerIP:    &RateLimitItem{Limit: 100, Window: 60},
		PerRoute: []RateLimitRouteItem{},
	}
}

// NewConfig 便捷构造：返回已启用的配置。
// global / perIP 传 nil 表示不启用对应级别（不会继承 DefaultConfig 的默认值）。
func NewConfig(global, perIP *RateLimitItem, routes ...RateLimitRouteItem) *RateLimitConfig {
	cfg := &RateLimitConfig{
		Enabled: true,
		Global:  global,
		PerIP:   perIP,
	}
	if len(routes) > 0 {
		cfg.PerRoute = append([]RateLimitRouteItem(nil), routes...)
	} else {
		cfg.PerRoute = []RateLimitRouteItem{}
	}
	return cfg
}

// Clone 返回配置深拷贝，修改副本不影响原配置
func (c *RateLimitConfig) Clone() *RateLimitConfig {
	if c == nil {
		return nil
	}
	out := &RateLimitConfig{
		Enabled: c.Enabled,
	}
	if c.Global != nil {
		g := *c.Global
		out.Global = &g
	}
	if c.PerIP != nil {
		p := *c.PerIP
		out.PerIP = &p
	}
	if c.PerRoute != nil {
		out.PerRoute = append([]RateLimitRouteItem(nil), c.PerRoute...)
	} else {
		out.PerRoute = []RateLimitRouteItem{}
	}
	return out
}

// GetRouteRule 按路由查找规则；返回副本，不存在时返回 nil
func (c *RateLimitConfig) GetRouteRule(route string) *RateLimitRouteItem {
	if c == nil {
		return nil
	}
	for i := range c.PerRoute {
		if c.PerRoute[i].Route == route {
			item := c.PerRoute[i]
			return &item
		}
	}
	return nil
}

// maxWindow 返回配置中各规则 Window 的最大值，用于构造默认 Store
func (c *RateLimitConfig) maxWindow() int64 {
	var w int64
	if c.Global != nil && c.Global.Window > w {
		w = c.Global.Window
	}
	if c.PerIP != nil && c.PerIP.Window > w {
		w = c.PerIP.Window
	}
	for i := range c.PerRoute {
		if c.PerRoute[i].Window > w {
			w = c.PerRoute[i].Window
		}
	}
	if w <= 0 {
		return 60
	}
	return w
}
