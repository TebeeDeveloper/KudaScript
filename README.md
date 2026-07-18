# KudaScript
-   Cáo ống tre - Kuda-gitsune
-   Native C

## Examples
```rust
fn main() {
    printf("Hello World\n")
}
```

## KudaScript - Combines the speed & stability of C with the simplicity of scripting
-   Current version: 0x100.
-   License: __GNU GPL v3__.
-   Core language: Pure standard C.

## Intro
KudaScript is once of milions programming-language latest, within the target:
-   Pure C, and strength as same as absolute of C.
-   Lightweight syntax, easy to read & learn like a scripting-language.
-   Inherit all of ecosys of std C.
-   Compile to standalone execute file, non-dependend enviroment.

## Standard Libs
In the first version, you can use:
-   stdio.h
-   stdin.h / stddef.h / stdbool.h
-   stdlib.h
-   string.h
-   stdarg.h
-   math.h
-   time.h

## Install & Use
### System require
Need one of some compiler on your computer:
-   gcc (most)
-   clang
-   cl (only MSVC on Windows)

### Build from source (free! trust me bro)
```bash
git clone https://github.com/TebeeDeveloper/KudaScript.git
cd KudaScript

make

gcc -o kuda src/*.c -std=c99
```
If you use Windows, add `kuda` to enviroment variables PATH.

### Checking version
```bash
kuda --version
```
The result must be:
```plaintext
KudaScript . ver 0x100.
Combines speed/stability of C with simplicity of scripting
Auto-detected compiler priority: gcc > clang > cl
```

### Syntax Compiler
```bash
kuda <source.kuda> <output_name> [compiler_flags]
```

### Basic Examples
```bash
kuda main.kuda main
./main
```

### Compiler Flags
All params write back of your output file.
```bash
kuda main.kuda main -O3 -march=native
kuda main.kuda main_debug -g -ggdb
kuda main.kuda main -Wall -Wextra -Wpedantic
kuda game.kuda game -lSDL2 -lpthread

kuda main.kuda main -cc=gcc -O2
kuda main.kuda main -cc=clang -O2
kuda main.kuda main.exe -cc=cl /O2
```
$\color{red}Caution$: The -cc parameter must always be at position 4, right after kuda, source file, output file. If omitted compiler will be auto-detected.

## Examples of Program
### 1. Basic Output
```c
fn main() {
    let name string = "KudaScript"
    let version int = 0x100

    printf("=== %s ===\n", name)
    printf("Version: ver 0x%03X\n", version)
    printf("Speed of C, simplicity of scripting!\n")

    return 0
}
```

### 2. Loop & Math
```c
fn main() {
    printf("Squares & square roots from 1 to 5:\n")

    for (let i int = 0; i <= 5; i++) {
        let square int = i * i
        let root int = sqrt(i)
        printf("%d → square: %d | root: %.2f\n", i, square, root)
    }

    return 0
}
```

### 3. String Operations
```rs
fn main() {
    let text string = "KudaScript Open Source"
    let keyword string = "Open Source"

    printf("Original text: %s\n", text)
    printf("Length: %u characters\n", strlen(text))

    if (strstr(text, keyword) != NULL) {
        printf("Found substring: '%s'\n", keyword)
    } else {
        printf("Substring not found\n")
    }

    return 0
}
```

### 4. File I/O
```rs
fn main() {
    let path string = "test.txt"

    let write_ptr FILE* = fopen(path, "w")
    if write_ptr != NULL {
        fputs("Hello from KudaScript!\n", write_ptr)
        fputs("Written via stdio.h directly\n", write_ptr)
        fclose(write_ptr)
        printf("Write successful!\n")
    }

    let read_ptr FILE* = fopen(path, "r")
    if read_ptr != NULL {
        printf("\nFile content:\n")
        let line: char[256]
        while fgets(line, 256, read_ptr) != NULL {
            printf("%s", line)
        }
        fclose(read_ptr)
    }

    return 0
}
```

### 5. Custom Function & Random
```rs
fn random_range(min: int, max: int) int {
    return min + (rand() % (max - min + 1))
}

fn main() {
    srand(time(NULL))

    printf("5 random numbers (1–100):\n")
    for (let i int = 0; i < 5; i++) {
        printf("- %d\n", random_range(1, 100))
    }

    return 0
}
```

## Contribute
This project released under the GNU GPL v3.0, all contributions are welcome!

The only rule:

-    $\color{red}Must not reduce execution speed$ compared to the orriginal version.
Upholding the spirit of "fast as C, simple as scripting."

## License
See details in the accompanying LICENSE file.
