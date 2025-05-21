package db

import "back/model"

func (crh *commentReportService) SelectAll() ([]model.CommentReport, error) {
	var commentReports []model.CommentReport
	result := crh.Db.Find(&commentReports)
	if result.Error != nil {
		return nil, result.Error
	}
	return commentReports, nil
}

func (crh *commentReportService) SelectAllByUserIdAndCommentId(userId uint, commentId uint) ([]model.CommentReport, error) {
	var commentReports []model.CommentReport
	result := crh.Db.Where("user_id = ? AND comment_id = ?", userId, commentId).Find(&commentReports)
	if result.Error != nil {
		return nil, result.Error
	}
	return commentReports, nil
}

func (crh *commentReportService) Create(commentReport model.CommentReport) (uint, error) {
	result := crh.Db.Create(&commentReport)
	return commentReport.Id, result.Error
}