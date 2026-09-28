package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_clear_relation_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(128)
	m.G0 = v5
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)) = uint8(v2)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v9
	v13 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+26)) = uint16(v13)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+32)) = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+40)) = uint8(v17)
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+48)) = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+120)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+112)) = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+104)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+96)) = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+88)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+80)) = int64(-1082130432)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+72)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+64)) = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+56)) = uint8(v21)
	v41 = F_relation_statistics_update(m, v5+int32(8))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return int64(0)
	} else {
		m.G0 = v5 + int32(128)
		return int64(0)
	}
}
func F_pg_relation_filepath(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int64
	_ = v220
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_SearchSysCache1(m, int32(57), v15&int64(4294967295))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v220
L2:
	;
	return int64(0)
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
	v220 = int64(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
	switch v30 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L9
	default:
		goto L8
	}
L7:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
	switch v155 - int32(112) {
	case 0, 5:
		v197 = int32(-1)
		goto L51
	default:
		goto L53
	case 4:
		goto L54
	}
L8:
	;
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L50
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[1]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	if v38 != 0 {
		v153 = v38
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v39 = base.I32_wrap_i64(v15)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+117)))
	v41 = int32(0)
	if v40 == v41 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v143 != 0 {
		v153 = v143
		goto L7
	} else {
		goto L49
	}
L12:
	;
	goto L11
L13:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v143 = v138
	goto L12
L14:
	;
	v46 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[2]))
	if v46 < v48 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v89 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[3]))
	if v89 < v91 {
		goto L35
	} else {
		goto L36
	}
L17:
	;
	v52 = v46
	goto L20
L18:
	;
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[4]))
	if v71 <= int32(0) {
		v143 = v41
		goto L12
	} else {
		goto L26
	}
L20:
	;
	v57 = v52 << (uint(int32(3)) % 32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_pg_relation_filepath[5])))
	if v58 == v39 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v137 = v57 + int32(_a_F_pg_relation_filepath_0)
	goto L13
L23:
	;
	goto L24
L24:
	;
	v63 = v52 + int32(1)
	if v63 != v48 {
		v52 = v63
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v76 = int32(0)
	goto L27
L27:
	;
	v81 = v76 << (uint(int32(3)) % 32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_pg_relation_filepath[6])))
	if v82 != v39 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v137 = v81 + int32(_a_F_pg_relation_filepath_1)
	goto L13
L29:
	;
	v85 = v76 + int32(1)
	if v71 != v85 {
		v76 = v85
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v143 = v41
	goto L12
L33:
	;
	v119 = int32(0)
	goto L43
L34:
	;
	v137 = v100 + int32(_a_F_pg_relation_filepath_2)
	goto L13
L35:
	;
	v95 = v89
	goto L38
L36:
	;
	goto L37
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[7]))
	if v112 <= int32(0) {
		v143 = v41
		goto L12
	} else {
		goto L42
	}
L38:
	;
	v100 = v95 << (uint(int32(3)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_pg_relation_filepath[8])))
	if v39 == v101 {
		goto L34
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	v104 = v95 + int32(1)
	if v104 != v91 {
		v95 = v104
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L33
L43:
	;
	v124 = v119 << (uint(int32(3)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+uint32(_c_F_pg_relation_filepath[9])))
	if v125 != v39 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v137 = v124 + int32(_a_F_pg_relation_filepath_3)
	goto L13
L45:
	;
	v128 = v119 + int32(1)
	if v112 != v128 {
		v119 = v128
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v143 = v41
	goto L12
L49:
	;
	goto L8
L50:
	;
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v150)
	v220 = int64(0)
	goto L1
L51:
	;
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L70
	}
L52:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v195 = F_GetTempNamespaceProcNumber(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L69
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L66
	}
L54:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[10]))
	if v162 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if v170 == int32(0) {
		goto L52
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v163 = int32(1)
	if v158 == v162 {
		v170 = v163
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v170 = int32(0)
	goto L56
L60:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[11]))
	if v166 == v158 {
		v170 = v163
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[12]))
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[13]))
	if v176 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v179 = v174
	goto L65
L64:
	;
	v179 = v176
	goto L65
L65:
	;
	v197 = v179
	goto L51
L66:
	;
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v184
	F_errmsg_internal(m, int32(_a_F_pg_relation_filepath_4), v12)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_pg_relation_filepath_5), int32(1036), int32(_a_F_pg_relation_filepath_6))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
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
	v197 = v195
	goto L51
L70:
	;
	v201 = v12 + int32(8)
	if v35 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v203 = v35
	goto L73
L72:
	;
	v203 = v34
	goto L73
L73:
	;
	if v203 != int32(1664) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v206 = v37
	goto L76
L75:
	;
	v206 = int32(0)
	goto L76
L76:
	;
	F_GetRelationPath(m, v201, v206, v203, v153, v197, int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v210 = F_cstring_to_text(m, v201)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v220 = base.I64_extend_i32_u(v210)
	goto L1
}
