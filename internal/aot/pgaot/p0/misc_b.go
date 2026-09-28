package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BecomeLockGroupLeader(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupLeader[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+364))
	if v7 != v6 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupLeader[1]))
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupLeader[2]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v16 = base.I32_div_s(v6-v13, int32(768))
		v18 = base.I32_rem_s(v16, int32(16))
		v23 = v10 + v18<<(uint(int32(7))%32) + int32(_a_F_BecomeLockGroupLeader_0)
		v25 = F_LWLockAcquire(m, v23, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupLeader[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v28)+364)) = v28
			v31 = v28 + int32(368)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+372))
			if v32 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v31
				v36 = v31
			} else {
				v36 = v32
			}
			*(*int32)(unsafe.Add(mBase, uint32(v28)+376)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v28)+380)) = v36
			v40 = v28 + int32(376)
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v28)+372)) = v40
			F_LWLockRelease(m, v23)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_BlockSampler_Init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = l0 + int32(16)
	v11 = base.I64_extend_i32_u(l3)
	v14 = v11 + int64(4354685564936845354)
	v15 = int64(30)
	v18 = int64(-4658895280553007687)
	v19 = (int64(base.Ui64(v14)>>(uint(v15)%64)) ^ v14) * v18
	v20 = int64(27)
	v23 = int64(-7723592293110705685)
	v24 = (int64(base.Ui64(v19)>>(uint(v20)%64)) ^ v19) * v23
	v25 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(base.Ui64(v24)>>(uint(v25)%64)) ^ v24
	v30 = v11 - int64(7046029254386353131)
	v35 = (int64(base.Ui64(v30)>>(uint(v15)%64)) ^ v30) * v18
	v40 = (int64(base.Ui64(v35)>>(uint(v20)%64)) ^ v35) * v23
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(base.Ui64(v40)>>(uint(v25)%64)) ^ v40
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v45) < base.Ui32(v46) {
		v48 = v45
	} else {
		v48 = v46
	}
	return v48
}
func F_base32hex_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
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
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v128 int64
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v4 = int32(0)
	v11 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L22
	} else {
		goto L43
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L22
	} else {
		goto L38
	}
L3:
	;
	v16 = l0 + l1
	v17 = l0
	v21 = v4
	v23 = v4
	v24 = v4
	v25 = v4
	v27 = v11
	goto L6
L4:
	;
	v128 = v11
	goto L5
L5:
	;
	m.G0 = v14 + int32(32)
	return v128
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v30 = v28 - int32(9)
	if int32(1)<<(uint(v30)%32)&int32(_a_F_base32hex_decode_0) != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v128 = v114
	goto L5
L8:
	;
	v116 = v17 + int32(1)
	if base.Ui32(v116) < base.Ui32(v16) {
		v17 = v116
		v21 = v109
		v23 = v111
		v24 = v112
		v25 = v113
		v27 = v114
		goto L6
	} else {
		goto L37
	}
L9:
	;
	v38 = base.B2i32(base.Ui32(v30) <= base.Ui32(int32(23)))
	goto L11
L10:
	;
	v38 = int32(0)
	goto L11
L11:
	;
	if v38 != 0 {
		v109 = v21
		v111 = v23
		v112 = v24
		v113 = v25
		v114 = v27
		goto L8
	} else {
		goto L12
	}
L12:
	;
	if v28 == int32(61) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if int32(1)<<(uint(v21)%32)&int32(180) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v25 != 0 {
		goto L2
	} else {
		goto L27
	}
L16:
	;
	v48 = base.B2i32(base.Ui32(v21) <= base.Ui32(int32(7)))
	goto L18
L17:
	;
	v48 = int32(0)
	goto L18
L18:
	;
	if v25|v48 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v70 = int32(1)
	v109 = v21 + v70
	v111 = v23
	v112 = v24
	v113 = v70
	v114 = v27
	goto L8
L22:
	;
	return int64(0)
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_base32hex_decode_1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_base32hex_decode_2), int32(929), int32(_a_F_base32hex_decode_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	if base.I32_extend8_s(v28) < int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_base32hex_decode[0]))))
	if v76 < int32(0) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v81 = v24<<(uint(int32(5))%32) | v76
	if v23 < int32(3) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v102 = int32(0)
	v104 = v21 + int32(1)
	if v104 != int32(8) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v99 = v23 + int32(5)
	v100 = v81
	v101 = v27
	goto L30
L32:
	;
	goto L33
L33:
	;
	v89 = v23 - int32(3)
	v90 = int32(base.Ui32(v81) >> (uint(v89) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(l2+base.I32_wrap_i64(v27)))) = uint8(v90)
	v94 = int32(-1)
	v99 = v89
	v100 = v81 & (v94<<(uint(v89)%32) ^ v94)
	v101 = v27 + int64(1)
	goto L30
L34:
	;
	v108 = v104
	goto L36
L35:
	;
	v108 = v102
	goto L36
L36:
	;
	v109 = v108
	v111 = v99
	v112 = v100
	v113 = v102
	v114 = v101
	goto L8
L37:
	;
	goto L7
L38:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	v140 = F_pg_mblen_range(m, v17, v16)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v140
	F_errmsg(m, int32(_a_F_base32hex_decode_4), v14)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_base32hex_decode_2), int32(941), int32(_a_F_base32hex_decode_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v160 = F_pg_mblen_range(m, v17, v16)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v160
	F_errmsg(m, int32(_a_F_base32hex_decode_4), v14+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_base32hex_decode_2), int32(951), int32(_a_F_base32hex_decode_3))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_base32hex_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v150 int64
	_ = v150
	v4 = int32(0)
	v9 = int64(0)
	if l1 == v4 {
		v150 = v9
	} else {
		v15 = v4
		v17 = v4
		v18 = v4
		v20 = v9
		for {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v18))))
			v23 = int32(8)
			v25 = v22 | v17<<(uint(v23)%32)
			v27 = v15 + v23
			if v15 < int32(-3) {
				v99 = v27
				v101 = v25
				v104 = v20
			} else {
				v31 = v15 + int32(3)
				v33 = base.I32_div_u_s(v31, int32(5))
				if v33&int32(1) == int32(0) {
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v25)>>(uint(v31)%32))&int32(31))+uint32(_c_F_base32hex_encode[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l2+base.I32_wrap_i64(v20)))) = uint8(v43)
					v47 = int32(-1)
					v52 = v31
					v53 = v25 & (v47<<(uint(v31)%32) ^ v47)
					v54 = v20 + int64(1)
				} else {
					v52 = v27
					v53 = v25
					v54 = v20
				}
				if base.Ui32(v31) < base.Ui32(int32(5)) {
					v99 = v31
					v101 = v53
					v104 = v54
				} else {
					v61 = v52
					v62 = v53
					v65 = v54
					for {
						v67 = l2 + base.I32_wrap_i64(v65)
						v69 = v61 - int32(5)
						v71 = int32(31)
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v62)>>(uint(v69)%32))&v71)+uint32(_c_F_base32hex_encode[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v73)
						v75 = int32(-1)
						v79 = v62 & (v75<<(uint(v69)%32) ^ v75)
						v81 = v61 - int32(10)
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v79)>>(uint(v81)%32))&v71)+uint32(_c_F_base32hex_encode[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)) = uint8(v85)
						v91 = v79 & (v75<<(uint(v81)%32) ^ v75)
						v93 = v65 + int64(2)
						if int32(14) < v61 {
							v61 = v81
							v62 = v91
							v65 = v93
							continue
						} else {
							break
						}
						break
					}
					v99 = v81
					v101 = v91
					v104 = v93
				}
			}
			v106 = v18 + int32(1)
			if v106 != l1 {
				v15 = v99
				v17 = v101
				v18 = v106
				v20 = v104
				continue
			} else {
				break
			}
			break
		}
		if int32(0) < v99 {
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101<<(uint(int32(5)-v99)%32)&int32(31))+uint32(_c_F_base32hex_encode[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l2+base.I32_wrap_i64(v104)))) = uint8(v117)
			v121 = v104 + int64(1)
		} else {
			v121 = v104
		}
		if v121&int64(7) == int64(0) {
			v150 = v121
		} else {
			v126 = base.I32_wrap_i64(v121)
			v127 = int32(7)
			v128 = v126 ^ v127
			v132 = v128&v127 + int32(1)
			if v132 != 0 {
				base.MemoryFill(m, v126+l2, int32(61), v132)
			} else {
			}
			v150 = v121 + base.I64_extend_i32_u(v128)&int64(7) + int64(1)
		}
	}
	return v150
}
func F_basque_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v560 int32
	_ = v560
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v683 int32
	_ = v683
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v801 int32
	_ = v801
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v920 int32
	_ = v920
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v964 int32
	_ = v964
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1041 int32
	_ = v1041
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1079 int32
	_ = v1079
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1226 int32
	_ = v1226
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1259 int32
	_ = v1259
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1297 int32
	_ = v1297
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1348 int32
	_ = v1348
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1382 int32
	_ = v1382
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1421 int32
	_ = v1421
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1472 int32
	_ = v1472
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1505 int32
	_ = v1505
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1543 int32
	_ = v1543
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1594 int32
	_ = v1594
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1628 int32
	_ = v1628
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L7
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1175 = v10
	goto L263
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1148
	goto L1
L3:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1148 = v1146 + v1144
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L136
L5:
	;
	if v128 != 0 {
		goto L4
	} else {
		goto L29
	}
L6:
	;
	v128 = v121
	goto L5
L7:
	;
	if v6 <= v10 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v121 = int32(0)
	goto L6
L9:
	;
	v128 = int32(-1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v24))))
	if base.Ui32(v41) < base.Ui32(int32(192)) {
		v98 = v41
		v99 = v39
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if int32(117) < v98 {
		v121 = v99
		goto L6
	} else {
		goto L25
	}
L13:
	;
	v45 = v10 + int32(1)
	if v45 == v6 {
		v98 = v41
		v99 = v39
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v24))))
	v50 = v48 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v41) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v24))))
	v66 = v64 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v41) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v54 = v10 + int32(2)
	if v54 != v6 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v98 = v41<<(uint(int32(6))%32)&int32(1984) | v50
	v99 = int32(2)
	goto L12
L19:
	;
	goto L18
L20:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v70))))
	v98 = v83&int32(63) | (v41<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v50<<(uint(int32(12))%32) | v66<<(uint(int32(6))%32))
	v99 = int32(4)
	goto L12
L21:
	;
	v70 = v10 + int32(3)
	if v70 != v6 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v98 = v41<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v50<<(uint(int32(6))%32) | v66
	v99 = int32(3)
	goto L12
L24:
	;
	goto L23
L25:
	;
	v103 = v98 - int32(97)
	if v103 < int32(0) {
		v121 = v99
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v103)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v109)>>(uint(v103&int32(7))%32))&int32(1) == int32(0) {
		v121 = v99
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99 + v10
	goto L28
L28:
	;
	goto L8
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L30:
	;
	if v246 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L31:
	;
	v246 = v239
	goto L30
L32:
	;
	if v142 <= v129 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v239 = int32(0)
	goto L31
L34:
	;
	v246 = int32(-1)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v158 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v143))))
	if base.Ui32(v160) < base.Ui32(int32(192)) {
		v217 = v160
		v218 = v158
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if int32(117) < v217 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v164 = v129 + int32(1)
	if v164 == v142 {
		v217 = v160
		v218 = v158
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v143))))
	v169 = v167 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v160) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v143))))
	v185 = v183 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v160) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v173 = v129 + int32(2)
	if v173 != v142 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v217 = v160<<(uint(int32(6))%32)&int32(1984) | v169
	v218 = int32(2)
	goto L37
L44:
	;
	goto L43
L45:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v189))))
	v217 = v202&int32(63) | (v160<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v169<<(uint(int32(12))%32) | v185<<(uint(int32(6))%32))
	v218 = int32(4)
	goto L37
L46:
	;
	v189 = v129 + int32(3)
	if v189 != v142 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v217 = v160<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v169<<(uint(int32(6))%32) | v185
	v218 = int32(3)
	goto L37
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v218 + v129
	goto L54
L51:
	;
	v222 = v217 - int32(97)
	if v222 < int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v222)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v228)>>(uint(v222&int32(7))%32))&int32(1) != 0 {
		v239 = v218
		goto L31
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L33
L55:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v270 = v260
	goto L60
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v129
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L58:
	;
	if int32(0) <= v365 {
		v1144 = v365
		goto L3
	} else {
		goto L83
	}
L59:
	;
	v365 = v337
	goto L58
L60:
	;
	if v261 <= v270 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v365 = int32(-1)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v277 = int32(1)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v262))))
	if base.Ui32(v279) < base.Ui32(int32(192)) {
		v336 = v279
		v337 = v277
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if int32(117) < v336 {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v283 = v270 + int32(1)
	if v283 == v261 {
		v336 = v279
		v337 = v277
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v262))))
	v288 = v286 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v279) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v262))))
	v304 = v302 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v279) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v292 = v270 + int32(2)
	if v292 != v261 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v336 = v279<<(uint(int32(6))%32)&int32(1984) | v288
	v337 = int32(2)
	goto L65
L72:
	;
	goto L71
L73:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v308))))
	v336 = v321&int32(63) | (v279<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v288<<(uint(int32(12))%32) | v304<<(uint(int32(6))%32))
	v337 = int32(4)
	goto L65
L74:
	;
	v308 = v270 + int32(3)
	if v308 != v261 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v336 = v279<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v288<<(uint(int32(6))%32) | v304
	v337 = int32(3)
	goto L65
L77:
	;
	goto L76
L78:
	;
	v354 = v337 + v270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v270 = v354
	goto L60
L79:
	;
	v341 = v336 - int32(97)
	if v341 < int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v341)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v347)>>(uint(v341&int32(7))%32))&int32(1) != 0 {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	goto L78
L83:
	;
	goto L57
L84:
	;
	if v487 != 0 {
		goto L4
	} else {
		goto L108
	}
L85:
	;
	v487 = v480
	goto L84
L86:
	;
	if v382 <= v129 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v480 = int32(0)
	goto L85
L88:
	;
	v487 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v398 = int32(1)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v383))))
	if base.Ui32(v400) < base.Ui32(int32(192)) {
		v457 = v400
		v458 = v398
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(117) < v457 {
		v480 = v458
		goto L85
	} else {
		goto L104
	}
L92:
	;
	v404 = v129 + int32(1)
	if v404 == v382 {
		v457 = v400
		v458 = v398
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v383))))
	v409 = v407 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v400) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v383))))
	v425 = v423 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v400) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v413 = v129 + int32(2)
	if v413 != v382 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v457 = v400<<(uint(int32(6))%32)&int32(1984) | v409
	v458 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+v429))))
	v457 = v442&int32(63) | (v400<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v409<<(uint(int32(12))%32) | v425<<(uint(int32(6))%32))
	v458 = int32(4)
	goto L91
L100:
	;
	v429 = v129 + int32(3)
	if v429 != v382 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v457 = v400<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v409<<(uint(int32(6))%32) | v425
	v458 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	v462 = v457 - int32(97)
	if v462 < int32(0) {
		v480 = v458
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v462)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v468)>>(uint(v462&int32(7))%32))&int32(1) == int32(0) {
		v480 = v458
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v458 + v129
	goto L107
L107:
	;
	goto L87
L108:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v509 = v499
	goto L111
L109:
	;
	if int32(0) <= v605 {
		v1144 = v605
		goto L3
	} else {
		goto L133
	}
L110:
	;
	v605 = v576
	goto L109
L111:
	;
	if v500 <= v509 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v605 = int32(-1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v516 = int32(1)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v501))))
	if base.Ui32(v518) < base.Ui32(int32(192)) {
		v575 = v518
		v576 = v516
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if int32(117) < v575 {
		goto L110
	} else {
		goto L129
	}
L117:
	;
	v522 = v509 + int32(1)
	if v522 == v500 {
		v575 = v518
		v576 = v516
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522+v501))))
	v527 = v525 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v518) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531+v501))))
	v543 = v541 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v518) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v531 = v509 + int32(2)
	if v531 != v500 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v575 = v518<<(uint(int32(6))%32)&int32(1984) | v527
	v576 = int32(2)
	goto L116
L123:
	;
	goto L122
L124:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v547))))
	v575 = v560&int32(63) | (v518<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v527<<(uint(int32(12))%32) | v543<<(uint(int32(6))%32))
	v576 = int32(4)
	goto L116
L125:
	;
	v547 = v509 + int32(3)
	if v547 != v500 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v575 = v518<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v527<<(uint(int32(6))%32) | v543
	v576 = int32(3)
	goto L116
L128:
	;
	goto L127
L129:
	;
	v580 = v575 - int32(97)
	if v580 < int32(0) {
		goto L110
	} else {
		goto L130
	}
L130:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v580)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v586)>>(uint(v580&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L131
	}
L131:
	;
	v594 = v576 + v509
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v594
	v509 = v594
	goto L111
L133:
	;
	goto L4
L134:
	;
	if v727 != 0 {
		goto L1
	} else {
		goto L159
	}
L135:
	;
	v727 = v720
	goto L134
L136:
	;
	if v623 <= v10 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v720 = int32(0)
	goto L135
L138:
	;
	v727 = int32(-1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v639 = int32(1)
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v624))))
	if base.Ui32(v641) < base.Ui32(int32(192)) {
		v698 = v641
		v699 = v639
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if int32(117) < v698 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v645 = v10 + int32(1)
	if v645 == v623 {
		v698 = v641
		v699 = v639
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v645+v624))))
	v650 = v648 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v641) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654+v624))))
	v666 = v664 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v641) {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v654 = v10 + int32(2)
	if v654 != v623 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v698 = v641<<(uint(int32(6))%32)&int32(1984) | v650
	v699 = int32(2)
	goto L141
L148:
	;
	goto L147
L149:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624+v670))))
	v698 = v683&int32(63) | (v641<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v650<<(uint(int32(12))%32) | v666<<(uint(int32(6))%32))
	v699 = int32(4)
	goto L141
L150:
	;
	v670 = v10 + int32(3)
	if v670 != v623 {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v698 = v641<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v650<<(uint(int32(6))%32) | v666
	v699 = int32(3)
	goto L141
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v699 + v10
	goto L158
L155:
	;
	v703 = v698 - int32(97)
	if v703 < int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v703)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v709)>>(uint(v703&int32(7))%32))&int32(1) != 0 {
		v720 = v699
		goto L135
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	goto L137
L159:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L162
L160:
	;
	if v845 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L161:
	;
	v845 = v838
	goto L160
L162:
	;
	if v741 <= v728 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v838 = int32(0)
	goto L161
L164:
	;
	v845 = int32(-1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v757 = int32(1)
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728+v742))))
	if base.Ui32(v759) < base.Ui32(int32(192)) {
		v816 = v759
		v817 = v757
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if int32(117) < v816 {
		goto L180
	} else {
		goto L181
	}
L168:
	;
	v763 = v728 + int32(1)
	if v763 == v741 {
		v816 = v759
		v817 = v757
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763+v742))))
	v768 = v766 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v759) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772+v742))))
	v784 = v782 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v759) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v772 = v728 + int32(2)
	if v772 != v741 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v816 = v759<<(uint(int32(6))%32)&int32(1984) | v768
	v817 = int32(2)
	goto L167
L174:
	;
	goto L173
L175:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742+v788))))
	v816 = v801&int32(63) | (v759<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v768<<(uint(int32(12))%32) | v784<<(uint(int32(6))%32))
	v817 = int32(4)
	goto L167
L176:
	;
	v788 = v728 + int32(3)
	if v788 != v741 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v816 = v759<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v768<<(uint(int32(6))%32) | v784
	v817 = int32(3)
	goto L167
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v817 + v728
	goto L184
L181:
	;
	v821 = v816 - int32(97)
	if v821 < int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v821)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v827)>>(uint(v821&int32(7))%32))&int32(1) != 0 {
		v838 = v817
		goto L161
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	goto L163
L185:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v869 = v859
	goto L190
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v728
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L216
L188:
	;
	if int32(0) <= v964 {
		v1144 = v964
		goto L3
	} else {
		goto L213
	}
L189:
	;
	v964 = v936
	goto L188
L190:
	;
	if v860 <= v869 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v964 = int32(-1)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v876 = int32(1)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869+v861))))
	if base.Ui32(v878) < base.Ui32(int32(192)) {
		v935 = v878
		v936 = v876
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if int32(117) < v935 {
		goto L208
	} else {
		goto L209
	}
L196:
	;
	v882 = v869 + int32(1)
	if v882 == v860 {
		v935 = v878
		v936 = v876
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882+v861))))
	v887 = v885 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v878) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891+v861))))
	v903 = v901 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v878) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v891 = v869 + int32(2)
	if v891 != v860 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v935 = v878<<(uint(int32(6))%32)&int32(1984) | v887
	v936 = int32(2)
	goto L195
L202:
	;
	goto L201
L203:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861+v907))))
	v935 = v920&int32(63) | (v878<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v887<<(uint(int32(12))%32) | v903<<(uint(int32(6))%32))
	v936 = int32(4)
	goto L195
L204:
	;
	v907 = v869 + int32(3)
	if v907 != v860 {
		goto L203
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v935 = v878<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v887<<(uint(int32(6))%32) | v903
	v936 = int32(3)
	goto L195
L207:
	;
	goto L206
L208:
	;
	v953 = v936 + v869
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v953
	v869 = v953
	goto L190
L209:
	;
	v940 = v935 - int32(97)
	if v940 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v940)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v946)>>(uint(v940&int32(7))%32))&int32(1) != 0 {
		goto L189
	} else {
		goto L211
	}
L211:
	;
	goto L208
L213:
	;
	goto L187
L214:
	;
	if v1086 != 0 {
		goto L1
	} else {
		goto L238
	}
L215:
	;
	v1086 = v1079
	goto L214
L216:
	;
	if v981 <= v728 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1079 = int32(0)
	goto L215
L218:
	;
	v1086 = int32(-1)
	goto L214
L219:
	;
	goto L220
L220:
	;
	v997 = int32(1)
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728+v982))))
	if base.Ui32(v999) < base.Ui32(int32(192)) {
		v1056 = v999
		v1057 = v997
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if int32(117) < v1056 {
		v1079 = v1057
		goto L215
	} else {
		goto L234
	}
L222:
	;
	v1003 = v728 + int32(1)
	if v1003 == v981 {
		v1056 = v999
		v1057 = v997
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003+v982))))
	v1008 = v1006 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v999) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012+v982))))
	v1024 = v1022 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v999) {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v1012 = v728 + int32(2)
	if v1012 != v981 {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1056 = v999<<(uint(int32(6))%32)&int32(1984) | v1008
	v1057 = int32(2)
	goto L221
L228:
	;
	goto L227
L229:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982+v1028))))
	v1056 = v1041&int32(63) | (v999<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1008<<(uint(int32(12))%32) | v1024<<(uint(int32(6))%32))
	v1057 = int32(4)
	goto L221
L230:
	;
	v1028 = v728 + int32(3)
	if v1028 != v981 {
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1056 = v999<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1008<<(uint(int32(6))%32) | v1024
	v1057 = int32(3)
	goto L221
L233:
	;
	goto L232
L234:
	;
	v1061 = v1056 - int32(97)
	if v1061 < int32(0) {
		v1079 = v1057
		goto L215
	} else {
		goto L235
	}
L235:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1061)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v1067)>>(uint(v1061&int32(7))%32))&int32(1) == int32(0) {
		v1079 = v1057
		goto L215
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1057 + v728
	goto L237
L237:
	;
	goto L217
L238:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L241
L239:
	;
	if int32(0) <= v1141 {
		v1148 = v1141
		goto L2
	} else {
		goto L259
	}
L241:
	;
	goto L242
L242:
	;
	goto L243
L243:
	;
	v1096 = v1088
	v1098 = int32(1)
	goto L246
L245:
	;
	v1141 = v1126
	goto L239
L246:
	;
	if v1089 <= v1096 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	goto L245
L248:
	;
	v1141 = int32(-1)
	goto L239
L249:
	;
	goto L250
L250:
	;
	v1103 = v1096 + int32(1)
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087+v1096))))
	if base.Ui32(v1105) < base.Ui32(int32(192)) {
		v1126 = v1103
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1127 = int32(1)
	if v1127 < v1098 {
		v1096 = v1126
		v1098 = v1098 - v1127
		goto L246
	} else {
		goto L258
	}
L252:
	;
	if v1089 <= v1103 {
		v1126 = v1103
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1112 = v1103
	goto L254
L254:
	;
	v1115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1087+v1112))))
	if int32(-65) < v1115 {
		v1126 = v1112
		goto L251
	} else {
		goto L256
	}
L255:
	;
	v1126 = v1089
	goto L251
L256:
	;
	v1119 = v1112 + int32(1)
	if v1119 != v1089 {
		v1112 = v1119
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	goto L247
L259:
	;
	goto L1
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1647
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1647
	v1651 = v1647 - int32(1)
	if v1651 <= v10 {
		goto L365
	} else {
		goto L366
	}
L261:
	;
	if v1270 < int32(0) {
		goto L260
	} else {
		goto L286
	}
L262:
	;
	v1270 = v1242
	goto L261
L263:
	;
	if v1166 <= v1175 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1270 = int32(-1)
	goto L261
L266:
	;
	goto L267
L267:
	;
	v1182 = int32(1)
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175+v1167))))
	if base.Ui32(v1184) < base.Ui32(int32(192)) {
		v1241 = v1184
		v1242 = v1182
		goto L268
	} else {
		goto L269
	}
L268:
	;
	if int32(117) < v1241 {
		goto L281
	} else {
		goto L282
	}
L269:
	;
	v1188 = v1175 + int32(1)
	if v1188 == v1166 {
		v1241 = v1184
		v1242 = v1182
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1188+v1167))))
	v1193 = v1191 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1184) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197+v1167))))
	v1209 = v1207 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1184) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v1197 = v1175 + int32(2)
	if v1197 != v1166 {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1241 = v1184<<(uint(int32(6))%32)&int32(1984) | v1193
	v1242 = int32(2)
	goto L268
L275:
	;
	goto L274
L276:
	;
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1167+v1213))))
	v1241 = v1226&int32(63) | (v1184<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1193<<(uint(int32(12))%32) | v1209<<(uint(int32(6))%32))
	v1242 = int32(4)
	goto L268
L277:
	;
	v1213 = v1175 + int32(3)
	if v1213 != v1166 {
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1241 = v1184<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1193<<(uint(int32(6))%32) | v1209
	v1242 = int32(3)
	goto L268
L280:
	;
	goto L279
L281:
	;
	v1259 = v1242 + v1175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1259
	v1175 = v1259
	goto L263
L282:
	;
	v1246 = v1241 - int32(97)
	if v1246 < int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1246)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v1252)>>(uint(v1246&int32(7))%32))&int32(1) != 0 {
		goto L262
	} else {
		goto L284
	}
L284:
	;
	goto L281
L286:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1274 = v1273 + v1270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1274
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1297 = v1274
	goto L289
L287:
	;
	if v1393 < int32(0) {
		goto L260
	} else {
		goto L311
	}
L288:
	;
	v1393 = v1364
	goto L287
L289:
	;
	if v1288 <= v1297 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1393 = int32(-1)
	goto L287
L292:
	;
	goto L293
L293:
	;
	v1304 = int32(1)
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297+v1289))))
	if base.Ui32(v1306) < base.Ui32(int32(192)) {
		v1363 = v1306
		v1364 = v1304
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(117) < v1363 {
		goto L288
	} else {
		goto L307
	}
L295:
	;
	v1310 = v1297 + int32(1)
	if v1310 == v1288 {
		v1363 = v1306
		v1364 = v1304
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310+v1289))))
	v1315 = v1313 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1306) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319+v1289))))
	v1331 = v1329 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1306) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	v1319 = v1297 + int32(2)
	if v1319 != v1288 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1363 = v1306<<(uint(int32(6))%32)&int32(1984) | v1315
	v1364 = int32(2)
	goto L294
L301:
	;
	goto L300
L302:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289+v1335))))
	v1363 = v1348&int32(63) | (v1306<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1315<<(uint(int32(12))%32) | v1331<<(uint(int32(6))%32))
	v1364 = int32(4)
	goto L294
L303:
	;
	v1335 = v1297 + int32(3)
	if v1335 != v1288 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1363 = v1306<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1315<<(uint(int32(6))%32) | v1331
	v1364 = int32(3)
	goto L294
L306:
	;
	goto L305
L307:
	;
	v1368 = v1363 - int32(97)
	if v1368 < int32(0) {
		goto L288
	} else {
		goto L308
	}
L308:
	;
	v1374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1368)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v1374)>>(uint(v1368&int32(7))%32))&int32(1) == int32(0) {
		goto L288
	} else {
		goto L309
	}
L309:
	;
	v1382 = v1364 + v1297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1382
	v1297 = v1382
	goto L289
L311:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1397 = v1396 + v1393
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1421 = v1397
	goto L314
L312:
	;
	if v1516 < int32(0) {
		goto L260
	} else {
		goto L337
	}
L313:
	;
	v1516 = v1488
	goto L312
L314:
	;
	if v1412 <= v1421 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1516 = int32(-1)
	goto L312
L317:
	;
	goto L318
L318:
	;
	v1428 = int32(1)
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421+v1413))))
	if base.Ui32(v1430) < base.Ui32(int32(192)) {
		v1487 = v1430
		v1488 = v1428
		goto L319
	} else {
		goto L320
	}
L319:
	;
	if int32(117) < v1487 {
		goto L332
	} else {
		goto L333
	}
L320:
	;
	v1434 = v1421 + int32(1)
	if v1434 == v1412 {
		v1487 = v1430
		v1488 = v1428
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434+v1413))))
	v1439 = v1437 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1430) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443+v1413))))
	v1455 = v1453 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1430) {
		goto L328
	} else {
		goto L329
	}
L323:
	;
	v1443 = v1421 + int32(2)
	if v1443 != v1412 {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1487 = v1430<<(uint(int32(6))%32)&int32(1984) | v1439
	v1488 = int32(2)
	goto L319
L326:
	;
	goto L325
L327:
	;
	v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1413+v1459))))
	v1487 = v1472&int32(63) | (v1430<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1439<<(uint(int32(12))%32) | v1455<<(uint(int32(6))%32))
	v1488 = int32(4)
	goto L319
L328:
	;
	v1459 = v1421 + int32(3)
	if v1459 != v1412 {
		goto L327
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1487 = v1430<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1439<<(uint(int32(6))%32) | v1455
	v1488 = int32(3)
	goto L319
L331:
	;
	goto L330
L332:
	;
	v1505 = v1488 + v1421
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1505
	v1421 = v1505
	goto L314
L333:
	;
	v1492 = v1487 - int32(97)
	if v1492 < int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1492)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v1498)>>(uint(v1492&int32(7))%32))&int32(1) != 0 {
		goto L313
	} else {
		goto L335
	}
L335:
	;
	goto L332
L337:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1520 = v1519 + v1516
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1520
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1543 = v1520
	goto L340
L338:
	;
	if v1639 < int32(0) {
		goto L260
	} else {
		goto L362
	}
L339:
	;
	v1639 = v1610
	goto L338
L340:
	;
	if v1534 <= v1543 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1639 = int32(-1)
	goto L338
L343:
	;
	goto L344
L344:
	;
	v1550 = int32(1)
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1543+v1535))))
	if base.Ui32(v1552) < base.Ui32(int32(192)) {
		v1609 = v1552
		v1610 = v1550
		goto L345
	} else {
		goto L346
	}
L345:
	;
	if int32(117) < v1609 {
		goto L339
	} else {
		goto L358
	}
L346:
	;
	v1556 = v1543 + int32(1)
	if v1556 == v1534 {
		v1609 = v1552
		v1610 = v1550
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556+v1535))))
	v1561 = v1559 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1552) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565+v1535))))
	v1577 = v1575 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1552) {
		goto L354
	} else {
		goto L355
	}
L349:
	;
	v1565 = v1543 + int32(2)
	if v1565 != v1534 {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1609 = v1552<<(uint(int32(6))%32)&int32(1984) | v1561
	v1610 = int32(2)
	goto L345
L352:
	;
	goto L351
L353:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535+v1581))))
	v1609 = v1594&int32(63) | (v1552<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1561<<(uint(int32(12))%32) | v1577<<(uint(int32(6))%32))
	v1610 = int32(4)
	goto L345
L354:
	;
	v1581 = v1543 + int32(3)
	if v1581 != v1534 {
		goto L353
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1609 = v1552<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1561<<(uint(int32(6))%32) | v1577
	v1610 = int32(3)
	goto L345
L357:
	;
	goto L356
L358:
	;
	v1614 = v1609 - int32(97)
	if v1614 < int32(0) {
		goto L339
	} else {
		goto L359
	}
L359:
	;
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1614)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
	if int32(base.Ui32(v1620)>>(uint(v1614&int32(7))%32))&int32(1) == int32(0) {
		goto L339
	} else {
		goto L360
	}
L360:
	;
	v1628 = v1610 + v1543
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1628
	v1543 = v1628
	goto L340
L362:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1642 + v1639
	goto L260
L363:
	;
	return v1842
L364:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1710 = v1708 + (v1704 - v1705)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1710
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1710
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1710 <= v1713 {
		v1787 = v1710
		v1790 = v1708
		goto L382
	} else {
		goto L383
	}
L365:
	;
	v1704 = v1647
	v1705 = v1647
	goto L364
L366:
	;
	goto L367
L367:
	;
	v1654 = v1647
	v1655 = v1647
	v1657 = v1651
	goto L368
L368:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1658+v1657))))
	if base.B2i32(v1660&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1660)%32)&int32(70566434) == int32(0)) != 0 {
		v1704 = v1654
		v1705 = v1655
		goto L364
	} else {
		goto L370
	}
L369:
	;
	v1704 = v1696
	v1705 = v1698
	goto L364
L370:
	;
	v1675 = F_find_among_b(m, l0, int32(_a_F_basque_UTF_8_stem_2), int32(109), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	return int32(0)
L372:
	;
	if v1675 == int32(0) {
		v1704 = v1654
		v1705 = v1655
		goto L364
	} else {
		goto L373
	}
L373:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1681
	switch v1675 - int32(1) {
	case 0:
		goto L376
	case 1:
		goto L375
	default:
		goto L374
	}
L374:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1696
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1700 = v1696 - int32(1)
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1701 < v1700 {
		v1654 = v1696
		v1655 = v1698
		v1657 = v1700
		goto L368
	} else {
		goto L381
	}
L375:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1681 < v1690 {
		v1704 = v1654
		v1705 = v1655
		goto L364
	} else {
		goto L379
	}
L376:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1681 < v1685 {
		v1704 = v1654
		v1705 = v1655
		goto L364
	} else {
		goto L377
	}
L377:
	;
	v1687 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1687 {
		goto L374
	} else {
		goto L378
	}
L378:
	;
	v1842 = v1687
	goto L363
L379:
	;
	v1692 = F_slice_del(m, l0)
	mBase = m.M
	if v1692 < int32(0) {
		v1842 = v1692
		goto L363
	} else {
		goto L380
	}
L380:
	;
	goto L374
L381:
	;
	goto L369
L382:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1793 = v1791 + (v1787 - v1790)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1793
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1793
	v1797 = v1793 - int32(1)
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1797 <= v1798 {
		goto L409
	} else {
		goto L410
	}
L383:
	;
	v1716 = v1710
	v1719 = v1708
	goto L384
L384:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1722 = int32(1)
	v1724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1720+v1716-v1722))))
	if base.B2i32(v1724&int32(224) != int32(96))|base.B2i32(v1722<<(uint(v1724)%32)&int32(71162402) == int32(0)) != 0 {
		v1787 = v1716
		v1790 = v1719
		goto L382
	} else {
		goto L386
	}
L385:
	;
	v1787 = v1781
	v1790 = v1783
	goto L382
L386:
	;
	v1739 = F_find_among_b(m, l0, int32(_a_F_basque_UTF_8_stem_3), int32(295), int32(0))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L371
	} else {
		goto L387
	}
L387:
	;
	if v1739 == int32(0) {
		v1787 = v1716
		v1790 = v1719
		goto L382
	} else {
		goto L388
	}
L388:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1743
	switch v1739 - int32(1) {
	case 0:
		goto L395
	case 1:
		goto L394
	case 2:
		goto L393
	case 3:
		goto L392
	case 4:
		goto L391
	case 5:
		goto L390
	default:
		goto L389
	}
L389:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1781
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1784 < v1781 {
		v1716 = v1781
		v1719 = v1783
		goto L384
	} else {
		goto L408
	}
L390:
	;
	v1776 = F_slice_from_s(m, l0, int32(6), int32(_a_F_basque_UTF_8_stem_4))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L371
	} else {
		goto L406
	}
L391:
	;
	v1770 = F_slice_from_s(m, l0, int32(3), int32(_a_F_basque_UTF_8_stem_5))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L371
	} else {
		goto L404
	}
L392:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1743 < v1763 {
		v1787 = v1716
		v1790 = v1719
		goto L382
	} else {
		goto L402
	}
L393:
	;
	v1759 = F_slice_from_s(m, l0, int32(3), int32(_a_F_basque_UTF_8_stem_6))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L371
	} else {
		goto L400
	}
L394:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1743 < v1752 {
		v1787 = v1716
		v1790 = v1719
		goto L382
	} else {
		goto L398
	}
L395:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1743 < v1747 {
		v1787 = v1716
		v1790 = v1719
		goto L382
	} else {
		goto L396
	}
L396:
	;
	v1749 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1749 {
		goto L389
	} else {
		goto L397
	}
L397:
	;
	v1842 = v1749
	goto L363
L398:
	;
	v1754 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1754 {
		goto L389
	} else {
		goto L399
	}
L399:
	;
	v1842 = v1754
	goto L363
L400:
	;
	if int32(0) <= v1759 {
		goto L389
	} else {
		goto L401
	}
L401:
	;
	v1842 = v1759
	goto L363
L402:
	;
	v1765 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1765 {
		goto L389
	} else {
		goto L403
	}
L403:
	;
	v1842 = v1765
	goto L363
L404:
	;
	if int32(0) <= v1770 {
		goto L389
	} else {
		goto L405
	}
L405:
	;
	v1842 = v1770
	goto L363
L406:
	;
	if v1776 < int32(0) {
		v1842 = v1776
		goto L363
	} else {
		goto L407
	}
L407:
	;
	goto L389
L408:
	;
	goto L385
L409:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1838
	v1842 = int32(1)
	goto L363
L410:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1800+v1797))))
	if base.B2i32(v1802&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1802)%32)&int32(_a_F_basque_UTF_8_stem_7) == int32(0)) != 0 {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v1817 = F_find_among_b(m, l0, int32(_a_F_basque_UTF_8_stem_8), int32(19), int32(0))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L371
	} else {
		goto L412
	}
L412:
	;
	if v1817 == int32(0) {
		goto L409
	} else {
		goto L413
	}
L413:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1821
	switch v1817 - int32(1) {
	case 0:
		goto L415
	case 1:
		goto L414
	default:
		goto L409
	}
L414:
	;
	v1832 = F_slice_from_s(m, l0, int32(1), int32(_a_F_basque_UTF_8_stem_9))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L371
	} else {
		goto L418
	}
L415:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1821 < v1825 {
		goto L409
	} else {
		goto L416
	}
L416:
	;
	v1827 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1827 {
		goto L409
	} else {
		goto L417
	}
L417:
	;
	v1842 = v1827
	goto L363
L418:
	;
	if v1832 < int32(0) {
		v1842 = v1832
		goto L363
	} else {
		goto L419
	}
L419:
	;
	goto L409
}
func F_bernoulli_initsamplescan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_palloc0(m, int32(16))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v4
		return
	}
}
func F_bernoulli_nextsampletuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
	v15 = v11
	goto L1
L1:
	;
	v21 = v15 + int32(1)
	v23 = v21 & int32(_a_F_bernoulli_nextsampletuple_0)
	if base.Ui32(l2) < base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v297)
	m.G0 = v8 + int32(16)
	return v297 & int32(_a_F_bernoulli_nextsampletuple_0)
L3:
	;
	goto L2
L4:
	;
	v297 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v23
	v28 = v8 + int32(4)
	v29 = int32(12)
	v35 = int32(-1636608420)
	if v28&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(v294) <= base.Ui64(base.I64_extend_i32_u(v289^v281-base.I32_rotl(v289, int32(24)))) {
		v15 = v21
		goto L1
	} else {
		goto L47
	}
L8:
	;
	v267 = int32(14)
	v269 = v263 ^ v264 - base.I32_rotl(v263, v267)
	v273 = v269 ^ v262 - base.I32_rotl(v269, int32(11))
	v277 = v273 ^ v263 - base.I32_rotl(v273, int32(25))
	v281 = v277 ^ v269 - base.I32_rotl(v277, int32(16))
	v285 = v281 ^ v273 - base.I32_rotl(v281, int32(4))
	v289 = v285 ^ v277 - base.I32_rotl(v285, v267)
	goto L7
L9:
	;
	switch v189 - int32(1) {
	case 0:
		v255 = v180
		v256 = v181
		v257 = v185
		goto L36
	case 1:
		v248 = v180
		v249 = v181
		v250 = v185
		goto L37
	case 2:
		v241 = v180
		v242 = v181
		v243 = v185
		goto L38
	case 3:
		v235 = v181
		v236 = v185
		goto L39
	case 4:
		v231 = v181
		v232 = v185
		goto L40
	case 5:
		v225 = v181
		v226 = v185
		goto L41
	case 6:
		v219 = v181
		v220 = v185
		goto L42
	case 7:
		v214 = v185
		goto L43
	case 8:
		v209 = v185
		goto L44
	case 9:
		v204 = v185
		goto L45
	case 10:
		goto L46
	default:
		v262 = v180
		v263 = v181
		v264 = v185
		goto L8
	}
L10:
	;
	v144 = v28
	v145 = v29
	v146 = v35
	v147 = v35
	v148 = v35
	goto L33
L11:
	;
	goto L10
L12:
	;
	goto L13
L13:
	;
	goto L17
L15:
	;
	switch v87 - int32(1) {
	case 0:
		v141 = v78
		goto L22
	case 1:
		v136 = v78
		goto L23
	case 2:
		goto L24
	case 3:
		v129 = v79
		goto L25
	case 4:
		v126 = v79
		goto L26
	case 5:
		v121 = v79
		goto L27
	case 6:
		goto L28
	case 7:
		v112 = v83
		goto L29
	case 8:
		v107 = v83
		goto L30
	case 9:
		v102 = v83
		goto L31
	case 10:
		goto L32
	default:
		v262 = v78
		v263 = v79
		v264 = v83
		goto L8
	}
L17:
	;
	goto L18
L18:
	;
	v42 = v28
	v43 = v29
	v44 = v35
	v45 = v35
	v46 = v35
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v49 = v48 + v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v53 = v52 + v46
	v55 = int32(4)
	v57 = v50 + v44 - v53 ^ base.I32_rotl(v53, v55)
	v61 = v49 - v57 ^ base.I32_rotl(v57, int32(6))
	v62 = v53 + v49
	v63 = v57 + v62
	v64 = v61 + v63
	v68 = v62 - v61 ^ base.I32_rotl(v61, int32(8))
	v72 = v63 - v68 ^ base.I32_rotl(v68, int32(16))
	v76 = v64 - v72 ^ base.I32_rotl(v72, int32(19))
	v77 = v68 + v64
	v78 = v72 + v77
	v79 = v76 + v78
	v83 = v77 - v76 ^ base.I32_rotl(v76, v55)
	v84 = int32(12)
	v85 = v42 + v84
	v87 = v43 - v84
	if base.Ui32(int32(11)) < base.Ui32(v87) {
		v42 = v85
		v43 = v87
		v44 = v78
		v45 = v79
		v46 = v83
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	goto L20
L22:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v262 = v141 + v142
	v263 = v79
	v264 = v83
	goto L8
L23:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v141 = v137<<(uint(int32(8))%32) + v136
	goto L22
L24:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+2)))
	v136 = v132<<(uint(int32(16))%32) + v78
	goto L23
L25:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v262 = v130 + v78
	v263 = v129
	v264 = v83
	goto L8
L26:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	v129 = v126 + v127
	goto L25
L27:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+5)))
	v126 = v122<<(uint(int32(8))%32) + v121
	goto L26
L28:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+6)))
	v121 = v117<<(uint(int32(16))%32) + v79
	goto L27
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v262 = v113 + v78
	v263 = v115 + v79
	v264 = v112
	goto L8
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
	v112 = v108<<(uint(int32(8))%32) + v107
	goto L29
L31:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+9)))
	v107 = v103<<(uint(int32(16))%32) + v102
	goto L30
L32:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+10)))
	v102 = v98<<(uint(int32(24))%32) + v83
	goto L31
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v151 = v150 + v147
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v155 = v154 + v148
	v157 = int32(4)
	v159 = v152 + v146 - v155 ^ base.I32_rotl(v155, v157)
	v163 = v151 - v159 ^ base.I32_rotl(v159, int32(6))
	v164 = v155 + v151
	v165 = v159 + v164
	v166 = v163 + v165
	v170 = v164 - v163 ^ base.I32_rotl(v163, int32(8))
	v174 = v165 - v170 ^ base.I32_rotl(v170, int32(16))
	v178 = v166 - v174 ^ base.I32_rotl(v174, int32(19))
	v179 = v170 + v166
	v180 = v174 + v179
	v181 = v178 + v180
	v185 = v179 - v178 ^ base.I32_rotl(v178, v157)
	v186 = int32(12)
	v187 = v144 + v186
	v189 = v145 - v186
	if base.Ui32(int32(11)) < base.Ui32(v189) {
		v144 = v187
		v145 = v189
		v146 = v180
		v147 = v181
		v148 = v185
		goto L33
	} else {
		goto L35
	}
L34:
	;
	goto L9
L35:
	;
	goto L34
L36:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v262 = v255 + v258
	v263 = v256
	v264 = v257
	goto L8
L37:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v255 = v251<<(uint(int32(8))%32) + v248
	v256 = v249
	v257 = v250
	goto L36
L38:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
	v248 = v244<<(uint(int32(16))%32) + v241
	v249 = v242
	v250 = v243
	goto L37
L39:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
	v241 = v237<<(uint(int32(24))%32) + v180
	v242 = v235
	v243 = v236
	goto L38
L40:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
	v235 = v231 + v233
	v236 = v232
	goto L39
L41:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
	v231 = v227<<(uint(int32(8))%32) + v225
	v232 = v226
	goto L40
L42:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
	v225 = v221<<(uint(int32(16))%32) + v219
	v226 = v220
	goto L41
L43:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+7)))
	v219 = v215<<(uint(int32(24))%32) + v181
	v220 = v214
	goto L42
L44:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
	v214 = v210<<(uint(int32(8))%32) + v209
	goto L43
L45:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
	v209 = v205<<(uint(int32(16))%32) + v204
	goto L44
L46:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
	v204 = v200<<(uint(int32(24))%32) + v185
	goto L45
L47:
	;
	v297 = v21
	goto L3
}
func F_bitgetbit(m *base.Module, l0 int32) int64 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = int32(0)
		if base.B2i32(v16 <= v15)&base.B2i32(v15 < v14) == v16 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15
					F_errmsg(m, int32(_a_F_bitgetbit_0), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_bitgetbit_1), int32(1883), int32(_a_F_bitgetbit_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(base.Ui32(v15)>>(uint(int32(3))%32)))+8)))
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(int32(base.Ui32(v44)>>(uint((v15^int32(-1))&int32(7))%32)) & int32(1))
		}
	}
}
func F_bitne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 == v17 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = int32(8)
	v20 = v8 + v19
	v22 = v13 + v19
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v24 = int32(2)
	v25 = int32(base.Ui32(v23) >> (uint(v24) % 32))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v28 = int32(base.Ui32(v26) >> (uint(v24) % 32))
	if base.Ui32(v25) < base.Ui32(v28) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v100 = int64(1)
	goto L6
L6:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v101 != v8 {
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v30 = v25
	goto L9
L8:
	;
	v30 = v28
	goto L9
L9:
	;
	v32 = v30 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v100 = base.I64_extend_i32_u(base.B2i32(v94 != int32(0)))
	goto L6
L11:
	;
	v94 = int32(0)
	goto L10
L12:
	;
	v68 = v63
	v69 = v64
	v70 = v65
	goto L22
L13:
	;
	if (v20|v22)&int32(3) != 0 {
		v63 = v20
		v64 = v22
		v65 = v32
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v56 = v20
	v57 = v22
	v58 = v32
	goto L15
L15:
	;
	if v58 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v40 = v20
	v41 = v22
	v42 = v32
	goto L17
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v45 != v46 {
		v63 = v40
		v64 = v41
		v65 = v42
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v56 = v51
	v57 = v49
	v58 = v53
	goto L15
L19:
	;
	v48 = int32(4)
	v49 = v41 + v48
	v51 = v40 + v48
	v53 = v42 - v48
	if base.Ui32(int32(3)) < base.Ui32(v53) {
		v40 = v51
		v41 = v49
		v42 = v53
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v63 = v56
	v64 = v57
	v65 = v58
	goto L12
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 == v74 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v94 = v73 - v74
	goto L10
L24:
	;
	v76 = int32(1)
	v81 = v70 - v76
	if v81 != 0 {
		v68 = v68 + v76
		v69 = v69 + v76
		v70 = v81
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L11
L28:
	;
	F_pfree(m, v8)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v105 != v13 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	F_pfree(m, v13)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v100
L35:
	;
	goto L34
}
func F_bitposition(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v105 int32
	_ = v105
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v32 < v35 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v35 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int64(1)
L8:
	;
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v42 = int32(2)
	v43 = int32(base.Ui32(v41) >> (uint(v42) % 32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v46 = int32(base.Ui32(v44) >> (uint(v42) % 32))
	v47 = v43 - v46
	if v47 == int32(-1) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v50 = int32(255)
	v51 = int32(3)
	v54 = int32(-64)
	v56 = v50 << (uint(v46<<(uint(v51)%32)-v35+v54) % 32)
	v59 = int32(8)
	v63 = v25 + v43
	v64 = int32(1)
	v65 = v63 - v64
	v66 = v30 + v46
	v75 = v50 << (uint(v43<<(uint(v51)%32)-v32+v54) % 32)
	v77 = v75 ^ int32(-1)
	v85 = int32(0)
	goto L11
L11:
	;
	v105 = int32(0)
	goto L13
L12:
	;
	goto L4
L13:
	;
	v132 = int32(8) - v105
	v133 = v56 << (uint(v132) % 32)
	v134 = v30 + v59
	v135 = v85 + (v25 + v59)
	v137 = int32(-256) >> (uint(v105) % 32)
	v138 = int32(base.Ui32(int32(255)) >> (uint(v105) % 32))
	goto L15
L14:
	;
	if v85 != v47 {
		v85 = v85 + int32(1)
		goto L11
	} else {
		goto L35
	}
L15:
	;
	if base.Ui32(v134) < base.Ui32(v66) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v207 = v105 + int32(1)
	if v207 != int32(8) {
		v105 = v207
		goto L13
	} else {
		goto L34
	}
L17:
	;
	goto L16
L18:
	;
	if v134 != v66-v64 {
		v191 = v137
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v159 = base.B2i32(v134 != v66-v64)
	if v134 != v66-v64 {
		v164 = v138
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	return base.I64_extend_i32_s(v105 + v85<<(uint(int32(3))%32) + int32(1))
L22:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v164&(v165^int32(base.Ui32(v158)>>(uint(v105)%32))) != 0 {
		goto L17
	} else {
		goto L26
	}
L23:
	;
	v160 = v138 & int32(base.Ui32(v56&v50)>>(uint(v105)%32))
	if v135 != v65 {
		v164 = v160
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v160&v77 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v164 = v160 & v75
	goto L22
L26:
	;
	v170 = v135 + int32(1)
	if v170 != v63 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	if v133&int32(254) != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v191&(v194^v158<<(uint(v132)%32))&int32(255) == int32(0) {
		v134 = v134 + int32(1)
		v135 = v170
		v137 = v191
		v138 = v164
		goto L15
	} else {
		goto L33
	}
L30:
	;
	v185 = v137 & v133
	if v170 != v65 {
		v191 = v185
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v185&v77&int32(255) != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v191 = v185 & v75
	goto L29
L33:
	;
	goto L17
L34:
	;
	goto L14
L35:
	;
	goto L12
}
func F_bitsetbit(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = int32(0)
		if base.B2i32(v18 < v19)|base.B2i32(v17 <= v18) == v19 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if base.Ui32(int32(2)) <= base.Ui32(v25) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bitsetbit_0), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bitsetbit_1), int32(1833), int32(_a_F_bitsetbit_2))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v31 = F_palloc(m, int32(base.Ui32(v28)>>(uint(int32(2))%32)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v28 & int32(-4)
					v37 = int32(8)
					v38 = v31 + v37
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v43 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - v37
					if v43 != 0 {
						base.MemoryCopy(m, v38, v13+int32(8), v43)
					} else {
					}
					v49 = v38 + int32(base.Ui32(v18)>>(uint(int32(3))%32))
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					v55 = (v18 ^ int32(-1)) & int32(7)
					if v25 != 0 {
						v61 = v50 | int32(1)<<(uint(v55)%32)
					} else {
						v61 = v50 & base.I32_rotl(int32(-2), v55)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v61)
					m.G0 = v10 + int32(16)
					return base.I64_extend_i32_u(v31)
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
					F_errmsg(m, int32(_a_F_bitsetbit_3), v10)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_bitsetbit_1), int32(1825), int32(_a_F_bitsetbit_2))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
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
func F_bitsubstr_no_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_bitsubstring(m, v3, v7, int32(-1), int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v10)
		}
	}
}
func F_bitxor(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v15 == v16 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v21 = F_palloc(m, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v15
					v24 = int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18 & v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v27&v24 != int32(32) {
						v32 = int32(8)
						v38 = v8 + v32
						v39 = v13 + v32
						v41 = v21 + v32
						v43 = int32(0)
						for {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
							v46 = v44 ^ v45
							*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
							v48 = int32(1)
							v55 = v43 + v48
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							if base.Ui32(v55) < base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(8)) {
								v38 = v38 + v48
								v39 = v39 + v48
								v41 = v41 + v48
								v43 = v55
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return base.I64_extend_i32_u(v21)
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bitxor_0), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bitxor_1), int32(1342), int32(_a_F_bitxor_2))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
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
func F_blackhole_get_sink(m *base.Module, l0 int32, l1 int32) int32 {
	return l0
}
func F_blbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(1172))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			F_initBloomState(m, v9+int32(4), v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
				return v4
			}
		}
	}
}
func F_blendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 != 0 {
		F_pfree(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(0)
		return
	}
}
func F_blockreftable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
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
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v309 int64
	_ = v309
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v330 int64
	_ = v330
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v357 int64
	_ = v357
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int64
	_ = v982
	var v984 int64
	_ = v984
	var v986 int64
	_ = v986
	var v988 int64
	_ = v988
	var v990 int64
	_ = v990
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1059 int64
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1062 int64
	_ = v1062
	var v1063 int64
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1370 int64
	_ = v1370
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int64
	_ = v1420
	var v1422 int64
	_ = v1422
	var v1424 int64
	_ = v1424
	var v1426 int64
	_ = v1426
	var v1428 int64
	_ = v1428
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int64
	_ = v1452
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1469 int32
	_ = v1469
	var v1484 int32
	_ = v1484
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int64
	_ = v1495
	var v1497 int64
	_ = v1497
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1550 int32
	_ = v1550
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	v16 = m.G0
	v17 = int32(16)
	v18 = v16 - v17
	m.G0 = v18
	v26 = int32(-1636608416)
	if l1&int32(3) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v285 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v293 = base.B2i32(base.Ui32(v289) < base.Ui32(v290))
	goto L41
L2:
	;
	v258 = int32(14)
	v260 = v254 ^ v255 - base.I32_rotl(v254, v258)
	v264 = v260 ^ v253 - base.I32_rotl(v260, int32(11))
	v268 = v264 ^ v254 - base.I32_rotl(v264, int32(25))
	v272 = v268 ^ v260 - base.I32_rotl(v268, int32(16))
	v276 = v272 ^ v264 - base.I32_rotl(v272, int32(4))
	v280 = v276 ^ v268 - base.I32_rotl(v276, v258)
	goto L1
L3:
	;
	switch v180 - int32(1) {
	case 0:
		v246 = v171
		v247 = v172
		v248 = v176
		goto L30
	case 1:
		v239 = v171
		v240 = v172
		v241 = v176
		goto L31
	case 2:
		v232 = v171
		v233 = v172
		v234 = v176
		goto L32
	case 3:
		v226 = v172
		v227 = v176
		goto L33
	case 4:
		v222 = v172
		v223 = v176
		goto L34
	case 5:
		v216 = v172
		v217 = v176
		goto L35
	case 6:
		v210 = v172
		v211 = v176
		goto L36
	case 7:
		v205 = v176
		goto L37
	case 8:
		v200 = v176
		goto L38
	case 9:
		v195 = v176
		goto L39
	case 10:
		goto L40
	default:
		v253 = v171
		v254 = v172
		v255 = v176
		goto L2
	}
L4:
	;
	v135 = l1
	v136 = v17
	v137 = v26
	v138 = v26
	v139 = v26
	goto L27
L5:
	;
	goto L4
L6:
	;
	goto L7
L7:
	;
	goto L11
L9:
	;
	switch v78 - int32(1) {
	case 0:
		v132 = v69
		goto L16
	case 1:
		v127 = v69
		goto L17
	case 2:
		goto L18
	case 3:
		v120 = v70
		goto L19
	case 4:
		v117 = v70
		goto L20
	case 5:
		v112 = v70
		goto L21
	case 6:
		goto L22
	case 7:
		v103 = v74
		goto L23
	case 8:
		v98 = v74
		goto L24
	case 9:
		v93 = v74
		goto L25
	case 10:
		goto L26
	default:
		v253 = v69
		v254 = v70
		v255 = v74
		goto L2
	}
L11:
	;
	goto L12
L12:
	;
	v33 = l1
	v34 = v17
	v35 = v26
	v36 = v26
	v37 = v26
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v40 = v39 + v36
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v44 = v43 + v37
	v46 = int32(4)
	v48 = v41 + v35 - v44 ^ base.I32_rotl(v44, v46)
	v52 = v40 - v48 ^ base.I32_rotl(v48, int32(6))
	v53 = v44 + v40
	v54 = v48 + v53
	v55 = v52 + v54
	v59 = v53 - v52 ^ base.I32_rotl(v52, int32(8))
	v63 = v54 - v59 ^ base.I32_rotl(v59, int32(16))
	v67 = v55 - v63 ^ base.I32_rotl(v63, int32(19))
	v68 = v59 + v55
	v69 = v63 + v68
	v70 = v67 + v69
	v74 = v68 - v67 ^ base.I32_rotl(v67, v46)
	v75 = int32(12)
	v76 = v33 + v75
	v78 = v34 - v75
	if base.Ui32(int32(11)) < base.Ui32(v78) {
		v33 = v76
		v34 = v78
		v35 = v69
		v36 = v70
		v37 = v74
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L9
L15:
	;
	goto L14
L16:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v253 = v132 + v133
	v254 = v70
	v255 = v74
	goto L2
L17:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v132 = v128<<(uint(int32(8))%32) + v127
	goto L16
L18:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
	v127 = v123<<(uint(int32(16))%32) + v69
	goto L17
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v253 = v121 + v69
	v254 = v120
	v255 = v74
	goto L2
L20:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
	v120 = v117 + v118
	goto L19
L21:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
	v117 = v113<<(uint(int32(8))%32) + v112
	goto L20
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
	v112 = v108<<(uint(int32(16))%32) + v70
	goto L21
L23:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v253 = v104 + v69
	v254 = v106 + v70
	v255 = v103
	goto L2
L24:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
	v103 = v99<<(uint(int32(8))%32) + v98
	goto L23
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
	v98 = v94<<(uint(int32(16))%32) + v93
	goto L24
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)))
	v93 = v89<<(uint(int32(24))%32) + v74
	goto L25
L27:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v142 = v141 + v138
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v146 = v145 + v139
	v148 = int32(4)
	v150 = v143 + v137 - v146 ^ base.I32_rotl(v146, v148)
	v154 = v142 - v150 ^ base.I32_rotl(v150, int32(6))
	v155 = v146 + v142
	v156 = v150 + v155
	v157 = v154 + v156
	v161 = v155 - v154 ^ base.I32_rotl(v154, int32(8))
	v165 = v156 - v161 ^ base.I32_rotl(v161, int32(16))
	v169 = v157 - v165 ^ base.I32_rotl(v165, int32(19))
	v170 = v161 + v157
	v171 = v165 + v170
	v172 = v169 + v171
	v176 = v170 - v169 ^ base.I32_rotl(v169, v148)
	v177 = int32(12)
	v178 = v135 + v177
	v180 = v136 - v177
	if base.Ui32(int32(11)) < base.Ui32(v180) {
		v135 = v178
		v136 = v180
		v137 = v171
		v138 = v172
		v139 = v176
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	goto L28
L30:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v253 = v246 + v249
	v254 = v247
	v255 = v248
	goto L2
L31:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	v246 = v242<<(uint(int32(8))%32) + v239
	v247 = v240
	v248 = v241
	goto L30
L32:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
	v239 = v235<<(uint(int32(16))%32) + v232
	v240 = v233
	v241 = v234
	goto L31
L33:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
	v232 = v228<<(uint(int32(24))%32) + v171
	v233 = v226
	v234 = v227
	goto L32
L34:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
	v226 = v222 + v224
	v227 = v223
	goto L33
L35:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
	v222 = v218<<(uint(int32(8))%32) + v216
	v223 = v217
	goto L34
L36:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
	v216 = v212<<(uint(int32(16))%32) + v210
	v217 = v211
	goto L35
L37:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
	v210 = v206<<(uint(int32(24))%32) + v172
	v211 = v205
	goto L36
L38:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
	v205 = v201<<(uint(int32(8))%32) + v200
	goto L37
L39:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
	v200 = v196<<(uint(int32(16))%32) + v195
	goto L38
L40:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
	v195 = v191<<(uint(int32(24))%32) + v176
	goto L39
L41:
	;
	if v293 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L60
	} else {
		goto L249
	}
L43:
	;
	goto L42
L44:
	;
	v1550 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1550
	v293 = v1550
	goto L41
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L60
	} else {
		goto L246
	}
L46:
	;
	v309 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v309 == int64(4294967296) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1037 = v1036 & (v280 ^ v272 - base.I32_rotl(v280, int32(24)))
	v1040 = v1035 + v1037*int32(40)
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040)+20)))
	if v1041 != 0 {
		goto L172
	} else {
		goto L173
	}
L49:
	;
	v312 = int32(0)
	v314 = int64(2)
	v316 = v309 << (uint(int64(1)) % 64)
	if base.Ui64(v316) <= base.Ui64(v314) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v293 = int32(1)
	goto L41
L51:
	;
	v319 = v314
	goto L53
L52:
	;
	v319 = v316
	goto L53
L53:
	;
	v320 = int64(1)
	if v319&(v319-v320) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v330 = v319
	goto L56
L55:
	;
	v330 = v320 << (uint(int64(64)-base.I64_clz(v319)) % 64)
	goto L56
L56:
	;
	if base.Ui64(v330*int64(40)) < base.Ui64(int64(2147483647)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v342 = F_MemoryContextAllocExtended(m, v337, base.I32_wrap_i64(v330)*int32(40), int32(5))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	goto L43
L60:
	;
	return int32(0)
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v342
	v347 = int64(1)
	if v330&(v330-v347) == int64(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v357 = v330
	goto L64
L63:
	;
	v357 = v347 << (uint(int64(64)-base.I64_clz(v330)) % 64)
	goto L64
L64:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v357*int64(40)) {
		goto L43
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v357
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = base.I32_wrap_i64(v357) - int32(1)
	if v357 == int64(4294967296) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v374 = int32(-85899346)
	goto L68
L67:
	;
	v374 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v357), float64(0.9)))
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v374
	if v336 != int64(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v384 = v312
	goto L73
L70:
	;
	goto L71
L71:
	;
	F_pfree(m, v335)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L60
	} else {
		goto L170
	}
L72:
	;
	v678 = v312
	v679 = v672
	goto L118
L73:
	;
	v395 = v335 + v384*int32(40)
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)))
	if v396 != int32(1) {
		v672 = v384
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v672 = int32(0)
	goto L72
L75:
	;
	v399 = int32(16)
	v405 = int32(-1636608416)
	if v395&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (v659^v651-base.I32_rotl(v659, int32(24)))&v664 == v384 {
		v672 = v384
		goto L72
	} else {
		goto L116
	}
L77:
	;
	v637 = int32(14)
	v639 = v633 ^ v634 - base.I32_rotl(v633, v637)
	v643 = v639 ^ v632 - base.I32_rotl(v639, int32(11))
	v647 = v643 ^ v633 - base.I32_rotl(v643, int32(25))
	v651 = v647 ^ v639 - base.I32_rotl(v647, int32(16))
	v655 = v651 ^ v643 - base.I32_rotl(v651, int32(4))
	v659 = v655 ^ v647 - base.I32_rotl(v655, v637)
	goto L76
L78:
	;
	switch v559 - int32(1) {
	case 0:
		v625 = v550
		v626 = v551
		v627 = v555
		goto L105
	case 1:
		v618 = v550
		v619 = v551
		v620 = v555
		goto L106
	case 2:
		v611 = v550
		v612 = v551
		v613 = v555
		goto L107
	case 3:
		v605 = v551
		v606 = v555
		goto L108
	case 4:
		v601 = v551
		v602 = v555
		goto L109
	case 5:
		v595 = v551
		v596 = v555
		goto L110
	case 6:
		v589 = v551
		v590 = v555
		goto L111
	case 7:
		v584 = v555
		goto L112
	case 8:
		v579 = v555
		goto L113
	case 9:
		v574 = v555
		goto L114
	case 10:
		goto L115
	default:
		v632 = v550
		v633 = v551
		v634 = v555
		goto L77
	}
L79:
	;
	v514 = v395
	v515 = v399
	v516 = v405
	v517 = v405
	v518 = v405
	goto L102
L80:
	;
	goto L79
L81:
	;
	goto L82
L82:
	;
	goto L86
L84:
	;
	switch v457 - int32(1) {
	case 0:
		v511 = v448
		goto L91
	case 1:
		v506 = v448
		goto L92
	case 2:
		goto L93
	case 3:
		v499 = v449
		goto L94
	case 4:
		v496 = v449
		goto L95
	case 5:
		v491 = v449
		goto L96
	case 6:
		goto L97
	case 7:
		v482 = v453
		goto L98
	case 8:
		v477 = v453
		goto L99
	case 9:
		v472 = v453
		goto L100
	case 10:
		goto L101
	default:
		v632 = v448
		v633 = v449
		v634 = v453
		goto L77
	}
L86:
	;
	goto L87
L87:
	;
	v412 = v395
	v413 = v399
	v414 = v405
	v415 = v405
	v416 = v405
	goto L88
L88:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v419 = v418 + v415
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v412)+8))
	v423 = v422 + v416
	v425 = int32(4)
	v427 = v420 + v414 - v423 ^ base.I32_rotl(v423, v425)
	v431 = v419 - v427 ^ base.I32_rotl(v427, int32(6))
	v432 = v423 + v419
	v433 = v427 + v432
	v434 = v431 + v433
	v438 = v432 - v431 ^ base.I32_rotl(v431, int32(8))
	v442 = v433 - v438 ^ base.I32_rotl(v438, int32(16))
	v446 = v434 - v442 ^ base.I32_rotl(v442, int32(19))
	v447 = v438 + v434
	v448 = v442 + v447
	v449 = v446 + v448
	v453 = v447 - v446 ^ base.I32_rotl(v446, v425)
	v454 = int32(12)
	v455 = v412 + v454
	v457 = v413 - v454
	if base.Ui32(int32(11)) < base.Ui32(v457) {
		v412 = v455
		v413 = v457
		v414 = v448
		v415 = v449
		v416 = v453
		goto L88
	} else {
		goto L90
	}
L89:
	;
	goto L84
L90:
	;
	goto L89
L91:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	v632 = v511 + v512
	v633 = v449
	v634 = v453
	goto L77
L92:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+1)))
	v511 = v507<<(uint(int32(8))%32) + v506
	goto L91
L93:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+2)))
	v506 = v502<<(uint(int32(16))%32) + v448
	goto L92
L94:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v632 = v500 + v448
	v633 = v499
	v634 = v453
	goto L77
L95:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+4)))
	v499 = v496 + v497
	goto L94
L96:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+5)))
	v496 = v492<<(uint(int32(8))%32) + v491
	goto L95
L97:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+6)))
	v491 = v487<<(uint(int32(16))%32) + v449
	goto L96
L98:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v632 = v483 + v448
	v633 = v485 + v449
	v634 = v482
	goto L77
L99:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+8)))
	v482 = v478<<(uint(int32(8))%32) + v477
	goto L98
L100:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+9)))
	v477 = v473<<(uint(int32(16))%32) + v472
	goto L99
L101:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+10)))
	v472 = v468<<(uint(int32(24))%32) + v453
	goto L100
L102:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v521 = v520 + v517
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v525 = v524 + v518
	v527 = int32(4)
	v529 = v522 + v516 - v525 ^ base.I32_rotl(v525, v527)
	v533 = v521 - v529 ^ base.I32_rotl(v529, int32(6))
	v534 = v525 + v521
	v535 = v529 + v534
	v536 = v533 + v535
	v540 = v534 - v533 ^ base.I32_rotl(v533, int32(8))
	v544 = v535 - v540 ^ base.I32_rotl(v540, int32(16))
	v548 = v536 - v544 ^ base.I32_rotl(v544, int32(19))
	v549 = v540 + v536
	v550 = v544 + v549
	v551 = v548 + v550
	v555 = v549 - v548 ^ base.I32_rotl(v548, v527)
	v556 = int32(12)
	v557 = v514 + v556
	v559 = v515 - v556
	if base.Ui32(int32(11)) < base.Ui32(v559) {
		v514 = v557
		v515 = v559
		v516 = v550
		v517 = v551
		v518 = v555
		goto L102
	} else {
		goto L104
	}
L103:
	;
	goto L78
L104:
	;
	goto L103
L105:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v632 = v625 + v628
	v633 = v626
	v634 = v627
	goto L77
L106:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	v625 = v621<<(uint(int32(8))%32) + v618
	v626 = v619
	v627 = v620
	goto L105
L107:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+2)))
	v618 = v614<<(uint(int32(16))%32) + v611
	v619 = v612
	v620 = v613
	goto L106
L108:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+3)))
	v611 = v607<<(uint(int32(24))%32) + v550
	v612 = v605
	v613 = v606
	goto L107
L109:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+4)))
	v605 = v601 + v603
	v606 = v602
	goto L108
L110:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+5)))
	v601 = v597<<(uint(int32(8))%32) + v595
	v602 = v596
	goto L109
L111:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+6)))
	v595 = v591<<(uint(int32(16))%32) + v589
	v596 = v590
	goto L110
L112:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+7)))
	v589 = v585<<(uint(int32(24))%32) + v551
	v590 = v584
	goto L111
L113:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+8)))
	v584 = v580<<(uint(int32(8))%32) + v579
	goto L112
L114:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+9)))
	v579 = v575<<(uint(int32(16))%32) + v574
	goto L113
L115:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+10)))
	v574 = v570<<(uint(int32(24))%32) + v555
	goto L114
L116:
	;
	v668 = v384 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v668)) < base.Ui64(v336) {
		v384 = v668
		goto L73
	} else {
		goto L117
	}
L117:
	;
	goto L74
L118:
	;
	v690 = v335 + v679*int32(40)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690)+20)))
	if v691 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L71
L120:
	;
	v694 = int32(16)
	v700 = int32(-1636608416)
	if v690&int32(3) != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L122
L122:
	;
	v1008 = v679 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1008)) < base.Ui64(v336) {
		goto L166
	} else {
		goto L167
	}
L123:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v973 = v954 ^ v946 - base.I32_rotl(v954, int32(24))
	goto L163
L124:
	;
	v932 = int32(14)
	v934 = v928 ^ v929 - base.I32_rotl(v928, v932)
	v938 = v934 ^ v927 - base.I32_rotl(v934, int32(11))
	v942 = v938 ^ v928 - base.I32_rotl(v938, int32(25))
	v946 = v942 ^ v934 - base.I32_rotl(v942, int32(16))
	v950 = v946 ^ v938 - base.I32_rotl(v946, int32(4))
	v954 = v950 ^ v942 - base.I32_rotl(v950, v932)
	goto L123
L125:
	;
	switch v854 - int32(1) {
	case 0:
		v920 = v845
		v921 = v846
		v922 = v850
		goto L152
	case 1:
		v913 = v845
		v914 = v846
		v915 = v850
		goto L153
	case 2:
		v906 = v845
		v907 = v846
		v908 = v850
		goto L154
	case 3:
		v900 = v846
		v901 = v850
		goto L155
	case 4:
		v896 = v846
		v897 = v850
		goto L156
	case 5:
		v890 = v846
		v891 = v850
		goto L157
	case 6:
		v884 = v846
		v885 = v850
		goto L158
	case 7:
		v879 = v850
		goto L159
	case 8:
		v874 = v850
		goto L160
	case 9:
		v869 = v850
		goto L161
	case 10:
		goto L162
	default:
		v927 = v845
		v928 = v846
		v929 = v850
		goto L124
	}
L126:
	;
	v809 = v690
	v810 = v694
	v811 = v700
	v812 = v700
	v813 = v700
	goto L149
L127:
	;
	goto L126
L128:
	;
	goto L129
L129:
	;
	goto L133
L131:
	;
	switch v752 - int32(1) {
	case 0:
		v806 = v743
		goto L138
	case 1:
		v801 = v743
		goto L139
	case 2:
		goto L140
	case 3:
		v794 = v744
		goto L141
	case 4:
		v791 = v744
		goto L142
	case 5:
		v786 = v744
		goto L143
	case 6:
		goto L144
	case 7:
		v777 = v748
		goto L145
	case 8:
		v772 = v748
		goto L146
	case 9:
		v767 = v748
		goto L147
	case 10:
		goto L148
	default:
		v927 = v743
		v928 = v744
		v929 = v748
		goto L124
	}
L133:
	;
	goto L134
L134:
	;
	v707 = v690
	v708 = v694
	v709 = v700
	v710 = v700
	v711 = v700
	goto L135
L135:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v714 = v713 + v710
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v707)+8))
	v718 = v717 + v711
	v720 = int32(4)
	v722 = v715 + v709 - v718 ^ base.I32_rotl(v718, v720)
	v726 = v714 - v722 ^ base.I32_rotl(v722, int32(6))
	v727 = v718 + v714
	v728 = v722 + v727
	v729 = v726 + v728
	v733 = v727 - v726 ^ base.I32_rotl(v726, int32(8))
	v737 = v728 - v733 ^ base.I32_rotl(v733, int32(16))
	v741 = v729 - v737 ^ base.I32_rotl(v737, int32(19))
	v742 = v733 + v729
	v743 = v737 + v742
	v744 = v741 + v743
	v748 = v742 - v741 ^ base.I32_rotl(v741, v720)
	v749 = int32(12)
	v750 = v707 + v749
	v752 = v708 - v749
	if base.Ui32(int32(11)) < base.Ui32(v752) {
		v707 = v750
		v708 = v752
		v709 = v743
		v710 = v744
		v711 = v748
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L131
L137:
	;
	goto L136
L138:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	v927 = v806 + v807
	v928 = v744
	v929 = v748
	goto L124
L139:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	v806 = v802<<(uint(int32(8))%32) + v801
	goto L138
L140:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+2)))
	v801 = v797<<(uint(int32(16))%32) + v743
	goto L139
L141:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v927 = v795 + v743
	v928 = v794
	v929 = v748
	goto L124
L142:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+4)))
	v794 = v791 + v792
	goto L141
L143:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+5)))
	v791 = v787<<(uint(int32(8))%32) + v786
	goto L142
L144:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+6)))
	v786 = v782<<(uint(int32(16))%32) + v744
	goto L143
L145:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v927 = v778 + v743
	v928 = v780 + v744
	v929 = v777
	goto L124
L146:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+8)))
	v777 = v773<<(uint(int32(8))%32) + v772
	goto L145
L147:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+9)))
	v772 = v768<<(uint(int32(16))%32) + v767
	goto L146
L148:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+10)))
	v767 = v763<<(uint(int32(24))%32) + v748
	goto L147
L149:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	v816 = v815 + v812
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v809)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v809)+8))
	v820 = v819 + v813
	v822 = int32(4)
	v824 = v817 + v811 - v820 ^ base.I32_rotl(v820, v822)
	v828 = v816 - v824 ^ base.I32_rotl(v824, int32(6))
	v829 = v820 + v816
	v830 = v824 + v829
	v831 = v828 + v830
	v835 = v829 - v828 ^ base.I32_rotl(v828, int32(8))
	v839 = v830 - v835 ^ base.I32_rotl(v835, int32(16))
	v843 = v831 - v839 ^ base.I32_rotl(v839, int32(19))
	v844 = v835 + v831
	v845 = v839 + v844
	v846 = v843 + v845
	v850 = v844 - v843 ^ base.I32_rotl(v843, v822)
	v851 = int32(12)
	v852 = v809 + v851
	v854 = v810 - v851
	if base.Ui32(int32(11)) < base.Ui32(v854) {
		v809 = v852
		v810 = v854
		v811 = v845
		v812 = v846
		v813 = v850
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L125
L151:
	;
	goto L150
L152:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	v927 = v920 + v923
	v928 = v921
	v929 = v922
	goto L124
L153:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+1)))
	v920 = v916<<(uint(int32(8))%32) + v913
	v921 = v914
	v922 = v915
	goto L152
L154:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+2)))
	v913 = v909<<(uint(int32(16))%32) + v906
	v914 = v907
	v915 = v908
	goto L153
L155:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+3)))
	v906 = v902<<(uint(int32(24))%32) + v845
	v907 = v900
	v908 = v901
	goto L154
L156:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+4)))
	v900 = v896 + v898
	v901 = v897
	goto L155
L157:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+5)))
	v896 = v892<<(uint(int32(8))%32) + v890
	v897 = v891
	goto L156
L158:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+6)))
	v890 = v886<<(uint(int32(16))%32) + v884
	v891 = v885
	goto L157
L159:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+7)))
	v884 = v880<<(uint(int32(24))%32) + v846
	v885 = v879
	goto L158
L160:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+8)))
	v879 = v875<<(uint(int32(8))%32) + v874
	goto L159
L161:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+9)))
	v874 = v870<<(uint(int32(16))%32) + v869
	goto L160
L162:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+10)))
	v869 = v865<<(uint(int32(24))%32) + v850
	goto L161
L163:
	;
	v975 = v959 & v973
	v980 = v342 + v975*int32(40)
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+20)))
	if v981 != 0 {
		v973 = v975 + int32(1)
		goto L163
	} else {
		goto L165
	}
L164:
	;
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v690)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+32)) = v982
	v984 = *(*int64)(unsafe.Add(mBase, uint32(v690)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+24)) = v984
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v690)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+16)) = v986
	v988 = *(*int64)(unsafe.Add(mBase, uint32(v690)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+8)) = v988
	v990 = *(*int64)(unsafe.Add(mBase, uint32(v690)))
	*(*int64)(unsafe.Add(mBase, uint32(v980))) = v990
	goto L122
L165:
	;
	goto L164
L166:
	;
	v1012 = v1008
	goto L168
L167:
	;
	v1012 = int32(0)
	goto L168
L168:
	;
	v1014 = v678 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1014)) < base.Ui64(v336) {
		v678 = v1014
		v679 = v1012
		goto L118
	} else {
		goto L169
	}
L169:
	;
	goto L119
L170:
	;
	goto L50
L171:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1512)
	m.G0 = v18 + int32(16)
	return v1510
L172:
	;
	v1051 = int32(0)
	v1052 = v1040
	v1055 = v1037
	goto L176
L173:
	;
	v1484 = v1040
	goto L174
L174:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1492 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1491 + v1492
	v1495 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1484)+8)) = v1495
	v1497 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v1484))) = v1497
	*(*uint8)(unsafe.Add(mBase, uint32(v1484)+20)) = uint8(v1492)
	v1510 = v1484
	v1512 = int32(0)
	goto L171
L175:
	;
	v1484 = v1469
	goto L174
L176:
	;
	v1059 = *(*int64)(unsafe.Add(mBase, uint32(v1052)))
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1052)+8))
	v1063 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v1059^v1060|(v1062^v1063) == int64(0) {
		v1510 = v1052
		v1512 = int32(1)
		goto L171
	} else {
		goto L178
	}
L177:
	;
	v1469 = v1459
	goto L175
L178:
	;
	v1068 = int32(16)
	v1074 = int32(-1636608416)
	if v1052&int32(3) != 0 {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1334 = (v1328 ^ v1320 - base.I32_rotl(v1328, int32(24))) & v1333
	if base.Ui32(v1055) < base.Ui32(v1334) {
		goto L219
	} else {
		goto L220
	}
L180:
	;
	v1306 = int32(14)
	v1308 = v1302 ^ v1303 - base.I32_rotl(v1302, v1306)
	v1312 = v1308 ^ v1301 - base.I32_rotl(v1308, int32(11))
	v1316 = v1312 ^ v1302 - base.I32_rotl(v1312, int32(25))
	v1320 = v1316 ^ v1308 - base.I32_rotl(v1316, int32(16))
	v1324 = v1320 ^ v1312 - base.I32_rotl(v1320, int32(4))
	v1328 = v1324 ^ v1316 - base.I32_rotl(v1324, v1306)
	goto L179
L181:
	;
	switch v1228 - int32(1) {
	case 0:
		v1294 = v1219
		v1295 = v1220
		v1296 = v1224
		goto L208
	case 1:
		v1287 = v1219
		v1288 = v1220
		v1289 = v1224
		goto L209
	case 2:
		v1280 = v1219
		v1281 = v1220
		v1282 = v1224
		goto L210
	case 3:
		v1274 = v1220
		v1275 = v1224
		goto L211
	case 4:
		v1270 = v1220
		v1271 = v1224
		goto L212
	case 5:
		v1264 = v1220
		v1265 = v1224
		goto L213
	case 6:
		v1258 = v1220
		v1259 = v1224
		goto L214
	case 7:
		v1253 = v1224
		goto L215
	case 8:
		v1248 = v1224
		goto L216
	case 9:
		v1243 = v1224
		goto L217
	case 10:
		goto L218
	default:
		v1301 = v1219
		v1302 = v1220
		v1303 = v1224
		goto L180
	}
L182:
	;
	v1183 = v1052
	v1184 = v1068
	v1185 = v1074
	v1186 = v1074
	v1187 = v1074
	goto L205
L183:
	;
	goto L182
L184:
	;
	goto L185
L185:
	;
	goto L189
L187:
	;
	switch v1126 - int32(1) {
	case 0:
		v1180 = v1117
		goto L194
	case 1:
		v1175 = v1117
		goto L195
	case 2:
		goto L196
	case 3:
		v1168 = v1118
		goto L197
	case 4:
		v1165 = v1118
		goto L198
	case 5:
		v1160 = v1118
		goto L199
	case 6:
		goto L200
	case 7:
		v1151 = v1122
		goto L201
	case 8:
		v1146 = v1122
		goto L202
	case 9:
		v1141 = v1122
		goto L203
	case 10:
		goto L204
	default:
		v1301 = v1117
		v1302 = v1118
		v1303 = v1122
		goto L180
	}
L189:
	;
	goto L190
L190:
	;
	v1081 = v1052
	v1082 = v1068
	v1083 = v1074
	v1084 = v1074
	v1085 = v1074
	goto L191
L191:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	v1088 = v1087 + v1084
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+8))
	v1092 = v1091 + v1085
	v1094 = int32(4)
	v1096 = v1089 + v1083 - v1092 ^ base.I32_rotl(v1092, v1094)
	v1100 = v1088 - v1096 ^ base.I32_rotl(v1096, int32(6))
	v1101 = v1092 + v1088
	v1102 = v1096 + v1101
	v1103 = v1100 + v1102
	v1107 = v1101 - v1100 ^ base.I32_rotl(v1100, int32(8))
	v1111 = v1102 - v1107 ^ base.I32_rotl(v1107, int32(16))
	v1115 = v1103 - v1111 ^ base.I32_rotl(v1111, int32(19))
	v1116 = v1107 + v1103
	v1117 = v1111 + v1116
	v1118 = v1115 + v1117
	v1122 = v1116 - v1115 ^ base.I32_rotl(v1115, v1094)
	v1123 = int32(12)
	v1124 = v1081 + v1123
	v1126 = v1082 - v1123
	if base.Ui32(int32(11)) < base.Ui32(v1126) {
		v1081 = v1124
		v1082 = v1126
		v1083 = v1117
		v1084 = v1118
		v1085 = v1122
		goto L191
	} else {
		goto L193
	}
L192:
	;
	goto L187
L193:
	;
	goto L192
L194:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	v1301 = v1180 + v1181
	v1302 = v1118
	v1303 = v1122
	goto L180
L195:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+1)))
	v1180 = v1176<<(uint(int32(8))%32) + v1175
	goto L194
L196:
	;
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+2)))
	v1175 = v1171<<(uint(int32(16))%32) + v1117
	goto L195
L197:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	v1301 = v1169 + v1117
	v1302 = v1168
	v1303 = v1122
	goto L180
L198:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+4)))
	v1168 = v1165 + v1166
	goto L197
L199:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+5)))
	v1165 = v1161<<(uint(int32(8))%32) + v1160
	goto L198
L200:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+6)))
	v1160 = v1156<<(uint(int32(16))%32) + v1118
	goto L199
L201:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+4))
	v1301 = v1152 + v1117
	v1302 = v1154 + v1118
	v1303 = v1151
	goto L180
L202:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+8)))
	v1151 = v1147<<(uint(int32(8))%32) + v1146
	goto L201
L203:
	;
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+9)))
	v1146 = v1142<<(uint(int32(16))%32) + v1141
	goto L202
L204:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+10)))
	v1141 = v1137<<(uint(int32(24))%32) + v1122
	goto L203
L205:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+4))
	v1190 = v1189 + v1186
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+8))
	v1194 = v1193 + v1187
	v1196 = int32(4)
	v1198 = v1191 + v1185 - v1194 ^ base.I32_rotl(v1194, v1196)
	v1202 = v1190 - v1198 ^ base.I32_rotl(v1198, int32(6))
	v1203 = v1194 + v1190
	v1204 = v1198 + v1203
	v1205 = v1202 + v1204
	v1209 = v1203 - v1202 ^ base.I32_rotl(v1202, int32(8))
	v1213 = v1204 - v1209 ^ base.I32_rotl(v1209, int32(16))
	v1217 = v1205 - v1213 ^ base.I32_rotl(v1213, int32(19))
	v1218 = v1209 + v1205
	v1219 = v1213 + v1218
	v1220 = v1217 + v1219
	v1224 = v1218 - v1217 ^ base.I32_rotl(v1217, v1196)
	v1225 = int32(12)
	v1226 = v1183 + v1225
	v1228 = v1184 - v1225
	if base.Ui32(int32(11)) < base.Ui32(v1228) {
		v1183 = v1226
		v1184 = v1228
		v1185 = v1219
		v1186 = v1220
		v1187 = v1224
		goto L205
	} else {
		goto L207
	}
L206:
	;
	goto L181
L207:
	;
	goto L206
L208:
	;
	v1297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226))))
	v1301 = v1294 + v1297
	v1302 = v1295
	v1303 = v1296
	goto L180
L209:
	;
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+1)))
	v1294 = v1290<<(uint(int32(8))%32) + v1287
	v1295 = v1288
	v1296 = v1289
	goto L208
L210:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+2)))
	v1287 = v1283<<(uint(int32(16))%32) + v1280
	v1288 = v1281
	v1289 = v1282
	goto L209
L211:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+3)))
	v1280 = v1276<<(uint(int32(24))%32) + v1219
	v1281 = v1274
	v1282 = v1275
	goto L210
L212:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+4)))
	v1274 = v1270 + v1272
	v1275 = v1271
	goto L211
L213:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+5)))
	v1270 = v1266<<(uint(int32(8))%32) + v1264
	v1271 = v1265
	goto L212
L214:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+6)))
	v1264 = v1260<<(uint(int32(16))%32) + v1258
	v1265 = v1259
	goto L213
L215:
	;
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+7)))
	v1258 = v1254<<(uint(int32(24))%32) + v1220
	v1259 = v1253
	goto L214
L216:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+8)))
	v1253 = v1249<<(uint(int32(8))%32) + v1248
	goto L215
L217:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+9)))
	v1248 = v1244<<(uint(int32(16))%32) + v1243
	goto L216
L218:
	;
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+10)))
	v1243 = v1239<<(uint(int32(24))%32) + v1224
	goto L217
L219:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1338 = v1055 + v1336
	goto L221
L220:
	;
	v1338 = v1055
	goto L221
L221:
	;
	v1341 = (v1055 + int32(1)) & v1333
	if base.Ui32(v1338-v1334) < base.Ui32(v1051) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1346 = v1035 + v1341*int32(40)
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346)+20)))
	if v1347 != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v1447 = v1051 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1447) {
		goto L241
	} else {
		goto L242
	}
L225:
	;
	v1350 = v1341
	v1354 = int32(0)
	goto L228
L226:
	;
	v1383 = v1341
	v1389 = v1346
	goto L227
L227:
	;
	if v1383 != v1055 {
		goto L235
	} else {
		goto L236
	}
L228:
	;
	v1365 = v1354 + int32(1)
	if int32(151) <= v1365 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1383 = v1377
	v1389 = v1380
	goto L227
L230:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1370 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1368), base.F64_convert_i64_u(v1370)), float64(0.1)) != 0 {
		goto L44
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1377 = (v1350 + int32(1)) & v1333
	v1380 = v1035 + v1377*int32(40)
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1380)+20)))
	if v1381 != 0 {
		v1350 = v1377
		v1354 = v1365
		goto L228
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	goto L229
L235:
	;
	v1399 = v1383
	v1405 = v1389
	goto L238
L236:
	;
	goto L237
L237:
	;
	v1469 = v1052
	goto L175
L238:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1416 = v1413 & (v1399 - int32(1))
	v1419 = v1035 + v1416*int32(40)
	v1420 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+32)) = v1420
	v1422 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+24)) = v1422
	v1424 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+16)) = v1424
	v1426 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+8)) = v1426
	v1428 = *(*int64)(unsafe.Add(mBase, uint32(v1419)))
	*(*int64)(unsafe.Add(mBase, uint32(v1405))) = v1428
	if v1416 != v1055 {
		v1399 = v1416
		v1405 = v1419
		goto L238
	} else {
		goto L240
	}
L239:
	;
	goto L237
L240:
	;
	goto L239
L241:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1452 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1450), base.F64_convert_i64_u(v1452)), float64(0.1)) != 0 {
		goto L44
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1459 = v1035 + v1341*int32(40)
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459)+20)))
	if v1460 != 0 {
		v1051 = v1447
		v1052 = v1459
		v1055 = v1341
		goto L176
	} else {
		goto L245
	}
L244:
	;
	goto L243
L245:
	;
	goto L177
L246:
	;
	F_errmsg_internal(m, int32(_a_F_blockreftable_insert_0), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L60
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_blockreftable_insert_1), int32(635), int32(_a_F_blockreftable_insert_2))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L60
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errmsg_internal(m, int32(_a_F_blockreftable_insert_3), int32(0))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L60
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_blockreftable_insert_1), int32(332), int32(_a_F_blockreftable_insert_4))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L60
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bloptions(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_bloptions[0]))
	v8 = F_build_reloptions(m, l0, l1, v4, int32(136), int32(_a_F_bloptions_0), int32(33))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = base.I32_div_s(v12+int32(15), int32(16))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v16
		} else {
		}
		return v8
	}
}
func F_blvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
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
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
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
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v463 int32
	_ = v463
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	v18 = m.G0
	v20 = v18 - int32(128)
	m.G0 = v20
	v24 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if int32(0) < v206 {
		goto L47
	} else {
		goto L48
	}
L2:
	;
	return int32(0)
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+92))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	v34 = F_get_opfamily_name(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L44
	}
L7:
	;
	v38 = base.I64_extend_i32_u(v33)
	v39 = int64(0)
	v41 = F_SearchSysCacheList(m, int32(4), int32(1), v38, v39, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v43 = int32(1)
	v46 = int64(0)
	v48 = F_SearchSysCacheList(m, int32(5), v43, v38, v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v50 <= int32(0) {
		v192 = v43
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v32
	goto L13
L12:
	;
	v53 = v31
	goto L13
L13:
	;
	v59 = int32(0)
	v60 = v43
	goto L14
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(-64)+v59<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+72))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
	v80 = v78 + v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v81 == v82 {
		v109 = v60
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v192 = v169
	goto L1
L16:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if v110 != v31 {
		v169 = v109
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v84 = int32(0)
	v87 = F_errstart(m, int32(17), v84)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v87 == int32(0) {
		v109 = v84
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v95 = F_format_procedure(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v34
	F_errmsg(m, int32(_a_F_blvalidate_0), v20+int32(112))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(84), int32(_a_F_blvalidate_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v109 = v84
	goto L16
L24:
	;
	v173 = v59 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v173 < v174 {
		v59 = v173
		v60 = v169
		goto L14
	} else {
		goto L43
	}
L25:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+16)))
	switch v112 - int32(1) {
	case 0:
		goto L30
	case 1:
		goto L28
	default:
		goto L29
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L39
	}
L27:
	;
	v139 = int32(0)
	v142 = F_errstart(m, int32(17), v139)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L37
	}
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v137 = F_check_amoptsproc_signature(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L35
	}
L29:
	;
	v127 = int32(0)
	v130 = F_errstart(m, int32(17), v127)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L33
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v53
	v119 = int32(1)
	v123 = F_check_amproc_signature(m, v115, int32(23), int32(0), v119, v119, v20+int32(96))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v123 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v169 = v109
	goto L24
L33:
	;
	if v130 == int32(0) {
		v169 = v127
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v148 = int32(_a_F_blvalidate_3)
	v149 = int32(111)
	goto L26
L35:
	;
	if v137 != 0 {
		v169 = v109
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	if v142 == int32(0) {
		v169 = v139
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v148 = int32(_a_F_blvalidate_4)
	v149 = int32(123)
	goto L26
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v80)+20))
	v154 = F_format_procedure(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v34
	F_errmsg(m, v148, v20+int32(80))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), v149, int32(_a_F_blvalidate_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v169 = int32(0)
	goto L24
L43:
	;
	goto L15
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_blvalidate_5), v20)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(50), int32(_a_F_blvalidate_2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v214 = int32(0)
	v215 = v192
	goto L50
L48:
	;
	v339 = v192
	goto L49
L49:
	;
	v353 = F_identify_opfamily_groups(m, v41, v48)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L2
	} else {
		goto L83
	}
L50:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(-64)+v214<<(uint(int32(2))%32))))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+72))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+22)))
	v235 = v233 + v234
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235)+16)))
	if v236 == int32(1) {
		v266 = v215
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v339 = v331
	goto L49
L52:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+18)))
	if v268 == int32(115) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v239 = int32(0)
	v242 = F_errstart(m, int32(17), v239)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v242 == int32(0) {
		v266 = v239
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	v250 = F_format_operator(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v252 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v34
	F_errmsg(m, int32(_a_F_blvalidate_6), v20-int32(-64))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(143), int32(_a_F_blvalidate_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v266 = v239
	goto L52
L60:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	v304 = F_check_amop_signature(m, v300, int32(16), v302, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L72
	}
L61:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v235)+28))
	if v271 == int32(0) {
		v299 = v266
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v274 = int32(0)
	v277 = F_errstart(m, int32(17), v274)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v277 == int32(0) {
		v299 = v274
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	v285 = F_format_operator(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v34
	F_errmsg(m, int32(_a_F_blvalidate_7), v20+int32(48))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(155), int32(_a_F_blvalidate_2))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v299 = v274
	goto L60
L71:
	;
	v333 = v214 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if v333 < v334 {
		v214 = v333
		v215 = v331
		goto L50
	} else {
		goto L80
	}
L72:
	;
	if v304 != 0 {
		v331 = v299
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v306 = int32(0)
	v309 = F_errstart(m, int32(17), v306)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	if v309 == int32(0) {
		v331 = v306
		goto L71
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	v317 = F_format_operator(m, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v34
	F_errmsg(m, int32(_a_F_blvalidate_8), v20+int32(32))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(168), int32(_a_F_blvalidate_2))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v331 = v306
	goto L71
L80:
	;
	goto L51
L81:
	;
	F_ReleaseCatCacheList(m, v48)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L2
	} else {
		goto L122
	}
L82:
	;
	v483 = int32(0)
	v486 = F_errstart(m, int32(17), v483)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L117
	}
L83:
	;
	if v353 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v357 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v446 == int32(0) {
		goto L82
	} else {
		goto L115
	}
L86:
	;
	v446 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v361 = int32(0)
	if v357 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v365 = int32(0)
	if v365 < v357 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v417 = v361
	v419 = v361
	goto L91
L91:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434+v417<<(uint(int32(2))%32))))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	if v439 != v31 {
		v446 = v419
		goto L85
	} else {
		goto L111
	}
L92:
	;
	v368 = v357
	goto L94
L93:
	;
	v368 = v365
	goto L94
L94:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v374 = int32(0)
	v376 = v374
	v378 = v361
	v380 = v374
	goto L95
L95:
	;
	v395 = v373 + v376<<(uint(int32(2))%32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	if v31 == v397 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v368&int32(1) == int32(0) {
		v446 = v409
		goto L85
	} else {
		goto L110
	}
L97:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v399 == v31 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v402 = v378
	goto L99
L99:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if v31 == v404 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v401 = v396
	goto L102
L101:
	;
	v401 = v378
	goto L102
L102:
	;
	v402 = v401
	goto L99
L103:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v406 == v31 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v409 = v402
	goto L105
L105:
	;
	v410 = int32(2)
	v411 = v376 + v410
	v413 = v380 + v410
	if v413 != v368&int32(2147483646) {
		v376 = v411
		v378 = v409
		v380 = v413
		goto L95
	} else {
		goto L109
	}
L106:
	;
	v408 = v403
	goto L108
L107:
	;
	v408 = v402
	goto L108
L108:
	;
	v409 = v408
	goto L105
L109:
	;
	goto L96
L110:
	;
	v417 = v411
	v419 = v409
	goto L91
L111:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v441 == v31 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v443 = v438
	goto L114
L113:
	;
	v443 = v419
	goto L114
L114:
	;
	v446 = v443
	goto L85
L115:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+16)))
	if v463&int32(2) != 0 {
		v511 = v339
		goto L81
	} else {
		goto L116
	}
L116:
	;
	goto L82
L117:
	;
	if v486 == int32(0) {
		v511 = v483
		goto L81
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v30 + int32(8)
	F_errmsg(m, int32(_a_F_blvalidate_9), v20+int32(16))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(206), int32(_a_F_blvalidate_2))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v511 = v483
	goto L81
L122:
	;
	F_ReleaseCatCacheList(m, v41)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_ReleaseCatCache(m, v24)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	m.G0 = v20 + int32(128)
	return v511 & int32(1)
}
func F_boollt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3) & base.B2i32(v5 != v3))
}
func F_boolrecv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pq_getmsgbyte(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v3 != int32(0)))
	}
}
func F_bounds_adjacent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v7 == int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
		if v6&int32(1) != 0 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
			if v10 == v13 {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
				return base.B2i32(v79 != v80)
			} else {
				if v10&int32(1) != 0 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v43 == int32(0) {
						return int32(0)
					} else {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
						v49 = int32(1)
						v50 = v48 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v50)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
						v54 = v52 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)) = uint8(v54)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v49)
						v58 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)) = uint8(v58)
						v62 = F_make_range(m, l0, l1, l2, v58, v58)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							v68 = int32(1)
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
							return v70 & v68
						}
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if v10&int32(1) != 0 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v43 == int32(0) {
					return int32(0)
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
					v49 = int32(1)
					v50 = v48 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
					v54 = v52 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)) = uint8(v54)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v49)
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)) = uint8(v58)
					v62 = F_make_range(m, l0, l1, l2, v58, v58)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						v68 = int32(1)
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
						return v70 & v68
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		if v6&int32(1) != 0 {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
			if v25 == int32(0) {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v43 == int32(0) {
					return int32(0)
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
					v49 = int32(1)
					v50 = v48 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
					v54 = v52 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)) = uint8(v54)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v49)
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)) = uint8(v58)
					v62 = F_make_range(m, l0, l1, l2, v58, v58)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						v68 = int32(1)
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
						return v70 & v68
					}
				}
			} else {
				return int32(0)
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
			v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v34 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			v35 = F_FunctionCall2Coll(m, l0+int32(212), v32, v33, v34)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = base.I32_wrap_i64(v35)
				if int32(0) <= v39 {
					if v39 == int32(0) {
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
						return base.B2i32(v79 != v80)
					} else {
						return int32(0)
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v43 == int32(0) {
						return int32(0)
					} else {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
						v49 = int32(1)
						v50 = v48 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)) = uint8(v50)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
						v54 = v52 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)) = uint8(v54)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v49)
						v58 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)) = uint8(v58)
						v62 = F_make_range(m, l0, l1, l2, v58, v58)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							v68 = int32(1)
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
							return v70 & v68
						}
					}
				}
			}
		}
	}
}
func F_bpcharcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v23 + v11
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = v51
	goto L18
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v30 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = int32(16)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(4)
	goto L16
L15:
	;
	v40 = v33
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v57 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v78 = v76 & v74
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v51 & (v51 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = v57 - int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v67))))
	if v69 == int32(32) {
		v57 = v67
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v73 = v57
	goto L20
L25:
	;
	v79 = v74
	goto L27
L26:
	;
	v79 = int32(4)
	goto L27
L27:
	;
	v80 = v79 + v16
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v107
	goto L39
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v86 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v97 = int32(1)
	if v78 != 0 {
		v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v89 = int32(16)
	goto L34
L33:
	;
	v89 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v96 = int32(4)
	goto L37
L36:
	;
	v96 = v89
	goto L37
L37:
	;
	v107 = v96
	goto L28
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v113 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = F_varstr_cmp(m, v24, v73, v80, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v128 = v107 & (v107 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v123 = v113 - int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v123))))
	if v125 == int32(32) {
		v113 = v123
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v128 = v113
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v133 != v11 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v137 != v16 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v16)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return base.I64_extend_i32_s(v131)
L54:
	;
	goto L53
}
func F_bpcharlt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v23 + v11
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = v51
	goto L18
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v30 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = int32(16)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(4)
	goto L16
L15:
	;
	v40 = v33
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v57 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v78 = v76 & v74
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v51 & (v51 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = v57 - int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v67))))
	if v69 == int32(32) {
		v57 = v67
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v73 = v57
	goto L20
L25:
	;
	v79 = v74
	goto L27
L26:
	;
	v79 = int32(4)
	goto L27
L27:
	;
	v80 = v79 + v16
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v107
	goto L39
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v86 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v97 = int32(1)
	if v78 != 0 {
		v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v89 = int32(16)
	goto L34
L33:
	;
	v89 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v96 = int32(4)
	goto L37
L36:
	;
	v96 = v89
	goto L37
L37:
	;
	v107 = v96
	goto L28
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v113 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = F_varstr_cmp(m, v24, v73, v80, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v128 = v107 & (v107 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v123 = v113 - int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v123))))
	if v125 == int32(32) {
		v113 = v123
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v128 = v113
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v133 != v11 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v137 != v16 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v16)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return base.I64_extend_i32_u(int32(base.Ui32(v131) >> (uint(int32(31)) % 32)))
L54:
	;
	goto L53
}
func F_brincostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v150 int32
	_ = v150
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
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v237 float64
	_ = v237
	var v238 int32
	_ = v238
	var v244 float64
	_ = v244
	var v245 float64
	_ = v245
	var v248 float64
	_ = v248
	var v259 float64
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int64
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 float32
	_ = v396
	var v399 float64
	_ = v399
	var v400 float64
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 float64
	_ = v447
	var v448 int32
	_ = v448
	var v449 float64
	_ = v449
	var v456 float64
	_ = v456
	var v458 float64
	_ = v458
	var v459 float64
	_ = v459
	var v460 float64
	_ = v460
	var v461 float64
	_ = v461
	var v469 float64
	_ = v469
	var v471 float64
	_ = v471
	var v472 int32
	_ = v472
	var v473 float64
	_ = v473
	var v474 int32
	_ = v474
	var v475 float64
	_ = v475
	var v478 float64
	_ = v478
	var v480 float64
	_ = v480
	var v485 float64
	_ = v485
	var v488 float64
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	v9 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(96)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v30 == v9 {
		v150 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v161 != 0 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v33 <= int32(0) {
		v150 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v45 = v9
	v50 = v33
	v51 = v9
	goto L4
L4:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v45<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v65 == int32(0) {
		v122 = v50
		v123 = v51
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v150 = v123
	goto L1
L6:
	;
	v133 = v45 + int32(1)
	if v133 < v122 {
		v45 = v133
		v50 = v122
		v51 = v123
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v68 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v69 <= v68 {
		v122 = v50
		v123 = v51
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v82 = v68
	v87 = v51
	goto L9
L9:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v82<<(uint(int32(2))%32))))
	v101 = F_lappend(m, v87, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v122 = v107
	v123 = v101
	goto L6
L11:
	;
	return
L12:
	;
	v104 = v82 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v104 < v105 {
		v82 = v104
		v87 = v101
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	F_get_tablespace_page_costs(m, v177, v27+int32(80), v27+int32(88))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L19
	}
L16:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	v175 = v161 + v162<<(uint(int32(2))%32)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+52))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	v175 = v168 + v169<<(uint(int32(2))%32) - int32(4)
	goto L15
L19:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+105)))
	if v184 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6))) = int64(0)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v262 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L21:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v189 = F_index_open(m, v187, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v159)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = int32(128)
	v244 = base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v238), float64(0.0078125)))
	v245 = float64(1)
	if base.F64_gt(v244, v245) != 0 {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v192 = v27 + int32(72)
	v194 = F_ReadBuffer(m, v189, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	F_LockBufferInternal(m, v194, int32(1))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	if v194 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v219 - int32(1)
	F_UnlockReleaseBuffer(m, v194)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L31
	}
L28:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[0]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(v194^int32(-1))<<(uint(int32(2))%32))))
	v216 = v208
	goto L27
L29:
	;
	goto L30
L30:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[1]))
	v216 = v210 + v194<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	F_relation_close(m, v189, int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v159)+124))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
	v233 = base.F64_ceil(base.F64_div(base.F64_convert_i32_u(v228), base.F64_convert_i32_u(v230)))
	v234 = float64(1)
	if base.F64_gt(v233, v234) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v237 = v233
	goto L35
L34:
	;
	v237 = v234
	goto L35
L35:
	;
	v259 = v237
	goto L20
L36:
	;
	v248 = v244
	goto L38
L37:
	;
	v248 = v245
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = base.I32_trunc_sat_f64_u(base.F64_add(base.F64_div(v248, float64(1360)), float64(1)))
	v259 = v248
	goto L20
L39:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	v445 = int32(0)
	v447 = F_clauselist_selectivity(m, l0, v150, v444, v445, v445)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L11
	} else {
		goto L86
	}
L40:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v265 <= int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v269 = v29 + int32(4)
	v271 = v176 + int32(16)
	v288 = v9
	goto L42
L42:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v298 = int32(2)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297+v288<<(uint(v298)%32))))
	v302 = int32(*(*int16)(unsafe.Add(mBase, uint32(v301)+14)))
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296+v302<<(uint(v298)%32)))))
	if v306 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L39
L44:
	;
	if v380 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L45:
	;
	v370 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v366))))
	v374 = F_SearchSysCache3(m, int32(65), v370, base.I64_extend16_s(base.I64_extend_i32_u(v367)), int64(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L11
	} else {
		goto L69
	}
L46:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[2]))
	if v308 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v337 = base.I32_extend16_s(v302 + int32(1))
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[3]))
	if v339 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v366 = v271
	v367 = v306
	goto L45
L50:
	;
	goto L51
L51:
	;
	v311 = base.I32_extend16_s(v306)
	v314 = m.T0[v308].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v176, v311, v27+int32(40))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	if v314 == int32(0) {
		v366 = v271
		v367 = v311
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v318 == int32(0) {
		v380 = v318
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v321 != 0 {
		v380 = v318
		goto L44
	} else {
		goto L55
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_brincostestimate_0), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_brincostestimate_1), int32(_a_F_brincostestimate_2), int32(_a_F_brincostestimate_3))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v366 = v269
	v367 = v337
	goto L45
L60:
	;
	goto L61
L61:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v345 = m.T0[v339].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v342, v337, v27+int32(40))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	if v345 == int32(0) {
		v366 = v269
		v367 = v337
		goto L45
	} else {
		goto L63
	}
L63:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v349 == int32(0) {
		v380 = v349
		goto L44
	} else {
		goto L64
	}
L64:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v352 != 0 {
		v380 = v349
		goto L44
	} else {
		goto L65
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errmsg_internal(m, int32(_a_F_brincostestimate_0), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_brincostestimate_1), int32(_a_F_brincostestimate_4), int32(_a_F_brincostestimate_3))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = int32(1704)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v374
	v380 = v374
	goto L44
L70:
	;
	v417 = v288 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v417 < v418 {
		v288 = v417
		goto L42
	} else {
		goto L85
	}
L71:
	;
	v389 = F_get_attstatsslot(m, v27+int32(4), v380, int32(3), int32(0), int32(2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	if v389 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v391 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v408 == int32(0) {
		goto L70
	} else {
		goto L83
	}
L76:
	;
	v399 = float64(0)
	goto L78
L77:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	v396 = *(*float32)(unsafe.Add(mBase, uint32(v395)))
	v399 = base.F64_promote_f32(base.F32_abs(v396))
	goto L78
L78:
	;
	v400 = *(*float64)(unsafe.Add(mBase, uint32(l6)))
	if base.F64_gt(v399, v400) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l6))) = v399
	goto L81
L80:
	;
	goto L81
L81:
	;
	F_free_attstatsslot(m, v27+int32(4))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	goto L75
L83:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	m.T0[v411].(func(*base.Module, int32))(m, v408)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	goto L70
L85:
	;
	goto L43
L86:
	;
	v449 = *(*float64)(unsafe.Add(mBase, uint32(l6)))
	if base.F64_lt(v449, float64(1e-10)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v469
	v471 = F_index_other_operands_eval_cost(m, l0, v150)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L11
	} else {
		goto L96
	}
L88:
	;
	v456 = base.F64_div(base.F64_ceil(base.F64_mul(v259, v447)), v449)
	if base.F64_gt(v259, v456) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v459 = v259
	goto L90
L90:
	;
	v460 = float64(0)
	v461 = base.F64_div(v459, v259)
	if base.F64_lt(v461, v460) != 0 {
		v469 = v460
		goto L87
	} else {
		goto L94
	}
L91:
	;
	v458 = v456
	goto L93
L92:
	;
	v458 = v259
	goto L93
L93:
	;
	v459 = v458
	goto L90
L94:
	;
	if base.F64_gt(v461, float64(1)) == int32(0) {
		v469 = v461
		goto L87
	} else {
		goto L95
	}
L95:
	;
	v469 = float64(1)
	goto L87
L96:
	;
	v473 = *(*float64)(unsafe.Add(mBase, uint32(v27)+88))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v475 = base.F64_convert_i32_u(v474)
	v478 = base.F64_add(v471, base.F64_mul(l2, base.F64_mul(v473, v475)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v478
	v480 = *(*float64)(unsafe.Add(mBase, uint32(v27)+80))
	v485 = base.F64_add(base.F64_mul(base.F64_mul(v480, base.F64_sub(base.F64_convert_i32_u(v160), v475)), l2), v478)
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v485
	v488 = *(*float64)(unsafe.Add(mBase, _c_F_brincostestimate[4]))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v459, base.F64_mul(v488, float64(0.1))), base.F64_convert_i32_u(v492)), v485)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = base.F64_convert_i32_u(v497)
	m.G0 = v27 + int32(96)
	return
}
func F_bringetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
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
	var v73 int32
	_ = v73
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int64
	_ = v345
	var v348 int64
	_ = v348
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v373 int64
	_ = v373
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v419 int32
	_ = v419
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int64
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int64
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v672 int32
	_ = v672
	var v684 int64
	_ = v684
	var v685 int64
	_ = v685
	var v687 int64
	_ = v687
	var v717 int64
	_ = v717
	var v720 int64
	_ = v720
	var v725 int32
	_ = v725
	var v728 int64
	_ = v728
	var v729 int64
	_ = v729
	var v731 int64
	_ = v731
	var v732 int64
	_ = v732
	var v733 int64
	_ = v733
	var v735 int64
	_ = v735
	var v744 int32
	_ = v744
	var v757 int32
	_ = v757
	var v768 int64
	_ = v768
	var v769 int64
	_ = v769
	var v770 int64
	_ = v770
	var v805 int64
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	v3 = int32(0)
	v25 = int64(0)
	v31 = m.G0
	v33 = v31 - int32(16)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v3
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	if v42 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v59 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)))
	if v45 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v53 = v42
	goto L4
L4:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v54 + int64(1)
	goto L1
L5:
	;
	F_pgstat_assoc_relation(m, v35)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int64(0)
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	v53 = v52
	goto L4
L8:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = v60 + int64(1)
	goto L10
L9:
	;
	goto L10
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v66 = F_IndexGetRelation(m, v64, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v69 = F_table_open(m, v66, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v72 = F_RelationGetNumberOfBlocksInFork(m, v69, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_relation_close(m, v69, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v80 = F_palloc0_mul(m, int32(28), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v101 = F_palloc(m, ((v83<<(uint(int32(3))%32)+int32(14))&int32(2147483632)+(v90<<(uint(int32(2))%32)+int32(7))&int32(2147483640)*v83)<<(uint(int32(1))%32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v106 = v104 << (uint(int32(2)) % 32)
	v110 = (v106 + int32(7)) & int32(-8)
	v111 = v101 + v110
	v112 = v111 + v110
	v113 = v112 + v110
	if int32(0) < v104 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v120 = int32(0)
	v121 = v110 + v113
	goto L20
L18:
	;
	v179 = v106
	goto L19
L19:
	;
	if v179 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v148 = int32(2)
	v149 = v120 << (uint(v148) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v101+v149))) = v121
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = int32(7)
	v158 = int32(-8)
	v160 = v121 + (v153<<(uint(v148)%32)+v156)&v158
	*(*int32)(unsafe.Add(mBase, uint32(v149+v111))) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v171 = v120 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v171 < v173 {
		v120 = v171
		v121 = v160 + (v162<<(uint(v148)%32)+v156)&v158
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v179 = v173 << (uint(int32(2)) % 32)
	goto L19
L22:
	;
	goto L21
L23:
	;
	base.MemoryFill(m, v112, int32(0), v179)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v212 = v210 << (uint(int32(2)) % 32)
	if v212 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	base.MemoryFill(m, v113, int32(0), v212)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v215 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v222 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	v331 = F_brin_new_memtuple(m, v41)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L6
	} else {
		goto L46
	}
L32:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v252 = v249 + v222*int32(56)
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+4)))
	v255 = v253 - int32(1)
	v258 = v80 + v255*int32(28)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v259 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v263 = F_index_getprocinfo(m, v35, v253, int32(3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v279 = v255 << (uint(int32(2)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v282 = v280 & int32(1)
	if v282 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0]))
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+16)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+24)) = v269
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v271
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v258)+20)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v258)+16)) = int32(0)
	goto L38
L38:
	;
	goto L36
L39:
	;
	v283 = v111
	goto L41
L40:
	;
	v283 = v101
	goto L41
L41:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279+v283)))
	if v282 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v286 = v113
	goto L44
L43:
	;
	v286 = v112
	goto L44
L44:
	;
	v287 = v286 + v279
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+v288<<(uint(int32(2))%32)))) = v252
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v294 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v293 + v294
	v298 = v222 + v294
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v298 < v299 {
		v222 = v298
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0]))
	v339 = F_AllocSetContextCreateInternal(m, v334, int32(_a_F_bringetbitmap_0), int32(0), int32(_a_F_bringetbitmap_1), int32(_a_F_bringetbitmap_2))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v341 = int32(_a_F_bringetbitmap_3)
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v339
	if v72 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v345 = base.I64_extend_i32_u(v72)
	v348 = base.I64_extend_i32_u(v41)
	v354 = v331
	v367 = v3
	v373 = v25
	v378 = v25
	goto L51
L49:
	;
	v805 = int64(0)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v342
	F_MemoryContextDelete(m, v339)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L6
	} else {
		goto L117
	}
L51:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[1]))
	if v380 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v805 = v768 * int64(10)
	goto L50
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_MemoryContextReset(m, v339)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v391 = F_brinGetTupleForHeapBlock(m, v385, base.I32_wrap_i64(v373), v33+int32(12), v33+int32(6), v33)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	v769 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v770 = v373 + v769
	if base.Ui64(v770) < base.Ui64(v345) {
		v354 = v744
		v367 = v757
		v373 = v770
		v378 = v768
		goto L51
	} else {
		goto L116
	}
L59:
	;
	v684 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v685 = v373 + v684
	if base.Ui64(v685) < base.Ui64(v345) {
		goto L105
	} else {
		goto L106
	}
L60:
	;
	if v391 == int32(0) {
		v659 = v354
		v672 = v367
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v398 = F_brin_copy_tuple(m, v391, v395, v367, v33+int32(8))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	F_UnlockBuffer(m, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v403 = F_brin_deform_tuple(m, v41, v398, v354)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v405 != 0 {
		v659 = v403
		v672 = v398
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	if v407 <= int32(0) {
		v659 = v403
		v672 = v398
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v419 = int32(1)
	goto L67
L67:
	;
	v442 = v419 - int32(1)
	v444 = v442 << (uint(int32(2)) % 32)
	v445 = v112 + v444
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	if v446 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v659 = v403
	v672 = v398
	goto L59
L69:
	;
	v650 = v419 + int32(1)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	if v650 <= v652 {
		v419 = v650
		goto L67
	} else {
		goto L104
	}
L70:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v444+v113)))
	if v450 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	if v453 != 0 {
		v744 = v403
		v757 = v398
		v768 = v378
		goto L58
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v456 = v403 + v419*int32(24)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v444+(v41+int32(20)))))
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+2)))
	if v459 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v446 == int32(0) {
		goto L69
	} else {
		goto L90
	}
L76:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v444+v113)))
	if v463 <= int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v444+v111)))
	v471 = int32(0)
	goto L78
L78:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v467+v471<<(uint(int32(2))%32))))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v503&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L75
L80:
	;
	v518 = v471 + int32(1)
	if v518 != v463 {
		v471 = v518
		goto L78
	} else {
		goto L89
	}
L81:
	;
	if v503&int32(64) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+3)))
	if v510 != 0 {
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v503&int32(128) == int32(0) {
		v744 = v403
		v757 = v398
		v768 = v378
		goto L58
	} else {
		goto L87
	}
L85:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+2)))
	if v511 != 0 {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v744 = v403
	v757 = v398
	v768 = v378
	goto L58
L87:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+3)))
	if v516 != 0 {
		v744 = v403
		v757 = v398
		v768 = v378
		goto L58
	} else {
		goto L88
	}
L88:
	;
	goto L80
L89:
	;
	goto L79
L90:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+3)))
	if v552 != 0 {
		v744 = v403
		v757 = v398
		v768 = v378
		goto L58
	} else {
		goto L91
	}
L91:
	;
	v553 = v444 + v101
	v556 = v80 + v442*int32(28)
	v557 = int32(*(*int16)(unsafe.Add(mBase, uint32(v556)+8)))
	if v557 <= int32(3) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v576 = int32(0)
	goto L99
L93:
	;
	if v446 <= int32(0) {
		goto L69
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+12))
	v570 = F_FunctionCall4Coll(m, v556, v566, v348, base.I64_extend_i32_u(v456), base.I64_extend_i32_u(v564), base.I64_extend_i32_s(v446))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L97
	}
L96:
	;
	goto L92
L97:
	;
	if v570 == int64(0) {
		v744 = v403
		v757 = v398
		v768 = v378
		goto L58
	} else {
		goto L98
	}
L98:
	;
	goto L69
L99:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v604+v576<<(uint(int32(2))%32))))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	v611 = F_FunctionCall3Coll(m, v556, v609, v348, base.I64_extend_i32_u(v456), base.I64_extend_i32_u(v608))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L101
	}
L100:
	;
	goto L69
L101:
	;
	if v611 == int64(0) {
		v744 = v403
		v757 = v398
		v768 = v378
		goto L58
	} else {
		goto L102
	}
L102:
	;
	v616 = v576 + int32(1)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	if v616 < v617 {
		v576 = v616
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	goto L68
L105:
	;
	v687 = v685
	goto L107
L106:
	;
	v687 = v345
	goto L107
L107:
	;
	if base.Ui64(v687-int64(1)) < base.Ui64(v373) {
		v744 = v659
		v757 = v672
		v768 = v378
		goto L58
	} else {
		goto L108
	}
L108:
	;
	v717 = v373
	v720 = v378
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v342
	F_tbm_add_page(m, l1, base.I32_wrap_i64(v717))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L6
	} else {
		goto L111
	}
L110:
	;
	v744 = v659
	v757 = v672
	v768 = v729
	goto L58
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v339
	v728 = int64(1)
	v729 = v720 + v728
	v731 = v717 + v728
	v732 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v733 = v373 + v732
	if base.Ui64(v733) < base.Ui64(v345) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v735 = v733
	goto L114
L113:
	;
	v735 = v345
	goto L114
L114:
	;
	if base.Ui64(v731) <= base.Ui64(v735-int64(1)) {
		v717 = v731
		v720 = v729
		goto L109
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	goto L52
L117:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v810 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_ReleaseBuffer(m, v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	m.G0 = v33 + int32(16)
	return v805
L121:
	;
	goto L120
}
func F_brininsertcleanup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = int32(0)
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_brinRevmapTerminate(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_pfree(m, v3)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_btboolcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3)) - base.I64_extend_i32_u(base.B2i32(v6 != v3))
}
func F_btcharcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return v2 - v3
}
func F_btestimateparallelscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v10 = l1<<(uint(int32(2))%32) + int32(40)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+10)))
	if v12 == int32(1) {
		v61 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v61
L2:
	;
	v19 = F_datumEstimateSpace(m, int64(0), int32(0), int32(1), int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v12 < int32(2) {
		v61 = v10
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = int32(1)
	v28 = v10
	goto L6
L6:
	;
	v33 = F_add_size(m, v28, int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v61 = v55
	goto L1
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v38 = v35 + v27<<(uint(int32(3))%32)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
	if v39 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v57 = v27 + int32(1)
	if v57 != v12 {
		v27 = v57
		v28 = v55
		goto L6
	} else {
		goto L17
	}
L10:
	;
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+22)))
	v46 = F_datumEstimateSpace(m, int64(0), int32(0), int32(1), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v50 = F_add_size(m, v33, v19)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	v48 = F_add_size(m, v33, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v55 = v48
	goto L9
L15:
	;
	v53 = F_add_size(m, v50, int32(2704))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v55 = v53
	goto L9
L17:
	;
	goto L7
}
func F_btfloat4cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 float32
	_ = v11
	var v12 float32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v35 int64
	_ = v35
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int32(2147483647)
	v9 = v7 & v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = base.F32_reinterpret_i32(v10)
	v12 = base.F32_reinterpret_i32(v7)
	v14 = v10 & v8
	if base.Ui32(int32(2139095041)) <= base.Ui32(v14) {
		v24 = base.B2i32(base.Ui32(v9) < base.Ui32(int32(2139095041)))
		v35 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v14))|base.F32_gt(v11, v12))&v24)
	} else {
		v19 = int64(1)
		if base.F32_lt(v11, v12) != 0 {
			v35 = v19
		} else {
			if base.Ui32(int32(2139095040)) < base.Ui32(v9) {
				v35 = v19
			} else {
				v24 = int32(1)
				v35 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v14))|base.F32_gt(v11, v12))&v24)
			}
		}
	}
	return v35
}
func F_btfloat8cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 float64
	_ = v11
	var v14 int64
	_ = v14
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v35 int64
	_ = v35
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = base.I64_reinterpret_f64(v11) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v14) {
		v24 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v35 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v14))|base.F64_lt(v7, v11))&v24)
	} else {
		v19 = int64(1)
		if base.F64_gt(v7, v11) != 0 {
			v35 = v19
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v35 = v19
			} else {
				v24 = int32(1)
				v35 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v14))|base.F64_lt(v7, v11))&v24)
			}
		}
	}
	return v35
}
func F_btfloat8fastcmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	v7 = int64(9223372036854775807)
	v8 = l0 & v7
	v9 = base.F64_reinterpret_i64(l1)
	v10 = base.F64_reinterpret_i64(l0)
	v12 = l1 & v7
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v12) {
		v22 = base.B2i32(base.Ui64(v8) < base.Ui64(int64(9218868437227405313)))
		v31 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v12))|base.F64_gt(v9, v10))&v22
	} else {
		v17 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v8))|base.F64_lt(v9, v10) != 0 {
			v31 = v17
		} else {
			v22 = v17
			v31 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v12))|base.F64_gt(v9, v10))&v22
		}
	}
	return v31
}
func F_btint4sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(199)
	return int64(0)
}
func F_btint8sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(202)
	return int64(0)
}
func F_btoid8fastcmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	return base.B2i32(base.Ui64(l1) < base.Ui64(l0)) - base.B2i32(base.Ui64(l0) < base.Ui64(l1))
}
func F_btrim(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn14239(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_bttextsortsupport(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14238(m, l0, int32(25))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_bttranslatecmptype(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0-int32(1)) < base.Ui32(int32(5)) {
		v8 = l0
	} else {
		v8 = int32(0)
	}
	return v8 & int32(_a_F_bttranslatecmptype_0)
}
func F_btvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
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
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v428 int64
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int64
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	v17 = m.G0
	v19 = v17 - int32(240)
	m.G0 = v19
	v23 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	if int32(0) < v235 {
		goto L47
	} else {
		goto L48
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v33 = F_get_opfamily_name(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L44
	}
L7:
	;
	v37 = base.I64_extend_i32_u(v32)
	v38 = int64(0)
	v40 = F_SearchSysCacheList(m, int32(4), int32(1), v37, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(1)
	v45 = int64(0)
	v47 = F_SearchSysCacheList(m, int32(5), v42, v37, v45, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v49 <= int32(0) {
		v223 = v42
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(0)
	v58 = v42
	goto L11
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v47-int32(-64)+v54<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+72))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v76 = v74 + v75
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	switch v77 - int32(1) {
	case 0:
		goto L22
	case 1:
		goto L16
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	case 5:
		goto L18
	default:
		goto L17
	}
L12:
	;
	v223 = v199
	goto L1
L13:
	;
	v203 = v54 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v203 < v204 {
		v54 = v203
		v58 = v199
		goto L11
	} else {
		goto L43
	}
L14:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L39
	}
L15:
	;
	v164 = int32(0)
	v167 = F_errstart(m, int32(17), v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L37
	}
L16:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = int32(2281)
	v155 = int32(1)
	v160 = F_check_amproc_signature(m, v151, int32(2278), v155, v155, v155, v19+int32(160))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L35
	}
L17:
	;
	v142 = int32(0)
	v145 = F_errstart(m, int32(17), v142)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L33
	}
L18:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = int32(2281)
	v133 = int32(1)
	v138 = F_check_amproc_signature(m, v129, int32(2278), v133, v133, v133, v19+int32(224))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L31
	}
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v125 = F_check_amoptsproc_signature(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L29
	}
L20:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = int32(26)
	v115 = int32(1)
	v120 = F_check_amproc_signature(m, v111, int32(16), v115, v115, v115, v19+int32(208))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L27
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+188)) = int64(68719476752)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v94
	v103 = int32(5)
	v107 = F_check_amproc_signature(m, v93, int32(16), int32(1), v103, v103, v19+int32(176))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L25
	}
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+144)) = v81
	v85 = int32(2)
	v89 = F_check_amproc_signature(m, v80, int32(23), int32(1), v85, v85, v19+int32(144))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v89 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v199 = v58
	goto L13
L25:
	;
	if v107 == int32(0) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v199 = v58
	goto L13
L27:
	;
	if v120 == int32(0) {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v199 = v58
	goto L13
L29:
	;
	if v125 == int32(0) {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v199 = v58
	goto L13
L31:
	;
	if v138 == int32(0) {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v199 = v58
	goto L13
L33:
	;
	if v145 == int32(0) {
		v199 = v142
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v174 = int32(119)
	v176 = int32(_a_F_btvalidate_0)
	goto L14
L35:
	;
	if v160 != 0 {
		v199 = v58
		goto L13
	} else {
		goto L36
	}
L36:
	;
	goto L15
L37:
	;
	if v167 == int32(0) {
		v199 = v164
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v174 = int32(131)
	v176 = int32(_a_F_btvalidate_1)
	goto L14
L39:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v181 = F_format_procedure(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+140)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v33
	F_errmsg(m, v176, v19+int32(128))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), v174, int32(_a_F_btvalidate_4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v199 = int32(0)
	goto L13
L43:
	;
	goto L12
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg_internal(m, int32(_a_F_btvalidate_5), v19)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(61), int32(_a_F_btvalidate_4))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v243 = int32(0)
	v245 = v223
	goto L50
L48:
	;
	v378 = v223
	goto L49
L49:
	;
	v390 = F_identify_opfamily_groups(m, v40, v47)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L83
	}
L50:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(-64)+v243<<(uint(int32(2))%32))))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+72))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+22)))
	v263 = v261 + v262
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+16)))
	if base.Ui32(int32(_a_F_btvalidate_6)) < base.Ui32((v264-int32(6))&int32(_a_F_btvalidate_7)) {
		v300 = v245
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v378 = v369
	goto L49
L52:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+18)))
	if v302 == int32(115) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v271 = int32(0)
	v274 = F_errstart(m, int32(17), v271)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v274 == int32(0) {
		v300 = v271
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	v282 = F_format_operator(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v263)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+124)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v33
	F_errmsg(m, int32(_a_F_btvalidate_8), v19+int32(112))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(151), int32(_a_F_btvalidate_4))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v300 = v271
	goto L52
L60:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v340 = F_check_amop_signature(m, v336, int32(16), v338, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L2
	} else {
		goto L72
	}
L61:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v263)+28))
	if v305 == int32(0) {
		v335 = v300
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v308 = int32(0)
	v311 = F_errstart(m, int32(17), v308)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v311 == int32(0) {
		v335 = v308
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	v319 = F_format_operator(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v33
	F_errmsg(m, int32(_a_F_btvalidate_9), v19+int32(96))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(163), int32(_a_F_btvalidate_4))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v335 = v308
	goto L60
L71:
	;
	v371 = v243 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	if v371 < v372 {
		v243 = v371
		v245 = v369
		goto L50
	} else {
		goto L80
	}
L72:
	;
	if v340 != 0 {
		v369 = v335
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v342 = int32(0)
	v345 = F_errstart(m, int32(17), v342)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	if v345 == int32(0) {
		v369 = v342
		goto L71
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v263)+20))
	v353 = F_format_operator(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v33
	F_errmsg(m, int32(_a_F_btvalidate_10), v19+int32(80))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(176), int32(_a_F_btvalidate_4))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v369 = v342
	goto L71
L80:
	;
	goto L51
L81:
	;
	if v585 != 0 {
		goto L133
	} else {
		goto L134
	}
L82:
	;
	v561 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L2
	} else {
		goto L125
	}
L83:
	;
	if v390 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v394 = int32(0)
	v546 = v394
	v548 = v394
	goto L82
L85:
	;
	goto L86
L86:
	;
	v397 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v398 <= v397 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v525 == int32(0) {
		v585 = v528
		v586 = v529
		v587 = v530
		goto L81
	} else {
		goto L124
	}
L88:
	;
	v525 = int32(1)
	v528 = int32(0)
	v529 = v378
	v530 = v397
	goto L87
L89:
	;
	goto L90
L90:
	;
	v402 = int32(0)
	v406 = v402
	v407 = v402
	v408 = v378
	v409 = v397
	v415 = int32(0)
	goto L91
L91:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v406<<(uint(int32(2))%32))))
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v424)+8))
	if v425 == int64(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v525 = base.B2i32(v518 == int32(0))
	v528 = v514
	v529 = v515
	v530 = v516
	goto L87
L93:
	;
	v520 = v406 + int32(1)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v520 < v521 {
		v406 = v520
		v407 = v514
		v408 = v515
		v409 = v516
		v415 = v518
		goto L91
	} else {
		goto L123
	}
L94:
	;
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v424)+16))
	if v428 == int64(8) {
		v514 = v407
		v515 = v408
		v516 = v409
		v518 = v415
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	if v30 == v431 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v433 == v30 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v436 = v415
	goto L100
L100:
	;
	v437 = F_list_append_unique_oid(m, v407, v431)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L2
	} else {
		goto L104
	}
L101:
	;
	v435 = v424
	goto L103
L102:
	;
	v435 = v415
	goto L103
L103:
	;
	v436 = v435
	goto L100
L104:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v440 = F_list_append_unique_oid(m, v437, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v424)+8))
	if v442 == int64(62) {
		v476 = v408
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v479 = v409 + int32(1)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+16)))
	if v480&int32(2) != 0 {
		v514 = v440
		v515 = v476
		v516 = v479
		v518 = v436
		goto L93
	} else {
		goto L115
	}
L107:
	;
	v445 = int32(0)
	v448 = F_errstart(m, int32(17), v445)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	if v448 == int32(0) {
		v476 = v445
		goto L106
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v456 = F_format_type_be(m, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v459 = F_format_type_be(m, v458)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v33
	F_errmsg(m, int32(_a_F_btvalidate_11), v19-int32(-64))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(235), int32(_a_F_btvalidate_4))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v476 = v445
	goto L106
L115:
	;
	v483 = int32(0)
	v486 = F_errstart(m, int32(17), v483)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	if v486 == int32(0) {
		v514 = v440
		v515 = v483
		v516 = v479
		v518 = v436
		goto L93
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v494 = F_format_type_be(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v497 = F_format_type_be(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v33
	F_errmsg(m, int32(_a_F_btvalidate_12), v19+int32(48))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(245), int32(_a_F_btvalidate_4))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v514 = v440
	v515 = v483
	v516 = v479
	v518 = v436
	goto L93
L123:
	;
	goto L92
L124:
	;
	v546 = v528
	v548 = v530
	goto L82
L125:
	;
	if v561 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v585 = v546
	v586 = int32(0)
	v587 = v548
	goto L81
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v29 + int32(8)
	F_errmsg(m, int32(_a_F_btvalidate_13), v19+int32(32))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(257), int32(_a_F_btvalidate_4))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	F_ReleaseCatCacheList(m, v47)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L2
	} else {
		goto L142
	}
L133:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	v602 = v598 * v598
	goto L135
L134:
	;
	v602 = int32(0)
	goto L135
L135:
	;
	if v602 == v587 {
		v627 = v586
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v604 = int32(0)
	v607 = F_errstart(m, int32(17), v604)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	if v607 == int32(0) {
		v627 = v604
		goto L132
	} else {
		goto L138
	}
L138:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v33
	F_errmsg(m, int32(_a_F_btvalidate_14), v19+int32(16))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(273), int32(_a_F_btvalidate_4))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v627 = v604
	goto L132
L142:
	;
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	F_ReleaseCatCache(m, v23)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	m.G0 = v19 + int32(240)
	return v627 & int32(1)
}
func F_build_merged_partition_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	if l1 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v13 = v11
	} else {
		v13 = int32(0)
	}
	v15 = F_palloc(m, int32(36))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
		v22 = F_palloc_mul(m, int32(4), v13)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v22
			if l1 == int32(0) {
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v27 <= int32(0) {
				} else {
					v37 = int32(0)
					for {
						v41 = v37 << (uint(int32(2)) % 32)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v44+v41)))
						*(*int32)(unsafe.Add(mBase, uint32(v41+v42))) = v46
						v49 = v37 + int32(1)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v49 < v50 {
							v37 = v49
							continue
						} else {
							break
						}
						break
					}
				}
			}
			if l0 == int32(114) {
				v65 = F_palloc_mul(m, int32(4), v13)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v65
					if l2 == int32(0) {
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if v70 <= int32(0) {
						} else {
							v81 = int32(0)
							for {
								v85 = v81 << (uint(int32(2)) % 32)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v88+v85)))
								*(*int32)(unsafe.Add(mBase, uint32(v85+v86))) = v90
								v93 = v81 + int32(1)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								if v93 < v94 {
									v81 = v93
									continue
								} else {
									break
								}
								break
							}
						}
					}
					v109 = F_lappend_int(m, l3, int32(-1))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v116 = v109
						v121 = v13 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v121
						v124 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v124
						v128 = F_palloc_mul(m, int32(4), v121)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v128
							if v116 == int32(0) {
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
								if v133 <= int32(0) {
								} else {
									v143 = v124
									for {
										v147 = v143 << (uint(int32(2)) % 32)
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
										v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147)))
										*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
										v155 = v143 + int32(1)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
										if v155 < v156 {
											v143 = v155
											continue
										} else {
											break
										}
										break
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l4
							return v15
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
				v116 = l3
				v121 = v13
				*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v121
				v124 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v124
				v128 = F_palloc_mul(m, int32(4), v121)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v128
					if v116 == int32(0) {
					} else {
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
						if v133 <= int32(0) {
						} else {
							v143 = v124
							for {
								v147 = v143 << (uint(int32(2)) % 32)
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v147)))
								*(*int32)(unsafe.Add(mBase, uint32(v147+v148))) = v152
								v155 = v143 + int32(1)
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
								if v155 < v156 {
									v143 = v155
									continue
								} else {
									break
								}
								break
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = l4
					return v15
				}
			}
		}
	}
}
func F_build_subplan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v138 int32
	_ = v138
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 float64
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 float64
	_ = v348
	var v350 float64
	_ = v350
	var v353 float64
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 float64
	_ = v364
	var v366 int32
	_ = v366
	var v370 float64
	_ = v370
	var v371 float64
	_ = v371
	var v374 float64
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
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
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v554 int32
	_ = v554
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
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
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	v10 = l9
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v22 = F_palloc0(m, int32(72))
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(23)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v29
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+38)) = uint8(v10)
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v55
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+39)) = uint8(v60)
	if l4 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = int64(-4294965018)
	v55 = int32(0)
	goto L3
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+26)))
	if v38 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v40 = F_exprType(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v44 = F_exprTypmod(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v48 = F_exprCollation(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v55 = v48
	goto L3
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L150
	}
L11:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	v589 = F_lappend(m, v588, v572)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L134
	}
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if int32(0) < v66 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v158 = int32(1)
	goto L14
L14:
	;
	v159 = int32(0)
	if l5|base.B2i32(v158 == v159) == v159 {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	v82 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v158 = base.B2i32(v138 == int32(0))
	goto L14
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v86 = int32(2)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v82<<(uint(v86)%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if base.B2i32(base.Ui32(v86) <= base.Ui32(v91-int32(9)))&base.B2i32(v91 != int32(61)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+44)) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = l0
	v106 = F_process_sublinks_mutator(m, v90, v19+int32(40))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v108 = v90
	goto L22
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v111 = F_lappend_int(m, v109, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v108 = v106
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v115 = F_lappend(m, v114, v108)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v115
	v119 = v82 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v119 < v120 {
		v82 = v119
		goto L18
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	v167 = F_generate_new_exec_param(m, l0, int32(16), int32(-1), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v181 = v158 ^ int32(1)
	if v181|base.B2i32(l5 != int32(4)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v169
	v175 = F_list_make1_impl(m, int32(479), v19+int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v177)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v175
	v572 = l1
	v582 = v167
	goto L11
L32:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v191 = F_exprType(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.B2i32(l5 != int32(6))|v181 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v194 = F_exprTypmod(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v197 = F_exprCollation(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v199 = F_generate_new_exec_param(m, l0, v191, v194, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v201
	v207 = F_list_make1_impl(m, int32(479), v19+int32(12))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v207
	v572 = l1
	v582 = v199
	goto L11
L40:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v221 = F_exprType(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v247 = v22 + int32(12)
	if v158^int32(1)|base.B2i32(l5 != int32(3)) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v223 = F_get_promoted_array_type(m, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v223 == int32(0) {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v228 = F_exprTypmod(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v231 = F_exprCollation(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v233 = F_generate_new_exec_param(m, l0, v223, v228, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v235
	v241 = F_list_make1_impl(m, int32(479), v19+int32(24))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v243 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v241
	v572 = l1
	v582 = v233
	goto L11
L50:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v256 = F_generate_subquery_params(m, l0, v255, v247)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if l5 == int32(5) {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = l0
	v262 = F_convert_testexpr_mutator(m, l7, v19+int32(40))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v265 = F_list_copy(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v267)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v265
	v572 = l1
	v582 = v262
	goto L11
L56:
	;
	v569 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v569)
	v572 = v554
	v582 = v22
	goto L11
L57:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v275 = F_generate_subquery_params(m, l0, v272, v22+int32(40))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v319 = int32(0)
	if base.B2i32(l7 == v319)|l8 == v319 {
		goto L75
	} else {
		goto L76
	}
L60:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v289 = v277
	goto L61
L61:
	;
	if v289 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v302+l6<<(uint(int32(2))%32)-int32(4)))) = v275
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v309 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v296 = v294
	goto L65
L64:
	;
	v296 = int32(0)
	goto L65
L65:
	;
	if v296 < l6 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v299 = F_lappend(m, v289, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L62
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v299
	v289 = v299
	goto L61
L70:
	;
	v312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v312)
	v317 = F_makeNullConst(m, int32(2249), int32(-1), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v554 = l1
	goto L56
L73:
	;
	v572 = l1
	v582 = v317
	goto L11
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v334
	if l5 != int32(2) {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v325 = F_generate_subquery_params(m, l0, v324, v247)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = l8
	v334 = l7
	goto L74
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = l0
	v331 = F_convert_testexpr_mutator(m, l7, v19+int32(40))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v334 = v331
	goto L74
L80:
	;
	v554 = v538
	goto L56
L81:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v520 != 0 {
		v538 = l1
		goto L80
	} else {
		goto L129
	}
L82:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v338 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v339 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v343 = F_EstimateTupleHashTableSpace(m, v339, v340, int32(0))
	mBase = m.M
	if v10|base.B2i32(v343 == int32(-1)) != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v364 = *(*float64)(unsafe.Add(mBase, _c_F_build_subplan[0]))
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_build_subplan[1]))
	v370 = base.F64_mul(base.F64_mul(v364, base.F64_convert_i32_s(v366)), float64(1024))
	v371 = float64(4.294967295e+09)
	if base.F64_lt(v370, v371) != 0 {
		goto L95
	} else {
		goto L96
	}
L85:
	;
	v361 = v343
	goto L87
L86:
	;
	v348 = float64(1)
	v350 = base.F64_mul(v339, float64(0.0625))
	if base.F64_lt(v350, v348) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L84
L88:
	;
	v353 = v348
	goto L90
L89:
	;
	v353 = v350
	goto L90
L90:
	;
	v355 = F_EstimateTupleHashTableSpace(m, v353, v340, int32(0))
	mBase = m.M
	v356 = v355 + v343
	if base.Ui32(v356) < base.Ui32(v343) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v358 = int32(-1)
	goto L93
L92:
	;
	v358 = v356
	goto L93
L93:
	;
	v361 = v358
	goto L87
L94:
	;
	if base.Ui32(base.I32_trunc_sat_f64_u(v374)) <= base.Ui32(v361) {
		goto L81
	} else {
		goto L98
	}
L95:
	;
	v374 = v370
	goto L97
L96:
	;
	v374 = v371
	goto L97
L97:
	;
	goto L94
L98:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v378 = int32(0)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v379 == v378 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v499 == int32(0) {
		goto L81
	} else {
		goto L128
	}
L100:
	;
	v499 = int32(0)
	goto L99
L101:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	switch v382 - int32(17) {
	case 0:
		goto L104
	default:
		goto L100
	case 4:
		goto L103
	}
L102:
	;
	v499 = v472
	goto L99
L103:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v407 != 0 {
		goto L100
	} else {
		goto L113
	}
L104:
	;
	v385 = F_hash_ok_operator(m, v379)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v385 == int32(0) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v379)+28))
	if v389 == int32(0) {
		goto L100
	} else {
		goto L107
	}
L107:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v392 != int32(2) {
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = F_contain_exec_param(m, v396, v377)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v397 != 0 {
		goto L100
	} else {
		goto L110
	}
L110:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v379)+28))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v403 = F_contain_var_clause(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v403 == int32(0) {
		v472 = int32(1)
		goto L102
	} else {
		goto L112
	}
L112:
	;
	goto L100
L113:
	;
	v408 = int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	if v409 == int32(0) {
		v472 = v408
		goto L102
	} else {
		goto L114
	}
L114:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if v412 <= int32(0) {
		v472 = v408
		goto L102
	} else {
		goto L115
	}
L115:
	;
	v421 = v378
	goto L116
L116:
	;
	v431 = int32(0)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+v421<<(uint(int32(2))%32))))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	if v437 != int32(17) {
		v472 = v431
		goto L102
	} else {
		goto L118
	}
L117:
	;
	v472 = v459
	goto L102
L118:
	;
	v440 = F_hash_ok_operator(m, v436)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v440 == int32(0) {
		v472 = v431
		goto L102
	} else {
		goto L120
	}
L120:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v436)+28))
	if v444 == int32(0) {
		v472 = v431
		goto L102
	} else {
		goto L121
	}
L121:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	if v447 != int32(2) {
		v472 = v431
		goto L102
	} else {
		goto L122
	}
L122:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v444)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	v452 = F_contain_exec_param(m, v451, v377)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v452 != 0 {
		v472 = v431
		goto L102
	} else {
		goto L124
	}
L124:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v436)+28))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v457 = F_contain_var_clause(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v457 != 0 {
		v472 = v431
		goto L102
	} else {
		goto L126
	}
L126:
	;
	v459 = int32(1)
	v461 = v421 + v459
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if v461 < v462 {
		v421 = v461
		goto L116
	} else {
		goto L127
	}
L127:
	;
	goto L117
L128:
	;
	v502 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)) = uint8(v502)
	v538 = l1
	goto L80
L129:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_subplan[2])))
	if v522&int32(1) == int32(0) {
		v538 = l1
		goto L80
	} else {
		goto L130
	}
L130:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v529 = v527 - int32(352)
	goto L131
L131:
	;
	if base.B2i32(base.Ui32(v529) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_build_subplan_0))>>(uint(v529)%32)) != 0 {
		v538 = l1
		goto L80
	} else {
		goto L132
	}
L132:
	;
	v535 = F_materialize_finished_plan(m, l1)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v538 = v535
	goto L80
L134:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+8)) = v589
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	v595 = F_lappend(m, v594, l2)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v597)+12)) = v595
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+16))
	v601 = F_lappend(m, v600, l3)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+16)) = v601
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+8))
	if v606 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+4))
	v609 = v607
	goto L139
L138:
	;
	v609 = int32(0)
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v609
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)))
	if v611 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v615 = F_lappend(m, v614, v22)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v618 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v615
	goto L142
L144:
	;
	F_cost_subplan(m, v22, v572)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L149
	}
L145:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)))
	if v619 != 0 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)))
	if v620 != 0 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)+24))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v624 = F_bms_add_member(m, v622, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v626)+24)) = v624
	goto L144
L149:
	;
	m.G0 = v19 + int32(48)
	return v582
L150:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v640 = F_exprType(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v642 = F_format_type_be(m, v640)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v642
	F_errmsg_internal(m, int32(_a_F_build_subplan_1), v19+int32(16))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_build_subplan_2), int32(437), int32(_a_F_build_subplan_3))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_builtin_locale_encoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 != int32(67) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v90
L2:
	;
	v13 = int32(6)
	v14 = int32(_a_F_builtin_locale_encoding_0)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_locale_encoding[0])))
	if base.B2i32(v17 == int32(0))|base.B2i32(v17 != v20) != 0 {
		v38 = v17
		v39 = v20
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v11 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v90 = int32(-1)
	goto L1
L5:
	;
	if v38-v39 == int32(0) {
		v90 = v13
		goto L1
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v23 = l0
	v24 = v14
	goto L8
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v28
		v39 = v27
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v38 = v28
	v39 = v27
	goto L6
L10:
	;
	v31 = int32(1)
	if v28 == v27 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v43 = int32(_a_F_builtin_locale_encoding_1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_locale_encoding[1])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v67-v68 == int32(0) {
		v90 = v13
		goto L1
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v52 = l0
	v53 = v43
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v67 = v57
	v68 = v56
	goto L14
L18:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_builtin_locale_encoding_2), v6)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_builtin_locale_encoding_3), int32(1793), int32(_a_F_builtin_locale_encoding_4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_byteale(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v160 = int32(0)
	return base.I64_extend_i32_u(base.B2i32(v151 == v160)&base.B2i32(v45 <= v75) | base.B2i32(v151 < v160))
L60:
	;
	goto L59
}
