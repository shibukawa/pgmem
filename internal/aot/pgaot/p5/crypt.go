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
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(int32(18))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v26)
					v28 = int32(63)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v31)
					v33 = int32(12)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v33)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v38)
					v40 = int32(6)
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v40)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v45)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v50 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v50)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49)>>(uint(int32(2))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v55)
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v60)
					v65 = v47 << (uint(int32(8)) % 32)
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49<<(uint(int32(16))%32)|v65)>>(uint(v33)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v72)
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v48|v65)>>(uint(v40)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v80)
					return l3
				}
			} else {
				v19 = int32(725)
				v20 = int32(95)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v20)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(int32(18))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v26)
				v28 = int32(63)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v31)
				v33 = int32(12)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v33)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v38)
				v40 = int32(6)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v19)>>(uint(v40)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v45)
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v50 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v50)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49)>>(uint(int32(2))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v55)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v60)
				v65 = v47 << (uint(int32(8)) % 32)
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v49<<(uint(int32(16))%32)|v65)>>(uint(v33)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v72)
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v48|v65)>>(uint(v40)%32))&v28)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
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
	v11 = int32(_a_F__crypt_gensalt_sha256_rn_0)
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
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
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
	return v90
L2:
	;
	return int32(0)
L3:
	;
	if v6 == int32(0) {
		v90 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = F_strlen(m, v6)
	mBase = m.M
	if base.Ui32(l3-int32(1)) < base.Ui32(v12) {
		v90 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if (v6^l2)&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v90 = l2
	goto L1
L7:
	;
	goto L6
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v69)
	if v69&int32(255) == int32(0) {
		goto L7
	} else {
		goto L23
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v68 = v6
	v69 = v21
	v70 = l2
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v6&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v25 = v6
	v27 = l2
	goto L15
L13:
	;
	v39 = v6
	v41 = l2
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 != v46 {
		v68 = v39
		v69 = v43
		v70 = v41
		goto L8
	} else {
		goto L19
	}
L15:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v28)
	if v28 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	v39 = v35
	v41 = v33
	goto L14
L17:
	;
	v32 = int32(1)
	v33 = v27 + v32
	v35 = v25 + v32
	if v35&int32(3) != 0 {
		v25 = v35
		v27 = v33
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v51 = v39
	v52 = v43
	v53 = v41
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v52
	v55 = int32(4)
	v56 = v53 + v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v59 = v51 + v55
	v63 = int32(-2139062144)
	if (v57|(int32(16843008)-v57))&v63 == v63 {
		v51 = v59
		v52 = v57
		v53 = v56
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v68 = v59
	v69 = v57
	v70 = v56
	goto L8
L22:
	;
	goto L21
L23:
	;
	v77 = v68
	v79 = v70
	goto L24
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)) = uint8(v80)
	v82 = int32(1)
	if v80 != 0 {
		v77 = v77 + v82
		v79 = v79 + v82
		goto L24
	} else {
		goto L26
	}
L25:
	;
	goto L7
L26:
	;
	goto L25
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
