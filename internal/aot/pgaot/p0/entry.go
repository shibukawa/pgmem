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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
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
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v36 == int32(1) {
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
	return v189
L6:
	;
	v39 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v42) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v60) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v52 = int32(base.Ui32(v42+int32(_a_F_entryLocateEntry_0))>>(uint(int32(2))%32)) & int32(_a_F_entryLocateEntry_1)
	goto L11
L10:
	;
	v52 = int32(0)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v41 * v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v56 = m.T0[v55].(func(*base.Module, int32, int32) int32)(m, l0, v35)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v189 = v56
	goto L5
L14:
	;
	v68 = int32(base.Ui32(v60+int32(_a_F_entryLocateEntry_0)) >> (uint(int32(2)) % 32))
	goto L16
L15:
	;
	v68 = int32(0)
	goto L16
L16:
	;
	v70 = v68 + int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v70&int32(_a_F_entryLocateEntry_1)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v81 = v70
	v83 = int32(1)
	goto L20
L18:
	;
	v153 = v70
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v153)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v35+v153&int32(_a_F_entryLocateEntry_1)<<(uint(int32(2))%32))+20))
	v171 = v35 + v168&int32(_a_F_entryLocateEntry_2)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171))))
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+2)))
	v189 = v172<<(uint(int32(16))%32) | v175
	goto L5
L20:
	;
	v95 = int32(base.Ui32((v81-v83)&int32(_a_F_entryLocateEntry_3))>>(uint(int32(1))%32)) + v83
	v96 = int32(_a_F_entryLocateEntry_1)
	v97 = v95 & v96
	if v97 == v68&v96 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v153 = v141
	goto L19
L22:
	;
	v140 = base.B2i32(int32(0) < v137)
	if int32(0) < v137 {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v101 = int32(-1)
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v35+v102)))
	if v104 == v101 {
		v137 = v101
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(20)+v97<<(uint(int32(2))%32))))
	v115 = v35 + v112&int32(_a_F_entryLocateEntry_2)
	v116 = F_gintuple_get_attrnum(m, v108, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v121 = F_gintuple_get_key(m, v118, v115, v15+int32(15))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+60)))
	v127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+15)))
	v128 = F_ginCompareAttEntries(m, v123, v124, v125, v126, v116, v121, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	if v128 != 0 {
		v137 = v128
		goto L22
	} else {
		goto L30
	}
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v95)
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+2)))
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v189 = v131 | v132<<(uint(int32(16))%32)
	goto L5
L31:
	;
	v141 = v81
	goto L33
L32:
	;
	v141 = v95
	goto L33
L33:
	;
	if int32(0) < v137 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v146 = v95 + int32(1)
	goto L36
L35:
	;
	v146 = v83
	goto L36
L36:
	;
	if base.Ui32(v146&int32(_a_F_entryLocateEntry_1)) < base.Ui32(v141&int32(_a_F_entryLocateEntry_1)) {
		v81 = v141
		v83 = v146
		goto L20
	} else {
		goto L37
	}
L37:
	;
	goto L21
}
