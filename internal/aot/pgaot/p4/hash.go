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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	v4 = l3
	v10 = base.F64_div(l1, base.F64_convert_i32_u(v4))
	if base.F64_le(v10, float64(2)) != 0 {
		v69 = int32(2)
	} else {
		if base.F64_ge(v10, float64(1.073741824e+09)) != 0 {
			v69 = int32(1073741824)
		} else {
			if base.F64_lt(v10, float64(4.294967296e+09))&base.F64_ge(v10, float64(0)) != 0 {
				v21 = base.I32_trunc_f64_u(v10)
				v23 = v21
			} else {
				v23 = int32(0)
			}
			v27 = v23 - int32(1)
			if base.Ui32(int32(2)) <= base.Ui32(v23) {
				v33 = int32(32) - base.I32_clz(v27)
			} else {
				v33 = int32(0)
			}
			if base.Ui32(int32(10)) <= base.Ui32(v33) {
				v36 = int32(3)
				v46 = int32(base.Ui32(v27)>>(uint(v33-v36)%32))&v36 | v33<<(uint(int32(2))%32) - int32(30)
			} else {
				v46 = v33
			}
			if base.Ui32(v46) <= base.Ui32(int32(9)) {
				v68 = int32(1) << (uint(v46) % 32)
			} else {
				v54 = v46 - int32(10)
				v55 = int32(2)
				v57 = int32(512) << (uint(int32(base.Ui32(v54)>>(uint(v55)%32))) % 32)
				v68 = v57>>(uint(v55)%32)*(v54&int32(3)+int32(1)) + v57
			}
			v69 = v68
		}
	}
	v73 = v69 - int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v69) {
		v79 = int32(32) - base.I32_clz(v73)
	} else {
		v79 = int32(0)
	}
	if base.Ui32(int32(10)) <= base.Ui32(v79) {
		v82 = int32(3)
		v92 = int32(base.Ui32(v73)>>(uint(v79-v82)%32))&v82 | v79<<(uint(int32(2))%32) - int32(30)
	} else {
		v92 = v79
	}
	if l0 < int32(0) {
		v96 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+(l0^int32(-1))<<(uint(int32(2))%32))))
		v110 = v102
	} else {
		v104 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v110 = v104 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if l4 != 0 {
		if v110&int32(3) != 0 {
		} else {
		}
		v137 = F___memset(m, v110, int32(0), int32(8192))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v110)+10)) = int32(1572864)
		v143 = int32(8196)
		*(*uint16)(unsafe.Add(mBase, uint32(v110)+18)) = uint16(v143)
		v149 = int32(8176)
		*(*uint16)(unsafe.Add(mBase, uint32(v110)+16)) = uint16(v149)
		*(*uint16)(unsafe.Add(mBase, uint32(v110)+14)) = uint16(v149)
	} else {
	}
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+16)))
	v153 = v110 + v152
	*(*int64)(unsafe.Add(mBase, uint32(v153)+8)) = int64(-36028758364258305)
	*(*int64)(unsafe.Add(mBase, uint32(v153))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v110)+32)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v110)+24)) = int64(17284990528)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+40)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+72)) = l2
	v166 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+48)) = v69 - v166
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+19)))
	v173 = v169<<(uint(int32(8))%32) - int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+42)) = uint16(v173)
	v175 = int32(-1)
	v178 = v69 + v166
	if v178&v69 != 0 {
		v185 = v175<<(uint(int32(32)-base.I32_clz(v178))%32) ^ v175
	} else {
		v185 = v69
	}
	*(*int32)(unsafe.Add(mBase, uint32(v110)+52)) = v185
	v187 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+56)) = int32(base.Ui32(v185) >> (uint(v187) % 32))
	v194 = base.I32_clz(v173&int32(65496)) ^ int32(31)
	v195 = int32(3)
	v196 = v194 + v195
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+46)) = uint16(v196)
	v199 = v187 << (uint(v194) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+44)) = uint16(v199)
	v202 = v110 + int32(76)
	if v202&v195 == int32(0) {
		v208 = v110 + int32(468)
		if base.Ui32(v208) <= base.Ui32(v202) {
		} else {
			v212 = v110 + int32(80)
			if base.Ui32(v212) < base.Ui32(v208) {
				v214 = v208
			} else {
				v214 = v212
			}
			v223 = F__emscripten_memset_bulkmem(m, v202, base.I32_extend8_s(int32(0)), (v214-v110-int32(77))&int32(-4)+int32(4))
			mBase = m.M
		}
	} else {
		v227 = F__emscripten_memset_bulkmem(m, v202, base.I32_extend8_s(int32(0)), int32(392))
		mBase = m.M
	}
	v235 = F__emscripten_memset_bulkmem(m, v110+int32(468), base.I32_extend8_s(int32(0)), int32(4096))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v202+v92<<(uint(int32(2))%32)))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v110-int32(-64)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+60)) = v92
	v246 = int32(4568)
	*(*uint16)(unsafe.Add(mBase, uint32(v110)+12)) = uint16(v246)
	return
}
func F_hashRowType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(711645284)
	v14 = v6 - int32(1636608428) ^ v11 - int32(1455628627)
	v19 = v14 ^ int32(-1636608428) - base.I32_rotl(v14, int32(25))
	v24 = v19 ^ v11 - base.I32_rotl(v19, int32(16))
	v28 = v24 ^ v14 - base.I32_rotl(v24, int32(4))
	v32 = v28 ^ v19 - base.I32_rotl(v28, int32(14))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = int32(711645284)
	v45 = v37 - int32(1636608428) ^ v42 - int32(1455628627)
	v50 = v45 ^ int32(-1636608428) - base.I32_rotl(v45, int32(25))
	v55 = v50 ^ v42 - base.I32_rotl(v50, int32(16))
	v59 = v55 ^ v45 - base.I32_rotl(v55, int32(4))
	v63 = v59 ^ v50 - base.I32_rotl(v59, int32(14))
	v68 = int32(1640531527)
	v69 = v32 ^ v24 - base.I32_rotl(v32, int32(24)) - v68
	v78 = v63 ^ v55 - base.I32_rotl(v63, int32(24)) + v69<<(uint(int32(6))%32) + int32(base.Ui32(v69)>>(uint(int32(2))%32)) - v68 ^ v69
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v79 {
		v85 = v78
		v86 = int32(0)
		v87 = v79
		for {
			v89 = int32(4)
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(88)+v87<<(uint(v89)%32)+v86*int32(100))))
			v100 = int32(711645284)
			v103 = v95 - int32(1636608428) ^ v100 - int32(1455628627)
			v108 = v103 ^ int32(-1636608428) - base.I32_rotl(v103, int32(25))
			v113 = v108 ^ v100 - base.I32_rotl(v108, int32(16))
			v117 = v113 ^ v103 - base.I32_rotl(v113, v89)
			v121 = v117 ^ v108 - base.I32_rotl(v117, int32(14))
			v134 = v121 ^ v113 - base.I32_rotl(v121, int32(24)) + (v85<<(uint(int32(6))%32) + int32(base.Ui32(v85)>>(uint(int32(2))%32))) - int32(1640531527) ^ v85
			v136 = v86 + int32(1)
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v136 < v137 {
				v85 = v134
				v86 = v136
				v87 = v137
				continue
			} else {
				break
			}
			break
		}
		v140 = v134
	} else {
		v140 = v78
	}
	return v140
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
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
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
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
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
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
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	v17 = int32(49152)
	if v16&v17 == v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v23 = base.I32_extend16_s(v16)
	if v23 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v41 = int32(2)
	v43 = v39 + int32(base.Ui32(v40)>>(uint(v41)%32))
	if base.Ui32(v43) < base.Ui32(v41) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v38 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
	v39 = int32(-6)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	v38 = v36
	v39 = int32(-8)
	goto L6
L10:
	;
	return int32(-1)
L11:
	;
	goto L12
L12:
	;
	v48 = int32(0)
	v50 = v12 + int32(6)
	v52 = v12 + int32(8)
	if v23 < v48 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = v50
	goto L15
L14:
	;
	v55 = v52
	goto L15
L15:
	;
	v56 = int32(1)
	v58 = int32(base.Ui32(v43) >> (uint(v56) % 32))
	if base.Ui32(v58) <= base.Ui32(v56) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v61 = v56
	goto L18
L17:
	;
	v61 = v58
	goto L18
L18:
	;
	v63 = v48
	v66 = v38
	goto L20
L19:
	;
	if v82 == v58 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+v63<<(uint(int32(1))%32)))))
	if v76 != 0 {
		v82 = v63
		v83 = v66
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v82 = v61
	v83 = v38 - v61
	goto L19
L22:
	;
	v77 = int32(1)
	v80 = v63 + v77
	if v80 != v61 {
		v63 = v80
		v66 = v66 - v77
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	return int32(-1)
L25:
	;
	goto L26
L26:
	;
	v87 = int32(1)
	if v58 <= v87 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v90 = v87
	goto L29
L28:
	;
	v90 = v58
	goto L29
L29:
	;
	v94 = int32(0)
	v96 = v58
	goto L31
L30:
	;
	if int32(0) <= v23 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v102 = int32(1)
	v103 = v96 - v102
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+v103<<(uint(v102)%32)))))
	if v107 != 0 {
		v111 = v94
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v111 = v90
	goto L30
L33:
	;
	v109 = v94 + int32(1)
	if v109 != v90 {
		v94 = v109
		v96 = v103
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v116 = v52
	goto L37
L36:
	;
	v116 = v50
	goto L37
L37:
	;
	v117 = v82<<(uint(int32(1))%32) + v116
	v123 = (v43 - (v82+v111)<<(uint(int32(1))%32)) & int32(-2)
	v129 = v123 - int32(1636608432)
	if v117&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	return v383 ^ v375 - base.I32_rotl(v383, int32(24)) ^ v83
L39:
	;
	v361 = int32(14)
	v363 = v357 ^ v358 - base.I32_rotl(v357, v361)
	v367 = v363 ^ v356 - base.I32_rotl(v363, int32(11))
	v371 = v367 ^ v357 - base.I32_rotl(v367, int32(25))
	v375 = v371 ^ v363 - base.I32_rotl(v371, int32(16))
	v379 = v375 ^ v367 - base.I32_rotl(v375, int32(4))
	v383 = v379 ^ v371 - base.I32_rotl(v379, v361)
	goto L38
L40:
	;
	switch v287 - int32(1) {
	case 0:
		v349 = v288
		v350 = v289
		v351 = v290
		goto L67
	case 1:
		v342 = v288
		v343 = v289
		v344 = v290
		goto L68
	case 2:
		v335 = v288
		v336 = v289
		v337 = v290
		goto L69
	case 3:
		v329 = v289
		v330 = v290
		goto L70
	case 4:
		v325 = v289
		v326 = v290
		goto L71
	case 5:
		v319 = v289
		v320 = v290
		goto L72
	case 6:
		v313 = v289
		v314 = v290
		goto L73
	case 7:
		v308 = v290
		goto L74
	case 8:
		v303 = v290
		goto L75
	case 9:
		v298 = v290
		goto L76
	case 10:
		goto L77
	default:
		v356 = v288
		v357 = v289
		v358 = v290
		goto L39
	}
L41:
	;
	v238 = v117
	v239 = v123
	v240 = v129
	v241 = v129
	v242 = v129
	goto L64
L42:
	;
	if base.Ui32(int32(11)) < base.Ui32(v123) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(v123) < base.Ui32(int32(12)) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v286 = v117
	v287 = v123
	v288 = v129
	v289 = v129
	v290 = v129
	goto L40
L46:
	;
	switch v185 - int32(1) {
	case 0:
		v235 = v186
		goto L53
	case 1:
		v230 = v186
		goto L54
	case 2:
		goto L55
	case 3:
		v223 = v187
		goto L56
	case 4:
		v220 = v187
		goto L57
	case 5:
		v215 = v187
		goto L58
	case 6:
		goto L59
	case 7:
		v206 = v188
		goto L60
	case 8:
		v201 = v188
		goto L61
	case 9:
		v196 = v188
		goto L62
	case 10:
		goto L63
	default:
		v356 = v186
		v357 = v187
		v358 = v188
		goto L39
	}
L47:
	;
	v184 = v117
	v185 = v123
	v186 = v129
	v187 = v129
	v188 = v129
	goto L46
L48:
	;
	goto L49
L49:
	;
	v136 = v117
	v137 = v123
	v138 = v129
	v139 = v129
	v140 = v129
	goto L50
L50:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v143 = v142 + v139
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v147 = v146 + v140
	v149 = int32(4)
	v151 = v144 + v138 - v147 ^ base.I32_rotl(v147, v149)
	v155 = v143 - v151 ^ base.I32_rotl(v151, int32(6))
	v156 = v147 + v143
	v157 = v151 + v156
	v158 = v155 + v157
	v162 = v156 - v155 ^ base.I32_rotl(v155, int32(8))
	v166 = v157 - v162 ^ base.I32_rotl(v162, int32(16))
	v170 = v158 - v166 ^ base.I32_rotl(v166, int32(19))
	v171 = v162 + v158
	v172 = v166 + v171
	v173 = v170 + v172
	v177 = v171 - v170 ^ base.I32_rotl(v170, v149)
	v178 = int32(12)
	v179 = v136 + v178
	v181 = v137 - v178
	if base.Ui32(int32(11)) < base.Ui32(v181) {
		v136 = v179
		v137 = v181
		v138 = v172
		v139 = v173
		v140 = v177
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v184 = v179
	v185 = v181
	v186 = v172
	v187 = v173
	v188 = v177
	goto L46
L52:
	;
	goto L51
L53:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v356 = v235 + v236
	v357 = v187
	v358 = v188
	goto L39
L54:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	v235 = v231<<(uint(int32(8))%32) + v230
	goto L53
L55:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+2)))
	v230 = v226<<(uint(int32(16))%32) + v186
	goto L54
L56:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v356 = v224 + v186
	v357 = v223
	v358 = v188
	goto L39
L57:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+4)))
	v223 = v220 + v221
	goto L56
L58:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+5)))
	v220 = v216<<(uint(int32(8))%32) + v215
	goto L57
L59:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+6)))
	v215 = v211<<(uint(int32(16))%32) + v187
	goto L58
L60:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v356 = v207 + v186
	v357 = v209 + v187
	v358 = v206
	goto L39
L61:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+8)))
	v206 = v202<<(uint(int32(8))%32) + v201
	goto L60
L62:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+9)))
	v201 = v197<<(uint(int32(16))%32) + v196
	goto L61
L63:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+10)))
	v196 = v192<<(uint(int32(24))%32) + v188
	goto L62
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v245 = v244 + v241
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v249 = v248 + v242
	v251 = int32(4)
	v253 = v246 + v240 - v249 ^ base.I32_rotl(v249, v251)
	v257 = v245 - v253 ^ base.I32_rotl(v253, int32(6))
	v258 = v249 + v245
	v259 = v253 + v258
	v260 = v257 + v259
	v264 = v258 - v257 ^ base.I32_rotl(v257, int32(8))
	v268 = v259 - v264 ^ base.I32_rotl(v264, int32(16))
	v272 = v260 - v268 ^ base.I32_rotl(v268, int32(19))
	v273 = v264 + v260
	v274 = v268 + v273
	v275 = v272 + v274
	v279 = v273 - v272 ^ base.I32_rotl(v272, v251)
	v280 = int32(12)
	v281 = v238 + v280
	v283 = v239 - v280
	if base.Ui32(int32(11)) < base.Ui32(v283) {
		v238 = v281
		v239 = v283
		v240 = v274
		v241 = v275
		v242 = v279
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v286 = v281
	v287 = v283
	v288 = v274
	v289 = v275
	v290 = v279
	goto L40
L66:
	;
	goto L65
L67:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	v356 = v349 + v352
	v357 = v350
	v358 = v351
	goto L39
L68:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	v349 = v345<<(uint(int32(8))%32) + v342
	v350 = v343
	v351 = v344
	goto L67
L69:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+2)))
	v342 = v338<<(uint(int32(16))%32) + v335
	v343 = v336
	v344 = v337
	goto L68
L70:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+3)))
	v335 = v331<<(uint(int32(24))%32) + v288
	v336 = v329
	v337 = v330
	goto L69
L71:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+4)))
	v329 = v325 + v327
	v330 = v326
	goto L70
L72:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+5)))
	v325 = v321<<(uint(int32(8))%32) + v319
	v326 = v320
	goto L71
L73:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+6)))
	v319 = v315<<(uint(int32(16))%32) + v313
	v320 = v314
	goto L72
L74:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+7)))
	v313 = v309<<(uint(int32(24))%32) + v289
	v314 = v308
	goto L73
L75:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+8)))
	v308 = v304<<(uint(int32(8))%32) + v303
	goto L74
L76:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+9)))
	v303 = v299<<(uint(int32(16))%32) + v298
	goto L75
L77:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+10)))
	v298 = v294<<(uint(int32(24))%32) + v290
	goto L76
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
												F_errmsg(m, int32(182234), v9+int32(16))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(477399), int32(1426), int32(387479))
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
								F_errmsg_internal(m, int32(357498), v9)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(477399), int32(1776), int32(384828))
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
														F_errmsg(m, int32(182234), v9+int32(16))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(477399), int32(1426), int32(387479))
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
							F_errmsg_internal(m, int32(357498), v9)
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(477399), int32(1776), int32(384828))
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
													F_errmsg(m, int32(182234), v9+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(477399), int32(1426), int32(387479))
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
	var v7 int64
	_ = v7
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
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
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	F_range_deserialize(m, v34, v14, v11+int32(40), v11+int32(32), v11+int32(31))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 == v21 {
		v34 = v23
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v27 = F_lookup_type_cache(m, v21, int32(2048))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
	v34 = v27
	goto L6
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v43)>>(uint(int32(2))%32))-int32(1)))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+200))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+164))
	if v51 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = F_lookup_type_cache(m, v54, int32(32768))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v61 = v50
	goto L16
L16:
	;
	if v49&int32(41) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+164))
	if v58 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v61 = v56
	goto L16
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v70 = F_FunctionCall2Coll(m, v61+int32(160), v68, v69, v18)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v73 = v7
	goto L21
L21:
	;
	if v49&int32(81) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	v73 = v72
	goto L21
L23:
	;
	v83 = v7
	goto L25
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v34)+208))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v80 = F_FunctionCall2Coll(m, v61+int32(160), v78, v79, v18)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
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
	v170 = F_Int64GetDatum(m, base.I64_extend_i32_u(v160)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v160^v152-base.I32_rotl(v160, int32(24))))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v137 = int32(14)
	v139 = v130 ^ v135 - base.I32_rotl(v130, v137)
	v144 = v139 ^ (v49 + v132) - base.I32_rotl(v139, int32(11))
	v148 = v144 ^ v130 - base.I32_rotl(v144, int32(25))
	v152 = v148 ^ v139 - base.I32_rotl(v148, int32(16))
	v156 = v152 ^ v144 - base.I32_rotl(v152, int32(4))
	v160 = v156 ^ v148 - base.I32_rotl(v156, v137)
	goto L27
L29:
	;
	v92 = int32(-1636608428)
	v130 = v92
	v132 = v92
	v135 = v92
	goto L28
L30:
	;
	goto L31
L31:
	;
	v95 = base.I32_wrap_i64(v84)
	v100 = base.I32_wrap_i64(int64(base.Ui64(v84)>>(uint(int64(32))%64))) ^ int32(-415931063)
	v106 = v95 - v100 - int32(1636608428) ^ base.I32_rotl(v100, int32(6))
	v108 = v95 + int32(1021750440)
	v109 = v100 + v108
	v110 = v106 + v109
	v114 = v108 - v106 ^ base.I32_rotl(v106, int32(8))
	v118 = v109 - v114 ^ base.I32_rotl(v114, int32(16))
	v122 = v110 - v118 ^ base.I32_rotl(v118, int32(19))
	v123 = v114 + v110
	v124 = v118 + v123
	v130 = v122 + v124
	v132 = v124
	v135 = v123 - v122 ^ base.I32_rotl(v122, int32(4))
	goto L28
L32:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	v173 = v172 ^ v73
	v184 = F_Int64GetDatum(m, v83^(v173<<(uint(int64(1))%64)&int64(-4294967298)|int64(base.Ui64(v173)>>(uint(int64(31))%64))&int64(4294967297)))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v11 + int32(48)
	return v184
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v21
	F_errmsg_internal(m, int32(357498), v11)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(477399), int32(1776), int32(384828))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v211 = F_format_type_be(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v211
	F_errmsg(m, int32(182234), v11+int32(16))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(477399), int32(1490), int32(445478))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
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
	F_check_stack_depth(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v30 = F_lookup_rowtype_tupdesc(m, v28, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v22
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v35
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+52)) = uint16(v35)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = int32(base.Ui32(v33) >> (uint(int32(2)) % 32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v45 == v35 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v66 != v28 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v56 = F_MemoryContextAlloc(m, v51, v32<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v48 < v32 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v65 = v45
	v66 = v50
	goto L5
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v61)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v32
	v65 = v61
	v66 = int32(0)
	goto L5
L10:
	;
	v112 = F_palloc(m, v106)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v73 = v32 << (uint(int32(2)) % 32)
	v75 = v65 + int32(20)
	if v75&int32(3) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v68 != v29 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v106 = v32 << (uint(int32(2)) % 32)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v28
	v106 = v73
	goto L10
L15:
	;
	v101 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), v73)
	mBase = m.M
	goto L23
L16:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v73) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v73+v75) <= base.Ui32(v75) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v85 = v65 + v73 + int32(20)
	v87 = v65 + int32(24)
	if base.Ui32(v87) < base.Ui32(v85) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = v85
	goto L21
L20:
	;
	v89 = v87
	goto L21
L21:
	;
	v98 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), (v89-v65-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L22
L22:
	;
	goto L14
L23:
	;
	goto L14
L24:
	;
	v114 = F_palloc(m, v32)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_heap_deform_tuple(m, v17+int32(-20), v30, v112, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v32 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_pfree(m, v112)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L55
	}
L28:
	;
	v234 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v121 = int32(20)
	v125 = int32(0)
	v128 = v125
	v131 = v125
	goto L31
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v149 = v30 + v121 + v143<<(uint(int32(4))%32) + v128*int32(100)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+91)))
	if v150 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L50
	}
L33:
	;
	goto L32
L34:
	;
	v154 = v128 << (uint(int32(2)) % 32)
	v155 = v65 + v121 + v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v156 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v203 = v131
	goto L36
L36:
	;
	v209 = v128 + int32(1)
	if v32 != v209 {
		v128 = v209
		v131 = v203
		goto L31
	} else {
		goto L49
	}
L37:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+v114))))
	if v175 != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v165 = F_lookup_type_cache(m, v163, int32(128))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L39:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v163 = v159
	goto L38
L40:
	;
	goto L41
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v160 == v161 {
		v171 = v156
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v163 = v160
	goto L38
L43:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)+136))
	if v167 == int32(0) {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v165
	v171 = v165
	goto L37
L45:
	;
	v198 = int32(0)
	goto L47
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v171 + int32(132)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v149)+96))
	v182 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+34)) = uint16(v182)
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)) = uint8(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v181
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v112+v154)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+40)) = uint8(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v188
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v171)+132))
	v195 = m.T0[v194].(func(*base.Module, int32) int32)(m, v17+int32(-48))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v203 = v198 + v131*int32(31)
	goto L36
L48:
	;
	v198 = v195
	goto L47
L49:
	;
	v234 = v203
	goto L27
L50:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v219 = F_format_type_be(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v219
	F_errmsg(m, int32(182234), v19)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(477369), int32(1894), int32(405587))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	F_pfree(m, v114)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if int32(0) <= v250 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_DecrTupleDescRefCount(m, v30)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v255 != v22 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	F_pfree(m, v22)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	m.G0 = v19 - int32(-64)
	return v234
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
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v23 = v19
	v24 = v20
	goto L5
L3:
	;
	v37 = v19
	goto L4
L4:
	;
	goto L1
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v26 = v25 + v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v27 != 0 {
		v23 = v26
		v24 = v27
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v37 = v26
	goto L4
L7:
	;
	v29 = v24
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	if v32 != 0 {
		v23 = v26
		v24 = v32
		goto L5
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v33 != v15 {
		v29 = v33
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v69 != 0 {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	v47 = v43
	v48 = v44
	goto L16
L14:
	;
	v61 = v43
	goto L15
L15:
	;
	goto L12
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v50 = v49 + v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v51 != 0 {
		v47 = v50
		v48 = v51
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v61 = v50
	goto L15
L18:
	;
	v53 = v48
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	if v56 != 0 {
		v47 = v50
		v48 = v56
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v57 != v39 {
		v53 = v57
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+320))
	if v88 == int64(0) {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v72 = v68
	v73 = v69
	goto L27
L25:
	;
	v86 = v68
	goto L26
L26:
	;
	goto L23
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v75 = v74 + v72
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	if v76 != 0 {
		v72 = v75
		v73 = v76
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v86 = v75
	goto L26
L29:
	;
	v78 = v73
	goto L30
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+28))
	if v81 != 0 {
		v72 = v75
		v73 = v81
		goto L27
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v82 != v64 {
		v78 = v82
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
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if base.Ui32(v37+v61+v86) <= base.Ui32(v91) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+288))
	if base.Ui64(v13) <= base.Ui64(v95) {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)) = uint8(v97)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v102 != int32(2) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v105 = int32(48)
	goto L42
L41:
	;
	v105 = int32(0)
	goto L42
L42:
	;
	v106 = v99 + v105
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+240)))
	v116 = v106 + (v107<<(uint(int32(3))%32)|int32(4))&int32(12) + int32(32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v120 = int32(1)
	v121 = v107 ^ v120
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v121&v120 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v143 = v117
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+28)) = v143
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+276)))
	if v148 != 0 {
		goto L34
	} else {
		goto L51
	}
L46:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(1583036)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v135 = int32(1)
	v137 = F_ExecBuildAggTrans(m, l0, v106, v121&base.B2i32(v102 == int32(3)), v135, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v137
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+97)) = uint8(v125)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v126
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v143 = v142
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
