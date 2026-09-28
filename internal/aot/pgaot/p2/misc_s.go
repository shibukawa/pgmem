package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_SanityCheckBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v10&int32(1) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return v214
L2:
	;
	v15 = int32(0)
	v17 = F_errstart(m, l1, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v10&int32(2) != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	return int32(0)
L6:
	;
	if v17 == int32(0) {
		v214 = v15
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_0), v8)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(664), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v214 = v15
	goto L1
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if base.Ui32(v82-int32(86400001)) <= base.Ui32(int32(-86400003)) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v37 != 0 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v10&int32(4) == int32(0) {
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v38 = int32(0)
	v40 = F_errstart(m, l1, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if v40 == int32(0) {
		v214 = v38
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_3), v8-int32(-64))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(675), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v214 = v38
	goto L1
L21:
	;
	v62 = int32(0)
	v64 = F_errstart(m, l1, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if v64 == int32(0) {
		v214 = v62
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_4), v8+int32(48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(689), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v214 = v62
	goto L1
L27:
	;
	v87 = int32(0)
	v89 = F_errstart(m, l1, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v109 = int32(0)
	if base.B2i32(v10&int32(16) == v109)|base.B2i32(v82 == int32(-1)) == v109 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	if v89 == int32(0) {
		v214 = v87
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_5), v8+int32(16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(700), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v214 = v87
	goto L1
L35:
	;
	v116 = int32(0)
	v118 = F_errstart(m, l1, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v136 = int32(1)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v137 != 0 {
		v214 = v136
		goto L1
	} else {
		goto L43
	}
L38:
	;
	if v118 == int32(0) {
		v214 = v116
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	F_errmsg(m, int32(_a_F_SanityCheckBackgroundWorker_6), v8+int32(32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_SanityCheckBackgroundWorker_1), int32(715), int32(_a_F_SanityCheckBackgroundWorker_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v214 = v116
	goto L1
L43:
	;
	v139 = l0 + int32(96)
	if (l0^v139)&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v214 = v136
	goto L1
L45:
	;
	goto L44
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v193)
	if v193&int32(255) == int32(0) {
		goto L45
	} else {
		goto L61
	}
L47:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v192 = l0
	v193 = v145
	v194 = v139
	goto L46
L48:
	;
	goto L49
L49:
	;
	if l0&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v149 = l0
	v151 = v139
	goto L53
L51:
	;
	v163 = l0
	v165 = v139
	goto L52
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v170 = int32(-2139062144)
	if (int32(16843008)-v167|v167)&v170 != v170 {
		v192 = v163
		v193 = v167
		v194 = v165
		goto L46
	} else {
		goto L57
	}
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v152)
	if v152 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L54:
	;
	v163 = v159
	v165 = v157
	goto L52
L55:
	;
	v156 = int32(1)
	v157 = v151 + v156
	v159 = v149 + v156
	if v159&int32(3) != 0 {
		v149 = v159
		v151 = v157
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v175 = v163
	v176 = v167
	v177 = v165
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v176
	v179 = int32(4)
	v180 = v177 + v179
	v182 = v175 + v179
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v187 = int32(-2139062144)
	if (int32(16843008)-v184|v184)&v187 == v187 {
		v175 = v182
		v176 = v184
		v177 = v180
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v192 = v182
	v193 = v184
	v194 = v180
	goto L46
L60:
	;
	goto L59
L61:
	;
	v201 = v192
	v203 = v194
	goto L62
L62:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)) = uint8(v204)
	v206 = int32(1)
	if v204 != 0 {
		v201 = v201 + v206
		v203 = v203 + v206
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L45
L64:
	;
	goto L63
}
func F_ScanKeywordLookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v7 = int32(-1)
	v8 = F_strlen(m, l0)
	mBase = m.M
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v9) < base.Ui32(v8) {
		v56 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v56
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, l0, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v12 < int32(0) {
		v56 = v7
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v18 <= v12 {
		v56 = v7
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+v12<<(uint(int32(1))%32)))))
	v27 = l0
	v28 = v20 + v25
	goto L7
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v35 = int32(1)
	if base.Ui32((v33-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v47 = v33 | int32(32)
	goto L14
L13:
	;
	v47 = v33
	goto L14
L14:
	;
	if v34 == v47 {
		v27 = v27 + v35
		v28 = v28 + v35
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v56 = v7
	goto L1
L16:
	;
	v51 = int32(-1)
	goto L18
L17:
	;
	v51 = v12
	goto L18
L18:
	;
	v56 = v51
	goto L1
}
func F_SetLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	v2 = int32(0)
	v5 = base.AtomicRmwOr32(m, v2, int32(_a_F_SetLatch_0), v2)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v9 = int32(0)
	v12 = base.AtomicRmwOr32(m, v9, int32(_a_F_SetLatch_0), v9)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v9 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[0]))
	if v20 == v16 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[1]))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v50 = F_pgmem_kill(m, v16, int32(23))
	mBase = m.M
	goto L1
L8:
	;
	m.G0 = v24 + int32(16)
	return
L9:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+15)) = uint8(v30)
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[2]))
	v38 = F_write(m, v34, v24+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v38 {
		goto L8
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_SetLatch[3]))
	if v42 == int32(27) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_SetVarReturningType_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v3 = int32(0)
	if l0 == v3 {
		v41 = v3
		return v41
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(67) {
			if v7 != int32(6) {
				v38 = F_expression_tree_walker_impl(m, l0, int32(1135), l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = v38
					return v41
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v12 != v13 {
					v41 = v3
					return v41
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v15 != v16 {
						v41 = v3
						return v41
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
						return int32(0)
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 + int32(1)
			v28 = F_query_tree_walker_impl(m, l0, int32(1135), l1, int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v32 - int32(1)
				return v28
			}
		}
	}
}
func F_SetupHistoricSnapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_SetupHistoricSnapshot[0])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_SetupHistoricSnapshot[1])) = l0
	return
}
func F_ShowUsage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	v9 = m.G0
	v11 = v9 - int32(368)
	m.G0 = v11
	v14 = v11 + int32(184)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1)
	v24 = F___syscall_ret(m, int32(0))
	mBase = m.M
	F_gettimeofday(m, v11+int32(336))
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+344))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[0]))
	if v29 < v31 {
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v11)+336))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+336)) = v33 - int64(1)
		v39 = v29 + int32(_a_F_ShowUsage_0)
	} else {
		v39 = v29
	}
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+208))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v11)+184))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[1]))
	if v28 < v43 {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v28 + int32(_a_F_ShowUsage_0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+184)) = v41 - int64(1)
	} else {
	}
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v11)+200))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[2]))
	if v40 < v53 {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v40 + int32(_a_F_ShowUsage_0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+200)) = v51 - int64(1)
	} else {
	}
	v62 = v11 + int32(352)
	F_initStringInfo(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		return
	} else {
		F_appendStringInfoString(m, v62, int32(_a_F_ShowUsage_1))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			v68 = *(*int64)(unsafe.Add(mBase, uint32(v11)+336))
			v70 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[3]))
			v71 = v68 - v70
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+176)) = uint32(v71)
			v74 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+180)) = v39 - v74
			v77 = *(*int64)(unsafe.Add(mBase, uint32(v11)+184))
			v79 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[4]))
			v80 = v77 - v79
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+160)) = uint32(v80)
			v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+192))
			v84 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v82 - v84
			v87 = *(*int64)(unsafe.Add(mBase, uint32(v11)+200))
			v89 = *(*int64)(unsafe.Add(mBase, _c_F_ShowUsage[5]))
			v90 = v87 - v89
			*(*uint32)(unsafe.Add(mBase, uint32(v11)+168)) = uint32(v90)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+208))
			v94 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+172)) = v92 - v94
			F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_2), v11+int32(160))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = v40
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+152)) = uint32(v51)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+148)) = v28
				*(*uint32)(unsafe.Add(mBase, uint32(v11)+144)) = uint32(v41)
				F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_3), v11+int32(144))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v111
					F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_4), v11+int32(128))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+244))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = v118
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+248))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = v120
						v123 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[6]))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v118 - v123
						v127 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[7]))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v120 - v127
						F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_5), v11+int32(112))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+240))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v135
							v138 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[8]))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v135 - v138
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+236))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v141
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+232))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v143
							v146 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[9]))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v141 - v146
							v150 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[10]))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v143 - v150
							F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_6), v11+int32(80))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v11)+256))
								*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v160
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+252))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v162
								v164 = *(*int32)(unsafe.Add(mBase, uint32(v11)+260))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v164
								v167 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[11]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v164 - v167
								v171 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[12]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v160 - v171
								v175 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[13]))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v162 - v175
								F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_7), v11+int32(48))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return
								} else {
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v11)+264))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v183
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v11)+268))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v185
									v188 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[14]))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v183 - v188
									v192 = *(*int32)(unsafe.Add(mBase, _c_F_ShowUsage[15]))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v185 - v192
									F_appendStringInfo(m, v62, int32(_a_F_ShowUsage_8), v11+int32(32))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return
									} else {
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
										v201 = *(*int32)(unsafe.Add(mBase, uint32(v11)+356))
										v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v201-int32(1)))))
										if v205 == int32(10) {
											v209 = v201 - int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+356)) = v209
											v212 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v209+v200))) = uint8(v212)
										} else {
										}
										v217 = F_errstart(m, int32(15), int32(0))
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return
										} else {
											if v217 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
												F_errmsg_internal(m, int32(_a_F_ShowUsage_9), v11+int32(16))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													v225 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v225
													F_errdetail_internal(m, int32(_a_F_ShowUsage_9), v11)
													mBase = m.M
													v229 = m.ExcPending
													if v229 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ShowUsage_10), int32(_a_F_ShowUsage_11), int32(_a_F_ShowUsage_12))
														mBase = m.M
														v234 = m.ExcPending
														if v234 != 0 {
															return
														} else {
															v235 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
															F_pfree(m, v235)
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
																return
															} else {
																m.G0 = v11 + int32(368)
																return
															}
														}
													}
												}
											} else {
												v235 = *(*int32)(unsafe.Add(mBase, uint32(v11)+352))
												F_pfree(m, v235)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return
												} else {
													m.G0 = v11 + int32(368)
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
			}
		}
	}
}
func F_SignalHandlerForCrashExit(m *base.Module, l0 int32, l1 int32) {
	F__Exit(m, int32(2))
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SplitDirectoriesString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v9 = l0
	goto L1
L1:
	;
	v17 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
	goto L3
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v17 == int32(32))|base.B2i32(base.Ui32((v17-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v9 = v9 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	return int32(1)
L6:
	;
	v33 = v9
	v34 = v27
	goto L7
L7:
	;
	if v34 != int32(34) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v167)
	v169 = F_strlen(m, v109)
	mBase = m.M
	if base.Ui32(int32(1024)) <= base.Ui32(v169) {
		goto L55
	} else {
		goto L56
	}
L10:
	;
	return int32(0)
L11:
	;
	v112 = v106
	goto L43
L12:
	;
	if v34 == int32(0) {
		v66 = v33
		v71 = v33
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v74 = v33 + int32(1)
	v75 = int32(34)
	v76 = F___strchrnul(m, v74, v75)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 == v75 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	if v33 == v71 {
		goto L10
	} else {
		goto L26
	}
L16:
	;
	if v34 == int32(44) {
		v66 = v33
		v71 = v33
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = v33
	v46 = v34
	v47 = v33
	goto L18
L18:
	;
	v49 = v44 + int32(1)
	v50 = base.I32_extend8_s(v46)
	goto L20
L19:
	;
	v66 = v49
	v71 = v60
	goto L15
L20:
	;
	if base.B2i32(v50 == int32(32))|base.B2i32(base.Ui32((v50-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v60 = v47
	goto L23
L22:
	;
	v60 = v49
	goto L23
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v61 == int32(0) {
		v66 = v49
		v71 = v60
		goto L15
	} else {
		goto L24
	}
L24:
	;
	if v61 != int32(44) {
		v44 = v49
		v46 = v61
		v47 = v60
		goto L18
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	v106 = v66
	v109 = v33
	v111 = v71
	goto L11
L27:
	;
	if v82 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L28:
	;
	v82 = v76
	goto L30
L29:
	;
	v82 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	v90 = v82
	goto L32
L32:
	;
	v92 = v90 + int32(1)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v93 != int32(34) {
		v106 = v92
		v109 = v74
		v111 = v90
		goto L11
	} else {
		goto L34
	}
L33:
	;
	goto L10
L34:
	;
	v96 = F_strlen(m, v90)
	mBase = m.M
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryCopy(m, v90, v92, v96)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v98 = int32(34)
	v99 = F___strchrnul(m, v92, v98)
	mBase = m.M
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v101 == v98 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v105 != 0 {
		v90 = v105
		goto L32
	} else {
		goto L42
	}
L39:
	;
	v105 = v99
	goto L41
L40:
	;
	v105 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L33
L43:
	;
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	goto L45
L44:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v130 == int32(44) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if base.B2i32(v120 == int32(32))|base.B2i32(base.Ui32((v120-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v112 = v112 + int32(1)
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v135 = v112
	goto L50
L48:
	;
	goto L49
L49:
	;
	if v130 == int32(0) {
		v163 = v112
		goto L9
	} else {
		goto L54
	}
L50:
	;
	v140 = v135 + int32(1)
	v141 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135)+1)))
	goto L52
L52:
	;
	if base.B2i32(v141 == int32(32))|base.B2i32(base.Ui32((v141-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v135 = v140
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v163 = v140
	goto L9
L54:
	;
	goto L10
L55:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+1023)) = uint8(v172)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v174 = F_pstrdup(m, v109)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	return int32(0)
L59:
	;
	F_canonicalize_path_enc(m, v174)
	mBase = m.M
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v180 = F_lappend(m, v179, v174)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v180
	if v130 != int32(44) {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v33 = v163
	v34 = v185
	goto L7
}
func F_StrategyNotifyBgWriter(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(0), int32(1))
	if v7 != 0 {
		F_s_lock(m, v4, int32(_a_F_StrategyNotifyBgWriter_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
			v14 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v12))), uint32(v14))
			return
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_StrategyNotifyBgWriter[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
		v14 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v12))), uint32(v14))
		return
	}
}
func F___shgetc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v15 = v10 + base.I64_extend_i32_s(v11-v12)
	if base.B2i32(v7 != int64(0))&base.B2i32(v7 <= v15) == int32(0) {
		v20 = F___uflow(m, l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v20 {
				v41 = v15 + int64(1)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v44 == int64(0) {
					v53 = v43
				} else {
					v47 = v44 - v41
					if base.I64_extend_i32_s(v43-v42) <= v47 {
						v53 = v43
					} else {
						v53 = v42 + base.I32_wrap_i64(v47)
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v53
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v41 + base.I64_extend_i32_s(v56-v42)
				if base.Ui32(v42) <= base.Ui32(v56) {
					*(*uint8)(unsafe.Add(mBase, uint32(v42-int32(1)))) = uint8(v20)
				} else {
				}
				return v20
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v28 = v27
				v29 = v26
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v28
				*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v15 + base.I64_extend_i32_s(v29-v28)
				return int32(-1)
			}
		}
	} else {
		v28 = v11
		v29 = v12
		*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(-1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v28
		*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v15 + base.I64_extend_i32_s(v29-v28)
		return int32(-1)
	}
}
func F___stdio_close(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3 = m.Wasi_snapshot_preview1.Fd_close(m, v2)
	mBase = m.M
	if v3 == int32(0) {
		v10 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___stdio_close[0])) = v3
		v10 = int32(-1)
	}
	return v10
}
func F_scanGetCandidate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(-1)
	v13 = l1 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = v14
	goto L2
L1:
	;
	return base.B2i32(base.Ui32(v22) <= base.Ui32(v49&int32(_a_F_scanGetCandidate_0)))
L2:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v17 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v77 = v40 + int32(20)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v22<<(uint(int32(2))%32))))
	v84 = v40 + v81&int32(_a_F_scanGetCandidate_1)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v89)+6)))
	if v91&int32(32) != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v41) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_scanGetCandidate[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v17^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L4
L6:
	;
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_scanGetCandidate[1]))
	v40 = v34 + v17<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v49 = int32(base.Ui32(v41+int32(_a_F_scanGetCandidate_2)) >> (uint(int32(2)) % 32))
	goto L10
L9:
	;
	v49 = int32(0)
	goto L10
L10:
	;
	if base.Ui32(v49&int32(_a_F_scanGetCandidate_0)) < base.Ui32(v22) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40+v53)))
	if v55 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	goto L3
L14:
	;
	F_UnlockReleaseBuffer(m, v17)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = F_ReadBuffer(m, v64, v55)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L1
L19:
	;
	F_LockBufferInternal(m, v65, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_UnlockReleaseBuffer(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v73 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v73)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v65
	v17 = v65
	goto L2
L22:
	;
	v95 = v22 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v95)
	v97 = int32(_a_F_scanGetCandidate_0)
	v98 = v95 & v97
	if base.Ui32(v49&v97) < base.Ui32(v98) {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v169 = v49 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v169)
	goto L1
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v77+v98<<(uint(int32(2))%32))))
	v108 = v40 + v105&int32(_a_F_scanGetCandidate_1)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v111 = int32(16)
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+2)))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108))))
	if v109|v110<<(uint(v111)%32) == v114|v115<<(uint(v111)%32) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v125 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	goto L26
L28:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)))
	if v121 == v122 {
		v125 = int32(1)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v125 = int32(0)
	goto L27
L31:
	;
	goto L30
L32:
	;
	goto L33
L33:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v137 = v135 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v137)
	v139 = int32(_a_F_scanGetCandidate_0)
	v140 = v137 & v139
	if base.Ui32(v49&v139) < base.Ui32(v140) {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L1
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v77+v140<<(uint(int32(2))%32))))
	v150 = v40 + v147&int32(_a_F_scanGetCandidate_1)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v153 = int32(16)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+2)))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150))))
	if v151|v152<<(uint(v153)%32) == v156|v157<<(uint(v153)%32) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v167 != 0 {
		goto L33
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+4)))
	if v163 == v164 {
		v167 = int32(1)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v167 = int32(0)
	goto L37
L41:
	;
	goto L40
L42:
	;
	goto L34
}
func F_scanexp(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int64
	_ = v115
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v156 int64
	_ = v156
	var v161 int64
	_ = v161
	var v164 int32
	_ = v164
	var v175 int64
	_ = v175
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v7 != v8 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	return v175
L2:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v161 < int64(0) {
		v175 = int64(-9223372036854775807 - 1)
		goto L1
	} else {
		goto L54
	}
L3:
	;
	if base.Ui32(v49) < base.Ui32(int32(-10)) {
		goto L2
	} else {
		goto L19
	}
L4:
	;
	v49 = v18 - int32(58)
	v50 = v18
	v51 = int32(0)
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v21 != v22 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	switch v18 - int32(43) {
	case 0, 2:
		goto L5
	default:
		goto L4
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 + int32(1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v18 = v13
	goto L6
L8:
	;
	goto L9
L9:
	;
	v14 = F___shgetc(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int64(0)
L11:
	;
	v18 = v14
	goto L6
L12:
	;
	v36 = v30 - int32(58)
	if base.B2i32(l1 == int32(0))|base.B2i32(base.Ui32(int32(-11)) < base.Ui32(v36)) != 0 {
		v49 = v36
		v50 = v30
		v51 = base.B2i32(v18 == int32(45))
		goto L3
	} else {
		goto L17
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v21 + int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v30 = v27
	goto L12
L14:
	;
	goto L15
L15:
	;
	v28 = F___shgetc(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v30 = v28
	goto L12
L17:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if v40 < int64(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 - int32(1)
	goto L2
L19:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v50-int32(48)) {
		v146 = int64(0)
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if int64(0) <= v147 {
		goto L48
	} else {
		goto L49
	}
L21:
	;
	v61 = int32(0)
	v62 = v50
	goto L22
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v68 != v69 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v87 = base.I64_extend_i32_s(v83)
	if base.Ui32(int32(10)) <= base.Ui32(v79) {
		v146 = v87
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v78 = int32(48)
	v79 = v77 - v78
	v83 = v62 + v61*int32(10) - v78
	if base.B2i32(base.Ui32(v79) <= base.Ui32(int32(9)))&base.B2i32(v83 < int32(214748364)) != 0 {
		v61 = v83
		v62 = v77
		goto L22
	} else {
		goto L29
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68 + int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v77 = v74
	goto L24
L26:
	;
	goto L27
L27:
	;
	v75 = F___shgetc(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v77 = v75
	goto L24
L29:
	;
	goto L23
L30:
	;
	v93 = v77
	v95 = v87
	goto L31
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v100 != v101 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v111) {
		v146 = v115
		goto L20
	} else {
		goto L39
	}
L33:
	;
	v111 = v109 - int32(48)
	v115 = base.I64_extend_i32_u(v93) + v95*int64(10) - int64(48)
	if base.B2i32(base.Ui32(v111) <= base.Ui32(int32(9)))&base.B2i32(v115 < int64(92233720368547758)) != 0 {
		v93 = v109
		v95 = v115
		goto L31
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v109 = v106
	goto L33
L35:
	;
	goto L36
L36:
	;
	v107 = F___shgetc(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v109 = v107
	goto L33
L38:
	;
	goto L32
L39:
	;
	goto L40
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v127 != v128 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v146 = v115
	goto L20
L42:
	;
	if base.Ui32(v136-int32(48)) < base.Ui32(int32(10)) {
		goto L40
	} else {
		goto L47
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127 + int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v136 = v133
	goto L42
L44:
	;
	goto L45
L45:
	;
	v134 = F___shgetc(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v136 = v134
	goto L42
L47:
	;
	goto L41
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v150 - int32(1)
	goto L50
L49:
	;
	goto L50
L50:
	;
	if v51 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v156 = int64(0) - v146
	goto L53
L52:
	;
	v156 = v146
	goto L53
L53:
	;
	v175 = v156
	goto L1
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v164 - int32(1)
	return int64(-9223372036854775807 - 1)
}
func F_search_indexed_tlist_for_phv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
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
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L63
	} else {
		goto L65
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v233
L3:
	;
	v233 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v19 <= v18 {
		v233 = v18
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = v19
	goto L9
L8:
	;
	v25 = v22
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v28 = v18
	goto L12
L10:
	;
	v224 = F_makeVarFromTargetEntry(m, l2, v39)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L63
	} else {
		goto L64
	}
L11:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v168 = int32(0)
	if v166 == v168 {
		goto L49
	} else {
		goto L50
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v233 = int32(0)
	goto L2
L14:
	;
	v163 = v28 + int32(1)
	if v163 != v25 {
		v28 = v163
		goto L12
	} else {
		goto L47
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v43 != int32(321) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v46 != v47 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	switch l3 - int32(1) {
	case 0:
		goto L11
	case 1:
		goto L19
	default:
		goto L18
	}
L18:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = int32(0)
	if base.B2i32(v109 == v111)|base.B2i32(v110 == v111) != 0 {
		v157 = base.B2i32(v109|v110 == v111)
		goto L36
	} else {
		goto L37
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = int32(0)
	if v51 == v53 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v106 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L21:
	;
	v106 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v52 == int32(0) {
		v99 = v53
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v106 = v99
	goto L20
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v63 < v62 {
		v99 = v53
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v65 = int32(1)
	if v62 <= v65 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v68 = v65
	goto L29
L28:
	;
	v68 = v62
	goto L29
L29:
	;
	v69 = int32(8)
	v74 = int32(0)
	goto L30
L30:
	;
	v81 = v74 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v51+v69+v81)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v52+v69+v81)))
	v88 = v83 & (v85 ^ int32(-1))
	v90 = base.B2i32(v88 == int32(0))
	if v88 != 0 {
		v99 = v90
		goto L24
	} else {
		goto L32
	}
L31:
	;
	v99 = v90
	goto L24
L32:
	;
	v92 = v74 + int32(1)
	if v92 != v68 {
		v74 = v92
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L10
L35:
	;
	if v157 != 0 {
		goto L10
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v125 != v126 {
		v157 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v128 = int32(1)
	if v125 <= v128 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v131 = v128
	goto L41
L40:
	;
	v131 = v125
	goto L41
L41:
	;
	v132 = int32(8)
	v137 = int32(0)
	goto L42
L42:
	;
	v145 = v137 << (uint(int32(2)) % 32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v109+v132+v145)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v110+v132+v145)))
	v150 = base.B2i32(v147 == v149)
	if v147 != v149 {
		v157 = v150
		goto L36
	} else {
		goto L44
	}
L43:
	;
	v157 = v150
	goto L36
L44:
	;
	v153 = v137 + int32(1)
	if v153 != v131 {
		v137 = v153
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L1
L47:
	;
	goto L13
L48:
	;
	if v221 == int32(0) {
		goto L1
	} else {
		goto L62
	}
L49:
	;
	v221 = int32(1)
	goto L48
L50:
	;
	goto L51
L51:
	;
	if v167 == int32(0) {
		v214 = v168
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v221 = v214
	goto L48
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v178 < v177 {
		v214 = v168
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v180 = int32(1)
	if v177 <= v180 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = v180
	goto L57
L56:
	;
	v183 = v177
	goto L57
L57:
	;
	v184 = int32(8)
	v189 = int32(0)
	goto L58
L58:
	;
	v196 = v189 << (uint(int32(2)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v166+v184+v196)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v167+v184+v196)))
	v203 = v198 & (v200 ^ int32(-1))
	v205 = base.B2i32(v203 == int32(0))
	if v203 != 0 {
		v214 = v205
		goto L52
	} else {
		goto L60
	}
L59:
	;
	v214 = v205
	goto L52
L60:
	;
	v207 = v189 + int32(1)
	if v207 != v183 {
		v189 = v207
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L10
L63:
	;
	return int32(0)
L64:
	;
	v228 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+40)) = uint16(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+36)) = v228
	v233 = v224
	goto L2
L65:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v250 = F_bmsToString(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v253 = F_bmsToString(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v250
	F_errmsg_internal(m, int32(_a_F_search_indexed_tlist_for_phv_0), v12)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_search_indexed_tlist_for_phv_1), int32(3043), int32(_a_F_search_indexed_tlist_for_phv_2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_session_user(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_session_user[0]))
	v7 = F_GetUserNameFromId(m, v5, v3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(534), v3, base.I64_extend_i32_u(v7))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
func F_set_upper_references(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v229 int32
	_ = v229
	var v237 float64
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
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
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 float64
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = int32(12)
	v29 = v23*v24 + v24
	goto L3
L2:
	;
	v29 = int32(12)
	goto L3
L3:
	;
	v30 = F_palloc(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v32 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+8)) = uint16(v32)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v22
	v36 = v30 + int32(12)
	if v22 == v32 {
		v98 = v36
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v110 = base.I32_div_s(v98-v36, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v112 != int32(369) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v39 <= int32(0) {
		v98 = v36
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v46 = v4
	v48 = v36
	goto L9
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v46<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v63 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v98 = v86
	goto L6
L11:
	;
	v89 = v46 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v89 < v90 {
		v46 = v89
		v48 = v86
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+9)) = uint8(v84)
	v86 = v48
	goto L11
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v66 != int32(321) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v66 != int32(6) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+8)) = uint8(v81)
	v86 = v48
	goto L11
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v71
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)) = uint16(v73)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+6)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v77
	v86 = v48 + int32(12)
	goto L11
L18:
	;
	goto L10
L19:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v137 == int32(0) {
		v293 = v4
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	if v115 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v118 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v122 = F_bms_make_singleton(m, v115)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v125 = F_remove_nulling_relids(m, v121, v122, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v130 = F_bms_make_singleton(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v133 = F_remove_nulling_relids(m, v128, v130, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v133
	goto L19
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v293
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v298 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = base.F64_add(v298, v298)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	v308 = F_fix_upper_expr_mutator(m, v297, v19)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L53
	}
L28:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v140 <= int32(0) {
		v293 = v4
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v155 = v4
	v156 = v4
	goto L30
L30:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v155<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v165 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v293 = v274
	goto L27
L32:
	;
	v271 = F_flatCopyTargetEntry(m, v163)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L50
	}
L33:
	;
	v249 = F_makeVarFromTargetEntry(m, int32(-2), v193)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L49
	}
L34:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v166 == int32(0) {
		v213 = v164
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v229 = v164
	goto L36
L36:
	;
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	v246 = F_fix_upper_expr_mutator(m, v229, v19)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L48
	}
L37:
	;
	v229 = v213
	goto L36
L38:
	;
	v169 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v170 <= v169 {
		v213 = v164
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v177 = v169
	v183 = v170
	goto L40
L40:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v177<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	if v165 == v194 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v213 = v204
	goto L37
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v197 = F_equal(m, v164, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	v200 = v183
	goto L44
L44:
	;
	v202 = v177 + int32(1)
	if v202 < v200 {
		v177 = v202
		v183 = v200
		goto L40
	} else {
		goto L47
	}
L45:
	;
	if v197 != 0 {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v200 = v199
	goto L44
L47:
	;
	goto L41
L48:
	;
	v259 = v246
	goto L32
L49:
	;
	v251 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+40)) = uint16(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+36)) = v251
	v259 = v249
	goto L32
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271)+4)) = v259
	v274 = F_lappend(m, v156, v271)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v277 = v155 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v277 < v278 {
		v155 = v277
		v156 = v274
		goto L30
	} else {
		goto L52
	}
L52:
	;
	goto L31
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v308
	F_pfree(m, v30)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v19 + int32(32)
	return
}
func F_setup_parser_errposition_callback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l0
	v9 = int32(_a_F_setup_parser_errposition_callback_0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_setup_parser_errposition_callback[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	*(*int32)(unsafe.Add(mBase, _c_F_setup_parser_errposition_callback[0])) = l0 + int32(8)
	return
}
func F_setup_pct_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v60 int64
	_ = v60
	var v61 float64
	_ = v61
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v77 float64
	_ = v77
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v100 int32
	_ = v100
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v20 = F_palloc(m, l0<<(uint(int32(5))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < l0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v36 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pg_qsort(m, v20, l0, int32(32), int32(1607))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L21
	}
L7:
	;
	v45 = v20 + v36<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = v36
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v36))))
	if v48 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v100 = v36 + int32(1)
	if v100 != l0 {
		v36 = v100
		goto L7
	} else {
		goto L20
	}
L10:
	;
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v51
	goto L9
L11:
	;
	goto L12
L12:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1+v36<<(uint(int32(3))%32))))
	v61 = base.F64_reinterpret_i64(v60)
	if base.F64_lt(v61, float64(0))|base.F64_gt(v61, float64(1))|base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v60&int64(9223372036854775807))) != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if l4 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = base.F64_mul(base.F64_convert_i64_s(l3-int64(1)), v61)
	v73 = base.F64_floor(v72)
	*(*float64)(unsafe.Add(mBase, uint32(v45)+16)) = base.F64_sub(v72, v73)
	v77 = float64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = base.I64_trunc_sat_f64_s(base.F64_add(base.F64_ceil(v72), v77))
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = base.I64_trunc_sat_f64_s(base.F64_add(v73, v77))
	goto L9
L15:
	;
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = int64(0)
	v87 = int64(1)
	v90 = base.I64_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(base.F64_convert_i64_s(l3), v61)))
	if v90 <= v87 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v93 = v87
	goto L19
L18:
	;
	v93 = v90
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v93
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v93
	goto L9
L20:
	;
	goto L8
L21:
	;
	m.G0 = v16 + int32(16)
	return v20
L22:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16))) = v61
	F_errmsg(m, int32(_a_F_setup_pct_info_0), v16)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_setup_pct_info_1), int32(693), int32(_a_F_setup_pct_info_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shell_archive_file(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	if l2 != 0 {
		v10 = F_pstrdup(m, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = v10
			*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[0]))
			v23 = F_replace_percent_placeholders(m, v18, int32(_a_F_shell_archive_file_0), int32(_a_F_shell_archive_file_1), v8+int32(96))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v27 = F_errstart(m, int32(12), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v23
						F_errmsg_internal(m, int32(_a_F_shell_archive_file_2), v8+int32(80))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(77), int32(_a_F_shell_archive_file_4))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v41 = F_fflush(m, int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
									v47 = F_pgl_system(m, v23)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
										*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
										if v47 != 0 {
											v63 = int32(255)
											if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
												v75 = int32(22)
											} else {
												v75 = int32(15)
											}
											v77 = v47 & int32(127)
											if v77 == int32(0) {
												v81 = F_errstart(m, v75, int32(0))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													if v81 == int32(0) {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
														F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															v145 = int32(102)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(112)
																		return base.B2i32(v47 == int32(0))
																	}
																}
															}
														}
													}
												}
											} else {
												v97 = F_errstart(m, v75, int32(0))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
														if v97 == int32(0) {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														} else {
															v109 = int32(_a_F_shell_archive_file_8)
															if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
																v114 = v77
																v115 = v109
																for {
																	v118 = v115 + int32(1)
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
																	if v119 != 0 {
																		v115 = v118
																		continue
																	} else {
																	}
																	v121 = v114 - int32(1)
																	if v121 != 0 {
																		v114 = v121
																		v115 = v118
																		continue
																	} else {
																		break
																	}
																	break
																}
																v123 = v118
															} else {
																v123 = v109
															}
															if v123 != 0 {
																v126 = v123
															} else {
																v126 = int32(_a_F_shell_archive_file_9)
															}
															*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
															*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
															F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																v145 = int32(118)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
																v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v23)
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v8 + int32(112)
																			return base.B2i32(v47 == int32(0))
																		}
																	}
																}
															}
														}
													} else {
														if v97 == int32(0) {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
															F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																v145 = int32(127)
																*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
																v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v23)
																		mBase = m.M
																		v158 = m.ExcPending
																		if v158 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v8 + int32(112)
																			return base.B2i32(v47 == int32(0))
																		}
																	}
																}
															}
														}
													}
												}
											}
										} else {
											F_pfree(m, v23)
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int32(0)
											} else {
												v163 = F_errstart(m, int32(14), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													if v163 == int32(0) {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
														F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(135), int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
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
						v41 = F_fflush(m, int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
							v47 = F_pgl_system(m, v23)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
								if v47 != 0 {
									v63 = int32(255)
									if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
										v75 = int32(22)
									} else {
										v75 = int32(15)
									}
									v77 = v47 & int32(127)
									if v77 == int32(0) {
										v81 = F_errstart(m, v75, int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											if v81 == int32(0) {
												F_pfree(m, v23)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
												F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													v145 = int32(102)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														}
													}
												}
											}
										}
									} else {
										v97 = F_errstart(m, v75, int32(0))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
												if v97 == int32(0) {
													F_pfree(m, v23)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													v109 = int32(_a_F_shell_archive_file_8)
													if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
														v114 = v77
														v115 = v109
														for {
															v118 = v115 + int32(1)
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
															if v119 != 0 {
																v115 = v118
																continue
															} else {
															}
															v121 = v114 - int32(1)
															if v121 != 0 {
																v114 = v121
																v115 = v118
																continue
															} else {
																break
															}
															break
														}
														v123 = v118
													} else {
														v123 = v109
													}
													if v123 != 0 {
														v126 = v123
													} else {
														v126 = int32(_a_F_shell_archive_file_9)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
													*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
													F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														v145 = int32(118)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(112)
																	return base.B2i32(v47 == int32(0))
																}
															}
														}
													}
												}
											} else {
												if v97 == int32(0) {
													F_pfree(m, v23)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
													F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int32(0)
													} else {
														v145 = int32(127)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(112)
																	return base.B2i32(v47 == int32(0))
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									F_pfree(m, v23)
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										v163 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											if v163 == int32(0) {
												m.G0 = v8 + int32(112)
												return base.B2i32(v47 == int32(0))
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
												F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(135), int32(_a_F_shell_archive_file_4))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
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
	} else {
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v14
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[0]))
		v23 = F_replace_percent_placeholders(m, v18, int32(_a_F_shell_archive_file_0), int32(_a_F_shell_archive_file_1), v8+int32(96))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v27 = F_errstart(m, int32(12), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v23
					F_errmsg_internal(m, int32(_a_F_shell_archive_file_2), v8+int32(80))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(77), int32(_a_F_shell_archive_file_4))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = F_fflush(m, int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
								v47 = F_pgl_system(m, v23)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
									if v47 != 0 {
										v63 = int32(255)
										if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
											v75 = int32(22)
										} else {
											v75 = int32(15)
										}
										v77 = v47 & int32(127)
										if v77 == int32(0) {
											v81 = F_errstart(m, v75, int32(0))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												if v81 == int32(0) {
													F_pfree(m, v23)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(112)
														return base.B2i32(v47 == int32(0))
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
													F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														v145 = int32(102)
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
														v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v23)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(112)
																	return base.B2i32(v47 == int32(0))
																}
															}
														}
													}
												}
											}
										} else {
											v97 = F_errstart(m, v75, int32(0))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
													if v97 == int32(0) {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														v109 = int32(_a_F_shell_archive_file_8)
														if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
															v114 = v77
															v115 = v109
															for {
																v118 = v115 + int32(1)
																v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
																if v119 != 0 {
																	v115 = v118
																	continue
																} else {
																}
																v121 = v114 - int32(1)
																if v121 != 0 {
																	v114 = v121
																	v115 = v118
																	continue
																} else {
																	break
																}
																break
															}
															v123 = v118
														} else {
															v123 = v109
														}
														if v123 != 0 {
															v126 = v123
														} else {
															v126 = int32(_a_F_shell_archive_file_9)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
														*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
														F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															v145 = int32(118)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(112)
																		return base.B2i32(v47 == int32(0))
																	}
																}
															}
														}
													}
												} else {
													if v97 == int32(0) {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
														F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															v145 = int32(127)
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
															v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v23)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v8 + int32(112)
																		return base.B2i32(v47 == int32(0))
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										F_pfree(m, v23)
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											v163 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												if v163 == int32(0) {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
													F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(135), int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
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
					v41 = F_fflush(m, int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(134217730)
						v47 = F_pgl_system(m, v23)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_file[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
							if v47 != 0 {
								v63 = int32(255)
								if base.B2i32(v47&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v47)>>(uint(int32(8))%32))&v63))|base.B2i32(base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) < base.Ui32(v63)) != 0 {
									v75 = int32(22)
								} else {
									v75 = int32(15)
								}
								v77 = v47 & int32(127)
								if v77 == int32(0) {
									v81 = F_errstart(m, v75, int32(0))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										if v81 == int32(0) {
											F_pfree(m, v23)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(112)
												return base.B2i32(v47 == int32(0))
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(base.Ui32(v47)>>(uint(int32(8))%32)) & int32(255)
											F_errmsg(m, int32(_a_F_shell_archive_file_6), v8+int32(32))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												v145 = int32(102)
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
												v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v23)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(112)
															return base.B2i32(v47 == int32(0))
														}
													}
												}
											}
										}
									}
								} else {
									v97 = F_errstart(m, v75, int32(0))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										if base.Ui32(v47&int32(_a_F_shell_archive_file_5)-int32(1)) <= base.Ui32(int32(254)) {
											if v97 == int32(0) {
												F_pfree(m, v23)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												v109 = int32(_a_F_shell_archive_file_8)
												if base.Ui32(int32(-64)) <= base.Ui32(v77-int32(65)) {
													v114 = v77
													v115 = v109
													for {
														v118 = v115 + int32(1)
														v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
														if v119 != 0 {
															v115 = v118
															continue
														} else {
														}
														v121 = v114 - int32(1)
														if v121 != 0 {
															v114 = v121
															v115 = v118
															continue
														} else {
															break
														}
														break
													}
													v123 = v118
												} else {
													v123 = v109
												}
												if v123 != 0 {
													v126 = v123
												} else {
													v126 = int32(_a_F_shell_archive_file_9)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v126
												*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v77
												F_errmsg(m, int32(_a_F_shell_archive_file_10), v8+int32(48))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													v145 = int32(118)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														}
													}
												}
											}
										} else {
											if v97 == int32(0) {
												F_pfree(m, v23)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v47
												F_errmsg(m, int32(_a_F_shell_archive_file_11), v8-int32(-64))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													v145 = int32(127)
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
													v150 = F_errdetail(m, int32(_a_F_shell_archive_file_7), v8+int32(16))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_shell_archive_file_3), v145, int32(_a_F_shell_archive_file_4))
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																m.G0 = v8 + int32(112)
																return base.B2i32(v47 == int32(0))
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_pfree(m, v23)
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									v163 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										if v163 == int32(0) {
											m.G0 = v8 + int32(112)
											return base.B2i32(v47 == int32(0))
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
											F_errmsg_internal(m, int32(_a_F_shell_archive_file_12), v8)
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_shell_archive_file_3), int32(135), int32(_a_F_shell_archive_file_4))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													m.G0 = v8 + int32(112)
													return base.B2i32(v47 == int32(0))
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
func F_shell_in(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errmsg(m, int32(_a_F_shell_in_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_shell_in_1), int32(307), int32(_a_F_shell_in_2))
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_shift_jis_2004_to_utf8(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14231(m, l0, int32(41), int32(0), int32(25), int32(_a_F_shift_jis_2004_to_utf8_0), int32(_a_F_shift_jis_2004_to_utf8_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_shimTriConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
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
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v398
L2:
	;
	v375 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v375)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v379 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v380 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v382 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v383 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)))
	v387 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+64)))
	v388 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+68)))
	v389 = F_FunctionCall8Coll(m, v377, v378, v379, v380, v381, v382, v383, base.I64_extend_i32_u(l0+int32(95)), v387, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L26
	} else {
		goto L60
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = v2
	v23 = v2
	goto L4
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v20))))
	if v34 == int32(2) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v46 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L6:
	;
	if int32(3) < v23 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v46 = v23
	goto L8
L8:
	;
	v48 = v22 + int32(1)
	if v48 != v17 {
		v22 = v48
		v23 = v46
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v398 = int32(2)
	goto L1
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v23<<(uint(int32(2))%32)))) = v22
	v46 = v23 + int32(1)
	goto L8
L12:
	;
	goto L5
L13:
	;
	v53 = base.B2i32(v46 <= int32(0))
	if v46 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v154 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v154)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v158 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v159 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v161 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v162 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)))
	v165 = base.I64_extend_i32_u(l0 + int32(95))
	v166 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+64)))
	v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+68)))
	v168 = F_FunctionCall8Coll(m, v156, v157, v158, v159, v160, v161, v162, v165, v166, v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v55 = v46 & int32(3)
	v56 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v46) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = v56
	v66 = int32(0)
	goto L19
L17:
	;
	v105 = v56
	goto L18
L18:
	;
	v118 = v105
	v120 = int32(0)
	goto L23
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v77 = v15 + v63<<(uint(int32(2))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v78))) = uint8(v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v82+v83))) = uint8(v80)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v88))) = uint8(v80)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v80)
	v97 = int32(4)
	v98 = v63 + v97
	v100 = v66 + v97
	if v100 != v46&int32(2147483644) {
		v63 = v98
		v66 = v100
		goto L19
	} else {
		goto L21
	}
L20:
	;
	if v55 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v105 = v98
	goto L18
L23:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v15+v118<<(uint(int32(2))%32))))
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v129+v133))) = uint8(v135)
	v137 = int32(1)
	v140 = v120 + v137
	if v140 != v55 {
		v118 = v118 + v137
		v120 = v140
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L14
L25:
	;
	goto L24
L26:
	;
	return int32(0)
L27:
	;
	v173 = base.B2i32(v168 != int64(0))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	v179 = v174
	goto L29
L28:
	;
	if v258&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L29:
	;
	v187 = int32(0)
	if v53 == v187 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v257 = int32(2)
	v258 = v247
	goto L28
L31:
	;
	v191 = v187
	goto L34
L32:
	;
	goto L33
L33:
	;
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v229)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v233 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v234 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v236 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v237 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)))
	v238 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+64)))
	v239 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+68)))
	v240 = F_FunctionCall8Coll(m, v231, v232, v233, v234, v235, v236, v237, v165, v238, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L26
	} else {
		goto L41
	}
L34:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v15+v191<<(uint(int32(2))%32))))
	v207 = v202 + v206
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v208 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v214 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v214)
	if v191 == v46 {
		v257 = v173
		v258 = v179
		goto L28
	} else {
		goto L40
	}
L36:
	;
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v209)
	v212 = v191 + int32(1)
	if v212 != v46 {
		v191 = v212
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	v257 = v173
	v258 = v179
	goto L28
L40:
	;
	goto L33
L41:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	v247 = base.B2i32(v242|v179&int32(1) != int32(0))
	v248 = int64(0)
	if base.B2i32(v168 != v248) == base.B2i32(v240 != v248) {
		v179 = v247
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	v269 = int32(2)
	goto L45
L44:
	;
	v269 = v257
	goto L45
L45:
	;
	if v257 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v272 = v269
	goto L48
L47:
	;
	v272 = v257
	goto L48
L48:
	;
	if v46 <= int32(0) {
		v398 = v272
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v276 = v46 & int32(3)
	v277 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v46) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v285 = v277
	v288 = int32(0)
	goto L53
L51:
	;
	v327 = v277
	goto L52
L52:
	;
	v339 = v327
	v341 = v277
	goto L57
L53:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v297 = int32(2)
	v299 = v15 + v285<<(uint(v297)%32)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	*(*uint8)(unsafe.Add(mBase, uint32(v296+v300))) = uint8(v297)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v304+v305))) = uint8(v297)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v299)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v309+v310))) = uint8(v297)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v314+v315))) = uint8(v297)
	v319 = int32(4)
	v320 = v285 + v319
	v322 = v288 + v319
	if v322 != v46&int32(2147483644) {
		v285 = v320
		v288 = v322
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if v276 == int32(0) {
		v398 = v272
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v327 = v320
	goto L52
L57:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v351 = int32(2)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v15+v339<<(uint(v351)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v350+v354))) = uint8(v351)
	v358 = int32(1)
	v361 = v341 + v358
	if v361 != v276 {
		v339 = v339 + v358
		v341 = v361
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v398 = v272
	goto L1
L59:
	;
	goto L58
L60:
	;
	v398 = base.B2i32(v389 != int64(0))
	goto L1
}
func F_shim_pclose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = int32(-1)
	if l0 == int32(0) {
		v32 = v3
		return v32
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_shim_pclose[0]))
		if l0 != v7 {
			v32 = v3
			return v32
		} else {
			v9 = F_fclose(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_shim_pclose[0])) = int32(0)
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_shim_pclose[1])))
				if v17 != 0 {
					v19 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_shim_pclose[1])) = uint8(v19)
					v24 = m.Env.Pgmem_run(m, int32(_a_F_shim_pclose_0), int32(_a_F_shim_pclose_1), int32(_a_F_shim_pclose_2))
					mBase = m.M
					return v24 << (uint(int32(8)) % 32) & int32(_a_F_shim_pclose_3)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_shim_pclose[2]))
					v32 = v31
					return v32
				}
			}
		}
	}
}
func F_shim_popen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	v3 = int32(_a_F_shim_popen_0)
	v4 = int32(_a_F_shim_popen_1)
	if (l0^v3)&int32(3) != 0 {
		v75 = l0
		v76 = v4
		v77 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[0])) = uint8(v115)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v118 == int32(119) {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	F___memset(m, v110, int32(0), v109)
	mBase = m.M
	goto L1
L3:
	;
	v109 = int32(0)
	v110 = v104
	goto L2
L4:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L22
L5:
	;
	if v76 == int32(0) {
		v104 = v77
		goto L3
	} else {
		goto L21
	}
L6:
	;
	if base.B2i32(l0&int32(3) == int32(0))|int32(0) != 0 {
		v41 = l0
		v42 = v4
		v43 = v3
		v44 = int32(1)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v44 == int32(0) {
		v104 = v43
		goto L3
	} else {
		goto L14
	}
L8:
	;
	v20 = l0
	v21 = v4
	v22 = v3
	goto L9
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v24)
	if v24 == int32(0) {
		v109 = v21
		v110 = v22
		goto L2
	} else {
		goto L11
	}
L10:
	;
	v41 = v35
	v42 = v31
	v43 = v29
	v44 = v33
	goto L7
L11:
	;
	v28 = int32(1)
	v29 = v22 + v28
	v31 = v21 - v28
	v32 = int32(0)
	v33 = base.B2i32(v31 != v32)
	v35 = v20 + v28
	if v35&int32(3) == v32 {
		v41 = v35
		v42 = v31
		v43 = v29
		v44 = v33
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v31 != 0 {
		v20 = v35
		v21 = v31
		v22 = v29
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v47 == int32(0) {
		v109 = v42
		v110 = v43
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v42) < base.Ui32(int32(4)) {
		v75 = v41
		v76 = v42
		v77 = v43
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v53 = v41
	v54 = v42
	v55 = v43
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 != v61 {
		v82 = v53
		v83 = v54
		v84 = v55
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v75 = v69
	v76 = v71
	v77 = v67
	goto L5
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v58
	v66 = int32(4)
	v67 = v55 + v66
	v69 = v53 + v66
	v71 = v54 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v53 = v69
		v54 = v71
		v55 = v67
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L4
L22:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v91)
	if v91 == int32(0) {
		v109 = v88
		v110 = v89
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v104 = v96
	goto L3
L24:
	;
	v95 = int32(1)
	v96 = v89 + v95
	v100 = v88 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v100
		v89 = v96
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_shim_popen[1])) = v142
	return v142
L27:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[2])) = uint8(v122)
	v126 = F_fopen(m, int32(_a_F_shim_popen_2), int32(_a_F_shim_popen_3))
	mBase = m.M
	v142 = v126
	goto L26
L28:
	;
	goto L29
L29:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_shim_popen[2])) = uint8(v128)
	v132 = int32(_a_F_shim_popen_4)
	v133 = m.Env.Pgmem_run(m, l0, int32(_a_F_shim_popen_5), v132)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_shim_popen[3])) = v133 << (uint(int32(8)) % 32) & int32(_a_F_shim_popen_6)
	v141 = F_fopen(m, v132, int32(_a_F_shim_popen_7))
	mBase = m.M
	v142 = v141
	goto L26
}
func F_should_apply_changes_for_rel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v11 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_should_apply_changes_for_rel_0), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_apply_changes_for_rel_1), int32(720), int32(_a_F_should_apply_changes_for_rel_2))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v60 = base.B2i32(v56 == v57)
		m.G0 = v7 + int32(16)
		return v60
	case 2:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_should_apply_changes_for_rel_3), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_apply_changes_for_rel_1), int32(715), int32(_a_F_should_apply_changes_for_rel_2))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		switch v20 - int32(114) {
		case 0:
			v60 = int32(1)
		case 1:
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
			v25 = *(*int64)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[1]))
			v60 = base.B2i32(base.Ui64(v23) <= base.Ui64(v25))
		default:
			v60 = int32(0)
		}
		m.G0 = v7 + int32(16)
		return v60
	case 4:
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v12 != int32(114) {
			v16 = v12
		} else {
			v16 = int32(0)
		}
		if v16 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_should_apply_changes_for_rel[2]))
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v74
					F_errmsg(m, int32(_a_F_should_apply_changes_for_rel_4), v7)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v81 = F_errdetail(m, int32(_a_F_should_apply_changes_for_rel_5), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_should_apply_changes_for_rel_1), int32(704), int32(_a_F_should_apply_changes_for_rel_2))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
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
			v60 = base.B2i32(v12 == int32(114))
			m.G0 = v7 + int32(16)
			return v60
		}
	default:
		v60 = int32(0)
		m.G0 = v7 + int32(16)
		return v60
	}
}
func F_should_refetch_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l0 {
	case 0:
		v85 = int32(0)
		m.G0 = v7 + int32(16)
		return v85
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_should_refetch_tuple_0), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_refetch_tuple_1), int32(165), int32(_a_F_should_refetch_tuple_2))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_should_refetch_tuple_3), v7)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_should_refetch_tuple_1), int32(168), int32(_a_F_should_refetch_tuple_2))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9 != int32(_a_F_should_refetch_tuple_4) {
			v29 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v29 == int32(0) {
					v85 = int32(1)
					m.G0 = v7 + int32(16)
					return v85
				} else {
					v69 = int32(_a_F_should_refetch_tuple_5)
					v70 = int32(154)
					F_errcode(m, int32(16777220))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, v69, int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v85 = int32(1)
								m.G0 = v7 + int32(16)
								return v85
							}
						}
					}
				}
			}
		} else {
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			if v12&v13 != int32(_a_F_should_refetch_tuple_6) {
				v29 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v29 == int32(0) {
						v85 = int32(1)
						m.G0 = v7 + int32(16)
						return v85
					} else {
						v69 = int32(_a_F_should_refetch_tuple_5)
						v70 = int32(154)
						F_errcode(m, int32(16777220))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, v69, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v85 = int32(1)
									m.G0 = v7 + int32(16)
									return v85
								}
							}
						}
					}
				}
			} else {
				v19 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v85 = int32(1)
						m.G0 = v7 + int32(16)
						return v85
					} else {
						v69 = int32(_a_F_should_refetch_tuple_7)
						v70 = int32(150)
						F_errcode(m, int32(16777220))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, v69, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v85 = int32(1)
									m.G0 = v7 + int32(16)
									return v85
								}
							}
						}
					}
				}
			}
		}
	case 4:
		v37 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			if v37 == int32(0) {
				v85 = int32(1)
				m.G0 = v7 + int32(16)
				return v85
			} else {
				v69 = int32(_a_F_should_refetch_tuple_8)
				v70 = int32(161)
				F_errcode(m, int32(16777220))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, v69, int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_should_refetch_tuple_1), v70, int32(_a_F_should_refetch_tuple_2))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v85 = int32(1)
							m.G0 = v7 + int32(16)
							return v85
						}
					}
				}
			}
		}
	}
}
func F_show_archive_command(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_show_archive_command[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_show_archive_command[1]))
	if v5 <= int32(0) {
		v8 = int32(_a_F_show_archive_command_0)
	} else {
		v8 = v3
	}
	return v8
}
func F_sigaddset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = l1 - int32(1)
	if base.B2i32(base.Ui32(v5) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1-int32(32))) == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_sigaddset[0])) = int32(28)
		return
	} else {
		v22 = l0 + int32(base.Ui32(v5)>>(uint(int32(3))%32))&int32(536870908)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 | int32(1)<<(uint(v5)%32)
		return
	}
}
func F_sigdelset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = l1 - int32(1)
	if base.B2i32(base.Ui32(v5) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1-int32(32))) == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_sigdelset[0])) = int32(28)
		return
	} else {
		v22 = l0 + int32(base.Ui32(v5)>>(uint(int32(3))%32))&int32(536870908)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 & base.I32_rotl(int32(-2), v5)
		return
	}
}
func F_signValue(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v9 = int32(_a_F_signValue_0)
	v11 = base.I32_rem_s(l3, int32(2147483646))
	*(*int32)(unsafe.Add(mBase, _c_F_signValue[0])) = v11 + int32(1)
	v21 = l0 + l3<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+896))
	v23 = F_FunctionCall1Coll(m, l0+l3*int32(28), v22, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_signValue[0]))
		v30 = int32(_a_F_signValue_1)
		v31 = base.I32_div_s(v29, v30)
		v39 = v31*int32(-2836) + (v29-v31*v30)*int32(_a_F_signValue_2)
		if int32(0) <= v39 {
			v42 = int32(-1)
		} else {
			v42 = int32(2147483646)
		}
		v46 = base.I32_rem_s(base.I32_wrap_i64(v23)^(v42+v39), int32(2147483646))
		v48 = v46 + int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_signValue[0])) = v48
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(1032))))
		if int32(0) < v52 {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1028))
			v59 = int32(0)
			v62 = v48
			for {
				v67 = int32(_a_F_signValue_1)
				v68 = base.I32_div_s(v62, v67)
				v76 = v68*int32(-2836) + (v62-v68*v67)*int32(_a_F_signValue_2)
				if v76 < int32(0) {
					v81 = v76 + int32(2147483647)
				} else {
					v81 = v76
				}
				v82 = int32(1)
				v84 = base.I32_rem_s(v81-v82, v55<<(uint(int32(4))%32))
				v86 = base.I32_div_s(v84, int32(16))
				v89 = l1 + v86<<(uint(v82)%32)
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
				v95 = v90 | v82<<(uint(v84&int32(15))%32)
				*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v95)
				v98 = v59 + v82
				if v98 != v52 {
					v59 = v98
					v62 = v81
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, _c_F_signValue[0])) = v81
		} else {
		}
		return
	}
}
func F_similar_to_escape_2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = F_similar_escape_internal(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v10)
			}
		}
	}
}
func F_similarity_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_DirectFunctionCall2Coll(m, int32(_a_F_similarity_dist_0), int32(0), v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_sub(float32(1), base.F32_reinterpret_i32(base.I32_wrap_i64(v7)))))
	}
}
func F_sin(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v24 float64
	_ = v24
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v73 float64
	_ = v73
	var v89 float64
	_ = v89
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v127 float64
	_ = v127
	var v147 float64
	_ = v147
	var v163 float64
	_ = v163
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v202 float64
	_ = v202
	var v219 float64
	_ = v219
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v14) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v14) < base.Ui32(int32(1045430272)) {
			v219 = l0
		} else {
			v24 = base.F64_mul(l0, l0)
			v219 = base.F64_add(base.F64_mul(base.F64_mul(l0, v24), base.F64_add(base.F64_mul(v24, base.F64_add(base.F64_mul(base.F64_mul(v24, base.F64_mul(v24, v24)), base.F64_add(base.F64_mul(v24, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v24, base.F64_add(base.F64_mul(v24, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))), float64(-0.16666666666666632))), l0)
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v14) {
			v219 = base.F64_sub(l0, l0)
		} else {
			v62 = F___rem_pio2(m, l0, v7)
			mBase = m.M
			v63 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v64 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			switch v62&int32(3) - int32(1) {
			case 0:
				v111 = float64(1)
				v112 = base.F64_mul(v64, v64)
				v114 = base.F64_mul(v112, float64(0.5))
				v115 = base.F64_sub(v111, v114)
				v127 = base.F64_mul(v112, v112)
				v219 = base.F64_add(v115, base.F64_add(base.F64_sub(base.F64_sub(v111, v115), v114), base.F64_sub(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v127, v127), base.F64_add(base.F64_mul(v112, base.F64_add(base.F64_mul(v112, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v64, v63))))
			case 1:
				v147 = base.F64_mul(v64, v64)
				v163 = base.F64_mul(v64, v147)
				v219 = base.F64_neg(base.F64_sub(v64, base.F64_add(base.F64_sub(base.F64_mul(v147, base.F64_sub(base.F64_mul(v63, float64(0.5)), base.F64_mul(v163, base.F64_add(base.F64_mul(base.F64_mul(v147, base.F64_mul(v147, v147)), base.F64_add(base.F64_mul(v147, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v147, base.F64_add(base.F64_mul(v147, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v63), base.F64_mul(v163, float64(0.16666666666666632)))))
			case 2:
				v186 = float64(1)
				v187 = base.F64_mul(v64, v64)
				v189 = base.F64_mul(v187, float64(0.5))
				v190 = base.F64_sub(v186, v189)
				v202 = base.F64_mul(v187, v187)
				v219 = base.F64_neg(base.F64_add(v190, base.F64_add(base.F64_sub(base.F64_sub(v186, v190), v189), base.F64_sub(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, float64(2.480158728947673e-05)), float64(-0.001388888888887411))), float64(0.0416666666666666))), base.F64_mul(base.F64_mul(v202, v202), base.F64_add(base.F64_mul(v187, base.F64_add(base.F64_mul(v187, float64(-1.1359647557788195e-11)), float64(2.087572321298175e-09))), float64(-2.7557314351390663e-07))))), base.F64_mul(v64, v63)))))
			default:
				v73 = base.F64_mul(v64, v64)
				v89 = base.F64_mul(v64, v73)
				v219 = base.F64_sub(v64, base.F64_add(base.F64_sub(base.F64_mul(v73, base.F64_sub(base.F64_mul(v63, float64(0.5)), base.F64_mul(v89, base.F64_add(base.F64_mul(base.F64_mul(v73, base.F64_mul(v73, v73)), base.F64_add(base.F64_mul(v73, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))))), v63), base.F64_mul(v89, float64(0.16666666666666632))))
			}
		}
	}
	m.G0 = v7 + int32(16)
	return v219
}
func F_single_bound_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(8)
	v8 = F_range_cmp_bounds(m, l2, l0+v4, l1+v4)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_skip_b_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	if l3 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = l1
	v13 = l3
	goto L7
L5:
	;
	v48 = l1
	goto L6
L6:
	;
	return v48
L7:
	;
	if v11 <= l2 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v48 = v39
	goto L6
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v19 = v11 - int32(1)
	v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v19))))
	if base.B2i32(int32(0) <= v21)|base.B2i32(v19 <= l2) != 0 {
		v39 = v19
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(1)
	if v43 < v13 {
		v11 = v39
		v13 = v13 - v43
		goto L7
	} else {
		goto L18
	}
L13:
	;
	v27 = v19
	goto L14
L14:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v27))))
	if base.Ui32(int32(191)) < base.Ui32(v32) {
		v39 = v27
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v39 = l2
	goto L12
L16:
	;
	v36 = v27 - int32(1)
	if l2 < v36 {
		v27 = v36
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L8
}
func F_slice_from_v(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
	v6 = F_slice_from_s(m, l0, v5, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_slotsync_failure_callback(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[0]))
	if v4 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(1))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])))
				if v11 != 0 {
					v13 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					v16 = base.AtomicRmwXchg32(m, v13, int32(16), int32(1))
					if v16 != 0 {
						F_s_lock(m, v13+int32(16), int32(_a_F_slotsync_failure_callback_0))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v22 = int32(_a_F_slotsync_failure_callback_1)
							v23 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(-1)
							v26 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)) = uint8(v26)
							v29 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+16)), uint32(v26))
							*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v26)
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
							m.T0[v40].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v22 = int32(_a_F_slotsync_failure_callback_1)
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(-1)
						v26 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)) = uint8(v26)
						v29 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+16)), uint32(v26))
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v26)
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
						m.T0[v40].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
					m.T0[v40].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_ReplicationSlotCleanup(m, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])))
			if v11 != 0 {
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
				v16 = base.AtomicRmwXchg32(m, v13, int32(16), int32(1))
				if v16 != 0 {
					F_s_lock(m, v13+int32(16), int32(_a_F_slotsync_failure_callback_0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = int32(_a_F_slotsync_failure_callback_1)
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(-1)
						v26 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)) = uint8(v26)
						v29 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+16)), uint32(v26))
						*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v26)
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
						m.T0[v40].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v22 = int32(_a_F_slotsync_failure_callback_1)
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(-1)
					v26 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v23)+5)) = uint8(v26)
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[2]))
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+16)), uint32(v26))
					*(*uint8)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[1])) = uint8(v26)
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
					m.T0[v40].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_slotsync_failure_callback[3]))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
				m.T0[v40].(func(*base.Module, int32))(m, base.I32_wrap_i64(l1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_smgrdounlinkall(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
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
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v184 int64
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int64
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int64
	_ = v376
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int64
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int64
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int64
	_ = v498
	var v522 int32
	_ = v522
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int64
	_ = v603
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v631 int64
	_ = v631
	var v633 int64
	_ = v633
	var v653 int64
	_ = v653
	var v662 int64
	_ = v662
	var v687 int32
	_ = v687
	var v688 int64
	_ = v688
	var v691 int64
	_ = v691
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v735 int64
	_ = v735
	var v737 int64
	_ = v737
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v771 int64
	_ = v771
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v810 int32
	_ = v810
	var v828 int32
	_ = v828
	var v846 int32
	_ = v846
	var v857 int32
	_ = v857
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int64
	_ = v891
	var v894 int32
	_ = v894
	var v895 int64
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v936 int32
	_ = v936
	var v937 int64
	_ = v937
	var v939 int64
	_ = v939
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int64
	_ = v968
	var v970 int64
	_ = v970
	var v976 int32
	_ = v976
	var v977 int64
	_ = v977
	var v979 int64
	_ = v979
	var v985 int32
	_ = v985
	var v986 int64
	_ = v986
	var v988 int64
	_ = v988
	var v994 int32
	_ = v994
	var v995 int64
	_ = v995
	var v997 int64
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1039 int32
	_ = v1039
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(_a_F_smgrdounlinkall_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0])) = v23 + int32(1)
	v27 = m.G0
	v29 = v27 - int32(112)
	m.G0 = v29
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v1039 = v19
	goto L3
L3:
	;
	m.G0 = v1039 + int32(80)
	return
L4:
	;
	v32 = F_palloc_mul(m, int32(4), l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v857 = v19
	goto L6
L6:
	;
	m.G0 = v29 + int32(112)
	v867 = F_palloc_mul(m, int32(16), l1)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L7
	} else {
		goto L177
	}
L7:
	;
	return
L8:
	;
	if l1 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v32)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L7
	} else {
		goto L176
	}
L10:
	;
	v40 = v4
	v45 = v4
	goto L11
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0+v40<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v56 != int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v150 == int32(0) {
		goto L9
	} else {
		goto L31
	}
L13:
	;
	v158 = v40 + int32(1)
	if v158 != l1 {
		v40 = v158
		v45 = v150
		goto L11
	} else {
		goto L30
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[1]))
	if v56 != v60 {
		v150 = v45
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+v45<<(uint(int32(2))%32)))) = v55
	v150 = v45 + int32(1)
	goto L13
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+72)) = v64
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[2]))
	if v66 < v68 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v72 = v29 + int32(72)
	v76 = v66
	goto L21
L19:
	;
	goto L20
L20:
	;
	v150 = v45
	goto L13
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[3]))
	v93 = v90 + v76*int32(56)
	v94 = int64(0)
	v97 = base.AtomicRmwCmpxchg64(m, v93, int32(24), v94, v94)
	if v97&int64(33554432) == v94 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v115 = v76 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[2]))
	if v115 < v117 {
		v76 = v115
		goto L21
	} else {
		goto L29
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v102 != v103 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v105 != v106 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v108 != v109 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_InvalidateLocalBuffer(m, v93, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L23
L29:
	;
	goto L22
L30:
	;
	goto L12
L31:
	;
	v164 = F_palloc(m, v150<<(uint(int32(4))%32))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	if int32(0) < v150 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_pfree(m, v164)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L7
	} else {
		goto L175
	}
L34:
	;
	F_pfree(m, v164)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L7
	} else {
		goto L110
	}
L35:
	;
	v172 = int32(0)
	v184 = int64(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if base.Ui32(int32(63)) <= base.Ui32(v389+int32(31)) {
		goto L33
	} else {
		goto L109
	}
L38:
	;
	v187 = v164 + v172<<(uint(int32(4))%32)
	v190 = v32 + v172<<(uint(int32(2))%32)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v194 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	v300 = base.I32_div_s(v298, int32(32))
	if base.Ui64(base.I64_extend_i32_s(v300)) <= base.Ui64(v292) {
		goto L34
	} else {
		goto L89
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v205
	if v205 == int32(-1) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v191+int32(0))+20))
	if v200 != int32(-1) {
		v205 = v200
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v205 = int32(-1)
	goto L41
L45:
	;
	goto L44
L46:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v221 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v211 = F_smgrexists(m, v209, int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v217 = v184 + base.I64_extend_i32_u(v205)
	goto L46
L50:
	;
	if v211 == int32(0) {
		v217 = v184
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L34
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v232
	if v232 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v218+int32(4))+20))
	if v227 != int32(-1) {
		v232 = v227
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v232 = int32(-1)
	goto L53
L57:
	;
	goto L56
L58:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v246 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v242 = v217 + base.I64_extend_i32_u(v232)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v240 = F_smgrexists(m, v238, int32(1))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	if v240 != 0 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	v242 = v217
	goto L58
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = v257
	if v257 != int32(-1) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	goto L64
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v243+int32(8))+20))
	if v252 != int32(-1) {
		v257 = v252
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v257 = int32(-1)
	goto L65
L69:
	;
	goto L68
L70:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_smgrdounlinkall[5])))
	if v271 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v267 = v242 + base.I64_extend_i32_u(v257)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v265 = F_smgrexists(m, v263, int32(2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	if v265 != 0 {
		goto L34
	} else {
		goto L75
	}
L75:
	;
	v267 = v242
	goto L70
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+12)) = v282
	if v282 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v268+int32(12))+20))
	if v277 != int32(-1) {
		v282 = v277
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v282 = int32(-1)
	goto L77
L81:
	;
	goto L80
L82:
	;
	v294 = v172 + int32(1)
	if v294 < v150 {
		v172 = v294
		v184 = v292
		goto L38
	} else {
		goto L88
	}
L83:
	;
	v292 = v267 + base.I64_extend_i32_u(v282)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v290 = F_smgrexists(m, v288, int32(3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	if v290 != 0 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	v292 = v267
	goto L82
L88:
	;
	goto L39
L89:
	;
	v307 = int32(0)
	goto L90
L90:
	;
	v321 = v32 + v307<<(uint(int32(2))%32)
	v324 = v164 + v307<<(uint(int32(4))%32)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	if v325 != int32(-1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L33
L92:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v328)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = v331
	v335 = int32(0)
	F_FindAndDropRelationBuffers(m, v29+int32(56), v335, v325, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L7
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v340 != int32(-1) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v344
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v346
	F_FindAndDropRelationBuffers(m, v29+int32(40), int32(1), v340, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	if v355 != int32(-1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v359
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v358)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v361
	F_FindAndDropRelationBuffers(m, v29+int32(24), int32(2), v355, int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	if v370 != int32(-1) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v374
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v373)))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v376
	F_FindAndDropRelationBuffers(m, v29+int32(8), int32(3), v370, int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v386 = v307 + int32(1)
	if v386 != v150 {
		v307 = v386
		goto L90
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	goto L91
L109:
	;
	goto L34
L110:
	;
	v413 = F_palloc_mul(m, int32(12), v150)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	v415 = int32(0)
	if v150 <= v415 {
		v540 = v415
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if int32(0) < v542 {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	v418 = int32(0)
	if v150 != int32(1) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if base.Ui32(v150) < base.Ui32(int32(21)) {
		v540 = int32(0)
		goto L112
	} else {
		goto L122
	}
L115:
	;
	v429 = int32(0)
	v430 = v418
	goto L118
L116:
	;
	v477 = v418
	goto L117
L117:
	;
	v491 = v413 + v477*int32(12)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v32+v477<<(uint(int32(2))%32))))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v491)+8)) = v496
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v495)))
	*(*int64)(unsafe.Add(mBase, uint32(v491))) = v498
	goto L114
L118:
	;
	v442 = int32(12)
	v444 = v413 + v430*v442
	v445 = int32(2)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v32+v430<<(uint(v445)%32))))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+8)) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v448)))
	*(*int64)(unsafe.Add(mBase, uint32(v444))) = v451
	v454 = v430 | int32(1)
	v457 = v413 + v454*v442
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v32+v454<<(uint(v445)%32))))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v457)+8)) = v462
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v461)))
	*(*int64)(unsafe.Add(mBase, uint32(v457))) = v464
	v467 = v430 + v445
	v469 = v429 + v445
	if v469 != v150&int32(2147483646) {
		v429 = v469
		v430 = v467
		goto L118
	} else {
		goto L120
	}
L119:
	;
	if v150&int32(1) == int32(0) {
		goto L114
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v477 = v467
	goto L117
L122:
	;
	F_pg_qsort(m, v413, v150, int32(12), int32(1165))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	v540 = int32(1)
	goto L112
L124:
	;
	v554 = int32(0)
	goto L127
L125:
	;
	goto L126
L126:
	;
	F_pfree(m, v413)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L7
	} else {
		goto L174
	}
L127:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[6]))
	v566 = v563 + v554*int32(56)
	if v540 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	goto L126
L129:
	;
	v789 = v554 + int32(1)
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[4]))
	if v789 < v791 {
		v554 = v789
		goto L127
	} else {
		goto L173
	}
L130:
	;
	v631 = int64(4194304)
	v633 = base.AtomicRmwOr64(m, v566, int32(24), v631)
	if v633&v631 != int64(0) {
		goto L144
	} else {
		goto L145
	}
L131:
	;
	if v150 <= int32(0) {
		goto L129
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v566)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v604
	*(*int64)(unsafe.Add(mBase, uint32(v29)+88)) = v603
	v611 = F_bsearch(m, v29+int32(88), v413, v150, int32(12), int32(1165))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L7
	} else {
		goto L142
	}
L134:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v577 = int32(0)
	goto L135
L135:
	;
	v591 = v413 + v577*int32(12)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	if v571 != v592 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L129
L137:
	;
	v601 = v577 + int32(1)
	if v601 != v150 {
		v577 = v601
		goto L135
	} else {
		goto L141
	}
L138:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	if v594 != v595 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v591)+8))
	if v597 == v598 {
		v618 = v591
		goto L130
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	goto L136
L142:
	;
	if v611 == int32(0) {
		goto L129
	} else {
		goto L143
	}
L143:
	;
	v618 = v611
	goto L130
L144:
	;
	v653 = v633
	goto L147
L145:
	;
	goto L146
L146:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	if v758 != v759 {
		goto L168
	} else {
		goto L169
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+108)) = int32(_a_F_smgrdounlinkall_1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+104)) = int32(_a_F_smgrdounlinkall_2)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = int32(_a_F_smgrdounlinkall_3)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = int32(0)
	v662 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+88)) = v662
	if v653&int64(4194304) != v662 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	goto L146
L149:
	;
	goto L152
L150:
	;
	goto L151
L151:
	;
	v715 = int32(_a_F_smgrdounlinkall_4)
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[7]))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(88))+8))
	if v718 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L152:
	;
	F_perform_spin_delay(m, v29+int32(88))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L7
	} else {
		goto L154
	}
L153:
	;
	goto L151
L154:
	;
	v688 = int64(0)
	v691 = base.AtomicRmwCmpxchg64(m, v566, int32(24), v688, v688)
	if v691&int64(4194304) != v688 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v735 = int64(4194304)
	v737 = base.AtomicRmwOr64(m, v566, int32(24), v735)
	if v737&v735 != int64(0) {
		v653 = v737
		goto L147
	} else {
		goto L167
	}
L157:
	;
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[7])) = v733
	goto L157
L159:
	;
	if int32(999) < v716 {
		goto L157
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if v716 < int32(11) {
		goto L157
	} else {
		goto L166
	}
L162:
	;
	v723 = int32(900)
	if v723 <= v716 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v726 = v723
	goto L165
L164:
	;
	v726 = v716
	goto L165
L165:
	;
	v733 = v726 + int32(100)
	goto L158
L166:
	;
	v733 = v716 - int32(1)
	goto L158
L167:
	;
	goto L148
L168:
	;
	v771 = base.AtomicRmwSub64(m, v566, int32(24), int64(4194304))
	goto L129
L169:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	if v761 != v762 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v618)+8))
	if v764 != v765 {
		goto L168
	} else {
		goto L171
	}
L171:
	;
	F_InvalidateBuffer(m, v566)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L7
	} else {
		goto L172
	}
L172:
	;
	goto L129
L173:
	;
	goto L128
L174:
	;
	goto L9
L175:
	;
	goto L9
L176:
	;
	v857 = v19
	goto L6
L177:
	;
	if int32(0) < l1 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v884 = v4
	goto L181
L179:
	;
	goto L180
L180:
	;
	F_pfree(m, v867)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L7
	} else {
		goto L199
	}
L181:
	;
	v889 = l0 + v884<<(uint(int32(2))%32)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	v891 = *(*int64)(unsafe.Add(mBase, uint32(v890)))
	v894 = v867 + v884<<(uint(int32(4))%32)
	v895 = *(*int64)(unsafe.Add(mBase, uint32(v890)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v894)+8)) = v895
	*(*int64)(unsafe.Add(mBase, uint32(v894))) = v891
	v898 = int32(0)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	F_mdclose(m, v899, v898)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L7
	} else {
		goto L183
	}
L182:
	;
	v921 = v898
	goto L188
L183:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	F_mdclose(m, v903, int32(1))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L7
	} else {
		goto L184
	}
L184:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	F_mdclose(m, v907, int32(2))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	F_mdclose(m, v911, int32(3))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	v916 = v884 + int32(1)
	if v916 != l1 {
		v884 = v916
		goto L181
	} else {
		goto L187
	}
L187:
	;
	goto L182
L188:
	;
	v936 = v867 + v921<<(uint(int32(4))%32)
	v937 = *(*int64)(unsafe.Add(mBase, uint32(v936)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+72)) = v937
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v936)))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+64)) = v939
	F_CacheInvalidateSmgr(m, v857-int32(-64))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L7
	} else {
		goto L190
	}
L189:
	;
	v962 = int32(0)
	goto L192
L190:
	;
	v946 = v921 + int32(1)
	if v946 != l1 {
		v921 = v946
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v967 = v867 + v962<<(uint(int32(4))%32)
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v967)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+56)) = v968
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v967)))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+48)) = v970
	F_mdunlink(m, v857+int32(48), int32(0), l2)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L7
	} else {
		goto L194
	}
L193:
	;
	goto L180
L194:
	;
	v977 = *(*int64)(unsafe.Add(mBase, uint32(v967)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+40)) = v977
	v979 = *(*int64)(unsafe.Add(mBase, uint32(v967)))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+32)) = v979
	F_mdunlink(m, v857+int32(32), int32(1), l2)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L7
	} else {
		goto L195
	}
L195:
	;
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v967)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+24)) = v986
	v988 = *(*int64)(unsafe.Add(mBase, uint32(v967)))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+16)) = v988
	F_mdunlink(m, v857+int32(16), int32(2), l2)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L7
	} else {
		goto L196
	}
L196:
	;
	v995 = *(*int64)(unsafe.Add(mBase, uint32(v967)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v857)+8)) = v995
	v997 = *(*int64)(unsafe.Add(mBase, uint32(v967)))
	*(*int64)(unsafe.Add(mBase, uint32(v857))) = v997
	F_mdunlink(m, v857, int32(3), l2)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	v1003 = v962 + int32(1)
	if v1003 != l1 {
		v962 = v1003
		goto L192
	} else {
		goto L198
	}
L198:
	;
	goto L193
L199:
	;
	v1023 = int32(_a_F_smgrdounlinkall_0)
	v1025 = *(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrdounlinkall[0])) = v1025 - int32(1)
	v1039 = v857
	goto L3
}
func F_smgrpin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v3 != 0 {
		v11 = v3
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v5
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v11 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v11 + int32(1)
	return
}
func F_smgrwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	v17 = int32(_a_F_smgrwritev_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0])) = v19 + int32(1)
	v23 = m.G0
	v25 = v23 - int32(1040)
	m.G0 = v25
	v28 = F__mdfd_getseg(m, l0, l1, l2, l4, int32(9))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = int32(1)
	v34 = l2 & int32(_a_F_smgrwritev_1)
	v35 = int32(_a_F_smgrwritev_2) - v34
	if base.Ui32(v31) < base.Ui32(v35) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v269 = int32(_a_F_smgrwritev_0)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[0])) = v271 - int32(1)
	return
L4:
	;
	v38 = v31
	goto L6
L5:
	;
	v38 = v35
	goto L6
L6:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v38) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = int32(128)
	goto L9
L8:
	;
	v41 = v38
	goto L9
L9:
	;
	if v41 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = base.I64_extend_i32_u(v34 << (uint(int32(13)) % 32))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = int32(_a_F_smgrwritev_3)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v47
	v51 = int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v38) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L59
	}
L13:
	;
	v57 = v47
	v62 = v25 + int32(16)
	v64 = v51
	v68 = int32(1)
	v71 = int32(0)
	goto L16
L14:
	;
	v121 = v51
	goto L15
L15:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v134 = F_FileWriteV(m, v130, v25+int32(16), v121, v46, int32(167772186))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L28
	}
L16:
	;
	v75 = l3 + v68<<(uint(int32(2))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v76 == v57+v77 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v121 = v109
	goto L15
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v93 != v90+v94 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v77 - int32(-8192)
	v90 = v57
	v91 = v62
	v92 = v64
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = int32(_a_F_smgrwritev_3)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v76
	v90 = v76
	v91 = v62 + int32(8)
	v92 = v64 + int32(1)
	goto L18
L22:
	;
	v110 = int32(2)
	v113 = v71 + v110
	if v113 != 0 {
		v57 = v107
		v62 = v108
		v64 = v109
		v68 = v68 + v110
		v71 = v113
		goto L16
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = int32(_a_F_smgrwritev_3)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v93
	v107 = v93
	v108 = v91 + int32(8)
	v109 = v92 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v94 - int32(-8192)
	v107 = v90
	v108 = v91
	v109 = v92
	goto L22
L26:
	;
	goto L17
L27:
	;
	if l4 != 0 {
		goto L55
	} else {
		goto L56
	}
L28:
	;
	if int32(0) <= v134 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v139 = int32(0)
	v144 = v134
	v146 = v121
	v154 = v46
	goto L32
L30:
	;
	goto L31
L31:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[1]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L46
	}
L32:
	;
	v155 = v139 + v144
	if v155 == int32(_a_F_smgrwritev_3) {
		goto L27
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	v159 = v25 + int32(16)
	v162 = v159
	v163 = v146
	v164 = v144
	goto L37
L35:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v194 = v154 + base.I64_extend_i32_u(v144)
	v196 = F_FileWriteV(m, v192, v159, v191, v194, int32(167772186))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L44
	}
L36:
	;
	if v159 == v162 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if base.Ui32(v164) < base.Ui32(v166) {
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v191 = int32(0)
	goto L35
L39:
	;
	v172 = v163 - int32(1)
	if v172 != 0 {
		v162 = v162 + int32(8)
		v163 = v172
		v164 = v164 - v166
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v181 + v164
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v184 - v164
	v191 = v163
	goto L35
L42:
	;
	v176 = v163 << (uint(int32(3)) % 32)
	if v176 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	base.MemoryCopy(m, v159, v162, v176)
	goto L41
L44:
	;
	if int32(0) <= v196 {
		v139 = v155
		v144 = v196
		v146 = v191
		v154 = v194
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_smgrwritev[2]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226+v224*int32(48))+32))
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l2
	F_errmsg(m, int32(_a_F_smgrwritev_4), v25)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v217 == int32(51) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_errhint(m, int32(_a_F_smgrwritev_5), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_errfinish(m, int32(_a_F_smgrwritev_6), int32(1145), int32(_a_F_smgrwritev_7))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	m.G0 = v25 + int32(1040)
	goto L3
L56:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v248 != int32(-1) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	F_register_dirty_segment(m, l0, l1, v28)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	F_errmsg_internal(m, int32(_a_F_smgrwritev_8), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_smgrwritev_6), int32(1103), int32(_a_F_smgrwritev_7))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	v11 = F_vsnprintf(m, l0, l1, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v11
	}
}
func F_sort_asc(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14377(m, l0, int32(_a_F_sort_asc_0), int32(235), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_sort_checkpoint_bufferids(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int64
	_ = v553
	var v555 int64
	_ = v555
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v561 int64
	_ = v561
	var v563 int32
	_ = v563
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int64
	_ = v616
	var v618 int64
	_ = v618
	var v620 int32
	_ = v620
	var v622 int64
	_ = v622
	var v624 int64
	_ = v624
	var v626 int32
	_ = v626
	var v628 int64
	_ = v628
	var v630 int64
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int64
	_ = v685
	var v687 int64
	_ = v687
	var v689 int32
	_ = v689
	var v691 int64
	_ = v691
	var v693 int64
	_ = v693
	var v695 int32
	_ = v695
	var v697 int64
	_ = v697
	var v699 int64
	_ = v699
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int64
	_ = v752
	var v754 int64
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int64
	_ = v759
	var v761 int64
	_ = v761
	var v763 int32
	_ = v763
	var v765 int64
	_ = v765
	var v767 int64
	_ = v767
	var v770 int32
	_ = v770
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int64
	_ = v817
	var v819 int64
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int64
	_ = v824
	var v826 int64
	_ = v826
	var v828 int32
	_ = v828
	var v830 int64
	_ = v830
	var v832 int64
	_ = v832
	var v835 int32
	_ = v835
	var v854 int32
	_ = v854
	var v865 int32
	_ = v865
	var v867 int64
	_ = v867
	var v869 int64
	_ = v869
	var v871 int32
	_ = v871
	var v873 int64
	_ = v873
	var v875 int64
	_ = v875
	var v877 int32
	_ = v877
	var v879 int64
	_ = v879
	var v881 int64
	_ = v881
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v915 int32
	_ = v915
	var v925 int32
	_ = v925
	var v937 int32
	_ = v937
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v974 int64
	_ = v974
	var v976 int64
	_ = v976
	var v978 int32
	_ = v978
	var v980 int64
	_ = v980
	var v982 int64
	_ = v982
	var v984 int32
	_ = v984
	var v986 int64
	_ = v986
	var v988 int64
	_ = v988
	var v1005 int32
	_ = v1005
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v887 = l0
		v888 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return
L2:
	;
	if base.Ui32(v902) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L272
	}
L3:
	;
	v900 = v887
	v902 = v888
	goto L2
L4:
	;
	v20 = l0
	v21 = l1
	goto L5
L5:
	;
	v34 = v20 + int32(20)
	v36 = v21
	goto L7
L7:
	;
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v52 = v20 + v36*int32(20)
	v55 = v34
	goto L10
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(20))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if base.Ui32(v68) < base.Ui32(v69) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v100 = v20 + int32(base.Ui32(v36)>>(uint(int32(1))%32))*int32(20)
	if v36 != int32(7) {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	goto L11
L13:
	;
	v92 = v55 + int32(20)
	if base.Ui32(v92) < base.Ui32(v52) {
		v55 = v92
		goto L10
	} else {
		goto L21
	}
L14:
	;
	if base.Ui32(v69) < base.Ui32(v68) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(16))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if base.Ui32(v74) < base.Ui32(v75) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v75) < base.Ui32(v74) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(12))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v80 < v81 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v81 < v80 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v55-int32(8))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if base.Ui32(v87) < base.Ui32(v86) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	goto L1
L22:
	;
	v104 = v52 - int32(20)
	if base.Ui32(v36) < base.Ui32(int32(41)) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v546 = v100
	goto L24
L24:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v551
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v553
	v555 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v555
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v546)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v557
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v546)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v559
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v546)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v561
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v546)+16)) = v563
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v546)+8)) = v565
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v546))) = v567
	v570 = v52 - int32(20)
	v573 = v570
	v574 = v34
	v576 = v34
	v579 = v570
	goto L217
L25:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	if base.Ui32(v445) < base.Ui32(v446) {
		goto L175
	} else {
		goto L176
	}
L26:
	;
	v435 = v20
	v436 = v100
	v437 = v104
	goto L25
L27:
	;
	goto L28
L28:
	;
	v108 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
	v110 = v108 * int32(20)
	v111 = v20 + v110
	v114 = v20 + v108*int32(40)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if base.Ui32(v119) < base.Ui32(v120) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v220 = v108 * int32(-20)
	v221 = v100 + v220
	v222 = v100 + v110
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if base.Ui32(v227) < base.Ui32(v228) {
		goto L81
	} else {
		goto L82
	}
L30:
	;
	v218 = v20
	goto L29
L31:
	;
	v218 = v114
	goto L29
L32:
	;
	v218 = v198
	goto L29
L33:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if base.Ui32(v120) < base.Ui32(v168) {
		goto L59
	} else {
		goto L60
	}
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if base.Ui32(v120) < base.Ui32(v136) {
		v198 = v111
		goto L32
	} else {
		goto L42
	}
L35:
	;
	if base.Ui32(v120) < base.Ui32(v119) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if base.Ui32(v123) < base.Ui32(v124) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	if base.Ui32(v124) < base.Ui32(v123) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v127 < v128 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	if v128 < v127 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	if base.Ui32(v132) <= base.Ui32(v131) {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L34
L42:
	;
	if base.Ui32(v136) < base.Ui32(v120) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.Ui32(v119) < base.Ui32(v136) {
		goto L31
	} else {
		goto L50
	}
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v139) < base.Ui32(v140) {
		v198 = v111
		goto L32
	} else {
		goto L45
	}
L45:
	;
	if base.Ui32(v140) < base.Ui32(v139) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v143 < v144 {
		v198 = v111
		goto L32
	} else {
		goto L47
	}
L47:
	;
	if v144 < v143 {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if base.Ui32(v147) < base.Ui32(v148) {
		v198 = v111
		goto L32
	} else {
		goto L49
	}
L49:
	;
	goto L43
L50:
	;
	if base.Ui32(v136) < base.Ui32(v119) {
		goto L30
	} else {
		goto L51
	}
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v154) < base.Ui32(v155) {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(v155) < base.Ui32(v154) {
		goto L30
	} else {
		goto L53
	}
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v158 < v159 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	if v159 < v158 {
		goto L30
	} else {
		goto L55
	}
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if base.Ui32(v162) < base.Ui32(v163) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v165 = v114
	goto L58
L57:
	;
	v165 = v20
	goto L58
L58:
	;
	v218 = v165
	goto L29
L59:
	;
	if base.Ui32(v119) < base.Ui32(v168) {
		goto L30
	} else {
		goto L67
	}
L60:
	;
	if base.Ui32(v168) < base.Ui32(v120) {
		v198 = v111
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v171) < base.Ui32(v172) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	if base.Ui32(v172) < base.Ui32(v171) {
		v198 = v111
		goto L32
	} else {
		goto L63
	}
L63:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v175 < v176 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	if v176 < v175 {
		v198 = v111
		goto L32
	} else {
		goto L65
	}
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if base.Ui32(v180) < base.Ui32(v179) {
		v198 = v111
		goto L32
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	if base.Ui32(v168) < base.Ui32(v119) {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v186) < base.Ui32(v187) {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	if base.Ui32(v187) < base.Ui32(v186) {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v190 < v191 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	if v191 < v190 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if base.Ui32(v194) < base.Ui32(v195) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v197 = v20
	goto L75
L74:
	;
	v197 = v114
	goto L75
L75:
	;
	v198 = v197
	goto L32
L76:
	;
	v329 = v104 + v108*int32(-40)
	v330 = v104 + v220
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if base.Ui32(v335) < base.Ui32(v336) {
		goto L128
	} else {
		goto L129
	}
L77:
	;
	v326 = v221
	goto L76
L78:
	;
	v326 = v222
	goto L76
L79:
	;
	v326 = v306
	goto L76
L80:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if base.Ui32(v228) < base.Ui32(v276) {
		goto L106
	} else {
		goto L107
	}
L81:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if base.Ui32(v228) < base.Ui32(v244) {
		v306 = v100
		goto L79
	} else {
		goto L89
	}
L82:
	;
	if base.Ui32(v228) < base.Ui32(v227) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if base.Ui32(v231) < base.Ui32(v232) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	if base.Ui32(v232) < base.Ui32(v231) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v235 < v236 {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	if v236 < v235 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if base.Ui32(v240) <= base.Ui32(v239) {
		goto L80
	} else {
		goto L88
	}
L88:
	;
	goto L81
L89:
	;
	if base.Ui32(v244) < base.Ui32(v228) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if base.Ui32(v227) < base.Ui32(v244) {
		goto L78
	} else {
		goto L97
	}
L91:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if base.Ui32(v247) < base.Ui32(v248) {
		v306 = v100
		goto L79
	} else {
		goto L92
	}
L92:
	;
	if base.Ui32(v248) < base.Ui32(v247) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if v251 < v252 {
		v306 = v100
		goto L79
	} else {
		goto L94
	}
L94:
	;
	if v252 < v251 {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if base.Ui32(v255) < base.Ui32(v256) {
		v306 = v100
		goto L79
	} else {
		goto L96
	}
L96:
	;
	goto L90
L97:
	;
	if base.Ui32(v244) < base.Ui32(v227) {
		goto L77
	} else {
		goto L98
	}
L98:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if base.Ui32(v262) < base.Ui32(v263) {
		goto L78
	} else {
		goto L99
	}
L99:
	;
	if base.Ui32(v263) < base.Ui32(v262) {
		goto L77
	} else {
		goto L100
	}
L100:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if v266 < v267 {
		goto L78
	} else {
		goto L101
	}
L101:
	;
	if v267 < v266 {
		goto L77
	} else {
		goto L102
	}
L102:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if base.Ui32(v270) < base.Ui32(v271) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v273 = v222
	goto L105
L104:
	;
	v273 = v221
	goto L105
L105:
	;
	v326 = v273
	goto L76
L106:
	;
	if base.Ui32(v227) < base.Ui32(v276) {
		goto L77
	} else {
		goto L114
	}
L107:
	;
	if base.Ui32(v276) < base.Ui32(v228) {
		v306 = v100
		goto L79
	} else {
		goto L108
	}
L108:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if base.Ui32(v279) < base.Ui32(v280) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	if base.Ui32(v280) < base.Ui32(v279) {
		v306 = v100
		goto L79
	} else {
		goto L110
	}
L110:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if v283 < v284 {
		goto L106
	} else {
		goto L111
	}
L111:
	;
	if v284 < v283 {
		v306 = v100
		goto L79
	} else {
		goto L112
	}
L112:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if base.Ui32(v288) < base.Ui32(v287) {
		v306 = v100
		goto L79
	} else {
		goto L113
	}
L113:
	;
	goto L106
L114:
	;
	if base.Ui32(v276) < base.Ui32(v227) {
		goto L78
	} else {
		goto L115
	}
L115:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if base.Ui32(v294) < base.Ui32(v295) {
		goto L77
	} else {
		goto L116
	}
L116:
	;
	if base.Ui32(v295) < base.Ui32(v294) {
		goto L78
	} else {
		goto L117
	}
L117:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if v298 < v299 {
		goto L77
	} else {
		goto L118
	}
L118:
	;
	if v299 < v298 {
		goto L78
	} else {
		goto L119
	}
L119:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	if base.Ui32(v302) < base.Ui32(v303) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v305 = v221
	goto L122
L121:
	;
	v305 = v222
	goto L122
L122:
	;
	v306 = v305
	goto L79
L123:
	;
	v435 = v218
	v436 = v326
	v437 = v434
	goto L25
L124:
	;
	v434 = v329
	goto L123
L125:
	;
	v434 = v104
	goto L123
L126:
	;
	v434 = v414
	goto L123
L127:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if base.Ui32(v336) < base.Ui32(v384) {
		goto L153
	} else {
		goto L154
	}
L128:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if base.Ui32(v336) < base.Ui32(v352) {
		v414 = v330
		goto L126
	} else {
		goto L136
	}
L129:
	;
	if base.Ui32(v336) < base.Ui32(v335) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(v339) < base.Ui32(v340) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	if base.Ui32(v340) < base.Ui32(v339) {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	if v343 < v344 {
		goto L128
	} else {
		goto L133
	}
L133:
	;
	if v344 < v343 {
		goto L127
	} else {
		goto L134
	}
L134:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	if base.Ui32(v348) <= base.Ui32(v347) {
		goto L127
	} else {
		goto L135
	}
L135:
	;
	goto L128
L136:
	;
	if base.Ui32(v352) < base.Ui32(v336) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	if base.Ui32(v335) < base.Ui32(v352) {
		goto L125
	} else {
		goto L144
	}
L138:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if base.Ui32(v355) < base.Ui32(v356) {
		v414 = v330
		goto L126
	} else {
		goto L139
	}
L139:
	;
	if base.Ui32(v356) < base.Ui32(v355) {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v359 < v360 {
		v414 = v330
		goto L126
	} else {
		goto L141
	}
L141:
	;
	if v360 < v359 {
		goto L137
	} else {
		goto L142
	}
L142:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if base.Ui32(v363) < base.Ui32(v364) {
		v414 = v330
		goto L126
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	if base.Ui32(v352) < base.Ui32(v335) {
		goto L124
	} else {
		goto L145
	}
L145:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if base.Ui32(v370) < base.Ui32(v371) {
		goto L125
	} else {
		goto L146
	}
L146:
	;
	if base.Ui32(v371) < base.Ui32(v370) {
		goto L124
	} else {
		goto L147
	}
L147:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v374 < v375 {
		goto L125
	} else {
		goto L148
	}
L148:
	;
	if v375 < v374 {
		goto L124
	} else {
		goto L149
	}
L149:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if base.Ui32(v378) < base.Ui32(v379) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v381 = v104
	goto L152
L151:
	;
	v381 = v329
	goto L152
L152:
	;
	v434 = v381
	goto L123
L153:
	;
	if base.Ui32(v335) < base.Ui32(v384) {
		goto L124
	} else {
		goto L161
	}
L154:
	;
	if base.Ui32(v384) < base.Ui32(v336) {
		v414 = v330
		goto L126
	} else {
		goto L155
	}
L155:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if base.Ui32(v387) < base.Ui32(v388) {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	if base.Ui32(v388) < base.Ui32(v387) {
		v414 = v330
		goto L126
	} else {
		goto L157
	}
L157:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v391 < v392 {
		goto L153
	} else {
		goto L158
	}
L158:
	;
	if v392 < v391 {
		v414 = v330
		goto L126
	} else {
		goto L159
	}
L159:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if base.Ui32(v396) < base.Ui32(v395) {
		v414 = v330
		goto L126
	} else {
		goto L160
	}
L160:
	;
	goto L153
L161:
	;
	if base.Ui32(v384) < base.Ui32(v335) {
		goto L125
	} else {
		goto L162
	}
L162:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if base.Ui32(v402) < base.Ui32(v403) {
		goto L124
	} else {
		goto L163
	}
L163:
	;
	if base.Ui32(v403) < base.Ui32(v402) {
		goto L125
	} else {
		goto L164
	}
L164:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v406 < v407 {
		goto L124
	} else {
		goto L165
	}
L165:
	;
	if v407 < v406 {
		goto L125
	} else {
		goto L166
	}
L166:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if base.Ui32(v410) < base.Ui32(v411) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v413 = v329
	goto L169
L168:
	;
	v413 = v104
	goto L169
L169:
	;
	v414 = v413
	goto L126
L170:
	;
	v546 = v544
	goto L24
L171:
	;
	v544 = v435
	goto L170
L172:
	;
	v544 = v437
	goto L170
L173:
	;
	v544 = v524
	goto L170
L174:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if base.Ui32(v446) < base.Ui32(v494) {
		goto L200
	} else {
		goto L201
	}
L175:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if base.Ui32(v446) < base.Ui32(v462) {
		v524 = v436
		goto L173
	} else {
		goto L183
	}
L176:
	;
	if base.Ui32(v446) < base.Ui32(v445) {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	if base.Ui32(v449) < base.Ui32(v450) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	if base.Ui32(v450) < base.Ui32(v449) {
		goto L174
	} else {
		goto L179
	}
L179:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	if v453 < v454 {
		goto L175
	} else {
		goto L180
	}
L180:
	;
	if v454 < v453 {
		goto L174
	} else {
		goto L181
	}
L181:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v435)+12))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v436)+12))
	if base.Ui32(v458) <= base.Ui32(v457) {
		goto L174
	} else {
		goto L182
	}
L182:
	;
	goto L175
L183:
	;
	if base.Ui32(v462) < base.Ui32(v446) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if base.Ui32(v445) < base.Ui32(v462) {
		goto L172
	} else {
		goto L191
	}
L185:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if base.Ui32(v465) < base.Ui32(v466) {
		v524 = v436
		goto L173
	} else {
		goto L186
	}
L186:
	;
	if base.Ui32(v466) < base.Ui32(v465) {
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	if v469 < v470 {
		v524 = v436
		goto L173
	} else {
		goto L188
	}
L188:
	;
	if v470 < v469 {
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v436)+12))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	if base.Ui32(v473) < base.Ui32(v474) {
		v524 = v436
		goto L173
	} else {
		goto L190
	}
L190:
	;
	goto L184
L191:
	;
	if base.Ui32(v462) < base.Ui32(v445) {
		goto L171
	} else {
		goto L192
	}
L192:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if base.Ui32(v480) < base.Ui32(v481) {
		goto L172
	} else {
		goto L193
	}
L193:
	;
	if base.Ui32(v481) < base.Ui32(v480) {
		goto L171
	} else {
		goto L194
	}
L194:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	if v484 < v485 {
		goto L172
	} else {
		goto L195
	}
L195:
	;
	if v485 < v484 {
		goto L171
	} else {
		goto L196
	}
L196:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v435)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	if base.Ui32(v488) < base.Ui32(v489) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v491 = v437
	goto L199
L198:
	;
	v491 = v435
	goto L199
L199:
	;
	v544 = v491
	goto L170
L200:
	;
	if base.Ui32(v445) < base.Ui32(v494) {
		goto L171
	} else {
		goto L208
	}
L201:
	;
	if base.Ui32(v494) < base.Ui32(v446) {
		v524 = v436
		goto L173
	} else {
		goto L202
	}
L202:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if base.Ui32(v497) < base.Ui32(v498) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	if base.Ui32(v498) < base.Ui32(v497) {
		v524 = v436
		goto L173
	} else {
		goto L204
	}
L204:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v436)+8))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	if v501 < v502 {
		goto L200
	} else {
		goto L205
	}
L205:
	;
	if v502 < v501 {
		v524 = v436
		goto L173
	} else {
		goto L206
	}
L206:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v436)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	if base.Ui32(v506) < base.Ui32(v505) {
		v524 = v436
		goto L173
	} else {
		goto L207
	}
L207:
	;
	goto L200
L208:
	;
	if base.Ui32(v494) < base.Ui32(v445) {
		goto L172
	} else {
		goto L209
	}
L209:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if base.Ui32(v512) < base.Ui32(v513) {
		goto L171
	} else {
		goto L210
	}
L210:
	;
	if base.Ui32(v513) < base.Ui32(v512) {
		goto L172
	} else {
		goto L211
	}
L211:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v435)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	if v516 < v517 {
		goto L171
	} else {
		goto L212
	}
L212:
	;
	if v517 < v516 {
		goto L172
	} else {
		goto L213
	}
L213:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v435)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	if base.Ui32(v520) < base.Ui32(v521) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v523 = v435
	goto L216
L215:
	;
	v523 = v437
	goto L216
L216:
	;
	v524 = v523
	goto L173
L217:
	;
	if base.Ui32(v573) < base.Ui32(v574) {
		v643 = v574
		v645 = v576
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if base.Ui32(v643) <= base.Ui32(v573) {
		goto L234
	} else {
		goto L235
	}
L220:
	;
	v588 = v574
	v590 = v576
	goto L221
L221:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v588)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if base.Ui32(v598) < base.Ui32(v599) {
		v635 = v590
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v643 = v638
	v645 = v635
	goto L219
L223:
	;
	v638 = v588 + int32(20)
	if base.Ui32(v638) <= base.Ui32(v573) {
		v588 = v638
		v590 = v635
		goto L221
	} else {
		goto L232
	}
L224:
	;
	if base.Ui32(v599) < base.Ui32(v598) {
		v643 = v588
		v645 = v590
		goto L219
	} else {
		goto L225
	}
L225:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.Ui32(v602) < base.Ui32(v603) {
		v635 = v590
		goto L223
	} else {
		goto L226
	}
L226:
	;
	if base.Ui32(v603) < base.Ui32(v602) {
		v643 = v588
		v645 = v590
		goto L219
	} else {
		goto L227
	}
L227:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v588)+8))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v606 < v607 {
		v635 = v590
		goto L223
	} else {
		goto L228
	}
L228:
	;
	if v607 < v606 {
		v643 = v588
		v645 = v590
		goto L219
	} else {
		goto L229
	}
L229:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v588)+12))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(v610) < base.Ui32(v611) {
		v635 = v590
		goto L223
	} else {
		goto L230
	}
L230:
	;
	if base.Ui32(v611) < base.Ui32(v610) {
		v643 = v588
		v645 = v590
		goto L219
	} else {
		goto L231
	}
L231:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v590)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v614
	v616 = *(*int64)(unsafe.Add(mBase, uint32(v590)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v616
	v618 = *(*int64)(unsafe.Add(mBase, uint32(v590)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v618
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v588)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v590)+16)) = v620
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v588)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v590)+8)) = v622
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v588)))
	*(*int64)(unsafe.Add(mBase, uint32(v590))) = v624
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+16)) = v626
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v588)+8)) = v628
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v588))) = v630
	v635 = v590 + int32(20)
	goto L223
L232:
	;
	goto L222
L233:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v643)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v865
	v867 = *(*int64)(unsafe.Add(mBase, uint32(v643)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v867
	v869 = *(*int64)(unsafe.Add(mBase, uint32(v643)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v869
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v656)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v643)+16)) = v871
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v656)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v643)+8)) = v873
	v875 = *(*int64)(unsafe.Add(mBase, uint32(v656)))
	*(*int64)(unsafe.Add(mBase, uint32(v643))) = v875
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v656)+16)) = v877
	v879 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v656)+8)) = v879
	v881 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v656))) = v881
	v883 = int32(20)
	v573 = v656 - v883
	v574 = v643 + v883
	v576 = v645
	v579 = v662
	goto L217
L234:
	;
	v656 = v573
	v662 = v579
	goto L237
L235:
	;
	v711 = v573
	v717 = v579
	goto L236
L236:
	;
	v723 = int32(20)
	v724 = base.I32_div_s(v645-v20, v723)
	v727 = base.I32_div_s(v643-v645, v723)
	if v724 < v727 {
		goto L249
	} else {
		goto L250
	}
L237:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if base.Ui32(v667) < base.Ui32(v668) {
		goto L233
	} else {
		goto L239
	}
L238:
	;
	v711 = v707
	v717 = v705
	goto L236
L239:
	;
	if base.Ui32(v668) < base.Ui32(v667) {
		v705 = v662
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v707 = v656 - int32(20)
	if base.Ui32(v643) <= base.Ui32(v707) {
		v656 = v707
		v662 = v705
		goto L237
	} else {
		goto L248
	}
L241:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.Ui32(v671) < base.Ui32(v672) {
		goto L233
	} else {
		goto L242
	}
L242:
	;
	if base.Ui32(v672) < base.Ui32(v671) {
		v705 = v662
		goto L240
	} else {
		goto L243
	}
L243:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v656)+8))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v675 < v676 {
		goto L233
	} else {
		goto L244
	}
L244:
	;
	if v676 < v675 {
		v705 = v662
		goto L240
	} else {
		goto L245
	}
L245:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v656)+12))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if base.Ui32(v679) < base.Ui32(v680) {
		goto L233
	} else {
		goto L246
	}
L246:
	;
	if base.Ui32(v680) < base.Ui32(v679) {
		v705 = v662
		goto L240
	} else {
		goto L247
	}
L247:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v656)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v683
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v656)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v685
	v687 = *(*int64)(unsafe.Add(mBase, uint32(v656)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v687
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v662)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v656)+16)) = v689
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v662)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v656)+8)) = v691
	v693 = *(*int64)(unsafe.Add(mBase, uint32(v662)))
	*(*int64)(unsafe.Add(mBase, uint32(v656))) = v693
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v662)+16)) = v695
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v662)+8)) = v697
	v699 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v662))) = v699
	v705 = v662 - int32(20)
	goto L240
L248:
	;
	goto L238
L249:
	;
	v729 = v724
	goto L251
L250:
	;
	v729 = v727
	goto L251
L251:
	;
	if v729 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v739 = int32(0)
	goto L255
L253:
	;
	goto L254
L254:
	;
	v786 = int32(20)
	v787 = base.I32_div_s(v717-v711, v786)
	v790 = base.I32_div_s(v52-v717, v786)
	v792 = v790 - int32(1)
	if v787 < v792 {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v748 = v739 * int32(20)
	v749 = v20 + v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v750
	v752 = *(*int64)(unsafe.Add(mBase, uint32(v749)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v752
	v754 = *(*int64)(unsafe.Add(mBase, uint32(v749)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v754
	v756 = v748 + (v643 + v729*int32(-20))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v749)+16)) = v757
	v759 = *(*int64)(unsafe.Add(mBase, uint32(v756)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v749)+8)) = v759
	v761 = *(*int64)(unsafe.Add(mBase, uint32(v756)))
	*(*int64)(unsafe.Add(mBase, uint32(v749))) = v761
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+16)) = v763
	v765 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v756)+8)) = v765
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v756))) = v767
	v770 = v739 + int32(1)
	if v770 != v729 {
		v739 = v770
		goto L255
	} else {
		goto L257
	}
L256:
	;
	goto L254
L257:
	;
	goto L256
L258:
	;
	v794 = v787
	goto L260
L259:
	;
	v794 = v792
	goto L260
L260:
	;
	if v794 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v811 = int32(0)
	goto L264
L262:
	;
	goto L263
L263:
	;
	if base.Ui32(v727) <= base.Ui32(v787) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v813 = v811 * int32(20)
	v814 = v643 + v813
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v815
	v817 = *(*int64)(unsafe.Add(mBase, uint32(v814)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v817
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v814)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v819
	v821 = v813 + (v52 + v794*int32(-20))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v814)+16)) = v822
	v824 = *(*int64)(unsafe.Add(mBase, uint32(v821)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v814)+8)) = v824
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v821)))
	*(*int64)(unsafe.Add(mBase, uint32(v814))) = v826
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v821)+16)) = v828
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v821)+8)) = v830
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v821))) = v832
	v835 = v811 + int32(1)
	if v835 != v794 {
		v811 = v835
		goto L264
	} else {
		goto L266
	}
L265:
	;
	goto L263
L266:
	;
	goto L265
L267:
	;
	F_sort_checkpoint_bufferids(m, v20, v727)
	mBase = m.M
	v854 = v52 + v787*int32(-20)
	if base.Ui32(v787) < base.Ui32(int32(7)) {
		v900 = v854
		v902 = v787
		goto L2
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_sort_checkpoint_bufferids(m, v52+v787*int32(-20), v787)
	mBase = m.M
	if base.Ui32(v727) < base.Ui32(int32(7)) {
		v887 = v20
		v888 = v727
		goto L3
	} else {
		goto L271
	}
L270:
	;
	v20 = v854
	v21 = v787
	goto L5
L271:
	;
	v36 = v727
	goto L7
L272:
	;
	v915 = int32(20)
	v925 = v900 + v915
	goto L273
L273:
	;
	if base.Ui32(v925) <= base.Ui32(v900) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	goto L1
L275:
	;
	v1005 = v925 + int32(20)
	if base.Ui32(v1005) < base.Ui32(v900+v902*v915) {
		v925 = v1005
		goto L273
	} else {
		goto L288
	}
L276:
	;
	v937 = v925
	goto L277
L277:
	;
	v948 = v937 - int32(20)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if base.Ui32(v949) < base.Ui32(v950) {
		goto L275
	} else {
		goto L279
	}
L278:
	;
	goto L275
L279:
	;
	if base.Ui32(v950) < base.Ui32(v949) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v937)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v972
	v974 = *(*int64)(unsafe.Add(mBase, uint32(v937)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v974
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v937)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v976
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v948)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v937)+16)) = v978
	v980 = *(*int64)(unsafe.Add(mBase, uint32(v948)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v937)+8)) = v980
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v948)))
	*(*int64)(unsafe.Add(mBase, uint32(v937))) = v982
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v948)+16)) = v984
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v948)+8)) = v986
	v988 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v948))) = v988
	if base.Ui32(v900) < base.Ui32(v948) {
		v937 = v948
		goto L277
	} else {
		goto L287
	}
L281:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v937-int32(16))))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v937)+4))
	if base.Ui32(v955) < base.Ui32(v956) {
		goto L275
	} else {
		goto L282
	}
L282:
	;
	if base.Ui32(v956) < base.Ui32(v955) {
		goto L280
	} else {
		goto L283
	}
L283:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v937-int32(12))))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v937)+8))
	if v961 < v962 {
		goto L275
	} else {
		goto L284
	}
L284:
	;
	if v962 < v961 {
		goto L280
	} else {
		goto L285
	}
L285:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v937-int32(8))))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v937)+12))
	if base.Ui32(v967) <= base.Ui32(v968) {
		goto L275
	} else {
		goto L286
	}
L286:
	;
	goto L280
L287:
	;
	goto L278
L288:
	;
	goto L274
}
func F_sort_checkpoint_bufferids_med3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v88 int32
	_ = v88
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v8) < base.Ui32(v9) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return l0
L2:
	;
	return l2
L3:
	;
	return v88
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v9) < base.Ui32(v58) {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v9) < base.Ui32(v25) {
		v88 = l1
		goto L3
	} else {
		goto L13
	}
L6:
	;
	if base.Ui32(v9) < base.Ui32(v8) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v12) < base.Ui32(v13) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(v13) < base.Ui32(v12) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 < v17 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v17 < v16 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v21) <= base.Ui32(v20) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	if base.Ui32(v25) < base.Ui32(v9) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui32(v8) < base.Ui32(v25) {
		goto L2
	} else {
		goto L21
	}
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v28) < base.Ui32(v29) {
		v88 = l1
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v29) < base.Ui32(v28) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v32 < v33 {
		v88 = l1
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v33 < v32 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v36) < base.Ui32(v37) {
		v88 = l1
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	if base.Ui32(v25) < base.Ui32(v8) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v43) < base.Ui32(v44) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v44) < base.Ui32(v43) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v47 < v48 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	if v48 < v47 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v51) < base.Ui32(v52) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v54 = l2
	goto L29
L28:
	;
	v54 = l0
	goto L29
L29:
	;
	return v54
L30:
	;
	if base.Ui32(v8) < base.Ui32(v58) {
		goto L1
	} else {
		goto L38
	}
L31:
	;
	if base.Ui32(v58) < base.Ui32(v9) {
		v88 = l1
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v61) < base.Ui32(v62) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v62) < base.Ui32(v61) {
		v88 = l1
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v65 < v66 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v66 < v65 {
		v88 = l1
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v70) < base.Ui32(v69) {
		v88 = l1
		goto L3
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	if base.Ui32(v58) < base.Ui32(v8) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v76) < base.Ui32(v77) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v77) < base.Ui32(v76) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v80 < v81 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v81 < v80 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v84) < base.Ui32(v85) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v87 = l0
	goto L46
L45:
	;
	v87 = l2
	goto L46
L46:
	;
	v88 = v87
	goto L3
}
func F_sort_desc(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14377(m, l0, int32(_a_F_sort_desc_0), int32(245), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_sort_order_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 float32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 float32
	_ = v15
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
	v10 = *(*float32)(unsafe.Add(mBase, uint32(v7+v8)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v15 = *(*float32)(unsafe.Add(mBase, uint32(v12+v13)+8))
	if base.F32_lt(v10, v15) != 0 {
		v18 = int32(-1)
	} else {
		v18 = base.F32_gt(v10, v15)
	}
	return v18
}
func F_sortins(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(2) <= v10 {
		v13 = int32(2)
		v16 = F_palloc_extended(m, v10<<(uint(v13)%32), v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(101)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				if v24 != 0 {
					v26 = v24
				} else {
					v26 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v26
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				if v28 != 0 {
					v29 = v28
					v33 = int32(0)
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v16+v33<<(uint(int32(2))%32)))) = v29
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
						if v44 != 0 {
							v29 = v44
							v33 = v33 + int32(1)
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				F_pg_qsort(m, v16, v10, int32(4), int32(1035))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v60
					if v10 == int32(2) {
						v140 = int32(1)
					} else {
						v67 = int32(1)
						v69 = v10 - v67
						if v10 != int32(3) {
							v80 = int32(0)
							v83 = v67
							for {
								v88 = int32(2)
								v90 = v16 + v83<<(uint(v88)%32)
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v92 = int32(4)
								v93 = v90 + v92
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
								*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = v94
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v90-v92)))
								*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v98
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
								v102 = v83 + v88
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v16+v102<<(uint(v88)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v100)+24)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								*(*int32)(unsafe.Add(mBase, uint32(v100)+28)) = v108
								if v80 != v10&int32(2147483646)-int32(4) {
									v80 = v80 + v88
									v83 = v102
									continue
								} else {
									break
								}
								break
							}
							if v10&int32(1) == int32(0) {
								v140 = v69
							} else {
								v119 = v102
								v126 = v16 + v119<<(uint(int32(2))%32)
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v128
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v126-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v132
								v140 = v69
							}
						} else {
							v119 = v67
							v126 = v16 + v119<<(uint(int32(2))%32)
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+24)) = v128
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v126-int32(4))))
							*(*int32)(unsafe.Add(mBase, uint32(v127)+28)) = v132
							v140 = v69
						}
					}
					v145 = v16 + v140<<(uint(int32(2))%32)
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
					*(*int32)(unsafe.Add(mBase, uint32(v146)+24)) = int32(0)
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v145-int32(4))))
					*(*int32)(unsafe.Add(mBase, uint32(v146)+28)) = v151
					F_pfree(m, v16)
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_spgadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v13 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(0)
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = l0
	v29 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+24)) = uint16(v29)
	v32 = v18 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 < v33 {
		v18 = v32
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(16)
	return
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v43 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = int32(0)
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v48<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui32(int32(5)) <= base.Ui32(v58-int32(1)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L7
L12:
	;
	v92 = v48 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v92 < v93 {
		v48 = v92
		goto L10
	} else {
		goto L24
	}
L13:
	;
	if base.Ui32(v58-int32(6)) < base.Ui32(int32(2)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+24)) = uint8(v89)
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = l0
	v68 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+24)) = uint16(v68)
	goto L12
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_spgadjustmembers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
	F_errmsg(m, int32(_a_F_spgadjustmembers_1), v9)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_spgadjustmembers_2), int32(379), int32(_a_F_spgadjustmembers_3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L11
}
func F_spgbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v276 int32
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	v8 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = F_palloc0(m, int32(_a_F_spgbeginscan_0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = F_palloc_mul(m, int32(56), l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v21 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	F_initSpGistState(m, v13, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v21 = v18
	goto L6
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
	v32 = F_AllocSetContextCreateInternal(m, v27, int32(_a_F_spgbeginscan_1), int32(0), int32(_a_F_spgbeginscan_2), int32(_a_F_spgbeginscan_3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v32
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_spgbeginscan_4), int32(0), int32(_a_F_spgbeginscan_2), int32(_a_F_spgbeginscan_3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = v41
	v45 = v13 + int32(20)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48<<(uint(int32(3))%32))+96))
	if v46 != v52 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = F_CreateTupleDescCopy(m, v47)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v166 = v47
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v13)+212)) = v166
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v172 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v60 = v54 + v57<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+104)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+96)) = v56
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+100)) = uint16(v64)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+110)) = uint8(v66)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+111)) = uint8(v68)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+8)))
	v71 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+124)) = v71
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+113)) = uint8(v71)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+112)) = uint8(v70)
	F_populate_compact_attribute(m, v54, v71)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v79 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v79 < v88 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v166 = v54
	goto L13
L17:
	;
	v92 = v54 + int32(28)
	v99 = v79
	v100 = v88
	v102 = v79
	goto L21
L18:
	;
	v156 = v79
	v163 = v88
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v156
	goto L16
L20:
	;
	v156 = v150
	v163 = v129
	goto L19
L21:
	;
	v108 = v92 + v88<<(uint(int32(3))%32) + v99*int32(100)
	v111 = v92 + v99<<(uint(int32(3))%32)
	if v88 != v100 {
		v129 = v100
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v150 = v88
	goto L20
L23:
	;
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+2)))
	if v130 <= int32(0) {
		v150 = v99
		goto L20
	} else {
		goto L31
	}
L24:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+7)))
	if v113 != int32(118) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v129 = v99
	goto L23
L26:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	if v116 != int32(1) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+6)))
	if v119&int32(6) != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+2)))
	if v122 <= int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+90)))
	if v125 != int32(118) {
		v129 = v88
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+90)))
	if v133 == int32(118) {
		v150 = v99
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
	v142 = (v102 + v136 - int32(1)) & (int32(0) - v136)
	if int32(_a_F_spgbeginscan_5) < v142 {
		v150 = v99
		goto L20
	} else {
		goto L33
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v142)
	v148 = v99 + int32(1)
	if v148 != v88 {
		v99 = v148
		v100 = v129
		v102 = v142 + v130
		goto L21
	} else {
		goto L34
	}
L34:
	;
	goto L22
L35:
	;
	v248 = v13 + int32(132)
	v251 = F_index_getprocinfo(m, l0, int32(1), int32(4))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L50
	}
L36:
	;
	v176 = F_palloc_mul(m, int32(4), v172)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v176
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v181 = F_palloc_mul(m, int32(4), v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v181
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v186 = F_palloc_mul(m, int32(8), v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v186
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v191 = F_palloc_mul(m, int32(8), v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = v191
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if int32(0) < v194 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v199 = int32(0)
	goto L44
L42:
	;
	v223 = v194
	goto L43
L43:
	;
	v227 = F_palloc0_mul(m, int32(8), v223)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L47
	}
L44:
	;
	v206 = v199 << (uint(int32(3)) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
	*(*int64)(unsafe.Add(mBase, uint32(v206+v207))) = int64(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v211+v206))) = int64(9218868437227405312)
	v216 = v199 + int32(1)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v216 < v217 {
		v199 = v216
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v223 = v217
	goto L43
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v232 = F_palloc_mul(m, int32(1), v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v235 == int32(0) {
		goto L35
	} else {
		goto L49
	}
L49:
	;
	base.MemoryFill(m, v232, int32(1), v235)
	goto L35
L50:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v251)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+16)) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+24)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v251)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v248)+8)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
	*(*int64)(unsafe.Add(mBase, uint32(v248))) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v248)+20)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = int32(0)
	goto L51
L51:
	;
	v267 = v13 + int32(160)
	v270 = F_index_getprocinfo(m, l0, int32(1), int32(5))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_spgbeginscan[0]))
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v270)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+16)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+24)) = v276
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v270)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+8)) = v278
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
	*(*int64)(unsafe.Add(mBase, uint32(v267))) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v267)+20)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v267)+16)) = int32(0)
	goto L53
L53:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v13
	return v8
}
func F_spgoptions(m *base.Module, l0 int64, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(256), int32(8), int32(_a_F_spgoptions_0), int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_spgvacuumscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 float64
	_ = v265
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 float64
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v784 int32
	_ = v784
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	v19 = m.G0
	v21 = v19 - int32(880)
	m.G0 = v21
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	F_initSpGistState(m, l0+int32(16), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+4)) = uint8(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = int64(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = v39
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	if v47 == v39 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v53 = base.B2i32(v50 == int32(0))
	goto L6
L5:
	;
	v53 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v61 = int32(0)
	v66 = F_read_stream_begin_relation(m, int32(13), v60, v26, v61, int32(3), v21+int32(40), v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v68 = l0
	v71 = v21
	v77 = v26
	v81 = l0 + int32(100)
	v83 = v66
	v85 = v53
	goto L8
L8:
	;
	if v85 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	if base.Ui32(v101) < base.Ui32(v100) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v89 = F_RelationGetNumberOfBlocksInFork(m, v77, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LockRelationForExtension(m, v77, int32(7))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v100 = v89
	goto L10
L15:
	;
	v95 = F_RelationGetNumberOfBlocksInFork(m, v77, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_UnlockRelationForExtension(m, v77, int32(7))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v100 = v95
	goto L10
L18:
	;
	F_read_stream_reset(m, v119)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L189
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+44)) = v100
	v104 = v68
	v107 = v71
	v113 = v77
	v117 = v81
	v119 = v83
	v121 = v85
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_read_stream_end(m, v83)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L183
	}
L22:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L180
	}
L24:
	;
	v126 = F_read_stream_next_buffer(m, v119, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v126 == int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v126 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_LockBufferInternal(m, v126, int32(3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[1]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135+(v126^int32(-1))*int32(56))+16))
	v150 = v141
	goto L27
L29:
	;
	goto L30
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[2]))
	v144 = int32(56)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+v126*v144-v144)+16))
	v150 = v149
	goto L27
L31:
	;
	v154 = int32(0)
	v155 = base.B2i32(v154 <= v126)
	if v155 == v154 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v126)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L92
	}
L33:
	;
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+14)))
	if v420 != 0 {
		goto L86
	} else {
		goto L87
	}
L34:
	;
	if base.Ui32(v150-int32(1)) < base.Ui32(int32(2)) {
		goto L32
	} else {
		goto L84
	}
L35:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+14)))
	if v174 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L36:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+(v126^int32(-1))<<(uint(int32(2))%32))))
	v173 = v165
	goto L35
L37:
	;
	goto L38
L38:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v173 = v167 + v126<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+12)))
	if base.Ui32(v177) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+16)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v180))))
	if v182&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v185 = int32(1)
	if base.Ui32(v150-v185) <= base.Ui32(v185) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	F_vacuumRedirectAndPlaceholder(m, v131, v377, v126)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L83
	}
L44:
	;
	if v155 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	goto L46
L46:
	;
	F_vacuumLeafPage(m, v104, v131, v126, int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L81
	}
L47:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+12)))
	v208 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+868)) = uint16(v208)
	if base.Ui32(v207) < base.Ui32(int32(25)) {
		goto L34
	} else {
		goto L51
	}
L48:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v126^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L47
L49:
	;
	goto L50
L50:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v206 = v200 + v126<<(uint(int32(13))%32) + int32(-8192)
	goto L47
L51:
	;
	v217 = int32(base.Ui32(v207+int32(_a_F_spgvacuumscan_0))>>(uint(int32(2))%32)) & int32(_a_F_spgvacuumscan_1)
	if v217 == int32(0) {
		goto L34
	} else {
		goto L52
	}
L52:
	;
	v228 = int32(1)
	v230 = int32(0)
	goto L53
L53:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(20)+v228&int32(_a_F_spgvacuumscan_1)<<(uint(int32(2))%32))))
	v250 = v206 + v247&int32(_a_F_spgvacuumscan_2)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v251&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v309 = v302 & int32(_a_F_spgvacuumscan_1)
	if v309 == int32(0) {
		goto L34
	} else {
		goto L66
	}
L55:
	;
	v304 = v228 + int32(1)
	if base.Ui32(v304&int32(_a_F_spgvacuumscan_1)) <= base.Ui32(v217) {
		v228 = v304
		v230 = v302
		goto L53
	} else {
		goto L65
	}
L56:
	;
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v262)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v262)+8)) = base.F64_add(v298, float64(1))
	v302 = v230
	goto L55
L57:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v260 = m.T0[v259].(func(*base.Module, int32, int32) int32)(m, v250+int32(6), v258)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v260 == int32(0) {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v265 = *(*float64)(unsafe.Add(mBase, uint32(v262)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v262)+16)) = base.F64_add(v265, float64(1))
	v273 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v107+int32(48)+v230&int32(_a_F_spgvacuumscan_1)<<(uint(v273)%32)))) = uint16(v228)
	v278 = v230 + v273
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+868)) = uint16(v278)
	v302 = v278
	goto L55
L62:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+32)) = v284 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_3), v107+int32(32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(445), int32(_a_F_spgvacuumscan_5))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	goto L54
L66:
	;
	v312 = int32(_a_F_spgvacuumscan_6)
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5])) = v314 + int32(1)
	F_PageIndexMultiDelete(m, v206, v107+int32(48), v309)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_MarkBufferDirty(m, v126)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+118)))
	if v325 != int32(112) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v363 = int32(_a_F_spgvacuumscan_6)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[5])) = v365 - int32(1)
	goto L34
L70:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[6]))
	if v329 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v131)+32))
	if v332 != 0 {
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v131)+40))
	if v333 != 0 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v104)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+872)) = v336
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+876)) = uint8(v338)
	F_XLogRegisterData(m, v107+int32(868), int32(12))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+868)))
	F_XLogRegisterData(m, v107+int32(48), v347<<(uint(int32(1))%32))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_XLogRegisterBuffer(m, int32(0), v126, int32(8))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v358 = F_XLogInsert(m, int32(16), int32(112))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v206))) = base.I64_rotl(v358, int64(32))
	goto L69
L81:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	F_vacuumRedirectAndPlaceholder(m, v131, v373, v126)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L33
L83:
	;
	goto L34
L84:
	;
	goto L33
L85:
	;
	F_SpGistSetLastUsedPage(m, v131, v126)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v421) {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_RecordFreeIndexPage(m, v131, v150)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v426)+28)) = v427 + int32(1)
	goto L32
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+108)) = v150
	goto L32
L92:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v454 == int32(0) {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v459 = v104
	v462 = v107
	v464 = v458
	v465 = v454
	v468 = v113
	v472 = v117
	v474 = v119
	v476 = v121
	goto L95
L94:
	;
	goto L23
L95:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+6)))
	if v477 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	if v905 != 0 {
		goto L173
	} else {
		goto L174
	}
L97:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	v886 = v459
	v889 = v462
	v891 = v464
	v895 = v468
	v899 = v472
	v901 = v474
	v903 = v476
	goto L99
L99:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v904 != 0 {
		v459 = v886
		v462 = v889
		v464 = v891
		v465 = v904
		v468 = v895
		v472 = v899
		v474 = v901
		v476 = v903
		goto L95
	} else {
		goto L172
	}
L100:
	;
	v483 = int32(0)
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+2)))
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465))))
	v488 = v484 | v485<<(uint(int32(16))%32)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+24))
	v492 = F_ReadBufferExtended(m, v464, v483, v488, v483, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_LockBufferInternal(m, v492, int32(3))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v492 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	F_UnlockReleaseBuffer(m, v492)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L171
	}
L104:
	;
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514)+14)))
	if v515 == int32(0) {
		goto L103
	} else {
		goto L108
	}
L105:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[3]))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v500+(v492^int32(-1))<<(uint(int32(2))%32))))
	v514 = v506
	goto L104
L106:
	;
	goto L107
L107:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_spgvacuumscan[4]))
	v514 = v508 + v492<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L108:
	;
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514)+16)))
	v520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+v518))))
	if v520&int32(2) != 0 {
		goto L103
	} else {
		goto L109
	}
L109:
	;
	if v520&int32(4) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v540 = v465
	goto L113
L111:
	;
	goto L112
L112:
	;
	v821 = int32(1)
	if base.Ui32(v488-v821) <= base.Ui32(v821) {
		goto L94
	} else {
		goto L160
	}
L113:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+6)))
	if v547 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L157
	}
L115:
	;
	goto L114
L116:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	if v804 != 0 {
		v540 = v804
		goto L113
	} else {
		goto L156
	}
L117:
	;
	v548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+2)))
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540))))
	if v548|v549<<(uint(int32(16))%32) != v488 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+4)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v514+int32(20)+v554<<(uint(int32(2))%32))))
	v561 = v514 + v558&int32(_a_F_spgvacuumscan_2)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	switch v562 & int32(3) {
	case 0:
		goto L121
	case 1:
		goto L120
	default:
		goto L115
	}
L119:
	;
	v784 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+6)) = uint8(v784)
	goto L116
L120:
	;
	v696 = v561 + int32(6)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v697 != 0 {
		goto L142
	} else {
		goto L143
	}
L121:
	;
	if v562&int32(_a_F_spgvacuumscan_7) == int32(0) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v577 = v561 + int32(base.Ui32(v562)>>(uint(int32(16))%32)) + int32(8)
	v582 = int32(0)
	goto L123
L123:
	;
	v593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	if v593 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L119
L125:
	;
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+6)))
	v684 = int32(_a_F_spgvacuumscan_8)
	v688 = v582 + int32(1)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	if base.Ui32(v688) < base.Ui32(int32(base.Ui32(v689)>>(uint(int32(3))%32))&v684) {
		v577 = v577 + v683&v684
		v582 = v688
		goto L123
	} else {
		goto L141
	}
L126:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v596 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v598 = v596
	goto L130
L128:
	;
	v639 = v472
	goto L129
L129:
	;
	v654 = F_palloc(m, int32(12))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L140
	}
L130:
	;
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+2)))
	v616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577))))
	v617 = int32(16)
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+2)))
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598))))
	if v615|v616<<(uint(v617)%32) == v620|v621<<(uint(v617)%32) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v639 = v598 + int32(8)
	goto L129
L132:
	;
	if v631 != 0 {
		goto L125
	} else {
		goto L138
	}
L133:
	;
	goto L132
L134:
	;
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598)+4)))
	if v627 == v628 {
		v631 = int32(1)
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v631 = int32(0)
	goto L133
L137:
	;
	goto L136
L138:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v598)+8))
	if v632 != 0 {
		v598 = v632
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v577)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v654)+4)) = uint16(v656)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	*(*int32)(unsafe.Add(mBase, uint32(v654))) = v658
	v660 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v654)+8)) = v660
	*(*uint8)(unsafe.Add(mBase, uint32(v654)+6)) = uint8(v660)
	*(*int32)(unsafe.Add(mBase, uint32(v639))) = v654
	goto L125
L141:
	;
	goto L124
L142:
	;
	v699 = v697
	goto L145
L143:
	;
	v740 = v472
	goto L144
L144:
	;
	v755 = F_palloc(m, int32(12))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L155
	}
L145:
	;
	v716 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696)+2)))
	v717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696))))
	v718 = int32(16)
	v721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v699)+2)))
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v699))))
	if v716|v717<<(uint(v718)%32) == v721|v722<<(uint(v718)%32) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v740 = v699 + int32(8)
	goto L144
L147:
	;
	if v732 != 0 {
		goto L119
	} else {
		goto L153
	}
L148:
	;
	goto L147
L149:
	;
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696)+4)))
	v729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v699)+4)))
	if v728 == v729 {
		v732 = int32(1)
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v732 = int32(0)
	goto L148
L152:
	;
	goto L151
L153:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	if v733 != 0 {
		v699 = v733
		goto L145
	} else {
		goto L154
	}
L154:
	;
	goto L146
L155:
	;
	v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v755)+4)) = uint16(v757)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v696)))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v759
	v761 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v755)+8)) = v761
	*(*uint8)(unsafe.Add(mBase, uint32(v755)+6)) = uint8(v761)
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = v755
	goto L119
L156:
	;
	goto L103
L157:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v809 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_3), v462)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(783), int32(_a_F_spgvacuumscan_9))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_vacuumLeafPage(m, v459, v464, v492, int32(1))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	F_vacuumRedirectAndPlaceholder(m, v464, v829, v492)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_SpGistSetLastUsedPage(m, v464, v492)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v834 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+6)) = uint8(v834)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v465)+8))
	if v836 == int32(0) {
		goto L103
	} else {
		goto L164
	}
L164:
	;
	v840 = v836
	goto L165
L165:
	;
	v857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v840)+2)))
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v840))))
	if v488 == v857|v858<<(uint(int32(16))%32) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L103
L167:
	;
	v863 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v840)+6)) = uint8(v863)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v840)+8))
	if v865 != 0 {
		v840 = v865
		goto L165
	} else {
		goto L170
	}
L170:
	;
	goto L166
L171:
	;
	v886 = v459
	v889 = v462
	v891 = v464
	v895 = v468
	v899 = v472
	v901 = v474
	v903 = v476
	goto L99
L172:
	;
	goto L96
L173:
	;
	v907 = v905
	goto L176
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = int32(0)
	v104 = v886
	v107 = v889
	v113 = v895
	v117 = v899
	v119 = v901
	v121 = v903
	goto L22
L176:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v907)+8))
	F_pfree(m, v907)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	goto L175
L178:
	;
	if v924 != 0 {
		v907 = v924
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v464)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+16)) = v951 + int32(4)
	F_errmsg_internal(m, int32(_a_F_spgvacuumscan_10), v462+int32(16))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_spgvacuumscan_4), int32(722), int32(_a_F_spgvacuumscan_9))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_SpGistUpdateMetaPage(m, v77)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)+28))
	if v970 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_FreeSpaceMapVacuum(m, v77)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	v974 = v969
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = v100
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v976)+24)) = v977
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v979)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v979)+32)) = v980
	m.G0 = v71 + int32(880)
	return
L188:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v974 = v973
	goto L187
L189:
	;
	v68 = v104
	v71 = v107
	v77 = v113
	v81 = v117
	v83 = v119
	v85 = v121
	goto L8
}
func F_spgvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int64
	_ = v243
	var v250 int32
	_ = v250
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
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
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int64
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v956 int32
	_ = v956
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(336)
	m.G0 = v27
	v31 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if int32(0) < v473 {
		goto L84
	} else {
		goto L85
	}
L2:
	;
	return int32(0)
L3:
	;
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	v37 = v35 + v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+92))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+84))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+80))
	v41 = F_get_opfamily_name(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L2
	} else {
		goto L81
	}
L7:
	;
	v45 = base.I64_extend_i32_u(v40)
	v46 = int64(0)
	v48 = F_SearchSysCacheList(m, int32(4), int32(1), v45, v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v52 = int64(0)
	v54 = F_SearchSysCacheList(m, int32(5), int32(1), v45, v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v56 = F_identify_opfamily_groups(m, v48, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v58 = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	if v59 <= int32(0) {
		v449 = v58
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v70 = v58
	v75 = v2
	v76 = v2
	v86 = v2
	v88 = v2
	goto L12
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(-64)+v75<<(uint(int32(2))%32))))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+72))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+22)))
	v100 = v98 + v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if v101 == v102 {
		v131 = v70
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v449 = v408
	goto L1
L14:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+16)))
	switch v132 - int32(1) {
	case 0:
		goto L30
	case 1, 2, 3:
		goto L25
	case 4:
		goto L26
	case 5:
		goto L27
	case 6:
		goto L28
	default:
		goto L29
	}
L15:
	;
	v104 = int32(0)
	v107 = F_errstart(m, int32(17), v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v107 == int32(0) {
		v131 = v104
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v115 = F_format_procedure(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_1), v27+int32(288))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(96), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v131 = v104
	goto L14
L22:
	;
	v433 = v75 + int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	if v433 < v434 {
		v70 = v408
		v75 = v433
		v76 = v414
		v86 = v424
		v88 = v426
		goto L12
	} else {
		goto L80
	}
L23:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L2
	} else {
		goto L76
	}
L24:
	;
	v353 = int32(0)
	v356 = F_errstart(m, int32(17), v353)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L2
	} else {
		goto L74
	}
L25:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = int64(9796820404457)
	v323 = int32(2)
	v327 = F_check_amproc_signature(m, v318, int32(2278), int32(1), v323, v323, v27+int32(240))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L72
	}
L26:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+256)) = int64(9796820404457)
	v310 = int32(2)
	v314 = F_check_amproc_signature(m, v305, int32(16), int32(1), v310, v310, v27+int32(256))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L70
	}
L27:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v86 != v290 {
		v335 = v76
		v345 = v86
		v347 = v88
		goto L24
	} else {
		goto L66
	}
L28:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v286 = F_check_amoptsproc_signature(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L2
	} else {
		goto L64
	}
L29:
	;
	v276 = int32(0)
	v279 = F_errstart(m, int32(17), v276)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L62
	}
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = int64(9796820404457)
	v140 = int32(2)
	v144 = F_check_amproc_signature(m, v135, int32(2278), int32(1), v140, v140, v27+int32(224))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v147 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+312)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v27)+332)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v27)+320)) = v147
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v154 = F_OidFunctionCall2Coll(m, v152, int32(0), base.I64_extend_i32_u(v27+int32(332)), base.I64_extend_i32_u(v27+int32(312)))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v27)+320))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v38 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v161 = v38
	goto L35
L34:
	;
	v161 = v160
	goto L35
L35:
	;
	if base.B2i32(v157 == int32(0))|base.B2i32(v157 == v161) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v168 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	v193 = v131
	v194 = v161
	goto L38
L38:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v27)+332))
	if base.B2i32(v56 == int32(0))|base.B2i32(v194 != v197) != 0 {
		goto L48
	} else {
		goto L49
	}
L39:
	;
	if v168 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v27)+320))
	v193 = int32(0)
	v194 = v191
	goto L38
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v27)+320))
	v174 = F_format_type_be(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v176 = F_format_type_be(m, v161)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+212)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v174
	F_errmsg(m, int32(_a_F_spgvalidate_4), v27+int32(208))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(130), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	if v144 != 0 {
		v408 = v193
		v414 = v194
		v424 = v160
		v426 = v156
		goto L22
	} else {
		goto L61
	}
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v200 <= int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v203 = int32(0)
	if v203 < v200 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v206 = v200
	goto L53
L52:
	;
	v206 = v203
	goto L53
L53:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v213 = int32(0)
	goto L54
L54:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v208+v213<<(uint(int32(2))%32))))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v238 != v207 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L48
L56:
	;
	v250 = v213 + int32(1)
	if v250 != v206 {
		v213 = v250
		goto L54
	} else {
		goto L60
	}
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if v240 != v241 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v237)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v237)+16)) = v243 | int64(64)
	if v144 == int32(0) {
		v335 = v194
		v345 = v160
		v347 = v156
		goto L24
	} else {
		goto L59
	}
L59:
	;
	v408 = v193
	v414 = v194
	v424 = v160
	v426 = v156
	goto L22
L60:
	;
	goto L55
L61:
	;
	v335 = v194
	v345 = v160
	v347 = v156
	goto L24
L62:
	;
	if v279 == int32(0) {
		v408 = v276
		v414 = v76
		v424 = v86
		v426 = v88
		goto L22
	} else {
		goto L63
	}
L63:
	;
	v362 = int32(184)
	v368 = v76
	v378 = v86
	v380 = v88
	v386 = int32(_a_F_spgvalidate_5)
	goto L23
L64:
	;
	if v286 == int32(0) {
		v335 = v76
		v345 = v86
		v347 = v88
		goto L24
	} else {
		goto L65
	}
L65:
	;
	v408 = v131
	v414 = v76
	v424 = v86
	v426 = v88
	goto L22
L66:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if v88 != v292 {
		v335 = v76
		v345 = v86
		v347 = v88
		goto L24
	} else {
		goto L67
	}
L67:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = v86
	v296 = int32(1)
	v301 = F_check_amproc_signature(m, v294, v76, v296, v296, v296, v27+int32(272))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v301 == int32(0) {
		v335 = v76
		v345 = v86
		v347 = v88
		goto L24
	} else {
		goto L69
	}
L69:
	;
	v408 = v131
	v414 = v76
	v424 = v86
	v426 = v88
	goto L22
L70:
	;
	if v314 == int32(0) {
		v335 = v76
		v345 = v86
		v347 = v88
		goto L24
	} else {
		goto L71
	}
L71:
	;
	v408 = v131
	v414 = v76
	v424 = v86
	v426 = v88
	goto L22
L72:
	;
	if v327 != 0 {
		v408 = v131
		v414 = v76
		v424 = v86
		v426 = v88
		goto L22
	} else {
		goto L73
	}
L73:
	;
	v335 = v76
	v345 = v86
	v347 = v88
	goto L24
L74:
	;
	if v356 == int32(0) {
		v408 = v353
		v414 = v335
		v424 = v345
		v426 = v347
		goto L22
	} else {
		goto L75
	}
L75:
	;
	v362 = int32(196)
	v368 = v335
	v378 = v345
	v380 = v347
	v386 = int32(_a_F_spgvalidate_6)
	goto L23
L76:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v391 = F_format_procedure(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v393 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+204)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v27)+200)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v27)+196)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v41
	F_errmsg(m, v386, v27+int32(192))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), v362, int32(_a_F_spgvalidate_3))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v408 = int32(0)
	v414 = v368
	v424 = v378
	v426 = v380
	goto L22
L80:
	;
	goto L13
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
	F_errmsg_internal(m, int32(_a_F_spgvalidate_7), v27)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(63), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v479 = v449
	v484 = int32(0)
	goto L87
L85:
	;
	v624 = v449
	goto L86
L86:
	;
	if v56 != 0 {
		goto L119
	} else {
		goto L120
	}
L87:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(-64)+v484<<(uint(int32(2))%32))))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+72))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+22)))
	v509 = v507 + v508
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v509)+16)))
	if base.Ui32(int32(_a_F_spgvalidate_8)) < base.Ui32((v510+int32(-64))&int32(_a_F_spgvalidate_9)) {
		v546 = v479
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v624 = v619
	goto L86
L89:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+18)))
	if v549 == int32(115) {
		v585 = v546
		v586 = int32(16)
		goto L97
	} else {
		goto L98
	}
L90:
	;
	v517 = int32(0)
	v520 = F_errstart(m, int32(17), v517)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v520 == int32(0) {
		v546 = v517
		goto L89
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	v528 = F_format_operator(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v530 = int32(*(*int16)(unsafe.Add(mBase, uint32(v509)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+188)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v27)+184)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v27)+180)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_10), v27+int32(176))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(216), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v546 = v517
	goto L89
L97:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v509)+8))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	v590 = F_check_amop_signature(m, v587, v586, v588, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L2
	} else {
		goto L109
	}
L98:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	v553 = F_get_op_rettype(m, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v509)+28))
	v556 = F_opfamily_can_sort_type(m, v555, v553)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v556 != 0 {
		v585 = v546
		v586 = v553
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v558 = int32(0)
	v561 = F_errstart(m, int32(17), v558)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	if v561 == int32(0) {
		v585 = v558
		v586 = v553
		goto L97
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	v569 = F_format_operator(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+168)) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v27)+164)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_11), v27+int32(160))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(231), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	v585 = v558
	v586 = v553
	goto L97
L108:
	;
	v621 = v484 + int32(1)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v621 < v622 {
		v479 = v619
		v484 = v621
		goto L87
	} else {
		goto L117
	}
L109:
	;
	if v590 != 0 {
		v619 = v585
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v592 = int32(0)
	v595 = F_errstart(m, int32(17), v592)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	if v595 == int32(0) {
		v619 = v592
		goto L108
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	v603 = F_format_operator(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+152)) = v603
	*(*int32)(unsafe.Add(mBase, uint32(v27)+148)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_12), v27+int32(144))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(247), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	v619 = v592
	goto L108
L117:
	;
	goto L88
L118:
	;
	F_ReleaseCatCacheList(m, v54)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L2
	} else {
		goto L198
	}
L119:
	;
	v648 = int32(0)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v649 <= v648 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	v983 = int32(0)
	v986 = F_errstart(m, int32(17), v983)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L2
	} else {
		goto L193
	}
L122:
	;
	v932 = v624
	v956 = int32(1)
	goto L124
L123:
	;
	v654 = v624
	v657 = v648
	v659 = int32(0)
	goto L125
L124:
	;
	if v956 == int32(0) {
		v1008 = v932
		goto L118
	} else {
		goto L192
	}
L125:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v678+v659<<(uint(int32(2))%32))))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	if v39 == v683 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v932 = v925
	v956 = base.B2i32(v688 == int32(0))
	goto L124
L127:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	if v685 == v39 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v688 = v657
	goto L129
L129:
	;
	v689 = *(*int64)(unsafe.Add(mBase, uint32(v682)+8))
	if v689 != int64(0) {
		v723 = v654
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v687 = v682
	goto L132
L131:
	;
	v687 = v657
	goto L132
L132:
	;
	v688 = v687
	goto L129
L133:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	if v725 != v726 {
		v925 = v723
		goto L142
	} else {
		goto L143
	}
L134:
	;
	v692 = int32(0)
	v695 = F_errstart(m, int32(17), v692)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	if v695 == int32(0) {
		v723 = v692
		goto L133
	} else {
		goto L136
	}
L136:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v703 = F_format_type_be(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L2
	} else {
		goto L138
	}
L138:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
	v706 = F_format_type_be(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+140)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v27)+136)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v27)+132)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_13), v27+int32(128))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(275), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v723 = v692
	goto L133
L142:
	;
	v927 = v659 + int32(1)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v927 < v928 {
		v654 = v925
		v657 = v688
		v659 = v927
		goto L125
	} else {
		goto L191
	}
L143:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	if v728&int32(2) != 0 {
		v760 = v723
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	if v761&int32(4) != 0 {
		v793 = v760
		goto L152
	} else {
		goto L153
	}
L145:
	;
	v731 = int32(0)
	v734 = F_errstart(m, int32(17), v731)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	if v734 == int32(0) {
		v760 = v731
		goto L144
	} else {
		goto L147
	}
L147:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v742 = F_format_type_be(m, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+124)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v27)+120)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+116)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27+int32(112))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L2
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	v760 = v731
	goto L144
L152:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	if v794&int32(8) != 0 {
		v826 = v793
		goto L160
	} else {
		goto L161
	}
L153:
	;
	v764 = int32(0)
	v767 = F_errstart(m, int32(17), v764)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	if v767 == int32(0) {
		v793 = v764
		goto L152
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v775 = F_format_type_be(m, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27+int32(96))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	v793 = v764
	goto L152
L160:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	if v827&int32(16) != 0 {
		v859 = v826
		goto L168
	} else {
		goto L169
	}
L161:
	;
	v797 = int32(0)
	v800 = F_errstart(m, int32(17), v797)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L2
	} else {
		goto L162
	}
L162:
	;
	if v800 == int32(0) {
		v826 = v797
		goto L160
	} else {
		goto L163
	}
L163:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v808 = F_format_type_be(m, v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27+int32(80))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	v826 = v797
	goto L160
L168:
	;
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	if v860&int32(32) != 0 {
		v892 = v859
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v830 = int32(0)
	v833 = F_errstart(m, int32(17), v830)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	if v833 == int32(0) {
		v859 = v830
		goto L168
	} else {
		goto L171
	}
L171:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L2
	} else {
		goto L172
	}
L172:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v841 = F_format_type_be(m, v840)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27-int32(-64))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L2
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	v859 = v830
	goto L168
L176:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+16)))
	if v893&int32(64) != 0 {
		v925 = v892
		goto L142
	} else {
		goto L184
	}
L177:
	;
	v863 = int32(0)
	v866 = F_errstart(m, int32(17), v863)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	if v866 == int32(0) {
		v892 = v863
		goto L176
	} else {
		goto L179
	}
L179:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L2
	} else {
		goto L180
	}
L180:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v874 = F_format_type_be(m, v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L2
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v874
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27+int32(48))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L2
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L2
	} else {
		goto L183
	}
L183:
	;
	v892 = v863
	goto L176
L184:
	;
	v896 = int32(0)
	v899 = F_errstart(m, int32(17), v896)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L2
	} else {
		goto L185
	}
L185:
	;
	if v899 == int32(0) {
		v925 = v896
		goto L142
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	v907 = F_format_type_be(m, v906)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v41
	F_errmsg(m, int32(_a_F_spgvalidate_14), v27+int32(32))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L2
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(296), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L2
	} else {
		goto L190
	}
L190:
	;
	v925 = v896
	goto L142
L191:
	;
	goto L126
L192:
	;
	goto L121
L193:
	;
	if v986 == int32(0) {
		v1008 = v983
		goto L118
	} else {
		goto L194
	}
L194:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(_a_F_spgvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v37 + int32(8)
	F_errmsg(m, int32(_a_F_spgvalidate_15), v27+int32(16))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L2
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_spgvalidate_2), int32(308), int32(_a_F_spgvalidate_3))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L2
	} else {
		goto L197
	}
L197:
	;
	v1008 = v983
	goto L118
L198:
	;
	F_ReleaseCatCacheList(m, v48)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L2
	} else {
		goto L199
	}
L199:
	;
	F_ReleaseCatCache(m, v31)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L2
	} else {
		goto L200
	}
L200:
	;
	m.G0 = v27 + int32(336)
	return v1008 & int32(1)
}
func F_split_part(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	v10 = m.G0
	v12 = v10 - int32(1072)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	m.G0 = v12 + int32(1072)
	return base.I64_extend_i32_u(v299)
L5:
	;
	v285 = v276 - v279
	v287 = v285 + int32(4)
	v288 = F_palloc(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L96
	}
L6:
	;
	v276 = v233
	v279 = v215 + v226
	goto L5
L7:
	;
	if v239 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L8:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1052))
	if v205 <= int32(1) {
		v276 = v212
		v279 = v206
		goto L5
	} else {
		goto L80
	}
L9:
	;
	v196 = int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v198&v196 != 0 {
		goto L77
	} else {
		goto L78
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v23 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L73
	}
L13:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v53 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v29 == int32(18) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v40 = int32(1)
	if v23&v40 != 0 {
		v52 = int32(base.Ui32(v23)>>(uint(v40)%32)) - v40
		goto L13
	} else {
		goto L23
	}
L17:
	;
	v32 = int32(16)
	goto L19
L18:
	;
	v32 = int32(0)
	goto L19
L19:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v39 = int32(4)
	goto L22
L21:
	;
	v39 = v32
	goto L22
L22:
	;
	v52 = v39
	goto L13
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L13
L24:
	;
	if v52 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v59 == int32(18) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v70 = int32(1)
	if v53&v70 != 0 {
		v82 = int32(base.Ui32(v53)>>(uint(v70)%32)) - v70
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v62 = int32(16)
	goto L30
L29:
	;
	v62 = int32(0)
	goto L30
L30:
	;
	if base.Ui32((v59-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v69 = int32(4)
	goto L33
L32:
	;
	v69 = v62
	goto L33
L33:
	;
	v82 = v69
	goto L24
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L35:
	;
	v86 = F_palloc(m, int32(4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v82 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(16)
	v299 = v86
	goto L4
L39:
	;
	switch v22 + int32(1) {
	case 0, 2:
		v299 = v15
		goto L4
	default:
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_text_position_setup(m, v15, v20, v99, v12)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v95 = F_palloc(m, int32(4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(16)
	v299 = v95
	goto L4
L44:
	;
	v102 = F_text_position_next(m, v12)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v102 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	switch v22 + int32(1) {
	case 0, 2:
		v299 = v15
		goto L4
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if int32(0) <= v22 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	v109 = F_palloc(m, int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(16)
	v299 = v109
	goto L4
L51:
	;
	v116 = int32(2)
	goto L52
L52:
	;
	v127 = F_text_position_next(m, v12)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	if v22 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v127 != 0 {
		v116 = v116 + int32(1)
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v131 = int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v133&v131 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v157 = v22 + v116 + int32(1)
	if v157 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v136 = v131
	goto L61
L60:
	;
	v136 = int32(4)
	goto L61
L61:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1052))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1056))
	v141 = v139 + v140
	v142 = v15 + v136 + v52 - v141
	v144 = v142 + int32(4)
	v145 = F_palloc(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v144 << (uint(int32(2)) % 32)
	if v142 == int32(0) {
		v299 = v145
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.MemoryCopy(m, v145+int32(4), v141, v142)
	v299 = v145
	goto L4
L64:
	;
	v161 = F_cstring_to_text(m, int32(_a_F_split_part_0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v163 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1052)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1068)) = v163
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1064)) = v167
	v169 = F_text_position_next(m, v12)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v299 = v161
	goto L4
L68:
	;
	v171 = int32(1)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v173&v171 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v176 = v171
	goto L71
L70:
	;
	v176 = int32(4)
	goto L71
L71:
	;
	v177 = v15 + v176
	if v169 == int32(0) {
		v239 = v157
		v246 = v177
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v205 = v157
	v206 = v177
	goto L8
L73:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_split_part_1), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_split_part_2), int32(3526), int32(_a_F_split_part_3))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v201 = v196
	goto L79
L78:
	;
	v201 = int32(4)
	goto L79
L79:
	;
	v205 = v22
	v206 = v15 + v201
	goto L8
L80:
	;
	v215 = v212
	v217 = v205
	goto L81
L81:
	;
	v225 = v217 - int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1056))
	v227 = F_text_position_next(m, v12)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v239 = v225
	v246 = v215 + v226
	goto L7
L83:
	;
	goto L82
L84:
	;
	if v227 == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1052))
	if base.B2i32(v217 < int32(3)) == int32(0) {
		v215 = v233
		v217 = v225
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L6
L87:
	;
	v249 = int32(1)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v251&v249 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v271 = F_palloc(m, int32(4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L95
	}
L90:
	;
	v254 = v249
	goto L92
L91:
	;
	v254 = int32(4)
	goto L92
L92:
	;
	v257 = v15 + v254 - v246 + v52
	v259 = v257 + int32(4)
	v260 = F_palloc(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v259 << (uint(int32(2)) % 32)
	if v257 == int32(0) {
		v299 = v260
		goto L4
	} else {
		goto L94
	}
L94:
	;
	base.MemoryCopy(m, v260+int32(4), v246, v257)
	v299 = v260
	goto L4
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = int32(16)
	v299 = v271
	goto L4
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v287 << (uint(int32(2)) % 32)
	if v285 == int32(0) {
		v299 = v288
		goto L4
	} else {
		goto L97
	}
L97:
	;
	base.MemoryCopy(m, v288+int32(4), v279, v285)
	v299 = v288
	goto L4
}
func F_sq(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	v6 = base.F64_mul(l2, l2)
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v6
	v9 = base.F64_mul(l2, float64(1.34217729e+08))
	v11 = base.F64_add(v9, base.F64_sub(l2, v9))
	v12 = base.F64_sub(l2, v11)
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(base.F64_mul(v12, v12), base.F64_add(base.F64_mul(base.F64_add(v11, v11), v12), base.F64_sub(base.F64_mul(v11, v11), v6)))
	return
}
func F_sqlfunction_receive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v5 != 0 {
		v6 = F_ExecFilterJunk(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			F_tuplestore_puttupleslot(m, v10, v6)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
		if v16&int32(2) != 0 {
			v19 = F_ExecFilterJunk(m, v4, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
				m.T0[v22].(func(*base.Module, int32))(m, v19)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			}
		} else {
			return int32(1)
		}
	}
}
func F_sscanf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	v11 = m.G0
	v12 = int32(144)
	v13 = v11 - v12
	m.G0 = v13
	base.MemoryFill(m, v13, int32(0), v12)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_sscanf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = l0
	v24 = F_vfscanf(m, v13, l1, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		m.G0 = v13 + int32(144)
		m.G0 = v8 + int32(16)
		return v24
	}
}
func F_stop_repack_decoding_worker(m *base.Module) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v5 != 0 {
			F_TerminateBackgroundWorker(m, v5)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
				v10 = v9
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v11 != 0 {
					F_shm_mq_detach(m, v11)
					mBase = m.M
					v13 = m.ExcPending
					if v13 != 0 {
						return
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
						F_ConditionVariableCancelSleep(m)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v22 != 0 {
								v23 = int32(_a_F_stop_repack_decoding_worker_0)
								v25 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v25 + int32(1)
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								v30 = F_WaitForBackgroundWorkerShutdown(m, v29)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									v32 = int32(_a_F_stop_repack_decoding_worker_0)
									v34 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v34 - int32(1)
									if v30 == int32(3) {
										F_errstart_cold(m, int32(22), int32(0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											F_errcode(m, int32(16908741))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_stop_repack_decoding_worker_1), int32(0))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_stop_repack_decoding_worker_2), int32(3959), int32(_a_F_stop_repack_decoding_worker_3))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
										v42 = v41
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
										if v43 != 0 {
											F_dsm_detach(m, v43)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
												*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
												v50 = v47
												F_pfree(m, v50)
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
													return
												}
											}
										} else {
											v50 = v42
											F_pfree(m, v50)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
												return
											}
										}
									}
								}
							} else {
								v42 = v21
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
								if v43 != 0 {
									F_dsm_detach(m, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
										v50 = v47
										F_pfree(m, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
											return
										}
									}
								} else {
									v50 = v42
									F_pfree(m, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
										return
									}
								}
							}
						}
					}
				} else {
					F_ConditionVariableCancelSleep(m)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v22 != 0 {
							v23 = int32(_a_F_stop_repack_decoding_worker_0)
							v25 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v25 + int32(1)
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v30 = F_WaitForBackgroundWorkerShutdown(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								v32 = int32(_a_F_stop_repack_decoding_worker_0)
								v34 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v34 - int32(1)
								if v30 == int32(3) {
									F_errstart_cold(m, int32(22), int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_errcode(m, int32(16908741))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_stop_repack_decoding_worker_1), int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_stop_repack_decoding_worker_2), int32(3959), int32(_a_F_stop_repack_decoding_worker_3))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
									v42 = v41
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
									if v43 != 0 {
										F_dsm_detach(m, v43)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
											*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
											v50 = v47
											F_pfree(m, v50)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
												return
											}
										}
									} else {
										v50 = v42
										F_pfree(m, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
											return
										}
									}
								}
							}
						} else {
							v42 = v21
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
							if v43 != 0 {
								F_dsm_detach(m, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
									v50 = v47
									F_pfree(m, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
										return
									}
								}
							} else {
								v50 = v42
								F_pfree(m, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v10 = v4
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			if v11 != 0 {
				F_shm_mq_detach(m, v11)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(0)
					F_ConditionVariableCancelSleep(m)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						if v22 != 0 {
							v23 = int32(_a_F_stop_repack_decoding_worker_0)
							v25 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v25 + int32(1)
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v30 = F_WaitForBackgroundWorkerShutdown(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								v32 = int32(_a_F_stop_repack_decoding_worker_0)
								v34 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v34 - int32(1)
								if v30 == int32(3) {
									F_errstart_cold(m, int32(22), int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_errcode(m, int32(16908741))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_stop_repack_decoding_worker_1), int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_stop_repack_decoding_worker_2), int32(3959), int32(_a_F_stop_repack_decoding_worker_3))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
									v42 = v41
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
									if v43 != 0 {
										F_dsm_detach(m, v43)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
											*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
											v50 = v47
											F_pfree(m, v50)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
												return
											}
										}
									} else {
										v50 = v42
										F_pfree(m, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
											return
										}
									}
								}
							}
						} else {
							v42 = v21
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
							if v43 != 0 {
								F_dsm_detach(m, v43)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
									v50 = v47
									F_pfree(m, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
										return
									}
								}
							} else {
								v50 = v42
								F_pfree(m, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
									return
								}
							}
						}
					}
				}
			} else {
				F_ConditionVariableCancelSleep(m)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
					if v22 != 0 {
						v23 = int32(_a_F_stop_repack_decoding_worker_0)
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
						*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v25 + int32(1)
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						v30 = F_WaitForBackgroundWorkerShutdown(m, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v32 = int32(_a_F_stop_repack_decoding_worker_0)
							v34 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[1])) = v34 - int32(1)
							if v30 == int32(3) {
								F_errstart_cold(m, int32(22), int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									F_errcode(m, int32(16908741))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_stop_repack_decoding_worker_1), int32(0))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_stop_repack_decoding_worker_2), int32(3959), int32(_a_F_stop_repack_decoding_worker_3))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
								v42 = v41
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
								if v43 != 0 {
									F_dsm_detach(m, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
										v50 = v47
										F_pfree(m, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
											return
										}
									}
								} else {
									v50 = v42
									F_pfree(m, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
										return
									}
								}
							}
						}
					} else {
						v42 = v21
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						if v43 != 0 {
							F_dsm_detach(m, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
								v50 = v47
								F_pfree(m, v50)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
									return
								}
							}
						} else {
							v50 = v42
							F_pfree(m, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_stop_repack_decoding_worker[0])) = int32(0)
								return
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_storeOperators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v21 = F_table_open(m, int32(2602), int32(3))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L61
	}
L4:
	;
	F_relation_close(m, v21, int32(3))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L60
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v25 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = base.I64_extend_i32_u(l2)
	v39 = int32(0)
	goto L7
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v39<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+12)))
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+16)))
	v52 = int64(*(*int16)(unsafe.Add(mBase, uint32(v48)+8)))
	v53 = F_SearchSysCacheExists(m, int32(4), v29, v50, v51, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v57 = v17 - int32(-64)
	v58 = int32(0)
	base.MemoryFill(m, v57, v58, int32(72))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+56)) = uint8(v58)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = int64(0)
	v67 = F_GetNewOidWithIndex(m, v21, int32(2756), int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	if v53 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = base.I64_extend_i32_u(v67)
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v72
	v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v74
	v76 = int64(*(*int16)(unsafe.Add(mBase, uint32(v48)+8)))
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = int64(111)
	goto L17
L16:
	;
	v79 = int64(115)
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v76
	v82 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v82
	v85 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+20)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v90 = F_heap_form_tuple(m, v87, v57, v17+int32(48))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_CatalogTupleInsert(m, v21, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(2602)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(2617)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v103
	v108 = v17 + int32(36)
	v110 = v17 + int32(24)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+24)))
	if v113 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v114 = int32(110)
	goto L23
L22:
	;
	v114 = int32(97)
	goto L23
L23:
	;
	F_recordDependencyOn(m, v108, v110, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+25)))
	if v119 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v120 = int32(2753)
	goto L27
L26:
	;
	v120 = int32(2616)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v122
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+24)))
	if v128 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v129 = int32(105)
	goto L30
L29:
	;
	v129 = int32(97)
	goto L30
L30:
	;
	F_recordDependencyOn(m, v108, v110, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v133 = F_typeDepNeeded(m, v132, v48)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v133 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(1247)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v137
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+24)))
	if v143 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v148 == v149 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v144 = int32(110)
	goto L38
L37:
	;
	v144 = int32(97)
	goto L38
L38:
	;
	F_recordDependencyOn(m, v108, v110, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v172 != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v151 = F_typeDepNeeded(m, v148, v48)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v151 == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(1247)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v157
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+24)))
	if v167 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v168 = int32(110)
	goto L46
L45:
	;
	v168 = int32(97)
	goto L46
L46:
	;
	F_recordDependencyOn(m, v17+int32(36), v17+int32(24), v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(2753)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+24)))
	if v184 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_storeOperators[0]))
	if v189 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v185 = int32(110)
	goto L53
L52:
	;
	v185 = int32(97)
	goto L53
L53:
	;
	F_recordDependencyOn(m, v17+int32(36), v17+int32(24), v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v191 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2602), v67, v191, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v196 = v39 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v196 < v197 {
		v39 = v196
		goto L7
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	goto L8
L60:
	;
	m.G0 = v17 + int32(144)
	return
L61:
	;
	F_errcode(m, int32(_a_F_storeOperators_0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v228 = F_format_type_be(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v231 = F_format_type_be(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v233 = F_NameListToString(m, l0)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v226
	F_errmsg(m, int32(_a_F_storeOperators_1), v17)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_storeOperators_2), int32(1507), int32(_a_F_storeOperators_3))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_store_flush_position(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[0]))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)))
	if v7 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v10 == int32(4) {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v15
			v18 = F_palloc(m, int32(24))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = l1
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3]))
				if v23 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4]))
					v30 = v25
				} else {
					v27 = int32(_a_F_store_flush_position_0)
					*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3])) = v27
					v30 = v27
				}
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v30
				v32 = int32(_a_F_store_flush_position_0)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v18
				*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4])) = v18
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[5]))
				*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v39
				return
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v15
		v18 = F_palloc(m, int32(24))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = l1
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3]))
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4]))
				v30 = v25
			} else {
				v27 = int32(_a_F_store_flush_position_0)
				*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[3])) = v27
				v30 = v27
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v30
			v32 = int32(_a_F_store_flush_position_0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v18
			*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[4])) = v18
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[5]))
			*(*int32)(unsafe.Add(mBase, _c_F_store_flush_position[2])) = v39
			return
		}
	}
}
func F_strchr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v7 = l1 & int32(255)
	if v7 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v99 == l1&int32(255) {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v97 = v87
	goto L1
L3:
	;
	v77 = v72
	goto L20
L4:
	;
	v72 = v64
	goto L3
L5:
	;
	if l0&int32(3) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v62 = F_strlen(m, l0)
	mBase = m.M
	v97 = v62 + l0
	goto L1
L8:
	;
	v10 = l0
	goto L11
L9:
	;
	v24 = l0
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v33 = int32(-2139062144)
	if (int32(16843008)-v30|v30)&v33 != v33 {
		v64 = v24
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v7 == v15) != 0 {
		v87 = v10
		goto L2
	} else {
		goto L13
	}
L12:
	;
	v24 = v21
	goto L10
L13:
	;
	v21 = v10 + int32(1)
	if v21&int32(3) != 0 {
		v10 = v21
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v39 = v24
	v41 = v30
	goto L16
L16:
	;
	v45 = v41 ^ v7*int32(16843009)
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 != v48 {
		v64 = v39
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v72 = v54
	goto L3
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v54 = v39 + int32(4)
	v58 = int32(-2139062144)
	if (v52|(int32(16843008)-v52))&v58 == v58 {
		v39 = v54
		v41 = v52
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == int32(0) {
		v87 = v77
		goto L2
	} else {
		goto L22
	}
L21:
	;
	v87 = v77
	goto L2
L22:
	;
	if v79 != l1&int32(255) {
		v77 = v77 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v103 = v97
	goto L26
L25:
	;
	v103 = int32(0)
	goto L26
L26:
	;
	return v103
}
func F_strdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v4 = F_strlen(m, l0)
	mBase = m.M
	v6 = v4 + int32(1)
	v7 = F_emscripten_builtin_malloc(m, v6)
	mBase = m.M
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v6) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v7
L5:
	;
	if v6 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v18 = v7 + v6
	if (v7^l0)&int32(3) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	base.MemoryCopy(m, v7, l0, v6)
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v150) < base.Ui32(v18) {
		goto L45
	} else {
		goto L46
	}
L12:
	;
	if v7&int32(3) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v18) < base.Ui32(int32(4)) {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v54 = v18 & int32(-4)
	if base.Ui32(v18) < base.Ui32(int32(64)) {
		v104 = v48
		v105 = v49
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v48 = l0
	v49 = v7
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v6 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = l0
	v49 = v7
	goto L15
L20:
	;
	goto L21
L21:
	;
	v31 = l0
	v32 = v7
	goto L22
L22:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v36)
	v38 = int32(1)
	v39 = v31 + v38
	v41 = v32 + v38
	if v41&int32(3) == int32(0) {
		v48 = v39
		v49 = v41
		goto L15
	} else {
		goto L24
	}
L23:
	;
	v48 = v39
	v49 = v41
	goto L15
L24:
	;
	if base.Ui32(v41) < base.Ui32(v18) {
		v31 = v39
		v32 = v41
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if base.Ui32(v54) <= base.Ui32(v105) {
		v149 = v104
		v150 = v105
		goto L11
	} else {
		goto L32
	}
L27:
	;
	v58 = v54 + int32(-64)
	if base.Ui32(v58) < base.Ui32(v49) {
		v104 = v48
		v105 = v49
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v61 = v48
	v62 = v49
	goto L29
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+28)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+32)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+36)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+40)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+48)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+52)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+56)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v61)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+60)) = v96
	v98 = int32(-64)
	v99 = v61 - v98
	v101 = v62 - v98
	if base.Ui32(v101) <= base.Ui32(v58) {
		v61 = v99
		v62 = v101
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v104 = v99
	v105 = v101
	goto L26
L31:
	;
	goto L30
L32:
	;
	v111 = v104
	v112 = v105
	goto L33
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v116
	v118 = int32(4)
	v119 = v111 + v118
	v121 = v112 + v118
	if base.Ui32(v121) < base.Ui32(v54) {
		v111 = v119
		v112 = v121
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v149 = v119
	v150 = v121
	goto L11
L35:
	;
	goto L34
L36:
	;
	v149 = l0
	v150 = v7
	goto L11
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v6) < base.Ui32(int32(4)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v149 = l0
	v150 = v7
	goto L11
L40:
	;
	goto L41
L41:
	;
	v130 = l0
	v131 = v7
	goto L42
L42:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)) = uint8(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+2)) = uint8(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+3)) = uint8(v141)
	v143 = int32(4)
	v144 = v130 + v143
	v146 = v131 + v143
	if base.Ui32(v146) <= base.Ui32(v18-int32(4)) {
		v130 = v144
		v131 = v146
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v149 = v144
	v150 = v146
	goto L11
L44:
	;
	goto L43
L45:
	;
	v156 = v149
	v157 = v150
	goto L48
L46:
	;
	goto L47
L47:
	;
	goto L4
L48:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v161)
	v163 = int32(1)
	v166 = v157 + v163
	if v166 != v18 {
		v156 = v156 + v163
		v157 = v166
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
}
func F_strfold_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	v8 = int32(0)
	v10 = F_convert_case(m, l0, l1, l2, l3, int32(3), v7, v8, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_strlower_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v6 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	v10 = F_convert_case(m, l0, l1, l2, l3, v6, v7, v6, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_strstr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
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
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
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
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v542 int32
	_ = v542
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v728 int32
	_ = v728
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v756 int32
	_ = v756
	var v773 int32
	_ = v773
	v3 = int32(0)
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v15 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	goto L3
L3:
	;
	v19 = F___strchrnul(m, l0, v15)
	mBase = m.M
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 == v15&int32(255) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v773
L5:
	;
	if v25 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L9
	}
L6:
	;
	v25 = v19
	goto L8
L7:
	;
	v25 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v28 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return v25
L11:
	;
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v32 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v39 = int32(0)
	v40 = base.B2i32(v38 != v39)
	if v38 == v39 {
		v82 = v25
		v85 = v40
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	if v99 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L27
	}
L17:
	;
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v44 = int32(8)
	v46 = v43<<(uint(v44)%32) | v38
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v51 = v47 | v48<<(uint(v44)%32)
	if v46 == v51 {
		v82 = v25
		v85 = v40
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v56 = v25 + int32(1)
	v60 = v46
	goto L20
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v70 = int32(0)
	v71 = base.B2i32(v69 != v70)
	if v69 == v70 {
		v82 = v56
		v85 = v71
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v82 = v56
	v85 = v71
	goto L17
L22:
	;
	v80 = v60<<(uint(int32(8))%32)&int32(_a_F_strstr_0) | v69
	if v80 != v51 {
		v56 = v56 + int32(1)
		v60 = v80
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v97 = v82
	goto L26
L25:
	;
	v97 = int32(0)
	goto L26
L26:
	;
	return v97
L27:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	if v102 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v106 = v25 + int32(2)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v108 = int32(0)
	if v107 == v108 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
	if v180 == int32(0) {
		v773 = v3
		goto L4
	} else {
		goto L42
	}
L31:
	;
	if v163 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v162 = v106
	v163 = base.B2i32(v107 != v108)
	goto L31
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v113 = int32(16)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v116 = int32(24)
	v119 = int32(8)
	v121 = v112<<(uint(v113)%32) | v115<<(uint(v116)%32) | v107<<(uint(v119)%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	v132 = v122<<(uint(v113)%32) | v125<<(uint(v116)%32) | v129<<(uint(v119)%32)
	if v121 == v132 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v137 = v106
	v138 = v121
	goto L35
L35:
	;
	v149 = v137 + int32(1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v151 = int32(0)
	v152 = base.B2i32(v150 != v151)
	if v150 == v151 {
		v162 = v149
		v163 = v152
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v162 = v149
	v163 = v152
	goto L31
L37:
	;
	v157 = (v150 | v138) << (uint(int32(8)) % 32)
	if v157 != v132 {
		v137 = v149
		v138 = v157
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v178 = v162 - int32(2)
	goto L41
L40:
	;
	v178 = int32(0)
	goto L41
L41:
	;
	return v178
L42:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v183 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v187 = v25 + int32(3)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
	v189 = int32(0)
	if v188 == v189 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	v263 = int32(0)
	v264 = m.G0
	v266 = v264 - int32(1056)
	m.G0 = v266
	v268 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1048)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1040)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1032)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v266)+1024)) = v268
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v276 == v263 {
		goto L62
	} else {
		goto L63
	}
L46:
	;
	if v249 != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v246 = v187
	v249 = base.B2i32(v188 != v189)
	goto L46
L48:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v197 = int32(24)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v201 = int32(8)
	v204 = v193<<(uint(int32(16))%32) | v196<<(uint(v197)%32) | v200<<(uint(v201)%32) | v188
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v206 = int32(16711935)
	v214 = base.I32_rotr(v205&v206, v201) | base.I32_rotr(v205, v197)&v206
	if v204 == v214 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v219 = v187
	v220 = v204
	goto L50
L50:
	;
	v231 = v219 + int32(1)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	v233 = int32(0)
	v234 = base.B2i32(v232 != v233)
	if v232 == v233 {
		v246 = v231
		v249 = v234
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v246 = v231
	v249 = v234
	goto L46
L52:
	;
	v239 = v220<<(uint(int32(8))%32) | v232
	if v239 != v214 {
		v219 = v231
		v220 = v239
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v261 = v246 - int32(3)
	goto L56
L55:
	;
	v261 = int32(0)
	goto L56
L56:
	;
	return v261
L57:
	;
	m.G0 = v266 + int32(1056)
	v773 = v756
	goto L4
L58:
	;
	v439 = int32(1)
	v443 = base.B2i32(base.Ui32(v430+v439) < base.Ui32(v429+v439))
	if base.Ui32(v430+v439) < base.Ui32(v429+v439) {
		goto L97
	} else {
		goto L98
	}
L59:
	;
	v342 = int32(1)
	v344 = v321
	v345 = v342
	v346 = v263
	v349 = v322
	v354 = v342
	goto L70
L60:
	;
	v756 = int32(0)
	goto L57
L61:
	;
	v425 = v325
	v429 = int32(-1)
	v430 = v330
	v431 = v331
	v432 = int32(1)
	goto L58
L62:
	;
	v325 = int32(1)
	v330 = int32(-1)
	v331 = v3
	goto L61
L63:
	;
	goto L64
L64:
	;
	v282 = v276
	v287 = v3
	goto L65
L65:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v287))))
	if v296 == int32(0) {
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v321 = int32(1)
	v322 = int32(-1)
	if base.Ui32(v321) < base.Ui32(v305) {
		goto L59
	} else {
		goto L69
	}
L67:
	;
	v304 = int32(1)
	v305 = v287 + v304
	*(*int32)(unsafe.Add(mBase, uint32(v266+v282&int32(255)<<(uint(int32(2))%32)))) = v305
	v313 = v266 + int32(1024) + int32(base.Ui32(v282)>>(uint(int32(3))%32))&int32(28)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v314 | v304<<(uint(v282)%32)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+l1))))
	if v320 != 0 {
		v282 = v320
		v287 = v305
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v325 = v321
	v330 = v322
	v331 = v305
	goto L61
L70:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349+l1+v345))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+l1))))
	if v360 == v362 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v384 = int32(1)
	v387 = int32(0)
	v388 = v384
	v389 = v384
	v391 = int32(-1)
	v394 = v384
	goto L83
L72:
	;
	v380 = v379 + v376
	if base.Ui32(v380) < base.Ui32(v305) {
		v344 = v380
		v345 = v379
		v346 = v376
		v349 = v377
		v354 = v378
		goto L70
	} else {
		goto L82
	}
L73:
	;
	if v345 == v354 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v362) < base.Ui32(v360) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v376 = v346 + v354
	v377 = v349
	v378 = v354
	v379 = int32(1)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v376 = v346
	v377 = v349
	v378 = v354
	v379 = v345 + int32(1)
	goto L72
L79:
	;
	v376 = v344
	v377 = v349
	v378 = v344 - v349
	v379 = int32(1)
	goto L72
L80:
	;
	goto L81
L81:
	;
	v372 = int32(1)
	v376 = v346 + v372
	v377 = v346
	v378 = v372
	v379 = v372
	goto L72
L82:
	;
	goto L71
L83:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391+l1+v388))))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+l1))))
	if v403 == v405 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v425 = v378
	v429 = v420
	v430 = v377
	v431 = v305
	v432 = v421
	goto L58
L85:
	;
	v423 = v422 + v419
	if base.Ui32(v423) < base.Ui32(v305) {
		v387 = v419
		v388 = v422
		v389 = v423
		v391 = v420
		v394 = v421
		goto L83
	} else {
		goto L95
	}
L86:
	;
	if v388 == v394 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(v403) < base.Ui32(v405) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v419 = v387 + v394
	v420 = v391
	v421 = v394
	v422 = int32(1)
	goto L85
L90:
	;
	goto L91
L91:
	;
	v419 = v387
	v420 = v391
	v421 = v394
	v422 = v388 + int32(1)
	goto L85
L92:
	;
	v419 = v389
	v420 = v391
	v421 = v389 - v391
	v422 = int32(1)
	goto L85
L93:
	;
	goto L94
L94:
	;
	v415 = int32(1)
	v419 = v387 + v415
	v420 = v387
	v421 = v415
	v422 = v415
	goto L85
L95:
	;
	goto L84
L96:
	;
	v524 = v431 | int32(63)
	v527 = int32(0)
	v529 = v25
	v530 = v25
	goto L127
L97:
	;
	v444 = v432
	goto L99
L98:
	;
	v444 = v425
	goto L99
L99:
	;
	v445 = l1 + v444
	if base.Ui32(v430+v439) < base.Ui32(v429+v439) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v446 = v429
	goto L102
L101:
	;
	v446 = v430
	goto L102
L102:
	;
	v448 = v446 + int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v448) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	if v510 != 0 {
		goto L121
	} else {
		goto L122
	}
L104:
	;
	v510 = int32(0)
	goto L103
L105:
	;
	v484 = v479
	v485 = v480
	v486 = v481
	goto L115
L106:
	;
	if (l1|v445)&int32(3) != 0 {
		v479 = l1
		v480 = v445
		v481 = v448
		goto L105
	} else {
		goto L109
	}
L107:
	;
	v472 = l1
	v473 = v445
	v474 = v448
	goto L108
L108:
	;
	if v474 == int32(0) {
		goto L104
	} else {
		goto L114
	}
L109:
	;
	v456 = l1
	v457 = v445
	v458 = v448
	goto L110
L110:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	if v461 != v462 {
		v479 = v456
		v480 = v457
		v481 = v458
		goto L105
	} else {
		goto L112
	}
L111:
	;
	v472 = v467
	v473 = v465
	v474 = v469
	goto L108
L112:
	;
	v464 = int32(4)
	v465 = v457 + v464
	v467 = v456 + v464
	v469 = v458 - v464
	if base.Ui32(int32(3)) < base.Ui32(v469) {
		v456 = v467
		v457 = v465
		v458 = v469
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v479 = v472
	v480 = v473
	v481 = v474
	goto L105
L115:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v489 == v490 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v510 = v489 - v490
	goto L103
L117:
	;
	v492 = int32(1)
	v497 = v486 - v492
	if v497 != 0 {
		v484 = v484 + v492
		v485 = v485 + v492
		v486 = v497
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	goto L104
L121:
	;
	v513 = v431 + (v446 ^ int32(-1))
	if base.Ui32(v513) < base.Ui32(v446) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v521 = v444
	v522 = v431 - v444
	goto L96
L124:
	;
	v515 = v446
	goto L126
L125:
	;
	v515 = v513
	goto L126
L126:
	;
	v521 = v515 + int32(1)
	v522 = int32(0)
	goto L96
L127:
	;
	if base.Ui32(v431) <= base.Ui32(v530-v529) {
		v657 = v530
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v658 = int32(0)
	v661 = v529 + v431
	v662 = int32(1)
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661-v662))))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v266+int32(1024)+int32(base.Ui32(v664)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v670)>>(uint(v664)%32))&v662 == v658 {
		v527 = v658
		v529 = v661
		v530 = v657
		goto L127
	} else {
		goto L161
	}
L130:
	;
	v542 = int32(0)
	if base.B2i32(v530&int32(3) == v542)|base.B2i32(v524 == v542) != 0 {
		v574 = v530
		v576 = v524
		v577 = base.B2i32(v524 != v542)
		goto L134
	} else {
		goto L135
	}
L131:
	;
	if v648 != 0 {
		goto L156
	} else {
		goto L157
	}
L132:
	;
	v648 = int32(0)
	goto L131
L133:
	;
	v626 = v619
	v628 = v621
	goto L150
L134:
	;
	if v577 == int32(0) {
		goto L132
	} else {
		goto L141
	}
L135:
	;
	v557 = v530
	v559 = v524
	goto L136
L136:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v562 == int32(0) {
		v619 = v557
		v621 = v559
		goto L133
	} else {
		goto L138
	}
L137:
	;
	v574 = v569
	v576 = v565
	v577 = v567
	goto L134
L138:
	;
	v564 = int32(1)
	v565 = v559 - v564
	v566 = int32(0)
	v567 = base.B2i32(v565 != v566)
	v569 = v557 + v564
	if v569&int32(3) == v566 {
		v574 = v569
		v576 = v565
		v577 = v567
		goto L134
	} else {
		goto L139
	}
L139:
	;
	if v565 != 0 {
		v557 = v569
		v559 = v565
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v582 = int32(0)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if base.B2i32(v582 == v583)|base.B2i32(base.Ui32(v576) < base.Ui32(int32(4))) == v582 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v592 = v574
	v594 = v576
	goto L145
L143:
	;
	v612 = v574
	v614 = v576
	goto L144
L144:
	;
	if v614 == int32(0) {
		goto L132
	} else {
		goto L149
	}
L145:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	v599 = v598 ^ int32(0)
	v602 = int32(-2139062144)
	if (int32(16843008)-v599|v599)&v602 != v602 {
		v619 = v592
		v621 = v594
		goto L133
	} else {
		goto L147
	}
L146:
	;
	v612 = v607
	v614 = v609
	goto L144
L147:
	;
	v606 = int32(4)
	v607 = v592 + v606
	v609 = v594 - v606
	if base.Ui32(int32(3)) < base.Ui32(v609) {
		v592 = v607
		v594 = v609
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v619 = v612
	v621 = v614
	goto L133
L150:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if int32(0) == v631 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L132
L152:
	;
	v648 = v626
	goto L131
L153:
	;
	goto L154
L154:
	;
	v633 = int32(1)
	v636 = v628 - v633
	if v636 != 0 {
		v626 = v626 + v633
		v628 = v636
		goto L150
	} else {
		goto L155
	}
L155:
	;
	goto L151
L156:
	;
	v650 = v648
	goto L158
L157:
	;
	v650 = v530 + v524
	goto L158
L158:
	;
	if v648 == int32(0) {
		v657 = v650
		goto L129
	} else {
		goto L159
	}
L159:
	;
	if base.Ui32(v648-v529) < base.Ui32(v431) {
		v756 = v542
		goto L57
	} else {
		goto L160
	}
L160:
	;
	v657 = v650
	goto L129
L161:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v266+v664<<(uint(int32(2))%32))))
	if v679 != v431 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v681 = v431 - v679
	if base.Ui32(v527) < base.Ui32(v681) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	if base.Ui32(v527) < base.Ui32(v448) {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	v683 = v681
	goto L167
L166:
	;
	v683 = v527
	goto L167
L167:
	;
	v527 = int32(0)
	v529 = v529 + v683
	v530 = v657
	goto L127
L168:
	;
	v527 = int32(0)
	v529 = v529 + (v693 - v446)
	v530 = v657
	goto L127
L169:
	;
	v687 = v448
	goto L171
L170:
	;
	v687 = v527
	goto L171
L171:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v687))))
	if v689 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v692 = v689
	v693 = v687
	goto L175
L173:
	;
	goto L174
L174:
	;
	v728 = v448
	goto L179
L175:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v693))))
	if v705 != v692&int32(255) {
		goto L168
	} else {
		goto L177
	}
L176:
	;
	goto L174
L177:
	;
	v710 = v693 + int32(1)
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v710))))
	if v712 != 0 {
		v692 = v712
		v693 = v710
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	if base.Ui32(v728) <= base.Ui32(v527) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v527 = v522
	v529 = v529 + v521
	v530 = v657
	goto L127
L181:
	;
	v756 = v529
	goto L57
L182:
	;
	goto L183
L183:
	;
	v743 = v728 - int32(1)
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v743))))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v743))))
	if v745 == v747 {
		v728 = v743
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
}
func F_strtitle_libc_mb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
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
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	v10 = l3 + int32(1)
	if base.Ui32(v10) < base.Ui32(int32(536870912)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v15 = F_palloc_mul(m, int32(4), v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L144
	}
L4:
	;
	return int32(0)
L5:
	;
	F_char2wchar(m, v15, v10, l2, l3, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = int32(0)
	v27 = v21
	v29 = v22
	goto L10
L8:
	;
	v59 = v21
	goto L9
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_strtitle_libc_mb[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66*int32(28))+uint32(_c_F_strtitle_libc_mb[1])))
	goto L23
L10:
	;
	if v26 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v59 = v51
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v27<<(uint(int32(2))%32)))) = v39
	if base.Ui32(int32(10)) <= base.Ui32(v39-int32(48)) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v36 = F_casemap(m, v29, int32(0))
	mBase = m.M
	goto L16
L14:
	;
	goto L15
L15:
	;
	v38 = F_casemap(m, v29, int32(1))
	mBase = m.M
	goto L17
L16:
	;
	v39 = v36
	goto L12
L17:
	;
	v39 = v38
	goto L12
L18:
	;
	v51 = v27 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15+v51<<(uint(int32(2))%32))))
	if v55 != 0 {
		v26 = v49
		v27 = v51
		v29 = v55
		goto L10
	} else {
		goto L22
	}
L19:
	;
	v45 = F_iswalpha(m, v39)
	mBase = m.M
	v49 = base.B2i32(v45 != int32(0))
	goto L21
L20:
	;
	v49 = int32(1)
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L11
L23:
	;
	v74 = v71*v59 + int32(1)
	v75 = F_palloc(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v74 == int32(0) {
		v461 = int32(0)
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32(v461+int32(1)) <= base.Ui32(l1) {
		goto L136
	} else {
		goto L137
	}
L26:
	;
	if v13 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = int32(0)
	v87 = m.G0
	v88 = int32(16)
	v89 = v87 - v88
	m.G0 = v89
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v15
	v93 = v89 + int32(12)
	v94 = m.G0
	v96 = v94 - v88
	m.G0 = v96
	if v75 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	goto L29
L29:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_strtitle_libc_mb[2]))
	if v13 != 0 {
		goto L74
	} else {
		goto L75
	}
L30:
	;
	v461 = v248
	goto L25
L31:
	;
	v252 = int32(16)
	m.G0 = v96 + v252
	m.G0 = v89 + v252
	goto L30
L32:
	;
	v248 = v74 - v233
	goto L31
L33:
	;
	if v170 != 0 {
		goto L58
	} else {
		goto L59
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v129 = v74
	v130 = v75
	v134 = v128
	goto L47
L35:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v74) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 == int32(0) {
		v248 = v81
		goto L31
	} else {
		goto L39
	}
L38:
	;
	v170 = v74
	v171 = v75
	goto L33
L39:
	;
	v104 = v101
	v105 = v100
	v107 = v81
	goto L40
L40:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v104) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v248 = v127
	goto L31
L42:
	;
	v116 = int32(-1)
	v119 = F_wcrtomb(m, v96+int32(12), v104)
	mBase = m.M
	if v119 == v116 {
		v248 = v116
		goto L31
	} else {
		goto L45
	}
L43:
	;
	v122 = int32(1)
	goto L44
L44:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v127 = v122 + v107
	if v124 != 0 {
		v104 = v124
		v105 = v105 + int32(4)
		v107 = v127
		goto L40
	} else {
		goto L46
	}
L45:
	;
	v122 = v119
	goto L44
L46:
	;
	goto L41
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if base.Ui32(v138-int32(128)) <= base.Ui32(int32(-128)) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v170 = v160
	v171 = v163
	goto L33
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v166 = v164 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v166
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v129 = v160
		v130 = v163
		v134 = v166
		goto L47
	} else {
		goto L57
	}
L50:
	;
	if v138 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v138)
	v156 = int32(1)
	v160 = v129 - v156
	v163 = v130 + v156
	goto L49
L53:
	;
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v145)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v145
	v233 = v129
	goto L32
L54:
	;
	goto L55
L55:
	;
	v149 = int32(-1)
	v150 = F_wcrtomb(m, v130, v138)
	mBase = m.M
	if v150 == v149 {
		v248 = v149
		goto L31
	} else {
		goto L56
	}
L56:
	;
	v160 = v129 - v150
	v163 = v130 + v150
	goto L49
L57:
	;
	goto L48
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v180 = v170
	v181 = v171
	v183 = v179
	goto L61
L59:
	;
	goto L60
L60:
	;
	v248 = v74
	goto L31
L61:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if base.Ui32(v189-int32(128)) <= base.Ui32(int32(-128)) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L60
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v222 = v220 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v222
	if v216 != 0 {
		v180 = v216
		v181 = v219
		v183 = v222
		goto L61
	} else {
		goto L72
	}
L64:
	;
	if v189 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v189)
	v212 = int32(1)
	v216 = v180 - v212
	v219 = v181 + v212
	goto L63
L67:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v196)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v196
	v233 = v180
	goto L32
L68:
	;
	goto L69
L69:
	;
	v200 = int32(-1)
	v203 = F_wcrtomb(m, v96+int32(12), v189)
	mBase = m.M
	if v203 == v200 {
		v248 = v200
		goto L31
	} else {
		goto L70
	}
L70:
	;
	if base.Ui32(v180) < base.Ui32(v203) {
		v233 = v180
		goto L32
	} else {
		goto L71
	}
L71:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v208 = F_wcrtomb(m, v181, v207)
	mBase = m.M
	v216 = v180 - v203
	v219 = v181 + v203
	goto L63
L72:
	;
	goto L62
L73:
	;
	v271 = int32(0)
	v277 = m.G0
	v278 = int32(16)
	v279 = v277 - v278
	m.G0 = v279
	*(*int32)(unsafe.Add(mBase, uint32(v279)+12)) = v15
	v283 = v279 + int32(12)
	v284 = m.G0
	v286 = v284 - v278
	m.G0 = v286
	if v75 != 0 {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	if v13 == int32(-1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v260 == int32(_a_F_strtitle_libc_mb_0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v265 = int32(_a_F_strtitle_libc_mb_0)
	goto L79
L78:
	;
	v265 = v13
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtitle_libc_mb[2])) = v265
	goto L76
L80:
	;
	v270 = int32(-1)
	goto L82
L81:
	;
	v270 = v260
	goto L82
L82:
	;
	goto L73
L83:
	;
	if v270 != 0 {
		goto L127
	} else {
		goto L128
	}
L84:
	;
	v442 = int32(16)
	m.G0 = v286 + v442
	m.G0 = v279 + v442
	goto L83
L85:
	;
	v438 = v74 - v423
	goto L84
L86:
	;
	if v360 != 0 {
		goto L111
	} else {
		goto L112
	}
L87:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v319 = v74
	v320 = v75
	v324 = v318
	goto L100
L88:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v74) {
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v291 == int32(0) {
		v438 = v271
		goto L84
	} else {
		goto L92
	}
L91:
	;
	v360 = v74
	v361 = v75
	goto L86
L92:
	;
	v294 = v291
	v295 = v290
	v297 = v271
	goto L93
L93:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v294) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v438 = v317
	goto L84
L95:
	;
	v306 = int32(-1)
	v309 = F_wcrtomb(m, v286+int32(12), v294)
	mBase = m.M
	if v309 == v306 {
		v438 = v306
		goto L84
	} else {
		goto L98
	}
L96:
	;
	v312 = int32(1)
	goto L97
L97:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v317 = v312 + v297
	if v314 != 0 {
		v294 = v314
		v295 = v295 + int32(4)
		v297 = v317
		goto L93
	} else {
		goto L99
	}
L98:
	;
	v312 = v309
	goto L97
L99:
	;
	goto L94
L100:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	if base.Ui32(v328-int32(128)) <= base.Ui32(int32(-128)) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v360 = v350
	v361 = v353
	goto L86
L102:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v356 = v354 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v356
	if base.Ui32(int32(3)) < base.Ui32(v350) {
		v319 = v350
		v320 = v353
		v324 = v356
		goto L100
	} else {
		goto L110
	}
L103:
	;
	if v328 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v328)
	v346 = int32(1)
	v350 = v319 - v346
	v353 = v320 + v346
	goto L102
L106:
	;
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v320))) = uint8(v335)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v335
	v423 = v319
	goto L85
L107:
	;
	goto L108
L108:
	;
	v339 = int32(-1)
	v340 = F_wcrtomb(m, v320, v328)
	mBase = m.M
	if v340 == v339 {
		v438 = v339
		goto L84
	} else {
		goto L109
	}
L109:
	;
	v350 = v319 - v340
	v353 = v320 + v340
	goto L102
L110:
	;
	goto L101
L111:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v370 = v360
	v371 = v361
	v373 = v369
	goto L114
L112:
	;
	goto L113
L113:
	;
	v438 = v74
	goto L84
L114:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	if base.Ui32(v379-int32(128)) <= base.Ui32(int32(-128)) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L113
L116:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v412 = v410 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v412
	if v406 != 0 {
		v370 = v406
		v371 = v409
		v373 = v412
		goto L114
	} else {
		goto L125
	}
L117:
	;
	if v379 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v379)
	v402 = int32(1)
	v406 = v370 - v402
	v409 = v371 + v402
	goto L116
L120:
	;
	v386 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v386)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v386
	v423 = v370
	goto L85
L121:
	;
	goto L122
L122:
	;
	v390 = int32(-1)
	v393 = F_wcrtomb(m, v286+int32(12), v379)
	mBase = m.M
	if v393 == v390 {
		v438 = v390
		goto L84
	} else {
		goto L123
	}
L123:
	;
	if base.Ui32(v370) < base.Ui32(v393) {
		v423 = v370
		goto L85
	} else {
		goto L124
	}
L124:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v398 = F_wcrtomb(m, v371, v397)
	mBase = m.M
	v406 = v370 - v393
	v409 = v371 + v393
	goto L116
L125:
	;
	goto L115
L126:
	;
	v461 = v438
	goto L25
L127:
	;
	if v270 == int32(-1) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	goto L133
L130:
	;
	v455 = int32(_a_F_strtitle_libc_mb_0)
	goto L132
L131:
	;
	v455 = v270
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strtitle_libc_mb[2])) = v455
	goto L129
L133:
	;
	goto L135
L135:
	;
	goto L126
L136:
	;
	if v461 != 0 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L138
L138:
	;
	F_pfree(m, v15)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L142
	}
L139:
	;
	base.MemoryCopy(m, l0, v75, v461)
	goto L141
L140:
	;
	goto L141
L141:
	;
	v467 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v461))) = uint8(v467)
	goto L138
L142:
	;
	F_pfree(m, v75)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	return v461
L144:
	;
	F_errcode(m, int32(_a_F_strtitle_libc_mb_1))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_strtitle_libc_mb_2), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_strtitle_libc_mb_3), int32(650), int32(_a_F_strtitle_libc_mb_4))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strtof(m *base.Module, l0 int32, l1 int32) float32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_strtox_1(m, v7, l0, l1, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float32(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v23 = m.G0
		v25 = v23 - int32(32)
		m.G0 = v25
		v28 = v15 & int64(281474976710655)
		v32 = int64(base.Ui64(v15)>>(uint(int64(48))%64)) & int64(32767)
		v33 = base.I32_wrap_i64(v32)
		if base.Ui32(v33-int32(_a_F_strtof_0)) <= base.Ui32(int32(253)) {
			v40 = base.I32_wrap_i64(int64(base.Ui64(v28) >> (uint(int64(25)) % 64)))
			v44 = v15 & int64(33554431)
			v45 = int64(16777216)
			if v44 == v45 {
				v49 = base.B2i32(v14 == int64(0))
			} else {
				v49 = base.B2i32(base.Ui64(v44) < base.Ui64(v45))
			}
			if v49 == int32(0) {
				v62 = v40 + int32(1)
			} else {
				if v14|(v44^int64(16777216)) != int64(0) {
					v62 = v40
				} else {
					v62 = v40&int32(1) + v40
				}
			}
			v65 = base.B2i32(base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v62))
			if base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v62) {
				v66 = int32(0)
			} else {
				v66 = v62
			}
			if base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v62) {
				v69 = int32(-16255)
			} else {
				v69 = int32(-16256)
			}
			v150 = v66
			v151 = v69 + v33
		} else {
			if base.B2i32(v14|v28 == int64(0))|base.B2i32(v32 != int64(32767)) == int32(0) {
				v150 = base.I32_wrap_i64(int64(base.Ui64(v28)>>(uint(int64(25))%64))) | int32(_a_F_strtof_2)
				v151 = int32(255)
			} else {
				if base.Ui32(int32(_a_F_strtof_3)) < base.Ui32(v33) {
					v150 = int32(0)
					v151 = int32(255)
				} else {
					v91 = base.B2i32(v32 == int64(0))
					if v32 == int64(0) {
						v92 = int32(_a_F_strtof_4)
					} else {
						v92 = int32(_a_F_strtof_0)
					}
					v93 = v92 - v33
					if int32(112) < v93 {
						v96 = int32(0)
						v150 = v96
						v151 = v96
					} else {
						if v32 == int64(0) {
							v100 = v28
						} else {
							v100 = v28 | int64(281474976710656)
						}
						if v33 != v92 {
							F___ashlti3(m, v25+int32(16), v14, v100, int32(128)-v93)
							mBase = m.M
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
							v113 = base.B2i32(v108|v109 != int64(0))
						} else {
							v113 = int32(0)
						}
						F___lshrti3(m, v25, v14, v100, v93)
						mBase = m.M
						v115 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
						v118 = base.I32_wrap_i64(int64(base.Ui64(v115) >> (uint(int64(25)) % 64)))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
						v121 = v119 | base.I64_extend_i32_u(v113)
						v125 = v115 & int64(33554431)
						v126 = int64(16777216)
						if v125 == v126 {
							v130 = base.B2i32(v121 == int64(0))
						} else {
							v130 = base.B2i32(base.Ui64(v125) < base.Ui64(v126))
						}
						if v130 == int32(0) {
							v143 = v118 + int32(1)
						} else {
							if v121|(v125^int64(16777216)) != int64(0) {
								v143 = v118
							} else {
								v143 = v118&int32(1) + v118
							}
						}
						v147 = base.B2i32(base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v143))
						if base.Ui32(int32(_a_F_strtof_1)) < base.Ui32(v143) {
							v148 = v143 ^ int32(_a_F_strtof_5)
						} else {
							v148 = v143
						}
						v150 = v148
						v151 = v147
					}
				}
			}
		}
		m.G0 = v25 + int32(32)
		m.G0 = v7 + int32(16)
		return base.F32_reinterpret_i32(base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(32))%64)))&int32(-2147483648) | v151<<(uint(int32(23))%32) | v150)
	}
}
func F_strtol(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(2147483648))
	return base.I32_wrap_i64(v5)
}
func F_strupper_libc_mb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	v9 = l3 + int32(1)
	if base.Ui32(v9) < base.Ui32(int32(536870912)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v14 = F_palloc_mul(m, int32(4), v9)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L135
	}
L4:
	;
	return int32(0)
L5:
	;
	F_char2wchar(m, v14, v9, l2, l3, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v20
	v28 = v22
	goto L10
L8:
	;
	v45 = v20
	goto L9
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_strupper_libc_mb[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(28))+uint32(_c_F_strupper_libc_mb[1])))
	goto L14
L10:
	;
	v34 = F_casemap(m, v28, int32(1))
	mBase = m.M
	goto L12
L11:
	;
	v45 = v37
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v26<<(uint(int32(2))%32)))) = v34
	v37 = v26 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14+v37<<(uint(int32(2))%32))))
	if v41 != 0 {
		v26 = v37
		v28 = v41
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v59 = v56*v45 + int32(1)
	v60 = F_palloc(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v59 == int32(0) {
		v446 = v20
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui32(v446+int32(1)) <= base.Ui32(l1) {
		goto L127
	} else {
		goto L128
	}
L17:
	;
	if v12 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v66 = int32(0)
	v72 = m.G0
	v73 = int32(16)
	v74 = v72 - v73
	m.G0 = v74
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v14
	v78 = v74 + int32(12)
	v79 = m.G0
	v81 = v79 - v73
	m.G0 = v81
	if v60 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L20
L20:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_strupper_libc_mb[2]))
	if v12 != 0 {
		goto L65
	} else {
		goto L66
	}
L21:
	;
	v446 = v233
	goto L16
L22:
	;
	v237 = int32(16)
	m.G0 = v81 + v237
	m.G0 = v74 + v237
	goto L21
L23:
	;
	v233 = v59 - v218
	goto L22
L24:
	;
	if v155 != 0 {
		goto L49
	} else {
		goto L50
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v114 = v59
	v115 = v60
	v119 = v113
	goto L38
L26:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 == int32(0) {
		v233 = v66
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v155 = v59
	v156 = v60
	goto L24
L30:
	;
	v89 = v86
	v90 = v85
	v92 = v66
	goto L31
L31:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v89) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v233 = v112
	goto L22
L33:
	;
	v101 = int32(-1)
	v104 = F_wcrtomb(m, v81+int32(12), v89)
	mBase = m.M
	if v104 == v101 {
		v233 = v101
		goto L22
	} else {
		goto L36
	}
L34:
	;
	v107 = int32(1)
	goto L35
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v112 = v107 + v92
	if v109 != 0 {
		v89 = v109
		v90 = v90 + int32(4)
		v92 = v112
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v107 = v104
	goto L35
L37:
	;
	goto L32
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if base.Ui32(v123-int32(128)) <= base.Ui32(int32(-128)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v155 = v145
	v156 = v148
	goto L24
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v151 = v149 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v151
	if base.Ui32(int32(3)) < base.Ui32(v145) {
		v114 = v145
		v115 = v148
		v119 = v151
		goto L38
	} else {
		goto L48
	}
L41:
	;
	if v123 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v123)
	v141 = int32(1)
	v145 = v114 - v141
	v148 = v115 + v141
	goto L40
L44:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v130)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v130
	v218 = v114
	goto L23
L45:
	;
	goto L46
L46:
	;
	v134 = int32(-1)
	v135 = F_wcrtomb(m, v115, v123)
	mBase = m.M
	if v135 == v134 {
		v233 = v134
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v145 = v114 - v135
	v148 = v115 + v135
	goto L40
L48:
	;
	goto L39
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v165 = v155
	v166 = v156
	v168 = v164
	goto L52
L50:
	;
	goto L51
L51:
	;
	v233 = v59
	goto L22
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if base.Ui32(v174-int32(128)) <= base.Ui32(int32(-128)) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L51
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v207 = v205 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v207
	if v201 != 0 {
		v165 = v201
		v166 = v204
		v168 = v207
		goto L52
	} else {
		goto L63
	}
L55:
	;
	if v174 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v174)
	v197 = int32(1)
	v201 = v165 - v197
	v204 = v166 + v197
	goto L54
L58:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v181)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v181
	v218 = v165
	goto L23
L59:
	;
	goto L60
L60:
	;
	v185 = int32(-1)
	v188 = F_wcrtomb(m, v81+int32(12), v174)
	mBase = m.M
	if v188 == v185 {
		v233 = v185
		goto L22
	} else {
		goto L61
	}
L61:
	;
	if base.Ui32(v165) < base.Ui32(v188) {
		v218 = v165
		goto L23
	} else {
		goto L62
	}
L62:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v193 = F_wcrtomb(m, v166, v192)
	mBase = m.M
	v201 = v165 - v188
	v204 = v166 + v188
	goto L54
L63:
	;
	goto L53
L64:
	;
	v256 = int32(0)
	v262 = m.G0
	v263 = int32(16)
	v264 = v262 - v263
	m.G0 = v264
	*(*int32)(unsafe.Add(mBase, uint32(v264)+12)) = v14
	v268 = v264 + int32(12)
	v269 = m.G0
	v271 = v269 - v263
	m.G0 = v271
	if v60 != 0 {
		goto L79
	} else {
		goto L80
	}
L65:
	;
	if v12 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v245 == int32(_a_F_strupper_libc_mb_0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v250 = int32(_a_F_strupper_libc_mb_0)
	goto L70
L69:
	;
	v250 = v12
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strupper_libc_mb[2])) = v250
	goto L67
L71:
	;
	v255 = int32(-1)
	goto L73
L72:
	;
	v255 = v245
	goto L73
L73:
	;
	goto L64
L74:
	;
	if v255 != 0 {
		goto L118
	} else {
		goto L119
	}
L75:
	;
	v427 = int32(16)
	m.G0 = v271 + v427
	m.G0 = v264 + v427
	goto L74
L76:
	;
	v423 = v59 - v408
	goto L75
L77:
	;
	if v345 != 0 {
		goto L102
	} else {
		goto L103
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v304 = v59
	v305 = v60
	v309 = v303
	goto L91
L79:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v276 == int32(0) {
		v423 = v256
		goto L75
	} else {
		goto L83
	}
L82:
	;
	v345 = v59
	v346 = v60
	goto L77
L83:
	;
	v279 = v276
	v280 = v275
	v282 = v256
	goto L84
L84:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v279) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v423 = v302
	goto L75
L86:
	;
	v291 = int32(-1)
	v294 = F_wcrtomb(m, v271+int32(12), v279)
	mBase = m.M
	if v294 == v291 {
		v423 = v291
		goto L75
	} else {
		goto L89
	}
L87:
	;
	v297 = int32(1)
	goto L88
L88:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v302 = v297 + v282
	if v299 != 0 {
		v279 = v299
		v280 = v280 + int32(4)
		v282 = v302
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v297 = v294
	goto L88
L90:
	;
	goto L85
L91:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if base.Ui32(v313-int32(128)) <= base.Ui32(int32(-128)) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v345 = v335
	v346 = v338
	goto L77
L93:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v341 = v339 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v341
	if base.Ui32(int32(3)) < base.Ui32(v335) {
		v304 = v335
		v305 = v338
		v309 = v341
		goto L91
	} else {
		goto L101
	}
L94:
	;
	if v313 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v305))) = uint8(v313)
	v331 = int32(1)
	v335 = v304 - v331
	v338 = v305 + v331
	goto L93
L97:
	;
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v305))) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v320
	v408 = v304
	goto L76
L98:
	;
	goto L99
L99:
	;
	v324 = int32(-1)
	v325 = F_wcrtomb(m, v305, v313)
	mBase = m.M
	if v325 == v324 {
		v423 = v324
		goto L75
	} else {
		goto L100
	}
L100:
	;
	v335 = v304 - v325
	v338 = v305 + v325
	goto L93
L101:
	;
	goto L92
L102:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v355 = v345
	v356 = v346
	v358 = v354
	goto L105
L103:
	;
	goto L104
L104:
	;
	v423 = v59
	goto L75
L105:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if base.Ui32(v364-int32(128)) <= base.Ui32(int32(-128)) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L104
L107:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v397 = v395 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v397
	if v391 != 0 {
		v355 = v391
		v356 = v394
		v358 = v397
		goto L105
	} else {
		goto L116
	}
L108:
	;
	if v364 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L110
L110:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v364)
	v387 = int32(1)
	v391 = v355 - v387
	v394 = v356 + v387
	goto L107
L111:
	;
	v371 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v371
	v408 = v355
	goto L76
L112:
	;
	goto L113
L113:
	;
	v375 = int32(-1)
	v378 = F_wcrtomb(m, v271+int32(12), v364)
	mBase = m.M
	if v378 == v375 {
		v423 = v375
		goto L75
	} else {
		goto L114
	}
L114:
	;
	if base.Ui32(v355) < base.Ui32(v378) {
		v408 = v355
		goto L76
	} else {
		goto L115
	}
L115:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v383 = F_wcrtomb(m, v356, v382)
	mBase = m.M
	v391 = v355 - v378
	v394 = v356 + v378
	goto L107
L116:
	;
	goto L106
L117:
	;
	v446 = v423
	goto L16
L118:
	;
	if v255 == int32(-1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	goto L124
L121:
	;
	v440 = int32(_a_F_strupper_libc_mb_0)
	goto L123
L122:
	;
	v440 = v255
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strupper_libc_mb[2])) = v440
	goto L120
L124:
	;
	goto L126
L126:
	;
	goto L117
L127:
	;
	if v446 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	F_pfree(m, v14)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L133
	}
L130:
	;
	base.MemoryCopy(m, l0, v60, v446)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v452 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v446))) = uint8(v452)
	goto L129
L133:
	;
	F_pfree(m, v60)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	return v446
L135:
	;
	F_errcode(m, int32(_a_F_strupper_libc_mb_1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_strupper_libc_mb_2), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_strupper_libc_mb_3), int32(737), int32(_a_F_strupper_libc_mb_4))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_strupper_libc_sb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	if base.Ui32(l1) < base.Ui32(l3+int32(1)) {
	} else {
		if l3 != 0 {
			base.MemoryCopy(m, l0, l2, l3)
		} else {
		}
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+l3))) = uint8(v12)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v14 == v12 {
		} else {
			v17 = l0
			v19 = v14
			for {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
				if v22 == int32(1) {
					if base.Ui32((v19-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
						v61 = v19 - int32(32)
						*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v61)
					} else {
						if int32(0) <= base.I32_extend8_s(v19) {
						} else {
							if base.B2i32(base.Ui32(v19&int32(255)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
								if base.Ui32(v44-int32(97)) < base.Ui32(int32(26)) {
									v51 = v44 & int32(95)
								} else {
									v51 = v44
								}
								v61 = v51
								*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v61)
							}
						}
					}
				} else {
					v53 = v19 & int32(255)
					if base.Ui32(v53-int32(97)) < base.Ui32(int32(26)) {
						v60 = v53 & int32(95)
					} else {
						v60 = v53
					}
					v61 = v60
					*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v61)
				}
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				if v65 != 0 {
					v17 = v17 + int32(1)
					v19 = v65
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return l3
}
func F_subcolorcvec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v230 int32
	_ = v230
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v300 int32
	_ = v300
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int64
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int64
	_ = v448
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int64
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v754 int32
	_ = v754
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v817 int32
	_ = v817
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v869 int32
	_ = v869
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1031 int32
	_ = v1031
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1076 int32
	_ = v1076
	var v1096 int32
	_ = v1096
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1144 int32
	_ = v1144
	var v1160 int32
	_ = v1160
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v25 = int32(_a_F_subcolorcvec_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v27 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v1160 + int32(16)
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v35 = v27
	v36 = v30
	goto L5
L3:
	;
	goto L4
L4:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v81 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_subcoloronechr(m, l0, v50, l2, l3, v22+int32(14))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v55 != 0 {
		v1160 = v22
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v58 = int32(1)
	if v58 < v35 {
		v35 = v35 - v58
		v36 = v36 + int32(4)
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v100 = v84
	v101 = v81
	goto L14
L12:
	;
	v662 = v22
	goto L13
L13:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v667 < int32(0) {
		v1160 = v662
		goto L1
	} else {
		goto L133
	}
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if base.Ui32(int32(2047)) < base.Ui32(v105) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v662 = v22
	goto L13
L16:
	;
	if base.Ui32(v300) < base.Ui32(v104) {
		goto L63
	} else {
		goto L64
	}
L17:
	;
	v300 = v105
	goto L16
L18:
	;
	goto L19
L19:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)))
	if base.Ui32(v104) < base.Ui32(v105) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)) = uint16(v279)
	v300 = v280
	goto L16
L21:
	;
	v279 = v108
	v280 = v105
	goto L20
L22:
	;
	goto L23
L23:
	;
	v110 = int32(2047)
	if base.Ui32(v110) <= base.Ui32(v104) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v113 = v110
	goto L26
L25:
	;
	v113 = v104
	goto L26
L26:
	;
	v119 = v105
	v120 = v108
	goto L27
L27:
	;
	v133 = F_subcolor(m, v24, v119)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L29
	}
L28:
	;
	v279 = v257
	v280 = v272
	goto L20
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v135 != 0 {
		v1160 = v22
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v136 = int32(_a_F_subcolorcvec_0)
	v137 = v133 & v136
	if v137 != v120&v136 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_subcolorcvec[0]))
	if v143 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v257 = v120
	goto L33
L33:
	;
	v272 = v119 + int32(1)
	if base.Ui32(v119) < base.Ui32(v113) {
		v119 = v272
		v120 = v257
		goto L27
	} else {
		goto L61
	}
L34:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v146 <= v147 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L36
L38:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v250 != 0 {
		v1160 = v22
		goto L1
	} else {
		goto L60
	}
L39:
	;
	F_createarc(m, v141, int32(112), v133, l2, l3)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L59
	}
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v149 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v179 == int32(0) {
		goto L39
	} else {
		goto L51
	}
L43:
	;
	v156 = v149
	goto L44
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	if v171 != l3 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L39
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v178 != 0 {
		v156 = v178
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	if v173 != v137 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v175 == int32(112) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L45
L51:
	;
	v186 = v179
	goto L52
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v201 != l2 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L39
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v186)+24))
	if v208 != 0 {
		v186 = v208
		goto L52
	} else {
		goto L58
	}
L55:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+4)))
	if v203 != v137 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	if v205 == int32(112) {
		goto L38
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	goto L53
L59:
	;
	goto L38
L60:
	;
	v257 = v133
	goto L33
L61:
	;
	goto L28
L62:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v641 != 0 {
		v1160 = v22
		goto L1
	} else {
		goto L131
	}
L63:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+84))
	v316 = int32(1)
	v320 = F_palloc_mul_extended(m, int32(12), v315<<(uint(v316)%32)|v316)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v300 != v104 {
		goto L62
	} else {
		goto L129
	}
L66:
	;
	if v320 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+24)) = int32(101)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	if v328 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v314)+88))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v314)+84))
	if v333 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v330 = v328
	goto L72
L71:
	;
	v330 = int32(12)
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+12)) = v330
	goto L62
L73:
	;
	if v385 <= v379 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v378 = v332
	v379 = int32(0)
	v385 = v333
	goto L73
L75:
	;
	goto L76
L76:
	;
	v337 = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	if base.Ui32(v300) <= base.Ui32(v338) {
		v378 = v332
		v379 = v337
		v385 = v333
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v345 = v337
	v346 = v332
	goto L78
L78:
	;
	v359 = int32(12)
	v361 = v320 + v345*v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+8)) = v362
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v346)))
	*(*int64)(unsafe.Add(mBase, uint32(v361))) = v364
	v367 = v346 + v359
	v369 = v345 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v314)+84))
	if v370 <= v369 {
		v378 = v367
		v379 = v369
		v385 = v370
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v378 = v367
	v379 = v369
	v385 = v370
	goto L73
L80:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v346)+16))
	if base.Ui32(v372) < base.Ui32(v300) {
		v345 = v369
		v346 = v367
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if base.Ui32(v526) <= base.Ui32(v104) {
		goto L114
	} else {
		goto L115
	}
L83:
	;
	v523 = v378
	v524 = v379
	v525 = v379
	v526 = v300
	v530 = v385
	goto L82
L84:
	;
	goto L85
L85:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	if base.Ui32(v104) < base.Ui32(v394) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v523 = v378
	v524 = v379
	v525 = v379
	v526 = v300
	v530 = v385
	goto L82
L87:
	;
	goto L88
L88:
	;
	v402 = v378
	v403 = v379
	v404 = v379
	v405 = v300
	v411 = v394
	goto L89
L89:
	;
	if base.Ui32(v405) < base.Ui32(v411) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v523 = v510
	v524 = v503
	v525 = v514
	v526 = v512
	v530 = v515
	goto L82
L91:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	F_subcoloronerow(m, l0, v500, l2, l3, v22+int32(14))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L7
	} else {
		goto L111
	}
L92:
	;
	v469 = v320 + v464*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v465
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if base.Ui32(v104) < base.Ui32(v471) {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	v455 = v320 + v403*int32(12)
	v456 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v405 - v456
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = v411
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+8)) = v460
	v464 = v403 + v456
	v465 = v405
	goto L92
L94:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if base.Ui32(v104) < base.Ui32(v441) {
		v464 = v438
		v465 = v439
		goto L92
	} else {
		goto L101
	}
L95:
	;
	v420 = v320 + v403*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v405
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = v422 - int32(1)
	v427 = F_newhicolorrow(m, v314, int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L7
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if base.Ui32(v411) < base.Ui32(v405) {
		goto L93
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420)+8)) = v427
	F_subcoloronerow(m, l0, v427, l2, l3, v22+int32(14))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v438 = v403 + int32(1)
	v439 = v436
	goto L94
L100:
	;
	v438 = v403
	v439 = v405
	goto L94
L101:
	;
	v445 = v320 + v438*int32(12)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v446
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v402)))
	*(*int64)(unsafe.Add(mBase, uint32(v445))) = v448
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v500 = v450
	v503 = v438 + int32(1)
	goto L91
L102:
	;
	v473 = v104
	goto L104
L103:
	;
	v473 = v471
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+4)) = v473
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v476 = F_newhicolorrow(m, v314, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+8)) = v476
	v480 = v464 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if base.Ui32(v481) <= base.Ui32(v104) {
		v500 = v476
		v503 = v480
		goto L91
	} else {
		goto L106
	}
L106:
	;
	v485 = v320 + v480*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v104 + int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+4)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	if base.Ui32(v490) < base.Ui32(v465) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v492 = F_newhicolorrow(m, v314, v489)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L7
	} else {
		goto L110
	}
L108:
	;
	v494 = v489
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+8)) = v494
	v500 = v476
	v503 = v464 + int32(2)
	goto L91
L110:
	;
	v494 = v492
	goto L109
L111:
	;
	v510 = v402 + int32(12)
	v511 = int32(1)
	v512 = v504 + v511
	v514 = v404 + v511
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v314)+84))
	if v515 <= v514 {
		v523 = v510
		v524 = v503
		v525 = v514
		v526 = v512
		v530 = v515
		goto L82
	} else {
		goto L112
	}
L112:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	if base.Ui32(v517) <= base.Ui32(v104) {
		v402 = v510
		v403 = v503
		v404 = v514
		v405 = v512
		v411 = v517
		goto L89
	} else {
		goto L113
	}
L113:
	;
	goto L90
L114:
	;
	v541 = v320 + v524*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v541)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v526
	v545 = F_newhicolorrow(m, v314, int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L7
	} else {
		goto L117
	}
L115:
	;
	v555 = v524
	v557 = v530
	goto L116
L116:
	;
	if v525 < v557 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v541)+8)) = v545
	F_subcoloronerow(m, l0, v545, l2, l3, v22+int32(14))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v314)+84))
	v555 = v524 + int32(1)
	v557 = v552
	goto L116
L119:
	;
	v563 = v523
	v564 = v555
	v565 = v525
	goto L122
L120:
	;
	v598 = v555
	goto L121
L121:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v314)+88))
	if v612 != 0 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v578 = int32(12)
	v580 = v320 + v564*v578
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+8)) = v581
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v563)))
	*(*int64)(unsafe.Add(mBase, uint32(v580))) = v583
	v587 = int32(1)
	v588 = v564 + v587
	v590 = v565 + v587
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v314)+84))
	if v590 < v591 {
		v563 = v563 + v578
		v564 = v588
		v565 = v590
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v598 = v588
	goto L121
L124:
	;
	goto L123
L125:
	;
	F_pfree(m, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L7
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+84)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v314)+88)) = v320
	goto L62
L128:
	;
	goto L127
L129:
	;
	F_subcoloronechr(m, l0, v104, l2, l3, v22+int32(14))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	goto L62
L131:
	;
	v644 = int32(1)
	if v644 < v101 {
		v100 = v100 + int32(8)
		v101 = v101 - v644
		goto L14
	} else {
		goto L132
	}
L132:
	;
	goto L15
L133:
	;
	v671 = v24 + int32(28)
	v674 = v671 + v667<<(uint(int32(2))%32)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	if v675 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = v678
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v24)+96))
	v684 = base.I32_div_s(int32(2147483647), v681<<(uint(int32(1))%32))
	if v684 <= v678 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v869 = v675
	goto L136
L136:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	if v878 <= int32(0) {
		v1160 = v662
		goto L1
	} else {
		goto L164
	}
L137:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v853 != 0 {
		v1160 = v662
		goto L1
	} else {
		goto L163
	}
L138:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v686)+24)) = int32(101)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v689)+12))
	if v690 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v699 = F_repalloc_mul_extended(m, v694, int32(2), v678*v681<<(uint(int32(1))%32))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L7
	} else {
		goto L144
	}
L141:
	;
	v692 = v690
	goto L143
L142:
	;
	v692 = int32(12)
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v689)+12)) = v692
	goto L137
L144:
	;
	if v699 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v703)+24)) = int32(101)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+12))
	if v707 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v699
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v715 = v713 - int32(1)
	if int32(0) <= v715 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v709 = v707
	goto L150
L149:
	;
	v709 = int32(12)
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+12)) = v709
	goto L137
L151:
	;
	v722 = v715
	v723 = v712
	goto L154
L152:
	;
	v817 = v712
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v817 << (uint(int32(1)) % 32)
	goto L137
L154:
	;
	if int32(0) < v723 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v817 = v794
	goto L153
L156:
	;
	v739 = v722 * v723
	v742 = v699 + v739<<(uint(int32(2))%32)
	v743 = int32(1)
	v754 = int32(0)
	goto L159
L157:
	;
	v794 = v723
	goto L158
L158:
	;
	if int32(0) < v722 {
		v722 = v722 - int32(1)
		v723 = v794
		goto L154
	} else {
		goto L162
	}
L159:
	;
	v769 = int32(1)
	v770 = v754 << (uint(v769) % 32)
	v773 = int32(*(*int16)(unsafe.Add(mBase, uint32(v770+(v699+v739<<(uint(v743)%32))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v742+v723<<(uint(v743)%32)+v770))) = uint16(v773)
	*(*uint16)(unsafe.Add(mBase, uint32(v770+v742))) = uint16(v773)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v780 = v777 + v773*int32(24)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v780)+4)) = v781 + v769
	v786 = v754 + v769
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	if v786 < v787 {
		v754 = v786
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v794 = v787
	goto L158
L161:
	;
	goto L160
L162:
	;
	goto L155
L163:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v671+v854<<(uint(int32(2))%32))))
	v869 = v858
	goto L136
L164:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v888 = v881
	v890 = v878
	v891 = v882
	v899 = int32(0)
	goto L165
L165:
	;
	v903 = int32(0)
	if v903 < v888 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v1160 = v662
	goto L1
L167:
	;
	v911 = v903
	v913 = v891
	goto L170
L168:
	;
	v1128 = v888
	v1130 = v890
	v1131 = v891
	goto L169
L169:
	;
	v1144 = v899 + int32(1)
	if v1144 < v1130 {
		v888 = v1128
		v890 = v1130
		v891 = v1131
		v899 = v1144
		goto L165
	} else {
		goto L214
	}
L170:
	;
	if v911&v869 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	v1128 = v1121
	v1130 = v1123
	v1131 = v1118
	goto L169
L172:
	;
	v1118 = v913 + int32(2)
	v1120 = v911 + int32(1)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	if v1120 < v1121 {
		v911 = v1120
		v913 = v1118
		goto L170
	} else {
		goto L213
	}
L173:
	;
	v928 = int32(*(*int16)(unsafe.Add(mBase, uint32(v913))))
	v930 = v928 * int32(24)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v932 = v930 + v931
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v932)+8)))
	if v933 != int32(_a_F_subcolorcvec_0) {
		v954 = v933
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)+12))
	if v957 != 0 {
		v980 = int32(_a_F_subcolorcvec_0)
		goto L181
	} else {
		goto L182
	}
L175:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v932)+4))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v932)))
	if v936+v937 == int32(1) {
		v954 = v928
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v941 = F_newcolor(m, v24)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	if v941 == int32(-1) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v954 = int32(_a_F_subcolorcvec_0)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v946+v930)+8)) = uint16(v941)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v949+v941*int32(24))+8)) = uint16(v941)
	v954 = v941
	goto L174
L181:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v981 != 0 {
		v1160 = v662
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v958 = int32(_a_F_subcolorcvec_0)
	if v928&v958 == v954&v958 {
		v980 = v928
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v964 = v963 + v930
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)+4))
	v966 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v964)+4)) = v965 - v966
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v973 = v969 + base.I32_extend16_s(v954)*int32(24)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v973)+4)) = v974 + v966
	*(*uint16)(unsafe.Add(mBase, uint32(v913))) = uint16(v954)
	v980 = v954
	goto L181
L184:
	;
	v983 = v980 & int32(_a_F_subcolorcvec_0)
	v984 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v662)+14)))
	if v983 == v984 {
		goto L172
	} else {
		goto L185
	}
L185:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v988 = *(*int32)(unsafe.Add(mBase, _c_F_subcolorcvec[0]))
	if v988 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L7
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v991 <= v992 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L188
L190:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1096 != 0 {
		v1160 = v662
		goto L1
	} else {
		goto L212
	}
L191:
	;
	F_createarc(m, v986, int32(112), base.I32_extend16_s(v980), l2, l3)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L7
	} else {
		goto L211
	}
L192:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v994 == int32(0) {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v1024 == int32(0) {
		goto L191
	} else {
		goto L203
	}
L195:
	;
	v1001 = v994
	goto L196
L196:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	if v1016 != l3 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L191
L198:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+16))
	if v1023 != 0 {
		v1001 = v1023
		goto L196
	} else {
		goto L202
	}
L199:
	;
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1001)+4)))
	if v1018 != v983 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	if v1020 == int32(112) {
		goto L190
	} else {
		goto L201
	}
L201:
	;
	goto L198
L202:
	;
	goto L197
L203:
	;
	v1031 = v1024
	goto L204
L204:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+8))
	if v1046 != l2 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L191
L206:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+24))
	if v1053 != 0 {
		v1031 = v1053
		goto L204
	} else {
		goto L210
	}
L207:
	;
	v1048 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1031)+4)))
	if v1048 != v983 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	if v1050 == int32(112) {
		goto L190
	} else {
		goto L209
	}
L209:
	;
	goto L206
L210:
	;
	goto L205
L211:
	;
	goto L190
L212:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v662)+14)) = uint16(v980)
	goto L172
L213:
	;
	goto L171
L214:
	;
	goto L166
}
func F_subcoloronechr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
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
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	v6 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if base.Ui32(l1) <= base.Ui32(int32(2047)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v17 = F_subcolor(m, v14, l1)
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
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v118 = F_palloc_mul_extended(m, int32(12), v115+int32(2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L35
	}
L5:
	;
	return
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
	if v20 == v17&int32(_a_F_subcoloronechr_0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_subcoloronechr[0]))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v29 <= v30 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L11
L13:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v17)
	return
L14:
	;
	F_createarc(m, v24, int32(112), v17, l2, l3)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L34
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v32 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v58 == int32(0) {
		goto L14
	} else {
		goto L26
	}
L18:
	;
	v42 = v32
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v50 != l3 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L14
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v57 != 0 {
		v42 = v57
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	if v52 != v17&int32(_a_F_subcoloronechr_0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v54 == int32(112) {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	goto L20
L26:
	;
	v68 = v58
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v76 != l2 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L14
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if v83 != 0 {
		v68 = v83
		goto L27
	} else {
		goto L33
	}
L30:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	if v78 != v17&int32(_a_F_subcoloronechr_0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v80 == int32(112) {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L28
L34:
	;
	goto L13
L35:
	;
	if v118 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+24)) = int32(101)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	if v126 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v131 <= int32(0) {
		v169 = v130
		v170 = v6
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v128 = v126
	goto L41
L40:
	;
	v128 = int32(12)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = v128
	return
L42:
	;
	F_subcoloronerow(m, l0, v261, l2, l3, l4)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L68
	}
L43:
	;
	if v162 == v147 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v179 = v118 + v170*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = l1
	v183 = F_newhicolorrow(m, v14, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L53
	}
L45:
	;
	v139 = v130
	v140 = v6
	goto L46
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if base.Ui32(v147) < base.Ui32(l1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if base.Ui32(v162) <= base.Ui32(l1) {
		goto L43
	} else {
		goto L52
	}
L48:
	;
	v149 = int32(12)
	v151 = v118 + v140*v149
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v154
	v157 = v139 + v149
	v159 = v140 + int32(1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v159 < v160 {
		v139 = v157
		v140 = v159
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v169 = v157
	v170 = v159
	goto L44
L52:
	;
	v169 = v139
	v170 = v140
	goto L44
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v183
	v255 = v169
	v256 = v170
	v257 = v170 + int32(1)
	v261 = v183
	goto L42
L54:
	;
	v189 = int32(12)
	v191 = v118 + v140*v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v200 = v140 + int32(1)
	v255 = v139 + v189
	v256 = v200
	v257 = v200
	v261 = v196
	goto L42
L55:
	;
	goto L56
L56:
	;
	if base.Ui32(v162) < base.Ui32(l1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v204 = v118 + v140*int32(12)
	v205 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = l1 - v205
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v162
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+8)) = v209
	v213 = v140 + v205
	goto L59
L58:
	;
	v213 = v140
	goto L59
L59:
	;
	v216 = v118 + v213*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = l1
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v220 = F_newhicolorrow(m, v14, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v220
	v224 = v213 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if base.Ui32(l1) < base.Ui32(v225) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v229 = v118 + v224*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = l1 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if base.Ui32(v236) < base.Ui32(l1) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v244 = v224
	goto L63
L63:
	;
	v255 = v139 + int32(12)
	v256 = v140 + int32(1)
	v257 = v244
	v261 = v220
	goto L42
L64:
	;
	v238 = F_newhicolorrow(m, v14, v235)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	v240 = v235
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+8)) = v240
	v244 = v213 + int32(2)
	goto L63
L67:
	;
	v240 = v238
	goto L66
L68:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v256 < v265 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v272 = v255
	v273 = v256
	v274 = v257
	goto L72
L70:
	;
	v302 = v257
	goto L71
L71:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v308 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v280 = int32(12)
	v282 = v118 + v274*v280
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+8)) = v283
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v272)))
	*(*int64)(unsafe.Add(mBase, uint32(v282))) = v285
	v289 = int32(1)
	v290 = v274 + v289
	v292 = v273 + v289
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	if v292 < v293 {
		v272 = v272 + v280
		v273 = v292
		v274 = v290
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v302 = v290
	goto L71
L74:
	;
	goto L73
L75:
	;
	F_pfree(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v118
	goto L1
L78:
	;
	goto L77
}
func F_substitute_phv_relids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 == int32(321) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v11 != v12 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = v32
				if v34 == int32(67) {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
					v43 = F_query_tree_walker_impl(m, l0, int32(900), l1, int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
						return v43
					}
				} else {
					v51 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						return v51
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v16 = F_bms_is_member(m, v14, v15)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v16 == int32(0) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v34 = v32
						if v34 == int32(67) {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
							v43 = F_query_tree_walker_impl(m, l0, int32(900), l1, int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
								return v43
							}
						} else {
							v51 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								return v51
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v24 = F_bms_union(m, v22, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v28 = F_bms_del_member(m, v24, v27)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v28
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v34 = v32
								if v34 == int32(67) {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
									v43 = F_query_tree_walker_impl(m, l0, int32(900), l1, int32(0))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
										return v43
									}
								} else {
									v51 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										return v51
									}
								}
							}
						}
					}
				}
			}
		} else {
			v34 = v8
			if v34 == int32(67) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
				v43 = F_query_tree_walker_impl(m, l0, int32(900), l1, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v45 - int32(1)
					return v43
				}
			} else {
				v51 = F_expression_tree_walker_impl(m, l0, int32(900), l1)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v51
				}
			}
		}
	}
}
func F_swedish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v8 + int32(3)
	if v6 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v128 < v130 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v20 < v19 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v60 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v22 = v19
	goto L6
L5:
	;
	v22 = v20
	goto L6
L6:
	;
	v29 = v19
	goto L8
L7:
	;
	v60 = v40
	goto L3
L8:
	;
	if v29 == v22 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v60 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v29))))
	if int32(246) < v35 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v29 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v52
	v29 = v52
	goto L8
L14:
	;
	v37 = v35 - int32(97)
	if v37 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = int32(1)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v37)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v44)>>(uint(v37&int32(7))%32))&v40 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = v63 + v60
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v75 < v64 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v118 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v77 = v64
	goto L22
L21:
	;
	v77 = v75
	goto L22
L22:
	;
	v83 = v64
	goto L24
L23:
	;
	v118 = int32(1)
	goto L19
L24:
	;
	if v83 == v77 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v118 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v83))))
	if int32(246) < v92 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v94 = v92 - int32(97)
	if v94 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v94)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v100)>>(uint(v94&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v109 = v83 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
	v83 = v109
	goto L24
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v122 = v121 + v118
	if v10 < v122 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v124 = v122
	goto L36
L35:
	;
	v124 = v10
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v124
	goto L1
L37:
	;
	return v423
L38:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v266 <= v264 {
		goto L75
	} else {
		goto L76
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v130
	if v128 <= v130 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	goto L38
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v137 = int32(1)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v128-v137))))
	if base.B2i32(v139&int32(224) != int32(96))|base.B2i32(v137<<(uint(v139)%32)&int32(_a_F_swedish_ISO_8859_1_stem_0) == int32(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v154 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_1), int32(38), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v154 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v161
	switch v154 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	default:
		goto L38
	}
L46:
	;
	v253 = F_r_et_condition_1(m, l0)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L43
	} else {
		goto L72
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v169 = int32(2)
	v171 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v173-v174 < v169 {
		v184 = v171
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v165 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v165 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v423 = v165
	goto L37
L50:
	;
	v250 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v250 {
		goto L38
	} else {
		goto L71
	}
L51:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v195 = v193 + (v161 - v168)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v195
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L61
L52:
	;
	if v184 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L53:
	;
	goto L52
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v180 = F_memcmp(m, v177+v173-v169, int32(_a_F_swedish_ISO_8859_1_stem_2), v169)
	mBase = m.M
	if v180 != 0 {
		v184 = v171
		goto L53
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v173 - v169
	v184 = int32(1)
	goto L53
L56:
	;
	v187 = F_r_et_condition_1(m, l0)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	if v187 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v191
	goto L50
L59:
	;
	if v249 != 0 {
		goto L38
	} else {
		goto L70
	}
L60:
	;
	v249 = v245
	goto L59
L61:
	;
	if v195 <= v206 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v245 = int32(0)
	goto L60
L63:
	;
	v249 = int32(-1)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v218 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219+v195-v218))))
	if int32(121) < v223 {
		v245 = v218
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v225 = v223 - int32(98)
	if v225 < int32(0) {
		v245 = v218
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v231)>>(uint(v225&int32(7))%32))&int32(1) == int32(0) {
		v245 = v218
		goto L60
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v195 - int32(1)
	goto L69
L69:
	;
	goto L62
L70:
	;
	goto L50
L71:
	;
	v423 = v250
	goto L37
L72:
	;
	if v253 == int32(0) {
		goto L38
	} else {
		goto L73
	}
L73:
	;
	v257 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v257 {
		goto L38
	} else {
		goto L74
	}
L74:
	;
	v423 = v257
	goto L37
L75:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v266
	v271 = v264 - int32(1)
	if v271 <= v266 {
		v306 = v264
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v310 = v264
	v311 = v266
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310
	if v310 < v311 {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v268
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v310 = v306
	v311 = v309
	goto L77
L79:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v271))))
	if base.B2i32(v275&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v275)%32)&int32(_a_F_swedish_ISO_8859_1_stem_3) == int32(0)) != 0 {
		v306 = v264
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v290 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_4), int32(7), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L43
	} else {
		goto L81
	}
L81:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v290 == int32(0) {
		v306 = v292
		goto L78
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v292
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v292 <= v296 {
		v306 = v292
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v299 = v292 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v299
	v302 = F_slice_del(m, l0)
	mBase = m.M
	if v302 < int32(0) {
		v423 = v302
		goto L37
	} else {
		goto L84
	}
L84:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v306 = v305
	goto L78
L85:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v420
	v423 = int32(1)
	goto L37
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v310
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v311
	v320 = v310 - int32(1)
	if v320 <= v311 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v317
	goto L85
L88:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v320))))
	if base.B2i32(v324&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v324)%32)&int32(_a_F_swedish_ISO_8859_1_stem_5) == int32(0)) != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v339 = F_find_among_b(m, l0, int32(_a_F_swedish_ISO_8859_1_stem_6), int32(5), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L43
	} else {
		goto L90
	}
L90:
	;
	if v339 == int32(0) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v317
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v344
	switch v339 - int32(1) {
	case 0:
		goto L94
	case 1:
		goto L93
	case 2:
		goto L92
	default:
		goto L85
	}
L92:
	;
	v412 = F_slice_from_s(m, l0, int32(4), int32(_a_F_swedish_ISO_8859_1_stem_7))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L43
	} else {
		goto L110
	}
L93:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L98
L94:
	;
	v348 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v348 {
		goto L85
	} else {
		goto L95
	}
L95:
	;
	v423 = v348
	goto L37
L96:
	;
	if v403 != 0 {
		goto L85
	} else {
		goto L107
	}
L97:
	;
	v403 = v399
	goto L96
L98:
	;
	if v359 <= v360 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v399 = int32(0)
	goto L97
L100:
	;
	v403 = int32(-1)
	goto L96
L101:
	;
	goto L102
L102:
	;
	v372 = int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v359-v372))))
	if int32(118) < v377 {
		v399 = v372
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v379 = v377 - int32(105)
	if v379 < int32(0) {
		v399 = v372
		goto L97
	} else {
		goto L104
	}
L104:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v379)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v385)>>(uint(v379&int32(7))%32))&int32(1) == int32(0) {
		v399 = v372
		goto L97
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v359 - int32(1)
	goto L106
L106:
	;
	goto L99
L107:
	;
	v406 = F_slice_from_s(m, l0, int32(2), int32(_a_F_swedish_ISO_8859_1_stem_8))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L43
	} else {
		goto L108
	}
L108:
	;
	if int32(0) <= v406 {
		goto L85
	} else {
		goto L109
	}
L109:
	;
	v423 = v406
	goto L37
L110:
	;
	if int32(0) <= v412 {
		goto L85
	} else {
		goto L111
	}
L111:
	;
	v423 = v412
	goto L37
}
