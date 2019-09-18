package service

//QueryFragment 碎片
type QueryFragment struct {
	Fragment
	LastID int `json:"lastId"`
}

//GetFragments 查询碎片
func GetFragments(fragment QueryFragment) []Fragment {
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

//GetFragmentByID 查询碎片
func GetFragmentByID(fragmentID int) Fragment {
	if fragmentID > 0 {
		return queryFragmentByID(fragmentID)
	}
	var fragment Fragment
	return fragment
}

//DeleteFragmentByID 删除碎片
func DeleteFragmentByID(fragmentID int) int {
	if fragmentID > 0 {
		return deleteFragmentByID(fragmentID)
	}
	return 0
}
