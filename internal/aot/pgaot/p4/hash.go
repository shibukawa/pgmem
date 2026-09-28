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
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0+v83<<(uint(int32(3))%32)+v84*int32(100))+96))
			v96 = int32(711645284)
			v99 = v91 - int32(1636608428) ^ v96 - int32(1455628627)
			v104 = v99 ^ int32(-1636608428) - base.I32_rotl(v99, int32(25))
			v109 = v104 ^ v96 - base.I32_rotl(v104, int32(16))
			v113 = v109 ^ v99 - base.I32_rotl(v109, int32(4))
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
func F_hash_metapage_info(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v150 int64
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v8 = m.G0
	v10 = v8 - int32(_a_F_hash_metapage_info_0)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[0]))) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[1]))) = v17
		v21 = F_superuser(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			if v21 != 0 {
				v24 = F_verify_hash_page(m, v13, int32(8))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v29 = F_get_call_result_type(m, l0, int32(0), v10+int32(_a_F_hash_metapage_info_1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						if v29 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_hash_metapage_info_2), int32(0))
								mBase = m.M
								v199 = m.ExcPending
								if v199 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_metapage_info_3), int32(542), int32(_a_F_hash_metapage_info_4))
									mBase = m.M
									v204 = m.ExcPending
									if v204 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[2])))
							v34 = F_BlessTupleDesc(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[2]))) = v34
								v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+24)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[3]))) = v37
								v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+28)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[4]))) = v39
								v41 = *(*int64)(unsafe.Add(mBase, uint32(v24)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[5]))) = v41
								v43 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v24)+40)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[6]))) = v43
								v45 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v24)+42)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[7]))) = v45
								v47 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v24)+44)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[8]))) = v47
								v49 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v24)+46)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[9]))) = v49
								v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+48)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[10]))) = v51
								v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+52)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[11]))) = v53
								v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+56)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[12]))) = v55
								v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+60)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[13]))) = v57
								v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+64)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[14]))) = v59
								v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+68)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[15]))) = v61
								v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+72)))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[16]))) = v63
								v66 = v24 + int32(76)
								v70 = int32(0)
								for {
									v75 = v10 - int32(-8192)
									v76 = int32(3)
									v79 = int32(2)
									v82 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v70<<(uint(v79)%32)))))
									*(*int64)(unsafe.Add(mBase, uint32(v75+v70<<(uint(v76)%32)))) = v82
									v85 = v70 | int32(1)
									v92 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v85<<(uint(v79)%32)))))
									*(*int64)(unsafe.Add(mBase, uint32(v85<<(uint(v76)%32)+v75))) = v92
									v95 = v70 + v79
									if v95 != int32(98) {
										v70 = v95
										continue
									} else {
										break
									}
									break
								}
								v100 = F_construct_array_builtin(m, v75, int32(98), int32(20))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[17]))) = base.I64_extend_i32_u(v100)
									v105 = v24 + int32(468)
									v110 = int32(0)
									for {
										v114 = int32(3)
										v117 = int32(2)
										v120 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v105+v110<<(uint(v117)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v10+v110<<(uint(v114)%32)))) = v120
										v123 = v110 | int32(1)
										v130 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v105+v123<<(uint(v117)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v10+v123<<(uint(v114)%32)))) = v130
										v133 = v110 | v117
										v140 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v105+v133<<(uint(v117)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v10+v133<<(uint(v114)%32)))) = v140
										v143 = v110 | v114
										v150 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v105+v143<<(uint(v117)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v10+v143<<(uint(v114)%32)))) = v150
										v153 = v110 + int32(4)
										if v153 != int32(1024) {
											v110 = v153
											continue
										} else {
											break
										}
										break
									}
									v158 = F_construct_array_builtin(m, v10, int32(1024), int32(20))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[18]))) = base.I64_extend_i32_u(v158)
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_hash_metapage_info[2])))
										v167 = F_heap_form_tuple(m, v162, v10+int32(_a_F_hash_metapage_info_5), v10+int32(_a_F_hash_metapage_info_6))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return int64(0)
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
											v170 = F_HeapTupleHeaderGetDatum(m, v169)
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return int64(0)
											} else {
												m.G0 = v10 + int32(_a_F_hash_metapage_info_0)
												return v170
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
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_hash_metapage_info_7), int32(0))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_hash_metapage_info_3), int32(536), int32(_a_F_hash_metapage_info_4))
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
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
func F_hash_numeric(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
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
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
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
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v389 int64
	_ = v389
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v389
L2:
	;
	return int64(0)
L3:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v18 = int32(_a_F_hash_numeric_0)
	if v17&v18 == v18 {
		v389 = int64(0)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = int64(4294967295)
	v23 = base.I32_extend16_s(v17)
	if v23 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v43 = v39 + int32(base.Ui32(v40)>>(uint(int32(2))%32))
	v45 = int32(base.Ui32(v43) >> (uint(int32(1)) % 32))
	if v45 == int32(0) {
		v389 = v22
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v38 = v17<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v17&int32(63)
	v39 = int32(-6)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
	v38 = v36
	v39 = int32(-8)
	goto L5
L9:
	;
	v48 = int32(0)
	v50 = v13 + int32(6)
	v52 = v13 + int32(8)
	if v23 < v48 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v55 = v50
	goto L12
L11:
	;
	v55 = v52
	goto L12
L12:
	;
	v56 = v48
	v59 = v38
	goto L13
L13:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+v56<<(uint(int32(1))%32)))))
	if v70 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v56 == v45 {
		v389 = v22
		goto L1
	} else {
		goto L19
	}
L15:
	;
	v73 = int32(1)
	v76 = v56 + v73
	if v76 != v45 {
		v56 = v76
		v59 = v59 - v73
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
	v389 = v22
	goto L1
L19:
	;
	v80 = v45
	v83 = int32(0)
	goto L21
L20:
	;
	if int32(0) <= v23 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v90 = int32(1)
	v91 = v80 - v90
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+v91<<(uint(v90)%32)))))
	if v95 != 0 {
		v99 = v83
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v99 = v45
	goto L20
L23:
	;
	v97 = v83 + int32(1)
	if v97 != v45 {
		v80 = v91
		v83 = v97
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v105 = v52
	goto L27
L26:
	;
	v105 = v50
	goto L27
L27:
	;
	v106 = v56<<(uint(int32(1))%32) + v105
	v112 = (v43 - (v56+v99)<<(uint(int32(1))%32)) & int32(-2)
	v118 = v112 - int32(1636608432)
	if v106&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v389 = base.I64_extend_i32_s(v59) ^ base.I64_extend_i32_u(v372^v364-base.I32_rotl(v372, int32(24)))
	goto L1
L29:
	;
	v350 = int32(14)
	v352 = v346 ^ v347 - base.I32_rotl(v346, v350)
	v356 = v352 ^ v345 - base.I32_rotl(v352, int32(11))
	v360 = v356 ^ v346 - base.I32_rotl(v356, int32(25))
	v364 = v360 ^ v352 - base.I32_rotl(v360, int32(16))
	v368 = v364 ^ v356 - base.I32_rotl(v364, int32(4))
	v372 = v368 ^ v360 - base.I32_rotl(v368, v350)
	goto L28
L30:
	;
	switch v276 - int32(1) {
	case 0:
		v338 = v277
		v339 = v278
		v340 = v279
		goto L57
	case 1:
		v331 = v277
		v332 = v278
		v333 = v279
		goto L58
	case 2:
		v324 = v277
		v325 = v278
		v326 = v279
		goto L59
	case 3:
		v318 = v278
		v319 = v279
		goto L60
	case 4:
		v314 = v278
		v315 = v279
		goto L61
	case 5:
		v308 = v278
		v309 = v279
		goto L62
	case 6:
		v302 = v278
		v303 = v279
		goto L63
	case 7:
		v297 = v279
		goto L64
	case 8:
		v292 = v279
		goto L65
	case 9:
		v287 = v279
		goto L66
	case 10:
		goto L67
	default:
		v345 = v277
		v346 = v278
		v347 = v279
		goto L29
	}
L31:
	;
	v227 = v106
	v228 = v112
	v229 = v118
	v230 = v118
	v231 = v118
	goto L54
L32:
	;
	if base.Ui32(int32(11)) < base.Ui32(v112) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v112) < base.Ui32(int32(12)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v275 = v106
	v276 = v112
	v277 = v118
	v278 = v118
	v279 = v118
	goto L30
L36:
	;
	switch v174 - int32(1) {
	case 0:
		v224 = v175
		goto L43
	case 1:
		v219 = v175
		goto L44
	case 2:
		goto L45
	case 3:
		v212 = v176
		goto L46
	case 4:
		v209 = v176
		goto L47
	case 5:
		v204 = v176
		goto L48
	case 6:
		goto L49
	case 7:
		v195 = v177
		goto L50
	case 8:
		v190 = v177
		goto L51
	case 9:
		v185 = v177
		goto L52
	case 10:
		goto L53
	default:
		v345 = v175
		v346 = v176
		v347 = v177
		goto L29
	}
L37:
	;
	v173 = v106
	v174 = v112
	v175 = v118
	v176 = v118
	v177 = v118
	goto L36
L38:
	;
	goto L39
L39:
	;
	v125 = v106
	v126 = v112
	v127 = v118
	v128 = v118
	v129 = v118
	goto L40
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v132 = v131 + v128
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v136 = v135 + v129
	v138 = int32(4)
	v140 = v133 + v127 - v136 ^ base.I32_rotl(v136, v138)
	v144 = v132 - v140 ^ base.I32_rotl(v140, int32(6))
	v145 = v136 + v132
	v146 = v140 + v145
	v147 = v144 + v146
	v151 = v145 - v144 ^ base.I32_rotl(v144, int32(8))
	v155 = v146 - v151 ^ base.I32_rotl(v151, int32(16))
	v159 = v147 - v155 ^ base.I32_rotl(v155, int32(19))
	v160 = v151 + v147
	v161 = v155 + v160
	v162 = v159 + v161
	v166 = v160 - v159 ^ base.I32_rotl(v159, v138)
	v167 = int32(12)
	v168 = v125 + v167
	v170 = v126 - v167
	if base.Ui32(int32(11)) < base.Ui32(v170) {
		v125 = v168
		v126 = v170
		v127 = v161
		v128 = v162
		v129 = v166
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v173 = v168
	v174 = v170
	v175 = v161
	v176 = v162
	v177 = v166
	goto L36
L42:
	;
	goto L41
L43:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v345 = v224 + v225
	v346 = v176
	v347 = v177
	goto L29
L44:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	v224 = v220<<(uint(int32(8))%32) + v219
	goto L43
L45:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
	v219 = v215<<(uint(int32(16))%32) + v175
	goto L44
L46:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v345 = v213 + v175
	v346 = v212
	v347 = v177
	goto L29
L47:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
	v212 = v209 + v210
	goto L46
L48:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
	v209 = v205<<(uint(int32(8))%32) + v204
	goto L47
L49:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
	v204 = v200<<(uint(int32(16))%32) + v176
	goto L48
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v345 = v196 + v175
	v346 = v198 + v176
	v347 = v195
	goto L29
L51:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
	v195 = v191<<(uint(int32(8))%32) + v190
	goto L50
L52:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+9)))
	v190 = v186<<(uint(int32(16))%32) + v185
	goto L51
L53:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+10)))
	v185 = v181<<(uint(int32(24))%32) + v177
	goto L52
L54:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v234 = v233 + v230
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v238 = v237 + v231
	v240 = int32(4)
	v242 = v235 + v229 - v238 ^ base.I32_rotl(v238, v240)
	v246 = v234 - v242 ^ base.I32_rotl(v242, int32(6))
	v247 = v238 + v234
	v248 = v242 + v247
	v249 = v246 + v248
	v253 = v247 - v246 ^ base.I32_rotl(v246, int32(8))
	v257 = v248 - v253 ^ base.I32_rotl(v253, int32(16))
	v261 = v249 - v257 ^ base.I32_rotl(v257, int32(19))
	v262 = v253 + v249
	v263 = v257 + v262
	v264 = v261 + v263
	v268 = v262 - v261 ^ base.I32_rotl(v261, v240)
	v269 = int32(12)
	v270 = v227 + v269
	v272 = v228 - v269
	if base.Ui32(int32(11)) < base.Ui32(v272) {
		v227 = v270
		v228 = v272
		v229 = v263
		v230 = v264
		v231 = v268
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v275 = v270
	v276 = v272
	v277 = v263
	v278 = v264
	v279 = v268
	goto L30
L56:
	;
	goto L55
L57:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v345 = v338 + v341
	v346 = v339
	v347 = v340
	goto L29
L58:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	v338 = v334<<(uint(int32(8))%32) + v331
	v339 = v332
	v340 = v333
	goto L57
L59:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+2)))
	v331 = v327<<(uint(int32(16))%32) + v324
	v332 = v325
	v333 = v326
	goto L58
L60:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+3)))
	v324 = v320<<(uint(int32(24))%32) + v277
	v325 = v318
	v326 = v319
	goto L59
L61:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+4)))
	v318 = v314 + v316
	v319 = v315
	goto L60
L62:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+5)))
	v314 = v310<<(uint(int32(8))%32) + v308
	v315 = v309
	goto L61
L63:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+6)))
	v308 = v304<<(uint(int32(16))%32) + v302
	v309 = v303
	goto L62
L64:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+7)))
	v302 = v298<<(uint(int32(24))%32) + v278
	v303 = v297
	goto L63
L65:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+8)))
	v297 = v293<<(uint(int32(8))%32) + v292
	goto L64
L66:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+9)))
	v292 = v288<<(uint(int32(16))%32) + v287
	goto L65
L67:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+10)))
	v287 = v283<<(uint(int32(24))%32) + v279
	goto L66
}
func F_hash_page_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int64
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v180 int64
	_ = v180
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v2 = int32(0)
	v13 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(112)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int64(0)
	} else {
		v29 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)) = uint8(v29)
		*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = int64(0)
		v33 = F_superuser(m)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int64(0)
		} else {
			if v33 != 0 {
				v36 = F_verify_hash_page(m, v25, int32(3))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+16)))
					v39 = v36 + v38
					v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
					if base.Ui32(v41) < base.Ui32(int32(25)) {
						v180 = v13
						v187 = int64(0)
					} else {
						v48 = int32(base.Ui32(v41+int32(_a_F_hash_page_stats_0)) >> (uint(int32(2)) % 32))
						v50 = v48 & int32(_a_F_hash_page_stats_1)
						if v50 == int32(0) {
							v180 = v13
							v187 = int64(0)
						} else {
							v54 = v36 + int32(20)
							v55 = int32(1)
							if v50 != v55 {
								v65 = v2
								v66 = v2
								v67 = v55
								v70 = int32(0)
								for {
									v84 = v54 + v67<<(uint(int32(2))%32)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
									v86 = int32(_a_F_hash_page_stats_2)
									if v85&v86 != v86 {
										v94 = v65 + int32(1)
										v95 = v66
									} else {
										v94 = v65
										v95 = v66 + int32(1)
									}
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
									v97 = int32(_a_F_hash_page_stats_2)
									if v96&v97 != v97 {
										v105 = v94 + int32(1)
										v106 = v95
									} else {
										v105 = v94
										v106 = v95 + int32(1)
									}
									v107 = int32(2)
									v108 = v67 + v107
									v110 = v70 + v107
									if v110 != v48&int32(_a_F_hash_page_stats_3) {
										v65 = v105
										v66 = v106
										v67 = v108
										v70 = v110
										continue
									} else {
										break
									}
									break
								}
								if v48&int32(1) == int32(0) {
									v149 = v105
									v150 = v106
								} else {
									v116 = v105
									v117 = v106
									v118 = v108
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v54+v118<<(uint(int32(2))%32))))
									v139 = int32(_a_F_hash_page_stats_2)
									v142 = base.B2i32(v138&v139 == v139)
									if v138&v139 == v139 {
										v143 = v116
									} else {
										v143 = v116 + int32(1)
									}
									if v138&v139 == v139 {
										v146 = v117 + int32(1)
									} else {
										v146 = v117
									}
									v149 = v143
									v150 = v146
								}
							} else {
								v116 = v2
								v117 = v2
								v118 = v55
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v54+v118<<(uint(int32(2))%32))))
								v139 = int32(_a_F_hash_page_stats_2)
								v142 = base.B2i32(v138&v139 == v139)
								if v138&v139 == v139 {
									v143 = v116
								} else {
									v143 = v116 + int32(1)
								}
								if v138&v139 == v139 {
									v146 = v117 + int32(1)
								} else {
									v146 = v117
								}
								v149 = v143
								v150 = v146
							}
							v180 = base.I64_extend_i32_s(v150)
							v187 = base.I64_extend_i32_s(v149)
						}
					}
					v188 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v39)+14)))
					v189 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
					v190 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+8)))
					v191 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+4)))
					v192 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39))))
					v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+18)))
					v194 = int32(4)
					v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+14)))
					v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
					v197 = v195 - v196
					if v197 <= v194 {
						v200 = v194
					} else {
						v200 = v197
					}
					v206 = F_get_call_result_type(m, l0, int32(0), v22+int32(12))
					mBase = m.M
					v207 = m.ExcPending
					if v207 != 0 {
						return int64(0)
					} else {
						if v206 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_hash_page_stats_4), int32(0))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_page_stats_5), int32(263), int32(_a_F_hash_page_stats_6))
									mBase = m.M
									v268 = m.ExcPending
									if v268 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
							v211 = F_BlessTupleDesc(m, v210)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v188
								*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v189
								*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v190
								*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v191
								*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v192
								*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = base.I64_extend_i32_s(v200 - int32(4))
								*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v180
								*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v187
								*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v211
								*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = base.I64_extend_i32_u(v193) & int64(65280)
								v231 = F_heap_form_tuple(m, v211, v22+int32(32), v22+int32(16))
								mBase = m.M
								v232 = m.ExcPending
								if v232 != 0 {
									return int64(0)
								} else {
									v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
									v234 = F_HeapTupleHeaderGetDatum(m, v233)
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
										return int64(0)
									} else {
										m.G0 = v22 + int32(112)
										return v234
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v246 = m.ExcPending
					if v246 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_hash_page_stats_7), int32(0))
						mBase = m.M
						v250 = m.ExcPending
						if v250 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_hash_page_stats_5), int32(251), int32(_a_F_hash_page_stats_6))
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
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
func F_hash_range(m *base.Module, l0 int32) int64 {
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
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if v20 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v21 == v18 {
					v31 = v20
					F_range_deserialize(m, v31, v12, v7+int32(-16), v7+int32(-32), v7+int32(-33))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
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
								return int64(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
								if v55 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(52461700))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int64(0)
										} else {
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
											v144 = F_format_type_be(m, v143)
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v144
												F_errmsg(m, int32(_a_F_hash_range_0), v7+int32(-48))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_hash_range_1), int32(1596), int32(_a_F_hash_range_2))
													mBase = m.M
													v156 = m.ExcPending
													if v156 != 0 {
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
									v58 = v53
									v59 = int32(0)
									if v46&int32(41) == v59 {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
										v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int64(0)
										} else {
											v71 = base.I32_wrap_i64(v68)
											if v46&int32(81) == int32(0) {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
												v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int64(0)
												} else {
													v83 = base.I32_wrap_i64(v80)
													v88 = int32(711645284)
													v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
													v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
													v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
													v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
													v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
													m.G0 = v9 - int32(-64)
													return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
												}
											} else {
												v83 = v59
												v88 = int32(711645284)
												v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
												v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
												v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
												v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
												v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
												m.G0 = v9 - int32(-64)
												return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
											}
										}
									} else {
										v71 = int32(0)
										if v46&int32(81) == int32(0) {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
											v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int64(0)
											} else {
												v83 = base.I32_wrap_i64(v80)
												v88 = int32(711645284)
												v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
												v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
												v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
												v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
												v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
												m.G0 = v9 - int32(-64)
												return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
											}
										} else {
											v83 = v59
											v88 = int32(711645284)
											v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
											v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
											v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
											v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
											v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
											m.G0 = v9 - int32(-64)
											return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
										}
									}
								}
							}
						} else {
							v58 = v47
							v59 = int32(0)
							if v46&int32(41) == v59 {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
								v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int64(0)
								} else {
									v71 = base.I32_wrap_i64(v68)
									if v46&int32(81) == int32(0) {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
										v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int64(0)
										} else {
											v83 = base.I32_wrap_i64(v80)
											v88 = int32(711645284)
											v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
											v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
											v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
											v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
											v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
											m.G0 = v9 - int32(-64)
											return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
										}
									} else {
										v83 = v59
										v88 = int32(711645284)
										v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
										v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
										v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
										v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
										v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
										m.G0 = v9 - int32(-64)
										return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
									}
								}
							} else {
								v71 = int32(0)
								if v46&int32(81) == int32(0) {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
									v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
									v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int64(0)
									} else {
										v83 = base.I32_wrap_i64(v80)
										v88 = int32(711645284)
										v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
										v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
										v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
										v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
										v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
										m.G0 = v9 - int32(-64)
										return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
									}
								} else {
									v83 = v59
									v88 = int32(711645284)
									v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
									v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
									v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
									v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
									v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
									m.G0 = v9 - int32(-64)
									return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
								}
							}
						}
					}
				} else {
					v24 = F_lookup_type_cache(m, v18, int32(2048))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
								F_errmsg_internal(m, int32(_a_F_hash_range_3), v9)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_range_1), int32(1946), int32(_a_F_hash_range_4))
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int64(0)
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
							F_range_deserialize(m, v31, v12, v7+int32(-16), v7+int32(-32), v7+int32(-33))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
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
										return int64(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
										if v55 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(52461700))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int64(0)
												} else {
													v143 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
													v144 = F_format_type_be(m, v143)
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v144
														F_errmsg(m, int32(_a_F_hash_range_0), v7+int32(-48))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_hash_range_1), int32(1596), int32(_a_F_hash_range_2))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
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
											v58 = v53
											v59 = int32(0)
											if v46&int32(41) == v59 {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
												v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int64(0)
												} else {
													v71 = base.I32_wrap_i64(v68)
													if v46&int32(81) == int32(0) {
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
														v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
														v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int64(0)
														} else {
															v83 = base.I32_wrap_i64(v80)
															v88 = int32(711645284)
															v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
															v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
															v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
															v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
															v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
															m.G0 = v9 - int32(-64)
															return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
														}
													} else {
														v83 = v59
														v88 = int32(711645284)
														v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
														v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
														v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
														v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
														v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
														m.G0 = v9 - int32(-64)
														return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
													}
												}
											} else {
												v71 = int32(0)
												if v46&int32(81) == int32(0) {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
													v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
													v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int64(0)
													} else {
														v83 = base.I32_wrap_i64(v80)
														v88 = int32(711645284)
														v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
														v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
														v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
														v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
														v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
														m.G0 = v9 - int32(-64)
														return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
													}
												} else {
													v83 = v59
													v88 = int32(711645284)
													v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
													v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
													v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
													v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
													v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
													m.G0 = v9 - int32(-64)
													return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
												}
											}
										}
									}
								} else {
									v58 = v47
									v59 = int32(0)
									if v46&int32(41) == v59 {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
										v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int64(0)
										} else {
											v71 = base.I32_wrap_i64(v68)
											if v46&int32(81) == int32(0) {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
												v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int64(0)
												} else {
													v83 = base.I32_wrap_i64(v80)
													v88 = int32(711645284)
													v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
													v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
													v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
													v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
													v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
													m.G0 = v9 - int32(-64)
													return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
												}
											} else {
												v83 = v59
												v88 = int32(711645284)
												v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
												v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
												v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
												v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
												v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
												m.G0 = v9 - int32(-64)
												return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
											}
										}
									} else {
										v71 = int32(0)
										if v46&int32(81) == int32(0) {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
											v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int64(0)
											} else {
												v83 = base.I32_wrap_i64(v80)
												v88 = int32(711645284)
												v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
												v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
												v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
												v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
												v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
												m.G0 = v9 - int32(-64)
												return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
											}
										} else {
											v83 = v59
											v88 = int32(711645284)
											v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
											v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
											v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
											v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
											v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
											m.G0 = v9 - int32(-64)
											return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
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
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v18
							F_errmsg_internal(m, int32(_a_F_hash_range_3), v9)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_hash_range_1), int32(1946), int32(_a_F_hash_range_4))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int64(0)
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
						F_range_deserialize(m, v31, v12, v7+int32(-16), v7+int32(-32), v7+int32(-33))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
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
									return int64(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
									if v55 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(52461700))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int64(0)
											} else {
												v143 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
												v144 = F_format_type_be(m, v143)
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v144
													F_errmsg(m, int32(_a_F_hash_range_0), v7+int32(-48))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_hash_range_1), int32(1596), int32(_a_F_hash_range_2))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
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
										v58 = v53
										v59 = int32(0)
										if v46&int32(41) == v59 {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
											v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int64(0)
											} else {
												v71 = base.I32_wrap_i64(v68)
												if v46&int32(81) == int32(0) {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
													v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
													v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int64(0)
													} else {
														v83 = base.I32_wrap_i64(v80)
														v88 = int32(711645284)
														v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
														v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
														v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
														v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
														v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
														m.G0 = v9 - int32(-64)
														return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
													}
												} else {
													v83 = v59
													v88 = int32(711645284)
													v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
													v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
													v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
													v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
													v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
													m.G0 = v9 - int32(-64)
													return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
												}
											}
										} else {
											v71 = int32(0)
											if v46&int32(81) == int32(0) {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
												v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
												v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int64(0)
												} else {
													v83 = base.I32_wrap_i64(v80)
													v88 = int32(711645284)
													v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
													v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
													v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
													v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
													v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
													m.G0 = v9 - int32(-64)
													return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
												}
											} else {
												v83 = v59
												v88 = int32(711645284)
												v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
												v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
												v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
												v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
												v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
												m.G0 = v9 - int32(-64)
												return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
											}
										}
									}
								}
							} else {
								v58 = v47
								v59 = int32(0)
								if v46&int32(41) == v59 {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
									v68 = F_FunctionCall1Coll(m, v58+int32(132), v66, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int64(0)
									} else {
										v71 = base.I32_wrap_i64(v68)
										if v46&int32(81) == int32(0) {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
											v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int64(0)
											} else {
												v83 = base.I32_wrap_i64(v80)
												v88 = int32(711645284)
												v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
												v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
												v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
												v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
												v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
												m.G0 = v9 - int32(-64)
												return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
											}
										} else {
											v83 = v59
											v88 = int32(711645284)
											v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
											v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
											v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
											v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
											v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
											m.G0 = v9 - int32(-64)
											return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
										}
									}
								} else {
									v71 = int32(0)
									if v46&int32(81) == int32(0) {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
										v80 = F_FunctionCall1Coll(m, v58+int32(132), v78, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int64(0)
										} else {
											v83 = base.I32_wrap_i64(v80)
											v88 = int32(711645284)
											v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
											v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
											v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
											v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
											v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
											m.G0 = v9 - int32(-64)
											return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
										}
									} else {
										v83 = v59
										v88 = int32(711645284)
										v91 = v46 - int32(1636608428) ^ v88 - int32(1455628627)
										v96 = v91 ^ int32(-1636608428) - base.I32_rotl(v91, int32(25))
										v101 = v96 ^ v88 - base.I32_rotl(v96, int32(16))
										v105 = v101 ^ v91 - base.I32_rotl(v101, int32(4))
										v109 = v105 ^ v96 - base.I32_rotl(v105, int32(14))
										m.G0 = v9 - int32(-64)
										return base.I64_extend_i32_s(base.I32_rotl(v109^v101-base.I32_rotl(v109, int32(24))^v71, int32(1)) ^ v83)
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
func F_hash_range_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v170 int64
	_ = v170
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v6 = int64(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		F_check_stack_depth(m)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
			if v23 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				if v24 == v21 {
					v34 = v23
					F_range_deserialize(m, v34, v14, v9+int32(-16), v9+int32(-32), v9+int32(-33))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v43)>>(uint(int32(2))%32))-int32(1)))))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+200))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+164))
						if v51 == int32(0) {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							v56 = F_lookup_type_cache(m, v54, int32(_a_F_hash_range_extended_0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+164))
								if v58 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(52461700))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int64(0)
										} else {
											v202 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											v203 = F_format_type_be(m, v202)
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v203
												F_errmsg(m, int32(_a_F_hash_range_extended_1), v9+int32(-48))
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1660), int32(_a_F_hash_range_extended_3))
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
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
									v61 = v56
									if v49&int32(41) == int32(0) {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
										v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int64(0)
										} else {
											v72 = v70
											if v49&int32(81) == int32(0) {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
												v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int64(0)
												} else {
													v83 = v81
													if v18 == int64(0) {
														v90 = int32(-1636608428)
														v129 = v90
														v130 = v90
														v133 = int32(0)
													} else {
														v93 = base.I32_wrap_i64(v18)
														v95 = v93 + int32(1021750440)
														v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
														v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
														v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
														v111 = v100 + v95
														v112 = v106 + v111
														v113 = v110 + v112
														v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
														v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
														v126 = v117 + v113
														v127 = v121 + v126
														v129 = v127
														v130 = v126
														v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
													}
													v134 = int32(14)
													v136 = v133 - base.I32_rotl(v129, v134)
													v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
													v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
													v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
													v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
													v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
													m.G0 = v11 - int32(-64)
													v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
													return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
												}
											} else {
												v83 = v6
												if v18 == int64(0) {
													v90 = int32(-1636608428)
													v129 = v90
													v130 = v90
													v133 = int32(0)
												} else {
													v93 = base.I32_wrap_i64(v18)
													v95 = v93 + int32(1021750440)
													v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
													v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
													v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
													v111 = v100 + v95
													v112 = v106 + v111
													v113 = v110 + v112
													v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
													v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
													v126 = v117 + v113
													v127 = v121 + v126
													v129 = v127
													v130 = v126
													v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
												}
												v134 = int32(14)
												v136 = v133 - base.I32_rotl(v129, v134)
												v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
												v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
												v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
												v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
												v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
												m.G0 = v11 - int32(-64)
												v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
												return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
											}
										}
									} else {
										v72 = v6
										if v49&int32(81) == int32(0) {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
											v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int64(0)
											} else {
												v83 = v81
												if v18 == int64(0) {
													v90 = int32(-1636608428)
													v129 = v90
													v130 = v90
													v133 = int32(0)
												} else {
													v93 = base.I32_wrap_i64(v18)
													v95 = v93 + int32(1021750440)
													v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
													v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
													v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
													v111 = v100 + v95
													v112 = v106 + v111
													v113 = v110 + v112
													v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
													v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
													v126 = v117 + v113
													v127 = v121 + v126
													v129 = v127
													v130 = v126
													v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
												}
												v134 = int32(14)
												v136 = v133 - base.I32_rotl(v129, v134)
												v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
												v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
												v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
												v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
												v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
												m.G0 = v11 - int32(-64)
												v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
												return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
											}
										} else {
											v83 = v6
											if v18 == int64(0) {
												v90 = int32(-1636608428)
												v129 = v90
												v130 = v90
												v133 = int32(0)
											} else {
												v93 = base.I32_wrap_i64(v18)
												v95 = v93 + int32(1021750440)
												v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
												v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
												v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
												v111 = v100 + v95
												v112 = v106 + v111
												v113 = v110 + v112
												v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
												v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
												v126 = v117 + v113
												v127 = v121 + v126
												v129 = v127
												v130 = v126
												v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
											}
											v134 = int32(14)
											v136 = v133 - base.I32_rotl(v129, v134)
											v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
											v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
											v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
											v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
											v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
											m.G0 = v11 - int32(-64)
											v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
											return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
										}
									}
								}
							}
						} else {
							v61 = v50
							if v49&int32(41) == int32(0) {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
								v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									v72 = v70
									if v49&int32(81) == int32(0) {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int64(0)
										} else {
											v83 = v81
											if v18 == int64(0) {
												v90 = int32(-1636608428)
												v129 = v90
												v130 = v90
												v133 = int32(0)
											} else {
												v93 = base.I32_wrap_i64(v18)
												v95 = v93 + int32(1021750440)
												v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
												v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
												v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
												v111 = v100 + v95
												v112 = v106 + v111
												v113 = v110 + v112
												v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
												v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
												v126 = v117 + v113
												v127 = v121 + v126
												v129 = v127
												v130 = v126
												v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
											}
											v134 = int32(14)
											v136 = v133 - base.I32_rotl(v129, v134)
											v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
											v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
											v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
											v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
											v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
											m.G0 = v11 - int32(-64)
											v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
											return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
										}
									} else {
										v83 = v6
										if v18 == int64(0) {
											v90 = int32(-1636608428)
											v129 = v90
											v130 = v90
											v133 = int32(0)
										} else {
											v93 = base.I32_wrap_i64(v18)
											v95 = v93 + int32(1021750440)
											v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
											v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
											v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
											v111 = v100 + v95
											v112 = v106 + v111
											v113 = v110 + v112
											v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
											v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
											v126 = v117 + v113
											v127 = v121 + v126
											v129 = v127
											v130 = v126
											v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
										}
										v134 = int32(14)
										v136 = v133 - base.I32_rotl(v129, v134)
										v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
										v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
										v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
										v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
										v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
										m.G0 = v11 - int32(-64)
										v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
										return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
									}
								}
							} else {
								v72 = v6
								if v49&int32(81) == int32(0) {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
									v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int64(0)
									} else {
										v83 = v81
										if v18 == int64(0) {
											v90 = int32(-1636608428)
											v129 = v90
											v130 = v90
											v133 = int32(0)
										} else {
											v93 = base.I32_wrap_i64(v18)
											v95 = v93 + int32(1021750440)
											v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
											v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
											v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
											v111 = v100 + v95
											v112 = v106 + v111
											v113 = v110 + v112
											v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
											v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
											v126 = v117 + v113
											v127 = v121 + v126
											v129 = v127
											v130 = v126
											v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
										}
										v134 = int32(14)
										v136 = v133 - base.I32_rotl(v129, v134)
										v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
										v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
										v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
										v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
										v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
										m.G0 = v11 - int32(-64)
										v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
										return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
									}
								} else {
									v83 = v6
									if v18 == int64(0) {
										v90 = int32(-1636608428)
										v129 = v90
										v130 = v90
										v133 = int32(0)
									} else {
										v93 = base.I32_wrap_i64(v18)
										v95 = v93 + int32(1021750440)
										v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
										v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
										v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
										v111 = v100 + v95
										v112 = v106 + v111
										v113 = v110 + v112
										v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
										v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
										v126 = v117 + v113
										v127 = v121 + v126
										v129 = v127
										v130 = v126
										v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
									}
									v134 = int32(14)
									v136 = v133 - base.I32_rotl(v129, v134)
									v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
									v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
									v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
									v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
									v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
									m.G0 = v11 - int32(-64)
									v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
									return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
								}
							}
						}
					}
				} else {
					v27 = F_lookup_type_cache(m, v21, int32(2048))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v21
								F_errmsg_internal(m, int32(_a_F_hash_range_extended_4), v11)
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1946), int32(_a_F_hash_range_extended_5))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
							v34 = v27
							F_range_deserialize(m, v34, v14, v9+int32(-16), v9+int32(-32), v9+int32(-33))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
								v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v43)>>(uint(int32(2))%32))-int32(1)))))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+200))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+164))
								if v51 == int32(0) {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									v56 = F_lookup_type_cache(m, v54, int32(_a_F_hash_range_extended_0))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int64(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+164))
										if v58 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(52461700))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int64(0)
												} else {
													v202 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
													v203 = F_format_type_be(m, v202)
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v203
														F_errmsg(m, int32(_a_F_hash_range_extended_1), v9+int32(-48))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1660), int32(_a_F_hash_range_extended_3))
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
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
											v61 = v56
											if v49&int32(41) == int32(0) {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
												v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
												v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int64(0)
												} else {
													v72 = v70
													if v49&int32(81) == int32(0) {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
														v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
														v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
														mBase = m.M
														v82 = m.ExcPending
														if v82 != 0 {
															return int64(0)
														} else {
															v83 = v81
															if v18 == int64(0) {
																v90 = int32(-1636608428)
																v129 = v90
																v130 = v90
																v133 = int32(0)
															} else {
																v93 = base.I32_wrap_i64(v18)
																v95 = v93 + int32(1021750440)
																v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
																v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
																v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
																v111 = v100 + v95
																v112 = v106 + v111
																v113 = v110 + v112
																v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
																v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
																v126 = v117 + v113
																v127 = v121 + v126
																v129 = v127
																v130 = v126
																v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
															}
															v134 = int32(14)
															v136 = v133 - base.I32_rotl(v129, v134)
															v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
															v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
															v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
															v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
															v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
															m.G0 = v11 - int32(-64)
															v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
															return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
														}
													} else {
														v83 = v6
														if v18 == int64(0) {
															v90 = int32(-1636608428)
															v129 = v90
															v130 = v90
															v133 = int32(0)
														} else {
															v93 = base.I32_wrap_i64(v18)
															v95 = v93 + int32(1021750440)
															v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
															v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
															v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
															v111 = v100 + v95
															v112 = v106 + v111
															v113 = v110 + v112
															v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
															v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
															v126 = v117 + v113
															v127 = v121 + v126
															v129 = v127
															v130 = v126
															v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
														}
														v134 = int32(14)
														v136 = v133 - base.I32_rotl(v129, v134)
														v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
														v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
														v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
														v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
														v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
														m.G0 = v11 - int32(-64)
														v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
														return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
													}
												}
											} else {
												v72 = v6
												if v49&int32(81) == int32(0) {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
													v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
													v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int64(0)
													} else {
														v83 = v81
														if v18 == int64(0) {
															v90 = int32(-1636608428)
															v129 = v90
															v130 = v90
															v133 = int32(0)
														} else {
															v93 = base.I32_wrap_i64(v18)
															v95 = v93 + int32(1021750440)
															v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
															v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
															v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
															v111 = v100 + v95
															v112 = v106 + v111
															v113 = v110 + v112
															v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
															v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
															v126 = v117 + v113
															v127 = v121 + v126
															v129 = v127
															v130 = v126
															v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
														}
														v134 = int32(14)
														v136 = v133 - base.I32_rotl(v129, v134)
														v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
														v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
														v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
														v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
														v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
														m.G0 = v11 - int32(-64)
														v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
														return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
													}
												} else {
													v83 = v6
													if v18 == int64(0) {
														v90 = int32(-1636608428)
														v129 = v90
														v130 = v90
														v133 = int32(0)
													} else {
														v93 = base.I32_wrap_i64(v18)
														v95 = v93 + int32(1021750440)
														v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
														v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
														v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
														v111 = v100 + v95
														v112 = v106 + v111
														v113 = v110 + v112
														v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
														v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
														v126 = v117 + v113
														v127 = v121 + v126
														v129 = v127
														v130 = v126
														v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
													}
													v134 = int32(14)
													v136 = v133 - base.I32_rotl(v129, v134)
													v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
													v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
													v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
													v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
													v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
													m.G0 = v11 - int32(-64)
													v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
													return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
												}
											}
										}
									}
								} else {
									v61 = v50
									if v49&int32(41) == int32(0) {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
										v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int64(0)
										} else {
											v72 = v70
											if v49&int32(81) == int32(0) {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
												v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int64(0)
												} else {
													v83 = v81
													if v18 == int64(0) {
														v90 = int32(-1636608428)
														v129 = v90
														v130 = v90
														v133 = int32(0)
													} else {
														v93 = base.I32_wrap_i64(v18)
														v95 = v93 + int32(1021750440)
														v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
														v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
														v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
														v111 = v100 + v95
														v112 = v106 + v111
														v113 = v110 + v112
														v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
														v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
														v126 = v117 + v113
														v127 = v121 + v126
														v129 = v127
														v130 = v126
														v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
													}
													v134 = int32(14)
													v136 = v133 - base.I32_rotl(v129, v134)
													v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
													v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
													v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
													v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
													v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
													m.G0 = v11 - int32(-64)
													v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
													return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
												}
											} else {
												v83 = v6
												if v18 == int64(0) {
													v90 = int32(-1636608428)
													v129 = v90
													v130 = v90
													v133 = int32(0)
												} else {
													v93 = base.I32_wrap_i64(v18)
													v95 = v93 + int32(1021750440)
													v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
													v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
													v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
													v111 = v100 + v95
													v112 = v106 + v111
													v113 = v110 + v112
													v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
													v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
													v126 = v117 + v113
													v127 = v121 + v126
													v129 = v127
													v130 = v126
													v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
												}
												v134 = int32(14)
												v136 = v133 - base.I32_rotl(v129, v134)
												v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
												v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
												v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
												v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
												v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
												m.G0 = v11 - int32(-64)
												v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
												return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
											}
										}
									} else {
										v72 = v6
										if v49&int32(81) == int32(0) {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
											v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int64(0)
											} else {
												v83 = v81
												if v18 == int64(0) {
													v90 = int32(-1636608428)
													v129 = v90
													v130 = v90
													v133 = int32(0)
												} else {
													v93 = base.I32_wrap_i64(v18)
													v95 = v93 + int32(1021750440)
													v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
													v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
													v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
													v111 = v100 + v95
													v112 = v106 + v111
													v113 = v110 + v112
													v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
													v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
													v126 = v117 + v113
													v127 = v121 + v126
													v129 = v127
													v130 = v126
													v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
												}
												v134 = int32(14)
												v136 = v133 - base.I32_rotl(v129, v134)
												v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
												v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
												v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
												v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
												v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
												m.G0 = v11 - int32(-64)
												v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
												return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
											}
										} else {
											v83 = v6
											if v18 == int64(0) {
												v90 = int32(-1636608428)
												v129 = v90
												v130 = v90
												v133 = int32(0)
											} else {
												v93 = base.I32_wrap_i64(v18)
												v95 = v93 + int32(1021750440)
												v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
												v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
												v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
												v111 = v100 + v95
												v112 = v106 + v111
												v113 = v110 + v112
												v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
												v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
												v126 = v117 + v113
												v127 = v121 + v126
												v129 = v127
												v130 = v126
												v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
											}
											v134 = int32(14)
											v136 = v133 - base.I32_rotl(v129, v134)
											v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
											v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
											v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
											v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
											v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
											m.G0 = v11 - int32(-64)
											v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
											return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
										}
									}
								}
							}
						}
					}
				}
			} else {
				v27 = F_lookup_type_cache(m, v21, int32(2048))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
					if v29 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v21
							F_errmsg_internal(m, int32(_a_F_hash_range_extended_4), v11)
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1946), int32(_a_F_hash_range_extended_5))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
						v34 = v27
						F_range_deserialize(m, v34, v14, v9+int32(-16), v9+int32(-32), v9+int32(-33))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v43)>>(uint(int32(2))%32))-int32(1)))))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+200))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+164))
							if v51 == int32(0) {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
								v56 = F_lookup_type_cache(m, v54, int32(_a_F_hash_range_extended_0))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int64(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+164))
									if v58 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(52461700))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int64(0)
											} else {
												v202 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
												v203 = F_format_type_be(m, v202)
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v203
													F_errmsg(m, int32(_a_F_hash_range_extended_1), v9+int32(-48))
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_hash_range_extended_2), int32(1660), int32(_a_F_hash_range_extended_3))
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
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
										v61 = v56
										if v49&int32(41) == int32(0) {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
											v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
											v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int64(0)
											} else {
												v72 = v70
												if v49&int32(81) == int32(0) {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
													v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
													v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int64(0)
													} else {
														v83 = v81
														if v18 == int64(0) {
															v90 = int32(-1636608428)
															v129 = v90
															v130 = v90
															v133 = int32(0)
														} else {
															v93 = base.I32_wrap_i64(v18)
															v95 = v93 + int32(1021750440)
															v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
															v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
															v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
															v111 = v100 + v95
															v112 = v106 + v111
															v113 = v110 + v112
															v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
															v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
															v126 = v117 + v113
															v127 = v121 + v126
															v129 = v127
															v130 = v126
															v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
														}
														v134 = int32(14)
														v136 = v133 - base.I32_rotl(v129, v134)
														v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
														v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
														v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
														v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
														v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
														m.G0 = v11 - int32(-64)
														v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
														return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
													}
												} else {
													v83 = v6
													if v18 == int64(0) {
														v90 = int32(-1636608428)
														v129 = v90
														v130 = v90
														v133 = int32(0)
													} else {
														v93 = base.I32_wrap_i64(v18)
														v95 = v93 + int32(1021750440)
														v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
														v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
														v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
														v111 = v100 + v95
														v112 = v106 + v111
														v113 = v110 + v112
														v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
														v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
														v126 = v117 + v113
														v127 = v121 + v126
														v129 = v127
														v130 = v126
														v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
													}
													v134 = int32(14)
													v136 = v133 - base.I32_rotl(v129, v134)
													v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
													v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
													v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
													v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
													v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
													m.G0 = v11 - int32(-64)
													v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
													return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
												}
											}
										} else {
											v72 = v6
											if v49&int32(81) == int32(0) {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
												v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
												v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int64(0)
												} else {
													v83 = v81
													if v18 == int64(0) {
														v90 = int32(-1636608428)
														v129 = v90
														v130 = v90
														v133 = int32(0)
													} else {
														v93 = base.I32_wrap_i64(v18)
														v95 = v93 + int32(1021750440)
														v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
														v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
														v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
														v111 = v100 + v95
														v112 = v106 + v111
														v113 = v110 + v112
														v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
														v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
														v126 = v117 + v113
														v127 = v121 + v126
														v129 = v127
														v130 = v126
														v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
													}
													v134 = int32(14)
													v136 = v133 - base.I32_rotl(v129, v134)
													v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
													v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
													v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
													v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
													v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
													m.G0 = v11 - int32(-64)
													v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
													return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
												}
											} else {
												v83 = v6
												if v18 == int64(0) {
													v90 = int32(-1636608428)
													v129 = v90
													v130 = v90
													v133 = int32(0)
												} else {
													v93 = base.I32_wrap_i64(v18)
													v95 = v93 + int32(1021750440)
													v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
													v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
													v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
													v111 = v100 + v95
													v112 = v106 + v111
													v113 = v110 + v112
													v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
													v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
													v126 = v117 + v113
													v127 = v121 + v126
													v129 = v127
													v130 = v126
													v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
												}
												v134 = int32(14)
												v136 = v133 - base.I32_rotl(v129, v134)
												v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
												v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
												v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
												v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
												v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
												m.G0 = v11 - int32(-64)
												v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
												return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
											}
										}
									}
								}
							} else {
								v61 = v50
								if v49&int32(41) == int32(0) {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
									v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										v72 = v70
										if v49&int32(81) == int32(0) {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
											v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
											v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int64(0)
											} else {
												v83 = v81
												if v18 == int64(0) {
													v90 = int32(-1636608428)
													v129 = v90
													v130 = v90
													v133 = int32(0)
												} else {
													v93 = base.I32_wrap_i64(v18)
													v95 = v93 + int32(1021750440)
													v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
													v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
													v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
													v111 = v100 + v95
													v112 = v106 + v111
													v113 = v110 + v112
													v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
													v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
													v126 = v117 + v113
													v127 = v121 + v126
													v129 = v127
													v130 = v126
													v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
												}
												v134 = int32(14)
												v136 = v133 - base.I32_rotl(v129, v134)
												v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
												v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
												v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
												v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
												v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
												m.G0 = v11 - int32(-64)
												v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
												return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
											}
										} else {
											v83 = v6
											if v18 == int64(0) {
												v90 = int32(-1636608428)
												v129 = v90
												v130 = v90
												v133 = int32(0)
											} else {
												v93 = base.I32_wrap_i64(v18)
												v95 = v93 + int32(1021750440)
												v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
												v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
												v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
												v111 = v100 + v95
												v112 = v106 + v111
												v113 = v110 + v112
												v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
												v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
												v126 = v117 + v113
												v127 = v121 + v126
												v129 = v127
												v130 = v126
												v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
											}
											v134 = int32(14)
											v136 = v133 - base.I32_rotl(v129, v134)
											v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
											v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
											v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
											v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
											v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
											m.G0 = v11 - int32(-64)
											v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
											return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
										}
									}
								} else {
									v72 = v6
									if v49&int32(81) == int32(0) {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
										v80 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										v81 = F_FunctionCall2Coll(m, v61+int32(160), v79, v80, v18)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int64(0)
										} else {
											v83 = v81
											if v18 == int64(0) {
												v90 = int32(-1636608428)
												v129 = v90
												v130 = v90
												v133 = int32(0)
											} else {
												v93 = base.I32_wrap_i64(v18)
												v95 = v93 + int32(1021750440)
												v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
												v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
												v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
												v111 = v100 + v95
												v112 = v106 + v111
												v113 = v110 + v112
												v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
												v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
												v126 = v117 + v113
												v127 = v121 + v126
												v129 = v127
												v130 = v126
												v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
											}
											v134 = int32(14)
											v136 = v133 - base.I32_rotl(v129, v134)
											v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
											v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
											v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
											v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
											v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
											m.G0 = v11 - int32(-64)
											v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
											return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
										}
									} else {
										v83 = v6
										if v18 == int64(0) {
											v90 = int32(-1636608428)
											v129 = v90
											v130 = v90
											v133 = int32(0)
										} else {
											v93 = base.I32_wrap_i64(v18)
											v95 = v93 + int32(1021750440)
											v100 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ int32(-415931063)
											v106 = v93 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
											v110 = v95 - v106 ^ base.I32_rotl(v106, int32(8))
											v111 = v100 + v95
											v112 = v106 + v111
											v113 = v110 + v112
											v117 = v111 - v110 ^ base.I32_rotl(v110, int32(16))
											v121 = v112 - v117 ^ base.I32_rotl(v117, int32(19))
											v126 = v117 + v113
											v127 = v121 + v126
											v129 = v127
											v130 = v126
											v133 = v113 - v121 ^ base.I32_rotl(v121, int32(4)) ^ v127
										}
										v134 = int32(14)
										v136 = v133 - base.I32_rotl(v129, v134)
										v141 = v136 ^ (v49 + v130) - base.I32_rotl(v136, int32(11))
										v145 = v129 ^ v141 - base.I32_rotl(v141, int32(25))
										v149 = v145 ^ v136 - base.I32_rotl(v145, int32(16))
										v153 = v149 ^ v141 - base.I32_rotl(v149, int32(4))
										v157 = v153 ^ v145 - base.I32_rotl(v153, v134)
										m.G0 = v11 - int32(-64)
										v170 = base.I64_extend_i32_u(v157)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v157^v149-base.I32_rotl(v157, int32(24))) ^ v72
										return v170<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v170)>>(uint(int64(31))%64))&int64(4294967297) ^ v83
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
func F_hash_record(m *base.Module, l0 int32) int64 {
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
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int64
	_ = v192
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v252 int64
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v21
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v34
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+68)) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = int32(base.Ui32(v32) >> (uint(int32(2)) % 32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v44 == v34 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v27 == v65 {
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
	v114 = F_palloc_mul(m, int32(8), v31)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v67 == v28 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v70 = v64 + int32(20)
	v74 = v31 << (uint(int32(2)) % 32)
	if v70&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v74)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v27
	goto L10
L16:
	;
	if v74 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v74 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L19:
	;
	v84 = v64 + v74 + int32(20)
	v86 = v64 + int32(24)
	if base.Ui32(v86) < base.Ui32(v84) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v88 = v84
	goto L22
L21:
	;
	v88 = v86
	goto L22
L22:
	;
	v95 = (v88-v64-int32(21))&int32(-4) + int32(4)
	if v95 == int32(0) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	base.MemoryFill(m, v70, int32(0), v95)
	goto L15
L24:
	;
	base.MemoryFill(m, v70, int32(0), v74)
	goto L15
L25:
	;
	v117 = F_palloc_mul(m, int32(1), v31)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v18+int32(60), v29, v114, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v31 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v252 = int64(0)
	goto L30
L29:
	;
	v126 = int32(0)
	v129 = v126
	v134 = v126
	goto L32
L30:
	;
	F_pfree(m, v114)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L56
	}
L31:
	;
	v252 = base.I64_extend_i32_u(v210)
	goto L30
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v149 = v29 + v143<<(uint(int32(3))%32) + v129*int32(100)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+119)))
	if v150 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L34:
	;
	goto L33
L35:
	;
	v154 = v149 + int32(28)
	v157 = v64 + int32(20) + v129<<(uint(int32(2))%32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v158 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v210 = v134
	goto L37
L37:
	;
	v215 = v129 + int32(1)
	if v31 != v215 {
		v129 = v215
		v134 = v210
		goto L32
	} else {
		goto L50
	}
L38:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v117))))
	if v176 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v167 = F_lookup_type_cache(m, v165, int32(128))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L44
	}
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+68))
	v165 = v161
	goto L39
L41:
	;
	goto L42
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)+68))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v162 == v163 {
		v173 = v158
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v165 = v162
	goto L39
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)+136))
	if v169 == int32(0) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v167
	v173 = v167
	goto L38
L46:
	;
	v204 = int32(0)
	goto L48
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v173 + int32(132)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v154)+96))
	v184 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+34)) = uint16(v184)
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)) = uint8(v186)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v183
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v114+v129<<(uint(int32(3))%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+48)) = uint8(v186)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v192
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v173)+132))
	v199 = m.T0[v198].(func(*base.Module, int32) int64)(m, v18+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v210 = v204 + v134*int32(31)
	goto L37
L49:
	;
	v204 = base.I32_wrap_i64(v199)
	goto L48
L50:
	;
	goto L31
L51:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v225 = F_format_type_be(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v225
	F_errmsg(m, int32(_a_F_hash_record_0), v18)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_hash_record_1), int32(1894), int32(_a_F_hash_record_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_pfree(m, v117)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if int32(0) <= v257 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_DecrTupleDescRefCount(m, v29)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v262 != v21 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	F_pfree(m, v21)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	m.G0 = v18 + int32(80)
	return v252
L65:
	;
	goto L64
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
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
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v205 == int32(0) {
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v159 = F_palloc_mul(m, int32(24), v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v159
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v162 <= int32(0) {
		goto L34
	} else {
		goto L54
	}
L54:
	;
	v169 = int32(0)
	goto L55
L55:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v169*int32(52))+48))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v186)+96))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v176+v169*int32(24), v180, int32(0), v187, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L34
L57:
	;
	v192 = v169 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v192 < v193 {
		v169 = v192
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
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v208 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v212 = v209 - v208
	goto L63
L62:
	;
	v212 = int32(0)
	goto L63
L63:
	;
	if v205 <= int32(0) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v219 = int32(0)
	goto L65
L65:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	F_initialize_aggregate(m, l0, v226+v219*int32(240), v212+v219<<(uint(int32(4))%32))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L49
	} else {
		goto L67
	}
L66:
	;
	goto L59
L67:
	;
	v236 = v219 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v236 < v237 {
		v219 = v236
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
	var v139 int32
	_ = v139
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
	var v167 int32
	_ = v167
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
	var v210 int32
	_ = v210
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
	var v235 int32
	_ = v235
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
							v139 = F_errdetail(m, int32(_a_F_verify_hash_page_3), v8+int32(96))
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
								v167 = F_errdetail(m, int32(_a_F_verify_hash_page_7), v8-int32(-64))
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
												v210 = F_errdetail(m, int32(_a_F_verify_hash_page_14), v8+int32(32))
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
													v235 = F_errdetail(m, int32(_a_F_verify_hash_page_16), v8+int32(16))
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
