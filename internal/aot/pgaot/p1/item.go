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
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if l3 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return v167
L2:
	;
	v157 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, l4, (l3^int32(1))&v14)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L47
	}
L3:
	;
	if v14&int32(1) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = int64(0)
	v27 = F_executeItemOptUnwrapTarget(m, l0, l1, l2, v10+int32(-16), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v27 == int32(2) {
		v167 = int32(2)
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v33 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v34 != 0 {
		v49 = v34
		v50 = v33
		v51 = v6
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v55 = v49
	v57 = v50
	goto L16
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if v35 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(0)
	v49 = v38
	v50 = v33
	v51 = v38
	goto L8
L11:
	;
	goto L12
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(1) < v44 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v40 + int32(4)
	goto L15
L14:
	;
	v47 = int32(0)
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v49 = v48
	v50 = v47
	v51 = v35
	goto L8
L16:
	;
	v61 = int32(0)
	if v57 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v65 = v57 + int32(4)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if base.Ui32(v65) < base.Ui32(v67+v68<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v75 = v61
	v76 = v61
	goto L20
L20:
	;
	if v55 == int32(0) {
		v167 = v61
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v73 = v65
	goto L23
L22:
	;
	v73 = int32(0)
	goto L23
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v75 = v74
	v76 = v73
	goto L20
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	switch v79 - int32(16) {
	case 0:
		goto L28
	default:
		goto L26
	case 2:
		goto L29
	}
L25:
	;
	v55 = v75
	v57 = v76
	goto L16
L26:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v130 != 0 {
		goto L39
	} else {
		goto L40
	}
L27:
	;
	v120 = int32(0)
	v121 = int32(1)
	v126 = F_executeAnyItem(m, l0, v120, v82, l4, v121, v121, v121, v120, v120)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L38
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L35
	}
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83&int32(536870912) != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if v83&int32(1073741824) != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v92
	F_errmsg_internal(m, int32(28728), v10+int32(-48))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(488788), int32(3629), int32(363441))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
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
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v108
	F_errmsg_internal(m, int32(474733), v10+int32(-32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(488788), int32(1680), int32(25522))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v55
	v139 = F_list_make2_impl(m, v10+int32(-52), v10+int32(-56))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v144 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v139
	goto L25
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v55
	goto L25
L44:
	;
	goto L45
L45:
	;
	v148 = F_lappend(m, v144, v55)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v148
	goto L25
L47:
	;
	v167 = v157
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v18 != 0 {
				v20 = F_fwrite(m, l0, int32(1), v18, l2)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					if v20 != v18 {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(288263), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									F_errfinish(m, int32(487996), int32(6806), int32(284494))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
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
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(288263), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(487996), int32(6802), int32(284494))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
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
