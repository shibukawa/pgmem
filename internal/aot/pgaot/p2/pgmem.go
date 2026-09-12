package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgmem_dlsym(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
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
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(288)
	m.G0 = v9
	if l0 == v3 {
		v146 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(288)
	return v146
L2:
	;
	if l1 == int32(0) {
		v146 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(468881)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1295])))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v83 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L5:
	;
	if v39-v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = l1
	v24 = v15
	goto L9
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v38 = v27
	v39 = v28
	goto L6
L11:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v41 = int32(94505)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1296])))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v45 == int32(0) {
		v64 = v44
		v65 = v45
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	v76 = F_snprintf(m, v9+int32(32), int32(256), int32(166393), v9+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	if v65-v64 != 0 {
		v82 = l1
		goto L4
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v44 != v45 {
		v64 = v44
		v65 = v45
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v49 = l1
	v50 = v41
	goto L20
L20:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v64 = v53
		v65 = v54
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v64 = v53
	v65 = v54
	goto L17
L22:
	;
	v57 = int32(1)
	if v53 == v54 {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L15
L25:
	;
	return int32(0)
L26:
	;
	v82 = v9 + int32(32)
	goto L4
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v82
	v137 = F_snprintf(m, int32(4536528), int32(512), int32(192585), v9)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L25
	} else {
		goto L43
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = int32(0)
	goto L29
L29:
	;
	v96 = v86 + v88<<(uint(int32(3))%32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v101 == int32(0) {
		v120 = v100
		v121 = v101
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v146 = v126
	goto L1
L31:
	;
	if v121-v120 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	goto L31
L33:
	;
	if v100 != v101 {
		v120 = v100
		v121 = v101
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v105 = v97
	v106 = v82
	goto L35
L35:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	if v110 == int32(0) {
		v120 = v109
		v121 = v110
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v120 = v109
	v121 = v110
	goto L32
L37:
	;
	v113 = int32(1)
	if v109 == v110 {
		v105 = v105 + v113
		v106 = v106 + v113
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v124 = v88 + int32(1)
	if v83 != v124 {
		v88 = v124
		goto L29
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L30
L42:
	;
	goto L27
L43:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1297])) = uint8(v141)
	v146 = int32(0)
	goto L1
}
func F_pgmem_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	v4 = int32(0)
	if l1 == v4 {
	} else {
		v12 = l1 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(l1) {
			v20 = v4
			v23 = v4
			for {
				v27 = l0 + v20<<(uint(int32(3))%32)
				v28 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+14)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+22)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+30)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+38)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+46)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+54)) = uint16(v28)
				*(*uint16)(unsafe.Add(mBase, uint32(v27)+62)) = uint16(v28)
				v44 = int32(8)
				v45 = v20 + v44
				v47 = v23 + v44
				if v47 != l1&int32(-8) {
					v20 = v45
					v23 = v47
					continue
				} else {
					break
				}
				break
			}
			v52 = v45
		} else {
			v52 = v4
		}
		if v12 == int32(0) {
		} else {
			v62 = v52
			v64 = v4
			for {
				v70 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0+v62<<(uint(int32(3))%32))+6)) = uint16(v70)
				v72 = int32(1)
				v75 = v64 + v72
				if v75 != v12 {
					v62 = v62 + v72
					v64 = v75
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v85 = m.Env.Pgmem_poll(m, l2)
	mBase = m.M
	return int32(0)
}
