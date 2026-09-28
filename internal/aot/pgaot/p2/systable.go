package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_systable_beginscan_ordered(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[0]))
	if v21 != v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[1]))
	v25 = F_list_member_ptr(m, v24, v19)
	mBase = m.M
	v27 = v25
	goto L4
L3:
	;
	v27 = int32(1)
	goto L4
L4:
	;
	goto L1
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[2])))
	if v31 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L49
	}
L8:
	;
	v57 = F_palloc(m, int32(24))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L15
	}
L9:
	;
	v36 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v36 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v42 + int32(4)
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_ordered_0), v17+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_ordered_1), int32(669), int32(_a_F_systable_beginscan_ordered_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
	v62 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v62
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v68 = F_GetCatalogSnapshot(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v72 = l2
	v73 = v6
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v73
	v76 = F_palloc_mul(m, int32(56), l3)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	v70 = F_RegisterSnapshot(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v72 = v70
	v73 = v70
	goto L19
L22:
	;
	if l3 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[3]))
	if v208 != 0 {
		goto L43
	} else {
		goto L44
	}
L24:
	;
	v92 = v6
	goto L25
L25:
	;
	v95 = v92 * int32(56)
	v96 = v76 + v95
	v97 = l4 + v95
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+48)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v97)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+40)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v97)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+32)) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v97)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+24)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v97)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+8)))
	if v113 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L40
	}
L27:
	;
	goto L26
L28:
	;
	if v153 == v159 {
		goto L27
	} else {
		goto L38
	}
L29:
	;
	v153 = int32(0)
	v159 = v113
	goto L28
L30:
	;
	goto L31
L31:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)))
	v126 = int32(0)
	goto L32
L32:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112+int32(48)+v126<<(uint(int32(1))%32)))))
	if v138 == v120 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v141 = v126 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+4)) = uint16(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+8)))
	v153 = v126
	v159 = v144
	goto L28
L35:
	;
	goto L36
L36:
	;
	v146 = v126 + int32(1)
	if v146 != v113 {
		v126 = v146
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v164 = v92 + int32(1)
	if l3 != v164 {
		v92 = v164
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_ordered_3), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_ordered_1), int32(708), int32(_a_F_systable_beginscan_ordered_2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[4])) = uint8(v210)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v212 = int32(0)
	v215 = F_index_beginscan(m, l0, l1, v72, v212, l3, v212, v212)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v215
	v218 = int32(0)
	F_index_rescan(m, v215, v76, l3, v218, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
	F_pfree(m, v76)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v17 + int32(32)
	return v57
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v237 + int32(4)
	F_errmsg(m, int32(_a_F_systable_beginscan_ordered_4), v17)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_ordered_1), int32(665), int32(_a_F_systable_beginscan_ordered_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
