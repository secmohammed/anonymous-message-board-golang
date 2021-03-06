package services

import (
    "errors"
    "strconv"
    "time"

    "github.com/secmohammed/anonymous-message-board-golang/database"
    "github.com/secmohammed/anonymous-message-board-golang/models"
    "github.com/secmohammed/anonymous-message-board-golang/utils"
    "gorm.io/gorm"
)

type ReplyService interface {
    ListByThreadID(tid string, page int) (error, *[]models.Reply)
    GetByID(id string) (error, *models.Reply)
    Create(tid string, r models.Reply) (error, *models.Reply)
    DeleteWithPassword(id, password string) error
    Report(id string) error
    Delete(id string) error
}
type replyService struct {
    db *gorm.DB
}

func NewReplyService(conn database.DatabaseConnection) ReplyService {
    return &replyService{db: conn.Get()}
}
func (rs *replyService) ListByThreadID(tid string, page int) (error, *[]models.Reply) {
    var r []models.Reply
    offset := 0
    if page > 0 {
        offset = page - 1
    }
    result := rs.db.Where("thread_id = ?", tid).Limit(10).Offset(offset).Find(&r)
    return result.Error, &r
}
func (rs *replyService) GetByID(id string) (error, *models.Reply) {
    var r models.Reply
    result := rs.db.Where("id = ? ", id).First(&r)
    return result.Error, &r
}
func (rs *replyService) Create(tid string, r models.Reply) (error, *models.Reply) {
    password, err := utils.HashPassword(r.DeletePassword)
    if err != nil {
        return err, nil
    }
    threadID, err := strconv.Atoi(tid)
    if err != nil {
        return err, nil
    }
    r.DeletePassword = password
    r.ThreadID = uint(threadID)
    transaction := rs.db.Transaction(func(tx *gorm.DB) error {
        var t models.Thread
        if result := tx.Where("id = ?", tid).First(&t); result.Error != nil {
            return result.Error
        }
        if result := tx.Create(&r); result.Error != nil {
            return result.Error
        }
        if result := tx.Model(&t).Where("id = ?", tid).Update("bumped_on", time.Now()); result.Error != nil {
            return result.Error
        }
        return nil
    })
    return transaction, &r
}
func (rs *replyService) DeleteWithPassword(id, password string) error {
    return rs.db.Transaction(func(tx *gorm.DB) error {
        var r models.Reply
        if result := tx.Where("id = ? ", id).First(&r); result.Error != nil {
            return result.Error
        }
        if !utils.CheckPassword(password, r.DeletePassword) {
            return errors.New("Incorrect Password")
        }
        if result := tx.Model(&r).Where("id = ?", id).Update("text", "[deleted]"); result.Error != nil {
            return result.Error
        }
        return nil
    })
}
func (rs *replyService) Delete(id string) error {
    return rs.db.Transaction(func(tx *gorm.DB) error {
        var r models.Reply
        if result := tx.Where("id = ? ", id).First(&r); result.Error != nil {
            return result.Error
        }
        if result := tx.Model(&r).Where("id = ?", id).Update("text", "[deleted]"); result.Error != nil {
            return result.Error
        }
        return nil
    })
}

func (rs *replyService) Report(id string) error {
    return rs.db.Transaction(func(tx *gorm.DB) error {
        var r models.Reply
        if result := tx.Where("id = ? ", id).First(&r); result.Error != nil {
            return result.Error
        }
        if result := tx.Model(&r).Where("id = ?", id).Update("reported", true); result.Error != nil {
            return result.Error
        }
        return nil
    })
}
