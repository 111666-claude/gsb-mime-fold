// Package mimefold 解析邮件头块。
package mimefold

import "strings"

// Counter 记录比较次数。
type Counter struct {
	Scanned int
}

// Header 是一个头部。
type Header struct {
	Name  string
	Value string
}

// Index 是解析后的头块。
type Index struct {
	Headers []Header
	ByName  map[string][]string
}

// Build 解析头块并按缺省层补齐。
func Build(text string, defaults []Header, c *Counter) *Index {
	ix := &Index{}
	for _, raw := range strings.Split(text, "\n") {
		c.Scanned++
		if raw == "" {
			continue
		}
		if raw[0] == ' ' || raw[0] == '\t' {
			continue
		}
		parts := strings.SplitN(raw, ":", 2)
		if len(parts) != 2 {
			continue
		}
		ix.Headers = append(ix.Headers, Header{
			Name:  strings.TrimSpace(parts[0]),
			Value: strings.TrimSpace(parts[1]),
		})
	}
	for _, d := range defaults {
		c.Scanned++
		ix.Headers = append(ix.Headers, Header{Name: d.Name, Value: d.Value})
	}
	return ix
}

// Values 取某个头的全部值。
func (ix *Index) Values(name string, c *Counter) []string {
	var out []string
	for _, h := range ix.Headers {
		c.Scanned++
		if h.Name == name {
			out = append(out, h.Value)
		}
	}
	if len(out) > 1 {
		out = out[len(out)-1:]
	}
	return out
}
