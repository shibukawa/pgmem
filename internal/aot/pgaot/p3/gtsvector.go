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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
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
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
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
	v61 = F__emscripten_memset_bulkmem(m, v50+int32(8), base.I32_extend8_s(int32(0)), v47)
	mBase = m.M
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v62 <= int32(0) {
		v290 = v55
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(base.Ui32(v290) >> (uint(int32(2)) % 32))
	return v50
L15:
	;
	v67 = int32(3)
	v68 = v47 & v67
	v83 = v2
	goto L16
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(4)+v83<<(uint(int32(4))%32))))
	v95 = v93 + int32(8)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96&int32(2) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v290 = v55
	goto L14
L18:
	;
	v276 = v83 + int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v276 < v277 {
		v83 = v276
		goto L16
	} else {
		goto L39
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(25769803808)
	v290 = int32(32)
	goto L14
L20:
	;
	if v96&int32(4) != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v208&int32(-16) == int32(32) {
		goto L18
	} else {
		goto L35
	}
L23:
	;
	if v47 <= int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v103 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v47) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v108 = v103
	v113 = v103
	goto L28
L26:
	;
	v160 = v103
	goto L27
L27:
	;
	if v68 == int32(0) {
		goto L18
	} else {
		goto L31
	}
L28:
	;
	v125 = v108 + v61
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v95))))
	v129 = v126 | v128
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v129)
	v132 = v108 | int32(1)
	v133 = v61 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v95))))
	v137 = v134 | v136
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v137)
	v140 = v108 | int32(2)
	v141 = v61 + v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v95))))
	v145 = v142 | v144
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v145)
	v148 = v108 | int32(3)
	v149 = v61 + v148
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v95))))
	v153 = v150 | v152
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v153)
	v155 = int32(4)
	v156 = v108 + v155
	v158 = v113 + v155
	if v158 != v47&int32(2147483644) {
		v108 = v156
		v113 = v158
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v160 = v156
	goto L27
L30:
	;
	goto L29
L31:
	;
	v179 = v160
	v187 = v103
	goto L32
L32:
	;
	v196 = v179 + v61
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v95))))
	v200 = v197 | v199
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v200)
	v202 = int32(1)
	v205 = v187 + v202
	if v205 != v68 {
		v179 = v179 + v202
		v187 = v205
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L18
L34:
	;
	goto L33
L35:
	;
	v213 = int32(0)
	goto L36
L36:
	;
	v230 = int32(2)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v95+v213<<(uint(v230)%32))))
	v234 = base.I32_rem_u_s(v233, v47<<(uint(v67)%32))
	v237 = v61 + int32(base.Ui32(v234)>>(uint(int32(3))%32))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v239 = int32(1)
	v243 = v238 | v239<<(uint(v234&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v243)
	v246 = v213 + v239
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if base.Ui32(v246) < base.Ui32(int32(base.Ui32(int32(base.Ui32(v247)>>(uint(v230)%32))-int32(8))>>(uint(v230)%32))) {
		v213 = v246
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L18
L38:
	;
	goto L37
L39:
	;
	goto L17
}
