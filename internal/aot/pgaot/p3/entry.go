package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryFindChildPtr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v118 int32
	_ = v118
	v9 = l3 - int32(1)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v12) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v118 & int32(_a_F_entryFindChildPtr_0)
L2:
	;
	if v66&int32(_a_F_entryFindChildPtr_0) != 0 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v20 = int32(base.Ui32(v12+int32(_a_F_entryFindChildPtr_1)) >> (uint(int32(2)) % 32))
	goto L5
L4:
	;
	v20 = int32(0)
	goto L5
L5:
	;
	if base.Ui32(v20&int32(_a_F_entryFindChildPtr_0)) <= base.Ui32(v9&int32(_a_F_entryFindChildPtr_0)) {
		v66 = v20
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = l1 + int32(20)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+l3<<(uint(int32(2))%32))))
	v32 = l1 + v29&int32(_a_F_entryFindChildPtr_2)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+2)))
	if v33<<(uint(int32(16))%32)|v36 == l2 {
		v118 = l3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v42 = l3
	goto L8
L8:
	;
	v47 = v42 + int32(1)
	v48 = int32(_a_F_entryFindChildPtr_0)
	v49 = v47 & v48
	if base.Ui32(v20&v48) < base.Ui32(v49) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v118 = v47
	goto L1
L10:
	;
	v66 = v9
	goto L2
L11:
	;
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v25+v49<<(uint(int32(2))%32))))
	v59 = l1 + v56&int32(_a_F_entryFindChildPtr_2)
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+2)))
	if v60<<(uint(int32(16))%32)|v63 != l2 {
		v42 = v47
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	v81 = int32(1)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v118 = int32(0)
	goto L1
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v81&int32(_a_F_entryFindChildPtr_0)<<(uint(int32(2))%32))))
	v93 = l1 + v90&int32(_a_F_entryFindChildPtr_2)
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
	if v94<<(uint(int32(16))%32)|v97 == l2 {
		v118 = v81
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v101 = v81 + int32(1)
	v102 = int32(_a_F_entryFindChildPtr_0)
	if base.Ui32(v101&v102) <= base.Ui32(v66&v102) {
		v81 = v101
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
}
func F_entryIndexByFrequencyCmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(2)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(v6)%32))))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+656))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4+v11<<(uint(v6)%32))))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+656))
	return base.B2i32(base.Ui32(v16) < base.Ui32(v10)) - base.B2i32(base.Ui32(v10) < base.Ui32(v16))
}
func F_entryPrepareDownlink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	if l1 < int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_entryPrepareDownlink[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9+(l1^int32(-1))<<(uint(int32(2))%32))))
		v23 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_entryPrepareDownlink[1]))
		v23 = v17 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	if l1 < int32(0) {
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_entryPrepareDownlink[2]))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(l1^int32(-1))<<(uint(int32(6))%32))+16))
		v42 = v33
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_entryPrepareDownlink[3]))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v35+l1<<(uint(int32(6))%32)+int32(-64))+16))
		v42 = v41
	}
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v43) {
		v53 = int32(base.Ui32(v43+int32(_a_F_entryPrepareDownlink_0))>>(uint(int32(2))%32)) & int32(_a_F_entryPrepareDownlink_1)
	} else {
		v53 = int32(0)
	}
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v23+v53<<(uint(int32(2))%32))+20))
	v60 = v23 + v57&int32(_a_F_entryPrepareDownlink_2)
	v62 = F_palloc(m, int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		return int32(0)
	} else {
		v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)))
		v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v66)+6)))
		if v68&int32(2) == int32(0) {
			v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)))
			v98 = F_palloc(m, v95&int32(_a_F_entryPrepareDownlink_3))
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int32(0)
			} else {
				v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)))
				v102 = v100 & int32(_a_F_entryPrepareDownlink_3)
				if v102 == int32(0) {
					v106 = v98
				} else {
					base.MemoryCopy(m, v98, v60, v102)
					v106 = v98
				}
				v108 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)) = uint16(v108)
				*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)) = uint16(v42)
				v112 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v112)
				*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v108)
				*(*int32)(unsafe.Add(mBase, uint32(v62))) = v106
				return v62
			}
		} else {
			v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
			if v73 == int32(_a_F_entryPrepareDownlink_1) {
				v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)))
				v98 = F_palloc(m, v95&int32(_a_F_entryPrepareDownlink_3))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)))
					v102 = v100 & int32(_a_F_entryPrepareDownlink_3)
					if v102 == int32(0) {
						v106 = v98
					} else {
						base.MemoryCopy(m, v98, v60, v102)
						v106 = v98
					}
					v108 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)) = uint16(v108)
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)) = uint16(v42)
					v112 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v112)
					*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v108)
					*(*int32)(unsafe.Add(mBase, uint32(v62))) = v106
					return v62
				}
			} else {
				v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)))
				v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
				v86 = (v76 | v77<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
				v87 = F_palloc(m, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					if v86 != 0 {
						base.MemoryCopy(m, v87, v60, v86)
					} else {
					}
					v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
					v93 = v90&int32(_a_F_entryPrepareDownlink_4) | v86
					*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)) = uint16(v93)
					v106 = v87
					v108 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)) = uint16(v108)
					*(*uint16)(unsafe.Add(mBase, uint32(v106)+2)) = uint16(v42)
					v112 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v112)
					*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v108)
					*(*int32)(unsafe.Add(mBase, uint32(v62))) = v106
					return v62
				}
			}
		}
	}
}
func F_entry_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+256))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+256))
	if base.F64_lt(v7, v9) != 0 {
		v12 = int32(-1)
	} else {
		v12 = base.F64_gt(v7, v9)
	}
	return v12
}
