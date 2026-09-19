package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MultiXactAdvanceOldest(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactAdvanceOldest[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5-l0 < int32(0) {
		F_SetMultiXactIdLimit(m, l0, l1, int32(0))
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
func F_RecordNewMultiXact(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
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
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v14 = int32(1)
	v16 = l0 + v14
	if base.Ui32(v16) <= base.Ui32(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = v14
	goto L3
L2:
	;
	v19 = v16
	goto L3
L3:
	;
	v20 = int32(11)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = base.I64_extend_i32_u(v21)
	v26 = int32(base.Ui32(l0) >> (uint(v20) % 32))
	v27 = base.I64_extend_i32_u(v26)
	if v26 == v21 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v83 = base.I32_rem_u_s(v26, v82)
	v86 = v80 + v83<<(uint(int32(7))%32)
	v88 = F_LWLockAcquire(m, v86, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L19
	}
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[2])))
	if v30&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v37 = int64(0)
	v40 = base.AtomicRmwCmpxchg64(m, v36, int32(48), v37, v37)
	if v40 != v27 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v44 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_errmsg_internal(m, int32(_a_F_RecordNewMultiXact_0), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v60 = base.I32_rem_u_s(v21, v59)
	v63 = v57 + v60<<(uint(int32(7))%32)
	v65 = F_LWLockAcquire(m, v63, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	F_errfinish(m, int32(_a_F_RecordNewMultiXact_1), int32(951), int32(_a_F_RecordNewMultiXact_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v67 = int32(_a_F_RecordNewMultiXact_3)
	v69 = F_SimpleLruZeroPage(m, v67, v22)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_SimpleLruWritePage(m, v67, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_LWLockRelease(m, v63)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[3])) = v22
	goto L4
L19:
	;
	v92 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_3), v27, int32(1), l0)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v97 = int32(2)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v92<<(uint(v97)%32))))
	v103 = v100 + l0&int32(2047)<<(uint(v97)%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if l1 != v104 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = l1
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v109+v92))) = uint8(v111)
	goto L23
L22:
	;
	goto L23
L23:
	;
	if v26 == v21 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v146 = v92
	v147 = v86
	v148 = v103 + int32(4)
	goto L26
L25:
	;
	F_LWLockRelease(m, v86)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	v149 = int32(1)
	v150 = l1 + l2
	if base.Ui32(v150) <= base.Ui32(v149) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v123 = base.I32_rem_u_s(v21, v122)
	v126 = v120 + v123<<(uint(int32(7))%32)
	v128 = F_LWLockAcquire(m, v126, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v132 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_3), v22, int32(1), v19)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v137 = int32(2)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v132<<(uint(v137)%32))))
	v146 = v132
	v147 = v126
	v148 = v140 + v19&int32(2047)<<(uint(v137)%32)
	goto L26
L30:
	;
	v153 = v149
	goto L32
L31:
	;
	v153 = v150
	goto L32
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v153 != v154 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v153
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159+v146))) = uint8(v161)
	goto L35
L34:
	;
	goto L35
L35:
	;
	F_LWLockRelease(m, v147)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if l2 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return
L38:
	;
	v170 = l1
	v173 = v146
	v175 = int32(0)
	v176 = int32(0)
	v180 = int64(-1)
	goto L39
L39:
	;
	v183 = base.I32_div_u_s(v170, int32(1636))
	v184 = base.I64_extend_i32_u(v183)
	if v184 != v180 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v207 == int32(0) {
		goto L37
	} else {
		goto L54
	}
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+28))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[5])))
	v191 = base.I32_rem_u_s(v183, v190)
	v194 = v188 + v191<<(uint(int32(7))%32)
	if v176 != v194 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v206 = v173
	v207 = v176
	v208 = v180
	goto L43
L43:
	;
	v209 = int32(2)
	v212 = base.I32_rem_u_s(int32(base.Ui32(v170)>>(uint(v209)%32)), int32(409))
	v214 = v212 * int32(20)
	v216 = v206 << (uint(v209) % 32)
	v217 = int32(_a_F_RecordNewMultiXact_4)
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216+v219)))
	v228 = int32(3)
	v230 = l3 + v175<<(uint(v228)%32)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	*(*int32)(unsafe.Add(mBase, uint32(v214+(v221+v170<<(uint(v209)%32)&int32(12)))+4)) = v231
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v235+v216)))
	v238 = v237 + v214
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v244 = v170 << (uint(v228) % 32) & int32(24)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v239&(int32(255)<<(uint(v244)%32)^int32(-1)) | v249<<(uint(v244)%32)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v257 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v255+v206))) = uint8(v257)
	v262 = v175 + v257
	if v262 != l2 {
		v170 = v170 + v257
		v173 = v206
		v175 = v262
		v176 = v207
		v180 = v208
		goto L39
	} else {
		goto L53
	}
L44:
	;
	if v176 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v201 = v176
	goto L46
L46:
	;
	v204 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_4), v184, int32(1), l0)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L52
	}
L47:
	;
	F_LWLockRelease(m, v176)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v199 = F_LWLockAcquire(m, v194, int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v201 = v194
	goto L46
L52:
	;
	v206 = v204
	v207 = v201
	v208 = v184
	goto L43
L53:
	;
	goto L40
L54:
	;
	F_LWLockRelease(m, v207)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	goto L37
}
