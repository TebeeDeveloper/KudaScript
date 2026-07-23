# KudaScript
-   Cáo ống tre - Kuda-gitsune
-   Native C

## Spirit
-    Fast as C
-    Simple as Script

## Examples
```rust
fn main() int
    printf("Hello World\n")
end
```

## KudaScript - Combines the speed & stability of C with the simplicity of scripting
-   Current version: 0x103.
-   License: __GNU GPL v3__.
-   Core language: Pure standard C.

## Why KudaScript
-    Only upgrade C, no garbage
-    Determinism syntax, no abstract
-    Absolute secure your source
-    Native C speed
## Unlike Cython/Nuitka
    no intermediate .c files written to disk — all translation happens in memory.

## Intro
KudaScript is a new programming language built with clear goals:
-   Full power and performance of pure C.
-   Lightweight syntax, easy to read & learn like a scripting-language.
-   Inherit all of ecosys of std C.
-   Compile to standalone execute file, no external dependencies required.

## Standard Libs
In the first version, you can use:
-   stdio.h
-   stdint.h / stddef.h / stdbool.h
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
cd KudaScript/src

go build -o kuda main.go
```
If you use Windows, add `kuda` to enviroment variables PATH.

### Checking version
```bash
kuda --version
```
The result must be:
```plaintext
KudaScript | Faster | Simplier | More Convenient
KudaScript | 0x103  | 1.0.3
```

### Syntax Compiler
```bash
kuda [kdflag] <source.kuda> <output_name>
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
kuda -cc=cl -flags="/O2" main.kuda main.exe
```

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
