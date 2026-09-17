package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_executeItemOptUnwrapResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if base.B2i32(l3 == v6)|base.B2i32(v16&int32(1) == v6) == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return v170
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = int64(0)
	v30 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, v10+int32(-16), int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v161 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, l4, (l3^int32(1))&v16)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L5
	} else {
		goto L47
	}
L5:
	;
	return int32(0)
L6:
	;
	if v30 == int32(2) {
		v170 = int32(2)
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v36 = int32(0)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v37 != 0 {
		v53 = v37
		v54 = v36
		v55 = v6
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = v53
	v62 = v54
	goto L16
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = int32(0)
	v53 = v41
	v54 = v36
	v55 = v41
	goto L8
L11:
	;
	goto L12
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if int32(1) < v47 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = v43 + int32(4)
	goto L15
L14:
	;
	v50 = int32(0)
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v53 = v51
	v54 = v50
	v55 = v38
	goto L8
L16:
	;
	v65 = int32(0)
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = v62 + int32(4)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if base.Ui32(v69) < base.Ui32(v71+v72<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v79 = v65
	v80 = v65
	goto L20
L20:
	;
	if v59 == int32(0) {
		v170 = v65
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v77 = v69
	goto L23
L22:
	;
	v77 = int32(0)
	goto L23
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v79 = v78
	v80 = v77
	goto L20
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch v83 - int32(16) {
	case 0:
		goto L28
	default:
		goto L26
	case 2:
		goto L29
	}
L25:
	;
	v59 = v79
	v62 = v80
	goto L16
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v134 != 0 {
		goto L39
	} else {
		goto L40
	}
L27:
	;
	v124 = int32(0)
	v125 = int32(1)
	v130 = F_executeAnyItem(m, l0, v124, v86, l4, v125, v125, v125, v124, v124)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L38
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L35
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87&int32(536870912) != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if v87&int32(1073741824) != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v96
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapResult_0), v10+int32(-48))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapResult_1), int32(3629), int32(_a_F_executeItemOptUnwrapResult_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v112
	F_errmsg_internal(m, int32(_a_F_executeItemOptUnwrapResult_3), v10+int32(-32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_executeItemOptUnwrapResult_1), int32(1680), int32(_a_F_executeItemOptUnwrapResult_4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	goto L25
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v59
	v143 = F_list_make2_impl(m, v10+int32(-52), v10+int32(-56))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v148 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v143
	goto L25
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v59
	goto L25
L44:
	;
	goto L45
L45:
	;
	v152 = F_lappend(m, v148, v59)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v152
	goto L25
L47:
	;
	v170 = v161
	goto L1
}
func F_write_item(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v14 = F_fwrite(m, v7+int32(12), int32(1), int32(4), l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(4) {
			if l1 != 0 {
				v19 = F_fwrite(m, l0, int32(1), l1, l2)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					if v19 != l1 {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_write_item_0), int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_write_item_1), int32(_a_F_write_item_2), int32(_a_F_write_item_3))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_write_item_0), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_write_item_1), int32(_a_F_write_item_4), int32(_a_F_write_item_3))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
