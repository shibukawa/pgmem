package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsToastRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+68))
	if v3 != int32(99) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_IsToastRelation[0]))
		v14 = base.B2i32(v8 != int32(0)) & base.B2i32(v3 == v8)
	} else {
		v14 = int32(1)
	}
	return v14
}
func F_get_toast_snapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_get_toast_snapshot[0]))
	if v3 != 0 {
		v19 = int32(1)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_get_toast_snapshot[1]))
		v7 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_get_toast_snapshot[2]))
		if base.B2i32(v6 == v7)|base.B2i32(v10 == v7) != 0 {
			v19 = base.B2i32(v10 != int32(0))
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v14 != 0 {
				v19 = base.B2i32(v10 != int32(0))
			} else {
				v19 = int32(0)
			}
		}
	}
	if v19 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_get_toast_snapshot_0), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_get_toast_snapshot_1), int32(653), int32(_a_F_get_toast_snapshot_2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return int32(_a_F_get_toast_snapshot_3)
	}
}
func F_toast_close_indexes(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	v3 = int32(0)
	if v3 < l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_pfree(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v8<<(uint(int32(2))%32))))
	F_relation_close(m, v12, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v17 = v8 + int32(1)
	if v17 != l1 {
		v8 = v17
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	return
}
func F_toast_tuple_find_biggest_attribute(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 <= int32(0) {
		return int32(-1)
	} else {
		if l1 != 0 {
			v22 = int32(48)
		} else {
			v22 = int32(16)
		}
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v33 = int32(0)
		v35 = int32(24)
		v36 = int32(-1)
		for {
			v44 = v26 + v33*int32(12)
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)))
			if v22&v45 != 0 {
				v81 = v35
				v82 = v36
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v48 = int32(2)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v33<<(uint(v48)%32))))
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v52&int32(3) == v48 {
					v60 = l1
				} else {
					v60 = int32(0)
				}
				if base.B2i32(v52 == int32(1))|v60 != 0 {
					v81 = v35
					v82 = v36
				} else {
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(4))%32)+v33*int32(100))+104)))
					if l2 != 0 {
						if v65 != int32(109) {
							v81 = v35
							v82 = v36
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v77 = base.B2i32(v35 < v76)
							if v35 < v76 {
								v78 = v33
							} else {
								v78 = v36
							}
							if v35 < v76 {
								v79 = v76
							} else {
								v79 = v35
							}
							v81 = v79
							v82 = v78
						}
					} else {
						v69 = v65 - int32(101)
						if base.B2i32(v69 == int32(0))|base.B2i32(v69 == int32(19)) != 0 {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v77 = base.B2i32(v35 < v76)
							if v35 < v76 {
								v78 = v33
							} else {
								v78 = v36
							}
							if v35 < v76 {
								v79 = v76
							} else {
								v79 = v35
							}
							v81 = v79
							v82 = v78
						} else {
							v81 = v35
							v82 = v36
						}
					}
				}
			}
			v85 = v33 + int32(1)
			if v85 != v15 {
				v33 = v85
				v35 = v81
				v36 = v82
				continue
			} else {
				break
			}
			break
		}
		return v82
	}
}
