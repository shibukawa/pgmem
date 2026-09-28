package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_dependencies_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 == v2 {
		v153 = v2
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L18
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L18
	} else {
		goto L69
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L18
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L18
	} else {
		goto L63
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L18
	} else {
		goto L47
	}
L6:
	;
	m.G0 = v12 - int32(-64)
	return v153
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui32(v43) <= base.Ui32(int32(11)) {
		goto L5
	} else {
		goto L17
	}
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = int32(1)
	if v16&v31 != 0 {
		v43 = int32(base.Ui32(v16)>>(uint(v31)%32)) - v31
		goto L8
	} else {
		goto L16
	}
L12:
	;
	if v19 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v30 = int32(16)
	goto L15
L14:
	;
	v30 = int32(0)
	goto L15
L15:
	;
	v43 = v30
	goto L8
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L17:
	;
	v47 = F_palloc0(m, int32(12))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v51 = int32(1)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v53&v51 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v56 = v51
	goto L22
L21:
	;
	v56 = int32(4)
	goto L22
L22:
	;
	v57 = l0 + v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v62
	if v58 != int32(-1269523924) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v60 != int32(1) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	if v62 == int32(0) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v73 = v62*int32(14) + int32(12)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v74 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if base.Ui32(v103) < base.Ui32(v73) {
		goto L1
	} else {
		goto L37
	}
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v80 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v91 = int32(1)
	if v74&v91 != 0 {
		v103 = int32(base.Ui32(v74)>>(uint(v91)%32)) - v91
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v83 = int32(16)
	goto L32
L31:
	;
	v83 = int32(0)
	goto L32
L32:
	;
	if base.Ui32((v80-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = int32(4)
	goto L35
L34:
	;
	v90 = v83
	goto L35
L35:
	;
	v103 = v90
	goto L26
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v109 = F_repalloc(m, v47, v62<<(uint(int32(2))%32)+int32(12))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if v111 == int32(0) {
		v153 = v109
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v114 = int32(12)
	v119 = v57 + v114
	v124 = int32(0)
	goto L40
L40:
	;
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v119)))
	v129 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+8)))
	v131 = v129 << (uint(int32(1)) % 32)
	v134 = F_palloc0(m, v131+int32(10))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L42
	}
L41:
	;
	v153 = v109
	goto L6
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+8)) = uint16(v129)
	*(*float64)(unsafe.Add(mBase, uint32(v134))) = v128
	v139 = v119 + int32(10)
	if v131 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	base.MemoryCopy(m, v134+int32(10), v139, v131)
	goto L45
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109+v114+v124<<(uint(int32(2))%32)))) = v134
	v149 = v124 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if base.Ui32(v149) < base.Ui32(v150) {
		v119 = v139 + v131
		v124 = v149
		goto L40
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v171 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v200
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_0), v12)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L18
	} else {
		goto L61
	}
L49:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v177 == int32(18) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if v171&int32(1) != 0 {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v180 = int32(16)
	goto L54
L53:
	;
	v180 = int32(0)
	goto L54
L54:
	;
	if base.Ui32((v177-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v187 = int32(4)
	goto L57
L56:
	;
	v187 = v180
	goto L57
L57:
	;
	v200 = v187
	goto L48
L58:
	;
	v190 = int32(1)
	v200 = int32(base.Ui32(v171)>>(uint(v190)%32)) - v190
	goto L48
L59:
	;
	goto L60
L60:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v200 = int32(base.Ui32(v194)>>(uint(int32(2))%32)) - int32(4)
	goto L48
L61:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(504), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(-1269523924)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v216
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_3), v10+int32(-16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(522), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v234
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_4), v10+int32(-32))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(526), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_5), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(529), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v266 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v295
	F_errmsg_internal(m, int32(_a_F_statext_dependencies_deserialize_6), v10+int32(-48))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L18
	} else {
		goto L86
	}
L74:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v272 == int32(18) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v266&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v275 = int32(16)
	goto L79
L78:
	;
	v275 = int32(0)
	goto L79
L79:
	;
	if base.Ui32((v272-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v282 = int32(4)
	goto L82
L81:
	;
	v282 = v275
	goto L82
L82:
	;
	v295 = v282
	goto L73
L83:
	;
	v285 = int32(1)
	v295 = int32(base.Ui32(v266)>>(uint(v285)%32)) - v285
	goto L73
L84:
	;
	goto L85
L85:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v295 = int32(base.Ui32(v289)>>(uint(int32(2))%32)) - int32(4)
	goto L73
L86:
	;
	F_errfinish(m, int32(_a_F_statext_dependencies_deserialize_1), int32(536), int32(_a_F_statext_dependencies_deserialize_2))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_statext_mcv_load(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v12 = F_SearchSysCache2(m, int32(62), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v20 = F_SysCacheGetAttr(m, int32(62), v12, int32(5), v7+int32(31))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)))
				if v22 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(109)
						F_errmsg_internal(m, int32(_a_F_statext_mcv_load_0), v7+int32(16))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_statext_mcv_load_1), int32(573), int32(_a_F_statext_mcv_load_2))
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
					v26 = F_pg_detoast_datum(m, base.I32_wrap_i64(v20))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = F_statext_mcv_deserialize(m, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v28
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_statext_mcv_load_3), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_statext_mcv_load_1), int32(565), int32(_a_F_statext_mcv_load_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_statext_ndistinct_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 == v2 {
		v93 = int32(16)
	} else {
		v15 = v10 & int32(3)
		v16 = int32(16)
		if base.Ui32(int32(4)) <= base.Ui32(v10) {
			v23 = v2
			v24 = v16
			v27 = v2
			for {
				v30 = int32(4)
				v32 = l0 + v23<<(uint(v30)%32)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
				v34 = int32(1)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
				v50 = v24 + v33<<(uint(v34)%32) + v37<<(uint(v34)%32) + v41<<(uint(v34)%32) + v45<<(uint(v34)%32) + int32(48)
				v52 = v23 + v30
				v54 = v27 + v30
				if v54 != v10&int32(-4) {
					v23 = v52
					v24 = v50
					v27 = v54
					continue
				} else {
					break
				}
				break
			}
			if v15 == int32(0) {
				v93 = v50
			} else {
				v60 = v52
				v61 = v50
				v69 = v60
				v70 = v61
				v74 = v2
				for {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0+v69<<(uint(int32(4))%32))+24))
					v80 = int32(1)
					v84 = v70 + v79<<(uint(v80)%32) + int32(12)
					v88 = v74 + v80
					if v88 != v15 {
						v69 = v69 + v80
						v70 = v84
						v74 = v88
						continue
					} else {
						break
					}
					break
				}
				v93 = v84
			}
		} else {
			v60 = v2
			v61 = v16
			v69 = v60
			v70 = v61
			v74 = v2
			for {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0+v69<<(uint(int32(4))%32))+24))
				v80 = int32(1)
				v84 = v70 + v79<<(uint(v80)%32) + int32(12)
				v88 = v74 + v80
				if v88 != v15 {
					v69 = v69 + v80
					v70 = v84
					v74 = v88
					continue
				} else {
					break
				}
				break
			}
			v93 = v84
		}
	}
	v99 = F_palloc(m, v93)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v99))) = v93 << (uint(int32(2)) % 32)
		v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v106
		v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v108
		v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v110
		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v112 != 0 {
			v113 = int32(16)
			v120 = v99 + v113
			v121 = int32(0)
			for {
				v129 = l0 + v113 + v121<<(uint(int32(4))%32)
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
				v131 = *(*float64)(unsafe.Add(mBase, uint32(v129)))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v132
				*(*float64)(unsafe.Add(mBase, uint32(v120))) = v131
				v136 = v120 + int32(12)
				v138 = v132 << (uint(int32(1)) % 32)
				if v138 != 0 {
					base.MemoryCopy(m, v136, v130, v138)
				} else {
				}
				v142 = v121 + int32(1)
				v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(v142) < base.Ui32(v143) {
					v120 = v136 + v138
					v121 = v142
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		return v99
	}
}
