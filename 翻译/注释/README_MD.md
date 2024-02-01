
<原文开始>
grab

[![GoDoc](https://godoc.org/github.com/cavaliercoder/grab?status.svg)](https://godoc.org/github.com/cavaliercoder/grab) [![Build Status](https://travis-ci.org/cavaliercoder/grab.svg?branch=master)](https://travis-ci.org/cavaliercoder/grab) [![Go Report Card](https://goreportcard.com/badge/github.com/cavaliercoder/grab)](https://goreportcard.com/report/github.com/cavaliercoder/grab)

*Downloading the internet, one goroutine at a time!*

	$ go get github.com/cavaliergopher/grab/v3

Grab is a Go package for downloading files from the internet with the following
rad features:

* Monitor download progress concurrently
* Auto-resume incomplete downloads
* Guess filename from content header or URL path
* Safely cancel downloads using context.Context
* Validate downloads using checksums
* Download batches of files concurrently
* Apply rate limiters

Requires Go v1.7+


<原文结束>

# <翻译开始>
# [![](https://godoc.org/github.com/cavaliercoder/grab?status.svg)](https://godoc.org/github.com/cavaliercoder/grab) [![](https://travis-ci.org/cavaliercoder/grab.svg?branch=master)](https://travis-ci.org/cavaliercoder/grab) [![](https://goreportcard.com/badge/github.com/cavaliercoder/grab)](https://goreportcard.com/report/github.com/cavaliercoder/grab)

* 一次一个goroutine，下载整个互联网！

	$ go get github.com/cavaliergopher/grab/v3

Grab是一个用于从互联网下载文件的Go语言包，具有以下炫酷特性：

* 并发监控下载进度
* 自动续传未完成的下载
* 根据内容头或URL路径猜测文件名
* 安全地使用context.Context取消下载
* 使用校验和验证下载内容
* 并发批量下载多个文件
* 应用速率限制器

要求Go v1.7及以上版本

# <翻译结束>


<原文开始>
Example

The following example downloads a PDF copy of the free eBook, "An Introduction
to Programming in Go" into the current working directory.

```go
resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
if err != nil {
	log.Fatal(err)
}

fmt.Println("Download saved to", resp.Filename)
```

The following, more complete example allows for more granular control and
periodically prints the download progress until it is complete.

The second time you run the example, it will auto-resume the previous download
and exit sooner.

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", "http://www.golang-book.com/public/pdf/gobook.pdf")

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	// Output:
	// Downloading http://www.golang-book.com/public/pdf/gobook.pdf...
	//   200 OK
	//   transferred 42970 / 2893557 bytes (1.49%)
	//   transferred 1207474 / 2893557 bytes (41.73%)
	//   transferred 2758210 / 2893557 bytes (95.32%)
	// Download saved to ./gobook.pdf
}
```


<原文结束>

# <翻译开始>
Example

以下示例将免费电子书“An Introduction to Programming in Go”的 PDF 副本下载到当前工作目录中。

```go
resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
if err != nil {
	log.Fatal(err)
}

fmt.Println("Download saved to", resp.Filename)
```

以下更完整的示例允许更精细的控制，并定期打印下载进度，直到完成。
第二次运行该示例时，它将自动恢复之前的下载并更快地退出。

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", "http://www.golang-book.com/public/pdf/gobook.pdf")

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	// Output:
	// Downloading http://www.golang-book.com/public/pdf/gobook.pdf...
	//   200 OK
	//   transferred 42970 / 2893557 bytes (1.49%)
	//   transferred 1207474 / 2893557 bytes (41.73%)
	//   transferred 2758210 / 2893557 bytes (95.32%)
	// Download saved to ./gobook.pdf
}
```


# <翻译结束>


<原文开始>
Design trade-offs

The primary use case for Grab is to concurrently downloading thousands of large
files from remote file repositories where the remote files are immutable.
Examples include operating system package repositories or ISO libraries.

Grab aims to provide robust, sane defaults. These are usually determined using
the HTTP specifications, or by mimicking the behavior of common web clients like
cURL, wget and common web browsers.

Grab aims to be stateless. The only state that exists is the remote files you
wish to download and the local copy which may be completed, partially completed
or not yet created. The advantage to this is that the local file system is not
cluttered unnecessarily with addition state files (like a `.crdownload` file).
The disadvantage of this approach is that grab must make assumptions about the
local and remote state; specifically, that they have not been modified by
another program.

If the local or remote file are modified outside of grab, and you download the
file again with resuming enabled, the local file will likely become corrupted.
In this case, you might consider making remote files immutable, or disabling
resume.

Grab aims to enable best-in-class functionality for more complex features
through extensible interfaces, rather than reimplementation. For example,
you can provide your own Hash algorithm to compute file checksums, or your
own rate limiter implementation (with all the associated trade-offs) to rate
limit downloads.

<原文结束>

# <翻译开始>
# 设计权衡

Grab 的主要使用场景是在远程文件仓库并发下载数千个大型不可变文件，例如操作系统包存储库或 ISO 库。

Grab 力求提供健壮且合理的默认设置。这些设置通常通过遵循 HTTP 规范确定，或者模仿 cURL、wget 等常见网络客户端以及主流网络浏览器的行为来实现。

Grab 目标是实现无状态化。仅存在的状态是您希望下载的远程文件和可能已完成、部分完成或尚未创建的本地副本。这样做的优点是可以避免本地文件系统因额外状态文件（如 .crdownload 文件）而变得杂乱无章。这种方法的缺点在于 Grab 必须对本地状态和远程状态做出假设，即它们没有被其他程序修改。

如果在 Grab 之外修改了本地或远程文件，并在启用了续传的情况下再次下载该文件，则本地文件可能会遭到损坏。在这种情况下，您可以考虑使远程文件变为不可变，或者禁用续传功能。

Grab 旨在通过可扩展接口实现复杂功能的最佳实践，而不是重新实现。例如，您可以提供自己的哈希算法来计算文件校验和，或者提供您自定义的带宽限制器实现（并接受所有相关的权衡），以限制下载速度。

# <翻译结束>

