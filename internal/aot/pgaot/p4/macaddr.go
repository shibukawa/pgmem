package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int64
	_ = v203
	v11 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(288)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = v14 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v19
	v22 = v14 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = v22
	v25 = v14 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+248)) = v25
	v28 = v14 + int32(284)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v28
	v31 = v14 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v31
	v34 = v14 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v34
	v37 = v14 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v37
	v42 = F_sscanf(m, v17, int32(_a_F_macaddr_in_0), v14+int32(224))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(288)
	return v203
L2:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v14)+284))
	if base.Ui32(int32(255)) < base.Ui32(v152) {
		goto L24
	} else {
		goto L25
	}
L3:
	;
	return int64(0)
L4:
	;
	if v42 == int32(6) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+216)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+204)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+200)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v28
	v58 = F_sscanf(m, v17, int32(_a_F_macaddr_in_4), v14+int32(192))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v58 == int32(6) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+184)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+172)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+168)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v28
	v72 = F_sscanf(m, v17, int32(_a_F_macaddr_in_5), v14+int32(160))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v72 == int32(6) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+148)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+140)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v28
	v86 = F_sscanf(m, v17, int32(_a_F_macaddr_in_6), v14+int32(128))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v86 == int32(6) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v28
	v100 = F_sscanf(m, v17, int32(_a_F_macaddr_in_7), v14+int32(96))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v100 == int32(6) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v28
	v114 = F_sscanf(m, v17, int32(_a_F_macaddr_in_8), v14-int32(-64))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	if v114 == int32(6) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v28
	v128 = F_sscanf(m, v17, int32(_a_F_macaddr_in_9), v14+int32(32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if v128 == int32(6) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v132 = F_errsave_start(m, v16)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v132 == int32(0) {
		v203 = v11
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_macaddr_in_10)
	F_errmsg(m, int32(_a_F_macaddr_in_11), v14+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v16, int32(_a_F_macaddr_in_2), int32(84), int32(_a_F_macaddr_in_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v203 = v11
	goto L1
L23:
	;
	v187 = F_palloc(m, int32(6))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L36
	}
L24:
	;
	v170 = F_errsave_start(m, v16)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+280))
	if base.Ui32(int32(255)) < base.Ui32(v155) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	if base.Ui32(int32(255)) < base.Ui32(v158) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	if base.Ui32(int32(255)) < base.Ui32(v161) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v14)+268))
	if base.Ui32(int32(255)) < base.Ui32(v164) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v14)+264))
	if base.Ui32(v167) < base.Ui32(int32(256)) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	if v170 == int32(0) {
		v203 = v11
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
	F_errmsg(m, int32(_a_F_macaddr_in_1), v14)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errsave_finish(m, v16, int32(_a_F_macaddr_in_2), int32(91), int32(_a_F_macaddr_in_3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v203 = v11
	goto L1
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v14)+284))
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v189)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)+280))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)) = uint8(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+276))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+272))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)) = uint8(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+268))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)) = uint8(v197)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v14)+264))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)) = uint8(v199)
	v203 = base.I64_extend_i32_u(v187)
	goto L1
}
func F_macaddr_not(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(6))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
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
		return base.I64_extend_i32_u(v5)
	}
}
func F_macaddr_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1586)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+20)))
	if v7 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v2)+32)) = int32(1586)
		*(*int32)(unsafe.Add(mBase, uint32(v2)+28)) = int32(1587)
		*(*int32)(unsafe.Add(mBase, uint32(v2)+24)) = int32(1588)
		*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(118)
	} else {
	}
	return int64(0)
}
