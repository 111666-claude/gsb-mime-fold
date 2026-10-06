package mimefold

import "testing"

func TestSingleHeader(t *testing.T) {
	ix := Build("subject: a\n", nil, &Counter{})
	got := ix.Values("subject", &Counter{})
	if len(got) != 1 || got[0] != "a" {
		t.Fatalf("应取到 a：%+v", got)
	}
}

func TestMissingIsEmpty(t *testing.T) {
	ix := Build("subject: a\n", nil, &Counter{})
	if got := ix.Values("nope", &Counter{}); len(got) != 0 {
		t.Fatalf("缺失的头应返回空：%+v", got)
	}
}

func TestDefaultsFillGap(t *testing.T) {
	ix := Build("a: 1\n", []Header{{Name: "b", Value: "2"}}, &Counter{})
	got := ix.Values("b", &Counter{})
	if len(got) != 1 || got[0] != "2" {
		t.Fatalf("缺省层应补上 b=2：%+v", got)
	}
}

func TestHeaderCount(t *testing.T) {
	ix := Build("a: 1\nb: 2\n", nil, &Counter{})
	if len(ix.Headers) != 2 {
		t.Fatalf("应有两行头：%+v", ix.Headers)
	}
}

func TestCounterCounts(t *testing.T) {
	c := &Counter{}
	Build("a: 1\n", nil, c)
	ix := Build("a: 1\n", nil, &Counter{})
	ix.Values("a", c)
	if c.Scanned == 0 {
		t.Fatal("应统计比较次数")
	}
}

func TestBlankLinesSkipped(t *testing.T) {
	ix := Build("a: 1\n\nb: 2\n", nil, &Counter{})
	if len(ix.Headers) != 2 {
		t.Fatalf("空行不该产生头：%+v", ix.Headers)
	}
}
