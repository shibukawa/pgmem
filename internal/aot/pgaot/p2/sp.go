package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistInitMetapage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v2 = int32(0)
	v3 = int32(_a_F_SpGistInitMetapage_0)
	if v2|(l0&int32(3)|int32(1)) == v2 {
		v21 = l0 + v3
		v23 = l0 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_SpGistInitMetapage_1)
	v44 = int32(_a_F_SpGistInitMetapage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = int32(_a_F_SpGistInitMetapage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	v55 = int32(_a_F_SpGistInitMetapage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v55)
	v57 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v54))) = uint16(v57)
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1173640210)
	v77 = int32(92)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v77)
	v79 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v79
	return
}
func F_SpGistSetLastUsedPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	v5 = F_spgGetCache(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l1 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[0]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10+(l1^int32(-1))<<(uint(int32(2))%32))))
			v24 = v16
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[1]))
			v24 = v18 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		if l1 < int32(0) {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[2]))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v43 = v34
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[3]))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v43 = v42
		}
		if base.Ui32(v43) < base.Ui32(int32(3)) {
		} else {
			v46 = int32(3)
			v48 = base.I32_rem_u_s(v43, v46)
			v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)))
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+v49))))
			if v51&int32(4) != 0 {
				v54 = v46
			} else {
				v54 = v48
			}
			v62 = v5 + v54<<(uint(int32(3))%32) + v51<<(uint(int32(2))%32)&int32(32)
			v64 = v62 - int32(-64)
			v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)))
			v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)))
			v67 = v65 - v66
			v68 = int32(0)
			if v68 < v67 {
				v71 = v67
			} else {
				v71 = v68
			}
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+64))
			if base.B2i32(v72 == int32(-1))|base.B2i32(v43 == v72) == int32(0) {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
				if v71 <= v79 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v71
					*(*int32)(unsafe.Add(mBase, uint32(v64))) = v43
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v71
				*(*int32)(unsafe.Add(mBase, uint32(v64))) = v43
			}
		}
		return
	}
}
func F_add_sp_item_to_pathtarget(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v80 = F_copyObjectImpl(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L27
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v16 = v3
	goto L5
L5:
	;
	v22 = v16 << (uint(int32(2)) % 32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22+v23)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v69 = v16 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v69 < v70 {
		v16 = v69
		goto L5
	} else {
		goto L26
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+v26)))
	v30 = v28
	goto L10
L9:
	;
	v30 = int32(0)
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = base.B2i32(v30 == v31) | base.B2i32(v30 == int32(0))
	goto L13
L12:
	;
	v37 = int32(1)
	goto L13
L13:
	;
	if v37 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v41 = F_equal(m, v40, v25)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	if v41 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v45 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v48 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v62 = v45
	v63 = v48
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63+v16<<(uint(int32(2))%32)))) = v62
	return
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v56 = v52 << (uint(int32(2)) % 32)
	goto L24
L23:
	;
	v56 = int32(0)
	goto L24
L24:
	;
	v57 = F_palloc0(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v62 = v60
	v63 = v57
	goto L21
L26:
	;
	goto L6
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v84 = F_lappend(m, v83, v80)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v114 != int32(2) {
		goto L1
	} else {
		goto L43
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106+v107-int32(4)))) = v82
	goto L29
L31:
	;
	if v84 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v82 == int32(0) {
		goto L29
	} else {
		goto L38
	}
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v90 = v88
	goto L36
L35:
	;
	v90 = int32(0)
	goto L36
L36:
	;
	v92 = v90 << (uint(int32(2)) % 32)
	v93 = F_repalloc(m, v87, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93
	v106 = v92
	v107 = v93
	goto L30
L38:
	;
	if v84 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v100 = v98
	goto L41
L40:
	;
	v100 = int32(0)
	goto L41
L41:
	;
	v102 = v100 << (uint(int32(2)) % 32)
	v103 = F_palloc0(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v103
	v106 = v102
	v107 = v103
	goto L30
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	goto L1
}
