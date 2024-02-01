package 下载类

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cavaliergopher/grab/v3/pkg/grabtest"
)

// TestFilenameResolutions 测试对于 Requests，其目标文件名能够被正确地确定，通过显式请求的路径、Content-Disposition 头部信息或 URL 路径来实现——无论目标目录是否存在与否。
func TestFilenameResolution(t *testing.T) {
	tests := []struct {
		Name               string
		Filename           string
		URL                string
		AttachmentFilename string
		Expect             string
	}{
		{"Using Request.Filename", ".testWithFilename", "/url-filename", "header-filename", ".testWithFilename"},
		{"Using Content-Disposition Header", "", "/url-filename", ".testWithHeaderFilename", ".testWithHeaderFilename"},
		{"Using Content-Disposition Header with target directory", ".test", "/url-filename", "header-filename", ".test\\header-filename"},//2024-01-02 原此处是/, 在win平台单元测试不过,改成\\
		{"Using URL Path", "", "/.testWithURLFilename?params-filename", "", ".testWithURLFilename"},
		{"Using URL Path with target directory", ".test", "/url-filename?garbage", "", ".test\\url-filename"},//2024-01-02 原此处是/, 在win平台单元测试不过,改成\\
		{"Failure", "", "", "", ""},
	}

	err := os.Mkdir(".test", 0777)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(".test")

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			opts := []grabtest.HandlerOption{}
			if test.AttachmentFilename != "" {
				opts = append(opts, grabtest.AttachmentFilename(test.AttachmentFilename))
			}
			grabtest.WithTestServer(t, func(url string) {
				req := mustNewRequest(test.Filename, url+test.URL)
				resp := X默认全局客户端.X下载(req)
				defer os.Remove(resp.X文件名)
				if err := resp.X等待错误(); err != nil {
					if test.Expect != "" || err != ERR_无法确定文件名 {
						panic(err)
					}
				} else {
					if test.Expect == "" {
						t.Errorf("expected: %v, got: %v", ERR_无法确定文件名, err)
					}
				}
				if resp.X文件名 != test.Expect {
					t.Errorf("Filename mismatch. Expected '%s', got '%s'.", test.Expect, resp.X文件名)
				}
				testComplete(t, resp)
			}, opts...)
		})
	}
}

// TestChecksums 测试校验和验证功能，确保其对有效下载和损坏下载的行为符合预期。
func TestChecksums(t *testing.T) {
	tests := []struct {
		size  int
		hash  hash.Hash
		sum   string
		match bool
	}{
		{128, md5.New(), "37eff01866ba3f538421b30b7cbefcac", true},
		{128, md5.New(), "37eff01866ba3f538421b30b7cbefcad", false},
		{1024, md5.New(), "b2ea9f7fcea831a4a63b213f41a8855b", true},
		{1024, md5.New(), "b2ea9f7fcea831a4a63b213f41a8855c", false},
		{1048576, md5.New(), "c35cc7d8d91728a0cb052831bc4ef372", true},
		{1048576, md5.New(), "c35cc7d8d91728a0cb052831bc4ef373", false},
		{128, sha1.New(), "e6434bc401f98603d7eda504790c98c67385d535", true},
		{128, sha1.New(), "e6434bc401f98603d7eda504790c98c67385d536", false},
		{1024, sha1.New(), "5b00669c480d5cffbdfa8bdba99561160f2d1b77", true},
		{1024, sha1.New(), "5b00669c480d5cffbdfa8bdba99561160f2d1b78", false},
		{1048576, sha1.New(), "ecfc8e86fdd83811f9cc9bf500993b63069923be", true},
		{1048576, sha1.New(), "ecfc8e86fdd83811f9cc9bf500993b63069923bf", false},
		{128, sha256.New(), "471fb943aa23c511f6f72f8d1652d9c880cfa392ad80503120547703e56a2be5", true},
		{128, sha256.New(), "471fb943aa23c511f6f72f8d1652d9c880cfa392ad80503120547703e56a2be4", false},
		{1024, sha256.New(), "785b0751fc2c53dc14a4ce3d800e69ef9ce1009eb327ccf458afe09c242c26c9", true},
		{1024, sha256.New(), "785b0751fc2c53dc14a4ce3d800e69ef9ce1009eb327ccf458afe09c242c26c8", false},
		{1048576, sha256.New(), "fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83", true},
		{1048576, sha256.New(), "fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c82", false},
		{128, sha512.New(), "1dffd5e3adb71d45d2245939665521ae001a317a03720a45732ba1900ca3b8351fc5c9b4ca513eba6f80bc7b1d1fdad4abd13491cb824d61b08d8c0e1561b3f7", true},
		{128, sha512.New(), "1dffd5e3adb71d45d2245939665521ae001a317a03720a45732ba1900ca3b8351fc5c9b4ca513eba6f80bc7b1d1fdad4abd13491cb824d61b08d8c0e1561b3f8", false},
		{1024, sha512.New(), "37f652be867f28ed033269cbba201af2112c2b3fd334a89fd2f757938ddee815787cc61d6e24a8a33340d0f7e86ffc058816b88530766ba6e231620a130b566c", true},
		{1024, sha512.New(), "37f652bf867f28ed033269cbba201af2112c2b3fd334a89fd2f757938ddee815787cc61d6e24a8a33340d0f7e86ffc058816b88530766ba6e231620a130b566d", false},
		{1048576, sha512.New(), "ac1d097b4ea6f6ad7ba640275b9ac290e4828cd760a0ebf76d555463a4f505f95df4f611629539a2dd1848e7c1304633baa1826462b3c87521c0c6e3469b67af", true},
		{1048576, sha512.New(), "ac1d097c4ea6f6ad7ba640275b9ac290e4828cd760a0ebf76d555463a4f505f95df4f611629539a2dd1848e7c1304633baa1826462b3c87521c0c6e3469b67af", false},
	}

	for _, test := range tests {
		var expect error
		comparison := "Match"
		if !test.match {
			comparison = "Mismatch"
			expect = ERR_文件校验失败
		}

		t.Run(fmt.Sprintf("With%s%s", comparison, test.sum[:8]), func(t *testing.T) {
			filename := fmt.Sprintf(".testChecksum-%s-%s", comparison, test.sum[:8])
			defer os.Remove(filename)

			grabtest.WithTestServer(t, func(url string) {
				req := mustNewRequest(filename, url)
				req.X设置完成后效验(test.hash, grabtest.MustHexDecodeString(test.sum), true)

				resp := X默认全局客户端.X下载(req)
				err := resp.X等待错误()
				if err != expect {
					t.Errorf("expected error: %v, got: %v", expect, err)
				}

// 确保不匹配的文件已被删除
				if !test.match {
					if _, err := os.Stat(filename); err == nil {
						t.Errorf("checksum failure not cleaned up: %s", filename)
					} else if !os.IsNotExist(err) {
						panic(err)
					}
				}

				testComplete(t, resp)
			}, grabtest.ContentLength(test.size))
		})
	}
}

// TestContentLength 确保当服务器响应与请求的长度不匹配时返回 ErrBadLength 错误。
func TestContentLength(t *testing.T) {
	size := int64(32768)
	testCases := []struct {
		Name   string
		NoHead bool
		Size   int64
		Expect int64
		Match  bool
	}{
		{"Good size in HEAD request", false, size, size, true},
		{"Good size in GET request", true, size, size, true},
		{"Bad size in HEAD request", false, size - 1, size, false},
		{"Bad size in GET request", true, size - 1, size, false},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			opts := []grabtest.HandlerOption{
				grabtest.ContentLength(int(test.Size)),
			}
			if test.NoHead {
				opts = append(opts, grabtest.MethodWhitelist("GET"))
			}

			grabtest.WithTestServer(t, func(url string) {
				req := mustNewRequest(".testSize-mismatch-head", url)
				req.X预期文件大小 = size
				resp := X默认全局客户端.X下载(req)
				defer os.Remove(resp.X文件名)
				err := resp.X等待错误()
				if test.Match {
					if err == ERR_文件长度不匹配 {
						t.Errorf("error: %v", err)
					} else if err != nil {
						panic(err)
					} else if resp.X取总字节() != size {
						t.Errorf("expected %v bytes, got %v bytes", size, resp.X取总字节())
					}
				} else {
					if err == nil {
						t.Errorf("expected: %v, got %v", ERR_文件长度不匹配, err)
					} else if err != ERR_文件长度不匹配 {
						panic(err)
					}
				}
				testComplete(t, resp)
			}, opts...)
		})
	}
}

// TestAutoResume 测试大文件的分段下载功能。
func TestAutoResume(t *testing.T) {
	segs := 8
	size := 1048576
	sum := grabtest.DefaultHandlerSHA256ChecksumBytes // grab/v3test包中的MustHexDecodeString函数将给定的十六进制字符串"fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83"进行解码，必须成功解码并返回对应的字节切片。如果解码失败，则panic。
	filename := ".testAutoResume"

	defer os.Remove(filename)

	for i := 0; i < segs; i++ {
		segsize := (i + 1) * (size / segs)
		t.Run(fmt.Sprintf("With%vBytes", segsize), func(t *testing.T) {
			grabtest.WithTestServer(t, func(url string) {
				req := mustNewRequest(filename, url)
				if i == segs-1 {
					req.X设置完成后效验(sha256.New(), sum, false)
				}
				resp := mustDo(req)
				if i > 0 && !resp.DidResume {
					t.Errorf("expected Response.DidResume to be true")
				}
				testComplete(t, resp)
			},
				grabtest.ContentLength(segsize),
			)
		})
	}

	t.Run("WithFailure", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
// 请求更小的段落
			req := mustNewRequest(filename, url)
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != ERR_文件长度不匹配 {
				t.Errorf("expected ErrBadLength for smaller request, got: %v", err)
			}
		},
			grabtest.ContentLength(size-128),
		)
	})

	t.Run("WithNoResume", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			req.X不续传 = true
			resp := mustDo(req)
			if resp.DidResume {
				t.Errorf("expected Response.DidResume to be false")
			}
			testComplete(t, resp)
		},
			grabtest.ContentLength(size+128),
		)
	})

	t.Run("WithNoResumeAndTruncate", func(t *testing.T) {
		size := size - 128
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			req.X不续传 = true
			resp := mustDo(req)
			if resp.DidResume {
				t.Errorf("expected Response.DidResume to be false")
			}
			if v := resp.X已完成字节(); v != int64(size) {
				t.Errorf("expected Response.BytesComplete: %d, got: %d", size, v)
			}
			testComplete(t, resp)
		},
			grabtest.ContentLength(size),
		)
	})

	t.Run("WithNoContentLengthHeader", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			req.X设置完成后效验(sha256.New(), sum, false)
			resp := mustDo(req)
			if !resp.DidResume {
				t.Errorf("expected Response.DidResume to be true")
			}
			if actual := resp.X取总字节(); actual != int64(size) {
				t.Errorf("expected Response.Size: %d, got: %d", size, actual)
			}
			testComplete(t, resp)
		},
			grabtest.ContentLength(size),
			grabtest.HeaderBlacklist("Content-Length"),
		)
	})

	t.Run("WithNoContentLengthHeaderAndChecksumFailure", func(t *testing.T) {
// 参考：https://github.com/cavaliergopher/grab/v3/pull/27
// （这段代码引用了GitHub上的一个项目cavaliergopher/grab的v3版本中的第27个pull request）
		size := size * 2
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			req.X设置完成后效验(sha256.New(), sum, false)
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != ERR_文件校验失败 {
				t.Errorf("expected error: %v, got: %v", ERR_文件校验失败, err)
			}
			if !resp.DidResume {
				t.Errorf("expected Response.DidResume to be true")
			}
			if actual := resp.X已完成字节(); actual != int64(size) {
				t.Errorf("expected Response.BytesComplete: %d, got: %d", size, actual)
			}
			if actual := resp.X取总字节(); actual != int64(size) {
				t.Errorf("expected Response.Size: %d, got: %d", size, actual)
			}
			testComplete(t, resp)
		},
			grabtest.ContentLength(size),
			grabtest.HeaderBlacklist("Content-Length"),
		)
	})
// TODO: 当现有文件损坏时进行测试
}

func TestSkipExisting(t *testing.T) {
	filename := ".testSkipExisting"
	defer os.Remove(filename)

	// download a file
	grabtest.WithTestServer(t, func(url string) {
		resp := mustDo(mustNewRequest(filename, url))
		testComplete(t, resp)
	})

	// redownload
	grabtest.WithTestServer(t, func(url string) {
		resp := mustDo(mustNewRequest(filename, url))
		testComplete(t, resp)

// 确保下载已恢复
		if !resp.DidResume {
			t.Fatalf("Expected download to skip existing file, but it did not")
		}

// 确保所有字节已恢复
		if resp.X取总字节() == 0 || resp.X取总字节() != resp.bytesResumed {
			t.Fatalf("Expected to skip %d bytes in redownload; got %d", resp.X取总字节(), resp.bytesResumed)
		}
	})

// 确保对预先存在的文件执行校验和检查
	grabtest.WithTestServer(t, func(url string) {
		req := mustNewRequest(filename, url)
		req.X设置完成后效验(sha256.New(), []byte{0x01, 0x02, 0x03, 0x04}, true)
		resp := X默认全局客户端.X下载(req)
		if err := resp.X等待错误(); err != ERR_文件校验失败 {
			t.Fatalf("Expected checksum error, got: %v", err)
		}
	})
}

// TestBatch 同时执行多个请求并验证响应。
func TestBatch(t *testing.T) {
	tests := 32
	size := 32768
	sum := grabtest.MustHexDecodeString("e11360251d1173650cdcd20f111d8f1ca2e412f572e8b36a4dc067121c1799b8")

// 使用4个worker，并且每个请求使用一个worker进行测试
	grabtest.WithTestServer(t, func(url string) {
		for _, workerCount := range []int{4, 0} {
// 创建请求
			reqs := make([]*X下载参数, tests)
			for i := 0; i < len(reqs); i++ {
				filename := fmt.Sprintf(".testBatch.%d", i+1)
				reqs[i] = mustNewRequest(filename, url+fmt.Sprintf("/request_%d?", i+1))
				reqs[i].X名称 = fmt.Sprintf("Test %d", i+1)
				reqs[i].X设置完成后效验(sha256.New(), sum, false)
			}

			// batch run
			responses := X默认全局客户端.X多线程下载(workerCount, reqs...)

// 监听响应
		Loop:
			for i := 0; i < len(reqs); {
				select {
				case resp := <-responses:
					if resp == nil {
						break Loop
					}
					testComplete(t, resp)
					if err := resp.X等待错误(); err != nil {
						t.Errorf("%s: %v", resp.X文件名, err)
					}

// 移除测试文件
					if resp.X是否已完成() {
						os.Remove(resp.X文件名) // ignore errors
					}
					i++
				}
			}
		}
	},
		grabtest.ContentLength(size),
	)
}

// TestCancelContext 测试在使用 context.Context 取消功能时，一批请求能够被取消。
// 请求在多种状态下被取消：正在进行中和未开始。
func TestCancelContext(t *testing.T) {
	fileSize := 134217728
	tests := 256
	client := X创建客户端()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grabtest.WithTestServer(t, func(url string) {
		reqs := make([]*X下载参数, tests)
		for i := 0; i < tests; i++ {
			req := mustNewRequest("", fmt.Sprintf("%s/.testCancelContext%d", url, i))
			reqs[i] = req.WithContext(ctx)
		}

		respch := client.X多线程下载(8, reqs...)
		time.Sleep(time.Millisecond * 500)
		cancel()
		for resp := range respch {
			defer os.Remove(resp.X文件名)

// err 应该是 context.Canceled 或 http.errRequestCanceled
			if resp.X等待错误() == nil || !strings.Contains(resp.X等待错误().Error(), "canceled") {
				t.Errorf("expected '%v', got '%v'", context.Canceled, resp.X等待错误())
			}
			if resp.X已完成字节() >= int64(fileSize) {
				t.Errorf("expected Response.BytesComplete: < %d, got: %d", fileSize, resp.X已完成字节())
			}
		}
	},
		grabtest.ContentLength(fileSize),
	)
}

// TestCancelHangingResponse 测试一个永远不会结束的请求在响应被取消时能够被终止。
func TestCancelHangingResponse(t *testing.T) {
	fileSize := 10
	client := X创建客户端()

	grabtest.WithTestServer(t, func(url string) {
		req := mustNewRequest("", fmt.Sprintf("%s/.testCancelHangingResponse", url))

		resp := client.X下载(req)
		defer os.Remove(resp.X文件名)

// 等待传输一些字节
		for resp.X已完成字节() == 0 {
			time.Sleep(50 * time.Millisecond)
		}

		done := make(chan error)
		go func() {
			done <- resp.X取消()
		}()

		select {
		case err := <-done:
			if err != context.Canceled {
				t.Errorf("Expected context.Canceled error, go: %v", err)
			}
		case <-time.After(time.Second):
			t.Fatal("response was not cancelled within 1s")
		}
		if resp.X已完成字节() == int64(fileSize) {
			t.Error("download was not supposed to be complete")
		}
	},
		grabtest.RateLimiter(1),
		grabtest.ContentLength(fileSize),
	)
}

// TestNestedDirectory 测试缺失的子目录是否会被创建。
func TestNestedDirectory(t *testing.T) {
	dir := "./.testNested/one/two/three"
	filename := ".testNestedFile"
	expect := dir + "/" + filename

	t.Run("Create", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
			resp := mustDo(mustNewRequest(expect, url+"/"+filename))
			defer os.RemoveAll("./.testNested/")
			if resp.X文件名 != expect {
				t.Errorf("expected nested Request.Filename to be %v, got %v", expect, resp.X文件名)
			}
		})
	})

	t.Run("No create", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(expect, url+"/"+filename)
			req.X不自动创建目录 = true
			resp := X默认全局客户端.X下载(req)
			err := resp.X等待错误()
			if !os.IsNotExist(err) {
				t.Errorf("expected: %v, got: %v", os.ErrNotExist, err)
			}
		})
	})
}

// TestRemoteTime 测试从远程下载的文件的时间戳能否根据远程文件的时间戳进行设置
func TestRemoteTime(t *testing.T) {
	filename := "./.testRemoteTime"
	defer os.Remove(filename)

// 随机时间，范围在 Unix 时间纪元（epoch）和当前时间之间
	expect := time.Unix(rand.Int63n(time.Now().Unix()), 0)
	grabtest.WithTestServer(t, func(url string) {
		resp := mustDo(mustNewRequest(filename, url))
		fi, err := os.Stat(resp.X文件名)
		if err != nil {
			panic(err)
		}
		actual := fi.ModTime()
		if !actual.Equal(expect) {
			t.Errorf("expected %v, got %v", expect, actual)
		}
	},
		grabtest.LastModified(expect),
	)
}

func TestResponseCode(t *testing.T) {
	filename := "./.testResponseCode"

	t.Run("With404", func(t *testing.T) {
		defer os.Remove(filename)
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			resp := X默认全局客户端.X下载(req)
			expect := X状态码错误(http.StatusNotFound)
			err := resp.X等待错误()
			if err != expect {
				t.Errorf("expected %v, got '%v'", expect, err)
			}
			if !X是否为状态码错误(err) {
				t.Errorf("expected IsStatusCodeError to return true for %T: %v", err, err)
			}
		},
			grabtest.StatusCodeStatic(http.StatusNotFound),
		)
	})

	t.Run("WithIgnoreNon2XX", func(t *testing.T) {
		defer os.Remove(filename)
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			req.X忽略错误状态码 = true
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != nil {
				t.Errorf("expected nil, got '%v'", err)
			}
		},
			grabtest.StatusCodeStatic(http.StatusNotFound),
		)
	})
}

func TestBeforeCopyHook(t *testing.T) {
	filename := "./.testBeforeCopy"
	t.Run("Noop", func(t *testing.T) {
		defer os.RemoveAll(filename)
		grabtest.WithTestServer(t, func(url string) {
			called := false
			req := mustNewRequest(filename, url)
			req.X传输开始之前回调 = func(resp *X响应) error {
				called = true
				if resp.X是否已完成() {
					t.Error("Response object passed to BeforeCopy hook has already been closed")
				}
				if resp.X取进度() != 0 {
					t.Error("Download progress already > 0 when BeforeCopy hook was called")
				}
				if resp.X取下载已持续时间() == 0 {
					t.Error("Duration was zero when BeforeCopy was called")
				}
				if resp.X已完成字节() != 0 {
					t.Error("BytesComplete already > 0 when BeforeCopy hook was called")
				}
				return nil
			}
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != nil {
				t.Errorf("unexpected error using BeforeCopy hook: %v", err)
			}
			testComplete(t, resp)
			if !called {
				t.Error("BeforeCopy hook was never called")
			}
		})
	})

	t.Run("WithError", func(t *testing.T) {
		defer os.RemoveAll(filename)
		grabtest.WithTestServer(t, func(url string) {
			testError := errors.New("test")
			req := mustNewRequest(filename, url)
			req.X传输开始之前回调 = func(resp *X响应) error {
				return testError
			}
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != testError {
				t.Errorf("expected error '%v', got '%v'", testError, err)
			}
			if resp.X已完成字节() != 0 {
				t.Errorf("expected 0 bytes completed for canceled BeforeCopy hook, got %d",
					resp.X已完成字节())
			}
			testComplete(t, resp)
		})
	})

// 断言在`BeforeCopy`钩子有机会取消请求之前，现有本地文件不会被截断
	t.Run("NoTruncate", func(t *testing.T) {
		tfile, err := ioutil.TempFile("", "grab_client_test.*.file")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tfile.Name())

		const size = 128
		_, err = tfile.Write(bytes.Repeat([]byte("x"), size))
		if err != nil {
			t.Fatal(err)
		}

		grabtest.WithTestServer(t, func(url string) {
			called := false
			req := mustNewRequest(tfile.Name(), url)
			req.X不续传 = true
			req.X传输开始之前回调 = func(resp *X响应) error {
				called = true
				fi, err := tfile.Stat()
				if err != nil {
					t.Errorf("failed to stat temp file: %v", err)
					return nil
				}
				if fi.Size() != size {
					t.Errorf("expected existing file size of %d bytes "+
						"prior to BeforeCopy hook, got %d", size, fi.Size())
				}
				return nil
			}
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != nil {
				t.Errorf("unexpected error using BeforeCopy hook: %v", err)
			}
			testComplete(t, resp)
			if !called {
				t.Error("BeforeCopy hook was never called")
			}
		})
	})
}

func TestAfterCopyHook(t *testing.T) {
	filename := "./.testAfterCopy"
	t.Run("Noop", func(t *testing.T) {
		defer os.RemoveAll(filename)
		grabtest.WithTestServer(t, func(url string) {
			called := false
			req := mustNewRequest(filename, url)
			req.X传输完成之后回调 = func(resp *X响应) error {
				called = true
				if resp.X是否已完成() {
					t.Error("Response object passed to AfterCopy hook has already been closed")
				}
				if resp.X取进度() <= 0 {
					t.Error("Download progress was 0 when AfterCopy hook was called")
				}
				if resp.X取下载已持续时间() == 0 {
					t.Error("Duration was zero when AfterCopy was called")
				}
				if resp.X已完成字节() <= 0 {
					t.Error("BytesComplete was 0 when AfterCopy hook was called")
				}
				return nil
			}
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != nil {
				t.Errorf("unexpected error using AfterCopy hook: %v", err)
			}
			testComplete(t, resp)
			if !called {
				t.Error("AfterCopy hook was never called")
			}
		})
	})

	t.Run("WithError", func(t *testing.T) {
		defer os.RemoveAll(filename)
		grabtest.WithTestServer(t, func(url string) {
			testError := errors.New("test")
			req := mustNewRequest(filename, url)
			req.X传输完成之后回调 = func(resp *X响应) error {
				return testError
			}
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != testError {
				t.Errorf("expected error '%v', got '%v'", testError, err)
			}
			if resp.X已完成字节() <= 0 {
				t.Errorf("ByteCompleted was %d after AfterCopy hook was called",
					resp.X已完成字节())
			}
			testComplete(t, resp)
		})
	})
}

func TestIssue37(t *testing.T) {
// 参考：https://github.com/cavaliergopher/grab/v3/issues/37
	filename := "./.testIssue37"
	largeSize := int64(2097152)
	smallSize := int64(1048576)
	defer os.RemoveAll(filename)

// 下载大文件
	grabtest.WithTestServer(t, func(url string) {
		resp := mustDo(mustNewRequest(filename, url))
		if resp.X取总字节() != largeSize {
			t.Errorf("expected response size: %d, got: %d", largeSize, resp.X取总字节())
		}
	}, grabtest.ContentLength(int(largeSize)))

// 下载相同文件的新版本，体积更小
	grabtest.WithTestServer(t, func(url string) {
		req := mustNewRequest(filename, url)
		req.X不续传 = true
		resp := mustDo(req)
		if resp.X取总字节() != smallSize {
			t.Errorf("expected response size: %d, got: %d", smallSize, resp.X取总字节())
		}

// 本地文件应被截断且不应恢复
		if resp.DidResume {
			t.Errorf("expected download to truncate, resumed instead")
		}
	}, grabtest.ContentLength(int(smallSize)))

	fi, err := os.Stat(filename)
	if err != nil {
		t.Fatal(err)
	}
	if fi.Size() != int64(smallSize) {
		t.Errorf("expected file size %d, got %d", smallSize, fi.Size())
	}
}

// TestHeadBadStatus 验证了如果 HEAD 请求返回非 200 状态码时，可以被忽略，
// 并且只要后续的 GET 请求成功，整个流程就算成功。
//
// 解决问题：https://github.com/cavaliergopher/grab/v3/issues/43
func TestHeadBadStatus(t *testing.T) {
	expect := http.StatusOK
	filename := ".testIssue43"

	statusFunc := func(r *http.Request) int {
		if r.Method == "HEAD" {
			return http.StatusForbidden
		}
		return http.StatusOK
	}

	grabtest.WithTestServer(t, func(url string) {
		testURL := fmt.Sprintf("%s/%s", url, filename)
		resp := mustDo(mustNewRequest("", testURL))
		if resp.HTTP响应.StatusCode != expect {
			t.Errorf(
				"expected status code: %d, got:% d",
				expect,
				resp.HTTP响应.StatusCode)
		}
	},
		grabtest.StatusCode(statusFunc),
	)
}

// TestMissingContentLength ensures that the Response.Size is correct for
// transfers where the remote server does not send a Content-Length header.
//
// TestAutoResume also covers cases with checksum validation.
//
// Kudos to Setnička Jiří <Jiri.Setnicka@ysoft.com> for identifying and raising
// a solution to this issue. Ref: https://github.com/cavaliergopher/grab/v3/pull/27
func TestMissingContentLength(t *testing.T) {
// expectSize 必须足够大，以确保 DefaultClient.Do 不会在返回 Response 之前预加载整个 body 并计算 ContentLength。
	expectSize := 1048576
	opts := []grabtest.HandlerOption{
		grabtest.ContentLength(expectSize),
		grabtest.HeaderBlacklist("Content-Length"),
		grabtest.TimeToFirstByte(time.Millisecond * 100), // 延迟初始读取
	}
	grabtest.WithTestServer(t, func(url string) {
		req := mustNewRequest(".testMissingContentLength", url)
		req.X设置完成后效验(
			md5.New(),
			grabtest.DefaultHandlerMD5ChecksumBytes,
			false)
		resp := X默认全局客户端.X下载(req)

// 确保远程服务器未发送内容长度头
		if v := resp.HTTP响应.Header.Get("Content-Length"); v != "" {
			panic(fmt.Sprintf("http header content length must be empty, got: %s", v))
		}
		if v := resp.HTTP响应.ContentLength; v != -1 {
			panic(fmt.Sprintf("http response content length must be -1, got: %d", v))
		}

// 在完成之前，响应大小应为-1
		if resp.X取总字节() != -1 {
			t.Errorf("expected response size: -1, got: %d", resp.X取总字节())
		}

// 等待完成进行阻塞
		if err := resp.X等待错误(); err != nil {
			panic(err)
		}

// 完成时，响应大小应为实际传输大小
		if resp.X取总字节() != int64(expectSize) {
			t.Errorf("expected response size: %d, got: %d", expectSize, resp.X取总字节())
		}
	}, opts...)
}

func TestNoStore(t *testing.T) {
	filename := ".testSubdir/testNoStore"
	t.Run("DefaultCase", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest(filename, url)
			req.X不写入本地文件系统 = true
			req.X设置完成后效验(md5.New(), grabtest.DefaultHandlerMD5ChecksumBytes, true)
			resp := mustDo(req)

// 确保 Response.Bytes 的内容正确无误，并且可以重新读取
			b, err := resp.X等待完成后取字节集()
			if err != nil {
				panic(err)
			}
			grabtest.AssertSHA256Sum(
				t,
				grabtest.DefaultHandlerSHA256ChecksumBytes,
				bytes.NewReader(b),
			)

// 确保 Response.Open 流是正确的，并且可以重新读取
			r, err := resp.X等待完成后打开文件()
			if err != nil {
				panic(err)
			}
			defer r.Close()
			grabtest.AssertSHA256Sum(
				t,
				grabtest.DefaultHandlerSHA256ChecksumBytes,
				r,
			)

// Response.Filename 应该仍然被设置
			if resp.X文件名 != filename {
				t.Errorf("expected Response.Filename: %s, got: %s", filename, resp.X文件名)
			}

// 确保没有文件被写入
			paths := []string{
				filename,
				filepath.Base(filename),
				filepath.Dir(filename),
				resp.X文件名,
				filepath.Base(resp.X文件名),
				filepath.Dir(resp.X文件名),
			}
			for _, path := range paths {
				_, err := os.Stat(path)
				if !os.IsNotExist(err) {
					t.Errorf(
						"expect error: %v, got: %v, for path: %s",
						os.ErrNotExist,
						err,
						path)
				}
			}
		})
	})

	t.Run("ChecksumValidation", func(t *testing.T) {
		grabtest.WithTestServer(t, func(url string) {
			req := mustNewRequest("", url)
			req.X不写入本地文件系统 = true
			req.X设置完成后效验(
				md5.New(),
				grabtest.MustHexDecodeString("deadbeefcafebabe"),
				true)
			resp := X默认全局客户端.X下载(req)
			if err := resp.X等待错误(); err != ERR_文件校验失败 {
				t.Errorf("expected error: %v, got: %v", ERR_文件校验失败, err)
			}
		})
	})
}
