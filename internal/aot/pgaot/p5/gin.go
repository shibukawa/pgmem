package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GinBufferShouldTrim(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 <= int32(0) {
		v48 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v48 <= v49 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = int32(6)
	v15 = v10 + v7*v11 - v11
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = (l1 + v16 + int32(25)) & int32(-2)
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+2)))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	v27 = int32(16)
	v29 = v25 | v26<<(uint(v27)%32)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v34 = v30 | v31<<(uint(v27)%32)
	if base.Ui32(v29) < base.Ui32(v34) {
		v45 = int32(-1)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v45 != 0 {
		v48 = v46
		goto L1
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	if base.Ui32(v34) < base.Ui32(v29) {
		v45 = int32(1)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	if base.Ui32(v39) < base.Ui32(v40) {
		v45 = int32(-1)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v45 = base.B2i32(base.Ui32(v40) < base.Ui32(v39))
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v46
	v48 = v46
	goto L1
L9:
	;
	if int32(1024) <= v105 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v105 = v49
	goto L9
L11:
	;
	goto L12
L12:
	;
	v56 = v49
	goto L13
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v62 = v59 + v56*int32(6)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v68 = (l1 + int32(24) + v63 + int32(1)) & int32(-2)
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
	v74 = int32(16)
	v76 = v72 | v73<<(uint(v74)%32)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+2)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68))))
	v81 = v77 | v78<<(uint(v74)%32)
	if base.Ui32(v76) < base.Ui32(v81) {
		v92 = int32(-1)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v105 = v97
	goto L9
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(0) < v92 {
		v105 = v93
		goto L9
	} else {
		goto L20
	}
L16:
	;
	goto L15
L17:
	;
	if base.Ui32(v81) < base.Ui32(v76) {
		v92 = int32(1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	if base.Ui32(v86) < base.Ui32(v87) {
		v92 = int32(-1)
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v92 = base.B2i32(base.Ui32(v87) < base.Ui32(v86))
	goto L16
L20:
	;
	v96 = int32(1)
	v97 = v93 + v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v97
	v100 = v56 + v96
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v100 < v101 {
		v56 = v100
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v117 = base.B2i32(v111 <= v112+v113)
	goto L24
L23:
	;
	v117 = int32(0)
	goto L24
L24:
	;
	return v117
}
func F_GinDataLeafPageGetItems(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v10 = l0 + int32(32)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v12 = l0 + v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	if v13&int32(128) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v17 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	if v17 == int64(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v104
	v108 = F_palloc(m, v104*int32(6))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L19
	}
L4:
	;
	v90 = v10
	v94 = v16 - int32(32)
	goto L6
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+38)))
	v27 = v10 + (v22+int32(1))&int32(_a_F_GinDataLeafPageGetItems_0)
	v29 = v27 + int32(8)
	v30 = l0 + v16
	if base.Ui32(v30) <= base.Ui32(v29) {
		v81 = v10
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v94 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v90 = v81
	v94 = v30 - v81
	goto L6
L8:
	;
	v32 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	v33 = int64(48)
	v36 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v37 = int64(32)
	v39 = v32<<(uint(v33)%64) | v17 | v36<<(uint(v37)%64)
	v40 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+10)))
	v44 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
	if base.Ui64(v39) < base.Ui64(v40|(v41<<(uint(v37)%64)|v44<<(uint(v33)%64))) {
		v81 = v10
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v53 = v27
	v55 = v29
	goto L10
L10:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+14)))
	v63 = v55 + (v58+int32(1))&int32(_a_F_GinDataLeafPageGetItems_0)
	v65 = v63 + int32(8)
	if base.Ui32(v30) <= base.Ui32(v65) {
		v81 = v55
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v81 = v55
	goto L7
L12:
	;
	v67 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)))
	v68 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+10)))
	v71 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v63)+8)))
	if base.Ui64(v67|(v68<<(uint(int64(32))%64)|v71<<(uint(int64(48))%64))) <= base.Ui64(v39) {
		v53 = v63
		v55 = v65
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v95 = F_ginPostingListDecodeAllSegments(m, v90, v94, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
	return v100
L17:
	;
	return int32(0)
L18:
	;
	return v95
L19:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v112 = v110 * int32(6)
	if v112 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v108, v10, v112)
	goto L22
L21:
	;
	goto L22
L22:
	;
	return v108
}
func F_GinInitMetabuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v91 int32
	_ = v91
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitMetabuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GinInitMetabuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v21 = int32(_a_F_GinInitMetabuffer_0)
	v23 = int32(0)
	if v23|(v20&int32(3)|int32(1)) == v23 {
		v39 = v20 + v21
		v41 = v20 + int32(4)
		if base.Ui32(v41) < base.Ui32(v39) {
			v43 = v39
		} else {
			v43 = v41
		}
		v48 = (v20^int32(-1)+v43)&int32(-4) + int32(4)
		if v48 == int32(0) {
		} else {
			base.MemoryFill(m, v20, int32(0), v48)
		}
	} else {
		base.MemoryFill(m, v20, int32(0), v21)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_GinInitMetabuffer_1)
	v62 = int32(_a_F_GinInitMetabuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v62)
	v68 = int32(_a_F_GinInitMetabuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v68)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v68)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v72 = v20 + v71
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(-1)
	v75 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v75)
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+64)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = int32(2)
	v91 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)) = uint16(v91)
	return
}
func F_GinInitPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
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
	v2 = l1
	v7 = int32(3)
	if l2&v7|(l0&v7|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l2))) == int32(0) {
		if l2 == int32(0) {
		} else {
			v21 = l0 + l2
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
		}
	} else {
		if l2 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), l2)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_GinInitPage_0)
	v44 = l2 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = l2 - int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v54 = l0 + v53
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)) = uint16(v2)
	return
}
func F_ginAllocEntryAccumulator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.Ui32(v5) < base.Ui32(int32(2048)) {
			v23 = v5
			v24 = v4
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23 + int32(1)
			return v24 + v23*int32(48)
		} else {
			v11 = F_palloc_mul(m, int32(48), int32(2048))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
				v16 = F_GetMemoryChunkSpace(m, v11)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 + v18
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v23 = int32(0)
					v24 = v21
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23 + int32(1)
					return v24 + v23*int32(48)
				}
			}
		}
	} else {
		v11 = F_palloc_mul(m, int32(48), int32(2048))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
			v16 = F_GetMemoryChunkSpace(m, v11)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 + v18
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v23 = int32(0)
				v24 = v21
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23 + int32(1)
				return v24 + v23*int32(48)
			}
		}
	}
}
func F_ginBeginBAScan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v10 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v4
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(base.B2i32(v14 == int32(_a_F_ginBeginBAScan_0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(836)
	m.G0 = v7 + v6
	return
}
func F_ginCombineData(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v6) <= base.Ui32(v7) {
		if v6 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ginCombineData_0), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_ginCombineData_1), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ginCombineData_2), int32(45), int32(_a_F_ginCombineData_3))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v12 = F_GetMemoryChunkSpace(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v14 - v12
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v17 << (uint(int32(1)) % 32)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v24 = F_repalloc_huge(m, v21, v17*int32(12))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v24
					v27 = F_GetMemoryChunkSpace(m, v24)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v27 + v29
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v35 != 0 {
						} else {
							v36 = int32(6)
							v38 = v34 + v33*v36
							v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(4)))))
							v42 = int64(32)
							v46 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-v36))))
							v47 = int64(48)
							v52 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(2)))))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							v55 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
							v58 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
							v62 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
							if base.Ui64(v41<<(uint(v42)%64)|v46<<(uint(v47)%64)|v52) <= base.Ui64(v55<<(uint(v42)%64)|v58<<(uint(v47)%64)|v62) {
							} else {
								v65 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v65)
							}
						}
						v70 = v34 + v33*int32(6)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v76 + int32(1)
						return
					}
				}
			}
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if v35 != 0 {
		} else {
			v36 = int32(6)
			v38 = v34 + v33*v36
			v41 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(4)))))
			v42 = int64(32)
			v46 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-v36))))
			v47 = int64(48)
			v52 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38-int32(2)))))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			v55 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+2)))
			v58 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54))))
			v62 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v54)+4)))
			if base.Ui64(v41<<(uint(v42)%64)|v46<<(uint(v47)%64)|v52) <= base.Ui64(v55<<(uint(v42)%64)|v58<<(uint(v47)%64)|v62) {
			} else {
				v65 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v65)
			}
		}
		v70 = v34 + v33*int32(6)
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v72)
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
		*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v76 + int32(1)
		return
	}
}
func F_ginEntryFillRoot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
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
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v3 = l2
	v5 = l4
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v8) {
		v18 = int32(base.Ui32(v8+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
	} else {
		v18 = int32(0)
	}
	v19 = int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3+v18<<(uint(v19)%32))+20))
	v25 = l3 + v22&int32(_a_F_ginEntryFillRoot_2)
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+16)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v26)+6)))
	if v28&v19 == int32(0) {
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
		v58 = F_palloc(m, v55&int32(_a_F_ginEntryFillRoot_3))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
			v62 = v60 & int32(_a_F_ginEntryFillRoot_3)
			if v62 != 0 {
				base.MemoryCopy(m, v58, v25, v62)
			} else {
			}
			v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
			v65 = v64
			v66 = v58
			v68 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v68)
			*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v3)
			v72 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v72)
			v78 = F_PageAddItemExtended(m, l1, v66, v65&int32(_a_F_ginEntryFillRoot_3), v68, v68)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				if v78 != 0 {
					F_pfree(m, v66)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
						if base.Ui32(int32(25)) <= base.Ui32(v82) {
							v92 = int32(base.Ui32(v82+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
						} else {
							v92 = int32(0)
						}
						v93 = int32(2)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l5+v92<<(uint(v93)%32))+20))
						v99 = l5 + v96&int32(_a_F_ginEntryFillRoot_2)
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v100)+6)))
						if v102&v93 == int32(0) {
							v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
							v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return
							} else {
								v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
								if v136 != 0 {
									base.MemoryCopy(m, v132, v99, v136)
								} else {
								}
								v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
								v140 = v132
								v141 = v138
								v142 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
								*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
								v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
								*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
								v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									if v152 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										F_pfree(m, v140)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
							if v107 == int32(_a_F_ginEntryFillRoot_1) {
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
									if v136 != 0 {
										base.MemoryCopy(m, v132, v99, v136)
									} else {
									}
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
									v140 = v132
									v141 = v138
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
								v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
								v120 = (v110 | v111<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
								v121 = F_palloc(m, v120)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									if v120 != 0 {
										base.MemoryCopy(m, v121, v99, v120)
									} else {
									}
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
									v127 = v124&int32(-8192) | v120
									*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v127)
									v140 = v121
									v141 = v127
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(731), int32(_a_F_ginEntryFillRoot_6))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
		if v33 == int32(_a_F_ginEntryFillRoot_1) {
			v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
			v58 = F_palloc(m, v55&int32(_a_F_ginEntryFillRoot_3))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
				v62 = v60 & int32(_a_F_ginEntryFillRoot_3)
				if v62 != 0 {
					base.MemoryCopy(m, v58, v25, v62)
				} else {
				}
				v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
				v65 = v64
				v66 = v58
				v68 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v68)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v3)
				v72 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v72)
				v78 = F_PageAddItemExtended(m, l1, v66, v65&int32(_a_F_ginEntryFillRoot_3), v68, v68)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if v78 != 0 {
						F_pfree(m, v66)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
							if base.Ui32(int32(25)) <= base.Ui32(v82) {
								v92 = int32(base.Ui32(v82+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
							} else {
								v92 = int32(0)
							}
							v93 = int32(2)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l5+v92<<(uint(v93)%32))+20))
							v99 = l5 + v96&int32(_a_F_ginEntryFillRoot_2)
							v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v100)+6)))
							if v102&v93 == int32(0) {
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
									if v136 != 0 {
										base.MemoryCopy(m, v132, v99, v136)
									} else {
									}
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
									v140 = v132
									v141 = v138
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
								if v107 == int32(_a_F_ginEntryFillRoot_1) {
									v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
										v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
										if v136 != 0 {
											base.MemoryCopy(m, v132, v99, v136)
										} else {
										}
										v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
										v140 = v132
										v141 = v138
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
									v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
									v120 = (v110 | v111<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
									v121 = F_palloc(m, v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										if v120 != 0 {
											base.MemoryCopy(m, v121, v99, v120)
										} else {
										}
										v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
										v127 = v124&int32(-8192) | v120
										*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v127)
										v140 = v121
										v141 = v127
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(731), int32(_a_F_ginEntryFillRoot_6))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+2)))
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
			v46 = (v36 | v37<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
			v47 = F_palloc(m, v46)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				if v46 != 0 {
					base.MemoryCopy(m, v47, v25, v46)
				} else {
				}
				v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)))
				v53 = v50&int32(-8192) | v46
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)) = uint16(v53)
				v65 = v53
				v66 = v47
				v68 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v68)
				*(*uint16)(unsafe.Add(mBase, uint32(v66)+2)) = uint16(v3)
				v72 = int32(base.Ui32(v3) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v72)
				v78 = F_PageAddItemExtended(m, l1, v66, v65&int32(_a_F_ginEntryFillRoot_3), v68, v68)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if v78 != 0 {
						F_pfree(m, v66)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+12)))
							if base.Ui32(int32(25)) <= base.Ui32(v82) {
								v92 = int32(base.Ui32(v82+int32(_a_F_ginEntryFillRoot_0))>>(uint(int32(2))%32)) & int32(_a_F_ginEntryFillRoot_1)
							} else {
								v92 = int32(0)
							}
							v93 = int32(2)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l5+v92<<(uint(v93)%32))+20))
							v99 = l5 + v96&int32(_a_F_ginEntryFillRoot_2)
							v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+16)))
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v100)+6)))
							if v102&v93 == int32(0) {
								v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
								v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
									if v136 != 0 {
										base.MemoryCopy(m, v132, v99, v136)
									} else {
									}
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
									v140 = v132
									v141 = v138
									v142 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
									*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
									v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
									*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
									v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										if v152 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											F_pfree(m, v140)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
								if v107 == int32(_a_F_ginEntryFillRoot_1) {
									v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
									v132 = F_palloc(m, v129&int32(_a_F_ginEntryFillRoot_3))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
										v136 = v134 & int32(_a_F_ginEntryFillRoot_3)
										if v136 != 0 {
											base.MemoryCopy(m, v132, v99, v136)
										} else {
										}
										v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+6)))
										v140 = v132
										v141 = v138
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								} else {
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
									v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
									v120 = (v110 | v111<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
									v121 = F_palloc(m, v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										if v120 != 0 {
											base.MemoryCopy(m, v121, v99, v120)
										} else {
										}
										v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
										v127 = v124&int32(-8192) | v120
										*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v127)
										v140 = v121
										v141 = v127
										v142 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v142)
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
										v146 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v146)
										v152 = F_PageAddItemExtended(m, l1, v140, v141&int32(_a_F_ginEntryFillRoot_3), v142, v142)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											if v152 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(736), int32(_a_F_ginEntryFillRoot_6))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												F_pfree(m, v140)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ginEntryFillRoot_4), int32(0))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ginEntryFillRoot_5), int32(731), int32(_a_F_ginEntryFillRoot_6))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ginEntryInsert(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	v2 = l1
	v4 = l3
	v8 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v8)
	v22 = v17 + int32(8)
	base.MemoryFill(m, v22, v8, int32(72))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(50)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(51)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(53)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+64)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = l2
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+54)) = uint16(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+52)) = uint16(v8)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+36)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = int32(54)
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+61)) = uint8(base.B2i32(l6 != v57))
	v62 = F_ginFindLeafPage(m, v22, v57, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		return
	} else {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
		if v64 < int32(0) {
			v68 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[0]))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v68+(v64^int32(-1))<<(uint(int32(2))%32))))
			v82 = v74
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[1]))
			v82 = v76 + v64<<(uint(int32(13))%32) + int32(-8192)
		}
		v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
		v86 = m.T0[v85].(func(*base.Module, int32, int32) int32)(m, v17+int32(8), v62)
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return
		} else {
			if v86 != 0 {
				v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+8)))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v82+v88<<(uint(int32(2))%32))+20))
				v95 = v82 + v92&int32(_a_F_ginEntryInsert_0)
				v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+4)))
				if v96 == int32(_a_F_ginEntryInsert_1) {
					v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+2)))
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95))))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					F_UnlockBuffer(m, v101)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						F_freeGinBtreeStack(m, v62)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_ginInsertItemPointers(m, v106, v99|v100<<(uint(int32(16))%32), l4, l5, l6)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return
							} else {
								m.G0 = v17 + int32(96)
								return
							}
						}
					}
				} else {
					v112 = int32(0)
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					if v115 < v112 {
						v119 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[2]))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+(v115^int32(-1))*int32(56))+16))
						v134 = v125
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[3]))
						v128 = int32(56)
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v127+v115*v128-v128)+16))
						v134 = v133
					}
					F_CheckForSerializableConflictIn(m, v113, v112, v134)
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
						v138 = F_gintuple_get_attrnum(m, l0, v95)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							v142 = F_gintuple_get_key(m, l0, v95, v17+int32(91))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								v146 = F_ginReadTuple(m, v95, v17+int32(92))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return
								} else {
									v148 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
									v151 = F_ginMergeItemPointers(m, l4, l5, v146, v148, v17+int32(84))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
										v157 = F_ginCompressPostingList(m, v151, v153, int32(2712), v17+int32(80))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
											v160 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
											if v159 == v160 {
												v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+91)))
												v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+6)))
												v171 = F_GinFormTuple(m, l0, v138, v142, v162, v157, (v163+int32(1))&int32(_a_F_ginEntryInsert_2)+int32(8), v159, int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return
												} else {
													v173 = v171
													F_pfree(m, v151)
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return
													} else {
														F_pfree(m, v157)
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return
														} else {
															if v173 == int32(0) {
																v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v181 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
																v182 = F_createPostingTree(m, v180, v146, v181, l6, v137)
																mBase = m.M
																v183 = m.ExcPending
																if v183 != 0 {
																	return
																} else {
																	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	F_ginInsertItemPointers(m, v184, v182, l4, l5, l6)
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return
																	} else {
																		v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+91)))
																		v188 = int32(0)
																		v192 = F_GinFormTuple(m, l0, v138, v142, v187, v188, v188, v188, int32(1))
																		mBase = m.M
																		v193 = m.ExcPending
																		if v193 != 0 {
																			return
																		} else {
																			*(*uint16)(unsafe.Add(mBase, uint32(v192)+2)) = uint16(v182)
																			v196 = int32(base.Ui32(v182) >> (uint(int32(16)) % 32))
																			*(*uint16)(unsafe.Add(mBase, uint32(v192))) = uint16(v196)
																			v198 = int32(_a_F_ginEntryInsert_1)
																			*(*uint16)(unsafe.Add(mBase, uint32(v192)+4)) = uint16(v198)
																			v201 = v192
																			F_pfree(m, v146)
																			mBase = m.M
																			v203 = m.ExcPending
																			if v203 != 0 {
																				return
																			} else {
																				v204 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v204)
																				v280 = v201
																				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
																				F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
																				mBase = m.M
																				v289 = m.ExcPending
																				if v289 != 0 {
																					return
																				} else {
																					F_pfree(m, v280)
																					mBase = m.M
																					v291 = m.ExcPending
																					if v291 != 0 {
																						return
																					} else {
																						m.G0 = v17 + int32(96)
																						return
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v201 = v173
																F_pfree(m, v146)
																mBase = m.M
																v203 = m.ExcPending
																if v203 != 0 {
																	return
																} else {
																	v204 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v204)
																	v280 = v201
																	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
																	F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
																	mBase = m.M
																	v289 = m.ExcPending
																	if v289 != 0 {
																		return
																	} else {
																		F_pfree(m, v280)
																		mBase = m.M
																		v291 = m.ExcPending
																		if v291 != 0 {
																			return
																		} else {
																			m.G0 = v17 + int32(96)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v173 = v112
												F_pfree(m, v151)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													F_pfree(m, v157)
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return
													} else {
														if v173 == int32(0) {
															v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															v181 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
															v182 = F_createPostingTree(m, v180, v146, v181, l6, v137)
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
																return
															} else {
																v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_ginInsertItemPointers(m, v184, v182, l4, l5, l6)
																mBase = m.M
																v186 = m.ExcPending
																if v186 != 0 {
																	return
																} else {
																	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+91)))
																	v188 = int32(0)
																	v192 = F_GinFormTuple(m, l0, v138, v142, v187, v188, v188, v188, int32(1))
																	mBase = m.M
																	v193 = m.ExcPending
																	if v193 != 0 {
																		return
																	} else {
																		*(*uint16)(unsafe.Add(mBase, uint32(v192)+2)) = uint16(v182)
																		v196 = int32(base.Ui32(v182) >> (uint(int32(16)) % 32))
																		*(*uint16)(unsafe.Add(mBase, uint32(v192))) = uint16(v196)
																		v198 = int32(_a_F_ginEntryInsert_1)
																		*(*uint16)(unsafe.Add(mBase, uint32(v192)+4)) = uint16(v198)
																		v201 = v192
																		F_pfree(m, v146)
																		mBase = m.M
																		v203 = m.ExcPending
																		if v203 != 0 {
																			return
																		} else {
																			v204 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v204)
																			v280 = v201
																			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
																			F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
																			mBase = m.M
																			v289 = m.ExcPending
																			if v289 != 0 {
																				return
																			} else {
																				F_pfree(m, v280)
																				mBase = m.M
																				v291 = m.ExcPending
																				if v291 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(96)
																					return
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v201 = v173
															F_pfree(m, v146)
															mBase = m.M
															v203 = m.ExcPending
															if v203 != 0 {
																return
															} else {
																v204 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v204)
																v280 = v201
																*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
																F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
																mBase = m.M
																v289 = m.ExcPending
																if v289 != 0 {
																	return
																} else {
																	F_pfree(m, v280)
																	mBase = m.M
																	v291 = m.ExcPending
																	if v291 != 0 {
																		return
																	} else {
																		m.G0 = v17 + int32(96)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v207 = int32(0)
				v208 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
				if v208 < v207 {
					v212 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[2]))
					v218 = *(*int32)(unsafe.Add(mBase, uint32(v212+(v208^int32(-1))*int32(56))+16))
					v227 = v218
				} else {
					v220 = *(*int32)(unsafe.Add(mBase, _c_F_ginEntryInsert[3]))
					v221 = int32(56)
					v226 = *(*int32)(unsafe.Add(mBase, uint32(v220+v208*v221-v221)+16))
					v227 = v226
				}
				F_CheckForSerializableConflictIn(m, v206, v207, v227)
				mBase = m.M
				v229 = m.ExcPending
				if v229 != 0 {
					return
				} else {
					v230 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					v234 = F_ginCompressPostingList(m, l4, l5, int32(2712), v17+int32(92))
					mBase = m.M
					v235 = m.ExcPending
					if v235 != 0 {
						return
					} else {
						v236 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
						if l5 != v236 {
							F_pfree(m, v234)
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return
							} else {
								v253 = int32(0)
								v257 = F_GinFormTuple(m, l0, v2, l2, v4, v253, v253, v253, int32(1))
								mBase = m.M
								v258 = m.ExcPending
								if v258 != 0 {
									return
								} else {
									v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v260 = F_createPostingTree(m, v259, l4, l5, l6, v230)
									mBase = m.M
									v261 = m.ExcPending
									if v261 != 0 {
										return
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v257)+2)) = uint16(v260)
										v263 = int32(_a_F_ginEntryInsert_1)
										*(*uint16)(unsafe.Add(mBase, uint32(v257)+4)) = uint16(v263)
										v266 = int32(base.Ui32(v260) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v257))) = uint16(v266)
										v269 = v257
										if l6 == int32(0) {
											v280 = v269
										} else {
											v272 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
											*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v272 + int64(1)
											v280 = v269
										}
										*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
										F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
										mBase = m.M
										v289 = m.ExcPending
										if v289 != 0 {
											return
										} else {
											F_pfree(m, v280)
											mBase = m.M
											v291 = m.ExcPending
											if v291 != 0 {
												return
											} else {
												m.G0 = v17 + int32(96)
												return
											}
										}
									}
								}
							}
						} else {
							v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+6)))
							v248 = F_GinFormTuple(m, l0, v2, l2, v4, v234, (v240+int32(1))&int32(_a_F_ginEntryInsert_2)+int32(8), l5, int32(0))
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return
							} else {
								F_pfree(m, v234)
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return
								} else {
									if v248 != 0 {
										v269 = v248
										if l6 == int32(0) {
											v280 = v269
										} else {
											v272 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
											*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v272 + int64(1)
											v280 = v269
										}
										*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
										F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
										mBase = m.M
										v289 = m.ExcPending
										if v289 != 0 {
											return
										} else {
											F_pfree(m, v280)
											mBase = m.M
											v291 = m.ExcPending
											if v291 != 0 {
												return
											} else {
												m.G0 = v17 + int32(96)
												return
											}
										}
									} else {
										v253 = int32(0)
										v257 = F_GinFormTuple(m, l0, v2, l2, v4, v253, v253, v253, int32(1))
										mBase = m.M
										v258 = m.ExcPending
										if v258 != 0 {
											return
										} else {
											v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v260 = F_createPostingTree(m, v259, l4, l5, l6, v230)
											mBase = m.M
											v261 = m.ExcPending
											if v261 != 0 {
												return
											} else {
												*(*uint16)(unsafe.Add(mBase, uint32(v257)+2)) = uint16(v260)
												v263 = int32(_a_F_ginEntryInsert_1)
												*(*uint16)(unsafe.Add(mBase, uint32(v257)+4)) = uint16(v263)
												v266 = int32(base.Ui32(v260) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v257))) = uint16(v266)
												v269 = v257
												if l6 == int32(0) {
													v280 = v269
												} else {
													v272 = *(*int64)(unsafe.Add(mBase, uint32(l6)+16))
													*(*int64)(unsafe.Add(mBase, uint32(l6)+16)) = v272 + int64(1)
													v280 = v269
												}
												*(*int32)(unsafe.Add(mBase, uint32(v17))) = v280
												F_ginInsertValue(m, v17+int32(8), v62, v17, l6)
												mBase = m.M
												v289 = m.ExcPending
												if v289 != 0 {
													return
												} else {
													F_pfree(m, v280)
													mBase = m.M
													v291 = m.ExcPending
													if v291 != 0 {
														return
													} else {
														m.G0 = v17 + int32(96)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ginFreeScanKeys(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[0])))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[1])))
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v10 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[2])))
	F_MemoryContextReset(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L27
	}
L7:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[3])))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+v10<<(uint(int32(2))%32))))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ReleaseBuffer(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+648))
	if v19 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	goto L11
L14:
	;
	F_pfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	if v22 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v25 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_tbm_free(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v29 = v10 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[1])))
	if base.Ui32(v29) < base.Ui32(v30) {
		v10 = v29
		goto L7
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	goto L8
L27:
	;
	v41 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[3]))) = v41
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_ginFreeScanKeys[0]))) = v41
	goto L3
}
func F_ginGetBAEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	v7 = l0 + int32(20)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v8 != 0 {
		v15 = int32(0)
		if v15 == int32(0) {
			return int32(0)
		} else {
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+26)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v20)
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
			if base.B2i32(v29 != int32(1))|base.B2i32(base.Ui32(v27) < base.Ui32(int32(2))) == int32(0) {
				F_pg_qsort(m, v26, v27, int32(6), int32(37))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					return v26
				}
			} else {
				return v26
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v11 = m.T0[v10].(func(*base.Module, int32) int32)(m, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			if v15 == int32(0) {
				return int32(0)
			} else {
				v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+26)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v20)
				v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l2))) = v22
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
				if base.B2i32(v29 != int32(1))|base.B2i32(base.Ui32(v27) < base.Ui32(int32(2))) == int32(0) {
					F_pg_qsort(m, v26, v27, int32(6), int32(37))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v26
					}
				} else {
					return v26
				}
			}
		}
	}
}
func F_ginInsertCleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v423 int32
	_ = v423
	var v424 int64
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v512 int32
	_ = v512
	var v516 int64
	_ = v516
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v541 int64
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int64
	_ = v590
	var v591 int64
	_ = v591
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v618 int64
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int64
	_ = v643
	var v646 int32
	_ = v646
	var v653 int64
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v830 int64
	_ = v830
	var v832 int64
	_ = v832
	var v834 int64
	_ = v834
	var v836 int64
	_ = v836
	var v838 int64
	_ = v838
	var v840 int64
	_ = v840
	var v842 int64
	_ = v842
	var v848 int32
	_ = v848
	var v851 int64
	_ = v851
	var v852 int32
	_ = v852
	var v854 int64
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v869 int32
	_ = v869
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v925 int32
	_ = v925
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v955 int32
	_ = v955
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1054 int32
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1088 int32
	_ = v1088
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	v25 = m.G0
	v27 = v25 - int32(272)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v27 + int32(272)
	return
L2:
	;
	v74 = F_ReadBuffer(m, v29, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L16
	}
L3:
	;
	F_LockPage(m, v29, int32(0), int32(7))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v46 = m.G0
	v48 = v46 - int32(16)
	m.G0 = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = int32(16973824)
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v52
	v61 = F_LockAcquire(m, v48, int32(7), v55, int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L14
	}
L6:
	;
	return
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[0]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[1]))
	if v35 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = v35
	goto L10
L9:
	;
	v40 = v37
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[2]))
	if v42 == int32(4) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v40
	goto L13
L12:
	;
	v45 = v37
	goto L13
L13:
	;
	v72 = v45
	goto L2
L14:
	;
	m.G0 = v48 + int32(16)
	if v61 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[3]))
	v72 = v69
	goto L2
L16:
	;
	F_LockBufferInternal(m, v74, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v74 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	if v97 == int32(-1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82+(v74^int32(-1))<<(uint(int32(2))%32))))
	v96 = v88
	goto L18
L20:
	;
	goto L21
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v96 = v90 + v74<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L22:
	;
	F_UnlockReleaseBuffer(m, v74)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	v107 = F_ReadBuffer(m, v29, v97)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	F_UnlockPage(m, v29, int32(0), int32(7))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	F_LockBufferInternal(m, v107, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v107 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_UnlockBuffer(m, v74)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L33
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115+(v107^int32(-1))<<(uint(int32(2))%32))))
	v129 = v121
	goto L29
L31:
	;
	goto L32
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v129 = v123 + v107<<(uint(int32(13))%32) + int32(-8192)
	goto L29
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6]))
	v138 = F_AllocSetContextCreateInternal(m, v133, int32(_a_F_ginInsertCleanup_0), int32(0), int32(_a_F_ginInsertCleanup_1), int32(_a_F_ginInsertCleanup_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v140 = int32(_a_F_ginInsertCleanup_3)
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6])) = v138
	v146 = F_palloc_mul(m, int32(8), int32(128))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v146
	v151 = F_palloc_mul(m, int32(1), int32(128))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+36)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v151
	F_ginInitBA(m, v27+int32(44))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = l0
	v165 = int32(-1)
	v177 = v129
	v178 = v107
	v180 = v97
	v185 = int32(0)
	goto L38
L38:
	;
	F_processPendingPage(m, v27+int32(44), v27+int32(28), v177, int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	F_UnlockPage(m, v29, int32(0), int32(7))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L6
	} else {
		goto L183
	}
L40:
	;
	goto L39
L41:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v207 = (l1^v165)&base.B2i32(v180 == v106) | v185
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+16)))
	v209 = v177 + v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v210 != int32(-1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L6
	} else {
		goto L177
	}
L44:
	;
	F_UnlockReleaseBuffer(m, v178)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L6
	} else {
		goto L176
	}
L45:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
	if v213&int32(32) == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+12)))
	F_UnlockBuffer(m, v178)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if base.Ui32(v218) < base.Ui32(v72<<(uint(int32(10))%32)) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v220) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v230 = int32(base.Ui32(v220+int32(_a_F_ginInsertCleanup_4)) >> (uint(int32(2)) % 32))
	goto L53
L52:
	;
	v230 = int32(0)
	goto L53
L53:
	;
	v232 = v27 + int32(44)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
	v236 = m.G0
	v237 = int32(16)
	v238 = v236 - v237
	m.G0 = v238
	v241 = v27 + int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v235
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+12)) = uint8(base.B2i32(v245 == int32(_a_F_ginInsertCleanup_5)))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = int32(836)
	m.G0 = v238 + v237
	goto L54
L54:
	;
	v262 = F_ginGetBAEntry(m, v232, v27+int32(12), v27+int32(16), v27+int32(15), v27+int32(24))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	if v262 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v267 = v262
	goto L59
L57:
	;
	goto L58
L58:
	;
	F_LockBufferInternal(m, v74, int32(3))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L65
	}
L59:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v290 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+15)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	F_ginEntryInsert(m, l0, v288, v289, v290, v267, v291, int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v308 = F_ginGetBAEntry(m, v27+int32(44), v27+int32(12), v27+int32(16), v27+int32(15), v27+int32(24))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	if v308 != 0 {
		v267 = v308
		goto L59
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	F_LockBufferInternal(m, v178, int32(1))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v340) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+16)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v177+v466)))
	F_UnlockReleaseBuffer(m, v178)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L6
	} else {
		goto L82
	}
L68:
	;
	v348 = int32(base.Ui32(v340+int32(_a_F_ginInsertCleanup_4)) >> (uint(int32(2)) % 32))
	goto L70
L69:
	;
	v348 = int32(0)
	goto L70
L70:
	;
	v349 = int32(_a_F_ginInsertCleanup_6)
	if v348&v349 == v230&v349 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v355 = v27 + int32(44)
	F_ginInitBA(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	F_processPendingPage(m, v355, v27+int32(28), v177, (v230+int32(1))&int32(_a_F_ginInsertCleanup_6))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v369 = m.G0
	v370 = int32(16)
	v371 = v369 - v370
	m.G0 = v371
	v374 = v27 + int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v374))) = v368
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+12)) = uint8(base.B2i32(v378 == int32(_a_F_ginInsertCleanup_5)))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = int32(836)
	m.G0 = v371 + v370
	goto L74
L74:
	;
	v395 = F_ginGetBAEntry(m, v355, v27+int32(12), v27+int32(16), v27+int32(15), v27+int32(24))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	if v395 == int32(0) {
		goto L67
	} else {
		goto L76
	}
L76:
	;
	v402 = v395
	goto L77
L77:
	;
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v425 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+15)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	F_ginEntryInsert(m, l0, v423, v424, v425, v402, v426, int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L79
	}
L78:
	;
	goto L67
L79:
	;
	v440 = F_ginGetBAEntry(m, v27+int32(44), v27+int32(12), v27+int32(16), v27+int32(15), v27+int32(24))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	if v440 != 0 {
		v402 = v440
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	if v74 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v484 = v482 + int32(32)
	v486 = v482 + int32(24)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v482)+24))
	v489 = v487
	goto L87
L84:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v474+(v74^v165)<<(uint(int32(2))%32))))
	v482 = v476
	goto L83
L85:
	;
	goto L86
L86:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v482 = v478 + v74<<(uint(int32(13))%32) + int32(-8192)
	goto L83
L87:
	;
	v512 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v512
	v516 = int64(0)
	if v489 == v468 {
		v596 = v468
		v601 = v512
		v611 = v512
		v618 = v516
		goto L89
	} else {
		goto L90
	}
L88:
	;
	F_UnlockBuffer(m, v74)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L6
	} else {
		goto L170
	}
L89:
	;
	if l4 != 0 {
		goto L103
	} else {
		goto L104
	}
L90:
	;
	v521 = v489
	v524 = v512
	v541 = v516
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(80)+v524<<(uint(int32(2))%32)))) = v521
	v548 = F_ReadBuffer(m, v29, v521)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L93
	}
L92:
	;
	v596 = v588
	v601 = v584
	v611 = v589
	v618 = v591
	goto L89
L93:
	;
	v551 = v27 + int32(144)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v551+v552<<(uint(int32(2))%32)))) = v548
	F_LockBufferInternal(m, v548, int32(3))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v551+v560<<(uint(int32(2))%32))))
	if v564 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v584 = v560 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+264)) = v584
	v586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v582)+16)))
	v587 = v582 + v586
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v589 = base.B2i32(v588 != v468)
	v590 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v587)+4)))
	v591 = v541 + v590
	if int32(15) < v584 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v568+(v564^int32(-1))<<(uint(int32(2))%32))))
	v582 = v574
	goto L95
L97:
	;
	goto L98
L98:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v582 = v576 + v564<<(uint(int32(13))%32) + int32(-8192)
	goto L95
L99:
	;
	v596 = v588
	v601 = v584
	v611 = v589
	v618 = v591
	goto L89
L100:
	;
	goto L101
L101:
	;
	if v588 != v468 {
		v521 = v588
		v524 = v584
		v541 = v591
		goto L91
	} else {
		goto L102
	}
L102:
	;
	goto L92
L103:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+28)) = v619 + v601
	goto L105
L104:
	;
	goto L105
L105:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622)+118)))
	if v623 != int32(112) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v635 = int32(_a_F_ginInsertCleanup_7)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7])) = v637 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+24)) = v596
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v482)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v482)+40)) = v643 - v618
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v482)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+36)) = v646 - v642
	if v596 == int32(-1) {
		goto L114
	} else {
		goto L115
	}
L107:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[8]))
	if v627 <= int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v630 != 0 {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	F_XLogEnsureRecordSpace(m, v601, int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if v631 != 0 {
		goto L106
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L106
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+28)) = int32(-1)
	v653 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v484)+8)) = v653
	*(*int64)(unsafe.Add(mBase, uint32(v484))) = v653
	goto L116
L115:
	;
	goto L116
L116:
	;
	v657 = int32(80)
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+12)) = uint16(v657)
	F_MarkBufferDirty(m, v74)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v661 = int32(0)
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v661 < v662 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v668 = v661
	goto L121
L119:
	;
	v729 = v662
	goto L120
L120:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747)+118)))
	if v748 != int32(112) {
		v925 = v729
		goto L130
	} else {
		goto L131
	}
L121:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(144)+v668<<(uint(int32(2))%32))))
	if v694 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v729 = v721
	goto L120
L123:
	;
	v713 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v712)+16)))
	v715 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v713+v712)+6)) = uint16(v715)
	F_MarkBufferDirty(m, v694)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L127
	}
L124:
	;
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v698+(v694^int32(-1))<<(uint(int32(2))%32))))
	v712 = v704
	goto L123
L125:
	;
	goto L126
L126:
	;
	v706 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v712 = v706 + v694<<(uint(int32(13))%32) + int32(-8192)
	goto L123
L127:
	;
	v720 = v668 + int32(1)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v720 < v721 {
		v668 = v720
		goto L121
	} else {
		goto L128
	}
L128:
	;
	goto L122
L129:
	;
	if v611 != 0 {
		v489 = v596
		goto L87
	} else {
		goto L169
	}
L130:
	;
	v943 = int32(0)
	v944 = int32(_a_F_ginInsertCleanup_7)
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7])) = v946 - int32(1)
	if v925 <= v943 {
		goto L129
	} else {
		goto L158
	}
L131:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[8]))
	if v752 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v755 != 0 {
		v925 = v729
		goto L130
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if v756 != 0 {
		v925 = v729
		goto L130
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v759 = int32(0)
	F_XLogRegisterBuffer(m, v759, v74, int32(14))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if int32(0) < v764 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v770 = v759
	goto L142
L140:
	;
	goto L141
L141:
	;
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v486)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+256)) = v830
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v486)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+248)) = v832
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v486)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+240)) = v834
	v836 = *(*int64)(unsafe.Add(mBase, uint32(v486)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+232)) = v836
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v486)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+224)) = v838
	v840 = *(*int64)(unsafe.Add(mBase, uint32(v486)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+216)) = v840
	v842 = *(*int64)(unsafe.Add(mBase, uint32(v486)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+208)) = v842
	F_XLogRegisterData(m, v27+int32(208), int32(64))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L6
	} else {
		goto L146
	}
L142:
	;
	v792 = v770 + int32(1)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(144)+v770<<(uint(int32(2))%32))))
	F_XLogRegisterBuffer(m, v792&int32(255), v800, int32(6))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L6
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v792 < v804 {
		v770 = v792
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v851 = F_XLogInsert(m, int32(13), int32(128))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	v854 = base.I64_rotl(v851, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v482))) = v854
	v856 = int32(0)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v857 <= v856 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v860 = int32(_a_F_ginInsertCleanup_7)
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[7])) = v862 - int32(1)
	goto L129
L149:
	;
	goto L150
L150:
	;
	v869 = v856
	goto L151
L151:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(144)+v869<<(uint(int32(2))%32))))
	if v895 < int32(0) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v925 = v917
	goto L130
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v913))) = v854
	v916 = v869 + int32(1)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v916 < v917 {
		v869 = v916
		goto L151
	} else {
		goto L157
	}
L154:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v899+(v895^int32(-1))<<(uint(int32(2))%32))))
	v913 = v905
	goto L153
L155:
	;
	goto L156
L156:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v913 = v907 + v895<<(uint(int32(13))%32) + int32(-8192)
	goto L153
L157:
	;
	goto L152
L158:
	;
	v955 = v943
	goto L159
L159:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(144)+v955<<(uint(int32(2))%32))))
	F_UnlockReleaseBuffer(m, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L6
	} else {
		goto L161
	}
L160:
	;
	if l2 == int32(0) {
		goto L129
	} else {
		goto L163
	}
L161:
	;
	v985 = v955 + int32(1)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v985 < v986 {
		v955 = v985
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v990 = int32(0)
	if v986 <= v990 {
		goto L129
	} else {
		goto L164
	}
L164:
	;
	v996 = v990
	goto L165
L165:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(80)+v996<<(uint(int32(2))%32))))
	F_RecordFreeIndexPage(m, v29, v1022)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L6
	} else {
		goto L167
	}
L166:
	;
	goto L129
L167:
	;
	v1026 = v996 + int32(1)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	if v1026 < v1027 {
		v996 = v1026
		goto L165
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	goto L88
L170:
	;
	if (base.B2i32(v468 == int32(-1))|v207)&int32(1) != 0 {
		goto L40
	} else {
		goto L171
	}
L171:
	;
	F_MemoryContextReset(m, v138)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	v1064 = F_palloc_mul(m, int32(8), v1063)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v1064
	v1068 = F_palloc_mul(m, int32(1), v1063)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v1068
	F_ginInitBA(m, v27+int32(44))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	v1088 = v468
	goto L43
L176:
	;
	v1088 = v210
	goto L43
L177:
	;
	v1106 = F_ReadBuffer(m, v29, v1088)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	F_LockBufferInternal(m, v1106, int32(1))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	if v1106 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[4]))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1114+(v1106^int32(-1))<<(uint(int32(2))%32))))
	v1128 = v1120
	goto L182
L181:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[5]))
	v1128 = v1122 + v1106<<(uint(int32(13))%32) + int32(-8192)
	goto L182
L182:
	;
	v177 = v1128
	v178 = v1106
	v180 = v1088
	v185 = v207
	goto L38
L183:
	;
	F_ReleaseBuffer(m, v74)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	if l2 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_FreeSpaceMapVacuum(m, v29)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L6
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginInsertCleanup[6])) = v141
	F_MemoryContextDelete(m, v138)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L6
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	goto L1
}
func F_ginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v211 int64
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v713 int64
	_ = v713
	var v714 int32
	_ = v714
	var v716 int64
	_ = v716
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 < v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v39
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2]))
	v51 = F_AllocSetContextCreateInternal(m, v46, int32(_a_F_ginPlaceToPage_0), v39, int32(_a_F_ginPlaceToPage_1), int32(_a_F_ginPlaceToPage_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
	v38 = v30
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v55 = int32(_a_F_ginPlaceToPage_3)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2])) = v51
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+v59)+6)))
	v63 = v61 & int32(1)
	if v61&int32(2) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v87 = int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, v88, l1, l2, l3, v16+int32(-12), v16+int32(-4), v16+int32(-8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L18
	}
L8:
	;
	v85 = v63 | int32(2)
	v86 = v7
	goto L7
L9:
	;
	goto L10
L10:
	;
	if l4 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71+(l4^int32(-1))<<(uint(int32(2))%32))))
	v85 = v63
	v86 = v77
	goto L7
L12:
	;
	goto L13
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v85 = v63
	v86 = v79 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L5
	} else {
		goto L193
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[2])) = v56
	F_MemoryContextDelete(m, v51)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L192
	}
L16:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v224 = F_GinNewBuffer(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L59
	}
L17:
	;
	v98 = int32(_a_F_ginPlaceToPage_4)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v100 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+48))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+118)))
	if v106 != int32(112) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	switch v96 {
	case 0:
		v783 = v87
		goto L15
	case 1:
		goto L17
	case 2:
		goto L16
	default:
		goto L14
	}
L19:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	m.T0[v120].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, v118, l1, l2, l3, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L28
	}
L20:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
	if v110 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	if v113 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v115 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+40))
	if v114 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	if l4 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+118)))
	if v152 != int32(112) {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v126 = v86 + v125
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+6)))
	v129 = v127 & int32(_a_F_ginPlaceToPage_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+6)) = uint16(v129)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+118)))
	if v135 != int32(112) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
	if v139 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)+32))
	if v142 != 0 {
		goto L29
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v144 != 0 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if v143 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_XLogRegisterBuffer(m, int32(1), l4, int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	goto L29
L40:
	;
	v217 = int32(_a_F_ginPlaceToPage_4)
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v219 - int32(1)
	v783 = v87
	goto L15
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
	if v156 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	if v159 != 0 {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v161 != 0 {
		goto L40
	} else {
		goto L47
	}
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+40))
	if v160 != 0 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+50)) = uint16(v85)
	F_XLogRegisterData(m, v16+int32(-14), int32(2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if l4 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if l4 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v211 = F_XLogInsert(m, int32(13), int32(32))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L58
	}
L52:
	;
	v187 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = base.I32_rotr(v186, v187)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v86+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = base.I32_rotr(v192, v187)
	F_XLogRegisterData(m, v16+int32(-48), int32(8))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L5
	} else {
		goto L56
	}
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+(l4^int32(-1))*int32(56))+16))
	v186 = v177
	goto L52
L54:
	;
	goto L55
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v180 = int32(56)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179+l4*v180-v180)+16))
	v186 = v185
	goto L52
L56:
	;
	v203 = F_XLogInsert(m, int32(13), int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v206 = base.I64_rotl(v203, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v206
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v206
	goto L40
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_rotl(v211, int64(32))
	goto L40
L59:
	;
	if l5 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v38+v239)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v243
	if l4 != 0 {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v228 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v231 + int32(1)
	goto L60
L63:
	;
	goto L64
L64:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v235 + int32(1)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v276 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L66:
	;
	if l4 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L68
L68:
	;
	v271 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v271
	v274 = v271
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v266
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v86+v268)))
	v274 = v270
	goto L65
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251+(l4^int32(-1))*int32(56))+16))
	v266 = v257
	goto L69
L71:
	;
	goto L72
L72:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v260 = int32(56)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259+l4*v260-v260)+16))
	v266 = v265
	goto L69
L73:
	;
	v574 = int32(_a_F_ginPlaceToPage_4)
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v576 + int32(1)
	F_MarkBufferDirty(m, v224)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L132
	}
L74:
	;
	v569 = v565
	v573 = v568
	goto L73
L75:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v522 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L76:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v280 = F_GinNewBuffer(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v241
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v457+v458))) = v241
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461)+16)))
	v463 = v461 + v462
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v463)+6)))
	v466 = v464 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v463)+6)) = uint16(v466)
	if v224 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L79:
	;
	if l5 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v295 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v295
	v298 = v85 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v298)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v300+v301))) = v295
	if v224 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v284 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v287 + int32(1)
	goto L80
L83:
	;
	goto L84
L84:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = v291 + int32(1)
	goto L80
L85:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v324+v325))) = v323
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v329 = F_PageGetTempPage(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L89
	}
L86:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308+(v224^int32(-1))*int32(56))+16))
	v323 = v314
	goto L85
L87:
	;
	goto L88
L88:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v317 = int32(56)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v316+v224*v317-v317)+16))
	v323 = v322
	goto L85
L89:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331)+16)))
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v331+v332)+6)))
	v336 = v334 & int32(_a_F_ginPlaceToPage_6)
	F_PageInit(m, v329, int32(_a_F_ginPlaceToPage_1), int32(8))
	mBase = m.M
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+16)))
	v341 = v329 + v340
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v341)+6)) = uint16(v336)
	goto L90
L90:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v280 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v224 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v349+(v280^int32(-1))*int32(56))+16))
	v364 = v355
	goto L91
L93:
	;
	goto L94
L94:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v358 = int32(56)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357+v280*v358-v358)+16))
	v364 = v363
	goto L91
L95:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	m.T0[v345].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, v329, v364, v365, v384, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L99
	}
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v369+(v224^int32(-1))*int32(56))+16))
	v384 = v375
	goto L95
L97:
	;
	goto L98
L98:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v378 = int32(56)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v377+v224*v378-v378)+16))
	v384 = v383
	goto L95
L99:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v388 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v406)+16)))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407+v406)+6)))
	if v409&int32(2) == int32(0) {
		v565 = v280
		v568 = v329
		goto L74
	} else {
		goto L104
	}
L101:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v392+(v388^int32(-1))<<(uint(int32(2))%32))))
	v406 = v398
	goto L100
L102:
	;
	goto L103
L103:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v406 = v400 + v388<<(uint(int32(13))%32) + int32(-8192)
	goto L100
L104:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v388 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v280 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v418+(v388^int32(-1))*int32(56))+16))
	v433 = v424
	goto L105
L107:
	;
	goto L108
L108:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v427 = int32(56)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v426+v388*v427-v427)+16))
	v433 = v432
	goto L105
L109:
	;
	F_PredicateLockPageSplit(m, v414, v433, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L5
	} else {
		goto L113
	}
L110:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v437+(v280^int32(-1))*int32(56))+16))
	v452 = v443
	goto L109
L111:
	;
	goto L112
L112:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v446 = int32(56)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v445+v280*v446-v446)+16))
	v452 = v451
	goto L109
L113:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v520 = v280
	v522 = v455
	v523 = v329
	goto L75
L114:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v487+v488))) = v486
	v491 = int32(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v494 < v491 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v471+(v224^int32(-1))*int32(56))+16))
	v486 = v477
	goto L114
L116:
	;
	goto L117
L117:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v480 = int32(56)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479+v224*v480-v480)+16))
	v486 = v485
	goto L114
L118:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v512)+16)))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v512)+6)))
	if v515&int32(2) == int32(0) {
		v569 = v491
		v573 = v491
		goto L73
	} else {
		goto L122
	}
L119:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v498+(v494^int32(-1))<<(uint(int32(2))%32))))
	v512 = v504
	goto L118
L120:
	;
	goto L121
L121:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v512 = v506 + v494<<(uint(int32(13))%32) + int32(-8192)
	goto L118
L122:
	;
	v520 = v491
	v522 = v494
	v523 = v491
	goto L75
L123:
	;
	if v224 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v528+(v522^int32(-1))*int32(56))+16))
	v543 = v534
	goto L123
L125:
	;
	goto L126
L126:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v537 = int32(56)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v536+v522*v537-v537)+16))
	v543 = v542
	goto L123
L127:
	;
	F_PredicateLockPageSplit(m, v524, v543, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L5
	} else {
		goto L131
	}
L128:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[5]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v547+(v224^int32(-1))*int32(56))+16))
	v562 = v553
	goto L127
L129:
	;
	goto L130
L130:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[6]))
	v556 = int32(56)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v555+v224*v556-v556)+16))
	v562 = v561
	goto L127
L131:
	;
	v565 = v520
	v568 = v523
	goto L74
L132:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_MarkBufferDirty(m, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v585 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	base.MemoryCopy(m, v650, v651, int32(_a_F_ginPlaceToPage_1))
	if l4 != 0 {
		goto L149
	} else {
		goto L150
	}
L135:
	;
	F_MarkBufferDirty(m, v569)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L5
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	base.MemoryCopy(m, v38, v630, int32(_a_F_ginPlaceToPage_1))
	if v224 < int32(0) {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	base.MemoryCopy(m, v38, v573, int32(_a_F_ginPlaceToPage_1))
	if v569 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	base.MemoryCopy(m, v609, v610, int32(_a_F_ginPlaceToPage_1))
	if v224 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v595+(v569^int32(-1))<<(uint(int32(2))%32))))
	v609 = v601
	goto L139
L141:
	;
	goto L142
L142:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v609 = v603 + v569<<(uint(int32(13))%32) + int32(-8192)
	goto L139
L143:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v616+(v224^int32(-1))<<(uint(int32(2))%32))))
	v650 = v622
	goto L134
L144:
	;
	goto L145
L145:
	;
	v624 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v650 = v624 + v224<<(uint(int32(13))%32) + int32(-8192)
	goto L134
L146:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v636+(v224^int32(-1))<<(uint(int32(2))%32))))
	v650 = v642
	goto L134
L147:
	;
	goto L148
L148:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v650 = v644 + v224<<(uint(int32(13))%32) + int32(-8192)
	goto L134
L149:
	;
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+16)))
	v655 = v86 + v654
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v655)+6)))
	v658 = v656 & int32(_a_F_ginPlaceToPage_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v655)+6)) = uint16(v658)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L5
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+48))
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+118)))
	if v665 != int32(112) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v764 = int32(_a_F_ginPlaceToPage_4)
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[3])) = v766 - int32(1)
	F_UnlockReleaseBuffer(m, v224)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L5
	} else {
		goto L189
	}
L154:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[4]))
	if v669 <= int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v663)+32))
	if v672 != 0 {
		goto L153
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v674 != 0 {
		goto L153
	} else {
		goto L160
	}
L158:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v663)+40))
	if v673 != 0 {
		goto L153
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v677 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	if l4 != 0 {
		goto L171
	} else {
		goto L172
	}
L163:
	;
	F_XLogRegisterBuffer(m, int32(0), v569, int32(9))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L5
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_XLogRegisterBuffer(m, int32(0), v694, int32(9))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L5
	} else {
		goto L169
	}
L166:
	;
	F_XLogRegisterBuffer(m, int32(1), v224, int32(9))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L5
	} else {
		goto L167
	}
L167:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_XLogRegisterBuffer(m, int32(2), v689, int32(9))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	goto L162
L169:
	;
	F_XLogRegisterBuffer(m, int32(1), v224, int32(9))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L170
	}
L170:
	;
	goto L162
L171:
	;
	F_XLogRegisterBuffer(m, int32(3), l4, int32(8))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	F_XLogRegisterData(m, v16+int32(-48), int32(28))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L5
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	v713 = F_XLogInsert(m, int32(13), int32(48))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L5
	} else {
		goto L176
	}
L176:
	;
	v716 = base.I64_rotl(v713, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v716
	if v224 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v735))) = v716
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v737 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v721+(v224^int32(-1))<<(uint(int32(2))%32))))
	v735 = v727
	goto L177
L179:
	;
	goto L180
L180:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v735 = v729 + v224<<(uint(int32(13))%32) + int32(-8192)
	goto L177
L181:
	;
	if v569 < int32(0) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	goto L183
L183:
	;
	if l4 == int32(0) {
		goto L153
	} else {
		goto L188
	}
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v757))) = v716
	goto L183
L185:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[0]))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v743+(v569^int32(-1))<<(uint(int32(2))%32))))
	v757 = v749
	goto L184
L186:
	;
	goto L187
L187:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_ginPlaceToPage[1]))
	v757 = v751 + v569<<(uint(int32(13))%32) + int32(-8192)
	goto L184
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v716
	goto L153
L189:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v772 != 0 {
		v783 = int32(0)
		goto L15
	} else {
		goto L190
	}
L190:
	;
	F_UnlockReleaseBuffer(m, v569)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L5
	} else {
		goto L191
	}
L191:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v783 = base.B2i32(v775 == int32(0))
	goto L15
L192:
	;
	m.G0 = v18 - int32(-64)
	return v783
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v96
	F_errmsg_internal(m, int32(_a_F_ginPlaceToPage_7), v18)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_ginPlaceToPage_8), int32(650), int32(_a_F_ginPlaceToPage_9))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_cmp_tslexeme(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = int32(1)
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v18 = v16 & v14
			if v18 != 0 {
				v19 = v14
			} else {
				v19 = int32(4)
			}
			if v16 == int32(1) {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v26 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v36 = int32(4)
				} else {
					v36 = v29
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v18 != 0 {
					v47 = int32(base.Ui32(v16)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v48 = int32(1)
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v52 = v50 & v48
			if v52 != 0 {
				v53 = v48
			} else {
				v53 = int32(4)
			}
			if v50 == int32(1) {
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
				if v60 == int32(18) {
					v63 = int32(16)
				} else {
					v63 = int32(0)
				}
				if base.Ui32((v60-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v70 = int32(4)
				} else {
					v70 = v63
				}
				v81 = v70
			} else {
				v71 = int32(1)
				if v52 != 0 {
					v81 = int32(base.Ui32(v50)>>(uint(v71)%32)) - v71
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v47 == int32(0) {
				v87 = int32(0)
				if v87 < v81 {
					v90 = int32(-1)
				} else {
					v90 = v87
				}
				v107 = v90
			} else {
				if v81 == int32(0) {
					v107 = base.B2i32(int32(0) < v47)
				} else {
					if base.Ui32(v47) < base.Ui32(v81) {
						v96 = v47
					} else {
						v96 = v81
					}
					v97 = F_memcmp(m, v19+v7, v12+v53, v96)
					mBase = m.M
					if v97 != 0 {
						v105 = v97
						v107 = v105
					} else {
						if v47 == v81 {
							v107 = int32(0)
						} else {
							if v47 < v81 {
								v104 = int32(-1)
							} else {
								v104 = int32(1)
							}
							v105 = v104
							v107 = v105
						}
					}
				}
			}
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v108 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int64(0)
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v112 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(v107)
						}
					} else {
						return base.I64_extend_i32_s(v107)
					}
				}
			} else {
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v112 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_s(v107)
					}
				} else {
					return base.I64_extend_i32_s(v107)
				}
			}
		}
	}
}
func F_gin_compare_prefix_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_CallerFInfoFunctionCall2(m, v9, v10, v11, v12, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = base.I32_wrap_i64(v14)
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
		switch v19&int32(15) - int32(1) {
		case 0:
			v53 = base.B2i32(v18 <= int32(0))
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_s(v53)
		case 1:
			v53 = int32(base.Ui32(v18) >> (uint(int32(31)) % 32))
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_s(v53)
		case 2:
			v53 = base.B2i32(v18 != int32(0))
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_s(v53)
		case 3:
			v29 = int32(0)
			if v29 < v18 {
				v32 = int32(-1)
			} else {
				v32 = v29
			}
			v53 = v32
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_s(v53)
		case 4:
			v53 = (v18 ^ int32(-1)) >> (uint(int32(31)) % 32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_s(v53)
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v41
				F_errmsg_internal(m, int32(_a_F_gin_compare_prefix_int2_0), v6)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_gin_compare_prefix_int2_1), int32(228), int32(_a_F_gin_compare_prefix_int2_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_gin_extract_jsonb(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v101 int64
	_ = v101
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v20 = v18 & int32(268435455)
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v10 + int32(48)
	return v101
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
	v101 = int64(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = v20 << (uint(int32(1)) % 32)
	v28 = F_palloc_mul(m, int32(8), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v32 = F_JsonbIteratorInit(m, v13+int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v32
	v36 = int32(0)
	v37 = v27
	v39 = v28
	goto L9
L9:
	;
	v48 = F_JsonbIteratorNext(m, v10+int32(44), v10+int32(8), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87+v36<<(uint(int32(3))%32)))) = v88
	v36 = v36 + int32(1)
	v37 = v86
	v39 = v87
	goto L9
L12:
	;
	v83 = v37 << (uint(int32(1)) % 32)
	v84 = F_repalloc_mul(m, v39, int32(8), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v36
	v101 = base.I64_extend_i32_u(v39)
	goto L3
L14:
	;
	v73 = int32(8)
	v76 = F_palloc_mul(m, v73, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L15:
	;
	v69 = F_make_scalar_key(m, v10+int32(8), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L25
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v63 = F_make_scalar_key(m, v10+int32(8), base.B2i32(v60 == int32(1)))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	v53 = F_make_scalar_key(m, v10+int32(8), int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	switch v48 {
	case 0:
		goto L13
	case 1:
		goto L17
	case 2:
		goto L15
	case 3:
		goto L16
	default:
		goto L9
	}
L19:
	;
	if v36 < v37 {
		v86 = v37
		v87 = v39
		v88 = v53
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v37 == int32(0) {
		v72 = v53
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v80 = v53
	goto L12
L22:
	;
	if v36 < v37 {
		v86 = v37
		v87 = v39
		v88 = v63
		goto L11
	} else {
		goto L23
	}
L23:
	;
	if v37 != 0 {
		v80 = v63
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v72 = v63
	goto L14
L25:
	;
	if v36 < v37 {
		v86 = v37
		v87 = v39
		v88 = v69
		goto L11
	} else {
		goto L26
	}
L26:
	;
	if v37 != 0 {
		v80 = v69
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v72 = v69
	goto L14
L28:
	;
	v86 = v73
	v87 = v76
	v88 = v72
	goto L11
L29:
	;
	v86 = v83
	v87 = v84
	v88 = v80
	goto L11
}
func F_gin_extract_query_name(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_name_0), int32(_a_F_gin_extract_query_name_1), int32(_a_F_gin_extract_query_name_2), int32(_a_F_gin_extract_query_name_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_tsvector_2args(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v2 <= int32(2) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_gin_extract_tsvector_2args_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_gin_extract_tsvector_2args_1), int32(313), int32(_a_F_gin_extract_tsvector_2args_2))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v20 = F_gin_extract_tsvector(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			return v20
		}
	}
}
func F_gin_extract_value_inet(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc(m, int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = F_pg_detoast_datum(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = base.I64_extend_i32_u(v11)
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
			return base.I64_extend_i32_u(v7)
		}
	}
}
func F_gin_leafpage_items(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int64
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v291 int64
	_ = v291
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	v14 = int64(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		v24 = F_superuser(m)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			if v24 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				if v27 == int32(0) {
					v30 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v32 = int32(_a_F_gin_leafpage_items_0)
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0]))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0])) = v35
						v37 = F_get_page_from_raw(m, v20)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+14)))
							if v39 == int32(0) {
								*(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0])) = v33
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
								v291 = v14
								m.G0 = v17 - int32(-64)
								return v291
							} else {
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+19)))
								v47 = int32(8)
								v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)))
								if (v46<<(uint(v47)%32)-v49)&int32(_a_F_gin_leafpage_items_1) != v47 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v315 = m.ExcPending
									if v315 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v318 = m.ExcPending
										if v318 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_gin_leafpage_items_2), int32(0))
											mBase = m.M
											v322 = m.ExcPending
											if v322 != 0 {
												return int64(0)
											} else {
												v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)))
												v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+19)))
												v325 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v325
												*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = (v324<<(uint(v325)%32) - v323) & int32(_a_F_gin_leafpage_items_1)
												v336 = F_errdetail(m, int32(_a_F_gin_leafpage_items_3), v15+int32(-48))
												mBase = m.M
												v337 = m.ExcPending
												if v337 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(214), int32(_a_F_gin_leafpage_items_5))
													mBase = m.M
													v342 = m.ExcPending
													if v342 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									}
								} else {
									v55 = v37 + v49
									v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)))
									if v56 != int32(131) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v346 = m.ExcPending
										if v346 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v349 = m.ExcPending
											if v349 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_gin_leafpage_items_6), int32(0))
												mBase = m.M
												v353 = m.ExcPending
												if v353 != 0 {
													return int64(0)
												} else {
													v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)))
													*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(131)
													*(*int32)(unsafe.Add(mBase, uint32(v17))) = v354
													v359 = F_errdetail(m, int32(_a_F_gin_leafpage_items_7), v17)
													mBase = m.M
													v360 = m.ExcPending
													if v360 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(223), int32(_a_F_gin_leafpage_items_5))
														mBase = m.M
														v365 = m.ExcPending
														if v365 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									} else {
										v60 = F_palloc(m, int32(12))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int64(0)
										} else {
											v65 = F_get_call_result_type(m, l0, int32(0), v15+int32(-32))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												if v65 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v369 = m.ExcPending
													if v369 != 0 {
														return int64(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_gin_leafpage_items_8), int32(0))
														mBase = m.M
														v373 = m.ExcPending
														if v373 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(229), int32(_a_F_gin_leafpage_items_5))
															mBase = m.M
															v378 = m.ExcPending
															if v378 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
													*(*int32)(unsafe.Add(mBase, uint32(v60))) = v69
													v71 = int32(32)
													v72 = v37 + v71
													*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v72
													v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)))
													*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v72 + v74 - v71
													*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v60
													*(*int32)(unsafe.Add(mBase, _c_F_gin_leafpage_items[0])) = v33
													v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
													if v90 != v91 {
														v93 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v17)+28)) = uint16(v93)
														*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)) = uint8(v93)
														*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = base.I64_extend_i32_u(v90)
														v100 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
														*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v100
														v104 = F_ginPostingListDecode(m, v90, v15+int32(-40))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return int64(0)
														} else {
															v106 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
															v109 = F_palloc(m, v106<<(uint(int32(3))%32))
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return int64(0)
															} else {
																v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
																if v111 <= int32(0) {
																} else {
																	v115 = v111 & int32(3)
																	if base.Ui32(int32(4)) <= base.Ui32(v111) {
																		v122 = v93
																		v127 = int32(0)
																		for {
																			v135 = int32(3)
																			v138 = int32(6)
																			*(*int64)(unsafe.Add(mBase, uint32(v109+v122<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v122*v138)
																			v144 = v122 | int32(1)
																			*(*int64)(unsafe.Add(mBase, uint32(v109+v144<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v144*v138)
																			v154 = v122 | int32(2)
																			*(*int64)(unsafe.Add(mBase, uint32(v109+v154<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v154*v138)
																			v164 = v122 | v135
																			*(*int64)(unsafe.Add(mBase, uint32(v109+v164<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v164*v138)
																			v173 = int32(4)
																			v174 = v122 + v173
																			v176 = v127 + v173
																			if v176 != v111&int32(2147483644) {
																				v122 = v174
																				v127 = v176
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v115 == int32(0) {
																		} else {
																			v181 = v174
																			v196 = v181
																			v201 = int32(0)
																			for {
																				*(*int64)(unsafe.Add(mBase, uint32(v109+v196<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v104 + v196*int32(6))
																				v217 = int32(1)
																				v220 = v201 + v217
																				if v220 != v115 {
																					v196 = v196 + v217
																					v201 = v220
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v181 = v93
																		v196 = v181
																		v201 = int32(0)
																		for {
																			*(*int64)(unsafe.Add(mBase, uint32(v109+v196<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v104 + v196*int32(6))
																			v217 = int32(1)
																			v220 = v201 + v217
																			if v220 != v115 {
																				v196 = v196 + v217
																				v201 = v220
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
																v237 = F_construct_array_builtin(m, v109, v111, int32(27))
																mBase = m.M
																v238 = m.ExcPending
																if v238 != 0 {
																	return int64(0)
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = base.I64_extend_i32_u(v237)
																	F_pfree(m, v109)
																	mBase = m.M
																	v242 = m.ExcPending
																	if v242 != 0 {
																		return int64(0)
																	} else {
																		F_pfree(m, v104)
																		mBase = m.M
																		v244 = m.ExcPending
																		if v244 != 0 {
																			return int64(0)
																		} else {
																			v245 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
																			v250 = F_heap_form_tuple(m, v245, v15+int32(-32), v15+int32(-36))
																			mBase = m.M
																			v251 = m.ExcPending
																			if v251 != 0 {
																				return int64(0)
																			} else {
																				v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
																				v253 = F_HeapTupleHeaderGetDatum(m, v252)
																				mBase = m.M
																				v254 = m.ExcPending
																				if v254 != 0 {
																					return int64(0)
																				} else {
																					v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
																					v256 = int32(1)
																					*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v90 + (v255+v256)&int32(_a_F_gin_leafpage_items_9) + int32(8)
																					v264 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
																					*(*int64)(unsafe.Add(mBase, uint32(v88))) = v264 + int64(1)
																					v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v256
																					v291 = v253
																					m.G0 = v17 - int32(-64)
																					return v291
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v272 = m.ExcPending
														if v272 != 0 {
															return int64(0)
														} else {
															v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v273)+20)) = int32(2)
															v276 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v276)
															v291 = v14
															m.G0 = v17 - int32(-64)
															return v291
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
					if v90 != v91 {
						v93 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v17)+28)) = uint16(v93)
						*(*uint8)(unsafe.Add(mBase, uint32(v17)+30)) = uint8(v93)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = base.I64_extend_i32_u(v90)
						v100 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
						*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v100
						v104 = F_ginPostingListDecode(m, v90, v15+int32(-40))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int64(0)
						} else {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
							v109 = F_palloc(m, v106<<(uint(int32(3))%32))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int64(0)
							} else {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
								if v111 <= int32(0) {
								} else {
									v115 = v111 & int32(3)
									if base.Ui32(int32(4)) <= base.Ui32(v111) {
										v122 = v93
										v127 = int32(0)
										for {
											v135 = int32(3)
											v138 = int32(6)
											*(*int64)(unsafe.Add(mBase, uint32(v109+v122<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v122*v138)
											v144 = v122 | int32(1)
											*(*int64)(unsafe.Add(mBase, uint32(v109+v144<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v144*v138)
											v154 = v122 | int32(2)
											*(*int64)(unsafe.Add(mBase, uint32(v109+v154<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v154*v138)
											v164 = v122 | v135
											*(*int64)(unsafe.Add(mBase, uint32(v109+v164<<(uint(v135)%32)))) = base.I64_extend_i32_u(v104 + v164*v138)
											v173 = int32(4)
											v174 = v122 + v173
											v176 = v127 + v173
											if v176 != v111&int32(2147483644) {
												v122 = v174
												v127 = v176
												continue
											} else {
												break
											}
											break
										}
										if v115 == int32(0) {
										} else {
											v181 = v174
											v196 = v181
											v201 = int32(0)
											for {
												*(*int64)(unsafe.Add(mBase, uint32(v109+v196<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v104 + v196*int32(6))
												v217 = int32(1)
												v220 = v201 + v217
												if v220 != v115 {
													v196 = v196 + v217
													v201 = v220
													continue
												} else {
													break
												}
												break
											}
										}
									} else {
										v181 = v93
										v196 = v181
										v201 = int32(0)
										for {
											*(*int64)(unsafe.Add(mBase, uint32(v109+v196<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v104 + v196*int32(6))
											v217 = int32(1)
											v220 = v201 + v217
											if v220 != v115 {
												v196 = v196 + v217
												v201 = v220
												continue
											} else {
												break
											}
											break
										}
									}
								}
								v237 = F_construct_array_builtin(m, v109, v111, int32(27))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = base.I64_extend_i32_u(v237)
									F_pfree(m, v109)
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v104)
										mBase = m.M
										v244 = m.ExcPending
										if v244 != 0 {
											return int64(0)
										} else {
											v245 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
											v250 = F_heap_form_tuple(m, v245, v15+int32(-32), v15+int32(-36))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int64(0)
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
												v253 = F_HeapTupleHeaderGetDatum(m, v252)
												mBase = m.M
												v254 = m.ExcPending
												if v254 != 0 {
													return int64(0)
												} else {
													v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
													v256 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v90 + (v255+v256)&int32(_a_F_gin_leafpage_items_9) + int32(8)
													v264 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
													*(*int64)(unsafe.Add(mBase, uint32(v88))) = v264 + int64(1)
													v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v268)+20)) = v256
													v291 = v253
													m.G0 = v17 - int32(-64)
													return v291
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int64(0)
						} else {
							v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v273)+20)) = int32(2)
							v276 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v276)
							v291 = v14
							m.G0 = v17 - int32(-64)
							return v291
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v299 = m.ExcPending
				if v299 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v302 = m.ExcPending
					if v302 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_gin_leafpage_items_10), int32(0))
						mBase = m.M
						v306 = m.ExcPending
						if v306 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_leafpage_items_4), int32(188), int32(_a_F_gin_leafpage_items_5))
							mBase = m.M
							v311 = m.ExcPending
							if v311 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gin_numeric_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v12 int64
	_ = v12
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = base.I32_wrap_i64(v5)
	v7 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	if v7 == int64(0) {
		if v6 != 0 {
			v12 = int64(-1)
		} else {
			v12 = int64(0)
		}
		return v12
	} else {
		if v6 == int32(0) {
			return int64(1)
		} else {
			v22 = F_DirectFunctionCall2Coll(m, int32(1467), int32(0), v7, v5&int64(4294967295))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				return base.I64_extend32_s(v22)
			}
		}
	}
}
func F_gin_page_opaque_info(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v203 int64
	_ = v203
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v7 = m.G0
	v9 = v7 - int32(192)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = F_superuser(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L65
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L60
	}
L5:
	;
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = F_get_page_from_raw(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L56
	}
L9:
	;
	m.G0 = v9 + int32(192)
	return v203
L10:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+14)))
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
	v203 = int64(0)
	goto L9
L12:
	;
	goto L13
L13:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)))
	v27 = int32(8)
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	if (v26<<(uint(v27)%32)-v29)&int32(_a_F_gin_page_opaque_info_0) != v27 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v38 = F_get_call_result_type(m, l0, int32(0), v9+int32(188))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v38 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v43 = v9 + int32(16)
	v44 = v18 + v29
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+6)))
	if v45&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v51 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_6))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v56 = v43
	v57 = int32(0)
	goto L19
L19:
	;
	if v45&int32(2) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = base.I64_extend_i32_u(v51)
	v56 = v43 | int32(8)
	v57 = int32(1)
	goto L19
L21:
	;
	v61 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_7))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v67 = v57
	goto L23
L23:
	;
	if v45&int32(4) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = base.I64_extend_i32_u(v61)
	v67 = v57 + int32(1)
	goto L23
L25:
	;
	v76 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_8))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v82 = v67
	goto L27
L27:
	;
	if v45&int32(8) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v67<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v76)
	v82 = v67 + int32(1)
	goto L27
L29:
	;
	v91 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_9))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v97 = v82
	goto L31
L31:
	;
	if v45&int32(16) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v82<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v91)
	v97 = v82 + int32(1)
	goto L31
L33:
	;
	v106 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_10))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v112 = v97
	goto L35
L35:
	;
	if v45&int32(32) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v97<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v106)
	v112 = v97 + int32(1)
	goto L35
L37:
	;
	v121 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_11))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v127 = v112
	goto L39
L39:
	;
	if v45&int32(64) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v112<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v121)
	v127 = v112 + int32(1)
	goto L39
L41:
	;
	v136 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_12))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v142 = v127
	goto L43
L43:
	;
	if v45&int32(128) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v127<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v136)
	v142 = v127 + int32(1)
	goto L43
L45:
	;
	v151 = F_cstring_to_text(m, int32(_a_F_gin_page_opaque_info_13))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v157 = v142
	goto L47
L47:
	;
	v159 = v45 & int32(_a_F_gin_page_opaque_info_14)
	if v159 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v142<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v151)
	v157 = v142 + int32(1)
	goto L47
L49:
	;
	v168 = F_DirectFunctionCall1Coll(m, int32(3109), int32(0), base.I64_extend_i32_u(v159))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v173 = v157
	goto L51
L51:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+158)) = uint8(v174)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+156)) = uint16(v174)
	v178 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44))))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v178
	v180 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = v180
	v185 = F_construct_array_builtin(m, v9+int32(16), v173, int32(25))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)+v157<<(uint(int32(3))%32)))) = v168
	v173 = v157 + int32(1)
	goto L51
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = base.I64_extend_i32_u(v185)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
	v194 = F_heap_form_tuple(m, v189, v9+int32(160), v9+int32(156))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	v197 = F_HeapTupleHeaderGetDatum(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v203 = v197
	goto L9
L56:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_gin_page_opaque_info_15), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_gin_page_opaque_info_3), int32(112), int32(_a_F_gin_page_opaque_info_4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_gin_page_opaque_info_1), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)))
	v237 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = (v236<<(uint(v237)%32) - v235) & int32(_a_F_gin_page_opaque_info_0)
	v246 = F_errdetail(m, int32(_a_F_gin_page_opaque_info_2), v9)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_gin_page_opaque_info_3), int32(125), int32(_a_F_gin_page_opaque_info_4))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_gin_page_opaque_info_5), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_gin_page_opaque_info_3), int32(131), int32(_a_F_gin_page_opaque_info_4))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
