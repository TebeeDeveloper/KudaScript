package main

import (
	"flag"
	"fmt"
	"kudaproject/kuda"
	"os"
	"time"
)

func ShowHeader() {
	fmt.Println("KudaScript | Faster | Simplier | More Convenient")
}

func ShowVersion() {
	ShowHeader()
	fmt.Println("KudaScript | 0x103  | 1.0.3")
}

func ShowHelp() {
	ShowVersion()
	fmt.Println("Usage:")
	fmt.Println("    -v -version       Show the version")
	fmt.Println("    -h -help          Show the help")
	fmt.Println("    -cc               Choose compiler")
	fmt.Println("    -flags=\"...\"      Choose the flags")
	fmt.Println("    -files=\"...\"      Include more files")
}

func ReadFile(in string) string {
	var data []byte
	var err error
	data, err = os.ReadFile(in)
	if err != nil {
		fmt.Printf("[Kuda-IO] Error: %s does not exist", in)
		return ""
	}

	var content string = string(data)
	return content
}

func main() {
	var v *bool = flag.Bool("v", false, "Show the version")
	var ver *bool = flag.Bool("version", false, "Show the version")
	var help *bool = flag.Bool("help", false, "Show the help")
	var h *bool = flag.Bool("h", false, "Show the help")
	
	var cc *string = flag.String("cc", "", "Choose compiler")
	var flags *string = flag.String("flags", "", "Choose flags")

	var files *string = flag.String("files", "", "Include more files")

	flag.Parse()

	if *v || *ver {
		ShowVersion()
		return
	}

	if *help || *h {
		ShowHelp()
		return
	}

	var args []string = flag.Args()
	if len(args) < 2 {
		ShowHelp()
		return
	}

	var inputFile string = args[0]
	var outputFile string = args[1]

	var kudaCode string = ReadFile(inputFile)

	fmt.Println("[Kuda] Translating KudaScript → C...")
	var start time.Time = time.Now()
	var cCode string
	var ccm bool 
	cCode, ccm = kuda.KudaTranslate(kudaCode, inputFile)

	var exitCode int = 0
	if ccm {
		exitCode = kuda.KudaCompile(cCode, *cc, *flags, *files, outputFile)
	}

	elapsed := time.Since(start)
	elapsedUs := elapsed.Microseconds()

	fmt.Printf("\n[Kuda] Compile time: %d µs (%.2f ms)\n", elapsedUs, elapsed.Seconds()*1000)

	
	os.Exit(exitCode)
}