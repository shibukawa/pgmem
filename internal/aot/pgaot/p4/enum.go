package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_cmp(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_enum_first(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13873(m, l0, int32(_a_F_enum_first_0), int32(460), int32(451), int32(1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_enum_last(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13873(m, l0, int32(_a_F_enum_last_0), int32(489), int32(480), int32(-1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_enum_lt(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v4) >> (uint(int32(31)) % 32))
	}
}
func F_enum_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_enum_cmp_internal(m, v4, v5, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 < int32(0) {
			v12 = v4
		} else {
			v12 = v5
		}
		return v12
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 float32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 float32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v172 float32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 float32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 float32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
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
	v27 = F_palloc(m, int32(512))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
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
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L68
	}
L4:
	;
	return
L5:
	;
	v30 = v19 + int32(-48)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v30, int32(2), int32(3), int32(184), v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = F_table_open(m, int32(3501), int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v42 = int32(1)
	v45 = F_systable_beginscan(m, v39, int32(3503), v42, int32(0), v42, v30)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v47 = F_systable_getnext(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v51 = v47
	v52 = int32(64)
	v53 = v2
	v55 = v27
	goto L13
L11:
	;
	v94 = v2
	v96 = v27
	goto L12
L12:
	;
	F_systable_endscan(m, v45)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L21
	}
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
	v70 = v68 + v69
	if v52 <= v53 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v94 = v88
	v96 = v79
	goto L12
L15:
	;
	v74 = F_repalloc(m, v55, v52<<(uint(int32(4))%32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	v78 = v52
	v79 = v55
	goto L17
L17:
	;
	v82 = v79 + v53<<(uint(int32(3))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v83
	v85 = *(*float32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v82)+4)) = v85
	v88 = v53 + int32(1)
	v89 = F_systable_getnext(m, v45)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v78 = v52 << (uint(int32(1)) % 32)
	v79 = v74
	goto L17
L19:
	;
	if v89 != 0 {
		v51 = v89
		v52 = v78
		v53 = v88
		v55 = v79
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	F_relation_close(m, v39, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_pg_qsort(m, v96, v94, int32(8), int32(1608))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v119 = v94 - int32(1)
	v120 = int32(0)
	if v120 < v119 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v123 = v119
	goto L26
L25:
	;
	v123 = v120
	goto L26
L26:
	;
	v131 = v2
	v133 = int32(1)
	v135 = v2
	v137 = v2
	goto L27
L27:
	;
	if v123 != v131 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v240 = int32(_a_F_load_enum_cache_data_0)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[0]))
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[0])) = v244
	v247 = v94 << (uint(int32(3)) % 32)
	v250 = F_palloc(m, v247+int32(12))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L57
	}
L29:
	;
	v144 = int32(1)
	v146 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v232 = v135
	v234 = v137
	goto L31
L31:
	;
	goto L28
L32:
	;
	v150 = v96 + v131<<(uint(int32(3))%32)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v153 = v131 + int32(1)
	if v94 <= v153 {
		v197 = v144
		v200 = v146
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v211 = base.B2i32(v133 < v197)
	if v133 < v197 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v155 = *(*float32)(unsafe.Add(mBase, uint32(v150)+4))
	v157 = v153
	v160 = v144
	v163 = v146
	v172 = v155
	goto L35
L35:
	;
	v176 = v96 + v157<<(uint(int32(3))%32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v178 = v177 - v151
	if base.Ui32(int32(_a_F_load_enum_cache_data_1)) < base.Ui32(v178) {
		v197 = v160
		v200 = v163
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v197 = v187
	v200 = v188
	goto L33
L37:
	;
	v181 = *(*float32)(unsafe.Add(mBase, uint32(v176)+4))
	if base.F32_lt(v172, v181) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v185 = F_bms_add_member(m, v163, v178)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	v187 = v160
	v188 = v163
	v189 = v172
	goto L40
L40:
	;
	v191 = v157 + int32(1)
	if v191 != v94 {
		v157 = v191
		v160 = v187
		v163 = v188
		v172 = v189
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v187 = v160 + int32(1)
	v188 = v185
	v189 = v181
	goto L40
L42:
	;
	goto L36
L43:
	;
	v212 = v135
	goto L45
L44:
	;
	v212 = v200
	goto L45
L45:
	;
	F_bms_free(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v133 < v197 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v215 = v151
	goto L49
L48:
	;
	v215 = v137
	goto L49
L49:
	;
	if v133 < v197 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v216 = v200
	goto L52
L51:
	;
	v216 = v135
	goto L52
L52:
	;
	if v133 < v197 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v217 = v197
	goto L55
L54:
	;
	v217 = v133
	goto L55
L55:
	;
	if v217 < v94+(v131^int32(-1)) {
		v131 = v153
		v133 = v217
		v135 = v216
		v137 = v215
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v232 = v216
	v234 = v215
	goto L31
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v234
	v253 = F_bms_copy(m, v232)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v253
	if v247 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v250+int32(12), v96, v247)
	goto L61
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_enum_cache_data[0])) = v241
	F_pfree(m, v96)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_bms_free(m, v232)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	if v266 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v250
	m.G0 = v21 - int32(-64)
	return
L67:
	;
	goto L66
L68:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = F_format_type_be(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v281
	F_errmsg(m, int32(_a_F_load_enum_cache_data_2), v21)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_load_enum_cache_data_3), int32(2758), int32(_a_F_load_enum_cache_data_4))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
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
