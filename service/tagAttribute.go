package service

//TagAttribute 标签比较属性
type TagAttribute struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	TagID int    `json:"tagId"`
}

const (
	queryTagAttributesSQL = "select id, name, tag_id from tag_attribute "
	addTagAttributeSQL    = "insert into tag_attribute(name, tag_id) values($1, $2)"
	updateTagAttributeSQL = "update tag_attribute name=$1, tag_id=$2 where id=$3"
	deleteTagAttributeSQL = "delete from tag_attribute where id=$1"
)

func queryTagAttributes(attribute TagAttribute) []TagAttribute {
	var tagAttributes []TagAttribute
	whereSQL := " where 1=1 "
	if attribute.TagID > 0 {
		whereSQL += " and tag_id=" + intToSafeString(attribute.TagID)
	}

	if attribute.Name != "" {
		whereSQL += " and name like '" + strToSafeString(attribute.Name) + "%' "
	}

	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryTagAttributesSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return tagAttributes
	}
	if err != nil {
		panic(err)
	}

	var temp TagAttribute
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.TagID)
		tagAttributes = append(tagAttributes, temp)
	}

	return tagAttributes
}

func addTagAttribute(attribute TagAttribute) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagAttributeSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(attribute.Name, attribute.TagID)
	if err != nil {
		return -1

	}
	return 0
}

func updateTagAttribute(attribute TagAttribute) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateTagAttributeSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(attribute.Name, attribute.TagID, attribute.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteTagAttributeByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagAttributeSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
