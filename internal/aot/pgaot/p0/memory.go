package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MemoryContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2*int32(36) + int32(1768672)
	if l3 != 0 {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v22
		if v22 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l0
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = l0
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v26)
		return
	} else {
		v28 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v28
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v28)
		return
	}
}
func F_MemoryContextMemAllocated(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v5
	v10 = v6
	goto L4
L2:
	;
	v23 = v5
	goto L3
L3:
	;
	return v23
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = v11 + v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v13 != 0 {
		v9 = v12
		v10 = v13
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v23 = v12
	goto L3
L6:
	;
	v15 = v10
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v18 != 0 {
		v9 = v12
		v10 = v18
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v19 != l0 {
		v15 = v19
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_MemoryContextStatsInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = l1
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	m.T0[v18].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(1787), v11+int32(92), l4, l5)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if l2 < v22 {
		v66 = v21
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v66 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v30 = m.G0
	v33 = v29 - (v30 - int32(1))
	v35 = v33 >> (uint(int32(31)) % 32)
	goto L5
L5:
	;
	if base.B2i32(v27 < v33^v35-v35)&base.B2i32(v29 != int32(0)) != 0 {
		v66 = v21
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v21 == int32(0) {
		v66 = v21
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if l3 <= int32(0) {
		v66 = v21
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v47 = v21
	v53 = int32(0)
	goto L9
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_MemoryContextStatsInternal(m, v47, v54+int32(1), l2, l3, l4, l5)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v66 = v59
	goto L3
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v59 == int32(0) {
		v66 = v59
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v63 = v53 + int32(1)
	if v63 < l3 {
		v47 = v59
		v53 = v63
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v73 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v73
	v79 = v66
	v85 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	m.G0 = v11 + int32(96)
	return
L17:
	;
	v86 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	m.T0[v92].(func(*base.Module, int32, int32, int32, int32, int32))(m, v79, v86, v86, v11+int32(72), v86)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	if l5 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v96 = v85 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v97 != 0 {
		v79 = v97
		v85 = v96
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v101 = v79
	goto L21
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+28))
	if v106 != 0 {
		v79 = v106
		v85 = v96
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v107 != l0 {
		v101 = v107
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v194 + v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v198 + v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v202 + v203
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v206 + v207
	goto L16
L26:
	;
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[900]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if v109 < v112 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v155 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L29:
	;
	v118 = v109
	goto L32
L30:
	;
	goto L31
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v141 - v142
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v141
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v142
	v151 = F_pg_fprintf(m, v111, int32(750985), v11)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L36
	}
L32:
	;
	v125 = F_pg_fprintf(m, v111, int32(746690), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v128 = v118 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if v128 < v129 {
		v118 = v128
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
	if v155 == int32(0) {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	F_errhidestmt(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errhidecontext(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v167 - v163
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v96
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v167
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v174
	F_errmsg_internal(m, int32(449514), v11+int32(32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(492961), int32(956), int32(312802))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L25
}
func F_MemoryContextStrdup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = int32(0)
	v4 = F_strlen(m, l1)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v3)
	v8 = v4 + int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, v8, v3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v16 = F__emscripten_memcpy_bulkmem(m, v12, l1, v8)
			mBase = m.M
			v17 = v16
		} else {
			v17 = v12
		}
		return v17
	}
}
