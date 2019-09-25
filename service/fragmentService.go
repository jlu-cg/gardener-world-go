package service

//QueryFragment 碎片
type QueryFragment struct {
	Fragment
	LastID int `json:"lastId"`
}

//QueryFragments 查询碎片
func QueryFragments(fragment QueryFragment) []Fragment {
	fragments := queryFragments(fragment.Fragment, fragment.LastID)
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

//QueryFragmentByID 查询碎片
func QueryFragmentByID(fragmentID int) Fragment {
	if fragmentID > 0 {
		var queryFragment Fragment
		queryFragment.ID = fragmentID
		fragments := queryFragments(queryFragment, 0)
		if len(fragments) > 0 {
			return fragments[0]
		}
	}
	var fragment Fragment
	return fragment
}

//DeleteFragmentByID 删除碎片
func DeleteFragmentByID(fragmentID int) int {
	if fragmentID <= 0 {
		return 0
	}
	code := deleteFragmentByID(fragmentID)

	var deleteFragmentTagRelation FragmentTagRelation
	deleteFragmentTagRelation.FragmentID = fragmentID
	code = DeleteFragmentTagRelations(deleteFragmentTagRelation)
	return code
}
