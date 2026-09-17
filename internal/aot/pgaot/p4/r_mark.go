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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
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
		if v33 < v36 {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
			if v43 == v2 {
				v72 = int32(0)
			} else {
				v46 = int32(0)
				v48 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
				mBase = m.M
				if v48 < v46 {
					v72 = v46
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
					v56 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
					mBase = m.M
					if v56 != 0 {
						v72 = v46
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v66 = v57 + v35
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
						v72 = int32(1)
					}
				}
			}
		} else {
			v46 = int32(0)
			v48 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
			mBase = m.M
			if v48 < v46 {
				v72 = v46
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
				v56 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
				mBase = m.M
				if v56 != 0 {
					v72 = v46
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v66 = v57 + v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
					v72 = int32(1)
				}
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
			if v33 < v36 {
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
				if v43 == v2 {
					v72 = int32(0)
				} else {
					v46 = int32(0)
					v48 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
					mBase = m.M
					if v48 < v46 {
						v72 = v46
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
						v56 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
						mBase = m.M
						if v56 != 0 {
							v72 = v46
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v66 = v57 + v35
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
							v72 = int32(1)
						}
					}
				}
			} else {
				v46 = int32(0)
				v48 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
				mBase = m.M
				if v48 < v46 {
					v72 = v46
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
					v56 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
					mBase = m.M
					if v56 != 0 {
						v72 = v46
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v66 = v57 + v35
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
						v72 = int32(1)
					}
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
				v66 = v20 - v9 + v27
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
				v72 = int32(1)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v32 = v27
				v33 = v31
				v34 = v30
				v35 = v11 - v9
				v36 = v32 + v35
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
				if v33 < v36 {
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v34-int32(1)))))
					if v43 == v2 {
						v72 = int32(0)
					} else {
						v46 = int32(0)
						v48 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
						mBase = m.M
						if v48 < v46 {
							v72 = v46
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
							v56 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
							mBase = m.M
							if v56 != 0 {
								v72 = v46
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v66 = v57 + v35
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
								v72 = int32(1)
							}
						}
					}
				} else {
					v46 = int32(0)
					v48 = F_skip_b_utf8(m, v34, v36, v33, int32(1))
					mBase = m.M
					if v48 < v46 {
						v72 = v46
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
						v56 = F_in_grouping_b_U(m, l0, int32(_a_F_r_mark_suffix_with_optional_y_consonant_0), int32(97), int32(305), int32(0))
						mBase = m.M
						if v56 != 0 {
							v72 = v46
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v66 = v57 + v35
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
							v72 = int32(1)
						}
					}
				}
			}
		}
	}
	return v72
}
func F_r_mark_yUm(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13982(m, l0, int32(_a_F_r_mark_yUm_0), int32(109))
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v30 = v2
		return v30
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8-int32(3) <= v7 {
			v30 = v2
			return v30
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v8-int32(1)))))
			if v16 != int32(159) {
				v30 = v2
				return v30
			} else {
				v21 = F_find_among_b(m, l0, int32(_a_F_r_mark_ymUs__0), int32(4))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						v30 = v2
					} else {
						v28 = Fn13981(m, l0, int32(121))
						mBase = m.M
						v30 = v28
					}
					return v30
				}
			}
		}
	}
}
func F_r_mark_ysA(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = v4 - int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v7 {
		v34 = v2
		return v34
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v6))))
		if base.B2i32(v11&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v11)%32)&int32(_a_F_r_mark_ysA_0) == int32(0)) != 0 {
			v34 = v2
			return v34
		} else {
			v25 = F_find_among_b(m, l0, int32(_a_F_r_mark_ysA_1), int32(8))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					v34 = v2
				} else {
					v32 = Fn13981(m, l0, int32(121))
					mBase = m.M
					v34 = v32
				}
				return v34
			}
		}
	}
}
