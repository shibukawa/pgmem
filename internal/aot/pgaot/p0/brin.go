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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapInitialize[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v7^int32(-1))<<(uint(int32(2))%32))))
				v31 = v23
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapInitialize[1]))
				v31 = v25 + v7<<(uint(int32(13))%32) + int32(-8192)
			}
			v33 = F_palloc(m, int32(20))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = l0
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
				v39 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v38
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
				F_LockBuffer(m, v7, v39)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
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
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v250 int64
	_ = v250
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v269 int64
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
		v306 = int32(1)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v18 + int32(16)
	return v306
L4:
	;
	v40 = int32(0)
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23+v40<<(uint(int32(2))%32))))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)))
	if v55 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v306 = v301
	goto L3
L7:
	;
	v301 = int32(1)
	v303 = v40 + v301
	if v303 != v20 {
		v40 = v303
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
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L30
	}
L11:
	;
	v62 = F_FunctionCall1Coll(m, v60, v22, v58)
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
	v148 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)))
	v149 = base.I64_rem_u_s(base.I64_extend_i32_u(v138)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v138^v130-base.I32_rotl(v138, int32(24))), v148)
	goto L21
L14:
	;
	v115 = int32(14)
	v117 = v94 - v102 ^ base.I32_rotl(v102, int32(4)) ^ v108 - base.I32_rotl(v108, v115)
	v122 = v117 ^ (v62 + v107) - base.I32_rotl(v117, int32(11))
	v126 = v108 ^ v122 - base.I32_rotl(v122, int32(25))
	v130 = v126 ^ v117 - base.I32_rotl(v126, int32(16))
	v134 = v130 ^ v122 - base.I32_rotl(v130, int32(4))
	v138 = v134 ^ v126 - base.I32_rotl(v134, v115)
	goto L13
L16:
	;
	goto L17
L17:
	;
	v74 = base.I32_wrap_i64(int64(1910056111))
	v76 = v74 + int32(1021750440)
	v81 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
	v87 = v74 - v81 - int32(1636608428) ^ base.I32_rotl(v81, int32(6))
	v91 = v76 - v87 ^ base.I32_rotl(v87, int32(8))
	v92 = v81 + v76
	v93 = v87 + v92
	v94 = v91 + v93
	v98 = v92 - v91 ^ base.I32_rotl(v91, int32(16))
	v102 = v93 - v98 ^ base.I32_rotl(v98, int32(19))
	v107 = v98 + v94
	v108 = v102 + v107
	goto L14
L18:
	;
	v234 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)))
	v235 = base.I64_rem_u_s(base.I64_extend_i32_u(v224)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v224^v216-base.I32_rotl(v224, int32(24))), v234)
	v236 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	if v236 == int64(0) {
		goto L7
	} else {
		goto L23
	}
L19:
	;
	v201 = int32(14)
	v203 = v180 - v188 ^ base.I32_rotl(v188, int32(4)) ^ v194 - base.I32_rotl(v194, v201)
	v208 = v203 ^ (v62 + v193) - base.I32_rotl(v203, int32(11))
	v212 = v194 ^ v208 - base.I32_rotl(v208, int32(25))
	v216 = v212 ^ v203 - base.I32_rotl(v212, int32(16))
	v220 = v216 ^ v208 - base.I32_rotl(v216, int32(4))
	v224 = v220 ^ v212 - base.I32_rotl(v220, v201)
	goto L18
L21:
	;
	goto L22
L22:
	;
	v160 = base.I32_wrap_i64(int64(3125326612))
	v162 = v160 + int32(1021750440)
	v167 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
	v173 = v160 - v167 - int32(1636608428) ^ base.I32_rotl(v167, int32(6))
	v177 = v162 - v173 ^ base.I32_rotl(v173, int32(8))
	v178 = v167 + v162
	v179 = v173 + v178
	v180 = v177 + v179
	v184 = v178 - v177 ^ base.I32_rotl(v177, int32(16))
	v188 = v179 - v184 ^ base.I32_rotl(v184, int32(19))
	v193 = v184 + v180
	v194 = v188 + v193
	goto L19
L23:
	;
	v250 = int64(0)
	goto L24
L24:
	;
	v257 = base.I64_rem_u_s(v250*v235+v149, v234)
	v258 = base.I32_wrap_i64(v257)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(16)+int32(base.Ui32(v258)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v262)>>(uint(v258&int32(7))%32))&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v306 = int32(0)
	goto L3
L26:
	;
	v269 = v250 + int64(1)
	if v236 != v269 {
		v250 = v269
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
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v276
	F_errmsg_internal(m, int32(_a_F_brin_bloom_consistent_0), v18)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_brin_bloom_consistent_1), int32(644), int32(_a_F_brin_bloom_consistent_2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_brin_bloom_summary_in_0), int32(784), int32(_a_F_brin_bloom_summary_in_1), int32(_a_F_brin_bloom_summary_in_2), int32(_a_F_brin_bloom_summary_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brin_build_desc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_brin_build_desc[0]))
	v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F_brin_build_desc_0), v2, int32(1024), int32(_a_F_brin_build_desc_1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(_a_F_brin_build_desc_2)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_brin_build_desc[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_build_desc[0])) = v16
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v28 = F_palloc(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if int32(0) < v30 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = v30
	v35 = v2
	v40 = v2
	goto L7
L5:
	;
	v67 = v30
	v73 = v2
	goto L6
L6:
	;
	v79 = F_palloc(m, v67<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v45 = int32(1)
	v46 = v35 + v45
	v49 = F_index_getprocinfo(m, l0, base.I32_extend16_s(v46), v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v67 = v64
	v73 = v63
	goto L6
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v24+v34<<(uint(int32(4))%32)+v35*int32(100))+88))
	v59 = F_FunctionCall1Coll(m, v49, int32(0), v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+v35<<(uint(int32(2))%32)))) = v59
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	v63 = v40 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v46 < v64 {
		v34 = v64
		v35 = v46
		v40 = v63
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v73
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v16
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v82 < v87 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v95 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_pfree(m, v28)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v103 = v95 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103+v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v79+int32(20)+v103))) = v106
	v109 = v95 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v109 < v110 {
		v95 = v109
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brin_build_desc[0])) = v21
	return v79
}
func F_brin_can_do_samepage_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	if base.Ui32(l1) < base.Ui32(l2) {
		if l0 < int32(0) {
			v8 = *(*int32)(unsafe.Add(mBase, _c_F_brin_can_do_samepage_update[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+(l0^int32(-1))<<(uint(int32(2))%32))))
			v22 = v14
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_brin_can_do_samepage_update[1]))
			v22 = v16 + l0<<(uint(int32(13))%32) + int32(-8192)
		}
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)))
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+12)))
		v25 = v23 - v24
		v26 = int32(0)
		if v26 < v25 {
			v29 = v25
		} else {
			v29 = v26
		}
		v33 = base.B2i32(base.Ui32(l2-l1) <= base.Ui32(v29))
	} else {
		v33 = int32(1)
	}
	return v33
}
func F_brin_desummarize_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
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
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_brin_desummarize_range[0])))
	if v23 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L12
	} else {
		goto L142
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L12
	} else {
		goto L138
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L12
	} else {
		goto L134
	}
L4:
	;
	if v33 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[1]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+316))
	v31 = base.B2i32(v29 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_brin_desummarize_range[0])) = uint8(v31)
	v33 = v31
	goto L7
L6:
	;
	v33 = int32(0)
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
	v470 = m.ExcPending
	if v470 != 0 {
		goto L12
	} else {
		goto L129
	}
L11:
	;
	v39 = F_IndexGetRelation(m, v18, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = F_table_open(m, v39, int32(4))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v47 = int32(0)
	goto L16
L16:
	;
	v49 = F_index_open(m, v18, int32(4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v47 = v44
	goto L16
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
	if v52 != int32(105) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
	if v55 != int32(3580) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[2]))
	v61 = F_object_ownercheck(m, int32(1259), v18, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v61 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	F_aclcheck_error(m, int32(2), int32(20), v67+int32(4))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v47 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v75 = F_IndexGetRelation(m, v18, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v75 != v39 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+192))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+18)))
	if v79 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_relation_close(m, v49, int32(4))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L12
	} else {
		goto L127
	}
L30:
	;
	v82 = base.I32_wrap_i64(v20)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v422 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L12
	} else {
		goto L122
	}
L33:
	;
	v96 = m.G0
	v98 = v96 - int32(16)
	m.G0 = v98
	v101 = F_ReadBuffer(m, v49, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L36
	}
L34:
	;
	goto L29
L35:
	;
	if v378 == int32(0) {
		goto L33
	} else {
		goto L121
	}
L36:
	;
	F_LockBuffer(m, v101, int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v101 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v125 = F_palloc(m, int32(20))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L42
	}
L39:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109+(v101^int32(-1))<<(uint(int32(2))%32))))
	v123 = v115
	goto L38
L40:
	;
	goto L41
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v123 = v117 + v101<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v49
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v131 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v130
	F_LockBuffer(m, v101, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v140 = base.I32_div_u_s(v82, v139)
	v142 = base.I32_div_u_s(v140, int32(1360))
	if base.Ui32(v138) <= base.Ui32(v142) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L12
	} else {
		goto L117
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L12
	} else {
		goto L113
	}
L46:
	;
	m.G0 = v98 + int32(16)
	goto L35
L47:
	;
	F_pfree(m, v125)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L12
	} else {
		goto L112
	}
L48:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	F_ReleaseBuffer(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v152 = F_brinLockRevmapPageForUpdate(m, v125, v82)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L55
	}
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v147 == int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	F_ReleaseBuffer(m, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v173 = base.I32_div_u_s(v82, v172)
	v175 = base.I32_rem_u_s(v173, int32(1360))
	v178 = v171 + v175*int32(6)
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+28)))
	if v179 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	if v152 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157+(v152^int32(-1))<<(uint(int32(2))%32))))
	v171 = v163
	goto L54
L57:
	;
	goto L58
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v171 = v165 + v152<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L59:
	;
	F_LockBuffer(m, v152, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v194 = v178 + int32(24)
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194))))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+2)))
	v200 = F_ReadBuffer(m, v49, v195<<(uint(int32(16))%32)|v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L66
	}
L62:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	F_ReleaseBuffer(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v188 == int32(0) {
		goto L47
	} else {
		goto L64
	}
L64:
	;
	F_ReleaseBuffer(m, v188)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	goto L47
L66:
	;
	F_LockBuffer(m, v200, int32(2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	if v200 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+16)))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223+v222)+6)))
	if v225 != int32(_a_F_brin_desummarize_range_0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v208+(v200^int32(-1))<<(uint(int32(2))%32))))
	v222 = v214
	goto L68
L70:
	;
	goto L71
L71:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v222 = v216 + v200<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	v228 = int32(0)
	F_LockBuffer(m, v152, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+4)))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v244) {
		goto L83
	} else {
		goto L84
	}
L75:
	;
	F_LockBuffer(m, v200, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	F_ReleaseBuffer(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v238 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_ReleaseBuffer(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_pfree(m, v125)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L12
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v378 = v228
	goto L46
L83:
	;
	v252 = int32(base.Ui32(v244+int32(_a_F_brin_desummarize_range_1)) >> (uint(int32(2)) % 32))
	goto L85
L84:
	;
	v252 = int32(0)
	goto L85
L85:
	;
	if base.Ui32(v252&int32(_a_F_brin_desummarize_range_2)) < base.Ui32(v243) {
		goto L45
	} else {
		goto L86
	}
L86:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222+v243<<(uint(int32(2))%32))+21)))
	if v259&int32(384) == int32(0) {
		goto L44
	} else {
		goto L87
	}
L87:
	;
	v264 = int32(_a_F_brin_desummarize_range_3)
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5])) = v266 + int32(1)
	if v152 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v289 = base.I32_div_u_s(v82, v288)
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
		goto L92
	}
L89:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+(v152^int32(-1))<<(uint(int32(2))%32))))
	v287 = v279
	goto L88
L90:
	;
	goto L91
L91:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v287 = v281 + v152<<(uint(int32(13))%32) + int32(-8192)
	goto L88
L92:
	;
	F_MarkBufferDirty(m, v200)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	F_MarkBufferDirty(m, v152)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L94
	}
L94:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+118)))
	if v306 != int32(112) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v347 = int32(_a_F_brin_desummarize_range_3)
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5])) = v349 - int32(1)
	F_UnlockReleaseBuffer(m, v200)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L12
	} else {
		goto L107
	}
L96:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[6]))
	if v310 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
	if v313 != 0 {
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+12)) = uint16(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v315
	F_XLogBeginInsert(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L102
	}
L100:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v49)+40))
	if v314 != 0 {
		goto L95
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	F_XLogRegisterData(m, v98+int32(4), int32(10))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	v326 = int32(0)
	F_XLogRegisterBuffer(m, v326, v152, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_XLogRegisterBuffer(m, int32(1), v200, int32(8))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v336 = F_XLogInsert(m, int32(17), int32(80))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v338 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = base.I64_rotr(v336, v338)
	*(*uint32)(unsafe.Add(mBase, uint32(v222)+4)) = uint32(v336)
	v343 = int64(base.Ui64(v336) >> (uint(v338) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v222))) = uint32(v343)
	goto L95
L107:
	;
	F_LockBuffer(m, v152, int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	F_ReleaseBuffer(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v361 == int32(0) {
		goto L47
	} else {
		goto L110
	}
L110:
	;
	F_ReleaseBuffer(m, v361)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L111
	}
L111:
	;
	goto L47
L112:
	;
	v378 = int32(1)
	goto L46
L113:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_brin_desummarize_range_4), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_5), int32(383), int32(_a_F_brin_desummarize_range_6))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_brin_desummarize_range_4), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_5), int32(389), int32(_a_F_brin_desummarize_range_6))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	goto L34
L122:
	;
	if v422 == int32(0) {
		goto L29
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v429 + int32(4)
	F_errmsg(m, int32(_a_F_brin_desummarize_range_7), v14+int32(-32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1569), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	goto L29
L127:
	;
	F_relation_close(m, v47, int32(4))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	m.G0 = v16 - int32(-64)
	return int32(0)
L129:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_brin_desummarize_range_10), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	F_errhint(m, int32(_a_F_brin_desummarize_range_11), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1505), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L12
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v20
	F_errmsg(m, int32(_a_F_brin_desummarize_range_12), v16)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L12
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1511), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L12
	} else {
		goto L139
	}
L139:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v510 + int32(4)
	F_errmsg(m, int32(_a_F_brin_desummarize_range_13), v14+int32(-16))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1537), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v531 + int32(4)
	F_errmsg(m, int32(_a_F_brin_desummarize_range_14), v14+int32(-48))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1553), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
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
	var v12 int32
	_ = v12
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
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int64
	_ = v534
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	v6 = l5
	v12 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	if base.Ui32(l9) < base.Ui32(int32(_a_F_brin_doupdate_0)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L6
	} else {
		goto L202
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L6
	} else {
		goto L199
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
	v609 = m.ExcPending
	if v609 != 0 {
		goto L6
	} else {
		goto L195
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
	return v595
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
	v595 = int32(0)
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
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v33^int32(-1))<<(uint(int32(6))%32))+16))
	v58 = v49
	goto L19
L21:
	;
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[1]))
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
	v135 = v87 + v108&int32(_a_F_brin_doupdate_1)
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
	if v90 != int32(_a_F_brin_doupdate_2) {
		goto L25
	} else {
		goto L30
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73+(l4^int32(-1))<<(uint(int32(2))%32))))
	v87 = v79
	goto L26
L28:
	;
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
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
	v101 = int32(base.Ui32(v93+int32(_a_F_brin_doupdate_3)) >> (uint(int32(2)) % 32))
	goto L33
L32:
	;
	v101 = int32(0)
	goto L33
L33:
	;
	if base.Ui32(v101&int32(_a_F_brin_doupdate_4)) < base.Ui32(v6) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v87+v6<<(uint(int32(2))%32))+20))
	if v108&int32(_a_F_brin_doupdate_5) == int32(_a_F_brin_doupdate_6) {
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
		v595 = v114
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
	v595 = v114
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
	v595 = v114
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
		v595 = v203
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
	v595 = v203
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
	v595 = v203
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
	v255 = int32(_a_F_brin_doupdate_7)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v257 + int32(1)
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
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231+(l4^int32(-1))<<(uint(int32(2))%32))))
	v245 = v237
	goto L83
L85:
	;
	goto L86
L86:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
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
	v299 = int32(_a_F_brin_doupdate_7)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v301 - int32(1)
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
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[5]))
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
		v595 = v308
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
	v595 = v308
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
	v595 = v308
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
	v595 = v326
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
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333+(v68^int32(-1))<<(uint(int32(2))%32))))
	v347 = v339
	goto L120
L122:
	;
	goto L123
L123:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v347 = v341 + v68<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L124:
	;
	v350 = int32(_a_F_brin_doupdate_7)
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	v353 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v352 + v353
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v356 == v353 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v359 = int32(_a_F_brin_doupdate_8)
	v361 = int32(0)
	if v361|(v347&int32(3)|int32(1)) == v361 {
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
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L139
	}
L128:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)))
	v411 = int32(_a_F_brin_doupdate_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v347+v409)+6)) = uint16(v411)
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+10)) = int32(_a_F_brin_doupdate_9)
	v400 = int32(_a_F_brin_doupdate_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+18)) = uint16(v400)
	v406 = int32(_a_F_brin_doupdate_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)) = uint16(v406)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+14)) = uint16(v406)
	goto L128
L130:
	;
	goto L133
L131:
	;
	goto L132
L132:
	;
	goto L138
L133:
	;
	v377 = v347 + v359
	v379 = v347 + int32(4)
	if base.Ui32(v379) < base.Ui32(v377) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v381 = v377
	goto L136
L135:
	;
	v381 = v379
	goto L136
L136:
	;
	v386 = (v347^int32(-1)+v381)&int32(-4) + int32(4)
	if v386 == int32(0) {
		goto L129
	} else {
		goto L137
	}
L137:
	;
	base.MemoryFill(m, v347, int32(0), v386)
	goto L129
L138:
	;
	base.MemoryFill(m, v347, int32(0), v359)
	goto L129
L139:
	;
	v415 = int32(0)
	v417 = F_PageAddItemExtended(m, v347, l8, l9, v415, v415)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	if v417 == int32(0) {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	F_MarkBufferDirty(m, v68)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	if v356 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+16)))
	v426 = v347 + v425
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v426)+6)))
	if v427 != int32(_a_F_brin_doupdate_2) {
		v442 = v12
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v444 = v12
	goto L146
L146:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v417)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v417)
	v447 = int32(16)
	v448 = base.I32_rotr(v69, v447)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v448
	v452 = v21 + v447
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	v456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+2)))
	if v348 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L147:
	;
	v444 = v442
	goto L146
L148:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+4)))
	if v430&int32(1) != 0 {
		v442 = v12
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v433 = int32(4)
	v434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+14)))
	v435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+12)))
	v436 = v434 - v435
	if v436 <= v433 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v442 = v439 - int32(4)
	goto L147
L151:
	;
	v439 = v433
	goto L153
L152:
	;
	v439 = v436
	goto L153
L153:
	;
	goto L150
L154:
	;
	F_MarkBufferDirty(m, v348)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L6
	} else {
		goto L165
	}
L155:
	;
	v475 = base.I32_div_u_s(l3, l1)
	v477 = base.I32_rem_u_s(v475, int32(1360))
	v480 = v474 + v477*int32(6)
	v481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v480)+28)) = uint16(v481)
	if v481 != 0 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460+(v348^int32(-1))<<(uint(int32(2))%32))))
	v474 = v466
	goto L155
L157:
	;
	goto L158
L158:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v474 = v468 + v348<<(uint(int32(13))%32) + int32(-8192)
	goto L155
L159:
	;
	v484 = v456
	goto L161
L160:
	;
	v484 = int32(-1)
	goto L161
L161:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v480)+26)) = uint16(v484)
	if v481 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v487 = v455
	goto L164
L163:
	;
	v487 = int32(-1)
	goto L164
L164:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v480)+24)) = uint16(v487)
	goto L154
L165:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+118)))
	if v492 != int32(112) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v569 = int32(_a_F_brin_doupdate_7)
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v571 - int32(1)
	F_LockBuffer(m, v348, int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L6
	} else {
		goto L187
	}
L167:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[5]))
	if v496 <= int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v499 != 0 {
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v417)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+24)) = uint16(v6)
	F_XLogBeginInsert(m)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L173
	}
L171:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v500 != 0 {
		goto L166
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	F_XLogRegisterData(m, v21+int32(24), int32(14))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	if v356 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v515 = int32(14)
	goto L177
L176:
	;
	v515 = int32(8)
	goto L177
L177:
	;
	F_XLogRegisterBuffer(m, int32(0), v68, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_XLogRegisterBufData(m, int32(0), l8, l9)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	F_XLogRegisterBuffer(m, int32(1), v348, int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_XLogRegisterBuffer(m, int32(2), l4, int32(8))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	v534 = F_XLogInsert(m, int32(17), v356<<(uint(int32(7))%32)|int32(32))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	v536 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = base.I64_rotr(v534, v536)
	v539 = base.I32_wrap_i64(v534)
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v539
	v543 = base.I32_wrap_i64(int64(base.Ui64(v534) >> (uint(v536) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v543
	if v348 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562)+4)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v543
	goto L166
L184:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v548+(v348^int32(-1))<<(uint(int32(2))%32))))
	v562 = v554
	goto L183
L185:
	;
	goto L186
L186:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v562 = v556 + v348<<(uint(int32(13))%32) + int32(-8192)
	goto L183
L187:
	;
	F_LockBuffer(m, l4, int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	if v356 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_RecordPageWithFreeSpace(m, l0, v69, v444)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v595 = int32(1)
	goto L8
L193:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(_a_F_brin_doupdate_12)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v613 + int32(4)
	F_errmsg(m, int32(_a_F_brin_doupdate_13), v21)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_brin_doupdate_14), int32(76), int32(_a_F_brin_doupdate_15))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errmsg_internal(m, int32(_a_F_brin_doupdate_16), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_brin_doupdate_14), int32(180), int32(_a_F_brin_doupdate_15))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errmsg_internal(m, int32(_a_F_brin_doupdate_17), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_brin_doupdate_14), int32(256), int32(_a_F_brin_doupdate_15))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L6
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(175)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_brin_identify[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_brin_inclusion_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v18 != 0 {
		v26 = v17
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		if v20 == int32(0) {
			v26 = v17
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(1)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v26 = v25
		}
	}
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v27 == int32(0) {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		if v31 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(1)
			return int32(0)
		} else {
			v36 = base.I32_extend16_s(v12)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v13+v12<<(uint(int32(2))%32))+16))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+113)))
			if v42 != 0 {
				v102 = F_inclusion_get_procinfo(m, v13, v36&int32(_a_F_brin_inclusion_union_0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
					v108 = F_FunctionCall2Coll(m, v102, v15, v105, v107)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = int32(4)
						v114 = v14 + v36<<(uint(v110)%32) + v110
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
						if v115 != 0 {
							v128 = v108
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
							return int32(0)
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
							if v108 == v117 {
								v128 = v108
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
								return int32(0)
							} else {
								F_pfree(m, v117)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
									if v108 != v122 {
										v128 = v108
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
										return int32(0)
									} else {
										v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
										v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+4)))
										v126 = F_datumCopy(m, v108, v124, v125)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											v128 = v126
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v44 = v41 + int32(28)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
				if v45 != 0 {
					v85 = v30
					v86 = v26
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v89 = F_FunctionCall2Coll(m, v44, v15, v87, v88)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						if v89 != 0 {
							v102 = F_inclusion_get_procinfo(m, v13, v36&int32(_a_F_brin_inclusion_union_0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
								v108 = F_FunctionCall2Coll(m, v102, v15, v105, v107)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									v110 = int32(4)
									v114 = v14 + v36<<(uint(v110)%32) + v110
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
									if v115 != 0 {
										v128 = v108
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
										return int32(0)
									} else {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
										if v108 == v117 {
											v128 = v108
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
											return int32(0)
										} else {
											F_pfree(m, v117)
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
												if v108 != v122 {
													v128 = v108
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
													return int32(0)
												} else {
													v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
													v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+4)))
													v126 = F_datumCopy(m, v108, v124, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = v126
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
														return int32(0)
													}
												}
											}
										}
									}
								}
							}
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = int32(1)
							return int32(0)
						}
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+216))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+204))
					v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+6)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v48+v50*(v36-int32(1))<<(uint(int32(2))%32)+int32(48)-int32(4))))
					if v62 == int32(0) {
						v96 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v41)+113)) = uint8(v96)
						v102 = F_inclusion_get_procinfo(m, v13, v36&int32(_a_F_brin_inclusion_union_0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
							v108 = F_FunctionCall2Coll(m, v102, v15, v105, v107)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								v110 = int32(4)
								v114 = v14 + v36<<(uint(v110)%32) + v110
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
								if v115 != 0 {
									v128 = v108
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
									return int32(0)
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
									if v108 == v117 {
										v128 = v108
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
										return int32(0)
									} else {
										F_pfree(m, v117)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
											if v108 != v122 {
												v128 = v108
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
												return int32(0)
											} else {
												v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
												v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+4)))
												v126 = F_datumCopy(m, v108, v124, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													v128 = v126
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v67 = F_index_getprocinfo(m, v65, v36, int32(12))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v72 = *(*int64)(unsafe.Add(mBase, uint32(v67)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v72
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v44)+8)) = v76
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
							*(*int64)(unsafe.Add(mBase, uint32(v44))) = v78
							*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = int32(0)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v85 = v83
							v86 = v84
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
							v89 = F_FunctionCall2Coll(m, v44, v15, v87, v88)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								if v89 != 0 {
									v102 = F_inclusion_get_procinfo(m, v13, v36&int32(_a_F_brin_inclusion_union_0))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
										v108 = F_FunctionCall2Coll(m, v102, v15, v105, v107)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v110 = int32(4)
											v114 = v14 + v36<<(uint(v110)%32) + v110
											v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
											if v115 != 0 {
												v128 = v108
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
												return int32(0)
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												if v108 == v117 {
													v128 = v108
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
													return int32(0)
												} else {
													F_pfree(m, v117)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
														if v108 != v122 {
															v128 = v108
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
															return int32(0)
														} else {
															v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
															v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114)+4)))
															v126 = F_datumCopy(m, v108, v124, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = v126
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v130))) = v128
																return int32(0)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = int32(1)
									return int32(0)
								}
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
	var v55 int32
	_ = v55
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
	var v84 int32
	_ = v84
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
		v46 = v16 & int32(_a_F_brin_minmax_add_value_0)
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
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
					if v55 != 0 {
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
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
										if v84 != 0 {
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
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
											if v84 != 0 {
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
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
								if v84 != 0 {
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
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
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
	var v77 int32
	_ = v77
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
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
	var v221 int32
	_ = v221
	var v242 int32
	_ = v242
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
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int64
	_ = v410
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
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
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 float64
	_ = v547
	var v550 int32
	_ = v550
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
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
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v877 int32
	_ = v877
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = F_pg_detoast_datum(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v43 = F_brin_range_deserialize(m, v42, v34)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v46 = F_brin_range_deserialize(m, v45, v40)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0]))
	v58 = F_AllocSetContextCreateInternal(m, v53, int32(_a_F_brin_minmax_multi_union_0), int32(0), int32(_a_F_brin_minmax_multi_union_1), int32(_a_F_brin_minmax_multi_union_2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v60 = int32(_a_F_brin_minmax_multi_union_3)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0])) = v58
	v66 = v48 + (v49 + (v51 + v50))
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if int32(0) < v71 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v77 = v2
	goto L11
L9:
	;
	v112 = v71
	v113 = v2
	goto L10
L10:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	if int32(0) < v132 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v98 = v69 + v77*int32(12)
	v101 = v43 + int32(36) + v77<<(uint(int32(3))%32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+8)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v104
	v109 = v77 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v109 < v110 {
		v77 = v109
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v112 = v110
	v113 = v109
	goto L10
L13:
	;
	goto L12
L14:
	;
	v136 = v43 + int32(36)
	v138 = int32(0)
	v139 = v113
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
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v213 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v160 = v69 + v139*int32(12)
	v162 = v138 << (uint(int32(2)) % 32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v164 = int32(3)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v162+(v136+v163<<(uint(v164)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v136+v170<<(uint(v164)%32)+v162)))
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+8)) = uint8(v176)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v175
	v182 = v138 + v176
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	if v182 < v183 {
		v138 = v182
		v139 = v139 + v176
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v189 = v183
	v206 = v185
	goto L16
L19:
	;
	goto L18
L20:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if int32(0) < v276 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v257 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v221 = int32(0)
	goto L24
L24:
	;
	v242 = v212 + v221*int32(12)
	v245 = v46 + int32(36) + v221<<(uint(int32(3))%32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+8)) = uint8(v249)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v248
	v253 = v221 + int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v253 < v254 {
		v221 = v253
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v257 = v253
	goto L20
L26:
	;
	goto L25
L27:
	;
	v280 = v46 + int32(36)
	v282 = int32(0)
	v283 = v257
	goto L30
L28:
	;
	goto L29
L29:
	;
	v349 = int32(1)
	v351 = v30 & int32(_a_F_brin_minmax_multi_union_4)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(4))%32)+v27+v30*int32(100)-int32(12))))
	v362 = F_minmax_multi_get_strategy_procinfo(m, v26, v351, v360, v349)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v304 = v212 + v283*int32(12)
	v306 = v282 << (uint(int32(2)) % 32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v308 = int32(3)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v280+v307<<(uint(v308)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v280+v314<<(uint(v308)%32)+v306)))
	v320 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v304)+8)) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+4)) = v319
	v326 = v282 + v320
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v326 < v327 {
		v282 = v326
		v283 = v283 + v320
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
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v25
	F_qsort_arg(m, v69, v66, int32(12), int32(22), v23+int32(8))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
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
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	v602 = F_reduce_expanded_ranges(m, v69, v585, v583, v601, v362, v25)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L79
	}
L36:
	;
	v375 = v349
	v376 = int32(1)
	goto L39
L37:
	;
	goto L38
L38:
	;
	v576 = F_minmax_multi_get_procinfo(m, v26, v351)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L78
	}
L39:
	;
	v395 = int32(12)
	v397 = v69 + v376*v395
	v402 = F_compare_expanded_ranges(m, v397-v395, v397, v23+int32(8))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v420 = int32(1)
	v422 = v415 - v420
	if int32(0) < v422 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	if v402 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v375 != v376 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v415 = v375
	goto L44
L44:
	;
	v418 = v376 + int32(1)
	if v418 != v66 {
		v375 = v415
		v376 = v418
		goto L39
	} else {
		goto L48
	}
L45:
	;
	v407 = v69 + v375*int32(12)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = v408
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = v410
	goto L47
L46:
	;
	goto L47
L47:
	;
	v415 = v375 + int32(1)
	goto L44
L48:
	;
	goto L40
L49:
	;
	v426 = v415
	v427 = int32(0)
	v428 = v422
	goto L52
L50:
	;
	v480 = v415
	goto L51
L51:
	;
	v503 = F_minmax_multi_get_procinfo(m, v26, v30&int32(_a_F_brin_minmax_multi_union_4))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L67
	}
L52:
	;
	v446 = int32(12)
	v448 = v69 + v427*v446
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v451 = v448 + v446
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v453 = F_FunctionCall2Coll(m, v362, v25, v449, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v480 = v474
	goto L51
L54:
	;
	v478 = v474 - int32(1)
	if v475 < v478 {
		v426 = v474
		v427 = v475
		v428 = v478
		goto L52
	} else {
		goto L66
	}
L55:
	;
	if v453 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v474 = v426
	v475 = v427 + int32(1)
	goto L54
L57:
	;
	goto L58
L58:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v459 = F_FunctionCall2Coll(m, v362, v25, v457, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v459 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v461
	goto L62
L61:
	;
	goto L62
L62:
	;
	v463 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+8)) = uint8(v463)
	v466 = v427 + int32(2)
	v469 = (v426 - v466) * int32(12)
	if v469 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	base.MemoryCopy(m, v451, v69+v466*int32(12), v469)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v474 = v428
	v475 = v427
	goto L54
L66:
	;
	goto L53
L67:
	;
	if v480 == int32(1) {
		v583 = int32(0)
		v585 = v420
		goto L35
	} else {
		goto L68
	}
L68:
	;
	v508 = v480 - int32(1)
	v511 = F_palloc0(m, v508<<(uint(int32(4))%32))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if int32(0) < v508 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v517 = int32(0)
	goto L73
L71:
	;
	goto L72
L72:
	;
	F_pg_qsort(m, v511, v508, int32(16), int32(20))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L77
	}
L73:
	;
	v538 = v69 + v517*int32(12)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	v541 = F_FunctionCall2Coll(m, v503, v25, v539, v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	goto L72
L75:
	;
	v545 = v511 + v517<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v517
	v547 = *(*float64)(unsafe.Add(mBase, uint32(v541)))
	*(*float64)(unsafe.Add(mBase, uint32(v545)+8)) = v547
	v550 = v517 + int32(1)
	if v550 != v508 {
		v517 = v550
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v583 = v511
	v585 = v480
	goto L35
L78:
	;
	v583 = int32(0)
	v585 = int32(1)
	goto L35
L79:
	;
	v604 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v604
	if v604 < v602 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v877
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0])) = v61
	F_MemoryContextDelete(m, v58)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L115
	}
L81:
	;
	v609 = v43 + int32(36)
	v611 = v602 - int32(1)
	if v611 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v856 = int32(0)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v856
	v877 = v856
	goto L80
L84:
	;
	v746 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v746
	if v611 == v746 {
		goto L101
	} else {
		goto L102
	}
L85:
	;
	v711 = v69 + v689*int32(12)
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711)+8)))
	if v712 != 0 {
		v727 = v690
		goto L84
	} else {
		goto L99
	}
L86:
	;
	v614 = int32(0)
	v689 = v614
	v690 = v614
	goto L85
L87:
	;
	goto L88
L88:
	;
	v620 = int32(0)
	v623 = v620
	v624 = v620
	v626 = v620
	goto L89
L89:
	;
	v645 = v69 + v623*int32(12)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+8)))
	if v646 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v602&int32(1) == int32(0) {
		v727 = v680
		goto L84
	} else {
		goto L98
	}
L91:
	;
	v649 = int32(2)
	v651 = v609 + v624<<(uint(v649)%32)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	*(*int32)(unsafe.Add(mBase, uint32(v651))) = v652
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v645)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v654
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v656 + int32(1)
	v662 = v624 + v649
	goto L93
L92:
	;
	v662 = v624
	goto L93
L93:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645)+20)))
	if v664 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v667 = int32(2)
	v669 = v609 + v662<<(uint(v667)%32)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v645)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = v670
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v645)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v669)+4)) = v672
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v674 + int32(1)
	v680 = v662 + v667
	goto L96
L95:
	;
	v680 = v662
	goto L96
L96:
	;
	v682 = int32(2)
	v683 = v623 + v682
	v685 = v626 + v682
	if v685 != v602&int32(2147483646) {
		v623 = v683
		v624 = v680
		v626 = v685
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v689 = v683
	v690 = v680
	goto L85
L99:
	;
	v713 = int32(2)
	v715 = v609 + v690<<(uint(v713)%32)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	*(*int32)(unsafe.Add(mBase, uint32(v715))) = v716
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v715)+4)) = v718
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v720 + int32(1)
	v727 = v690 + v713
	goto L84
L100:
	;
	v842 = v69 + v820*int32(12)
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+8)))
	if v843 != int32(1) {
		v877 = v822
		goto L80
	} else {
		goto L114
	}
L101:
	;
	v820 = int32(0)
	v821 = v727
	v822 = v746
	goto L100
L102:
	;
	goto L103
L103:
	;
	v756 = int32(0)
	v758 = v756
	v759 = v727
	v760 = v746
	v761 = v756
	goto L104
L104:
	;
	v780 = v69 + v758*int32(12)
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+8)))
	if v781 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v602&int32(1) == int32(0) {
		v877 = v812
		goto L80
	} else {
		goto L113
	}
L106:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	*(*int32)(unsafe.Add(mBase, uint32(v609+v759<<(uint(int32(2))%32)))) = v787
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v790 = int32(1)
	v791 = v789 + v790
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v791
	v795 = v759 + v790
	v796 = v791
	goto L108
L107:
	;
	v795 = v759
	v796 = v760
	goto L108
L108:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+20)))
	if v797 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v780)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v609+v795<<(uint(int32(2))%32)))) = v803
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v806 = int32(1)
	v807 = v805 + v806
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v807
	v811 = v795 + v806
	v812 = v807
	goto L111
L110:
	;
	v811 = v795
	v812 = v796
	goto L111
L111:
	;
	v813 = int32(2)
	v814 = v758 + v813
	v816 = v761 + v813
	if v816 != v602&int32(2147483646) {
		v758 = v814
		v759 = v811
		v760 = v812
		v761 = v816
		goto L104
	} else {
		goto L112
	}
L112:
	;
	goto L105
L113:
	;
	v820 = v814
	v821 = v811
	v822 = v812
	goto L100
L114:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	*(*int32)(unsafe.Add(mBase, uint32(v609+v821<<(uint(int32(2))%32)))) = v849
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v856 = v851 + int32(1)
	goto L83
L115:
	;
	F_pfree(m, v34)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v902 = F_brin_range_serialize(m, v43)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v904))) = v902
	m.G0 = v23 + int32(16)
	return int32(0)
}
func F_brin_page_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_superuser(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				v15 = F_get_page_from_raw(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+14)))
					if v17 == int32(0) {
						v20 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
						v56 = int32(0)
						m.G0 = v6 + int32(48)
						return v56
					} else {
						v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
						v24 = int32(8)
						v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+16)))
						if (v23<<(uint(v24)%32)-v26)&int32(_a_F_brin_page_type_0) != v24 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_brin_page_type_1)
									F_errmsg(m, int32(_a_F_brin_page_type_2), v6+int32(32))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+16)))
										v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+19)))
										v93 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v93
										*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = (v92<<(uint(v93)%32) - v91) & int32(_a_F_brin_page_type_0)
										F_errdetail(m, int32(_a_F_brin_page_type_3), v6+int32(16))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_brin_page_type_4), int32(67), int32(_a_F_brin_page_type_5))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
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
							v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26+v15)+6)))
							v35 = v33 + int32(3951)
							if base.Ui32(int32(3)) <= base.Ui32(v35&int32(_a_F_brin_page_type_0)) {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v33
								v42 = F_psprintf(m, int32(_a_F_brin_page_type_6), v6)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v51 = v42
									v52 = F_cstring_to_text(m, v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v56 = v52
										m.G0 = v6 + int32(48)
										return v56
									}
								}
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v35&int32(_a_F_brin_page_type_0)<<(uint(int32(2))%32))+uint32(_c_F_brin_page_type[0])))
								v51 = v50
								v52 = F_cstring_to_text(m, v51)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v56 = v52
									m.G0 = v6 + int32(48)
									return v56
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_brin_page_type_7), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_brin_page_type_4), int32(53), int32(_a_F_brin_page_type_5))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
