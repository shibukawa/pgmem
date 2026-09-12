package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AfterTriggerBeginQuery(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(4367492)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[502]))
	*(*int32)(unsafe.Add(mBase, _consts[502])) = v3 + int32(1)
	return
}
func F_AfterTriggerEndXact(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	if v3 != 0 {
		F_MemoryContextDelete(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v7 = int64(0)
			*(*int64)(unsafe.Add(mBase, _consts[505])) = v7
			*(*int64)(unsafe.Add(mBase, _consts[503])) = v7
			*(*int64)(unsafe.Add(mBase, _consts[506])) = int64(0)
			*(*int64)(unsafe.Add(mBase, _consts[507])) = int64(-4294967296)
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[508])) = v19
			*(*int32)(unsafe.Add(mBase, _consts[509])) = v19
			return
		}
	} else {
		*(*int64)(unsafe.Add(mBase, _consts[506])) = int64(0)
		*(*int64)(unsafe.Add(mBase, _consts[507])) = int64(-4294967296)
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[508])) = v19
		*(*int32)(unsafe.Add(mBase, _consts[509])) = v19
		return
	}
}
func F_AfterTriggerFireDeferred(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, _consts[503]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	goto L7
L4:
	;
	return
L5:
	;
	F_PushActiveSnapshot(m, v5)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v12 = int32(0)
	v14 = F_afterTriggerMarkEvents(m, int32(4367472), v12, v12)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	if v4 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	if v14 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v16 = int32(4367464)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[504]))
	v19 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[504])) = v18 + v19
	v25 = F_afterTriggerInvokeEvents(m, int32(4367472), v18, int32(0), v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L8
L13:
	;
	if v25 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	return
L18:
	;
	goto L17
}
func F_afterTriggerAddEvent(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = v11 & int32(939524096)
	if v13 == int32(134217728) {
		v24 = int32(24)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	if v13 == int32(268435456) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = int32(12)
	goto L5
L4:
	;
	v20 = int32(4)
	goto L5
L5:
	;
	if v13 != int32(805306368) {
		v24 = v20
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = int32(16)
	goto L1
L7:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if base.Ui32(v86) <= base.Ui32(v82) {
		v174 = v82
		v176 = v86
		goto L30
	} else {
		goto L31
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.Ui32(v24+int32(28)) <= base.Ui32(v28-v29) {
		v82 = v28
		v83 = v25
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	v44 = F_AllocSetContextCreateInternal(m, v39, int32(118240), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v47 = v34
	goto L14
L14:
	;
	if v25 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[501])) = v44
	v47 = v44
	goto L14
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v50 = v49 - v25
	v51 = int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if base.Ui32(v49-v55) < base.Ui32(int32(2801)) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v66 = int32(1024)
	goto L19
L19:
	;
	v67 = F_MemoryContextAlloc(m, v47, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L26
	}
L20:
	;
	v59 = v50 << (uint(v51) % 32)
	goto L22
L21:
	;
	v59 = int32(base.Ui32(v50) >> (uint(v51) % 32))
	goto L22
L22:
	;
	if base.Ui32(int32(1048576)) <= base.Ui32(v59) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v62 = int32(1048576)
	goto L25
L24:
	;
	v62 = v59
	goto L25
L25:
	;
	v66 = v62
	goto L19
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(0)
	v71 = v66 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v67 + int32(16)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v77 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = v77
	goto L29
L28:
	;
	v78 = l0
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v67
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v67
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v82 = v81
	v83 = v67
	goto L7
L30:
	;
	if base.Ui32(v176) <= base.Ui32(v174) {
		goto L55
	} else {
		goto L56
	}
L31:
	;
	v91 = v82
	v93 = v86
	goto L32
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v97 != v98 {
		v166 = v93
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v174 = v169
	v176 = v166
	goto L30
L34:
	;
	v169 = v91 + int32(28)
	if base.Ui32(v169) < base.Ui32(v166) {
		v91 = v169
		v93 = v166
		goto L32
	} else {
		goto L54
	}
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v100 != v101 {
		v166 = v93
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v103 != 0 {
		v166 = v93
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v104 != v105 {
		v166 = v93
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v107 != v108 {
		v166 = v93
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v110 != v111 {
		v166 = v93
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v115 = int32(0)
	v122 = base.B2i32(v113|v114 == v115)
	if v113 == v115 {
		v161 = v122
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v161 != 0 {
		v174 = v91
		v176 = v165
		goto L30
	} else {
		goto L53
	}
L42:
	;
	goto L41
L43:
	;
	if v114 == int32(0) {
		v161 = v122
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v128 != v129 {
		v161 = int32(0)
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v131 = int32(1)
	if v128 <= v131 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v134 = v131
	goto L48
L47:
	;
	v134 = v128
	goto L48
L48:
	;
	v135 = int32(8)
	v140 = int32(0)
	goto L49
L49:
	;
	v148 = v140 << (uint(int32(2)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v113+v135+v148)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+(v114+v135))))
	v153 = base.B2i32(v150 == v152)
	if v152 != v150 {
		v161 = v153
		goto L42
	} else {
		goto L51
	}
L50:
	;
	v161 = v153
	goto L42
L51:
	;
	v156 = v140 + int32(1)
	if v156 != v134 {
		v140 = v156
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v166 = v165
	goto L34
L54:
	;
	goto L33
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v183 = v181 - int32(28)
	v184 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v183))) = v184
	v187 = l2 + int32(24)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+24)) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v183)+16)) = v190
	v192 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v183)+8)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v195 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v217 = v174
	goto L57
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v24 != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v196 = int32(4470400)
	v197 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v200
	v202 = F_bms_copy(m, v195)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L15
	} else {
		goto L61
	}
L59:
	;
	v206 = int32(0)
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181-int32(12)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v181-int32(4)))) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v183
	v217 = v183
	goto L57
L61:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v197
	v206 = v202
	goto L60
L62:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v217 - v223 | v225&int32(-134217728)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v231 = v230 + v24
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v231
	return
L63:
	;
	v222 = F__emscripten_memcpy_bulkmem(m, v221, l1, v24)
	mBase = m.M
	v223 = v222
	goto L65
L64:
	;
	v223 = v221
	goto L65
L65:
	;
	goto L62
}
