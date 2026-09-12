package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogTupleDelete(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_simple_heap_delete(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_CatalogTupleInsert(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v5 = F_palloc0(m, int32(216))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+52)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(388)
		F_ExecOpenIndices(m, v5, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_simple_heap_insert(m, l0, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_CatalogIndexInsert(m, v5, l1, int32(1))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_ExecCloseIndices(m, v5)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						F_pfree(m, v5)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_ResetCatalogCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v2 < v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v7
	v14 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v57 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v19 = v16 + v14<<(uint(int32(3))%32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v20 == int32(0) {
		v44 = v12
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v49 = v14 + int32(1)
	if v49 < v44 {
		v12 = v44
		v14 = v49
		goto L4
	} else {
		goto L18
	}
L7:
	;
	if v20 == v19 {
		v44 = v12
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = v20
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if int32(0) < v31 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v44 = v41
	goto L6
L11:
	;
	if v19 != v30 {
		v25 = v30
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+28)) = uint8(v34)
	goto L11
L13:
	;
	goto L14
L14:
	;
	F_CatCacheRemoveCList(m, l0, v25-int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	goto L11
L17:
	;
	goto L10
L18:
	;
	goto L5
L19:
	;
	v62 = v57
	v65 = v2
	goto L22
L20:
	;
	goto L21
L21:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[1359]))
	if v116 != 0 {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v69 = v66 + v65<<(uint(int32(3))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v70 == int32(0) {
		v102 = v62
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v107 = v65 + int32(1)
	if v107 < v102 {
		v62 = v102
		v65 = v107
		goto L22
	} else {
		goto L38
	}
L25:
	;
	if v70 == v69 {
		v102 = v62
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v75 = v70
	goto L27
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	if v81 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = v99
	goto L24
L29:
	;
	if v69 != v80 {
		v75 = v80
		goto L27
	} else {
		goto L37
	}
L30:
	;
	F_CatCacheRemoveCTup(m, l0, v75-int32(24))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L15
	} else {
		goto L36
	}
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+36))
	if v84 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+12)) = uint8(v91)
	goto L29
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
	if v87 <= int32(0) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L29
L37:
	;
	goto L28
L38:
	;
	goto L23
L39:
	;
	v118 = v116
	goto L42
L40:
	;
	goto L41
L41:
	;
	return
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if l0 == v123 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+9)) = uint8(v125)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	if v127 != 0 {
		v118 = v127
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
}
