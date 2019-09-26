package service

type IntroductionEnvironmentRelation struct {
	ID                   int `json:"id"`
	DetailIntroductionID int `json:"detailIntroductionId"`
	EnvironmentLabelID   int `json:"environmentLabelId"`
}

type IntroductionEnvironmentRelationDetail struct {
	IntroductionEnvironmentRelation
	Name    string `json:"name"`
	Version string `json:"version"`
}

const (
	queryIntroductionEnvironmentRelationsSQL = "select a.id, a.detail_introduction_id, a.environment_label_id, b.name, b.version from introduction_environment_relation a inner join environment_label b on a.tag_id=b.id "
	addIntroductionEnvironmentRelationSQL    = "insert into introduction_environment_relation(detail_introduction_id, environment_label_id) values($1, $2)"
	deleteIntroductionEnvironmentRelationSQL = "delete from introduction_environment_relation "
)

func queryIntroductionEnvironmentRelations(relationDetail IntroductionEnvironmentRelationDetail) []IntroductionEnvironmentRelationDetail {
	hasCondition := false
	whereSQL := " where 1=1 "

	if relationDetail.DetailIntroductionID > 0 {
		whereSQL += " and a.detail_introduction_id=" + intToSafeString(relationDetail.DetailIntroductionID)
		hasCondition = true
	}

	if relationDetail.EnvironmentLabelID > 0 {
		whereSQL += " and a.environment_label_id=" + intToSafeString(relationDetail.EnvironmentLabelID)
		hasCondition = true
	}

	var relationDetails []IntroductionEnvironmentRelationDetail
	if !hasCondition {
		return relationDetails
	}

	whereSQL += " order by b.id asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryIntroductionEnvironmentRelationsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return relationDetails
	}
	if err != nil {
		panic(err)
	}

	var temp IntroductionEnvironmentRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.DetailIntroductionID, &temp.EnvironmentLabelID, &temp.Name, &temp.Version)
		relationDetails = append(relationDetails, temp)
	}

	return relationDetails
}

func addIntroductionEnvironmentRelation(relation IntroductionEnvironmentRelation) int {

	if relation.DetailIntroductionID <= 0 || relation.EnvironmentLabelID <= 0 {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addIntroductionEnvironmentRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.DetailIntroductionID, relation.EnvironmentLabelID)
	if err != nil {
		return -1

	}
	return 0
}

func deleteIntroductionEnvironmentRelations(relation IntroductionEnvironmentRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}
	if relation.DetailIntroductionID > 0 {
		whereSQL += " and detail_introduction_id=" + intToSafeString(relation.DetailIntroductionID)
		hasCondition = true
	}

	if relation.EnvironmentLabelID > 0 {
		whereSQL += " and environment_label_id=" + intToSafeString(relation.EnvironmentLabelID)
		hasCondition = true
	}

	if !hasCondition {
		return -1
	}

	connection := connect()
	defer release(connection)
	stmt, err := connection.Prepare(deleteIntroductionEnvironmentRelationSQL + whereSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}
	return 0
}
