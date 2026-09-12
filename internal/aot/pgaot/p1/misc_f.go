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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[800]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[801]))
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
	*(*int32)(unsafe.Add(mBase, _consts[801])) = v14 + int32(1)
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
	*(*int32)(unsafe.Add(mBase, _consts[802])) = l1
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
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
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
	v265 = m.ExcPending
	if v265 != 0 {
		goto L59
	} else {
		goto L74
	}
L2:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v232+int32(3584))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L59
	} else {
		goto L67
	}
L3:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v202+int32(3584))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L59
	} else {
		goto L60
	}
L4:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[816]))
	if v146 == l0 {
		goto L45
	} else {
		goto L46
	}
L5:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[200]))
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
		v140 = v14
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
		v140 = v14
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
		v140 = v14
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v90 = l0 + int32(40)
	if v86 == v90 {
		v140 = v14
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
	v140 = v14
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
	v140 = v134
	goto L4
L44:
	;
	m.G0 = v12 + int32(16)
	return
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v140 | int32(512)
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
	v157 = *(*int32)(unsafe.Add(mBase, _consts[817]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v158 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v158 == v157 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = l0
	v170 = l0 + int32(32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v171 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v170
	goto L55
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+4)) = v170
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v158
	v182 = l1 + int32(40)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v183 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v182
	goto L58
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = v182
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v189
	v192 = v158 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v192
	goto L44
L59:
	;
	return
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(142199), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	F_errdetail_internal(m, int32(630460), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	F_errhint(m, int32(650717), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(498955), int32(4668), int32(363852))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L59
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L59
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L59
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(142199), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L59
	} else {
		goto L70
	}
L70:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v248
	F_errdetail_internal(m, int32(652795), v12)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L59
	} else {
		goto L71
	}
L71:
	;
	F_errhint(m, int32(650717), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L59
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(498955), int32(4680), int32(363852))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L59
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
	F_errcode(m, int32(8389))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L59
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(109653), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L59
	} else {
		goto L76
	}
L76:
	;
	F_errhint(m, int32(664015), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L59
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(498955), int32(654), int32(109877))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L59
	} else {
		goto L78
	}
L78:
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
	switch v61 - int32(65530) {
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
	v102 = int32(4515600)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v105
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v103
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
	F_errmsg_internal(m, int32(145162), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(492721), int32(2772), int32(286650))
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
	F_errmsg_internal(m, int32(145162), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(492721), int32(2783), int32(286650))
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	v10 = v8 - int32(1)
	if int32(0) <= v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[753]))
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
	F_errmsg_internal(m, int32(213525), int32(0))
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
	F_errfinish(m, int32(500025), int32(3050), int32(213581))
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v91
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
		v91 = int32(0)
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if v43 < int32(0) {
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
	if v76 != 0 {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v76 = l1
	v78 = int32(0)
	v79 = v34
	v80 = l0
	goto L15
L17:
	;
	if l1 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v51 = l1
	goto L19
L19:
	;
	v54 = l0 + v51
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54-int32(1)))))
	if v57 != int32(10) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v63 = m.T0[v62].(func(*base.Module, int32, int32, int32) int32)(m, l2, l0, v51)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L25
	}
L21:
	;
	v61 = v51 - int32(1)
	if v61 != 0 {
		v51 = v61
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L16
L25:
	;
	if base.Ui32(v63) < base.Ui32(v51) {
		v91 = v63
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v76 = l1 - v51
	v78 = v51
	v79 = v67
	v80 = v54
	goto L15
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v83 + v76
	v91 = v76 + v78
	goto L1
L28:
	;
	v81 = F__emscripten_memcpy_bulkmem(m, v79, v80, v76)
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
}
func F_fastgetattr_3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = int32(1)
	v15 = l1 - v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v17&v14 == v5 {
		v26 = l2 + v15<<(uint(int32(4))%32) + int32(20)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		if v27 < int32(0) {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v10 + int32(16)
				return v79
			}
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v32 = v16 + v30 + v27
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)))
			if v33 != int32(1) {
				v79 = v32
				m.G0 = v10 + int32(16)
				return v79
			} else {
				v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
				switch v36&int32(65535) - int32(1) {
				case 0:
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
					v79 = v41
					m.G0 = v10 + int32(16)
					return v79
				case 1:
					v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32))))
					v79 = v42
					m.G0 = v10 + int32(16)
					return v79
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v36
						F_errmsg_internal(m, int32(483562), v10)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(326784), int32(70), int32(67821))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					v79 = v43
					m.G0 = v10 + int32(16)
					return v79
				}
			}
		}
	} else {
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v15)>>(uint(int32(3))%32)))+23)))
		if int32(base.Ui32(v62)>>(uint(v15&int32(7))%32))&int32(1) != 0 {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v10 + int32(16)
				return v79
			}
		} else {
			v68 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v68)
			v79 = int32(0)
			m.G0 = v10 + int32(16)
			return v79
		}
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
	v13 = F_psprintf(m, int32(176001), v5+int32(-16))
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
							F_errmsg_internal(m, int32(703911), v5+int32(-48))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495426), int32(481), int32(420736))
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
									F_errmsg(m, int32(703852), v5+int32(-32))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495426), int32(491), int32(420736))
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
						F_errmsg(m, int32(703798), v7)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(662041), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495426), int32(472), int32(420736))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v12 = l0 + l1<<(uint(int32(2))%32) + int32(192)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v126
L2:
	;
	v95 = F_palloc0(m, int32(272))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v16 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 <= v16 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = v16
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v21<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = int32(0)
	v40 = base.B2i32(v32|l2 == v33)
	if v32 == v33 {
		v79 = v40
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	if v79 != 0 {
		v126 = v31
		goto L1
	} else {
		goto L19
	}
L8:
	;
	goto L7
L9:
	;
	if l2 == int32(0) {
		v79 = v40
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 != v47 {
		v79 = int32(0)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v49 = int32(1)
	if v46 <= v49 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = v49
	goto L14
L13:
	;
	v52 = v46
	goto L14
L14:
	;
	v53 = int32(8)
	v58 = int32(0)
	goto L15
L15:
	;
	v66 = v58 << (uint(int32(2)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v32+v53+v66)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+(l2+v53))))
	v71 = base.B2i32(v68 == v70)
	if v70 != v68 {
		v79 = v71
		goto L8
	} else {
		goto L17
	}
L16:
	;
	v79 = v71
	goto L8
L17:
	;
	v74 = v58 + int32(1)
	if v74 != v52 {
		v58 = v74
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v84 = v21 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v84 < v85 {
		v21 = v84
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L6
L21:
	;
	return int32(0)
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v95))) = int64(17179869452)
	v101 = F_bms_copy(m, l2)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v101
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l0)+296))
	v105 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v95)+25)) = uint16(v105)
	v108 = base.F64_gt(v104, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v108)
	v110 = F_create_empty_pathtarget(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v95)+44)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v95)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v95)+52)) = v112
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v120 = F_lappend(m, v119, v95)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v120
	v126 = v95
	goto L1
}
func F_fgets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	if int32(0) <= v9 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v235
L2:
	;
	v46 = l0
	v48 = l1 - int32(1)
	goto L13
L3:
	;
	if l1 != int32(1) {
		v235 = int32(0)
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v25 - int32(1) | v25
	goto L10
L5:
	;
	if l1 < int32(2) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v17 = int32(1)
	if v17 < l1 {
		v42 = v17
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(0)
	goto L2
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v20 - int32(1) | v20
	goto L3
L10:
	;
	goto L3
L11:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v38)
	return l0
L12:
	;
	if v42 != 0 {
		v235 = v228
		goto L1
	} else {
		goto L69
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v53 == v54 {
		v187 = v46
		v188 = v48
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if l0 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L15:
	;
	goto L14
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v205)
	v210 = v203 + int32(1)
	if v205&int32(255) == int32(10) {
		v217 = v210
		goto L15
	} else {
		goto L64
	}
L17:
	;
	v192 = F___uflow(m, l2)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L59
	} else {
		goto L60
	}
L18:
	;
	v57 = v54 - v53
	v58 = int32(0)
	v61 = base.B2i32(v57 != v58)
	if v53&int32(3) == v58 {
		v87 = v53
		v89 = v57
		v90 = v61
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if base.Ui32(v169) < base.Ui32(v48) {
		goto L49
	} else {
		goto L50
	}
L20:
	;
	if v160 != 0 {
		goto L46
	} else {
		goto L47
	}
L21:
	;
	v160 = int32(0)
	goto L20
L22:
	;
	v138 = v131
	v140 = v133
	goto L40
L23:
	;
	if v90 == int32(0) {
		goto L21
	} else {
		goto L31
	}
L24:
	;
	if v57 == int32(0) {
		v87 = v53
		v89 = v57
		v90 = v61
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v70 = v53
	v72 = v57
	goto L26
L26:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 == int32(10) {
		v131 = v70
		v133 = v72
		goto L22
	} else {
		goto L28
	}
L27:
	;
	v87 = v82
	v89 = v78
	v90 = v80
	goto L23
L28:
	;
	v77 = int32(1)
	v78 = v72 - v77
	v79 = int32(0)
	v80 = base.B2i32(v78 != v79)
	v82 = v70 + v77
	if v82&int32(3) == v79 {
		v87 = v82
		v89 = v78
		v90 = v80
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if v78 != 0 {
		v70 = v82
		v72 = v78
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v94 == int32(10) {
		v124 = v87
		v126 = v89
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v126 == int32(0) {
		goto L21
	} else {
		goto L39
	}
L33:
	;
	if base.Ui32(v89) < base.Ui32(int32(4)) {
		v124 = v87
		v126 = v89
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v104 = v87
	v106 = v89
	goto L35
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v111 = v110 ^ int32(168430090)
	v114 = int32(-2139062144)
	if (int32(16843008)-v111|v111)&v114 != v114 {
		v131 = v104
		v133 = v106
		goto L22
	} else {
		goto L37
	}
L36:
	;
	v124 = v119
	v126 = v121
	goto L32
L37:
	;
	v118 = int32(4)
	v119 = v104 + v118
	v121 = v106 - v118
	if base.Ui32(int32(3)) < base.Ui32(v121) {
		v104 = v119
		v106 = v121
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v131 = v124
	v133 = v126
	goto L22
L40:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if int32(10) == v143 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L21
L42:
	;
	v160 = v138
	goto L20
L43:
	;
	goto L44
L44:
	;
	v145 = int32(1)
	v148 = v140 - v145
	if v148 != 0 {
		v138 = v138 + v145
		v140 = v148
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v168 = v161
	v169 = v160 - v161 + int32(1)
	goto L19
L47:
	;
	goto L48
L48:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v168 = v166
	v169 = v165 - v166
	goto L19
L49:
	;
	v171 = v169
	goto L51
L50:
	;
	v171 = v48
	goto L51
L51:
	;
	if v171 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v175 = v174 + v171
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v175
	v177 = v46 + v171
	if v160 != 0 {
		v217 = v177
		goto L15
	} else {
		goto L56
	}
L53:
	;
	v172 = F__emscripten_memcpy_bulkmem(m, v46, v168, v171)
	mBase = m.M
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v178 = v48 - v171
	if v178 == int32(0) {
		v217 = v177
		goto L15
	} else {
		goto L57
	}
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v175 == v181 {
		v187 = v177
		v188 = v178
		goto L17
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v175 + int32(1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v203 = v177
	v204 = v178
	v205 = v186
	goto L16
L59:
	;
	return int32(0)
L60:
	;
	if int32(0) <= v192 {
		v203 = v187
		v204 = v188
		v205 = v192
		goto L16
	} else {
		goto L61
	}
L61:
	;
	v198 = int32(0)
	if l0 == v187 {
		v228 = v198
		goto L12
	} else {
		goto L62
	}
L62:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v200&int32(16) != 0 {
		v217 = v187
		goto L15
	} else {
		goto L63
	}
L63:
	;
	v228 = v198
	goto L12
L64:
	;
	v216 = v204 - int32(1)
	if v216 != 0 {
		v46 = v210
		v48 = v216
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v217 = v210
	goto L15
L66:
	;
	v228 = int32(0)
	goto L12
L67:
	;
	goto L68
L68:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v225)
	v228 = l0
	goto L12
L69:
	;
	v235 = v228
	goto L1
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
	var v22 int32
	_ = v22
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
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
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
			v22 = v20 + v21
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v24)
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
					v184 = v72
					v186 = v49
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
					m.G0 = v12 + int32(16)
					return
				case 1:
					*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v6)
					v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
					v184 = v54
					v186 = v49
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
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
						F_errmsg_internal(m, int32(483562), v12)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(326784), int32(230), int32(308869))
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
					v184 = v56
					v186 = v49
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
					m.G0 = v12 + int32(16)
					return
				}
			} else {
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				switch v73 - int32(65534) {
				case 0:
					v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
					v166 = v164 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v166)
					v168 = F_strlen(m, v6)
					mBase = m.M
					v170 = v168 + int32(1)
					if v170 != 0 {
						v171 = F__emscripten_memcpy_bulkmem(m, v14, v6, v170)
						mBase = m.M
					} else {
					}
					v184 = v170
					v186 = v14
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
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
									v184 = v96
									v186 = v95
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
									m.G0 = v12 + int32(16)
									return
								}
							}
						} else {
							v100 = int32(6)
							v101 = v76 | v100
							*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v101)
							v105 = int32(18)
							v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
							if v107 == v105 {
								v110 = v105
							} else {
								v110 = int32(2)
							}
							if v107&int32(254) == int32(2) {
								v115 = v100
							} else {
								v115 = v110
							}
							if v107 == int32(1) {
								v118 = v100
							} else {
								v118 = v115
							}
							if v118 != 0 {
								v119 = F__emscripten_memcpy_bulkmem(m, v14, v6, v118)
								mBase = m.M
							} else {
							}
							v184 = v118
							v186 = v14
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						if v80&int32(1) != 0 {
							v124 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if v124 != 0 {
								v125 = F__emscripten_memcpy_bulkmem(m, v14, v6, v124)
								mBase = m.M
							} else {
							}
							v184 = v124
							v186 = v14
						} else {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							v128 = int32(2)
							v129 = int32(base.Ui32(v127) >> (uint(v128) % 32))
							if v80&v128 != 0 {
								v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
								v161 = (v14 + v155 - int32(1)) & (int32(0) - v155)
								if v129 != 0 {
									v162 = F__emscripten_memcpy_bulkmem(m, v161, v6, v129)
									mBase = m.M
								} else {
								}
								v184 = v129
								v186 = v161
							} else {
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
								if v132&int32(1) == int32(0) {
									v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
									v161 = (v14 + v155 - int32(1)) & (int32(0) - v155)
									if v129 != 0 {
										v162 = F__emscripten_memcpy_bulkmem(m, v161, v6, v129)
										mBase = m.M
									} else {
									}
									v184 = v129
									v186 = v161
								} else {
									v138 = v129 - int32(3)
									if base.Ui32(int32(127)) < base.Ui32(v138) {
										v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
										v161 = (v14 + v155 - int32(1)) & (int32(0) - v155)
										if v129 != 0 {
											v162 = F__emscripten_memcpy_bulkmem(m, v161, v6, v129)
											mBase = m.M
										} else {
										}
										v184 = v129
										v186 = v161
									} else {
										v141 = int32(1)
										v144 = v138<<(uint(v141)%32) | v141
										*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v144)
										v148 = int32(4)
										v151 = v129 - v148
										if v151 != 0 {
											v152 = F__emscripten_memcpy_bulkmem(m, v14+v141, v6+v148, v151)
											mBase = m.M
										} else {
										}
										v184 = v138
										v186 = v14
									}
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
						m.G0 = v12 + int32(16)
						return
					}
				default:
					v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					v179 = (v14 + v173 - int32(1)) & (int32(0) - v173)
					v180 = base.I32_extend16_s(v73)
					if v180 != 0 {
						v181 = F__emscripten_memcpy_bulkmem(m, v179, v6, v180)
						mBase = m.M
					} else {
					}
					v184 = v180
					v186 = v179
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
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
				v184 = v72
				v186 = v49
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
				m.G0 = v12 + int32(16)
				return
			case 1:
				*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v6)
				v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				v184 = v54
				v186 = v49
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
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
					F_errmsg_internal(m, int32(483562), v12)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(326784), int32(230), int32(308869))
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
				v184 = v56
				v186 = v49
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			switch v73 - int32(65534) {
			case 0:
				v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
				v166 = v164 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v166)
				v168 = F_strlen(m, v6)
				mBase = m.M
				v170 = v168 + int32(1)
				if v170 != 0 {
					v171 = F__emscripten_memcpy_bulkmem(m, v14, v6, v170)
					mBase = m.M
				} else {
				}
				v184 = v170
				v186 = v14
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
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
								v184 = v96
								v186 = v95
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
								m.G0 = v12 + int32(16)
								return
							}
						}
					} else {
						v100 = int32(6)
						v101 = v76 | v100
						*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v101)
						v105 = int32(18)
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
						if v107 == v105 {
							v110 = v105
						} else {
							v110 = int32(2)
						}
						if v107&int32(254) == int32(2) {
							v115 = v100
						} else {
							v115 = v110
						}
						if v107 == int32(1) {
							v118 = v100
						} else {
							v118 = v115
						}
						if v118 != 0 {
							v119 = F__emscripten_memcpy_bulkmem(m, v14, v6, v118)
							mBase = m.M
						} else {
						}
						v184 = v118
						v186 = v14
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					if v80&int32(1) != 0 {
						v124 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
						if v124 != 0 {
							v125 = F__emscripten_memcpy_bulkmem(m, v14, v6, v124)
							mBase = m.M
						} else {
						}
						v184 = v124
						v186 = v14
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v128 = int32(2)
						v129 = int32(base.Ui32(v127) >> (uint(v128) % 32))
						if v80&v128 != 0 {
							v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
							v161 = (v14 + v155 - int32(1)) & (int32(0) - v155)
							if v129 != 0 {
								v162 = F__emscripten_memcpy_bulkmem(m, v161, v6, v129)
								mBase = m.M
							} else {
							}
							v184 = v129
							v186 = v161
						} else {
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
							if v132&int32(1) == int32(0) {
								v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
								v161 = (v14 + v155 - int32(1)) & (int32(0) - v155)
								if v129 != 0 {
									v162 = F__emscripten_memcpy_bulkmem(m, v161, v6, v129)
									mBase = m.M
								} else {
								}
								v184 = v129
								v186 = v161
							} else {
								v138 = v129 - int32(3)
								if base.Ui32(int32(127)) < base.Ui32(v138) {
									v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
									v161 = (v14 + v155 - int32(1)) & (int32(0) - v155)
									if v129 != 0 {
										v162 = F__emscripten_memcpy_bulkmem(m, v161, v6, v129)
										mBase = m.M
									} else {
									}
									v184 = v129
									v186 = v161
								} else {
									v141 = int32(1)
									v144 = v138<<(uint(v141)%32) | v141
									*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v144)
									v148 = int32(4)
									v151 = v129 - v148
									if v151 != 0 {
										v152 = F__emscripten_memcpy_bulkmem(m, v14+v141, v6+v148, v151)
										mBase = m.M
									} else {
									}
									v184 = v138
									v186 = v14
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
					m.G0 = v12 + int32(16)
					return
				}
			default:
				v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				v179 = (v14 + v173 - int32(1)) & (int32(0) - v173)
				v180 = base.I32_extend16_s(v73)
				if v180 != 0 {
					v181 = F__emscripten_memcpy_bulkmem(m, v179, v6, v180)
					mBase = m.M
				} else {
				}
				v184 = v180
				v186 = v179
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v184 + v186
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(276672)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(993)
	v16 = int32(4508504)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v7 + int32(4)
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
		*(*int32)(unsafe.Add(mBase, _consts[337])) = v36
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
	var v52 int32
	_ = v52
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
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
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v869 int32
	_ = v869
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1183 int32
	_ = v1183
	var v1201 int32
	_ = v1201
	var v1211 int32
	_ = v1211
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1249 int32
	_ = v1249
	var v1260 int32
	_ = v1260
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
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
	v52 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if int32(0) < v858 {
		goto L128
	} else {
		goto L129
	}
L4:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v62 = v59 + v52*int32(224)
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
	v831 = v52 + int32(1)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v831 < v832 {
		v52 = v831
		goto L4
	} else {
		goto L127
	}
L7:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v62)+124))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v62)+192))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v62)+212))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v62)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = int32(0)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v490+v491<<(uint(int32(2))%32))))
	F_tuplesort_performsort(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L12
	} else {
		goto L77
	}
L8:
	;
	v68 = l2 + v52<<(uint(int32(3))%32)
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
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v62)+124))
	if v458 <= int32(0) {
		goto L6
	} else {
		goto L67
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
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v443+v444<<(uint(int32(2))%32))))
	F_tuplesort_end(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L12
	} else {
		goto L66
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
	v105 = v62 + int32(144)
	v106 = int32(1)
	v107 = int32(0)
	v119 = v107
	v122 = v106
	v123 = v106
	v126 = v107
	goto L17
L17:
	;
	F_MemoryContextReset(m, v74)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	if v401&int32(1) != 0 {
		goto L14
	} else {
		goto L63
	}
L19:
	;
	goto L18
L20:
	;
	v138 = int32(4515600)
	v139 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v74
	if v73 <= v107 {
		v269 = v139
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v376 = int32(0)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v378+v379<<(uint(int32(2))%32))))
	v388 = F_tuplesort_getdatum(m, v383, int32(1), v376, v95, v97, v27+int32(12))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L12
	} else {
		goto L61
	}
L22:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v358 = v350
	goto L21
L23:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v320 != 0 {
		v358 = v119
		goto L21
	} else {
		goto L59
	}
L24:
	;
	F_advance_transition_function(m, l0, v62, v68)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L55
	}
L25:
	;
	if v122&int32(1) != 0 {
		v269 = v139
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v146 = v123 & int32(1)
	if v146 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	F_advance_transition_function(m, l0, v62, v68)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L53
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v139
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163+v164<<(uint(int32(2))%32))))
	v173 = F_tuplesort_getdatum(m, v168, int32(1), int32(0), v95, v97, v27+int32(12))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L12
	} else {
		goto L37
	}
L29:
	;
	if v144&int32(1) != 0 {
		v269 = v139
		goto L24
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v144&int32(1) == int32(0) {
		v240 = v139
		goto L27
	} else {
		goto L36
	}
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v126 != v151 {
		v269 = v139
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v62)+116))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v155 = F_FunctionCall2Coll(m, v105, v153, v119, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v155 != 0 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v269 = v139
	goto L24
L36:
	;
	goto L28
L37:
	;
	if v173 == int32(0) {
		v397 = v119
		v401 = v123
		goto L19
	} else {
		goto L38
	}
L38:
	;
	goto L39
L39:
	;
	F_MemoryContextReset(m, v74)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L41
	}
L40:
	;
	v397 = v119
	v401 = v123
	goto L19
L41:
	;
	v203 = int32(4515600)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v74
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v146 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v204
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222+v223<<(uint(int32(2))%32))))
	v232 = F_tuplesort_getdatum(m, v227, int32(1), int32(0), v95, v97, v27+int32(12))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L51
	}
L43:
	;
	if v207&int32(1) != 0 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v207&int32(1) != 0 {
		v269 = v204
		goto L24
	} else {
		goto L47
	}
L46:
	;
	v240 = v204
	goto L27
L47:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v126 != v212 {
		v269 = v204
		goto L24
	} else {
		goto L48
	}
L48:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v62)+116))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v216 = F_FunctionCall2Coll(m, v105, v214, v119, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	if v216 == int32(0) {
		v269 = v204
		goto L24
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	if v232 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v240
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v262 != 0 {
		goto L22
	} else {
		goto L54
	}
L54:
	;
	goto L23
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v269
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v291 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	if v123&int32(1) != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v119)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	goto L23
L59:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	v323 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+182)))
	v324 = F_datumCopy(m, v321, v322, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	v358 = v324
	goto L21
L61:
	;
	if v388 != 0 {
		v119 = v358
		v122 = v376
		v123 = v375
		v126 = v377
		goto L17
	} else {
		goto L62
	}
L62:
	;
	v397 = v358
	v401 = v375
	goto L19
L63:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v416 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	F_pfree(m, v397)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	goto L14
L66:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v451+v452<<(uint(int32(2))%32)))) = int32(0)
	goto L6
L67:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+205)))
	if v461 != int32(1) {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+205)) = uint8(v464)
	if v458 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+186)))
	if v468 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v62)+192))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	m.T0[v479].(func(*base.Module, int32))(m, v477)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L76
	}
L72:
	;
	v473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+200)) = v473
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+204)) = uint8(v473)
	goto L6
L73:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+204)))
	if v469 != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v62)+200))
	F_pfree(m, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L6
L77:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v486)+8))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	m.T0[v499].(func(*base.Module, int32))(m, v486)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	if v484 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+12))
	m.T0[v503].(func(*base.Module, int32))(m, v484)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L12
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v506+v507<<(uint(int32(2))%32))))
	v512 = int32(1)
	v516 = F_tuplesort_gettupleslot(m, v511, v512, v512, v486, v27+int32(12))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	if v516 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v523 = v485 + int32(20)
	v524 = int32(0)
	v529 = v484
	v530 = v486
	v539 = v524
	v545 = v524
	goto L87
L85:
	;
	v765 = v484
	goto L86
L86:
	;
	if v765 != 0 {
		goto L122
	} else {
		goto L123
	}
L87:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v551 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v765 = v722
	goto L86
L89:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L12
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v530
	v556 = int32(0)
	if base.B2i32(v482 != v556)&v539 == v556 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L91
L93:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	F_MemoryContextReset(m, v743)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L12
	} else {
		goto L118
	}
L94:
	;
	v722 = v529
	v723 = v530
	v732 = int32(1)
	v738 = v545
	goto L93
L95:
	;
	v580 = int32(*(*int16)(unsafe.Add(mBase, uint32(v530)+6)))
	if v580 < v483 {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v561 != v545 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v62)+172))
	if v563 == int32(0) {
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v566 = int32(4515600)
	v567 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v569
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v563)+20))
	v574 = m.T0[v573].(func(*base.Module, int32, int32, int32) int32)(m, v563, v69, v27+int32(11))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v567
	if v574 != 0 {
		goto L94
	} else {
		goto L100
	}
L100:
	;
	goto L95
L101:
	;
	F_slot_getsomeattrs_int(m, v530, v483)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L12
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v483 <= int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L103
L105:
	;
	F_advance_transition_function(m, l0, v62, v68)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L12
	} else {
		goto L114
	}
L106:
	;
	v586 = int32(0)
	if v483 != int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v593 = v586
	v599 = v586
	goto L110
L108:
	;
	v650 = v586
	goto L109
L109:
	;
	if v483&int32(1) == int32(0) {
		goto L105
	} else {
		goto L113
	}
L110:
	;
	v615 = v593 | int32(1)
	v616 = int32(3)
	v618 = v523 + v615<<(uint(v616)%32)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v530)+16))
	v620 = int32(2)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619+v593<<(uint(v620)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v618))) = v623
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v530)+20))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v593))))
	*(*uint8)(unsafe.Add(mBase, uint32(v618)+4)) = uint8(v627)
	v630 = v593 + v620
	v633 = v523 + v630<<(uint(v616)%32)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v530)+16))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v634+v615<<(uint(v620)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v633))) = v638
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v530)+20))
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615+v640))))
	*(*uint8)(unsafe.Add(mBase, uint32(v633)+4)) = uint8(v642)
	v645 = v599 + v620
	if v645 != v483&int32(2147483646) {
		v593 = v630
		v599 = v645
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v650 = v630
	goto L109
L112:
	;
	goto L111
L113:
	;
	v675 = v650<<(uint(int32(3))%32) + v523
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v530)+16))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676+v650<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v675)+8)) = v680
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v530)+20))
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682+v650))))
	*(*uint8)(unsafe.Add(mBase, uint32(v675)+12)) = uint8(v684)
	goto L105
L114:
	;
	if v482 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v722 = v529
	v723 = v530
	v732 = v539
	v738 = v545
	goto L93
L116:
	;
	goto L117
L117:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v722 = v530
	v723 = v529
	v732 = int32(1)
	v738 = v715
	goto L93
L118:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+12))
	m.T0[v747].(func(*base.Module, int32))(m, v723)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v750+v751<<(uint(int32(2))%32))))
	v756 = int32(1)
	v760 = F_tuplesort_gettupleslot(m, v755, v756, v756, v723, v27+int32(12))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	if v760 != 0 {
		v529 = v722
		v530 = v723
		v539 = v732
		v545 = v738
		goto L87
	} else {
		goto L121
	}
L121:
	;
	goto L88
L122:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)+12))
	m.T0[v787].(func(*base.Module, int32))(m, v765)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L12
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v790+v791<<(uint(int32(2))%32))))
	F_tuplesort_end(m, v795)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L12
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v62)+208))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v798+v799<<(uint(int32(2))%32)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v489
	goto L6
L127:
	;
	goto L5
L128:
	;
	v862 = v27 + int32(32)
	v869 = int32(0)
	goto L131
L129:
	;
	goto L130
L130:
	;
	m.G0 = v27 + int32(832)
	return
L131:
	;
	v888 = v869 + v30
	v889 = int32(2)
	v891 = v31 + v869<<(uint(v889)%32)
	v894 = l1 + v869*int32(52)
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)+4))
	v898 = l2 + v895<<(uint(int32(3))%32)
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v899&v889 != 0 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L130
L133:
	;
	v1386 = v869 + int32(1)
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1386 < v1387 {
		v869 = v1386
		goto L131
	} else {
		goto L232
	}
L134:
	;
	v902 = int32(4515600)
	v903 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v907
	v911 = v904 + v895*int32(224)
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+20))
	if v912 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v1003 = int32(0)
	v1004 = int32(4515600)
	v1005 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1009
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v894)+44))
	if v1013 == v1003 {
		goto L175
	} else {
		goto L176
	}
L137:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911)+70)))
	if v914 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	goto L139
L139:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	if v979 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v949)+20)) = v950
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	v953 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v949)+16)) = uint8(v953)
	*(*uint8)(unsafe.Add(mBase, uint32(v949)+24)) = uint8(v952)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v949)))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	v958 = m.T0[v957].(func(*base.Module, int32) int32)(m, v949)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L12
	} else {
		goto L156
	}
L141:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938))))
	if v939 != int32(1) {
		v948 = v938
		goto L153
	} else {
		goto L154
	}
L142:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v949 = v936
	v950 = v937
	goto L140
L143:
	;
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v911)+184)))
	if v933 == int32(65535) {
		goto L141
	} else {
		goto L151
	}
L144:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v911)+216))
	if v913&int32(1) == int32(0) {
		v932 = v917
		goto L143
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if v913&int32(1) != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v936 = v917
	goto L142
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = int32(0)
	v926 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v926)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v903
	goto L133
L149:
	;
	goto L150
L150:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v911)+216))
	v932 = v930
	goto L143
L151:
	;
	v936 = v932
	goto L142
L152:
	;
	v949 = v932
	v950 = v948
	goto L140
L153:
	;
	goto L152
L154:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938)+1)))
	if v942 != int32(3) {
		v948 = v938
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v938)+2))
	v948 = v945 + int32(18)
	goto L153
L156:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v960)
	if v960 != 0 {
		v975 = v958
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v975
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v903
	goto L133
L158:
	;
	v962 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v894)+48)))
	if v962 != int32(65535) {
		v975 = v958
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
	if v965 != int32(1) {
		v974 = v958
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v975 = v974
	goto L157
L161:
	;
	goto L160
L162:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+1)))
	if v968 != int32(3) {
		v974 = v958
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v958)+2))
	v974 = v971 + int32(18)
	goto L161
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v997
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v999)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v903
	goto L133
L165:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986))))
	if v987 != int32(1) {
		v996 = v986
		goto L171
	} else {
		goto L172
	}
L166:
	;
	v982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v911)+184)))
	if v982 == int32(65535) {
		goto L165
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v997 = v985
	goto L164
L169:
	;
	goto L168
L170:
	;
	v997 = v996
	goto L164
L171:
	;
	goto L170
L172:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)))
	if v990 != int32(3) {
		v996 = v986
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v986)+2))
	v996 = v993 + int32(18)
	goto L171
L174:
	;
	v1097 = v895*int32(224) + v1006
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v894)+8))
	if v1098 != 0 {
		goto L184
	} else {
		goto L185
	}
L175:
	;
	v1077 = int32(1)
	v1082 = v1003
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1017 = int32(1)
	v1018 = int32(0)
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+4))
	if v1019 <= v1018 {
		v1077 = v1017
		v1082 = v1003
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v1026 = v1017
	v1028 = v1018
	v1031 = v1003
	goto L179
L179:
	;
	v1048 = v862 + v1026<<(uint(int32(3))%32)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+12))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1049+v1028<<(uint(int32(2))%32))))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+20))
	v1058 = m.T0[v1057].(func(*base.Module, int32, int32, int32) int32)(m, v1053, v1054, v1048+int32(4))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L12
	} else {
		goto L181
	}
L180:
	;
	v1077 = v1068
	v1082 = v1066
	goto L174
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1048))) = v1058
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048)+4)))
	v1062 = int32(1)
	v1066 = base.B2i32(v1061|v1031&v1062 != int32(0))
	v1068 = v1026 + v1062
	v1070 = v1028 + v1062
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+4))
	if v1070 < v1071 {
		v1026 = v1068
		v1028 = v1070
		v1031 = v1066
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1005
	goto L133
L184:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v894)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v894
	v1101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v894 + int32(12)
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+116))
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+30)) = uint16(v1099)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+28)) = uint8(v1101)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v1107
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	if v1112 == v1101 {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	goto L186
L186:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	if v1313 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L187:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+36)) = uint8(v1132)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v1131
	if v1099 <= v1077 {
		goto L198
	} else {
		goto L199
	}
L188:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119))))
	if v1120 != int32(1) {
		v1129 = v1119
		goto L194
	} else {
		goto L195
	}
L189:
	;
	v1115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1097)+184)))
	if v1115 == int32(65535) {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v1131 = v1118
	v1132 = v1112
	goto L187
L192:
	;
	goto L191
L193:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	v1131 = v1129
	v1132 = v1130
	goto L187
L194:
	;
	goto L193
L195:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119)+1)))
	if v1123 != int32(3) {
		v1129 = v1119
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+2))
	v1129 = v1126 + int32(18)
	goto L194
L197:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275)+10)))
	if v1276 != int32(1) {
		goto L211
	} else {
		goto L212
	}
L198:
	;
	v1260 = v1132 | v1082
	goto L197
L199:
	;
	goto L200
L200:
	;
	v1140 = (v1099 - v1077) & int32(3)
	if v1140 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1144 = int32(0)
	v1147 = v1077
	goto L204
L202:
	;
	v1183 = v1077
	goto L203
L203:
	;
	v1201 = int32(1)
	if base.Ui32(int32(-4)) < base.Ui32(v1077-v1099) {
		v1260 = v1201
		goto L197
	} else {
		goto L207
	}
L204:
	;
	v1167 = v862 + v1147<<(uint(int32(3))%32)
	v1168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1167)+4)) = uint8(v1168)
	*(*int32)(unsafe.Add(mBase, uint32(v1167))) = int32(0)
	v1173 = v1147 + v1168
	v1175 = v1144 + v1168
	if v1175 != v1140 {
		v1144 = v1175
		v1147 = v1173
		goto L204
	} else {
		goto L206
	}
L205:
	;
	v1183 = v1173
	goto L203
L206:
	;
	goto L205
L207:
	;
	v1211 = v1183
	goto L208
L208:
	;
	v1231 = v862 + v1211<<(uint(int32(3))%32)
	v1232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1231)+4)) = uint8(v1232)
	v1234 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1231))) = v1234
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+8)) = v1234
	*(*uint8)(unsafe.Add(mBase, uint32(v1231)+12)) = uint8(v1232)
	*(*uint8)(unsafe.Add(mBase, uint32(v1231)+20)) = uint8(v1232)
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+16)) = v1234
	*(*uint8)(unsafe.Add(mBase, uint32(v1231)+28)) = uint8(v1232)
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+24)) = v1234
	v1249 = v1211 + int32(4)
	if v1249 != v1099 {
		v1211 = v1249
		goto L208
	} else {
		goto L210
	}
L209:
	;
	v1260 = v1201
	goto L197
L210:
	;
	goto L209
L211:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1275)))
	v1292 = m.T0[v1291].(func(*base.Module, int32) int32)(m, v27+int32(12))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L12
	} else {
		goto L214
	}
L212:
	;
	if v1260&int32(1) == int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v1283 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v1283
	v1285 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v1285)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1283
	goto L183
L214:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v1294)
	if v1294 != 0 {
		v1309 = v1292
		goto L215
	} else {
		goto L216
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L183
L216:
	;
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v894)+48)))
	if v1296 != int32(65535) {
		v1309 = v1292
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292))))
	if v1299 != int32(1) {
		v1308 = v1292
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1309 = v1308
	goto L215
L219:
	;
	goto L218
L220:
	;
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+1)))
	if v1302 != int32(3) {
		v1308 = v1292
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+2))
	v1308 = v1305 + int32(18)
	goto L219
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v1331
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v1333)
	goto L183
L223:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v1321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1320))))
	if v1321 != int32(1) {
		v1330 = v1320
		goto L229
	} else {
		goto L230
	}
L224:
	;
	v1316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1097)+184)))
	if v1316 == int32(65535) {
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v1331 = v1319
	goto L222
L227:
	;
	goto L226
L228:
	;
	v1331 = v1330
	goto L222
L229:
	;
	goto L228
L230:
	;
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1320)+1)))
	if v1324 != int32(3) {
		v1330 = v1320
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+2))
	v1330 = v1327 + int32(18)
	goto L229
L232:
	;
	goto L132
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
	var v13 int32
	_ = v13
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
	v13 = int32(0)
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
	v22 = F_bms_add_member(m, v13, v21)
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
		v13 = v22
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
	F_errmsg_internal(m, int32(479031), v8)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(499718), int32(426), int32(307858))
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v7 = F_pgstat_fetch_pending_entry(m, int32(3), v5, base.I64_extend_i32_u(l0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v12 = v11
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	v10 = int32(-101)
	if base.Ui32(l1) <= base.Ui32(l0) {
		v334 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v334
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
		v334 = v10
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
	v21 = int32(543027)
	goto L9
L8:
	;
	v21 = int32(530675)
	goto L9
L9:
	;
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21))))
	v26 = l0
	goto L10
L10:
	;
	v32 = l1 - v26
	v33 = int32(0)
	v36 = base.B2i32(v32 != v33)
	if v26&int32(3) == v33 {
		v62 = v26
		v64 = v32
		v65 = v36
		goto L15
	} else {
		goto L16
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
	if base.Ui32(l1) <= base.Ui32(v138) {
		v236 = v138
		goto L67
	} else {
		goto L68
	}
L12:
	;
	if v135 == int32(0) {
		v334 = v10
		goto L1
	} else {
		goto L38
	}
L13:
	;
	v135 = int32(0)
	goto L12
L14:
	;
	v113 = v106
	v115 = v108
	goto L32
L15:
	;
	if v65 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L16:
	;
	if v32 == int32(0) {
		v62 = v26
		v64 = v32
		v65 = v36
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v45 = v26
	v47 = v32
	goto L18
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v50 == v22&int32(255) {
		v106 = v45
		v108 = v47
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v62 = v57
	v64 = v53
	v65 = v55
	goto L15
L20:
	;
	v52 = int32(1)
	v53 = v47 - v52
	v54 = int32(0)
	v55 = base.B2i32(v53 != v54)
	v57 = v45 + v52
	if v57&int32(3) == v54 {
		v62 = v57
		v64 = v53
		v65 = v55
		goto L15
	} else {
		goto L21
	}
L21:
	;
	if v53 != 0 {
		v45 = v57
		v47 = v53
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v69 == v22&int32(255) {
		v99 = v62
		v101 = v64
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v101 == int32(0) {
		goto L13
	} else {
		goto L31
	}
L25:
	;
	if base.Ui32(v64) < base.Ui32(int32(4)) {
		v99 = v62
		v101 = v64
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v79 = v62
	v81 = v64
	goto L27
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v86 = v85 ^ v22&int32(255)*int32(16843009)
	v89 = int32(-2139062144)
	if (int32(16843008)-v86|v86)&v89 != v89 {
		v106 = v79
		v108 = v81
		goto L14
	} else {
		goto L29
	}
L28:
	;
	v99 = v94
	v101 = v96
	goto L24
L29:
	;
	v93 = int32(4)
	v94 = v79 + v93
	v96 = v81 - v93
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v79 = v94
		v81 = v96
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v106 = v99
	v108 = v101
	goto L14
L32:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v22&int32(255) == v118 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L13
L34:
	;
	v135 = v113
	goto L12
L35:
	;
	goto L36
L36:
	;
	v120 = int32(1)
	v123 = v115 - v120
	if v123 != 0 {
		v113 = v113 + v120
		v115 = v123
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v138 = v14 + v135
	if base.Ui32(l1) < base.Ui32(v138) {
		v334 = v10
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v14) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if v201 != 0 {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v201 = int32(0)
	goto L40
L42:
	;
	v175 = v170
	v176 = v171
	v177 = v172
	goto L52
L43:
	;
	if (v135|v21)&int32(3) != 0 {
		v170 = v135
		v171 = v21
		v172 = v14
		goto L42
	} else {
		goto L46
	}
L44:
	;
	v163 = v135
	v164 = v21
	v165 = v14
	goto L45
L45:
	;
	if v165 == int32(0) {
		goto L41
	} else {
		goto L51
	}
L46:
	;
	v147 = v135
	v148 = v21
	v149 = v14
	goto L47
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v152 != v153 {
		v170 = v147
		v171 = v148
		v172 = v149
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v163 = v158
	v164 = v156
	v165 = v160
	goto L45
L49:
	;
	v155 = int32(4)
	v156 = v148 + v155
	v158 = v147 + v155
	v160 = v149 - v155
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v147 = v158
		v148 = v156
		v149 = v160
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v170 = v163
	v171 = v164
	v172 = v165
	goto L42
L52:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v180 == v181 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v201 = v180 - v181
	goto L40
L54:
	;
	v183 = int32(1)
	v188 = v177 - v183
	if v188 != 0 {
		v175 = v175 + v183
		v176 = v176 + v183
		v177 = v188
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L41
L58:
	;
	v203 = v135 + int32(1)
	if base.Ui32(v203) < base.Ui32(l1) {
		v26 = v203
		goto L10
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if l0 == v135 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v334 = v10
	goto L1
L62:
	;
	goto L11
L63:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135-int32(1)))))
	if v208 == int32(10) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if base.Ui32(l1) <= base.Ui32(v138) {
		v334 = v10
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v14 <= l1-v138 {
		v26 = v138
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v334 = v10
	goto L1
L67:
	;
	if l1-v236 < int32(5) {
		v334 = v10
		goto L1
	} else {
		goto L74
	}
L68:
	;
	v219 = v138
	goto L69
L69:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v225 == int32(45) {
		v236 = v219
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v236 = l1
	goto L67
L71:
	;
	if base.Ui32(v225) < base.Ui32(int32(32)) {
		v334 = v10
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v231 = v219 + int32(1)
	if base.Ui32(v231) < base.Ui32(l1) {
		v219 = v231
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v245 = int32(5)
	goto L78
L75:
	;
	if v307 != 0 {
		v334 = v10
		goto L1
	} else {
		goto L93
	}
L76:
	;
	v307 = int32(0)
	goto L75
L77:
	;
	v281 = v276
	v282 = v277
	v283 = v278
	goto L87
L78:
	;
	if (v236|v21)&int32(3) != 0 {
		v276 = v236
		v277 = v21
		v278 = v245
		goto L77
	} else {
		goto L81
	}
L80:
	;
	if v266 == int32(0) {
		goto L76
	} else {
		goto L86
	}
L81:
	;
	v253 = v236
	v254 = v21
	v255 = v245
	goto L82
L82:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v258 != v259 {
		v276 = v253
		v277 = v254
		v278 = v255
		goto L77
	} else {
		goto L84
	}
L83:
	;
	goto L80
L84:
	;
	v261 = int32(4)
	v262 = v254 + v261
	v264 = v253 + v261
	v266 = v255 - v261
	if base.Ui32(int32(3)) < base.Ui32(v266) {
		v253 = v264
		v254 = v262
		v255 = v266
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v276 = v264
	v277 = v262
	v278 = v266
	goto L77
L87:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v286 == v287 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v307 = v286 - v287
	goto L75
L89:
	;
	v289 = int32(1)
	v294 = v283 - v289
	if v294 != 0 {
		v281 = v281 + v289
		v282 = v282 + v289
		v283 = v294
		goto L87
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	goto L88
L92:
	;
	goto L76
L93:
	;
	v309 = v236 + int32(5)
	if base.Ui32(l1) <= base.Ui32(v309) {
		v324 = v309
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v334 = v324 - v135
	goto L1
L95:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	switch v311 - int32(10) {
	case 0, 3:
		goto L96
	default:
		v334 = v10
		goto L1
	}
L96:
	;
	if v311 == int32(13) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v318 = v236 + int32(6)
	goto L99
L98:
	;
	v318 = v309
	goto L99
L99:
	;
	if base.Ui32(l1) <= base.Ui32(v318) {
		v324 = v318
		goto L94
	} else {
		goto L100
	}
L100:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v324 = v318 + base.B2i32(v320 == int32(10))
	goto L94
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
	F_errmsg_internal(m, int32(486247), v8)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(499674), int32(4358), int32(307695))
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l2 == int32(0) {
		v147 = l1
		v149 = l3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v147 - v14
	if v147 == v14 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	if base.Ui32(l1) <= base.Ui32(v14) {
		v147 = l1
		v149 = l3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = l1
	v23 = l3
	goto L4
L4:
	;
	v31 = v21 - int32(1)
	v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
	goto L10
L5:
	;
	v147 = v14
	v149 = l3 + v14 - l1
	goto L1
L6:
	;
	if v136 == int32(0) {
		v147 = v21
		v149 = v23
		goto L1
	} else {
		goto L32
	}
L7:
	;
	v136 = int32(0)
	goto L6
L8:
	;
	v114 = v107
	v116 = v109
	goto L26
L9:
	;
	if base.B2i32(v54 != v55) == int32(0) {
		goto L7
	} else {
		goto L17
	}
L10:
	;
	goto L11
L11:
	;
	v46 = int32(690097)
	v48 = int32(4)
	goto L12
L12:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v51 == v32&int32(255) {
		v107 = v46
		v109 = v48
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v53 = int32(1)
	v54 = v48 - v53
	v55 = int32(0)
	v58 = v46 + v53
	if v58&int32(3) == v55 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v54 != 0 {
		v46 = v58
		v48 = v54
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v70 == v32&int32(255) {
		v100 = v58
		v102 = v54
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v102 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L19:
	;
	if base.Ui32(v54) < base.Ui32(int32(4)) {
		v100 = v58
		v102 = v54
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v80 = v58
	v82 = v54
	goto L21
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v87 = v86 ^ v32&int32(255)*int32(16843009)
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 != v90 {
		v107 = v80
		v109 = v82
		goto L8
	} else {
		goto L23
	}
L22:
	;
	v100 = v95
	v102 = v97
	goto L18
L23:
	;
	v94 = int32(4)
	v95 = v80 + v94
	v97 = v82 - v94
	if base.Ui32(int32(3)) < base.Ui32(v97) {
		v80 = v95
		v82 = v97
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v107 = v100
	v109 = v102
	goto L8
L26:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v32&int32(255) == v119 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L7
L28:
	;
	v136 = v114
	goto L6
L29:
	;
	goto L30
L30:
	;
	v121 = int32(1)
	v124 = v116 - v121
	if v124 != 0 {
		v114 = v114 + v121
		v116 = v124
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v140 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v139 - v140
	if base.Ui32(v14) < base.Ui32(v31) {
		v21 = v31
		v23 = v23 - v140
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L5
L34:
	;
	m.G0 = v12 + int32(32)
	return v224
L35:
	;
	F_errsave_finish(m, l4, int32(496428), v214, int32(290896))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L39
	} else {
		goto L54
	}
L36:
	;
	v159 = F_errsave_start(m, l4)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v183 < int32(1001) {
		v224 = int32(1)
		goto L34
	} else {
		goto L48
	}
L39:
	;
	return int32(0)
L40:
	;
	if v159 == int32(0) {
		v224 = int32(0)
		goto L34
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v149
	if l2 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v172 = int32(471241)
	goto L45
L44:
	;
	v172 = int32(471277)
	goto L45
L45:
	;
	F_errmsg(m, int32(0)+v172, v12)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	F_errdetail(m, int32(647135), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v214 = int32(612)
	goto L35
L48:
	;
	v187 = F_errsave_start(m, l4)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	if v187 == int32(0) {
		v224 = int32(0)
		goto L34
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(34103428))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(327967), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1000)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v200
	F_errdetail(m, int32(653675), v12+int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v214 = int32(618)
	goto L35
L54:
	;
	v224 = int32(0)
	goto L34
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
	F_errmsg_internal(m, int32(273665), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(496998), int32(5295), int32(427751))
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
	F_errmsg_internal(m, int32(273665), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(496998), int32(5322), int32(427751))
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
	F_errmsg_internal(m, int32(273665), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(496998), int32(5329), int32(427751))
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
	F_errmsg_internal(m, int32(74522), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(496998), int32(5305), int32(427751))
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
				F_errmsg(m, int32(467188), v9)
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
						F_errfinish(m, int32(497512), int32(112), int32(315562))
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
					F_errmsg(m, int32(467188), v9)
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
							F_errfinish(m, int32(497512), int32(112), int32(315562))
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
						F_errmsg(m, int32(467188), v9)
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
								F_errfinish(m, int32(497512), int32(112), int32(315562))
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
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v12 float32
	_ = v12
	var v13 float64
	_ = v13
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v12 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = base.F64_promote_f32(v12)
		v22 = base.F64_lt(v6, v13) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)))
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
	var v5 float32
	_ = v5
	var v6 float64
	_ = v6
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v22 int32
	_ = v22
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = base.F64_promote_f32(v5)
	if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		v22 = base.F64_ge(v6, v13) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
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
func F_format_elog_string(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4515600)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[425]))
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v30 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = v30
	goto L7
L5:
	;
	goto L6
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v56 = F_pstrdup(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	F_enlargeStringInfo(m, v9+int32(16), v36)
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
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v47 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v47 != 0 {
		v36 = v47
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
	m.G0 = v9 + int32(32)
	return v56
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	return v123
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
		v123 = v3
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
	F_appendStringInfo(m, v11+int32(32), int32(685689), v11+int32(16))
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
	v107 = m.ExcPending
	if v107 != 0 {
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
	F_appendStringInfoString(m, v11+int32(32), v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L25
	}
L20:
	;
	v54 = F_format_type_be_qualified(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v56 = F_format_type_be(m, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v58 = v54
	goto L19
L24:
	;
	v58 = v56
	goto L19
L25:
	;
	v61 = int32(1)
	if v23 == v61 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v66 = v61
	goto L27
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(136)+v66<<(uint(int32(2))%32))))
	F_appendStringInfoChar(m, v11+int32(32), int32(44))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	F_appendStringInfoString(m, v11+int32(32), v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L36
	}
L31:
	;
	v85 = F_format_type_be_qualified(m, v77)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v87 = F_format_type_be(m, v77)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L35
	}
L34:
	;
	v89 = v85
	goto L30
L35:
	;
	v89 = v87
	goto L30
L36:
	;
	v93 = v66 + int32(1)
	if v93 != v23 {
		v66 = v93
		goto L27
	} else {
		goto L37
	}
L37:
	;
	goto L28
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v123 = v108
	goto L1
L40:
	;
	v114 = F_palloc(m, int32(64))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v119 = F_pg_snprintf(m, v114, int32(64), int32(59441), v11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v123 = v114
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	v18 = l1 + int32(20)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v19 != 0 {
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	F_pfree(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v31)
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v33
	if l0 == v31 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_pfree(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_pfree(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
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
func F_fsm_set_avail(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
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
	v3 = l2
	v8 = l0 + int32(28)
	v10 = l1 + int32(4095)
	v11 = v8 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 != v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v3)
	v23 = v10
	goto L4
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.Ui32(v14) < base.Ui32(v3) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	v25 = int32(1)
	v26 = v23 - v25
	v27 = int32(2)
	v28 = base.I32_div_s(v26, v27)
	v30 = v28 << (uint(v25) % 32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v30)+1)))
	v34 = v30 + v27
	if base.Ui32(v34) <= base.Ui32(int32(8163)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.Ui32(v53) < base.Ui32(v3) {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v38 = v32 & int32(255)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v34))))
	if base.Ui32(v40) < base.Ui32(v38) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v43 = v32
	goto L8
L8:
	;
	v45 = v8 + v28
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != v43&int32(255) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v42 = v38
	goto L11
L10:
	;
	v42 = v40
	goto L11
L11:
	;
	v43 = v42
	goto L8
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v43)
	if int32(1) < v26 {
		v23 = v28
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
	v59 = int32(4094)
	goto L19
L17:
	;
	goto L18
L18:
	;
	return int32(1)
L19:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v59) {
		v81 = int32(0)
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v82 = v8 + v59
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v83 != v81&int32(255) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v66 = v59 << (uint(int32(1)) % 32)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v66)+1)))
	if v59 == int32(4081) {
		v81 = v68
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v72 = v68 & int32(255)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+(v66+int32(2))))))
	if base.Ui32(v76) < base.Ui32(v72) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = v72
	goto L26
L25:
	;
	v78 = v76
	goto L26
L26:
	;
	v81 = v78
	goto L21
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v81)
	goto L29
L28:
	;
	goto L29
L29:
	;
	if v59 != 0 {
		v59 = v59 - int32(1)
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
		*(*int32)(unsafe.Add(mBase, _consts[137])) = v2
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
