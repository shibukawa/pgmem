package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AllocateFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_reserveAllocatedDesc(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[0]))
	if v15 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v55 = F_fopen(m, l0, l1)
	mBase = m.M
	if v55 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[1]))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[2]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	if v21+(v23+v15) < v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[4]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[0]))
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[1]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[2]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	if v42 <= v44+(v46+v38) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	m.G0 = v8 + int32(16)
	return v130
L15:
	;
	goto L18
L16:
	;
	v102 = v55
	goto L17
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[5]))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	v111 = v106 + v108*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[6]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	goto L33
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[7]))
	switch v64 - int32(33) {
	case 0, 8:
		goto L20
	default:
		v130 = v3
		goto L14
	}
L19:
	;
	v102 = v97
	goto L17
L20:
	;
	v69 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v69 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[7])) = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[0]))
	if v87 <= v84 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	F_errmsg(m, int32(_a_F_AllocateFile_0), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_AllocateFile_1), int32(2684), int32(_a_F_AllocateFile_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[7])) = v64
	v130 = v3
	goto L14
L29:
	;
	goto L30
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[4]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	F_LruDelete(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v97 = F_fopen(m, l0, l1)
	mBase = m.M
	if v97 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v117
	v119 = int32(_a_F_AllocateFile_3)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3])) = v121 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v130 = v125
	goto L14
L34:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v144
	F_errmsg(m, int32(_a_F_AllocateFile_4), v8)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_AllocateFile_1), int32(2661), int32(_a_F_AllocateFile_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FileSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_FileSize[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4+l0*int32(48))))
	if v8 == int32(-1) {
		v11 = F_FileAccess(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			if v11 < int32(0) {
				return int64(-1)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_FileSize[0]))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0*int32(48))))
				v25 = v24
				v28 = F___lseek(m, v25, int64(0), int32(2))
				mBase = m.M
				return v28
			}
		}
	} else {
		v25 = v8
		v28 = F___lseek(m, v25, int64(0), int32(2))
		mBase = m.M
		return v28
	}
}
