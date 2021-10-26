package controller

import (
	"github.com/gin-gonic/gin"
	"goapi/service"
	"strconv"
)

/**
获得课程列表
*/
func CourseList(ctx *gin.Context) {
	infos, err := service.GetCourseList()
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, infos)
}

func Lessons(ctx *gin.Context) {
	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	infos, err := service.GetLessonsList(courseId)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, infos)
}

func LessonDetail(ctx *gin.Context) {
	lessonId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	info, err := service.GetLessonsDetail(lessonId)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, info)
}

func CreatePdf(ctx *gin.Context) {
	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	str, err := service.CreatePdf(courseId)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, str)
}

func CreatePdfAll(ctx *gin.Context) {
	str, err := service.CreatePdfAll()
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, str)
}
