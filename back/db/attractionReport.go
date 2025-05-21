package db

import "back/model"

func (arh *AttractionReportService) SelectAll() ([]model.AttractionReport, error) {
	var attractionReports []model.AttractionReport
	result := arh.Db.Find(&attractionReports)
	if result.Error != nil {
		return nil, result.Error
	}
	return attractionReports, nil
}

func (arh *AttractionReportService) SelectAllByUserIdAndAttractionId(userId uint, attractionId uint) ([]model.AttractionReport, error) {
	var attractionReports []model.AttractionReport
	result := arh.Db.Where("user_id = ? AND attraction_id = ?", userId, attractionId).Find(&attractionReports)
	if result.Error != nil {
		return nil, result.Error
	}
	return attractionReports, nil
}

func (arh *AttractionReportService) Create(attractionReport model.AttractionReport) (uint, error) {
	result := arh.Db.Create(&attractionReport)
	return attractionReport.Id, result.Error
}