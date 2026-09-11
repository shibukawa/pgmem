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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v15 = int32(1)
	v17 = l0 + v15
	if base.Ui32(v17) <= base.Ui32(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v15
	goto L3
L2:
	;
	v20 = v17
	goto L3
L3:
	;
	v21 = int32(11)
	v22 = int32(base.Ui32(v20) >> (uint(v21) % 32))
	v23 = base.I64_extend_i32_u(v22)
	v27 = int32(base.Ui32(l0) >> (uint(v21) % 32))
	v28 = base.I64_extend_i32_u(v27)
	if v27 == v22 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+28))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v83 = base.I32_rem_u_s(v27, v82)
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[2])))
	if v31&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+48)) = v38
	if v28 != v38 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v43 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_errmsg_internal(m, int32(_a_F_RecordNewMultiXact_0), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v59 = base.I32_rem_u_s(v22, v58)
	v62 = v56 + v59<<(uint(int32(7))%32)
	v64 = F_LWLockAcquire(m, v62, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	F_errfinish(m, int32(_a_F_RecordNewMultiXact_1), int32(951), int32(_a_F_RecordNewMultiXact_2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v66 = int32(_a_F_RecordNewMultiXact_3)
	v68 = F_SimpleLruZeroPage(m, v66, v23)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_SimpleLruWritePage(m, v66, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	F_LWLockRelease(m, v62)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[3])) = v23
	goto L4
L19:
	;
	v92 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_3), v28, int32(1), l0)
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
	if v27 == v22 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v146 = v86
	v147 = v92
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
	v123 = base.I32_rem_u_s(v22, v122)
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
	v132 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_3), v23, int32(1), v20)
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
	v146 = v126
	v147 = v132
	v148 = v140 + v20&int32(2047)<<(uint(v137)%32)
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
	*(*uint8)(unsafe.Add(mBase, uint32(v159+v147))) = uint8(v161)
	goto L35
L34:
	;
	goto L35
L35:
	;
	F_LWLockRelease(m, v146)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v165 = int32(0)
	if l2 <= v165 {
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
	v173 = v165
	v176 = v147
	v178 = int32(0)
	v180 = int64(-1)
	goto L39
L39:
	;
	v186 = v170 << (uint(int32(3)) % 32) & int32(24)
	v187 = int32(2)
	v194 = base.I32_rem_u_s(int32(base.Ui32(v170)>>(uint(v187)%32)), int32(409))
	v196 = v194 * int32(20)
	v198 = base.I32_div_u_s(v170, int32(1636))
	v199 = base.I64_extend_i32_u(v198)
	if v199 != v180 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v222 == int32(0) {
		goto L37
	} else {
		goto L54
	}
L41:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+28))
	v205 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[5])))
	v206 = base.I32_rem_u_s(v198, v205)
	v209 = v203 + v206<<(uint(int32(7))%32)
	if v178 != v209 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v221 = v176
	v222 = v178
	v223 = v180
	goto L43
L43:
	;
	v225 = v221 << (uint(int32(2)) % 32)
	v226 = int32(_a_F_RecordNewMultiXact_4)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225+v228)))
	v235 = l3 + v173<<(uint(int32(3))%32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v230+v170<<(uint(v187)%32)&int32(12)+v196)+4)) = v236
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240+v225)))
	v243 = v242 + v196
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v244&(int32(255)<<(uint(v186)%32)^int32(-1)) | v250<<(uint(v186)%32)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[4]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v256+v221))) = uint8(v258)
	v263 = v173 + v258
	if v263 != l2 {
		v170 = v170 + v258
		v173 = v263
		v176 = v221
		v178 = v222
		v180 = v223
		goto L39
	} else {
		goto L53
	}
L44:
	;
	if v178 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v216 = v178
	goto L46
L46:
	;
	v219 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_4), v199, int32(1), l0)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L52
	}
L47:
	;
	F_LWLockRelease(m, v178)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v214 = F_LWLockAcquire(m, v209, int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v216 = v209
	goto L46
L52:
	;
	v221 = v219
	v222 = v216
	v223 = v199
	goto L43
L53:
	;
	goto L40
L54:
	;
	F_LWLockRelease(m, v222)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	goto L37
}
