package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecHashTableDetachBatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v8 < int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = v8 * int32(36)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v14 = v12 + v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	F_sts_end_parallel_scan(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19+v12)+32))
	F_sts_end_parallel_scan(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = v15 + int32(4)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 != int32(3) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v34 == int32(3) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v12)+25)))
	if v31 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+61)) = uint8(v32)
	goto L7
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(-1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = v75 + v79<<(uint(int32(2))%32)
	if base.Ui32(v82) < base.Ui32(v78) {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	v37 = F_BarrierArriveAndDetachExceptLast(m, v25)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v41 = F_BarrierArriveAndDetach(m, v25)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	if v37 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	if v41 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v47 = v45
	goto L21
L19:
	;
	goto L20
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v63 == int32(0) {
		goto L10
	} else {
		goto L26
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v51 = F_dsa_get_address(m, v50, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_dsa_free(m, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v53
	if v53 != 0 {
		v47 = v53
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_dsa_free(m, v66, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
	goto L10
L28:
	;
	v84 = v78
	goto L30
L29:
	;
	v84 = v82
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v84
	goto L1
}
func F__hash_get_totalbuckets(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	if base.Ui32(l0) <= base.Ui32(int32(9)) {
		return int32(1) << (uint(l0) % 32)
	} else {
		v10 = l0 - int32(10)
		v11 = int32(2)
		v13 = int32(512) << (uint(int32(base.Ui32(v10)>>(uint(v11)%32))) % 32)
		return v13>>(uint(v11)%32)*(v10&int32(3)+int32(1)) + v13
	}
}
func F__hash_init_metabuffer(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 float64
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	v4 = l3
	v10 = base.F64_div(l1, base.F64_convert_i32_u(v4))
	if base.F64_le(v10, float64(2)) != 0 {
		v62 = int32(2)
	} else {
		if base.F64_ge(v10, float64(1.073741824e+09)) != 0 {
			v62 = int32(1073741824)
		} else {
			v16 = base.I32_trunc_sat_f64_u(v10)
			v20 = v16 - int32(1)
			if base.Ui32(int32(2)) <= base.Ui32(v16) {
				v26 = int32(32) - base.I32_clz(v20)
			} else {
				v26 = int32(0)
			}
			if base.Ui32(int32(10)) <= base.Ui32(v26) {
				v29 = int32(3)
				v39 = int32(base.Ui32(v20)>>(uint(v26-v29)%32))&v29 | v26<<(uint(int32(2))%32) - int32(30)
			} else {
				v39 = v26
			}
			if base.Ui32(v39) <= base.Ui32(int32(9)) {
				v61 = int32(1) << (uint(v39) % 32)
			} else {
				v47 = v39 - int32(10)
				v48 = int32(2)
				v50 = int32(512) << (uint(int32(base.Ui32(v47)>>(uint(v48)%32))) % 32)
				v61 = v50>>(uint(v48)%32)*(v47&int32(3)+int32(1)) + v50
			}
			v62 = v61
		}
	}
	v66 = v62 - int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v62) {
		v72 = int32(32) - base.I32_clz(v66)
	} else {
		v72 = int32(0)
	}
	if base.Ui32(int32(10)) <= base.Ui32(v72) {
		v75 = int32(3)
		v85 = int32(base.Ui32(v66)>>(uint(v72-v75)%32))&v75 | v72<<(uint(int32(2))%32) - int32(30)
	} else {
		v85 = v72
	}
	if l0 < int32(0) {
		v89 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init_metabuffer[0]))
		v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(l0^int32(-1))<<(uint(int32(2))%32))))
		v103 = v95
	} else {
		v97 = *(*int32)(unsafe.Add(mBase, _c_F__hash_init_metabuffer[1]))
		v103 = v97 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if l4 != 0 {
		v104 = int32(_a_F__hash_init_metabuffer_0)
		v106 = int32(0)
		if v106|(v103&int32(3)|int32(1)) == v106 {
			v122 = v103 + v104
			v124 = v103 + int32(4)
			if base.Ui32(v124) < base.Ui32(v122) {
				v126 = v122
			} else {
				v126 = v124
			}
			v131 = (v103^int32(-1)+v126)&int32(-4) + int32(4)
			if v131 == int32(0) {
			} else {
				base.MemoryFill(m, v103, int32(0), v131)
			}
		} else {
			base.MemoryFill(m, v103, int32(0), v104)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v103)+10)) = int32(_a_F__hash_init_metabuffer_1)
		v145 = int32(_a_F__hash_init_metabuffer_2)
		*(*uint16)(unsafe.Add(mBase, uint32(v103)+18)) = uint16(v145)
		v151 = int32(_a_F__hash_init_metabuffer_3)
		*(*uint16)(unsafe.Add(mBase, uint32(v103)+16)) = uint16(v151)
		*(*uint16)(unsafe.Add(mBase, uint32(v103)+14)) = uint16(v151)
	} else {
	}
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+16)))
	v155 = v103 + v154
	*(*int64)(unsafe.Add(mBase, uint32(v155)+8)) = int64(-36028758364258305)
	*(*int64)(unsafe.Add(mBase, uint32(v155))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v103)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v103)+32)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v103)+24)) = int64(17284990528)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+40)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v103)+72)) = l2
	v168 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v103)+48)) = v62 - v168
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+19)))
	v175 = v171<<(uint(int32(8))%32) - int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+42)) = uint16(v175)
	v177 = int32(-1)
	v180 = v62 + v168
	if v180&v62 != 0 {
		v187 = v177<<(uint(int32(32)-base.I32_clz(v180))%32) ^ v177
	} else {
		v187 = v62
	}
	*(*int32)(unsafe.Add(mBase, uint32(v103)+52)) = v187
	v189 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v103)+56)) = int32(base.Ui32(v187) >> (uint(v189) % 32))
	v196 = base.I32_clz(v175&int32(_a_F__hash_init_metabuffer_4)) ^ int32(31)
	v198 = v196 + int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+46)) = uint16(v198)
	v201 = v189 << (uint(v196) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+44)) = uint16(v201)
	v204 = v103 + int32(76)
	v205 = int32(0)
	base.MemoryFill(m, v204, v205, int32(392))
	base.MemoryFill(m, v103+int32(468), v205, int32(_a_F__hash_init_metabuffer_5))
	*(*int32)(unsafe.Add(mBase, uint32(v204+v85<<(uint(int32(2))%32)))) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v103)+64)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v103)+60)) = v85
	v221 = int32(_a_F__hash_init_metabuffer_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+12)) = uint16(v221)
	return
}
func F__hash_ovflblkno_to_bitno(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v12 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v70 - int32(1)
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v18 = int32(1)
	goto L4
L4:
	;
	if base.Ui32(v18) <= base.Ui32(int32(9)) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L2
L6:
	;
	if base.Ui32(l1) <= base.Ui32(v46) {
		goto L2
	} else {
		goto L10
	}
L7:
	;
	v46 = int32(1) << (uint(v18) % 32)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v32 = v18 - int32(10)
	v33 = int32(2)
	v35 = int32(512) << (uint(int32(base.Ui32(v32)>>(uint(v33)%32))) % 32)
	v46 = v35>>(uint(v33)%32)*(v32&int32(3)+int32(1)) + v35
	goto L6
L10:
	;
	if base.Ui32(v18) <= base.Ui32(int32(9)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v70 = l1 - v69
	v73 = l0 + int32(52) + v18<<(uint(int32(2))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73-int32(4))))
	if base.Ui32(v76) < base.Ui32(v70) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v69 = int32(1) << (uint(v18) % 32)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v55 = v18 - int32(10)
	v56 = int32(2)
	v58 = int32(512) << (uint(int32(base.Ui32(v55)>>(uint(v56)%32))) % 32)
	v69 = v58>>(uint(v56)%32)*(v55&int32(3)+int32(1)) + v58
	goto L11
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if base.Ui32(v70) <= base.Ui32(v78) {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v81 = v18 + int32(1)
	if base.Ui32(v81) <= base.Ui32(v12) {
		v18 = v81
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	goto L5
L20:
	;
	return int32(0)
L21:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg(m, int32(_a_F__hash_ovflblkno_to_bitno_0), v10)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F__hash_ovflblkno_to_bitno_1), int32(88), int32(_a_F__hash_ovflblkno_to_bitno_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashRowType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(711645284)
	v13 = v5 - int32(1636608428) ^ v10 - int32(1455628627)
	v18 = v13 ^ int32(-1636608428) - base.I32_rotl(v13, int32(25))
	v23 = v18 ^ v10 - base.I32_rotl(v18, int32(16))
	v27 = v23 ^ v13 - base.I32_rotl(v23, int32(4))
	v31 = v27 ^ v18 - base.I32_rotl(v27, int32(14))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = int32(711645284)
	v44 = v36 - int32(1636608428) ^ v41 - int32(1455628627)
	v49 = v44 ^ int32(-1636608428) - base.I32_rotl(v44, int32(25))
	v54 = v49 ^ v41 - base.I32_rotl(v49, int32(16))
	v58 = v54 ^ v44 - base.I32_rotl(v54, int32(4))
	v62 = v58 ^ v49 - base.I32_rotl(v58, int32(14))
	v67 = int32(1640531527)
	v68 = v31 ^ v23 - base.I32_rotl(v31, int32(24)) - v67
	v77 = v62 ^ v54 - base.I32_rotl(v62, int32(24)) + v68<<(uint(int32(6))%32) + int32(base.Ui32(v68)>>(uint(int32(2))%32)) - v67 ^ v68
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v78 {
		v82 = v77
		v83 = v78
		v84 = int32(0)
		for {
			v85 = int32(4)
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0+v83<<(uint(v85)%32)+v84*int32(100))+88))
			v96 = int32(711645284)
			v99 = v91 - int32(1636608428) ^ v96 - int32(1455628627)
			v104 = v99 ^ int32(-1636608428) - base.I32_rotl(v99, int32(25))
			v109 = v104 ^ v96 - base.I32_rotl(v104, int32(16))
			v113 = v109 ^ v99 - base.I32_rotl(v109, v85)
			v117 = v113 ^ v104 - base.I32_rotl(v113, int32(14))
			v130 = v117 ^ v109 - base.I32_rotl(v117, int32(24)) + (v82<<(uint(int32(6))%32) + int32(base.Ui32(v82)>>(uint(int32(2))%32))) - int32(1640531527) ^ v82
			v132 = v84 + int32(1)
			v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v132 < v133 {
				v82 = v130
				v83 = v133
				v84 = v132
				continue
			} else {
				break
			}
			break
		}
		v136 = v130
	} else {
		v136 = v77
	}
	return v136
}
func F_hash_metapage_info(m *base.Module, l0 int32) int32 {
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
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v7 = m.G0
	v9 = v7 - int32(_a_F_hash_metapage_info_0)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[0]))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[1]))) = v16
	v20 = F_superuser(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L37
	}
L4:
	;
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = F_verify_hash_page(m, v12, int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L33
	}
L8:
	;
	v28 = F_get_call_result_type(m, l0, int32(0), v9+int32(_a_F_hash_metapage_info_1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v28 != int32(1) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[2])))
	v33 = F_BlessTupleDesc(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[2]))) = v33
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+24)))
	v37 = F_Int64GetDatum(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[3]))) = v37
	v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+28)))
	v41 = F_Int64GetDatum(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[4]))) = v41
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v23)+32))
	v45 = F_Float8GetDatum(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[5]))) = v45
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+40)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[6]))) = v48
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+42)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[7]))) = v50
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[8]))) = v52
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+46)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[9]))) = v54
	v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+48)))
	v57 = F_Int64GetDatum(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[10]))) = v57
	v60 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+52)))
	v61 = F_Int64GetDatum(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[11]))) = v61
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+56)))
	v65 = F_Int64GetDatum(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[12]))) = v65
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+60)))
	v69 = F_Int64GetDatum(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[13]))) = v69
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+64)))
	v73 = F_Int64GetDatum(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[14]))) = v73
	v76 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+68)))
	v77 = F_Int64GetDatum(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[15]))) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[16]))) = v80
	v87 = int32(0)
	goto L21
L21:
	;
	v91 = v87 << (uint(int32(2)) % 32)
	v93 = v9 + int32(_a_F_hash_metapage_info_2)
	v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23+int32(76)+v91))))
	v97 = F_Int64GetDatum(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v106 = F_construct_array_builtin(m, v93, int32(98), int32(20))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v93))) = v97
	v101 = v87 + int32(1)
	if v101 != int32(98) {
		v87 = v101
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[17]))) = v106
	v115 = int32(0)
	goto L26
L26:
	;
	v119 = v115 << (uint(int32(2)) % 32)
	v122 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23+int32(468)+v119))))
	v123 = F_Int64GetDatum(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v132 = F_construct_array_builtin(m, v9, int32(1024), int32(20))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+v119))) = v123
	v127 = v115 + int32(1)
	if v127 != int32(1024) {
		v115 = v127
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[18]))) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_hash_metapage_info[2])))
	v140 = F_heap_form_tuple(m, v135, v9+int32(_a_F_hash_metapage_info_3), v9+int32(_a_F_hash_metapage_info_4))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v143 = F_HeapTupleHeaderGetDatum(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v9 + int32(_a_F_hash_metapage_info_0)
	return v143
L33:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_hash_metapage_info_5), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_hash_metapage_info_6), int32(532), int32(_a_F_hash_metapage_info_7))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_errmsg_internal(m, int32(_a_F_hash_metapage_info_8), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_hash_metapage_info_6), int32(538), int32(_a_F_hash_metapage_info_7))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hash_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
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
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
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
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v377
L2:
	;
	return int32(0)
L3:
	;
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	v17 = int32(_a_F_hash_numeric_0)
	if v16&v17 == v17 {
		v377 = int32(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(-1)
	v22 = base.I32_extend16_s(v16)
	if v22 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v42 = v38 + int32(base.Ui32(v39)>>(uint(int32(2))%32))
	v44 = int32(base.Ui32(v42) >> (uint(int32(1)) % 32))
	if v44 == int32(0) {
		v377 = v21
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v37 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
	v38 = int32(-6)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	v37 = v35
	v38 = int32(-8)
	goto L5
L9:
	;
	v47 = int32(0)
	v49 = v12 + int32(6)
	v51 = v12 + int32(8)
	if v22 < v47 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v54 = v49
	goto L12
L11:
	;
	v54 = v51
	goto L12
L12:
	;
	v55 = v47
	v59 = v37
	goto L13
L13:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54+v55<<(uint(int32(1))%32)))))
	if v68 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v55 == v44 {
		v377 = v21
		goto L1
	} else {
		goto L19
	}
L15:
	;
	v71 = int32(1)
	v74 = v55 + v71
	if v74 != v44 {
		v55 = v74
		v59 = v59 - v71
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	v377 = v21
	goto L1
L19:
	;
	v79 = v44
	v80 = int32(0)
	goto L21
L20:
	;
	if int32(0) <= v22 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v88 = int32(1)
	v89 = v79 - v88
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54+v89<<(uint(v88)%32)))))
	if v93 != 0 {
		v97 = v80
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v97 = v44
	goto L20
L23:
	;
	v95 = v80 + int32(1)
	if v95 != v44 {
		v79 = v89
		v80 = v95
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v102 = v51
	goto L27
L26:
	;
	v102 = v49
	goto L27
L27:
	;
	v103 = v55<<(uint(int32(1))%32) + v102
	v109 = (v42 - (v55+v97)<<(uint(int32(1))%32)) & int32(-2)
	v115 = v109 - int32(1636608432)
	if v103&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v377 = v369 ^ v361 - base.I32_rotl(v369, int32(24)) ^ v59
	goto L1
L29:
	;
	v347 = int32(14)
	v349 = v343 ^ v344 - base.I32_rotl(v343, v347)
	v353 = v349 ^ v342 - base.I32_rotl(v349, int32(11))
	v357 = v353 ^ v343 - base.I32_rotl(v353, int32(25))
	v361 = v357 ^ v349 - base.I32_rotl(v357, int32(16))
	v365 = v361 ^ v353 - base.I32_rotl(v361, int32(4))
	v369 = v365 ^ v357 - base.I32_rotl(v365, v347)
	goto L28
L30:
	;
	switch v273 - int32(1) {
	case 0:
		v335 = v274
		v336 = v275
		v337 = v276
		goto L57
	case 1:
		v328 = v274
		v329 = v275
		v330 = v276
		goto L58
	case 2:
		v321 = v274
		v322 = v275
		v323 = v276
		goto L59
	case 3:
		v315 = v275
		v316 = v276
		goto L60
	case 4:
		v311 = v275
		v312 = v276
		goto L61
	case 5:
		v305 = v275
		v306 = v276
		goto L62
	case 6:
		v299 = v275
		v300 = v276
		goto L63
	case 7:
		v294 = v276
		goto L64
	case 8:
		v289 = v276
		goto L65
	case 9:
		v284 = v276
		goto L66
	case 10:
		goto L67
	default:
		v342 = v274
		v343 = v275
		v344 = v276
		goto L29
	}
L31:
	;
	v224 = v103
	v225 = v109
	v226 = v115
	v227 = v115
	v228 = v115
	goto L54
L32:
	;
	if base.Ui32(int32(11)) < base.Ui32(v109) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v109) < base.Ui32(int32(12)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v272 = v103
	v273 = v109
	v274 = v115
	v275 = v115
	v276 = v115
	goto L30
L36:
	;
	switch v171 - int32(1) {
	case 0:
		v221 = v172
		goto L43
	case 1:
		v216 = v172
		goto L44
	case 2:
		goto L45
	case 3:
		v209 = v173
		goto L46
	case 4:
		v206 = v173
		goto L47
	case 5:
		v201 = v173
		goto L48
	case 6:
		goto L49
	case 7:
		v192 = v174
		goto L50
	case 8:
		v187 = v174
		goto L51
	case 9:
		v182 = v174
		goto L52
	case 10:
		goto L53
	default:
		v342 = v172
		v343 = v173
		v344 = v174
		goto L29
	}
L37:
	;
	v170 = v103
	v171 = v109
	v172 = v115
	v173 = v115
	v174 = v115
	goto L36
L38:
	;
	goto L39
L39:
	;
	v122 = v103
	v123 = v109
	v124 = v115
	v125 = v115
	v126 = v115
	goto L40
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v129 = v128 + v125
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v133 = v132 + v126
	v135 = int32(4)
	v137 = v130 + v124 - v133 ^ base.I32_rotl(v133, v135)
	v141 = v129 - v137 ^ base.I32_rotl(v137, int32(6))
	v142 = v133 + v129
	v143 = v137 + v142
	v144 = v141 + v143
	v148 = v142 - v141 ^ base.I32_rotl(v141, int32(8))
	v152 = v143 - v148 ^ base.I32_rotl(v148, int32(16))
	v156 = v144 - v152 ^ base.I32_rotl(v152, int32(19))
	v157 = v148 + v144
	v158 = v152 + v157
	v159 = v156 + v158
	v163 = v157 - v156 ^ base.I32_rotl(v156, v135)
	v164 = int32(12)
	v165 = v122 + v164
	v167 = v123 - v164
	if base.Ui32(int32(11)) < base.Ui32(v167) {
		v122 = v165
		v123 = v167
		v124 = v158
		v125 = v159
		v126 = v163
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v170 = v165
	v171 = v167
	v172 = v158
	v173 = v159
	v174 = v163
	goto L36
L42:
	;
	goto L41
L43:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v342 = v221 + v222
	v343 = v173
	v344 = v174
	goto L29
L44:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
	v221 = v217<<(uint(int32(8))%32) + v216
	goto L43
L45:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
	v216 = v212<<(uint(int32(16))%32) + v172
	goto L44
L46:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v342 = v210 + v172
	v343 = v209
	v344 = v174
	goto L29
L47:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
	v209 = v206 + v207
	goto L46
L48:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
	v206 = v202<<(uint(int32(8))%32) + v201
	goto L47
L49:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
	v201 = v197<<(uint(int32(16))%32) + v173
	goto L48
L50:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v342 = v193 + v172
	v343 = v195 + v173
	v344 = v192
	goto L29
L51:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
	v192 = v188<<(uint(int32(8))%32) + v187
	goto L50
L52:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+9)))
	v187 = v183<<(uint(int32(16))%32) + v182
	goto L51
L53:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+10)))
	v182 = v178<<(uint(int32(24))%32) + v174
	goto L52
L54:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v231 = v230 + v227
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v235 = v234 + v228
	v237 = int32(4)
	v239 = v232 + v226 - v235 ^ base.I32_rotl(v235, v237)
	v243 = v231 - v239 ^ base.I32_rotl(v239, int32(6))
	v244 = v235 + v231
	v245 = v239 + v244
	v246 = v243 + v245
	v250 = v244 - v243 ^ base.I32_rotl(v243, int32(8))
	v254 = v245 - v250 ^ base.I32_rotl(v250, int32(16))
	v258 = v246 - v254 ^ base.I32_rotl(v254, int32(19))
	v259 = v250 + v246
	v260 = v254 + v259
	v261 = v258 + v260
	v265 = v259 - v258 ^ base.I32_rotl(v258, v237)
	v266 = int32(12)
	v267 = v224 + v266
	v269 = v225 - v266
	if base.Ui32(int32(11)) < base.Ui32(v269) {
		v224 = v267
		v225 = v269
		v226 = v260
		v227 = v261
		v228 = v265
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v272 = v267
	v273 = v269
	v274 = v260
	v275 = v261
	v276 = v265
	goto L30
L56:
	;
	goto L55
L57:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v342 = v335 + v338
	v343 = v336
	v344 = v337
	goto L29
L58:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v335 = v331<<(uint(int32(8))%32) + v328
	v336 = v329
	v337 = v330
	goto L57
L59:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+2)))
	v328 = v324<<(uint(int32(16))%32) + v321
	v329 = v322
	v330 = v323
	goto L58
L60:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+3)))
	v321 = v317<<(uint(int32(24))%32) + v274
	v322 = v315
	v323 = v316
	goto L59
L61:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+4)))
	v315 = v311 + v313
	v316 = v312
	goto L60
L62:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+5)))
	v311 = v307<<(uint(int32(8))%32) + v305
	v312 = v306
	goto L61
L63:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+6)))
	v305 = v301<<(uint(int32(16))%32) + v299
	v306 = v300
	goto L62
L64:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+7)))
	v299 = v295<<(uint(int32(24))%32) + v275
	v300 = v294
	goto L63
L65:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+8)))
	v294 = v290<<(uint(int32(8))%32) + v289
	goto L64
L66:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+9)))
	v289 = v285<<(uint(int32(16))%32) + v284
	goto L65
L67:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+10)))
	v284 = v280<<(uint(int32(24))%32) + v276
	goto L66
}
func F_hash_page_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
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
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v25)
		*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = int64(0)
		v29 = F_superuser(m)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v32 = F_verify_hash_page(m, v21, int32(3))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+16)))
					v35 = v32 + v34
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
					if base.Ui32(v36) < base.Ui32(int32(25)) {
						v139 = v2
						v140 = v2
					} else {
						v42 = int32(base.Ui32(v36+int32(_a_F_hash_page_stats_0)) >> (uint(int32(2)) % 32))
						v44 = v42 & int32(_a_F_hash_page_stats_1)
						if v44 == int32(0) {
							v139 = v2
							v140 = v2
						} else {
							v48 = v32 + int32(20)
							v49 = int32(1)
							if v44 != v49 {
								v59 = v49
								v62 = int32(0)
								v63 = v2
								v64 = v2
								for {
									v74 = v48 + v59<<(uint(int32(2))%32)
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
									v76 = int32(_a_F_hash_page_stats_2)
									if v75&v76 != v76 {
										v84 = v63 + int32(1)
										v85 = v64
									} else {
										v84 = v63
										v85 = v64 + int32(1)
									}
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
									v87 = int32(_a_F_hash_page_stats_2)
									if v86&v87 != v87 {
										v95 = v84 + int32(1)
										v96 = v85
									} else {
										v95 = v84
										v96 = v85 + int32(1)
									}
									v97 = int32(2)
									v98 = v59 + v97
									v100 = v62 + v97
									if v100 != v42&int32(_a_F_hash_page_stats_3) {
										v59 = v98
										v62 = v100
										v63 = v95
										v64 = v96
										continue
									} else {
										break
									}
									break
								}
								if v42&int32(1) == int32(0) {
									v139 = v95
									v140 = v96
								} else {
									v106 = v98
									v110 = v95
									v111 = v96
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v48+v106<<(uint(int32(2))%32))))
									v125 = int32(_a_F_hash_page_stats_2)
									v128 = base.B2i32(v124&v125 == v125)
									if v124&v125 == v125 {
										v129 = v110
									} else {
										v129 = v110 + int32(1)
									}
									if v124&v125 == v125 {
										v132 = v111 + int32(1)
									} else {
										v132 = v111
									}
									v139 = v129
									v140 = v132
								}
							} else {
								v106 = v49
								v110 = v2
								v111 = v2
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v48+v106<<(uint(int32(2))%32))))
								v125 = int32(_a_F_hash_page_stats_2)
								v128 = base.B2i32(v124&v125 == v125)
								if v124&v125 == v125 {
									v129 = v110
								} else {
									v129 = v110 + int32(1)
								}
								if v124&v125 == v125 {
									v132 = v111 + int32(1)
								} else {
									v132 = v111
								}
								v139 = v129
								v140 = v132
							}
						}
					}
					v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+14)))
					v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
					v150 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+8)))
					v151 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+4)))
					v152 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35))))
					v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+18)))
					v154 = int32(4)
					v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
					v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
					v157 = v155 - v156
					if v157 <= v154 {
						v160 = v154
					} else {
						v160 = v157
					}
					v166 = F_get_call_result_type(m, l0, int32(0), v18+int32(12))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						if v166 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_hash_page_stats_4), int32(0))
								mBase = m.M
								v228 = m.ExcPending
								if v228 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_page_stats_5), int32(263), int32(_a_F_hash_page_stats_6))
									mBase = m.M
									v233 = m.ExcPending
									if v233 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
							v171 = F_BlessTupleDesc(m, v170)
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v160 - int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v153 & int32(_a_F_hash_page_stats_7)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v140
								*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v139
								*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v171
								v180 = F_Int64GetDatum(m, v152)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v180
									v183 = F_Int64GetDatum(m, v151)
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v183
										v186 = F_Int64GetDatum(m, v150)
										mBase = m.M
										v187 = m.ExcPending
										if v187 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v148
											*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v149
											*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v186
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
											v196 = F_heap_form_tuple(m, v191, v18+int32(32), v18+int32(16))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
												v199 = F_HeapTupleHeaderGetDatum(m, v198)
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(80)
													return v199
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v211 = m.ExcPending
					if v211 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_hash_page_stats_8), int32(0))
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hash_page_stats_5), int32(251), int32(_a_F_hash_page_stats_6))
							mBase = m.M
							v220 = m.ExcPending
							if v220 != 0 {
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
func F_hash_range(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
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
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if v20 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v21 == v18 {
					v31 = v20
					F_range_deserialize(m, v31, v12, v9+int32(40), v9+int32(32), v9+int32(31))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v40)>>(uint(int32(2))%32))-int32(1)))))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+200))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
						if v48 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
							v53 = F_lookup_type_cache(m, v51, int32(128))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
								if v55 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(52461700))
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
											v141 = F_format_type_be(m, v140)
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v141
												F_errmsg(m, int32(_a_F_hash_range_0), v9+int32(16))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_hash_range_1), int32(1426), int32(_a_F_hash_range_2))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
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
									v58 = v53
									v59 = int32(0)
									if v46&int32(41) == v59 {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
										v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = v68
											if v46&int32(81) == int32(0) {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = v79
													v86 = int32(711645284)
													v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
													v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
													v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
													v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
													v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
													m.G0 = v9 + int32(48)
													return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
												}
											} else {
												v81 = v59
												v86 = int32(711645284)
												v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
												v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
												v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
												v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
												v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
												m.G0 = v9 + int32(48)
												return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
											}
										}
									} else {
										v70 = int32(0)
										if v46&int32(81) == int32(0) {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v79
												v86 = int32(711645284)
												v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
												v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
												v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
												v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
												v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
												m.G0 = v9 + int32(48)
												return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
											}
										} else {
											v81 = v59
											v86 = int32(711645284)
											v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
											v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
											v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
											v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
											v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
											m.G0 = v9 + int32(48)
											return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
										}
									}
								}
							}
						} else {
							v58 = v47
							v59 = int32(0)
							if v46&int32(41) == v59 {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
								v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = v68
									if v46&int32(81) == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = v79
											v86 = int32(711645284)
											v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
											v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
											v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
											v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
											v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
											m.G0 = v9 + int32(48)
											return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
										}
									} else {
										v81 = v59
										v86 = int32(711645284)
										v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
										v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
										v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
										v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
										v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
										m.G0 = v9 + int32(48)
										return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
									}
								}
							} else {
								v70 = int32(0)
								if v46&int32(81) == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = v79
										v86 = int32(711645284)
										v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
										v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
										v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
										v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
										v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
										m.G0 = v9 + int32(48)
										return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
									}
								} else {
									v81 = v59
									v86 = int32(711645284)
									v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
									v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
									v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
									v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
									v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
									m.G0 = v9 + int32(48)
									return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
								}
							}
						}
					}
				} else {
					v24 = F_lookup_type_cache(m, v18, int32(2048))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
								F_errmsg_internal(m, int32(_a_F_hash_range_3), v9)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_range_1), int32(1776), int32(_a_F_hash_range_4))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
							v31 = v24
							F_range_deserialize(m, v31, v12, v9+int32(40), v9+int32(32), v9+int32(31))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v40)>>(uint(int32(2))%32))-int32(1)))))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+200))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
								if v48 == int32(0) {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
									v53 = F_lookup_type_cache(m, v51, int32(128))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
										if v55 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(52461700))
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													v140 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
													v141 = F_format_type_be(m, v140)
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v141
														F_errmsg(m, int32(_a_F_hash_range_0), v9+int32(16))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_hash_range_1), int32(1426), int32(_a_F_hash_range_2))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
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
											v58 = v53
											v59 = int32(0)
											if v46&int32(41) == v59 {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
												v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v70 = v68
													if v46&int32(81) == int32(0) {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
														v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v81 = v79
															v86 = int32(711645284)
															v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
															v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
															v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
															v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
															v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
															m.G0 = v9 + int32(48)
															return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
														}
													} else {
														v81 = v59
														v86 = int32(711645284)
														v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
														v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
														v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
														v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
														v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
														m.G0 = v9 + int32(48)
														return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
													}
												}
											} else {
												v70 = int32(0)
												if v46&int32(81) == int32(0) {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
													v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v81 = v79
														v86 = int32(711645284)
														v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
														v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
														v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
														v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
														v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
														m.G0 = v9 + int32(48)
														return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
													}
												} else {
													v81 = v59
													v86 = int32(711645284)
													v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
													v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
													v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
													v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
													v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
													m.G0 = v9 + int32(48)
													return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
												}
											}
										}
									}
								} else {
									v58 = v47
									v59 = int32(0)
									if v46&int32(41) == v59 {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
										v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = v68
											if v46&int32(81) == int32(0) {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = v79
													v86 = int32(711645284)
													v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
													v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
													v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
													v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
													v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
													m.G0 = v9 + int32(48)
													return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
												}
											} else {
												v81 = v59
												v86 = int32(711645284)
												v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
												v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
												v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
												v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
												v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
												m.G0 = v9 + int32(48)
												return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
											}
										}
									} else {
										v70 = int32(0)
										if v46&int32(81) == int32(0) {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v79
												v86 = int32(711645284)
												v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
												v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
												v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
												v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
												v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
												m.G0 = v9 + int32(48)
												return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
											}
										} else {
											v81 = v59
											v86 = int32(711645284)
											v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
											v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
											v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
											v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
											v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
											m.G0 = v9 + int32(48)
											return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
										}
									}
								}
							}
						}
					}
				}
			} else {
				v24 = F_lookup_type_cache(m, v18, int32(2048))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
							F_errmsg_internal(m, int32(_a_F_hash_range_3), v9)
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_hash_range_1), int32(1776), int32(_a_F_hash_range_4))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
						v31 = v24
						F_range_deserialize(m, v31, v12, v9+int32(40), v9+int32(32), v9+int32(31))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v40)>>(uint(int32(2))%32))-int32(1)))))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+200))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
							if v48 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								v53 = F_lookup_type_cache(m, v51, int32(128))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
									if v55 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(52461700))
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return int32(0)
											} else {
												v140 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
												v141 = F_format_type_be(m, v140)
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v141
													F_errmsg(m, int32(_a_F_hash_range_0), v9+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_hash_range_1), int32(1426), int32(_a_F_hash_range_2))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
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
										v58 = v53
										v59 = int32(0)
										if v46&int32(41) == v59 {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
											v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = v68
												if v46&int32(81) == int32(0) {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
													v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v81 = v79
														v86 = int32(711645284)
														v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
														v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
														v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
														v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
														v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
														m.G0 = v9 + int32(48)
														return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
													}
												} else {
													v81 = v59
													v86 = int32(711645284)
													v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
													v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
													v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
													v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
													v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
													m.G0 = v9 + int32(48)
													return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
												}
											}
										} else {
											v70 = int32(0)
											if v46&int32(81) == int32(0) {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													v81 = v79
													v86 = int32(711645284)
													v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
													v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
													v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
													v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
													v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
													m.G0 = v9 + int32(48)
													return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
												}
											} else {
												v81 = v59
												v86 = int32(711645284)
												v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
												v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
												v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
												v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
												v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
												m.G0 = v9 + int32(48)
												return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
											}
										}
									}
								}
							} else {
								v58 = v47
								v59 = int32(0)
								if v46&int32(41) == v59 {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
									v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = v68
										if v46&int32(81) == int32(0) {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = v79
												v86 = int32(711645284)
												v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
												v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
												v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
												v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
												v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
												m.G0 = v9 + int32(48)
												return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
											}
										} else {
											v81 = v59
											v86 = int32(711645284)
											v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
											v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
											v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
											v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
											v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
											m.G0 = v9 + int32(48)
											return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
										}
									}
								} else {
									v70 = int32(0)
									if v46&int32(81) == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v79 = F_FunctionCall1Coll(m, v58+int32(132), v77, v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = v79
											v86 = int32(711645284)
											v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
											v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
											v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
											v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
											v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
											m.G0 = v9 + int32(48)
											return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
										}
									} else {
										v81 = v59
										v86 = int32(711645284)
										v89 = v46 - int32(1636608428) ^ v86 - int32(1455628627)
										v94 = v89 ^ int32(-1636608428) - base.I32_rotl(v89, int32(25))
										v99 = v94 ^ v86 - base.I32_rotl(v94, int32(16))
										v103 = v99 ^ v89 - base.I32_rotl(v99, int32(4))
										v107 = v103 ^ v94 - base.I32_rotl(v103, int32(14))
										m.G0 = v9 + int32(48)
										return base.I32_rotl(v107^v99-base.I32_rotl(v107, int32(24))^v70, int32(1)) ^ v81
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
func F_hash_range_extended(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	F_range_deserialize(m, v33, v13, v10+int32(40), v10+int32(32), v10+int32(31))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 == v20 {
		v33 = v22
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v26 = F_lookup_type_cache(m, v20, int32(2048))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
	if v28 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
	v33 = v26
	goto L6
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13+int32(base.Ui32(v42)>>(uint(int32(2))%32))-int32(1)))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+200))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+164))
	if v50 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v55 = F_lookup_type_cache(m, v53, int32(_a_F_hash_range_extended_0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v60 = v49
	goto L16
L16:
	;
	if v48&int32(41) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+164))
	if v57 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v60 = v55
	goto L16
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v69 = F_FunctionCall2Coll(m, v60+int32(160), v67, v68, v17)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v72 = int64(0)
	goto L21
L21:
	;
	if v48&int32(81) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	v72 = v71
	goto L21
L23:
	;
	v83 = int64(0)
	goto L25
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v80 = F_FunctionCall2Coll(m, v60+int32(160), v78, v79, v17)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if v84 == int64(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	v83 = v82
	goto L25
L27:
	;
	v168 = F_Int64GetDatum(m, base.I64_extend_i32_u(v158)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v158^v150-base.I32_rotl(v158, int32(24))))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v135 = int32(14)
	v137 = v134 - base.I32_rotl(v130, v135)
	v142 = v137 ^ (v48 + v131) - base.I32_rotl(v137, int32(11))
	v146 = v130 ^ v142 - base.I32_rotl(v142, int32(25))
	v150 = v146 ^ v137 - base.I32_rotl(v146, int32(16))
	v154 = v150 ^ v142 - base.I32_rotl(v150, int32(4))
	v158 = v154 ^ v146 - base.I32_rotl(v154, v135)
	goto L27
L29:
	;
	v91 = int32(-1636608428)
	v130 = v91
	v131 = v91
	v134 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v94 = base.I32_wrap_i64(v84)
	v96 = v94 + int32(1021750440)
	v101 = base.I32_wrap_i64(int64(base.Ui64(v84)>>(uint(int64(32))%64))) ^ int32(-415931063)
	v107 = v94 - v101 - int32(1636608428) ^ base.I32_rotl(v101, int32(6))
	v111 = v96 - v107 ^ base.I32_rotl(v107, int32(8))
	v112 = v101 + v96
	v113 = v107 + v112
	v114 = v111 + v113
	v118 = v112 - v111 ^ base.I32_rotl(v111, int32(16))
	v122 = v113 - v118 ^ base.I32_rotl(v118, int32(19))
	v127 = v118 + v114
	v128 = v122 + v127
	v130 = v128
	v131 = v127
	v134 = v114 - v122 ^ base.I32_rotl(v122, int32(4)) ^ v128
	goto L28
L32:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	v171 = v170 ^ v72
	v182 = F_Int64GetDatum(m, v83^(v171<<(uint(int64(1))%64)&int64(-4294967298)|int64(base.Ui64(v171)>>(uint(int64(31))%64))&int64(4294967297)))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v10 + int32(48)
	return v182
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
	F_errmsg_internal(m, int32(_a_F_hash_range_extended_4), v10)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1776), int32(_a_F_hash_range_extended_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v209 = F_format_type_be(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v209
	F_errmsg(m, int32(_a_F_hash_range_extended_1), v10+int32(16))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1490), int32(_a_F_hash_range_extended_3))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hash_record(m *base.Module, l0 int32) int32 {
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
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
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
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v29 = F_lookup_rowtype_tupdesc(m, v27, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v21
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v34
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+52)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = int32(base.Ui32(v32) >> (uint(int32(2)) % 32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v44 == v34 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v65 != v27 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v55 = F_MemoryContextAlloc(m, v50, v31<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v47 < v31 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v64 = v44
	v65 = v49
	goto L5
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v60)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v31
	v64 = v60
	v65 = int32(0)
	goto L5
L10:
	;
	v116 = F_palloc(m, v110)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v72 = v64 + int32(20)
	v76 = v31 << (uint(int32(2)) % 32)
	if v72&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v76)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v67 != v28 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v110 = v31 << (uint(int32(2)) % 32)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v27
	v110 = v76
	goto L10
L15:
	;
	if v76 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v76 == int32(0) {
		goto L14
	} else {
		goto L23
	}
L18:
	;
	v86 = v64 + v76 + int32(20)
	v88 = v64 + int32(24)
	if base.Ui32(v88) < base.Ui32(v86) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v90 = v86
	goto L21
L20:
	;
	v90 = v88
	goto L21
L21:
	;
	v97 = (v90-v64-int32(21))&int32(-4) + int32(4)
	if v97 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	base.MemoryFill(m, v72, int32(0), v97)
	goto L14
L23:
	;
	base.MemoryFill(m, v72, int32(0), v76)
	goto L14
L24:
	;
	v118 = F_palloc(m, v31)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_heap_deform_tuple(m, v16+int32(-20), v29, v116, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v31 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_pfree(m, v116)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L55
	}
L28:
	;
	v238 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v127 = int32(0)
	v130 = v127
	v134 = v127
	goto L31
L31:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v150 = v29 + v144<<(uint(int32(4))%32) + v130*int32(100)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+111)))
	if v151 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L50
	}
L33:
	;
	goto L32
L34:
	;
	v155 = v150 + int32(20)
	v157 = v130 << (uint(int32(2)) % 32)
	v158 = v64 + int32(20) + v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v207 = v134
	goto L36
L36:
	;
	v212 = v130 + int32(1)
	if v31 != v212 {
		v130 = v212
		v134 = v207
		goto L31
	} else {
		goto L49
	}
L37:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v118))))
	if v177 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v168 = F_lookup_type_cache(m, v166, int32(128))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L43
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	v166 = v162
	goto L38
L40:
	;
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v163 == v164 {
		v174 = v159
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v166 = v163
	goto L38
L43:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v168)+136))
	if v170 == int32(0) {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v168
	v174 = v168
	goto L37
L45:
	;
	v201 = int32(0)
	goto L47
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v174 + int32(132)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v155)+96))
	v185 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+34)) = uint16(v185)
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v184
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v116+v157)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+40)) = uint8(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v191
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v174)+132))
	v198 = m.T0[v197].(func(*base.Module, int32) int32)(m, v16+int32(-48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v207 = v201 + v134*int32(31)
	goto L36
L48:
	;
	v201 = v198
	goto L47
L49:
	;
	v238 = v207
	goto L27
L50:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v222 = F_format_type_be(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v222
	F_errmsg(m, int32(_a_F_hash_record_0), v18)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_hash_record_1), int32(1894), int32(_a_F_hash_record_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_pfree(m, v118)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if int32(0) <= v252 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_DecrTupleDescRefCount(m, v29)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v257 != v21 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	F_pfree(m, v21)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	m.G0 = v18 - int32(-64)
	return v238
L64:
	;
	goto L63
}
func F_initialize_hash_entry(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 float64
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	v13 = v11 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+320)) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	goto L3
L1:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	goto L14
L2:
	;
	goto L1
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v22 == int32(0) {
		v41 = v19
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v27 = v19
	v28 = v22
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v30 = v29 + v27
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v31 != 0 {
		v27 = v30
		v28 = v31
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v41 = v30
	goto L2
L7:
	;
	v33 = v28
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	if v36 != 0 {
		v27 = v30
		v28 = v36
		goto L5
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v37 != v15 {
		v33 = v37
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	goto L25
L13:
	;
	goto L12
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v50 == int32(0) {
		v69 = v47
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v55 = v47
	v56 = v50
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = v57 + v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v59 != 0 {
		v55 = v58
		v56 = v59
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v69 = v58
	goto L13
L18:
	;
	v61 = v56
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	if v64 != 0 {
		v55 = v58
		v56 = v64
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v65 != v43 {
		v61 = v65
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v100 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	goto L23
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	if v79 == int32(0) {
		v98 = v76
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v84 = v76
	v85 = v79
	goto L27
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v87 = v86 + v84
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	if v88 != 0 {
		v84 = v87
		v85 = v88
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v98 = v87
	goto L24
L29:
	;
	v90 = v85
	goto L30
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	if v93 != 0 {
		v84 = v87
		v85 = v93
		goto L27
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v94 != v72 {
		v90 = v94
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v207 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if base.Ui32(v41+v69+v98) <= base.Ui32(v103) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(l0)+288))
	if base.Ui64(v13) <= base.Ui64(v107) {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v114 != int32(2) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v117 = int32(48)
	goto L42
L41:
	;
	v117 = int32(0)
	goto L42
L42:
	;
	v118 = v111 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)))
	v122 = v118 + v119<<(uint(int32(3))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	if v123 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v119 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v145 = v123
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+28)) = v145
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+276)))
	if v148 != 0 {
		goto L34
	} else {
		goto L51
	}
L46:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(_a_F_initialize_hash_entry_0)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v132 = int32(1)
	v139 = F_ExecBuildAggTrans(m, l0, v118, (v119^v132)&base.B2i32(v114 == int32(3)), v132, v132)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+36)) = v139
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v126)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v127
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	v145 = v144
	goto L45
L51:
	;
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+276)) = uint8(v149)
	v154 = F_LogicalTapeSetCreate(m, v149, int32(0), int32(-1))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v160 = F_palloc(m, v157*int32(24))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v163 <= int32(0) {
		goto L34
	} else {
		goto L54
	}
L54:
	;
	v170 = int32(0)
	goto L55
L55:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v170*int32(52))+48))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+96))
	v190 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v177+v170*int32(24), v181, int32(0), base.F64_convert_i32_s(v188), v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L34
L57:
	;
	v194 = v170 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v194 < v195 {
		v170 = v194
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	return
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v210 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v214 = v211 - v210
	goto L63
L62:
	;
	v214 = int32(0)
	goto L63
L63:
	;
	if v207 <= int32(0) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v221 = int32(0)
	goto L65
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_initialize_aggregate(m, l0, v228+v221*int32(224), v214+v221<<(uint(int32(3))%32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L49
	} else {
		goto L67
	}
L66:
	;
	goto L59
L67:
	;
	v238 = v221 + int32(1)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v238 < v239 {
		v221 = v238
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
}
func F_verify_hash_page(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v10 = F_get_page_from_raw(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)))
		if v14 != 0 {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)))
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
			if (v15<<(uint(int32(8))%32)-v18)&int32(_a_F_verify_hash_page_0) != int32(16) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = int32(_a_F_verify_hash_page_1)
						F_errmsg(m, int32(_a_F_verify_hash_page_2), v8+int32(112))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = (v127<<(uint(int32(8))%32) - v126) & int32(_a_F_verify_hash_page_0)
							F_errdetail(m, int32(_a_F_verify_hash_page_3), v8+int32(96))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(75), int32(_a_F_verify_hash_page_5))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
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
				v24 = v10 + v18
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)))
				if v25 != int32(_a_F_verify_hash_page_6) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = int32(_a_F_verify_hash_page_1)
							F_errmsg(m, int32(_a_F_verify_hash_page_2), v8+int32(80))
							mBase = m.M
							v159 = m.ExcPending
							if v159 != 0 {
								return int32(0)
							} else {
								v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v160
								*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = int32(_a_F_verify_hash_page_6)
								F_errdetail(m, int32(_a_F_verify_hash_page_7), v8-int32(-64))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(83), int32(_a_F_verify_hash_page_5))
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
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
					v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)))
					v30 = v28 & int32(15)
					v33 = int32(0)
					if base.B2i32(base.B2i32(v28&int32(7) == v33)|base.B2i32(base.Ui32(v30-int32(1)) < base.Ui32(int32(2))) == v33)&base.B2i32(v30 != int32(4)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v30
								F_errmsg(m, int32(_a_F_verify_hash_page_8), v8+int32(48))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(94), int32(_a_F_verify_hash_page_5))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
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
						if l1&v30 != 0 {
							v47 = int32(0)
						} else {
							v47 = l1
						}
						if v47 != 0 {
							v58 = v30
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								switch l1 - int32(1) {
								case 0:
									F_errcode(m, int32(50856066))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_verify_hash_page_9), int32(0))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(114), int32(_a_F_verify_hash_page_5))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								default:
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v58
									F_errmsg_internal(m, int32(_a_F_verify_hash_page_10), v8)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(119), int32(_a_F_verify_hash_page_5))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								case 2:
									F_errcode(m, int32(50856066))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_verify_hash_page_11), int32(0))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(109), int32(_a_F_verify_hash_page_5))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								case 7:
									F_errcode(m, int32(50856066))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_verify_hash_page_12), int32(0))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(104), int32(_a_F_verify_hash_page_5))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
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
							if v30 != int32(8) {
								m.G0 = v8 + int32(128)
								return v10
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								if v50 != int32(105121344) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_verify_hash_page_13), int32(0))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												v203 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v203
												*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(105121344)
												F_errdetail(m, int32(_a_F_verify_hash_page_14), v8+int32(32))
												mBase = m.M
												v211 = m.ExcPending
												if v211 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(136), int32(_a_F_verify_hash_page_5))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
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
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
									if v53 != int32(4) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v220 = m.ExcPending
										if v220 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v223 = m.ExcPending
											if v223 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_verify_hash_page_15), int32(0))
												mBase = m.M
												v227 = m.ExcPending
												if v227 != 0 {
													return int32(0)
												} else {
													v228 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
													*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v228
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(4)
													F_errdetail(m, int32(_a_F_verify_hash_page_16), v8+int32(16))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(143), int32(_a_F_verify_hash_page_5))
														mBase = m.M
														v241 = m.ExcPending
														if v241 != 0 {
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
										m.G0 = v8 + int32(128)
										return v10
									}
								}
							}
						}
					}
				}
			}
		} else {
			if l1 == int32(0) {
				m.G0 = v8 + int32(128)
				return v10
			} else {
				v58 = int32(0)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					switch l1 - int32(1) {
					case 0:
						F_errcode(m, int32(50856066))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_verify_hash_page_9), int32(0))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(114), int32(_a_F_verify_hash_page_5))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					default:
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v58
						F_errmsg_internal(m, int32(_a_F_verify_hash_page_10), v8)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(119), int32(_a_F_verify_hash_page_5))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					case 2:
						F_errcode(m, int32(50856066))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_verify_hash_page_11), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(109), int32(_a_F_verify_hash_page_5))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 7:
						F_errcode(m, int32(50856066))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_verify_hash_page_12), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_hash_page_4), int32(104), int32(_a_F_verify_hash_page_5))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
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
