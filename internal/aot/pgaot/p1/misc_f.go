package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FindLockCycleRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v8
	goto L3
L2:
	;
	v9 = l0
	goto L3
L3:
	;
	v10 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurse[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurse[1]))
	if v10 < v14 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = v10
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurse[1])) = v14 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12+v14<<(uint(int32(2))%32)))) = v9
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12+v17<<(uint(int32(2))%32))))
	if v9 == v27 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	if v17 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v36 = v17 + int32(1)
	if v36 != v14 {
		v17 = v36
		goto L7
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FindLockCycleRecurse[2])) = l1
	return int32(1)
L15:
	;
	goto L8
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+372))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v56 = F_FindLockCycleRecurseMember(m, v9, v9, l1, l2, l3)
	mBase = m.M
	if v56 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	return int32(1)
L19:
	;
	return int32(0)
L20:
	;
	v65 = v9 + int32(368)
	if v61 == v65 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v67 = v61
	goto L22
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v74 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L19
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v89 != v65 {
		v67 = v89
		goto L22
	} else {
		goto L29
	}
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v77 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v81 = v67 - int32(376)
	if v81 == v9 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v83 = F_FindLockCycleRecurseMember(m, v81, v9, l1, l2, l3)
	mBase = m.M
	if v83 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	return int32(1)
L29:
	;
	goto L23
}
func F_FlagRWConflict(m *base.Module, l0 int32, l1 int32) {
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v14&int32(1024) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L58
	} else {
		goto L73
	}
L2:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_FlagRWConflict[0]))
	F_LWLockRelease(m, v234+int32(3584))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L58
	} else {
		goto L66
	}
L3:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_FlagRWConflict[0]))
	F_LWLockRelease(m, v204+int32(3584))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L58
	} else {
		goto L59
	}
L4:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_FlagRWConflict[1]))
	if v146 == l0 {
		goto L45
	} else {
		goto L46
	}
L5:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_FlagRWConflict[2]))
	if v129 == l1 {
		goto L3
	} else {
		goto L42
	}
L6:
	;
	v18 = v14 & int32(1)
	v19 = int32(0)
	if base.B2i32(v18 == v19)|base.B2i32(v14&int32(1040) == v19) == v19 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v14&int32(2) == int32(0) {
		v139 = v14
		goto L4
	} else {
		goto L26
	}
L9:
	;
	v32 = l1 + int32(32)
	if v28 == v32 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v36 = v28
	goto L11
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+108)))
	if v44&int32(2) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v66 != v32 {
		v36 = v66
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v49&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v53) < base.Ui64(v52) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v18 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui64(v56) < base.Ui64(v55) {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v49&int32(32) == int32(0) {
		goto L5
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v62) <= base.Ui64(v63) {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	goto L12
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v81&int32(32) != 0 {
		v139 = v14
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v81&int32(512) != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v86 == int32(0) {
		v139 = v14
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v90 = l0 + int32(40)
	if v86 == v90 {
		v139 = v14
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v94 = v86
	goto L31
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+108))
	if v102&int32(8) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v139 = v14
	goto L4
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v117 != v90 {
		v94 = v117
		goto L31
	} else {
		goto L41
	}
L34:
	;
	if v102&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v101)+16))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(v107) < base.Ui64(v108) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v102&int32(32) == int32(0) {
		goto L5
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v101)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(v115) <= base.Ui64(v114) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	goto L32
L42:
	;
	if v14&int32(2) != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v134 = v14 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v134
	v139 = v134
	goto L4
L44:
	;
	m.G0 = v12 + int32(16)
	return
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v139 | int32(512)
	goto L44
L46:
	;
	goto L47
L47:
	;
	if l1 == v146 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v152 | int32(1024)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_FlagRWConflict[3]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if base.B2i32(v158 == int32(0))|base.B2i32(v158 == v157) != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = l0
	v171 = l0 + int32(32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v172 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v171
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+4)) = v171
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v158
	v183 = l1 + int32(40)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v184 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v183
	goto L57
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = v183
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v190
	v193 = v158 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v193
	goto L44
L58:
	;
	return
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_FlagRWConflict_0), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	F_errdetail_internal(m, int32(_a_F_FlagRWConflict_1), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	F_errhint(m, int32(_a_F_FlagRWConflict_2), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_FlagRWConflict_3), int32(_a_F_FlagRWConflict_4), int32(_a_F_FlagRWConflict_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L58
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L58
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L58
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_FlagRWConflict_0), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L58
	} else {
		goto L69
	}
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v250
	F_errdetail_internal(m, int32(_a_F_FlagRWConflict_6), v12)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L58
	} else {
		goto L70
	}
L70:
	;
	F_errhint(m, int32(_a_F_FlagRWConflict_2), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L58
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_FlagRWConflict_3), int32(_a_F_FlagRWConflict_7), int32(_a_F_FlagRWConflict_5))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L58
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(_a_F_FlagRWConflict_8))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L58
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_FlagRWConflict_9), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L58
	} else {
		goto L75
	}
L75:
	;
	F_errhint(m, int32(_a_F_FlagRWConflict_10), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L58
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_FlagRWConflict_3), int32(668), int32(_a_F_FlagRWConflict_11))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L58
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FormIndexDatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v37 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v35 = v32
	v36 = v33
	goto L1
L3:
	;
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v27 = v20
	goto L5
L5:
	;
	v29 = l0 + int32(80)
	if v27 != 0 {
		v31 = v27
		v32 = v29
		goto L2
	} else {
		goto L11
	}
L6:
	;
	v31 = v20
	v32 = l0 + int32(80)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v24 = F_ExecPrepareExprList(m, v21, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v24
	v27 = v24
	goto L5
L11:
	;
	v35 = v29
	v36 = int32(0)
	goto L1
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L48
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L45
	}
L14:
	;
	v52 = int32(0)
	v53 = v36
	goto L17
L15:
	;
	v153 = v36
	goto L16
L16:
	;
	if v153 != 0 {
		goto L12
	} else {
		goto L44
	}
L17:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(12)+v52<<(uint(int32(1))%32)))))
	v65 = base.I32_extend16_s(v64)
	if v65 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v153 = v131
	goto L16
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3+v52<<(uint(int32(3))%32)))) = v134
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v52))) = uint8(v140)
	v143 = v52 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v143 < v144 {
		v52 = v143
		v53 = v131
		goto L17
	} else {
		goto L43
	}
L20:
	;
	switch v64 - int32(_a_F_FormIndexDatum_0) {
	case 0:
		goto L25
	default:
		goto L23
	case 5:
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v65 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v79 = m.T0[v78].(func(*base.Module, int32, int32, int32) int64)(m, l1, v65, v18+int32(15))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v73)
	v131 = v53
	v134 = base.I64_extend_i32_u(l1 + int32(32))
	goto L19
L25:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v70)
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+40)))
	v131 = v53
	v134 = v72
	goto L19
L26:
	;
	v131 = v53
	v134 = v79
	goto L19
L27:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v81 < v65 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v53 == int32(0) {
		goto L13
	} else {
		goto L34
	}
L30:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	m.T0[v84].(func(*base.Module, int32, int32))(m, l1, v65)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v88 = v65 - int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v89))))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v91)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v93+v88<<(uint(int32(3))%32))))
	v131 = v53
	v134 = v97
	goto L19
L33:
	;
	goto L32
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v101 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	v106 = v101
	goto L37
L37:
	;
	v107 = int32(_a_F_FormIndexDatum_1)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_FormIndexDatum[0]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_FormIndexDatum[0])) = v110
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32) int64)(m, v100, v106, v18+int32(15))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L39
	}
L38:
	;
	v106 = v104
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FormIndexDatum[0])) = v108
	v120 = v53 + int32(4)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if base.Ui32(v120) < base.Ui32(v123+v124<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v129 = v120
	goto L42
L41:
	;
	v129 = int32(0)
	goto L42
L42:
	;
	v131 = v129
	v134 = v115
	goto L19
L43:
	;
	goto L18
L44:
	;
	m.G0 = v18 + int32(16)
	return
L45:
	;
	F_errmsg_internal(m, int32(_a_F_FormIndexDatum_2), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_FormIndexDatum_3), int32(2891), int32(_a_F_FormIndexDatum_4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(_a_F_FormIndexDatum_2), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_FormIndexDatum_3), int32(2902), int32(_a_F_FormIndexDatum_4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FreeDir(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDir[0]))
	v10 = v8 - int32(1)
	if int32(0) <= v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDir[1]))
	v16 = v10
	goto L7
L5:
	;
	goto L6
L6:
	;
	v39 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L15
	}
L7:
	;
	v21 = v14 + v16*int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 != int32(2) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	if int32(0) < v16 {
		v16 = v16 - int32(1)
		goto L7
	} else {
		goto L14
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v25 != l0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = F_FreeDesc(m, v21)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	return
L14:
	;
	goto L8
L15:
	;
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_errmsg_internal(m, int32(_a_F_FreeDir_0), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = F_close(m, v51)
	mBase = m.M
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	goto L21
L19:
	;
	F_errfinish(m, int32(_a_F_FreeDir_1), int32(3034), int32(_a_F_FreeDir_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return
}
func F___fwritex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v259
L2:
	;
	v33 = v7
	goto L4
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v9 - int32(1) | v9
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v14&int32(8) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if base.Ui32(v33-v34) < base.Ui32(l1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	if v31 != 0 {
		v259 = int32(0)
		goto L1
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v14 | int32(32)
	v31 = int32(-1)
	goto L5
L7:
	;
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v23 + v26
	v31 = int32(0)
	goto L5
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v33 = v32
	goto L4
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v38 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, l2, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if base.B2i32(l1 == v43)|base.B2i32(v45 < v43) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	return int32(0)
L14:
	;
	return v38
L15:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v77) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v77 = l1
	v79 = int32(0)
	v80 = v34
	v81 = l0
	goto L15
L17:
	;
	v52 = l1
	goto L18
L18:
	;
	v55 = l0 + v52
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55-int32(1)))))
	if v58 != int32(10) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v64 = m.T0[v63].(func(*base.Module, int32, int32, int32) int32)(m, l2, l0, v52)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L24
	}
L20:
	;
	v62 = v52 - int32(1)
	if v62 != 0 {
		v52 = v62
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	goto L16
L24:
	;
	if base.Ui32(v64) < base.Ui32(v52) {
		v259 = v64
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v77 = l1 - v52
	v79 = v52
	v80 = v68
	v81 = v55
	goto L15
L26:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v251 + v77
	v259 = v77 + v79
	goto L1
L27:
	;
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v88 = v80 + v77
	if (v80^v81)&int32(3) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	base.MemoryCopy(m, v80, v81, v77)
	goto L32
L31:
	;
	goto L32
L32:
	;
	goto L26
L33:
	;
	if base.Ui32(v220) < base.Ui32(v88) {
		goto L67
	} else {
		goto L68
	}
L34:
	;
	if v80&int32(3) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(v88) < base.Ui32(int32(4)) {
		goto L58
	} else {
		goto L59
	}
L37:
	;
	v124 = v88 & int32(-4)
	if base.Ui32(v88) < base.Ui32(int32(64)) {
		v174 = v118
		v175 = v119
		goto L48
	} else {
		goto L49
	}
L38:
	;
	v118 = v81
	v119 = v80
	goto L37
L39:
	;
	goto L40
L40:
	;
	if v77 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v118 = v81
	v119 = v80
	goto L37
L42:
	;
	goto L43
L43:
	;
	v101 = v81
	v102 = v80
	goto L44
L44:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v106)
	v108 = int32(1)
	v109 = v101 + v108
	v111 = v102 + v108
	if v111&int32(3) == int32(0) {
		v118 = v109
		v119 = v111
		goto L37
	} else {
		goto L46
	}
L45:
	;
	v118 = v109
	v119 = v111
	goto L37
L46:
	;
	if base.Ui32(v111) < base.Ui32(v88) {
		v101 = v109
		v102 = v111
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	if base.Ui32(v124) <= base.Ui32(v175) {
		v219 = v174
		v220 = v175
		goto L33
	} else {
		goto L54
	}
L49:
	;
	v128 = v124 + int32(-64)
	if base.Ui32(v128) < base.Ui32(v119) {
		v174 = v118
		v175 = v119
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v131 = v118
	v132 = v119
	goto L51
L51:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+12)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+16)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+20)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v131)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+28)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v131)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+32)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v131)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+36)) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v131)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+40)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+44)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v131)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+48)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v131)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+52)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v131)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+56)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v131)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+60)) = v166
	v168 = int32(-64)
	v169 = v131 - v168
	v171 = v132 - v168
	if base.Ui32(v171) <= base.Ui32(v128) {
		v131 = v169
		v132 = v171
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v174 = v169
	v175 = v171
	goto L48
L53:
	;
	goto L52
L54:
	;
	v181 = v174
	v182 = v175
	goto L55
L55:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v186
	v188 = int32(4)
	v189 = v181 + v188
	v191 = v182 + v188
	if base.Ui32(v191) < base.Ui32(v124) {
		v181 = v189
		v182 = v191
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v219 = v189
	v220 = v191
	goto L33
L57:
	;
	goto L56
L58:
	;
	v219 = v81
	v220 = v80
	goto L33
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(v77) < base.Ui32(int32(4)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v219 = v81
	v220 = v80
	goto L33
L62:
	;
	goto L63
L63:
	;
	v200 = v81
	v201 = v80
	goto L64
L64:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v205)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)) = uint8(v207)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+2)) = uint8(v209)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+3)) = uint8(v211)
	v213 = int32(4)
	v214 = v200 + v213
	v216 = v201 + v213
	if base.Ui32(v216) <= base.Ui32(v88-int32(4)) {
		v200 = v214
		v201 = v216
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v219 = v214
	v220 = v216
	goto L33
L66:
	;
	goto L65
L67:
	;
	v226 = v219
	v227 = v220
	goto L70
L68:
	;
	goto L69
L69:
	;
	goto L26
L70:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v227))) = uint8(v231)
	v233 = int32(1)
	v236 = v227 + v233
	if v236 != v88 {
		v226 = v226 + v233
		v227 = v236
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	goto L71
}
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14263(m, l0, l1, l2, l3, int32(_a_F_fastgetattr_3_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_fetch_finfo_record(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
	v13 = F_psprintf(m, int32(_a_F_fetch_finfo_record_0), v5+int32(-16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_pgmem_dlsym(m, l0, v13)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v19 = m.T0[v17].(func(*base.Module) int32)(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v13
							F_errmsg_internal(m, int32(_a_F_fetch_finfo_record_1), v5+int32(-48))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_fetch_finfo_record_2), int32(483), int32(_a_F_fetch_finfo_record_3))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
						if v23 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v13
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v74
									F_errmsg(m, int32(_a_F_fetch_finfo_record_4), v5+int32(-32))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_fetch_finfo_record_2), int32(493), int32(_a_F_fetch_finfo_record_3))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_pfree(m, v13)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 - int32(-64)
								return v19
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(52461700))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg(m, int32(_a_F_fetch_finfo_record_5), v7)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_fetch_finfo_record_6), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_fetch_finfo_record_2), int32(474), int32(_a_F_fetch_finfo_record_3))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
			}
		}
	}
}
func F_fetch_upper_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v10 = l0 + l1<<(uint(int32(2))%32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+212))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v130
L2:
	;
	v94 = v10 + int32(212)
	v96 = F_palloc0(m, int32(304))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v14 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v15 <= v14 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = v14
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v19<<(uint(int32(2))%32))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v31 = int32(0)
	if base.B2i32(v30 == v31)|base.B2i32(l2 == v31) != 0 {
		v77 = base.B2i32(v30|l2 == v31)
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	if v77 != 0 {
		v130 = v29
		goto L1
	} else {
		goto L18
	}
L8:
	;
	goto L7
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v45 != v46 {
		v77 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v48 = int32(1)
	if v45 <= v48 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v48
	goto L13
L12:
	;
	v51 = v45
	goto L13
L13:
	;
	v52 = int32(8)
	v57 = int32(0)
	goto L14
L14:
	;
	v65 = v57 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v30+v52+v65)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2+v52+v65)))
	v70 = base.B2i32(v67 == v69)
	if v67 != v69 {
		v77 = v70
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v77 = v70
	goto L8
L16:
	;
	v73 = v57 + int32(1)
	if v73 != v51 {
		v57 = v73
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v83 = v19 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v83 < v84 {
		v19 = v83
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L6
L20:
	;
	return int32(0)
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = int64(17179869454)
	v102 = F_bms_copy(m, l2)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+32)) = v106
	v108 = *(*float64)(unsafe.Add(mBase, uint32(l0)+312))
	v109 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+25)) = uint16(v109)
	v112 = base.F64_gt(v108, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)) = uint8(v112)
	v114 = F_create_empty_pathtarget(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+64)) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v96)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+44)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v96)+40)) = v114
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v124 = F_lappend(m, v123, v96)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v124
	v130 = v96
	goto L1
}
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
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
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	if l1 <= int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v380
L2:
	;
	v377 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v374))) = uint8(v377)
	v380 = l0
	goto L1
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	v11 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v10 - v11 | v10
	if l1 == v11 {
		v374 = l0
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v22 = l1 - int32(1)
	v25 = l0
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v28 == v29 {
		v331 = v22
		v333 = v25
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if l0 != 0 {
		v374 = v365
		goto L2
	} else {
		goto L102
	}
L9:
	;
	goto L8
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v350)
	v356 = v351 + int32(1)
	if v350&int32(255) == int32(10) {
		v365 = v356
		goto L9
	} else {
		goto L100
	}
L11:
	;
	v336 = F___uflow(m, l2)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L95
	} else {
		goto L96
	}
L12:
	;
	v32 = v29 - v28
	v33 = int32(0)
	if base.B2i32(v28&int32(3) == v33)|base.B2i32(v32 == v33) != 0 {
		v63 = v28
		v65 = v32
		v66 = base.B2i32(v32 != v33)
		goto L17
	} else {
		goto L18
	}
L13:
	;
	if base.Ui32(v146) < base.Ui32(v22) {
		goto L42
	} else {
		goto L43
	}
L14:
	;
	if v137 != 0 {
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v137 = int32(0)
	goto L14
L16:
	;
	v115 = v108
	v117 = v110
	goto L33
L17:
	;
	if v66 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v46 = v28
	v48 = v32
	goto L19
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v51 == int32(10) {
		v108 = v46
		v110 = v48
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v63 = v58
	v65 = v54
	v66 = v56
	goto L17
L21:
	;
	v53 = int32(1)
	v54 = v48 - v53
	v55 = int32(0)
	v56 = base.B2i32(v54 != v55)
	v58 = v46 + v53
	if v58&int32(3) == v55 {
		v63 = v58
		v65 = v54
		v66 = v56
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if v54 != 0 {
		v46 = v58
		v48 = v54
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if base.B2i32(int32(10) == v72)|base.B2i32(base.Ui32(v65) < base.Ui32(int32(4))) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = v63
	v83 = v65
	goto L28
L26:
	;
	v101 = v63
	v103 = v65
	goto L27
L27:
	;
	if v103 == int32(0) {
		goto L15
	} else {
		goto L32
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v88 = v87 ^ int32(168430090)
	v91 = int32(-2139062144)
	if (int32(16843008)-v88|v88)&v91 != v91 {
		v108 = v81
		v110 = v83
		goto L16
	} else {
		goto L30
	}
L29:
	;
	v101 = v96
	v103 = v98
	goto L27
L30:
	;
	v95 = int32(4)
	v96 = v81 + v95
	v98 = v83 - v95
	if base.Ui32(int32(3)) < base.Ui32(v98) {
		v81 = v96
		v83 = v98
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v108 = v101
	v110 = v103
	goto L16
L33:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if int32(10) == v120 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L15
L35:
	;
	v137 = v115
	goto L14
L36:
	;
	goto L37
L37:
	;
	v122 = int32(1)
	v125 = v117 - v122
	if v125 != 0 {
		v115 = v115 + v122
		v117 = v125
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v145 = v138
	v146 = v137 - v138 + int32(1)
	goto L13
L40:
	;
	goto L41
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v145 = v143
	v146 = v142 - v143
	goto L13
L42:
	;
	v148 = v146
	goto L44
L43:
	;
	v148 = v22
	goto L44
L44:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v148) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v319 = v318 + v148
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v319
	v321 = v148 + v25
	if v137 != 0 {
		v365 = v321
		goto L9
	} else {
		goto L92
	}
L46:
	;
	if v148 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v155 = v25 + v148
	if (v25^v145)&int32(3) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	base.MemoryCopy(m, v25, v145, v148)
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L45
L52:
	;
	if base.Ui32(v287) < base.Ui32(v155) {
		goto L86
	} else {
		goto L87
	}
L53:
	;
	if v25&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v155) < base.Ui32(int32(4)) {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	v191 = v155 & int32(-4)
	if base.Ui32(v155) < base.Ui32(int32(64)) {
		v241 = v185
		v242 = v186
		goto L67
	} else {
		goto L68
	}
L57:
	;
	v185 = v145
	v186 = v25
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v148 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v185 = v145
	v186 = v25
	goto L56
L61:
	;
	goto L62
L62:
	;
	v168 = v145
	v169 = v25
	goto L63
L63:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v173)
	v175 = int32(1)
	v176 = v168 + v175
	v178 = v169 + v175
	if v178&int32(3) == int32(0) {
		v185 = v176
		v186 = v178
		goto L56
	} else {
		goto L65
	}
L64:
	;
	v185 = v176
	v186 = v178
	goto L56
L65:
	;
	if base.Ui32(v178) < base.Ui32(v155) {
		v168 = v176
		v169 = v178
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	if base.Ui32(v191) <= base.Ui32(v242) {
		v286 = v241
		v287 = v242
		goto L52
	} else {
		goto L73
	}
L68:
	;
	v195 = v191 + int32(-64)
	if base.Ui32(v195) < base.Ui32(v186) {
		v241 = v185
		v242 = v186
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v198 = v185
	v199 = v186
	goto L70
L70:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+16)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+20)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v198)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v198)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+28)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v198)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+32)) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v198)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+36)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v198)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+40)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v198)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+44)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v198)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+48)) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v198)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+52)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v198)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+56)) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v198)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+60)) = v233
	v235 = int32(-64)
	v236 = v198 - v235
	v238 = v199 - v235
	if base.Ui32(v238) <= base.Ui32(v195) {
		v198 = v236
		v199 = v238
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v241 = v236
	v242 = v238
	goto L67
L72:
	;
	goto L71
L73:
	;
	v248 = v241
	v249 = v242
	goto L74
L74:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v253
	v255 = int32(4)
	v256 = v248 + v255
	v258 = v249 + v255
	if base.Ui32(v258) < base.Ui32(v191) {
		v248 = v256
		v249 = v258
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v286 = v256
	v287 = v258
	goto L52
L76:
	;
	goto L75
L77:
	;
	v286 = v145
	v287 = v25
	goto L52
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(v148) < base.Ui32(int32(4)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v286 = v145
	v287 = v25
	goto L52
L81:
	;
	goto L82
L82:
	;
	v267 = v145
	v268 = v25
	goto L83
L83:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	*(*uint8)(unsafe.Add(mBase, uint32(v268))) = uint8(v272)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)) = uint8(v274)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+2)) = uint8(v276)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+3)) = uint8(v278)
	v280 = int32(4)
	v281 = v267 + v280
	v283 = v268 + v280
	if base.Ui32(v283) <= base.Ui32(v155-int32(4)) {
		v267 = v281
		v268 = v283
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v286 = v281
	v287 = v283
	goto L52
L85:
	;
	goto L84
L86:
	;
	v293 = v286
	v294 = v287
	goto L89
L87:
	;
	goto L88
L88:
	;
	goto L45
L89:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	*(*uint8)(unsafe.Add(mBase, uint32(v294))) = uint8(v298)
	v300 = int32(1)
	v303 = v294 + v300
	if v303 != v155 {
		v293 = v293 + v300
		v294 = v303
		goto L89
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	goto L90
L92:
	;
	v322 = v22 - v148
	if v322 == int32(0) {
		v365 = v321
		goto L9
	} else {
		goto L93
	}
L93:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v319 == v325 {
		v331 = v322
		v333 = v321
		goto L11
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v319 + int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v349 = v322
	v350 = v330
	v351 = v321
	goto L10
L95:
	;
	return int32(0)
L96:
	;
	if int32(0) <= v336 {
		v349 = v331
		v350 = v336
		v351 = v333
		goto L10
	} else {
		goto L97
	}
L97:
	;
	v342 = int32(0)
	if l0 == v333 {
		v380 = v342
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v344&int32(16) == int32(0) {
		v380 = v342
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v365 = v333
	goto L9
L100:
	;
	v362 = v349 - int32(1)
	if v362 != 0 {
		v22 = v362
		v25 = v356
		goto L7
	} else {
		goto L101
	}
L101:
	;
	v365 = v356
	goto L9
L102:
	;
	return int32(0)
}
func F_fill_val(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v6 = l5
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if l1 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v15 != int32(128) {
			v28 = v15 << (uint(int32(1)) % 32)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v21 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20 + v21
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v24)
			v28 = v21
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
		if l6 != 0 {
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
			v32 = v30 | int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v32)
			m.G0 = v12 + int32(16)
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
			v36 = v35 | v28
			*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v36)
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			if v40 == int32(1) {
				v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
				if base.I32_popcnt(v43) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
						F_errmsg_internal(m, int32(_a_F_fill_val_0), v12)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_fill_val_1), int32(474), int32(_a_F_fill_val_2))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					v53 = (v14 + v47 - int32(1)) & (int32(0) - v47)
					switch base.I32_ctz(v43) {
					case 0:
						*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v6)
						v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						v197 = v77
						v200 = v53
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
						m.G0 = v12 + int32(16)
						return
					case 1:
						*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v6)
						v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						v197 = v56
						v200 = v53
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
						m.G0 = v12 + int32(16)
						return
					case 2:
						*(*uint32)(unsafe.Add(mBase, uint32(v53))) = uint32(v6)
						v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						v197 = v58
						v200 = v53
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
						m.G0 = v12 + int32(16)
						return
					case 3:
						*(*int64)(unsafe.Add(mBase, uint32(v53))) = v6
						v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
						v197 = v60
						v200 = v53
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
						m.G0 = v12 + int32(16)
						return
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
							F_errmsg_internal(m, int32(_a_F_fill_val_0), v12)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_fill_val_1), int32(474), int32(_a_F_fill_val_2))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
				v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
				switch v78 - int32(_a_F_fill_val_3) {
				case 0:
					v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
					v175 = v173 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v175)
					v177 = base.I32_wrap_i64(v6)
					v178 = F_strlen(m, v177)
					mBase = m.M
					v180 = v178 + int32(1)
					if v180 == int32(0) {
						v197 = v180
						v200 = v14
					} else {
						base.MemoryCopy(m, v14, v177, v180)
						v197 = v180
						v200 = v14
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				case 1:
					v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
					v83 = v81 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v83)
					v85 = base.I32_wrap_i64(v6)
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
					if v86 == int32(1) {
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
						if v89&int32(254) == int32(2) {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v6))+2))
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							v102 = (v14 + v96 - int32(1)) & (int32(0) - v96)
							v103 = F_EOH_get_flat_size(m, v95)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								F_EOH_flatten_into(m, v95, v102, v103)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									v197 = v103
									v200 = v102
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
									m.G0 = v12 + int32(16)
									return
								}
							}
						} else {
							v107 = int32(6)
							v108 = v81 | v107
							*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v108)
							v111 = int32(18)
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
							if v113 == v111 {
								v116 = v111
							} else {
								v116 = int32(2)
							}
							if base.Ui32((v113-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v123 = v107
							} else {
								v123 = v116
							}
							if v123 == int32(0) {
								v197 = v123
								v200 = v14
							} else {
								base.MemoryCopy(m, v14, v85, v123)
								v197 = v123
								v200 = v14
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						if v86&int32(1) != 0 {
							v130 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
							if v130 == int32(0) {
								v197 = v130
								v200 = v14
							} else {
								base.MemoryCopy(m, v14, v85, v130)
								v197 = v130
								v200 = v14
							}
						} else {
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
							v135 = int32(2)
							v136 = int32(base.Ui32(v134) >> (uint(v135) % 32))
							if v86&v135 != 0 {
								v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
								v167 = int32(0)
								v169 = (v14 + v163 - int32(1)) & (v167 - v163)
								if v136 == v167 {
									v197 = v136
									v200 = v169
								} else {
									base.MemoryCopy(m, v169, v85, v136)
									v197 = v136
									v200 = v169
								}
							} else {
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
								if v139&int32(1) == int32(0) {
									v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
									v167 = int32(0)
									v169 = (v14 + v163 - int32(1)) & (v167 - v163)
									if v136 == v167 {
										v197 = v136
										v200 = v169
									} else {
										base.MemoryCopy(m, v169, v85, v136)
										v197 = v136
										v200 = v169
									}
								} else {
									v145 = v136 - int32(3)
									if base.Ui32(int32(127)) < base.Ui32(v145) {
										v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
										v167 = int32(0)
										v169 = (v14 + v163 - int32(1)) & (v167 - v163)
										if v136 == v167 {
											v197 = v136
											v200 = v169
										} else {
											base.MemoryCopy(m, v169, v85, v136)
											v197 = v136
											v200 = v169
										}
									} else {
										v148 = int32(1)
										v151 = v145<<(uint(v148)%32) | v148
										*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v151)
										v154 = v136 - int32(4)
										if v154 == int32(0) {
											v197 = v145
											v200 = v14
										} else {
											base.MemoryCopy(m, v14+int32(1), v85+int32(4), v154)
											v197 = v145
											v200 = v14
										}
									}
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
						m.G0 = v12 + int32(16)
						return
					}
				default:
					v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
					v188 = int32(0)
					v190 = (v14 + v184 - int32(1)) & (v188 - v184)
					v191 = base.I32_extend16_s(v78)
					if v191 == v188 {
						v197 = v191
						v200 = v190
					} else {
						base.MemoryCopy(m, v190, base.I32_wrap_i64(v6), v191)
						v197 = v191
						v200 = v190
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	} else {
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v40 == int32(1) {
			v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
			if base.I32_popcnt(v43) != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
					F_errmsg_internal(m, int32(_a_F_fill_val_0), v12)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_fill_val_1), int32(474), int32(_a_F_fill_val_2))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				v53 = (v14 + v47 - int32(1)) & (int32(0) - v47)
				switch base.I32_ctz(v43) {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v6)
					v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					v197 = v77
					v200 = v53
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v6)
					v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					v197 = v56
					v200 = v53
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				case 2:
					*(*uint32)(unsafe.Add(mBase, uint32(v53))) = uint32(v6)
					v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					v197 = v58
					v200 = v53
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				case 3:
					*(*int64)(unsafe.Add(mBase, uint32(v53))) = v6
					v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
					v197 = v60
					v200 = v53
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v43
						F_errmsg_internal(m, int32(_a_F_fill_val_0), v12)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_fill_val_1), int32(474), int32(_a_F_fill_val_2))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
			v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			switch v78 - int32(_a_F_fill_val_3) {
			case 0:
				v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
				v175 = v173 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v175)
				v177 = base.I32_wrap_i64(v6)
				v178 = F_strlen(m, v177)
				mBase = m.M
				v180 = v178 + int32(1)
				if v180 == int32(0) {
					v197 = v180
					v200 = v14
				} else {
					base.MemoryCopy(m, v14, v177, v180)
					v197 = v180
					v200 = v14
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
				m.G0 = v12 + int32(16)
				return
			case 1:
				v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
				v83 = v81 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v83)
				v85 = base.I32_wrap_i64(v6)
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
				if v86 == int32(1) {
					v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
					if v89&int32(254) == int32(2) {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v6))+2))
						v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
						v102 = (v14 + v96 - int32(1)) & (int32(0) - v96)
						v103 = F_EOH_get_flat_size(m, v95)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							F_EOH_flatten_into(m, v95, v102, v103)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								v197 = v103
								v200 = v102
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
								m.G0 = v12 + int32(16)
								return
							}
						}
					} else {
						v107 = int32(6)
						v108 = v81 | v107
						*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v108)
						v111 = int32(18)
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
						if v113 == v111 {
							v116 = v111
						} else {
							v116 = int32(2)
						}
						if base.Ui32((v113-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v123 = v107
						} else {
							v123 = v116
						}
						if v123 == int32(0) {
							v197 = v123
							v200 = v14
						} else {
							base.MemoryCopy(m, v14, v85, v123)
							v197 = v123
							v200 = v14
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					if v86&int32(1) != 0 {
						v130 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
						if v130 == int32(0) {
							v197 = v130
							v200 = v14
						} else {
							base.MemoryCopy(m, v14, v85, v130)
							v197 = v130
							v200 = v14
						}
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
						v135 = int32(2)
						v136 = int32(base.Ui32(v134) >> (uint(v135) % 32))
						if v86&v135 != 0 {
							v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
							v167 = int32(0)
							v169 = (v14 + v163 - int32(1)) & (v167 - v163)
							if v136 == v167 {
								v197 = v136
								v200 = v169
							} else {
								base.MemoryCopy(m, v169, v85, v136)
								v197 = v136
								v200 = v169
							}
						} else {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
							if v139&int32(1) == int32(0) {
								v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
								v167 = int32(0)
								v169 = (v14 + v163 - int32(1)) & (v167 - v163)
								if v136 == v167 {
									v197 = v136
									v200 = v169
								} else {
									base.MemoryCopy(m, v169, v85, v136)
									v197 = v136
									v200 = v169
								}
							} else {
								v145 = v136 - int32(3)
								if base.Ui32(int32(127)) < base.Ui32(v145) {
									v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
									v167 = int32(0)
									v169 = (v14 + v163 - int32(1)) & (v167 - v163)
									if v136 == v167 {
										v197 = v136
										v200 = v169
									} else {
										base.MemoryCopy(m, v169, v85, v136)
										v197 = v136
										v200 = v169
									}
								} else {
									v148 = int32(1)
									v151 = v145<<(uint(v148)%32) | v148
									*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v151)
									v154 = v136 - int32(4)
									if v154 == int32(0) {
										v197 = v145
										v200 = v14
									} else {
										base.MemoryCopy(m, v14+int32(1), v85+int32(4), v154)
										v197 = v145
										v200 = v14
									}
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
					m.G0 = v12 + int32(16)
					return
				}
			default:
				v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				v188 = int32(0)
				v190 = (v14 + v184 - int32(1)) & (v188 - v184)
				v191 = base.I32_extend16_s(v78)
				if v191 == v188 {
					v197 = v191
					v200 = v190
				} else {
					base.MemoryCopy(m, v190, base.I32_wrap_i64(v6), v191)
					v197 = v191
					v200 = v190
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v197 + v200
				m.G0 = v12 + int32(16)
				return
			}
		}
	}
}
func F_filter_by_origin_cb_wrapper(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_filter_by_origin_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(1058)
	v16 = int32(_a_F_filter_by_origin_cb_wrapper_1)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_filter_by_origin_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_filter_by_origin_cb_wrapper[0])) = v7 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v7 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+164)) = uint8(v3)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+147)) = uint8(v3)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v31 = m.T0[v30].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_filter_by_origin_cb_wrapper[0])) = v36
		m.G0 = v7 + int32(32)
		return v31
	}
}
func F_finalize_aggregates(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int64
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int64
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
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
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v390 int64
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int64
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int64
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int64
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int64
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v560 int32
	_ = v560
	var v564 int64
	_ = v564
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v590 int64
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v729 int32
	_ = v729
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int64
	_ = v789
	var v790 int64
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int64
	_ = v803
	var v804 int32
	_ = v804
	var v805 int64
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int64
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v832 int64
	_ = v832
	var v833 int64
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int64
	_ = v843
	var v844 int64
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v857 int64
	_ = v857
	var v858 int64
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int64
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int64
	_ = v981
	var v982 int64
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v995 int64
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int64
	_ = v998
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1069 int32
	_ = v1069
	var v1078 int32
	_ = v1078
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int64
	_ = v1103
	var v1118 int32
	_ = v1118
	var v1128 int32
	_ = v1128
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1165 int64
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1184 int64
	_ = v1184
	var v1185 int64
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int64
	_ = v1195
	var v1196 int64
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1209 int64
	_ = v1209
	var v1210 int64
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	v4 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(1632)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+36))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v4 < v33 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if int32(0) < v710 {
		goto L115
	} else {
		goto L116
	}
L4:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v64 = v61 + v48*int32(240)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+5)))
	if v65 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v682 = v48 + int32(1)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v682 < v683 {
		v48 = v682
		goto L4
	} else {
		goto L114
	}
L7:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v64)+124))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v64)+196))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v64)+224))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v64)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = int64(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v331+v332<<(uint(int32(2))%32))))
	F_tuplesort_performsort(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L12
	} else {
		goto L64
	}
L8:
	;
	v70 = l2 + v48<<(uint(int32(4))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v72 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v64)+124))
	if v299 <= int32(0) {
		goto L6
	} else {
		goto L54
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+124))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = int64(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+224))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81<<(uint(int32(2))%32))))
	F_tuplesort_performsort(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88+v89<<(uint(int32(2))%32))))
	v97 = v79 + int32(40)
	v99 = v79 + int32(48)
	v102 = F_tuplesort_getdatum(m, v93, int32(1), int32(0), v97, v99, v28+int32(8))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(int32(2))%32))))
	F_tuplesort_end(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L12
	} else {
		goto L53
	}
L15:
	;
	if v102 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v108 = int32(0)
	v110 = int64(0)
	v119 = v108
	v122 = int32(1)
	v137 = v110
	v138 = v110
	goto L17
L17:
	;
	goto L24
L18:
	;
	if v249&int32(1) != 0 {
		goto L14
	} else {
		goto L50
	}
L19:
	;
	goto L18
L20:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v234 = int32(1)
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236+v237<<(uint(int32(2))%32))))
	v246 = F_tuplesort_getdatum(m, v241, v234, int32(0), v97, v99, v28+int32(8))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L12
	} else {
		goto L48
	}
L21:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v225 != 0 {
		v232 = v138
		goto L20
	} else {
		goto L46
	}
L22:
	;
	if v122&int32(1) != 0 {
		goto L21
	} else {
		goto L44
	}
L23:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	v232 = v218
	goto L20
L24:
	;
	F_MemoryContextReset(m, v76)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L12
	} else {
		goto L27
	}
L25:
	;
	F_advance_transition_function(m, l0, v64, v70)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L42
	}
L26:
	;
	goto L25
L27:
	;
	v167 = int32(_a_F_finalize_aggregates_0)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v76
	if v119&base.B2i32(v108 < v75) == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v122&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v168
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197+v198<<(uint(int32(2))%32))))
	v207 = F_tuplesort_getdatum(m, v202, int32(1), int32(0), v97, v99, v28+int32(8))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L12
	} else {
		goto L40
	}
L30:
	;
	if v173&int32(1) != 0 {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v173&int32(1) != 0 {
		goto L26
	} else {
		goto L36
	}
L33:
	;
	F_advance_transition_function(m, l0, v64, v70)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v168
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+190)))
	if v182 == int32(0) {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	if v137 != v187 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v64)+116))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	v191 = F_FunctionCall2Coll(m, v64+int32(144), v189, v138, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	if v191 == int64(0) {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	goto L29
L40:
	;
	if v207 != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	v249 = v122
	v252 = v138
	goto L19
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v168
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+190)))
	if v214 == int32(0) {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	goto L23
L44:
	;
	F_pfree(m, base.I32_wrap_i64(v138))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	goto L21
L46:
	;
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+190)))
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+186)))
	v229 = F_datumCopy(m, v226, v227, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v232 = v229
	goto L20
L48:
	;
	if v246 != 0 {
		v119 = v234
		v122 = v233
		v137 = v235
		v138 = v232
		goto L17
	} else {
		goto L49
	}
L49:
	;
	v249 = v233
	v252 = v232
	goto L19
L50:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+190)))
	if v255 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, base.I32_wrap_i64(v252))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	goto L14
L53:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v292+v293<<(uint(int32(2))%32)))) = int32(0)
	goto L6
L54:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+217)))
	if v302 != int32(1) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+217)) = uint8(v305)
	if v299 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+190)))
	if v309 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v64)+196))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	m.T0[v320].(func(*base.Module, int32))(m, v318)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L12
	} else {
		goto L63
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v64)+208)) = int64(0)
	v316 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+216)) = uint8(v316)
	goto L6
L60:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+216)))
	if v310 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v64)+208))
	F_pfree(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L6
L64:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	m.T0[v340].(func(*base.Module, int32))(m, v327)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	if v325 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	m.T0[v344].(func(*base.Module, int32))(m, v325)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L12
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347+v348<<(uint(int32(2))%32))))
	v353 = int32(1)
	v357 = F_tuplesort_gettupleslot(m, v352, v353, v353, v327, v28+int32(8))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L12
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v357 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v364 = v326 + int32(24)
	v370 = v325
	v373 = v327
	v381 = int32(0)
	v390 = int64(0)
	goto L74
L72:
	;
	v614 = v325
	goto L73
L73:
	;
	if v614 != 0 {
		goto L109
	} else {
		goto L110
	}
L74:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[1]))
	if v393 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v614 = v570
	goto L73
L76:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L12
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v373
	v398 = int32(0)
	if base.B2i32(v323 != v398)&v381 == v398 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L78
L80:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	F_MemoryContextReset(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L12
	} else {
		goto L105
	}
L81:
	;
	v570 = v370
	v573 = v373
	v581 = int32(1)
	v590 = v390
	goto L80
L82:
	;
	v424 = int32(*(*int16)(unsafe.Add(mBase, uint32(v373)+6)))
	if v424 < v324 {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	if v403 != v390 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v64)+172))
	if v405 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v408 = int32(_a_F_finalize_aggregates_0)
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v411
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v405)+24))
	v416 = m.T0[v415].(func(*base.Module, int32, int32, int32) int64)(m, v405, v71, v28+int32(7))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v409
	if v416 != int64(0) {
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L82
L88:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+16))
	m.T0[v427].(func(*base.Module, int32, int32))(m, v373, v324)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L12
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v324 <= int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	F_advance_transition_function(m, l0, v64, v70)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L101
	}
L93:
	;
	v432 = int32(0)
	if v324 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v439 = v432
	v444 = v432
	goto L97
L95:
	;
	v499 = v432
	goto L96
L96:
	;
	v523 = v364 + v499<<(uint(int32(4))%32)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v524+v499<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v523)+16)) = v528
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530+v499))))
	*(*uint8)(unsafe.Add(mBase, uint32(v523)+24)) = uint8(v532)
	goto L92
L97:
	;
	v462 = v439 | int32(1)
	v463 = int32(4)
	v465 = v364 + v462<<(uint(v463)%32)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v467 = int32(3)
	v470 = *(*int64)(unsafe.Add(mBase, uint32(v466+v439<<(uint(v467)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v465))) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v439))))
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+8)) = uint8(v474)
	v476 = int32(2)
	v477 = v439 + v476
	v480 = v364 + v477<<(uint(v463)%32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v481+v462<<(uint(v467)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v480))) = v485
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462+v487))))
	*(*uint8)(unsafe.Add(mBase, uint32(v480)+8)) = uint8(v489)
	v492 = v444 + v476
	if v492 != v324&int32(2147483646) {
		v439 = v477
		v444 = v492
		goto L97
	} else {
		goto L99
	}
L98:
	;
	if v324&int32(1) == int32(0) {
		goto L92
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v499 = v477
	goto L96
L101:
	;
	if v323 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v570 = v370
	v573 = v373
	v581 = v381
	v590 = v390
	goto L80
L103:
	;
	goto L104
L104:
	;
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v570 = v373
	v573 = v370
	v581 = int32(1)
	v590 = v564
	goto L80
L105:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v573)+8))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	m.T0[v596].(func(*base.Module, int32))(m, v573)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v599+v600<<(uint(int32(2))%32))))
	v605 = int32(1)
	v609 = F_tuplesort_gettupleslot(m, v604, v605, v605, v573, v28+int32(8))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	if v609 != 0 {
		v370 = v570
		v373 = v573
		v381 = v581
		v390 = v590
		goto L74
	} else {
		goto L108
	}
L108:
	;
	goto L75
L109:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)+12))
	m.T0[v637].(func(*base.Module, int32))(m, v614)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v640+v641<<(uint(int32(2))%32))))
	F_tuplesort_end(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L12
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v64)+220))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v648+v649<<(uint(int32(2))%32)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v330
	goto L6
L114:
	;
	goto L5
L115:
	;
	v714 = v28 + int32(32)
	v729 = int32(0)
	goto L118
L116:
	;
	goto L117
L117:
	;
	m.G0 = v28 + int32(1632)
	return
L118:
	;
	v741 = v729 + v31
	v744 = v32 + v729<<(uint(int32(3))%32)
	v747 = l1 + v729*int32(52)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	v751 = l2 + v748<<(uint(int32(4))%32)
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v752&int32(2) != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L117
L120:
	;
	v1267 = v729 + int32(1)
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1267 < v1268 {
		v729 = v1267
		goto L118
	} else {
		goto L219
	}
L121:
	;
	v755 = int32(_a_F_finalize_aggregates_0)
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v760
	v764 = v757 + v748*int32(240)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+20))
	if v765 != 0 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v864 = int32(0)
	v865 = int32(_a_F_finalize_aggregates_0)
	v866 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v869)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v870
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v747)+44))
	if v874 == v864 {
		goto L162
	} else {
		goto L163
	}
L124:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764)+70)))
	if v767 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	goto L126
L126:
	;
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	if v837 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v804)+24)) = v805
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	v808 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v804)+16)) = uint8(v808)
	*(*uint8)(unsafe.Add(mBase, uint32(v804)+32)) = uint8(v807)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v804)))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)))
	v813 = m.T0[v812].(func(*base.Module, int32) int64)(m, v804)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L12
	} else {
		goto L143
	}
L128:
	;
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v792 = base.I32_wrap_i64(v790)
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792))))
	if v793 != int32(1) {
		v803 = v790
		goto L140
	} else {
		goto L141
	}
L129:
	;
	v789 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v804 = v788
	v805 = v789
	goto L127
L130:
	;
	v785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v764)+188)))
	if v785 == int32(_a_F_finalize_aggregates_1) {
		goto L128
	} else {
		goto L138
	}
L131:
	;
	if v766&int32(1) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v764)+228))
	if v766&int32(1) != 0 {
		v788 = v781
		goto L129
	} else {
		goto L137
	}
L134:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v764)+228))
	v784 = v774
	goto L130
L135:
	;
	goto L136
L136:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v744))) = int64(0)
	v777 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v777)
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v756
	goto L120
L137:
	;
	v784 = v781
	goto L130
L138:
	;
	v788 = v784
	goto L129
L139:
	;
	v804 = v784
	v805 = v803
	goto L127
L140:
	;
	goto L139
L141:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792)+1)))
	if v796 != int32(3) {
		v803 = v790
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v792)+2))
	v803 = base.I64_extend_i32_u(v799 + int32(18))
	goto L140
L143:
	;
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v815)
	if v815 != 0 {
		v833 = v813
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v744))) = v833
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v756
	goto L120
L145:
	;
	v817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v747)+48)))
	if v817 != int32(_a_F_finalize_aggregates_1) {
		v833 = v813
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v821 = base.I32_wrap_i64(v813)
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	if v822 != int32(1) {
		v832 = v813
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v833 = v832
	goto L144
L148:
	;
	goto L147
L149:
	;
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821)+1)))
	if v825 != int32(3) {
		v832 = v813
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v821)+2))
	v832 = base.I64_extend_i32_u(v828 + int32(18))
	goto L148
L151:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v744))) = v858
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v860)
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v756
	goto L120
L152:
	;
	v844 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v846 = base.I32_wrap_i64(v844)
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846))))
	if v847 != int32(1) {
		v857 = v844
		goto L158
	} else {
		goto L159
	}
L153:
	;
	v840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v764)+188)))
	if v840 == int32(_a_F_finalize_aggregates_1) {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v843 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v858 = v843
	goto L151
L156:
	;
	goto L155
L157:
	;
	v858 = v857
	goto L151
L158:
	;
	goto L157
L159:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+1)))
	if v850 != int32(3) {
		v857 = v844
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v846)+2))
	v857 = base.I64_extend_i32_u(v853 + int32(18))
	goto L158
L161:
	;
	v960 = v867 + v748*int32(240)
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v747)+8))
	if v961 != 0 {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	v941 = int32(1)
	v943 = v864
	goto L161
L163:
	;
	goto L164
L164:
	;
	v878 = int32(1)
	v879 = int32(0)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	if v880 <= v879 {
		v941 = v878
		v943 = v864
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v888 = v879
	v889 = v878
	v891 = v864
	goto L166
L166:
	;
	v910 = v714 + v889<<(uint(int32(4))%32)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v874)+12))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v911+v888<<(uint(int32(2))%32))))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v915)+24))
	v920 = m.T0[v919].(func(*base.Module, int32, int32, int32) int64)(m, v915, v916, v910+int32(8))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L12
	} else {
		goto L168
	}
L167:
	;
	v941 = v930
	v943 = v928
	goto L161
L168:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v910))) = v920
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910)+8)))
	v924 = int32(1)
	v928 = base.B2i32(v923|v891&v924 != int32(0))
	v930 = v889 + v924
	v932 = v888 + v924
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v874)+4))
	if v932 < v933 {
		v888 = v932
		v889 = v930
		v891 = v928
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v866
	goto L120
L171:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v747)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v747
	v964 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v747 + int32(12)
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v960)+116))
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+26)) = uint16(v962)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)) = uint8(v964)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v970
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	if v975 == v964 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	goto L173
L173:
	;
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	if v1189 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+40)) = uint8(v997)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v998
	if v962 <= v941 {
		goto L185
	} else {
		goto L186
	}
L175:
	;
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v984 = base.I32_wrap_i64(v982)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984))))
	if v985 != int32(1) {
		v995 = v982
		goto L181
	} else {
		goto L182
	}
L176:
	;
	v978 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v960)+188)))
	if v978 == int32(_a_F_finalize_aggregates_1) {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v981 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v997 = v975
	v998 = v981
	goto L174
L179:
	;
	goto L178
L180:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	v997 = v996
	v998 = v995
	goto L174
L181:
	;
	goto L180
L182:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984)+1)))
	if v988 != int32(3) {
		v995 = v982
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v984)+2))
	v995 = base.I64_extend_i32_u(v991 + int32(18))
	goto L181
L184:
	;
	v1145 = int32(1)
	v1147 = int32(0)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1149)+10)))
	if base.B2i32(v1128&v1145 == v1147)|base.B2i32(v1150 != v1145) == v1147 {
		goto L198
	} else {
		goto L199
	}
L185:
	;
	v1128 = v997 | v943
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1006 = (v962 - v941) & int32(3)
	if v1006 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1010 = int32(0)
	v1012 = v941
	goto L191
L189:
	;
	v1049 = v941
	goto L190
L190:
	;
	v1069 = int32(1)
	if base.Ui32(int32(-4)) < base.Ui32(v941-v962) {
		v1128 = v1069
		goto L184
	} else {
		goto L194
	}
L191:
	;
	v1034 = v714 + v1012<<(uint(int32(4))%32)
	v1035 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1034)+8)) = uint8(v1035)
	*(*int64)(unsafe.Add(mBase, uint32(v1034))) = int64(0)
	v1040 = v1012 + v1035
	v1042 = v1010 + v1035
	if v1042 != v1006 {
		v1010 = v1042
		v1012 = v1040
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v1049 = v1040
	goto L190
L193:
	;
	goto L192
L194:
	;
	v1078 = v1049
	goto L195
L195:
	;
	v1098 = int32(4)
	v1100 = v714 + v1078<<(uint(v1098)%32)
	v1101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1100)+8)) = uint8(v1101)
	v1103 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1100))) = v1103
	*(*uint8)(unsafe.Add(mBase, uint32(v1100)+56)) = uint8(v1101)
	*(*int64)(unsafe.Add(mBase, uint32(v1100)+48)) = v1103
	*(*uint8)(unsafe.Add(mBase, uint32(v1100)+40)) = uint8(v1101)
	*(*int64)(unsafe.Add(mBase, uint32(v1100)+32)) = v1103
	*(*uint8)(unsafe.Add(mBase, uint32(v1100)+24)) = uint8(v1101)
	*(*int64)(unsafe.Add(mBase, uint32(v1100)+16)) = v1103
	v1118 = v1078 + v1098
	if v1118 != v962 {
		v1078 = v1118
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v1128 = v1069
	goto L184
L197:
	;
	goto L196
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v744))) = int64(0)
	v1158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v1158)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L170
L199:
	;
	goto L200
L200:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1149)))
	v1165 = m.T0[v1164].(func(*base.Module, int32) int64)(m, v28+int32(8))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L12
	} else {
		goto L201
	}
L201:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v1167)
	if v1167 != 0 {
		v1185 = v1165
		goto L202
	} else {
		goto L203
	}
L202:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v744))) = v1185
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L170
L203:
	;
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v747)+48)))
	if v1169 != int32(_a_F_finalize_aggregates_1) {
		v1185 = v1165
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1173 = base.I32_wrap_i64(v1165)
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173))))
	if v1174 != int32(1) {
		v1184 = v1165
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v1185 = v1184
	goto L202
L206:
	;
	goto L205
L207:
	;
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173)+1)))
	if v1177 != int32(3) {
		v1184 = v1165
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+2))
	v1184 = base.I64_extend_i32_u(v1180 + int32(18))
	goto L206
L209:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v744))) = v1210
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v1212)
	goto L170
L210:
	;
	v1196 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v1198 = base.I32_wrap_i64(v1196)
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1198))))
	if v1199 != int32(1) {
		v1209 = v1196
		goto L216
	} else {
		goto L217
	}
L211:
	;
	v1192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v960)+188)))
	if v1192 == int32(_a_F_finalize_aggregates_1) {
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1195 = *(*int64)(unsafe.Add(mBase, uint32(v751)))
	v1210 = v1195
	goto L209
L214:
	;
	goto L213
L215:
	;
	v1210 = v1209
	goto L209
L216:
	;
	goto L215
L217:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1198)+1)))
	if v1202 != int32(3) {
		v1209 = v1196
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+2))
	v1209 = base.I64_extend_i32_u(v1205 + int32(18))
	goto L216
L219:
	;
	goto L119
}
func F_findNotNullConstraintAttnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v14 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v8 + int32(-56)
	F_ScanKeyInit(m, v19, int32(9), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(1)
	v30 = F_systable_beginscan(m, v14, int32(2665), v27, int32(0), v27, v19)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_systable_endscan(m, v30)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L5:
	;
	v32 = F_systable_getnext(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v32 == int32(0) {
		v62 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = v32
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v44)+72)))
	if v46 != int32(110) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = v3
	goto L4
L10:
	;
	v54 = F_systable_getnext(m, v30)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v49 = F_extractNotNullColumn(m, v36)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v49 != l1 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v52 = F_heap_copytuple(m, v36)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v62 = v52
	goto L4
L15:
	;
	if v54 != 0 {
		v36 = v54
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 - int32(-64)
	return v62
}
func F_find_childrel_parents(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l1
	v14 = int32(0)
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L9
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16<<(uint(int32(2))%32))))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v22 = F_bms_add_member(m, v14, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return v22
L4:
	;
	return int32(0)
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v26) <= base.Ui32(v21) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v21<<(uint(int32(2))%32))))
	if v32 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 == int32(2) {
		v11 = v32
		v14 = v22
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v21
	F_errmsg_internal(m, int32(_a_F_find_childrel_parents_0), v8)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_find_childrel_parents_1), int32(556), int32(_a_F_find_childrel_parents_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_funcstat_entry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_find_funcstat_entry[0]))
	v6 = F_pgstat_fetch_pending_entry(m, int32(3), v4, base.I64_extend_i32_u(l0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v12 = v10
		} else {
			v12 = int32(0)
		}
		return v12
	}
}
func F_find_header(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	v10 = int32(-101)
	if base.Ui32(l1) <= base.Ui32(l0) {
		v278 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v278
L2:
	;
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = int32(8)
	goto L5
L4:
	;
	v14 = int32(10)
	goto L5
L5:
	;
	if l1-l0 < v14 {
		v278 = v10
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = int32(_a_F_find_header_0)
	goto L9
L8:
	;
	v19 = int32(_a_F_find_header_1)
	goto L9
L9:
	;
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19))))
	v24 = l0
	goto L10
L10:
	;
	v30 = l1 - v24
	v31 = int32(0)
	if base.B2i32(v24&int32(3) == v31)|base.B2i32(v30 == v31) != 0 {
		v61 = v24
		v63 = v30
		v64 = base.B2i32(v30 != v31)
		goto L15
	} else {
		goto L16
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
	if base.Ui32(l1) <= base.Ui32(v138) {
		v236 = v138
		goto L66
	} else {
		goto L67
	}
L12:
	;
	if v135 == int32(0) {
		v278 = v10
		goto L1
	} else {
		goto L37
	}
L13:
	;
	v135 = int32(0)
	goto L12
L14:
	;
	v113 = v106
	v115 = v108
	goto L31
L15:
	;
	if v64 == int32(0) {
		goto L13
	} else {
		goto L22
	}
L16:
	;
	v44 = v24
	v46 = v30
	goto L17
L17:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v49 == v20&int32(255) {
		v106 = v44
		v108 = v46
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v61 = v56
	v63 = v52
	v64 = v54
	goto L15
L19:
	;
	v51 = int32(1)
	v52 = v46 - v51
	v53 = int32(0)
	v54 = base.B2i32(v52 != v53)
	v56 = v44 + v51
	if v56&int32(3) == v53 {
		v61 = v56
		v63 = v52
		v64 = v54
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v52 != 0 {
		v44 = v56
		v46 = v52
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v69 = v20 & int32(255)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if base.B2i32(v69 == v70)|base.B2i32(base.Ui32(v63) < base.Ui32(int32(4))) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = v61
	v81 = v63
	goto L26
L24:
	;
	v99 = v61
	v101 = v63
	goto L25
L25:
	;
	if v101 == int32(0) {
		goto L13
	} else {
		goto L30
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v86 = v85 ^ v69*int32(16843009)
	v89 = int32(-2139062144)
	if (int32(16843008)-v86|v86)&v89 != v89 {
		v106 = v79
		v108 = v81
		goto L14
	} else {
		goto L28
	}
L27:
	;
	v99 = v94
	v101 = v96
	goto L25
L28:
	;
	v93 = int32(4)
	v94 = v79 + v93
	v96 = v81 - v93
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v79 = v94
		v81 = v96
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v106 = v99
	v108 = v101
	goto L14
L31:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v20&int32(255) == v118 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L13
L33:
	;
	v135 = v113
	goto L12
L34:
	;
	goto L35
L35:
	;
	v120 = int32(1)
	v123 = v115 - v120
	if v123 != 0 {
		v113 = v113 + v120
		v115 = v123
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v138 = v135 + v14
	if base.Ui32(l1) < base.Ui32(v138) {
		v278 = v10
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v14) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v201 != 0 {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	v201 = int32(0)
	goto L39
L41:
	;
	v175 = v170
	v176 = v171
	v177 = v172
	goto L51
L42:
	;
	if (v135|v19)&int32(3) != 0 {
		v170 = v135
		v171 = v19
		v172 = v14
		goto L41
	} else {
		goto L45
	}
L43:
	;
	v163 = v135
	v164 = v19
	v165 = v14
	goto L44
L44:
	;
	if v165 == int32(0) {
		goto L40
	} else {
		goto L50
	}
L45:
	;
	v147 = v135
	v148 = v19
	v149 = v14
	goto L46
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v152 != v153 {
		v170 = v147
		v171 = v148
		v172 = v149
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v163 = v158
	v164 = v156
	v165 = v160
	goto L44
L48:
	;
	v155 = int32(4)
	v156 = v148 + v155
	v158 = v147 + v155
	v160 = v149 - v155
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v147 = v158
		v148 = v156
		v149 = v160
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v170 = v163
	v171 = v164
	v172 = v165
	goto L41
L51:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v180 == v181 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v201 = v180 - v181
	goto L39
L53:
	;
	v183 = int32(1)
	v188 = v177 - v183
	if v188 != 0 {
		v175 = v175 + v183
		v176 = v176 + v183
		v177 = v188
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L40
L57:
	;
	v203 = v135 + int32(1)
	if base.Ui32(v203) < base.Ui32(l1) {
		v24 = v203
		goto L10
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if l0 == v135 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v278 = v10
	goto L1
L61:
	;
	goto L11
L62:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135-int32(1)))))
	if v208 == int32(10) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(l1) <= base.Ui32(v138) {
		v278 = v10
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v14 <= l1-v138 {
		v24 = v138
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v278 = v10
	goto L1
L66:
	;
	if l1-v236 < int32(5) {
		v278 = v10
		goto L1
	} else {
		goto L73
	}
L67:
	;
	v219 = v138
	goto L68
L68:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v225 == int32(45) {
		v236 = v219
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v236 = l1
	goto L66
L70:
	;
	if base.Ui32(v225) < base.Ui32(int32(32)) {
		v278 = v10
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v231 = v219 + int32(1)
	if base.Ui32(v231) < base.Ui32(l1) {
		v219 = v231
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+4)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)))
	if v245^v246|(v248^v249) != 0 {
		v278 = v10
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v253 = v236 + int32(5)
	if base.Ui32(l1) <= base.Ui32(v253) {
		v268 = v253
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v278 = v268 - v135
	goto L1
L76:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	switch v255 - int32(10) {
	case 0, 3:
		goto L77
	default:
		v278 = v10
		goto L1
	}
L77:
	;
	if v255 == int32(13) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v262 = v236 + int32(6)
	goto L80
L79:
	;
	v262 = v253
	goto L80
L80:
	;
	if base.Ui32(l1) <= base.Ui32(v262) {
		v268 = v262
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	v268 = v262 + base.B2i32(v264 == int32(10))
	goto L75
}
func F_find_jointree_node_for_rel(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		v74 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v74
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(63) {
	case 0:
		goto L4
	case 1:
		goto L6
	case 2:
		goto L7
	default:
		goto L5
	}
L3:
	;
	v74 = int32(0)
	goto L1
L4:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 == v64 {
		v74 = l0
		goto L1
	} else {
		goto L26
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L23
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v40 == l1 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= v18 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = v18
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v22<<(uint(int32(2))%32))))
	v32 = F_find_jointree_node_for_rel(m, v31, l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L3
L12:
	;
	return int32(0)
L13:
	;
	if v32 != 0 {
		v74 = v32
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v37 = v22 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v37 < v38 {
		v22 = v37
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v74 = l0
	goto L1
L17:
	;
	goto L18
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = F_find_jointree_node_for_rel(m, v42, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if v43 != 0 {
		v74 = v43
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v46 = F_find_jointree_node_for_rel(m, v45, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v74 = v46
	goto L1
L23:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v54
	F_errmsg_internal(m, int32(_a_F_find_jointree_node_for_rel_0), v8)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_find_jointree_node_for_rel_1), int32(_a_F_find_jointree_node_for_rel_2), int32(_a_F_find_jointree_node_for_rel_3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	goto L3
}
func F_finish_nodeitem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(l2 == int32(0))|base.B2i32(base.Ui32(l1) <= base.Ui32(v16)) != 0 {
		v150 = l1
		v152 = l3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v150 - v16
	if v150 == v16 {
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v22 = l1
	v24 = l3
	goto L3
L3:
	;
	v32 = v22 - int32(1)
	v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
	goto L9
L4:
	;
	v150 = v16
	v152 = l3 + v16 - l1
	goto L1
L5:
	;
	if v139 == int32(0) {
		v150 = v22
		v152 = v24
		goto L1
	} else {
		goto L30
	}
L6:
	;
	v139 = int32(0)
	goto L5
L7:
	;
	v117 = v110
	v119 = v112
	goto L24
L8:
	;
	if base.B2i32(v56 != v57) == int32(0) {
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v48 = int32(_a_F_finish_nodeitem_0)
	v50 = int32(4)
	goto L10
L10:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v53 == v33&int32(255) {
		v110 = v48
		v112 = v50
		goto L7
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v55 = int32(1)
	v56 = v50 - v55
	v57 = int32(0)
	v60 = v48 + v55
	if v60&int32(3) == v57 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v56 != 0 {
		v48 = v60
		v50 = v56
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v73 = v33 & int32(255)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if base.B2i32(v73 == v74)|base.B2i32(base.Ui32(v56) < base.Ui32(int32(4))) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v83 = v60
	v85 = v56
	goto L19
L17:
	;
	v103 = v60
	v105 = v56
	goto L18
L18:
	;
	if v105 == int32(0) {
		goto L6
	} else {
		goto L23
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v90 = v89 ^ v73*int32(16843009)
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 != v93 {
		v110 = v83
		v112 = v85
		goto L7
	} else {
		goto L21
	}
L20:
	;
	v103 = v98
	v105 = v100
	goto L18
L21:
	;
	v97 = int32(4)
	v98 = v83 + v97
	v100 = v85 - v97
	if base.Ui32(int32(3)) < base.Ui32(v100) {
		v83 = v98
		v85 = v100
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v110 = v103
	v112 = v105
	goto L7
L24:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v33&int32(255) == v122 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L6
L26:
	;
	v139 = v117
	goto L5
L27:
	;
	goto L28
L28:
	;
	v124 = int32(1)
	v127 = v119 - v124
	if v127 != 0 {
		v117 = v117 + v124
		v119 = v127
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v142 - v143
	if base.Ui32(v16) < base.Ui32(v32) {
		v22 = v32
		v24 = v24 - v143
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L4
L32:
	;
	m.G0 = v12 + int32(32)
	return v217
L33:
	;
	F_errsave_finish(m, l4, int32(_a_F_finish_nodeitem_1), v212, int32(_a_F_finish_nodeitem_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L37
	} else {
		goto L52
	}
L34:
	;
	v162 = int32(0)
	v163 = F_errsave_start(m, l4)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v184 < int32(1001) {
		v217 = int32(1)
		goto L32
	} else {
		goto L46
	}
L37:
	;
	return int32(0)
L38:
	;
	if v163 == int32(0) {
		v217 = v162
		goto L32
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v152
	if l2 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v175 = int32(_a_F_finish_nodeitem_3)
	goto L43
L42:
	;
	v175 = int32(_a_F_finish_nodeitem_4)
	goto L43
L43:
	;
	F_errmsg(m, v175, v12)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v180 = F_errdetail(m, int32(_a_F_finish_nodeitem_5), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v211 = v162
	v212 = int32(627)
	goto L33
L46:
	;
	v187 = int32(0)
	v188 = F_errsave_start(m, l4)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L37
	} else {
		goto L47
	}
L47:
	;
	if v188 == int32(0) {
		v217 = v187
		goto L32
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(34103428))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L37
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_finish_nodeitem_6), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L37
	} else {
		goto L50
	}
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1000)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v199
	v207 = F_errdetail(m, int32(_a_F_finish_nodeitem_7), v12+int32(16))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v211 = v187
	v212 = int32(633)
	goto L33
L52:
	;
	v217 = v211
	goto L32
}
func F_fix_indexqual_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v8 = F_strip_noop_phvs(m, l0)
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v12 == int32(27) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = v8
	goto L6
L4:
	;
	v26 = v12
	v30 = v8
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+l2<<(uint(int32(2))%32))))
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 == int32(27) {
		v19 = v22
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v26 = v23
	v30 = v22
	goto L5
L8:
	;
	goto L7
L9:
	;
	if v26 != int32(6) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
	if v40 != v42 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+8)))
	if v37 != v44 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v46 = F_copyObjectImpl(m, v30)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v49 = l2 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v46)+8)) = uint16(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(-3)
	return v46
L17:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_0), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_2), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v70 = v68
	goto L22
L21:
	;
	v70 = int32(0)
	goto L22
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if int32(0) < v71 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L58
	}
L24:
	;
	v75 = int32(0)
	v78 = v70
	goto L27
L25:
	;
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L55
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v33+v75<<(uint(int32(2))%32))))
	if v85 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	if v78 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L30:
	;
	v140 = v78
	goto L31
L31:
	;
	v142 = v75 + int32(1)
	if v142 != v71 {
		v75 = v142
		v78 = v140
		goto L27
	} else {
		goto L54
	}
L32:
	;
	if v75 == l2 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v91 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v131 = v78 + int32(4)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if base.Ui32(v131) < base.Ui32(v133+v134<<(uint(int32(2))%32)) {
		goto L51
	} else {
		goto L52
	}
L36:
	;
	v100 = F_equal(m, v30, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L41
	}
L37:
	;
	v99 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v95 != int32(27) {
		v99 = v91
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v99 = v98
	goto L36
L41:
	;
	if v100 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v107 = F_exprType(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v111 = F_exprCollation(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v114 = F_makeVar(m, int32(-3), base.I32_extend16_s(l2+int32(1)), v107, int32(-1), v111, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	return v114
L48:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_0), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_4), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v139 = v131
	goto L53
L52:
	;
	v139 = int32(0)
	goto L53
L53:
	;
	v140 = v139
	goto L31
L54:
	;
	goto L28
L55:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_0), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_5), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_6), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_7), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fixed_paramref_hook(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685636))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
				F_errmsg(m, int32(_a_F_fixed_paramref_hook_0), v9)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					F_parser_errposition(m, l0, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_fixed_paramref_hook_1), int32(112), int32(_a_F_fixed_paramref_hook_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v15 < v11 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33685636))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					F_errmsg(m, int32(_a_F_fixed_paramref_hook_0), v9)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						F_parser_errposition(m, l0, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_fixed_paramref_hook_1), int32(112), int32(_a_F_fixed_paramref_hook_2))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
			v20 = (v11 - int32(1)) << (uint(int32(2)) % 32)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21)))
			if v23 != 0 {
				v48 = F_palloc0(m, int32(28))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v11
					*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(8)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v53+v20)))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v55
					v59 = F_get_typcollation(m, v55)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v59
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v62
						m.G0 = v9 + int32(16)
						return v48
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33685636))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg(m, int32(_a_F_fixed_paramref_hook_0), v9)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							F_parser_errposition(m, l0, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_fixed_paramref_hook_1), int32(112), int32(_a_F_fixed_paramref_hook_2))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
			}
		}
	}
}
func F_float48gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	var v22 int64
	_ = v22
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
		v11 = base.F64_promote_f32(v10)
		v22 = base.I64_extend_i32_u(base.F64_lt(v4, v11) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807))))
	} else {
		v22 = int64(0)
	}
	return v22
}
func F_float4_cmp_internal(m *base.Module, l0 float32, l1 float32) int32 {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	v7 = int32(2147483647)
	v8 = base.I32_reinterpret_f32(l0) & v7
	v11 = base.I32_reinterpret_f32(l1) & v7
	if base.Ui32(int32(2139095041)) <= base.Ui32(v11) {
		v21 = base.B2i32(base.Ui32(v8) < base.Ui32(int32(2139095041)))
		v30 = int32(0) - v21&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v11))|base.F32_lt(l0, l1))
	} else {
		v16 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v8))|base.F32_gt(l0, l1) != 0 {
			v30 = v16
		} else {
			v21 = v16
			v30 = int32(0) - v21&(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v11))|base.F32_lt(l0, l1))
		}
	}
	return v30
}
func F_float4larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = base.I32_wrap_i64(v5)
	if base.Ui32(v6&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = base.I32_wrap_i64(v11)
		if base.F32_gt(base.F32_reinterpret_i32(v12), base.F32_reinterpret_i32(v6))|base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v12&int32(2147483647))) != 0 {
			v21 = v11
		} else {
			v21 = v5
		}
		v24 = v21
	} else {
		v24 = v5
	}
	return base.I64_extend32_s(v24)
}
func F_float84le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float64
	_ = v11
	var v22 int64
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v22 = base.I64_extend_i32_u(base.F64_ge(v5, v11) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))))
	} else {
		v22 = int64(1)
	}
	return v22
}
func F_float84ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 float64
	_ = v10
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = base.F64_promote_f32(v5)
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v10)&v8) {
		return base.I64_extend_i32_u(base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313))))
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9)) | base.F64_ne(v6, v10))
	}
}
func F_forkname_to_number(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v2 = int32(_a_F_forkname_to_number_0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_to_number[0])))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 != v8) != 0 {
		v26 = v5
		v27 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L1
L3:
	;
	v11 = v2
	v12 = l0
	goto L4
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v16 == int32(0) {
		v26 = v16
		v27 = v15
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v26 = v16
	v27 = v15
	goto L2
L6:
	;
	v19 = int32(1)
	if v16 == v15 {
		v11 = v11 + v19
		v12 = v12 + v19
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	v33 = int32(_a_F_forkname_to_number_1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_to_number[1])))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v36 == int32(0))|base.B2i32(v36 != v39) != 0 {
		v57 = v36
		v58 = v39
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v57-v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v42 = v33
	v43 = l0
	goto L14
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v47
		v58 = v46
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v57 = v47
	v58 = v46
	goto L12
L16:
	;
	v50 = int32(1)
	if v47 == v46 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	v64 = int32(_a_F_forkname_to_number_2)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_to_number[2])))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v88-v89 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	v73 = v64
	v74 = l0
	goto L24
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v88 = v78
	v89 = v77
	goto L22
L26:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(2)
L29:
	;
	goto L30
L30:
	;
	v95 = int32(_a_F_forkname_to_number_3)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_forkname_to_number[3])))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v98 == int32(0))|base.B2i32(v98 != v101) != 0 {
		v119 = v98
		v120 = v101
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v119-v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	v104 = v95
	v105 = l0
	goto L34
L34:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v109
		v120 = v108
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v119 = v109
	v120 = v108
	goto L32
L36:
	;
	v112 = int32(1)
	if v109 == v108 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	return int32(3)
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return int32(0)
L42:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_forkname_to_number_4), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	F_errhint(m, int32(_a_F_forkname_to_number_5), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_forkname_to_number_6), int32(63), int32(_a_F_forkname_to_number_7))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(_a_F_format_elog_string_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[0]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[0])) = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[2]))
	v21 = v10 + int32(16)
	F_initStringInfo(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[3])) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v29 = F_appendStringInfoVA(m, v21, l0, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = v29
	goto L7
L5:
	;
	goto L6
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v55 = F_pstrdup(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v39 = v10 + int32(16)
	F_enlargeStringInfo(m, v39, v34)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[3])) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v45 = F_appendStringInfoVA(m, v39, l0, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v45 != 0 {
		v34 = v45
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_elog_string[0])) = v13
	m.G0 = v10 + int32(32)
	return v55
}
func F_format_procedure(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_format_procedure_extended(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_format_procedure_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v15 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v120
L2:
	;
	return int32(0)
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
	v21 = v19 + v20
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+104)))
	F_initStringInfo(m, v11+int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l1&int32(1) != 0 {
		v120 = v3
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v30 = l1 & int32(2)
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v40 = F_quote_qualified_identifier(m, v39, v21+int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L15
	}
L9:
	;
	v34 = F_FunctionIsVisibleExt(m, l0, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v37 = F_get_namespace_name(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	if v34 != 0 {
		v39 = v3
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v39 = v37
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v40
	v44 = v11 + int32(32)
	F_appendStringInfo(m, v44, int32(_a_F_format_procedure_extended_0), v11+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v24 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(41))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L38
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+136))
	if v30 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_appendStringInfoString(m, v44, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L25
	}
L20:
	;
	v53 = F_format_type_be_qualified(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v55 = F_format_type_be(m, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v57 = v53
	goto L19
L24:
	;
	v57 = v55
	goto L19
L25:
	;
	v60 = int32(1)
	if v24 == v60 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v65 = v60
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(136)+v65<<(uint(int32(2))%32))))
	v78 = v11 + int32(32)
	F_appendStringInfoChar(m, v78, int32(44))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L29
	}
L28:
	;
	goto L17
L29:
	;
	if v30 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_appendStringInfoString(m, v78, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L36
	}
L31:
	;
	v82 = F_format_type_be_qualified(m, v76)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v84 = F_format_type_be(m, v76)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L35
	}
L34:
	;
	v86 = v82
	goto L30
L35:
	;
	v86 = v84
	goto L30
L36:
	;
	v90 = v65 + int32(1)
	if v90 != v24 {
		v65 = v90
		goto L27
	} else {
		goto L37
	}
L37:
	;
	goto L28
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v120 = v105
	goto L1
L40:
	;
	v111 = F_palloc(m, int32(64))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v116 = F_pg_snprintf(m, v111, int32(64), int32(_a_F_format_procedure_extended_1), v11)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v120 = v111
	goto L1
}
func F_free_attstatsslot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v6 != 0 {
				F_pfree(m, v6)
				mBase = m.M
				v8 = m.ExcPending
				if v8 != 0 {
					return
				} else {
					v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v9 != 0 {
						F_pfree(m, v9)
						mBase = m.M
						v11 = m.ExcPending
						if v11 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v9 != 0 {
					F_pfree(m, v9)
					mBase = m.M
					v11 = m.ExcPending
					if v11 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v6 != 0 {
			F_pfree(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v9 != 0 {
					F_pfree(m, v9)
					mBase = m.M
					v11 = m.ExcPending
					if v11 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v9 != 0 {
				F_pfree(m, v9)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
}
func F_freesubre(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v5 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v8 = v5
	goto L7
L5:
	;
	goto L6
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v17 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	F_freesubre(m, l0, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	if v10 != 0 {
		v8 = v10
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	F_pfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v29)
	v32 = l1 + int32(20)
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33
	if l0 == v29 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pfree(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_pfree(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	goto L14
L18:
	;
	F_pfree(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L21
	}
L19:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v39 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = l1
	return
L21:
	;
	goto L3
}
func F_fsm_page_contents(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = F_superuser(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = F_get_page_from_raw(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
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
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	m.G0 = v9 + int32(48)
	return v72
L8:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+14)))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
	v72 = int64(0)
	goto L7
L10:
	;
	goto L11
L11:
	;
	F_initStringInfo(m, v9+int32(32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+(v18+int32(28))))))
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v54
	F_appendStringInfo(m, v9+int32(32), int32(_a_F_fsm_page_contents_0), v9)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L20
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v33
	F_appendStringInfo(m, v9+int32(32), int32(_a_F_fsm_page_contents_1), v9+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v51 = v33 + int32(1)
	if v51 != int32(_a_F_fsm_page_contents_2) {
		v33 = v51
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	goto L14
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v63 = F_cstring_to_text_with_len(m, v61, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v72 = base.I64_extend_i32_u(v63)
	goto L7
L22:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_fsm_page_contents_3), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_fsm_page_contents_4), int32(46), int32(_a_F_fsm_page_contents_5))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
func F_fsm_set_avail(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v3 = l2
	v9 = l0 + int32(28)
	v11 = l1 + int32(4095)
	v12 = v9 + v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v3)
	v23 = v11
	goto L4
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.Ui32(v15) < base.Ui32(v3) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	v27 = int32(1)
	v28 = v23 - v27
	v29 = int32(2)
	v30 = base.I32_div_s(v28, v29)
	v32 = v30 << (uint(v27) % 32)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v32)+1)))
	v36 = v32 + v29
	if base.Ui32(v36) <= base.Ui32(int32(_a_F_fsm_set_avail_0)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.Ui32(v55) < base.Ui32(v3) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v40 = v34 & int32(255)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v36))))
	if base.Ui32(v42) < base.Ui32(v40) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v45 = v34
	goto L8
L8:
	;
	v47 = v30 + v9
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != v45&int32(255) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v44 = v40
	goto L11
L10:
	;
	v44 = v42
	goto L11
L11:
	;
	v45 = v44
	goto L8
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v45)
	if int32(1) < v28 {
		v23 = v30
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L5
L15:
	;
	goto L14
L16:
	;
	v61 = int32(4094)
	goto L19
L17:
	;
	goto L18
L18:
	;
	return int32(1)
L19:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v61) {
		v82 = int32(0)
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v83 = v61 + v9
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != v82&int32(255) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v71 = v61 << (uint(int32(1)) % 32)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(29)+v71))))
	if v61 == int32(4081) {
		v82 = v73
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v9)+2)))
	if base.Ui32(v77) < base.Ui32(v73) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = v73
	goto L26
L25:
	;
	v79 = v77
	goto L26
L26:
	;
	v82 = v79
	goto L21
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v82)
	goto L29
L28:
	;
	goto L29
L29:
	;
	if v61 != 0 {
		v61 = v61 - int32(1)
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L20
}
func F_fsync(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	v2 = m.Wasi_snapshot_preview1.Fd_sync(m, l0)
	mBase = m.M
	if v2 == int32(0) {
		v9 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_fsync[0])) = v2
		v9 = int32(-1)
	}
	return v9
}
func F_ftod(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 float32
	_ = v2
	v2 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_reinterpret_f64(base.F64_promote_f32(v2))
}
