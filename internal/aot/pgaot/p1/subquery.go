package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_subquery_params(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v4 = int32(0)
	if l1 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v10
	return v10
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
	return v17
L5:
	;
	goto L6
L6:
	;
	v25 = v4
	v26 = v4
	v27 = v4
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+26)))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v55
	return v56
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v37 = F_exprType(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v55 = v25
	v56 = v27
	goto L11
L11:
	;
	v58 = v26 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v58 < v59 {
		v25 = v55
		v26 = v58
		v27 = v56
		goto L7
	} else {
		goto L19
	}
L12:
	;
	return int32(0)
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v42 = F_exprTypmod(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v45 = F_exprCollation(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v47 = F_generate_new_exec_param(m, l0, v37, v42, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v49 = F_lappend(m, v27, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v52 = F_lappend_int(m, v25, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v55 = v52
	v56 = v49
	goto L11
L19:
	;
	goto L8
}
func F_subquery_push_qual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v6 != 0 {
		F_recurse_push_qual(m, v6, l0, l1, l2, l3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v11 = int32(0)
		v15 = F_ReplaceVarsFromTargetList(m, l3, l2, l1, v9, v10, v11, v11, l0+int32(39))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
			if v17 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				v24 = F_make_and_qual(m, v23, v15)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v24
					return
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				if v18 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					v24 = F_make_and_qual(m, v23, v15)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v24
						return
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
					if v19 != 0 {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						v24 = F_make_and_qual(m, v23, v15)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v24
							return
						}
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v20 == int32(0) {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							v29 = F_make_and_qual(m, v28, v15)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v29
								return
							}
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							v24 = F_make_and_qual(m, v23, v15)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v24
								return
							}
						}
					}
				}
			}
		}
	}
}
