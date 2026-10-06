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
	last := -1
	for _, raw := range strings.Split(text, "\n") {
		c.Scanned++
		if raw == "" {
			last = -1
			continue
		}
		if raw[0] == ' ' || raw[0] == '\t' {
			if last < 0 {
				continue
			}
			cont := strings.TrimSpace(raw)
			if cont == "" {
				continue
			}
			ix.Headers[last].Value += " " + cont
			name := ix.Headers[last].Name
			vals := ix.ByName[name]
			vals[len(vals)-1] += " " + cont
			ix.ByName[name] = vals
			continue
		}
		parts := strings.SplitN(raw, ":", 2)
		if len(parts) != 2 {
			last = -1
			continue
		}
		name := strings.ToLower(strings.TrimSpace(parts[0]))
		value := strings.TrimSpace(parts[1])
		ix.Headers = append(ix.Headers, Header{Name: name, Value: value})
		ix.ByName[name] = append(ix.ByName[name], value)
		last = len(ix.Headers) - 1
	}
	for _, d := range defaults {
		c.Scanned++
		name := strings.ToLower(strings.TrimSpace(d.Name))
		if _, ok := ix.ByName[name]; ok {
			continue
		}
		value := strings.TrimSpace(d.Value)
		ix.Headers = append(ix.Headers, Header{Name: name, Value: value})
		ix.ByName[name] = append(ix.ByName[name], value)
	}
	return ix
}

// Values 取某个头的全部值。
func (ix *Index) Values(name string, c *Counter) []string {
	c.Scanned++
	return ix.ByName[strings.ToLower(name)]
}
