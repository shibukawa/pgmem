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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_reserveAllocatedDesc(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[0]))
	if v16 <= int32(0) {
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
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v58 = F_fopen(m, l0, l1)
	mBase = m.M
	if v58 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[1]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[2]))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	if v22+(v24+v16) < v20 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[4]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	F_LruDelete(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[0]))
	if v40 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[1]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[2]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	if v44 <= v46+(v48+v40) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	m.G0 = v9 + int32(16)
	return v136
L15:
	;
	goto L18
L16:
	;
	v108 = v58
	goto L17
L17:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[5]))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	v116 = v111 + v113*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[6]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	goto L33
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[7]))
	switch v68 - int32(33) {
	case 0, 8:
		goto L20
	default:
		v136 = v3
		goto L14
	}
L19:
	;
	v108 = v101
	goto L17
L20:
	;
	v73 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[7])) = v88
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[0]))
	if v91 <= v88 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	F_errmsg(m, int32(_a_F_AllocateFile_0), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_AllocateFile_1), int32(2684), int32(_a_F_AllocateFile_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[7])) = v68
	v136 = v3
	goto L14
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[4]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	F_LruDelete(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v101 = F_fopen(m, l0, l1)
	mBase = m.M
	if v101 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v122
	v124 = int32(_a_F_AllocateFile_3)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[3])) = v126 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v136 = v130
	goto L14
L34:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateFile[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v150
	F_errmsg(m, int32(_a_F_AllocateFile_4), v9)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_AllocateFile_1), int32(2661), int32(_a_F_AllocateFile_2))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
