package mimefold

import "fmt"

// Sample 是一个固定场景，用来复现题面里的锚点。
type Sample struct {
	Name string
	Run  func() map[string]any
}

func values(ix *Index, name string) []string {
	v := ix.Values(name, &Counter{})
	if v == nil {
		return []string{}
	}
	return v
}

// Samples 返回全部场景。
func Samples() []Sample {
	return []Sample{
		{Name: "fold", Run: func() map[string]any {
			ix := Build("subject: hello\n world\n", nil, &Counter{})
			return map[string]any{"values": values(ix, "subject")}
		}},
		{Name: "case", Run: func() map[string]any {
			ix := Build("Subject: a\n", nil, &Counter{})
			return map[string]any{"values": values(ix, "subject")}
		}},
		{Name: "multi", Run: func() map[string]any {
			ix := Build("received: a\nreceived: b\n", nil, &Counter{})
			return map[string]any{"values": values(ix, "received")}
		}},
		{Name: "default", Run: func() map[string]any {
			ix := Build("x-mode: live\n", []Header{{Name: "x-mode", Value: "off"}}, &Counter{})
			return map[string]any{"values": values(ix, "x-mode")}
		}},
	}
}

// Work 跑规模线场景。
func Work(n int) map[string]any {
	var b []string
	for i := 0; i < n; i++ {
		b = append(b, fmt.Sprintf("x-%04d: v%d", i, i))
	}
	c := &Counter{}
	ix := Build(join(b), nil, c)
	for i := 0; i < n; i++ {
		ix.Values(fmt.Sprintf("x-%04d", i), c)
	}
	return map[string]any{"headers": n, "scanned": c.Scanned}
}

func join(lines []string) string {
	out := ""
	for i, line := range lines {
		if i > 0 {
			out += "\n"
		}
		out += line
	}
	return out
}
