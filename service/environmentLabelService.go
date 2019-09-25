package service

//QueryEnvironmentLabels 查询环境标签
func QueryEnvironmentLabels(label EnvironmentLabel, lastID int) []EnvironmentLabel {
	if lastID < 0 {
		lastID = 0
	}
	return queryEnvironmentLabels(label, lastID)
}

//SaveEnvironmentLabel 保存环境标签
func SaveEnvironmentLabel(label EnvironmentLabel) int {
	if label.ID > 0 {
		return updateEnvironmentLabel(label)
	}
	return addEnvironmentLabel(label)
}

//DeleteEnvironmentLabel 删除环境标签
func DeleteEnvironmentLabel(label EnvironmentLabel) int {
	return deleteEnvironmentLabel(label)
}
