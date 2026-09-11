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
	var v52 int32
	_ = v52
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
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
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
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
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	if v149 == int32(0) {
		goto L37
	} else {
		goto L38
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
	v52 = int32(0)
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
	v110 = v52 + v106
	if v58 != v45 {
		v50 = v58
		v52 = v110
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	if l1 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v117 = *(*int32)(unsafe.Add(mBase, 12))
	v118 = v117
	goto L30
L29:
	;
	v118 = l1
	goto L30
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+216)) = v119 - v110
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v123 - v110
	if v122 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v126 = v122
	goto L33
L32:
	;
	v126 = v118
	goto L33
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+220)) = v127 - v110
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v132 = v118 + int32(204)
	F_pairingheap_remove(m, v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+216))
	if v135 == int32(0) {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v138, v132)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L9
L37:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v178 != 0 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v153 = l1 + int32(136)
	if v149 == v153 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v157 = v149
	goto L40
L40:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	F_ReorderBufferFreeChange(m, l0, v157-int32(52), int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L42
	}
L41:
	;
	goto L37
L42:
	;
	if v165 != v153 {
		v157 = v165
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_SnapBuildSnapDecRefcount(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v187 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v184
	goto L46
L48:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v202&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+30)))
	if v190 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_pfree(m, v187)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_SnapBuildSnapDecRefcount(m, v187)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L54
	}
L53:
	;
	goto L48
L54:
	;
	goto L48
L55:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+4)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v210 - int32(1)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v217 = l1 + int32(4)
	v221 = F_hash_search(m, v215, v217, int32(2), v11+int32(15))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v223&int32(4) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_ReorderBufferRestoreCleanup(m, l1)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v228 == v229 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = int64(0)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v233 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_pfree(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v238 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	goto L68
L70:
	;
	F_hash_destroy(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v243 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+152)) = int32(0)
	goto L72
L74:
	;
	F_pfree(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v248 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+176)) = int32(0)
	goto L76
L78:
	;
	F_pfree(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_ReorderBufferToastReset(m, l0, l1)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L6
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+184)) = int32(0)
	goto L80
L82:
	;
	F_pfree(m, l1)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
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
	var v19 int32
	_ = v19
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
	F_ReorderBufferProcessTXN(m, l0, l1, int64(0), v56, v57, int32(1))
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
	v19 = v10
	goto L9
L9:
	;
	F_ReorderBufferTransferSnapToParent(m, l1, v19-int32(188))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v26 != v14 {
		v19 = v26
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
	v56 = v39
	v57 = v37
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
	v56 = v42
	v57 = v41
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	F_hash_seq_init(m, v9+int32(12), v11)
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
	v18 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = v18
	goto L10
L8:
	;
	goto L9
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	F_hash_destroy(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L25
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	F_pfree(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v29 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v60 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L23
	}
L17:
	;
	v33 = v23 + int32(16)
	if v29 == v33 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v37 = v29
	goto L19
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v44
	F_ReorderBufferFreeChange(m, l0, v37-int32(52), int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	if v33 != v42 {
		v37 = v42
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v60 != 0 {
		v23 = v60
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v11 == v3 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = F_hash_search(m, v16, v8+int32(4), int32(0), v8+int32(3))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29
				v76 = v3
				m.G0 = v8 + int32(16)
				return v76
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
				if v34 == int32(0) {
					v76 = v3
					m.G0 = v8 + int32(16)
					return v76
				} else {
					v39 = v33
					v40 = v34
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
					if v41&int32(2) == int32(0) {
						v69 = v40
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
						v76 = base.B2i32(v71 != int32(0))
						m.G0 = v8 + int32(16)
						return v76
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v46
						if v39 == v46 {
							v50 = v39
						} else {
							v50 = int32(0)
						}
						if v50 != 0 {
							v69 = v40
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
							v76 = base.B2i32(v71 != int32(0))
							m.G0 = v8 + int32(16)
							return v76
						} else {
							v51 = int32(0)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v58 = F_hash_search(m, v52, v8+int32(12), v51, v8+int32(11))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
								if v60 == int32(1) {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
									v64 = v63
								} else {
									v64 = v51
								}
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v65
								v69 = v64
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
								v76 = base.B2i32(v71 != int32(0))
								m.G0 = v8 + int32(16)
								return v76
							}
						}
					}
				}
			}
		}
	} else {
		if l1 != v11 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = F_hash_search(m, v16, v8+int32(4), int32(0), v8+int32(3))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)))
				if v26 == int32(0) {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v29
					v76 = v3
					m.G0 = v8 + int32(16)
					return v76
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v33
					if v34 == int32(0) {
						v76 = v3
						m.G0 = v8 + int32(16)
						return v76
					} else {
						v39 = v33
						v40 = v34
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
						if v41&int32(2) == int32(0) {
							v69 = v40
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
							v76 = base.B2i32(v71 != int32(0))
							m.G0 = v8 + int32(16)
							return v76
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v46
							if v39 == v46 {
								v50 = v39
							} else {
								v50 = int32(0)
							}
							if v50 != 0 {
								v69 = v40
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
								v76 = base.B2i32(v71 != int32(0))
								m.G0 = v8 + int32(16)
								return v76
							} else {
								v51 = int32(0)
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v58 = F_hash_search(m, v52, v8+int32(12), v51, v8+int32(11))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
									if v60 == int32(1) {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
										v64 = v63
									} else {
										v64 = v51
									}
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v65
									v69 = v64
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
									v76 = base.B2i32(v71 != int32(0))
									m.G0 = v8 + int32(16)
									return v76
								}
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v15 != 0 {
				v39 = l1
				v40 = v15
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
				if v41&int32(2) == int32(0) {
					v69 = v40
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
					v76 = base.B2i32(v71 != int32(0))
					m.G0 = v8 + int32(16)
					return v76
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v46
					if v39 == v46 {
						v50 = v39
					} else {
						v50 = int32(0)
					}
					if v50 != 0 {
						v69 = v40
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
						v76 = base.B2i32(v71 != int32(0))
						m.G0 = v8 + int32(16)
						return v76
					} else {
						v51 = int32(0)
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v58 = F_hash_search(m, v52, v8+int32(12), v51, v8+int32(11))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
							if v60 == int32(1) {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v64 = v63
							} else {
								v64 = v51
							}
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v65
							v69 = v64
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+80))
							v76 = base.B2i32(v71 != int32(0))
							m.G0 = v8 + int32(16)
							return v76
						}
					}
				}
			} else {
				v76 = v3
				m.G0 = v8 + int32(16)
				return v76
			}
		}
	}
}
