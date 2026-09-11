package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CachedPlanGetTargetList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v5 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = F_RevalidateCachedQuery(m, l0, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v15 == int32(0) {
		v39 = v2
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v42 = F_FetchStatementTargetList(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		v39 = v2
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v23 = int32(0)
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+v23<<(uint(int32(2))%32))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	if v31 == int32(1) {
		v39 = v30
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v39 = int32(0)
	goto L6
L11:
	;
	v35 = v23 + int32(1)
	if v18 != v35 {
		v23 = v35
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	return v42
}
func F_DropCachedPlan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v4 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v8
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)) = uint8(v12)
	} else {
	}
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v15 == int32(0) {
		v31 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v31
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
		if v33 == v31 {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			F_MemoryContextDelete(m, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
		v22 = v20 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v22
		if v22 != 0 {
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v31
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
			if v33 == v31 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				F_MemoryContextDelete(m, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
			if v26 != 0 {
				v31 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v31
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
				if v33 == v31 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					F_MemoryContextDelete(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
				F_MemoryContextDelete(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v31 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v31
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
					if v33 == v31 {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						F_MemoryContextDelete(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			}
		}
	}
}
