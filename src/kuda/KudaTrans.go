package kuda

import (
	"fmt"
	"strings"
)

var kdMtype = map[string]string{
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

	"[]str":  "char**",
	"[]chr":  "char*",
	"[]int":  "int*",
	"[]i8":   "int8_t*",
	"[]i16":  "int16_t*",
	"[]i32":  "int32_t*",
	"[]i64":  "int64_t*",
	"[]uint": "unsigned int*",
	"[]u8":   "uint8_t*",
	"[]u16":  "uint16_t*",
	"[]u32":  "uint32_t*",
	"[]u64":  "uint64_t*",
	"[]flt":  "float*",
	"[]dbl":  "double*",
	"[]void": "void*",
	"[]bool": "bool*",
}

type KudaTranspiler struct {
	idlv int
	dwhl int
	iss  int
	curStruct string
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

func KudaCtype(kdtype string) string {
	if c, ok := kdMtype[kdtype]; ok {
		return c
	}
	return kdtype
}

func Kudapln(line, indent string) string {
	return fmt.Sprintf("%s%s;\n", indent, line)
}

func Kudapvstrcd(line, indent string) string {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return "Syntax Error"
	}
	vname := parts[1]
	vtype := KudaCtype(parts[2])
	return fmt.Sprintf("%s%s %s;\n", indent, vtype, vname)
}

func Kudapvd(line, indent string) string {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return Kudapvstrcd(line, indent)
	}
	decl := strings.Fields(strings.TrimSpace(parts[0]))
	vval := strings.TrimSpace(parts[1])

	var vname, vtype string
	switch len(decl) {
	case 2: // var name = val → mặc định kiểu int
		vname = decl[1]
		vtype = "int"
	case 3: // var name type = val
		vname = decl[1]
		vtype = KudaCtype(decl[2])
	default:
		return "Syntax Error"
	}
	return fmt.Sprintf("%s%s %s = %s;\n", indent, vtype, vname, vval)
}

func Fnpat(args string) string {
	if args == "" {
		return "void"
	}
	parts := strings.Split(args, ",")
	var result []string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		tokens := strings.Fields(part)
		if len(tokens) >= 2 {
			lastIdx := len(tokens) - 1
			paramType := KudaCtype(tokens[lastIdx])
			for _, name := range tokens[:lastIdx] {
				result = append(result, fmt.Sprintf("%s %s", paramType, name))
			}
		} else if len(tokens) == 1 {
			result = append(result, tokens[0])
		}
	}
	return strings.Join(result, ", ")
}

func Fnprt(returns string) string {
	rtype := strings.TrimSpace(returns)
	if rtype == "" {
		return "void"
	}
	return KudaCtype(rtype)
}

func Kudapfn(line, indent string) string {
	body := strings.TrimSpace(line[3:])
	idxOpen := strings.Index(body, "(")
	idxClose := strings.LastIndex(body, ")")
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
	body := strings.TrimSpace(line[3:])
	idxThen := strings.Index(body, " then")
	if idxThen < 0 {
		return "Syntax Error"
	}
	condition := strings.TrimSpace(body[:idxThen])
	return fmt.Sprintf("%sif (%s) {\n", indent, condition)
}

func Kudapelif(line, indent string) string {
	body := strings.TrimSpace(line[7:])
	idxThen := strings.Index(body, " then")
	if idxThen < 0 {
		return "Syntax Error"
	}
	condition := strings.TrimSpace(body[:idxThen])
	return fmt.Sprintf("%selse if (%s) {\n", indent, condition)
}

func Kudapelse(line, indent string) string {
	return fmt.Sprintf("%s} else {\n", indent)
}

func Forpl(line, indent string) string {
	body := strings.TrimSpace(line[4:])
	parts := strings.SplitN(body, "=", 2)
	if len(parts) != 2 {
		return "Syntax Error"
	}
	vname := strings.TrimSpace(parts[0])
	right := strings.Fields(strings.TrimSpace(parts[1]))
	if len(right) != 3 {
		return "Syntax Error"
	}
	start, end, step := right[0], right[1], right[2]
	return fmt.Sprintf("%sfor(%s=%s;%s<%s;%s+=%s) {\n", indent, vname, start, vname, end, vname, step)
}

func Kudapfor(line, indent string) string {
	body := strings.TrimSpace(line[4:])
	parts := strings.SplitN(body, "=", 2)
	if len(parts) != 2 {
		return "Syntax Error"
	}

	left := strings.Fields(strings.TrimSpace(parts[0]))
	// Hỗ trợ cả dấu phẩy và dấu cách cho phần giá trị
	rightRaw := strings.ReplaceAll(parts[1], ",", " ")
	right := strings.Fields(strings.TrimSpace(rightRaw))

	if len(right) != 4 {
		return "Syntax Error"
	}
	start, end, step := right[0], right[1], right[2]

	var vname, vtype string
	if len(left) == 2 {
		vname = left[0]
		vtype = KudaCtype(left[1])
	} else if len(left) == 1 {
		return Forpl(line, indent)
	} else {
		return "Syntax Error"
	}

	return fmt.Sprintf("%sfor(%s %s=%s;%s<%s;%s+=%s) {\n", indent, vtype, vname, start, vname, end, vname, step)
}


func Kudapwhl(line, indent string) string {
	body := strings.TrimSpace(line[6:])
	return fmt.Sprintf("%swhile (%s) {\n", indent, body)
}

func Kudapdo(line, indent string) string {
	return fmt.Sprintf("%sdo {\n", indent)
}

func Kudapdwh(line, indent string) string {
	body := strings.TrimSpace(line[6:])
	return fmt.Sprintf("%s} while (%s);\n", indent, body)
}

func Kudapstrc(line, indent string) (string, string) {
	body := strings.TrimSpace(line[7:])
	if body == "" {
		return "Syntax Error", ""
	}
	strcName := strings.TrimSpace(body)
	return fmt.Sprintf("%stypedef struct %s {\n", indent, strcName), strcName
}

func Kudapimp(line, indent string) string {
	body := strings.TrimSpace(line[7:])
	return fmt.Sprintf("#include %s\n", body)
}

func Kudapend(line, indent string, iss bool, curStructName string) string {
	if iss {
		// Đóng struct ĐÚNG chuẩn: } TênStruct;
		return fmt.Sprintf("%s} %s;\n", indent, curStructName)
	}
	return fmt.Sprintf("%s}\n", indent)
}

func KudaTranslate(kudaInput, inputFile string) (string, bool) {
	lines := strings.Split(kudaInput, "\n")
	kudaCtx := KudaInit()
	canCompile := true

	var kudaOutput strings.Builder
	kudaOutput.WriteString(`#include <stddef.h>
#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include <stdbool.h>
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

		indent := kudaCtx.KudaIndent()
		parts := strings.Fields(line)
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
		case "import":
			result = Kudapimp(line, indent)
		case "end":
			kudaCtx.KudaDecr()
			indent = kudaCtx.KudaIndent()
			var strcnm string = kudaCtx.KudaGStr()
			result = Kudapend(line, indent, kudaCtx.IsStruct(), strcnm)
			if kudaCtx.IsStruct() {
				kudaCtx.KudaStrctDecr()
			}
		default:
			result = Kudapln(line, indent)
		}

		if result == "Syntax Error" {
			canCompile = false
			fmt.Printf("[Lỗi cú pháp] Dòng %d: %s\n", lineIdx+1, rawLine)
		}

		if canCompile {
			kudaOutput.WriteString(result)
		}
	}

	return kudaOutput.String(), canCompile
}