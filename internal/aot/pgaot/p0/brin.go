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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v7 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_LockBufferInternal(m, v7, int32(1))
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
				*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v38
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
				F_UnlockBuffer(m, v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					return v33
				}
			}
		}
	}
}
func F_brin_bloom_consistent(m *base.Module, l0 int32) int64 {
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
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
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
	var v249 int64
	_ = v249
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v271 int64
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v303 int32
	_ = v303
	var v335 int64
	_ = v335
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if int32(0) < v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v18 + int32(16)
	return v335
L4:
	;
	v39 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v335 = int64(1)
	goto L3
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22+v39<<(uint(int32(2))%32))))
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+6)))
	if v54 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v303 = v39 + int32(1)
	if v303 != v20 {
		v39 = v303
		goto L7
	} else {
		goto L33
	}
L10:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v53)+48))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)))
	v59 = F_bloom_get_procinfo(m, v23, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L30
	}
L13:
	;
	v61 = F_FunctionCall1Coll(m, v59, v21, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v63 = base.I32_wrap_i64(v61)
	goto L18
L15:
	;
	v148 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+8)))
	v149 = base.I64_rem_u_s(base.I64_extend_i32_u(v138)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v138^v130-base.I32_rotl(v138, int32(24))), v148)
	goto L23
L16:
	;
	v115 = int32(14)
	v117 = v94 - v102 ^ base.I32_rotl(v102, int32(4)) ^ v108 - base.I32_rotl(v108, v115)
	v122 = v117 ^ (v63 + v107) - base.I32_rotl(v117, int32(11))
	v126 = v108 ^ v122 - base.I32_rotl(v122, int32(25))
	v130 = v126 ^ v117 - base.I32_rotl(v126, int32(16))
	v134 = v130 ^ v122 - base.I32_rotl(v130, int32(4))
	v138 = v134 ^ v126 - base.I32_rotl(v134, v115)
	goto L15
L18:
	;
	goto L19
L19:
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
	goto L16
L20:
	;
	v234 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+8)))
	v235 = base.I64_rem_u_s(base.I64_extend_i32_u(v224)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v224^v216-base.I32_rotl(v224, int32(24))), v234)
	v236 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	if v236 == int64(0) {
		goto L9
	} else {
		goto L25
	}
L21:
	;
	v201 = int32(14)
	v203 = v180 - v188 ^ base.I32_rotl(v188, int32(4)) ^ v194 - base.I32_rotl(v194, v201)
	v208 = v203 ^ (v63 + v193) - base.I32_rotl(v203, int32(11))
	v212 = v194 ^ v208 - base.I32_rotl(v208, int32(25))
	v216 = v212 ^ v203 - base.I32_rotl(v212, int32(16))
	v220 = v216 ^ v208 - base.I32_rotl(v216, int32(4))
	v224 = v220 ^ v212 - base.I32_rotl(v220, v201)
	goto L20
L23:
	;
	goto L24
L24:
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
	goto L21
L25:
	;
	v249 = int64(0)
	goto L26
L26:
	;
	v257 = base.I64_rem_u_s(v249*v235+v149, v234)
	v258 = base.I32_wrap_i64(v257)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(16)+int32(base.Ui32(v258)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v262)>>(uint(v258&int32(7))%32))&int32(1) == int32(0) {
		v335 = int64(0)
		goto L3
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	v271 = v249 + int64(1)
	if v236 != v271 {
		v249 = v271
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v277
	F_errmsg_internal(m, int32(_a_F_brin_bloom_consistent_0), v18)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_brin_bloom_consistent_1), int32(646), int32(_a_F_brin_bloom_consistent_2))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	goto L8
}
func F_brin_bloom_summary_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_brin_bloom_summary_in_0), int32(786), int32(_a_F_brin_bloom_summary_in_1), int32(_a_F_brin_bloom_summary_in_2), int32(_a_F_brin_bloom_summary_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int64
	_ = v57
	var v58 int64
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v27 = F_palloc_mul(m, int32(4), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if int32(0) < v29 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = v29
	v34 = v2
	v39 = v2
	goto L7
L5:
	;
	v67 = v29
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
	v44 = int32(1)
	v45 = v34 + v44
	v48 = F_index_getprocinfo(m, l0, base.I32_extend16_s(v45), v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v25+v33<<(uint(int32(3))%32)+v34*int32(100))+96)))
	v58 = F_FunctionCall1Coll(m, v48, int32(0), v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v60 = base.I32_wrap_i64(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v27+v34<<(uint(int32(2))%32)))) = v60
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v63 = v39 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v45 < v64 {
		v33 = v64
		v34 = v45
		v39 = v63
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
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v16
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
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
	F_pfree(m, v27)
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103+v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v79+int32(20)+v103))) = v106
	v109 = v95 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
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
func F_brin_desummarize_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
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
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_brin_desummarize_range[0])))
	if v22 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L12
	} else {
		goto L142
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L12
	} else {
		goto L138
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L12
	} else {
		goto L134
	}
L4:
	;
	if v32 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[1]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+308))
	v30 = base.B2i32(v28 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_brin_desummarize_range[0])) = uint8(v30)
	v32 = v30
	goto L7
L6:
	;
	v32 = int32(0)
	goto L7
L7:
	;
	goto L4
L8:
	;
	if base.Ui64(int64(4294967295)) <= base.Ui64(v18) {
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
	v462 = m.ExcPending
	if v462 != 0 {
		goto L12
	} else {
		goto L129
	}
L11:
	;
	v38 = F_IndexGetRelation(m, v19, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int64(0)
L13:
	;
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = F_table_open(m, v38, int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v46 = int32(0)
	goto L16
L16:
	;
	v48 = F_index_open(m, v19, int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v46 = v43
	goto L16
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+119)))
	if v51 != int32(105) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+84))
	if v54 != int32(3580) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[2]))
	v60 = F_object_ownercheck(m, int32(1259), v19, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if v60 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	F_aclcheck_error(m, int32(2), int32(20), v66+int32(4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v46 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	v74 = F_IndexGetRelation(m, v19, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v74 != v38 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v48)+192))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+18)))
	if v78 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_relation_close(m, v48, int32(4))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L12
	} else {
		goto L127
	}
L30:
	;
	v81 = base.I32_wrap_i64(v18)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v414 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L12
	} else {
		goto L122
	}
L33:
	;
	v95 = m.G0
	v97 = v95 - int32(16)
	m.G0 = v97
	v100 = F_ReadBuffer(m, v48, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L36
	}
L34:
	;
	goto L29
L35:
	;
	if v374 == int32(0) {
		goto L33
	} else {
		goto L121
	}
L36:
	;
	F_LockBufferInternal(m, v100, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v100 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v124 = F_palloc(m, int32(20))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L42
	}
L39:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v100^int32(-1))<<(uint(int32(2))%32))))
	v122 = v114
	goto L38
L40:
	;
	goto L41
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v122 = v116 + v100<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v48
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v129
	F_UnlockBuffer(m, v100)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v138 = base.I32_div_u_s(v81, v137)
	v140 = base.I32_div_u_s(v138, int32(1360))
	if base.Ui32(v136) <= base.Ui32(v140) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L12
	} else {
		goto L117
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L12
	} else {
		goto L113
	}
L46:
	;
	m.G0 = v97 + int32(16)
	goto L35
L47:
	;
	F_pfree(m, v124)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L112
	}
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	F_ReleaseBuffer(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v150 = F_brinLockRevmapPageForUpdate(m, v124, v81)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L55
	}
L51:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	if v145 == int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	F_ReleaseBuffer(m, v145)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v171 = base.I32_div_u_s(v81, v170)
	v173 = base.I32_rem_u_s(v171, int32(1360))
	v176 = v169 + v173*int32(6)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176)+28)))
	if v177 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	if v150 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v155+(v150^int32(-1))<<(uint(int32(2))%32))))
	v169 = v161
	goto L54
L57:
	;
	goto L58
L58:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v169 = v163 + v150<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L59:
	;
	F_UnlockBuffer(m, v150)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L12
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v191 = v176 + int32(24)
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191))))
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+2)))
	v197 = F_ReadBuffer(m, v48, v192<<(uint(int32(16))%32)|v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L66
	}
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	F_ReleaseBuffer(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	if v185 == int32(0) {
		goto L47
	} else {
		goto L64
	}
L64:
	;
	F_ReleaseBuffer(m, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	goto L47
L66:
	;
	F_LockBufferInternal(m, v197, int32(3))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	if v197 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+16)))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220+v219)+6)))
	if v222 != int32(_a_F_brin_desummarize_range_0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v205+(v197^int32(-1))<<(uint(int32(2))%32))))
	v219 = v211
	goto L68
L70:
	;
	goto L71
L71:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v219 = v213 + v197<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	F_UnlockBuffer(m, v150)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+4)))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v239) {
		goto L83
	} else {
		goto L84
	}
L75:
	;
	F_UnlockBuffer(m, v197)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	F_ReleaseBuffer(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	if v232 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_ReleaseBuffer(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L12
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_pfree(m, v124)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L12
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v374 = int32(0)
	goto L46
L83:
	;
	v247 = int32(base.Ui32(v239+int32(_a_F_brin_desummarize_range_1)) >> (uint(int32(2)) % 32))
	goto L85
L84:
	;
	v247 = int32(0)
	goto L85
L85:
	;
	if base.Ui32(v247&int32(_a_F_brin_desummarize_range_2)) < base.Ui32(v238) {
		goto L45
	} else {
		goto L86
	}
L86:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219+v238<<(uint(int32(2))%32))+21)))
	if v254&int32(384) == int32(0) {
		goto L44
	} else {
		goto L87
	}
L87:
	;
	v259 = int32(_a_F_brin_desummarize_range_3)
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5])) = v261 + int32(1)
	if v150 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v284 = base.I32_div_u_s(v81, v283)
	v286 = base.I32_rem_u_s(v284, int32(1360))
	v289 = v282 + v286*int32(6)
	v290 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+28)) = uint16(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+24)) = int32(-1)
	F_PageIndexTupleDeleteNoCompact(m, v219, v238)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L12
	} else {
		goto L92
	}
L89:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[3]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v268+(v150^int32(-1))<<(uint(int32(2))%32))))
	v282 = v274
	goto L88
L90:
	;
	goto L91
L91:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[4]))
	v282 = v276 + v150<<(uint(int32(13))%32) + int32(-8192)
	goto L88
L92:
	;
	F_MarkBufferDirty(m, v197)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	F_MarkBufferDirty(m, v150)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L94
	}
L94:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+118)))
	if v301 != int32(112) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v339 = int32(_a_F_brin_desummarize_range_3)
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[5])) = v341 - int32(1)
	F_UnlockReleaseBuffer(m, v197)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L12
	} else {
		goto L107
	}
L96:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_brin_desummarize_range[6]))
	if v305 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	if v308 != 0 {
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+12)) = uint16(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v310
	F_XLogBeginInsert(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L12
	} else {
		goto L102
	}
L100:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v309 != 0 {
		goto L95
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	F_XLogRegisterData(m, v97+int32(4), int32(10))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	v321 = int32(0)
	F_XLogRegisterBuffer(m, v321, v150, v321)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_XLogRegisterBuffer(m, int32(1), v197, int32(8))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v331 = F_XLogInsert(m, int32(17), int32(80))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	v334 = base.I64_rotl(v331, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v219))) = v334
	goto L95
L107:
	;
	F_UnlockBuffer(m, v150)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	F_ReleaseBuffer(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	if v352 == int32(0) {
		goto L47
	} else {
		goto L110
	}
L110:
	;
	F_ReleaseBuffer(m, v352)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L111
	}
L111:
	;
	goto L47
L112:
	;
	v374 = int32(1)
	goto L46
L113:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_brin_desummarize_range_4), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_5), int32(383), int32(_a_F_brin_desummarize_range_6))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
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
	v400 = m.ExcPending
	if v400 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_brin_desummarize_range_4), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_5), int32(389), int32(_a_F_brin_desummarize_range_6))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
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
	if v414 == int32(0) {
		goto L29
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v421 + int32(4)
	F_errmsg(m, int32(_a_F_brin_desummarize_range_7), v14+int32(-32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1574), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	goto L29
L127:
	;
	F_relation_close(m, v46, int32(4))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	m.G0 = v16 - int32(-64)
	return int64(0)
L129:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_brin_desummarize_range_10), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	F_errhint(m, int32(_a_F_brin_desummarize_range_11), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1510), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
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
	v485 = m.ExcPending
	if v485 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v18
	F_errmsg(m, int32(_a_F_brin_desummarize_range_12), v16)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L12
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1516), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
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
	v501 = m.ExcPending
	if v501 != 0 {
		goto L12
	} else {
		goto L139
	}
L139:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v502 + int32(4)
	F_errmsg(m, int32(_a_F_brin_desummarize_range_13), v14+int32(-16))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1542), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
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
	v522 = m.ExcPending
	if v522 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v523 + int32(4)
	F_errmsg(m, int32(_a_F_brin_desummarize_range_14), v14+int32(-48))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_brin_desummarize_range_8), int32(1558), int32(_a_F_brin_desummarize_range_9))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
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
	var v52 int32
	_ = v52
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
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
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v530 int64
	_ = v530
	var v531 int32
	_ = v531
	var v533 int64
	_ = v533
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
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
	v628 = m.ExcPending
	if v628 != 0 {
		goto L6
	} else {
		goto L202
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
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
	v593 = m.ExcPending
	if v593 != 0 {
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
	return v576
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
	F_LockBufferInternal(m, l4, int32(3))
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
	v576 = int32(0)
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
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v33^int32(-1))*int32(56))+16))
	v58 = v49
	goto L19
L21:
	;
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[1]))
	v52 = int32(56)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+v33*v52-v52)+16))
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
	v134 = v87 + v108&int32(_a_F_brin_doupdate_1)
	v136 = int32(base.Ui32(v108) >> (uint(int32(17)) % 32))
	if l7 == v136 {
		goto L45
	} else {
		goto L46
	}
L25:
	;
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	v116 = int32(0)
	if v68 == v116 {
		v576 = v116
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v119 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
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
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L42
	}
L41:
	;
	v576 = v116
	goto L8
L42:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v576 = v116
	goto L8
L45:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v136) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v201 = int32(1)
	goto L47
L47:
	;
	if v201 != 0 {
		goto L66
	} else {
		goto L67
	}
L48:
	;
	v201 = v199
	goto L47
L49:
	;
	v199 = int32(0)
	goto L48
L50:
	;
	v173 = v168
	v174 = v169
	v175 = v170
	goto L60
L51:
	;
	if (v134|l6)&int32(3) != 0 {
		v168 = v134
		v169 = l6
		v170 = v136
		goto L50
	} else {
		goto L54
	}
L52:
	;
	v161 = v134
	v162 = l6
	v163 = v136
	goto L53
L53:
	;
	if v163 == int32(0) {
		goto L49
	} else {
		goto L59
	}
L54:
	;
	v145 = v134
	v146 = l6
	v147 = v136
	goto L55
L55:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v150 != v151 {
		v168 = v145
		v169 = v146
		v170 = v147
		goto L50
	} else {
		goto L57
	}
L56:
	;
	v161 = v156
	v162 = v154
	v163 = v158
	goto L53
L57:
	;
	v153 = int32(4)
	v154 = v146 + v153
	v156 = v145 + v153
	v158 = v147 - v153
	if base.Ui32(int32(3)) < base.Ui32(v158) {
		v145 = v156
		v146 = v154
		v147 = v158
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v168 = v161
	v169 = v162
	v170 = v163
	goto L50
L60:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v178 == v179 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v199 = v178 - v179
	goto L48
L62:
	;
	v181 = int32(1)
	v186 = v175 - v181
	if v186 != 0 {
		v173 = v173 + v181
		v174 = v174 + v181
		v175 = v186
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
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+16)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v220)+4)))
	if v222&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L69:
	;
	v204 = int32(0)
	if v68 == v204 {
		v576 = v204
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v207 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
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
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L75
	}
L74:
	;
	v576 = v204
	goto L8
L75:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v576 = v204
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
	v253 = int32(_a_F_brin_doupdate_7)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v255 + int32(1)
	v259 = F_PageIndexTupleOverwrite(m, v87, v6, l8, l9)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L92
	}
L83:
	;
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+14)))
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+12)))
	v246 = v244 - v245
	v247 = int32(0)
	if v247 < v246 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v229+(l4^int32(-1))<<(uint(int32(2))%32))))
	v243 = v235
	goto L83
L85:
	;
	goto L86
L86:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v243 = v237 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L83
L87:
	;
	if base.Ui32(v250) < base.Ui32(l9-l7) {
		goto L78
	} else {
		goto L91
	}
L88:
	;
	v250 = v246
	goto L90
L89:
	;
	v250 = v247
	goto L90
L90:
	;
	goto L87
L91:
	;
	goto L82
L92:
	;
	if v259 == int32(0) {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+118)))
	if v266 != int32(112) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v297 = int32(_a_F_brin_doupdate_7)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v299 - int32(1)
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L107
	}
L96:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[5]))
	if v270 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v273 != 0 {
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v274 != 0 {
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
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_XLogRegisterBuffer(m, int32(0), l4, int32(8))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_XLogRegisterBufData(m, int32(0), l8, l9)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v292 = F_XLogInsert(m, int32(17), int32(48))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = base.I64_rotl(v292, int64(32))
	goto L95
L107:
	;
	v305 = int32(1)
	if v68 == int32(0) {
		v576 = v305
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v308 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
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
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L113
	}
L112:
	;
	v576 = v305
	goto L8
L113:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	v576 = v305
	goto L8
L116:
	;
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
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
	v576 = int32(0)
	goto L8
L120:
	;
	v344 = F_brinLockRevmapPageForUpdate(m, l2, l3)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L124
	}
L121:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329+(v68^int32(-1))<<(uint(int32(2))%32))))
	v343 = v335
	goto L120
L122:
	;
	goto L123
L123:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v343 = v337 + v68<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L124:
	;
	v346 = int32(_a_F_brin_doupdate_7)
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	v349 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v348 + v349
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v352 == v349 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v355 = int32(_a_F_brin_doupdate_8)
	v357 = int32(0)
	if v357|(v343&int32(3)|int32(1)) == v357 {
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
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L139
	}
L128:
	;
	v405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+16)))
	v407 = int32(_a_F_brin_doupdate_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v343+v405)+6)) = uint16(v407)
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+10)) = int32(_a_F_brin_doupdate_9)
	v396 = int32(_a_F_brin_doupdate_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v343)+18)) = uint16(v396)
	v402 = int32(_a_F_brin_doupdate_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v343)+16)) = uint16(v402)
	*(*uint16)(unsafe.Add(mBase, uint32(v343)+14)) = uint16(v402)
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
	v373 = v343 + v355
	v375 = v343 + int32(4)
	if base.Ui32(v375) < base.Ui32(v373) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v377 = v373
	goto L136
L135:
	;
	v377 = v375
	goto L136
L136:
	;
	v382 = (v343^int32(-1)+v377)&int32(-4) + int32(4)
	if v382 == int32(0) {
		goto L129
	} else {
		goto L137
	}
L137:
	;
	base.MemoryFill(m, v343, int32(0), v382)
	goto L129
L138:
	;
	base.MemoryFill(m, v343, int32(0), v355)
	goto L129
L139:
	;
	v411 = int32(0)
	v413 = F_PageAddItemExtended(m, v343, l8, l9, v411, v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	if v413 == int32(0) {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	F_MarkBufferDirty(m, v68)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	if v352 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+16)))
	v422 = v343 + v421
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422)+6)))
	if v423 != int32(_a_F_brin_doupdate_2) {
		v438 = v12
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v440 = v12
	goto L146
L146:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v413)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v413)
	v443 = int32(16)
	v444 = base.I32_rotr(v69, v443)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v444
	v448 = v21 + v443
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448))))
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+2)))
	if v344 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L147:
	;
	v440 = v438
	goto L146
L148:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+4)))
	if v426&int32(1) != 0 {
		v438 = v12
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v429 = int32(4)
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+14)))
	v431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+12)))
	v432 = v430 - v431
	if v432 <= v429 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v438 = v435 - int32(4)
	goto L147
L151:
	;
	v435 = v429
	goto L153
L152:
	;
	v435 = v432
	goto L153
L153:
	;
	goto L150
L154:
	;
	F_MarkBufferDirty(m, v344)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L6
	} else {
		goto L165
	}
L155:
	;
	v471 = base.I32_div_u_s(l3, l1)
	v473 = base.I32_rem_u_s(v471, int32(1360))
	v476 = v470 + v473*int32(6)
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v448)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v476)+28)) = uint16(v477)
	if v477 != 0 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v456+(v344^int32(-1))<<(uint(int32(2))%32))))
	v470 = v462
	goto L155
L157:
	;
	goto L158
L158:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v470 = v464 + v344<<(uint(int32(13))%32) + int32(-8192)
	goto L155
L159:
	;
	v480 = v452
	goto L161
L160:
	;
	v480 = int32(-1)
	goto L161
L161:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v476)+26)) = uint16(v480)
	if v477 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v483 = v451
	goto L164
L163:
	;
	v483 = int32(-1)
	goto L164
L164:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v476)+24)) = uint16(v483)
	goto L154
L165:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+118)))
	if v488 != int32(112) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v556 = int32(_a_F_brin_doupdate_7)
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[4])) = v558 - int32(1)
	F_UnlockBuffer(m, v344)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L6
	} else {
		goto L187
	}
L167:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[5]))
	if v492 <= int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v495 != 0 {
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
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v413)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+24)) = uint16(v6)
	F_XLogBeginInsert(m)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L6
	} else {
		goto L173
	}
L171:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v496 != 0 {
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
	v507 = m.ExcPending
	if v507 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	if v352 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v511 = int32(14)
	goto L177
L176:
	;
	v511 = int32(8)
	goto L177
L177:
	;
	F_XLogRegisterBuffer(m, int32(0), v68, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_XLogRegisterBufData(m, int32(0), l8, l9)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	F_XLogRegisterBuffer(m, int32(1), v344, int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_XLogRegisterBuffer(m, int32(2), l4, int32(8))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	v530 = F_XLogInsert(m, int32(17), v352<<(uint(int32(7))%32)|int32(32))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	v533 = base.I64_rotl(v530, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = v533
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = v533
	if v344 < int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v553))) = v533
	goto L166
L184:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[2]))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v539+(v344^int32(-1))<<(uint(int32(2))%32))))
	v553 = v545
	goto L183
L185:
	;
	goto L186
L186:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doupdate[3]))
	v553 = v547 + v344<<(uint(int32(13))%32) + int32(-8192)
	goto L183
L187:
	;
	F_UnlockBuffer(m, l4)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	if v352 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_RecordPageWithFreeSpace(m, l0, v69, v440)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L6
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v576 = int32(1)
	goto L8
L193:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v69, v69+int32(1))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
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
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(_a_F_brin_doupdate_12)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v597 + int32(4)
	F_errmsg(m, int32(_a_F_brin_doupdate_13), v21)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_brin_doupdate_14), int32(76), int32(_a_F_brin_doupdate_15))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
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
	v619 = m.ExcPending
	if v619 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_brin_doupdate_14), int32(180), int32(_a_F_brin_doupdate_15))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
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
	v632 = m.ExcPending
	if v632 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_brin_doupdate_14), int32(255), int32(_a_F_brin_doupdate_15))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
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
func F_brin_inclusion_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	if v19 != int64(0) {
		v29 = v18
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
		if v23 == int64(0) {
			v29 = v18
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = int64(1)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v29 = v28
		}
	}
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	if v30 == int64(0) {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
		if v34 != int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(1)
			return int64(0)
		} else {
			v41 = base.I32_extend16_s(v13)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v14+v13<<(uint(int32(2))%32))+16))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+113)))
			if v47 != 0 {
				v109 = F_inclusion_get_procinfo(m, v14, v41&int32(_a_F_brin_inclusion_union_0))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int64(0)
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
					v115 = F_FunctionCall2Coll(m, v109, v16, v112, v114)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int64(0)
					} else {
						v121 = v15 + v41<<(uint(int32(3))%32) + int32(20)
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
						if v122 != 0 {
							v137 = v115
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
							return int64(0)
						} else {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
							if v124 == base.I32_wrap_i64(v115) {
								v137 = v115
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
								return int64(0)
							} else {
								F_pfree(m, v124)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int64(0)
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									v130 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
									if v115 != v130 {
										v137 = v115
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
										return int64(0)
									} else {
										v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
										v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+2)))
										v134 = F_datumCopy(m, v115, v132, v133)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int64(0)
										} else {
											v137 = v134
											v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
											return int64(0)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v49 = v46 + int32(28)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
				if v50 != 0 {
					v90 = v33
					v91 = v29
					v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
					v93 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
					v94 = F_FunctionCall2Coll(m, v49, v16, v92, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int64(0)
					} else {
						if v94 != int64(0) {
							v109 = F_inclusion_get_procinfo(m, v14, v41&int32(_a_F_brin_inclusion_union_0))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int64(0)
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
								v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
								v115 = F_FunctionCall2Coll(m, v109, v16, v112, v114)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									v121 = v15 + v41<<(uint(int32(3))%32) + int32(20)
									v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
									if v122 != 0 {
										v137 = v115
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
										return int64(0)
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
										if v124 == base.I32_wrap_i64(v115) {
											v137 = v115
											v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
											return int64(0)
										} else {
											F_pfree(m, v124)
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int64(0)
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
												v130 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
												if v115 != v130 {
													v137 = v115
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
													return int64(0)
												} else {
													v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
													v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+2)))
													v134 = F_datumCopy(m, v115, v132, v133)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int64(0)
													} else {
														v137 = v134
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
														*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
														return int64(0)
													}
												}
											}
										}
									}
								}
							}
						} else {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = int64(1)
							return int64(0)
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+216))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+204))
					v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v53+v55*(v41-int32(1))<<(uint(int32(2))%32)+int32(48)-int32(4))))
					if v67 == int32(0) {
						v103 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+113)) = uint8(v103)
						v109 = F_inclusion_get_procinfo(m, v14, v41&int32(_a_F_brin_inclusion_union_0))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int64(0)
						} else {
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
							v115 = F_FunctionCall2Coll(m, v109, v16, v112, v114)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int64(0)
							} else {
								v121 = v15 + v41<<(uint(int32(3))%32) + int32(20)
								v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
								if v122 != 0 {
									v137 = v115
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
									return int64(0)
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
									if v124 == base.I32_wrap_i64(v115) {
										v137 = v115
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
										return int64(0)
									} else {
										F_pfree(m, v124)
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int64(0)
										} else {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
											v130 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
											if v115 != v130 {
												v137 = v115
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
												return int64(0)
											} else {
												v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
												v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+2)))
												v134 = F_datumCopy(m, v115, v132, v133)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int64(0)
												} else {
													v137 = v134
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
													return int64(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						v72 = F_index_getprocinfo(m, v70, v41, int32(12))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int64(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v77
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v79
							v81 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v81
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
							*(*int64)(unsafe.Add(mBase, uint32(v49))) = v83
							*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v76
							*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v90 = v88
							v91 = v89
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
							v94 = F_FunctionCall2Coll(m, v49, v16, v92, v93)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int64(0)
							} else {
								if v94 != int64(0) {
									v109 = F_inclusion_get_procinfo(m, v14, v41&int32(_a_F_brin_inclusion_union_0))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int64(0)
									} else {
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
										v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)))
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
										v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
										v115 = F_FunctionCall2Coll(m, v109, v16, v112, v114)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int64(0)
										} else {
											v121 = v15 + v41<<(uint(int32(3))%32) + int32(20)
											v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
											if v122 != 0 {
												v137 = v115
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
												return int64(0)
											} else {
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
												if v124 == base.I32_wrap_i64(v115) {
													v137 = v115
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
													return int64(0)
												} else {
													F_pfree(m, v124)
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int64(0)
													} else {
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
														v130 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
														if v115 != v130 {
															v137 = v115
															v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
															*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
															return int64(0)
														} else {
															v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
															v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+2)))
															v134 = F_datumCopy(m, v115, v132, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int64(0)
															} else {
																v137 = v134
																v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
																*(*int64)(unsafe.Add(mBase, uint32(v138))) = v137
																return int64(0)
															}
														}
													}
												}
											}
										}
									}
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = int64(1)
									return int64(0)
								}
							}
						}
					}
				}
			}
		}
	} else {
		return int64(0)
	}
}
func F_brin_minmax_add_value(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
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
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15))))
	v21 = v10 + v11<<(uint(int32(3))%32) + v16*int32(100) - int32(72)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)))
	if v23 == int32(1) {
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
		v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
		v28 = F_datumCopy(m, v22, v26, v27)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int64(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v28
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
			v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
			v36 = F_datumCopy(m, v22, v34, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v36
				v40 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v40)
				return int64(1)
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
			return int64(0)
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
			v53 = F_FunctionCall2Coll(m, v49, v44, v22, v52)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int64(0)
			} else {
				if v53 != int64(0) {
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
					if v57 != 0 {
						v64 = int32(1)
						v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
						v68 = F_datumCopy(m, v22, v64&int32(1), v67)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int64(0)
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v70))) = v68
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
							v75 = F_minmax_get_strategy_procinfo(m, v9, v46, v73, int32(5))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int64(0)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = F_FunctionCall2Coll(m, v75, v44, v22, v78)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									if v79 == int64(0) {
										return base.I64_extend_i32_u(base.B2i32(v53 != int64(0)))
									} else {
										v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
										if v87 != 0 {
											v94 = int32(1)
											v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
											v98 = F_datumCopy(m, v22, v94&int32(1), v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int64(0)
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
												*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v98
												return int64(1)
											}
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
											F_pfree(m, v90)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int64(0)
											} else {
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
												v94 = v93
												v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
												v98 = F_datumCopy(m, v22, v94&int32(1), v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int64(0)
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v98
													return int64(1)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						F_pfree(m, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
							v64 = v63
							v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
							v68 = F_datumCopy(m, v22, v64&int32(1), v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int64(0)
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v70))) = v68
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
								v75 = F_minmax_get_strategy_procinfo(m, v9, v46, v73, int32(5))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
									v79 = F_FunctionCall2Coll(m, v75, v44, v22, v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int64(0)
									} else {
										if v79 == int64(0) {
											return base.I64_extend_i32_u(base.B2i32(v53 != int64(0)))
										} else {
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
											if v87 != 0 {
												v94 = int32(1)
												v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
												v98 = F_datumCopy(m, v22, v94&int32(1), v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int64(0)
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
													*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v98
													return int64(1)
												}
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
												F_pfree(m, v90)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int64(0)
												} else {
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
													v94 = v93
													v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
													v98 = F_datumCopy(m, v22, v94&int32(1), v97)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int64(0)
													} else {
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
														*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v98
														return int64(1)
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
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
					v75 = F_minmax_get_strategy_procinfo(m, v9, v46, v73, int32(5))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int64(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = F_FunctionCall2Coll(m, v75, v44, v22, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							if v79 == int64(0) {
								return base.I64_extend_i32_u(base.B2i32(v53 != int64(0)))
							} else {
								v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
								if v87 != 0 {
									v94 = int32(1)
									v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
									v98 = F_datumCopy(m, v22, v94&int32(1), v97)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int64(0)
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
										*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v98
										return int64(1)
									}
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
									F_pfree(m, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+82)))
										v94 = v93
										v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+72)))
										v98 = F_datumCopy(m, v22, v94&int32(1), v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int64(0)
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
											*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v98
											return int64(1)
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
func F_brin_minmax_multi_distance_int4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_reinterpret_f64(base.F64_sub(base.F64_convert_i32_s(v2), base.F64_convert_i32_s(v4)))
}
func F_brin_minmax_multi_distance_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int64
	_ = v4
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_reinterpret_f64(base.F64_sub(base.F64_convert_i64_s(v2), base.F64_convert_i64_s(v4)))
}
func F_brin_minmax_multi_distance_tid(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v5 = int32(16)
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v9 = int32(291)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	return base.I64_reinterpret_f64(base.F64_sub(base.F64_convert_i32_u((v4<<(uint(v5)%32)|v7)*v9+v11), base.F64_convert_i32_u((v15<<(uint(v5)%32)|v18)*v9+v22)))
}
func F_brin_minmax_multi_distance_uuid(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+15)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return base.I64_reinterpret_f64(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v4-v6), v9), base.F64_convert_i32_s(v11-v12)), v9), base.F64_convert_i32_s(v18-v19)), v9), base.F64_convert_i32_s(v25-v26)), v9), base.F64_convert_i32_s(v32-v33)), v9), base.F64_convert_i32_s(v39-v40)), v9), base.F64_convert_i32_s(v46-v47)), v9), base.F64_convert_i32_s(v53-v54)), v9), base.F64_convert_i32_s(v60-v61)), v9), base.F64_convert_i32_s(v67-v68)), v9), base.F64_convert_i32_s(v74-v75)), v9), base.F64_convert_i32_s(v81-v82)), v9), base.F64_convert_i32_s(v88-v89)), v9), base.F64_convert_i32_s(v95-v96)), v9), base.F64_convert_i32_s(v102-v103)), v9), base.F64_convert_i32_s(v109-v110)), v9))
}
func F_brin_minmax_multi_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int64
	_ = v320
	var v322 int32
	_ = v322
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int64
	_ = v462
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v466 int64
	_ = v466
	var v467 int32
	_ = v467
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v475 int32
	_ = v475
	var v478 int64
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int64
	_ = v560
	var v561 int64
	_ = v561
	var v562 int64
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int64
	_ = v672
	var v674 int64
	_ = v674
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int64
	_ = v690
	var v692 int64
	_ = v692
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v810 int64
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v826 int64
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v873 int64
	_ = v873
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v902 int32
	_ = v902
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = F_pg_detoast_datum(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v44 = F_brin_range_deserialize(m, v43, v35)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v47 = F_brin_range_deserialize(m, v46, v41)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0]))
	v59 = F_AllocSetContextCreateInternal(m, v54, int32(_a_F_brin_minmax_multi_union_0), int32(0), int32(_a_F_brin_minmax_multi_union_1), int32(_a_F_brin_minmax_multi_union_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v61 = int32(_a_F_brin_minmax_multi_union_3)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0])) = v59
	v67 = v49 + (v50 + (v52 + v51))
	v70 = F_palloc0(m, v67*int32(24))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if int32(0) < v72 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v78 = v2
	goto L11
L9:
	;
	v114 = v72
	v115 = v2
	goto L10
L10:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if int32(0) < v135 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v100 = v70 + v78*int32(24)
	v103 = v44 + int32(40) + v78<<(uint(int32(4))%32)
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v100))) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+16)) = uint8(v107)
	*(*int64)(unsafe.Add(mBase, uint32(v100)+8)) = v106
	v111 = v78 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v111 < v112 {
		v78 = v111
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v114 = v112
	v115 = v111
	goto L10
L13:
	;
	goto L12
L14:
	;
	v139 = v44 + int32(40)
	v141 = int32(0)
	v142 = v115
	goto L17
L15:
	;
	v193 = v135
	v211 = v114
	goto L16
L16:
	;
	v212 = int32(24)
	v217 = v70 + v211*v212 + v193*v212
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v218 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v164 = v70 + v142*int32(24)
	v166 = v141 << (uint(int32(3)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v168 = int32(4)
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v166+(v139+v167<<(uint(v168)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v139+v174<<(uint(v168)%32)+v166)))
	v180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+16)) = uint8(v180)
	*(*int64)(unsafe.Add(mBase, uint32(v164)+8)) = v179
	v186 = v141 + v180
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if v186 < v187 {
		v141 = v186
		v142 = v142 + v180
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v193 = v187
	v211 = v189
	goto L16
L19:
	;
	goto L18
L20:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if int32(0) < v283 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v263 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v226 = int32(0)
	goto L24
L24:
	;
	v248 = v217 + v226*int32(24)
	v251 = v47 + int32(40) + v226<<(uint(int32(4))%32)
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v251)))
	*(*int64)(unsafe.Add(mBase, uint32(v248))) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v251)+8))
	v255 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+16)) = uint8(v255)
	*(*int64)(unsafe.Add(mBase, uint32(v248)+8)) = v254
	v259 = v226 + int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v259 < v260 {
		v226 = v259
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v263 = v259
	goto L20
L26:
	;
	goto L25
L27:
	;
	v287 = v47 + int32(40)
	v289 = int32(0)
	v290 = v263
	goto L30
L28:
	;
	goto L29
L29:
	;
	v358 = int32(1)
	v360 = v31 & int32(_a_F_brin_minmax_multi_union_4)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(3))%32)+v28+v31*int32(100)-int32(4))))
	v371 = F_minmax_multi_get_strategy_procinfo(m, v27, v360, v369, v358)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v312 = v217 + v290*int32(24)
	v314 = v289 << (uint(int32(3)) % 32)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v316 = int32(4)
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v314+(v287+v315<<(uint(v316)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v312))) = v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v287+v322<<(uint(v316)%32)+v314)))
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v312)+16)) = uint8(v328)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+8)) = v327
	v334 = v289 + v328
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v334 < v335 {
		v289 = v334
		v290 = v290 + v328
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v26
	F_qsort_arg(m, v70, v67, int32(24), int32(23), v24+int32(8))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if int32(1) < v67 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	v621 = F_reduce_expanded_ranges(m, v70, v603, v601, v620, v371, v26)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L79
	}
L36:
	;
	v384 = v358
	v385 = int32(1)
	goto L39
L37:
	;
	goto L38
L38:
	;
	v594 = F_minmax_multi_get_procinfo(m, v27, v360)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L78
	}
L39:
	;
	v405 = int32(24)
	v407 = v70 + v385*v405
	v412 = F_compare_expanded_ranges(m, v407-v405, v407, v24+int32(8))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v432 = int32(1)
	v434 = v427 - v432
	if int32(0) < v434 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	if v412 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v384 != v385 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v427 = v384
	goto L44
L44:
	;
	v430 = v385 + int32(1)
	if v430 != v67 {
		v384 = v427
		v385 = v430
		goto L39
	} else {
		goto L48
	}
L45:
	;
	v417 = v70 + v384*int32(24)
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v407)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v417)+16)) = v418
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v407)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v417)+8)) = v420
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	*(*int64)(unsafe.Add(mBase, uint32(v417))) = v422
	goto L47
L46:
	;
	goto L47
L47:
	;
	v427 = v384 + int32(1)
	goto L44
L48:
	;
	goto L40
L49:
	;
	v438 = v427
	v439 = int32(0)
	v440 = v434
	goto L52
L50:
	;
	v497 = v427
	goto L51
L51:
	;
	v521 = F_minmax_multi_get_procinfo(m, v27, v31&int32(_a_F_brin_minmax_multi_union_4))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L67
	}
L52:
	;
	v459 = int32(24)
	v461 = v70 + v439*v459
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v461)+8))
	v464 = v461 + v459
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v464)))
	v466 = F_FunctionCall2Coll(m, v371, v26, v462, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v497 = v491
	goto L51
L54:
	;
	v495 = v491 - int32(1)
	if v492 < v495 {
		v438 = v491
		v439 = v492
		v440 = v495
		goto L52
	} else {
		goto L66
	}
L55:
	;
	if v466 != int64(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v491 = v438
	v492 = v439 + int32(1)
	goto L54
L57:
	;
	goto L58
L58:
	;
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v461)+8))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v464)+8))
	v474 = F_FunctionCall2Coll(m, v371, v26, v472, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v474 != int64(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v464)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v461)+8)) = v478
	goto L62
L61:
	;
	goto L62
L62:
	;
	v480 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v461)+16)) = uint8(v480)
	v483 = v439 + int32(2)
	v486 = (v438 - v483) * int32(24)
	if v486 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	base.MemoryCopy(m, v464, v70+v483*int32(24), v486)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v491 = v440
	v492 = v439
	goto L54
L66:
	;
	goto L53
L67:
	;
	if v497 == int32(1) {
		v601 = int32(0)
		v603 = v432
		goto L35
	} else {
		goto L68
	}
L68:
	;
	v527 = v497 - int32(1)
	v528 = F_palloc0_mul(m, int32(16), v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if int32(0) < v527 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v534 = int32(0)
	goto L73
L71:
	;
	goto L72
L72:
	;
	F_pg_qsort(m, v528, v527, int32(16), int32(21))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L77
	}
L73:
	;
	v556 = v528 + v534<<(uint(int32(4))%32)
	v559 = v70 + v534*int32(24)
	v560 = *(*int64)(unsafe.Add(mBase, uint32(v559)+8))
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v559)+24))
	v562 = F_FunctionCall2Coll(m, v521, v26, v560, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	goto L72
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v556)+8)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v556))) = v534
	v567 = v534 + int32(1)
	if v567 != v527 {
		v534 = v567
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v601 = v528
	v603 = v497
	goto L35
L78:
	;
	v601 = int32(0)
	v603 = int32(1)
	goto L35
L79:
	;
	v623 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v623
	if v623 < v621 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v902
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_union[0])) = v62
	F_MemoryContextDelete(m, v59)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L115
	}
L81:
	;
	v628 = v44 + int32(40)
	v630 = v621 - int32(1)
	if v630 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v880 = int32(0)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v880
	v902 = v880
	goto L80
L84:
	;
	v768 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v768
	if v630 == v768 {
		goto L101
	} else {
		goto L102
	}
L85:
	;
	v732 = v70 + v709*int32(24)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732)+16)))
	if v733 != 0 {
		v748 = v710
		goto L84
	} else {
		goto L99
	}
L86:
	;
	v633 = int32(0)
	v709 = v633
	v710 = v633
	goto L85
L87:
	;
	goto L88
L88:
	;
	v639 = int32(0)
	v642 = v639
	v643 = v639
	v645 = v639
	goto L89
L89:
	;
	v665 = v70 + v642*int32(24)
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665)+16)))
	if v666 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v621&int32(1) == int32(0) {
		v748 = v700
		goto L84
	} else {
		goto L98
	}
L91:
	;
	v671 = v628 + v643<<(uint(int32(3))%32)
	v672 = *(*int64)(unsafe.Add(mBase, uint32(v665)))
	*(*int64)(unsafe.Add(mBase, uint32(v671))) = v672
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v665)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v671)+8)) = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v676 + int32(1)
	v682 = v643 + int32(2)
	goto L93
L92:
	;
	v682 = v643
	goto L93
L93:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665)+40)))
	if v684 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v689 = v628 + v682<<(uint(int32(3))%32)
	v690 = *(*int64)(unsafe.Add(mBase, uint32(v665)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v689))) = v690
	v692 = *(*int64)(unsafe.Add(mBase, uint32(v665)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v689)+8)) = v692
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v694 + int32(1)
	v700 = v682 + int32(2)
	goto L96
L95:
	;
	v700 = v682
	goto L96
L96:
	;
	v702 = int32(2)
	v703 = v642 + v702
	v705 = v645 + v702
	if v705 != v621&int32(2147483646) {
		v642 = v703
		v643 = v700
		v645 = v705
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v709 = v703
	v710 = v700
	goto L85
L99:
	;
	v736 = v628 + v710<<(uint(int32(3))%32)
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v732)))
	*(*int64)(unsafe.Add(mBase, uint32(v736))) = v737
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v732)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v736)+8)) = v739
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v741 + int32(1)
	v748 = v710 + int32(2)
	goto L84
L100:
	;
	v866 = v70 + v843*int32(24)
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+16)))
	if v867 != int32(1) {
		v902 = v845
		goto L80
	} else {
		goto L114
	}
L101:
	;
	v843 = int32(0)
	v844 = v748
	v845 = v768
	goto L100
L102:
	;
	goto L103
L103:
	;
	v778 = int32(0)
	v780 = v778
	v781 = v748
	v782 = v768
	v783 = v778
	goto L104
L104:
	;
	v803 = v70 + v780*int32(24)
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+16)))
	if v804 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v621&int32(1) == int32(0) {
		v902 = v835
		goto L80
	} else {
		goto L113
	}
L106:
	;
	v810 = *(*int64)(unsafe.Add(mBase, uint32(v803)))
	*(*int64)(unsafe.Add(mBase, uint32(v628+v781<<(uint(int32(3))%32)))) = v810
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v813 = int32(1)
	v814 = v812 + v813
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v814
	v818 = v781 + v813
	v819 = v814
	goto L108
L107:
	;
	v818 = v781
	v819 = v782
	goto L108
L108:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+40)))
	if v820 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v803)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v628+v818<<(uint(int32(3))%32)))) = v826
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v829 = int32(1)
	v830 = v828 + v829
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v830
	v834 = v818 + v829
	v835 = v830
	goto L111
L110:
	;
	v834 = v818
	v835 = v819
	goto L111
L111:
	;
	v836 = int32(2)
	v837 = v780 + v836
	v839 = v783 + v836
	if v839 != v621&int32(2147483646) {
		v780 = v837
		v781 = v834
		v782 = v835
		v783 = v839
		goto L104
	} else {
		goto L112
	}
L112:
	;
	goto L105
L113:
	;
	v843 = v837
	v844 = v834
	v845 = v835
	goto L100
L114:
	;
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v866)))
	*(*int64)(unsafe.Add(mBase, uint32(v628+v844<<(uint(int32(3))%32)))) = v873
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v880 = v875 + int32(1)
	goto L83
L115:
	;
	F_pfree(m, v35)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v928 = F_brin_range_serialize(m, v44)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v930))) = base.I64_extend_i32_u(v928)
	m.G0 = v24 + int32(16)
	return int64(0)
}
func F_brin_page_type(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = F_superuser(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			if v14 != 0 {
				v16 = F_get_page_from_raw(m, v10)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+14)))
					if v18 == int32(0) {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
						v58 = int64(0)
						m.G0 = v7 + int32(48)
						return v58
					} else {
						v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+19)))
						v25 = int32(8)
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)))
						if (v24<<(uint(v25)%32)-v27)&int32(_a_F_brin_page_type_0) != v25 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_brin_page_type_1)
									F_errmsg(m, int32(_a_F_brin_page_type_2), v7+int32(32))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)))
										v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+19)))
										v95 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v95
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = (v94<<(uint(v95)%32) - v93) & int32(_a_F_brin_page_type_0)
										v106 = F_errdetail(m, int32(_a_F_brin_page_type_3), v7+int32(16))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_brin_page_type_4), int32(68), int32(_a_F_brin_page_type_5))
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
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
						} else {
							v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v16)+6)))
							v36 = v34 + int32(3951)
							if base.Ui32(int32(3)) <= base.Ui32(v36&int32(_a_F_brin_page_type_0)) {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34
								v43 = F_psprintf(m, int32(_a_F_brin_page_type_6), v7)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int64(0)
								} else {
									v52 = v43
									v53 = F_cstring_to_text(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										v58 = base.I64_extend_i32_u(v53)
										m.G0 = v7 + int32(48)
										return v58
									}
								}
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v36&int32(_a_F_brin_page_type_0)<<(uint(int32(2))%32))+uint32(_c_F_brin_page_type[0])))
								v52 = v51
								v53 = F_cstring_to_text(m, v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									v58 = base.I64_extend_i32_u(v53)
									m.G0 = v7 + int32(48)
									return v58
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_brin_page_type_7), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_brin_page_type_4), int32(54), int32(_a_F_brin_page_type_5))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
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
		}
	}
}
