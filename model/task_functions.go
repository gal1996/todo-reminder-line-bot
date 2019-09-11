package model

import (
	"time"
	"todo-reminder-line-bot/utils"
)

func (task *Task) Create(id int, content string, deadLine time.Time, userId int) (*Task, error) {
	// ユーザーIDは空であってはいけない
	if id == 0 {
		err := utils.MyError{Message: "#create Task id is empty"}
		return &Task{}, err
	} else {
		task.Id = id
	}

	// contentは空であってはいけない
	if content == "" {
		err := utils.MyError{Message: "#create Task content is empty"}
		return &Task{}, err
	} else {
		task.Content = content
	}

	// deadLineは当日より前の日付であってはならない
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().UTC().In(jst)
	// deadLineが今より後でなければエラー
	if !deadLine.After(now) {
		err := utils.MyError{Message: "#create Task dead line is before now"}
		return &Task{}, err
	} else {
		task.DeadLine = deadLine
	}

	// ユーザーIDは空であってはいけない
	if userId == 0 {
		err := utils.MyError{Message: "#create Task user id is empty"}
		return &Task{}, err
	} else {
		task.UserId = userId
	}

	return task, nil
}

func (task *Task) DoneTask() (*Task, error) {
	// 完了済みだったら何もしない
	if task.IsDone {
		err := utils.MyError{Message: "#doneTask Task is already done"}
		return task, err
	} else {
		task.IsDone = true
		return task, nil
	}
}

func (task *Task) ResetDoneTask() (*Task, error) {
	// 未完の状態からは実行しない
	if !task.IsDone {
		err := utils.MyError{Message: "#resetDoneTask Task is not already done"}
		return &Task{}, err
	} else {
		task.IsDone = false
		return task, nil
	}
}

func (task *Task) Delete() (*Task, error) {
	// todo これどうやって実装するのが正解や？からのタスクを返すんか？
	// todo とりあえず空にして返すのを正解にする
	return &Task{}, nil
}

func (task *Task) ChangeDeadLine(newDeadLine time.Time) (*Task, error) {
	// 新しいdead lineaは今よりあとでなければならない
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().UTC().In(jst)
	// deadLineが今より後でなければエラー
	if !newDeadLine.After(now) {
		err := utils.MyError{Message: "#changeDeadLine Task dead line is before now"}
		return &Task{}, err
	} else {
		task.DeadLine = newDeadLine
	}

	return task, nil
}


func (task *Task) ChangeContent(newContent string) (*Task, error) {
	// contentは空ではならない
	if newContent == "" {
		err := utils.MyError{Message: "#changeContent new content is empty"}
		return &Task{}, err
	} else {
		task.Content = newContent
		return task, nil
	}
}
