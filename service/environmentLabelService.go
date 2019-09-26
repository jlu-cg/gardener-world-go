package service

//QueryEnvironmentLabels 查询环境标签
func QueryEnvironmentLabels(label EnvironmentLabel, lastID int) []EnvironmentLabel {
	if lastID < 0 {
		lastID = 0
	}
	return queryEnvironmentLabels(label, lastID)
}

//QueryEnvironmentLabels 查询环境标签
func QueryEnvironmentLabelByID(environmentLabelID int) EnvironmentLabel {
	var queryEnvironmentLabel EnvironmentLabel
	queryEnvironmentLabel.ID = environmentLabelID
	var result EnvironmentLabel
	labels := QueryEnvironmentLabels(queryEnvironmentLabel, 0)
	if len(labels) > 0 {
		result = labels[0]
	}
	return result
}

//SaveEnvironmentLabel 保存环境标签
func SaveEnvironmentLabel(label EnvironmentLabel) int {
	if label.ID > 0 {
		return updateEnvironmentLabel(label)
	}
	return addEnvironmentLabel(label)
}

//DeleteEnvironmentLabel 删除环境标签
func DeleteEnvironmentLabels(label EnvironmentLabel) int {
	return deleteEnvironmentLabels(label)
}
