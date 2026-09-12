package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FindWord(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(1040)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v17 == v5 {
		v180 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(1040)
	return v180
L2:
	;
	v21 = l1
	v24 = v17
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32 == int32(0) {
		v180 = v5
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v180 = v5
	goto L1
L5:
	;
	v36 = v24 + int32(4)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v40 = v36 + v37<<(uint(int32(3))%32)
	if base.Ui32(v40) <= base.Ui32(v36) {
		v180 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = v40
	v48 = v36
	goto L7
L7:
	;
	v59 = v48 + (v46-v48)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v62 = v60 & int32(255)
	if v32 == v62 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v169 != 0 {
		v21 = v65
		v24 = v169
		goto L3
	} else {
		goto L45
	}
L9:
	;
	goto L8
L10:
	;
	v65 = v21 + int32(1)
	if v60&int32(256) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v153 = base.B2i32(base.Ui32(v62) < base.Ui32(v32))
	if base.Ui32(v62) < base.Ui32(v32) {
		goto L38
	} else {
		goto L39
	}
L13:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v70 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if l3 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v82 != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	if v60&int32(512) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if l3&int32(base.Ui32(v60)>>(uint(int32(9))%32)) == int32(0) {
		v180 = v5
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v180 = v5
	goto L1
L20:
	;
	goto L15
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(base.Ui32(v60)>>(uint(int32(11))%32))&int32(_a_F_FindWord_0))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1036)) = v89
	goto L24
L22:
	;
	goto L23
L23:
	;
	v180 = int32(1)
	goto L1
L24:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1036))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	F_getNextFlagFromString(m, l0, v15+int32(1036), v15)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v116 == int32(0) {
		v135 = v115
		v136 = v116
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v136-v135 != 0 {
		goto L24
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v115 != v116 {
		v135 = v115
		v136 = v116
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v120 = v15
	v121 = l2
	goto L33
L33:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v125 == int32(0) {
		v135 = v124
		v136 = v125
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v135 = v124
	v136 = v125
	goto L30
L35:
	;
	v128 = int32(1)
	if v124 == v125 {
		v120 = v120 + v128
		v121 = v121 + v128
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L25
L38:
	;
	v154 = v59 + int32(8)
	goto L40
L39:
	;
	v154 = v48
	goto L40
L40:
	;
	if base.Ui32(v62) < base.Ui32(v32) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v155 = v46
	goto L43
L42:
	;
	v155 = v59
	goto L43
L43:
	;
	if base.Ui32(v154) < base.Ui32(v155) {
		v46 = v155
		v48 = v154
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v180 = v5
	goto L1
L45:
	;
	goto L4
}
