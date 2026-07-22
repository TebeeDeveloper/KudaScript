# KudaScript
-   Cáo ống tre - Kuda-gitsune
-   Native C

## Examples
```rust
fn main() int
    printf("Hello World\n")
end
```

## KudaScript - Combines the speed & stability of C with the simplicity of scripting
-   Current version: 0x100.
-   License: __GNU GPL v3__.
-   Core language: Pure standard C.

## Intro
KudaScript is a new programming language built with clear goals:
-   Full power and performance of pure C.
-   Lightweight syntax, easy to read & learn like a scripting-language.
-   Inherit all of ecosys of std C.
-   Compile to standalone execute file, no external dependencies required.

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
kuda -flags="-O3 -march=native" main.kuda main
kuda -flags="-g -ggdb" main.kuda main_debug
kuda -flags="-Wall Wextra -Wpedantic" main.kuda main
kuda -flags="-lSDL2 -lpthread" game.kuda game

kuda -cc=gcc -flags="-O2" main.kuda main
kuda -cc=clang -flags="-O2" main.kuda main
kuda -cc=cl -flags="/O2 "main.kuda main.exe
```
Caution: The -cc parameter must always be at position 4, right after kuda, source file, output file. If omitted compiler will be auto-detected.

## Examples of Program
### 1. Basic Output
```c
fn main() int
    var name string = "KudaScript"
    var version int = 0x100

    printf("=== %s ===\n", name)
    printf("Version: ver 0x%03X\n", version)
    printf("Speed of C, simplicity of scripting!\n")
end
```

### 2. Loop & Math
```c
fn main() int
    printf("Squares & square roots from 1 to 5:\n")

    for i int = 0 5 2 do
        var square int = i * i
        var root int = sqrt(i)
        printf("%d → square: %d | root: %.2f\n", i, square, root)
    end
    return 0
end
```

### 3. String Operations
```rs
fn main() int
    var text str = "KudaScript Open Source"
    var keyword str = "Open Source"

    printf("Original text: %s\n", text)
    printf("Length: %u characters\n", strlen(text))

    if strstr(text, keyword) != NULL then
        printf("Found substring: '%s'\n", keyword)
    else
        printf("Substring not found\n")
    end
end
```

### 4. File I/O
```rs
fn main() int
    var path str = "test.txt"

    var write_ptr FILE* = fopen(path, "w")
    if write_ptr != NULL then
        fputs("Hello from KudaScript!\n", write_ptr)
        fputs("Written via stdio.h directly\n", write_ptr)
        fclose(write_ptr)
        printf("Write successful!\n")
    end

    var read_ptr FILE* = fopen(path, "r")
    if read_ptr != NULL then
        printf("\nFile content:\n")
        var line char[256]
        while fgets(line, 256, read_ptr) != NULL do
            printf("%s", line)
        end
        fclose(read_ptr)
    end
end
```

### 5. Custom Function & Random
```rs
fn random_range(min, max int) int
    return min + (rand() % (max - min + 1))
end

fn main() int
    srand(time(NULL))
    printf("5 random numbers (1–100):\n")
    for i int = 0 5 1 do
        printf("- %d\n", random_range(1, 100))
    end
end
```

## Contribute
This project released under the GNU GPL v3.0, all contributions are welcome!

The only rule:

-    Must not reduce execution speed compared to the orriginal version.
Upholding the spirit of "fast as C, simple as scripting."

## License
See details in the accompanying LICENSE file.
