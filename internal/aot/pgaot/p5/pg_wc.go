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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[491]))
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
	return v107
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v107 = v2
		goto L1
	} else {
		goto L25
	}
L3:
	;
	if base.Ui32(l0) <= base.Ui32(int32(131071)) {
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
		v107 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[494]))))
	v12 = int32(1)
	return int32(base.Ui32(v11)>>(uint(v12)%32)) & v12
L7:
	;
	return v63
L8:
	;
	v25 = int32(0)
	v26 = int32(1178)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v51 = int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v51)%32))+uint32(_consts[496]))))
	v63 = v55 & v51
	goto L7
L11:
	;
	v31 = base.I32_div_s(v25+v26, int32(2))
	v33 = v31 << (uint(int32(3)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[497])))
	if base.Ui32(v36) < base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v63 = int32(0)
	goto L7
L13:
	;
	if v47 <= v48 {
		v25 = v47
		v26 = v48
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v47 = v31 + int32(1)
	v48 = v26
	goto L13
L15:
	;
	goto L16
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[498])))
	if base.Ui32(v42) <= base.Ui32(l0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(1)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v47 = v25
	v48 = v31 - int32(1)
	goto L13
L20:
	;
	goto L12
L21:
	;
	return v92
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_consts[502]))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v78<<(uint(int32(5))%32))+uint32(_consts[502]))))
	v92 = int32(base.Ui32(v84)>>(uint(l0&int32(7))%32)) & int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v92 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(196606)))
	goto L21
L25:
	;
	goto L26
L26:
	;
	v107 = base.B2i32(base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
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
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	v3 = *(*int32)(unsafe.Add(mBase, _consts[491]))
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
	v119 = F_iswspace(m, l0)
	mBase = m.M
	if v119 != 0 {
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
	return v113
L6:
	;
	v113 = int32(0)
	goto L5
L7:
	;
	v113 = v99
	goto L5
L8:
	;
	v65 = int32(0)
	if int32(1)<<(uint(v64)%32)&int32(294913) != 0 {
		v99 = v65
		goto L7
	} else {
		goto L23
	}
L9:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[499]))))
	v64 = v61
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
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[495]))))
	if v44<<(uint(v50)%32)&int32(294913) != 0 {
		goto L6
	} else {
		goto L21
	}
L13:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 * int32(12)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[500])))
	if base.Ui32(v30) < base.Ui32(l0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v64 = int32(0)
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[501])))
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
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[496]))))
	if v56&int32(32) == int32(0) {
		v99 = v44
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	v72 = int32(10)
	v73 = v65
	goto L24
L24:
	;
	v78 = base.I32_div_s(v72+v73, int32(2))
	v80 = v78 << (uint(int32(3)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[503])))
	if base.Ui32(v83) < base.Ui32(l0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v99 = int32(1)
	goto L7
L26:
	;
	if v94 <= v93 {
		v72 = v93
		v73 = v94
		goto L24
	} else {
		goto L31
	}
L27:
	;
	v93 = v72
	v94 = v78 + int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[504])))
	if base.Ui32(v89) <= base.Ui32(l0) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v93 = v78 - int32(1)
	v94 = v73
	goto L26
L31:
	;
	goto L25
L32:
	;
	return v123
L33:
	;
	v123 = int32(0)
	goto L35
L34:
	;
	v120 = F_iswprint(m, l0)
	mBase = m.M
	v123 = base.B2i32(v120 != int32(0))
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L39
L37:
	;
	v137 = int32(0)
	goto L38
L38:
	;
	return v137
L39:
	;
	v137 = base.B2i32(base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94))) != int32(0))
	goto L38
}
func F_pg_wc_isword(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v175 int32
	_ = v175
	if l0 == int32(95) {
		v175 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v175
L2:
	;
	v6 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[491]))
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
		v175 = v6
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)))
	v27 = (v23 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v175 = v6
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[494]))))
	return base.B2i32(v15&int32(3) != int32(0))
L8:
	;
	return v141
L9:
	;
	v141 = v134
	goto L8
L10:
	;
	v134 = base.B2i32(v123&int32(255) == int32(9))
	goto L9
L11:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[495]))))
	v123 = v116
	goto L10
L12:
	;
	v141 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L8
L13:
	;
	v37 = int32(1178)
	v38 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v92 = int32(1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v92)%32))+uint32(_consts[496]))))
	if v97&v92 != 0 {
		v134 = v92
		goto L9
	} else {
		goto L36
	}
L16:
	;
	v43 = base.I32_div_s(v37+v38, int32(2))
	v45 = v43 << (uint(int32(3)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[497])))
	if base.Ui32(v48) < base.Ui32(l0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v27 != 0 {
		goto L12
	} else {
		goto L26
	}
L18:
	;
	if v60 <= v59 {
		v37 = v59
		v38 = v60
		goto L16
	} else {
		goto L25
	}
L19:
	;
	v59 = v37
	v60 = v43 + int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[498])))
	if base.Ui32(v54) <= base.Ui32(l0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v141 = int32(1)
	goto L8
L23:
	;
	goto L24
L24:
	;
	v59 = v43 - int32(1)
	v60 = v38
	goto L18
L25:
	;
	goto L17
L26:
	;
	v66 = int32(3367)
	v67 = int32(0)
	goto L28
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_consts[499]))))
	v123 = v91
	goto L10
L28:
	;
	v72 = base.I32_div_s(v66+v67, int32(2))
	v74 = v72 * int32(12)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_consts[500])))
	if base.Ui32(v77) < base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v123 = int32(0)
	goto L10
L30:
	;
	if v88 <= v87 {
		v66 = v87
		v67 = v88
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v87 = v66
	v88 = v72 + int32(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_consts[501])))
	if base.Ui32(v83) <= base.Ui32(l0) {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v87 = v72 - int32(1)
	v88 = v67
	goto L30
L35:
	;
	goto L29
L36:
	;
	if v27 == int32(0) {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	goto L12
L38:
	;
	return v155
L39:
	;
	v152 = F_iswalpha(m, l0)
	mBase = m.M
	v155 = base.B2i32(v152 != int32(0))
	goto L41
L40:
	;
	v155 = int32(1)
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L43
L43:
	;
	v175 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
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
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	v4 = *(*int32)(unsafe.Add(mBase, _consts[491]))
	switch v4 - int32(1) {
	case 0:
		if base.Ui32(l0) <= base.Ui32(int32(127)) {
			v124 = l0<<(uint(int32(2))%32) + int32(1907524)
		} else {
			v25 = int32(0)
			if base.Ui32(l0) < base.Ui32(int32(1416)) {
				v112 = l0
				v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
				v119 = v117
			} else {
				if base.Ui32(l0) <= base.Ui32(int32(43967)) {
					if base.Ui32(l0) <= base.Ui32(int32(8580)) {
						if base.Ui32(l0-int32(4256)) <= base.Ui32(int32(95)) {
							v112 = l0 - int32(2840)
							v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
							v119 = v117
						} else {
							if base.Ui32(l0) < base.Ui32(int32(5024)) {
								v119 = v25
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(5117)) {
									v112 = l0 - int32(3512)
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
									v119 = v117
								} else {
									if base.Ui32(l0) < base.Ui32(int32(7296)) {
										v119 = v25
									} else {
										v112 = l0 - int32(5690)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									}
								}
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(9398)) {
							v119 = v25
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(11565)) {
								if base.Ui32(l0) <= base.Ui32(int32(9449)) {
									v112 = l0 - int32(6507)
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
									v119 = v117
								} else {
									if base.Ui32(l0) < base.Ui32(int32(11264)) {
										v119 = v25
									} else {
										v112 = l0 - int32(8321)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(42560)) {
									v119 = v25
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(42998)) {
										v112 = l0 - int32(39315)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									} else {
										if base.Ui32(l0) < base.Ui32(int32(43859)) {
											v119 = v25
										} else {
											v112 = l0 - int32(40175)
											v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
											v119 = v117
										}
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(64256)) {
						v119 = v25
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(68997)) {
							if base.Ui32(l0) <= base.Ui32(int32(65370)) {
								if base.Ui32(l0) <= base.Ui32(int32(64279)) {
									v112 = l0 - int32(60463)
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
									v119 = v117
								} else {
									if base.Ui32(l0) < base.Ui32(int32(65313)) {
										v119 = v25
									} else {
										v112 = l0 - int32(61496)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									}
								}
							} else {
								if base.Ui32(l0) < base.Ui32(int32(66560)) {
									v119 = v25
								} else {
									if base.Ui32(l0) <= base.Ui32(int32(67004)) {
										v112 = l0 - int32(62685)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									} else {
										if base.Ui32(l0) < base.Ui32(int32(68736)) {
											v119 = v25
										} else {
											v112 = l0 - int32(64416)
											v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
											v119 = v117
										}
									}
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(71840)) {
								v119 = v25
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(93823)) {
									if base.Ui32(l0) <= base.Ui32(int32(71903)) {
										v112 = l0 - int32(67258)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									} else {
										if base.Ui32(l0) < base.Ui32(int32(93760)) {
											v119 = v25
										} else {
											v112 = l0 - int32(89114)
											v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
											v119 = v117
										}
									}
								} else {
									if base.Ui32(int32(67)) < base.Ui32(l0-int32(125184)) {
										v119 = v25
									} else {
										v112 = l0 - int32(120474)
										v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112<<(uint(int32(1))%32))+uint32(_consts[492]))))
										v119 = v117
									}
								}
							}
						}
					}
				}
			}
			v124 = v119<<(uint(int32(2))%32) + int32(1907520)
		}
		v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
		if v125 != 0 {
			v126 = v125
		} else {
			v126 = l0
		}
		return v126
	case 1:
		v129 = *(*int32)(unsafe.Add(mBase, _consts[493]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v149 = F_casemap(m, l0, int32(0))
			mBase = m.M
			return v149
		} else {
			v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
			if v132&int32(1) == int32(0) {
				v149 = F_casemap(m, l0, int32(0))
				mBase = m.M
				return v149
			} else {
				if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v145 = l0 | int32(32)
				} else {
					v145 = l0
				}
				return v145
			}
		}
	case 2:
		v152 = *(*int32)(unsafe.Add(mBase, _consts[493]))
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			if base.Ui32(int32(255)) < base.Ui32(l0) {
				v180 = l0
			} else {
				if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
					v179 = l0 | int32(32)
				} else {
					v179 = l0
				}
				v180 = v179
			}
			return v180
		} else {
			v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+4)))
			if v155&int32(1) == int32(0) {
				if base.Ui32(int32(255)) < base.Ui32(l0) {
					v180 = l0
				} else {
					if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
						v179 = l0 | int32(32)
					} else {
						v179 = l0
					}
					v180 = v179
				}
				return v180
			} else {
				if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v168 = l0 | int32(32)
				} else {
					v168 = l0
				}
				return v168
			}
		}
	default:
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v180 = l0
			return v180
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
