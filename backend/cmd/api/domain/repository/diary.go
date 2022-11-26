package repository

import "backend/cmd/api/domain/model"

type DiaryRepository interface {
	FindAll() (diaries []*model.Diary, err error)
	FindWritingDiary(userId int) (diary *model.Diary, err error)
	Create(diary *model.Diary) (err error)
	Update(diary *model.Diary) (err error)
}
