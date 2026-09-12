package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferFreeChange(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v89 {
	case 0, 1, 2, 8:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	case 5:
		goto L30
	default:
		goto L28
	case 11:
		goto L29
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v11 {
	case 0, 1, 2, 8:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	case 5:
		goto L6
	default:
		v50 = int32(64)
		goto L4
	case 11:
		goto L5
	}
L3:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v55 = v50
	goto L3
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v50 = v44<<(uint(int32(2))%32) - int32(-64)
	goto L4
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v55 = (v37+v38)<<(uint(int32(2))%32) + int32(136)
	goto L3
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v55 = v31<<(uint(int32(4))%32) - int32(-64)
	goto L3
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v26 = F_strlen(m, v25)
	mBase = m.M
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v55 = v26 + v27 + int32(73)
	goto L3
L9:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v13 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v18 = v14 + int32(84)
	goto L12
L11:
	;
	v18 = int32(64)
	goto L12
L12:
	;
	if v12 == int32(0) {
		v50 = v18
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v55 = v18 + v21 + int32(20)
	goto L3
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+216)) = v64 - v55
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+40))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v68 - v55
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v58 != int32(7) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	goto L14
L21:
	;
	v71 = v67
	goto L23
L22:
	;
	v71 = v63
	goto L23
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+220)) = v72 - v55
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v77 = v63 + int32(204)
	F_pairingheap_remove(m, v75, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v63)+216))
	if v80 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v83, v77)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	F_pfree(m, l1)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L24
	} else {
		goto L65
	}
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v137 == int32(0) {
		goto L28
	} else {
		goto L63
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v123 == int32(0) {
		goto L28
	} else {
		goto L55
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v116 != 0 {
		goto L50
	} else {
		goto L51
	}
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v104 != 0 {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v90 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_pfree(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v95 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(0)
	goto L36
L38:
	;
	F_pfree(m, v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	F_pfree(m, l1)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	return
L41:
	;
	F_pfree(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L24
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v109 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L24
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
	F_pfree(m, l1)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L24
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	return
L50:
	;
	F_pfree(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	F_pfree(m, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L24
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	return
L55:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+30)))
	if v126 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
	F_pfree(m, l1)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L24
	} else {
		goto L62
	}
L57:
	;
	F_pfree(m, v123)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L24
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_SnapBuildSnapDecRefcount(m, v123)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L24
	} else {
		goto L61
	}
L60:
	;
	goto L56
L61:
	;
	goto L56
L62:
	;
	return
L63:
	;
	F_pfree(m, v137)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L24
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
	goto L28
L65:
	;
	return
}
func F_ReorderBufferQueueMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v34 = int32(-1)
	v35 = v9
	v36 = v9
	v37 = v9
	v38 = v9
	v39 = v9
	v40 = v9
	v41 = v9
	v42 = v22
	goto L2
L1:
	;
	m.G0 = v22 + int32(32)
	return
L2:
	;
	goto L4
L3:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L4:
	;
	if v34 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v226 = int32(m.ExcTag)
	v227 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v226 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L7:
	;
	v47 = v42 - int32(16)
	m.G0 = v47
	v50 = v47 - int32(160)
	m.G0 = v50
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v155 = v35
	v156 = v36
	v157 = v37
	v158 = v38
	v159 = v39
	v160 = v40
	v161 = v41
	v162 = v42
	goto L9
L9:
	;
	if v161 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L10:
	;
	v52 = int32(4489440)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
	v65 = F_MemoryContextAlloc(m, v57, int32(64))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		v224 = v50
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v68 = v65 + int32(8)
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65)+56)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65)+48)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65)+40)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
	v93 = F_pstrdup(m, l5)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		v224 = v50
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+24)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
	v103 = F_palloc(m, l6)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		v224 = v50
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+28)) = v103
	if l6 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
	F_ReorderBufferQueueChange(m, l0, l1, l3, v65, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		v224 = v50
		goto L6
	} else {
		goto L20
	}
L17:
	;
	v106 = F__emscripten_memcpy_bulkmem(m, v103, l7, l6)
	mBase = m.M
	goto L19
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v53
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
	v127 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l3)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		v224 = v50
		goto L6
	} else {
		goto L24
	}
L22:
	;
	v130 = v40
	v131 = int32(0)
	goto L23
L23:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
	*(*int32)(unsafe.Add(mBase, _consts[80])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[655])) = v132
	goto L25
L24:
	;
	v130 = v127
	v131 = v127
	goto L23
L25:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v148 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v22 + int32(4)
	goto L29
L27:
	;
	v155 = v50
	v156 = v47
	v157 = v148
	v158 = v146
	v159 = v131
	v160 = v130
	v161 = int32(0)
	v162 = v50
	goto L9
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v155
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v156
	v174 = int32(0)
	m.T0[v167].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32))(m, l0, v159, l3, v174, l5, l6, l7)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v156
	*(*int32)(unsafe.Add(mBase, _consts[80])) = v174
	*(*int32)(unsafe.Add(mBase, _consts[655])) = v174
	goto L33
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v158
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v156
	v203 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[80])) = v203
	*(*int32)(unsafe.Add(mBase, _consts[655])) = v203
	goto L34
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v158
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v157
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v156
	F_pg_re_throw(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		v224 = v162
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L5
L36:
	;
	v231 = int32(v227)
	m.G0 = v224
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v22+int32(4) == v238 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	m.ExcPending = 1
	goto L45
L38:
	;
	if v241 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v241 = v240
	goto L41
L40:
	;
	v241 = int32(0)
	goto L41
L41:
	;
	goto L38
L42:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v34 = v241
	v35 = v243
	v36 = v242
	v37 = v246
	v38 = v247
	v39 = v245
	v40 = v244
	v41 = v233
	v42 = v224
	goto L2
L43:
	;
	goto L44
L44:
	;
	F___wasm_longjmp(m, v234, v233)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferSkipPrepare(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v10 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = F_hash_search(m, v15, v7+int32(12), int32(0), v7+int32(11))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)))
			if v23 == int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v26
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v30
				if v31 == int32(0) {
				} else {
					v36 = v31
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38 | int32(128)
				}
			}
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		if l1 != v10 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = F_hash_search(m, v15, v7+int32(12), int32(0), v7+int32(11))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)))
				if v23 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v26
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v30
					if v31 == int32(0) {
					} else {
						v36 = v31
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38 | int32(128)
					}
				}
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v14 != 0 {
				v36 = v14
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38 | int32(128)
			} else {
			}
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_ReorderBufferTXNSizeCompare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	return base.B2i32(base.Ui32(v5) < base.Ui32(v4)) - base.B2i32(base.Ui32(v4) < base.Ui32(v5))
}
func F_ReorderBufferTruncateTXN(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int64
	_ = v212
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v49 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v13 = l1 + int32(160)
	if v9 == v13 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v9
	goto L4
L4:
	;
	v24 = v18 - int32(188)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(148))))
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	F_ReorderBufferTruncateTXN(m, l0, v24, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v18-int32(68))))
	if v30 == int64(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v33 | int32(16)
	goto L6
L10:
	;
	goto L9
L11:
	;
	return
L12:
	;
	if v37 != v13 {
		v18 = v37
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	if l2 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L15:
	;
	v53 = l1 + int32(128)
	if v49 == v53 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v59 = v49
	v61 = int32(0)
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v67
	v70 = v59 - int32(52)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	switch v74 {
	case 0, 1, 2, 8:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	case 5:
		goto L22
	default:
		v113 = int32(64)
		goto L20
	case 11:
		goto L21
	}
L18:
	;
	if v122 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L19:
	;
	F_ReorderBufferFreeChange(m, l0, v70, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L11
	} else {
		goto L30
	}
L20:
	;
	v118 = v113
	goto L19
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v113 = v107<<(uint(int32(2))%32) - int32(-64)
	goto L20
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+24))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	v118 = (v100+v101)<<(uint(int32(2))%32) + int32(136)
	goto L19
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v118 = v94<<(uint(int32(4))%32) - int32(-64)
	goto L19
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v89 = F_strlen(m, v88)
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v118 = v89 + v90 + int32(73)
	goto L19
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+36))
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v81 = v77 + int32(84)
	goto L28
L27:
	;
	v81 = int32(64)
	goto L28
L28:
	;
	if v75 == int32(0) {
		v113 = v81
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v118 = v81 + v84 + int32(20)
	goto L19
L30:
	;
	v122 = v61 + v118
	if v53 != v65 {
		v59 = v65
		v61 = v122
		goto L17
	} else {
		goto L31
	}
L31:
	;
	goto L18
L32:
	;
	if l1 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, 12))
	v130 = v129
	goto L35
L34:
	;
	v130 = l1
	goto L35
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+216)) = v131 - v122
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+40))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v135 - v122
	if v134 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v138 = v134
	goto L38
L37:
	;
	v138 = v130
	goto L38
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+220)) = v139 - v122
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v144 = v130 + int32(204)
	F_pairingheap_remove(m, v142, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v130)+216))
	if v147 == int32(0) {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v150, v144)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	goto L14
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v196 != 0 {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v163 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v167 = l1 + int32(136)
	if v163 == v167 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v172 = v163
	goto L46
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v180
	F_ReorderBufferFreeChange(m, l0, v172-int32(52), int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L48
	}
L47:
	;
	goto L42
L48:
	;
	if v178 != v167 {
		v172 = v178
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	F_hash_destroy(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v201&int32(4) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+152)) = int32(0)
	goto L52
L54:
	;
	F_ReorderBufferRestoreCleanup(m, l1)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v212 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = v212
	return
L57:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v206&int32(-13) | int32(8)
	goto L56
}
