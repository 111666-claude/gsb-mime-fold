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
	ix := &Index{ByName: map[string][]string{}}
	add := func(name, value string) {
		ix.Headers = append(ix.Headers, Header{Name: name, Value: value})
		ix.ByName[name] = append(ix.ByName[name], value)
	}
	for _, raw := range strings.Split(text, "\n") {
		c.Scanned++
		if raw == "" {
			continue
		}
		if raw[0] == ' ' || raw[0] == '\t' {
			if len(ix.Headers) > 0 {
				cont := strings.TrimSpace(raw)
				if cont != "" {
					last := &ix.Headers[len(ix.Headers)-1]
					if last.Value != "" {
						last.Value += " " + cont
					} else {
						last.Value = cont
					}
					vals := ix.ByName[last.Name]
					vals[len(vals)-1] = last.Value
				}
			}
			continue
		}
		parts := strings.SplitN(raw, ":", 2)
		if len(parts) != 2 {
			continue
		}
		add(strings.ToLower(strings.TrimSpace(parts[0])), strings.TrimSpace(parts[1]))
	}
	for _, d := range defaults {
		c.Scanned++
		name := strings.ToLower(d.Name)
		if _, ok := ix.ByName[name]; ok {
			continue
		}
		add(name, d.Value)
	}
	return ix
}

// Values 取某个头的全部值。
func (ix *Index) Values(name string, c *Counter) []string {
	c.Scanned++
	return ix.ByName[strings.ToLower(name)]
}
