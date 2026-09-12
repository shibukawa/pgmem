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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	v9 = l3 - int32(1)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v12) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v124 & int32(65535)
L2:
	;
	if v70&int32(65535) != 0 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v20 = int32(base.Ui32(v12+int32(262120)) >> (uint(int32(2)) % 32))
	goto L5
L4:
	;
	v20 = int32(0)
	goto L5
L5:
	;
	if base.Ui32(v20&int32(65535)) <= base.Ui32(v9&int32(65535)) {
		v70 = v20
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = l1 + int32(24)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+l3<<(uint(int32(2))%32)-int32(4))))
	v34 = l1 + v31&int32(32767)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34))))
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+2)))
	if v35<<(uint(int32(16))%32)|v38 == l2 {
		v124 = l3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v44 = l3
	goto L8
L8:
	;
	v49 = v44 + int32(1)
	v50 = int32(65535)
	v51 = v49 & v50
	if base.Ui32(v20&v50) < base.Ui32(v51) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v124 = v49
	goto L1
L10:
	;
	v70 = v9
	goto L2
L11:
	;
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32)+v25-int32(4))))
	v63 = l1 + v60&int32(32767)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)))
	if v64<<(uint(int32(16))%32)|v67 != l2 {
		v44 = v49
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	v85 = int32(1)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v124 = int32(0)
	goto L1
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85&int32(65535)<<(uint(int32(2))%32)+(l1+int32(24))-int32(4))))
	v99 = l1 + v96&int32(32767)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+2)))
	if v100<<(uint(int32(16))%32)|v103 == l2 {
		v124 = v85
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v107 = v85 + int32(1)
	v108 = int32(65535)
	if base.Ui32(v107&v108) <= base.Ui32(v70&v108) {
		v85 = v107
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
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	if l1 < int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9+(l1^int32(-1))<<(uint(int32(2))%32))))
		v23 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		v23 = v17 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	if l1 < int32(0) {
		v27 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(l1^int32(-1))<<(uint(int32(6))%32))+16))
		v42 = v33
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v35+l1<<(uint(int32(6))%32)+int32(-64))+16))
		v42 = v41
	}
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)))
	if base.Ui32(v44) < base.Ui32(int32(25)) {
		v55 = int32(-1)
	} else {
		v55 = int32(base.Ui32(v44+int32(262120))>>(uint(int32(2))%32))&int32(65535) - int32(1)
	}
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v23+v55<<(uint(int32(2))%32))+24))
	v62 = v23 + v59&int32(32767)
	v64 = F_palloc(m, int32(8))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		return int32(0)
	} else {
		v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)))
		v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v68)+6)))
		if v70&int32(2) == int32(0) {
			v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
			v101 = F_palloc(m, v98&int32(8191))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
				v105 = v103 & int32(8191)
				if v105 != 0 {
					v106 = F__emscripten_memcpy_bulkmem(m, v101, v62, v105)
					mBase = m.M
				} else {
				}
				v108 = v101
				v111 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)) = uint16(v111)
				*(*uint16)(unsafe.Add(mBase, uint32(v108)+2)) = uint16(v42)
				v115 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v108))) = uint16(v115)
				*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)) = uint8(v111)
				*(*int32)(unsafe.Add(mBase, uint32(v64))) = v108
				return v64
			}
		} else {
			v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)))
			if v75 == int32(65535) {
				v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
				v101 = F_palloc(m, v98&int32(8191))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
					v105 = v103 & int32(8191)
					if v105 != 0 {
						v106 = F__emscripten_memcpy_bulkmem(m, v101, v62, v105)
						mBase = m.M
					} else {
					}
					v108 = v101
					v111 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)) = uint16(v111)
					*(*uint16)(unsafe.Add(mBase, uint32(v108)+2)) = uint16(v42)
					v115 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v108))) = uint16(v115)
					*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)) = uint8(v111)
					*(*int32)(unsafe.Add(mBase, uint32(v64))) = v108
					return v64
				}
			} else {
				v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+2)))
				v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
				v88 = (v78 | v79<<(uint(int32(16))%32)&int32(2147418112) + int32(7)) & int32(-8)
				v89 = F_palloc(m, v88)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					if v88 != 0 {
						v91 = F__emscripten_memcpy_bulkmem(m, v89, v62, v88)
						mBase = m.M
						v92 = v91
					} else {
						v92 = v89
					}
					v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)))
					v96 = v93&int32(57344) | v88
					*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)) = uint16(v96)
					v108 = v89
					v111 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)) = uint16(v111)
					*(*uint16)(unsafe.Add(mBase, uint32(v108)+2)) = uint16(v42)
					v115 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v108))) = uint16(v115)
					*(*uint8)(unsafe.Add(mBase, uint32(v64)+4)) = uint8(v111)
					*(*int32)(unsafe.Add(mBase, uint32(v64))) = v108
					return v64
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
