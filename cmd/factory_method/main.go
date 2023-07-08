package main

import "fmt"

type Getter interface {
	get(msgid string) string
}

type ChineseGetter struct{}
type EnglishGetter struct{}

func (p *EnglishGetter) get(msgid string) string {
	return msgid
}

func (p *ChineseGetter) get(msgid string) string {
	m := map[string]string{"dog": "gou", "cat": "mao"}
	if translated, ok := m[msgid]; ok {
		msgid = translated
	}
	return msgid
}

func get_localizer(lang_type string) Getter {
	m := map[string]Getter{"Chinese": new(ChineseGetter), "English": new(EnglishGetter)}
	return m[lang_type]
}

func main() {
	var chinese = get_localizer("Chinese")
	var english = get_localizer("English")
	animals := []string{"dog", "parrot", "cat", "bear"}
	for _, animal := range animals {
		fmt.Println(animal, "in Chinese is", chinese.get(animal))
		fmt.Println(animal, "in English is", english.get(animal))
	}
}
