package service

import "fmt"

type EnvironmentLabel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

const (
	queryEnvironmentLabelsSQL = "select id, name, version from environment_label "
	addEnvironmentLabelSQL    = "insert into environment_label(name, version)values($1, $2)"
	updateEnvironmentLabelSQL = "update environment_label set %s where id=$1"
	deleteEnvironmentLabelSQL = "delete from environment_label "
)

func queryEnvironmentLabels(label EnvironmentLabel, lastID int) []EnvironmentLabel {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "
	if label.ID > 0 {
		whereSQL += " and id=" + intToSafeString(label.ID)
	}

	if label.Name != "" {
		whereSQL += " and name like '" + strToSafeString(label.Name) + "%' "
	}

	if label.Version != "" {
		whereSQL += " and version like '" + strToSafeString(label.Version) + "%' "
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}

	rows, err := connection.Query(queryEnvironmentLabelsSQL + whereSQL)
	var environmentLabels []EnvironmentLabel
	if rows == nil {
		return environmentLabels
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp EnvironmentLabel
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.Version)
		environmentLabels = append(environmentLabels, temp)
	}

	return environmentLabels
}

func addEnvironmentLabel(label EnvironmentLabel) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addEnvironmentLabelSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(label.Name, label.Version)
	if err != nil {
		return -1
	}
	return 0
}

func updateEnvironmentLabel(label EnvironmentLabel) int {

	hasUpdate := false

	updateFieldSQL := ""

	if label.Name != "" {
		hasUpdate = true
		updateFieldSQL += " name='" + strToSafeString(label.Name) + "' "
	}

	if label.Version != "" {
		hasUpdate = true
		if updateFieldSQL != "" {
			updateFieldSQL += ","
		}
		updateFieldSQL += " version='" + strToSafeString(label.Version) + "' "
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(fmt.Sprintf(updateEnvironmentLabelSQL, updateFieldSQL))
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(label.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteEnvironmentLabels(label EnvironmentLabel) int {

	hasUpdate := false
	whereSQL := " where 1=1 "
	if label.ID > 0 {
		hasUpdate = true
		whereSQL += " and id=" + intToSafeString(label.ID)
	}

	if label.Name != "" {
		hasUpdate = true
		whereSQL += " and name = '" + strToSafeString(label.Name) + "' "
	}

	if label.Version != "" {
		hasUpdate = true
		whereSQL += " and version ='" + strToSafeString(label.Version) + "' "
	}

	if !hasUpdate {
		return -1
	}

	connection := connect()
	defer release(connection)
	stmt, err := connection.Prepare(deleteEnvironmentLabelSQL + whereSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}

	return 0
}
