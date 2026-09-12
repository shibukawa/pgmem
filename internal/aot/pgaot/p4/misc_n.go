package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NamedTuplestoreScanNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	F_tuplestore_select_read_pointer(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		v13 = F_tuplestore_gettupleslot(m, v10, int32(1), int32(0), v3)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v3
		}
	}
}
func F___netlink_enumerate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	v3 = l2
	v7 = m.G0
	v9 = v7 + int32(-8192)
	m.G0 = v9
	v14 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(int32(0)), int32(20))
	mBase = m.M
	goto L1
L1:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[1605]))) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[1606]))) = l1
	v18 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[1607]))) = uint16(v18)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[1608]))) = uint16(v3)
	v21 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_consts[922]))) = v21
	v26 = F_sendto(m, l0, v9, v21, v15, v15)
	mBase = m.M
	goto L3
L2:
	;
	m.G0 = v9 - int32(-8192)
	return v91
L3:
	;
	if v26 < int32(0) {
		v91 = v26
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = int32(0)
	v33 = F_recvfrom(m, l0, v9, int32(8192), int32(64), v31, v31)
	mBase = m.M
	goto L6
L5:
	;
	v91 = int32(-1)
	goto L2
L6:
	;
	if v33 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v41 = v33
	goto L8
L8:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v41) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v91 = int32(0)
	goto L2
L10:
	;
	goto L9
L11:
	;
	v46 = v9
	goto L14
L12:
	;
	goto L13
L13:
	;
	v73 = int32(0)
	v75 = F_recvfrom(m, l0, v9, int32(8192), int32(64), v73, v73)
	mBase = m.M
	goto L19
L14:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)))
	switch v52 - int32(2) {
	case 0:
		v91 = int32(-1)
		goto L2
	case 1:
		goto L10
	default:
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v55 = F_netlink_msg_to_ifaddr(m, l3, v46)
	mBase = m.M
	if v55 != 0 {
		v91 = v55
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v61 = v46 + (v56+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v9+v41-v61) {
		v46 = v61
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	if int32(0) < v75 {
		v41 = v75
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L5
}
func F_namele(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v60 <= int32(0))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v54 = F_strlen(m, v5)
	mBase = m.M
	v55 = F_strlen(m, v4)
	mBase = m.M
	v56 = F_varstr_cmp(m, v5, v54, v4, v55, v6)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v60 = v45 - v46
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v41 = v4
	v45 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	goto L5
L12:
	;
	v41 = v36
	v45 = v38
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 != v21 {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v36 = v30
	v38 = int32(0)
	goto L12
L15:
	;
	if v21 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v26 = v18 - int32(1)
	if v26 == int32(0) {
		v36 = v17
		v38 = v19
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v29 = int32(1)
	v30 = v17 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v31 != 0 {
		v16 = v16 + v29
		v17 = v30
		v18 = v26
		v19 = v31
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	return int32(0)
L20:
	;
	v60 = v56
	goto L1
}
func F_nepali_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
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
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v13 = F_find_among_b(m, l0, int32(4349872), int32(17))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v275
L2:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v76
	v82 = v75
	goto L22
L3:
	;
	return int32(0)
L4:
	;
	if v13 == int32(0) {
		v75 = int32(0)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
	v21 = int32(1)
	switch v13 - v21 {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		v75 = v21
		goto L2
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = int32(3)
	v31 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v33-v34 < v29 {
		v44 = v31
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v24 = F_slice_del(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if int32(0) <= v24 {
		v75 = v21
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v275 = v24
	goto L1
L10:
	;
	if v44 != 0 {
		v75 = v21
		goto L2
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = F_memcmp(m, v37+v33-v29, int32(2244082), v29)
	mBase = m.M
	if v40 != 0 {
		v44 = v31
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33 - v29
	v44 = int32(1)
	goto L11
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = v28 - v19
	v47 = v45 - v46
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47
	v49 = int32(3)
	v51 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v47-v54 < v49 {
		v64 = v51
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v64 != 0 {
		v75 = v21
		goto L2
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = F_memcmp(m, v57+v47-v49, int32(2244085), v49)
	mBase = m.M
	if v60 != 0 {
		v64 = v51
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47 - v49
	v64 = int32(1)
	goto L16
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65 - v46
	v68 = F_slice_del(m, l0)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if v68 < int32(0) {
		v275 = v68
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v75 = v21
	goto L2
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v86 = v85 - v83
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v83-int32(2) <= v87 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v270
	v275 = int32(1)
	goto L1
L24:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v254 = v253 - v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v259 = F_find_among_b(m, l0, int32(4350352), int32(91))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L67
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v83-int32(1)))))
	if v95&int32(224) != int32(128) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if int32(1)<<(uint(v95)%32)&int32(262) == int32(0) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v108 = F_find_among_b(m, l0, int32(4350224), int32(3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	if v108 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v115 = v114 - v86
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v115-int32(2) <= v118 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v115-int32(1)))))
	if v126&int32(224) != int32(128) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	if int32(1)<<(uint(v126)%32)&int32(262) == int32(0) {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v139 = F_find_among_b(m, l0, int32(4350288), int32(3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v139 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v143
	switch v139 - int32(1) {
	case 0:
		goto L36
	case 1:
		goto L35
	default:
		goto L24
	}
L35:
	;
	v229 = int32(9)
	v231 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v233-v234 < v229 {
		v244 = v231
		goto L61
	} else {
		goto L62
	}
L36:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v148 = int32(6)
	v150 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v152-v153 < v148 {
		v163 = v150
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v225 = F_slice_del(m, l0)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L58
	}
L38:
	;
	if v163 != 0 {
		goto L37
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = F_memcmp(m, v156+v152-v148, int32(2244266), v148)
	mBase = m.M
	if v159 != 0 {
		v163 = v150
		goto L39
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v152 - v148
	v163 = int32(1)
	goto L39
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v165 = v147 - v143
	v166 = v164 - v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v166
	v168 = int32(6)
	v170 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v166-v173 < v168 {
		v183 = v170
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v183 != 0 {
		goto L37
	} else {
		goto L47
	}
L44:
	;
	goto L43
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = F_memcmp(m, v176+v166-v168, int32(2244272), v168)
	mBase = m.M
	if v179 != 0 {
		v183 = v170
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v166 - v168
	v183 = int32(1)
	goto L44
L47:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v185 = v184 - v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v185
	v187 = int32(6)
	v189 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v185-v192 < v187 {
		v202 = v189
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v202 != 0 {
		goto L37
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = F_memcmp(m, v195+v185-v187, int32(2244278), v187)
	mBase = m.M
	if v198 != 0 {
		v202 = v189
		goto L49
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v185 - v187
	v202 = int32(1)
	goto L49
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v204 = v203 - v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v204
	v206 = int32(6)
	v208 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v204-v211 < v206 {
		v221 = v208
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v221 == int32(0) {
		goto L24
	} else {
		goto L57
	}
L54:
	;
	goto L53
L55:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v217 = F_memcmp(m, v214+v204-v206, int32(2244284), v206)
	mBase = m.M
	if v217 != 0 {
		v221 = v208
		goto L54
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v204 - v206
	v221 = int32(1)
	goto L54
L57:
	;
	goto L37
L58:
	;
	if int32(0) <= v225 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	v275 = v225
	goto L1
L60:
	;
	if v244 == int32(0) {
		goto L24
	} else {
		goto L64
	}
L61:
	;
	goto L60
L62:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = F_memcmp(m, v237+v233-v229, int32(2244290), v229)
	mBase = m.M
	if v240 != 0 {
		v244 = v231
		goto L61
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v233 - v229
	v244 = int32(1)
	goto L61
L64:
	;
	v247 = F_slice_del(m, l0)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	if v247 < int32(0) {
		v275 = v247
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L24
L67:
	;
	if v259 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v261
	v263 = F_slice_del(m, l0)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L23
L71:
	;
	if v263 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v267 = v263
	goto L74
L73:
	;
	v267 = v82
	goto L74
L74:
	;
	if int32(0) <= v263 {
		v82 = v267
		goto L22
	} else {
		goto L75
	}
L75:
	;
	v275 = v267
	goto L1
}
func F_netlink_msg_to_ifaddr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
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
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
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
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int64
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v14 == int32(16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v92 = int32(1)
	goto L23
L3:
	;
	v17 = int32(164)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v18&int32(-4) == int32(32) {
		v81 = v18
		v83 = v17
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0+v53&int32(63)<<(uint(int32(2))%32))+8))
	if v59 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v28 = l1 + int32(32)
	goto L7
L7:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+2)))
	if v40 == int32(7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v81 = v49
	v83 = v17
	goto L2
L9:
	;
	v81 = v28
	v83 = v39 + int32(160)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v45 = int32(3)
	v49 = v28 + (v39+v45)&int32(131068)
	if base.Ui32(v45) < base.Ui32(l1+v18-v49) {
		v28 = v49
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v64 = v59
	goto L14
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+140))
	if v53 == v75 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L1
L16:
	;
	v81 = v64
	v83 = int32(164)
	goto L2
L17:
	;
	goto L18
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	if v78 != 0 {
		v64 = v78
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	if v113 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	goto L20
L22:
	;
	v113 = F_emscripten_builtin_malloc(m, v112)
	mBase = m.M
	if v113 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L23:
	;
	v100 = base.I64_extend_i32_u(v92) * base.I64_extend_i32_u(v83)
	v101 = base.I32_wrap_i64(v100)
	if base.Ui32(v92|v83) < base.Ui32(int32(65536)) {
		v112 = v101
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(int64(32))%64))) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v109 = int32(-1)
	goto L27
L26:
	;
	v109 = v101
	goto L27
L27:
	;
	v112 = v109
	goto L22
L28:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113-int32(4)))))
	if v118&int32(3) == int32(0) {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v124 = F___memset(m, v113, int32(0), v112)
	mBase = m.M
	goto L21
L30:
	;
	return int32(-1)
L31:
	;
	goto L32
L32:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v129 == int32(16) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v592 != 0 {
		goto L167
	} else {
		goto L168
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+140)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v136&int32(-4) != int32(32) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v266&int32(-4) != int32(24) {
		goto L66
	} else {
		goto L67
	}
L37:
	;
	v142 = v113 + int32(144)
	v143 = int32(32)
	v144 = v113 + v143
	v148 = v113 + int32(104)
	v152 = v113 + int32(164)
	v157 = l1 + v143
	goto L40
L38:
	;
	goto L39
L39:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v250 == int32(0) {
		goto L33
	} else {
		goto L65
	}
L40:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+2)))
	switch v168 - int32(1) {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L46
	default:
		goto L42
	case 6:
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v229 = int32(3)
	v233 = v157 + (v228+v229)&int32(131068)
	if base.Ui32(v229) < base.Ui32(l1+v226-v233) {
		v157 = v233
		goto L40
	} else {
		goto L64
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = v152
	v218 = int32(4)
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v222 = v220 - v218
	if v222 != 0 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v199 = int32(4)
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v203 = v201 - v199
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v203) <= base.Ui32(int32(24)) {
		goto L57
	} else {
		goto L58
	}
L45:
	;
	v181 = int32(4)
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v185 = v183 - v181
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v185) <= base.Ui32(int32(24)) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v173 = v171 - int32(4)
	if base.Ui32(int32(16)) < base.Ui32(v173) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v173 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v142
	goto L42
L49:
	;
	v178 = F__emscripten_memcpy_bulkmem(m, v142, v157+int32(4), v173)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L42
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+11)) = uint8(v185)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+8)) = uint16(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v186
	v193 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v144))) = uint16(v193)
	v197 = F___memcpy(m, v113+int32(44), v157+v181, v185)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113+int32(12)))) = v144
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L42
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+11)) = uint8(v203)
	*(*uint16)(unsafe.Add(mBase, uint32(v148)+8)) = uint16(v205)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v204
	v211 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v211)
	v215 = F___memcpy(m, v113+int32(116), v157+v199, v203)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113+int32(20)))) = v148
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L42
L61:
	;
	v223 = F__emscripten_memcpy_bulkmem(m, v152, v157+v218, v222)
	mBase = m.M
	goto L63
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L41
L65:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v113)+140))
	v258 = l0 + v253&int32(63)<<(uint(int32(2))%32)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = v113
	goto L33
L66:
	;
	v272 = v113 + int32(20)
	v274 = v113 + int32(32)
	v276 = v113 + int32(104)
	v278 = v113 + int32(12)
	v280 = v113 + int32(144)
	v285 = l1 + int32(24)
	goto L69
L67:
	;
	goto L68
L68:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v504 == int32(0) {
		goto L33
	} else {
		goto L145
	}
L69:
	;
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285)+2)))
	switch v296 - int32(1) {
	case 0:
		goto L75
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L74
	default:
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v482 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	v483 = int32(3)
	v487 = v285 + (v482+v483)&int32(131068)
	if base.Ui32(v483) < base.Ui32(l1+v480-v487) {
		v285 = v487
		goto L69
	} else {
		goto L144
	}
L72:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	v468 = v466 - int32(4)
	if base.Ui32(int32(16)) < base.Ui32(v468) {
		goto L71
	} else {
		goto L139
	}
L73:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	if v416 != 0 {
		goto L118
	} else {
		goto L119
	}
L74:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v376 = int32(4)
	v377 = v285 + v376
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v375&int32(255) != int32(10) {
		goto L108
	} else {
		goto L109
	}
L75:
	;
	v299 = int32(4)
	v300 = v285 + v299
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	v303 = v301 - v299
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	if v306 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v305&int32(255) != int32(10) {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	goto L78
L78:
	;
	if v305&int32(255) != int32(10) {
		goto L95
	} else {
		goto L96
	}
L79:
	;
	goto L71
L80:
	;
	goto L79
L81:
	;
	if base.Ui32(v303) < base.Ui32(v333) {
		goto L80
	} else {
		goto L91
	}
L82:
	;
	if v305 != int32(2) {
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v319 = v113 + int32(112)
	v320 = int32(16)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	switch v321 - int32(254) {
	case 0:
		goto L88
	case 1:
		goto L87
	default:
		v333 = v320
		v334 = v319
		goto L81
	}
L85:
	;
	v333 = int32(4)
	v334 = v113 + int32(108)
	goto L81
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v304
	v333 = v320
	v334 = v319
	goto L81
L87:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v327&int32(15) != int32(2) {
		v333 = v320
		v334 = v319
		goto L81
	} else {
		goto L90
	}
L88:
	;
	v324 = int32(*(*int8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v324 < int32(-64) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v333 = v320
	v334 = v319
	goto L81
L90:
	;
	goto L86
L91:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v276))) = uint16(v305)
	v337 = F___memcpy(m, v334, v300, v333)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v276
	goto L80
L92:
	;
	goto L71
L93:
	;
	goto L92
L94:
	;
	if base.Ui32(v303) < base.Ui32(v367) {
		goto L93
	} else {
		goto L104
	}
L95:
	;
	if v305 != int32(2) {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v353 = v113 + int32(40)
	v354 = int32(16)
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	switch v355 - int32(254) {
	case 0:
		goto L101
	case 1:
		goto L100
	default:
		v367 = v354
		v368 = v353
		goto L94
	}
L98:
	;
	v367 = int32(4)
	v368 = v113 + int32(36)
	goto L94
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+24)) = v304
	v367 = v354
	v368 = v353
	goto L94
L100:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v361&int32(15) != int32(2) {
		v367 = v354
		v368 = v353
		goto L94
	} else {
		goto L103
	}
L101:
	;
	v358 = int32(*(*int8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v358 < int32(-64) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v367 = v354
	v368 = v353
	goto L94
L103:
	;
	goto L99
L104:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v274))) = uint16(v305)
	v371 = F___memcpy(m, v368, v300, v367)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = v274
	goto L93
L105:
	;
	goto L71
L106:
	;
	goto L105
L107:
	;
	if base.Ui32(v378-v376) < base.Ui32(v408) {
		goto L106
	} else {
		goto L117
	}
L108:
	;
	if v375 != int32(2) {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v394 = v113 + int32(112)
	v395 = int32(16)
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	switch v396 - int32(254) {
	case 0:
		goto L114
	case 1:
		goto L113
	default:
		v408 = v395
		v409 = v394
		goto L107
	}
L111:
	;
	v408 = int32(4)
	v409 = v113 + int32(108)
	goto L107
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v381
	v408 = v395
	v409 = v394
	goto L107
L113:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+1)))
	if v402&int32(15) != int32(2) {
		v408 = v395
		v409 = v394
		goto L107
	} else {
		goto L116
	}
L114:
	;
	v399 = int32(*(*int8)(unsafe.Add(mBase, uint32(v377)+1)))
	if v399 < int32(-64) {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v408 = v395
	v409 = v394
	goto L107
L116:
	;
	goto L112
L117:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v276))) = uint16(v375)
	v412 = F___memcpy(m, v409, v377, v408)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v276
	goto L106
L118:
	;
	goto L122
L119:
	;
	goto L120
L120:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v426 = int32(4)
	v427 = v285 + v426
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285))))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v425&int32(255) != int32(10) {
		goto L129
	} else {
		goto L130
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v418
	v424 = F__emscripten_memset_bulkmem(m, v274, base.I32_extend8_s(int32(0)), int32(36))
	mBase = m.M
	goto L125
L122:
	;
	v418 = F__emscripten_memcpy_bulkmem(m, v276, v274, int32(36))
	mBase = m.M
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L120
L126:
	;
	goto L71
L127:
	;
	goto L126
L128:
	;
	if base.Ui32(v428-v426) < base.Ui32(v458) {
		goto L127
	} else {
		goto L138
	}
L129:
	;
	if v425 != int32(2) {
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v444 = v113 + int32(40)
	v445 = int32(16)
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	switch v446 - int32(254) {
	case 0:
		goto L135
	case 1:
		goto L134
	default:
		v458 = v445
		v459 = v444
		goto L128
	}
L132:
	;
	v458 = int32(4)
	v459 = v113 + int32(36)
	goto L128
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+24)) = v431
	v458 = v445
	v459 = v444
	goto L128
L134:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	if v452&int32(15) != int32(2) {
		v458 = v445
		v459 = v444
		goto L128
	} else {
		goto L137
	}
L135:
	;
	v449 = int32(*(*int8)(unsafe.Add(mBase, uint32(v427)+1)))
	if v449 < int32(-64) {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v458 = v445
	v459 = v444
	goto L128
L137:
	;
	goto L133
L138:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v274))) = uint16(v425)
	v462 = F___memcpy(m, v459, v427, v458)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = v274
	goto L127
L139:
	;
	if v468 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v280
	goto L71
L141:
	;
	v473 = F__emscripten_memcpy_bulkmem(m, v280, v285+int32(4), v468)
	mBase = m.M
	goto L143
L142:
	;
	goto L143
L143:
	;
	goto L140
L144:
	;
	goto L70
L145:
	;
	v507 = int32(16)
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v511 = v113 + int32(68)
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	v513 = m.G0
	v515 = v513 - v507
	m.G0 = v515
	v517 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v515)+8)) = v517
	*(*int64)(unsafe.Add(mBase, uint32(v515))) = v517
	v522 = int32(128)
	if base.Ui32(v522) <= base.Ui32(v512) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v525 = v522
	goto L148
L147:
	;
	v525 = v512
	goto L148
L148:
	;
	v527 = int32(base.Ui32(v525) >> (uint(int32(3)) % 32))
	v529 = F__emscripten_memset_bulkmem(m, v515, base.I32_extend8_s(int32(255)), v527)
	mBase = m.M
	goto L149
L149:
	;
	if base.Ui32(v512) <= base.Ui32(int32(127)) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v538 = int32(255) << (uint(int32(8)-v525&int32(7)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v515+v527))) = uint8(v538)
	goto L152
L151:
	;
	goto L152
L152:
	;
	if v509&int32(255) != int32(10) {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	m.G0 = v515 + int32(16)
	goto L33
L154:
	;
	goto L153
L155:
	;
	if base.Ui32(int32(16)) < base.Ui32(v568) {
		goto L154
	} else {
		goto L165
	}
L156:
	;
	if v509 != int32(2) {
		goto L154
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v554 = v113 + int32(76)
	v555 = int32(16)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	switch v556 - int32(254) {
	case 0:
		goto L162
	case 1:
		goto L161
	default:
		v568 = v555
		v569 = v554
		goto L155
	}
L159:
	;
	v568 = int32(4)
	v569 = v113 + int32(72)
	goto L155
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v511)+24)) = int32(0)
	v568 = v555
	v569 = v554
	goto L155
L161:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	if v562&int32(15) != int32(2) {
		v568 = v555
		v569 = v554
		goto L155
	} else {
		goto L164
	}
L162:
	;
	v559 = int32(*(*int8)(unsafe.Add(mBase, uint32(v515)+1)))
	if v559 < int32(-64) {
		goto L160
	} else {
		goto L163
	}
L163:
	;
	v568 = v555
	v569 = v554
	goto L155
L164:
	;
	goto L160
L165:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v511))) = uint16(v509)
	v572 = F___memcpy(m, v569, v515, v568)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113+v507))) = v511
	goto L154
L166:
	;
	goto L1
L167:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v593 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	F_emscripten_builtin_free(m, v113)
	mBase = m.M
	goto L166
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113
	goto L172
L171:
	;
	goto L172
L172:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v597 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597))) = v113
	goto L175
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v113
	goto L166
}
func F_networkjoinsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 float64
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 float32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v97 int32
	_ = v97
	var v98 float64
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v135 float64
	_ = v135
	var v144 int32
	_ = v144
	var v145 float32
	_ = v145
	var v148 float32
	_ = v148
	var v151 float32
	_ = v151
	var v154 float32
	_ = v154
	var v156 float64
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v183 float64
	_ = v183
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v214 float64
	_ = v214
	var v221 int32
	_ = v221
	var v225 float32
	_ = v225
	var v227 float64
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 float64
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 float64
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v307 float64
	_ = v307
	var v310 float64
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 float32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 float64
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v378 float64
	_ = v378
	var v386 int32
	_ = v386
	var v387 float32
	_ = v387
	var v390 float32
	_ = v390
	var v393 float32
	_ = v393
	var v396 float32
	_ = v396
	var v398 float64
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v426 float64
	_ = v426
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v456 float64
	_ = v456
	var v465 float32
	_ = v465
	var v467 float64
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v495 float64
	_ = v495
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v530 int32
	_ = v530
	var v539 float64
	_ = v539
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v574 float64
	_ = v574
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 float32
	_ = v592
	var v594 float32
	_ = v594
	var v598 float64
	_ = v598
	var v600 int32
	_ = v600
	var v622 float64
	_ = v622
	var v631 int32
	_ = v631
	var v657 float64
	_ = v657
	var v663 float64
	_ = v663
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v702 float64
	_ = v702
	var v708 int32
	_ = v708
	var v710 float32
	_ = v710
	var v713 int32
	_ = v713
	var v714 float64
	_ = v714
	var v715 int32
	_ = v715
	var v717 float64
	_ = v717
	var v719 int32
	_ = v719
	var v744 float64
	_ = v744
	var v764 int32
	_ = v764
	var v771 float64
	_ = v771
	var v773 float64
	_ = v773
	var v785 float64
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v816 float64
	_ = v816
	var v822 int32
	_ = v822
	var v824 float32
	_ = v824
	var v827 int32
	_ = v827
	var v828 float64
	_ = v828
	var v829 int32
	_ = v829
	var v831 float64
	_ = v831
	var v833 int32
	_ = v833
	var v858 float64
	_ = v858
	var v885 float64
	_ = v885
	var v896 float64
	_ = v896
	var v903 float64
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v941 float64
	_ = v941
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v954 float64
	_ = v954
	var v955 int32
	_ = v955
	var v956 float64
	_ = v956
	var v957 int32
	_ = v957
	var v989 float64
	_ = v989
	var v993 int32
	_ = v993
	var v995 int64
	_ = v995
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1033 float64
	_ = v1033
	var v1038 float64
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1054 float64
	_ = v1054
	var v1055 float64
	_ = v1055
	var v1061 float64
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1098 float64
	_ = v1098
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 float64
	_ = v1114
	var v1122 float64
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	v2 = int32(0)
	v21 = float64(0)
	v29 = m.G0
	v31 = v29 - int32(272)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v37 - int32(931) {
	case 0:
		goto L5
	case 1:
		goto L6
	case 2:
		v61 = int32(-2)
		goto L1
	case 3:
		goto L3
	default:
		goto L4
	}
L1:
	;
	F_get_join_variables(m, v35, v34, v33, v31+int32(56), v31+int32(24), v31+int32(23))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L12
	}
L2:
	;
	v61 = int32(0)
	goto L1
L3:
	;
	v61 = int32(-1)
	goto L1
L4:
	;
	if v37 == int32(3552) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	v61 = int32(2)
	goto L1
L6:
	;
	v61 = int32(1)
	goto L1
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v37
	F_errmsg_internal(m, int32(9996), v31)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(517724), int32(870), int32(300443))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	switch v70 {
	case 0, 1, 2:
		goto L19
	default:
		goto L17
	case 4, 5:
		goto L18
	}
L13:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	if v1106 != 0 {
		goto L130
	} else {
		goto L131
	}
L14:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if v314 != 0 {
		goto L51
	} else {
		goto L52
	}
L15:
	;
	v266 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v266
	v268 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+232)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v31)+224)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v31)+216)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v31)+136)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v31)+144)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v31)+152)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v31)+160)) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v31)+208)) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v31)+128)) = v268
	v295 = v2
	v296 = v2
	v302 = v2
	v307 = v21
	v310 = v21
	goto L14
L16:
	;
	v256 = F_get_commutator(m, v37)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L48
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L45
	}
L18:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+23)))
	if v233 != 0 {
		goto L16
	} else {
		goto L43
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	if v71 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v77 = *(*float32)(unsafe.Add(mBase, uint32(v74+v75)+8))
	v83 = F_get_attstatsslot(m, v31+int32(208), v71, int32(1), int32(0), int32(3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v91 = F_get_attstatsslot(m, v31+int32(128), v87, int32(2), int32(0), int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v93 = int32(1024)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v31)+224))
	if v93 <= v94 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v97 = v93
	goto L25
L24:
	;
	v97 = v94
	goto L25
L25:
	;
	v98 = base.F64_promote_f32(v77)
	if v83 == int32(0) {
		v295 = v2
		v296 = v97
		v302 = v91
		v307 = v21
		v310 = v98
		goto L14
	} else {
		goto L26
	}
L26:
	;
	if v94 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v295 = int32(1)
	v296 = v97
	v302 = v91
	v307 = v21
	v310 = v98
	goto L14
L28:
	;
	goto L29
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v31)+228))
	v106 = v97 & int32(3)
	if v94 < int32(4) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v106 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v162 = int32(0)
	v183 = v21
	goto L30
L32:
	;
	goto L33
L33:
	;
	v114 = int32(0)
	v118 = v2
	v135 = v21
	goto L34
L34:
	;
	v144 = v104 + v114<<(uint(int32(2))%32)
	v145 = *(*float32)(unsafe.Add(mBase, uint32(v144)))
	v148 = *(*float32)(unsafe.Add(mBase, uint32(v144)+4))
	v151 = *(*float32)(unsafe.Add(mBase, uint32(v144)+8))
	v154 = *(*float32)(unsafe.Add(mBase, uint32(v144)+12))
	v156 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v135, base.F64_promote_f32(v145)), base.F64_promote_f32(v148)), base.F64_promote_f32(v151)), base.F64_promote_f32(v154))
	v157 = int32(4)
	v158 = v114 + v157
	v160 = v118 + v157
	if v160 != v97&int32(2044) {
		v114 = v158
		v118 = v160
		v135 = v156
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v162 = v158
	v183 = v156
	goto L30
L36:
	;
	goto L35
L37:
	;
	v295 = int32(1)
	v296 = v97
	v302 = v91
	v307 = v183
	v310 = v98
	goto L14
L38:
	;
	goto L39
L39:
	;
	v193 = v162
	v200 = int32(0)
	v214 = v183
	goto L40
L40:
	;
	v221 = int32(1)
	v225 = *(*float32)(unsafe.Add(mBase, uint32(v104+v193<<(uint(int32(2))%32))))
	v227 = base.F64_add(v214, base.F64_promote_f32(v225))
	v231 = v200 + v221
	if v231 != v106 {
		v193 = v193 + v221
		v200 = v231
		v214 = v227
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v295 = v221
	v296 = v97
	v302 = v91
	v307 = v227
	v310 = v98
	goto L14
L42:
	;
	goto L41
L43:
	;
	v238 = F_networkjoinsel_semi(m, v37, v61, v31+int32(56), v31+int32(24))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v1098 = v238
	goto L13
L45:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v244
	F_errmsg_internal(m, int32(506839), v31+int32(16))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(517724), int32(257), int32(321067))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v264 = F_networkjoinsel_semi(m, v256, int32(0)-v61, v31+int32(24), v31+int32(56))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	v1098 = v264
	goto L13
L50:
	;
	v1042 = int32(1)
	if (v295|v302)&v1042&((v1026|v1027)&v1042) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L51:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+22)))
	v318 = *(*float32)(unsafe.Add(mBase, uint32(v315+v316)+8))
	v324 = F_get_attstatsslot(m, v31+int32(168), v314, int32(1), int32(0), int32(3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v993 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+200)) = v993
	v995 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+192)) = v995
	*(*int64)(unsafe.Add(mBase, uint32(v31)+184)) = v995
	*(*int64)(unsafe.Add(mBase, uint32(v31)+176)) = v995
	*(*int64)(unsafe.Add(mBase, uint32(v31)+96)) = v995
	*(*int64)(unsafe.Add(mBase, uint32(v31)+104)) = v995
	*(*int64)(unsafe.Add(mBase, uint32(v31)+112)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v31)+120)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v31)+168)) = v995
	*(*int64)(unsafe.Add(mBase, uint32(v31)+88)) = v995
	v1026 = v2
	v1027 = v2
	v1033 = v21
	v1038 = v21
	goto L50
L54:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v332 = F_get_attstatsslot(m, v31+int32(88), v328, int32(2), int32(0), int32(1))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v334 = int32(1024)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v31)+184))
	if v334 <= v335 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v338 = v334
	goto L58
L57:
	;
	v338 = v335
	goto L58
L58:
	;
	v339 = base.F64_promote_f32(v318)
	if v324 == int32(0) {
		v657 = v21
		v663 = float64(0)
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v764&v302 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L60:
	;
	if v295&v332 == int32(0) {
		v764 = v324
		v771 = v663
		v773 = v657
		goto L59
	} else {
		goto L94
	}
L61:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v31)+188))
	if v335 <= int32(0) {
		v495 = v21
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v295 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	v347 = v338 & int32(3)
	if v335 < int32(4) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v347 == int32(0) {
		v495 = v426
		goto L62
	} else {
		goto L71
	}
L65:
	;
	v404 = int32(0)
	v426 = v21
	goto L64
L66:
	;
	goto L67
L67:
	;
	v354 = int32(0)
	v356 = v354
	v360 = v354
	v378 = v21
	goto L68
L68:
	;
	v386 = v343 + v356<<(uint(int32(2))%32)
	v387 = *(*float32)(unsafe.Add(mBase, uint32(v386)))
	v390 = *(*float32)(unsafe.Add(mBase, uint32(v386)+4))
	v393 = *(*float32)(unsafe.Add(mBase, uint32(v386)+8))
	v396 = *(*float32)(unsafe.Add(mBase, uint32(v386)+12))
	v398 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v378, base.F64_promote_f32(v387)), base.F64_promote_f32(v390)), base.F64_promote_f32(v393)), base.F64_promote_f32(v396))
	v399 = int32(4)
	v400 = v356 + v399
	v402 = v360 + v399
	if v402 != v338&int32(2044) {
		v356 = v400
		v360 = v402
		v378 = v398
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v404 = v400
	v426 = v398
	goto L64
L70:
	;
	goto L69
L71:
	;
	v434 = v404
	v437 = int32(0)
	v456 = v426
	goto L72
L72:
	;
	v465 = *(*float32)(unsafe.Add(mBase, uint32(v343+v434<<(uint(int32(2))%32))))
	v467 = base.F64_add(v456, base.F64_promote_f32(v465))
	v468 = int32(1)
	v471 = v437 + v468
	if v471 != v347 {
		v434 = v434 + v468
		v437 = v471
		v456 = v467
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v495 = v467
	goto L62
L74:
	;
	goto L73
L75:
	;
	v764 = int32(1)
	v771 = v21
	v773 = v495
	goto L59
L76:
	;
	goto L77
L77:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v31)+180))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v31)+228))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v31)+220))
	v507 = F_get_opcode(m, v37)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_fmgr_info(m, v507, v31+int32(244))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	if v296 <= int32(0) {
		v657 = v495
		v663 = float64(0)
		goto L60
	} else {
		goto L80
	}
L80:
	;
	v516 = int32(0)
	v530 = v516
	v539 = v21
	goto L81
L81:
	;
	if base.B2i32(v335 <= v516) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v657 = v495
	v663 = base.F64_add(v622, float64(0))
	goto L60
L83:
	;
	v550 = v530 << (uint(int32(2)) % 32)
	v554 = int32(0)
	v574 = v539
	goto L86
L84:
	;
	v622 = v539
	goto L85
L85:
	;
	v631 = v530 + int32(1)
	if v631 != v296 {
		v530 = v631
		v539 = v622
		goto L81
	} else {
		goto L93
	}
L86:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v550+v506)))
	v587 = v554 << (uint(int32(2)) % 32)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v504+v587)))
	v590 = F_FunctionCall2Coll(m, v31+int32(244), int32(0), v585, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L8
	} else {
		goto L88
	}
L87:
	;
	v622 = v598
	goto L85
L88:
	;
	if v590 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v592 = *(*float32)(unsafe.Add(mBase, uint32(v505+v550)))
	v594 = *(*float32)(unsafe.Add(mBase, uint32(v587+v343)))
	v598 = base.F64_add(v574, base.F64_promote_f32(base.F32_mul(v592, v594)))
	goto L91
L90:
	;
	v598 = v574
	goto L91
L91:
	;
	v600 = v554 + int32(1)
	if v600 != v338 {
		v554 = v600
		v574 = v598
		goto L86
	} else {
		goto L92
	}
L92:
	;
	goto L87
L93:
	;
	goto L82
L94:
	;
	if int32(0) < v296 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v672 = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v31)+100))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v31)+228))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v31)+220))
	v679 = v672
	v702 = v21
	goto L98
L96:
	;
	v744 = v21
	goto L97
L97:
	;
	v764 = v324
	v771 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v339), v657), v744), v663)
	v773 = v657
	goto L59
L98:
	;
	v708 = v679 << (uint(int32(2)) % 32)
	v710 = *(*float32)(unsafe.Add(mBase, uint32(v677+v708)))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v708+v678)))
	v714 = F_inet_hist_value_sel(m, v676, v675, v713, v672-v61)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L8
	} else {
		goto L100
	}
L99:
	;
	v744 = v717
	goto L97
L100:
	;
	v717 = base.F64_add(base.F64_mul(base.F64_promote_f32(v710), v714), v702)
	v719 = v679 + int32(1)
	if v719 != v296 {
		v679 = v719
		v702 = v717
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v785 = float64(0)
	if int32(0) < v335 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v885 = v771
	goto L104
L104:
	;
	if v332&v302 != int32(1) {
		v1026 = v764
		v1027 = v332
		v1033 = v885
		v1038 = v339
		goto L50
	} else {
		goto L112
	}
L105:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v31)+144))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v31)+188))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v31)+180))
	v793 = int32(0)
	v816 = v785
	goto L108
L106:
	;
	v858 = v785
	goto L107
L107:
	;
	v885 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v310), v307), v858), v771)
	goto L104
L108:
	;
	v822 = v793 << (uint(int32(2)) % 32)
	v824 = *(*float32)(unsafe.Add(mBase, uint32(v791+v822)))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v822+v792)))
	v828 = F_inet_hist_value_sel(m, v790, v789, v827, v61)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L8
	} else {
		goto L110
	}
L109:
	;
	v858 = v831
	goto L107
L110:
	;
	v831 = base.F64_add(base.F64_mul(base.F64_promote_f32(v824), v828), v816)
	v833 = v793 + int32(1)
	if v833 != v338 {
		v793 = v833
		v816 = v831
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v896 = float64(1)
	v903 = float64(0)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	if int32(3) <= v904 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v31)+100))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v31)+144))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v31)+140))
	v910 = int32(1)
	v920 = v910
	v922 = int32(0)
	v941 = v903
	goto L116
L114:
	;
	v989 = v903
	goto L115
L115:
	;
	v1026 = v764
	v1027 = int32(1)
	v1033 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_sub(v896, v310), v307), base.F64_sub(base.F64_sub(v896, v339), v773)), v989), v885)
	v1038 = v339
	goto L50
L116:
	;
	v949 = v922 + int32(1)
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v907+v920<<(uint(int32(2))%32))))
	v954 = F_inet_hist_value_sel(m, v909, v908, v953, v61)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L8
	} else {
		goto L118
	}
L117:
	;
	v989 = base.F64_div(v956, base.F64_convert_i32_s(v949))
	goto L115
L118:
	;
	v956 = base.F64_add(v941, v954)
	v957 = v920 + (int32(base.Ui32(v904-int32(3))>>(uint(int32(10))%32)) + v910)
	if v957 < v904-v910 {
		v920 = v957
		v922 = v949
		v941 = v956
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	if v37 == int32(3552) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v1061 = v1033
	goto L122
L122:
	;
	F_free_attstatsslot(m, v31+int32(208))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L8
	} else {
		goto L126
	}
L123:
	;
	v1054 = float64(0.01)
	goto L125
L124:
	;
	v1054 = float64(0.005)
	goto L125
L125:
	;
	v1055 = float64(1)
	v1061 = base.F64_mul(v1054, base.F64_mul(base.F64_sub(v1055, v310), base.F64_sub(v1055, v1038)))
	goto L122
L126:
	;
	F_free_attstatsslot(m, v31+int32(168))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	F_free_attstatsslot(m, v31+int32(128))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L8
	} else {
		goto L128
	}
L128:
	;
	F_free_attstatsslot(m, v31+int32(88))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	v1098 = v1061
	goto L13
L130:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
	m.T0[v1107].(func(*base.Module, int32))(m, v1106)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L8
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	if v1110 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	m.T0[v1111].(func(*base.Module, int32))(m, v1110)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L8
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v1114 = float64(0)
	if base.F64_lt(v1098, v1114) != 0 {
		v1122 = v1114
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L136
L138:
	;
	v1123 = F_Float8GetDatum(m, v1122)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L8
	} else {
		goto L141
	}
L139:
	;
	if base.F64_gt(v1098, float64(1)) == int32(0) {
		v1122 = v1098
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v1122 = float64(1)
	goto L138
L141:
	;
	m.G0 = v31 + int32(272)
	return v1123
}
func F_newLexeme(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = l4
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 < v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v28 = v10
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v30 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v29 + v30
		v35 = v28 + v29<<(uint(int32(3))%32)
		v36 = l2 - l1
		v39 = F_palloc(m, v36+v30)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = v39
			if v36 != 0 {
				v42 = F__emscripten_memcpy_bulkmem(m, v39, l1, v36)
				mBase = m.M
			} else {
			}
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			v46 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44+v36))) = uint8(v46)
			v49 = F_palloc(m, int32(16))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v54))) = l3
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v5)
				return
			}
		}
	} else {
		if v7 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(16)
			v16 = F_palloc(m, int32(128))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v26 = v16
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
				v28 = v26
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v30 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v29 + v30
				v35 = v28 + v29<<(uint(int32(3))%32)
				v36 = l2 - l1
				v39 = F_palloc(m, v36+v30)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = v39
					if v36 != 0 {
						v42 = F__emscripten_memcpy_bulkmem(m, v39, l1, v36)
						mBase = m.M
					} else {
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v44+v36))) = uint8(v46)
					v49 = F_palloc(m, int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = l3
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v5)
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7 << (uint(int32(1)) % 32)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = F_repalloc(m, v21, v7<<(uint(int32(4))%32))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = v24
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
				v28 = v26
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v30 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v29 + v30
				v35 = v28 + v29<<(uint(int32(3))%32)
				v36 = l2 - l1
				v39 = F_palloc(m, v36+v30)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v35))) = v39
					if v36 != 0 {
						v42 = F__emscripten_memcpy_bulkmem(m, v39, l1, v36)
						mBase = m.M
					} else {
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v44+v36))) = uint8(v46)
					v49 = F_palloc(m, int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = l3
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v5)
						return
					}
				}
			}
		}
	}
}
func F_newhicolorrow(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v6 < v7 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		v32 = v9
		v34 = v6
		v35 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34 + v35
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v42 = v32 + v38*v6<<(uint(v35)%32)
		v48 = v38 << (uint(v35) % 32)
		if v48 != 0 {
			v49 = F__emscripten_memcpy_bulkmem(m, v42, v32+l1*v38<<(uint(v35)%32), v48)
			mBase = m.M
			v50 = v49
		} else {
			v50 = v42
		}
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		if v51 <= int32(0) {
			v92 = v6
		} else {
			v57 = int32(0)
			for {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v61 = int32(1)
				v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50+v57<<(uint(v61)%32)))))
				v69 = v60 + v64*int32(24) + int32(4)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
				*(*int32)(unsafe.Add(mBase, uint32(v69))) = v70 + v61
				v75 = v57 + v61
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v75 < v76 {
					v57 = v75
					continue
				} else {
					break
				}
				break
			}
			v92 = v6
		}
		return v92
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v14 = base.I32_div_s(int32(2147483647), v11<<(uint(int32(1))%32))
		if v14 <= v7 {
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = int32(101)
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
			if v83 != 0 {
				v85 = v83
			} else {
				v85 = int32(12)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v85
			v92 = int32(0)
			return v92
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			v20 = F_repalloc_extended(m, v16, v7*v11<<(uint(int32(2))%32))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = int32(101)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
					if v83 != 0 {
						v85 = v83
					} else {
						v85 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v85
					v92 = int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v20
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v27 << (uint(int32(1)) % 32)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v32 = v20
					v34 = v31
					v35 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v34 + v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v42 = v32 + v38*v6<<(uint(v35)%32)
					v48 = v38 << (uint(v35) % 32)
					if v48 != 0 {
						v49 = F__emscripten_memcpy_bulkmem(m, v42, v32+l1*v38<<(uint(v35)%32), v48)
						mBase = m.M
						v50 = v49
					} else {
						v50 = v42
					}
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v51 <= int32(0) {
						v92 = v6
					} else {
						v57 = int32(0)
						for {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v61 = int32(1)
							v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50+v57<<(uint(v61)%32)))))
							v69 = v60 + v64*int32(24) + int32(4)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
							*(*int32)(unsafe.Add(mBase, uint32(v69))) = v70 + v61
							v75 = v57 + v61
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							if v75 < v76 {
								v57 = v75
								continue
							} else {
								break
							}
							break
						}
						v92 = v6
					}
				}
				return v92
			}
		}
	}
}
func F_nextval_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
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
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v148 int32
	_ = v148
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v169 int64
	_ = v169
	var v181 int64
	_ = v181
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v216 int64
	_ = v216
	var v226 int64
	_ = v226
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v253 int64
	_ = v253
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v263 int64
	_ = v263
	var v266 int64
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v374 int64
	_ = v374
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	v11 = int64(0)
	v25 = m.G0
	v27 = v25 - int32(112)
	m.G0 = v27
	F_init_sequence(m, l0, v27+int32(108), v27+int32(104))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L111
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L108
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L104
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v40 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v42 = F_pg_class_aclcheck(m, v38, v40, int64(260))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v42 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_PreventCommandIfReadOnly(m, int32(714093))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_PreventCommandIfParallelMode(m, int32(714093))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
	if v55 != v56 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	m.G0 = v27 + int32(112)
	return v374
L17:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v58 + v55
	F_sequence_close(m, v44, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v68 = F_SearchSysCache1(m, int32(61), l0)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[455])) = v54
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
	v374 = v66
	goto L16
L21:
	;
	if v68 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v74 = v72 + v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+48)))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v74)+40))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v74)+32))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v74)+24))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v74)+16))
	F_ReleaseCatCache(m, v68)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v86 = F_read_seq_tuple(m, v44, v27+int32(100), v27+int32(80))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v27)+100))
	if v88 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v86)+8))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+16)))
	v114 = base.I64_extend_i32_u(v109^int32(-1)) & int64(1)
	v115 = v76 - v114
	if base.B2i32(v115 <= v108)&(v109&int32(1)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92+(v88^int32(-1))<<(uint(int32(2))%32))))
	v106 = v98
	goto L25
L27:
	;
	goto L28
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v106 = v100 + v88<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	if v139 == int64(0) {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v124 = v115 + int64(32)
	v137 = int32(1)
	v139 = v124
	v140 = v124
	goto L29
L31:
	;
	goto L32
L32:
	;
	v125 = F_GetRedoRecPtr(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v127 = int64(32)
	v128 = v76 + v127
	v129 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v106)+4)))
	v130 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v106))))
	v134 = base.B2i32(base.Ui64(v129|v130<<(uint(v127)%64)) <= base.Ui64(v125))
	if base.Ui64(v129|v130<<(uint(v127)%64)) <= base.Ui64(v125) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v135 = v128
	goto L36
L35:
	;
	v135 = v108
	goto L36
L36:
	;
	if base.Ui64(v129|v130<<(uint(v127)%64)) <= base.Ui64(v125) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v136 = v128
	goto L39
L38:
	;
	v136 = v76
	goto L39
L39:
	;
	v137 = v134
	v139 = v136
	v140 = v135
	goto L29
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v258
	*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v79
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+12)) = uint8(v270)
	*(*int32)(unsafe.Add(mBase, _consts[455])) = v54
	if v137 != 0 {
		goto L78
	} else {
		goto L79
	}
L41:
	;
	v253 = v107
	v256 = v107
	v258 = v107
	v263 = v140
	v266 = v11
	goto L40
L42:
	;
	goto L43
L43:
	;
	v148 = v75 & int32(1)
	v160 = v107
	v161 = v114
	v162 = v107
	v163 = v139
	v164 = v107
	v169 = v140
	goto L44
L44:
	;
	if base.B2i32(v79 <= int64(0)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v253 = v226
	v256 = v236
	v258 = v237
	v263 = v238
	v266 = v11
	goto L40
L46:
	;
	if v161 < v76 {
		goto L70
	} else {
		goto L71
	}
L47:
	;
	if v148 == int32(0) {
		goto L3
	} else {
		goto L69
	}
L48:
	;
	v253 = v160
	v256 = v162
	v258 = v164
	v263 = v169
	v266 = v163
	goto L40
L49:
	;
	if base.B2i32(v78-v79 < v160)&base.B2i32(int64(0) <= v78) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.B2i32(v77 < int64(0))&base.B2i32(v160 < v77-v79) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	v181 = v160 + v79
	if int64(0) <= v78 {
		v226 = v181
		goto L46
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if int64(0) < v161 {
		goto L48
	} else {
		goto L57
	}
L55:
	;
	if v181 <= v78 {
		v226 = v181
		goto L46
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if v148 != 0 {
		v226 = v77
		goto L46
	} else {
		goto L58
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(402653314))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v195 + int32(4)
	F_errmsg(m, int32(707260), v27+int32(16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(523122), int32(750), int32(325981))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v216 = v160 + v79
	if v77 < int64(0) {
		v226 = v216
		goto L46
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v161 <= int64(0) {
		goto L47
	} else {
		goto L68
	}
L66:
	;
	if v77 <= v216 {
		v226 = v216
		goto L46
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	goto L48
L69:
	;
	v226 = v78
	goto L46
L70:
	;
	if v161 == int64(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v235 = v161
	v236 = v162
	v237 = v164
	v238 = v169
	goto L72
L72:
	;
	v240 = v163 - int64(1)
	if v240 != int64(0) {
		v160 = v226
		v161 = v235
		v162 = v236
		v163 = v240
		v164 = v237
		v169 = v238
		goto L44
	} else {
		goto L76
	}
L73:
	;
	v230 = v226
	goto L75
L74:
	;
	v230 = v164
	goto L75
L75:
	;
	v231 = int64(1)
	v235 = v161 + v231
	v236 = v226
	v237 = v230
	v238 = v169 - v231
	goto L72
L76:
	;
	goto L45
L77:
	;
	v344 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+16)) = uint8(v344)
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v263 - v266
	v348 = int32(4548900)
	v350 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v350 - v344
	F_UnlockReleaseBuffer(m, v88)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L102
	}
L78:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+118)))
	if v276 != int32(112) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v336 = int32(4548900)
	v338 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v338 + int32(1)
	F_MarkBufferDirty(m, v88)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L101
	}
L81:
	;
	v287 = int32(4548900)
	v289 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v289 + int32(1)
	F_MarkBufferDirty(m, v88)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L89
	}
L82:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v280 <= int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v283 != 0 {
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v285 = F_GetTopTransactionId(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v284 != 0 {
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	goto L81
L89:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+118)))
	if v296 != int32(112) {
		goto L77
	} else {
		goto L90
	}
L90:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v300 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v303 != 0 {
		goto L77
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v304 != 0 {
		goto L77
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	F_XLogRegisterBuffer(m, int32(0), v88, int32(6))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v311 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+16)) = uint8(v311)
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = int64(0)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v318
	F_XLogRegisterData(m, v27-int32(-64), int32(12))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	F_XLogRegisterData(m, v325, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v331 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = base.I64_rotr(v331, int64(32))
	goto L77
L101:
	;
	goto L77
L102:
	;
	F_sequence_close(m, v44, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v374 = v258
	goto L16
L104:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v395 + int32(4)
	F_errmsg(m, int32(206028), v27+int32(48))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(523122), int32(655), int32(325981))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
	F_errmsg_internal(m, int32(57640), v27)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(523122), int32(680), int32(325981))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(402653314))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v429 + int32(4)
	F_errmsg(m, int32(707315), v27+int32(32))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(523122), int32(769), int32(325981))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_nlevel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v10 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v9
			}
		} else {
			return v9
		}
	}
}
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = int32(1)
	v19 = l1 - v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = v20 + int32(23)
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+20)))
	if v23&v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(32)
	return v461
L2:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v461 = v459
	goto L1
L3:
	;
	v419 = v410 + v416
	v422 = l2 + v19<<(uint(int32(4))%32)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+26)))
	if v423 != int32(1) {
		v461 = v419
		goto L1
	} else {
		goto L83
	}
L4:
	;
	v280 = int32(0)
	v284 = v280
	v286 = int32(1)
	v287 = v280
	v288 = v23
	goto L44
L5:
	;
	v90 = int32(0)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
	v92 = v20 + v91
	v94 = l2 + int32(20)
	v97 = v94 + v19<<(uint(int32(4))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v90 <= v98 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v29 = v19 >> (uint(int32(3)) % 32)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v29))))
	v32 = int32(-1)
	if v31|v32<<(uint(v19&int32(7))%32) != v32 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
	v275 = v20 + v75
	goto L4
L8:
	;
	v39 = int32(0)
	if v29 <= v39 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v43 = v39
	goto L10
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v22))))
	if v56 != int32(255) {
		goto L7
	} else {
		goto L12
	}
L11:
	;
	goto L5
L12:
	;
	v60 = v43 + int32(1)
	if v29 != v60 {
		v43 = v60
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v101 = v98 + v92
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	if v102 != int32(1) {
		v461 = v101
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v23&int32(2) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+4)))
	switch v105&int32(65535) - int32(1) {
	case 0:
		goto L20
	case 1:
		goto L19
	default:
		goto L18
	case 3:
		v447 = v101
		goto L2
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101))))
	v461 = v111
	goto L1
L20:
	;
	v110 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101))))
	v461 = v110
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
	F_errmsg_internal(m, int32(504959), v16+int32(16))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(342394), int32(70), int32(73673))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(0)
	v172 = int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v173 < int32(2) {
		v199 = v172
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v19 < int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v136 = v90
	goto L28
L28:
	;
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94+v136<<(uint(int32(4))%32))+4)))
	if v151 <= int32(0) {
		v275 = v92
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	v155 = v136 + int32(1)
	if v155 <= v19 {
		v136 = v155
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v410 = v264
	v416 = v92
	goto L3
L33:
	;
	if v173 <= v199 {
		goto L32
	} else {
		goto L39
	}
L34:
	;
	v177 = v172
	goto L35
L35:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v94+v177<<(uint(int32(4))%32))))
	if v192 <= int32(0) {
		v199 = v177
		goto L33
	} else {
		goto L37
	}
L36:
	;
	goto L32
L37:
	;
	v196 = v177 + int32(1)
	if v196 != v173 {
		v177 = v196
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v216 = v199<<(uint(int32(4))%32) + v94 - int32(16)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216)+4)))
	v221 = v199
	v226 = v217 + v218
	goto L40
L40:
	;
	v235 = v94 + v221<<(uint(int32(4))%32)
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+4)))
	if v236 <= int32(0) {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	goto L32
L42:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+12)))
	v241 = int32(1)
	v245 = (v226 + v239 - v241) & (int32(0) - v239)
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v245
	v249 = v221 + v241
	if v249 != v173 {
		v221 = v249
		v226 = v236 + v245
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v288&int32(1) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v314 = l2 + int32(20) + v284<<(uint(int32(4))%32)
	if v286&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v284>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v303)>>(uint(v284&int32(7))%32))&int32(1) != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v284 = v284 + int32(1)
	v286 = int32(0)
	goto L44
L49:
	;
	if v284 == v19 {
		v410 = v358
		v416 = v275
		goto L3
	} else {
		goto L62
	}
L50:
	;
	v347 = int32(0)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v275))))
	if v349 != 0 {
		v358 = v287
		v359 = v347
		goto L49
	} else {
		goto L61
	}
L51:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+12)))
	v338 = (v287 + v332 - int32(1)) & (int32(0) - v332)
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314)+4)))
	if v339 != int32(65535) {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	v317 = int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if v318 < int32(0) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314)+4)))
	if v321 == int32(65535) {
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v358 = v318
	v359 = v317
	goto L49
L56:
	;
	v324 = int32(0)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+12)))
	v358 = (v287 + v325 - int32(1)) & (v324 - v325)
	v359 = v324
	goto L49
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v338
	v358 = v338
	v359 = v317
	goto L49
L58:
	;
	goto L59
L59:
	;
	if v338 != v287 {
		goto L50
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v287
	v358 = v287
	v359 = v317
	goto L49
L61:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+12)))
	v358 = (v287 + v350 - int32(1)) & (int32(0) - v350)
	v359 = v347
	goto L49
L62:
	;
	v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v314)+4)))
	if int32(0) < v361 {
		v396 = v361
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v402)+20)))
	v284 = v284 + int32(1)
	v286 = v359 & base.B2i32(int32(0) < v361)
	v287 = v396 + v358
	v288 = v403
	goto L44
L64:
	;
	v364 = v358 + v275
	if v361 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	if v367 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v393 = F_strlen(m, v364)
	mBase = m.M
	v396 = v393 + int32(1)
	goto L63
L68:
	;
	v370 = int32(6)
	v372 = int32(18)
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)))
	if v374 == v372 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if v367&int32(1) != 0 {
		goto L80
	} else {
		goto L81
	}
L71:
	;
	v377 = v372
	goto L73
L72:
	;
	v377 = int32(2)
	goto L73
L73:
	;
	if v374&int32(254) == int32(2) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v382 = v370
	goto L76
L75:
	;
	v382 = v377
	goto L76
L76:
	;
	if v374 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v385 = v370
	goto L79
L78:
	;
	v385 = v382
	goto L79
L79:
	;
	v396 = v385
	goto L63
L80:
	;
	v396 = int32(base.Ui32(v367) >> (uint(int32(1)) % 32))
	goto L63
L81:
	;
	goto L82
L82:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v396 = int32(base.Ui32(v390) >> (uint(int32(2)) % 32))
	goto L63
L83:
	;
	v426 = int32(*(*int16)(unsafe.Add(mBase, uint32(v422)+24)))
	switch v426&int32(65535) - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L85
	default:
		goto L84
	case 3:
		v447 = v419
		goto L2
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L21
	} else {
		goto L87
	}
L85:
	;
	v432 = int32(*(*int16)(unsafe.Add(mBase, uint32(v419))))
	v461 = v432
	goto L1
L86:
	;
	v431 = int32(*(*int8)(unsafe.Add(mBase, uint32(v419))))
	v461 = v431
	goto L1
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v426
	F_errmsg_internal(m, int32(504959), v16)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L21
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(342394), int32(70), int32(73673))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_now(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int64)(unsafe.Add(mBase, _consts[216]))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_numerictypmodout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc(m, int32(64))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(4) <= v8 {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(base.Ui32(v8-int32(4)) >> (uint(int32(16)) % 32))
			v21 = int32(21)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = (v8<<(uint(v21)%32) - int32(8388608)) >> (uint(v21) % 32)
			v30 = F_pg_snprintf(m, v10, int32(64), int32(708359), v6)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v10
			}
		} else {
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v32)
			m.G0 = v6 + int32(16)
			return v10
		}
	}
}
func F_numericvar_deserialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v7 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_pfree(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = F_palloc(m, v7<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v16
	v19 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16))) = uint16(v19)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v23 + int32(2)
	v28 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v28
	v32 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32
	v36 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v36
	if int32(0) < v7 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v19
	goto L14
L12:
	;
	goto L13
L13:
	;
	return
L14:
	;
	v47 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v50 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v43<<(uint(v50)%32)))) = uint16(v47)
	v55 = v43 + v50
	if v55 != v7 {
		v43 = v55
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
}
func F_numericvar_to_int64(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v194 int64
	_ = v194
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v218 int64
	_ = v218
	var v225 int64
	_ = v225
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v245 int64
	_ = v245
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v265 int64
	_ = v265
	var v268 int32
	_ = v268
	var v276 int64
	_ = v276
	var v280 int32
	_ = v280
	var v297 int32
	_ = v297
	var v323 int32
	_ = v323
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = F_palloc(m, v17<<(uint(int32(1))%32)+int32(2))
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
	v26 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22))) = uint16(v26)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v26 < v28 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = v28 << (uint(int32(1)) % 32)
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = v38 << (uint(int32(2)) % 32)
	if v40+int32(4) < int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v36 = F__emscripten_memcpy_bulkmem(m, v22+int32(2), v33, v35)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	m.G0 = v15 + int32(16)
	return v323
L11:
	;
	v323 = int32(1)
	goto L10
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	F_pfree(m, v22)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L63
	}
L13:
	;
	v46 = v22 + int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v51 = base.I32_div_s(v40+int32(7), int32(4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v52 <= v51 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if int32(0) < v105 {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v105 = v52
	v106 = v46
	v108 = v38
	goto L14
L16:
	;
	goto L17
L17:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46+v51<<(uint(int32(1))%32)))))
	if int32(5000) <= v57 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v51
	goto L21
L19:
	;
	v85 = v51
	goto L20
L20:
	;
	if int32(0) <= v85 {
		v105 = v51
		v106 = v46
		v108 = v38
		goto L14
	} else {
		goto L27
	}
L21:
	;
	v72 = int32(1)
	v74 = v22 + v60<<(uint(v72)%32)
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74))))
	v79 = base.B2i32(int32(9998) < v77)
	if int32(9998) < v77 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v85 = v84
	goto L20
L23:
	;
	v80 = int32(-9999)
	goto L25
L24:
	;
	v80 = v72
	goto L25
L25:
	;
	v81 = v80 + v77
	*(*uint16)(unsafe.Add(mBase, uint32(v74))) = uint16(v81)
	v84 = v60 - int32(1)
	if int32(9998) < v77 {
		v60 = v84
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v99 = int32(1)
	v105 = v51 + v99
	v106 = v22
	v108 = v38 + v99
	goto L14
L28:
	;
	v180 = int64(0) - base.I64_extend16_s(base.I64_extend_i32_u(v165))
	if int32(0) < v170 {
		goto L44
	} else {
		goto L45
	}
L29:
	;
	v119 = v105
	v120 = v106
	v122 = v108
	goto L32
L30:
	;
	goto L31
L31:
	;
	if v105 == int32(0) {
		goto L12
	} else {
		goto L42
	}
L32:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120))))
	if v129 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v144 = v119
	goto L38
L34:
	;
	v132 = int32(1)
	if v132 < v119 {
		v119 = v119 - v132
		v120 = v120 + int32(2)
		v122 = v122 - v132
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L12
L38:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120-int32(2)+v144<<(uint(int32(1))%32)))))
	if v157 != 0 {
		v165 = v129
		v167 = v144
		v168 = v120
		v170 = v122
		goto L28
	} else {
		goto L40
	}
L39:
	;
	goto L12
L40:
	;
	v158 = int32(1)
	if v158 < v144 {
		v144 = v144 - v158
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106))))
	v165 = v164
	v167 = v105
	v168 = v106
	v170 = v108
	goto L28
L43:
	;
	F_pfree(m, v22)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L62
	}
L44:
	;
	v184 = int32(1)
	v194 = v180
	goto L47
L45:
	;
	v265 = v180
	goto L46
L46:
	;
	F_pfree(m, v22)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L57
	}
L47:
	;
	v198 = int64(10000)
	v199 = int64(0)
	v204 = int64(32)
	v207 = int64(base.Ui64(v194) >> (uint(v204) % 64))
	v210 = int64(4294967295)
	v213 = v194 & v210
	v214 = v198 * v213
	v218 = int64(base.Ui64(v214)>>(uint(v204)%64)) + v198*v207
	v225 = v213*v199 + v218&v210
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v194*v199 + v194>>(uint(int64(63))%64)*v198 + v199*v207 + int64(base.Ui64(v218)>>(uint(v204)%64)) + int64(base.Ui64(v225)>>(uint(v204)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v214&v210 | v225<<(uint(v204)%64)
	goto L49
L48:
	;
	v265 = v251
	goto L46
L49:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	if v236 != v237>>(uint(int64(63))%64) {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	if v167 <= v184 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v253 = v184 + int32(1)
	if v253 <= v170 {
		v184 = v253
		v194 = v251
		goto L47
	} else {
		goto L56
	}
L52:
	;
	v251 = v237
	goto L51
L53:
	;
	goto L54
L54:
	;
	v245 = int64(*(*int16)(unsafe.Add(mBase, uint32(v168+v184<<(uint(int32(1))%32)))))
	v248 = v237 - v245
	if base.B2i32(int64(0) < v245) != base.B2i32(v248 < v237) {
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v251 = v248
	goto L51
L56:
	;
	goto L48
L57:
	;
	if v47 != int32(16384) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v265 == int64(-9223372036854775807-1) {
		v323 = int32(0)
		goto L10
	} else {
		goto L61
	}
L59:
	;
	v276 = v265
	goto L60
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v276
	goto L11
L61:
	;
	v276 = int64(0) - v265
	goto L60
L62:
	;
	v323 = int32(0)
	goto L10
L63:
	;
	goto L11
}
