package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistInitMetapage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	if l0&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l0, int32(0), int32(_a_F_SpGistInitMetapage_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_SpGistInitMetapage_1)
	v35 = int32(_a_F_SpGistInitMetapage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v35)
	v41 = int32(_a_F_SpGistInitMetapage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v45 = l0 + v44
	v46 = int32(_a_F_SpGistInitMetapage_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)) = uint16(v46)
	v48 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v48)
	v50 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v50
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1173640210)
	v70 = int32(92)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v70)
	v72 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v72
	return
}
func F_SpGistSetLastUsedPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v6 = F_spgGetCache(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if l1 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11+(l1^int32(-1))<<(uint(int32(2))%32))))
			v25 = v17
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[1]))
			v25 = v19 + l1<<(uint(int32(13))%32) + int32(-8192)
		}
		if l1 < int32(0) {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[2]))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v44 = v35
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistSetLastUsedPage[3]))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v44 = v43
		}
		if base.Ui32(v44) < base.Ui32(int32(3)) {
		} else {
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+16)))
			v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v47))))
			v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+14)))
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+12)))
			v52 = v50 - v51
			v53 = int32(0)
			if v53 < v52 {
				v56 = v52
			} else {
				v56 = v53
			}
			v59 = int32(4)
			v61 = int32(3)
			v63 = base.I32_rem_u_s(v44, v61)
			if v49&v59 != 0 {
				v66 = v61
			} else {
				v66 = v63
			}
			v72 = v6 + (int32(base.Ui32(v49)>>(uint(int32(1))%32))&v59|v66)<<(uint(int32(3))%32) - int32(-64)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			if v73 == int32(-1) {
				*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v72))) = v44
			} else {
				if v44 == v73 {
					*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v72))) = v44
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
					if v56 <= v77 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = v44
					}
				}
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
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
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v74 = F_copyObjectImpl(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
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
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L2
L7:
	;
	v63 = v16 + int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v63 < v64 {
		v16 = v63
		goto L5
	} else {
		goto L26
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = F_equal(m, v35, v25)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+v26)))
	v30 = v28
	goto L11
L10:
	;
	v30 = int32(0)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v30 == v31 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v31 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	return
L16:
	;
	if v36 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v43 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v46 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v56 = v43
	v57 = v40
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+v16<<(uint(int32(2))%32)))) = v57
	return
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v51 = v47 << (uint(int32(2)) % 32)
	goto L24
L23:
	;
	v51 = int32(0)
	goto L24
L24:
	;
	v52 = F_palloc0(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v56 = v52
	v57 = v55
	goto L21
L26:
	;
	goto L6
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v78 = F_lappend(m, v77, v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v81 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v110 != int32(2) {
		goto L1
	} else {
		goto L43
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104-int32(4)))) = v76
	goto L29
L31:
	;
	if v78 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v76 == int32(0) {
		goto L29
	} else {
		goto L38
	}
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v84 = v82
	goto L36
L35:
	;
	v84 = int32(0)
	goto L36
L36:
	;
	v86 = v84 << (uint(int32(2)) % 32)
	v87 = F_repalloc(m, v81, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
	v104 = v86 + v87
	goto L30
L38:
	;
	if v78 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v95 = v93
	goto L41
L40:
	;
	v95 = int32(0)
	goto L41
L41:
	;
	v97 = v95 << (uint(int32(2)) % 32)
	v98 = F_palloc0(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v98
	v104 = v97 + v98
	goto L30
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	goto L1
}
