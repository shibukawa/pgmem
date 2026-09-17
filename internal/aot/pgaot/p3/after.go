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
	v1 = int32(_a_F_AfterTriggerBeginQuery_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerBeginQuery[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerBeginQuery[0])) = v3 + int32(1)
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[0]))
	if v3 != 0 {
		F_MemoryContextDelete(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v7 = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[1])) = v7
			*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[2])) = v7
			*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[3])) = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[4])) = int64(-4294967296)
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[5])) = v19
			*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[6])) = v19
			return
		}
	} else {
		*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[3])) = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[4])) = int64(-4294967296)
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[5])) = v19
		*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerEndXact[6])) = v19
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerFireDeferred[0]))
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
	v14 = F_afterTriggerMarkEvents(m, int32(_a_F_AfterTriggerFireDeferred_0), v12, v12)
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
	v16 = int32(_a_F_AfterTriggerFireDeferred_1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerFireDeferred[1]))
	v19 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AfterTriggerFireDeferred[1])) = v18 + v19
	v25 = F_afterTriggerInvokeEvents(m, int32(_a_F_AfterTriggerFireDeferred_0), v18, int32(0), v19)
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = v11 & int32(939524096)
	if v13 == int32(134217728) {
		v23 = int32(24)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	if v13 != int32(805306368) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v13 == int32(268435456) {
		v23 = int32(12)
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = int32(16)
	goto L1
L6:
	;
	v23 = int32(4)
	goto L1
L7:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	if base.Ui32(v83) < base.Ui32(v85) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if base.Ui32(v23+int32(28)) <= base.Ui32(v27-v28) {
		v82 = v24
		v83 = v27
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[0]))
	if v33 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[1]))
	v43 = F_AllocSetContextCreateInternal(m, v38, int32(_a_F_afterTriggerAddEvent_0), int32(0), int32(_a_F_afterTriggerAddEvent_1), int32(_a_F_afterTriggerAddEvent_2))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v46 = v33
	goto L14
L14:
	;
	if v24 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[0])) = v43
	v46 = v43
	goto L14
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v49 = v48 - v24
	v50 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if base.Ui32(v48-v54) < base.Ui32(int32(2801)) {
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
	v67 = F_MemoryContextAlloc(m, v46, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L26
	}
L20:
	;
	v58 = v49 << (uint(v50) % 32)
	goto L22
L21:
	;
	v58 = int32(base.Ui32(v49) >> (uint(v50) % 32))
	goto L22
L22:
	;
	if base.Ui32(int32(_a_F_afterTriggerAddEvent_3)) <= base.Ui32(v58) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v61 = int32(_a_F_afterTriggerAddEvent_3)
	goto L25
L24:
	;
	v61 = v58
	goto L25
L25:
	;
	v66 = v61
	goto L19
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(0)
	v71 = v67 + v66
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
	v82 = v67
	v83 = v81
	goto L7
L30:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v23 != 0 {
		goto L62
	} else {
		goto L63
	}
L31:
	;
	v91 = v83
	v92 = v85
	goto L34
L32:
	;
	goto L33
L33:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v184 = v182 - int32(28)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v184)+24)) = v185
	v187 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v184)+16)) = v187
	v189 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v184)+8)) = v189
	v191 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v194 != 0 {
		goto L58
	} else {
		goto L59
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v96 != v97 {
		v166 = v92
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if base.Ui32(v170) < base.Ui32(v171) {
		v219 = v170
		goto L30
	} else {
		goto L57
	}
L36:
	;
	goto L35
L37:
	;
	v168 = v91 + int32(28)
	if base.Ui32(v168) < base.Ui32(v166) {
		v91 = v168
		v92 = v166
		goto L34
	} else {
		goto L56
	}
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v99 != v100 {
		v166 = v92
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v102 != 0 {
		v166 = v92
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v103 != v104 {
		v166 = v92
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v106 != v107 {
		v166 = v92
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v109 != v110 {
		v166 = v92
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v114 = int32(0)
	if base.B2i32(v112 == v114)|base.B2i32(v113 == v114) != 0 {
		v160 = base.B2i32(v112|v113 == v114)
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	if v160 != 0 {
		v170 = v91
		v171 = v165
		goto L36
	} else {
		goto L55
	}
L45:
	;
	goto L44
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v128 != v129 {
		v160 = int32(0)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v131 = int32(1)
	if v128 <= v131 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v134 = v131
	goto L50
L49:
	;
	v134 = v128
	goto L50
L50:
	;
	v135 = int32(8)
	v140 = int32(0)
	goto L51
L51:
	;
	v148 = v140 << (uint(int32(2)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v112+v135+v148)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v113+v135+v148)))
	v153 = base.B2i32(v150 == v152)
	if v150 != v152 {
		v160 = v153
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v160 = v153
	goto L45
L53:
	;
	v156 = v140 + int32(1)
	if v156 != v134 {
		v140 = v156
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v166 = v165
	goto L37
L56:
	;
	v170 = v168
	v171 = v166
	goto L36
L57:
	;
	goto L33
L58:
	;
	v195 = int32(_a_F_afterTriggerAddEvent_4)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[2]))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[2])) = v199
	v201 = F_bms_copy(m, v194)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L15
	} else {
		goto L61
	}
L59:
	;
	v205 = int32(0)
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182-int32(12)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v182-int32(4)))) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v184
	v219 = v184
	goto L30
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_afterTriggerAddEvent[2])) = v196
	v205 = v201
	goto L60
L62:
	;
	base.MemoryCopy(m, v224, l1, v23)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v219 - v224 | v227&int32(-134217728)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v233 = v232 + v23
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v233
	return
}
