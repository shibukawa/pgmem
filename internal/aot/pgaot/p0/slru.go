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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
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
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L67
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
	v231 = int32(0)
	goto L1
L8:
	;
	v32 = v25 + int32(19)
	if v32&int32(3) == int32(0) {
		v56 = v32
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v90 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v89 = v81 - v32
	goto L10
L12:
	;
	v60 = v56
	goto L21
L13:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v89 = int32(0)
	goto L10
L15:
	;
	goto L16
L16:
	;
	v45 = v32
	goto L17
L17:
	;
	v49 = v45 + int32(1)
	if v49&int32(3) == int32(0) {
		v56 = v49
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v81 = v49
	goto L11
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v54 != 0 {
		v45 = v49
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v69 = int32(-2139062144)
	if (int32(16843008)-v66|v66)&v69 == v69 {
		v60 = v60 + int32(4)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v75 = v60
	goto L24
L23:
	;
	goto L22
L24:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		v75 = v75 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v81 = v75
	goto L11
L26:
	;
	goto L25
L27:
	;
	v210 = F_ReadDir(m, v16, v15)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L65
	}
L28:
	;
	v99 = int32(505149)
	v103 = m.G0
	v105 = v103 - int32(32)
	v106 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v106
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _consts[140])))
	if v114 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v89 == int32(15) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(2)) < base.Ui32(v89-int32(4)) {
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L27
L33:
	;
	goto L28
L34:
	;
	if v182 != v89 {
		goto L27
	} else {
		goto L55
	}
L35:
	;
	v182 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _consts[141])))
	if v118 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v122 = v32
	goto L41
L39:
	;
	goto L40
L40:
	;
	v132 = v99
	v133 = v114
	goto L44
L41:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v128 == v114 {
		v122 = v122 + int32(1)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v182 = v122 - v32
	goto L34
L43:
	;
	goto L42
L44:
	;
	v140 = v105 + int32(base.Ui32(v133)>>(uint(int32(3))%32))&int32(28)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v142 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v141 | v142<<(uint(v133)%32)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v146 != 0 {
		v132 = v132 + v142
		v133 = v146
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v149 == int32(0) {
		v174 = v32
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v182 = v174 - v32
	goto L34
L48:
	;
	v153 = v32
	v154 = v149
	goto L49
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v105+int32(base.Ui32(v154)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v162)>>(uint(v154)%32))&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v174 = v170
	goto L47
L51:
	;
	v174 = v153
	goto L47
L52:
	;
	goto L53
L53:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	v170 = v153 + int32(1)
	if v168 != 0 {
		v153 = v170
		v154 = v168
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v187 = F_strtox_2(m, v32, int32(0), int32(16), int64(-9223372036854775807-1))
	mBase = m.M
	goto L56
L56:
	;
	v190 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v190 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	F_errmsg_internal(m, int32(165163), v12)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v204 = m.T0[l1].(func(*base.Module, int32, int32, int64, int32) int32)(m, l0, v32, v187<<(uint(int64(5))%64), l2)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L63
	}
L61:
	;
	F_errfinish(m, int32(462366), int32(1813), int32(12489))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	if v204 == int32(0) {
		goto L27
	} else {
		goto L64
	}
L64:
	;
	v231 = int32(1)
	goto L1
L65:
	;
	if v210 != 0 {
		v25 = v210
		goto L8
	} else {
		goto L66
	}
L66:
	;
	goto L9
L67:
	;
	m.G0 = v12 + int32(16)
	return v231
}
