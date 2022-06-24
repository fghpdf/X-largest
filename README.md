# X-largest
Find X-largest values in a large file

## Background
Get the X-largest values in a large file.
Each line in the file contains <unique identifier> and <number value>.

Example:
```
1426828011 9 
1426828028 350
1426828037 25
1426828056 231
1426828058 109
1426828066 111
```

## Plan

### Top-X
To find top-X values, we can use a heap.

First, we will maintain a min-heap of size X.

Then, we can put elements into the heap.

If the size of the heap is bigger than X, we can remove the top element of the heap.

So the time complexity and space complexity depends on the size of the heap.

| Complexity | Value    |
|------------|----------|
| Time       | O(NlogX) |
| Space      | O(X)     |


### Huge File
If we have a huge file, we can not read all the lines into memory.

We can split the huge file into many smaller files by hash.

Then we can read each file and put the elements into the heap.

Now we have the top x elements of each small file.

Finally, we can merge the top x elements and get the top x elements of the huge file.

![](https://s3.bmp.ovh/imgs/2022/06/23/5e3e82b41a900254.png)

By the way, I prepared a tool to generate a huge file.

It will generate a unique identifier by [snowflake](https://en.wikipedia.org/wiki/Snowflake_ID) and a random number into the huge file.

The tool will pick the line whose number >= 99900 and write it into a new file.

This file named `gen_top_x.txt` can help us to test more conveniently.

#### 🧮Calculate the number of files
Goroutine can use the cores of the machine.

So we can calculate the number of files by the number of cores.

Assume a core can take `16` times the memory of the original dataset.
And we can use 1GB of memory.

So our formula is:

`tmpFileSize := MaxMemorySize / (coreCount * 16)`

`processNum := fileSize / tmpFileSize`

e.g. we have 1 GB of the huge file, and we have 4 cores.

`temp file size = 1 GB / (16 * 4) = 0.015625 GB = 16 MB`

`process num = 1 GB / 16 MB = 64`

## Install and Run
First, make sure you installed the golang:
[golang install](https://go.dev/doc/install)

Second, you can run the following command to run:
```shell
make build
```

You will get two bin files in the `dist` folder.
1. generate.bin can generate a huge file to help you test.
2. main.bin is a cli tool to get the top x elements.
* run `./dist/main.bin -h` to get help for this command.
* run `./dist/main.bin -file=<file path> -x=<top xx>` to get the top x elements of the file.
* run `./dist/main.bin -x=<top x>` to get the top x elements of the stdin.
3. main.bin will print the top x elements to stdout.

## Integration Test and Unit Test
To test the program, you can run the following command:
```shell
make test
```

The test will generate a huge file and then get the top x elements of the huge file.

It will take a long time to generate a huge file and split it, just relax ☕️.

When everything is done, you can get the coverage of the program.
![](https://s3.bmp.ovh/imgs/2022/06/23/fea5c317c57a978e.png)

The test report is:
```
PASS
coverage: 66.8% of statements
total     (statements)		87.8%
```

The test program can cover the main process of the program.

Integration test is `main_test.go`.

Unit tests are `topx/*_test.go`.

## CI Integration Test
GitHub action is a great tool to test the program.

![](https://s3.bmp.ovh/imgs/2022/06/24/40325563289c1109.jpg)

## Performance
On my personal computer, I got the following performance using 1GB of huge file:

`./dist/main.bin -file=tmp/gen_records.txt -x=394`

| Step            | Time   |
|-----------------|--------|
| Split 165 Files | 4m1.5s |
| Top X           | 21.75s |

My computer info:
* OS: macOS 12.2.1 21D62 arm64
* CPU: Apple M1 Pro
* Memory: 32768MiB

## Tips
Due to limited time, I didn't do much optimization on splitting and test coverage.