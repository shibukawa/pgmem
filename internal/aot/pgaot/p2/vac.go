package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_tid_reaped(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	v3 = int32(0)
	v11 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v12 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v15 = v11 | v12<<(uint(int64(16))%64)
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v337
L2:
	;
	v337 = v324
	goto L1
L3:
	;
	v301 = int32(*(*int8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v301 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L4:
	;
	if v281 != 0 {
		v291 = v281
		goto L3
	} else {
		goto L55
	}
L5:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	if base.Ui64(v20) < base.Ui64(v15) {
		v324 = v3
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui64(v154) < base.Ui64(v15) {
		v324 = v3
		goto L2
	} else {
		goto L33
	}
L8:
	;
	v22 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+48)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v25 = v23
	v32 = v22
	goto L9
L9:
	;
	v35 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(v32) % 64)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v37 = F_dsa_get_address(m, v36, v25)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	if v146&int32(1) != 0 {
		v291 = v143
		goto L3
	} else {
		goto L31
	}
L11:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if int64(7) < v32 {
		v25 = v146
		v32 = v32 - int64(8)
		goto L9
	} else {
		goto L30
	}
L12:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(base.Ui32(v35)>>(uint(int32(3))%32))&int32(28))+4))
	if int32(base.Ui32(v120)>>(uint(v35)%32))&int32(1) == int32(0) {
		v324 = v3
		goto L2
	} else {
		goto L29
	}
L13:
	;
	v104 = int32(255)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v35&v104)+12)))
	if v107 == v104 {
		v324 = v3
		goto L2
	} else {
		goto L28
	}
L14:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v74 == int32(0) {
		v324 = v3
		goto L2
	} else {
		goto L23
	}
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v44 == int32(0) {
		v324 = v3
		goto L2
	} else {
		goto L18
	}
L16:
	;
	return int32(0)
L17:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	switch v41 - int32(1) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	default:
		goto L15
	}
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v37+int32(3))))))
	if v35&int32(255) == v68 {
		v143 = v37 + v52<<(uint(int32(2))%32) + int32(8)
		goto L11
	} else {
		goto L21
	}
L20:
	;
	v337 = int32(0)
	goto L1
L21:
	;
	v71 = v52 + int32(1)
	if v71 != v44 {
		v52 = v71
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v80 = int32(0)
	goto L24
L24:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+(v37+int32(3))))))
	if v96 == v35&int32(255) {
		v143 = v37 + v80<<(uint(int32(2))%32) + int32(36)
		goto L11
	} else {
		goto L26
	}
L25:
	;
	v337 = int32(0)
	goto L1
L26:
	;
	v101 = v80 + int32(1)
	if v101 != v74 {
		v80 = v101
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v143 = v37 + v107<<(uint(int32(2))%32) + int32(268)
	goto L11
L29:
	;
	v143 = v37 + v35&int32(255)<<(uint(int32(2))%32) + int32(36)
	goto L11
L30:
	;
	goto L10
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v152 = F_dsa_get_address(m, v151, v146)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v281 = v152
	goto L4
L33:
	;
	v156 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+24)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v159 = v157
	v166 = v156
	goto L34
L34:
	;
	v169 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(v166) % 64)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	switch v170 - int32(1) {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	default:
		goto L40
	}
L35:
	;
	if v275&int32(1) != 0 {
		v291 = v272
		goto L3
	} else {
		goto L54
	}
L36:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if int64(7) < v166 {
		v159 = v275
		v166 = v166 - int64(8)
		goto L34
	} else {
		goto L53
	}
L37:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(base.Ui32(v169)>>(uint(int32(3))%32))&int32(28))+4))
	if int32(base.Ui32(v249)>>(uint(v169)%32))&int32(1) == int32(0) {
		v324 = v3
		goto L2
	} else {
		goto L52
	}
L38:
	;
	v233 = int32(255)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v169&v233)+12)))
	if v236 == v233 {
		v324 = v3
		goto L2
	} else {
		goto L51
	}
L39:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+2)))
	if v203 == int32(0) {
		v324 = v3
		goto L2
	} else {
		goto L46
	}
L40:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+2)))
	if v173 == int32(0) {
		v324 = v3
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v181 = int32(0)
	goto L42
L42:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+(v159+int32(3))))))
	if v169&int32(255) == v197 {
		v272 = v159 + v181<<(uint(int32(2))%32) + int32(8)
		goto L36
	} else {
		goto L44
	}
L43:
	;
	v337 = int32(0)
	goto L1
L44:
	;
	v200 = v181 + int32(1)
	if v200 != v173 {
		v181 = v200
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v209 = int32(0)
	goto L47
L47:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+(v159+int32(3))))))
	if v225 == v169&int32(255) {
		v272 = v159 + v209<<(uint(int32(2))%32) + int32(36)
		goto L36
	} else {
		goto L49
	}
L48:
	;
	v337 = int32(0)
	goto L1
L49:
	;
	v230 = v209 + int32(1)
	if v230 != v203 {
		v209 = v230
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v272 = v159 + v236<<(uint(int32(2))%32) + int32(268)
	goto L36
L52:
	;
	v272 = v159 + v169&int32(255)<<(uint(int32(2))%32) + int32(36)
	goto L36
L53:
	;
	goto L35
L54:
	;
	v281 = v275
	goto L4
L55:
	;
	v337 = int32(0)
	goto L1
L56:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291)+2)))
	v337 = base.B2i32(v304 == v16)
	goto L1
L57:
	;
	goto L58
L58:
	;
	v308 = int32(base.Ui32(v16) >> (uint(int32(5)) % 32))
	if v301 <= v308 {
		v337 = int32(0)
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v291+v308<<(uint(int32(2))%32))+4))
	v324 = int32(base.Ui32(v313)>>(uint(v16)%32)) & int32(1)
	goto L2
}
