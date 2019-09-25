package service

//Tag 标签
type Tag struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	TagType int    `json:"tagType"`
}

const (
	queryTagsSQL     = "select id, name, tag_type from tag "
	addTagSQL        = "insert into tag(name, tag_type)values($1, $2)"
	updateTagSQL     = "update tag set name=$1, tag_type=$2 where id=$3"
	deleteTagByIDSQL = "delete from tag where id=$1"
)

//QueryTags 查询类别
func queryTags(tag Tag, lastID int) []Tag {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if tag.Name != "" {
		whereSQL += " and name like '" + strToSafeString(tag.Name) + "%' "
	}
	if tag.ID > 0 {
		whereSQL += " and id=" + intToSafeString(tag.ID)
	}
	if tag.TagType > 0 {
		whereSQL += " and tag_type=" + intToSafeString(tag.ID)
	}
	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
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
		rows.Scan(&temp.ID, &temp.Name, &temp.TagType)
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
	_, err = stmt.Exec(tag.Name, tag.TagType)
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
	_, err = stmt.Exec(tag.Name, tag.TagType)
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
