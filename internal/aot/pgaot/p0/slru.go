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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = m.T0[v7].(func(*base.Module, int64, int64) int32)(m, l2, v6)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v15 = m.T0[v14].(func(*base.Module, int64, int64) int32)(m, l2+int64(31), v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = v15
				return v18
			}
		} else {
			v18 = int32(0)
			return v18
		}
	}
}
func F_SlruScanDirectory(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_AllocateDir(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_FreeDir(m, v14)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = F_ReadDir(m, v14, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v19
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
	v30 = v24 + int32(19)
	v31 = F_strlen(m, v30)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v157 = F_ReadDir(m, v14, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L44
	}
L11:
	;
	v39 = base.B2i32(v31 == int32(15))
	goto L13
L12:
	;
	v39 = base.B2i32(base.Ui32(v31-int32(4)) < base.Ui32(int32(3)))
	goto L13
L13:
	;
	if v39 != int32(1) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v42 = int32(_a_F_SlruScanDirectory_0)
	v46 = m.G0
	v48 = v46 - int32(32)
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+24)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = v49
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruScanDirectory[0])))
	if v57 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v125 != v31 {
		goto L10
	} else {
		goto L34
	}
L16:
	;
	v125 = int32(0)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SlruScanDirectory[1])))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v65 = v30
	goto L22
L20:
	;
	goto L21
L21:
	;
	v75 = v42
	v76 = v57
	goto L25
L22:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v71 == v57 {
		v65 = v65 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v125 = v65 - v30
	goto L15
L24:
	;
	goto L23
L25:
	;
	v83 = v48 + int32(base.Ui32(v76)>>(uint(int32(3))%32))&int32(28)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v84 | v85<<(uint(v76)%32)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v89 != 0 {
		v75 = v75 + v85
		v76 = v89
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v92 == int32(0) {
		v115 = v30
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v125 = v115 - v30
	goto L15
L29:
	;
	v96 = v30
	v97 = v92
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(base.Ui32(v97)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v105)>>(uint(v97)%32))&int32(1) == int32(0) {
		v115 = v96
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v115 = v113
	goto L28
L32:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	v113 = v96 + int32(1)
	if v111 != 0 {
		v96 = v113
		v97 = v111
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v130 = F_strtox_2(m, v30, int32(0), int32(16), int64(-9223372036854775807-1))
	mBase = m.M
	goto L35
L35:
	;
	v135 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v135 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v137
	F_errmsg_internal(m, int32(_a_F_SlruScanDirectory_1), v11)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v149 = m.T0[l1].(func(*base.Module, int32, int32, int64, int32) int32)(m, l0, v30, v130<<(uint(int64(5))%64), l2)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	F_errfinish(m, int32(_a_F_SlruScanDirectory_2), int32(1866), int32(_a_F_SlruScanDirectory_3))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
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
	if v157 != 0 {
		v24 = v157
		goto L8
	} else {
		goto L45
	}
L45:
	;
	goto L9
L46:
	;
	m.G0 = v11 + int32(16)
	return v176
}
