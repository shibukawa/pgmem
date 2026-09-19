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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+616))
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+624))
	if v64 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
	if v56 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v59 = F_FindLockCycleRecurseMember(m, v9, v9, l1, l2, l3)
	mBase = m.M
	if v59 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	return int32(1)
L20:
	;
	return int32(0)
L21:
	;
	v68 = v9 + int32(620)
	if v64 == v68 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v70 = v64
	goto L23
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(624))))
	if v79 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v96 != v68 {
		v70 = v96
		goto L23
	} else {
		goto L30
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70-int32(536))))
	if v84 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v88 = v70 - int32(628)
	if v88 == v9 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v90 = F_FindLockCycleRecurseMember(m, v88, v9, l1, l2, l3)
	mBase = m.M
	if v90 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	return int32(1)
L30:
	;
	goto L24
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
	F_errfinish(m, int32(_a_F_FlagRWConflict_3), int32(654), int32(_a_F_FlagRWConflict_11))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v36 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v34 = v31
	v35 = v32
	goto L1
L3:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v26 = v19
	goto L5
L5:
	;
	v28 = l0 + int32(80)
	if v26 != 0 {
		v30 = v26
		v31 = v28
		goto L2
	} else {
		goto L11
	}
L6:
	;
	v30 = v19
	v31 = l0 + int32(80)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v23 = F_ExecPrepareExprList(m, v20, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v23
	v26 = v23
	goto L5
L11:
	;
	v34 = v28
	v35 = int32(0)
	goto L1
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L48
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L45
	}
L14:
	;
	v50 = int32(0)
	v52 = v35
	goto L17
L15:
	;
	v148 = v35
	goto L16
L16:
	;
	if v148 != 0 {
		goto L12
	} else {
		goto L44
	}
L17:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(12)+v50<<(uint(int32(1))%32)))))
	v62 = base.I32_extend16_s(v61)
	if v62 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v148 = v126
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v50<<(uint(int32(2))%32)))) = v125
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(l4+v50))) = uint8(v134)
	v137 = v50 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v137 < v138 {
		v50 = v137
		v52 = v126
		goto L17
	} else {
		goto L43
	}
L20:
	;
	switch v61 - int32(_a_F_FormIndexDatum_0) {
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
	if v62 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	v76 = m.T0[v75].(func(*base.Module, int32, int32, int32) int32)(m, l1, v62, v17+int32(15))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v70)
	v125 = l1 + int32(28)
	v126 = v52
	goto L19
L25:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v125 = v69
	v126 = v52
	goto L19
L26:
	;
	v125 = v76
	v126 = v52
	goto L19
L27:
	;
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v78 < v62 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v52 == int32(0) {
		goto L13
	} else {
		goto L34
	}
L30:
	;
	F_slot_getsomeattrs_int(m, l1, v62)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v83 = v62 - int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v86)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v83<<(uint(int32(2))%32))))
	v125 = v92
	v126 = v52
	goto L19
L33:
	;
	goto L32
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v96 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v99 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	v101 = v96
	goto L37
L37:
	;
	v102 = int32(_a_F_FormIndexDatum_1)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_FormIndexDatum[0]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_FormIndexDatum[0])) = v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v110 = m.T0[v109].(func(*base.Module, int32, int32, int32) int32)(m, v95, v101, v17+int32(15))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L39
	}
L38:
	;
	v101 = v99
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FormIndexDatum[0])) = v103
	v115 = v52 + int32(4)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if base.Ui32(v115) < base.Ui32(v118+v119<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v124 = v115
	goto L42
L41:
	;
	v124 = int32(0)
	goto L42
L42:
	;
	v125 = v110
	v126 = v124
	goto L19
L43:
	;
	goto L18
L44:
	;
	m.G0 = v17 + int32(16)
	return
L45:
	;
	F_errmsg_internal(m, int32(_a_F_FormIndexDatum_2), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_FormIndexDatum_3), int32(2772), int32(_a_F_FormIndexDatum_4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_FormIndexDatum_3), int32(2783), int32(_a_F_FormIndexDatum_4))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
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
	F_errfinish(m, int32(_a_F_FreeDir_1), int32(3050), int32(_a_F_FreeDir_2))
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
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13898(m, l0, l1, l2, l3, int32(_a_F_fastgetattr_3_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
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
								F_errfinish(m, int32(_a_F_fetch_finfo_record_2), int32(481), int32(_a_F_fetch_finfo_record_3))
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
										F_errfinish(m, int32(_a_F_fetch_finfo_record_2), int32(491), int32(_a_F_fetch_finfo_record_3))
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
								F_errfinish(m, int32(_a_F_fetch_finfo_record_2), int32(472), int32(_a_F_fetch_finfo_record_3))
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
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v10 = l0 + l1<<(uint(int32(2))%32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+192))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v127
L2:
	;
	v94 = v10 + int32(192)
	v96 = F_palloc0(m, int32(272))
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
		v127 = v29
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
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = int64(17179869452)
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
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	v106 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+25)) = uint16(v106)
	v109 = base.F64_gt(v105, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)) = uint8(v109)
	v111 = F_create_empty_pathtarget(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v113 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+44)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v96)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+28)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v96)+52)) = v113
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v121 = F_lappend(m, v120, v96)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v121
	v127 = v96
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
func F_fill_val(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
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
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
			if v40 == int32(1) {
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				v45 = int32(1)
				v49 = (v14 + v43 - v45) & (int32(0) - v43)
				v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				switch v50 - v45 {
				case 0:
					*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v6)
					v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
					v183 = v72
					v185 = v49
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
					m.G0 = v12 + int32(16)
					return
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v6)
					v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
					v183 = v54
					v185 = v49
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
					m.G0 = v12 + int32(16)
					return
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = base.I32_extend16_s(v50)
						F_errmsg_internal(m, int32(_a_F_fill_val_0), v12)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_fill_val_1), int32(230), int32(_a_F_fill_val_2))
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
				case 3:
					*(*int32)(unsafe.Add(mBase, uint32(v49))) = v6
					v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
					v183 = v56
					v185 = v49
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				switch v73 - int32(_a_F_fill_val_3) {
				case 0:
					v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
					v162 = v160 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v162)
					v164 = F_strlen(m, v6)
					mBase = m.M
					v166 = v164 + int32(1)
					if v166 != 0 {
						v178 = v166
						v180 = v14
						base.MemoryCopy(m, v180, v6, v178)
						v183 = v178
						v185 = v180
					} else {
						v183 = v166
						v185 = v14
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
					m.G0 = v12 + int32(16)
					return
				case 1:
					v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
					v78 = v76 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v78)
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
					if v80 == int32(1) {
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
						if v83&int32(254) == int32(2) {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v6)+2))
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
							v95 = (v14 + v89 - int32(1)) & (int32(0) - v89)
							v96 = F_EOH_get_flat_size(m, v88)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								F_EOH_flatten_into(m, v88, v95, v96)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v183 = v96
									v185 = v95
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
									m.G0 = v12 + int32(16)
									return
								}
							}
						} else {
							v100 = int32(6)
							v101 = v76 | v100
							*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v101)
							v104 = int32(18)
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
							if v106 == v104 {
								v109 = v104
							} else {
								v109 = int32(2)
							}
							if base.Ui32((v106-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v116 = v100
							} else {
								v116 = v109
							}
							if v116 != 0 {
								v178 = v116
								v180 = v14
								base.MemoryCopy(m, v180, v6, v178)
								v183 = v178
								v185 = v180
							} else {
								v183 = v116
								v185 = v14
							}
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						if v80&int32(1) != 0 {
							v120 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if v120 != 0 {
								v178 = v120
								v180 = v14
								base.MemoryCopy(m, v180, v6, v178)
								v183 = v178
								v185 = v180
							} else {
								v183 = v120
								v185 = v14
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							v122 = int32(2)
							v123 = int32(base.Ui32(v121) >> (uint(v122) % 32))
							if v80&v122 != 0 {
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
								v154 = int32(0)
								v156 = (v14 + v150 - int32(1)) & (v154 - v150)
								if v123 == v154 {
									v183 = v123
									v185 = v156
								} else {
									base.MemoryCopy(m, v156, v6, v123)
									v183 = v123
									v185 = v156
								}
							} else {
								v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
								if v126&int32(1) == int32(0) {
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
									v154 = int32(0)
									v156 = (v14 + v150 - int32(1)) & (v154 - v150)
									if v123 == v154 {
										v183 = v123
										v185 = v156
									} else {
										base.MemoryCopy(m, v156, v6, v123)
										v183 = v123
										v185 = v156
									}
								} else {
									v132 = v123 - int32(3)
									if base.Ui32(int32(127)) < base.Ui32(v132) {
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
										v154 = int32(0)
										v156 = (v14 + v150 - int32(1)) & (v154 - v150)
										if v123 == v154 {
											v183 = v123
											v185 = v156
										} else {
											base.MemoryCopy(m, v156, v6, v123)
											v183 = v123
											v185 = v156
										}
									} else {
										v135 = int32(1)
										v138 = v132<<(uint(v135)%32) | v135
										*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v138)
										v141 = v123 - int32(4)
										if v141 == int32(0) {
											v183 = v132
											v185 = v14
										} else {
											base.MemoryCopy(m, v14+int32(1), v6+int32(4), v141)
											v183 = v132
											v185 = v14
										}
									}
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
						m.G0 = v12 + int32(16)
						return
					}
				default:
					v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					v171 = int32(0)
					v173 = (v14 + v167 - int32(1)) & (v171 - v167)
					v174 = base.I32_extend16_s(v73)
					if v174 == v171 {
						v183 = v174
						v185 = v173
					} else {
						v178 = v174
						v180 = v173
						base.MemoryCopy(m, v180, v6, v178)
						v183 = v178
						v185 = v180
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	} else {
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
		if v40 == int32(1) {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			v45 = int32(1)
			v49 = (v14 + v43 - v45) & (int32(0) - v43)
			v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			switch v50 - v45 {
			case 0:
				*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v6)
				v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				v183 = v72
				v185 = v49
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
				m.G0 = v12 + int32(16)
				return
			case 1:
				*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v6)
				v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				v183 = v54
				v185 = v49
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
				m.G0 = v12 + int32(16)
				return
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = base.I32_extend16_s(v50)
					F_errmsg_internal(m, int32(_a_F_fill_val_0), v12)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_fill_val_1), int32(230), int32(_a_F_fill_val_2))
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
			case 3:
				*(*int32)(unsafe.Add(mBase, uint32(v49))) = v6
				v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				v183 = v56
				v185 = v49
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			switch v73 - int32(_a_F_fill_val_3) {
			case 0:
				v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
				v162 = v160 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v162)
				v164 = F_strlen(m, v6)
				mBase = m.M
				v166 = v164 + int32(1)
				if v166 != 0 {
					v178 = v166
					v180 = v14
					base.MemoryCopy(m, v180, v6, v178)
					v183 = v178
					v185 = v180
				} else {
					v183 = v166
					v185 = v14
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
				m.G0 = v12 + int32(16)
				return
			case 1:
				v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
				v78 = v76 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v78)
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
				if v80 == int32(1) {
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
					if v83&int32(254) == int32(2) {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v6)+2))
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
						v95 = (v14 + v89 - int32(1)) & (int32(0) - v89)
						v96 = F_EOH_get_flat_size(m, v88)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_EOH_flatten_into(m, v88, v95, v96)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v183 = v96
								v185 = v95
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
								m.G0 = v12 + int32(16)
								return
							}
						}
					} else {
						v100 = int32(6)
						v101 = v76 | v100
						*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v101)
						v104 = int32(18)
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
						if v106 == v104 {
							v109 = v104
						} else {
							v109 = int32(2)
						}
						if base.Ui32((v106-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v116 = v100
						} else {
							v116 = v109
						}
						if v116 != 0 {
							v178 = v116
							v180 = v14
							base.MemoryCopy(m, v180, v6, v178)
							v183 = v178
							v185 = v180
						} else {
							v183 = v116
							v185 = v14
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					if v80&int32(1) != 0 {
						v120 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
						if v120 != 0 {
							v178 = v120
							v180 = v14
							base.MemoryCopy(m, v180, v6, v178)
							v183 = v178
							v185 = v180
						} else {
							v183 = v120
							v185 = v14
						}
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v122 = int32(2)
						v123 = int32(base.Ui32(v121) >> (uint(v122) % 32))
						if v80&v122 != 0 {
							v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
							v154 = int32(0)
							v156 = (v14 + v150 - int32(1)) & (v154 - v150)
							if v123 == v154 {
								v183 = v123
								v185 = v156
							} else {
								base.MemoryCopy(m, v156, v6, v123)
								v183 = v123
								v185 = v156
							}
						} else {
							v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
							if v126&int32(1) == int32(0) {
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
								v154 = int32(0)
								v156 = (v14 + v150 - int32(1)) & (v154 - v150)
								if v123 == v154 {
									v183 = v123
									v185 = v156
								} else {
									base.MemoryCopy(m, v156, v6, v123)
									v183 = v123
									v185 = v156
								}
							} else {
								v132 = v123 - int32(3)
								if base.Ui32(int32(127)) < base.Ui32(v132) {
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
									v154 = int32(0)
									v156 = (v14 + v150 - int32(1)) & (v154 - v150)
									if v123 == v154 {
										v183 = v123
										v185 = v156
									} else {
										base.MemoryCopy(m, v156, v6, v123)
										v183 = v123
										v185 = v156
									}
								} else {
									v135 = int32(1)
									v138 = v132<<(uint(v135)%32) | v135
									*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v138)
									v141 = v123 - int32(4)
									if v141 == int32(0) {
										v183 = v132
										v185 = v14
									} else {
										base.MemoryCopy(m, v14+int32(1), v6+int32(4), v141)
										v183 = v132
										v185 = v14
									}
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
					m.G0 = v12 + int32(16)
					return
				}
			default:
				v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				v171 = int32(0)
				v173 = (v14 + v167 - int32(1)) & (v171 - v167)
				v174 = base.I32_extend16_s(v73)
				if v174 == v171 {
					v183 = v174
					v185 = v173
				} else {
					v178 = v174
					v180 = v173
					base.MemoryCopy(m, v180, v6, v178)
					v183 = v178
					v185 = v180
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v183 + v185
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(993)
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v711 int32
	_ = v711
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1014 int32
	_ = v1014
	var v1033 int32
	_ = v1033
	var v1042 int32
	_ = v1042
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1081 int32
	_ = v1081
	var v1090 int32
	_ = v1090
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	v4 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(832)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v4 < v32 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if int32(0) < v691 {
		goto L115
	} else {
		goto L116
	}
L4:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v62 = v59 + v47*int32(224)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+5)))
	if v63 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v664 = v47 + int32(1)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v664 < v665 {
		v47 = v664
		goto L4
	} else {
		goto L114
	}
L7:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v62)+124))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v62)+192))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v62)+212))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v62)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(0)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v324+v325<<(uint(int32(2))%32))))
	F_tuplesort_performsort(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L12
	} else {
		goto L64
	}
L8:
	;
	v68 = l2 + v47<<(uint(int32(3))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v70 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v62)+124))
	if v292 <= int32(0) {
		goto L6
	} else {
		goto L54
	}
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)+124))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)+212))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79<<(uint(int32(2))%32))))
	F_tuplesort_performsort(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32))))
	v95 = v77 + int32(28)
	v97 = v77 + int32(32)
	v100 = F_tuplesort_getdatum(m, v91, int32(1), int32(0), v95, v97, v27+int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(int32(2))%32))))
	F_tuplesort_end(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L53
	}
L15:
	;
	if v100 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v106 = int32(0)
	v117 = v106
	v119 = v106
	v121 = int32(1)
	v123 = v106
	goto L17
L17:
	;
	goto L24
L18:
	;
	if v245&int32(1) != 0 {
		goto L14
	} else {
		goto L50
	}
L19:
	;
	goto L18
L20:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v229 = int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231+v232<<(uint(int32(2))%32))))
	v241 = F_tuplesort_getdatum(m, v236, v229, int32(0), v95, v97, v27+int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L48
	}
L21:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v220 != 0 {
		v226 = v119
		goto L20
	} else {
		goto L46
	}
L22:
	;
	if v121&int32(1) != 0 {
		goto L21
	} else {
		goto L44
	}
L23:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v226 = v214
	goto L20
L24:
	;
	F_MemoryContextReset(m, v74)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L27
	}
L25:
	;
	F_advance_transition_function(m, l0, v62, v68)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L42
	}
L26:
	;
	goto L25
L27:
	;
	v163 = int32(_a_F_finalize_aggregates_0)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v74
	if v117&base.B2i32(v106 < v73) == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v121&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v164
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(int32(2))%32))))
	v203 = F_tuplesort_getdatum(m, v198, int32(1), int32(0), v95, v97, v27+int32(12))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L40
	}
L30:
	;
	if v169&int32(1) != 0 {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v169&int32(1) != 0 {
		goto L26
	} else {
		goto L36
	}
L33:
	;
	F_advance_transition_function(m, l0, v62, v68)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v164
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v178 == int32(0) {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v123 != v183 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v62)+116))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v187 = F_FunctionCall2Coll(m, v62+int32(144), v185, v119, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	if v187 == int32(0) {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	goto L29
L40:
	;
	if v203 != 0 {
		goto L24
	} else {
		goto L41
	}
L41:
	;
	v244 = v119
	v245 = v121
	goto L19
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v164
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v210 == int32(0) {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	goto L23
L44:
	;
	F_pfree(m, v119)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	goto L21
L46:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+182)))
	v224 = F_datumCopy(m, v221, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v226 = v224
	goto L20
L48:
	;
	if v241 != 0 {
		v117 = v229
		v119 = v226
		v121 = v228
		v123 = v230
		goto L17
	} else {
		goto L49
	}
L49:
	;
	v244 = v226
	v245 = v228
	goto L19
L50:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v250 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, v244)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	goto L14
L53:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v286<<(uint(int32(2))%32)))) = int32(0)
	goto L6
L54:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+205)))
	if v295 != int32(1) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+205)) = uint8(v298)
	if v292 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v302 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v62)+192))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	m.T0[v313].(func(*base.Module, int32))(m, v311)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L12
	} else {
		goto L63
	}
L59:
	;
	v307 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+200)) = v307
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+204)) = uint8(v307)
	goto L6
L60:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+204)))
	if v303 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v62)+200))
	F_pfree(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
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
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	m.T0[v333].(func(*base.Module, int32))(m, v320)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	if v318 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	m.T0[v337].(func(*base.Module, int32))(m, v318)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L12
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341<<(uint(int32(2))%32))))
	v346 = int32(1)
	v350 = F_tuplesort_gettupleslot(m, v345, v346, v346, v320, v27+int32(12))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v350 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v357 = v319 + int32(20)
	v358 = int32(0)
	v363 = v318
	v366 = v320
	v373 = v358
	v376 = v358
	goto L74
L72:
	;
	v598 = v318
	goto L73
L73:
	;
	if v598 != 0 {
		goto L109
	} else {
		goto L110
	}
L74:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[1]))
	if v385 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v598 = v555
	goto L73
L76:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L12
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v366
	v390 = int32(0)
	if base.B2i32(v316 != v390)&v373 == v390 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L78
L80:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	F_MemoryContextReset(m, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L12
	} else {
		goto L105
	}
L81:
	;
	v555 = v363
	v558 = v366
	v565 = int32(1)
	v568 = v376
	goto L80
L82:
	;
	v414 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366)+6)))
	if v414 < v317 {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v395 != v376 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v62)+172))
	if v397 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v400 = int32(_a_F_finalize_aggregates_0)
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v403
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v397)+20))
	v408 = m.T0[v407].(func(*base.Module, int32, int32, int32) int32)(m, v397, v69, v27+int32(11))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v401
	if v408 != 0 {
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L82
L88:
	;
	F_slot_getsomeattrs_int(m, v366, v317)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L12
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v317 <= int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	F_advance_transition_function(m, l0, v62, v68)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L12
	} else {
		goto L101
	}
L93:
	;
	v420 = int32(0)
	if v317 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v427 = v420
	v431 = v420
	goto L97
L95:
	;
	v486 = v420
	goto L96
L96:
	;
	v509 = v357 + v486<<(uint(int32(3))%32)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v510+v486<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v509)+8)) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v366)+20))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516+v486))))
	*(*uint8)(unsafe.Add(mBase, uint32(v509)+12)) = uint8(v518)
	goto L92
L97:
	;
	v449 = v427 | int32(1)
	v450 = int32(3)
	v452 = v357 + v449<<(uint(v450)%32)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	v454 = int32(2)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453+v427<<(uint(v454)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v366)+20))
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v427))))
	*(*uint8)(unsafe.Add(mBase, uint32(v452)+4)) = uint8(v461)
	v464 = v427 + v454
	v467 = v357 + v464<<(uint(v450)%32)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v468+v449<<(uint(v454)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v467))) = v472
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v366)+20))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449+v474))))
	*(*uint8)(unsafe.Add(mBase, uint32(v467)+4)) = uint8(v476)
	v479 = v431 + v454
	if v479 != v317&int32(2147483646) {
		v427 = v464
		v431 = v479
		goto L97
	} else {
		goto L99
	}
L98:
	;
	if v317&int32(1) == int32(0) {
		goto L92
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v486 = v464
	goto L96
L101:
	;
	if v316 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v555 = v363
	v558 = v366
	v565 = v373
	v568 = v376
	goto L80
L103:
	;
	goto L104
L104:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v555 = v366
	v558 = v363
	v565 = int32(1)
	v568 = v549
	goto L80
L105:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v558)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+12))
	m.T0[v580].(func(*base.Module, int32))(m, v558)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v583+v584<<(uint(int32(2))%32))))
	v589 = int32(1)
	v593 = F_tuplesort_gettupleslot(m, v588, v589, v589, v558, v27+int32(12))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	if v593 != 0 {
		v363 = v555
		v366 = v558
		v373 = v565
		v376 = v568
		goto L74
	} else {
		goto L108
	}
L108:
	;
	goto L75
L109:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v598)+8))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+12))
	m.T0[v620].(func(*base.Module, int32))(m, v598)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v623+v624<<(uint(int32(2))%32))))
	F_tuplesort_end(m, v628)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L12
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v631+v632<<(uint(int32(2))%32)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v323
	goto L6
L114:
	;
	goto L5
L115:
	;
	v695 = v27 + int32(32)
	v711 = int32(0)
	goto L118
L116:
	;
	goto L117
L117:
	;
	m.G0 = v27 + int32(832)
	return
L118:
	;
	v721 = v711 + v30
	v722 = int32(2)
	v724 = v31 + v711<<(uint(v722)%32)
	v727 = l1 + v711*int32(52)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	v731 = l2 + v728<<(uint(int32(3))%32)
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v732&v722 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L117
L120:
	;
	v1221 = v711 + int32(1)
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1221 < v1222 {
		v711 = v1221
		goto L118
	} else {
		goto L219
	}
L121:
	;
	v735 = int32(_a_F_finalize_aggregates_0)
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v740
	v744 = v737 + v728*int32(224)
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+20))
	if v745 != 0 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v835 = int32(0)
	v836 = int32(_a_F_finalize_aggregates_0)
	v837 = *(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0]))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v841
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v727)+44))
	if v845 == v835 {
		goto L162
	} else {
		goto L163
	}
L124:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+70)))
	if v747 == int32(1) {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	goto L126
L126:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	if v811 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v781)+20)) = v782
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	v785 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v781)+16)) = uint8(v785)
	*(*uint8)(unsafe.Add(mBase, uint32(v781)+24)) = uint8(v784)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	v790 = m.T0[v789].(func(*base.Module, int32) int32)(m, v781)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L12
	} else {
		goto L143
	}
L128:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v771 != int32(1) {
		v780 = v770
		goto L140
	} else {
		goto L141
	}
L129:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v781 = v768
	v782 = v769
	goto L127
L130:
	;
	v765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v744)+184)))
	if v765 == int32(_a_F_finalize_aggregates_1) {
		goto L128
	} else {
		goto L138
	}
L131:
	;
	if v746&int32(1) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L133
L133:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v744)+216))
	if v746&int32(1) != 0 {
		v768 = v761
		goto L129
	} else {
		goto L137
	}
L134:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v744)+216))
	v764 = v754
	goto L130
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = int32(0)
	v757 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v757)
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v736
	goto L120
L137:
	;
	v764 = v761
	goto L130
L138:
	;
	v768 = v764
	goto L129
L139:
	;
	v781 = v764
	v782 = v780
	goto L127
L140:
	;
	goto L139
L141:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+1)))
	if v774 != int32(3) {
		v780 = v770
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v770)+2))
	v780 = v777 + int32(18)
	goto L140
L143:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v792)
	if v792 != 0 {
		v807 = v790
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v807
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v736
	goto L120
L145:
	;
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v727)+48)))
	if v794 != int32(_a_F_finalize_aggregates_1) {
		v807 = v790
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790))))
	if v797 != int32(1) {
		v806 = v790
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v807 = v806
	goto L144
L148:
	;
	goto L147
L149:
	;
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790)+1)))
	if v800 != int32(3) {
		v806 = v790
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v790)+2))
	v806 = v803 + int32(18)
	goto L148
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v829
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v831)
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v736
	goto L120
L152:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818))))
	if v819 != int32(1) {
		v828 = v818
		goto L158
	} else {
		goto L159
	}
L153:
	;
	v814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v744)+184)))
	if v814 == int32(_a_F_finalize_aggregates_1) {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v829 = v817
	goto L151
L156:
	;
	goto L155
L157:
	;
	v829 = v828
	goto L151
L158:
	;
	goto L157
L159:
	;
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818)+1)))
	if v822 != int32(3) {
		v828 = v818
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v818)+2))
	v828 = v825 + int32(18)
	goto L158
L161:
	;
	v929 = v838 + v728*int32(224)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	if v930 != 0 {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	v911 = int32(1)
	v912 = v835
	goto L161
L163:
	;
	goto L164
L164:
	;
	v849 = int32(1)
	v850 = int32(0)
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	if v851 <= v850 {
		v911 = v849
		v912 = v835
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v859 = v850
	v860 = v849
	v861 = v835
	goto L166
L166:
	;
	v880 = v695 + v860<<(uint(int32(3))%32)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v845)+12))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v881+v859<<(uint(int32(2))%32))))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v885)+20))
	v890 = m.T0[v889].(func(*base.Module, int32, int32, int32) int32)(m, v885, v886, v880+int32(4))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L12
	} else {
		goto L168
	}
L167:
	;
	v911 = v900
	v912 = v898
	goto L161
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v880))) = v890
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+4)))
	v894 = int32(1)
	v898 = base.B2i32(v893|v861&v894 != int32(0))
	v900 = v860 + v894
	v902 = v859 + v894
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	if v902 < v903 {
		v859 = v902
		v860 = v900
		v861 = v898
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_finalize_aggregates[0])) = v837
	goto L120
L171:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v727)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v727
	v933 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v727 + int32(12)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v929)+116))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+30)) = uint16(v931)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+28)) = uint8(v933)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v939
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	if v944 == v933 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	goto L173
L173:
	;
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	if v1148 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+36)) = uint8(v964)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v963
	if v931 <= v911 {
		goto L185
	} else {
		goto L186
	}
L175:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951))))
	if v952 != int32(1) {
		v961 = v951
		goto L181
	} else {
		goto L182
	}
L176:
	;
	v947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+184)))
	if v947 == int32(_a_F_finalize_aggregates_1) {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v963 = v950
	v964 = v944
	goto L174
L179:
	;
	goto L178
L180:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	v963 = v961
	v964 = v962
	goto L174
L181:
	;
	goto L180
L182:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+1)))
	if v955 != int32(3) {
		v961 = v951
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v951)+2))
	v961 = v958 + int32(18)
	goto L181
L184:
	;
	v1107 = int32(1)
	v1109 = int32(0)
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111)+10)))
	if base.B2i32(v1090&v1107 == v1109)|base.B2i32(v1112 != v1107) == v1109 {
		goto L198
	} else {
		goto L199
	}
L185:
	;
	v1090 = v964 | v912
	goto L184
L186:
	;
	goto L187
L187:
	;
	v972 = (v931 - v911) & int32(3)
	if v972 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v976 = int32(0)
	v978 = v911
	goto L191
L189:
	;
	v1014 = v911
	goto L190
L190:
	;
	v1033 = int32(1)
	if base.Ui32(int32(-4)) < base.Ui32(v911-v931) {
		v1090 = v1033
		goto L184
	} else {
		goto L194
	}
L191:
	;
	v999 = v695 + v978<<(uint(int32(3))%32)
	v1000 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+4)) = uint8(v1000)
	*(*int32)(unsafe.Add(mBase, uint32(v999))) = int32(0)
	v1005 = v978 + v1000
	v1007 = v976 + v1000
	if v1007 != v972 {
		v976 = v1007
		v978 = v1005
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v1014 = v1005
	goto L190
L193:
	;
	goto L192
L194:
	;
	v1042 = v1014
	goto L195
L195:
	;
	v1063 = v695 + v1042<<(uint(int32(3))%32)
	v1064 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1063)+4)) = uint8(v1064)
	v1066 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1063))) = v1066
	*(*uint8)(unsafe.Add(mBase, uint32(v1063)+28)) = uint8(v1064)
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+24)) = v1066
	*(*uint8)(unsafe.Add(mBase, uint32(v1063)+20)) = uint8(v1064)
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+16)) = v1066
	*(*uint8)(unsafe.Add(mBase, uint32(v1063)+12)) = uint8(v1064)
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+8)) = v1066
	v1081 = v1042 + int32(4)
	if v1081 != v931 {
		v1042 = v1081
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v1090 = v1033
	goto L184
L197:
	;
	goto L196
L198:
	;
	v1118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v1118
	v1120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v1120)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1118
	goto L170
L199:
	;
	goto L200
L200:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1127 = m.T0[v1126].(func(*base.Module, int32) int32)(m, v27+int32(12))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L12
	} else {
		goto L201
	}
L201:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v1129)
	if v1129 != 0 {
		v1144 = v1127
		goto L202
	} else {
		goto L203
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L170
L203:
	;
	v1131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v727)+48)))
	if v1131 != int32(_a_F_finalize_aggregates_1) {
		v1144 = v1127
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127))))
	if v1134 != int32(1) {
		v1143 = v1127
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v1144 = v1143
	goto L202
L206:
	;
	goto L205
L207:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127)+1)))
	if v1137 != int32(3) {
		v1143 = v1127
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+2))
	v1143 = v1140 + int32(18)
	goto L206
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v1166
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v1168)
	goto L170
L210:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155))))
	if v1156 != int32(1) {
		v1165 = v1155
		goto L216
	} else {
		goto L217
	}
L211:
	;
	v1151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+184)))
	if v1151 == int32(_a_F_finalize_aggregates_1) {
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v1166 = v1154
	goto L209
L214:
	;
	goto L213
L215:
	;
	v1166 = v1165
	goto L209
L216:
	;
	goto L215
L217:
	;
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155)+1)))
	if v1159 != int32(3) {
		v1165 = v1155
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+2))
	v1165 = v1162 + int32(18)
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
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
	F_ScanKeyInit(m, v10, int32(9), int32(3), int32(184), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(1)
	v27 = F_systable_beginscan(m, v14, int32(2665), v24, int32(0), v24, v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_systable_endscan(m, v27)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L5:
	;
	v29 = F_systable_getnext(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		v59 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v33 = v29
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v41)+72)))
	if v43 != int32(110) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v59 = v3
	goto L4
L10:
	;
	v51 = F_systable_getnext(m, v27)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v46 = F_extractNotNullColumn(m, v33)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v46 != l1 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v49 = F_heap_copytuple(m, v33)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v59 = v49
	goto L4
L15:
	;
	if v51 != 0 {
		v33 = v51
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
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(48)
	return v59
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v26) <= base.Ui32(v21) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	F_errfinish(m, int32(_a_F_find_childrel_parents_1), int32(426), int32(_a_F_find_childrel_parents_2))
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
	F_errdetail(m, int32(_a_F_finish_nodeitem_5), int32(0))
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
	v212 = int32(612)
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
	F_errdetail(m, int32(_a_F_finish_nodeitem_7), v12+int32(16))
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
	v212 = int32(618)
	goto L33
L52:
	;
	v217 = v211
	goto L32
}
func F_fix_indexqual_operand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v4 = int32(0)
	if l0 == v4 {
		v41 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 == int32(27) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v52 = v41
	goto L1
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v22 = l0
	goto L12
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v18 = F_expression_tree_walker_impl(m, l0, int32(825), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	return int32(0)
L10:
	;
	if v18 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v52 = l0
	goto L1
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v29 != int32(319) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v36 = F_expression_tree_mutator_impl(m, v22, int32(826), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L18
	}
L14:
	;
	goto L13
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v32 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v33 != 0 {
		v22 = v33
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v41 = v4
	goto L2
L18:
	;
	v41 = v36
	goto L2
L19:
	;
	v60 = v52
	goto L22
L20:
	;
	v67 = v53
	v71 = v52
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+l2<<(uint(int32(2))%32))))
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 == int32(27) {
		v60 = v63
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v67 = v64
	v71 = v63
	goto L21
L24:
	;
	goto L23
L25:
	;
	if v67 != int32(6) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v108 != 0 {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L33
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+68))
	if v81 != v83 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+8)))
	if v78 != v85 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v87 = F_copyObjectImpl(m, v71)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v90 = l2 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v87)+8)) = uint16(v90)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(-3)
	return v87
L33:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_2), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v111 = v109
	goto L38
L37:
	;
	v111 = int32(0)
	goto L38
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if int32(0) < v112 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L74
	}
L40:
	;
	v116 = int32(0)
	v119 = v111
	goto L43
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L71
	}
L43:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v74+v116<<(uint(int32(2))%32))))
	if v126 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	if v119 == int32(0) {
		goto L39
	} else {
		goto L48
	}
L46:
	;
	v181 = v119
	goto L47
L47:
	;
	v183 = v116 + int32(1)
	if v183 != v112 {
		v116 = v183
		v119 = v181
		goto L43
	} else {
		goto L70
	}
L48:
	;
	if v116 == l2 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	if v132 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v172 = v119 + int32(4)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if base.Ui32(v172) < base.Ui32(v174+v175<<(uint(int32(2))%32)) {
		goto L67
	} else {
		goto L68
	}
L52:
	;
	v141 = F_equal(m, v71, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L57
	}
L53:
	;
	v140 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v136 != int32(27) {
		v140 = v132
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v140 = v139
	goto L52
L57:
	;
	if v141 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v148 = F_exprType(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L64
	}
L61:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v152 = F_exprCollation(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v155 = F_makeVar(m, int32(-3), base.I32_extend16_s(l2+int32(1)), v148, int32(-1), v152, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	return v155
L64:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_0), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_4), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v180 = v172
	goto L69
L68:
	;
	v180 = int32(0)
	goto L69
L69:
	;
	v181 = v180
	goto L47
L70:
	;
	goto L44
L71:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_0), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_5), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F_fix_indexqual_operand_6), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_fix_indexqual_operand_1), int32(_a_F_fix_indexqual_operand_7), int32(_a_F_fix_indexqual_operand_3))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
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
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
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
func F_float48gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = base.F64_promote_f32(v11)
		v22 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_float4larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v10 float32
	_ = v10
	var v18 float32
	_ = v18
	var v20 float32
	_ = v20
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v10)&int32(2147483647)))|base.F32_lt(v4, v10) != 0 {
			v18 = v10
		} else {
			v18 = v4
		}
		v20 = v18
	} else {
		v20 = v4
	}
	return base.I32_reinterpret_f32(v20)
}
func F_float84le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_ge(v5, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_float84ne(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = base.F64_promote_f32(v5)
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		return base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313)))
	} else {
		return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9)) | base.F64_ne(v6, v11)
	}
}
func F_float8up(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = F_Float8GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v119
L2:
	;
	return int32(0)
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
	v20 = v18 + v19
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+104)))
	F_initStringInfo(m, v11+int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
		v119 = v3
		goto L1
	} else {
		goto L40
	}
L7:
	;
	v29 = l1 & int32(2)
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v39 = F_quote_qualified_identifier(m, v38, v20+int32(4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L15
	}
L9:
	;
	v33 = F_FunctionIsVisibleExt(m, l0, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v36 = F_get_namespace_name(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	if v33 != 0 {
		v38 = v3
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v38 = v36
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v39
	v43 = v11 + int32(32)
	F_appendStringInfo(m, v43, int32(_a_F_format_procedure_extended_0), v11+int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v23 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_appendStringInfoChar(m, v11+int32(32), int32(41))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L38
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v20)+136))
	if v29 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_appendStringInfoString(m, v43, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L25
	}
L20:
	;
	v52 = F_format_type_be_qualified(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v54 = F_format_type_be(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v56 = v52
	goto L19
L24:
	;
	v56 = v54
	goto L19
L25:
	;
	v59 = int32(1)
	if v23 == v59 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v64 = v59
	goto L27
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(136)+v64<<(uint(int32(2))%32))))
	v77 = v11 + int32(32)
	F_appendStringInfoChar(m, v77, int32(44))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L29
	}
L28:
	;
	goto L17
L29:
	;
	if v29 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_appendStringInfoString(m, v77, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L36
	}
L31:
	;
	v81 = F_format_type_be_qualified(m, v75)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v83 = F_format_type_be(m, v75)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L35
	}
L34:
	;
	v85 = v81
	goto L30
L35:
	;
	v85 = v83
	goto L30
L36:
	;
	v89 = v64 + int32(1)
	if v89 != v23 {
		v64 = v89
		goto L27
	} else {
		goto L37
	}
L37:
	;
	goto L28
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v119 = v104
	goto L1
L40:
	;
	v110 = F_palloc(m, int32(64))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v115 = F_pg_snprintf(m, v110, int32(64), int32(_a_F_format_procedure_extended_1), v11)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v119 = v110
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
func F_ftod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_Float8GetDatum(m, base.F64_promote_f32(v2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
