package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+48)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v5)+56)) = v6
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[1]))
	if v2 < v23 {
		v27 = v2
		for {
			v30 = v27 * int32(40)
			v31 = int32(_a_F_AsyncShmemInit_0)
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
			v34 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v30+v32)+64)) = v34
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
			v39 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v37+v30)+68)) = v39
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v42+v30)+72)) = v34
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
			v48 = v47 + v30
			*(*int32)(unsafe.Add(mBase, uint32(v48)+88)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v48)+80)) = int64(0)
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v54+v30)+96)) = uint8(v39)
			v59 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[0]))
			*(*uint8)(unsafe.Add(mBase, uint32(v59+v30)+97)) = uint8(v39)
			v64 = v27 + int32(1)
			v66 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemInit[1]))
			if v64 < v66 {
				v27 = v64
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v74 = F_SlruScanDirectory(m, int32(_a_F_AsyncShmemInit_1), int32(302), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		return
	} else {
		return
	}
}
func F_Async_Notify(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
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
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[1]))
	if v17 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L10
	} else {
		goto L96
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L10
	} else {
		goto L92
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L10
	} else {
		goto L88
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_Async_Notify[2])))
	if v21 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L85
	}
L8:
	;
	if l0 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v26 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v26 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_Async_Notify_0), v11+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_Async_Notify_1), int32(906), int32(_a_F_Async_Notify_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v41 = F_strlen(m, l0)
	mBase = m.M
	v42 = v41
	goto L17
L16:
	;
	v42 = v3
	goto L17
L17:
	;
	if l1 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v43 = F_strlen(m, l1)
	mBase = m.M
	v44 = v43
	goto L20
L19:
	;
	v44 = v3
	goto L20
L20:
	;
	if v42 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v42) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(int32(_a_F_Async_Notify_3)) <= base.Ui32(v44) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v51 = int32(_a_F_Async_Notify_4)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[3]))
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[3])) = v55
	v60 = F_palloc(m, v42+v44+int32(6))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v44)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v42)
	v65 = v60 + int32(4)
	if (l0^v65)&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v142 = v65 + v42 + int32(1)
	if l1 != 0 {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	goto L25
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v119)
	if v119&int32(255) == int32(0) {
		goto L26
	} else {
		goto L42
	}
L28:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v118 = l0
	v119 = v71
	v120 = v65
	goto L27
L29:
	;
	goto L30
L30:
	;
	if l0&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v75 = l0
	v77 = v65
	goto L34
L32:
	;
	v89 = l0
	v91 = v65
	goto L33
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v96 = int32(-2139062144)
	if (int32(16843008)-v93|v93)&v96 != v96 {
		v118 = v89
		v119 = v93
		v120 = v91
		goto L27
	} else {
		goto L38
	}
L34:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v78)
	if v78 == int32(0) {
		goto L26
	} else {
		goto L36
	}
L35:
	;
	v89 = v85
	v91 = v83
	goto L33
L36:
	;
	v82 = int32(1)
	v83 = v77 + v82
	v85 = v75 + v82
	if v85&int32(3) != 0 {
		v75 = v85
		v77 = v83
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v101 = v89
	v102 = v93
	v103 = v91
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v102
	v105 = int32(4)
	v106 = v103 + v105
	v108 = v101 + v105
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v113 = int32(-2139062144)
	if (int32(16843008)-v110|v110)&v113 == v113 {
		v101 = v108
		v102 = v110
		v103 = v106
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v118 = v108
	v119 = v110
	v120 = v106
	goto L27
L41:
	;
	goto L40
L42:
	;
	v127 = v118
	v129 = v120
	goto L43
L43:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)) = uint8(v130)
	v132 = int32(1)
	if v130 != 0 {
		v127 = v127 + v132
		v129 = v129 + v132
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L26
L45:
	;
	goto L44
L46:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[5]))
	if v220 != 0 {
		goto L73
	} else {
		goto L74
	}
L47:
	;
	if (l1^v142)&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	goto L49
L49:
	;
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v217)
	goto L46
L50:
	;
	goto L46
L51:
	;
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v196)
	if v196&int32(255) == int32(0) {
		goto L51
	} else {
		goto L67
	}
L53:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v195 = l1
	v196 = v148
	v197 = v142
	goto L52
L54:
	;
	goto L55
L55:
	;
	if l1&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = l1
	v154 = v142
	goto L59
L57:
	;
	v166 = l1
	v168 = v142
	goto L58
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v173 = int32(-2139062144)
	if (int32(16843008)-v170|v170)&v173 != v173 {
		v195 = v166
		v196 = v170
		v197 = v168
		goto L52
	} else {
		goto L63
	}
L59:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v155)
	if v155 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L60:
	;
	v166 = v162
	v168 = v160
	goto L58
L61:
	;
	v159 = int32(1)
	v160 = v154 + v159
	v162 = v152 + v159
	if v162&int32(3) != 0 {
		v152 = v162
		v154 = v160
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v178 = v166
	v179 = v170
	v180 = v168
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v179
	v182 = int32(4)
	v183 = v180 + v182
	v185 = v178 + v182
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v190 = int32(-2139062144)
	if (int32(16843008)-v187|v187)&v190 == v190 {
		v178 = v185
		v179 = v187
		v180 = v183
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v195 = v185
	v196 = v187
	v197 = v183
	goto L52
L66:
	;
	goto L65
L67:
	;
	v204 = v195
	v206 = v197
	goto L68
L68:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)) = uint8(v207)
	v209 = int32(1)
	if v207 != 0 {
		v204 = v204 + v209
		v206 = v206 + v209
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L51
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[3])) = v52
	m.G0 = v11 + int32(32)
	return
L72:
	;
	v246 = F_AsyncExistsPendingNotify(m, v60)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L79
	}
L73:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v15 <= v221 {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[6]))
	v226 = F_MemoryContextAlloc(m, v224, int32(24))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L10
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v60
	v234 = F_list_make1_impl(m, int32(1), v11+int32(12))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v226)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = v234
	v241 = int32(_a_F_Async_Notify_5)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+20)) = v242
	*(*int32)(unsafe.Add(mBase, _c_F_Async_Notify[5])) = v226
	goto L71
L79:
	;
	if v246 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_pfree(m, v60)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_AddEventToPendingNotifies(m, v60)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L84
	}
L83:
	;
	goto L71
L84:
	;
	goto L71
L85:
	;
	F_errmsg_internal(m, int32(_a_F_Async_Notify_6), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_Async_Notify_1), int32(903), int32(_a_F_Async_Notify_2))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_Async_Notify_7), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_Async_Notify_1), int32(915), int32(_a_F_Async_Notify_2))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_Async_Notify_8), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_Async_Notify_1), int32(921), int32(_a_F_Async_Notify_2))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_Async_Notify_9), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_Async_Notify_1), int32(926), int32(_a_F_Async_Notify_2))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_Async_UnlistenAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[0])))
	if v7 != int32(1) {
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[1]))
		if v28 == int32(0) {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[2])))
			if v32&int32(1) == int32(0) {
				m.G0 = v4 + int32(16)
				return
			} else {
				F_queue_listen(m, int32(2), int32(_a_F_Async_UnlistenAll_0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					m.G0 = v4 + int32(16)
					return
				}
			}
		} else {
			F_queue_listen(m, int32(2), int32(_a_F_Async_UnlistenAll_0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				m.G0 = v4 + int32(16)
				return
			}
		}
	} else {
		v12 = F_errstart(m, int32(14), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 == int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[1]))
				if v28 == int32(0) {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[2])))
					if v32&int32(1) == int32(0) {
						m.G0 = v4 + int32(16)
						return
					} else {
						F_queue_listen(m, int32(2), int32(_a_F_Async_UnlistenAll_0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							m.G0 = v4 + int32(16)
							return
						}
					}
				} else {
					F_queue_listen(m, int32(2), int32(_a_F_Async_UnlistenAll_0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v4 + int32(16)
						return
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[3]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v17
				F_errmsg_internal(m, int32(_a_F_Async_UnlistenAll_1), v4)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_Async_UnlistenAll_2), int32(1078), int32(_a_F_Async_UnlistenAll_3))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[1]))
						if v28 == int32(0) {
							v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_Async_UnlistenAll[2])))
							if v32&int32(1) == int32(0) {
								m.G0 = v4 + int32(16)
								return
							} else {
								F_queue_listen(m, int32(2), int32(_a_F_Async_UnlistenAll_0))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									m.G0 = v4 + int32(16)
									return
								}
							}
						} else {
							F_queue_listen(m, int32(2), int32(_a_F_Async_UnlistenAll_0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								m.G0 = v4 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_Async_UnlistenOnExit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[0])))
	if v14 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[1]))
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v19 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v19 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v24
	F_errmsg_internal(m, int32(_a_F_Async_UnlistenOnExit_0), v11)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	F_errfinish(m, int32(_a_F_Async_UnlistenOnExit_1), int32(1867), int32(_a_F_Async_UnlistenOnExit_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	F_hash_destroy(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[3]))
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[1])) = int32(0)
	goto L10
L12:
	;
	v44 = v11 + int32(4)
	v45 = int32(1)
	v46 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v44)+12)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)) = uint8(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = int32(-1)
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_asyncQueueUnregister(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L44
	}
L15:
	;
	v54 = F_dshash_seq_next(m, v44)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v58 = v54
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_dshash_seq_term(m, v11+int32(4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L43
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[4]))
	if v64 != v66 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v127 = F_dshash_seq_next(m, v11+int32(4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L41
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[5]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)+68))
	v71 = F_dsa_get_address(m, v69, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+72))
	if v73 <= int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[6]))
	v79 = int32(0)
	goto L26
L26:
	;
	v89 = v71 + v79<<(uint(int32(3))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v78 == v90 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L22
L28:
	;
	v93 = v73 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+72)) = v93
	if v79 < v93 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v115 = v79 + int32(1)
	if v115 != v73 {
		v79 = v115
		goto L26
	} else {
		goto L40
	}
L31:
	;
	v98 = (v93 - v79) << (uint(int32(3)) % 32)
	if v98 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v104 = v93
	goto L33
L33:
	;
	if v104 != 0 {
		goto L22
	} else {
		goto L37
	}
L34:
	;
	base.MemoryCopy(m, v89, v89+int32(8), v98)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v58)+72))
	v104 = v102
	goto L33
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_Async_UnlistenOnExit[5]))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v58)+68))
	F_dsa_free(m, v106, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_dshash_delete_current(m, v11+int32(4))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L22
L40:
	;
	goto L27
L41:
	;
	if v127 != 0 {
		v58 = v127
		goto L20
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	goto L14
L44:
	;
	m.G0 = v11 + int32(32)
	return
}
func F_ExecAsyncAppendResponse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v9 != 0 {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
			if v10&int32(2) == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+136))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v19 + int32(1)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+132))
				*(*int32)(unsafe.Add(mBase, uint32(v23+v19<<(uint(int32(2))%32)))) = v9
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+148))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = F_bms_add_member(m, v28, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v30
					return
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v15 - int32(1)
				return
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v15 - int32(1)
			return
		}
	} else {
		return
	}
}
