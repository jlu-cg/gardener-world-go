package service

//QueryTagFragments 查询标签列表
func QueryTagFragments(queryTagFragment TagFragment, lastID int) []TagFragment {
	if lastID < 0 {
		lastID = 0
	}
	tagFragments := queryTagFragments(queryTagFragment, lastID)
	return tagFragments
}

//SaveTagFragment 保存类别
func SaveTagFragment(tagFragment TagFragment) int {
	if tagFragment.ID > 0 {
		updateTagFragment(tagFragment)
	} else {
		addTagFragment(tagFragment)
	}
	return 0
}

//QueryTagFragmentByID 通过ID查询
func QueryTagFragmentByID(tagID int) TagFragment {
	if tagID > 0 {
		var queryTagFragment TagFragment
		queryTagFragment.ID = tagID
		tags := queryTagFragments(queryTagFragment, 0)
		if len(tags) > 0 {
			return tags[0]
		}
	}
	var tagFragment TagFragment
	return tagFragment
}

//DeleteTagFragmentByID 通过ID删除标签
func DeleteTagFragmentByID(id int) int {
	code := deleteTagFragmentByID(id)

	var deleteFragmentTagRelation FragmentTagRelation
	deleteFragmentTagRelation.TagID = id
	code = DeleteFragmentTagRelations(deleteFragmentTagRelation)

	var deleteTagArticleFragmentRelation TagArticleFragmentRelation
	deleteTagArticleFragmentRelation.TagFragmentID = id
	code = DeleteTagArticleFragmentRelations(deleteTagArticleFragmentRelation)

	return code
}
