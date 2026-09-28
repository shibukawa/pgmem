package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_vacuum_buffer_usage_limit(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(0)
	v15 = base.B2i32(v8 == v9) | base.B2i32(base.Ui32(v8-int32(128)) < base.Ui32(int32(16777089)))
	if v15 == v9 {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_check_vacuum_buffer_usage_limit[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_vacuum_buffer_usage_limit[1])) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(72057594037928064)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_check_vacuum_buffer_usage_limit_0)
		v28 = F_format_elog_string(m, int32(_a_F_check_vacuum_buffer_usage_limit_1), v6)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_vacuum_buffer_usage_limit[2])) = v28
			m.G0 = v6 + int32(16)
			return v15
		}
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_vacuumRedirectAndPlaceholder(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(1648)
	m.G0 = v21
	if l2 < v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v41) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(l2^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L1
L3:
	;
	goto L4
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[1]))
	v40 = v34 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v49 = int32(base.Ui32(v41+int32(_a_F_vacuumRedirectAndPlaceholder_0)) >> (uint(int32(2)) % 32))
	goto L7
L6:
	;
	v49 = int32(0)
	goto L7
L7:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[2]))
	if v52 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v87
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v87)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v85)
	v92 = F_GlobalVisHorizonKindForRel(m, l1)
	mBase = m.M
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92<<(uint(int32(2))%32))+uint32(_c_F_vacuumRedirectAndPlaceholder[3])))
	goto L24
L9:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[4])))
	if v56&int32(1) == int32(0) {
		v85 = v4
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+118)))
	if v62 != int32(112) {
		v85 = v4
		goto L8
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	if int32(0) < v52 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L18
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v67 != 0 {
		v85 = v4
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v68 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v85 = v4
	goto L8
L18:
	;
	if base.Ui32(v72) < base.Ui32(int32(_a_F_vacuumRedirectAndPlaceholder_1)) {
		v85 = int32(1)
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v75 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v85 = int32(0)
	goto L8
L21:
	;
	goto L22
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+119)))
	switch v81 - int32(109) {
	case 0, 5:
		goto L23
	default:
		v85 = int32(0)
		goto L8
	}
L23:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+112)))
	v85 = v84
	goto L8
L24:
	;
	v96 = int32(_a_F_vacuumRedirectAndPlaceholder_2)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[5])) = v98 + int32(1)
	v103 = v49 & int32(_a_F_vacuumRedirectAndPlaceholder_3)
	if v103 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v374 = int32(_a_F_vacuumRedirectAndPlaceholder_2)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[5])) = v376 - int32(1)
	m.G0 = v21 + int32(1648)
	return
L26:
	;
	v106 = v40 + v50
	v110 = v103
	v118 = v4
	v119 = v49
	v120 = v4
	v121 = v4
	v123 = v4
	v126 = v4
	goto L27
L27:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)))
	v128 = int32(0)
	if base.B2i32(v127 == v128)&(v121&int32(1)) == v128 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v228 = v220 & int32(_a_F_vacuumRedirectAndPlaceholder_3)
	if v228 != 0 {
		goto L53
	} else {
		goto L54
	}
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(20)+v110<<(uint(int32(2))%32))))
	v141 = v40 + v138&int32(_a_F_vacuumRedirectAndPlaceholder_4)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142&int32(3) != int32(1) {
		v198 = v142
		v199 = v120
		v201 = v123
		v202 = v126
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v220 = v118
	v226 = v126
	goto L31
L31:
	;
	goto L28
L32:
	;
	v203 = int32(3)
	v207 = base.B2i32(v198&v203 != v203) | v121
	if v207&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	if v147 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v148 = F_GlobalVisTestIsRemovableXid(m, v95, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v153 = v142
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v153 | int32(3)
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)))
	v158 = int32(1)
	v159 = v157 - v158
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)) = uint16(v159)
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)))
	v163 = v161 + v158
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)) = uint16(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	if v120 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	return
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v148 == int32(0) {
		v198 = v150
		v199 = v120
		v201 = v123
		v202 = v126
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v153 = v150
	goto L36
L40:
	;
	v181 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+10)) = uint16(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+6)) = int32(-1)
	v185 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(832)+v123&int32(_a_F_vacuumRedirectAndPlaceholder_3)<<(uint(v185)%32)))) = uint16(v110)
	v195 = v123 + v185
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v198 = v197
	v199 = v180
	v201 = v195
	v202 = v185
	goto L32
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v165
	v180 = v165
	goto L40
L42:
	;
	v168 = int32(3)
	if base.B2i32(base.Ui32(v120) < base.Ui32(v168))|base.B2i32(base.Ui32(v165) < base.Ui32(v168)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v120-v165 < int32(0) {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(v165) <= base.Ui32(v120) {
		v180 = v120
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v180 = v120
	goto L40
L47:
	;
	goto L41
L48:
	;
	v210 = v118
	goto L50
L49:
	;
	v210 = v110
	goto L50
L50:
	;
	v211 = int32(1)
	v214 = v119 - v211
	if v214&int32(_a_F_vacuumRedirectAndPlaceholder_3) != 0 {
		v110 = v110 - v211
		v118 = v210
		v119 = v214
		v120 = v199
		v121 = v207
		v123 = v201
		v126 = v202
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v220 = v210
	v226 = v202
	goto L31
L52:
	;
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L37
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v228) <= base.Ui32(v49&int32(_a_F_vacuumRedirectAndPlaceholder_3)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v297 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v297)
	if v226 == v297 {
		goto L25
	} else {
		goto L63
	}
L56:
	;
	v233 = v220
	goto L59
L57:
	;
	goto L58
L58:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)))
	v287 = v49 - v220 + int32(1)
	v288 = v284 - v287
	*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)) = uint16(v288)
	F_PageIndexMultiDelete(m, v40, v21+int32(16), v287&int32(_a_F_vacuumRedirectAndPlaceholder_3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L37
	} else {
		goto L62
	}
L59:
	;
	v252 = int32(_a_F_vacuumRedirectAndPlaceholder_3)
	v255 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(16)+(v233&v252-v228)<<(uint(v255)%32)))) = uint16(v233)
	v260 = v233 + v255
	if base.Ui32(v260&v252) <= base.Ui32(v49&v252) {
		v233 = v260
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	goto L60
L62:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v220)
	goto L52
L63:
	;
	goto L52
L64:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+118)))
	if v322 != int32(112) {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[2]))
	if v326 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v329 != 0 {
		goto L25
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L37
	} else {
		goto L71
	}
L69:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v330 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	F_XLogRegisterData(m, v21+int32(4), int32(10))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L37
	} else {
		goto L72
	}
L72:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	F_XLogRegisterData(m, v21+int32(832), v340<<(uint(int32(1))%32))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L37
	} else {
		goto L73
	}
L73:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L37
	} else {
		goto L74
	}
L74:
	;
	v351 = F_XLogInsert(m, int32(16), int32(128))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L37
	} else {
		goto L75
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_rotl(v351, int64(32))
	goto L25
}
