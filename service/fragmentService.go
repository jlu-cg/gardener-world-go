package service

//QueryFragments 查询碎片
func QueryFragments(fragment FragmentWithTag, lastID int) []FragmentWithTag {
	if fragment.TagFragmentID > 0 {
		return queryFragmentWithTags(fragment, lastID)
	}
	fragmentWithTags := queryFragments(fragment.Fragment, lastID)
	return fragmentWithTags
}

//SaveFragment 保存碎片
func SaveFragment(fragment Fragment) int {
	if fragment.ID > 0 {
		updateFragment(fragment)
	} else {
		addFragment(fragment)
	}
	return 0
}

//QueryFragmentByID 查询碎片
func QueryFragmentByID(fragmentID int) FragmentWithTag {
	if fragmentID > 0 {
		var queryFragment Fragment
		queryFragment.ID = fragmentID
		fragmentWithTags := queryFragments(queryFragment, 0)
		if len(fragmentWithTags) > 0 {
			return fragmentWithTags[0]
		}
	}
	var fragmentWithTag FragmentWithTag
	return fragmentWithTag
}

//DeleteFragmentByID 删除碎片
func DeleteFragmentByID(fragmentID int) int {
	if fragmentID <= 0 {
		return 0
	}
	code := deleteFragmentByID(fragmentID)

	//删除碎片和标签对应关系
	var deleteFragmentTagRelation FragmentTagRelation
	deleteFragmentTagRelation.FragmentID = fragmentID
	code = DeleteFragmentTagRelations(deleteFragmentTagRelation)

	//删除碎片和详细介绍的对应关系
	var deleteFragmentIntroductionRelation FragmentIntroductionRelation
	deleteFragmentIntroductionRelation.FragmentID = fragmentID
	code = DeleteFragmentIntroductionRelations(deleteFragmentIntroductionRelation)

	return code
}
