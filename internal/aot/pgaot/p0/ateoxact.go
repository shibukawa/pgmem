package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_ApplyLauncher(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	if l0 == int32(0) {
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_ApplyLauncher[0])))
		if v5&int32(1) == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_ApplyLauncher[1]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v12 == int32(0) {
			} else {
				v16 = F_pgmem_kill(m, v12, int32(10))
				mBase = m.M
			}
		}
	}
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_ApplyLauncher[0])) = uint8(v19)
	return
}
func F_AtEOXact_RelationMap(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = int32(0)
	if l1|base.B2i32(l0 == v3) == v3 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[0]))
		if v9 != 0 {
			F_perform_relmap_update(m, int32(1), int32(_a_F_AtEOXact_RelationMap_0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[0])) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[1]))
				if v18 == int32(0) {
					return
				} else {
					F_perform_relmap_update(m, int32(0), int32(_a_F_AtEOXact_RelationMap_1))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[1])) = int32(0)
						return
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[1]))
			if v18 == int32(0) {
				return
			} else {
				F_perform_relmap_update(m, int32(0), int32(_a_F_AtEOXact_RelationMap_1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[1])) = int32(0)
					return
				}
			}
		}
	} else {
		v29 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[1])) = v29
		*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[0])) = v29
		*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[2])) = v29
		*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationMap[3])) = v29
		return
	}
}
func F_AtEOXact_SPI(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[0]))
	if v7 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[1]))
	v14 = v11 + v7<<(uint(int32(6))%32)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+41)))
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[2])) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[3])) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v25 = v7 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[0])) = v25
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[4])) = v22
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = v11 + v25<<(uint(int32(6))%32)
	goto L6
L5:
	;
	v34 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[5])) = v34
	if v7 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v41 = v25
	goto L9
L9:
	;
	v45 = v11 + v41<<(uint(int32(6))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+41)))
	if v46 != 0 {
		goto L7
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[2])) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[3])) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+60))
	v56 = v41 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[0])) = v56
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[4])) = v53
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v65 = v11 + v56<<(uint(int32(6))%32)
	goto L14
L13:
	;
	v65 = int32(0)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_SPI[5])) = v65
	if int32(0) < v41 {
		v41 = v56
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	v78 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	if v78 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(64))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_AtEOXact_SPI_0), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(_a_F_AtEOXact_SPI_1), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_SPI_2), int32(472), int32(_a_F_AtEOXact_SPI_3))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L1
}
