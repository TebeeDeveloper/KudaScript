package kuda

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"runtime"
)

var KudaCompilers map[string][]string = map[string][]string{
	"windows"   : {"gcc", "clang", "cl"},
	"unix-like" : {"gcc", "clang"},
}

var kudaFlags map[string]string = map[string]string{
	"gcc"  : "-Wall -O2",
	"clang": "-Wall -O2",
	"cl"   : "/Wall /O2",
}

func detectCompiler() (string, error) {
	var platform string
	if runtime.GOOS == "windows" {
		platform = "windows"
	} else {
		platform = "unix-lik"
	}

	for _, compiler := range KudaCompilers[platform] {
		var cmd *exec.Cmd

		if platform == "windows" {
			cmd = exec.Command("where.exe", compiler)
		} else {
			cmd = exec.Command("command", "-v", compiler)
		}

		if err := cmd.Run(); err == nil {
			fmt.Printf("[Kuda-Compiler] Found: %s\n", compiler)
			return compiler, nil
		}
	}

	return "", fmt.Errorf("[Kuda-Compiler] Error: Not Found")
}

func KudaCompile(Ccode, cc, flags, files, output string) int {
	if Ccode == "" {
		fmt.Println("[Kuda-Compiler] Error: Code not found")
		return 1
	}

	var Ccompiler string
	if cc != "" {
		Ccompiler = cc
	} else {
		var err error
		Ccompiler, err = detectCompiler()
		if err != nil {
			fmt.Println("[Kuda-Compiler] Error: No compiler found (gcc/clang/cl)")
			return 1
		}
	}

	var Flags string
	if flags != "" {
		Flags = flags
	} else {
		Flags = kudaFlags[Ccompiler]
	}

	fmt.Printf("[Kuda-Compiler] Flags: %s\n", Flags)

	var args []string
	if Ccompiler == "cl" {
		args = append(args, strings.Fields(Flags)...)
		args = append(args, strings.Fields(files)...)
		args = append(args, "/TC", "/Fe:"+output, "-")
	} else {
		args = append(args, strings.Fields(Flags)...)
		args = append(args, strings.Fields(files)...)
		args = append(args, "-xc", "-", "-o", output)
	}

	var cmder *exec.Cmd = exec.Command(Ccompiler, args...)
	cmder.Stdin = strings.NewReader(Ccode)
	cmder.Stdout = os.Stdout
	cmder.Stderr = os.Stderr

	var err error = cmder.Run()
	if err != nil {
		fmt.Printf("[Kuda-Compiler] Error: Compile failed -> %v\n", err)
		return 1
	}

	fmt.Printf("[Kuda-Compiler] Success: Output -> %s\n", output)
	return 0
}
