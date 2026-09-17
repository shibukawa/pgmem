package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_isalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v99 int32
	_ = v99
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isalpha[0]))
	switch v4 - int32(1) {
	case 0:
		goto L4
	case 1:
		goto L3
	case 2:
		goto L2
	default:
		goto L5
	}
L1:
	;
	return v99
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v99 = v2
		goto L1
	} else {
		goto L25
	}
L3:
	;
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_isalpha_0)) {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v99 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_wc_isalpha[1]))))
	v10 = int32(1)
	return int32(base.Ui32(v9)>>(uint(v10)%32)) & v10
L7:
	;
	return v59
L8:
	;
	v23 = int32(1178)
	v24 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v49 = int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v49)%32))+uint32(_c_F_pg_wc_isalpha[2]))))
	v59 = v51 & v49
	goto L7
L11:
	;
	v29 = base.I32_div_s(v23+v24, int32(2))
	v31 = v29 << (uint(int32(3)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_pg_wc_isalpha[3])))
	if base.Ui32(v34) < base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v59 = int32(0)
	goto L7
L13:
	;
	if v46 <= v45 {
		v23 = v45
		v24 = v46
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v45 = v23
	v46 = v29 + int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_pg_wc_isalpha[4])))
	if base.Ui32(v40) <= base.Ui32(l0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = int32(1)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v45 = v29 - int32(1)
	v46 = v24
	goto L13
L20:
	;
	goto L12
L21:
	;
	return v84
L22:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_pg_wc_isalpha[5]))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v72<<(uint(int32(5))%32))+uint32(_c_F_pg_wc_isalpha[5]))))
	v84 = int32(base.Ui32(v76)>>(uint(l0&int32(7))%32)) & int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v84 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_isalpha_1)))
	goto L21
L25:
	;
	goto L26
L26:
	;
	v99 = base.B2i32(base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L1
}
func F_pg_wc_isgraph(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v135 int32
	_ = v135
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isgraph[0]))
	switch v3 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L2
	case 2:
		goto L1
	default:
		goto L4
	}
L1:
	;
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	v116 = F_iswspace(m, l0)
	mBase = m.M
	if v116 != 0 {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94)))
L5:
	;
	return v111
L6:
	;
	v111 = int32(0)
	goto L5
L7:
	;
	v111 = v97
	goto L5
L8:
	;
	v63 = int32(0)
	if int32(1)<<(uint(v62)%32)&int32(_a_F_pg_wc_isgraph_0) != 0 {
		v97 = v63
		goto L7
	} else {
		goto L23
	}
L9:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_wc_isgraph[1]))))
	v62 = v59
	goto L8
L10:
	;
	v19 = int32(3367)
	v20 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v44 = int32(1)
	v47 = l0 << (uint(v44) % 32)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_pg_wc_isgraph[2]))))
	if v44<<(uint(v48)%32)&int32(_a_F_pg_wc_isgraph_0) != 0 {
		goto L6
	} else {
		goto L21
	}
L13:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 * int32(12)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_wc_isgraph[3])))
	if base.Ui32(v30) < base.Ui32(l0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v62 = int32(0)
	goto L8
L15:
	;
	if v41 <= v40 {
		v19 = v40
		v20 = v41
		goto L13
	} else {
		goto L20
	}
L16:
	;
	v40 = v19
	v41 = v25 + int32(1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_wc_isgraph[4])))
	if base.Ui32(v36) <= base.Ui32(l0) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v40 = v25 - int32(1)
	v41 = v20
	goto L15
L20:
	;
	goto L14
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_pg_wc_isgraph[5]))))
	if v54&int32(32) == int32(0) {
		v97 = v44
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	v70 = int32(10)
	v71 = v63
	goto L24
L24:
	;
	v76 = base.I32_div_s(v70+v71, int32(2))
	v78 = v76 << (uint(int32(3)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+uint32(_c_F_pg_wc_isgraph[6])))
	if base.Ui32(v81) < base.Ui32(l0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v97 = int32(1)
	goto L7
L26:
	;
	if v92 <= v91 {
		v70 = v91
		v71 = v92
		goto L24
	} else {
		goto L31
	}
L27:
	;
	v91 = v70
	v92 = v76 + int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)+uint32(_c_F_pg_wc_isgraph[7])))
	if base.Ui32(v87) <= base.Ui32(l0) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v91 = v76 - int32(1)
	v92 = v71
	goto L26
L31:
	;
	goto L25
L32:
	;
	return v121
L33:
	;
	v121 = int32(0)
	goto L35
L34:
	;
	v118 = F_iswprint(m, l0)
	mBase = m.M
	v121 = base.B2i32(v118 != int32(0))
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L39
L37:
	;
	v135 = int32(0)
	goto L38
L38:
	;
	return v135
L39:
	;
	v135 = base.B2i32(base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94))) != int32(0))
	goto L38
}
func F_pg_wc_isword(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v168 int32
	_ = v168
	if l0 == int32(95) {
		v168 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v168
L2:
	;
	v6 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isword[0]))
	switch v8 - int32(1) {
	case 0:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L3
	default:
		goto L6
	}
L3:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v168 = v6
		goto L1
	} else {
		goto L42
	}
L4:
	;
	if base.Ui32(int32(10)) <= base.Ui32(l0-int32(48)) {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isword[1]))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)))
	v25 = (v21 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v168 = v6
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_wc_isword[2]))))
	return base.B2i32(v13&int32(3) != int32(0))
L8:
	;
	return v135
L9:
	;
	v135 = v128
	goto L8
L10:
	;
	v128 = base.B2i32(v117&int32(255) == int32(9))
	goto L9
L11:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_pg_wc_isword[3]))))
	v117 = v110
	goto L10
L12:
	;
	v135 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L8
L13:
	;
	v35 = int32(1178)
	v36 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v90 = int32(1)
	v92 = l0 << (uint(v90) % 32)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_c_F_pg_wc_isword[4]))))
	if v93&v90 != 0 {
		v128 = v90
		goto L9
	} else {
		goto L36
	}
L16:
	;
	v41 = base.I32_div_s(v35+v36, int32(2))
	v43 = v41 << (uint(int32(3)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_pg_wc_isword[5])))
	if base.Ui32(v46) < base.Ui32(l0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v25 != 0 {
		goto L12
	} else {
		goto L26
	}
L18:
	;
	if v58 <= v57 {
		v35 = v57
		v36 = v58
		goto L16
	} else {
		goto L25
	}
L19:
	;
	v57 = v35
	v58 = v41 + int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_pg_wc_isword[6])))
	if base.Ui32(v52) <= base.Ui32(l0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v135 = int32(1)
	goto L8
L23:
	;
	goto L24
L24:
	;
	v57 = v41 - int32(1)
	v58 = v36
	goto L18
L25:
	;
	goto L17
L26:
	;
	v64 = int32(3367)
	v65 = int32(0)
	goto L28
L27:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_pg_wc_isword[7]))))
	v117 = v89
	goto L10
L28:
	;
	v70 = base.I32_div_s(v64+v65, int32(2))
	v72 = v70 * int32(12)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_pg_wc_isword[8])))
	if base.Ui32(v75) < base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v117 = int32(0)
	goto L10
L30:
	;
	if v86 <= v85 {
		v64 = v85
		v65 = v86
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v85 = v64
	v86 = v70 + int32(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_pg_wc_isword[9])))
	if base.Ui32(v81) <= base.Ui32(l0) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v85 = v70 - int32(1)
	v86 = v65
	goto L30
L35:
	;
	goto L29
L36:
	;
	if v25 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	goto L12
L38:
	;
	return v148
L39:
	;
	v144 = F_iswalpha(m, l0)
	mBase = m.M
	v148 = base.B2i32(v144 != int32(0))
	goto L41
L40:
	;
	v148 = int32(1)
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L43
L43:
	;
	v168 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L1
}
func F_pg_wc_tolower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_tolower[0]))
	switch v4 - int32(1) {
	case 0:
		if base.Ui32(l0) <= base.Ui32(int32(127)) {
			v162 = l0<<(uint(int32(2))%32) + int32(_a_F_pg_wc_tolower_0)
		} else {
			v25 = int32(0)
			if base.Ui32(l0) <= base.Ui32(int32(1415)) {
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[1]))))
				v157 = v30
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_1)) {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_2)) {
						if base.Ui32(l0-int32(_a_F_pg_wc_tolower_3)) <= base.Ui32(int32(95)) {
							v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[2]))))
							v157 = v43
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_4)) {
								v157 = v25
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_5)) {
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[3]))))
									v157 = v52
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_6)) {
										v157 = v25
									} else {
										v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[4]))))
										v157 = v59
									}
								}
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_7)) {
							v157 = v25
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_8)) {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_9)) {
									v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[5]))))
									v157 = v70
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_10)) {
										v157 = v25
									} else {
										v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[6]))))
										v157 = v77
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_11)) {
									v157 = v25
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_12)) {
										v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[7]))))
										v157 = v86
									} else {
										if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_13)) {
											v157 = v25
										} else {
											v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[8]))))
											v157 = v93
										}
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_14)) {
						v157 = v25
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_15)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_16)) {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_17)) {
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[9]))))
									v157 = v106
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_18)) {
										v157 = v25
									} else {
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[10]))))
										v157 = v113
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_19)) {
									v157 = v25
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_20)) {
										v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[11]))))
										v157 = v122
									} else {
										if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_21)) {
											v157 = v25
										} else {
											v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[12]))))
											v157 = v129
										}
									}
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_22)) {
								v157 = v25
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_23)) {
									if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_tolower_24)) {
										v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[13]))))
										v157 = v140
									} else {
										if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_tolower_25)) {
											v157 = v25
										} else {
											v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[14]))))
											v157 = v147
										}
									}
								} else {
									if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_pg_wc_tolower_26)) {
										v157 = v25
									} else {
										v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_tolower[15]))))
										v157 = v156
									}
								}
							}
						}
					}
				}
			}
			v162 = v157<<(uint(int32(2))%32) + int32(_a_F_pg_wc_tolower_27)
		}
		v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
		if v163 != 0 {
			v164 = v163
		} else {
			v164 = l0
		}
		return v164
	case 1:
		v167 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_tolower[16]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v187 = F_casemap(m, l0, int32(0))
			mBase = m.M
			return v187
		} else {
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			if v170&int32(1) == int32(0) {
				v187 = F_casemap(m, l0, int32(0))
				mBase = m.M
				return v187
			} else {
				if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v183 = l0 | int32(32)
				} else {
					v183 = l0
				}
				return v183
			}
		}
	case 2:
		v190 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_tolower[16]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			if base.Ui32(int32(255)) < base.Ui32(l0) {
				v218 = l0
			} else {
				if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
					v217 = l0 | int32(32)
				} else {
					v217 = l0
				}
				v218 = v217
			}
			return v218
		} else {
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
			if v193&int32(1) == int32(0) {
				if base.Ui32(int32(255)) < base.Ui32(l0) {
					v218 = l0
				} else {
					if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
						v217 = l0 | int32(32)
					} else {
						v217 = l0
					}
					v218 = v217
				}
				return v218
			} else {
				if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v206 = l0 | int32(32)
				} else {
					v206 = l0
				}
				return v206
			}
		}
	default:
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v218 = l0
			return v218
		} else {
			if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
				v17 = l0 | int32(32)
			} else {
				v17 = l0
			}
			return v17
		}
	}
}
