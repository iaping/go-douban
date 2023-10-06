package movie

import (
	"regexp"
	"strings"
)

type Celebrity struct {
	data []byte
}

func (sub *Celebrity) Name() string {
	reg := regexp.MustCompile(`<meta property="og:title" content="(.*?)"`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) FullName() string {
	reg := regexp.MustCompile(`<h1>(.*?)</h1>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Photo() string {
	reg := regexp.MustCompile(`<meta property="og:image" content="(.*?)"`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Sex() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?性别</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Constellation() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?星座</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Birthday() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?出生日期</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Birthplace() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?出生地</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) EnglishName() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?更多外文名</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) ChineseName() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?更多中文名</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Introduction() string {
	reg := regexp.MustCompile(`<meta property="og:description" content="([\s\S]*?)"`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Career() string {
	reg := regexp.MustCompile(`<li>[\s\S]*?职业</span>:([\s\S]*?)</li>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}

func (sub *Celebrity) Imdb() string {
	reg := regexp.MustCompile(`<a href=".*?imdb\.com.*?>(.*?)</a>`)
	result := reg.FindAllSubmatch(sub.data, 1)
	if len(result) == 0 {
		return ""
	}
	return strings.TrimSpace(string(result[0][1]))
}
