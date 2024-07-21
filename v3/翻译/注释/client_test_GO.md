
<原文开始>
// TestFilenameResolutions tests that the destination filename for Requests can
// be determined correctly, using an explicitly requested path,
// Content-Disposition headers or a URL path - with or without an existing
// target directory.
<原文结束>

# <翻译开始>
// TestFilenameResolutions 测试Requests的目标文件名可以正确确定，使用显式请求的路径，
// 内容Disposition头或URL路径 - 无论是否存在目标目录。
// md5:37bc33cfb322c6b2
# <翻译结束>


<原文开始>
// TestChecksums checks that checksum validation behaves as expected for valid
// and corrupted downloads.
<原文结束>

# <翻译开始>
// TestChecksums 验证校验和验证是否对有效和损坏的下载行为如预期。
// md5:28a6c4d2d2ab9742
# <翻译结束>


<原文开始>
// ensure mismatch file was deleted
<原文结束>

# <翻译开始>
// 确保不匹配的文件被删除 md5:993736392820d653
# <翻译结束>


<原文开始>
// TestContentLength ensures that ErrBadLength is returned if a server response
// does not match the requested length.
<原文结束>

# <翻译开始>
// TestContentLength 确保当服务器响应的长度与请求的长度不匹配时，会返回ErrBadLength错误。
// md5:fe1357f78b777a7b
# <翻译结束>


<原文开始>
// TestAutoResume tests segmented downloading of a large file.
<原文结束>

# <翻译开始>
// TestAutoResume 测试大文件的分段下载。 md5:c5ddc10d54e3f03a
# <翻译结束>


<原文开始>
//grab/v3test.MustHexDecodeString("fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83")
<原文结束>

# <翻译开始>
// 使用/v3test.MustHexDecodeString函数解码字符串"fbbab289f7f94b25736c58be46a994c441fd02552cc6022352e3d86d2fab7c83" md5:b0af387b03689666
# <翻译结束>


<原文开始>
// request smaller segment
<原文结束>

# <翻译开始>
// 请求更小的段 md5:dc957cfdddcbd506
# <翻译结束>


<原文开始>
// ref: https://github.com/cavaliergopher/grab/v3/pull/27
<原文结束>

# <翻译开始>
// 参考：https://github.com/cavaliergopher/grab/v3/pull/27 md5:12880e56476ca7f8
# <翻译结束>


<原文开始>
// TODO: test when existing file is corrupted
<原文结束>

# <翻译开始>
// 待办：测试当现有文件已损坏时的情况 md5:d8846b3f450d289b
# <翻译结束>


<原文开始>
// ensure download was resumed
<原文结束>

# <翻译开始>
// 确保下载已恢复 md5:05f6579b7aaf08cd
# <翻译结束>


<原文开始>
// ensure all bytes were resumed
<原文结束>

# <翻译开始>
// 确保所有字节都已恢复 md5:54f7e6b7bb977011
# <翻译结束>


<原文开始>
// ensure checksum is performed on pre-existing file
<原文结束>

# <翻译开始>
// 确保对已存在的文件进行校验和计算 md5:22b0cdf6b7369f29
# <翻译结束>


<原文开始>
// TestBatch executes multiple requests simultaneously and validates the
// responses.
<原文结束>

# <翻译开始>
// TestBatch 同时执行多个请求，并验证响应结果。
// md5:910be37f5c8be7a6
# <翻译结束>


<原文开始>
// test with 4 workers and with one per request
<原文结束>

# <翻译开始>
// 使用4个工人，并且每个请求一个工人进行测试 md5:5f16ec2ee1221663
# <翻译结束>


<原文开始>
// TestCancelContext tests that a batch of requests can be cancel using a
// context.Context cancellation. Requests are cancelled in multiple states:
// in-progress and unstarted.
<原文结束>

# <翻译开始>
// TestCancelContext 测试使用 context.Context 取消一批请求。在不同状态下取消请求：进行中和未开始。
// md5:0b5a48f04c400cd8
# <翻译结束>


<原文开始>
// err should be context.Canceled or http.errRequestCanceled
<原文结束>

# <翻译开始>
// err 应该是 context.Canceled 或者 http.errRequestCanceled md5:eaef91eab132bf47
# <翻译结束>


<原文开始>
// TestCancelHangingResponse tests that a never ending request is terminated
// when the response is cancelled.
<原文结束>

# <翻译开始>
// TestCancelHangingResponse 测试当响应被取消时，一个永无止境的请求是否会被终止。
// md5:338db5a2fa65b059
# <翻译结束>


<原文开始>
// Wait for some bytes to be transferred
<原文结束>

# <翻译开始>
// 等待一些字节被传输 md5:4a22648321a6fd5a
# <翻译结束>


<原文开始>
// TestNestedDirectory tests that missing subdirectories are created.
<原文结束>

# <翻译开始>
// TestNestedDirectory 测试缺失的子目录是否会被创建。 md5:bd508d1540240344
# <翻译结束>


<原文开始>
// TestRemoteTime tests that the timestamp of the downloaded file can be set
// according to the timestamp of the remote file.
<原文结束>

# <翻译开始>
// TestRemoteTime 测试下载的文件的时间戳可以根据远程文件的时间戳进行设置。
// md5:61219270ca4080c4
# <翻译结束>


<原文开始>
// random time between epoch and now
<原文结束>

# <翻译开始>
// 在从纪元到现在的随机时间之间 md5:27113058a5d039b7
# <翻译结束>


<原文开始>
	// Assert that an existing local file will not be truncated prior to the
	// BeforeCopy hook has a chance to cancel the request
<原文结束>

# <翻译开始>
	// 确保在BeforeCopy钩子有机会取消请求之前，现有的本地文件不会被截断
	// md5:a4a1cc70630a4253
# <翻译结束>


<原文开始>
// ref: https://github.com/cavaliergopher/grab/v3/issues/37
<原文结束>

# <翻译开始>
// 参考：https://github.com/cavaliergopher/grab/v3/issues/37 md5:b6b9c9f488d6f5b8
# <翻译结束>


<原文开始>
// download new, smaller version of same file
<原文结束>

# <翻译开始>
// 下载同一文件的更小版本 md5:7f051d22caa637fd
# <翻译结束>


<原文开始>
// local file should have truncated and not resumed
<原文结束>

# <翻译开始>
// 本地文件应被截断且不应恢复 md5:7ec9bcc05ac9b7b0
# <翻译结束>


<原文开始>
// TestHeadBadStatus validates that HEAD requests that return non-200 can be
// ignored and succeed if the GET requests succeeeds.
//
// Fixes: https://github.com/cavaliergopher/grab/v3/issues/43
<原文结束>

# <翻译开始>
// TestHeadBadStatus 验证非200状态码的HEAD请求可以被忽略，如果GET请求成功，则应成功。
// 
// 修复：https://github.com/cavaliergopher/grab/v3/issues/43
// md5:40f1894de2c9714b
# <翻译结束>


<原文开始>
// TestMissingContentLength ensures that the Response.Size is correct for
// transfers where the remote server does not send a Content-Length header.
//
// TestAutoResume also covers cases with checksum validation.
//
// Kudos to Setni?ka Ji?í <Jiri.Setnicka@ysoft.com> for identifying and raising
// a solution to this issue. Ref: https://github.com/cavaliergopher/grab/v3/pull/27
<原文结束>

# <翻译开始>
// TestMissingContentLength 确保在远程服务器不发送Content-Length标头的传输中，Response.Size是正确的。
// 
// TestAutoResume 还涵盖了带有校验和验证的情况。
// 
// 感谢Setni?ka Ji?í <Jiri.Setnicka@ysoft.com> 识别并提出了这个问题的解决方案。参考：https://github.com/cavaliergopher/grab/v3/pull/27
// md5:e7af6364b11bb9a2
# <翻译结束>


<原文开始>
	// expectSize must be sufficiently large that DefaultClient.Do won't prefetch
	// the entire body and compute ContentLength before returning a Response.
<原文结束>

# <翻译开始>
	// expectSize 必须足够大，以防止DefaultClient.Do在返回Response之前预读取整个正文并计算ContentLength。
	// md5:2fa7f5fccb977406
# <翻译结束>


<原文开始>
// ensure remote server is not sending content-length header
<原文结束>

# <翻译开始>
// 确保远程服务器没有发送内容长度头 md5:acf939fcfbd160e4
# <翻译结束>


<原文开始>
// before completion, response size should be -1
<原文结束>

# <翻译开始>
// 在完成之前，响应大小应该是 -1 md5:8c0220eb52158956
# <翻译结束>


<原文开始>
// on completion, response size should be actual transfer size
<原文结束>

# <翻译开始>
// 完成后，响应大小应为实际传输的大小 md5:e0f61b6db397fce5
# <翻译结束>


<原文开始>
// ensure Response.Bytes is correct and can be reread
<原文结束>

# <翻译开始>
// 确保Response.Bytes正确并且可以重新读取 md5:f0b71aedcf56f03c
# <翻译结束>


<原文开始>
// ensure Response.Open stream is correct and can be reread
<原文结束>

# <翻译开始>
// 确保Response.Open流正确并且可以重读 md5:e8cfd6aca307387f
# <翻译结束>


<原文开始>
// Response.Filename should still be set
<原文结束>

# <翻译开始>
// Response.Filename 应该仍然被设置 md5:8687d7c41c644b55
# <翻译结束>


<原文开始>
// ensure no files were written
<原文结束>

# <翻译开始>
// 确保没有文件被写入 md5:2eb385c25204adfd
# <翻译结束>

