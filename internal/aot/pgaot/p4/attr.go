package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAttrDefaultOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v12 = F_table_open(m, int32(2604), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v8, int32(2), int32(3), int32(184), l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v23 = int32(3)
			F_ScanKeyInit(m, v8+int32(48), v23, v23, int32(63), l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(0)
				v33 = F_systable_beginscan(m, v12, int32(2656), int32(1), v28, int32(2), v8)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = F_systable_getnext(m, v33)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 != 0 {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)))
							v41 = v40
						} else {
							v41 = v28
						}
						F_systable_endscan(m, v33)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_sequence_close(m, v12, int32(1))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(96)
								return v41
							}
						}
					}
				}
			}
		}
	}
}
func F_execute_attr_map_cols(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v3 = int32(0)
	if l1 == v3 {
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
	v10 = int32(-7)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v10
	v17 = v3
	goto L7
L5:
	;
	v54 = v3
	goto L6
L6:
	;
	return v54
L7:
	;
	if int32(0) <= v16 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v54 = v45
	goto L6
L9:
	;
	v48 = v16 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 <= v49 {
		v16 = v48
		v17 = v45
		goto L7
	} else {
		goto L19
	}
L10:
	;
	if v16 == int32(0) {
		v45 = v17
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v32 = v16
	goto L12
L12:
	;
	v35 = F_bms_is_member(m, v32+int32(7), l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23+v16<<(uint(int32(1))%32)-int32(2)))))
	if v29 == int32(0) {
		v45 = v17
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v32 = v29
	goto L12
L15:
	;
	return int32(0)
L16:
	;
	if v35 == int32(0) {
		v45 = v17
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v43 = F_bms_add_member(m, v17, v16+int32(7))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v45 = v43
	goto L9
L19:
	;
	goto L8
}
