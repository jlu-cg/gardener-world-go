package service

//Tag 标签
type Tag struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentId"`
	Path     string `json:"path"`
	NamePath string `json:"namePath"`
	Summary  string `json:"summary"`
}

const (
	queryTagsSQL     = "select id, name, parent_id, path, name_path, summary from tag "
	addTagSQL        = "insert into tag(name, parent_id, path, name_path, summary)values($1, $2, $3, $4, $5)"
	updateTagSQL     = "update tag set name=$1, parent_id=$2, path=$3, name_path=$4, summary=$5 where id=$6"
	queryTagByIDSQL  = "select id, name, parent_id,path, name_path, summary from tag where id=$1"
	deleteTagByIDSQL = "delete from tag where id=$1"
)

//QueryTags 查询类别
func queryTags(tag Tag, lastID int) []Tag {

	connection := connect()
	defer release(connection)

	whereSQL := " where "

	if tag.Name != "" {
		whereSQL += " name like '" + strToSafeString(tag.Name) + "%' and "
	}
	if tag.ID > 0 {
		whereSQL += " id=" + intToSafeString(tag.ID) + " and "
	}
	if tag.ParentID >= 0 {
		whereSQL += " parent_id=" + intToSafeString(tag.ParentID) + " and "
	}
	whereSQL += " id>" + intToSafeString(lastID) + " limit 20 "
	rows, err := connection.Query(queryTagsSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var tags []Tag
	if rows == nil {
		return tags
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp Tag
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.ParentID, &temp.Path, &temp.NamePath, &temp.Summary)
		tags = append(tags, temp)
	}

	return tags
}

//保存标签
func addTag(tag Tag) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(tag.Name, tag.ParentID, tag.Path, tag.NamePath, tag.Summary)
	if err != nil {
		return -1
	}
	return 0
}

//更新标签
func updateTag(tag Tag) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateTagSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(tag.Name, tag.ParentID, tag.Path, tag.NamePath, tag.Summary, tag.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteTagByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagByIDSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
