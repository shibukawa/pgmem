package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimezoneName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0&int32(3) == v4 {
		v34 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v69 = F_downcase_truncate_identifier(m, l0, v67, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v67 = v59 - l0
	goto L1
L3:
	;
	v38 = v34
	goto L12
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v23 = l0
	goto L8
L8:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v59 = v27
	goto L2
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v53 = v38
	goto L15
L14:
	;
	goto L13
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v53
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v77 = F_DecodeTimezoneAbbrev(m, v4, v69, v8+int32(12), l1, l2, v8+int32(4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = int32(0)
	F_DateTimeParseError(m, v77, v8+int32(4), v81, v81, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if base.Ui32(v86-int32(5)) < base.Ui32(int32(2)) {
		v100 = v4
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L18
	} else {
		goto L33
	}
L26:
	;
	m.G0 = v8 + int32(16)
	return v100
L27:
	;
	if v86 == int32(7) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = int32(1)
	goto L26
L29:
	;
	goto L30
L30:
	;
	v94 = F_pg_tzset(m, l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v94
	if v94 == int32(0) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v100 = int32(2)
	goto L26
L33:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(_a_F_DecodeTimezoneName_0), v8)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_DecodeTimezoneName_1), int32(3330), int32(_a_F_DecodeTimezoneName_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_timezone(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 float64
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = v12
	v18 = int32(_a_F_check_timezone_0)
	v19 = int32(8)
	goto L2
L1:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v64 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v64 = int32(0)
	goto L1
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 == v23 {
		v45 = v22
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v47 = int32(1)
	if v45 != 0 {
		v17 = v17 + v47
		v18 = v18 + v47
		v19 = v19 - v47
		goto L2
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = v22 | int32(32)
	goto L11
L10:
	;
	v33 = v22
	goto L11
L11:
	;
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v23 | int32(32)
	goto L14
L13:
	;
	v42 = v23
	goto L14
L14:
	;
	if v33 == v42 {
		v45 = v33
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v64 = v33 - v42
	goto L1
L16:
	;
	goto L6
L17:
	;
	m.G0 = v10 + int32(16)
	return v238
L18:
	;
	v227 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L30
	} else {
		goto L77
	}
L19:
	;
	if v203 != 0 {
		v221 = v203
		goto L18
	} else {
		goto L74
	}
L20:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
	v195 = base.I64_div_s(v193, int64(-1000000))
	v197 = F_pg_tzset_offset(m, base.I32_wrap_i64(v195))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L30
	} else {
		goto L72
	}
L21:
	;
	v70 = v65 + int32(8)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v147 = F_strtod(m, v65, v10+int32(12))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L30
	} else {
		goto L53
	}
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui32(v77-int32(9)) < base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v90 = F_pstrdup(m, v70+int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	goto L25
L27:
	;
	v70 = v70 + int32(1)
	goto L24
L28:
	;
	v82 = int32(0)
	switch v77 - int32(32) {
	case 0:
		goto L27
	default:
		v238 = v82
		goto L17
	case 7:
		goto L26
	}
L29:
	;
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v107)
	v113 = F_DirectFunctionCall3Coll(m, int32(584), v107, v90, v107, int32(-1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L30
	} else {
		goto L41
	}
L30:
	;
	return int32(0)
L31:
	;
	v94 = int32(39)
	v95 = F___strchrnul(m, v90, v94)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v97 == v94 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v101 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v101 = v95
	goto L35
L34:
	;
	v101 = int32(0)
	goto L35
L35:
	;
	goto L32
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v102 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v90)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L30
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v238 = v82
	goto L17
L41:
	;
	F_pfree(m, v90)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v117 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v119
	goto L46
L44:
	;
	goto L45
L45:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	if v130 == int32(0) {
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v125 = F_format_elog_string(m, int32(_a_F_check_timezone_1), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v125
	F_pfree(m, v113)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v238 = v82
	goto L17
L49:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v134
	goto L50
L50:
	;
	v140 = F_format_elog_string(m, int32(_a_F_check_timezone_2), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L30
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v140
	F_pfree(m, v113)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L30
	} else {
		goto L52
	}
L52:
	;
	v238 = v82
	goto L17
L53:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v149 == v150 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v164 = F_pg_tzset(m, v150)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L30
	} else {
		goto L62
	}
L55:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v152 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v154 = base.F64_mul(v147, float64(-3600))
	if base.F64_lt(base.F64_abs(v154), float64(2.147483648e+09)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v158 = base.I32_trunc_f64_s(v154)
	v159 = F_pg_tzset_offset(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L30
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v162 = F_pg_tzset_offset(m, int32(-2147483648))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L30
	} else {
		goto L61
	}
L60:
	;
	v203 = v159
	goto L19
L61:
	;
	v203 = v162
	goto L19
L62:
	;
	if v164 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v238 = int32(0)
	goto L17
L64:
	;
	goto L65
L65:
	;
	v169 = F_pg_tz_acceptable(m, v164)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L30
	} else {
		goto L66
	}
L66:
	;
	if v169 != 0 {
		v221 = v164
		goto L18
	} else {
		goto L67
	}
L67:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v173
	goto L68
L68:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v176
	v180 = F_format_elog_string(m, int32(_a_F_check_timezone_3), v10)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[3])) = v180
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v184
	goto L70
L70:
	;
	v190 = F_format_elog_string(m, int32(_a_F_check_timezone_4), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v190
	v238 = int32(0)
	goto L17
L72:
	;
	F_pfree(m, v113)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L30
	} else {
		goto L73
	}
L73:
	;
	v203 = v197
	goto L19
L74:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_check_timezone[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[1])) = v210
	goto L75
L75:
	;
	v216 = F_format_elog_string(m, int32(_a_F_check_timezone_5), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L30
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_timezone[2])) = v216
	v238 = int32(0)
	goto L17
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v227
	if v227 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v238 = int32(0)
	goto L17
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v221
	v238 = int32(1)
	goto L17
}
