package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	v13 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v60 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v19 = v16 + v13<<(uint(int32(3))%32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = int32(0)
	if base.B2i32(v20 == v21)|base.B2i32(v20 == v19) == v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v28 = v20
	goto L9
L7:
	;
	v47 = v12
	goto L8
L8:
	;
	v52 = v13 + int32(1)
	if v52 < v47 {
		v12 = v47
		v13 = v52
		goto L4
	} else {
		goto L18
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if int32(0) < v34 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v47 = v44
	goto L8
L11:
	;
	if v19 != v33 {
		v28 = v33
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+28)) = uint8(v37)
	goto L11
L13:
	;
	goto L14
L14:
	;
	F_CatCacheRemoveCList(m, l0, v28-int32(8))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	v65 = v60
	v68 = v2
	goto L22
L20:
	;
	goto L21
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ResetCatalogCache[0]))
	if v122 != 0 {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v72 = v69 + v68<<(uint(int32(3))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v74 = int32(0)
	if base.B2i32(v73 == v74)|base.B2i32(v73 == v72) == v74 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v81 = v73
	goto L27
L25:
	;
	v108 = v65
	goto L26
L26:
	;
	v113 = v68 + int32(1)
	if v113 < v108 {
		v65 = v108
		v68 = v113
		goto L22
	} else {
		goto L38
	}
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v87 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v108 = v105
	goto L26
L29:
	;
	if v72 != v86 {
		v81 = v86
		goto L27
	} else {
		goto L37
	}
L30:
	;
	F_CatCacheRemoveCTup(m, l0, v81-int32(24))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L15
	} else {
		goto L36
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+36))
	if v90 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+12)) = uint8(v97)
	goto L29
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+32))
	if v93 <= int32(0) {
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
	v124 = v122
	goto L42
L40:
	;
	goto L41
L41:
	;
	return
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if l0 == v129 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+9)) = uint8(v131)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v133 != 0 {
		v124 = v133
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
}
