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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v12 = F_table_open(m, int32(2604), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v8, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v24 = int32(3)
			F_ScanKeyInit(m, v8+int32(56), v24, v24, int32(63), base.I64_extend_i32_s(l1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v35 = F_systable_beginscan(m, v12, int32(2656), int32(1), v30, int32(2), v8)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = F_systable_getnext(m, v35)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v37 != 0 {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40)))
							v43 = v42
						} else {
							v43 = v30
						}
						F_systable_endscan(m, v35)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_relation_close(m, v12, int32(1))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(112)
								return v43
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
	var v18 int32
	_ = v18
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
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
	v18 = v3
	goto L7
L5:
	;
	v55 = v3
	goto L6
L6:
	;
	return v55
L7:
	;
	if int32(0) <= v16 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v55 = v46
	goto L6
L9:
	;
	v48 = v16 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 <= v49 {
		v16 = v48
		v18 = v46
		goto L7
	} else {
		goto L19
	}
L10:
	;
	if v16 == int32(0) {
		v46 = v18
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
		v46 = v18
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
		v46 = v18
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v43 = F_bms_add_member(m, v18, v16+int32(7))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v46 = v43
	goto L9
L19:
	;
	goto L8
}
