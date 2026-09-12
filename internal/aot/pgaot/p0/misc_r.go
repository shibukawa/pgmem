package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseAuxProcessResourcesCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, _consts[985]))
	v5 = int32(1)
	v7 = base.B2i32(l0 == int32(0))
	F_ResourceOwnerReleaseInternal(m, v4, v5, v7, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[985]))
		F_ResourceOwnerReleaseInternal(m, v12, int32(2), v7, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[985]))
			F_ResourceOwnerReleaseInternal(m, v18, int32(3), v7, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[985]))
				v25 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)) = uint16(v25)
				return
			}
		}
	}
}
func F_ReleaseDeletionLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(1259) {
		F_UnlockRelationOid(m, v3, int32(8))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		F_UnlockDatabaseObject(m, v4, v3, int32(8))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReleaseLockIfHeld(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v12 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(0)
	goto L3
L2:
	;
	v13 = v12
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = v15
	goto L5
L4:
	;
	return
L5:
	;
	v26 = v18 - int32(1)
	if v26 < int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v31)+8))
	if v35 < v34 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v31 = v14 + v26<<(uint(int32(4))%32)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != v13 {
		v18 = v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v38 = v15 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v34 - v35
	if v13 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v54 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v59 = F_LockRelease(m, l0, v58, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L18
	}
L12:
	;
	F_ResourceOwnerForgetLock(m, v13, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v45 = v38
	goto L14
L14:
	;
	if v45 <= v26 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v45 = v44
	goto L14
L17:
	;
	v49 = v14 + v45<<(uint(int32(4))%32)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v52
	return
L18:
	;
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v63 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v63 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errmsg_internal(m, int32(523176), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(476003), int32(2656), int32(412906))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L4
}
func F_RemoveLogrotateSignalFiles(m *base.Module) {
	var v2 int32
	_ = v2
	v2 = F_unlink(m, int32(337768))
	return
}
func F_RemoveNonParentXlogFiles(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v351 int32
	_ = v351
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	v21 = int64(*(*int32)(unsafe.Add(mBase, _consts[167])))
	v22 = base.I64_div_u_s(l0-int64(1), v21)
	v24 = base.I64_div_u_s(int64(4294967296), v21)
	v25 = base.I64_div_u_s(v22, v24)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+20)) = uint32(v25)
	v28 = v22 - v24*v25
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v28)
	v30 = base.I64_div_u_s(l0, v21)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v30
	v38 = F_pg_snprintf(m, v15+int32(48), int32(64), int32(487930), v15+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v42 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v42 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v15 + int32(48)
	F_errmsg_internal(m, int32(185313), v15)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v56 = F_AllocateDir(m, int32(294639))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errfinish(m, int32(476414), int32(3961), int32(155809))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v59 = F_ReadDir(m, v56, int32(294639))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v66 = v15 + int32(48) | int32(8)
	v72 = v59
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_FreeDir(m, v56)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L87
	}
L14:
	;
	v80 = v72 + int32(19)
	if v80&int32(3) == int32(0) {
		v104 = v80
		goto L19
	} else {
		goto L20
	}
L15:
	;
	goto L13
L16:
	;
	v336 = F_ReadDir(m, v56, int32(294639))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L85
	}
L17:
	;
	if v137 != int32(24) {
		goto L16
	} else {
		goto L34
	}
L18:
	;
	v137 = v129 - v80
	goto L17
L19:
	;
	v108 = v104
	goto L28
L20:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v88 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v137 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v93 = v80
	goto L24
L24:
	;
	v97 = v93 + int32(1)
	if v97&int32(3) == int32(0) {
		v104 = v97
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v129 = v97
	goto L18
L26:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v102 != 0 {
		v93 = v97
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v117 = int32(-2139062144)
	if (int32(16843008)-v114|v114)&v117 == v117 {
		v108 = v108 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v123 = v108
	goto L31
L30:
	;
	goto L29
L31:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v127 != 0 {
		v123 = v123 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v129 = v123
	goto L18
L33:
	;
	goto L32
L34:
	;
	v140 = int32(514719)
	v144 = m.G0
	v146 = v144 - int32(32)
	v147 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v146)+24)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v146)+16)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v146)+8)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v146))) = v147
	v155 = int32(*(*uint8)(unsafe.Add(mBase, _consts[140])))
	if v155 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v223 != int32(24) {
		goto L16
	} else {
		goto L56
	}
L36:
	;
	v223 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _consts[141])))
	if v159 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v163 = v80
	goto L42
L40:
	;
	goto L41
L41:
	;
	v173 = v140
	v174 = v155
	goto L45
L42:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v169 == v155 {
		v163 = v163 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v223 = v163 - v80
	goto L35
L44:
	;
	goto L43
L45:
	;
	v181 = v146 + int32(base.Ui32(v174)>>(uint(int32(3))%32))&int32(28)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v182 | v183<<(uint(v174)%32)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	if v187 != 0 {
		v173 = v173 + v183
		v174 = v187
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v190 == int32(0) {
		v215 = v80
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v223 = v215 - v80
	goto L35
L49:
	;
	v194 = v80
	v195 = v190
	goto L50
L50:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v146+int32(base.Ui32(v195)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v203)>>(uint(v195)%32))&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v215 = v211
	goto L48
L52:
	;
	v215 = v194
	goto L48
L53:
	;
	goto L54
L54:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	v211 = v194 + int32(1)
	if v209 != 0 {
		v194 = v211
		v195 = v209
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v227 = v15 + int32(48)
	goto L59
L57:
	;
	if int32(0) <= v264-v265 {
		goto L16
	} else {
		goto L71
	}
L59:
	;
	goto L60
L60:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v234 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v235 = v80
	v236 = v227
	v237 = int32(8)
	v238 = v234
	goto L65
L62:
	;
	v260 = v227
	v264 = int32(0)
	goto L63
L63:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	goto L57
L64:
	;
	v260 = v255
	v264 = v257
	goto L63
L65:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if v238 != v240 {
		v255 = v236
		v257 = v238
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v255 = v249
	v257 = int32(0)
	goto L64
L67:
	;
	if v240 == int32(0) {
		v255 = v236
		v257 = v238
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v245 = v237 - int32(1)
	if v245 == int32(0) {
		v255 = v236
		v257 = v238
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v248 = int32(1)
	v249 = v236 + v248
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	if v250 != 0 {
		v235 = v235 + v248
		v236 = v249
		v237 = v245
		v238 = v250
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	v276 = v72 + int32(27)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v280 == int32(0) {
		v299 = v279
		v300 = v280
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v300-v299 <= int32(0) {
		goto L16
	} else {
		goto L80
	}
L73:
	;
	goto L72
L74:
	;
	if v279 != v280 {
		v299 = v279
		v300 = v280
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v284 = v276
	v285 = v66
	goto L76
L76:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v289 == int32(0) {
		v299 = v288
		v300 = v289
		goto L73
	} else {
		goto L78
	}
L77:
	;
	v299 = v288
	v300 = v289
	goto L73
L78:
	;
	v292 = int32(1)
	if v288 == v289 {
		v284 = v284 + v292
		v285 = v285 + v292
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v304 = m.G0
	v306 = v304 - int32(1136)
	m.G0 = v306
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = int32(21379)
	v315 = F_pg_snprintf(m, v306+int32(112), int32(1024), int32(166127), v306)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v323 = F___fstatat(m, int32(-100), v306+int32(112), v306+int32(16), int32(0))
	mBase = m.M
	goto L82
L82:
	;
	m.G0 = v306 + int32(1136)
	if v323 == int32(0) {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	F_RemoveXlogFile(m, v72, v30+int64(10), v15+int32(40), l1)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L16
L85:
	;
	if v336 != 0 {
		v72 = v336
		goto L14
	} else {
		goto L86
	}
L86:
	;
	goto L15
L87:
	;
	m.G0 = v15 + int32(112)
	return
}
func F_ReqShutdownXLOG(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[439])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_ResolveCminCmaxDuringDecoding(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int64
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int64
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	v7 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(1264)
	m.G0 = v22
	if l0 == v7 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L13
	} else {
		goto L167
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L13
	} else {
		goto L163
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L13
	} else {
		goto L160
	}
L4:
	;
	m.G0 = v22 + int32(1264)
	return v720
L5:
	;
	v720 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v28 = v22 + int32(160)
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = v31
	v36 = v22 + int32(144)
	if l3 < v29 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28))) = uint16(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+156)) = v69
	v73 = int32(0)
	v75 = F_hash_search(m, l0, v22+int32(144), v73, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v59
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(140)))) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(136)))) = v65
	goto L8
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v58 = v45 + (l3^int32(-1))<<(uint(int32(6))%32)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v58 = v52 + l3<<(uint(int32(6))%32) + int32(-64)
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	if v75 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v84 = int32(1)
	if v81 <= int32(3591) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v695 = v75
	goto L17
L17:
	;
	if l4 != 0 {
		goto L156
	} else {
		goto L157
	}
L18:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v156 = F_AllocateDir(m, int32(147003))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L46
	}
L19:
	;
	goto L18
L20:
	;
	v152 = int32(0)
	goto L19
L21:
	;
	if base.Ui32(v81-int32(2964)) < base.Ui32(int32(4)) {
		v152 = v84
		goto L19
	} else {
		goto L44
	}
L22:
	;
	if v81 <= int32(2670) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v81 <= int32(5999) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	switch v81 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v152 = v84
		goto L19
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L20
	default:
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v96 = v81 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v96) {
		goto L21
	} else {
		goto L30
	}
L28:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v81-int32(2396)) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v152 = v84
	goto L19
L30:
	;
	if int32(1)<<(uint(v96)%32)&int32(226492515) == int32(0) {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v152 = v84
	goto L19
L32:
	;
	if base.Ui32(v81-int32(3592)) < base.Ui32(int32(2)) {
		v152 = v84
		goto L19
	} else {
		goto L42
	}
L33:
	;
	v108 = v81 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v108) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	switch v81 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v152 = v84
		goto L19
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L20
	default:
		goto L38
	}
L36:
	;
	if int32(1)<<(uint(v108)%32)&int32(963) == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v152 = v84
	goto L19
L38:
	;
	if base.Ui32(v81-int32(6000)) < base.Ui32(int32(3)) {
		v152 = v84
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v124 = v81 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v124) {
		goto L20
	} else {
		goto L40
	}
L40:
	;
	if int32(1)<<(uint(v124)%32)&int32(49153) != 0 {
		v152 = v84
		goto L19
	} else {
		goto L41
	}
L41:
	;
	goto L20
L42:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v81-int32(4060)) {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v152 = v84
	goto L19
L44:
	;
	if base.Ui32(v81-int32(2846)) < base.Ui32(int32(2)) {
		v152 = v84
		goto L19
	} else {
		goto L45
	}
L45:
	;
	goto L20
L46:
	;
	v159 = F_ReadDir(m, v156, int32(147003))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v159 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v152 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v396 = v7
	goto L50
L50:
	;
	F_FreeDir(m, v156)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L13
	} else {
		goto L110
	}
L51:
	;
	v162 = int32(0)
	goto L53
L52:
	;
	v162 = v154
	goto L53
L53:
	;
	v169 = v159
	v175 = v7
	goto L54
L54:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+19)))
	if v186 != int32(46) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v396 = v381
	goto L50
L56:
	;
	v386 = F_ReadDir(m, v156, int32(147003))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L13
	} else {
		goto L108
	}
L57:
	;
	v199 = v169 + int32(19)
	v200 = int32(630226)
	goto L64
L58:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+20)))
	if v189 == int32(0) {
		v381 = v175
		goto L56
	} else {
		goto L59
	}
L59:
	;
	if v189 != int32(46) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+21)))
	if v194 == int32(0) {
		v381 = v175
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	if v237-v238 != 0 {
		v381 = v175
		goto L56
	} else {
		goto L76
	}
L64:
	;
	goto L65
L65:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v207 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v208 = v199
	v209 = v200
	v210 = int32(4)
	v211 = v207
	goto L70
L67:
	;
	v233 = v200
	v237 = int32(0)
	goto L68
L68:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	goto L62
L69:
	;
	v233 = v228
	v237 = v230
	goto L68
L70:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v211 != v213 {
		v228 = v209
		v230 = v211
		goto L69
	} else {
		goto L72
	}
L71:
	;
	v228 = v222
	v230 = int32(0)
	goto L69
L72:
	;
	if v213 == int32(0) {
		v228 = v209
		v230 = v211
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v218 = v210 - int32(1)
	if v218 == int32(0) {
		v228 = v209
		v230 = v211
		goto L69
	} else {
		goto L74
	}
L74:
	;
	v221 = int32(1)
	v222 = v209 + v221
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v223 != 0 {
		v208 = v208 + v221
		v209 = v222
		v210 = v218
		v211 = v223
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(132)))) = v22 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(128)))) = v22 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v22 + int32(168)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v22 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v22 + int32(184)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v22 + int32(204)
	v267 = F_sscanf(m, v199, int32(27809), v22+int32(112))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	if v267 != int32(6) {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v22)+204))
	if v271 != v162 {
		v381 = v175
		goto L56
	} else {
		goto L79
	}
L79:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v22)+184))
	if v273 != v81 {
		v381 = v175
		goto L56
	} else {
		goto L80
	}
L80:
	;
	v275 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v22)+172)))
	v276 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v22)+168)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v22)+176))
	v278 = F_TransactionIdDidCommit(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	if v278 == int32(0) {
		v381 = v175
		goto L56
	} else {
		goto L82
	}
L82:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v22)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = v284
	v290 = F_bsearch(m, v22+int32(240), v283, v282, int32(4), int32(185))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	if v290 == int32(0) {
		v381 = v175
		goto L56
	} else {
		goto L84
	}
L84:
	;
	v295 = F_palloc(m, int32(1032))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v275<<(uint(int64(32))%64) | v276
	v302 = v295 + int32(8)
	if (v199^v302)&int32(3) != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v377 = F_lappend(m, v175, v295)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L13
	} else {
		goto L107
	}
L87:
	;
	goto L86
L88:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v356)
	if v356&int32(255) == int32(0) {
		goto L87
	} else {
		goto L103
	}
L89:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v355 = v199
	v356 = v308
	v357 = v302
	goto L88
L90:
	;
	goto L91
L91:
	;
	if v199&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v312 = v199
	v314 = v302
	goto L95
L93:
	;
	v326 = v199
	v328 = v302
	goto L94
L94:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v333 = int32(-2139062144)
	if (int32(16843008)-v330|v330)&v333 != v333 {
		v355 = v326
		v356 = v330
		v357 = v328
		goto L88
	} else {
		goto L99
	}
L95:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v315)
	if v315 == int32(0) {
		goto L87
	} else {
		goto L97
	}
L96:
	;
	v326 = v322
	v328 = v320
	goto L94
L97:
	;
	v319 = int32(1)
	v320 = v314 + v319
	v322 = v312 + v319
	if v322&int32(3) != 0 {
		v312 = v322
		v314 = v320
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v338 = v326
	v339 = v330
	v340 = v328
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v339
	v342 = int32(4)
	v343 = v340 + v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v346 = v338 + v342
	v350 = int32(-2139062144)
	if (v344|(int32(16843008)-v344))&v350 == v350 {
		v338 = v346
		v339 = v344
		v340 = v343
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v355 = v346
	v356 = v344
	v357 = v343
	goto L88
L102:
	;
	goto L101
L103:
	;
	v364 = v355
	v366 = v357
	goto L104
L104:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)) = uint8(v367)
	v369 = int32(1)
	if v367 != 0 {
		v364 = v364 + v369
		v366 = v366 + v369
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L87
L106:
	;
	goto L105
L107:
	;
	v381 = v377
	goto L56
L108:
	;
	if v386 != 0 {
		v169 = v386
		v175 = v381
		goto L54
	} else {
		goto L109
	}
L109:
	;
	goto L55
L110:
	;
	F_list_sort(m, v396, int32(1018))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	if v396 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v683 = int32(0)
	v688 = F_hash_search(m, l0, v22+int32(144), v683, v683)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L13
	} else {
		goto L154
	}
L113:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v414 <= int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v418 = v22 + int32(234)
	v420 = v22 + int32(216)
	v422 = v22 + int32(196)
	v424 = v22 + int32(228)
	v426 = v22 + int32(200)
	v441 = v7
	goto L115
L115:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v441<<(uint(int32(2))%32))))
	v453 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L117
	}
L116:
	;
	goto L112
L117:
	;
	if v453 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v450 + int32(8)
	F_errmsg_internal(m, int32(44752), v22+int32(80))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L13
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = int32(147003)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v450 + int32(8)
	v481 = F_pg_sprintf(m, v22+int32(240), int32(167509), v22-int32(-64))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L13
	} else {
		goto L123
	}
L121:
	;
	F_errfinish(m, int32(474010), int32(5516), int32(147109))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v486 = F_OpenTransientFile(m, v22+int32(240), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	if v486 < int32(0) {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v491 = v22 + int32(192)
	v492 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = v492
	v494 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v494
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v492
	v498 = int32(4063740)
	v499 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = int32(167772205)
	v505 = F_read(m, v486, v22+int32(204), int32(36))
	mBase = m.M
	v507 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = v494
	if v494 <= v505 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v656 = F_CloseTransientFile(m, v486)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L13
	} else {
		goto L150
	}
L127:
	;
	v514 = v505
	goto L130
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L146
	}
L130:
	;
	if v514 != int32(36) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L129
L132:
	;
	if v514 == int32(0) {
		goto L126
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v22)+212))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v557
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v22)+204))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v559
	v562 = v22 + int32(200)
	v563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v424)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v562))) = uint16(v563)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v565
	v569 = int32(0)
	v571 = F_hash_search(m, l0, v22+int32(184), v569, v569)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L13
	} else {
		goto L141
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L13
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v22 + int32(240)
	F_errmsg(m, int32(150146), v22+int32(32))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(474010), int32(5375), int32(372304))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v598 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = v598
	v600 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v600
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v598
	v604 = int32(4063740)
	v605 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v605))) = int32(167772205)
	v611 = F_read(m, v486, v22+int32(204), int32(36))
	mBase = m.M
	v613 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v613))) = v600
	if v600 <= v611 {
		v514 = v611
		goto L130
	} else {
		goto L145
	}
L141:
	;
	if v571 == int32(0) {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v575
	v577 = *(*int64)(unsafe.Add(mBase, uint32(v420)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v577
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v562))) = uint16(v579)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = v581
	v588 = F_hash_search(m, l0, v22+int32(184), int32(1), v22+int32(180))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+180)))
	if v590 != 0 {
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v571)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+20)) = v591
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v571)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+24)) = v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v571)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+28)) = v595
	goto L140
L145:
	;
	goto L131
L146:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v22 + int32(240)
	F_errmsg(m, int32(286109), v22+int32(16))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(474010), int32(5367), int32(372304))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L13
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	if v656 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_pfree(m, v450)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	v661 = v441 + int32(1)
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v661 < v662 {
		v441 = v661
		goto L115
	} else {
		goto L153
	}
L153:
	;
	goto L116
L154:
	;
	if v688 == int32(0) {
		v720 = v683
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v695 = v688
	goto L17
L156:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v695)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v711
	goto L158
L157:
	;
	goto L158
L158:
	;
	v713 = int32(1)
	if l5 == int32(0) {
		v720 = v713
		goto L4
	} else {
		goto L159
	}
L159:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v695)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v716
	v720 = v713
	goto L4
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v199
	F_errmsg_internal(m, int32(674519), v22+int32(96))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(474010), int32(5480), int32(147109))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v22 + int32(240)
	F_errmsg(m, int32(285190), v22)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(474010), int32(5346), int32(372304))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L13
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v22 + int32(240)
	F_errmsg(m, int32(285933), v22+int32(48))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L13
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(474010), int32(5418), int32(372304))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L13
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RestoreSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int32(2)
	v19 = v17 << (uint(v18) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v20 << (uint(v18) % 32)
	v26 = F_MemoryContextAlloc(m, v16, v19+v22+int32(72))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v14
		v33 = int32(1)
		v34 = v13 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+29)) = uint8(v34)
		v37 = v12 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v37)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v20
		v40 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v40
		*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v40
		*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v40
		v49 = l0 + int32(24)
		if v17 != 0 {
			v51 = v26 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v51
			if v19 != 0 {
				v53 = F__emscripten_memcpy_bulkmem(m, v51, v49, v19)
				mBase = m.M
			} else {
			}
		} else {
		}
		if int32(0) < v20 {
			v59 = v17 << (uint(int32(2)) % 32)
			v62 = v26 + v59 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v62
			if v22 != 0 {
				v65 = F__emscripten_memcpy_bulkmem(m, v62, v49+v59, v22)
				mBase = m.M
			} else {
			}
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(v26)+44)) = int64(0)
		v71 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+30)) = uint8(v71)
		return v26
	}
}
func F_RestoreUserContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != int32(-1) {
		F_AtEOXact_GUC(m, int32(0), v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _consts[122])) = v10
			*(*int32)(unsafe.Add(mBase, _consts[4])) = v9
			return
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, _consts[122])) = v10
		*(*int32)(unsafe.Add(mBase, _consts[4])) = v9
		return
	}
}
func F_r_LONG_2(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_find_among_b(m, l0, int32(4212464), int32(7))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_r_VOWEL_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 < v10 {
		v13 = v10
	} else {
		v13 = v11
	}
	if v10 == v13 {
		v54 = int32(-1)
	} else {
		v25 = int32(1)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v10))))
		if int32(117) < v28 {
			v50 = v25
		} else {
			v30 = v28 - int32(97)
			if v30 < int32(0) {
				v50 = v25
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v30)>>(uint(int32(3))%32)))+uint32(_consts[1062]))))
				if int32(base.Ui32(v36)>>(uint(v30&int32(7))%32))&int32(1) == int32(0) {
					v50 = v25
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
					v50 = int32(0)
				}
			}
		}
		v54 = v50
	}
	return base.B2i32(v54 == int32(0))
}
func F_r_e_ending_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 <= v11 {
		v145 = v2
		return v145
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v9-int32(1)))))
		if v17 != int32(101) {
			v145 = v2
			return v145
		} else {
			v21 = v9 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v9 <= v24 {
				v145 = v2
				return v145
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v34 <= v35 {
					v75 = int32(-1)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v34-int32(1)))))
					if int32(232) < v50 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - int32(1)
						v72 = int32(0)
					} else {
						v52 = v50 - int32(97)
						if v52 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - int32(1)
							v72 = int32(0)
						} else {
							v55 = int32(1)
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v52)>>(uint(int32(3))%32)))+uint32(_consts[1057]))))
							if int32(base.Ui32(v59)>>(uint(v52&int32(7))%32))&v55 != 0 {
								v72 = v55
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34 - int32(1)
								v72 = int32(0)
							}
						}
					}
					v75 = v72
				}
				if v75 != 0 {
					v145 = v2
					return v145
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v76 + (v21 - v26)
					v80 = F_slice_del(m, l0)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						if v80 < int32(0) {
							v145 = v80
							return v145
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v87 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v87
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v91 = v89 - v87
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v91 <= v92 {
								return int32(0)
							} else {
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v91))))
								if v98&int32(224) != int32(96) {
									return int32(0)
								} else {
									if int32(1)<<(uint(v98)%32)&int32(1050640) == int32(0) {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v116 = F_find_among_b(m, l0, int32(4140144), int32(3))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											if v116 == int32(0) {
												return int32(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v124 = v122 + (v89 - v113)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v124
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
												v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v124 <= v128 {
													v145 = int32(0)
													return v145
												} else {
													v130 = int32(1)
													v131 = v124 - v130
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
													v135 = F_slice_del(m, l0)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v135 {
															v142 = v130
														} else {
															v142 = v135 >> (uint(int32(31)) % 32) & v135
														}
														v145 = v142
														return v145
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
			}
		}
	}
}
func F_r_en_ending_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 < v7 {
		v143 = v2
		return v143
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v17 <= v18 {
			v58 = int32(-1)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v17-int32(1)))))
			if int32(232) < v33 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 - int32(1)
				v55 = int32(0)
			} else {
				v35 = v33 - int32(97)
				if v35 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 - int32(1)
					v55 = int32(0)
				} else {
					v38 = int32(1)
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32)))+uint32(_consts[1057]))))
					if int32(base.Ui32(v42)>>(uint(v35&int32(7))%32))&v38 != 0 {
						v55 = v38
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 - int32(1)
						v55 = int32(0)
					}
				}
			}
			v58 = v55
		}
		if v58 != 0 {
			v143 = v2
			return v143
		} else {
			v59 = v5 - v9
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v61 = v59 + v60
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
			v63 = int32(3)
			v65 = int32(0)
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v61-v68 < v63 {
				v78 = v65
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v74 = F_memcmp(m, v71+v61-v63, int32(2125025), v63)
				mBase = m.M
				if v74 != 0 {
					v78 = v65
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61 - v63
					v78 = int32(1)
				}
			}
			if v78 != 0 {
				v143 = v2
				return v143
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79 + v59
				v82 = F_slice_del(m, l0)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					if v82 < int32(0) {
						v143 = v82
						return v143
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v90 = v88 - int32(1)
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v90 <= v91 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v90))))
							if v97&int32(224) != int32(96) {
								return int32(0)
							} else {
								if int32(1)<<(uint(v97)%32)&int32(1050640) == int32(0) {
									return int32(0)
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v115 = F_find_among_b(m, l0, int32(4140144), int32(3))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										if v115 == int32(0) {
											return int32(0)
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v123 = v121 + (v88 - v112)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v123
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v123 <= v127 {
												v143 = int32(0)
												return v143
											} else {
												v129 = int32(1)
												v130 = v123 - v129
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v130
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130
												v134 = F_slice_del(m, l0)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v134 {
														v141 = v129
													} else {
														v141 = v134 >> (uint(int32(31)) % 32) & v134
													}
													v143 = v141
													return v143
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
		}
	}
}
func F_r_remove_suffix_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v6 {
		v36 = v2
		return v36
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v4-int32(1)))))
		switch v12 - int32(105) {
		case 0, 5:
			v17 = F_find_among_b(m, l0, int32(4241936), int32(3))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v36 = v2
					return v36
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					v25 = F_slice_del(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 < int32(0) {
							v36 = v25
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							v31 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30 - v31
							v36 = v31
						}
						return v36
					}
				}
			}
		default:
			v36 = v2
			return v36
		}
	}
}
func F_r_shortv_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v24 {
		v129 = int32(-1)
		v136 = v129
	} else {
		v41 = int32(1)
		v42 = v6 - v41
		v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+v42))))
		v46 = v44 & int32(255)
		if v42 == v24 {
			v101 = v46
			v102 = v41
		} else {
			if int32(0) <= v44 {
				v101 = v46
				v102 = v41
			} else {
				v52 = v46 & int32(63)
				v54 = v6 - int32(2)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v54))))
				v58 = v56 << (uint(int32(6)) % 32)
				if base.B2i32(v54 != v24)&base.B2i32(base.Ui32(v56) < base.Ui32(int32(192))) == int32(0) {
					v101 = v58&int32(1984) | v52
					v102 = int32(2)
				} else {
					v71 = v58&int32(4032) | v52
					v73 = v6 - int32(3)
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v73))))
					if base.B2i32(v73 != v24)&base.B2i32(base.Ui32(v75) < base.Ui32(int32(224))) == int32(0) {
						v101 = v75<<(uint(int32(12))%32)&int32(61440) | v71
						v102 = int32(3)
					} else {
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+(v20-int32(4))))))
						v101 = v75<<(uint(int32(12))%32)&int32(258048) | v93&int32(7)<<(uint(int32(18))%32) | v71
						v102 = int32(4)
					}
				}
			}
		}
		if int32(121) < v101 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v102
			v129 = int32(0)
			v136 = v129
		} else {
			v106 = v101 - int32(89)
			if v106 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v102
				v129 = int32(0)
				v136 = v129
			} else {
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v106)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
				if int32(base.Ui32(v112)>>(uint(v106&int32(7))%32))&int32(1) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - v102
					v129 = int32(0)
					v136 = v129
				} else {
					v136 = v102
				}
			}
		}
	}
	if v136 != 0 {
		v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v400 = v398 + (v6 - v5)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400
		v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v400 <= v419 {
			v524 = int32(-1)
			v531 = v524
		} else {
			v436 = int32(1)
			v437 = v400 - v436
			v439 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415+v437))))
			v441 = v439 & int32(255)
			if v437 == v419 {
				v496 = v441
				v497 = v436
			} else {
				if int32(0) <= v439 {
					v496 = v441
					v497 = v436
				} else {
					v447 = v441 & int32(63)
					v449 = v400 - int32(2)
					v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v449))))
					v453 = v451 << (uint(int32(6)) % 32)
					if base.B2i32(v449 != v419)&base.B2i32(base.Ui32(v451) < base.Ui32(int32(192))) == int32(0) {
						v496 = v453&int32(1984) | v447
						v497 = int32(2)
					} else {
						v466 = v453&int32(4032) | v447
						v468 = v400 - int32(3)
						v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v468))))
						if base.B2i32(v468 != v419)&base.B2i32(base.Ui32(v470) < base.Ui32(int32(224))) == int32(0) {
							v496 = v470<<(uint(int32(12))%32)&int32(61440) | v466
							v497 = int32(3)
						} else {
							v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+(v415-int32(4))))))
							v496 = v470<<(uint(int32(12))%32)&int32(258048) | v488&int32(7)<<(uint(int32(18))%32) | v466
							v497 = int32(4)
						}
					}
				}
			}
			if int32(121) < v496 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
				v524 = int32(0)
				v531 = v524
			} else {
				v501 = v496 - int32(97)
				if v501 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
					v524 = int32(0)
					v531 = v524
				} else {
					v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v501)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
					if int32(base.Ui32(v507)>>(uint(v501&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
						v524 = int32(0)
						v531 = v524
					} else {
						v531 = v497
					}
				}
			}
		}
		if v531 != 0 {
			v664 = v2
		} else {
			v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v548 <= v549 {
				v653 = int32(-1)
				v660 = v653
			} else {
				v566 = int32(1)
				v567 = v548 - v566
				v569 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545+v567))))
				v571 = v569 & int32(255)
				if v567 == v549 {
					v626 = v571
					v627 = v566
				} else {
					if int32(0) <= v569 {
						v626 = v571
						v627 = v566
					} else {
						v577 = v571 & int32(63)
						v579 = v548 - int32(2)
						v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v579))))
						v583 = v581 << (uint(int32(6)) % 32)
						if base.B2i32(v579 != v549)&base.B2i32(base.Ui32(v581) < base.Ui32(int32(192))) == int32(0) {
							v626 = v583&int32(1984) | v577
							v627 = int32(2)
						} else {
							v596 = v583&int32(4032) | v577
							v598 = v548 - int32(3)
							v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v598))))
							if base.B2i32(v598 != v549)&base.B2i32(base.Ui32(v600) < base.Ui32(int32(224))) == int32(0) {
								v626 = v600<<(uint(int32(12))%32)&int32(61440) | v596
								v627 = int32(3)
							} else {
								v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+(v545-int32(4))))))
								v626 = v600<<(uint(int32(12))%32)&int32(258048) | v618&int32(7)<<(uint(int32(18))%32) | v596
								v627 = int32(4)
							}
						}
					}
				}
				if int32(121) < v626 {
					v660 = v627
				} else {
					v631 = v626 - int32(97)
					if v631 < int32(0) {
						v660 = v627
					} else {
						v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v631)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
						if int32(base.Ui32(v637)>>(uint(v631&int32(7))%32))&int32(1) == int32(0) {
							v660 = v627
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v548 - v627
							v653 = int32(0)
							v660 = v653
						}
					}
				}
			}
			if v660 != 0 {
				v664 = v2
			} else {
				v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v664 = base.B2i32(v661 <= v662)
			}
		}
		return v664
	} else {
		v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v153 <= v154 {
			v258 = int32(-1)
			v265 = v258
		} else {
			v171 = int32(1)
			v172 = v153 - v171
			v174 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+v172))))
			v176 = v174 & int32(255)
			if v172 == v154 {
				v231 = v176
				v232 = v171
			} else {
				if int32(0) <= v174 {
					v231 = v176
					v232 = v171
				} else {
					v182 = v176 & int32(63)
					v184 = v153 - int32(2)
					v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v184))))
					v188 = v186 << (uint(int32(6)) % 32)
					if base.B2i32(v184 != v154)&base.B2i32(base.Ui32(v186) < base.Ui32(int32(192))) == int32(0) {
						v231 = v188&int32(1984) | v182
						v232 = int32(2)
					} else {
						v201 = v188&int32(4032) | v182
						v203 = v153 - int32(3)
						v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v203))))
						if base.B2i32(v203 != v154)&base.B2i32(base.Ui32(v205) < base.Ui32(int32(224))) == int32(0) {
							v231 = v205<<(uint(int32(12))%32)&int32(61440) | v201
							v232 = int32(3)
						} else {
							v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+(v150-int32(4))))))
							v231 = v205<<(uint(int32(12))%32)&int32(258048) | v223&int32(7)<<(uint(int32(18))%32) | v201
							v232 = int32(4)
						}
					}
				}
			}
			if int32(121) < v231 {
				v265 = v232
			} else {
				v236 = v231 - int32(97)
				if v236 < int32(0) {
					v265 = v232
				} else {
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v236)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
					if int32(base.Ui32(v242)>>(uint(v236&int32(7))%32))&int32(1) == int32(0) {
						v265 = v232
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v153 - v232
						v258 = int32(0)
						v265 = v258
					}
				}
			}
		}
		if v265 != 0 {
			v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v400 = v398 + (v6 - v5)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400
			v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v400 <= v419 {
				v524 = int32(-1)
				v531 = v524
			} else {
				v436 = int32(1)
				v437 = v400 - v436
				v439 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415+v437))))
				v441 = v439 & int32(255)
				if v437 == v419 {
					v496 = v441
					v497 = v436
				} else {
					if int32(0) <= v439 {
						v496 = v441
						v497 = v436
					} else {
						v447 = v441 & int32(63)
						v449 = v400 - int32(2)
						v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v449))))
						v453 = v451 << (uint(int32(6)) % 32)
						if base.B2i32(v449 != v419)&base.B2i32(base.Ui32(v451) < base.Ui32(int32(192))) == int32(0) {
							v496 = v453&int32(1984) | v447
							v497 = int32(2)
						} else {
							v466 = v453&int32(4032) | v447
							v468 = v400 - int32(3)
							v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v468))))
							if base.B2i32(v468 != v419)&base.B2i32(base.Ui32(v470) < base.Ui32(int32(224))) == int32(0) {
								v496 = v470<<(uint(int32(12))%32)&int32(61440) | v466
								v497 = int32(3)
							} else {
								v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+(v415-int32(4))))))
								v496 = v470<<(uint(int32(12))%32)&int32(258048) | v488&int32(7)<<(uint(int32(18))%32) | v466
								v497 = int32(4)
							}
						}
					}
				}
				if int32(121) < v496 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
					v524 = int32(0)
					v531 = v524
				} else {
					v501 = v496 - int32(97)
					if v501 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
						v524 = int32(0)
						v531 = v524
					} else {
						v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v501)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
						if int32(base.Ui32(v507)>>(uint(v501&int32(7))%32))&int32(1) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
							v524 = int32(0)
							v531 = v524
						} else {
							v531 = v497
						}
					}
				}
			}
			if v531 != 0 {
				v664 = v2
			} else {
				v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v548 <= v549 {
					v653 = int32(-1)
					v660 = v653
				} else {
					v566 = int32(1)
					v567 = v548 - v566
					v569 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545+v567))))
					v571 = v569 & int32(255)
					if v567 == v549 {
						v626 = v571
						v627 = v566
					} else {
						if int32(0) <= v569 {
							v626 = v571
							v627 = v566
						} else {
							v577 = v571 & int32(63)
							v579 = v548 - int32(2)
							v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v579))))
							v583 = v581 << (uint(int32(6)) % 32)
							if base.B2i32(v579 != v549)&base.B2i32(base.Ui32(v581) < base.Ui32(int32(192))) == int32(0) {
								v626 = v583&int32(1984) | v577
								v627 = int32(2)
							} else {
								v596 = v583&int32(4032) | v577
								v598 = v548 - int32(3)
								v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v598))))
								if base.B2i32(v598 != v549)&base.B2i32(base.Ui32(v600) < base.Ui32(int32(224))) == int32(0) {
									v626 = v600<<(uint(int32(12))%32)&int32(61440) | v596
									v627 = int32(3)
								} else {
									v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+(v545-int32(4))))))
									v626 = v600<<(uint(int32(12))%32)&int32(258048) | v618&int32(7)<<(uint(int32(18))%32) | v596
									v627 = int32(4)
								}
							}
						}
					}
					if int32(121) < v626 {
						v660 = v627
					} else {
						v631 = v626 - int32(97)
						if v631 < int32(0) {
							v660 = v627
						} else {
							v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v631)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
							if int32(base.Ui32(v637)>>(uint(v631&int32(7))%32))&int32(1) == int32(0) {
								v660 = v627
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v548 - v627
								v653 = int32(0)
								v660 = v653
							}
						}
					}
				}
				if v660 != 0 {
					v664 = v2
				} else {
					v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v664 = base.B2i32(v661 <= v662)
				}
			}
			return v664
		} else {
			v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v282 <= v283 {
				v388 = int32(-1)
				v395 = v388
			} else {
				v300 = int32(1)
				v301 = v282 - v300
				v303 = int32(*(*int8)(unsafe.Add(mBase, uint32(v279+v301))))
				v305 = v303 & int32(255)
				if v301 == v283 {
					v360 = v305
					v361 = v300
				} else {
					if int32(0) <= v303 {
						v360 = v305
						v361 = v300
					} else {
						v311 = v305 & int32(63)
						v313 = v282 - int32(2)
						v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v313))))
						v317 = v315 << (uint(int32(6)) % 32)
						if base.B2i32(v313 != v283)&base.B2i32(base.Ui32(v315) < base.Ui32(int32(192))) == int32(0) {
							v360 = v317&int32(1984) | v311
							v361 = int32(2)
						} else {
							v330 = v317&int32(4032) | v311
							v332 = v282 - int32(3)
							v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v332))))
							if base.B2i32(v332 != v283)&base.B2i32(base.Ui32(v334) < base.Ui32(int32(224))) == int32(0) {
								v360 = v334<<(uint(int32(12))%32)&int32(61440) | v330
								v361 = int32(3)
							} else {
								v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282+(v279-int32(4))))))
								v360 = v334<<(uint(int32(12))%32)&int32(258048) | v352&int32(7)<<(uint(int32(18))%32) | v330
								v361 = int32(4)
							}
						}
					}
				}
				if int32(121) < v360 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282 - v361
					v388 = int32(0)
					v395 = v388
				} else {
					v365 = v360 - int32(97)
					if v365 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282 - v361
						v388 = int32(0)
						v395 = v388
					} else {
						v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v365)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
						if int32(base.Ui32(v371)>>(uint(v365&int32(7))%32))&int32(1) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282 - v361
							v388 = int32(0)
							v395 = v388
						} else {
							v395 = v361
						}
					}
				}
			}
			if v395 != 0 {
				v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v400 = v398 + (v6 - v5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400
				v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v400 <= v419 {
					v524 = int32(-1)
					v531 = v524
				} else {
					v436 = int32(1)
					v437 = v400 - v436
					v439 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415+v437))))
					v441 = v439 & int32(255)
					if v437 == v419 {
						v496 = v441
						v497 = v436
					} else {
						if int32(0) <= v439 {
							v496 = v441
							v497 = v436
						} else {
							v447 = v441 & int32(63)
							v449 = v400 - int32(2)
							v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v449))))
							v453 = v451 << (uint(int32(6)) % 32)
							if base.B2i32(v449 != v419)&base.B2i32(base.Ui32(v451) < base.Ui32(int32(192))) == int32(0) {
								v496 = v453&int32(1984) | v447
								v497 = int32(2)
							} else {
								v466 = v453&int32(4032) | v447
								v468 = v400 - int32(3)
								v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v468))))
								if base.B2i32(v468 != v419)&base.B2i32(base.Ui32(v470) < base.Ui32(int32(224))) == int32(0) {
									v496 = v470<<(uint(int32(12))%32)&int32(61440) | v466
									v497 = int32(3)
								} else {
									v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+(v415-int32(4))))))
									v496 = v470<<(uint(int32(12))%32)&int32(258048) | v488&int32(7)<<(uint(int32(18))%32) | v466
									v497 = int32(4)
								}
							}
						}
					}
					if int32(121) < v496 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
						v524 = int32(0)
						v531 = v524
					} else {
						v501 = v496 - int32(97)
						if v501 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
							v524 = int32(0)
							v531 = v524
						} else {
							v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v501)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
							if int32(base.Ui32(v507)>>(uint(v501&int32(7))%32))&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v400 - v497
								v524 = int32(0)
								v531 = v524
							} else {
								v531 = v497
							}
						}
					}
				}
				if v531 != 0 {
					v664 = v2
				} else {
					v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v548 <= v549 {
						v653 = int32(-1)
						v660 = v653
					} else {
						v566 = int32(1)
						v567 = v548 - v566
						v569 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545+v567))))
						v571 = v569 & int32(255)
						if v567 == v549 {
							v626 = v571
							v627 = v566
						} else {
							if int32(0) <= v569 {
								v626 = v571
								v627 = v566
							} else {
								v577 = v571 & int32(63)
								v579 = v548 - int32(2)
								v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v579))))
								v583 = v581 << (uint(int32(6)) % 32)
								if base.B2i32(v579 != v549)&base.B2i32(base.Ui32(v581) < base.Ui32(int32(192))) == int32(0) {
									v626 = v583&int32(1984) | v577
									v627 = int32(2)
								} else {
									v596 = v583&int32(4032) | v577
									v598 = v548 - int32(3)
									v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v598))))
									if base.B2i32(v598 != v549)&base.B2i32(base.Ui32(v600) < base.Ui32(int32(224))) == int32(0) {
										v626 = v600<<(uint(int32(12))%32)&int32(61440) | v596
										v627 = int32(3)
									} else {
										v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+(v545-int32(4))))))
										v626 = v600<<(uint(int32(12))%32)&int32(258048) | v618&int32(7)<<(uint(int32(18))%32) | v596
										v627 = int32(4)
									}
								}
							}
						}
						if int32(121) < v626 {
							v660 = v627
						} else {
							v631 = v626 - int32(97)
							if v631 < int32(0) {
								v660 = v627
							} else {
								v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v631)>>(uint(int32(3))%32)))+uint32(_consts[1066]))))
								if int32(base.Ui32(v637)>>(uint(v631&int32(7))%32))&int32(1) == int32(0) {
									v660 = v627
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v548 - v627
									v653 = int32(0)
									v660 = v653
								}
							}
						}
					}
					if v660 != 0 {
						v664 = v2
					} else {
						v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v664 = base.B2i32(v661 <= v662)
					}
				}
				return v664
			} else {
				return int32(1)
			}
		}
	}
}
func F_rainbow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v187 int32
	_ = v187
	if l2 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(24)
	v17 = v13 + v14*v15
	if base.Ui32(v17+v15) <= base.Ui32(v13) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v124 != 0 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v26 = v13
	v27 = int32(0)
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L1
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v33&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if base.Ui32(v26) < base.Ui32(v17) {
		v26 = v26 + int32(24)
		v27 = v27 + int32(1)
		goto L6
	} else {
		goto L40
	}
L10:
	;
	if v33&int32(2) != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v38 = int32(65535)
	v39 = v27 & v38
	if v39 == l2&v38 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+8)))
	if v43 == v39 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v49 <= v50 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	return
L18:
	;
	goto L16
L19:
	;
	F_createarc(m, l0, int32(112), base.I32_extend16_s(v27), l3, l4)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L17
	} else {
		goto L39
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v52 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v73 == int32(0) {
		goto L19
	} else {
		goto L31
	}
L23:
	;
	v60 = v52
	goto L24
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v65 != l4 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L19
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v72 != 0 {
		v60 = v72
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	if v67 != v39 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v69 == int32(112) {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L25
L31:
	;
	v81 = v73
	goto L32
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v86 != l3 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L19
L34:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
	if v93 != 0 {
		v81 = v93
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+4)))
	if v88 != v39 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v90 == int32(112) {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	goto L9
L40:
	;
	goto L7
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v127 <= v128 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L43
L45:
	;
	F_createarc(m, l0, int32(112), int32(-2), l3, l4)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L17
	} else {
		goto L65
	}
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v130 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v152 == int32(0) {
		goto L45
	} else {
		goto L57
	}
L49:
	;
	v138 = v130
	goto L50
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	if v143 != l4 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L45
L52:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	if v151 != 0 {
		v138 = v151
		goto L50
	} else {
		goto L56
	}
L53:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)))
	if v145 != int32(65534) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v148 == int32(112) {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	goto L51
L57:
	;
	v160 = v152
	goto L58
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v165 != l3 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L45
L60:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v160)+24))
	if v173 != 0 {
		v160 = v173
		goto L58
	} else {
		goto L64
	}
L61:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+4)))
	if v167 != int32(65534) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v170 == int32(112) {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L59
L65:
	;
	goto L1
}
func F_recordMultipleDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v226 int32
	_ = v226
	v5 = int32(0)
	if l2 <= v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v25 = int32(2340)
	if base.Ui32(v25) <= base.Ui32(l2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = v25
	goto L8
L7:
	;
	v28 = l2
	goto L8
L8:
	;
	v31 = F_palloc(m, v28<<(uint(int32(2))%32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v34 = l1
	v37 = v5
	v41 = v5
	v42 = v5
	goto L11
L10:
	;
	if v177 != 0 {
		goto L43
	} else {
		goto L44
	}
L11:
	;
	v48 = v34
	v53 = int32(0)
	v55 = v41
	v56 = v42
	goto L13
L12:
	;
	if v158 <= int32(0) {
		v177 = v37
		v179 = v159
		goto L10
	} else {
		goto L37
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v62 = int32(0)
	if v60 == int32(2613) {
		v75 = v62
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L12
L15:
	;
	v164 = v56 + int32(1)
	if v164 != l2 {
		v48 = v48 + int32(12)
		v53 = v158
		v55 = v159
		v56 = v164
		goto L13
	} else {
		goto L36
	}
L16:
	;
	if v75 != 0 {
		v158 = v53
		v159 = v55
		goto L15
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui32(int32(11999)) < base.Ui32(v61) {
		v75 = v62
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v75 = (base.B2i32(v60 != int32(2615)) | base.B2i32(v61 != int32(2200))) & base.B2i32(v60 != int32(1262))
	goto L17
L20:
	;
	if v28 <= v55 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	m.T0[v94].(func(*base.Module, int32))(m, v91)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L26
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v31+v53<<(uint(int32(2))%32))))
	v91 = v80
	v92 = v55
	goto L21
L23:
	;
	goto L24
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v86 = F_MakeSingleTupleTableSlot(m, v84, int32(1577592))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31+v53<<(uint(int32(2))%32)))) = v86
	v91 = v86
	v92 = v55 + int32(1)
	goto L21
L26:
	;
	v99 = v31 + v53<<(uint(int32(2))%32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+16)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = base.I32_extend8_s(l3)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v133 = F__emscripten_memset_bulkmem(m, v128, base.I32_extend8_s(int32(0)), v131)
	mBase = m.M
	goto L27
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)))
	v137 = v135 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+4)) = uint16(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+6)) = uint16(v140)
	goto L28
L28:
	;
	v143 = v53 + int32(1)
	if v143 != v28 {
		v158 = v143
		v159 = v92
		goto L15
	} else {
		goto L29
	}
L29:
	;
	if v37 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v147 = F_CatalogOpenIndexes(m, v23)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	v149 = v37
	goto L32
L32:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, v23, v31, v28, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v149 = v147
	goto L32
L34:
	;
	v155 = v56 + int32(1)
	if v155 != l2 {
		v34 = v48 + int32(12)
		v37 = v149
		v41 = v92
		v42 = v155
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v177 = v149
	v179 = v92
	goto L10
L36:
	;
	goto L14
L37:
	;
	if v37 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v170 = F_CatalogOpenIndexes(m, v23)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	v172 = v37
	goto L40
L40:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, v23, v31, v158, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	v172 = v170
	goto L40
L42:
	;
	v177 = v172
	v179 = v159
	goto L10
L43:
	;
	F_CatalogCloseIndexes(m, v177)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_sequence_close(m, v23, int32(3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	if int32(0) < v179 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v191 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_pfree(m, v31)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L55
	}
L51:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v31+v191<<(uint(int32(2))%32))))
	F_ExecDropSingleTupleTableSlot(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v210 = v191 + int32(1)
	if v210 != v179 {
		v191 = v210
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L1
}
func F_record_image_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v598 int32
	_ = v598
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	v29 = m.G0
	v31 = v29 - int32(80)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v43 = F_lookup_rowtype_tupdesc(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v48 = F_lookup_rowtype_tupdesc(m, v46, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+76)) = v34
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v53
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+68)) = uint16(v53)
	v57 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v57
	v59 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = int32(base.Ui32(v51) >> (uint(v59) % 32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = v53
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+48)) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = int32(base.Ui32(v62) >> (uint(v59) % 32))
	if v50 < v45 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v74 = v45
	goto L8
L7:
	;
	v74 = v50
	goto L8
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v76 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v99 != v41 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	v87 = F_MemoryContextAlloc(m, v82, v74<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v79 < v74 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v98 = v76
	v99 = v81
	goto L9
L13:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v87
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v93 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v92)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v92)+12)) = v93
	v98 = v92
	v99 = int32(0)
	goto L9
L14:
	;
	v150 = F_palloc(m, v45<<(uint(int32(2))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v108 = v74 << (uint(int32(2)) % 32)
	v110 = v98 + int32(20)
	if v110&int32(3) != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v101 != v42 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	if v103 != v46 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	if v105 == v47 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v98)+12)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v41
	goto L14
L21:
	;
	v136 = F__emscripten_memset_bulkmem(m, v110, base.I32_extend8_s(int32(0)), v108)
	mBase = m.M
	goto L29
L22:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v108) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v110+v108) <= base.Ui32(v110) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v120 = v98 + v108 + int32(20)
	v122 = v98 + int32(24)
	if base.Ui32(v122) < base.Ui32(v120) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v124 = v120
	goto L27
L26:
	;
	v124 = v122
	goto L27
L27:
	;
	v133 = F__emscripten_memset_bulkmem(m, v110, base.I32_extend8_s(int32(0)), (v124-v98-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L28
L28:
	;
	goto L20
L29:
	;
	goto L20
L30:
	;
	v152 = F_palloc(m, v45)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_heap_deform_tuple(m, v31+int32(60), v43, v150, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v160 = F_palloc(m, v50<<(uint(int32(2))%32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v162 = F_palloc(m, v50)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v31+int32(40), v48, v160, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v166 = int32(0)
	v173 = int32(111)
	v178 = base.B2i32(v166 < v45)
	if v166 < v45 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L170
	}
L37:
	;
	F_pfree(m, v150)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L150
	}
L38:
	;
	if v565 != v45 {
		goto L36
	} else {
		goto L148
	}
L39:
	;
	v188 = v185
	v189 = base.B2i32(v166 < v50)
	v190 = v166
	v192 = v178
	v197 = v186
	goto L44
L40:
	;
	v179 = int32(0)
	v185 = v179
	v186 = v179
	goto L39
L41:
	;
	goto L42
L42:
	;
	v181 = int32(0)
	if v50 <= v181 {
		v563 = v181
		v565 = v166
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v185 = v181
	v186 = v181
	goto L39
L44:
	;
	if v192&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v563 = v547
	v565 = v549
	goto L38
L46:
	;
	v559 = base.B2i32(v547 < v50)
	v560 = base.B2i32(v549 < v45)
	if v549 < v45 {
		v188 = v547
		v189 = v559
		v190 = v549
		v192 = v560
		v197 = v552
		goto L44
	} else {
		goto L146
	}
L47:
	;
	if v189&int32(1) == int32(0) {
		v563 = v188
		v565 = v190
		goto L38
	} else {
		goto L50
	}
L48:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v173+v219<<(uint(int32(4))%32)+v190*int32(100)))))
	if v226 != int32(1) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v547 = v188
	v549 = v190 + int32(1)
	v552 = v197
	goto L46
L50:
	;
	v236 = v188 * int32(100)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v239 = v237 << (uint(int32(4)) % 32)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+(v48+v173+v239)))))
	if v242 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v547 = v188 + int32(1)
	v549 = v190
	v552 = v197
	goto L46
L52:
	;
	goto L53
L53:
	;
	if v192&int32(1) == int32(0) {
		v563 = v188
		v565 = v190
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v257 = v43 + int32(20) + v251<<(uint(int32(4))%32) + v190*int32(100)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+68))
	v260 = v239 + (v48 + int32(88)) + v236
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v258 == v261 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v541 = int32(1)
	v547 = v188 + v541
	v549 = v190 + v541
	v552 = v197 + v541
	goto L46
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L143
	}
L57:
	;
	v263 = int32(1)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v162))))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v152))))
	if v267 == v263 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L137
	}
L60:
	;
	if v265&int32(1) == int32(0) {
		v598 = v263
		goto L37
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v274 = int32(-1)
	if v265&int32(1) != 0 {
		v598 = v274
		goto L37
	} else {
		goto L64
	}
L63:
	;
	goto L55
L64:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+82)))
	if v277 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v280 = int32(2)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v150+v190<<(uint(v280)%32))))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v160+v188<<(uint(v280)%32))))
	if v283 == v287 {
		goto L55
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v293 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257)+72)))
	if int32(0) < v293 {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	if base.Ui32(v283) < base.Ui32(v287) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v292 = int32(-1)
	goto L71
L70:
	;
	v292 = int32(1)
	goto L71
L71:
	;
	v598 = v292
	goto L37
L72:
	;
	if v477 < int32(0) {
		v598 = v274
		goto L37
	} else {
		goto L135
	}
L73:
	;
	v296 = int32(2)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v150+v190<<(uint(v296)%32))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v160+v188<<(uint(v296)%32))))
	if base.Ui32(int32(4)) <= base.Ui32(v293) {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	goto L75
L75:
	;
	if v293 != int32(-1) {
		goto L56
	} else {
		goto L94
	}
L76:
	;
	v477 = v365
	goto L72
L77:
	;
	v365 = int32(0)
	goto L76
L78:
	;
	v339 = v334
	v340 = v335
	v341 = v336
	goto L88
L79:
	;
	if (v299|v303)&int32(3) != 0 {
		v334 = v299
		v335 = v303
		v336 = v293
		goto L78
	} else {
		goto L82
	}
L80:
	;
	v327 = v299
	v328 = v303
	v329 = v293
	goto L81
L81:
	;
	if v329 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v311 = v299
	v312 = v303
	v313 = v293
	goto L83
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v316 != v317 {
		v334 = v311
		v335 = v312
		v336 = v313
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v327 = v322
	v328 = v320
	v329 = v324
	goto L81
L85:
	;
	v319 = int32(4)
	v320 = v312 + v319
	v322 = v311 + v319
	v324 = v313 - v319
	if base.Ui32(int32(3)) < base.Ui32(v324) {
		v311 = v322
		v312 = v320
		v313 = v324
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v334 = v327
	v335 = v328
	v336 = v329
	goto L78
L88:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v344 == v345 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v365 = v344 - v345
	goto L76
L90:
	;
	v347 = int32(1)
	v352 = v341 - v347
	if v352 != 0 {
		v339 = v339 + v347
		v340 = v340 + v347
		v341 = v352
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L77
L94:
	;
	v370 = v150 + v190<<(uint(int32(2))%32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v372 = F_toast_raw_datum_size(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v376 = v160 + v188<<(uint(int32(2))%32)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v378 = F_toast_raw_datum_size(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v381 = base.B2i32(base.Ui32(v372) < base.Ui32(v378))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v384 = F_pg_detoast_datum_packed(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v387 = F_pg_detoast_datum_packed(m, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v389 = int32(1)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v391&v389 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v394 = v389
	goto L101
L100:
	;
	v394 = int32(4)
	goto L101
L101:
	;
	v395 = v384 + v394
	v396 = int32(1)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387))))
	if v398&v396 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v401 = v396
	goto L104
L103:
	;
	v401 = int32(4)
	goto L104
L104:
	;
	v402 = v387 + v401
	if base.Ui32(v372) < base.Ui32(v378) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v403 = v372
	goto L107
L106:
	;
	v403 = v378
	goto L107
L107:
	;
	v404 = int32(4)
	v405 = v403 - v404
	if base.Ui32(v404) <= base.Ui32(v405) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	if v468 != v384 {
		goto L126
	} else {
		goto L127
	}
L109:
	;
	v467 = int32(0)
	goto L108
L110:
	;
	v441 = v436
	v442 = v437
	v443 = v438
	goto L120
L111:
	;
	if (v395|v402)&int32(3) != 0 {
		v436 = v395
		v437 = v402
		v438 = v405
		goto L110
	} else {
		goto L114
	}
L112:
	;
	v429 = v395
	v430 = v402
	v431 = v405
	goto L113
L113:
	;
	if v431 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L114:
	;
	v413 = v395
	v414 = v402
	v415 = v405
	goto L115
L115:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	if v418 != v419 {
		v436 = v413
		v437 = v414
		v438 = v415
		goto L110
	} else {
		goto L117
	}
L116:
	;
	v429 = v424
	v430 = v422
	v431 = v426
	goto L113
L117:
	;
	v421 = int32(4)
	v422 = v414 + v421
	v424 = v413 + v421
	v426 = v415 - v421
	if base.Ui32(int32(3)) < base.Ui32(v426) {
		v413 = v424
		v414 = v422
		v415 = v426
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v436 = v429
	v437 = v430
	v438 = v431
	goto L110
L120:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441))))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	if v446 == v447 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v467 = v446 - v447
	goto L108
L122:
	;
	v449 = int32(1)
	v454 = v443 - v449
	if v454 != 0 {
		v441 = v441 + v449
		v442 = v442 + v449
		v443 = v454
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L109
L126:
	;
	F_pfree(m, v384)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v467 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v472 = v467
	goto L132
L131:
	;
	v472 = base.B2i32(base.Ui32(v378) < base.Ui32(v372)) - v381
	goto L132
L132:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	if v387 == v473 {
		v477 = v472
		goto L72
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v387)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v477 = v472
	goto L72
L135:
	;
	if v477 == int32(0) {
		goto L55
	} else {
		goto L136
	}
L136:
	;
	v598 = int32(1)
	goto L37
L137:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v257)+68))
	v498 = F_format_type_be(m, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v501 = F_format_type_be(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v197 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v498
	F_errmsg(m, int32(452677), v31+int32(16))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(472660), int32(1474), int32(224989))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v522
	F_errmsg_internal(m, int32(462144), v31)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(472660), int32(1538), int32(224989))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	if v547 < v50 {
		v188 = v547
		v189 = v559
		v190 = v549
		v192 = v560
		v197 = v552
		goto L44
	} else {
		goto L147
	}
L147:
	;
	goto L45
L148:
	;
	if v563 != v50 {
		goto L36
	} else {
		goto L149
	}
L149:
	;
	v598 = int32(0)
	goto L37
L150:
	;
	F_pfree(m, v152)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_pfree(m, v160)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_pfree(m, v162)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if int32(0) <= v629 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F_DecrTupleDescRefCount(m, v43)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if int32(0) <= v634 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L156
L158:
	;
	F_DecrTupleDescRefCount(m, v48)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v639 != v34 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	goto L160
L162:
	;
	F_pfree(m, v34)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v643 != v39 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L164
L166:
	;
	F_pfree(m, v39)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	m.G0 = v31 + int32(80)
	return v598
L169:
	;
	goto L168
L170:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(139119), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(472660), int32(1568), int32(224989))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_record_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_record_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v264 int32
	_ = v264
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = int32(0)
	if base.B2i32(v25 == int32(2249))&base.B2i32(v26 < v34) == v34 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v524 = F_heap_form_tuple(m, v39, v108, v110)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L100
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L96
	}
L5:
	;
	if v113 == int32(0) {
		goto L3
	} else {
		goto L95
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L91
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L87
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L81
	}
L9:
	;
	v39 = F_lookup_rowtype_tupdesc(m, v25, v26)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L77
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v43 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v108 = F_palloc(m, v41<<(uint(int32(2))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v65 == v25 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	v54 = F_MemoryContextAlloc(m, v49, v41*int32(44)+int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v46 != v41 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v64 = v43
	v65 = v48
	goto L14
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = int64(0)
	v64 = v59
	v65 = int32(0)
	goto L14
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 == v26 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v70 = v41 * int32(44)
	v72 = v70 + int32(12)
	if v64&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v25
	goto L13
L24:
	;
	v98 = F__emscripten_memset_bulkmem(m, v64, base.I32_extend8_s(int32(0)), v72)
	mBase = m.M
	goto L32
L25:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v72) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v72+v64) <= base.Ui32(v64) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v84 = v64 + v70 + int32(12)
	v86 = v64 + int32(4)
	if base.Ui32(v86) < base.Ui32(v84) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = v84
	goto L30
L29:
	;
	v88 = v86
	goto L30
L30:
	;
	v95 = F__emscripten_memset_bulkmem(m, v64, base.I32_extend8_s(int32(0)), (v64^int32(-1)+v88)&int32(-4)+int32(4))
	mBase = m.M
	goto L31
L31:
	;
	goto L23
L32:
	;
	goto L23
L33:
	;
	v110 = F_palloc(m, v41)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v113 = F_pq_getmsgint(m, v27, int32(4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v41 <= int32(0) {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v118 = v41 & int32(3)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v120 = int32(4)
	v122 = v39 + v119<<(uint(v120)%32)
	v124 = v122 + int32(111)
	v125 = int32(0)
	if base.Ui32(v120) <= base.Ui32(v41) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v139 = v125
	v140 = v125
	v149 = int32(0)
	goto L40
L38:
	;
	v186 = v125
	v187 = v125
	goto L39
L39:
	;
	if v118 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v159 = v139 * int32(100)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v159))))
	v162 = int32(1)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+(v122+int32(211))))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+(v122+int32(311))))))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+(v122+int32(411))))))
	v179 = v140 + (v161 ^ v162) + (v166 ^ v162) + (v171 ^ v162) + (v176 ^ v162)
	v180 = int32(4)
	v181 = v139 + v180
	v183 = v149 + v180
	if v183 != v41&int32(2147483644) {
		v139 = v181
		v140 = v179
		v149 = v183
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v186 = v181
	v187 = v179
	goto L39
L42:
	;
	goto L41
L43:
	;
	v206 = v186
	v207 = v187
	v211 = v125
	goto L46
L44:
	;
	v239 = v187
	goto L45
L45:
	;
	if v239 != v113 {
		v467 = v239
		goto L4
	} else {
		goto L49
	}
L46:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v206*int32(100)))))
	v229 = int32(1)
	v231 = v207 + (v228 ^ v229)
	v235 = v211 + v229
	if v235 != v118 {
		v206 = v206 + v229
		v207 = v231
		v211 = v235
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v239 = v231
	goto L45
L48:
	;
	goto L47
L49:
	;
	v264 = int32(0)
	goto L50
L50:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v289 = v39 + int32(20) + v283<<(uint(int32(4))%32) + v264*int32(100)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+91)))
	if v290 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L3
L52:
	;
	v377 = v264 + int32(1)
	if v377 != v41 {
		v264 = v377
		goto L50
	} else {
		goto L76
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108+v264<<(uint(int32(2))%32)))) = int32(0)
	v299 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264+v110))) = uint8(v299)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
	v303 = F_pq_getmsgint(m, v27, int32(4))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v301) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v311 = F_pq_getmsgint(m, v27, int32(4))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L61
	}
L58:
	;
	if v303 == v301 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v303) <= base.Ui32(int32(9999)) {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	if v311 < int32(-1) {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v315-v316 < v311 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v321 = v64 + int32(12) + v264*int32(44)
	if v311 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v264+v110))) = uint8(v337)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	if v301 != v341 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v337 = int32(1)
	v338 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v311 + v316
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v316 + v328
	v337 = int32(0)
	v338 = v23 - int32(-64)
	goto L64
L68:
	;
	F_getTypeBinaryInputInfo(m, v301, v321+int32(4), v321+int32(8))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v289)+76))
	v364 = F_ReceiveFunctionCall(m, v321+int32(16), v338, v362, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+20))
	F_fmgr_info_cxt(m, v349, v321+int32(16), v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v301
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108+v264<<(uint(int32(2))%32)))) = v364
	if v338 == int32(0) {
		goto L52
	} else {
		goto L74
	}
L74:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	if v369 != v311 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	goto L52
L76:
	;
	goto L51
L77:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(426148), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(472660), int32(509), int32(34410))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v404 = F_format_type_extended(m, v303, int32(-1), int32(2))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v408 = F_format_type_extended(m, v301, int32(-1), int32(2))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v264 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v303
	F_errmsg(m, int32(452789), v23)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(472660), int32(606), int32(34410))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(386717), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(472660), int32(613), int32(34410))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v264 + int32(1)
	F_errmsg(m, int32(452746), v23+int32(32))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(472660), int32(661), int32(34410))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	v467 = int32(0)
	goto L4
L96:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v113
	F_errmsg(m, int32(458405), v23+int32(48))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(472660), int32(559), int32(34410))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v527 = F_palloc(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	if v530 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	F_pfree(m, v524)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L106
	}
L103:
	;
	v531 = F__emscripten_memcpy_bulkmem(m, v527, v529, v530)
	mBase = m.M
	v532 = v531
	goto L105
L104:
	;
	v532 = v527
	goto L105
L105:
	;
	goto L102
L106:
	;
	F_pfree(m, v108)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_pfree(m, v110)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if int32(0) <= v539 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_DecrTupleDescRefCount(m, v39)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v544 = F_HeapTupleHeaderGetDatum(m, v532)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	m.G0 = v23 + int32(80)
	return v544
}
func F_recurse_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = l0
	goto L2
L1:
	;
	m.G0 = v8 + int32(16)
	return v58
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v15 != int32(142) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L14
	}
L4:
	;
	goto L3
L5:
	;
	if v15 != int32(63) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v35 == int32(3) {
		v58 = v34
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)-int32(4))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v30 = F_subquery_is_pushdown_safe(m, v29, l1, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v58 = v30
	goto L1
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v39 = F_recurse_pushdown_safe(m, v38, l1, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v39 == int32(0) {
		v58 = v34
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v10 = v43
	goto L2
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v48
	F_errmsg_internal(m, int32(465033), v8)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(472601), int32(3712), int32(391643))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_recvfrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v7 = m.Env.X__syscall_recvfrom(m, l0, l1, l2, l3, l4, l5)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_redirect_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	F_appendStringInfo(m, l0, int32(36280), v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_regclassout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(630250))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v53 = v14
			m.G0 = v8 + int32(16)
			return v53
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(57), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v25 = v23 + int32(4)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[205]))
				if v27 == int32(0) {
					v30 = F_pstrdup(m, v25)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v19)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v53 = v30
							m.G0 = v8 + int32(16)
							return v53
						}
					}
				} else {
					v34 = F_RelationIsVisible(m, v10)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							v40 = int32(0)
							v41 = F_quote_qualified_identifier(m, v40, v25)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v53 = v41
									m.G0 = v8 + int32(16)
									return v53
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
							v38 = F_get_namespace_name(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = v38
								v41 = F_quote_qualified_identifier(m, v40, v25)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v53 = v41
										m.G0 = v8 + int32(16)
										return v53
									}
								}
							}
						}
					}
				}
			} else {
				v46 = F_palloc(m, int32(64))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v51 = F_pg_snprintf(m, v46, int32(64), int32(57094), v8)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = v46
						m.G0 = v8 + int32(16)
						return v53
					}
				}
			}
		}
	}
}
func F_regex_selectivity_sub(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 float64
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v139 int32
	_ = v139
	var v143 float64
	_ = v143
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 float64
	_ = v161
	var v168 float64
	_ = v168
	var v171 float64
	_ = v171
	v5 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0)
L2:
	;
	v14 = float64(1)
	if l1 <= int32(0) {
		v161 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v168 = float64(1)
	if base.F64_gt(v161, v168) != 0 {
		goto L68
	} else {
		goto L69
	}
L4:
	;
	v19 = v14
	v21 = v5
	v22 = v5
	v24 = v5
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v21))))
	if v27 == int32(40) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v161 = v149
	goto L3
L7:
	;
	v157 = v151 + int32(1)
	if v157 < l1 {
		v19 = v149
		v21 = v157
		v22 = v152
		v24 = v154
		goto L5
	} else {
		goto L67
	}
L8:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if v27 != int32(41) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v30 = v24
	goto L13
L12:
	;
	v30 = v21
	goto L13
L13:
	;
	v149 = v19
	v151 = v21
	v152 = v22 + int32(1)
	v154 = v30
	goto L7
L14:
	;
	if v27 != int32(124) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	if v22 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = v22 - int32(1)
	if v38 != 0 {
		v149 = v19
		v151 = v21
		v152 = v38
		v154 = v24
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v40 = v24 + int32(1)
	v43 = F_regex_selectivity_sub(m, l0+v40, v21-v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v149 = base.F64_mul(v19, v43)
	v151 = v21
	v152 = int32(0)
	v154 = v24
	goto L7
L19:
	;
	switch v27 - int32(42) {
	case 0, 1, 21:
		goto L25
	case 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48:
		goto L23
	case 4:
		goto L26
	case 49:
		goto L27
	case 50:
		goto L24
	default:
		goto L28
	}
L20:
	;
	if v22 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v50 = v21 + int32(1)
	v53 = F_regex_selectivity_sub(m, l0+v50, l1-v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v161 = base.F64_add(v19, v53)
	goto L3
L23:
	;
	if v22 != 0 {
		goto L64
	} else {
		goto L65
	}
L24:
	;
	v139 = v21 + int32(1)
	if l1 <= v139 {
		v161 = v19
		goto L3
	} else {
		goto L60
	}
L25:
	;
	if v22 != 0 {
		goto L57
	} else {
		goto L58
	}
L26:
	;
	if v22 != 0 {
		goto L54
	} else {
		goto L55
	}
L27:
	;
	v91 = v21 + int32(1)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v91))))
	v95 = base.B2i32(v93 == int32(94))
	if v93 == int32(94) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	if v27 != int32(123) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if l1 <= v21 {
		v81 = v21
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v22 != 0 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v65 = v21
	goto L32
L32:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v65))))
	if v71 == int32(125) {
		v81 = v65
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v81 = l1
	goto L30
L34:
	;
	v75 = v65 + int32(1)
	if v75 != l1 {
		v65 = v75
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v87 = v19
	goto L38
L37:
	;
	v87 = base.F64_add(v19, v19)
	goto L38
L38:
	;
	v149 = v87
	v151 = v81
	v152 = v22
	v154 = v24
	goto L7
L39:
	;
	v96 = float64(0.75)
	goto L41
L40:
	;
	v96 = float64(0.25)
	goto L41
L41:
	;
	if v93 == int32(94) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v22 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v99 = v21 + int32(2)
	goto L45
L44:
	;
	v99 = v91
	goto L45
L45:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v99))))
	v104 = v99 + base.B2i32(v101 == int32(93))
	if l1 <= v104 {
		v126 = v104
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v110 = v104
	goto L47
L47:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v110))))
	if v116 == int32(93) {
		v126 = v110
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v126 = l1
	goto L42
L49:
	;
	v120 = v110 + int32(1)
	if v120 < l1 {
		v110 = v120
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v132 = v19
	goto L53
L52:
	;
	v132 = base.F64_mul(v19, v96)
	goto L53
L53:
	;
	v149 = v132
	v151 = v126
	v152 = v22
	v154 = v24
	goto L7
L54:
	;
	v135 = v19
	goto L56
L55:
	;
	v135 = base.F64_mul(v19, float64(0.9))
	goto L56
L56:
	;
	v149 = v135
	v151 = v21
	v152 = v22
	v154 = v24
	goto L7
L57:
	;
	v137 = v19
	goto L59
L58:
	;
	v137 = base.F64_add(v19, v19)
	goto L59
L59:
	;
	v149 = v137
	v151 = v21
	v152 = v22
	v154 = v24
	goto L7
L60:
	;
	if v22 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v143 = v19
	goto L63
L62:
	;
	v143 = base.F64_mul(v19, float64(0.2))
	goto L63
L63:
	;
	v149 = v143
	v151 = v139
	v152 = v22
	v154 = v24
	goto L7
L64:
	;
	v146 = v19
	goto L66
L65:
	;
	v146 = base.F64_mul(v19, float64(0.2))
	goto L66
L66:
	;
	v149 = v146
	v151 = v21
	v152 = v22
	v154 = v24
	goto L7
L67:
	;
	goto L6
L68:
	;
	v171 = v168
	goto L70
L69:
	;
	v171 = v161
	goto L70
L70:
	;
	return v171
}
func F_regexnesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_get_negator(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(199619), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(471451), int32(773), int32(293070))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = F_patternsel_common(m, v9, v11, int32(0), v7, v6, v8, int32(2), int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_Float8GetDatum(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_register_dirty_segment(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v171 int32
	_ = v171
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v223 int32
	_ = v223
	v2 = l1
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v14
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)) = uint16(v2)
	v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v17
	v23 = F_RegisterSyncRequest(m, v8+int32(8), v4, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		if v23 == int32(0) {
			v29 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 != 0 {
					F_errmsg_internal(m, int32(289836), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(477940), int32(1522), int32(89451))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
							v44 = m.G0
							v46 = v44 - int32(16)
							m.G0 = v46
							if v41 != 0 {
								F___clock_gettime(m, int32(1), v46)
								mBase = m.M
								v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
								v55 = v50 + v51*int64(1000000000)
							} else {
								v55 = int64(0)
							}
							m.G0 = v46 + int32(16)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v61 = F_FileSync(m, v59, int32(167772182))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								if int32(0) <= v61 {
									v95 = int32(1)
									v97 = int64(0)
									v101 = m.G0
									v103 = v101 - int32(16)
									m.G0 = v103
									if v55 != v97 {
										F___clock_gettime(m, int32(1), v103)
										mBase = m.M
										v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
										v114 = v109 + (v110*int64(1000000000) - v55)
										v166 = int32(4425760)
										v167 = *(*int64)(unsafe.Add(mBase, _consts[642]))
										*(*int64)(unsafe.Add(mBase, _consts[642])) = v167 + v114
										v171 = *(*int32)(unsafe.Add(mBase, _consts[172]))
										if base.Ui32(int32(16)) < base.Ui32(v171) {
										} else {
											if int32(1)<<(uint(v171)%32)&int32(115186) == int32(0) {
											} else {
												v189 = int32(4422656)
												v190 = *(*int64)(unsafe.Add(mBase, _consts[643]))
												*(*int64)(unsafe.Add(mBase, _consts[643])) = v190 + v114
												v194 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v194)
												*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v194)
											}
										}
									} else {
									}
									v211 = int32(4424800)
									v212 = *(*int64)(unsafe.Add(mBase, _consts[644]))
									*(*int64)(unsafe.Add(mBase, _consts[644])) = v212 + base.I64_extend_i32_u(v95)
									v217 = int32(4423840)
									v218 = *(*int64)(unsafe.Add(mBase, _consts[645]))
									*(*int64)(unsafe.Add(mBase, _consts[645])) = v218 + v97
									F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
									mBase = m.M
									v223 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v223)
									*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v223)
									m.G0 = v103 + int32(16)
									m.G0 = v8 + int32(32)
									return
								} else {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, _consts[39])))
									if v68 != 0 {
										v69 = int32(21)
									} else {
										v69 = int32(23)
									}
									v71 = F_errstart(m, v69, int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										if v71 == int32(0) {
											v95 = int32(1)
											v97 = int64(0)
											v101 = m.G0
											v103 = v101 - int32(16)
											m.G0 = v103
											if v55 != v97 {
												F___clock_gettime(m, int32(1), v103)
												mBase = m.M
												v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
												v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
												v114 = v109 + (v110*int64(1000000000) - v55)
												v166 = int32(4425760)
												v167 = *(*int64)(unsafe.Add(mBase, _consts[642]))
												*(*int64)(unsafe.Add(mBase, _consts[642])) = v167 + v114
												v171 = *(*int32)(unsafe.Add(mBase, _consts[172]))
												if base.Ui32(int32(16)) < base.Ui32(v171) {
												} else {
													if int32(1)<<(uint(v171)%32)&int32(115186) == int32(0) {
													} else {
														v189 = int32(4422656)
														v190 = *(*int64)(unsafe.Add(mBase, _consts[643]))
														*(*int64)(unsafe.Add(mBase, _consts[643])) = v190 + v114
														v194 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v194)
														*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v194)
													}
												}
											} else {
											}
											v211 = int32(4424800)
											v212 = *(*int64)(unsafe.Add(mBase, _consts[644]))
											*(*int64)(unsafe.Add(mBase, _consts[644])) = v212 + base.I64_extend_i32_u(v95)
											v217 = int32(4423840)
											v218 = *(*int64)(unsafe.Add(mBase, _consts[645]))
											*(*int64)(unsafe.Add(mBase, _consts[645])) = v218 + v97
											F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
											mBase = m.M
											v223 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v223)
											*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v223)
											m.G0 = v103 + int32(16)
											m.G0 = v8 + int32(32)
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v79 = *(*int32)(unsafe.Add(mBase, _consts[194]))
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77*int32(48))+32))
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
												F_errmsg(m, int32(286138), v8)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													F_errfinish(m, int32(477940), int32(1530), int32(89451))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														v95 = int32(1)
														v97 = int64(0)
														v101 = m.G0
														v103 = v101 - int32(16)
														m.G0 = v103
														if v55 != v97 {
															F___clock_gettime(m, int32(1), v103)
															mBase = m.M
															v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
															v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
															v114 = v109 + (v110*int64(1000000000) - v55)
															v166 = int32(4425760)
															v167 = *(*int64)(unsafe.Add(mBase, _consts[642]))
															*(*int64)(unsafe.Add(mBase, _consts[642])) = v167 + v114
															v171 = *(*int32)(unsafe.Add(mBase, _consts[172]))
															if base.Ui32(int32(16)) < base.Ui32(v171) {
															} else {
																if int32(1)<<(uint(v171)%32)&int32(115186) == int32(0) {
																} else {
																	v189 = int32(4422656)
																	v190 = *(*int64)(unsafe.Add(mBase, _consts[643]))
																	*(*int64)(unsafe.Add(mBase, _consts[643])) = v190 + v114
																	v194 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v194)
																	*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v194)
																}
															}
														} else {
														}
														v211 = int32(4424800)
														v212 = *(*int64)(unsafe.Add(mBase, _consts[644]))
														*(*int64)(unsafe.Add(mBase, _consts[644])) = v212 + base.I64_extend_i32_u(v95)
														v217 = int32(4423840)
														v218 = *(*int64)(unsafe.Add(mBase, _consts[645]))
														*(*int64)(unsafe.Add(mBase, _consts[645])) = v218 + v97
														F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
														mBase = m.M
														v223 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v223)
														*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v223)
														m.G0 = v103 + int32(16)
														m.G0 = v8 + int32(32)
														return
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
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
					v44 = m.G0
					v46 = v44 - int32(16)
					m.G0 = v46
					if v41 != 0 {
						F___clock_gettime(m, int32(1), v46)
						mBase = m.M
						v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
						v55 = v50 + v51*int64(1000000000)
					} else {
						v55 = int64(0)
					}
					m.G0 = v46 + int32(16)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v61 = F_FileSync(m, v59, int32(167772182))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						if int32(0) <= v61 {
							v95 = int32(1)
							v97 = int64(0)
							v101 = m.G0
							v103 = v101 - int32(16)
							m.G0 = v103
							if v55 != v97 {
								F___clock_gettime(m, int32(1), v103)
								mBase = m.M
								v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
								v114 = v109 + (v110*int64(1000000000) - v55)
								v166 = int32(4425760)
								v167 = *(*int64)(unsafe.Add(mBase, _consts[642]))
								*(*int64)(unsafe.Add(mBase, _consts[642])) = v167 + v114
								v171 = *(*int32)(unsafe.Add(mBase, _consts[172]))
								if base.Ui32(int32(16)) < base.Ui32(v171) {
								} else {
									if int32(1)<<(uint(v171)%32)&int32(115186) == int32(0) {
									} else {
										v189 = int32(4422656)
										v190 = *(*int64)(unsafe.Add(mBase, _consts[643]))
										*(*int64)(unsafe.Add(mBase, _consts[643])) = v190 + v114
										v194 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v194)
										*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v194)
									}
								}
							} else {
							}
							v211 = int32(4424800)
							v212 = *(*int64)(unsafe.Add(mBase, _consts[644]))
							*(*int64)(unsafe.Add(mBase, _consts[644])) = v212 + base.I64_extend_i32_u(v95)
							v217 = int32(4423840)
							v218 = *(*int64)(unsafe.Add(mBase, _consts[645]))
							*(*int64)(unsafe.Add(mBase, _consts[645])) = v218 + v97
							F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
							mBase = m.M
							v223 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v223)
							*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v223)
							m.G0 = v103 + int32(16)
							m.G0 = v8 + int32(32)
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _consts[39])))
							if v68 != 0 {
								v69 = int32(21)
							} else {
								v69 = int32(23)
							}
							v71 = F_errstart(m, v69, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								if v71 == int32(0) {
									v95 = int32(1)
									v97 = int64(0)
									v101 = m.G0
									v103 = v101 - int32(16)
									m.G0 = v103
									if v55 != v97 {
										F___clock_gettime(m, int32(1), v103)
										mBase = m.M
										v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
										v114 = v109 + (v110*int64(1000000000) - v55)
										v166 = int32(4425760)
										v167 = *(*int64)(unsafe.Add(mBase, _consts[642]))
										*(*int64)(unsafe.Add(mBase, _consts[642])) = v167 + v114
										v171 = *(*int32)(unsafe.Add(mBase, _consts[172]))
										if base.Ui32(int32(16)) < base.Ui32(v171) {
										} else {
											if int32(1)<<(uint(v171)%32)&int32(115186) == int32(0) {
											} else {
												v189 = int32(4422656)
												v190 = *(*int64)(unsafe.Add(mBase, _consts[643]))
												*(*int64)(unsafe.Add(mBase, _consts[643])) = v190 + v114
												v194 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v194)
												*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v194)
											}
										}
									} else {
									}
									v211 = int32(4424800)
									v212 = *(*int64)(unsafe.Add(mBase, _consts[644]))
									*(*int64)(unsafe.Add(mBase, _consts[644])) = v212 + base.I64_extend_i32_u(v95)
									v217 = int32(4423840)
									v218 = *(*int64)(unsafe.Add(mBase, _consts[645]))
									*(*int64)(unsafe.Add(mBase, _consts[645])) = v218 + v97
									F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
									mBase = m.M
									v223 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v223)
									*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v223)
									m.G0 = v103 + int32(16)
									m.G0 = v8 + int32(32)
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v79 = *(*int32)(unsafe.Add(mBase, _consts[194]))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77*int32(48))+32))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
										F_errmsg(m, int32(286138), v8)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											F_errfinish(m, int32(477940), int32(1530), int32(89451))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v95 = int32(1)
												v97 = int64(0)
												v101 = m.G0
												v103 = v101 - int32(16)
												m.G0 = v103
												if v55 != v97 {
													F___clock_gettime(m, int32(1), v103)
													mBase = m.M
													v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
													v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
													v114 = v109 + (v110*int64(1000000000) - v55)
													v166 = int32(4425760)
													v167 = *(*int64)(unsafe.Add(mBase, _consts[642]))
													*(*int64)(unsafe.Add(mBase, _consts[642])) = v167 + v114
													v171 = *(*int32)(unsafe.Add(mBase, _consts[172]))
													if base.Ui32(int32(16)) < base.Ui32(v171) {
													} else {
														if int32(1)<<(uint(v171)%32)&int32(115186) == int32(0) {
														} else {
															v189 = int32(4422656)
															v190 = *(*int64)(unsafe.Add(mBase, _consts[643]))
															*(*int64)(unsafe.Add(mBase, _consts[643])) = v190 + v114
															v194 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v194)
															*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v194)
														}
													}
												} else {
												}
												v211 = int32(4424800)
												v212 = *(*int64)(unsafe.Add(mBase, _consts[644]))
												*(*int64)(unsafe.Add(mBase, _consts[644])) = v212 + base.I64_extend_i32_u(v95)
												v217 = int32(4423840)
												v218 = *(*int64)(unsafe.Add(mBase, _consts[645]))
												*(*int64)(unsafe.Add(mBase, _consts[645])) = v218 + v97
												F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
												mBase = m.M
												v223 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v223)
												*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v223)
												m.G0 = v103 + int32(16)
												m.G0 = v8 + int32(32)
												return
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
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_regoperout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v198
L2:
	;
	v15 = F_pstrdup(m, int32(532178))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v20 = F_SearchSysCache1(m, int32(40), v11)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v198 = v15
	goto L1
L7:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v24 = v22 + v23
	v26 = v24 + int32(4)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v28 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v191 = F_palloc(m, int32(64))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L64
	}
L11:
	;
	v31 = F_pstrdup(m, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v35 = F_makeString(m, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v198 = v31
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v35
	v42 = F_list_make1_impl(m, int32(1), v9+int32(24))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v58 = F_get_namespace_name(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L25
	}
L18:
	;
	v44 = int32(0)
	v46 = F_OpernameGetCandidates(m, v42, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	if v46 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v51 != v11 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v53 = F_pstrdup(m, v26)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v198 = v53
	goto L1
L25:
	;
	v60 = F_quote_identifier(m, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v60&int32(3) == int32(0) {
		v85 = v60
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v26&int32(3) == int32(0) {
		v142 = v26
		goto L46
	} else {
		goto L47
	}
L28:
	;
	v118 = v110 - v60
	goto L27
L29:
	;
	v89 = v85
	goto L38
L30:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v69 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v118 = int32(0)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v74 = v60
	goto L34
L34:
	;
	v78 = v74 + int32(1)
	if v78&int32(3) == int32(0) {
		v85 = v78
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v110 = v78
	goto L28
L36:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v83 != 0 {
		v74 = v78
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v98 = int32(-2139062144)
	if (int32(16843008)-v95|v95)&v98 == v98 {
		v89 = v89 + int32(4)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v104 = v89
	goto L41
L40:
	;
	goto L39
L41:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 != 0 {
		v104 = v104 + int32(1)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v110 = v104
	goto L28
L43:
	;
	goto L42
L44:
	;
	v179 = F_palloc(m, v118+v175+int32(2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L61
	}
L45:
	;
	v175 = v167 - v26
	goto L44
L46:
	;
	v146 = v142
	goto L55
L47:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v126 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = int32(0)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v131 = v26
	goto L51
L51:
	;
	v135 = v131 + int32(1)
	if v135&int32(3) == int32(0) {
		v142 = v135
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v167 = v135
	goto L45
L53:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v140 != 0 {
		v131 = v135
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v155 = int32(-2139062144)
	if (int32(16843008)-v152|v152)&v155 == v155 {
		v146 = v146 + int32(4)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v161 = v146
	goto L58
L57:
	;
	goto L56
L58:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v165 != 0 {
		v161 = v161 + int32(1)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v167 = v161
	goto L45
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v60
	v186 = F_pg_sprintf(m, v179, int32(167685), v9+int32(16))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	v198 = v179
	goto L1
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
	v196 = F_pg_snprintf(m, v191, int32(64), int32(57094), v9)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v198 = v191
	goto L1
}
func F_regprocout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 == int32(0) {
		v15 = F_pstrdup(m, int32(630250))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v74 = v15
			m.G0 = v9 + int32(16)
			return v74
		}
	} else {
		v20 = F_SearchSysCache1(m, int32(47), v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
				v24 = v22 + v23
				v26 = v24 + int32(4)
				v28 = *(*int32)(unsafe.Add(mBase, _consts[205]))
				if v28 == int32(0) {
					v31 = F_pstrdup(m, v26)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v20)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v74 = v31
							m.G0 = v9 + int32(16)
							return v74
						}
					}
				} else {
					v35 = F_makeString(m, v26)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v35
						v42 = F_list_make1_impl(m, int32(1), v9+int32(8))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v45 = int32(0)
							v50 = F_FuncnameGetCandidates(m, v42, int32(-1), v45, v45, v45, v45, v45)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
									v59 = F_get_namespace_name(m, v58)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = v59
										v62 = F_quote_qualified_identifier(m, v61, v26)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v20)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v74 = v62
												m.G0 = v9 + int32(16)
												return v74
											}
										}
									}
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									if v54 != 0 {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
										v59 = F_get_namespace_name(m, v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = v59
											v62 = F_quote_qualified_identifier(m, v61, v26)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v74 = v62
													m.G0 = v9 + int32(16)
													return v74
												}
											}
										}
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
										if v56 == v11 {
											v61 = int32(0)
											v62 = F_quote_qualified_identifier(m, v61, v26)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v74 = v62
													m.G0 = v9 + int32(16)
													return v74
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
											v59 = F_get_namespace_name(m, v58)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v61 = v59
												v62 = F_quote_qualified_identifier(m, v61, v26)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v20)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														v74 = v62
														m.G0 = v9 + int32(16)
														return v74
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
				v67 = F_palloc(m, int32(64))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					v72 = F_pg_snprintf(m, v67, int32(64), int32(57094), v9)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = v67
						m.G0 = v9 + int32(16)
						return v74
					}
				}
			}
		}
	}
}
func F_regprocrecv(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_int4recv(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regtypeout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v9 == int32(0) {
		v13 = F_pstrdup(m, int32(630250))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v45 = v13
			m.G0 = v7 + int32(16)
			return v45
		}
	} else {
		v18 = F_SearchSysCache1(m, int32(82), v9)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[205]))
				if v21 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
					v29 = F_pstrdup(m, v24+v25+int32(4))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v18)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v45 = v29
							m.G0 = v7 + int32(16)
							return v45
						}
					}
				} else {
					v33 = F_format_type_be(m, v9)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v18)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v45 = v33
							m.G0 = v7 + int32(16)
							return v45
						}
					}
				}
			} else {
				v38 = F_palloc(m, int32(64))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
					v43 = F_pg_snprintf(m, v38, int32(64), int32(57094), v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v38
						m.G0 = v7 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_relmap_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	if base.Ui32(l0) < base.Ui32(int32(16)) {
		v6 = int32(516448)
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_relmap_redo(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v4 = m.G0
	v6 = v4 - int32(544)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+48)))
	v11 = v9 & int32(240)
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		if v15 != int32(524) {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v75
				F_errmsg_internal(m, int32(402333), v6)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					F_errfinish(m, int32(473789), int32(1111), int32(231177))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v23 = F__emscripten_memcpy_bulkmem(m, v6+int32(20), v14+int32(12), int32(524))
			mBase = m.M
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v27 = F_GetDatabasePath(m, v25, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[29]))
				v34 = F_LWLockAcquire(m, v30+int32(3200), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v38 = int32(0)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					F_write_relmap_file(m, v6+int32(20), v38, int32(1), v38, v41, v42, v27)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _consts[29]))
						F_LWLockRelease(m, v46+int32(3200))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								m.G0 = v6 + int32(544)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v11
			F_errmsg_internal(m, int32(50143), v6+int32(16))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errfinish(m, int32(473789), int32(1140), int32(231177))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_remove_dbtablespaces(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v17 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(112)
	return
L2:
	;
	v105 = v85 << (uint(int32(2)) % 32)
	v106 = F_palloc(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L34
	}
L3:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+188))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	m.T0[v98].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L32
	}
L4:
	;
	return
L5:
	;
	v19 = int32(0)
	v21 = F_table_beginscan_catalog(m, v17, v19, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = F_heap_getnext(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v28 = v23
	v30 = int32(0)
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)))
	if v40 != int32(1664) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v79 == int32(0) {
		goto L3
	} else {
		goto L30
	}
L11:
	;
	v43 = F_GetDatabasePath(m, l0, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L15
	}
L12:
	;
	v79 = v30
	goto L13
L13:
	;
	v81 = F_heap_getnext(m, v21)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L28
	}
L14:
	;
	F_pfree(m, v43)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L27
	}
L15:
	;
	v49 = F___fstatat(m, int32(-100), v43, v13+int32(16), int32(256))
	mBase = m.M
	goto L16
L16:
	;
	if v49 < int32(0) {
		v76 = v30
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v52&int32(61440) != int32(16384) {
		v76 = v30
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v57 = F_rmtree(m, v43)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v74 = F_lappend_oid(m, v30, v40)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L26
	}
L20:
	;
	if v57 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v61 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v61 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v43
	F_errmsg(m, int32(652328), v13)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(472768), int32(3040), int32(162341))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v76 = v74
	goto L14
L27:
	;
	v79 = v76
	goto L13
L28:
	;
	if v81 != 0 {
		v28 = v81
		v30 = v79
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L10
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v85 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L3
L32:
	;
	F_sequence_close(m, v17, int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L1
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if int32(0) < v108 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v113 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	F_XLogBeginInsert(m)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L41
	}
L38:
	;
	v123 = v113 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125+v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v106+v123))) = v127
	v130 = v113 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v130 < v131 {
		v113 = v130
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	goto L39
L41:
	;
	F_XLogRegisterData(m, v13+int32(16), int32(8))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_XLogRegisterData(m, v106, v105)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v156 = F_XLogInsert(m, int32(4), int32(33))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_list_free(m, v79)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_pfree(m, v106)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+188))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	m.T0[v164].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_sequence_close(m, v17, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L1
}
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v16 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1052), v7+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v16
	}
}
func F_rename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = int32(-100)
	v5 = m.Env.X__syscall_renameat(m, v3, l0, v3, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v5) {
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0) - v5
		v13 = int32(-1)
	} else {
		v13 = v5
	}
	return v13
}
func F_rename_constraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L62
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L58
	}
L3:
	;
	v36 = F_SearchSysCache1(m, int32(19), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L12
	}
L4:
	;
	v22 = F_get_domain_constraint_oid(m, l2, l3, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = F_relation_open(m, l1, int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v34 = int32(0)
	v35 = v22
	goto L3
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	F_renameatt_check(m, l1, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v32 = F_get_relation_constraint_oid(m, l1, l3, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v34 = v25
	v35 = v32
	goto L3
L12:
	;
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v40 = v38 + v39
	if l1 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L55
	}
L16:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v40)+88))
	if v138 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+72)))
	switch v43 - int32(99) {
	case 0, 11:
		goto L18
	default:
		goto L16
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+106)))
	if v46 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if l5 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+104)))
	if l6 < v121 {
		goto L1
	} else {
		goto L41
	}
L21:
	;
	v50 = F_find_all_inheritors(m, l1, int32(8), v16+int32(-4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if l6 != 0 {
		goto L20
	} else {
		goto L38
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v56 = int32(0)
	goto L25
L25:
	;
	v69 = int32(0)
	if v50 == v69 {
		v79 = v69
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v52 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v73 <= v56 {
		v79 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v79 = v75 + v56<<(uint(int32(2))%32)
	goto L27
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v82 <= v56 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	if v79 == int32(0) {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v89 = v86 + v56<<(uint(int32(2))%32)
	if v89 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if l1 != v92 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v96 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	F_rename_constraint_internal(m, v16+int32(-16), v92, v96, l3, l4, v96, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v56 = v56 + int32(1)
	goto L25
L37:
	;
	goto L36
L38:
	;
	v104 = F_find_inheritance_children(m, l1, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v104 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	goto L20
L41:
	;
	goto L16
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L49
	}
L43:
	;
	F_RenameConstraintById(m, v35, l4)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L48
	}
L44:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+72)))
	v143 = v141 - int32(112)
	if base.Ui32(int32(8)) < base.Ui32(v143) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	if int32(1)<<(uint(v143)%32)&int32(289) == int32(0) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	F_RenameRelationInternal(m, v138, l4, int32(0), int32(1))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	goto L42
L49:
	;
	if v34 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_CacheInvalidateRelcache(m, v34)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	m.G0 = v18 - int32(-64)
	return
L53:
	;
	F_relation_close(m, v34, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v35
	F_errmsg_internal(m, int32(38745), v18)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(473018), int32(4083), int32(296655))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l3
	F_errmsg(m, int32(229276), v16+int32(-32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(473018), int32(4119), int32(296655))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l3
	F_errmsg(m, int32(659193), v16+int32(-48))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(473018), int32(4126), int32(296655))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_reorderqueue_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v7 = F_pairingheap_remove_first(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = int32(0)
	v18 = v12
	goto L6
L4:
	;
	goto L5
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_pfree(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v17))))
	if v22 != 0 {
		v34 = v18
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v36 = v17 + int32(1)
	if v36 < v34 {
		v17 = v36
		v18 = v34
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v17))))
	if v25 != 0 {
		v34 = v18
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v17<<(uint(int32(2))%32))))
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v34 = v33
	goto L8
L12:
	;
	goto L7
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	F_pfree(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_pfree(m, v7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	return v11
}
func F_repalloc0(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(l2) < base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(34980), v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(471226), int32(1655), int32(531332))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v29&int32(15)*int32(36))+uint32(_consts[983])))
		v37 = m.T0[v36].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v43 = F__emscripten_memset_bulkmem(m, v37+l1, base.I32_extend8_s(int32(0)), l2-l1)
			mBase = m.M
			m.G0 = v7 + int32(16)
			return v37
		}
	}
}
func F_repalloc_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_consts[983])))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		return v14
	}
}
func F_repalloc_huge(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_consts[983])))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		return v14
	}
}
func F_reparameterize_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v162 int32
	_ = v162
	var v163 float64
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 float64
	_ = v309
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 float64
	_ = v320
	var v324 float64
	_ = v324
	var v332 float64
	_ = v332
	var v338 float64
	_ = v338
	var v344 float64
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 float64
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v396 float64
	_ = v396
	var v404 float64
	_ = v404
	var v408 float64
	_ = v408
	var v412 int32
	_ = v412
	var v414 float64
	_ = v414
	var v416 float64
	_ = v416
	var v419 float64
	_ = v419
	var v422 float64
	_ = v422
	var v428 int32
	_ = v428
	v5 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = v13
	goto L3
L2:
	;
	v14 = v5
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = int32(0)
	if v14 == v16 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v428
L5:
	;
	if v69 == int32(0) {
		v428 = v5
		goto L4
	} else {
		goto L19
	}
L6:
	;
	v69 = int32(1)
	goto L5
L7:
	;
	goto L8
L8:
	;
	if l2 == int32(0) {
		v60 = v16
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v69 = v60
	goto L5
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v26 < v25 {
		v60 = v16
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v28 = int32(1)
	if v25 <= v28 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = v28
	goto L14
L13:
	;
	v31 = v25
	goto L14
L14:
	;
	v32 = int32(8)
	v37 = int32(0)
	goto L15
L15:
	;
	v44 = v37 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14+v32+v44)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+(l2+v32))))
	v51 = v46 & (v48 ^ int32(-1))
	v53 = base.B2i32(v51 == int32(0))
	if v51 != 0 {
		v60 = v53
		goto L9
	} else {
		goto L17
	}
L16:
	;
	v60 = v53
	goto L9
L17:
	;
	v55 = v37 + int32(1)
	if v55 != v31 {
		v37 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v72 - int32(331) {
	case 0:
		goto L23
	default:
		v428 = v5
		goto L4
	case 3:
		goto L22
	case 8:
		goto L28
	case 9:
		goto L27
	case 10, 11:
		goto L26
	case 13:
		goto L25
	case 16:
		goto L24
	case 29:
		goto L21
	case 30:
		goto L20
	}
L20:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v356 = F_reparameterize_path(m, l0, v355, l2, l3)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L29
	} else {
		goto L87
	}
L21:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v278 = F_reparameterize_path(m, l0, v277, l2, l3)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L29
	} else {
		goto L77
	}
L22:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v218 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L23:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v193 != int32(279) {
		v428 = v5
		goto L4
	} else {
		goto L52
	}
L24:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v162)+56))
	v165 = F_palloc0(m, int32(80))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L29
	} else {
		goto L46
	}
L25:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v138 = F_palloc0(m, int32(80))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L29
	} else {
		goto L43
	}
L26:
	;
	v122 = F_palloc0(m, int32(112))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L36
	}
L27:
	;
	v100 = F_palloc0(m, int32(72))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L29
	} else {
		goto L33
	}
L28:
	;
	v76 = F_palloc0(m, int32(72))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = int64(1455993913623)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v83
	v85 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+20)) = uint8(v87)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v85
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+64)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = v87
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+21)) = uint8(v90)
	F_cost_seqscan(m, v76, l0, v15, v85)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	return v76
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v100))) = int64(1460288880919)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v105
	v107 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+20)) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+16)) = v107
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+64)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v100)+24)) = v109
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+21)) = uint8(v112)
	F_cost_samplescan(m, v100, l0, v15, v107)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	return v100
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(280)
	goto L38
L37:
	;
	v129 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L29
	} else {
		goto L41
	}
L38:
	;
	v127 = F__emscripten_memcpy_bulkmem(m, v122, l1, int32(112))
	mBase = m.M
	goto L40
L40:
	;
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v129
	F_cost_index(m, v127, l0, l3, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	return v127
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = int64(1477468750106)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+12)) = v143
	v145 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+20)) = uint8(v147)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+16)) = v145
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+72)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v138)+64)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v138)+24)) = v147
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+21)) = uint8(v150)
	F_cost_bitmap_heap_scan(m, v138, l0, v15, v145, v136, l3)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	return v138
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(1490353651999)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v170
	v172 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)) = uint8(v174)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v172
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v177 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+21)))
	v181 = v180
	goto L50
L49:
	;
	v181 = v5
	goto L50
L50:
	;
	v183 = v181 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+21)) = uint8(v183)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v162)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+72)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v165)+64)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = v185
	F_cost_subqueryscan(m, v165, l0, v15, v172, base.F64_eq(v161, v163))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	return v165
L52:
	;
	v197 = F_palloc0(m, int32(72))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L29
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(1421634175255)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = v202
	v204 = F_get_baserel_parampathinfo(m, l0, v15, l2)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v204
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+64)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v197)+24)) = v206
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+21)) = uint8(v209)
	F_cost_resultscan(m, v197, l0, v15, v204)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	return v197
L56:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v274 = F_create_append_path(m, l0, v15, v268, v267, v270, l2, v271, v272, float64(-1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L29
	} else {
		goto L76
	}
L57:
	;
	v267 = v5
	v268 = v5
	goto L56
L58:
	;
	goto L59
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v221 <= int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v267 = v5
	v268 = v5
	goto L56
L61:
	;
	goto L62
L62:
	;
	v231 = int32(0)
	v233 = v5
	v234 = v5
	goto L63
L63:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236+v231<<(uint(int32(2))%32))))
	v241 = F_reparameterize_path(m, l0, v240, l2, l3)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L29
	} else {
		goto L65
	}
L64:
	;
	v267 = v253
	v268 = v254
	goto L56
L65:
	;
	if v241 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	return int32(0)
L67:
	;
	goto L68
L68:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v231 < v247 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v256 = v231 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v256 < v257 {
		v231 = v256
		v233 = v253
		v234 = v254
		goto L63
	} else {
		goto L75
	}
L70:
	;
	v249 = F_lappend(m, v234, v241)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L29
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v251 = F_lappend(m, v233, v241)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L29
	} else {
		goto L74
	}
L73:
	;
	v253 = v233
	v254 = v249
	goto L69
L74:
	;
	v253 = v251
	v254 = v234
	goto L69
L75:
	;
	goto L64
L76:
	;
	return v274
L77:
	;
	if v278 == int32(0) {
		v428 = v5
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v283 = F_palloc0(m, int32(80))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L29
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v283))) = int64(1546188226853)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v291 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+20)) = uint8(v291)
	*(*int32)(unsafe.Add(mBase, uint32(v283)+16)) = v290
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v295 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+21)))
	v299 = v298
	goto L82
L81:
	;
	v299 = v291
	goto L82
L82:
	;
	v301 = v299 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+21)) = uint8(v301)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v278)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+24)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v278)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+72)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v283)+64)) = v305
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v278)+40))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v278)+48))
	v310 = *(*float64)(unsafe.Add(mBase, uint32(v278)+56))
	v311 = *(*float64)(unsafe.Add(mBase, uint32(v278)+32))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+32))
	v317 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	*(*float64)(unsafe.Add(mBase, uint32(v283)+32)) = v311
	v320 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v324 = base.F64_add(base.F64_mul(base.F64_add(v320, v320), v311), base.F64_sub(v310, v309))
	v332 = base.F64_mul(v311, base.F64_convert_i32_u((v313+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v332, base.F64_convert_i32_u(v317<<(uint(int32(10))%32))) != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	return v283
L84:
	;
	v338 = *(*float64)(unsafe.Add(mBase, _consts[386]))
	v344 = base.F64_add(base.F64_mul(v338, base.F64_ceil(base.F64_mul(v332, float64(0.0001220703125)))), v324)
	goto L86
L85:
	;
	v344 = v324
	goto L86
L86:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, _consts[390])))
	*(*float64)(unsafe.Add(mBase, uint32(v283)+56)) = base.F64_add(v309, v344)
	*(*float64)(unsafe.Add(mBase, uint32(v283)+48)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v283)+40)) = v308 + (v346 ^ int32(1))
	goto L83
L87:
	;
	if v356 == int32(0) {
		v428 = v5
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v360 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+85)))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v366 = F_palloc0(m, int32(104))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L29
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v366))) = int64(1550483194150)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v371
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v356)+16))
	v374 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+20)) = uint8(v374)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+16)) = v373
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+26)))
	if v378 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+21)))
	v382 = v381
	goto L92
L91:
	;
	v382 = v374
	goto L92
L92:
	;
	v384 = v382 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+21)) = uint8(v384)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v356)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v356)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+85)) = uint8(v361)
	*(*uint8)(unsafe.Add(mBase, uint32(v366)+84)) = uint8(v362)
	*(*int32)(unsafe.Add(mBase, uint32(v366)+80)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v366)+76)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v366)+72)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v366)+64)) = v388
	v396 = float64(1e+100)
	if base.F64_gt(v360, v396) != 0 {
		v408 = v396
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+96)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v366)+88)) = v408
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v356)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+40)) = v412
	v414 = *(*float64)(unsafe.Add(mBase, uint32(v356)+48))
	v416 = *(*float64)(unsafe.Add(mBase, _consts[380]))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+48)) = base.F64_add(v414, v416)
	v419 = *(*float64)(unsafe.Add(mBase, uint32(v356)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+56)) = base.F64_add(v416, v419)
	v422 = *(*float64)(unsafe.Add(mBase, uint32(v356)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v366)+32)) = v422
	v428 = v366
	goto L4
L94:
	;
	goto L93
L95:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v360)&int64(9223372036854775807)) {
		v408 = v396
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v404 = float64(1)
	if base.F64_le(v360, v404) != 0 {
		v408 = v404
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v408 = base.F64_nearest(v360)
	goto L94
}
func F_replace_vars_in_jointree(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(63) {
	case 0:
		goto L7
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L4
	}
L3:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v134 = F_replace_rte_variables(m, v130, v16, int32(0), int32(850), l1, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L36
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L14
	} else {
		goto L33
	}
L5:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_replace_vars_in_jointree(m, v96, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L27
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v60 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v15 == v16 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v15<<(uint(int32(2))%32)-int32(4))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+124)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	switch v31 {
	case 0:
		goto L3
	case 1:
		goto L13
	default:
		goto L1
	case 3:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v57 = F_replace_rte_variables(m, v53, v16, int32(0), int32(850), l1, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L18
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v50 = F_replace_rte_variables(m, v46, v16, int32(0), int32(850), l1, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L17
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v43 = F_replace_rte_variables(m, v39, v16, int32(0), int32(850), l1, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	v36 = F_replace_rte_variables(m, v32, v16, int32(1), int32(850), l1, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v36
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v43
	goto L1
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v50
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v57
	goto L1
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v92 = F_replace_rte_variables(m, v87, v88, int32(0), int32(850), l1, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L26
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v68 = v3
	goto L22
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v68<<(uint(int32(2))%32))))
	F_replace_vars_in_jointree(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v79 = v68 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v79 < v80 {
		v68 = v79
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92
	goto L1
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_replace_vars_in_jointree(m, v99, l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(2)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v112 = F_replace_rte_variables(m, v107, v108, int32(0), int32(850), l1, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v95
	goto L1
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v120
	F_errmsg_internal(m, int32(465033), v8)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(477665), int32(2613), int32(391871))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v134
	goto L1
}
func F_report_invalid_encoding_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l2 < l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = l2
	goto L3
L2:
	;
	v13 = l3
	goto L3
L3:
	;
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = int32(8)
	if v16 <= v13 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L19
	}
L7:
	;
	v19 = v16
	goto L9
L8:
	;
	v19 = v13
	goto L9
L9:
	;
	v27 = int32(0)
	v28 = v10 + int32(32)
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v27))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v33
	v38 = F_pg_sprintf(m, v28, int32(27622), v10+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L6
L12:
	;
	return
L13:
	;
	v40 = v38 + v28
	if v27 < v19-int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = F_pg_sprintf(m, v40, int32(706002), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v47 = v40
	goto L16
L16:
	;
	v49 = v27 + int32(1)
	if v49 != v19 {
		v27 = v49
		v28 = v47
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v47 = v44 + v40
	goto L16
L18:
	;
	goto L11
L19:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_consts[769])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v10 + int32(32)
	F_errmsg(m, int32(195929), v10)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(472485), int32(1853), int32(86828))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_report_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 float64
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 float64
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v16 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(80)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 <= int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = v16
	v32 = v4
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v40 = v37 + v32*int32(416)
	F_InstrEndLoop(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	return
L8:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	if base.F64_ne(v43, float64(0)) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v46 = int32(0)
	F_ExplainOpenGroup(m, int32(214413), v46, int32(1), l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v160 = v32 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v160 < v162 {
		v30 = v161
		v32 = v160
		goto L5
	} else {
		goto L55
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v56 = v36 + v32*int32(60)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	if v57 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = F_get_constraint_name(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v60 = v46
	goto L15
L15:
	;
	v62 = v53 + int32(4)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v60 = v58
	goto L15
L17:
	;
	if v60 != 0 {
		goto L50
	} else {
		goto L51
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v67 != 0 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	goto L20
L20:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	F_ExplainPropertyText(m, int32(364901), v119, l2)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L39
	}
L21:
	;
	if l1 != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v60
	F_appendStringInfo(m, v82, int32(169180), v14+int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L29
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v73
	F_appendStringInfo(m, v66, int32(172677), v14-int32(-64))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L27
	}
L24:
	;
	if v60 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_appendStringInfoString(m, v66, int32(214413))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	goto L21
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v62
	F_appendStringInfo(m, v90, int32(174887), v14+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v99 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v40)+208))
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v103
	*(*float64)(unsafe.Add(mBase, uint32(v14))) = base.F64_mul(v102, float64(1000))
	F_appendStringInfo(m, v98, int32(709238), v14)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = v111
	F_appendStringInfo(m, v98, int32(709262), v14+int32(16))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L38
	}
L37:
	;
	goto L17
L38:
	;
	goto L17
L39:
	;
	if v60 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_ExplainPropertyText(m, int32(364885), v60, l2)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_ExplainPropertyText(m, int32(252751), v62, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v128 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v40)+208))
	F_ExplainPropertyFloat(m, int32(359613), int32(142590), base.F64_mul(v133, float64(1000)), int32(3), l2)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v140 = int32(0)
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v40)+216))
	F_ExplainPropertyFloat(m, int32(143919), v140, v141, v140, l2)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	goto L17
L50:
	;
	F_pfree(m, v60)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_ExplainCloseGroup(m, int32(214413), int32(1), l2)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	goto L11
L55:
	;
	goto L6
}
func F_reserveAllocatedDesc(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, _consts[593]))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[585]))
	if v9 < v7 {
		v56 = int32(1)
		return v56
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[586]))
		if v12 == int32(0) {
			v17 = F_emscripten_builtin_malloc(m, int32(192))
			mBase = m.M
			if v17 != 0 {
				v47 = v17
				v49 = int32(16)
				*(*int32)(unsafe.Add(mBase, _consts[593])) = v49
				*(*int32)(unsafe.Add(mBase, _consts[586])) = v47
				v56 = int32(1)
				return v56
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(12890), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(478009), int32(2597), int32(467694))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
		} else {
			v36 = int32(0)
			v38 = *(*int32)(unsafe.Add(mBase, _consts[591]))
			v40 = base.I32_div_s(v38, int32(3))
			if v40 <= v7 {
				v56 = v36
			} else {
				v44 = F_emscripten_builtin_realloc(m, v12, v40*int32(12))
				mBase = m.M
				if v44 == int32(0) {
					v56 = v36
				} else {
					v47 = v44
					v49 = v40
					*(*int32)(unsafe.Add(mBase, _consts[593])) = v49
					*(*int32)(unsafe.Add(mBase, _consts[586])) = v47
					v56 = int32(1)
				}
			}
			return v56
		}
	}
}
func F_resolve_anymultirange_from_others(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = F_getBaseType(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = F_get_range_multirange(m, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				if v11 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v40 = F_format_type_be(m, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v40
								F_errmsg(m, int32(183839), v6)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_errfinish(m, int32(476064), int32(726), int32(125542))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v11
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(354256), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(476064), int32(730), int32(125542))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_restrict_and_check_grant(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v189 int64
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int64
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	v14 = m.G0
	v16 = v14 - int32(192)
	m.G0 = v16
	switch l6 - int32(6) {
	case 0:
		v58 = int64(167503724583)
		goto L1
	default:
		goto L3
	case 3:
		goto L11
	case 8:
		goto L5
	case 10, 11, 15, 43:
		goto L9
	case 13:
		goto L10
	case 16:
		goto L8
	case 21:
		goto L4
	case 30:
		goto L7
	case 31:
		goto L12
	case 35:
		goto L2
	case 36:
		goto L6
	}
L1:
	;
	if l1 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v58 = int64(70914205040767)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L17
	}
L4:
	;
	v58 = int64(52776558145536)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v58 = int64(2199023256064)
	goto L1
L7:
	;
	v58 = int64(3298534884096)
	goto L1
L8:
	;
	v58 = int64(25769803782)
	goto L1
L9:
	;
	v58 = int64(1099511628032)
	goto L1
L10:
	;
	v58 = int64(549755814016)
	goto L1
L11:
	;
	v58 = int64(15393162792448)
	goto L1
L12:
	;
	v58 = int64(1125281431814)
	goto L1
L13:
	;
	return int64(0)
L14:
	;
	F_errmsg_internal(m, int32(125664), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(475972), int32(285), int32(91259))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l6
	F_errmsg_internal(m, int32(463852), v16)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(475972), int32(295), int32(91259))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L144
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L13
	} else {
		goto L140
	}
L22:
	;
	v206 = l3 & int64(base.Ui64(l1)>>(uint(int64(32))%64))
	if l0 != 0 {
		goto L86
	} else {
		goto L87
	}
L23:
	;
	switch l6 - int32(6) {
	case 0:
		goto L25
	default:
		goto L26
	case 3:
		goto L38
	case 8:
		goto L28
	case 10:
		goto L30
	case 11:
		goto L29
	case 13:
		goto L37
	case 15:
		goto L36
	case 16:
		goto L35
	case 21:
		goto L34
	case 30:
		goto L33
	case 31, 35:
		goto L39
	case 33:
		goto L32
	case 36:
		goto L31
	case 43:
		goto L27
	}
L24:
	;
	if v189 != int64(0) {
		goto L22
	} else {
		goto L78
	}
L25:
	;
	v181 = F_pg_class_aclmask_ext(m, l4, l5, v58, int32(1), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L76
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L73
	}
L27:
	;
	v162 = F_object_aclmask_ext(m, int32(1247), l4, l5, v58, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L72
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L69
	}
L29:
	;
	v145 = F_object_aclmask_ext(m, int32(1417), l4, l5, v58, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L68
	}
L30:
	;
	v141 = F_object_aclmask_ext(m, int32(2328), l4, l5, v58, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L67
	}
L31:
	;
	v137 = F_object_aclmask_ext(m, int32(1213), l4, l5, v58, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L66
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L63
	}
L33:
	;
	v120 = F_object_aclmask_ext(m, int32(2615), l4, l5, v58, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L62
	}
L34:
	;
	v82 = F_superuser_arg(m, l5)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L45
	}
L35:
	;
	v80 = F_pg_largeobject_aclmask_snapshot(m, l4, l5, v58, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L44
	}
L36:
	;
	v77 = F_object_aclmask_ext(m, int32(2612), l4, l5, v58, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L43
	}
L37:
	;
	v73 = F_object_aclmask_ext(m, int32(1255), l4, l5, v58, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L42
	}
L38:
	;
	v69 = F_object_aclmask_ext(m, int32(1262), l4, l5, v58, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L41
	}
L39:
	;
	v65 = F_pg_class_aclmask_ext(m, l4, l5, v58, int32(1), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v189 = v65
	goto L24
L41:
	;
	v189 = v69
	goto L24
L42:
	;
	v189 = v73
	goto L24
L43:
	;
	v189 = v77
	goto L24
L44:
	;
	v189 = v80
	goto L24
L45:
	;
	if v82 != 0 {
		v189 = v58
		goto L24
	} else {
		goto L46
	}
L46:
	;
	v85 = F_SearchSysCache1(m, int32(44), l4)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v85 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v93 = F_SysCacheGetAttr(m, int32(44), v85, int32(3), v16+int32(191))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+191)))
	if v95 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v109 = F_aclmask(m, v106, l5, int32(10), v58, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L56
	}
L51:
	;
	v101 = F_acldefault(m, int32(27), int32(10))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v103 = F_pg_detoast_datum(m, v93)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L55
	}
L54:
	;
	v105 = int32(0)
	v106 = v101
	goto L50
L55:
	;
	v105 = v93
	v106 = v103
	goto L50
L56:
	;
	if v106 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_ReleaseCatCache(m, v85)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L13
	} else {
		goto L61
	}
L58:
	;
	if v106 == v105 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v106)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v189 = v109
	goto L24
L62:
	;
	v189 = v120
	goto L24
L63:
	;
	F_errmsg_internal(m, int32(117597), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(475972), int32(3012), int32(300809))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
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
	v189 = v137
	goto L24
L67:
	;
	v189 = v141
	goto L24
L68:
	;
	v189 = v145
	goto L24
L69:
	;
	F_errmsg_internal(m, int32(125664), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(475972), int32(3022), int32(300809))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
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
	v189 = v162
	goto L24
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l6
	F_errmsg_internal(m, int32(463852), v16+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(475972), int32(3029), int32(300809))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v184 = F_pg_attribute_aclmask_ext(m, l4, l8, l5, v58, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v189 = v181 | v184
	goto L24
L78:
	;
	if l9 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v196 = base.B2i32(l6 == int32(6))
	goto L81
L80:
	;
	v196 = int32(0)
	goto L81
L81:
	;
	if v196 != 0 {
		goto L20
	} else {
		goto L82
	}
L82:
	;
	F_aclcheck_error(m, int32(1), l6, l7)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	goto L22
L84:
	;
	m.G0 = v16 + int32(192)
	return v206
L85:
	;
	F_errfinish(m, int32(475972), v347, int32(91259))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L13
	} else {
		goto L139
	}
L86:
	;
	if v206 == int64(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if v206 == int64(0) {
		goto L114
	} else {
		goto L115
	}
L89:
	;
	v212 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if l2 != 0 {
		goto L84
	} else {
		goto L102
	}
L92:
	;
	if l6 != int32(6) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v212 == int32(0) {
		goto L84
	} else {
		goto L99
	}
L94:
	;
	if l9 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	if v212 == int32(0) {
		goto L84
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l9
	F_errmsg(m, int32(668137), v16+int32(48))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v347 = int32(334)
	goto L85
L99:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l7
	F_errmsg(m, int32(660292), v16-int32(-64))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v347 = int32(339)
	goto L85
L102:
	;
	if l3 == v206 {
		goto L84
	} else {
		goto L103
	}
L103:
	;
	v246 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	if l6 != int32(6) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v246 == int32(0) {
		goto L84
	} else {
		goto L111
	}
L106:
	;
	if l9 == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	if v246 == int32(0) {
		goto L84
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l9
	F_errmsg(m, int32(668197), v16+int32(80))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	v347 = int32(347)
	goto L85
L111:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = l7
	F_errmsg(m, int32(660328), v16+int32(96))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	v347 = int32(352)
	goto L85
L114:
	;
	v281 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L13
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if l2 != 0 {
		goto L84
	} else {
		goto L127
	}
L117:
	;
	if l6 != int32(6) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v281 == int32(0) {
		goto L84
	} else {
		goto L124
	}
L119:
	;
	if l9 == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	if v281 == int32(0) {
		goto L84
	} else {
		goto L121
	}
L121:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l9
	F_errmsg(m, int32(668262), v16+int32(112))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	v347 = int32(363)
	goto L85
L124:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = l7
	F_errmsg(m, int32(660397), v16+int32(128))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L13
	} else {
		goto L126
	}
L126:
	;
	v347 = int32(368)
	goto L85
L127:
	;
	if l3 == v206 {
		goto L84
	} else {
		goto L128
	}
L128:
	;
	v315 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L13
	} else {
		goto L129
	}
L129:
	;
	if l6 != int32(6) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	if v315 == int32(0) {
		goto L84
	} else {
		goto L136
	}
L131:
	;
	if l9 == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	if v315 == int32(0) {
		goto L84
	} else {
		goto L133
	}
L133:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = l9
	F_errmsg(m, int32(668326), v16+int32(144))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L13
	} else {
		goto L135
	}
L135:
	;
	v347 = int32(376)
	goto L85
L136:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L13
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = l7
	F_errmsg(m, int32(660437), v16+int32(160))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	v347 = int32(381)
	goto L85
L139:
	;
	goto L84
L140:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = l4
	F_errmsg(m, int32(66420), v16+int32(176))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(475972), int32(3495), int32(300784))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l9
	F_errmsg(m, int32(668395), v16+int32(32))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(475972), int32(2956), int32(288210))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_restriction_is_always_true(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v6 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v79
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v7 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v9 == int32(52) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v12 != int32(1) {
		v79 = v3
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L21
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v15 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v17 != int32(6) {
		v79 = v3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v20 != 0 {
		v79 = v3
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+8)))
	if v21 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(1)
L12:
	;
	goto L13
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v27 = F_find_base_rel(m, l0, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+8)))
	if int32(0) < v32 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v36 = F_bms_is_member(m, v32, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	return int32(0)
L19:
	;
	if v36 != 0 {
		v79 = int32(1)
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if base.B2i32(v40 != int32(0)) == int32(0) {
		v79 = v3
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v46 == int32(0) {
		v79 = v3
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		v79 = v3
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v54 = int32(0)
	goto L25
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v54<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != int32(318) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v79 = v3
	goto L1
L27:
	;
	v73 = v54 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v73 < v74 {
		v54 = v73
		goto L25
	} else {
		goto L31
	}
L28:
	;
	v66 = F_restriction_is_always_true(m, l0, v62)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	if v66 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	return int32(1)
L31:
	;
	goto L26
}
func F_rmdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_rmdir(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v2) {
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0) - v2
		v10 = int32(-1)
	} else {
		v10 = v2
	}
	return v10
}
func F_romanian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v650 int32
	_ = v650
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v769 int32
	_ = v769
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v888 int32
	_ = v888
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v932 int32
	_ = v932
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1009 int32
	_ = v1009
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1127 int32
	_ = v1127
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1161 int32
	_ = v1161
	var v1172 int32
	_ = v1172
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1250 int32
	_ = v1250
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1368 int32
	_ = v1368
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1436 int32
	_ = v1436
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1487 int32
	_ = v1487
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1520 int32
	_ = v1520
	var v1531 int32
	_ = v1531
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1608 int32
	_ = v1608
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1646 int32
	_ = v1646
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1743 int32
	_ = v1743
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1794 int32
	_ = v1794
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1820 int32
	_ = v1820
	var v1827 int32
	_ = v1827
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1916 int32
	_ = v1916
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1950 int32
	_ = v1950
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1990 int32
	_ = v1990
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2041 int32
	_ = v2041
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2067 int32
	_ = v2067
	var v2074 int32
	_ = v2074
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2112 int32
	_ = v2112
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2163 int32
	_ = v2163
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2189 int32
	_ = v2189
	var v2197 int32
	_ = v2197
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2520 int32
	_ = v2520
	var v2522 int32
	_ = v2522
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2539 int32
	_ = v2539
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2561 int32
	_ = v2561
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2574 int32
	_ = v2574
	var v2580 int32
	_ = v2580
	var v2597 int32
	_ = v2597
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2607 int32
	_ = v2607
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2615 int32
	_ = v2615
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2638 int32
	_ = v2638
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2658 int32
	_ = v2658
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2706 int32
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2713 int32
	_ = v2713
	var v2715 int32
	_ = v2715
	var v2722 int32
	_ = v2722
	var v2725 int32
	_ = v2725
	var v2729 int32
	_ = v2729
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2751 int32
	_ = v2751
	var v2757 int32
	_ = v2757
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2771 int32
	_ = v2771
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v7
	goto L2
L1:
	;
	return v2771
L2:
	;
	v15 = v10 + int32(1)
	goto L4
L3:
	;
	v117 = v7
	goto L43
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 <= v15 {
		v38 = v23
		v39 = v24
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L4
L8:
	;
	v108 = F_slice_from_s(m, l0, int32(2), int32(2157364))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L41
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
	v10 = v92
	goto L2
L10:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v95
	switch v32 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L38
	default:
		goto L7
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L19
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))))
	switch v27 - int32(159) {
	case 0, 4:
		goto L13
	default:
		v38 = v23
		v39 = v24
		goto L11
	}
L13:
	;
	v32 = F_find_among(m, l0, int32(4261680), int32(2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v32 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = v37
	v39 = v36
	goto L11
L17:
	;
	if int32(0) <= v92 {
		goto L9
	} else {
		goto L37
	}
L19:
	;
	goto L20
L20:
	;
	goto L21
L21:
	;
	v47 = v10
	v49 = int32(1)
	goto L24
L23:
	;
	v92 = v77
	goto L17
L24:
	;
	if v39 <= v47 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v92 = int32(-1)
	goto L17
L27:
	;
	goto L28
L28:
	;
	v54 = v47 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v47))))
	if base.Ui32(v56) < base.Ui32(int32(192)) {
		v77 = v54
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v78 = int32(1)
	if v78 < v49 {
		v47 = v77
		v49 = v49 - v78
		goto L24
	} else {
		goto L36
	}
L30:
	;
	if v39 <= v54 {
		v77 = v54
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v63 = v54
	goto L32
L32:
	;
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38+v63))))
	if int32(-65) < v66 {
		v77 = v63
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v77 = v39
	goto L29
L34:
	;
	v70 = v63 + int32(1)
	if v70 != v39 {
		v63 = v70
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L25
L37:
	;
	goto L6
L38:
	;
	v101 = F_slice_from_s(m, l0, int32(2), int32(2157366))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	if int32(0) <= v101 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v2771 = v101
	goto L1
L41:
	;
	if v108 < int32(0) {
		v2771 = v108
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L7
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L47
L44:
	;
	v2771 = v2766
	goto L1
L45:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v238 != 0 {
		v512 = v239
		goto L70
	} else {
		goto L71
	}
L46:
	;
	v238 = v231
	goto L45
L47:
	;
	if v133 <= v117 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v231 = int32(0)
	goto L46
L49:
	;
	v238 = int32(-1)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v149 = int32(1)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v134))))
	if base.Ui32(v151) < base.Ui32(int32(192)) {
		v208 = v151
		v209 = v149
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if int32(259) < v208 {
		v231 = v209
		goto L46
	} else {
		goto L65
	}
L53:
	;
	v155 = v117 + int32(1)
	if v155 == v133 {
		v208 = v151
		v209 = v149
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v134))))
	v160 = v158 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v151) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v134))))
	v176 = v174 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v151) {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v164 = v117 + int32(2)
	if v164 != v133 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v208 = v151<<(uint(int32(6))%32)&int32(1984) | v160
	v209 = int32(2)
	goto L52
L59:
	;
	goto L58
L60:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v180))))
	v208 = v193&int32(63) | (v151<<(uint(int32(18))%32)&int32(1835008) | v160<<(uint(int32(12))%32) | v176<<(uint(int32(6))%32))
	v209 = int32(4)
	goto L52
L61:
	;
	v180 = v117 + int32(3)
	if v180 != v133 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v208 = v151<<(uint(int32(12))%32)&int32(61440) | v160<<(uint(int32(6))%32) | v176
	v209 = int32(3)
	goto L52
L64:
	;
	goto L63
L65:
	;
	v213 = v208 - int32(97)
	if v213 < int32(0) {
		v231 = v209
		goto L46
	} else {
		goto L66
	}
L66:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v213)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v219)>>(uint(v213&int32(7))%32))&int32(1) == int32(0) {
		v231 = v209
		goto L46
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v209 + v117
	goto L68
L68:
	;
	goto L48
L69:
	;
	v2766 = F_slice_from_s(m, l0, int32(1), int32(2157398))
	mBase = m.M
	v2767 = m.ExcPending
	if v2767 != 0 {
		goto L14
	} else {
		goto L678
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L135
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v240
	if v239 == v240 {
		v379 = v239
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v240
	if v379 == v240 {
		goto L104
	} else {
		goto L105
	}
L73:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v240))))
	if v245 != int32(117) {
		v379 = v239
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v249 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L77
L75:
	;
	if v369 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L76:
	;
	v369 = v362
	goto L75
L77:
	;
	if v264 <= v249 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v362 = int32(0)
	goto L76
L79:
	;
	v369 = int32(-1)
	goto L75
L80:
	;
	goto L81
L81:
	;
	v280 = int32(1)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249+v265))))
	if base.Ui32(v282) < base.Ui32(int32(192)) {
		v339 = v282
		v340 = v280
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if int32(259) < v339 {
		v362 = v340
		goto L76
	} else {
		goto L95
	}
L83:
	;
	v286 = v240 + int32(2)
	if v286 == v264 {
		v339 = v282
		v340 = v280
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v265))))
	v291 = v289 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v282) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v265))))
	v307 = v305 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v282) {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	v295 = v240 + int32(3)
	if v295 != v264 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v339 = v282<<(uint(int32(6))%32)&int32(1984) | v291
	v340 = int32(2)
	goto L82
L89:
	;
	goto L88
L90:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v311))))
	v339 = v324&int32(63) | (v282<<(uint(int32(18))%32)&int32(1835008) | v291<<(uint(int32(12))%32) | v307<<(uint(int32(6))%32))
	v340 = int32(4)
	goto L82
L91:
	;
	v311 = v240 + int32(4)
	if v311 != v264 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v339 = v282<<(uint(int32(12))%32)&int32(61440) | v291<<(uint(int32(6))%32) | v307
	v340 = int32(3)
	goto L82
L94:
	;
	goto L93
L95:
	;
	v344 = v339 - int32(97)
	if v344 < int32(0) {
		v362 = v340
		goto L76
	} else {
		goto L96
	}
L96:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v344)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v350)>>(uint(v344&int32(7))%32))&int32(1) == int32(0) {
		v362 = v340
		goto L76
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v340 + v249
	goto L98
L98:
	;
	goto L78
L99:
	;
	v374 = F_slice_from_s(m, l0, int32(1), int32(2157397))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L14
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v379 = v378
	goto L72
L102:
	;
	if v374 < int32(0) {
		v2771 = v374
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L43
L104:
	;
	v512 = v240
	goto L70
L105:
	;
	goto L106
L106:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382+v240))))
	if v384 != int32(105) {
		v512 = v379
		goto L70
	} else {
		goto L107
	}
L107:
	;
	v388 = v240 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v388
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L110
L108:
	;
	if v508 == int32(0) {
		goto L69
	} else {
		goto L132
	}
L109:
	;
	v508 = v501
	goto L108
L110:
	;
	if v403 <= v388 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v501 = int32(0)
	goto L109
L112:
	;
	v508 = int32(-1)
	goto L108
L113:
	;
	goto L114
L114:
	;
	v419 = int32(1)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388+v404))))
	if base.Ui32(v421) < base.Ui32(int32(192)) {
		v478 = v421
		v479 = v419
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if int32(259) < v478 {
		v501 = v479
		goto L109
	} else {
		goto L128
	}
L116:
	;
	v425 = v240 + int32(2)
	if v425 == v403 {
		v478 = v421
		v479 = v419
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v404))))
	v430 = v428 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v421) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434+v404))))
	v446 = v444 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v421) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v434 = v240 + int32(3)
	if v434 != v403 {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v478 = v421<<(uint(int32(6))%32)&int32(1984) | v430
	v479 = int32(2)
	goto L115
L122:
	;
	goto L121
L123:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v450))))
	v478 = v463&int32(63) | (v421<<(uint(int32(18))%32)&int32(1835008) | v430<<(uint(int32(12))%32) | v446<<(uint(int32(6))%32))
	v479 = int32(4)
	goto L115
L124:
	;
	v450 = v240 + int32(4)
	if v450 != v403 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v478 = v421<<(uint(int32(12))%32)&int32(61440) | v430<<(uint(int32(6))%32) | v446
	v479 = int32(3)
	goto L115
L127:
	;
	goto L126
L128:
	;
	v483 = v478 - int32(97)
	if v483 < int32(0) {
		v501 = v479
		goto L109
	} else {
		goto L129
	}
L129:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v483)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v489)>>(uint(v483&int32(7))%32))&int32(1) == int32(0) {
		v501 = v479
		goto L109
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v479 + v388
	goto L131
L131:
	;
	goto L111
L132:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v512 = v511
	goto L70
L133:
	;
	if int32(0) <= v567 {
		v117 = v567
		goto L43
	} else {
		goto L153
	}
L135:
	;
	goto L136
L136:
	;
	goto L137
L137:
	;
	v522 = v117
	v524 = int32(1)
	goto L140
L139:
	;
	v567 = v552
	goto L133
L140:
	;
	if v512 <= v522 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v567 = int32(-1)
	goto L133
L143:
	;
	goto L144
L144:
	;
	v529 = v522 + int32(1)
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515+v522))))
	if base.Ui32(v531) < base.Ui32(int32(192)) {
		v552 = v529
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v553 = int32(1)
	if v553 < v524 {
		v522 = v552
		v524 = v524 - v553
		goto L140
	} else {
		goto L152
	}
L146:
	;
	if v512 <= v529 {
		v552 = v529
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v538 = v529
	goto L148
L148:
	;
	v541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v515+v538))))
	if int32(-65) < v541 {
		v552 = v538
		goto L145
	} else {
		goto L150
	}
L149:
	;
	v552 = v512
	goto L145
L150:
	;
	v545 = v538 + int32(1)
	if v545 != v512 {
		v538 = v545
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	goto L141
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v571)+8)) = v572
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L160
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v577
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1743 = v577
	goto L416
L155:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1717)+8)) = v1715
	goto L154
L156:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1715 = v1713 + v1711
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v577
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L289
L158:
	;
	if v695 != 0 {
		goto L157
	} else {
		goto L182
	}
L159:
	;
	v695 = v688
	goto L158
L160:
	;
	if v590 <= v577 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v688 = int32(0)
	goto L159
L162:
	;
	v695 = int32(-1)
	goto L158
L163:
	;
	goto L164
L164:
	;
	v606 = int32(1)
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v591))))
	if base.Ui32(v608) < base.Ui32(int32(192)) {
		v665 = v608
		v666 = v606
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if int32(259) < v665 {
		v688 = v666
		goto L159
	} else {
		goto L178
	}
L166:
	;
	v612 = v577 + int32(1)
	if v612 == v590 {
		v665 = v608
		v666 = v606
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612+v591))))
	v617 = v615 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v608) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v591))))
	v633 = v631 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v608) {
		goto L174
	} else {
		goto L175
	}
L169:
	;
	v621 = v577 + int32(2)
	if v621 != v590 {
		goto L168
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v665 = v608<<(uint(int32(6))%32)&int32(1984) | v617
	v666 = int32(2)
	goto L165
L172:
	;
	goto L171
L173:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591+v637))))
	v665 = v650&int32(63) | (v608<<(uint(int32(18))%32)&int32(1835008) | v617<<(uint(int32(12))%32) | v633<<(uint(int32(6))%32))
	v666 = int32(4)
	goto L165
L174:
	;
	v637 = v577 + int32(3)
	if v637 != v590 {
		goto L173
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v665 = v608<<(uint(int32(12))%32)&int32(61440) | v617<<(uint(int32(6))%32) | v633
	v666 = int32(3)
	goto L165
L177:
	;
	goto L176
L178:
	;
	v670 = v665 - int32(97)
	if v670 < int32(0) {
		v688 = v666
		goto L159
	} else {
		goto L179
	}
L179:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v670)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v676)>>(uint(v670&int32(7))%32))&int32(1) == int32(0) {
		v688 = v666
		goto L159
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v666 + v577
	goto L181
L181:
	;
	goto L161
L182:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L185
L183:
	;
	if v813 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L184:
	;
	v813 = v806
	goto L183
L185:
	;
	if v709 <= v696 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v806 = int32(0)
	goto L184
L187:
	;
	v813 = int32(-1)
	goto L183
L188:
	;
	goto L189
L189:
	;
	v725 = int32(1)
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v710))))
	if base.Ui32(v727) < base.Ui32(int32(192)) {
		v784 = v727
		v785 = v725
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if int32(259) < v784 {
		goto L203
	} else {
		goto L204
	}
L191:
	;
	v731 = v696 + int32(1)
	if v731 == v709 {
		v784 = v727
		v785 = v725
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731+v710))))
	v736 = v734 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v727) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740+v710))))
	v752 = v750 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v727) {
		goto L199
	} else {
		goto L200
	}
L194:
	;
	v740 = v696 + int32(2)
	if v740 != v709 {
		goto L193
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v784 = v727<<(uint(int32(6))%32)&int32(1984) | v736
	v785 = int32(2)
	goto L190
L197:
	;
	goto L196
L198:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710+v756))))
	v784 = v769&int32(63) | (v727<<(uint(int32(18))%32)&int32(1835008) | v736<<(uint(int32(12))%32) | v752<<(uint(int32(6))%32))
	v785 = int32(4)
	goto L190
L199:
	;
	v756 = v696 + int32(3)
	if v756 != v709 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v784 = v727<<(uint(int32(12))%32)&int32(61440) | v736<<(uint(int32(6))%32) | v752
	v785 = int32(3)
	goto L190
L202:
	;
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v785 + v696
	goto L207
L204:
	;
	v789 = v784 - int32(97)
	if v789 < int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v789)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v795)>>(uint(v789&int32(7))%32))&int32(1) != 0 {
		v806 = v785
		goto L184
	} else {
		goto L206
	}
L206:
	;
	goto L203
L207:
	;
	goto L186
L208:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v837 = v827
	goto L213
L209:
	;
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v696
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L239
L211:
	;
	if int32(0) <= v932 {
		v1711 = v932
		goto L156
	} else {
		goto L236
	}
L212:
	;
	v932 = v904
	goto L211
L213:
	;
	if v828 <= v837 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v932 = int32(-1)
	goto L211
L216:
	;
	goto L217
L217:
	;
	v844 = int32(1)
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837+v829))))
	if base.Ui32(v846) < base.Ui32(int32(192)) {
		v903 = v846
		v904 = v844
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if int32(259) < v903 {
		goto L231
	} else {
		goto L232
	}
L219:
	;
	v850 = v837 + int32(1)
	if v850 == v828 {
		v903 = v846
		v904 = v844
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850+v829))))
	v855 = v853 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v846) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859+v829))))
	v871 = v869 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v846) {
		goto L227
	} else {
		goto L228
	}
L222:
	;
	v859 = v837 + int32(2)
	if v859 != v828 {
		goto L221
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v903 = v846<<(uint(int32(6))%32)&int32(1984) | v855
	v904 = int32(2)
	goto L218
L225:
	;
	goto L224
L226:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v875))))
	v903 = v888&int32(63) | (v846<<(uint(int32(18))%32)&int32(1835008) | v855<<(uint(int32(12))%32) | v871<<(uint(int32(6))%32))
	v904 = int32(4)
	goto L218
L227:
	;
	v875 = v837 + int32(3)
	if v875 != v828 {
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v903 = v846<<(uint(int32(12))%32)&int32(61440) | v855<<(uint(int32(6))%32) | v871
	v904 = int32(3)
	goto L218
L230:
	;
	goto L229
L231:
	;
	v921 = v904 + v837
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v921
	v837 = v921
	goto L213
L232:
	;
	v908 = v903 - int32(97)
	if v908 < int32(0) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v908)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v914)>>(uint(v908&int32(7))%32))&int32(1) != 0 {
		goto L212
	} else {
		goto L234
	}
L234:
	;
	goto L231
L236:
	;
	goto L210
L237:
	;
	if v1054 != 0 {
		goto L157
	} else {
		goto L261
	}
L238:
	;
	v1054 = v1047
	goto L237
L239:
	;
	if v949 <= v696 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1047 = int32(0)
	goto L238
L241:
	;
	v1054 = int32(-1)
	goto L237
L242:
	;
	goto L243
L243:
	;
	v965 = int32(1)
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v950))))
	if base.Ui32(v967) < base.Ui32(int32(192)) {
		v1024 = v967
		v1025 = v965
		goto L244
	} else {
		goto L245
	}
L244:
	;
	if int32(259) < v1024 {
		v1047 = v1025
		goto L238
	} else {
		goto L257
	}
L245:
	;
	v971 = v696 + int32(1)
	if v971 == v949 {
		v1024 = v967
		v1025 = v965
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971+v950))))
	v976 = v974 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v967) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980+v950))))
	v992 = v990 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v967) {
		goto L253
	} else {
		goto L254
	}
L248:
	;
	v980 = v696 + int32(2)
	if v980 != v949 {
		goto L247
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1024 = v967<<(uint(int32(6))%32)&int32(1984) | v976
	v1025 = int32(2)
	goto L244
L251:
	;
	goto L250
L252:
	;
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950+v996))))
	v1024 = v1009&int32(63) | (v967<<(uint(int32(18))%32)&int32(1835008) | v976<<(uint(int32(12))%32) | v992<<(uint(int32(6))%32))
	v1025 = int32(4)
	goto L244
L253:
	;
	v996 = v696 + int32(3)
	if v996 != v949 {
		goto L252
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1024 = v967<<(uint(int32(12))%32)&int32(61440) | v976<<(uint(int32(6))%32) | v992
	v1025 = int32(3)
	goto L244
L256:
	;
	goto L255
L257:
	;
	v1029 = v1024 - int32(97)
	if v1029 < int32(0) {
		v1047 = v1025
		goto L238
	} else {
		goto L258
	}
L258:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1029)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1035)>>(uint(v1029&int32(7))%32))&int32(1) == int32(0) {
		v1047 = v1025
		goto L238
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1025 + v696
	goto L260
L260:
	;
	goto L240
L261:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1076 = v1066
	goto L264
L262:
	;
	if int32(0) <= v1172 {
		v1711 = v1172
		goto L156
	} else {
		goto L286
	}
L263:
	;
	v1172 = v1143
	goto L262
L264:
	;
	if v1067 <= v1076 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1172 = int32(-1)
	goto L262
L267:
	;
	goto L268
L268:
	;
	v1083 = int32(1)
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076+v1068))))
	if base.Ui32(v1085) < base.Ui32(int32(192)) {
		v1142 = v1085
		v1143 = v1083
		goto L269
	} else {
		goto L270
	}
L269:
	;
	if int32(259) < v1142 {
		goto L263
	} else {
		goto L282
	}
L270:
	;
	v1089 = v1076 + int32(1)
	if v1089 == v1067 {
		v1142 = v1085
		v1143 = v1083
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089+v1068))))
	v1094 = v1092 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1085) {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1098+v1068))))
	v1110 = v1108 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1085) {
		goto L278
	} else {
		goto L279
	}
L273:
	;
	v1098 = v1076 + int32(2)
	if v1098 != v1067 {
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1142 = v1085<<(uint(int32(6))%32)&int32(1984) | v1094
	v1143 = int32(2)
	goto L269
L276:
	;
	goto L275
L277:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068+v1114))))
	v1142 = v1127&int32(63) | (v1085<<(uint(int32(18))%32)&int32(1835008) | v1094<<(uint(int32(12))%32) | v1110<<(uint(int32(6))%32))
	v1143 = int32(4)
	goto L269
L278:
	;
	v1114 = v1076 + int32(3)
	if v1114 != v1067 {
		goto L277
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1142 = v1085<<(uint(int32(12))%32)&int32(61440) | v1094<<(uint(int32(6))%32) | v1110
	v1143 = int32(3)
	goto L269
L281:
	;
	goto L280
L282:
	;
	v1147 = v1142 - int32(97)
	if v1147 < int32(0) {
		goto L263
	} else {
		goto L283
	}
L283:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1147)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1153)>>(uint(v1147&int32(7))%32))&int32(1) == int32(0) {
		goto L263
	} else {
		goto L284
	}
L284:
	;
	v1161 = v1143 + v1076
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1161
	v1076 = v1161
	goto L264
L286:
	;
	goto L157
L287:
	;
	if v1294 != 0 {
		goto L154
	} else {
		goto L312
	}
L288:
	;
	v1294 = v1287
	goto L287
L289:
	;
	if v1190 <= v577 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1287 = int32(0)
	goto L288
L291:
	;
	v1294 = int32(-1)
	goto L287
L292:
	;
	goto L293
L293:
	;
	v1206 = int32(1)
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v1191))))
	if base.Ui32(v1208) < base.Ui32(int32(192)) {
		v1265 = v1208
		v1266 = v1206
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(259) < v1265 {
		goto L307
	} else {
		goto L308
	}
L295:
	;
	v1212 = v577 + int32(1)
	if v1212 == v1190 {
		v1265 = v1208
		v1266 = v1206
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1212+v1191))))
	v1217 = v1215 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1208) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221+v1191))))
	v1233 = v1231 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1208) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	v1221 = v577 + int32(2)
	if v1221 != v1190 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1265 = v1208<<(uint(int32(6))%32)&int32(1984) | v1217
	v1266 = int32(2)
	goto L294
L301:
	;
	goto L300
L302:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191+v1237))))
	v1265 = v1250&int32(63) | (v1208<<(uint(int32(18))%32)&int32(1835008) | v1217<<(uint(int32(12))%32) | v1233<<(uint(int32(6))%32))
	v1266 = int32(4)
	goto L294
L303:
	;
	v1237 = v577 + int32(3)
	if v1237 != v1190 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1265 = v1208<<(uint(int32(12))%32)&int32(61440) | v1217<<(uint(int32(6))%32) | v1233
	v1266 = int32(3)
	goto L294
L306:
	;
	goto L305
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1266 + v577
	goto L311
L308:
	;
	v1270 = v1265 - int32(97)
	if v1270 < int32(0) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1270)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1276)>>(uint(v1270&int32(7))%32))&int32(1) != 0 {
		v1287 = v1266
		goto L288
	} else {
		goto L310
	}
L310:
	;
	goto L307
L311:
	;
	goto L290
L312:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L315
L313:
	;
	if v1412 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L314:
	;
	v1412 = v1405
	goto L313
L315:
	;
	if v1308 <= v1295 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1405 = int32(0)
	goto L314
L317:
	;
	v1412 = int32(-1)
	goto L313
L318:
	;
	goto L319
L319:
	;
	v1324 = int32(1)
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1309))))
	if base.Ui32(v1326) < base.Ui32(int32(192)) {
		v1383 = v1326
		v1384 = v1324
		goto L320
	} else {
		goto L321
	}
L320:
	;
	if int32(259) < v1383 {
		goto L333
	} else {
		goto L334
	}
L321:
	;
	v1330 = v1295 + int32(1)
	if v1330 == v1308 {
		v1383 = v1326
		v1384 = v1324
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330+v1309))))
	v1335 = v1333 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1326) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1339+v1309))))
	v1351 = v1349 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1326) {
		goto L329
	} else {
		goto L330
	}
L324:
	;
	v1339 = v1295 + int32(2)
	if v1339 != v1308 {
		goto L323
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1383 = v1326<<(uint(int32(6))%32)&int32(1984) | v1335
	v1384 = int32(2)
	goto L320
L327:
	;
	goto L326
L328:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1309+v1355))))
	v1383 = v1368&int32(63) | (v1326<<(uint(int32(18))%32)&int32(1835008) | v1335<<(uint(int32(12))%32) | v1351<<(uint(int32(6))%32))
	v1384 = int32(4)
	goto L320
L329:
	;
	v1355 = v1295 + int32(3)
	if v1355 != v1308 {
		goto L328
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1383 = v1326<<(uint(int32(12))%32)&int32(61440) | v1335<<(uint(int32(6))%32) | v1351
	v1384 = int32(3)
	goto L320
L332:
	;
	goto L331
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1384 + v1295
	goto L337
L334:
	;
	v1388 = v1383 - int32(97)
	if v1388 < int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1388)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1394)>>(uint(v1388&int32(7))%32))&int32(1) != 0 {
		v1405 = v1384
		goto L314
	} else {
		goto L336
	}
L336:
	;
	goto L333
L337:
	;
	goto L316
L338:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1436 = v1426
	goto L343
L339:
	;
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1295
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L369
L341:
	;
	if int32(0) <= v1531 {
		v1711 = v1531
		goto L156
	} else {
		goto L366
	}
L342:
	;
	v1531 = v1503
	goto L341
L343:
	;
	if v1427 <= v1436 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1531 = int32(-1)
	goto L341
L346:
	;
	goto L347
L347:
	;
	v1443 = int32(1)
	v1445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436+v1428))))
	if base.Ui32(v1445) < base.Ui32(int32(192)) {
		v1502 = v1445
		v1503 = v1443
		goto L348
	} else {
		goto L349
	}
L348:
	;
	if int32(259) < v1502 {
		goto L361
	} else {
		goto L362
	}
L349:
	;
	v1449 = v1436 + int32(1)
	if v1449 == v1427 {
		v1502 = v1445
		v1503 = v1443
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1449+v1428))))
	v1454 = v1452 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1445) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458+v1428))))
	v1470 = v1468 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1445) {
		goto L357
	} else {
		goto L358
	}
L352:
	;
	v1458 = v1436 + int32(2)
	if v1458 != v1427 {
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1502 = v1445<<(uint(int32(6))%32)&int32(1984) | v1454
	v1503 = int32(2)
	goto L348
L355:
	;
	goto L354
L356:
	;
	v1487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1428+v1474))))
	v1502 = v1487&int32(63) | (v1445<<(uint(int32(18))%32)&int32(1835008) | v1454<<(uint(int32(12))%32) | v1470<<(uint(int32(6))%32))
	v1503 = int32(4)
	goto L348
L357:
	;
	v1474 = v1436 + int32(3)
	if v1474 != v1427 {
		goto L356
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1502 = v1445<<(uint(int32(12))%32)&int32(61440) | v1454<<(uint(int32(6))%32) | v1470
	v1503 = int32(3)
	goto L348
L360:
	;
	goto L359
L361:
	;
	v1520 = v1503 + v1436
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1520
	v1436 = v1520
	goto L343
L362:
	;
	v1507 = v1502 - int32(97)
	if v1507 < int32(0) {
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1507)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1513)>>(uint(v1507&int32(7))%32))&int32(1) != 0 {
		goto L342
	} else {
		goto L364
	}
L364:
	;
	goto L361
L366:
	;
	goto L340
L367:
	;
	if v1653 != 0 {
		goto L154
	} else {
		goto L391
	}
L368:
	;
	v1653 = v1646
	goto L367
L369:
	;
	if v1548 <= v1295 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	v1646 = int32(0)
	goto L368
L371:
	;
	v1653 = int32(-1)
	goto L367
L372:
	;
	goto L373
L373:
	;
	v1564 = int32(1)
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295+v1549))))
	if base.Ui32(v1566) < base.Ui32(int32(192)) {
		v1623 = v1566
		v1624 = v1564
		goto L374
	} else {
		goto L375
	}
L374:
	;
	if int32(259) < v1623 {
		v1646 = v1624
		goto L368
	} else {
		goto L387
	}
L375:
	;
	v1570 = v1295 + int32(1)
	if v1570 == v1548 {
		v1623 = v1566
		v1624 = v1564
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1570+v1549))))
	v1575 = v1573 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1566) {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1579+v1549))))
	v1591 = v1589 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1566) {
		goto L383
	} else {
		goto L384
	}
L378:
	;
	v1579 = v1295 + int32(2)
	if v1579 != v1548 {
		goto L377
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1623 = v1566<<(uint(int32(6))%32)&int32(1984) | v1575
	v1624 = int32(2)
	goto L374
L381:
	;
	goto L380
L382:
	;
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549+v1595))))
	v1623 = v1608&int32(63) | (v1566<<(uint(int32(18))%32)&int32(1835008) | v1575<<(uint(int32(12))%32) | v1591<<(uint(int32(6))%32))
	v1624 = int32(4)
	goto L374
L383:
	;
	v1595 = v1295 + int32(3)
	if v1595 != v1548 {
		goto L382
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1623 = v1566<<(uint(int32(12))%32)&int32(61440) | v1575<<(uint(int32(6))%32) | v1591
	v1624 = int32(3)
	goto L374
L386:
	;
	goto L385
L387:
	;
	v1628 = v1623 - int32(97)
	if v1628 < int32(0) {
		v1646 = v1624
		goto L368
	} else {
		goto L388
	}
L388:
	;
	v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1628)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1634)>>(uint(v1628&int32(7))%32))&int32(1) == int32(0) {
		v1646 = v1624
		goto L368
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1624 + v1295
	goto L390
L390:
	;
	goto L370
L391:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L394
L392:
	;
	if int32(0) <= v1708 {
		v1715 = v1708
		goto L155
	} else {
		goto L412
	}
L394:
	;
	goto L395
L395:
	;
	goto L396
L396:
	;
	v1663 = v1655
	v1665 = int32(1)
	goto L399
L398:
	;
	v1708 = v1693
	goto L392
L399:
	;
	if v1656 <= v1663 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	goto L398
L401:
	;
	v1708 = int32(-1)
	goto L392
L402:
	;
	goto L403
L403:
	;
	v1670 = v1663 + int32(1)
	v1672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1654+v1663))))
	if base.Ui32(v1672) < base.Ui32(int32(192)) {
		v1693 = v1670
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1694 = int32(1)
	if v1694 < v1665 {
		v1663 = v1693
		v1665 = v1665 - v1694
		goto L399
	} else {
		goto L411
	}
L405:
	;
	if v1656 <= v1670 {
		v1693 = v1670
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1679 = v1670
	goto L407
L407:
	;
	v1682 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1654+v1679))))
	if int32(-65) < v1682 {
		v1693 = v1679
		goto L404
	} else {
		goto L409
	}
L408:
	;
	v1693 = v1656
	goto L404
L409:
	;
	v1686 = v1679 + int32(1)
	if v1686 != v1656 {
		v1679 = v1686
		goto L407
	} else {
		goto L410
	}
L410:
	;
	goto L408
L411:
	;
	goto L400
L412:
	;
	goto L154
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v577
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2217
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2217
	v2221 = v2217 - int32(1)
	if v2221 <= v577 {
		goto L516
	} else {
		goto L517
	}
L414:
	;
	if v1838 < int32(0) {
		goto L413
	} else {
		goto L439
	}
L415:
	;
	v1838 = v1810
	goto L414
L416:
	;
	if v1734 <= v1743 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1838 = int32(-1)
	goto L414
L419:
	;
	goto L420
L420:
	;
	v1750 = int32(1)
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743+v1735))))
	if base.Ui32(v1752) < base.Ui32(int32(192)) {
		v1809 = v1752
		v1810 = v1750
		goto L421
	} else {
		goto L422
	}
L421:
	;
	if int32(259) < v1809 {
		goto L434
	} else {
		goto L435
	}
L422:
	;
	v1756 = v1743 + int32(1)
	if v1756 == v1734 {
		v1809 = v1752
		v1810 = v1750
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v1759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1756+v1735))))
	v1761 = v1759 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1752) {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1765+v1735))))
	v1777 = v1775 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1752) {
		goto L430
	} else {
		goto L431
	}
L425:
	;
	v1765 = v1743 + int32(2)
	if v1765 != v1734 {
		goto L424
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	v1809 = v1752<<(uint(int32(6))%32)&int32(1984) | v1761
	v1810 = int32(2)
	goto L421
L428:
	;
	goto L427
L429:
	;
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1735+v1781))))
	v1809 = v1794&int32(63) | (v1752<<(uint(int32(18))%32)&int32(1835008) | v1761<<(uint(int32(12))%32) | v1777<<(uint(int32(6))%32))
	v1810 = int32(4)
	goto L421
L430:
	;
	v1781 = v1743 + int32(3)
	if v1781 != v1734 {
		goto L429
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	v1809 = v1752<<(uint(int32(12))%32)&int32(61440) | v1761<<(uint(int32(6))%32) | v1777
	v1810 = int32(3)
	goto L421
L433:
	;
	goto L432
L434:
	;
	v1827 = v1810 + v1743
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1827
	v1743 = v1827
	goto L416
L435:
	;
	v1814 = v1809 - int32(97)
	if v1814 < int32(0) {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1814)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1820)>>(uint(v1814&int32(7))%32))&int32(1) != 0 {
		goto L415
	} else {
		goto L437
	}
L437:
	;
	goto L434
L439:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1842 = v1841 + v1838
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1842
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1865 = v1842
	goto L442
L440:
	;
	if v1961 < int32(0) {
		goto L413
	} else {
		goto L464
	}
L441:
	;
	v1961 = v1932
	goto L440
L442:
	;
	if v1856 <= v1865 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1961 = int32(-1)
	goto L440
L445:
	;
	goto L446
L446:
	;
	v1872 = int32(1)
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865+v1857))))
	if base.Ui32(v1874) < base.Ui32(int32(192)) {
		v1931 = v1874
		v1932 = v1872
		goto L447
	} else {
		goto L448
	}
L447:
	;
	if int32(259) < v1931 {
		goto L441
	} else {
		goto L460
	}
L448:
	;
	v1878 = v1865 + int32(1)
	if v1878 == v1856 {
		v1931 = v1874
		v1932 = v1872
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1878+v1857))))
	v1883 = v1881 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1874) {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	v1897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1887+v1857))))
	v1899 = v1897 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1874) {
		goto L456
	} else {
		goto L457
	}
L451:
	;
	v1887 = v1865 + int32(2)
	if v1887 != v1856 {
		goto L450
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1931 = v1874<<(uint(int32(6))%32)&int32(1984) | v1883
	v1932 = int32(2)
	goto L447
L454:
	;
	goto L453
L455:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1857+v1903))))
	v1931 = v1916&int32(63) | (v1874<<(uint(int32(18))%32)&int32(1835008) | v1883<<(uint(int32(12))%32) | v1899<<(uint(int32(6))%32))
	v1932 = int32(4)
	goto L447
L456:
	;
	v1903 = v1865 + int32(3)
	if v1903 != v1856 {
		goto L455
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v1931 = v1874<<(uint(int32(12))%32)&int32(61440) | v1883<<(uint(int32(6))%32) | v1899
	v1932 = int32(3)
	goto L447
L459:
	;
	goto L458
L460:
	;
	v1936 = v1931 - int32(97)
	if v1936 < int32(0) {
		goto L441
	} else {
		goto L461
	}
L461:
	;
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1936)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v1942)>>(uint(v1936&int32(7))%32))&int32(1) == int32(0) {
		goto L441
	} else {
		goto L462
	}
L462:
	;
	v1950 = v1932 + v1865
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1950
	v1865 = v1950
	goto L442
L464:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1965 = v1964 + v1961
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1965
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1967)+4)) = v1965
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1990 = v1980
	goto L467
L465:
	;
	if v2085 < int32(0) {
		goto L413
	} else {
		goto L490
	}
L466:
	;
	v2085 = v2057
	goto L465
L467:
	;
	if v1981 <= v1990 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2085 = int32(-1)
	goto L465
L470:
	;
	goto L471
L471:
	;
	v1997 = int32(1)
	v1999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1990+v1982))))
	if base.Ui32(v1999) < base.Ui32(int32(192)) {
		v2056 = v1999
		v2057 = v1997
		goto L472
	} else {
		goto L473
	}
L472:
	;
	if int32(259) < v2056 {
		goto L485
	} else {
		goto L486
	}
L473:
	;
	v2003 = v1990 + int32(1)
	if v2003 == v1981 {
		v2056 = v1999
		v2057 = v1997
		goto L472
	} else {
		goto L474
	}
L474:
	;
	v2006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2003+v1982))))
	v2008 = v2006 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1999) {
		goto L476
	} else {
		goto L477
	}
L475:
	;
	v2022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2012+v1982))))
	v2024 = v2022 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1999) {
		goto L481
	} else {
		goto L482
	}
L476:
	;
	v2012 = v1990 + int32(2)
	if v2012 != v1981 {
		goto L475
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v2056 = v1999<<(uint(int32(6))%32)&int32(1984) | v2008
	v2057 = int32(2)
	goto L472
L479:
	;
	goto L478
L480:
	;
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1982+v2028))))
	v2056 = v2041&int32(63) | (v1999<<(uint(int32(18))%32)&int32(1835008) | v2008<<(uint(int32(12))%32) | v2024<<(uint(int32(6))%32))
	v2057 = int32(4)
	goto L472
L481:
	;
	v2028 = v1990 + int32(3)
	if v2028 != v1981 {
		goto L480
	} else {
		goto L484
	}
L482:
	;
	goto L483
L483:
	;
	v2056 = v1999<<(uint(int32(12))%32)&int32(61440) | v2008<<(uint(int32(6))%32) | v2024
	v2057 = int32(3)
	goto L472
L484:
	;
	goto L483
L485:
	;
	v2074 = v2057 + v1990
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2074
	v1990 = v2074
	goto L467
L486:
	;
	v2061 = v2056 - int32(97)
	if v2061 < int32(0) {
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v2067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2061)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v2067)>>(uint(v2061&int32(7))%32))&int32(1) != 0 {
		goto L466
	} else {
		goto L488
	}
L488:
	;
	goto L485
L490:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2089 = v2088 + v2085
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2089
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2112 = v2089
	goto L493
L491:
	;
	if v2208 < int32(0) {
		goto L413
	} else {
		goto L515
	}
L492:
	;
	v2208 = v2179
	goto L491
L493:
	;
	if v2103 <= v2112 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v2208 = int32(-1)
	goto L491
L496:
	;
	goto L497
L497:
	;
	v2119 = int32(1)
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2112+v2104))))
	if base.Ui32(v2121) < base.Ui32(int32(192)) {
		v2178 = v2121
		v2179 = v2119
		goto L498
	} else {
		goto L499
	}
L498:
	;
	if int32(259) < v2178 {
		goto L492
	} else {
		goto L511
	}
L499:
	;
	v2125 = v2112 + int32(1)
	if v2125 == v2103 {
		v2178 = v2121
		v2179 = v2119
		goto L498
	} else {
		goto L500
	}
L500:
	;
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125+v2104))))
	v2130 = v2128 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2121) {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	v2144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2134+v2104))))
	v2146 = v2144 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2121) {
		goto L507
	} else {
		goto L508
	}
L502:
	;
	v2134 = v2112 + int32(2)
	if v2134 != v2103 {
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2178 = v2121<<(uint(int32(6))%32)&int32(1984) | v2130
	v2179 = int32(2)
	goto L498
L505:
	;
	goto L504
L506:
	;
	v2163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104+v2150))))
	v2178 = v2163&int32(63) | (v2121<<(uint(int32(18))%32)&int32(1835008) | v2130<<(uint(int32(12))%32) | v2146<<(uint(int32(6))%32))
	v2179 = int32(4)
	goto L498
L507:
	;
	v2150 = v2112 + int32(3)
	if v2150 != v2103 {
		goto L506
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v2178 = v2121<<(uint(int32(12))%32)&int32(61440) | v2130<<(uint(int32(6))%32) | v2146
	v2179 = int32(3)
	goto L498
L510:
	;
	goto L509
L511:
	;
	v2183 = v2178 - int32(97)
	if v2183 < int32(0) {
		goto L492
	} else {
		goto L512
	}
L512:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2183)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v2189)>>(uint(v2183&int32(7))%32))&int32(1) == int32(0) {
		goto L492
	} else {
		goto L513
	}
L513:
	;
	v2197 = v2179 + v2112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2197
	v2112 = v2197
	goto L493
L515:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2211))) = v2212 + v2208
	goto L413
L516:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2313
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2315)+12)) = int32(0)
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2318
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2321 = v2318 - v2320
	v2324 = F_find_among_b(m, l0, int32(4263296), int32(46))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L14
	} else {
		goto L550
	}
L517:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2223+v2221))))
	if v2225&int32(224) != int32(96) {
		goto L516
	} else {
		goto L518
	}
L518:
	;
	if int32(1)<<(uint(v2225)%32)&int32(266786) == int32(0) {
		goto L516
	} else {
		goto L519
	}
L519:
	;
	v2238 = F_find_among_b(m, l0, int32(4261728), int32(16))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L14
	} else {
		goto L520
	}
L520:
	;
	if v2238 == int32(0) {
		goto L516
	} else {
		goto L521
	}
L521:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2242
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+4))
	if v2242 < v2245 {
		goto L516
	} else {
		goto L522
	}
L522:
	;
	switch v2238 - int32(1) {
	case 0:
		goto L529
	case 1:
		goto L528
	case 2:
		goto L527
	case 3:
		goto L526
	case 4:
		goto L525
	case 5:
		goto L524
	case 6:
		goto L523
	default:
		goto L516
	}
L523:
	;
	v2306 = F_slice_from_s(m, l0, int32(4), int32(2157407))
	mBase = m.M
	v2307 = m.ExcPending
	if v2307 != 0 {
		goto L14
	} else {
		goto L547
	}
L524:
	;
	v2300 = F_slice_from_s(m, l0, int32(2), int32(2157405))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L14
	} else {
		goto L545
	}
L525:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2272 = int32(2)
	v2274 = int32(0)
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2276-v2277 < v2272 {
		v2287 = v2274
		goto L539
	} else {
		goto L540
	}
L526:
	;
	v2267 = F_slice_from_s(m, l0, int32(1), int32(2157401))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L14
	} else {
		goto L536
	}
L527:
	;
	v2261 = F_slice_from_s(m, l0, int32(1), int32(2157400))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L14
	} else {
		goto L534
	}
L528:
	;
	v2255 = F_slice_from_s(m, l0, int32(1), int32(2157399))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L14
	} else {
		goto L532
	}
L529:
	;
	v2249 = F_slice_del(m, l0)
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L14
	} else {
		goto L530
	}
L530:
	;
	if int32(0) <= v2249 {
		goto L516
	} else {
		goto L531
	}
L531:
	;
	v2771 = v2249
	goto L1
L532:
	;
	if int32(0) <= v2255 {
		goto L516
	} else {
		goto L533
	}
L533:
	;
	v2771 = v2255
	goto L1
L534:
	;
	if int32(0) <= v2261 {
		goto L516
	} else {
		goto L535
	}
L535:
	;
	v2771 = v2261
	goto L1
L536:
	;
	if int32(0) <= v2267 {
		goto L516
	} else {
		goto L537
	}
L537:
	;
	v2771 = v2267
	goto L1
L538:
	;
	if v2287 != 0 {
		goto L516
	} else {
		goto L542
	}
L539:
	;
	goto L538
L540:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2283 = F_memcmp(m, v2280+v2276-v2272, int32(2157402), v2272)
	mBase = m.M
	if v2283 != 0 {
		v2287 = v2274
		goto L539
	} else {
		goto L541
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2276 - v2272
	v2287 = int32(1)
	goto L539
L542:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2288 + (v2242 - v2271)
	v2294 = F_slice_from_s(m, l0, int32(1), int32(2157404))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L14
	} else {
		goto L543
	}
L543:
	;
	if int32(0) <= v2294 {
		goto L516
	} else {
		goto L544
	}
L544:
	;
	v2771 = v2294
	goto L1
L545:
	;
	if int32(0) <= v2300 {
		goto L516
	} else {
		goto L546
	}
L546:
	;
	v2771 = v2300
	goto L1
L547:
	;
	if v2306 < int32(0) {
		v2771 = v2306
		goto L1
	} else {
		goto L548
	}
L548:
	;
	goto L516
L549:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2396 = v2395 + v2321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2396
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2396
	v2401 = F_find_among_b(m, l0, int32(4262048), int32(62))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L14
	} else {
		goto L576
	}
L550:
	;
	if v2324 == int32(0) {
		goto L549
	} else {
		goto L551
	}
L551:
	;
	v2330 = v2324
	goto L552
L552:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2334
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+4))
	if v2334 < v2337 {
		goto L549
	} else {
		goto L554
	}
L553:
	;
	goto L549
L554:
	;
	switch v2330 - int32(1) {
	case 0:
		goto L561
	case 1:
		goto L560
	case 2:
		goto L559
	case 3:
		goto L558
	case 4:
		goto L557
	case 5:
		goto L556
	default:
		goto L555
	}
L555:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2378)+12)) = int32(1)
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2382 = v2381 + v2321
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2382
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2382
	v2387 = F_find_among_b(m, l0, int32(4263296), int32(46))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L14
	} else {
		goto L574
	}
L556:
	;
	v2373 = F_slice_from_s(m, l0, int32(2), int32(2157487))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L14
	} else {
		goto L572
	}
L557:
	;
	v2367 = F_slice_from_s(m, l0, int32(2), int32(2157485))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L14
	} else {
		goto L570
	}
L558:
	;
	v2361 = F_slice_from_s(m, l0, int32(2), int32(2157483))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L14
	} else {
		goto L568
	}
L559:
	;
	v2355 = F_slice_from_s(m, l0, int32(2), int32(2157481))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L14
	} else {
		goto L566
	}
L560:
	;
	v2349 = F_slice_from_s(m, l0, int32(4), int32(2157477))
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L14
	} else {
		goto L564
	}
L561:
	;
	v2343 = F_slice_from_s(m, l0, int32(4), int32(2157473))
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L14
	} else {
		goto L562
	}
L562:
	;
	if int32(0) <= v2343 {
		goto L555
	} else {
		goto L563
	}
L563:
	;
	v2771 = v2343
	goto L1
L564:
	;
	if int32(0) <= v2349 {
		goto L555
	} else {
		goto L565
	}
L565:
	;
	v2771 = v2349
	goto L1
L566:
	;
	if int32(0) <= v2355 {
		goto L555
	} else {
		goto L567
	}
L567:
	;
	v2771 = v2355
	goto L1
L568:
	;
	if int32(0) <= v2361 {
		goto L555
	} else {
		goto L569
	}
L569:
	;
	v2771 = v2361
	goto L1
L570:
	;
	if int32(0) <= v2367 {
		goto L555
	} else {
		goto L571
	}
L571:
	;
	v2771 = v2367
	goto L1
L572:
	;
	if v2373 < int32(0) {
		v2771 = v2373
		goto L1
	} else {
		goto L573
	}
L573:
	;
	goto L555
L574:
	;
	if v2387 != 0 {
		v2330 = v2387
		goto L552
	} else {
		goto L575
	}
L575:
	;
	goto L553
L576:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2401 == int32(0) {
		v2454 = v2403
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2456
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v2454)+12))
	if v2458 != 0 {
		goto L597
	} else {
		goto L598
	}
L578:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2406
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v2403)))
	if v2406 < v2408 {
		v2454 = v2403
		goto L577
	} else {
		goto L579
	}
L579:
	;
	switch v2401 - int32(1) {
	case 0:
		goto L583
	case 1:
		goto L582
	case 2:
		goto L581
	default:
		goto L580
	}
L580:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2450)+12)) = int32(1)
	v2454 = v2450
	goto L577
L581:
	;
	v2445 = F_slice_from_s(m, l0, int32(3), int32(2157470))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L14
	} else {
		goto L595
	}
L582:
	;
	v2416 = int32(2)
	v2418 = int32(0)
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2420-v2421 < v2416 {
		v2431 = v2418
		goto L587
	} else {
		goto L588
	}
L583:
	;
	v2412 = F_slice_del(m, l0)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L14
	} else {
		goto L584
	}
L584:
	;
	if int32(0) <= v2412 {
		goto L580
	} else {
		goto L585
	}
L585:
	;
	v2771 = v2412
	goto L1
L586:
	;
	if v2431 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L587:
	;
	goto L586
L588:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2427 = F_memcmp(m, v2424+v2420-v2416, int32(2157467), v2416)
	mBase = m.M
	if v2427 != 0 {
		v2431 = v2418
		goto L587
	} else {
		goto L589
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2420 - v2416
	v2431 = int32(1)
	goto L587
L590:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2454 = v2434
	goto L577
L591:
	;
	goto L592
L592:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2435
	v2439 = F_slice_from_s(m, l0, int32(1), int32(2157469))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L14
	} else {
		goto L593
	}
L593:
	;
	if int32(0) <= v2439 {
		goto L580
	} else {
		goto L594
	}
L594:
	;
	v2771 = v2439
	goto L1
L595:
	;
	if v2445 < int32(0) {
		v2771 = v2445
		goto L1
	} else {
		goto L596
	}
L596:
	;
	goto L580
L597:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2638
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2638
	v2643 = F_find_among_b(m, l0, int32(4266112), int32(5))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L14
	} else {
		goto L634
	}
L598:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2454)+8))
	if v2456 < v2459 {
		goto L597
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2456
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2459
	v2466 = F_find_among_b(m, l0, int32(4264224), int32(94))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L14
	} else {
		goto L601
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2462
	goto L597
L601:
	;
	if v2466 == int32(0) {
		goto L600
	} else {
		goto L602
	}
L602:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2470
	switch v2466 - int32(1) {
	case 0:
		goto L604
	case 1:
		goto L603
	default:
		goto L600
	}
L603:
	;
	v2626 = F_slice_del(m, l0)
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L14
	} else {
		goto L631
	}
L604:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L607
L605:
	;
	if v2604 != 0 {
		goto L624
	} else {
		goto L625
	}
L606:
	;
	v2604 = v2597
	goto L605
L607:
	;
	if v2491 <= v2492 {
		v2597 = int32(-1)
		goto L606
	} else {
		goto L609
	}
L608:
	;
	v2597 = int32(0)
	goto L606
L609:
	;
	v2509 = int32(1)
	v2510 = v2491 - v2509
	v2512 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2488+v2510))))
	v2514 = v2512 & int32(255)
	if v2510 == v2492 {
		v2569 = v2514
		v2570 = v2509
		goto L610
	} else {
		goto L611
	}
L610:
	;
	if int32(259) < v2569 {
		goto L619
	} else {
		goto L620
	}
L611:
	;
	if int32(0) <= v2512 {
		v2569 = v2514
		v2570 = v2509
		goto L610
	} else {
		goto L612
	}
L612:
	;
	v2520 = v2514 & int32(63)
	v2522 = v2491 - int32(2)
	v2524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2488+v2522))))
	v2526 = v2524 << (uint(int32(6)) % 32)
	if base.B2i32(v2522 != v2492)&base.B2i32(base.Ui32(v2524) < base.Ui32(int32(192))) == int32(0) {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	v2569 = v2526&int32(1984) | v2520
	v2570 = int32(2)
	goto L610
L614:
	;
	goto L615
L615:
	;
	v2539 = v2526&int32(4032) | v2520
	v2541 = v2491 - int32(3)
	v2543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2488+v2541))))
	if base.B2i32(v2541 != v2492)&base.B2i32(base.Ui32(v2543) < base.Ui32(int32(224))) == int32(0) {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v2569 = v2543<<(uint(int32(12))%32)&int32(61440) | v2539
	v2570 = int32(3)
	goto L610
L617:
	;
	goto L618
L618:
	;
	v2561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491+(v2488-int32(4))))))
	v2569 = v2543<<(uint(int32(12))%32)&int32(258048) | v2561&int32(7)<<(uint(int32(18))%32) | v2539
	v2570 = int32(4)
	goto L610
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2491 - v2570
	goto L623
L620:
	;
	v2574 = v2569 - int32(97)
	if v2574 < int32(0) {
		goto L619
	} else {
		goto L621
	}
L621:
	;
	v2580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2574)>>(uint(int32(3))%32)))+uint32(_consts[1069]))))
	if int32(base.Ui32(v2580)>>(uint(v2574&int32(7))%32))&int32(1) == int32(0) {
		goto L619
	} else {
		goto L622
	}
L622:
	;
	v2604 = v2570
	goto L605
L623:
	;
	goto L608
L624:
	;
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2607 = v2605 + (v2470 - v2474)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2607
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2607 <= v2609 {
		goto L600
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	v2622 = F_slice_del(m, l0)
	mBase = m.M
	v2623 = m.ExcPending
	if v2623 != 0 {
		goto L14
	} else {
		goto L629
	}
L627:
	;
	v2611 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2611+v2607-int32(1)))))
	if v2615 != int32(117) {
		goto L600
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2607 - int32(1)
	goto L626
L629:
	;
	if int32(0) <= v2622 {
		goto L600
	} else {
		goto L630
	}
L630:
	;
	v2771 = v2622
	goto L1
L631:
	;
	if v2626 < int32(0) {
		v2771 = v2626
		goto L1
	} else {
		goto L632
	}
L632:
	;
	goto L600
L633:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2658
	goto L640
L634:
	;
	if v2643 == int32(0) {
		goto L633
	} else {
		goto L635
	}
L635:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2647
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2649)+8))
	if v2647 < v2650 {
		goto L633
	} else {
		goto L636
	}
L636:
	;
	v2652 = F_slice_del(m, l0)
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L14
	} else {
		goto L637
	}
L637:
	;
	if v2652 < int32(0) {
		v2771 = v2652
		goto L1
	} else {
		goto L638
	}
L638:
	;
	goto L633
L639:
	;
	if v2757 < int32(0) {
		v2771 = v2757
		goto L1
	} else {
		goto L677
	}
L640:
	;
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2666
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2668 <= v2666 {
		goto L643
	} else {
		goto L644
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2666
	v2757 = int32(1)
	goto L639
L642:
	;
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L656
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2666
	v2697 = v2666
	v2698 = v2668
	goto L642
L644:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2670+v2666))))
	switch v2672 - int32(73) {
	case 0, 12:
		goto L645
	default:
		goto L643
	}
L645:
	;
	v2677 = F_find_among(m, l0, int32(4266224), int32(3))
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L14
	} else {
		goto L646
	}
L646:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2679
	switch v2677 - int32(1) {
	case 0:
		goto L648
	case 1:
		goto L647
	case 2:
		goto L649
	default:
		goto L640
	}
L647:
	;
	v2692 = F_slice_from_s(m, l0, int32(1), int32(2158464))
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L14
	} else {
		goto L652
	}
L648:
	;
	v2686 = F_slice_from_s(m, l0, int32(1), int32(2158463))
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L14
	} else {
		goto L650
	}
L649:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2697 = v2679
	v2698 = v2683
	goto L642
L650:
	;
	if int32(0) <= v2686 {
		goto L640
	} else {
		goto L651
	}
L651:
	;
	v2757 = v2686
	goto L639
L652:
	;
	if int32(0) <= v2692 {
		goto L640
	} else {
		goto L653
	}
L653:
	;
	v2757 = v2692
	goto L639
L654:
	;
	if int32(0) <= v2751 {
		goto L674
	} else {
		goto L675
	}
L656:
	;
	goto L657
L657:
	;
	goto L658
L658:
	;
	v2706 = v2697
	v2708 = int32(1)
	goto L661
L660:
	;
	v2751 = v2736
	goto L654
L661:
	;
	if v2698 <= v2706 {
		goto L663
	} else {
		goto L664
	}
L662:
	;
	goto L660
L663:
	;
	v2751 = int32(-1)
	goto L654
L664:
	;
	goto L665
L665:
	;
	v2713 = v2706 + int32(1)
	v2715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2699+v2706))))
	if base.Ui32(v2715) < base.Ui32(int32(192)) {
		v2736 = v2713
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v2737 = int32(1)
	if v2737 < v2708 {
		v2706 = v2736
		v2708 = v2708 - v2737
		goto L661
	} else {
		goto L673
	}
L667:
	;
	if v2698 <= v2713 {
		v2736 = v2713
		goto L666
	} else {
		goto L668
	}
L668:
	;
	v2722 = v2713
	goto L669
L669:
	;
	v2725 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2699+v2722))))
	if int32(-65) < v2725 {
		v2736 = v2722
		goto L666
	} else {
		goto L671
	}
L670:
	;
	v2736 = v2698
	goto L666
L671:
	;
	v2729 = v2722 + int32(1)
	if v2729 != v2698 {
		v2722 = v2729
		goto L669
	} else {
		goto L672
	}
L672:
	;
	goto L670
L673:
	;
	goto L662
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2751
	goto L640
L675:
	;
	goto L676
L676:
	;
	goto L641
L677:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2658
	v2771 = int32(1)
	goto L1
L678:
	;
	if int32(0) <= v2766 {
		goto L43
	} else {
		goto L679
	}
L679:
	;
	goto L44
}
func F_rstacktoodeep(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v7 = m.G0
	v10 = v6 - (v7 - int32(1))
	v12 = v10 >> (uint(int32(31)) % 32)
	return base.B2i32(v4 < v10^v12-v12) & base.B2i32(v6 != int32(0))
}
func F_rtrim(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = v7 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v20 = v18 & int32(1)
			if v20 != 0 {
				v21 = v12
			} else {
				v21 = v7 + int32(4)
			}
			if v18 == int32(1) {
				v24 = int32(4)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v26&int32(254) == int32(2) {
					v35 = v24
				} else {
					v35 = base.B2i32(v26 == int32(18)) << (uint(v24) % 32)
				}
				if v26 == int32(1) {
					v38 = v24
				} else {
					v38 = v35
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v20 != 0 {
					v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v14 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v14 + int32(4)
			}
			if v54 == int32(1) {
				v60 = int32(4)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v62&int32(254) == int32(2) {
					v71 = v60
				} else {
					v71 = base.B2i32(v62 == int32(18)) << (uint(v60) % 32)
				}
				if v62 == int32(1) {
					v74 = v60
				} else {
					v74 = v71
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v56 != 0 {
					v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_dotrim(m, v21, v49, v57, v85, int32(0), int32(1))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				return v88
			}
		}
	}
}
