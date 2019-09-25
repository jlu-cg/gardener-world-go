package service

//QueryTags 查询标签列表
func QueryTags(queryTag Tag, lastID int) []Tag {
	if lastID < 0 {
		lastID = 0
	}
	tags := queryTags(queryTag, lastID)
	return tags
}

//SaveTag 保存类别
func SaveTag(tag Tag) int {
	if tag.ID > 0 {
		updateTag(tag)
	} else {
		addTag(tag)
	}
	return 0
}

//QueryTagByID 通过ID查询
func QueryTagByID(tagID int) Tag {
	if tagID > 0 {
		var queryTag Tag
		queryTag.ID = tagID
		tags := queryTags(queryTag, 0)
		if len(tags) > 0 {
			return tags[0]
		}
	}
	var tag Tag
	return tag
}

//DeleteTagByID 通过ID删除标签
func DeleteTagByID(id int) int {
	code := deleteTagByID(id)

	var deleteArticleTagRelation ArticleTagRelation
	deleteArticleTagRelation.TagID = id
	code = DeleteArticleTagRelations(deleteArticleTagRelation)

	var deleteFragmentTagRelation FragmentTagRelation
	deleteFragmentTagRelation.TagID = id
	code = DeleteFragmentTagRelations(deleteFragmentTagRelation)

	var deleteTagTagRelation TagTagRelation
	deleteTagTagRelation.TagID = id
	code = DeleteFragmentTagRelations(deleteFragmentTagRelation)

	deleteTagTagRelation.TagID = -1
	deleteTagTagRelation.RelateTagID = id
	code = DeleteFragmentTagRelations(deleteFragmentTagRelation)

	return code
}
