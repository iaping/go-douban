package movie

import (
	"regexp"
	"strings"
)

type Subject struct {
	data []byte
}

func (sub *Subject) Name() string {
	reg := regexp.MustCompile(`<meta property="og:title" content="(.*?)"`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Poster() string {
	reg := regexp.MustCompile(`<meta property="og:image" content="(.*?)"`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Type() string {
	reg := regexp.MustCompile(`<meta property="og:type" content="(.*?)"`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Summary() string {
	reg := regexp.MustCompile(`<span.*?property="v:summary".*?>([\s\S]*?)</span>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func (sub *Subject) Average() string {
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Country() string {
	reg := regexp.MustCompile(`<span.*?class="pl">制片国家/地区:</span>(.*?)<br/>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Language() string {
	reg := regexp.MustCompile(`<span.*?class="pl">语言:</span>(.*?)<br/>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Release() string {
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)\(.*?\)</span>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Genres() (elems []string) {
	reg := regexp.MustCompile(`<span\s*property="v:genre">(.*?)</span>`)
	result := reg.FindAllSubmatch(sub.data, -1)
	if len(result) == 0 {
		return elems
	}
	for _, v := range result {
		elems = append(elems, string(v[1]))
	}
	return elems
}

func (sub *Subject) Celebrities() (elems []string) {
	reg := regexp.MustCompile(`<a href="/celebrity/([0-9]*?)/" rel="v:starring">`)
	result := reg.FindAllSubmatch(sub.data, -1)
	if len(result) == 0 {
		return elems
	}
	for _, v := range result {
		elems = append(elems, string(v[1]))
	}
	return elems
}

func (sub *Subject) OtherNames() string {
	reg := regexp.MustCompile(`<span.*?class="pl">又名:</span>(.*?)<br/>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Runtime() string {
	reg := regexp.MustCompile(`<span property="v:runtime"[\s\S]*?>(.*?)</span>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Imdb() string {
	reg := regexp.MustCompile(`<span class="pl">IMDb:</span>(.*?)<br`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) Episode() string {
	reg := regexp.MustCompile(`<span class="pl">集数:</span>(.*?)<br`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Subject) EpisodeRuntime() string {
	reg := regexp.MustCompile(`<span class="pl">单集片长:</span>(.*?)<br`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}
