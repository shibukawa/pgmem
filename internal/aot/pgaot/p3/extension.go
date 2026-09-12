package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ApplyExtensionUpdates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	v8 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	if l3 == v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(128)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v26 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = l2
	v45 = v8
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v45<<(uint(int32(2))%32))))
	v54 = F_palloc(m, int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+40)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v66
	F_parse_extension_control_file(m, v54, v52)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v72 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_ScanKeyInit(m, v22+int32(80), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v82 = int32(1)
	v87 = F_systable_beginscan(m, v72, int32(3080), v82, int32(0), v82, v22+int32(80))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v227 = int32(3079)
	v230 = F_deleteDependencyRecordsForClass(m, v227, l0, v227, int32(110))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L42
	}
L12:
	;
	v89 = F_systable_getnext(m, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v89 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91+v92)+72))
	v95 = F_get_namespace_name(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L39
	}
L17:
	;
	v97 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v97
	v100 = v22 - int32(-64)
	*(*int64)(unsafe.Add(mBase, uint32(v100))) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = int64(4294967296)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v111
	v113 = F_cstring_to_text(m, v52)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)) = uint8(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v113
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	v125 = F_heap_modify_tuple(m, v89, v118, v22+int32(48), v22+int32(40), v22+int32(32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	F_CatalogTupleUpdate(m, v72, v125+int32(4), v125)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_systable_endscan(m, v87)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	F_sequence_close(m, v72, int32(3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
	if v136 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v156 = v137
	v157 = v137
	v162 = v137
	goto L28
L24:
	;
	v137 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v137 < v140 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v146 = int32(0)
	v216 = v146
	v222 = v146
	goto L11
L27:
	;
	goto L26
L28:
	;
	v167 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169+v157<<(uint(int32(2))%32))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v176 = F_get_required_extension(m, v173, v174, l4, l5, v167, l6)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L30
	}
L29:
	;
	v216 = v187
	v222 = v189
	goto L11
L30:
	;
	v178 = F_SearchSysCache1(m, int32(28), v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	if v178 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+22)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180+v181)+72))
	F_ReleaseCatCache(m, v178)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	v186 = v167
	goto L34
L34:
	;
	v187 = F_lappend_oid(m, v156, v176)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v186 = v183
	goto L34
L36:
	;
	v189 = F_lappend_oid(m, v162, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v192 = v157 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v192 < v193 {
		v156 = v187
		v157 = v192
		v162 = v189
		goto L28
	} else {
		goto L38
	}
L38:
	;
	goto L29
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
	F_errmsg_internal(m, int32(46900), v22)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(496872), int32(3605), int32(160657))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v232 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(3079)
	if v216 == v232 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[442]))
	if v303 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	v239 = int32(0)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v240 <= v239 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v252 = v239
	goto L46
L46:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+v252<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3079)
	F_recordDependencyOn(m, v22+int32(20), v22+int32(8), int32(110))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L48
	}
L47:
	;
	goto L43
L48:
	;
	v280 = v252 + int32(1)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v280 < v281 {
		v252 = v280
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v305 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3079), l0, v305, v305, v305)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_execute_extension_script(m, l0, v54, v31, v52, v222, v95)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v313 = v45 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v313 < v314 {
		v31 = v52
		v45 = v313
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L5
}
