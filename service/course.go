package service

import (
	"goapi/entity/vo"
	"goapi/model/lagou"
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
