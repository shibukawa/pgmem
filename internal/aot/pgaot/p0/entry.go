package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryLocateEntry(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v37 == v36 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_entryLocateEntry[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v17^int32(-1))<<(uint(int32(2))%32))))
	v35 = v27
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_entryLocateEntry[1]))
	v35 = v29 + v17<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	m.G0 = v15 + int32(16)
	return v188
L6:
	;
	v40 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v43) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
	if base.Ui32(v61) < base.Ui32(int32(25)) {
		v153 = v36
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v53 = int32(base.Ui32(v43+int32(_a_F_entryLocateEntry_0))>>(uint(int32(2))%32)) & int32(_a_F_entryLocateEntry_1)
	goto L11
L10:
	;
	v53 = int32(0)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v42 * v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = m.T0[v56].(func(*base.Module, int32, int32) int32)(m, l0, v35)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v188 = v57
	goto L5
L14:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v153)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v153&int32(_a_F_entryLocateEntry_1)<<(uint(int32(2))%32)+v35)+20))
	v170 = v35 + v167&int32(_a_F_entryLocateEntry_2)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170))))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+2)))
	v188 = v171<<(uint(int32(16))%32) | v174
	goto L5
L15:
	;
	v67 = int32(base.Ui32(v61+int32(_a_F_entryLocateEntry_0)) >> (uint(int32(2)) % 32))
	v69 = v67 & int32(_a_F_entryLocateEntry_1)
	if v69 == int32(0) {
		v153 = v36
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v74 = int32(1)
	v79 = v74
	v81 = v67 + v74
	goto L17
L17:
	;
	v94 = int32(base.Ui32((v81-v79)&int32(_a_F_entryLocateEntry_3))>>(uint(int32(1))%32)) + v79
	v96 = v94 & int32(_a_F_entryLocateEntry_1)
	if v69 == v96 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v153 = v140
	goto L14
L19:
	;
	v139 = base.B2i32(int32(0) < v135)
	if int32(0) < v135 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v98 = int32(-1)
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v35+v99)))
	if v101 == v98 {
		v135 = v98
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v96<<(uint(int32(2))%32)+(v35+int32(24))-int32(4))))
	v114 = v35 + v111&int32(_a_F_entryLocateEntry_2)
	v115 = F_gintuple_get_attrnum(m, v105, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L12
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v120 = F_gintuple_get_key(m, v117, v114, v15+int32(15))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+60)))
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+15)))
	v127 = F_ginCompareAttEntries(m, v122, v123, v124, v125, v115, v120, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	if v127 != 0 {
		v135 = v127
		goto L19
	} else {
		goto L27
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v94)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+2)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v188 = v130 | v131<<(uint(int32(16))%32)
	goto L5
L28:
	;
	v140 = v81
	goto L30
L29:
	;
	v140 = v94
	goto L30
L30:
	;
	if int32(0) < v135 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v145 = v94 + int32(1)
	goto L33
L32:
	;
	v145 = v79
	goto L33
L33:
	;
	if base.Ui32(v145&int32(_a_F_entryLocateEntry_1)) < base.Ui32(v140&int32(_a_F_entryLocateEntry_1)) {
		v79 = v145
		v81 = v140
		goto L17
	} else {
		goto L34
	}
L34:
	;
	goto L18
}
