package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__crypt_gensalt_extended_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
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
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	if l2 < int32(3) {
		if int32(0) < l4 {
			v83 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v83)
			return v83
		} else {
			return int32(0)
		}
	} else {
		if l4 < int32(10) {
			if int32(0) < l4 {
				v83 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v83)
				return v83
			} else {
				return int32(0)
			}
		} else {
			if l0 != 0 {
				if l0&int32(-16777215) != int32(1) {
					v83 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v83)
					return v83
				} else {
					v19 = l0
					v20 = int32(95)
					*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v20)
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(int32(18))%32)))+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v26)
					v28 = int32(63)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v28)+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v31)
					v33 = int32(12)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v33)%32))&v28)+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v38)
					v40 = int32(6)
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v40)%32))&v28)+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v45)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v50 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v50)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49)>>(uint(int32(2))%32)))+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v55)
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48&v28)+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v60)
					v65 = v47 << (uint(int32(8)) % 32)
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49<<(uint(int32(16))%32)|v65)>>(uint(v33)%32))&v28)+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v72)
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v48|v65)>>(uint(v40)%32))&v28)+uint32(_consts[1452]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v80)
					return l3
				}
			} else {
				v19 = int32(725)
				v20 = int32(95)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v20)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(int32(18))%32)))+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v26)
				v28 = int32(63)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v28)+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v31)
				v33 = int32(12)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v33)%32))&v28)+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v38)
				v40 = int32(6)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v40)%32))&v28)+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v45)
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v50 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v50)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49)>>(uint(int32(2))%32)))+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v55)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48&v28)+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v60)
				v65 = v47 << (uint(int32(8)) % 32)
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49<<(uint(int32(16))%32)|v65)>>(uint(v33)%32))&v28)+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v72)
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v48|v65)>>(uint(v40)%32))&v28)+uint32(_consts[1452]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v80)
				return l3
			}
		}
	}
}
func F__crypt_gensalt_sha256_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v8 = F__emscripten_memset_bulkmem(m, l3, base.I32_extend8_s(int32(0)), l4)
	mBase = m.M
	v9 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)) = uint8(v9)
	v11 = int32(13604)
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v11)
	v13 = F__crypt_gensalt_sha(m, l0, l1, l2, v8, l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		return v13
	}
}
func F_run_crypt_des(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
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
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	v5 = int32(0)
	v6 = F_px_crypt_des(m, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v146
L2:
	;
	return int32(0)
L3:
	;
	if v6 == int32(0) {
		v146 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v6&int32(3) == int32(0) {
		v35 = v6
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if base.Ui32(l3-int32(1)) < base.Ui32(v68) {
		v146 = v5
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v68 = v60 - v6
	goto L5
L7:
	;
	v39 = v35
	goto L16
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v19 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v68 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v24 = v6
	goto L12
L12:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v60 = v28
	goto L6
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v54 = v39
	goto L19
L18:
	;
	goto L17
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v60 = v54
	goto L6
L21:
	;
	goto L20
L22:
	;
	if (v6^l2)&int32(3) != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v146 = l2
	goto L1
L24:
	;
	goto L23
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v125)
	if v125&int32(255) == int32(0) {
		goto L24
	} else {
		goto L40
	}
L26:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v124 = v6
	v125 = v77
	v126 = l2
	goto L25
L27:
	;
	goto L28
L28:
	;
	if v6&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = v6
	v83 = l2
	goto L32
L30:
	;
	v95 = v6
	v97 = l2
	goto L31
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v102 = int32(-2139062144)
	if (int32(16843008)-v99|v99)&v102 != v102 {
		v124 = v95
		v125 = v99
		v126 = v97
		goto L25
	} else {
		goto L36
	}
L32:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v84)
	if v84 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L33:
	;
	v95 = v91
	v97 = v89
	goto L31
L34:
	;
	v88 = int32(1)
	v89 = v83 + v88
	v91 = v81 + v88
	if v91&int32(3) != 0 {
		v81 = v91
		v83 = v89
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v107 = v95
	v108 = v99
	v109 = v97
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v108
	v111 = int32(4)
	v112 = v109 + v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v115 = v107 + v111
	v119 = int32(-2139062144)
	if (v113|(int32(16843008)-v113))&v119 == v119 {
		v107 = v115
		v108 = v113
		v109 = v112
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v124 = v115
	v125 = v113
	v126 = v112
	goto L25
L39:
	;
	goto L38
L40:
	;
	v133 = v124
	v135 = v126
	goto L41
L41:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)) = uint8(v136)
	v138 = int32(1)
	if v136 != 0 {
		v133 = v133 + v138
		v135 = v135 + v138
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L24
L43:
	;
	goto L42
}
func F_run_crypt_md5(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_px_crypt_md5(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
