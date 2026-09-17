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
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
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
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	case 5:
		goto L31
	default:
		goto L29
	case 11:
		goto L30
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L24
	} else {
		goto L63
	}
L29:
	;
	F_pfree(m, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L62
	}
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v129 == int32(0) {
		goto L29
	} else {
		goto L60
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v117 == int32(0) {
		goto L29
	} else {
		goto L53
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v112 != 0 {
		goto L49
	} else {
		goto L50
	}
L33:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v102 != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v90 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_pfree(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v95 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(0)
	goto L37
L39:
	;
	F_pfree(m, v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	goto L28
L41:
	;
	F_pfree(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v107 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
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
	goto L28
L48:
	;
	goto L47
L49:
	;
	F_pfree(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L24
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
	goto L28
L52:
	;
	goto L51
L53:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+30)))
	if v120 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = int32(0)
	goto L28
L55:
	;
	F_pfree(m, v117)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L24
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_SnapBuildSnapDecRefcount(m, v117)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L24
	} else {
		goto L59
	}
L58:
	;
	goto L54
L59:
	;
	goto L54
L60:
	;
	F_pfree(m, v129)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L24
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
	goto L29
L62:
	;
	return
L63:
	;
	return
}
func F_ReorderBufferQueueMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	v9 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(192)
	m.G0 = v19
	v31 = v9
	v32 = v9
	v33 = v9
	v34 = v9
	v35 = v9
	v36 = int32(-1)
	goto L2
L1:
	;
	m.G0 = v19 + int32(192)
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
	if v36 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v188 = int32(m.ExcTag)
	v189 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v188 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L7:
	;
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v129 = v31
	v130 = v32
	v131 = v33
	v132 = v34
	v133 = v35
	goto L9
L9:
	;
	if v129 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	v40 = int32(_a_F_ReorderBufferQueueMessage_0)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[0])) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v35
	v51 = F_MemoryContextAlloc(m, v45, int32(64))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+172)) = l2
	if l1 != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v53 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51)+56)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51)+48)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51)+40)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51)+32)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v35
	v75 = F_pstrdup(m, l5)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+24)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v35
	v83 = F_palloc(m, l6)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+28)) = v83
	if l6 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, v83, l7, l6)
	goto L18
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v35
	F_ReorderBufferQueueChange(m, l0, l1, l3, v51, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[0])) = v41
	goto L1
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v35
	v102 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l3)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v105 = v35
	v106 = int32(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v105
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[1])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[2])) = v111
	goto L24
L23:
	;
	v105 = v102
	v106 = v102
	goto L22
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[3]))
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[4]))
	goto L25
L25:
	;
	v123 = v19 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v19 + int32(12)
	goto L28
L26:
	;
	v129 = int32(0)
	v130 = v121
	v131 = v119
	v132 = v106
	v133 = v105
	goto L9
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[4])) = v19 + int32(16)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v133
	v145 = int32(0)
	m.T0[v140].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32))(m, l0, v132, l3, v145, l5, l6, l7)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v133
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[1])) = v145
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[2])) = v145
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[3])) = v131
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[4])) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v133
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[1])) = v170
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[2])) = v170
	goto L33
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[3])) = v131
	*(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueMessage[4])) = v130
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v133
	F_pg_re_throw(m)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	goto L5
L35:
	;
	v193 = int32(v189)
	m.G0 = v19
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v19+int32(12) == v199 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	m.ExcPending = 1
	goto L44
L37:
	;
	if v203 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v203 = v201
	goto L40
L39:
	;
	v203 = int32(0)
	goto L40
L40:
	;
	goto L37
L41:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v19)+184))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v19)+180))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	v31 = v195
	v32 = v206
	v33 = v207
	v34 = v205
	v35 = v204
	v36 = v203
	goto L2
L42:
	;
	goto L43
L43:
	;
	F___wasm_longjmp(m, v196, v195)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferSkipPrepare(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v10 == v3)|base.B2i32(l1 != v10) == v3 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v17 != 0 {
			v39 = v17
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41 | int32(128)
		} else {
		}
		m.G0 = v7 + int32(16)
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = F_hash_search(m, v18, v7+int32(12), int32(0), v7+int32(11))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
				if v34 == int32(0) {
				} else {
					v39 = v34
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41 | int32(128)
				}
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
	var v60 int32
	_ = v60
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int64
	_ = v207
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
		goto L39
	} else {
		goto L40
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
	v60 = int32(0)
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
	v122 = v60 + v118
	if v53 != v65 {
		v59 = v65
		v60 = v122
		goto L17
	} else {
		goto L31
	}
L31:
	;
	goto L18
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+216)) = v126 - v122
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v130 - v122
	if v129 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = v129
	goto L35
L34:
	;
	v133 = l1
	goto L35
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+220)) = v134 - v122
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v139 = l1 + int32(204)
	F_pairingheap_remove(m, v137, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v142 == int32(0) {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v145, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	goto L14
L39:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v191 != 0 {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v158 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v162 = l1 + int32(136)
	if v158 == v162 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v167 = v158
	goto L43
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v175
	F_ReorderBufferFreeChange(m, l0, v167-int32(52), int32(1))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	goto L39
L45:
	;
	if v173 != v162 {
		v167 = v173
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_hash_destroy(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L11
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v196&int32(4) != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+152)) = int32(0)
	goto L49
L51:
	;
	F_ReorderBufferRestoreCleanup(m, l1)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L11
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v207 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = v207
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v207
	return
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v201&int32(-13) | int32(8)
	goto L53
}
