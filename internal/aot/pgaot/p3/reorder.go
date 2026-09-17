package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferAllocTupleBuf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v6 = F_MemoryContextAlloc(m, v3, l1+int32(47))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v6 + int32(24)
		return v6
	}
}
func F_ReorderBufferCleanupTXN(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v17 = l1 + int32(160)
	if v13 == v17 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v13
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	F_ReorderBufferCleanupTXN(m, l0, v21-int32(188))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v29 != v17 {
		v21 = v29
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v144 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L10:
	;
	v45 = l1 + int32(128)
	if v41 == v45 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = v41
	v51 = int32(0)
	goto L12
L12:
	;
	v57 = v50 - int32(52)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	switch v62 {
	case 0, 1, 2, 8:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	case 5:
		goto L17
	default:
		v101 = int32(64)
		goto L15
	case 11:
		goto L16
	}
L13:
	;
	if v110 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L14:
	;
	F_ReorderBufferFreeChange(m, l0, v57, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L25
	}
L15:
	;
	v106 = v101
	goto L14
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v101 = v95<<(uint(int32(2))%32) - int32(-64)
	goto L15
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v106 = (v88+v89)<<(uint(int32(2))%32) + int32(136)
	goto L14
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v106 = v82<<(uint(int32(4))%32) - int32(-64)
	goto L14
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v77 = F_strlen(m, v76)
	mBase = m.M
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v106 = v77 + v78 + int32(73)
	goto L14
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v69 = v65 + int32(84)
	goto L23
L22:
	;
	v69 = int32(64)
	goto L23
L23:
	;
	if v63 == int32(0) {
		v101 = v69
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v106 = v69 + v72 + int32(20)
	goto L14
L25:
	;
	v110 = v51 + v106
	if v58 != v45 {
		v50 = v58
		v51 = v110
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+216)) = v114 - v110
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v118 - v110
	if v117 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v121 = v117
	goto L30
L29:
	;
	v121 = l1
	goto L30
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+220)) = v122 - v110
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v127 = l1 + int32(204)
	F_pairingheap_remove(m, v125, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	if v130 == int32(0) {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v133, v127)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	goto L9
L34:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v173 != 0 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v148 = l1 + int32(136)
	if v144 == v148 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v152 = v144
	goto L37
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	F_ReorderBufferFreeChange(m, l0, v152-int32(52), int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L39
	}
L38:
	;
	goto L34
L39:
	;
	if v160 != v148 {
		v152 = v160
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_SnapBuildSnapDecRefcount(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v182 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v179
	goto L43
L45:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v195
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v197&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+30)))
	if v185 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_pfree(m, v182)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_SnapBuildSnapDecRefcount(m, v182)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	goto L45
L51:
	;
	goto L45
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v205 - int32(1)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v212 = l1 + int32(4)
	v216 = F_hash_search(m, v210, v212, int32(2), v11+int32(15))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v218&int32(4) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_ReorderBufferRestoreCleanup(m, l1)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v223 == v224 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v228 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_pfree(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v233 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	goto L65
L67:
	;
	F_hash_destroy(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v238 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+152)) = int32(0)
	goto L69
L71:
	;
	F_pfree(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v243 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+176)) = int32(0)
	goto L73
L75:
	;
	F_pfree(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_ReorderBufferToastReset(m, l0, l1)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+184)) = int32(0)
	goto L77
L79:
	;
	F_pfree(m, l1)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_ReorderBufferProcessXid(m *base.Module, l0 int32, l1 int32, l2 int64) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	if l1 != 0 {
		v5 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l2)
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_ReorderBufferStreamTXN(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v7 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+220)))
	F_ReorderBufferProcessTXN(m, l0, l1, int64(0), v57, v56, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L23
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
	if v10 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v42 = F_ReorderBufferCopySnap(m, l0, v7, l1, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L16
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v34 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	v14 = l1 + int32(160)
	if v10 == v14 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v18 = v10
	goto L9
L9:
	;
	F_ReorderBufferTransferSnapToParent(m, l1, v18-int32(188))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	return
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v26 != v14 {
		v18 = v26
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v37 = int32(0)
	v39 = F_ReorderBufferCopySnap(m, l0, v34, l1, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v56 = v37
	v57 = v39
	goto L2
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+30)))
	if v45 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = int32(0)
	v56 = v41
	v57 = v42
	goto L2
L18:
	;
	F_pfree(m, v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_SnapBuildSnapDecRefcount(m, v44)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L22
	}
L21:
	;
	goto L17
L22:
	;
	goto L17
L23:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v66 + int64(1)
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v61 + v70
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v73 + base.I64_extend_i32_u(int32(base.Ui32(v60^int32(-1))>>(uint(int32(4))%32))&int32(1))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_UpdateDecodingStats(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L1
}
func F_ReorderBufferToastReset(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = v9 + int32(12)
	F_hash_seq_init(m, v13, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(32)
	return
L4:
	;
	return
L5:
	;
	v16 = F_hash_seq_search(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = v16
	goto L10
L8:
	;
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	F_hash_destroy(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L25
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	F_pfree(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v27 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v58 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L23
	}
L17:
	;
	v31 = v22 + int32(16)
	if v27 == v31 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v35 = v27
	goto L19
L19:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
	F_ReorderBufferFreeChange(m, l0, v35-int32(52), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	if v31 != v40 {
		v35 = v40
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v58 != 0 {
		v22 = v58
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L11
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+156)) = int32(0)
	goto L3
}
func F_ReorderBufferXidHasBaseSnapshot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v11 == v3)|base.B2i32(l1 != v11) == v3 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v18 != 0 {
			v42 = l1
			v43 = v18
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v44&int32(2) == int32(0) {
				v72 = v43
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
				v79 = base.B2i32(v74 != int32(0))
				m.G0 = v8 + int32(16)
				return v79
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v49
				if v42 == v49 {
					v53 = v42
				} else {
					v53 = int32(0)
				}
				if v53 != 0 {
					v72 = v43
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
					v79 = base.B2i32(v74 != int32(0))
					m.G0 = v8 + int32(16)
					return v79
				} else {
					v54 = int32(0)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v61 = F_hash_search(m, v55, v8+int32(12), v54, v8+int32(11))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
						if v63 == int32(1) {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
							v67 = v66
						} else {
							v67 = v54
						}
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v68
						v72 = v67
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
						v79 = base.B2i32(v74 != int32(0))
						m.G0 = v8 + int32(16)
						return v79
					}
				}
			}
		} else {
			v79 = v3
			m.G0 = v8 + int32(16)
			return v79
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v25 = F_hash_search(m, v19, v8+int32(4), int32(0), v8+int32(3))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)))
			if v29 == int32(0) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v32
				v79 = v3
				m.G0 = v8 + int32(16)
				return v79
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v37
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v36
				if v37 == int32(0) {
					v79 = v3
					m.G0 = v8 + int32(16)
					return v79
				} else {
					v42 = v36
					v43 = v37
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
					if v44&int32(2) == int32(0) {
						v72 = v43
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
						v79 = base.B2i32(v74 != int32(0))
						m.G0 = v8 + int32(16)
						return v79
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v49
						if v42 == v49 {
							v53 = v42
						} else {
							v53 = int32(0)
						}
						if v53 != 0 {
							v72 = v43
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
							v79 = base.B2i32(v74 != int32(0))
							m.G0 = v8 + int32(16)
							return v79
						} else {
							v54 = int32(0)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v61 = F_hash_search(m, v55, v8+int32(12), v54, v8+int32(11))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
								if v63 == int32(1) {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
									v67 = v66
								} else {
									v67 = v54
								}
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v68
								v72 = v67
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+80))
								v79 = base.B2i32(v74 != int32(0))
								m.G0 = v8 + int32(16)
								return v79
							}
						}
					}
				}
			}
		}
	}
}
