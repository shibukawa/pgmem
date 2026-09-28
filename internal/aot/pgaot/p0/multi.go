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
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactAdvanceOldest[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if v5-l0 < int32(0) {
		F_SetMultiXactIdLimit(m, l0, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_MultiXactMemberIoErrorDetail(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v9
		v15 = F_errdetail(m, int32(_a_F_MultiXactMemberIoErrorDetail_0), v6+int32(16))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v23 = v15
			m.G0 = v6 + int32(32)
			return v23
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v8
		v21 = F_errdetail(m, int32(_a_F_MultiXactMemberIoErrorDetail_1), v6)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v21
			m.G0 = v6 + int32(32)
			return v23
		}
	}
}
func F_RecordNewMultiXact(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v183 int64
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l0
	v19 = int32(1)
	v20 = l0 + v19
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = v20
	goto L3
L2:
	;
	v22 = v19
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v22
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v28 = int32(base.Ui32(l0) >> (uint(int32(10)) % 32))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v31 = base.I32_rem_u_s(v28, v30)
	v34 = v26 + v31<<(uint(int32(7))%32)
	v36 = F_LWLockAcquire(m, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v43 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_0), base.I64_extend_i32_u(v28), int32(1), v16+int32(28))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v43<<(uint(int32(2))%32))))
	v56 = v51 + l0&int32(1023)<<(uint(int32(3))%32)
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	if l1 != v57 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = l1
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v43))) = uint8(v64)
	goto L9
L8:
	;
	goto L9
L9:
	;
	v67 = int32(base.Ui32(v22) >> (uint(int32(10)) % 32))
	if v28 == v67 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v104 = v43
	v105 = v34
	v106 = v56 + int32(8)
	goto L12
L11:
	;
	F_LWLockRelease(m, v34)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v107 = int64(1)
	v109 = l1 + base.I64_extend_i32_s(l2)
	if base.Ui64(v109) <= base.Ui64(v107) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v77 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[1])))
	v78 = base.I32_rem_u_s(v67, v77)
	v81 = v75 + v78<<(uint(int32(7))%32)
	v83 = F_LWLockAcquire(m, v81, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v90 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_0), base.I64_extend_i32_u(v67), int32(1), v16+int32(24))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v90<<(uint(int32(2))%32))))
	v104 = v90
	v105 = v81
	v106 = v98 + v22&int32(1023)<<(uint(int32(3))%32)
	goto L12
L16:
	;
	v112 = v107
	goto L18
L17:
	;
	v112 = v109
	goto L18
L18:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v106)))
	if v112 != v113 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = v112
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[0]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118+v104))) = uint8(v120)
	goto L21
L20:
	;
	goto L21
L21:
	;
	F_LWLockRelease(m, v105)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if l2 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	m.G0 = v16 + int32(32)
	return
L24:
	;
	v128 = int32(0)
	v129 = l1
	v132 = v104
	v137 = int32(0)
	v140 = int64(-1)
	goto L25
L25:
	;
	v141 = base.I32_wrap_i64(v129)
	v145 = v141 << (uint(int32(3)) % 32) & int32(24)
	v147 = base.I64_div_u_s(v129, int64(1636))
	if v140 != v147 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v178 == int32(0) {
		goto L23
	} else {
		goto L40
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v129
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[2]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	v158 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[3])))
	v159 = base.I64_rem_u_s(v147, v158)
	v163 = v156 + base.I32_wrap_i64(v159)<<(uint(int32(7))%32)
	if v137 != v163 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v177 = v132
	v178 = v137
	v179 = v140
	goto L29
L29:
	;
	v183 = base.I64_rem_u_s(int64(base.Ui64(v129)>>(uint(int64(2))%64)), int64(409))
	v186 = base.I32_wrap_i64(v183) * int32(20)
	v187 = int32(2)
	v192 = v177 << (uint(v187) % 32)
	v193 = int32(_a_F_RecordNewMultiXact_1)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[2]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192+v195)))
	v202 = l3 + v128<<(uint(int32(3))%32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v186+(v141<<(uint(v187)%32)&int32(12)+v197))+4)) = v203
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[2]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207+v192)))
	v210 = v209 + v186
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v211&(int32(255)<<(uint(v145)%32)^int32(-1)) | v217<<(uint(v145)%32)
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_RecordNewMultiXact[2]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v223+v177))) = uint8(v225)
	v230 = v128 + v225
	if v230 != l2 {
		v128 = v230
		v129 = v129 + int64(1)
		v132 = v177
		v137 = v178
		v140 = v179
		goto L25
	} else {
		goto L39
	}
L30:
	;
	if v137 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v170 = v137
	goto L32
L32:
	;
	v175 = F_SimpleLruReadPage(m, int32(_a_F_RecordNewMultiXact_2), v147, int32(1), v16+int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L38
	}
L33:
	;
	F_LWLockRelease(m, v137)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v168 = F_LWLockAcquire(m, v163, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v170 = v163
	goto L32
L38:
	;
	v177 = v175
	v178 = v170
	v179 = v147
	goto L29
L39:
	;
	goto L26
L40:
	;
	F_LWLockRelease(m, v178)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	goto L23
}
