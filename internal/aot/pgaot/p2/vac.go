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
	var v26 int32
	_ = v26
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
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v277 int64
	_ = v277
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v343 int32
	_ = v343
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
	return v343
L2:
	;
	v343 = v329
	goto L1
L3:
	;
	v307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v307 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L4:
	;
	if v288 != 0 {
		v297 = v288
		goto L3
	} else {
		goto L71
	}
L5:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	if base.Ui64(v20) < base.Ui64(v15) {
		v329 = v3
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui64(v157) < base.Ui64(v15) {
		v329 = v3
		goto L2
	} else {
		goto L41
	}
L8:
	;
	v22 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+48)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v26 = v23
	v32 = v22
	goto L9
L9:
	;
	v35 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(v32) % 64)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v37 = F_dsa_get_address(m, v36, v26)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	if v147&int32(1) != 0 {
		v297 = v135
		goto L3
	} else {
		goto L39
	}
L11:
	;
	v145 = int64(8)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if base.B2i32(v32 < v145) == int32(0) {
		v26 = v147
		v32 = v32 - v145
		goto L9
	} else {
		goto L38
	}
L12:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(base.Ui32(v35)>>(uint(int32(3))%32))&int32(28))+4))
	if int32(base.Ui32(v120)>>(uint(v35)%32))&int32(1) == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L36
	}
L13:
	;
	v104 = int32(255)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v35&v104)+12)))
	if v107 == v104 {
		v329 = v3
		goto L2
	} else {
		goto L34
	}
L14:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v74 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v44 == int32(0) {
		v329 = v3
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
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v37+int32(3))))))
	if v35&int32(255) == v63 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v343 = int32(0)
	goto L1
L21:
	;
	v69 = v37 + v52<<(uint(int32(2))%32) + int32(8)
	if v69 != 0 {
		v135 = v69
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = v52 + int32(1)
	if v71 != v44 {
		v52 = v71
		goto L19
	} else {
		goto L25
	}
L24:
	;
	v329 = v3
	goto L2
L25:
	;
	goto L20
L26:
	;
	v80 = int32(0)
	goto L27
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+(v37+int32(3))))))
	if v91 == v35&int32(255) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v343 = int32(0)
	goto L1
L29:
	;
	v99 = v37 + v80<<(uint(int32(2))%32) + int32(36)
	if v99 != 0 {
		v135 = v99
		goto L11
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v101 = v80 + int32(1)
	if v101 != v74 {
		v80 = v101
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v329 = v3
	goto L2
L33:
	;
	goto L28
L34:
	;
	v114 = v37 + v107<<(uint(int32(2))%32) + int32(268)
	if v114 != 0 {
		v135 = v114
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v329 = v3
	goto L2
L36:
	;
	v132 = v37 + v35&int32(255)<<(uint(int32(2))%32) + int32(36)
	if v132 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v135 = v132
	goto L11
L38:
	;
	goto L10
L39:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v155 = F_dsa_get_address(m, v154, v147)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v288 = v155
	goto L4
L41:
	;
	v159 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+24)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v163 = v160
	v169 = v159
	goto L42
L42:
	;
	v172 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(v169) % 64)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	switch v173 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	default:
		goto L48
	}
L43:
	;
	if v279&int32(1) != 0 {
		v297 = v267
		goto L3
	} else {
		goto L70
	}
L44:
	;
	v277 = int64(8)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if base.B2i32(v169 < v277) == int32(0) {
		v163 = v279
		v169 = v169 - v277
		goto L42
	} else {
		goto L69
	}
L45:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v163+int32(base.Ui32(v172)>>(uint(int32(3))%32))&int32(28))+4))
	if int32(base.Ui32(v252)>>(uint(v172)%32))&int32(1) == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L67
	}
L46:
	;
	v236 = int32(255)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v172&v236)+12)))
	if v239 == v236 {
		v329 = v3
		goto L2
	} else {
		goto L65
	}
L47:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+2)))
	if v206 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L57
	}
L48:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+2)))
	if v176 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v184 = int32(0)
	goto L50
L50:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+(v163+int32(3))))))
	if v172&int32(255) == v195 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v343 = int32(0)
	goto L1
L52:
	;
	v201 = v163 + v184<<(uint(int32(2))%32) + int32(8)
	if v201 != 0 {
		v267 = v201
		goto L44
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v203 = v184 + int32(1)
	if v203 != v176 {
		v184 = v203
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v329 = v3
	goto L2
L56:
	;
	goto L51
L57:
	;
	v212 = int32(0)
	goto L58
L58:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+(v163+int32(3))))))
	if v223 == v172&int32(255) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v343 = int32(0)
	goto L1
L60:
	;
	v231 = v163 + v212<<(uint(int32(2))%32) + int32(36)
	if v231 != 0 {
		v267 = v231
		goto L44
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v233 = v212 + int32(1)
	if v233 != v206 {
		v212 = v233
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v329 = v3
	goto L2
L64:
	;
	goto L59
L65:
	;
	v246 = v163 + v239<<(uint(int32(2))%32) + int32(268)
	if v246 != 0 {
		v267 = v246
		goto L44
	} else {
		goto L66
	}
L66:
	;
	v329 = v3
	goto L2
L67:
	;
	v264 = v163 + v172&int32(255)<<(uint(int32(2))%32) + int32(36)
	if v264 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v267 = v264
	goto L44
L69:
	;
	goto L43
L70:
	;
	v288 = v279
	goto L4
L71:
	;
	v343 = int32(0)
	goto L1
L72:
	;
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+2)))
	v343 = base.B2i32(v310 == v16)
	goto L1
L73:
	;
	goto L74
L74:
	;
	v314 = int32(base.Ui32(v16) >> (uint(int32(5)) % 32))
	if v307 <= v314 {
		v343 = int32(0)
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v297+v314<<(uint(int32(2))%32))+4))
	v329 = int32(base.Ui32(v319)>>(uint(v16)%32)) & int32(1)
	goto L2
}
