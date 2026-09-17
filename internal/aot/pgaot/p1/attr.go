package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_execute_attr_map_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v14 < v13 {
		F_slot_getsomeattrs_int(m, l1, v13)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			m.T0[v21].(func(*base.Module, int32))(m, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if int32(0) < v11 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v32 = int32(0)
					for {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v32<<(uint(int32(1))%32)))))
						if v45 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v27+v32<<(uint(int32(2))%32)))) = int32(0)
							v67 = int32(1)
						} else {
							v54 = int32(2)
							v58 = v45 - int32(1)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v29+v58<<(uint(v54)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v27+v32<<(uint(v54)%32)))) = v62
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v28))))
							v67 = v65
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v32+v26))) = uint8(v67)
						v70 = v32 + int32(1)
						if v70 != v11 {
							v32 = v70
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				v83 = v81 & int32(_a_F_execute_attr_map_slot_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v83)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
				*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v86)
				return l2
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
		m.T0[v21].(func(*base.Module, int32))(m, l2)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if int32(0) < v11 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v32 = int32(0)
				for {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v32<<(uint(int32(1))%32)))))
					if v45 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v27+v32<<(uint(int32(2))%32)))) = int32(0)
						v67 = int32(1)
					} else {
						v54 = int32(2)
						v58 = v45 - int32(1)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v29+v58<<(uint(v54)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v27+v32<<(uint(v54)%32)))) = v62
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v28))))
						v67 = v65
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v32+v26))) = uint8(v67)
					v70 = v32 + int32(1)
					if v70 != v11 {
						v32 = v70
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			v83 = v81 & int32(_a_F_execute_attr_map_slot_0)
			*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v83)
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
			*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v86)
			return l2
		}
	}
}
func F_readAttrNumberCols(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L38
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v113 = int32(0)
		goto L6
	case 1:
		goto L7
	default:
		goto L2
	}
L4:
	;
	goto L5
L5:
	;
	goto L1
L6:
	;
	m.G0 = v8 + int32(16)
	return v113
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v20 = F_palloc(m, l0<<(uint(int32(1))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if int32(0) < l0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v28 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v101 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v101 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L14:
	;
	v33 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v33 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(41) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v45 = v33
	goto L19
L18:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(1))%32)))) = uint16(v89)
	v92 = v28 + int32(1)
	if v92 != l0 {
		v28 = v92
		goto L14
	} else {
		goto L34
	}
L19:
	;
	v50 = v45 + int32(1)
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45))))
	v52 = F___isspace(m, v51)
	mBase = m.M
	if v52 != 0 {
		v45 = v50
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v53 = int32(1)
	switch v51&int32(255) - int32(43) {
	case 0:
		v59 = v53
		goto L23
	default:
		v61 = v51
		v62 = v45
		v63 = v53
		goto L22
	case 2:
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	v64 = int32(0)
	v66 = v61 - int32(48)
	if base.Ui32(v66) <= base.Ui32(int32(9)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
	v61 = v60
	v62 = v50
	v63 = v59
	goto L22
L24:
	;
	v59 = int32(0)
	goto L23
L25:
	;
	v69 = v64
	v70 = v66
	v71 = v62
	goto L28
L26:
	;
	v83 = v64
	goto L27
L27:
	;
	if v63 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v73 = int32(10)
	v75 = v69*v73 - v70
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+1)))
	v80 = v76 - int32(48)
	if base.Ui32(v80) < base.Ui32(v73) {
		v69 = v75
		v70 = v80
		v71 = v71 + int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v83 = v75
	goto L27
L30:
	;
	goto L29
L31:
	;
	v89 = int32(0) - v83
	goto L33
L32:
	;
	v89 = v83
	goto L33
L33:
	;
	goto L18
L34:
	;
	goto L15
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v104 != int32(1) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v107 != int32(41) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v113 = v20
	goto L6
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v124
	F_errmsg_internal(m, int32(_a_F_readAttrNumberCols_0), v8)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_readAttrNumberCols_1), int32(692), int32(_a_F_readAttrNumberCols_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_readAttrNumberCols_3), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_readAttrNumberCols_1), int32(692), int32(_a_F_readAttrNumberCols_2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
