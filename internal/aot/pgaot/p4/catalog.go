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
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(394)
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v58 {
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
	v45 = v12
	goto L8
L8:
	;
	v50 = v13 + int32(1)
	if v50 < v45 {
		v12 = v45
		v13 = v50
		goto L4
	} else {
		goto L18
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+48))
	if int32(0) < v34 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v45 = v42
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
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+52)) = uint8(v37)
	goto L11
L13:
	;
	goto L14
L14:
	;
	F_CatCacheRemoveCList(m, l0, v28)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	v63 = v58
	v66 = v2
	goto L22
L20:
	;
	goto L21
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ResetCatalogCache[0]))
	if v118 != 0 {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = v67 + v66<<(uint(int32(3))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = int32(0)
	if base.B2i32(v71 == v72)|base.B2i32(v71 == v70) == v72 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v79 = v71
	goto L27
L25:
	;
	v104 = v63
	goto L26
L26:
	;
	v109 = v66 + int32(1)
	if v109 < v104 {
		v63 = v104
		v66 = v109
		goto L22
	} else {
		goto L38
	}
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	if v85 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v104 = v101
	goto L26
L29:
	;
	if v70 != v84 {
		v79 = v84
		goto L27
	} else {
		goto L37
	}
L30:
	;
	F_CatCacheRemoveCTup(m, l0, v79)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L15
	} else {
		goto L36
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	if v88 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+52)) = uint8(v95)
	goto L29
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+48))
	if v91 <= int32(0) {
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
	v120 = v118
	goto L42
L40:
	;
	goto L41
L41:
	;
	return
L42:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if l0 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+9)) = uint8(v127)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v129 != 0 {
		v120 = v129
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
}
