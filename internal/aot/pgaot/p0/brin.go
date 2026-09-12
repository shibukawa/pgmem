package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_brinRevmapInitialize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v7 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_LockBuffer(m, v7, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v7 < int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v7^int32(-1))<<(uint(int32(2))%32))))
				v31 = v23
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v31 = v25 + v7<<(uint(int32(13))%32) + int32(-8192)
			}
			v33 = F_palloc(m, int32(20))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
				v37 = v31 + int32(32)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
				v41 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v41
				*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v40
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v45
				F_LockBuffer(m, v7, v41)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					return v33
				}
			}
		}
	}
}
func F_brin_bloom_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v254 int64
	_ = v254
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v273 int64
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v20 <= int32(0) {
		v310 = int32(1)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v18 + int32(16)
	return v310
L4:
	;
	v40 = int32(0)
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22+v40<<(uint(int32(2))%32))))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)))
	if v55 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v310 = v305
	goto L3
L7:
	;
	v305 = int32(1)
	v307 = v40 + v305
	if v307 != v20 {
		v40 = v307
		goto L5
	} else {
		goto L33
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
	v60 = F_bloom_get_procinfo(m, v24, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L30
	}
L11:
	;
	v62 = F_FunctionCall1Coll(m, v60, v23, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L16
L13:
	;
	v150 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)))
	v151 = base.I64_rem_u_s(base.I64_extend_i32_u(v140)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v140^v132-base.I32_rotl(v140, int32(24))), v150)
	goto L21
L14:
	;
	v117 = int32(14)
	v119 = v105 ^ (v103 - v102 ^ base.I32_rotl(v102, int32(4))) - base.I32_rotl(v105, v117)
	v124 = v119 ^ (v62 + v104) - base.I32_rotl(v119, int32(11))
	v128 = v124 ^ v105 - base.I32_rotl(v124, int32(25))
	v132 = v128 ^ v119 - base.I32_rotl(v128, int32(16))
	v136 = v132 ^ v124 - base.I32_rotl(v132, int32(4))
	v140 = v136 ^ v128 - base.I32_rotl(v136, v117)
	goto L13
L16:
	;
	goto L17
L17:
	;
	v75 = base.I32_wrap_i64(int64(1910056111))
	v80 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
	v86 = v75 - v80 - int32(1636608428) ^ base.I32_rotl(v80, int32(6))
	v88 = v75 + int32(1021750440)
	v89 = v80 + v88
	v90 = v86 + v89
	v94 = v88 - v86 ^ base.I32_rotl(v86, int32(8))
	v98 = v89 - v94 ^ base.I32_rotl(v94, int32(16))
	v102 = v90 - v98 ^ base.I32_rotl(v98, int32(19))
	v103 = v94 + v90
	v104 = v98 + v103
	v105 = v102 + v104
	goto L14
L18:
	;
	v238 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)))
	v239 = base.I64_rem_u_s(base.I64_extend_i32_u(v228)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v228^v220-base.I32_rotl(v228, int32(24))), v238)
	v240 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	if v240 == int64(0) {
		goto L7
	} else {
		goto L23
	}
L19:
	;
	v205 = int32(14)
	v207 = v193 ^ (v191 - v190 ^ base.I32_rotl(v190, int32(4))) - base.I32_rotl(v193, v205)
	v212 = v207 ^ (v62 + v192) - base.I32_rotl(v207, int32(11))
	v216 = v212 ^ v193 - base.I32_rotl(v212, int32(25))
	v220 = v216 ^ v207 - base.I32_rotl(v216, int32(16))
	v224 = v220 ^ v212 - base.I32_rotl(v220, int32(4))
	v228 = v224 ^ v216 - base.I32_rotl(v224, v205)
	goto L18
L21:
	;
	goto L22
L22:
	;
	v163 = base.I32_wrap_i64(int64(3125326612))
	v168 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
	v174 = v163 - v168 - int32(1636608428) ^ base.I32_rotl(v168, int32(6))
	v176 = v163 + int32(1021750440)
	v177 = v168 + v176
	v178 = v174 + v177
	v182 = v176 - v174 ^ base.I32_rotl(v174, int32(8))
	v186 = v177 - v182 ^ base.I32_rotl(v182, int32(16))
	v190 = v178 - v186 ^ base.I32_rotl(v186, int32(19))
	v191 = v182 + v178
	v192 = v186 + v191
	v193 = v190 + v192
	goto L19
L23:
	;
	v254 = int64(0)
	goto L24
L24:
	;
	v261 = base.I64_rem_u_s(v254*v239+v151, v238)
	v262 = base.I32_wrap_i64(v261)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(16)+int32(base.Ui32(v262)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v266)>>(uint(v262&int32(7))%32))&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v310 = int32(0)
	goto L3
L26:
	;
	v273 = v254 + int64(1)
	if v240 != v273 {
		v254 = v273
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L7
L30:
	;
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v280
	F_errmsg_internal(m, int32(471508), v18)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(497268), int32(644), int32(92301))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	goto L6
}
func F_brin_bloom_summary_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(17894)
			F_errmsg(m, int32(192507), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(497268), int32(784), int32(279050))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_brin_build_desc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v17 = F_AllocSetContextCreateInternal(m, v12, int32(64276), v2, int32(1024), int32(8192))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(4515488)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v29 = F_palloc(m, v26<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v31 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v84 = F_palloc(m, v74<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v74 = v31
	v75 = v2
	goto L4
L6:
	;
	goto L7
L7:
	;
	v37 = v2
	v40 = v31
	v41 = v2
	goto L8
L8:
	;
	v49 = int32(1)
	v50 = v37 + v49
	v53 = F_index_getprocinfo(m, l0, base.I32_extend16_s(v50), v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v74 = v68
	v75 = v67
	goto L4
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(88)+v40<<(uint(int32(4))%32)+v37*int32(100))))
	v63 = F_FunctionCall1Coll(m, v53, int32(0), v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v37<<(uint(int32(2))%32)))) = v63
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
	v67 = v41 + v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v50 < v68 {
		v37 = v50
		v40 = v68
		v41 = v67
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = v75
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v17
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v87 < v92 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v99 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	F_pfree(m, v29)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v109 = v99 << (uint(int32(2)) % 32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v29+v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v84+int32(20)+v109))) = v112
	v115 = v99 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v115 < v116 {
		v99 = v115
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	return v84
}
func F_brin_can_do_samepage_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	if base.Ui32(l1) < base.Ui32(l2) {
		if l0 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[5]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10+(l0^int32(-1))<<(uint(int32(2))%32))))
			v24 = v16
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			v24 = v18 + l0<<(uint(int32(13))%32) + int32(-8192)
		}
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)))
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)))
		v27 = v25 - v26
		v28 = int32(0)
		if v28 < v27 {
			v31 = v27
		} else {
			v31 = v28
		}
		v34 = base.B2i32(base.Ui32(l2-l1) <= base.Ui32(v31))
	} else {
		v34 = int32(1)
	}
	return v34
}
func F_brin_desummarize_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
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
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
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
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v343 int64
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v24 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L12
	} else {
		goto L144
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L12
	} else {
		goto L140
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L12
	} else {
		goto L136
	}
L4:
	;
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+316))
	v32 = base.B2i32(v30 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v32)
	v34 = v32
	goto L7
L6:
	;
	v34 = int32(0)
	goto L7
L7:
	;
	goto L4
L8:
	;
	if base.Ui64(int64(4294967295)) <= base.Ui64(v20) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L12
	} else {
		goto L131
	}
L11:
	;
	v40 = F_IndexGetRelation(m, v21, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = F_table_open(m, v40, int32(4))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v48 = int32(0)
	goto L16
L16:
	;
	v50 = F_index_open(m, v21, int32(4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v48 = v45
	goto L16
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+119)))
	if v53 != int32(105) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+84))
	if v56 != int32(3580) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v62 = F_object_ownercheck(m, int32(1259), v21, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	F_aclcheck_error(m, int32(2), int32(20), v68+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v48 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v76 = F_IndexGetRelation(m, v21, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v76 != v40 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+192))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+18)))
	if v80 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_relation_close(m, v50, int32(4))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L12
	} else {
		goto L129
	}
L30:
	;
	v83 = base.I32_wrap_i64(v20)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v424 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L12
	} else {
		goto L124
	}
L33:
	;
	v98 = m.G0
	v100 = v98 - int32(16)
	m.G0 = v100
	v103 = F_ReadBuffer(m, v50, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L36
	}
L34:
	;
	goto L29
L35:
	;
	if v377 == int32(0) {
		goto L33
	} else {
		goto L123
	}
L36:
	;
	F_LockBuffer(m, v103, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v103 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v127 = F_palloc(m, int32(20))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L42
	}
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+(v103^int32(-1))<<(uint(int32(2))%32))))
	v125 = v117
	goto L38
L40:
	;
	goto L41
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v125 = v119 + v103<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v50
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+36))
	v133 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v132
	F_LockBuffer(m, v103, v133)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v142 = base.I32_div_u_s(v83, v141)
	v144 = base.I32_div_u_s(v142, int32(1360))
	if base.Ui32(v140) <= base.Ui32(v144) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L12
	} else {
		goto L119
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L12
	} else {
		goto L115
	}
L46:
	;
	m.G0 = v100 + int32(16)
	goto L35
L47:
	;
	F_pfree(m, v127)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L12
	} else {
		goto L114
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	F_ReleaseBuffer(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v154 = F_brinLockRevmapPageForUpdate(m, v127, v83)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L55
	}
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v149 == int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	F_ReleaseBuffer(m, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v175 = base.I32_div_u_s(v83, v174)
	v177 = base.I32_rem_u_s(v175, int32(1360))
	v182 = v173 + v177*int32(6) + int32(24)
	if v182 != 0 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	if v154 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+(v154^int32(-1))<<(uint(int32(2))%32))))
	v173 = v165
	goto L54
L57:
	;
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v173 = v167 + v154<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L59:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+2)))
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182))))
	v200 = F_ReadBuffer(m, v50, v195|v196<<(uint(int32(16))%32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L68
	}
L60:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+4)))
	if v183 != 0 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_LockBuffer(m, v154, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L12
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	F_ReleaseBuffer(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v190 == int32(0) {
		goto L47
	} else {
		goto L66
	}
L66:
	;
	F_ReleaseBuffer(m, v190)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	goto L47
L68:
	;
	F_LockBuffer(m, v200, int32(2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	if v200 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+16)))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223+v222)+6)))
	if v225 != int32(61587) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v208+(v200^int32(-1))<<(uint(int32(2))%32))))
	v222 = v214
	goto L70
L72:
	;
	goto L73
L73:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v222 = v216 + v200<<(uint(int32(13))%32) + int32(-8192)
	goto L70
L74:
	;
	v228 = int32(0)
	F_LockBuffer(m, v154, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+4)))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v244) {
		goto L85
	} else {
		goto L86
	}
L77:
	;
	F_LockBuffer(m, v200, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L12
	} else {
		goto L78
	}
L78:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	F_ReleaseBuffer(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v238 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_ReleaseBuffer(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_pfree(m, v127)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v377 = v228
	goto L46
L85:
	;
	v252 = int32(base.Ui32(v244+int32(262120)) >> (uint(int32(2)) % 32))
	goto L87
L86:
	;
	v252 = int32(0)
	goto L87
L87:
	;
	if base.Ui32(v252&int32(65535)) < base.Ui32(v243) {
		goto L45
	} else {
		goto L88
	}
L88:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243<<(uint(int32(2))%32)+v222)+21)))
	if v259&int32(384) == int32(0) {
		goto L44
	} else {
		goto L89
	}
L89:
	;
	v264 = int32(4510148)
	v266 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v266 + int32(1)
	if v154 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v289 = base.I32_div_u_s(v83, v288)
	v291 = base.I32_rem_u_s(v289, int32(1360))
	v294 = v287 + v291*int32(6)
	v295 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v294)+28)) = uint16(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+24)) = int32(-1)
	F_PageIndexTupleDeleteNoCompact(m, v222, v243)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L94
	}
L91:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+(v154^int32(-1))<<(uint(int32(2))%32))))
	v287 = v279
	goto L90
L92:
	;
	goto L93
L93:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v287 = v281 + v154<<(uint(int32(13))%32) + int32(-8192)
	goto L90
L94:
	;
	F_MarkBufferDirty(m, v200)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	F_MarkBufferDirty(m, v154)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+118)))
	if v306 != int32(112) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v347 = int32(4510148)
	v349 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v349 - int32(1)
	F_UnlockReleaseBuffer(m, v200)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L12
	} else {
		goto L109
	}
L98:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v310 <= int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	if v313 != 0 {
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v100)+12)) = uint16(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v315
	F_XLogBeginInsert(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L104
	}
L102:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	if v314 != 0 {
		goto L97
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	F_XLogRegisterData(m, v100+int32(4), int32(10))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v326 = int32(0)
	F_XLogRegisterBuffer(m, v326, v154, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	F_XLogRegisterBuffer(m, int32(1), v200, int32(8))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	v336 = F_XLogInsert(m, int32(17), int32(80))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	v338 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v173))) = base.I64_rotr(v336, v338)
	*(*uint32)(unsafe.Add(mBase, uint32(v222)+4)) = uint32(v336)
	v343 = int64(base.Ui64(v336) >> (uint(v338) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v222))) = uint32(v343)
	goto L97
L109:
	;
	F_LockBuffer(m, v154, int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L12
	} else {
		goto L110
	}
L110:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	F_ReleaseBuffer(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L111
	}
L111:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v361 == int32(0) {
		goto L47
	} else {
		goto L112
	}
L112:
	;
	F_ReleaseBuffer(m, v361)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	goto L47
L114:
	;
	v377 = int32(1)
	goto L46
L115:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(238561), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L12
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(496088), int32(383), int32(403395))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(238561), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(496088), int32(389), int32(403395))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L12
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	goto L34
L124:
	;
	if v424 == int32(0) {
		goto L29
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v431 + int32(4)
	F_errmsg(m, int32(436053), v15+int32(-32))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L12
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(496654), int32(1569), int32(401808))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	goto L29
L129:
	;
	F_relation_close(m, v48, int32(4))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	m.G0 = v17 - int32(-64)
	return int32(0)
L131:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(128103), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L12
	} else {
		goto L133
	}
L133:
	;
	F_errhint(m, int32(573127), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(496654), int32(1505), int32(401808))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v20
	F_errmsg(m, int32(430588), v17)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L12
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(496654), int32(1511), int32(401808))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L12
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v513 + int32(4)
	F_errmsg(m, int32(29085), v15+int32(-16))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L12
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(496654), int32(1537), int32(401808))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v534 + int32(4)
	F_errmsg(m, int32(696317), v15+int32(-48))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(496654), int32(1553), int32(401808))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_doupdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
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
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v527 int64
	_ = v527
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	v6 = l5
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	if base.Ui32(l9) < base.Ui32(int32(8153)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L6
	} else {
		goto L204
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L6
	} else {
		goto L201
	}
L3:
	;
	F_brinRevmapExtend(m, l2, l3)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L6
	} else {
		goto L197
	}
L6:
	;
	return int32(0)
L7:
	;
	if l10 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	m.G0 = v21 + int32(48)
	return v586
L9:
	;
	if l4 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v68 = v66
	v69 = int32(-1)
	goto L9
L11:
	;
	v33 = F_brin_getinsertbuffer(m, l0, l4, l9, v21+int32(47))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LockBuffer(m, l4, int32(2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L23
	}
L14:
	;
	if v33 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v586 = int32(0)
	goto L8
L16:
	;
	goto L17
L17:
	;
	if l4 == v33 {
		v66 = int32(0)
		goto L10
	} else {
		goto L18
	}
L18:
	;
	if v33 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v68 = v33
	v69 = v58
	goto L9
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v33^int32(-1))<<(uint(int32(6))%32))+16))
	v58 = v49
	goto L19
L21:
	;
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+v33<<(uint(int32(6))%32)+int32(-64))+16))
	v58 = v57
	goto L19
L23:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)) = uint8(v62)
	v66 = v62
	goto L10
L24:
	;
	v135 = v87 + v108&int32(32767)
	v137 = int32(base.Ui32(v108) >> (uint(int32(17)) % 32))
	if l7 == v137 {
		goto L45
	} else {
		goto L46
	}
L25:
	;
	v114 = int32(0)
	F_LockBuffer(m, l4, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L36
	}
L26:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+16)))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+v87)+6)))
	if v90 != int32(61587) {
		goto L25
	} else {
		goto L30
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73+(l4^int32(-1))<<(uint(int32(2))%32))))
	v87 = v79
	goto L26
L28:
	;
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v87 = v81 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v93) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = int32(base.Ui32(v93+int32(262120)) >> (uint(int32(2)) % 32))
	goto L33
L32:
	;
	v101 = int32(0)
	goto L33
L33:
	;
	if base.Ui32(v101&int32(65535)) < base.Ui32(v6) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32)+v87)+20))
	if v108&int32(98304) == int32(32768) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	goto L25
L36:
	;
	if v68 == int32(0) {
		v586 = v114
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v68)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L42
	}
L41:
	;
	v586 = v114
	goto L8
L42:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v586 = v114
	goto L8
L45:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v137) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v202 = int32(1)
	goto L47
L47:
	;
	if v202 != 0 {
		goto L66
	} else {
		goto L67
	}
L48:
	;
	v202 = v200
	goto L47
L49:
	;
	v200 = int32(0)
	goto L48
L50:
	;
	v174 = v169
	v175 = v170
	v176 = v171
	goto L60
L51:
	;
	if (v135|l6)&int32(3) != 0 {
		v169 = v135
		v170 = l6
		v171 = v137
		goto L50
	} else {
		goto L54
	}
L52:
	;
	v162 = v135
	v163 = l6
	v164 = v137
	goto L53
L53:
	;
	if v164 == int32(0) {
		goto L49
	} else {
		goto L59
	}
L54:
	;
	v146 = v135
	v147 = l6
	v148 = v137
	goto L55
L55:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v151 != v152 {
		v169 = v146
		v170 = v147
		v171 = v148
		goto L50
	} else {
		goto L57
	}
L56:
	;
	v162 = v157
	v163 = v155
	v164 = v159
	goto L53
L57:
	;
	v154 = int32(4)
	v155 = v147 + v154
	v157 = v146 + v154
	v159 = v148 - v154
	if base.Ui32(int32(3)) < base.Ui32(v159) {
		v146 = v157
		v147 = v155
		v148 = v159
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v169 = v162
	v170 = v163
	v171 = v164
	goto L50
L60:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v179 == v180 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v200 = v179 - v180
	goto L48
L62:
	;
	v182 = int32(1)
	v187 = v176 - v182
	if v187 != 0 {
		v174 = v174 + v182
		v175 = v175 + v182
		v176 = v187
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L61
L65:
	;
	goto L49
L66:
	;
	v203 = int32(0)
	F_LockBuffer(m, l4, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+16)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v222)+4)))
	if v224&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L69:
	;
	if v68 == int32(0) {
		v586 = v203
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v209 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v68)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L75
	}
L74:
	;
	v586 = v203
	goto L8
L75:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v586 = v203
	goto L8
L78:
	;
	if v68 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L79:
	;
	if base.Ui32(l7) < base.Ui32(l9) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if l4 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	v255 = int32(4510148)
	v257 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v257 + int32(1)
	v261 = F_PageIndexTupleOverwrite(m, v87, v6, l8, l9)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L92
	}
L83:
	;
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+14)))
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+12)))
	v248 = v246 - v247
	v249 = int32(0)
	if v249 < v248 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231+(l4^int32(-1))<<(uint(int32(2))%32))))
	v245 = v237
	goto L83
L85:
	;
	goto L86
L86:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v245 = v239 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L83
L87:
	;
	if base.Ui32(v252) < base.Ui32(l9-l7) {
		goto L78
	} else {
		goto L91
	}
L88:
	;
	v252 = v248
	goto L90
L89:
	;
	v252 = v249
	goto L90
L90:
	;
	goto L87
L91:
	;
	goto L82
L92:
	;
	if v261 == int32(0) {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+118)))
	if v268 != int32(112) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v299 = int32(4510148)
	v301 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v301 - int32(1)
	F_LockBuffer(m, l4, int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L107
	}
L96:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v272 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v275 != 0 {
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+24)) = uint16(v6)
	F_XLogBeginInsert(m)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v276 != 0 {
		goto L95
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	F_XLogRegisterData(m, v21+int32(24), int32(2))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_XLogRegisterBuffer(m, int32(0), l4, int32(8))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_XLogRegisterBufData(m, int32(0), l8, l9)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v294 = F_XLogInsert(m, int32(17), int32(48))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = base.I64_rotr(v294, int64(32))
	goto L95
L107:
	;
	v308 = int32(1)
	if v68 == int32(0) {
		v586 = v308
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v311 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L6
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_brin_initialize_empty_new_buffer(m, l0, v68)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L6
	} else {
		goto L113
	}
L112:
	;
	v586 = v308
	goto L8
L113:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	v586 = v308
	goto L8
L116:
	;
	v326 = int32(0)
	F_LockBuffer(m, l4, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v68 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v586 = v326
	goto L8
L120:
	;
	v348 = F_brinLockRevmapPageForUpdate(m, l2, l3)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L124
	}
L121:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333+(v68^int32(-1))<<(uint(int32(2))%32))))
	v347 = v339
	goto L120
L122:
	;
	goto L123
L123:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v347 = v341 + v68<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L124:
	;
	v350 = int32(4510148)
	v352 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v353 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v352 + v353
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v356 == v353 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	if v347&int32(3) != 0 {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	goto L127
L127:
	;
	F_PageIndexTupleDeleteNoCompact(m, v87, v6)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L138
	}
L128:
	;
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)))
	v402 = int32(61587)
	*(*uint16)(unsafe.Add(mBase, uint32(v347+v400)+6)) = uint16(v402)
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+10)) = int32(1572864)
	v391 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+18)) = uint16(v391)
	v397 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)) = uint16(v397)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+14)) = uint16(v397)
	goto L128
L130:
	;
	v385 = F___memset(m, v347, int32(0), int32(8192))
	mBase = m.M
	goto L129
L131:
	;
	goto L130
L138:
	;
	v406 = int32(0)
	v408 = F_PageAddItemExtended(m, v347, l8, l9, v406, v406)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	if v408 == int32(0) {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	F_MarkBufferDirty(m, v68)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	if v356 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v416 = int32(0)
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)))
	v418 = v347 + v417
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+6)))
	if v419 != int32(61587) {
		v434 = v416
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v436 = int32(0)
	goto L145
L145:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v408)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v408)
	v439 = int32(16)
	v440 = base.I32_rotr(v69, v439)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v440
	v444 = v21 + v439
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444))))
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+2)))
	if v348 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L146:
	;
	v436 = v434
	goto L145
L147:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+4)))
	if v422&int32(1) != 0 {
		v434 = v416
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v425 = int32(4)
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+14)))
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+12)))
	v428 = v426 - v427
	if v428 <= v425 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v434 = v431 - int32(4)
	goto L146
L150:
	;
	v431 = v425
	goto L152
L151:
	;
	v431 = v428
	goto L152
L152:
	;
	goto L149
L153:
	;
	F_MarkBufferDirty(m, v348)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L6
	} else {
		goto L164
	}
L154:
	;
	v467 = base.I32_div_u_s(l3, l1)
	v469 = base.I32_rem_u_s(v467, int32(1360))
	v472 = v466 + v469*int32(6)
	v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v444)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v472)+28)) = uint16(v473)
	if v473 != 0 {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v452+(v348^int32(-1))<<(uint(int32(2))%32))))
	v466 = v458
	goto L154
L156:
	;
	goto L157
L157:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v466 = v460 + v348<<(uint(int32(13))%32) + int32(-8192)
	goto L154
L158:
	;
	v476 = v448
	goto L160
L159:
	;
	v476 = int32(-1)
	goto L160
L160:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v472)+26)) = uint16(v476)
	if v473 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v479 = v447
	goto L163
L162:
	;
	v479 = int32(-1)
	goto L163
L163:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v472)+24)) = uint16(v479)
	goto L153
L164:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+118)))
	if v484 != int32(112) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v560 = int32(4510148)
	v562 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v562 - int32(1)
	F_LockBuffer(m, v348, int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L6
	} else {
		goto L189
	}
L166:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v488 <= int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v491 != 0 {
		goto L165
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v408)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+24)) = uint16(v6)
	F_XLogBeginInsert(m)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v492 != 0 {
		goto L165
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	F_XLogRegisterData(m, v21+int32(24), int32(14))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	if v356 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v507 = int32(14)
	goto L176
L175:
	;
	v507 = int32(8)
	goto L176
L176:
	;
	F_XLogRegisterBuffer(m, int32(0), v68, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	F_XLogRegisterBufData(m, int32(0), l8, l9)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_XLogRegisterBuffer(m, int32(1), v348, int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	F_XLogRegisterBuffer(m, int32(2), l4, int32(8))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	if v356 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v524 = int32(160)
	goto L183
L182:
	;
	v524 = int32(32)
	goto L183
L183:
	;
	v525 = F_XLogInsert(m, int32(17), v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	v527 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = base.I64_rotr(v525, v527)
	v530 = base.I32_wrap_i64(v525)
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v530
	v534 = base.I32_wrap_i64(int64(base.Ui64(v525) >> (uint(v527) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v534
	if v348 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = v534
	goto L165
L186:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v539+(v348^int32(-1))<<(uint(int32(2))%32))))
	v553 = v545
	goto L185
L187:
	;
	goto L188
L188:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v553 = v547 + v348<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	F_LockBuffer(m, l4, int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	if v356 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_RecordPageWithFreeSpace(m, l0, v69, v436)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L6
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v586 = int32(1)
	goto L8
L195:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(8152)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v604 + int32(4)
	F_errmsg(m, int32(693083), v21)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(493916), int32(76), int32(355968))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errmsg_internal(m, int32(384860), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(493916), int32(180), int32(355968))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errmsg_internal(m, int32(407637), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L6
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(493916), int32(256), int32(355968))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L6
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	if base.Ui32(l0) <= base.Ui32(int32(175)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_consts[56])))
		v12 = v11
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_brin_inclusion_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v99 int32
	_ = v99
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v16 != 0 {
		v24 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
		if v18 == int32(0) {
			v24 = v15
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(1)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v24 = v23
		}
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 == int32(0) {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
		if v29 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1)
			return int32(0)
		} else {
			v34 = base.I32_extend16_s(v10)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32)+v11)+16))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+113)))
			if v40 != 0 {
				v104 = F_inclusion_get_procinfo(m, v11, v34&int32(65535))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v110 = F_FunctionCall2Coll(m, v104, v14, v107, v109)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v112 = int32(4)
						v116 = v34<<(uint(v112)%32) + v12 + v112
						v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
						if v117 != 0 {
							v130 = v110
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
							return int32(0)
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
							if v110 == v119 {
								v130 = v110
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
								return int32(0)
							} else {
								F_pfree(m, v119)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
									if v110 != v124 {
										v130 = v110
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
										return int32(0)
									} else {
										v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
										v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
										v128 = F_datumCopy(m, v110, v126, v127)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											v130 = v128
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v42 = v39 + int32(28)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
				if v43 == int32(0) {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+216))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+204))
					v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+6)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50*(v34-int32(1))<<(uint(int32(2))%32)+int32(48)-int32(4))))
					if v62 == int32(0) {
						v99 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v39)+113)) = uint8(v99)
						v104 = F_inclusion_get_procinfo(m, v11, v34&int32(65535))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							v110 = F_FunctionCall2Coll(m, v104, v14, v107, v109)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = int32(4)
								v116 = v34<<(uint(v112)%32) + v12 + v112
								v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
								if v117 != 0 {
									v130 = v110
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
									return int32(0)
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
									if v110 == v119 {
										v130 = v110
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
										return int32(0)
									} else {
										F_pfree(m, v119)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
											if v110 != v124 {
												v130 = v110
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
												return int32(0)
											} else {
												v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
												v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
												v128 = F_datumCopy(m, v110, v126, v127)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = v128
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v67 = F_index_getprocinfo(m, v65, v34, int32(12))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v74 = v39 + int32(44)
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v67)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
							*(*int64)(unsafe.Add(mBase, uint32(v42))) = v77
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v79
							v81 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = v81
							*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
							if v42 == int32(0) {
								v104 = F_inclusion_get_procinfo(m, v11, v34&int32(65535))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
									v110 = F_FunctionCall2Coll(m, v104, v14, v107, v109)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v112 = int32(4)
										v116 = v34<<(uint(v112)%32) + v12 + v112
										v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
										if v117 != 0 {
											v130 = v110
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
											return int32(0)
										} else {
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
											if v110 == v119 {
												v130 = v110
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
												return int32(0)
											} else {
												F_pfree(m, v119)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
													if v110 != v124 {
														v130 = v110
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
														return int32(0)
													} else {
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
														v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
														v128 = F_datumCopy(m, v110, v126, v127)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															v130 = v128
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
															return int32(0)
														}
													}
												}
											}
										}
									}
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
								v92 = F_FunctionCall2Coll(m, v42, v14, v89, v91)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v92 != 0 {
										v104 = F_inclusion_get_procinfo(m, v11, v34&int32(65535))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
											v110 = F_FunctionCall2Coll(m, v104, v14, v107, v109)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												v112 = int32(4)
												v116 = v34<<(uint(v112)%32) + v12 + v112
												v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
												if v117 != 0 {
													v130 = v110
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
													if v110 == v119 {
														v130 = v110
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
														return int32(0)
													} else {
														F_pfree(m, v119)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
															if v110 != v124 {
																v130 = v110
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
																return int32(0)
															} else {
																v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
																v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
																v128 = F_datumCopy(m, v110, v126, v127)
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	v130 = v128
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
																	return int32(0)
																}
															}
														}
													}
												}
											}
										}
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(1)
										return int32(0)
									}
								}
							}
						}
					}
				} else {
					if v42 == int32(0) {
						v104 = F_inclusion_get_procinfo(m, v11, v34&int32(65535))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							v110 = F_FunctionCall2Coll(m, v104, v14, v107, v109)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = int32(4)
								v116 = v34<<(uint(v112)%32) + v12 + v112
								v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
								if v117 != 0 {
									v130 = v110
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
									return int32(0)
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
									if v110 == v119 {
										v130 = v110
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
										return int32(0)
									} else {
										F_pfree(m, v119)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
											if v110 != v124 {
												v130 = v110
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
												return int32(0)
											} else {
												v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
												v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
												v128 = F_datumCopy(m, v110, v126, v127)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = v128
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
						v92 = F_FunctionCall2Coll(m, v42, v14, v89, v91)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							if v92 != 0 {
								v104 = F_inclusion_get_procinfo(m, v11, v34&int32(65535))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
									v110 = F_FunctionCall2Coll(m, v104, v14, v107, v109)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v112 = int32(4)
										v116 = v34<<(uint(v112)%32) + v12 + v112
										v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
										if v117 != 0 {
											v130 = v110
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
											return int32(0)
										} else {
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
											if v110 == v119 {
												v130 = v110
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
												return int32(0)
											} else {
												F_pfree(m, v119)
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
													if v110 != v124 {
														v130 = v110
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
														return int32(0)
													} else {
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
														v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116)+4)))
														v128 = F_datumCopy(m, v110, v126, v127)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															v130 = v128
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v132))) = v130
															return int32(0)
														}
													}
												}
											}
										}
									}
								}
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(1)
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_brin_minmax_add_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15))))
	v21 = v10 + v11<<(uint(int32(4))%32) + v16*int32(100) - int32(80)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
	if v23 == int32(1) {
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
		v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
		v28 = F_datumCopy(m, v22, v26, v27)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v28
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
			v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
			v36 = F_datumCopy(m, v22, v34, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v36
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v40)
				return int32(1)
			}
		}
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v46 = v16 & int32(65535)
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
		v49 = F_minmax_get_strategy_procinfo(m, v9, v46, v47, int32(1))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
			v53 = F_FunctionCall2Coll(m, v49, v44, v22, v52)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				if v53 != 0 {
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
					if v56 != 0 {
						v62 = int32(1)
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
						v66 = F_datumCopy(m, v22, v62&int32(1), v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v68))) = v66
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
							v73 = F_minmax_get_strategy_procinfo(m, v9, v46, v71, int32(5))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
								v77 = F_FunctionCall2Coll(m, v73, v44, v22, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									if v77 == int32(0) {
										return base.B2i32(v53 != int32(0))
									} else {
										v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
										if v85 != 0 {
											v91 = int32(1)
											v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
											v95 = F_datumCopy(m, v22, v91&int32(1), v94)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v95
												return int32(1)
											}
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
											F_pfree(m, v87)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
												v91 = v90
												v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
												v95 = F_datumCopy(m, v22, v91&int32(1), v94)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v95
													return int32(1)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						F_pfree(m, v58)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
							v62 = v61
							v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
							v66 = F_datumCopy(m, v22, v62&int32(1), v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v68))) = v66
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
								v73 = F_minmax_get_strategy_procinfo(m, v9, v46, v71, int32(5))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
									v77 = F_FunctionCall2Coll(m, v73, v44, v22, v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										if v77 == int32(0) {
											return base.B2i32(v53 != int32(0))
										} else {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
											if v85 != 0 {
												v91 = int32(1)
												v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
												v95 = F_datumCopy(m, v22, v91&int32(1), v94)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v95
													return int32(1)
												}
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
												F_pfree(m, v87)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
													v91 = v90
													v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
													v95 = F_datumCopy(m, v22, v91&int32(1), v94)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v95
														return int32(1)
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
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
					v73 = F_minmax_get_strategy_procinfo(m, v9, v46, v71, int32(5))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
						v77 = F_FunctionCall2Coll(m, v73, v44, v22, v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							if v77 == int32(0) {
								return base.B2i32(v53 != int32(0))
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
								if v85 != 0 {
									v91 = int32(1)
									v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
									v95 = F_datumCopy(m, v22, v91&int32(1), v94)
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v95
										return int32(1)
									}
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
									F_pfree(m, v87)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
										v91 = v90
										v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
										v95 = F_datumCopy(m, v22, v91&int32(1), v94)
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v95
											return int32(1)
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
func F_brin_minmax_multi_distance_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i32_s(v2), base.F64_convert_i32_s(v4)))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brin_minmax_multi_distance_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v9 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i64_s(v3), base.F64_convert_i64_s(v6)))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_brin_minmax_multi_distance_tid(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v5 = int32(16)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v9 = int32(291)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v26 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i32_u((v4<<(uint(v5)%32)|v7)*v9+v11), base.F64_convert_i32_u((v15<<(uint(v5)%32)|v18)*v9+v22)))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		return v26
	}
}
func F_brin_minmax_multi_distance_uuid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+15)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
	v9 = float64(0.00390625)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+14)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+13)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+13)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+12)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+11)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+11)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+10)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+10)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+9)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+9)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v116 = F_Float8GetDatum(m, base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v4-v6), v9), base.F64_convert_i32_s(v11-v12)), v9), base.F64_convert_i32_s(v18-v19)), v9), base.F64_convert_i32_s(v25-v26)), v9), base.F64_convert_i32_s(v32-v33)), v9), base.F64_convert_i32_s(v39-v40)), v9), base.F64_convert_i32_s(v46-v47)), v9), base.F64_convert_i32_s(v53-v54)), v9), base.F64_convert_i32_s(v60-v61)), v9), base.F64_convert_i32_s(v67-v68)), v9), base.F64_convert_i32_s(v74-v75)), v9), base.F64_convert_i32_s(v81-v82)), v9), base.F64_convert_i32_s(v88-v89)), v9), base.F64_convert_i32_s(v95-v96)), v9), base.F64_convert_i32_s(v102-v103)), v9), base.F64_convert_i32_s(v109-v110)), v9))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		return int32(0)
	} else {
		return v116
	}
}
func F_brin_minmax_multi_union(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v112 int32
	_ = v112
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 float64
	_ = v682
	var v685 int32
	_ = v685
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v924 int32
	_ = v924
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24))))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v42 = F_brin_range_deserialize(m, v41, v33)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v45 = F_brin_range_deserialize(m, v44, v39)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v58 = F_AllocSetContextCreateInternal(m, v53, int32(60977), v51, int32(8192), int32(8388608))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v60 = int32(4515488)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v58
	v66 = v47 + (v48 + (v49 + v50))
	v69 = F_palloc0(m, v66*int32(12))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if int32(0) < v71 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v76 = v51
	goto L11
L9:
	;
	v111 = v51
	v112 = v71
	goto L10
L10:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	if int32(0) < v132 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v97 = v69 + v76*int32(12)
	v100 = v42 + int32(36) + v76<<(uint(int32(3))%32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+8)) = uint8(v104)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v103
	v108 = v76 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v108 < v109 {
		v76 = v108
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v111 = v108
	v112 = v109
	goto L10
L13:
	;
	goto L12
L14:
	;
	v136 = v42 + int32(36)
	v138 = v111
	v139 = int32(0)
	goto L17
L15:
	;
	v189 = v132
	v206 = v112
	goto L16
L16:
	;
	v207 = int32(12)
	v212 = v69 + v206*v207 + v189*v207
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v213 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v159 = v69 + v138*int32(12)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v161 = int32(1)
	v164 = int32(2)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v160<<(uint(v161)%32)+v139)<<(uint(v164)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v169<<(uint(v161)%32)+v139)<<(uint(v164)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+8)) = uint8(v161)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v176
	v183 = v139 + v161
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	if v183 < v184 {
		v138 = v138 + v161
		v139 = v183
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v189 = v184
	v206 = v186
	goto L16
L19:
	;
	goto L18
L20:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	if int32(0) < v275 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v255 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v220 = int32(0)
	goto L24
L24:
	;
	v241 = v212 + v220*int32(12)
	v244 = v45 + int32(36) + v220<<(uint(int32(3))%32)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v248 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+8)) = uint8(v248)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = v247
	v252 = v220 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v252 < v253 {
		v220 = v252
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v255 = v252
	goto L20
L26:
	;
	goto L25
L27:
	;
	v279 = v45 + int32(36)
	v281 = v255
	v282 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v348 = int32(1)
	v350 = v25 & int32(65535)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(4))%32)+v25*int32(100)-int32(12))))
	v358 = F_minmax_multi_get_strategy_procinfo(m, v28, v350, v356, v348)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v302 = v212 + v281*int32(12)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v304 = int32(1)
	v307 = int32(2)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v279+(v303<<(uint(v304)%32)+v282)<<(uint(v307)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v279+(v312<<(uint(v304)%32)+v282)<<(uint(v307)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v302)+8)) = uint8(v304)
	*(*int32)(unsafe.Add(mBase, uint32(v302)+4)) = v319
	v326 = v282 + v304
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	if v326 < v327 {
		v281 = v281 + v304
		v282 = v326
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v27
	F_qsort_arg(m, v69, v66, int32(12), int32(22), v22+int32(8))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if int32(1) < v66 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v735 = F_reduce_expanded_ranges(m, v69, v723, v722, v734, v358, v27)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L122
	}
L36:
	;
	v371 = v348
	v374 = int32(1)
	goto L39
L37:
	;
	goto L38
L38:
	;
	v710 = F_minmax_multi_get_procinfo(m, v28, v350)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L121
	}
L39:
	;
	v390 = int32(12)
	v392 = v69 + v371*v390
	v397 = F_compare_expanded_ranges(m, v392-v390, v392, v22+int32(8))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v415 = int32(1)
	v417 = v411 - v415
	if int32(0) < v417 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	if v397 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v371 != v374 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v411 = v374
	goto L44
L44:
	;
	v413 = v371 + int32(1)
	if v413 != v66 {
		v371 = v413
		v374 = v411
		goto L39
	} else {
		goto L48
	}
L45:
	;
	v402 = v69 + v374*int32(12)
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v392)))
	*(*int64)(unsafe.Add(mBase, uint32(v402))) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v402)+8)) = v405
	goto L47
L46:
	;
	goto L47
L47:
	;
	v411 = v374 + int32(1)
	goto L44
L48:
	;
	goto L40
L49:
	;
	v421 = int32(0)
	v424 = v411
	v426 = v417
	goto L52
L50:
	;
	v620 = v411
	goto L51
L51:
	;
	v639 = F_minmax_multi_get_procinfo(m, v28, v25&int32(65535))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L110
	}
L52:
	;
	v440 = int32(12)
	v442 = v69 + v421*v440
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v445 = v442 + v440
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	v447 = F_FunctionCall2Coll(m, v358, v27, v443, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v620 = v613
	goto L51
L54:
	;
	v615 = v613 - int32(1)
	if v611 < v615 {
		v421 = v611
		v424 = v613
		v426 = v615
		goto L52
	} else {
		goto L109
	}
L55:
	;
	if v447 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v611 = v421 + int32(1)
	v613 = v424
	goto L54
L57:
	;
	goto L58
L58:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	v453 = F_FunctionCall2Coll(m, v358, v27, v451, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v453 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+4)) = v455
	goto L62
L61:
	;
	goto L62
L62:
	;
	v457 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+8)) = uint8(v457)
	v460 = v421 + int32(2)
	v461 = int32(12)
	v463 = v69 + v460*v461
	v466 = (v424 - v460) * v461
	if v445 == v463 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v611 = v421
	v613 = v426
	goto L54
L64:
	;
	goto L63
L65:
	;
	v470 = v445 + v466
	if base.Ui32(v463-v470) <= base.Ui32(int32(0)-v466<<(uint(int32(1))%32)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v477 = F___memcpy(m, v445, v463, v466)
	mBase = m.M
	goto L63
L67:
	;
	goto L68
L68:
	;
	v480 = (v445 ^ v463) & int32(3)
	if base.Ui32(v445) < base.Ui32(v463) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v582 == int32(0) {
		goto L64
	} else {
		goto L105
	}
L70:
	;
	if base.Ui32(v560) <= base.Ui32(int32(3)) {
		v581 = v559
		v582 = v560
		v583 = v561
		goto L69
	} else {
		goto L101
	}
L71:
	;
	if v480 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v480 != 0 {
		v542 = v466
		goto L84
	} else {
		goto L85
	}
L74:
	;
	v581 = v463
	v582 = v466
	v583 = v445
	goto L69
L75:
	;
	goto L76
L76:
	;
	if v445&int32(3) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v559 = v463
	v560 = v466
	v561 = v445
	goto L70
L78:
	;
	goto L79
L79:
	;
	v487 = v463
	v488 = v466
	v489 = v445
	goto L80
L80:
	;
	if v488 == int32(0) {
		goto L64
	} else {
		goto L82
	}
L81:
	;
	v559 = v496
	v560 = v498
	v561 = v500
	goto L70
L82:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	*(*uint8)(unsafe.Add(mBase, uint32(v489))) = uint8(v493)
	v495 = int32(1)
	v496 = v487 + v495
	v498 = v488 - v495
	v500 = v489 + v495
	if v500&int32(3) != 0 {
		v487 = v496
		v488 = v498
		v489 = v500
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if v542 == int32(0) {
		goto L64
	} else {
		goto L97
	}
L85:
	;
	if v470&int32(3) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v507 = v466
	goto L89
L87:
	;
	v522 = v466
	goto L88
L88:
	;
	if base.Ui32(v522) <= base.Ui32(int32(3)) {
		v542 = v522
		goto L84
	} else {
		goto L93
	}
L89:
	;
	if v507 == int32(0) {
		goto L64
	} else {
		goto L91
	}
L90:
	;
	v522 = v513
	goto L88
L91:
	;
	v513 = v507 - int32(1)
	v514 = v445 + v513
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v513))))
	*(*uint8)(unsafe.Add(mBase, uint32(v514))) = uint8(v516)
	if v514&int32(3) != 0 {
		v507 = v513
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v529 = v522
	goto L94
L94:
	;
	v533 = v529 - int32(4)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v463+v533)))
	*(*int32)(unsafe.Add(mBase, uint32(v445+v533))) = v536
	if base.Ui32(int32(3)) < base.Ui32(v533) {
		v529 = v533
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v542 = v533
	goto L84
L96:
	;
	goto L95
L97:
	;
	v549 = v542
	goto L98
L98:
	;
	v553 = v549 - int32(1)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v553))))
	*(*uint8)(unsafe.Add(mBase, uint32(v445+v553))) = uint8(v556)
	if v553 != 0 {
		v549 = v553
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L64
L100:
	;
	goto L99
L101:
	;
	v566 = v559
	v567 = v560
	v568 = v561
	goto L102
L102:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v570
	v572 = int32(4)
	v573 = v566 + v572
	v575 = v568 + v572
	v577 = v567 - v572
	if base.Ui32(int32(3)) < base.Ui32(v577) {
		v566 = v573
		v567 = v577
		v568 = v575
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v581 = v573
	v582 = v577
	v583 = v575
	goto L69
L104:
	;
	goto L103
L105:
	;
	v588 = v581
	v589 = v582
	v590 = v583
	goto L106
L106:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
	*(*uint8)(unsafe.Add(mBase, uint32(v590))) = uint8(v592)
	v594 = int32(1)
	v599 = v589 - v594
	if v599 != 0 {
		v588 = v588 + v594
		v589 = v599
		v590 = v590 + v594
		goto L106
	} else {
		goto L108
	}
L107:
	;
	goto L64
L108:
	;
	goto L107
L109:
	;
	goto L53
L110:
	;
	if v620 == int32(1) {
		v722 = int32(0)
		v723 = v415
		goto L35
	} else {
		goto L111
	}
L111:
	;
	v644 = v620 - int32(1)
	v647 = F_palloc0(m, v644<<(uint(int32(4))%32))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if int32(0) < v644 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v652 = int32(0)
	goto L116
L114:
	;
	goto L115
L115:
	;
	F_pg_qsort(m, v647, v644, int32(16), int32(20))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L120
	}
L116:
	;
	v673 = v69 + v652*int32(12)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v673)+12))
	v676 = F_FunctionCall2Coll(m, v639, v27, v674, v675)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	goto L115
L118:
	;
	v680 = v647 + v652<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v680))) = v652
	v682 = *(*float64)(unsafe.Add(mBase, uint32(v676)))
	*(*float64)(unsafe.Add(mBase, uint32(v680)+8)) = v682
	v685 = v652 + int32(1)
	if v685 != v644 {
		v652 = v685
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v722 = v647
	v723 = v620
	goto L35
L121:
	;
	v722 = int32(0)
	v723 = int32(1)
	goto L35
L122:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v737
	if v737 < v735 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v924
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v61
	F_MemoryContextDelete(m, v58)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L148
	}
L124:
	;
	v742 = v42 + int32(36)
	v743 = int32(0)
	v745 = v743
	v748 = v743
	goto L127
L125:
	;
	v904 = int32(0)
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v904
	v924 = v904
	goto L123
L127:
	;
	v766 = v69 + v748*int32(12)
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+8)))
	if v767 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v788 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v788
	v791 = int32(1)
	if v735 == v791 {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v770 = int32(2)
	v772 = v742 + v745<<(uint(v770)%32)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v773
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v772)+4)) = v775
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v777 + int32(1)
	v783 = v745 + v770
	goto L131
L130:
	;
	v783 = v745
	goto L131
L131:
	;
	v786 = v748 + int32(1)
	if v786 != v735 {
		v745 = v783
		v748 = v786
		goto L127
	} else {
		goto L132
	}
L132:
	;
	goto L128
L133:
	;
	if v735&v791 == int32(0) {
		v924 = v869
		goto L123
	} else {
		goto L146
	}
L134:
	;
	v864 = v783
	v867 = int32(0)
	v869 = v788
	goto L133
L135:
	;
	goto L136
L136:
	;
	v798 = int32(0)
	v800 = v783
	v801 = v798
	v803 = v798
	v805 = v788
	goto L137
L137:
	;
	v821 = v69 + v803*int32(12)
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821)+8)))
	if v822 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v864 = v857
	v867 = v860
	v869 = v858
	goto L133
L139:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	*(*int32)(unsafe.Add(mBase, uint32(v742+v800<<(uint(int32(2))%32)))) = v828
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	v831 = int32(1)
	v832 = v830 + v831
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v832
	v836 = v800 + v831
	v837 = v832
	goto L141
L140:
	;
	v836 = v800
	v837 = v805
	goto L141
L141:
	;
	v838 = int32(1)
	v842 = v69 + (v803|v838)*int32(12)
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+8)))
	if v843 == v838 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	*(*int32)(unsafe.Add(mBase, uint32(v742+v836<<(uint(int32(2))%32)))) = v849
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	v852 = int32(1)
	v853 = v851 + v852
	*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v853
	v857 = v836 + v852
	v858 = v853
	goto L144
L143:
	;
	v857 = v836
	v858 = v837
	goto L144
L144:
	;
	v859 = int32(2)
	v860 = v803 + v859
	v862 = v801 + v859
	if v862 != v735&int32(2147483646) {
		v800 = v857
		v801 = v862
		v803 = v860
		v805 = v858
		goto L137
	} else {
		goto L145
	}
L145:
	;
	goto L138
L146:
	;
	v887 = v69 + v867*int32(12)
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+8)))
	if v888 != int32(1) {
		v924 = v869
		goto L123
	} else {
		goto L147
	}
L147:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	*(*int32)(unsafe.Add(mBase, uint32(v742+v864<<(uint(int32(2))%32)))) = v894
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	v904 = v896 + int32(1)
	goto L126
L148:
	;
	F_pfree(m, v33)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v945 = F_brin_range_serialize(m, v42)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v947))) = v945
	m.G0 = v22 + int32(16)
	return int32(0)
}
