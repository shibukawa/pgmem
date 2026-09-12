package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetCommonSlotOps(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	if l1 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+103)))
	if v11 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v43 = int32(1)
	if l1 == v43 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+99)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	if v14&int32(1) != 0 {
		v42 = v15
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
	if v23&int32(16) == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v42 = v28
	goto L3
L12:
	;
	if v14&int32(1) == int32(0) {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v14&int32(1) == int32(0) {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v42 = v34
	goto L3
L16:
	;
	v42 = int32(1592100)
	goto L3
L17:
	;
	return v42
L18:
	;
	goto L19
L19:
	;
	v49 = v43
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+v49<<(uint(int32(2))%32))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+103)))
	if v58 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return v42
L22:
	;
	if v79&int32(1) == int32(0) {
		goto L1
	} else {
		goto L30
	}
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v78 = v76
	v79 = v75
	goto L22
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+99)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+92))
	if v62 != 0 {
		v78 = v62
		v79 = v61
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+60))
	if v65 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+60))
	if v63 != 0 {
		v73 = v63
		v75 = v61
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(1592100)
	v79 = v61
	goto L22
L29:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	v73 = v65
	v75 = int32(base.Ui32(v68&int32(16)) >> (uint(int32(4)) % 32))
	goto L23
L30:
	;
	if v42 != v78 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v86 = v49 + int32(1)
	if v86 != l1 {
		v49 = v86
		goto L20
	} else {
		goto L32
	}
L32:
	;
	goto L21
}
func F_select_common_collation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	v16 = F_assign_collations_walker(m, l1, v7+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
		if v20 == int32(2) {
			if l2 != 0 {
				v56 = int32(0)
				m.G0 = v7 + int32(32)
				return v56
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(17432708))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v32 = F_get_collation_name(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
							v35 = F_get_collation_name(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v35
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v32
								F_errmsg(m, int32(697200), v7)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(567190), int32(0))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										F_parser_errposition(m, v46, v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(489967), int32(232), int32(258439))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
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
				}
			}
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v56 = v55
			m.G0 = v7 + int32(32)
			return v56
		}
	}
}
func F_select_common_type_from_oids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 != int32(705) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v158
L2:
	;
	if base.Ui32(l0) < base.Ui32(int32(2)) {
		v38 = v12
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v46 = v12
	goto L4
L4:
	;
	v48 = F_getBaseType(m, v13)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if l0 == v38 {
		v158 = v13
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v23 = v12
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1+v23<<(uint(int32(2))%32))))
	if v28 != v13 {
		v38 = v23
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v158 = v13
	goto L1
L9:
	;
	v31 = v23 + int32(1)
	if v31 != l0 {
		v23 = v31
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v46 = v38
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v48
	F_get_type_category_preferred(m, v48, v10+int32(27), v10+int32(26))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v46) < base.Ui32(l0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v158 = int32(0)
	goto L1
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v66 = v46
	v67 = v60
	goto L19
L17:
	;
	v146 = v48
	goto L18
L18:
	;
	if v146 == int32(705) {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1+v66<<(uint(int32(2))%32))))
	v72 = F_getBaseType(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v146 = v141
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v72
	if v72 == int32(705) {
		v137 = v67
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v139 = v66 + int32(1)
	if v139 != l0 {
		v66 = v139
		v67 = v137
		goto L19
	} else {
		goto L44
	}
L23:
	;
	if v72 == v67 {
		v137 = v67
		goto L22
	} else {
		goto L24
	}
L24:
	;
	F_get_type_category_preferred(m, v72, v10+int32(19), v10+int32(18))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	if v67 != int32(705) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+27)))
	if v86 != v87 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+27)) = uint8(v133)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+26)) = uint8(v135)
	v137 = v131
	goto L22
L29:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+26)))
	if v112 != 0 {
		v137 = v67
		goto L22
	} else {
		goto L39
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v97 = F_format_type_be(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v100 = F_format_type_be(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v97
	F_errmsg(m, int32(451597), v10)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(490929), int32(1536), int32(170563))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v119 = F_can_coerce_type(m, int32(1), v10+int32(28), v10+int32(20), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	if v119 == int32(0) {
		v137 = v67
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v129 = F_can_coerce_type(m, int32(1), v10+int32(20), v10+int32(28), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	if v129 != 0 {
		v137 = v67
		goto L22
	} else {
		goto L43
	}
L43:
	;
	goto L28
L44:
	;
	goto L20
L45:
	;
	v152 = int32(25)
	goto L47
L46:
	;
	v152 = v146
	goto L47
L47:
	;
	v158 = v152
	goto L1
}
