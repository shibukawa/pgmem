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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	v4 = int32(0)
	if l1 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v55
	return v55
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v4
	v18 = v4
	v19 = v4
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v18<<(uint(int32(2))%32))))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+26)))
	if v25 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v47
	return v48
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v29 = F_exprType(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v47 = v17
	v48 = v19
	goto L8
L8:
	;
	v50 = v18 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v50 < v51 {
		v17 = v47
		v18 = v50
		v19 = v48
		goto L4
	} else {
		goto L16
	}
L9:
	;
	return int32(0)
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v34 = F_exprTypmod(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v37 = F_exprCollation(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v39 = F_generate_new_exec_param(m, l0, v29, v34, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v41 = F_lappend(m, v19, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v44 = F_lappend_int(m, v17, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v47 = v44
	v48 = v41
	goto L8
L16:
	;
	goto L5
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
