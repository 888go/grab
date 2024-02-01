
<原文开始>
// TestFilenameResolutions tests that the destination filename for Requests can
// be determined correctly, using an explicitly requested path,
// Content-Disposition headers or a URL path - with or without an existing
// target directory.
<原文结束>

# <翻译开始>
// TestFilenameResolutions 测试对于 Requests，其目标文件名能够被正确地确定，通过显式请求的路径、Content-Disposition 头部信息或 URL 路径来实现——无论目标目录是否存在与否。
# <翻译结束>


<原文开始>
// TestChecksums checks that checksum validation behaves as expected for valid
// and corrupted downloads.
<原文结束>

# <翻译开始>
// TestChecksums 测试校验和验证功能，确保其对有效下载和损坏下载的行为符合预期。
# <翻译结束>


<原文开始>
				// ensure mismatch file was deleted
<原文结束>

# <翻译开始>
// 确保不匹配的文件已被删除
# <翻译结束>


<原文开始>
// TestContentLength ensures that ErrBadLength is returned if a server response
// does not match the requested length.
<原文结束>

# <翻译开始>
// TestContentLength 确保当服务器响应与请求的长度不匹配时返回 ErrBadLength 错误。
# <翻译结束>


<原文开始>
// TestAutoResume tests segmented downloading of a large file.
<原文结束>

# <翻译开始>
// TestAutoResume 测试大文件的分段下载功能。
# <翻译结束>


<原文开始>
			// request smaller segment
<原文结束>

# <翻译开始>
// 请求更小的段落
# <翻译结束>


<原文开始>
		// ref: https://github.com/cavaliergopher/grab/v3/pull/27
<原文结束>

# <翻译开始>
// 参考：https://github.com/cavaliergopher/grab/v3/pull/27
// （这段代码引用了GitHub上的一个项目cavaliergopher/grab的v3版本中的第27个pull request）
# <翻译结束>


<原文开始>
	// TODO: test when existing file is corrupted
<原文结束>

# <翻译开始>
// TODO: 当现有文件损坏时进行测试
# <翻译结束>


<原文开始>
		// ensure download was resumed
<原文结束>

# <翻译开始>
// 确保下载已恢复
# <翻译结束>


<原文开始>
		// ensure all bytes were resumed
<原文结束>

# <翻译开始>
// 确保所有字节已恢复
# <翻译结束>


<原文开始>
	// ensure checksum is performed on pre-existing file
<原文结束>

# <翻译开始>
// 确保对预先存在的文件执行校验和检查
# <翻译结束>


<原文开始>
// TestBatch executes multiple requests simultaneously and validates the
// responses.
<原文结束>

# <翻译开始>
// TestBatch 同时执行多个请求并验证响应。
# <翻译结束>


<原文开始>
	// test with 4 workers and with one per request
<原文结束>

# <翻译开始>
// 使用4个worker，并且每个请求使用一个worker进行测试
# <翻译结束>


<原文开始>
			// create requests
<原文结束>

# <翻译开始>
// 创建请求
# <翻译结束>


<原文开始>
			// listen for responses
<原文结束>

# <翻译开始>
// 监听响应
# <翻译结束>


<原文开始>
					// remove test file
<原文结束>

# <翻译开始>
// 移除测试文件
# <翻译结束>


<原文开始>
// TestCancelContext tests that a batch of requests can be cancel using a
// context.Context cancellation. Requests are cancelled in multiple states:
// in-progress and unstarted.
<原文结束>

# <翻译开始>
// TestCancelContext 测试在使用 context.Context 取消功能时，一批请求能够被取消。
// 请求在多种状态下被取消：正在进行中和未开始。
# <翻译结束>


<原文开始>
			// err should be context.Canceled or http.errRequestCanceled
<原文结束>

# <翻译开始>
// err 应该是 context.Canceled 或 http.errRequestCanceled
# <翻译结束>


<原文开始>
// TestCancelHangingResponse tests that a never ending request is terminated
// when the response is cancelled.
<原文结束>

# <翻译开始>
// TestCancelHangingResponse 测试一个永远不会结束的请求在响应被取消时能够被终止。
# <翻译结束>


<原文开始>
		// Wait for some bytes to be transferred
<原文结束>

# <翻译开始>
// 等待传输一些字节
# <翻译结束>


<原文开始>
// TestNestedDirectory tests that missing subdirectories are created.
<原文结束>

# <翻译开始>
// TestNestedDirectory 测试缺失的子目录是否会被创建。
# <翻译结束>


<原文开始>
// TestRemoteTime tests that the timestamp of the downloaded file can be set
// according to the timestamp of the remote file.
<原文结束>

# <翻译开始>
// TestRemoteTime 测试从远程下载的文件的时间戳能否根据远程文件的时间戳进行设置
# <翻译结束>


<原文开始>
	// random time between epoch and now
<原文结束>

# <翻译开始>
// 随机时间，范围在 Unix 时间纪元（epoch）和当前时间之间
# <翻译结束>


<原文开始>
	// Assert that an existing local file will not be truncated prior to the
	// BeforeCopy hook has a chance to cancel the request
<原文结束>

# <翻译开始>
// 断言在`BeforeCopy`钩子有机会取消请求之前，现有本地文件不会被截断
# <翻译结束>


<原文开始>
	// ref: https://github.com/cavaliergopher/grab/v3/issues/37
<原文结束>

# <翻译开始>
// 参考：https://github.com/cavaliergopher/grab/v3/issues/37
# <翻译结束>


<原文开始>
	// download large file
<原文结束>

# <翻译开始>
// 下载大文件
# <翻译结束>


<原文开始>
	// download new, smaller version of same file
<原文结束>

# <翻译开始>
// 下载相同文件的新版本，体积更小
# <翻译结束>


<原文开始>
		// local file should have truncated and not resumed
<原文结束>

# <翻译开始>
// 本地文件应被截断且不应恢复
# <翻译结束>


<原文开始>
// TestHeadBadStatus validates that HEAD requests that return non-200 can be
// ignored and succeed if the GET requests succeeeds.
//
// Fixes: https://github.com/cavaliergopher/grab/v3/issues/43
<原文结束>

# <翻译开始>
// TestHeadBadStatus 验证了如果 HEAD 请求返回非 200 状态码时，可以被忽略，
// 并且只要后续的 GET 请求成功，整个流程就算成功。
//
// 解决问题：https://github.com/cavaliergopher/grab/v3/issues/43
# <翻译结束>


<原文开始>
	// expectSize must be sufficiently large that DefaultClient.Do won't prefetch
	// the entire body and compute ContentLength before returning a Response.
<原文结束>

# <翻译开始>
// expectSize 必须足够大，以确保 DefaultClient.Do 不会在返回 Response 之前预加载整个 body 并计算 ContentLength。
# <翻译结束>


<原文开始>
// delay for initial read
<原文结束>

# <翻译开始>
// 延迟初始读取
# <翻译结束>


<原文开始>
		// ensure remote server is not sending content-length header
<原文结束>

# <翻译开始>
// 确保远程服务器未发送内容长度头
# <翻译结束>


<原文开始>
		// before completion, response size should be -1
<原文结束>

# <翻译开始>
// 在完成之前，响应大小应为-1
# <翻译结束>


<原文开始>
		// block for completion
<原文结束>

# <翻译开始>
// 等待完成进行阻塞
# <翻译结束>


<原文开始>
		// on completion, response size should be actual transfer size
<原文结束>

# <翻译开始>
// 完成时，响应大小应为实际传输大小
# <翻译结束>


<原文开始>
			// ensure Response.Bytes is correct and can be reread
<原文结束>

# <翻译开始>
// 确保 Response.Bytes 的内容正确无误，并且可以重新读取
# <翻译结束>


<原文开始>
			// ensure Response.Open stream is correct and can be reread
<原文结束>

# <翻译开始>
// 确保 Response.Open 流是正确的，并且可以重新读取
# <翻译结束>


<原文开始>
			// Response.Filename should still be set
<原文结束>

# <翻译开始>
// Response.Filename 应该仍然被设置
# <翻译结束>


<原文开始>
			// ensure no files were written
<原文结束>

# <翻译开始>
// 确保没有文件被写入
# <翻译结束>


<原文开始>
//grab/v3test.MustHexDecodeString("fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83")
<原文结束>

# <翻译开始>
// grab/v3test包中的MustHexDecodeString函数将给定的十六进制字符串"fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83"进行解码，必须成功解码并返回对应的字节切片。如果解码失败，则panic。
# <翻译结束>

