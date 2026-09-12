package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	v5 = m.G0
	v7 = v5 - int32(288)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+240)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+244)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+248)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+224)) = v7 + int32(284)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+228)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+232)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+236)) = v7 + int32(272)
	v35 = F_sscanf(m, v10, int32(173636), v7+int32(224))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v7 + int32(288)
	return v280
L2:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v7)+284))
	if base.Ui32(int32(255)) < base.Ui32(v230) {
		goto L24
	} else {
		goto L25
	}
L3:
	;
	return int32(0)
L4:
	;
	if v35 == int32(6) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+216)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+212)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+208)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+204)) = v7 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+200)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+196)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+192)) = v7 + int32(284)
	v65 = F_sscanf(m, v10, int32(173657), v7+int32(192))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v65 == int32(6) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+184)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+180)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+176)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = v7 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+168)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+164)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+160)) = v7 + int32(284)
	v93 = F_sscanf(m, v10, int32(173542), v7+int32(160))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v93 == int32(6) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+152)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+148)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+144)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+140)) = v7 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+136)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+132)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = v7 + int32(284)
	v121 = F_sscanf(m, v10, int32(173565), v7+int32(128))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v121 == int32(6) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+120)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+116)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+108)) = v7 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = v7 + int32(284)
	v149 = F_sscanf(m, v10, int32(173588), v7+int32(96))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v149 == int32(6) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = v7 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v7 + int32(284)
	v177 = F_sscanf(m, v10, int32(173612), v7-int32(-64))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	if v177 == int32(6) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v7 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v7 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v7 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v7 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v7 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v7 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v7 + int32(284)
	v205 = F_sscanf(m, v10, int32(173520), v7+int32(32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if v205 == int32(6) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v209 = int32(0)
	v210 = F_errsave_start(m, v9)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v210 == int32(0) {
		v280 = v209
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(226948)
	F_errmsg(m, int32(703551), v7+int32(16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v9, int32(494573), int32(95), int32(276510))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v280 = v209
	goto L1
L23:
	;
	v266 = F_palloc(m, int32(6))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L36
	}
L24:
	;
	v248 = int32(0)
	v249 = F_errsave_start(m, v9)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v7)+280))
	if base.Ui32(int32(255)) < base.Ui32(v233) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v7)+276))
	if base.Ui32(int32(255)) < base.Ui32(v236) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v7)+272))
	if base.Ui32(int32(255)) < base.Ui32(v239) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v7)+268))
	if base.Ui32(int32(255)) < base.Ui32(v242) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v7)+264))
	if base.Ui32(v245) < base.Ui32(int32(256)) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	if v249 == int32(0) {
		v280 = v248
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
	F_errmsg(m, int32(704961), v7)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errsave_finish(m, v9, int32(494573), int32(102), int32(276510))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v280 = v248
	goto L1
L36:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v7)+284))
	*(*uint8)(unsafe.Add(mBase, uint32(v266))) = uint8(v268)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v7)+280))
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)) = uint8(v270)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v7)+276))
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+2)) = uint8(v272)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v7)+272))
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+3)) = uint8(v274)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v7)+268))
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+4)) = uint8(v276)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v7)+264))
	*(*uint8)(unsafe.Add(mBase, uint32(v266)+5)) = uint8(v278)
	v280 = v266
	goto L1
}
func F_macaddr_not(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(6))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		v10 = int32(-1)
		v11 = v9 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v11)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
		v15 = v13 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v15)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
		v19 = v17 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v19)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
		v23 = v21 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v23)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
		v27 = v25 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v27)
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
		v31 = v29 ^ v10
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v31)
		return v5
	}
}
func F_macaddr_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(1455)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v9 == int32(1) {
		v12 = int32(4476144)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v15
		v18 = F_palloc(m, int32(40))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v22)
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
			F_initHyperLogLog(m, v18+int32(16), int32(10))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(1455)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(1456)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1457)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v18
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v13
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
