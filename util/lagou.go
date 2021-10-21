package util

import (
	"encoding/json"
	"goapi/config"
	"goapi/model/lagou"
	"goapi/pkg/client"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	headerMap        = make(map[string]string)
	courseListurl    = "https://gate.lagou.com/v1/neirong/edu/homepage/getCourseListV2?isPc=true&t=1634797429910"
	courseMenuUrl    = "https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons?courseId="
	courseContentUrl = "https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail?lessonId="
)

func init() {
	headerMap["accept"] = "application/json, text/plain, */*"
	headerMap["accept-encoding"] = "gzip, deflate, br"
	headerMap["accept-language"] = "zh-CN,zh;q=0.9"
	headerMap["authorization"] = "$$$EDU_eyJraWQiOiJiOGNjMDM0NC02NzgwLTRkNWMtYWJiYy1kNmY2ZjMwMGQxNTgiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ0Z3QiOiJfQ0FTX1RHVF9UR1QtMjI1NDAwYWU4YTU5NGIzNTk3Y2I5OGU2ZmViZGQwNjMtMjAyMTEwMjExMTEwMzEtX0NBU19UR1RfIiwic3ViIjoiODgzNzg0NyIsImlwIjoiNjAuMTkxLjU2LjcwIiwiaXNzIjoiZWR1LmxhZ291LmNvbSIsInRva2VuVHlwZSI6MSwiZXhwIjoxNjM1MzkwNjMxLCJpYXQiOjE2MzQ3ODU4MzF9.gU18vw3_Oo149s6tXjyoIF0ssXiKWHo0la9fHaIo8G7M6MWnPlhd7F58MDsHvbF2LbVTvBn9SXtF9hBzufdZC0RYSAcfS5Cb1YbWL_KHAAxNa_x6zKbOdJC_M0bmEtmi1xMpe2sgl95_a0t_ynX5_BFOs-plagOZValWYljB9m_TSmiOvwRdhFVGdjTMsvxcCXUaTFSqeSWVz6JTe9TNocvl_IOtIXUQZ9Aa897SjmF_e6P1IQ_BDB_uceU4bsDdL9gcC7pWAzehM2APZ3DhU0OH58vhOQ6xoQIVAGCf8QQYA7GvttNHzuGNNi8mBuWJXlRMsT643VWjzIE-d5GA5A"
	headerMap["cache-control"] = "no-cache"
	headerMap["cookie"] = "smidV2=20201103175011d685a5f9d671ed2d7c9dace93874b2490039189393321b8d0; LG_HAS_LOGIN=1; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1623397354; user_trace_token=20210611154239-177c69ab-2196-452f-bebd-e757e70b453d; LGUID=20210611154239-c80b8ddd-b9cd-4b02-a83e-96da53b2cdc3; _ga=GA1.2.1082295257.1623397354; EDUJSESSIONID=ABAAAECABCAAACD647144EBAB82A400133FB8CEF9415A9B; thirdDeviceIdInfo=%5B%7B%22channel%22%3A1%2C%22thirdDeviceId%22%3A%22WC39ZUyXRgdE4Qm16afmr3sIbO4rlUlytfujNPNld5cIXKj5kB4swLWvLzzq8O1En5iEsqIaeGhJ7ukzMeBKW57eCXKU60X5ntL/WmrP2Tav+DYF2YqyHqyJZY0ExL9NyO7dVu22KtO5ug38+tQRFiPbCeCS3dY9p8Ozw+p1E5euaKk0cJkZrpVxzI3vU2hPKBR3sT0Brt5HAdfRWw1vjtbnPaAHyTdeowHZphYaMYpqts7NHJdzwH6SmWcyEpK1i1487577677129%22%7D%2C%7B%22channel%22%3A2%7D%5D; sensorsdata2015session=%7B%7D; user-finger=69c80588be53fb911359d3b786117f39; gate_login_token=7c4612486affdf28d4c6e26309a01f0e4f8a62300ba7373f; LG_LOGIN_USER_ID=65f54e2d021a7fb4d67d012c4bcabc28826667a86333e9a5; _putrc=34B8E4019D91D9F5; edu_gate_login_token=$$$EDU_eyJraWQiOiJiOGNjMDM0NC02NzgwLTRkNWMtYWJiYy1kNmY2ZjMwMGQxNTgiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ0Z3QiOiJfQ0FTX1RHVF9UR1QtMjI1NDAwYWU4YTU5NGIzNTk3Y2I5OGU2ZmViZGQwNjMtMjAyMTEwMjExMTEwMzEtX0NBU19UR1RfIiwic3ViIjoiODgzNzg0NyIsImlwIjoiNjAuMTkxLjU2LjcwIiwiaXNzIjoiZWR1LmxhZ291LmNvbSIsInRva2VuVHlwZSI6MSwiZXhwIjoxNjM1MzkwNjMxLCJpYXQiOjE2MzQ3ODU4MzF9.gU18vw3_Oo149s6tXjyoIF0ssXiKWHo0la9fHaIo8G7M6MWnPlhd7F58MDsHvbF2LbVTvBn9SXtF9hBzufdZC0RYSAcfS5Cb1YbWL_KHAAxNa_x6zKbOdJC_M0bmEtmi1xMpe2sgl95_a0t_ynX5_BFOs-plagOZValWYljB9m_TSmiOvwRdhFVGdjTMsvxcCXUaTFSqeSWVz6JTe9TNocvl_IOtIXUQZ9Aa897SjmF_e6P1IQ_BDB_uceU4bsDdL9gcC7pWAzehM2APZ3DhU0OH58vhOQ6xoQIVAGCf8QQYA7GvttNHzuGNNi8mBuWJXlRMsT643VWjzIE-d5GA5A; login=true; unick=River_08; kw_login_authToken=\"iJMo1SUe/L4/XF6h9VGjJLe8drQ/HKHJrhETGtibtpedIPSsoFtMhujfIkcZJ54qL0w6hyJZTTCRFPncSxHWJTfZRI7F9c6heXBrNFaEiMt5jIyAWDmgQaAFqz9Py/hUglGRb16J5u9jgfoRM/9gNVfpNSIp16WUVeoyrHE6vVh4rucJXOpldXhUiavxhcCELWDotJ+bmNVwmAvQCptcy5e7czUcjiQC32Lco44BMYXrQ+AIOfEccJKHpj0vJ+ngq/27aqj1hWq8tEPFFjdnxMSfKgAnjbIEAX3F9CIW8BSiMHYmPBt7FDDY0CCVFICHr2dp5gQVGvhfbqg7VzvNsw==\"; X_HTTP_TOKEN=531269863769dae906958743619851ad49f0458ff6; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%228837847%22%2C%22first_id%22%3A%2217ca0d41b561fe-026dd4aa7ff0e6-b7a1438-2073600-17ca0d41b57fd0%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E8%87%AA%E7%84%B6%E6%90%9C%E7%B4%A2%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Fwww.google.com%2F%22%2C%22%24latest_utm_source%22%3A%22App%22%2C%22%24latest_utm_medium%22%3A%22%E8%AE%AD%E7%BB%83%E8%90%A5%E4%B8%93%E5%8C%BA%22%2C%22%24latest_utm_campaign%22%3A%22App%E8%AE%AD%E7%BB%83%E8%90%A5%E4%B8%93%E5%8C%BA%22%2C%22%24os%22%3A%22Windows%22%2C%22%24browser%22%3A%22Chrome%22%2C%22%24browser_version%22%3A%2294.0.4606.81%22%7D%2C%22%24device_id%22%3A%221758d839363aef-0b1f4da831fd42-5d16351b-2073600-1758d839364bde%22%7D"
	headerMap["edu-referer"] = "https://kaiwu.lagou.com/"
	headerMap["origin"] = "https://kaiwu.lagou.com"
	headerMap["pragma"] = "no-cache"
	headerMap["referer"] = "https://kaiwu.lagou.com/"
	headerMap["sec-ch-ua"] = "\"Chromium\";v=\"94\", \"Google Chrome\";v=\"94\", \";Not A Brand\";v=\"99\""
	headerMap["sec-ch-ua-mobile"] = "?0"
	headerMap["sec-ch-ua-platform"] = "\"Windows\""
	headerMap["sec-fetch-dest"] = "empty"
	headerMap["sec-fetch-mode"] = "cors"
	headerMap["sec-fetch-site"] = "same-site"
	headerMap["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36"
	headerMap["x-l-req-header"] = "{\"deviceType\":1,\"userToken\":\"$$$EDU_eyJraWQiOiJiOGNjMDM0NC02NzgwLTRkNWMtYWJiYy1kNmY2ZjMwMGQxNTgiLCJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ0Z3QiOiJfQ0FTX1RHVF9UR1QtMjI1NDAwYWU4YTU5NGIzNTk3Y2I5OGU2ZmViZGQwNjMtMjAyMTEwMjExMTEwMzEtX0NBU19UR1RfIiwic3ViIjoiODgzNzg0NyIsImlwIjoiNjAuMTkxLjU2LjcwIiwiaXNzIjoiZWR1LmxhZ291LmNvbSIsInRva2VuVHlwZSI6MSwiZXhwIjoxNjM1MzkwNjMxLCJpYXQiOjE2MzQ3ODU4MzF9.gU18vw3_Oo149s6tXjyoIF0ssXiKWHo0la9fHaIo8G7M6MWnPlhd7F58MDsHvbF2LbVTvBn9SXtF9hBzufdZC0RYSAcfS5Cb1YbWL_KHAAxNa_x6zKbOdJC_M0bmEtmi1xMpe2sgl95_a0t_ynX5_BFOs-plagOZValWYljB9m_TSmiOvwRdhFVGdjTMsvxcCXUaTFSqeSWVz6JTe9TNocvl_IOtIXUQZ9Aa897SjmF_e6P1IQ_BDB_uceU4bsDdL9gcC7pWAzehM2APZ3DhU0OH58vhOQ6xoQIVAGCf8QQYA7GvttNHzuGNNi8mBuWJXlRMsT643VWjzIE-d5GA5A\"}"
}

func GetData(url string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headerMap {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//所有专栏 strat
type ResponseStruct struct {
	State   int           `json:"state"`
	Message string        `json:"message"`
	Content ContentStruct `json:"content"`
}

type ContentStruct struct {
	ContentCardList []ContentCardStruct `json:"contentCardList"`
}

type ContentCardStruct struct {
	CourseList []CourseStruct `json:"courseList"`
}

type CourseStruct struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	Image        string `json:"image"`
	HasBuy       bool   `json:"hasBuy"`
	TeacherName  string `json:"teacherName"`
	TeacherTitle string `json:"teacherTitle"`
}

func GetAllCourse() {
	url := courseListurl
	log.Println("所有专栏列表", url)
	body, err := GetData(url)
	if err != nil {
		log.Println("获得所有专栏列表错误：", err.Error())
		return
	}

	var response ResponseStruct
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("获得所有专栏列表JSON解析错误：", err.Error())
		return
	}

	if response.State != 1 {
		log.Println("获得所有专栏列表返回异常：", response.Message)
		return
	}

	//courselen := len(response.Content.ContentCardList[4].CourseList)
	//wg := sync.WaitGroup{}
	//wg.Add(courselen)

	for _, courseInfo := range response.Content.ContentCardList[4].CourseList {
		//courseModel := lagou.CourseModel{
		//	CourseId:courseInfo.Id,
		//	Title:courseInfo.Title,
		//	Brief:courseInfo.Brief,
		//	Image:courseInfo.Image,
		//	TeacherName:courseInfo.TeacherName,
		//	TeacherTitle:courseInfo.TeacherTitle,
		//}
		//if err := courseModel.Create(); err != nil {
		//	fmt.Println(err.Error())
		//}

		GetCourseMenu(courseInfo.Id)
	}

	//wg.Wait()

}

//所有专栏 end

//专栏菜单列表 start
type MenuResponse struct {
	State   int         `json:"state"`
	Message string      `json:"message"`
	Content MenuContent `json:"content"`
}

type MenuContent struct {
	CourseName        string    `json:"courseName"`
	CourseSectionList []Section `json:"courseSectionList"`
}

type Section struct {
	Id             int       `json:"id"`
	CourseId       int       `json:"courseId"`
	SectionName    string    `json:"sectionName"`
	SectionSortNum int       `json:"sectionSortNum"`
	Description    string    `json:"description"`
	CourseLessons  []Lessons `json:"courseLessons"`
}

type Lessons struct {
	Id            int    `json:"id"`
	CourseId      int    `json:"courseId"`
	SectionId     int    `json:"sectionId"`
	Theme         string `json:"theme"`
	LessonSortNum int    `json:"lessonSortNum"`
}

func GetCourseMenu(courseId int) {
	//defer wg.Done()
	url := courseMenuUrl + strconv.Itoa(courseId)
	log.Println("获得专栏模块菜单url:", url, "；专栏编号：", courseId)
	body, err := GetData(url)
	if err != nil {
		log.Println("获得专栏模块菜单请求接口错误:", err.Error(), "；专栏编号：", courseId)
		return
	}

	var response MenuResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("获得专栏模块菜单JSON转换异常:", err.Error(), "；专栏编号：", courseId)
		return
	}

	if response.State != 1 {
		log.Println("获得专栏模块菜单请求数据错误:", response.Message, "；专栏编号：", courseId)
		return
	}

	for _, section := range response.Content.CourseSectionList {
		//创建section
		//sectionModel := lagou.SectionModel{
		//	SectionId:section.Id,
		//	CourseId:section.CourseId,
		//	SectionName:section.SectionName,
		//	Sort:section.SectionSortNum,
		//	Description:section.Description,
		//}
		//if err := sectionModel.Create(); err != nil {
		//	log.Println("获得专栏模块菜创建数据错误:",response.Message,"；",sectionModel)
		//}

		//lessonLen := len(section.CourseLessons)
		//wg := sync.WaitGroup{}
		//wg.Add(lessonLen)
		for _, lessons := range section.CourseLessons {
			//创建lesson
			//lessonModel := lagou.LessonModel{
			//	LessonId:lessons.Id,
			//	SectionId:lessons.SectionId,
			//	CourseId:lessons.CourseId,
			//	Theme:lessons.Theme,
			//	Sort:lessons.LessonSortNum,
			//}
			//if err := lessonModel.Create(); err != nil {
			//	log.Println("获得专栏模块菜创建数据错误:",response.Message,"；",lessonModel)
			//}

			GetCourseContent(lessons.Id)

		}
		//wg.Wait()
	}
}

//专栏菜单列表 end

//专栏内容 start

type ContentResponse struct {
	State   int            `json:"state"`
	Message string         `json:"message"`
	Content ContentContent `json:"content"`
}

type ContentContent struct {
	Id          int    `json:"id"`
	CourseId    int    `json:"courseId"`
	SectionId   int    `json:"sectionId"`
	Theme       string `json:"theme"`
	TextContent string `json:"textContent"`
}

func GetCourseContent(lessonsId int /*,wg *sync.WaitGroup*/) {
	//defer wg.Done()

	url := courseContentUrl + strconv.Itoa(lessonsId)
	log.Println("获得文章信息接口错误URL：", url)
	body, err := GetData(url)
	if err != nil {
		log.Println("获得文章信息请求接口错误：", err.Error(), "；文章id：", lessonsId)
		return
	}

	var response ContentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("获得文章信息JSON转换异常：", err.Error(), "；文章id：", lessonsId)
		return
	}

	if response.State != 1 {
		log.Println("获得文章信息请求数据错误：", response.Message, "；文章id：", lessonsId)
		return
	}

	//创建content
	contentModel := lagou.ContentModel{
		LessonId: response.Content.Id,
		Content:  response.Content.TextContent,
	}
	if err := contentModel.Create(); err != nil {
		log.Println("获得文章信息创建数据失败：", err.Error(), "；文章id：", lessonsId, contentModel)
		return
	}

	return
}

//专栏内容

func Run() {
	//init config
	if err := config.Init(""); err != nil {
		log.Fatal(err.Error())
	}

	//init db
	client.MySqlClients.Init()
	defer client.MySqlClients.Close()

	GetAllCourse()
}
