package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetRecoveryPauseState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+96)) = int32(1)
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[248]))
		F_s_lock(m, v9+int32(96), int32(470996), int32(3096), int32(338193))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[248]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
			return v23
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[248]))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
		return v23
	}
}
func F_check_recovery_target_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8&int32(3) == int32(0) {
		v32 = v8
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v67 = base.B2i32(base.Ui32(v65) < base.Ui32(int32(64)))
	if v67 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v65 = v57 - v8
	goto L1
L3:
	;
	v36 = v32
	goto L12
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v65 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v21 = v8
	goto L8
L8:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v57 = v25
	goto L2
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v51 = v36
	goto L15
L14:
	;
	goto L13
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v57 = v51
	goto L2
L17:
	;
	goto L16
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	*(*int32)(unsafe.Add(mBase, _consts[250])) = v71
	goto L21
L19:
	;
	goto L20
L20:
	;
	m.G0 = v6 + int32(16)
	return v67
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(361176)
	v80 = F_format_elog_string(m, int32(621104), v6)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[251])) = v80
	goto L20
}
func F_check_recovery_target_time(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
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
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v228 int32
	_ = v228
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	v4 = m.G0
	v6 = v4 - int32(464)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(464)
	return v269
L2:
	;
	v269 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = int32(0)
	v14 = int32(28995)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[253])))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v18 == v13 {
		v37 = v17
		v38 = v18
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v38-v37 == int32(0) {
		v269 = v13
		goto L1
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v17 != v18 {
		v37 = v17
		v38 = v18
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = v8
	v23 = v14
	goto L9
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v26
		v38 = v27
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v37 = v26
	v38 = v27
	goto L6
L11:
	;
	v30 = int32(1)
	if v26 == v27 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = int32(25182)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[254])))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v66-v65 == int32(0) {
		v269 = v13
		goto L1
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = v8
	v51 = v42
	goto L18
L18:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v65 = v54
	v66 = v55
	goto L15
L20:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v70 = int32(28024)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[255])))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v74 == int32(0) {
		v93 = v73
		v94 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v94-v93 == int32(0) {
		v269 = v13
		goto L1
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	if v73 != v74 {
		v93 = v73
		v94 = v74
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v78 = v8
	v79 = v70
	goto L27
L27:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v82
		v94 = v83
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v93 = v82
	v94 = v83
	goto L24
L29:
	;
	v86 = int32(1)
	if v82 == v83 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v98 = int32(25172)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[256])))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v102 == int32(0) {
		v121 = v101
		v122 = v102
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v122-v121 == int32(0) {
		v269 = v13
		goto L1
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	if v101 != v102 {
		v121 = v101
		v122 = v102
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v106 = v8
	v107 = v98
	goto L36
L36:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v110
		v122 = v111
		goto L33
	} else {
		goto L38
	}
L37:
	;
	v121 = v110
	v122 = v111
	goto L33
L38:
	;
	v114 = int32(1)
	if v110 == v111 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v135 = F_ParseDateTime(m, v8, v6+int32(32), int32(153), v6+int32(304), v6+int32(192), v6+int32(404))
	mBase = m.M
	if v135 != 0 {
		v269 = v13
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v6)+404))
	v151 = F_DecodeDateTime(m, v6+int32(304), v6+int32(192), v140, v6+int32(408), v6+int32(416), v6+int32(460), v6+int32(412), v6+int32(24))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return int32(0)
L43:
	;
	if v151 != 0 {
		v269 = v13
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v6)+408))
	if v155 != int32(2) {
		v269 = v13
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v159 = v6 + int32(416)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v6)+460))
	v162 = v6 + int32(412)
	v163 = int32(16)
	v164 = v6 + v163
	v171 = m.G0
	v173 = v171 - v163
	m.G0 = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	if v175 <= int32(-4713) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	if v252 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L47:
	;
	m.G0 = v173 + int32(16)
	goto L46
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = int64(0)
	v252 = int32(-1)
	goto L47
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v193 = F_date2j(m, v175, v191, v192)
	mBase = m.M
	v196 = base.I64_extend_i32_s(v193 - int32(2451545))
	v197 = int64(63)
	F___multi3(m, v173, v196, v196>>(uint(v197)%64), int64(86400000000), int64(0))
	mBase = m.M
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v173)+8))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v173)))
	if v202 != v203>>(uint(v197)%64) {
		goto L48
	} else {
		goto L60
	}
L50:
	;
	if v175 != int32(-4713) {
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v175 <= int32(5874897) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	if int32(10) < v180 {
		v191 = v180
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L48
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v191 = v185
	goto L49
L56:
	;
	goto L57
L57:
	;
	if v175 != int32(5874898) {
		goto L48
	} else {
		goto L58
	}
L58:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	if int32(5) < v188 {
		goto L48
	} else {
		goto L59
	}
L59:
	;
	v191 = v188
	goto L49
L60:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v211 = int32(60)
	v220 = base.I64_extend_i32_s(v160) + base.I64_extend_i32_s(v208+(v209+v210*v211)*v211)*int64(1000000)
	v221 = v203 + v220
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v221
	if base.B2i32(v220 < int64(0))^base.B2i32(v221 < v203) != 0 {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	if v162 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v233 = base.I64_extend_i32_s(int32(0)-v228)*int64(-1000000) + v221
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v233
	v235 = v233
	goto L64
L63:
	;
	v235 = v221
	goto L64
L64:
	;
	if base.Ui64(v235+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v252 = int32(0)
		goto L47
	} else {
		goto L65
	}
L65:
	;
	goto L48
L66:
	;
	v269 = int32(1)
	goto L1
L67:
	;
	goto L68
L68:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	*(*int32)(unsafe.Add(mBase, _consts[250])) = v260
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v266 = F_format_elog_string(m, int32(628223), v6)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L42
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[251])) = v266
	v269 = v13
	goto L1
}
func F_recovery_create_dbdir(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v4 = m.G0
	v6 = v4 - int32(160)
	m.G0 = v6
	v12 = F___fstatat(m, int32(-100), l0, v6-int32(-64), int32(0))
	mBase = m.M
	goto L4
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L16
	} else {
		goto L58
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L16
	} else {
		goto L55
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L16
	} else {
		goto L52
	}
L4:
	;
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	m.G0 = v6 + int32(160)
	return
L8:
	;
	v14 = F_strstr(m, l0, int32(532262))
	mBase = m.M
	if v14 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[154])))
	if v18 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v27 = int32(14)
	goto L14
L13:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[382])))
	if v23 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v29 = F_errstart(m, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v27 = int32(19)
	goto L14
L16:
	;
	return
L17:
	;
	if v29 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	F_errmsg_internal(m, int32(189173), v6+int32(32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v48 = m.G0
	v50 = v48 - int32(96)
	m.G0 = v50
	v53 = F_umask(m, int32(0))
	mBase = m.M
	v56 = F_umask(m, v53&int32(-193))
	mBase = m.M
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v63 = l0 + base.B2i32(v57 == int32(47))
	goto L25
L21:
	;
	F_errfinish(m, int32(472768), int32(3298), int32(203338))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v111 != 0 {
		goto L1
	} else {
		goto L51
	}
L24:
	;
	v112 = F_umask(m, v53)
	mBase = m.M
	m.G0 = v50 + int32(96)
	goto L23
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != int32(47) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v111 = v96 >> (uint(int32(31)) % 32)
	goto L24
L27:
	;
	goto L26
L28:
	;
	v63 = v63 + int32(1)
	goto L25
L29:
	;
	v79 = F_stat(m, l0, v50)
	mBase = m.M
	if v79 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v76 = F_umask(m, v53)
	mBase = m.M
	v78 = int32(0)
	goto L29
L31:
	;
	if v67 != 0 {
		goto L28
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v72)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v75 != 0 {
		v78 = int32(1)
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v70)
	goto L30
L35:
	;
	goto L30
L36:
	;
	v103 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v103)
	goto L28
L37:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v82&int32(61440) != int32(16384) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v78 != 0 {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	if v78 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v78 != 0 {
		goto L36
	} else {
		goto L46
	}
L43:
	;
	v90 = int32(54)
	goto L45
L44:
	;
	v90 = int32(20)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v90
	v111 = int32(-1)
	goto L24
L46:
	;
	v111 = int32(0)
	goto L24
L47:
	;
	v95 = int32(511)
	goto L49
L48:
	;
	v95 = v43
	goto L49
L49:
	;
	v96 = F_mkdir(m, l0, v95)
	mBase = m.M
	v97 = int32(0)
	if v78&base.B2i32(v97 <= v96) == v97 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	goto L36
L51:
	;
	goto L7
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg_internal(m, int32(189204), v6)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(472768), int32(3291), int32(203338))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
	F_errmsg(m, int32(652273), v6+int32(48))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(472768), int32(3295), int32(203338))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg(m, int32(283092), v6+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(472768), int32(3302), int32(203338))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
