package service

import "github.com/gardener/gardener-world-go/config"

//FragmentIntroductionRelation 碎片和详细介绍关联
type FragmentIntroductionRelation struct {
	ID                   int `json:"id"`
	FragmentID           int `json:"fragmentId"`
	DetailIntroductionID int `json:"detailIntroductionId"`
}

//FragmentIntroductionRelationWithIntroduction 碎片和详细介绍关联包含详细介绍
type FragmentIntroductionRelationWithIntroduction struct {
	FragmentIntroductionRelation
	Summary string `json:"summary"`
}

const (
	queryFragmentIntroductionRelationWithIntroductionsSQL = "select a.id, a.fragment_id, a.detail_introduction_id, b.summary from fragment_introduction_relation a inner join detail_introduction b on a.detail_introduction_id=b.id "
	addFragmentIntroductionRelationSQL                    = "insert into fragment_introduction_relation(fragment_id, detail_introduction_id)values($1, $2)"
	deleteFragmentIntroductionRelationsSQL                = "delete from fragment_introduction_relation "
)

func queryFragmentIntroductionRelationWithIntroductions(relation FragmentIntroductionRelationWithIntroduction, lastID int) []FragmentIntroductionRelationWithIntroduction {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if relation.ID > 0 {
		whereSQL += " and a.id=" + intToSafeString(relation.ID)
	}

	if relation.FragmentID > 0 {
		whereSQL += " and a.fragment_id=" + intToSafeString(relation.FragmentID)
	}

	if relation.DetailIntroductionID > 0 {
		whereSQL += " and a.detail_introduction_id=" + intToSafeString(relation.DetailIntroductionID)
	}

	if lastID >= 0 {
		whereSQL += " and a.id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryFragmentIntroductionRelationWithIntroductionsSQL + whereSQL)
	defer rows.Close()
	var relations []FragmentIntroductionRelationWithIntroduction
	if rows == nil {
		return relations
	}
	if err != nil {
		panic(err)
	}

	var temp FragmentIntroductionRelationWithIntroduction
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.FragmentID, &temp.DetailIntroductionID, &temp.Summary)
		relations = append(relations, temp)
	}

	return relations
}

func addFragmentIntroductionRelation(relation FragmentIntroductionRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addFragmentIntroductionRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(relation.FragmentID, relation.DetailIntroductionID)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteFragmentIntroductionRelations(relation FragmentIntroductionRelation) int {

	hasCondition := false
	whereSQL := " where 1=1 "
	if relation.ID > 0 {
		hasCondition = true
		whereSQL += " and id=" + intToSafeString(relation.ID)
	}
	if relation.FragmentID > 0 {
		hasCondition = true
		whereSQL += " and fragment_id=" + intToSafeString(relation.FragmentID)
	}
	if relation.DetailIntroductionID > 0 {
		hasCondition = true
		whereSQL += " and detail_introduction_id=" + intToSafeString(relation.DetailIntroductionID)
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteFragmentIntroductionRelationsSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
