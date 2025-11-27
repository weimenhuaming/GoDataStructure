package pool

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

// 定义一个结构体表示要返回的数据
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// 创建一个 Pool 用于缓存 bytes.Buffer 对象
var bufferPool = sync.Pool{
	New: func() interface{} {
		// 当 Pool 中没有可用对象时，调用 New 函数生成新对象
		return new(bytes.Buffer)
	},
}

// 处理 HTTP 请求的函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 从 Pool 中获取一个 bytes.Buffer（如果 Pool 为空，会自动调用 New 函数创建）
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf) // 处理完成后将对象放回 Pool

	// 重置 Buffer，避免复用时的残留数据（关键步骤！）
	buf.Reset()

	// 准备要返回的数据
	resp := Response{Message: "success", Code: 200}

	// 使用 Buffer 和 JSON 编码器
	encoder := json.NewEncoder(buf)
	encoder.Encode(resp) // 将数据编码到 Buffer

	// 将 Buffer 内容写入 HTTP 响应
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
