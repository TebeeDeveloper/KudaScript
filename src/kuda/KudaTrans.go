package kuda

import (
	"fmt"
	"strings"
	"unicode"
)

var kdMtype map[string]string = map[string]string{
	"str":   "char*",
	"chr":   "char",
	"int":   "int",
	"i8":    "int8_t",
	"i16":   "int16_t",
	"i32":   "int32_t",
	"i64":   "int64_t",
	"uint":  "unsigned int",
	"u8":    "uint8_t",
	"u16":   "uint16_t",
	"u32":   "uint32_t",
	"u64":   "uint64_t",
	"flt":   "float",
	"dbl":   "double",
	"void":  "void",
	"bool":  "bool",

	"[]str":  "char*[]",
	"[]chr":  "char[]",
	"[]int":  "int[]",
	"[]i8":   "int8_t[]",
	"[]i16":  "int16_t[]",
	"[]i32":  "int32_t[]",
	"[]i64":  "int64_t[]",
	"[]uint": "unsigned int[]",
	"[]u8":   "uint8_t[]",
	"[]u16":  "uint16_t[]",
	"[]u32":  "uint32_t[]",
	"[]u64":  "uint64_t[]",
	"[]flt":  "float[]",
	"[]dbl":  "double[]",
	"[]void": "void[]",
	"[]bool": "bool[]",
}

var baseKeys = []string{
	"str", "chr", "int", "i8", "i16", "i32", "i64",
	"uint", "u8", "u16", "u32", "u64", "flt", "dbl", "void", "bool",
}


type KudaTranspiler struct {
	idlv int
	dwhl int
	iss  int
	isn  int
	curStruct string
	curEnum string
}

func KudaInit() KudaTranspiler {
	return KudaTranspiler{idlv: 0, dwhl: 0, iss: 0, curStruct: ""}
}

func (s *KudaTranspiler) KudaIncr()  { s.idlv++ }
func (s *KudaTranspiler) KudaDecr()  { if s.idlv > 0 { s.idlv-- } }
func (s *KudaTranspiler) KudaDwh()   { s.dwhl++ }
func (s *KudaTranspiler) KudanDwh()  { if s.dwhl > 0 { s.dwhl-- } }
func (s *KudaTranspiler) IsStruct() bool  { return s.iss > 0 }
func (s *KudaTranspiler) IsDoWhile() bool { return s.dwhl > 0 }
func (s *KudaTranspiler) KudaStrctIncr() { s.iss++ }
func (s *KudaTranspiler) KudaStrctDecr() { if s.iss > 0 { s.iss-- } }
func (s *KudaTranspiler) KudaIndent() string { return strings.Repeat("    ", s.idlv) }
func (s *KudaTranspiler) KudaChgS(strc string) {s.curStruct = strc}
func (s *KudaTranspiler) KudaGStr() string {
	var tmp string = s.curStruct
	s.curStruct = ""
	return tmp
}
func (s *KudaTranspiler) KudaEnumIncr() { s.isn++ }
func (s *KudaTranspiler) KudaEnumDecr() { if s.isn > 0 { s.isn-- }}
func (s *KudaTranspiler) KudaIsEnum() bool { return s.isn > 0 }
func (s *KudaTranspiler) KudaChgEnum(enm string) {s.curEnum = enm}
func (s *KudaTranspiler) KudaGEnum() string {
	var tmp string = s.curEnum
	s.curEnum = ""
	return  tmp
}

var allKeys = []string{
	"[]str", "[]chr", "[]int", "[]i8", "[]i16", "[]i32", "[]i64",
	"[]uint", "[]u8", "[]u16", "[]u32", "[]u64", "[]flt", "[]dbl", "[]void", "[]bool",
	"str", "chr", "int", "i8", "i16", "i32", "i64",
	"uint", "u8", "u16", "u32", "u64", "flt", "dbl", "void", "bool",
}

func isTypeStart(s string, pos int) bool {
	if pos == 0 {
		return true
	}
	c := s[pos-1]
	return !unicode.IsLetter(rune(c)) && !unicode.IsDigit(rune(c)) && c != '_'
}

func isTypeEnd(s string, pos int) bool {
	if pos >= len(s) {
		return true
	}
	c := s[pos]
	return !unicode.IsLetter(rune(c)) && !unicode.IsDigit(rune(c)) && c != '_'
}

func KudaCtype(kdtype string) string {
	if kdtype == "" {
		return kdtype
	}

	var res []byte
	pos := 0
	n := len(kdtype)

	for pos < n {
		found := false

		if isTypeStart(kdtype, pos) {
			for _, k := range allKeys {
				klen := len(k)
				if pos+klen > n {
					continue
				}
				if strings.Index(kdtype[pos:], k) == 0 && isTypeEnd(kdtype, pos+klen) {
					res = append(res, kdMtype[k]...)
					pos += klen
					found = true
					break
				}
			}
		}

		if !found {
			res = append(res, kdtype[pos])
			pos++
		}
	}

	return string(res)
}

func Kudapln(line, indent string, isenum bool) string {
	if isenum {
		return fmt.Sprintf("%s%s,\n",indent, line)
	}
	return fmt.Sprintf("%s%s;\n", indent, line)
}

func Kudapvstrcd(line, indent string) string {
	var parts []string = strings.Fields(line)
	if len(parts) < 2 {
		return "Syntax Error"
	}
	var vname string = parts[1]
	var vtype string = strings.Join(parts[2:], " ")
	return fmt.Sprintf("%s%s %s;\n", indent, vtype, vname)
}

func Kudapvd(line, indent string) string {
	var parts []string = strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return Kudapvstrcd(line, indent)
	}
	var decl []string = strings.Fields(strings.TrimSpace(parts[0]))
	var vval string = strings.TrimSpace(parts[1])

	var vname, vtype string
	switch len(decl) {
	case 3:
		vname = decl[1]
		vtype = strings.Join(decl[2:], " ")
	default:
		return "Syntax Error"
	}
	return fmt.Sprintf("%s%s %s = %s;\n", indent, vtype, vname, vval)
}

func Fnpat(args string) string {
	trimmed := strings.TrimSpace(args)
	if trimmed == "" {
		return "void"
	}

	groups := strings.Split(trimmed, ",")
	var result []string
	var pendingNames []string

	for _, g := range groups {
		group := strings.TrimSpace(g)
		if group == "" {
			continue
		}

		tokens := strings.Fields(group)
		n := len(tokens)
		if n == 0 {
			continue
		}

		// Chỉ có tên biến → gom chờ
		if n == 1 {
			pendingNames = append(pendingNames, tokens[0])
			continue
		}
		varName := tokens[0]
		typeOnly := strings.Join(tokens[1:], " ")

		var qual, typ []string
		for _, t := range strings.Fields(typeOnly) {
			if t == "const" || t == "volatile" || t == "static" {
				qual = append(qual, t)
			} else {
				typ = append(typ, t)
			}
		}
		typeOnly = strings.Join(append(qual, typ...), " ")
		for _, nm := range pendingNames {
			result = append(result, typeOnly+" "+nm)
		}
		pendingNames = nil
		result = append(result, typeOnly+" "+varName)
	}
	if len(pendingNames) > 0 && len(result) > 0 {
		parts := strings.Fields(result[len(result)-1])
		lastType := strings.Join(parts[:len(parts)-1], " ")
		for _, nm := range pendingNames {
			result = append(result, lastType+" "+nm)
		}
	}
	return strings.Join(result, ", ")
}



func Fnprt(returns string) string {
	var rtype string = strings.TrimSpace(returns)
	if rtype == "" {
		return "void"
	}
	return rtype
}

func Kudapfn(line, indent string) string {
	var body string = strings.TrimSpace(line[3:])
	var idxOpen int = strings.Index(body, "(")
	var idxClose int = strings.LastIndex(body, ")")
	if idxOpen < 0 || idxClose < 0 || idxClose <= idxOpen {
		return "Syntax Error"
	}
	funcName := strings.TrimSpace(body[:idxOpen])
	argsRaw := strings.TrimSpace(body[idxOpen+1 : idxClose])
	argsList := Fnpat(argsRaw)
	rType := Fnprt(body[idxClose+1:])
	return fmt.Sprintf("%s %s(%s) {\n", rType, funcName, argsList)
}

func Kudapif(line, indent string) string {
	var body string = strings.TrimSpace(line[3:])
	var idxThen int = strings.Index(body, " then")
	if idxThen < 0 {
		return "Syntax Error"
	}
	var condition string = strings.TrimSpace(body[:idxThen])
	return fmt.Sprintf("%sif (%s) {\n", indent, condition)
}

func Kudapelif(line, indent string) string {
	var body string = strings.TrimSpace(line[7:])
	var idxThen int = strings.Index(body, " then")
	if idxThen < 0 {
		return "Syntax Error"
	}
	var condition string = strings.TrimSpace(body[:idxThen])
	return fmt.Sprintf("%selse if (%s) {\n", indent, condition)
}

func Kudapelse(line, indent string) string {
	return fmt.Sprintf("%s} else {\n", indent)
}

func Forpl(line, indent string) string {
	var body string = strings.TrimSpace(line[4:])
	var parts []string = strings.SplitN(body, "=", 2)
	if len(parts) != 2 {
		return "Syntax Error"
	}
	var vname string = strings.TrimSpace(parts[0])
	var right []string = strings.Fields(strings.TrimSpace(parts[1]))
	if len(right) != 3 {
		return "Syntax Error"
	}
	var start string
	var end string
	var step string
	start, end, step = right[0], right[1], right[2]
	return fmt.Sprintf("%sfor(%s=%s;%s<%s;%s+=%s) {\n", indent, vname, start, vname, end, vname, step)
}

func Kudapfor(line, indent string) string {
	var body string = strings.TrimSpace(line[4:])
	var parts []string = strings.SplitN(body, "=", 2)
	if len(parts) != 2 {
		return "Syntax Error"
	}

	var left []string = strings.Fields(strings.TrimSpace(parts[0]))

	var rightRaw string = strings.ReplaceAll(parts[1], ",", " ")
	var right []string = strings.Fields(strings.TrimSpace(rightRaw))

	if len(right) != 4 {
		return "Syntax Error"
	}
	var start string
	var end string
	var step string

	start, end, step = right[0], right[1], right[2]

	var vname, vtype string
	if len(left) == 2 {
		vname = left[0]
		vtype = left[1]
	} else if len(left) == 1 {
		return Forpl(line, indent)
	} else {
		return "Syntax Error"
	}

	var cond string = "<"
	if strings.HasPrefix(step, "-") {
		cond = ">"
	}

	return fmt.Sprintf("%sfor(%s %s=%s;%s%s%s;%s+=%s) {\n", indent, vtype, vname, start, vname, cond, end, vname, step)
}


func Kudapwhl(line, indent string) string {
	var body string = strings.TrimSpace(line[6:])
	return fmt.Sprintf("%swhile (%s) {\n", indent, body)
}

func Kudapdo(line, indent string) string {
	return fmt.Sprintf("%sdo {\n", indent)
}

func Kudapdwh(line, indent string) string {
	var body string = strings.TrimSpace(line[6:])
	return fmt.Sprintf("%s} while (%s);\n", indent, body)
}

func Kudapstrc(line, indent string) (string, string) {
	var body string = strings.TrimSpace(line[7:])
	if body == "" {
		return "Syntax Error", ""
	}
	var strcName string = strings.TrimSpace(body)
	return fmt.Sprintf("%stypedef struct %s {\n", indent, strcName), strcName
}

func Kudapimp(line, indent string) string {
	var body string = strings.TrimSpace(line[7:])
	return fmt.Sprintf("#include %s\n", body)
}

func kudapenum(line, indent string) (string, string) {
	var body string = strings.TrimSpace(line[5:])
	var parts []string = strings.Fields(body)
	if len(parts) != 1 {
		return "Syntax Error", ""
	}
	var name string = parts[0]
	return fmt.Sprintf("typedef enum %s {\n", name), name
}

func kudapswitch(line, indent string) string {
	var body string = strings.TrimSpace(line[7:])
	return fmt.Sprintf("%sswitch (%s) {\n", indent, body)
}

func kudapcase(line, indent string) string {
	var body string = strings.TrimSpace(line[5:])
	return fmt.Sprintf("%scase %s:\n", indent, body)
}

func kudapdefault(line, indent string) string {
	var body string = strings.TrimSpace(line[:7])
	if body == "default " {
		return fmt.Sprintf("%sdefault:\n", indent)
	} else {
		return  fmt.Sprintf("%sdefault:\n", indent)
	}
}

func Kudapend(line, indent string, iss bool, curStructName string) string {
	if iss {
		return fmt.Sprintf("%s} %s;\n", indent, curStructName)
	}
	return fmt.Sprintf("%s}\n", indent)
}

func KudaTranslate(kudaInput, inputFile string) (string, bool) {
	var lines []string = strings.Split(kudaInput, "\n")
	var kudaCtx KudaTranspiler = KudaInit()
	var canCompile bool = true

	var kudaOutput strings.Builder
	kudaOutput.WriteString(`#include <stddef.h>
#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <string.h>
#include <ctype.h>
#include <math.h>
#include <time.h>
`)
	fmt.Fprintf(&kudaOutput, "#line 1 \"%s\"\n", inputFile)

	for lineIdx, rawLine := range lines {
		line := strings.TrimSpace(strings.TrimRight(rawLine, "\r"))
		if line == "" {
			kudaOutput.WriteString("\n")
			continue
		}

		line = KudaCtype(line)

		var indent string = kudaCtx.KudaIndent()
		var parts []string = strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		var result string
		switch parts[0] {
		case "var":
			result = Kudapvd(line, indent)
		case "fn":
			result = Kudapfn(line, indent)
			kudaCtx.KudaIncr()
		case "if":
			result = Kudapif(line, indent)
			kudaCtx.KudaIncr()
		case "elseif":
			kudaCtx.KudaDecr()
			result = Kudapelif(line, indent)
			kudaCtx.KudaIncr()
		case "else":
			kudaCtx.KudaDecr()
			result = Kudapelse(line, indent)
			kudaCtx.KudaIncr()
		case "for":
			result = Kudapfor(line, indent)
			kudaCtx.KudaIncr()
		case "do":
			result = Kudapdo(line, indent)
			kudaCtx.KudaDwh()
			kudaCtx.KudaIncr()
		case "while":
			if kudaCtx.IsDoWhile() {
				kudaCtx.KudaDecr()
				result = Kudapdwh(line, indent)
				kudaCtx.KudanDwh()
			} else {
				result = Kudapwhl(line, indent)
				kudaCtx.KudaIncr()
			}
		case "struct":
			var sn string
			result, sn = Kudapstrc(line, indent)
			kudaCtx.KudaIncr()
			kudaCtx.KudaStrctIncr()
			kudaCtx.KudaChgS(sn)
		case "enum":
			var en string
			result, en = kudapenum(line, indent)
			kudaCtx.KudaEnumIncr()
			kudaCtx.KudaChgEnum(en)
		case "import":
			result = Kudapimp(line, indent)
		case "switch":
			result = kudapswitch(line, indent)
		case "case":
			result = kudapcase(line, indent)
			kudaCtx.KudaIncr()
		case "default":
			result = kudapdefault(line, indent)
			kudaCtx.KudaIncr()
		case "break":
			result = Kudapln(line, indent, false)
			kudaCtx.KudaDecr()
		case "continue":
			result = Kudapln(line, indent, false)
			kudaCtx.KudaDecr()
		case "return":
			result = Kudapln(line, indent, false)
			kudaCtx.KudaDecr()
		case "end":
			kudaCtx.KudaDecr()
			indent = kudaCtx.KudaIndent()
			var strcnm string
			if kudaCtx.IsStruct() {
				strcnm = kudaCtx.KudaGStr()
			} else if kudaCtx.KudaIsEnum() {
				strcnm = kudaCtx.KudaGEnum()
			}
			result = Kudapend(line, indent, kudaCtx.IsStruct() || kudaCtx.KudaIsEnum(), strcnm)
			if kudaCtx.IsStruct() {
				kudaCtx.KudaStrctDecr()
			} else if kudaCtx.KudaIsEnum() {
				kudaCtx.KudaEnumDecr()
			}
		case "//":
			continue
		default:
			result = Kudapln(line, indent, kudaCtx.KudaIsEnum())
		}

		if result == "Syntax Error" {
			canCompile = false
			fmt.Printf("[Kuda-Transpiler] Syntax Error Line %d: %s\n", lineIdx+1, rawLine)
		}

		if canCompile {
			kudaOutput.WriteString(result)
		}
	}

	return kudaOutput.String(), canCompile
}