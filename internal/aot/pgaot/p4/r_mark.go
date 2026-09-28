package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_r_mark_suffix_with_optional_y_consonant(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v2 = int32(121)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 <= v12 {
		v32 = v9
		v33 = v12
		v34 = v10
		v35 = v11 - v9
		v36 = v32 + v35
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
		if v36 <= v33 {
			v47 = int32(0)
			v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
			mBase = m.M
			if v49 < v47 {
				v80 = v47
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
				v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
				mBase = m.M
				if v57 != 0 {
					v80 = v47
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v67 = v58 + v35
					v69 = v67
					v74 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
					v80 = v74
				}
			}
		} else {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
			if v42 != v2 {
				v47 = int32(0)
				v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
				mBase = m.M
				if v49 < v47 {
					v80 = v47
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
					v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
					mBase = m.M
					if v57 != 0 {
						v80 = v47
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v67 = v58 + v35
						v69 = v67
						v74 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
						v80 = v74
					}
				}
			} else {
				v69 = v36 - int32(1)
				v74 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
				v80 = v74
			}
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11-int32(1)))))
		if v17 != v2 {
			v32 = v9
			v33 = v12
			v34 = v10
			v35 = v11 - v9
			v36 = v32 + v35
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
			if v36 <= v33 {
				v47 = int32(0)
				v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
				mBase = m.M
				if v49 < v47 {
					v80 = v47
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
					v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
					mBase = m.M
					if v57 != 0 {
						v80 = v47
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v67 = v58 + v35
						v69 = v67
						v74 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
						v80 = v74
					}
				}
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
				if v42 != v2 {
					v47 = int32(0)
					v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
					mBase = m.M
					if v49 < v47 {
						v80 = v47
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
						v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
						mBase = m.M
						if v57 != 0 {
							v80 = v47
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v67 = v58 + v35
							v69 = v67
							v74 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
							v80 = v74
						}
					}
				} else {
					v69 = v36 - int32(1)
					v74 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
					v80 = v74
				}
			}
		} else {
			v20 = v11 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
			v25 = int32(0)
			v26 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), v25)
			mBase = m.M
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v26 == v25 {
				v67 = v20 - v9 + v27
				v69 = v67
				v74 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
				v80 = v74
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v32 = v27
				v33 = v31
				v34 = v30
				v35 = v11 - v9
				v36 = v32 + v35
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
				if v36 <= v33 {
					v47 = int32(0)
					v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
					mBase = m.M
					if v49 < v47 {
						v80 = v47
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
						v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
						mBase = m.M
						if v57 != 0 {
							v80 = v47
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v67 = v58 + v35
							v69 = v67
							v74 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
							v80 = v74
						}
					}
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
					if v42 != v2 {
						v47 = int32(0)
						v49 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
						mBase = m.M
						if v49 < v47 {
							v80 = v47
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
							v57 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
							mBase = m.M
							if v57 != 0 {
								v80 = v47
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v67 = v58 + v35
								v69 = v67
								v74 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
								v80 = v74
							}
						}
					} else {
						v69 = v36 - int32(1)
						v74 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
						v80 = v74
					}
				}
			}
		}
	}
	return v80
}
func F_r_mark_yUm(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14365(m, l0, int32(_a_F_r_mark_yUm_0), int32(109))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_r_mark_ymUs_(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v31 = v2
		return v31
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8-int32(3) <= v7 {
			v31 = v2
			return v31
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v8-int32(1)))))
			if v16 != int32(159) {
				v31 = v2
				return v31
			} else {
				v22 = F_find_among_b(m, l0, int32(_a_F_r_mark_ymUs__0), int32(4), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v22 == int32(0) {
						v31 = v2
					} else {
						v29 = Fn14364(m, l0, int32(121))
						mBase = m.M
						v31 = v29
					}
					return v31
				}
			}
		}
	}
}
