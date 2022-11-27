package usecase

import (
	"backend/cmd/api/domain/model"
	"backend/cmd/api/domain/repository"
	"time"
)

type DiaryUseCase interface {
	Create(userId int, gameId int) (err error)
	EndPlaying(userId int) (err error)
}

type diaryUseCase struct {
	diaryRepository repository.DiaryRepository
	gameRepository  repository.GameRepository
}

func NewDiaryUseCase(dr repository.DiaryRepository, gr repository.GameRepository) DiaryUseCase {
	return &diaryUseCase{diaryRepository: dr, gameRepository: gr}
}

func (du diaryUseCase) Create(userId int, gameId int) (err error) {
	diary := &model.Diary{
		UserId:         userId,
		GameId:         gameId,
		StartPlayingAt: time.Now(),
		EndPlayingAt:   time.Now(),
	}

	err = du.diaryRepository.Create(diary)
	if err != nil {
		return err
	}
	return nil
}

func (du diaryUseCase) EndPlaying(userId int) (err error) {
	diary := &model.Diary{}
	diary, err = du.diaryRepository.FindWritingDiary(userId)
	if err != nil {
		return err
	}
	diary.EndPlayingAt = time.Now()

	err = du.diaryRepository.Update(diary)
	if err != nil {
		return err
	}
	return nil
}
