package service

//GetFragments 查询碎片
func GetFragments(fragment Fragment, lastID int) []Fragment {
	if lastID < 0 {
		lastID = 0
	}
	fragments := queryFragments(fragment, lastID)
	return fragments
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

//GetFragmentByID 查询碎片
func GetFragmentByID(fragmentID int, tagID int) Fragment {
	if fragmentID > 0 {
		return queryFragmentByID(fragmentID)
	}
	var fragment Fragment
	fragment.TagID = tagID
	return fragment
}

//DeleteFragmentByID 删除碎片
func DeleteFragmentByID(fragmentID int) int {
	if fragmentID > 0 {
		return deleteFragmentByID(fragmentID)
	}
	return 0
}
