package analysis

import (
	"testing"
	"fmt"
	"regexp"
)



func TestAnalysis(t *testing.T) {

	content := `2012.11 - 2013.01 上海博之光医疗科技医疗有限公司 （2个月）
	技术部 | 高级软件工程师 | 4001-6000元/月
			医药/生物工程 | 企业性质:民营 | 规模:20人以下
工作描述:	负责产品的软件研发，承担软件设计、代码编写、软件测试方案制订、代码调试和测试等.
	2011.08 至 2012.09 上海伍翼信息技术有限公司 （1年1个月）
	技术部 | 软件工程师 | 2001-4000元/月
			互联网/电子商务 | 企业性质:民营 | 规模:20人以下
工作描述:	ASP.NET程序员(前台页面与后台页面的编写与维护工作
	2011.08 - 当前 上海伍翼信息技术有限公司 （1年1个月）
	技术部 | 软件工程师 | 2001-4000元/月
			互联网/电子商务 | 企业性质:民营 | 规模:20人以下
工作描述:	ASP.NET程序员(前台页面与后台页面的编写与维护工作`

	fmt.Println(content)

	for _, exp := range timeRegex {
		reg, err := regexp.Compile(exp)
		if err != nil {
			fmt.Println("错误:", err)
		}else {
			all := reg.FindAllString(content, -1)
			fmt.Println(len(all))
			fmt.Println(all)
		}
	}
	t.Log("ok")
}
