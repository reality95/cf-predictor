package lib

import (
	"log"
	"regexp"
	"strings"
)

const prefixDefine string = `#define +`
const varSyntax string = `[A-Za-z_]+[0-9]*`
const argSyntax string = `( *\( *` + varSyntax + ` *( *, *` + varSyntax + ` *)* *\) *)?`
const middleDefine string = varSyntax + argSyntax
const suffixDefine string = `[^\n]*\n`
const notVarEnding string = `[ ,{}\n\t+-/*|&\[\];=\(\)%]`

func ReplaceDefine(code string) (ans string) {
	re := regexp.MustCompile(prefixDefine + middleDefine + suffixDefine)
	reVar := regexp.MustCompile(varSyntax)
	rePrefix := regexp.MustCompile(prefixDefine)
	reMiddle := regexp.MustCompile(middleDefine)

	definesIndex := re.FindAllIndex([]byte(code), -1)
	defines := make([]define, 0)

	for _, idxs := range definesIndex {
		expr := rePrefix.ReplaceAll([]byte(code[idxs[0]:idxs[1]]), nil)
		middle := reMiddle.Find(expr)
		args := reVar.FindAll(middle, -1)
		expr = []byte(strings.TrimPrefix(string(expr), string(middle)))
		defines = append(defines, define{
			Name:      string(args[0]),
			Args:      toStringArray(args[1:]),
			Expresion: string(expr[:len(expr)-1]),
		})
	}

	return parseDefines(re.ReplaceAll([]byte(code), nil), defines)
}

func parseDefines(code []byte, defines []define) string {
	Res := make([]*regexp.Regexp, 0)
	for _, define := range defines {
		Res = append(Res, regexp.MustCompile(notVarEnding+define.Name+argSyntax+notVarEnding))
	}
	reVar := regexp.MustCompile(varSyntax)
	found := true
	for found {
		found = false
		for i := len(defines) - 1; i >= 0; i-- {
			found = found || (Res[i].Find(code) != nil)
			code = Res[i].ReplaceAllFunc(code, func(s []byte) (ans []byte) {
				startingChar := make([]byte, 0)
				startingChar = append(startingChar, s[0])
				endingChar := s[len(s)-1]
				args := reVar.FindAll(s, -1)[1:]
				if len(args) != len(defines[i].Args) {
					log.Println(len(args), len(defines[i].Args))
					panic("Expected a different amount of arguments in define")
				}
				ans = make([]byte, len(defines[i].Expresion))
				copy(ans, defines[i].Expresion)
				for j, arg := range defines[i].Args {
					reArg := regexp.MustCompile(notVarEnding + arg + notVarEnding)

					ans = reArg.ReplaceAll(ans, args[j])
				}
				return append(append(startingChar, ans...), endingChar)
			})
		}
	}
	log.Println(string(code))
	return string(code)
}

func toStringArray(s [][]byte) (ans []string) {
	for _, v := range s {
		ans = append(ans, string(v))
	}
	return ans
}

type define struct {
	Name      string
	Args      []string
	Expresion string
}
