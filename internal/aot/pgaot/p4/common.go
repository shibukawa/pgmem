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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	v42 = int32(1)
	if l1 == v42 {
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
		v40 = v15
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
	v40 = v28
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
	v40 = v34
	goto L3
L16:
	;
	v40 = int32(_a_F_ExecGetCommonSlotOps_0)
	goto L3
L17:
	;
	return v40
L18:
	;
	goto L19
L19:
	;
	v50 = v42
	goto L20
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v50<<(uint(int32(2))%32))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+103)))
	if v57 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return v40
L22:
	;
	if base.B2i32(v78&int32(1) == int32(0))|base.B2i32(v40 != v77) != 0 {
		goto L1
	} else {
		goto L30
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v77 = v75
	v78 = v74
	goto L22
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+99)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+92))
	if v61 != 0 {
		v77 = v61
		v78 = v60
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+60))
	if v64 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+60))
	if v62 != 0 {
		v72 = v62
		v74 = v60
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v77 = int32(_a_F_ExecGetCommonSlotOps_0)
	v78 = v60
	goto L22
L29:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)))
	v72 = v64
	v74 = int32(base.Ui32(v67&int32(16)) >> (uint(int32(4)) % 32))
	goto L23
L30:
	;
	v86 = v50 + int32(1)
	if v86 != l1 {
		v50 = v86
		goto L20
	} else {
		goto L31
	}
L31:
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
								F_errmsg(m, int32(_a_F_select_common_collation_0), v7)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_select_common_collation_1), int32(0))
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
											F_errfinish(m, int32(_a_F_select_common_collation_2), int32(232), int32(_a_F_select_common_collation_3))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v15 != int32(705) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v166
L2:
	;
	if base.Ui32(l0) < base.Ui32(int32(2)) {
		v42 = v14
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v52 = v14
	goto L4
L4:
	;
	v56 = F_getBaseType(m, v15)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if l0 == v42 {
		v166 = v15
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v25 = v14
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+v25<<(uint(int32(2))%32))))
	if v32 != v15 {
		v42 = v25
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v166 = v15
	goto L1
L9:
	;
	v35 = v25 + int32(1)
	if v35 != l0 {
		v25 = v35
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v52 = v42
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v56
	F_get_type_category_preferred(m, v56, v12+int32(27), v12+int32(26))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v52) < base.Ui32(l0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v166 = int32(0)
	goto L1
L16:
	;
	v72 = v56
	v73 = v52
	goto L19
L17:
	;
	v152 = v56
	goto L18
L18:
	;
	if v152 == int32(705) {
		goto L44
	} else {
		goto L45
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1+v73<<(uint(int32(2))%32))))
	v81 = F_getBaseType(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	v152 = v142
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v81
	if base.B2i32(v81 == int32(705))|base.B2i32(v72 == v81) != 0 {
		v142 = v72
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v146 = v73 + int32(1)
	if v146 != l0 {
		v72 = v142
		v73 = v146
		goto L19
	} else {
		goto L43
	}
L23:
	;
	F_get_type_category_preferred(m, v81, v12+int32(19), v12+int32(18))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	if v72 != int32(705) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+19)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+27)))
	if v96 != v97 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v81
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+27)) = uint8(v138)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)) = uint8(v140)
	v142 = v81
	goto L22
L28:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)))
	if v120 != 0 {
		v142 = v72
		goto L22
	} else {
		goto L38
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v106 = F_format_type_be(m, v72)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v108 = F_format_type_be(m, v81)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v106
	F_errmsg(m, int32(_a_F_select_common_type_from_oids_0), v12)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_select_common_type_from_oids_1), int32(1536), int32(_a_F_select_common_type_from_oids_2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
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
	v123 = v12 + int32(28)
	v125 = v12 + int32(20)
	v127 = F_can_coerce_type(m, int32(1), v123, v125, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	if v127 == int32(0) {
		v142 = v72
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v133 = F_can_coerce_type(m, int32(1), v125, v123, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	if v133 != 0 {
		v142 = v72
		goto L22
	} else {
		goto L42
	}
L42:
	;
	goto L27
L43:
	;
	goto L20
L44:
	;
	v160 = int32(25)
	goto L46
L45:
	;
	v160 = v152
	goto L46
L46:
	;
	v166 = v160
	goto L1
}
