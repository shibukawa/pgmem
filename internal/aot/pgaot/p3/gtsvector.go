package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtsvector_decompress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v6 == v10 {
			return v4
		} else {
			v14 = F_palloc(m, int32(16))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v6
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v17
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v19
				v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+12)))
				v22 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v22)
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v21)
				return v14
			}
		}
	}
}
func F_gtsvector_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
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
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 == v2 {
		v38 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v38&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v25 == int32(0) {
		v38 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v28 != int32(7) {
		v38 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v31 != int32(17) {
		v38 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)))
	v38 = v34 ^ int32(1)
	goto L2
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = F_get_fn_opclass_options(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v47 = int32(124)
	goto L9
L9:
	;
	v49 = v47 + int32(8)
	v50 = F_palloc(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v47 = v46
	goto L9
L12:
	;
	v52 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v52
	v55 = v49 << (uint(v52) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v55
	v58 = v50 + int32(8)
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryFill(m, v58, int32(0), v47)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v61 <= int32(0) {
		v288 = v55
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v288) >> (uint(int32(2)) % 32))
	return v50
L17:
	;
	v66 = int32(3)
	v67 = v47 & v66
	v84 = v2
	goto L18
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v84<<(uint(int32(4))%32))))
	v94 = v92 + int32(8)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v95&int32(2) != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v288 = v55
	goto L16
L20:
	;
	v275 = v84 + int32(1)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v275 < v276 {
		v84 = v275
		goto L18
	} else {
		goto L41
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(25769803808)
	v288 = int32(32)
	goto L16
L22:
	;
	if v95&int32(4) != 0 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v207&int32(-16) == int32(32) {
		goto L20
	} else {
		goto L37
	}
L25:
	;
	if v47 <= int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v102 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v47) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = v102
	v113 = v102
	goto L30
L28:
	;
	v161 = v102
	goto L29
L29:
	;
	v178 = v161
	v182 = v102
	goto L34
L30:
	;
	v124 = v107 + v58
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v94))))
	v128 = v125 | v127
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v128)
	v131 = v107 | int32(1)
	v132 = v58 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v94))))
	v136 = v133 | v135
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v136)
	v139 = v107 | int32(2)
	v140 = v58 + v139
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+v94))))
	v144 = v141 | v143
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v144)
	v147 = v107 | int32(3)
	v148 = v58 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v94))))
	v152 = v149 | v151
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v152)
	v154 = int32(4)
	v155 = v107 + v154
	v157 = v113 + v154
	if v157 != v47&int32(2147483644) {
		v107 = v155
		v113 = v157
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v67 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v161 = v155
	goto L29
L34:
	;
	v195 = v178 + v58
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v94))))
	v199 = v196 | v198
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v199)
	v201 = int32(1)
	v204 = v182 + v201
	if v204 != v67 {
		v178 = v178 + v201
		v182 = v204
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L20
L36:
	;
	goto L35
L37:
	;
	v212 = int32(0)
	goto L38
L38:
	;
	v229 = int32(2)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v94+v212<<(uint(v229)%32))))
	v233 = base.I32_rem_u_s(v232, v47<<(uint(v66)%32))
	v236 = v58 + int32(base.Ui32(v233)>>(uint(int32(3))%32))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	v238 = int32(1)
	v242 = v237 | v238<<(uint(v233&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v242)
	v245 = v212 + v238
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if base.Ui32(v245) < base.Ui32(int32(base.Ui32(int32(base.Ui32(v246)>>(uint(v229)%32))-int32(8))>>(uint(v229)%32))) {
		v212 = v245
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L20
L40:
	;
	goto L39
L41:
	;
	goto L19
}
