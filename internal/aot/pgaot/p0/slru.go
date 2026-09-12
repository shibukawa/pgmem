package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruScanDirCbReportPresence(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = m.T0[v8].(func(*base.Module, int64, int64) int32)(m, l2, v7)
	mBase = m.M
	if v9 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = m.T0[v12].(func(*base.Module, int64, int64) int32)(m, l2+int64(31), v7)
		mBase = m.M
		v14 = v13
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
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
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
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L50
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
	v175 = int32(0)
	goto L1
L8:
	;
	v32 = v25 + int32(19)
	v33 = F_strlen(m, v32)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v34 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v154 = F_ReadDir(m, v16, v15)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L48
	}
L11:
	;
	v43 = int32(_a_F_SlruScanDirectory_0)
	v47 = m.G0
	v49 = v47 - int32(32)
	v50 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = v50
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruScanDirectory[0])))
	if v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	if v33 == int32(15) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(int32(2)) < base.Ui32(v33-int32(4)) {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L10
L16:
	;
	goto L11
L17:
	;
	if v126 != v33 {
		goto L10
	} else {
		goto L38
	}
L18:
	;
	v126 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruScanDirectory[1])))
	if v62 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v66 = v32
	goto L24
L22:
	;
	goto L23
L23:
	;
	v76 = v43
	v77 = v58
	goto L27
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v72 == v58 {
		v66 = v66 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v126 = v66 - v32
	goto L17
L26:
	;
	goto L25
L27:
	;
	v84 = v49 + int32(base.Ui32(v77)>>(uint(int32(3))%32))&int32(28)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v85 | v86<<(uint(v77)%32)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v90 != 0 {
		v76 = v76 + v86
		v77 = v90
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v93 == int32(0) {
		v118 = v32
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v126 = v118 - v32
	goto L17
L31:
	;
	v97 = v32
	v98 = v93
	goto L32
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(base.Ui32(v98)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v106)>>(uint(v98)%32))&int32(1) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v118 = v114
	goto L30
L34:
	;
	v118 = v97
	goto L30
L35:
	;
	goto L36
L36:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v114 = v97 + int32(1)
	if v112 != 0 {
		v97 = v114
		v98 = v112
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v131 = F_strtox_2(m, v32, int32(0), int32(16), int64(-9223372036854775807-1))
	mBase = m.M
	goto L39
L39:
	;
	v134 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v134 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	F_errmsg_internal(m, int32(_a_F_SlruScanDirectory_1), v12)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v148 = m.T0[l1].(func(*base.Module, int32, int32, int64, int32) int32)(m, l0, v32, v131<<(uint(int64(5))%64), l2)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	F_errfinish(m, int32(_a_F_SlruScanDirectory_2), int32(1813), int32(_a_F_SlruScanDirectory_3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v148 == int32(0) {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v175 = int32(1)
	goto L1
L48:
	;
	if v154 != 0 {
		v25 = v154
		goto L8
	} else {
		goto L49
	}
L49:
	;
	goto L9
L50:
	;
	m.G0 = v12 + int32(16)
	return v175
}
