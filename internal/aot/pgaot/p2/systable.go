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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v23 != v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	v27 = F_list_member_ptr(m, v26, v19)
	mBase = m.M
	v28 = v27
	goto L4
L3:
	;
	v28 = int32(1)
	goto L4
L4:
	;
	goto L1
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[124])))
	if v32 != int32(1) {
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
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L49
	}
L8:
	;
	v58 = F_palloc(m, int32(24))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L15
	}
L9:
	;
	v37 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v43 + int32(4)
	F_errmsg_internal(m, int32(157702), v17+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(497670), int32(668), int32(451591))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = l0
	v63 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v63
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v69 = F_GetCatalogSnapshot(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v73 = l2
	v74 = v6
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v74
	v78 = F_palloc(m, l3*int32(48))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	v71 = F_RegisterSnapshot(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v73 = v71
	v74 = v71
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
	v207 = int32(0)
	v209 = F_index_beginscan(m, l0, l1, v73, v207, l3, v207)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L43
	}
L24:
	;
	v93 = v6
	goto L25
L25:
	;
	v97 = v93 * int32(48)
	v98 = v78 + v97
	v99 = l4 + v97
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	*(*int64)(unsafe.Add(mBase, uint32(v98))) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v99)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v99)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v99)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+24)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v99)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v99)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v110
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
	if v153 == v160 {
		goto L27
	} else {
		goto L38
	}
L29:
	;
	v153 = int32(0)
	v160 = v113
	goto L28
L30:
	;
	goto L31
L31:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)))
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
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+4)) = uint16(v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+8)))
	v153 = v126
	v160 = v144
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
	v164 = v93 + int32(1)
	if l3 != v164 {
		v93 = v164
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	F_errmsg_internal(m, int32(28435), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(497670), int32(707), int32(451591))
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v209
	v212 = int32(0)
	F_index_rescan(m, v209, v78, l3, v212, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = int32(0)
	F_pfree(m, v78)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v221 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v223)
	goto L48
L47:
	;
	goto L48
L48:
	;
	m.G0 = v17 + int32(32)
	return v58
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v236 + int32(4)
	F_errmsg(m, int32(438997), v17)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(497670), int32(664), int32(451591))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
