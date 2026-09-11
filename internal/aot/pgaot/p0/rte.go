package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_markRTEForSelectPriv(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = l1 - int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15<<(uint(int32(2))%32))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v20 == int32(0) {
		v127 = l2
		v131 = v19
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L15
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L15
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L15
	} else {
		goto L35
	}
L4:
	;
	m.G0 = v10 + int32(32)
	return
L5:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v133 = F_getRTEPermissionInfo(m, v132, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L15
	} else {
		goto L33
	}
L6:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v20 != int32(2) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if l1 <= int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v27 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 < l1 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v15<<(uint(int32(2))%32))))
	if v36 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v39 = int32(36)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	switch v42 - int32(63) {
	case 0:
		goto L14
	case 1:
		v46 = v39
		goto L13
	default:
		v200 = v36
		goto L1
	}
L13:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46+v41)))
	F_markRTEForSelectPriv(m, l0, v49, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v46 = int32(4)
	goto L13
L15:
	;
	return
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v54 - int32(63) {
	case 0:
		goto L18
	case 1:
		v58 = v39
		goto L17
	default:
		v176 = v36
		goto L2
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v53)))
	v64 = v62 - int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60+v64<<(uint(int32(2))%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	if v69 == int32(0) {
		v127 = v47
		v131 = v68
		goto L5
	} else {
		goto L19
	}
L18:
	;
	v58 = int32(4)
	goto L17
L19:
	;
	v73 = v62
	v75 = v69
	v76 = v64
	goto L20
L20:
	;
	if v75 != int32(2) {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	v127 = v47
	v131 = v123
	goto L5
L22:
	;
	if v73 <= int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v83 == int32(0) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 < v73 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v76<<(uint(int32(2))%32))))
	if v92 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	switch v97 - int32(63) {
	case 0:
		v101 = int32(4)
		goto L27
	case 1:
		goto L28
	default:
		v200 = v92
		goto L1
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101+v96)))
	F_markRTEForSelectPriv(m, l0, v103, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L29
	}
L28:
	;
	v101 = int32(36)
	goto L27
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	switch v109 - int32(63) {
	case 0:
		goto L31
	case 1:
		v113 = int32(36)
		goto L30
	default:
		v176 = v92
		goto L2
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v108)))
	v119 = v117 - int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115+v119<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	if v124 != 0 {
		v73 = v117
		v75 = v124
		v76 = v119
		goto L20
	} else {
		goto L32
	}
L31:
	;
	v113 = int32(4)
	goto L30
L32:
	;
	goto L21
L33:
	;
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v133)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v133)+16)) = v135 | int64(2)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
	v142 = F_bms_add_member(m, v139, v127+int32(7))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = v142
	goto L4
L35:
	;
	F_errmsg_internal(m, int32(_a_F_markRTEForSelectPriv_0), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_markRTEForSelectPriv_1), int32(1112), int32(_a_F_markRTEForSelectPriv_2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L15
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
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v187
	F_errmsg_internal(m, int32(_a_F_markRTEForSelectPriv_3), v10+int32(16))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_markRTEForSelectPriv_1), int32(1144), int32(_a_F_markRTEForSelectPriv_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L15
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
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v211
	F_errmsg_internal(m, int32(_a_F_markRTEForSelectPriv_3), v10)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_markRTEForSelectPriv_1), int32(1129), int32(_a_F_markRTEForSelectPriv_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replace_rte_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	if l0 == int32(0) {
		v22 = int32(0)
		if l5 == v22 {
			v26 = v22
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
			v26 = v25
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v18 != int32(67) {
			v22 = int32(0)
			if l5 == v22 {
				v26 = v22
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
				v26 = v25
			}
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
			v26 = v21
		}
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)) = uint8(v26)
	v31 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1053), v10+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
		if v35 != int32(1) {
			m.G0 = v10 + int32(32)
			return v31
		} else {
			if v31 == int32(0) {
				if l5 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_replace_rte_variables_0), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_replace_rte_variables_1), int32(1475), int32(_a_F_replace_rte_variables_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v47 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v47)
					m.G0 = v10 + int32(32)
					return v31
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				if v40 != int32(67) {
					if l5 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_replace_rte_variables_0), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_replace_rte_variables_1), int32(1475), int32(_a_F_replace_rte_variables_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v47 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v47)
						m.G0 = v10 + int32(32)
						return v31
					}
				} else {
					v43 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v31)+39)) = uint8(v43)
					m.G0 = v10 + int32(32)
					return v31
				}
			}
		}
	}
}
