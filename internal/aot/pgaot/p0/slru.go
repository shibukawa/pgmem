package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruScanDirCbReportPresence(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = m.T0[v7].(func(*base.Module, int64, int64) int32)(m, l2, v6)
	mBase = m.M
	if v8 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = m.T0[v11].(func(*base.Module, int64, int64) int32)(m, l2+int64(31), v6)
		mBase = m.M
		v14 = v12
	} else {
		v14 = int32(0)
	}
	return v14
}
func F_SlruScanDirectory(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v132 int64
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	v15 = l0 + v11
	v16 = F_AllocateDir(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_FreeDir(m, v16)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L46
	}
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_ReadDir(m, v16, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = v20
	goto L8
L6:
	;
	goto L7
L7:
	;
	v176 = int32(0)
	goto L1
L8:
	;
	v32 = v25 + int32(19)
	v33 = F_strlen(m, v32)
	mBase = m.M
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v155 = F_ReadDir(m, v16, v15)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L44
	}
L11:
	;
	v41 = base.B2i32(v33 == int32(15))
	goto L13
L12:
	;
	v41 = base.B2i32(base.Ui32(v33-int32(4)) < base.Ui32(int32(3)))
	goto L13
L13:
	;
	if v41 != int32(1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v44 = int32(_a_F_SlruScanDirectory_0)
	v48 = m.G0
	v50 = v48 - int32(32)
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+24)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
	v59 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruScanDirectory[0])))
	if v59 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v127 != v33 {
		goto L10
	} else {
		goto L34
	}
L16:
	;
	v127 = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruScanDirectory[1])))
	if v63 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = v32
	goto L22
L20:
	;
	goto L21
L21:
	;
	v77 = v44
	v78 = v59
	goto L25
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v73 == v59 {
		v67 = v67 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v127 = v67 - v32
	goto L15
L24:
	;
	goto L23
L25:
	;
	v85 = v50 + int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v86 | v87<<(uint(v78)%32)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v91 != 0 {
		v77 = v77 + v87
		v78 = v91
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v94 == int32(0) {
		v117 = v32
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v127 = v117 - v32
	goto L15
L29:
	;
	v98 = v32
	v99 = v94
	goto L30
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v50+int32(base.Ui32(v99)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v107)>>(uint(v99)%32))&int32(1) == int32(0) {
		v117 = v98
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v117 = v115
	goto L28
L32:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v115 = v98 + int32(1)
	if v113 != 0 {
		v98 = v115
		v99 = v113
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v132 = F_strtox_2(m, v32, int32(0), int32(16), int64(-9223372036854775807-1))
	mBase = m.M
	goto L35
L35:
	;
	v137 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v137 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	F_errmsg_internal(m, int32(_a_F_SlruScanDirectory_1), v12)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v149 = m.T0[l1].(func(*base.Module, int32, int32, int64, int32) int32)(m, l0, v32, v132<<(uint(int64(5))%64), l2)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	F_errfinish(m, int32(_a_F_SlruScanDirectory_2), int32(1813), int32(_a_F_SlruScanDirectory_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v149 == int32(0) {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v176 = int32(1)
	goto L1
L44:
	;
	if v155 != 0 {
		v25 = v155
		goto L8
	} else {
		goto L45
	}
L45:
	;
	goto L9
L46:
	;
	m.G0 = v12 + int32(16)
	return v176
}
