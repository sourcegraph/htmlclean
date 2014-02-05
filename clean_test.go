package htmlclean

import (
	"fmt"
	"net/url"
	"testing"
)

func TestClean(t *testing.T) {
	tests := map[string]struct {
		html string
		url  *url.URL
		want string
	}{
		"text only":           {"a", nil, "a"},
		"allowed tag":         {"<a>foo</a>", nil, "<a>foo</a>"},
		"forbidden tag":       {"<script>foo</script>", nil, ""},
		"forbidden attribute": {"<a onclick='foo'>bar</a>", nil, "<a>bar</a>"},
	}
	for label, test := range tests {
		got, _ := Clean(test.html, test.url)
		if test.want != got {
			t.Errorf("%s: want %q, got %q", label, test.want, got)
			continue
		}
	}
}

func mustParseURL(urlStr string) *url.URL {
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(fmt.Sprintf("url.Parse(%q): %s", urlStr, err))
	}
	return u
}
