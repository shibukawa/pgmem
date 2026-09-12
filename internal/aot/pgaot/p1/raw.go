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
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L65
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v35 = (v31 + int32(7)) & int32(-8)
	if base.Ui32(v35) < base.Ui32(int32(8161)) {
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
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L61
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v46 = base.I32_div_s(int32(819200)-v41<<(uint(int32(13))%32), int32(100))
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
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v183 = F_PageAddItemExtended(m, v178, v179, v180, int32(0), int32(2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L52
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = F_smgr_bulk_get_buf(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L41
	}
L20:
	;
	if base.Ui32(v35+v48) <= base.Ui32(v118) {
		v178 = v49
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
	v118 = int32(0)
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
	v118 = v61
	goto L20
L28:
	;
	v72 = int32(base.Ui32(v55+int32(262120)) >> (uint(int32(2)) % 32))
	goto L30
L29:
	;
	v72 = int32(0)
	goto L30
L30:
	;
	if base.Ui32(v72&int32(65535)) < base.Ui32(int32(291)) {
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
	v118 = int32(0)
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
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86&int32(65535)<<(uint(int32(2))%32)+(v49+int32(24))-int32(3)))))
	if v97&int32(384) == int32(0) {
		goto L27
	} else {
		goto L37
	}
L36:
	;
	v118 = int32(0)
	goto L20
L37:
	;
	v103 = v86 + int32(1)
	v104 = int32(65535)
	if base.Ui32(v103&v104) <= base.Ui32(v72&v104) {
		v86 = v103
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_smgr_bulk_write(m, v121, v122, v123, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v129 + int32(1)
	goto L19
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v134
	if v134&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v178 = v134
	goto L16
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+10)) = int32(1572864)
	v169 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+18)) = uint16(v169)
	v175 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+16)) = uint16(v175)
	*(*uint16)(unsafe.Add(mBase, uint32(v134)+14)) = uint16(v175)
	goto L42
L44:
	;
	v163 = F___memset(m, v134, int32(0), int32(8192))
	mBase = m.M
	goto L43
L45:
	;
	goto L44
L52:
	;
	if v183 == int32(0) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v183)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v187)
	v191 = int32(base.Ui32(v187) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+16)))
	if v194 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v183<<(uint(int32(2))%32)+v178)+20))
	v203 = v178 + v200&int32(32767)
	v205 = l1 + int32(4)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v206
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v203)+16)) = uint16(v208)
	goto L56
L55:
	;
	goto L56
L56:
	;
	if l1 != v30 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v30)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	m.G0 = v10 + int32(16)
	return
L60:
	;
	goto L59
L61:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8160)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v35
	F_errmsg(m, int32(39015), v10)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(520510), int32(641), int32(87845))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(403534), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(520510), int32(679), int32(87845))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
