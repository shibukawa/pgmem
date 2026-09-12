package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v9)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v4
	v19 = F_LockAcquire(m, v7, l2, v4, v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_ReceiveSharedInvalidMessages(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[683]))
	if v15 < v17 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	goto L3
L3:
	;
	goto L9
L4:
	;
	v30 = int32(4471584)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	*(*int32)(unsafe.Add(mBase, _consts[682])) = v32 + int32(1)
	v37 = v32 << (uint(int32(4)) % 32)
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[684])))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(8)))) = v40
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v37)+uint32(_consts[685])))
	v45 = int32(4471592)
	v47 = *(*int64)(unsafe.Add(mBase, _consts[686]))
	*(*int64)(unsafe.Add(mBase, _consts[686])) = v47 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v44
	F_LocalExecuteInvalidationMessage(m, v12)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[683]))
	if v55 < v57 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[683])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[682])) = v80
	v87 = *(*int32)(unsafe.Add(mBase, _consts[687]))
	v89 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	v92 = v87 + v89<<(uint(int32(4))%32)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[688]))))
	if v95 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[689]))
	if v291 != 0 {
		goto L52
	} else {
		goto L53
	}
L11:
	;
	goto L10
L12:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v103 = F_LWLockAcquire(m, v99+int32(640), int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v189 = v80
	goto L14
L14:
	;
	if v189 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L15:
	;
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[688]))) = uint8(v107)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(1)
	if v109 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v87+int32(12), int32(515824), int32(511), int32(177856))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[690]))))
	if v122 == v119 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L18
L20:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v183+int32(640))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L33
	}
L21:
	;
	v134 = int32(0)
	v136 = v125
	goto L26
L22:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[691])))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v126 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[690]))) = uint16(v126)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[691]))) = v121
	v174 = int32(-1)
	goto L20
L25:
	;
	if v121 <= v167 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	if v121 <= v136 {
		v165 = v134
		v167 = v136
		goto L25
	} else {
		goto L28
	}
L27:
	;
	v165 = int32(32)
	v167 = v158
	goto L25
L28:
	;
	v143 = int32(4)
	v144 = v134 << (uint(v143) % 32)
	v148 = base.I32_rem_s(v136, int32(4096))
	v151 = v87 + int32(16) + v148<<(uint(v143)%32)
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
	*(*int64)(unsafe.Add(mBase, uint32(v144)+uint32(_consts[685]))) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v151)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v144)+uint32(_consts[684]))) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[691])))
	v157 = int32(1)
	v158 = v156 + v157
	*(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[691]))) = v158
	v161 = v134 + v157
	if v161 != int32(32) {
		v134 = v161
		v136 = v158
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[692]))) = uint8(v169)
	v174 = v165
	goto L20
L31:
	;
	goto L32
L32:
	;
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[688]))) = uint8(v171)
	v174 = v165
	goto L20
L33:
	;
	v189 = v174
	goto L14
L34:
	;
	v201 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v220 = int32(4471584)
	*(*int32)(unsafe.Add(mBase, _consts[682])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[683])) = v189
	v226 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	if v226 < v189 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if v201 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errmsg_internal(m, int32(112814), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v212 = int32(4471592)
	v214 = *(*int64)(unsafe.Add(mBase, _consts[686]))
	*(*int64)(unsafe.Add(mBase, _consts[686])) = v214 + int64(1)
	F_InvalidateSystemCaches(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(521071), int32(103), int32(180318))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L11
L44:
	;
	goto L47
L45:
	;
	goto L46
L46:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[683]))
	if v278 == int32(32) {
		goto L9
	} else {
		goto L51
	}
L47:
	;
	v239 = int32(4471584)
	v241 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	*(*int32)(unsafe.Add(mBase, _consts[682])) = v241 + int32(1)
	v246 = v241 << (uint(int32(4)) % 32)
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v246)+uint32(_consts[684])))
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(8)))) = v249
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v246)+uint32(_consts[685])))
	v254 = int32(4471592)
	v256 = *(*int64)(unsafe.Add(mBase, _consts[686]))
	*(*int64)(unsafe.Add(mBase, _consts[686])) = v256 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v253
	F_LocalExecuteInvalidationMessage(m, v12)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v266 = *(*int32)(unsafe.Add(mBase, _consts[683]))
	if v264 < v266 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L11
L52:
	;
	v293 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[689])) = v293
	v297 = F_errstart(m, int32(11), v293)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	m.G0 = v12 + int32(16)
	return
L55:
	;
	if v297 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errmsg_internal(m, int32(364853), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v308 = int32(0)
	F_SICleanupQueue(m, v308, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	F_errfinish(m, int32(521071), int32(138), int32(180318))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	goto L54
}
func F_SharedFileSetInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(4294967296)
	F_FileSetInit(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 != 0 {
			F_on_dsm_detach(m, l1, int32(1096), l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_SharedFileSetOnDetach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l1+int32(44), int32(515747), int32(101), int32(341813))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v17 = v15 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v17
			if v17 == v13 {
				F_FileSetDeleteAll(m, l1)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		v17 = v15 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v17
		if v17 == v13 {
			F_FileSetDeleteAll(m, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_shared_buffer_readv_complete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v203 int64
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int64
	_ = v326
	var v330 int32
	_ = v330
	v5 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(32)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = base.I64_extend_i32_u(v30)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v34)
	v39 = v34 & int32(448)
	v41 = l1 + int32(104)
	goto L1
L1:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(23)))) = uint8(v44)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	goto L2
L2:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+23)))
	if v53 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v39 == int32(256) {
		goto L60
	} else {
		goto L61
	}
L4:
	;
	v56 = int32(0)
	v261 = v56
	v262 = v56
	v267 = v5
	v268 = v5
	v269 = v5
	v272 = v5
	v284 = v56
	goto L3
L5:
	;
	goto L6
L6:
	;
	v67 = int32(0)
	v71 = v67
	v72 = v67
	v76 = v5
	v77 = v5
	v78 = v5
	v79 = v5
	v82 = v5
	v83 = v5
	goto L7
L7:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(3))%32)+v71<<(uint(int32(3))%32))))
	v104 = v97 + v101<<(uint(int32(6))%32)
	if v101 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v251 = int32(255)
	v261 = base.B2i32(v246&v251 != int32(0))
	v262 = v245
	v267 = v227
	v268 = v238
	v269 = v242
	v272 = v229
	v284 = v244 & v251 << (uint(int32(18)) % 32)
	goto L3
L9:
	;
	v126 = v104 + int32(-64)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v104-int32(48))))
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)) = uint8(v128)
	if base.B2i32(v39 == int32(256))|base.B2i32(v30 <= v71) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110+(v101^int32(-1))<<(uint(int32(2))%32))))
	v124 = v116
	goto L9
L11:
	;
	goto L12
L12:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v124 = v118 + v101<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	if v76&int32(255) != 0 {
		goto L47
	} else {
		goto L48
	}
L14:
	;
	v131 = int32(0)
	F_TerminateBufferIO(m, v126, v131, int32(134217728), v131, int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v140 = F_PageIsVerified(m, v124, v127, l3&int32(4)|int32(2), v28+int32(22))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return
L18:
	;
	v224 = v128
	v225 = v131
	v227 = v77
	v229 = v82
	goto L13
L19:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	if v140 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	if v180 != 0 {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = int64(0)
	v179 = v154
	v180 = v155
	v181 = int32(16777216)
	v182 = int32(194)
	v183 = int32(0)
	v184 = v155 << (uint(int32(11)) % 32)
	goto L20
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = int64(0)
	v169 = int32(0)
	v179 = v169
	v180 = v169
	v181 = int32(134217728)
	v182 = int32(258)
	v183 = int32(1)
	v184 = int32(2048)
	goto L20
L23:
	;
	if v142&int32(1) != 0 {
		goto L21
	} else {
		goto L29
	}
L24:
	;
	if l3&int32(1) == int32(0) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v154 = v142
	v155 = int32(0)
	goto L23
L27:
	;
	v147 = int32(0)
	v151 = F__emscripten_memset_bulkmem(m, v124, base.I32_extend8_s(v147), int32(8192))
	mBase = m.M
	goto L28
L28:
	;
	v154 = v147
	v155 = int32(1)
	goto L23
L29:
	;
	if v155 != 0 {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v158 = int32(0)
	F_TerminateBufferIO(m, v126, v158, int32(16777216), v158, int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v224 = v128
	v225 = v154
	v227 = v77
	v229 = v82
	goto L13
L32:
	;
	v190 = int32(512)
	goto L34
L33:
	;
	v190 = int32(0)
	goto L34
L34:
	;
	if v179&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v195 = int32(1024)
	goto L37
L36:
	;
	v195 = int32(0)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v142<<(uint(int32(18))%32) | v182 | (v184 | (v190 | v195)) | v71<<(uint(int32(25))%32)
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v203
	F_pgaio_result_report(m, v28+int32(8), v41, int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	v210 = int32(0)
	F_TerminateBufferIO(m, v126, v210, v181, v210, int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	if v183 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v224 = v180
	v225 = v179
	v227 = v77
	v229 = v82
	goto L13
L41:
	;
	goto L42
L42:
	;
	if v180 != 0 {
		v224 = int32(1)
		v225 = v179
		v227 = v77
		v229 = v82
		goto L13
	} else {
		goto L43
	}
L43:
	;
	if v77&int32(255) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v220 = v82
	goto L46
L45:
	;
	v220 = v71
	goto L46
L46:
	;
	v224 = int32(0)
	v225 = v179
	v227 = v77 + int32(1)
	v229 = v220
	goto L13
L47:
	;
	v235 = v78
	goto L49
L48:
	;
	v235 = v71
	goto L49
L49:
	;
	v237 = v225 & int32(1)
	if v237 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v238 = v235
	goto L52
L51:
	;
	v238 = v78
	goto L52
L52:
	;
	if v72&int32(255) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v241 = v79
	goto L55
L54:
	;
	v241 = v71
	goto L55
L55:
	;
	if v224 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v242 = v241
	goto L58
L57:
	;
	v242 = v79
	goto L58
L58:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v244 = v243 + v83
	v245 = v224 + v72
	v246 = v237 + v76
	v248 = v71 + int32(1)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+23)))
	if base.Ui32(v248) < base.Ui32(v249) {
		v71 = v248
		v72 = v245
		v76 = v246
		v77 = v227
		v78 = v238
		v79 = v242
		v82 = v229
		v83 = v244
		goto L7
	} else {
		goto L59
	}
L59:
	;
	goto L8
L60:
	;
	m.G0 = v28 + int32(32)
	return
L61:
	;
	v287 = int32(255)
	v288 = v267 & v287
	v289 = int32(0)
	if (base.B2i32(v288 != v289)|v261)&int32(1)|v262&v287 == v289 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v288 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v301 = int32(258)
	goto L65
L64:
	;
	v301 = int32(194)
	goto L65
L65:
	;
	v304 = v262 & int32(255)
	v305 = int32(0)
	if v261 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v311 = int32(1024)
	goto L68
L67:
	;
	v311 = v305
	goto L68
L68:
	;
	if v288 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v313 = v267
	goto L71
L70:
	;
	v313 = v262
	goto L71
L71:
	;
	if v304 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v320 = v269
	goto L74
L73:
	;
	v320 = v268
	goto L74
L74:
	;
	if v288 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v321 = v272
	goto L77
L76:
	;
	v321 = v320
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v284 | v301 | (base.B2i32(v304 != v305)<<(uint(int32(9))%32) | v311 | v313&int32(255)<<(uint(int32(11))%32)) | v321<<(uint(int32(25))%32)
	v326 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v326
	F_pgaio_result_report(m, v28, v41, int32(14))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L17
	} else {
		goto L78
	}
L78:
	;
	goto L60
}
func F_shared_buffer_write_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	if l0 != 0 {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_GetRelationPath(m, v6+int32(8), v14, v15, v16, int32(-1), v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v11
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v6 + int32(8)
				F_errcontext_msg(m, int32(740106), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v6 + int32(80)
					return
				}
			}
		}
	} else {
		m.G0 = v6 + int32(80)
		return
	}
}
