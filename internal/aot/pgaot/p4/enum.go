package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v4)
	}
}
func F_enum_first(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14259(m, l0, int32(_a_F_enum_first_0), int32(460), int32(451), int32(1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_enum_last(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14259(m, l0, int32(_a_F_enum_last_0), int32(489), int32(480), int32(-1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_enum_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(int32(base.Ui32(v4) >> (uint(int32(31)) % 32)))
	}
}
func F_enum_smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_enum_cmp_internal(m, base.I32_wrap_i64(v4), base.I32_wrap_i64(v5), l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		if v8 < int32(0) {
			v14 = v4
		} else {
			v14 = v5
		}
		return v14 & int64(4294967295)
	}
}
func F_load_enum_cache_data(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 float32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 float32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 float32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 float32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 float32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v23 == int32(101) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = int32(64)
	v29 = F_palloc_mul(m, int32(8), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L68
	}
L4:
	;
	return
L5:
	;
	v32 = v19 + int32(-56)
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	F_ScanKeyInit(m, v32, int32(2), int32(3), int32(184), v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v41 = F_table_open(m, int32(3501), int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v44 = int32(1)
	v47 = F_systable_beginscan(m, v41, int32(3503), v44, int32(0), v44, v32)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v49 = F_systable_getnext(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = v49
	v53 = v26
	v54 = v2
	v56 = v29
	goto L13
L11:
	;
	v95 = v2
	v97 = v29
	goto L12
L12:
	;
	F_systable_endscan(m, v47)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L21
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
	v71 = v69 + v70
	if v53 <= v54 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v95 = v89
	v97 = v80
	goto L12
L15:
	;
	v75 = F_repalloc(m, v56, v53<<(uint(int32(4))%32))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	v79 = v53
	v80 = v56
	goto L17
L17:
	;
	v83 = v80 + v54<<(uint(int32(3))%32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v84
	v86 = *(*float32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v83)+4)) = v86
	v89 = v54 + int32(1)
	v90 = F_systable_getnext(m, v47)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v79 = v53 << (uint(int32(1)) % 32)
	v80 = v75
	goto L17
L19:
	;
	if v90 != 0 {
		v52 = v90
		v53 = v79
		v54 = v89
		v56 = v80
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	F_relation_close(m, v41, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_pg_qsort(m, v97, v95, int32(8), int32(1822))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v120 = v95 - int32(1)
	v121 = int32(0)
	if v121 < v120 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v124 = v120
	goto L26
L25:
	;
	v124 = v121
	goto L26
L26:
	;
	v132 = v2
	v134 = int32(1)
	v136 = v2
	v138 = v2
	goto L27
L27:
	;
	if v124 != v132 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v241 = int32(_a_F_load_enum_cache_data_0)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[0]))
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[0])) = v245
	v248 = v95 << (uint(int32(3)) % 32)
	v251 = F_palloc(m, v248+int32(12))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L57
	}
L29:
	;
	v145 = int32(1)
	v147 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v233 = v136
	v235 = v138
	goto L31
L31:
	;
	goto L28
L32:
	;
	v151 = v97 + v132<<(uint(int32(3))%32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v154 = v132 + int32(1)
	if v95 <= v154 {
		v198 = v145
		v201 = v147
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v212 = base.B2i32(v134 < v198)
	if v134 < v198 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v156 = *(*float32)(unsafe.Add(mBase, uint32(v151)+4))
	v158 = v154
	v161 = v145
	v164 = v147
	v173 = v156
	goto L35
L35:
	;
	v177 = v97 + v158<<(uint(int32(3))%32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v179 = v178 - v152
	if base.Ui32(int32(_a_F_load_enum_cache_data_1)) < base.Ui32(v179) {
		v198 = v161
		v201 = v164
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v198 = v188
	v201 = v189
	goto L33
L37:
	;
	v182 = *(*float32)(unsafe.Add(mBase, uint32(v177)+4))
	if base.F32_lt(v173, v182) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v186 = F_bms_add_member(m, v164, v179)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	v188 = v161
	v189 = v164
	v190 = v173
	goto L40
L40:
	;
	v192 = v158 + int32(1)
	if v192 != v95 {
		v158 = v192
		v161 = v188
		v164 = v189
		v173 = v190
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v188 = v161 + int32(1)
	v189 = v186
	v190 = v182
	goto L40
L42:
	;
	goto L36
L43:
	;
	v213 = v136
	goto L45
L44:
	;
	v213 = v201
	goto L45
L45:
	;
	F_bms_free(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v134 < v198 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v216 = v152
	goto L49
L48:
	;
	v216 = v138
	goto L49
L49:
	;
	if v134 < v198 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v217 = v201
	goto L52
L51:
	;
	v217 = v136
	goto L52
L52:
	;
	if v134 < v198 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v218 = v198
	goto L55
L54:
	;
	v218 = v134
	goto L55
L55:
	;
	if v218 < v95+(v132^int32(-1)) {
		v132 = v154
		v134 = v218
		v136 = v217
		v138 = v216
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v233 = v217
	v235 = v216
	goto L31
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v235
	v254 = F_bms_copy(m, v233)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+8)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v254
	if v248 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v251+int32(12), v97, v248)
	goto L61
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[0])) = v242
	F_pfree(m, v97)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_bms_free(m, v233)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	if v267 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v251
	m.G0 = v21 - int32(-64)
	return
L67:
	;
	goto L66
L68:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v282 = F_format_type_be(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v282
	F_errmsg(m, int32(_a_F_load_enum_cache_data_2), v21)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_load_enum_cache_data_3), int32(2785), int32(_a_F_load_enum_cache_data_4))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
