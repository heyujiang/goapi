package vo

type CourseVo struct {
	CourseId     int    `json:"courseId"`
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	Image        string `json:"image"`
	TeacherName  string `json:"teacherName"`
	TeacherTitle string `json:"teacherTitle"`
}

type SectionVo struct {
	SectionId   int        `json:"sectionId"`
	CourseId    int        `json:"courseId"`
	SectionName string     `json:"sectionName"`
	Sort        int        `json:"sort"`
	Description string     `json:"description"`
	Lessons     []LessonVo `json:"lessons"`
}

type LessonVo struct {
	LessonId  int    `json:"lessonId"`
	SectionId int    `json:"sectionId"`
	CourseId  int    `json:"courseId"`
	Theme     string `json:"theme"`
	Sort      int    `json:"sort"`
}

type LessonContentVo struct {
	LessonId  int    `json:"lessonId"`
	Theme     string `json:"theme"`
	SectionId int    `json:"sectionId"`
	CourseId  int    `json:"courseId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
