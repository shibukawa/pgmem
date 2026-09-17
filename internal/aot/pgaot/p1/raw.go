package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_raw_heap_insert(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+119)))
	if v14 == int32(116) {
		v30 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L66
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v35 = (v31 + int32(7)) & int32(-8)
	if base.Ui32(v35) < base.Ui32(int32(_a_F_raw_heap_insert_0)) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)))
	if v18&int32(4) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v23) < base.Ui32(int32(2033)) {
		v30 = l1
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v28 = F_heap_toast_insert_or_update(m, v12, l1, int32(0), int32(10))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	return
L9:
	;
	v30 = v28
	goto L2
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L62
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v46 = base.I32_div_s(int32(_a_F_raw_heap_insert_1)-v41<<(uint(int32(13))%32), int32(100))
	v48 = v46
	goto L15
L14:
	;
	v48 = int32(0)
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v49 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v190 = F_PageAddItemExtended(m, v185, v186, v187, int32(0), int32(2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L53
	}
L17:
	;
	v53 = int32(4)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+14)))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+12)))
	v56 = v54 - v55
	if v56 <= v53 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v132 = F_smgr_bulk_get_buf(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L41
	}
L20:
	;
	if base.Ui32(v35+v48) <= base.Ui32(v116) {
		v185 = v49
		goto L16
	} else {
		goto L39
	}
L21:
	;
	v59 = v53
	goto L23
L22:
	;
	v59 = v56
	goto L23
L23:
	;
	v61 = v59 - int32(4)
	if v61 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v116 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v55) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v116 = v61
	goto L20
L28:
	;
	v72 = int32(base.Ui32(v55+int32(_a_F_raw_heap_insert_2)) >> (uint(int32(2)) % 32))
	goto L30
L29:
	;
	v72 = int32(0)
	goto L30
L30:
	;
	if base.Ui32(v72&int32(_a_F_raw_heap_insert_3)) < base.Ui32(int32(291)) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
	if v77&int32(1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v116 = int32(0)
	goto L20
L33:
	;
	goto L34
L34:
	;
	v86 = int32(1)
	goto L35
L35:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+int32(20)+v86&int32(_a_F_raw_heap_insert_3)<<(uint(int32(2))%32))+1)))
	if v95&int32(384) == int32(0) {
		goto L27
	} else {
		goto L37
	}
L36:
	;
	v116 = int32(0)
	goto L20
L37:
	;
	v101 = v86 + int32(1)
	v102 = int32(_a_F_raw_heap_insert_3)
	if base.Ui32(v101&v102) <= base.Ui32(v72&v102) {
		v86 = v101
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_smgr_bulk_write(m, v119, v120, v121, int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v127 + int32(1)
	goto L19
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v132
	v135 = int32(_a_F_raw_heap_insert_4)
	v136 = int32(0)
	if v136|(v132&int32(3)|int32(1)) == v136 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v185 = v132
	goto L16
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+10)) = int32(_a_F_raw_heap_insert_5)
	v176 = int32(_a_F_raw_heap_insert_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v132)+18)) = uint16(v176)
	v182 = int32(_a_F_raw_heap_insert_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v132)+16)) = uint16(v182)
	*(*uint16)(unsafe.Add(mBase, uint32(v132)+14)) = uint16(v182)
	goto L42
L44:
	;
	goto L47
L45:
	;
	goto L46
L46:
	;
	goto L52
L47:
	;
	v153 = v132 + v135
	v155 = v132 + int32(4)
	if base.Ui32(v155) < base.Ui32(v153) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v157 = v153
	goto L50
L49:
	;
	v157 = v155
	goto L50
L50:
	;
	v162 = (v132^int32(-1)+v157)&int32(-4) + int32(4)
	if v162 == int32(0) {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	base.MemoryFill(m, v132, int32(0), v162)
	goto L43
L52:
	;
	base.MemoryFill(m, v132, int32(0), v135)
	goto L43
L53:
	;
	if v190 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v190)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v194)
	v198 = int32(base.Ui32(v194) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+16)))
	if v201 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v185+v190<<(uint(int32(2))%32))+20))
	v210 = v185 + v207&int32(_a_F_raw_heap_insert_7)
	v212 = l1 + int32(4)
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v210)+16)) = uint16(v213)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+12)) = v215
	goto L57
L56:
	;
	goto L57
L57:
	;
	if l1 != v30 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_pfree(m, v30)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	m.G0 = v10 + int32(16)
	return
L61:
	;
	goto L60
L62:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_raw_heap_insert_8)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v35
	F_errmsg(m, int32(_a_F_raw_heap_insert_9), v10)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_raw_heap_insert_10), int32(641), int32(_a_F_raw_heap_insert_11))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errmsg_internal(m, int32(_a_F_raw_heap_insert_12), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_raw_heap_insert_10), int32(679), int32(_a_F_raw_heap_insert_11))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
