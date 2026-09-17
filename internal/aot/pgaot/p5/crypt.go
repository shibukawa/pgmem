package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__crypt_gensalt_extended_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	if base.B2i32(l2 < int32(3))|base.B2i32(l4 < int32(10)) == int32(0) {
		if l0 != 0 {
			if l0&int32(-16777215) != int32(1) {
				v93 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v93)
				return v93
			} else {
				v22 = l0
				v23 = int32(95)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v23)
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v22)>>(uint(int32(18))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v29)
				v31 = int32(63)
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v35)
				v37 = int32(12)
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v22)>>(uint(v37)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v43)
				v45 = int32(6)
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v22)>>(uint(v45)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v51)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
				v56 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v56)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v55)>>(uint(int32(2))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v62)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v68)
				v71 = v53 << (uint(int32(8)) % 32)
				v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v71|v55<<(uint(int32(16))%32))>>(uint(v37)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v81)
				v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v54|v71)>>(uint(v45)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v90)
				return l3
			}
		} else {
			v22 = int32(725)
			v23 = int32(95)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v23)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v22)>>(uint(int32(18))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v29)
			v31 = int32(63)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)) = uint8(v35)
			v37 = int32(12)
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v22)>>(uint(v37)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+3)) = uint8(v43)
			v45 = int32(6)
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v22)>>(uint(v45)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v51)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			v56 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v56)
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v55)>>(uint(int32(2))%32)))+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v62)
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v68)
			v71 = v53 << (uint(int32(8)) % 32)
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v71|v55<<(uint(int32(16))%32))>>(uint(v37)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)) = uint8(v81)
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v54|v71)>>(uint(v45)%32))&v31)+uint32(_c_F__crypt_gensalt_extended_rn[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v90)
			return l3
		}
	} else {
		if int32(0) < l4 {
			v93 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v93)
			return v93
		} else {
			return int32(0)
		}
	}
}
func F__crypt_gensalt_sha256_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l4 != 0 {
		base.MemoryFill(m, l3, int32(0), l4)
	} else {
	}
	v8 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v8)
	v10 = int32(_a_F__crypt_gensalt_sha256_rn_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v10)
	v12 = F__crypt_gensalt_sha(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	v58 = v51 + v55
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 == v63 {
		v51 = v58
		v52 = v60
		v53 = v56
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v68 = v58
	v69 = v60
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
