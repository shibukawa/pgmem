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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(1040)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v17 == v5 {
		v183 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(1040)
	return v183
L2:
	;
	v21 = l1
	v24 = v17
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32 == int32(0) {
		v183 = v5
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v183 = v5
	goto L1
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v35 == int32(0) {
		v183 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = v24 + int32(4)
	v47 = v39
	v50 = v39 + v35<<(uint(int32(3))%32)
	goto L7
L7:
	;
	v60 = v47 + (v50-v47)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v63 = v61 & int32(255)
	if v32 == v63 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v171 != 0 {
		v21 = v21 + int32(1)
		v24 = v171
		goto L3
	} else {
		goto L44
	}
L9:
	;
	goto L8
L10:
	;
	if v61&int32(256) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v153 = base.B2i32(base.Ui32(v63) < base.Ui32(v32))
	if base.Ui32(v63) < base.Ui32(v32) {
		goto L37
	} else {
		goto L38
	}
L13:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v69 != 0 {
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
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v81 != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	if v61&int32(512) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if l3&int32(base.Ui32(v61)>>(uint(int32(9))%32)) == int32(0) {
		v183 = v5
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v183 = v5
	goto L1
L20:
	;
	goto L15
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82+int32(base.Ui32(v61)>>(uint(int32(11))%32))&int32(_a_F_FindWord_0))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+1036)) = v88
	goto L24
L22:
	;
	goto L23
L23:
	;
	v183 = int32(1)
	goto L1
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1036))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v103 == int32(0) {
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
	v111 = m.ExcPending
	if v111 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v114 == int32(0))|base.B2i32(v114 != v117) != 0 {
		v135 = v114
		v136 = v117
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v135-v136 != 0 {
		goto L24
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v120 = v15
	v121 = l2
	goto L32
L32:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	if v125 == int32(0) {
		v135 = v125
		v136 = v124
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v135 = v125
	v136 = v124
	goto L30
L34:
	;
	v128 = int32(1)
	if v125 == v124 {
		v120 = v120 + v128
		v121 = v121 + v128
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L25
L37:
	;
	v154 = v60 + int32(8)
	goto L39
L38:
	;
	v154 = v47
	goto L39
L39:
	;
	if base.Ui32(v63) < base.Ui32(v32) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v155 = v50
	goto L42
L41:
	;
	v155 = v60
	goto L42
L42:
	;
	if base.Ui32(v154) < base.Ui32(v155) {
		v47 = v154
		v50 = v155
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v183 = v5
	goto L1
L44:
	;
	goto L4
}
