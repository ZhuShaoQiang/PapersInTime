// 定义使用的类型
package types

// Paper represents a paper record
type Paper struct {
	ID     string `json:"citeName" gorm:"primaryKey"` // 引用名字，就是引用里面的唯一名字
	Title  string `json:"title"`                      // 标题
	Author string `json:"author"`                     // 作者
	Year   string `json:"year"`                       // 发表年份
	Path   string `json:"path"`                       // 服务器存储路径
	Labels string `json:"labels"`                     // 标签, 多个用,隔开
	Cite   string `json:"cite"`                       // 引用了谁，就是其他人的引用名字，多个用,隔开
	Cited  string `json:"cited"`                      // 谁引用了这个，谁就是其中之一，多个用,隔开，这个非必填
}

// 查看详情的接收结构
type PaperDetailFunc struct {
	CiteName string `json:"citeName"`
}

// 从一个basepaper查找引用和非引用的结构，每一个内容就是一个paper
type BasepaperCitesAndCited struct {
	ID     string  `json:"citeName" gorm:"primaryKey"` // 引用名字，就是引用里面的唯一名字
	Title  string  `json:"title"`                      // 标题
	Author string  `json:"author"`                     // 作者
	Year   string  `json:"year"`                       // 发表年份
	Path   string  `json:"path"`                       // 服务器存储路径
	Labels string  `json:"labels"`                     // 标签, 多个用,隔开
	Cite   []Paper `json:"cite"`                       // 引用了谁，就是其他人的引用名字，多个用,隔开
	Cited  []Paper `json:"cited"`                      // 谁引用了这个，谁就是其中之一，多个用,隔开，这个非必填

}
