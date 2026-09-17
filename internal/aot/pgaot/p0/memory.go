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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2*int32(36) + int32(_a_F_MemoryContextCreate_0)
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
func F_MemoryContextMemAllocated(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == int32(0) {
		v27 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v27
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == int32(0) {
		v27 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v5
	v14 = v8
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = v15 + v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v17 != 0 {
		v13 = v16
		v14 = v17
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v27 = v16
	goto L1
L6:
	;
	v19 = v14
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v22 != 0 {
		v13 = v16
		v14 = v22
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v23 != l0 {
		v19 = v23
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
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int64
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = l1
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	m.T0[v18].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, int32(1771), v11+int32(92), l4, l5)
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
		v68 = v21
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v68 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_MemoryContextStatsInternal[0]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_MemoryContextStatsInternal[1]))
	v30 = m.G0
	v33 = v29 - (v30 - int32(1))
	v35 = v33 >> (uint(int32(31)) % 32)
	goto L5
L5:
	;
	v42 = int32(0)
	if base.B2i32(v27 < v33^v35-v35)&base.B2i32(v29 != int32(0))|base.B2i32(v21 == v42)|base.B2i32(l3 <= v42) != 0 {
		v68 = v21
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v49 = v21
	v55 = int32(0)
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_MemoryContextStatsInternal(m, v49, v56+int32(1), l2, l3, l4, l5)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v68 = v61
	goto L3
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
	if v61 == int32(0) {
		v68 = v61
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v65 = v55 + int32(1)
	if v65 < l3 {
		v49 = v61
		v55 = v65
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v75 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v75
	v81 = v68
	v87 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	m.G0 = v11 + int32(96)
	return
L15:
	;
	v88 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+32))
	m.T0[v94].(func(*base.Module, int32, int32, int32, int32, int32))(m, v81, v88, v88, v11+int32(72), v88)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	if l5 != 0 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v98 = v87 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v99 != 0 {
		v81 = v99
		v87 = v98
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v103 = v81
	goto L19
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	if v108 != 0 {
		v81 = v108
		v87 = v98
		goto L15
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v109 != l0 {
		v103 = v109
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v196 + v197
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v200 + v201
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v204 + v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v208 + v209
	goto L14
L24:
	;
	v111 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_MemoryContextStatsInternal[2]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if v111 < v114 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v157 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	v120 = v111
	goto L30
L28:
	;
	goto L29
L29:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v143 - v144
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v143
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v144
	v153 = F_pg_fprintf(m, v113, int32(_a_F_MemoryContextStatsInternal_0), v11)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v127 = F_pg_fprintf(m, v113, int32(_a_F_MemoryContextStatsInternal_1), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v130 = v120 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if v130 < v131 {
		v120 = v130
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L23
L35:
	;
	if v157 == int32(0) {
		goto L23
	} else {
		goto L36
	}
L36:
	;
	F_errhidestmt(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errhidecontext(m)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v169 - v165
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v98
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v169
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v176
	F_errmsg_internal(m, int32(_a_F_MemoryContextStatsInternal_2), v11+int32(32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_MemoryContextStatsInternal_3), int32(956), int32(_a_F_MemoryContextStatsInternal_4))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L23
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
			base.MemoryCopy(m, v12, l1, v8)
		} else {
		}
		return v12
	}
}
