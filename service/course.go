package service

import (
	"bytes"
	"goapi/entity/vo"
	"goapi/model/lagou"
	"goapi/util"
	"os"
)

func GetCourseList() ([]vo.CourseVo, error) {
	courselist, err := lagou.ListCourse()
	if err != nil {
		return nil, err
	}

	infos := make([]vo.CourseVo, len(courselist))
	for k, courseInfo := range courselist {
		infos[k] = vo.CourseVo{
			CourseId:     courseInfo.CourseId,
			Title:        courseInfo.Title,
			Brief:        courseInfo.Brief,
			Image:        courseInfo.Image,
			TeacherName:  courseInfo.TeacherName,
			TeacherTitle: courseInfo.TeacherTitle,
		}
	}

	return infos, nil
}

func GetLessonsList(courseId int) ([]vo.SectionVo, error) {
	sectionList, err := lagou.ListSection(courseId)
	if err != nil {
		return nil, err
	}

	courseIds := []int{}
	for _, section := range sectionList {
		courseIds = append(courseIds, section.SectionId)
	}

	//获得所有lesson
	lessonList, err := lagou.ListLesson(courseIds)
	lessonMap := make(map[int][]vo.LessonVo)

	for _, lesson := range lessonList {
		lessonMap[lesson.SectionId] = append(lessonMap[lesson.SectionId], vo.LessonVo{
			LessonId:  lesson.LessonId,
			SectionId: lesson.SectionId,
			CourseId:  lesson.CourseId,
			Theme:     lesson.Theme,
			Sort:      lesson.Sort,
		})
	}

	infos := make([]vo.SectionVo, len(sectionList))
	for k, section := range sectionList {
		infos[k] = vo.SectionVo{
			SectionId:   section.SectionId,
			CourseId:    section.CourseId,
			SectionName: section.SectionName,
			Sort:        section.Sort,
			Description: section.Description,
			Lessons:     lessonMap[section.SectionId],
		}
	}

	return infos, nil
}

func GetLessonsDetail(lessonId int) (*vo.LessonContentVo, error) {

	contentInfo, err := lagou.GetContent(lessonId)
	if err != nil {
		return nil, err
	}

	lessonInfo, err := lagou.GetLesson(lessonId)
	if err != nil {
		return nil, err
	}

	courseInfo, err := lagou.GetCourse(lessonInfo.CourseId)
	if err != nil {
		return nil, err
	}

	info := &vo.LessonContentVo{
		LessonId:  contentInfo.LessonId,
		Theme:     lessonInfo.Theme,
		SectionId: lessonInfo.SectionId,
		CourseId:  lessonInfo.CourseId,
		Title:     courseInfo.Title,
		Content:   contentInfo.Content,
	}

	return info, nil
}

func CreatePdf(courseId int) (string, error) {
	//文章信息
	course, err := lagou.GetCourse(courseId)
	if err != nil {
		return "", err
	}

	LessonList, err := GetLessonsList(course.CourseId)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>` + course.Title + `</title>
</head>
<style>
    .c{
        width: 90%;
        margin: auto;
    }
    .content_div{
        width: 100%;
        box-sizing: border-box;
        padding: 20px;
    }
    .content_div img {
        width: 100%;
    }
    .content_div p {
        margin-bottom: 20px;
        line-height: 20px;
    }
    .c h1{
        margin-top: 80px;
    }
    .c h2 {
        margin-top:  50px;
    }
    .c h3 {
        margin-top: 30px;
    }
    .c h4 {
        margin-top: 20px;
    }
    .c h1,h2,h3,h4{
        margin-bottom: 20px;
    }
</style>
<body>
<div class="c">
`)

	for _, sectionInfo := range LessonList {
		//菜单
		buffer.WriteString("<h1>")
		buffer.WriteString(sectionInfo.SectionName)
		buffer.WriteString("</h1>")
		buffer.WriteString("<hr>")

		for _, lessonInfo := range sectionInfo.Lessons {
			buffer.WriteString("<h2>")
			buffer.WriteString(lessonInfo.Theme)
			buffer.WriteString("</h2>")

			lessonContent, err := GetLessonsDetail(lessonInfo.LessonId)
			if err != nil {
				return "", err
			}

			buffer.WriteString(`<div class="content_div">`)
			buffer.WriteString(lessonContent.Content)
			buffer.WriteString("</div>")
		}
	}
	buffer.WriteString(`
</div>

</body>
</html>
`)

	htmlFile := "D:/develop/golang/goapi/course/" + course.Title + ".html"
	pdfFile := "D:/develop/golang/goapi/course/pdf/" + course.Title + ".pdf"

	file1, err := os.Create(htmlFile)
	defer file1.Close()
	if err != nil {
		return "", err
	}

	_, err = file1.Write(buffer.Bytes())
	if err != nil {
		return "", err
	}

	util.Pdf(htmlFile, pdfFile)

	return course.Title, nil
}

func CreatePdfAll() (string, error) {
	courseList, err := GetCourseList()
	if err != nil {
		return "", err
	}

	//wg := sync.WaitGroup{}
	for _, courseInfo := range courseList {
		//wg.Add(1)
		//go CreatePdf(courseInfo.CourseId,&wg)
		_, err := CreatePdf(courseInfo.CourseId)
		if err != nil {
			return "", err
		}
	}
	//wg.Wait()

	return "SUCCESS", nil
}
