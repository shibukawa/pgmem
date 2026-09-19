package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	v11 = m.G0
	v13 = v11 - int32(288)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = v13 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+240)) = v18
	v21 = v13 + int32(264)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+244)) = v21
	v24 = v13 + int32(262)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+248)) = v24
	v27 = v13 + int32(284)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+224)) = v27
	v30 = v13 + int32(280)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+228)) = v30
	v33 = v13 + int32(276)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+232)) = v33
	v36 = v13 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+236)) = v36
	v41 = F_sscanf(m, v16, int32(_a_F_macaddr_in_0), v13+int32(224))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(288)
	return v202
L2:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	if base.Ui32(int32(255)) < base.Ui32(v152) {
		goto L24
	} else {
		goto L25
	}
L3:
	;
	return int32(0)
L4:
	;
	if v41 == int32(6) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+216)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+212)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+208)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+204)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = v27
	v57 = F_sscanf(m, v16, int32(_a_F_macaddr_in_4), v13+int32(192))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v57 == int32(6) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+176)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+172)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v13)+168)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+164)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+160)) = v27
	v71 = F_sscanf(m, v16, int32(_a_F_macaddr_in_5), v13+int32(160))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v71 == int32(6) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+152)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+148)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v27
	v85 = F_sscanf(m, v16, int32(_a_F_macaddr_in_6), v13+int32(128))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v85 == int32(6) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+116)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v27
	v99 = F_sscanf(m, v16, int32(_a_F_macaddr_in_7), v13+int32(96))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v99 == int32(6) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v27
	v113 = F_sscanf(m, v16, int32(_a_F_macaddr_in_8), v13-int32(-64))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	if v113 == int32(6) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v27
	v127 = F_sscanf(m, v16, int32(_a_F_macaddr_in_9), v13+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if v127 == int32(6) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v131 = int32(0)
	v132 = F_errsave_start(m, v15)
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
		v202 = v131
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_macaddr_in_10)
	F_errmsg(m, int32(_a_F_macaddr_in_11), v13+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v15, int32(_a_F_macaddr_in_2), int32(95), int32(_a_F_macaddr_in_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v202 = v131
	goto L1
L23:
	;
	v188 = F_palloc(m, int32(6))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L36
	}
L24:
	;
	v170 = int32(0)
	v171 = F_errsave_start(m, v15)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+280))
	if base.Ui32(int32(255)) < base.Ui32(v155) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)+276))
	if base.Ui32(int32(255)) < base.Ui32(v158) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+272))
	if base.Ui32(int32(255)) < base.Ui32(v161) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v13)+268))
	if base.Ui32(int32(255)) < base.Ui32(v164) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v13)+264))
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
	if v171 == int32(0) {
		v202 = v170
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v16
	F_errmsg(m, int32(_a_F_macaddr_in_1), v13)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errsave_finish(m, v15, int32(_a_F_macaddr_in_2), int32(102), int32(_a_F_macaddr_in_3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v202 = v170
	goto L1
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v13)+284))
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v190)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v13)+280))
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)) = uint8(v192)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+276))
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)) = uint8(v194)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+272))
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)) = uint8(v196)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v13)+268))
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)) = uint8(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+264))
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)) = uint8(v200)
	v202 = v188
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13964(m, l0, int32(1441), int32(1440), int32(1439))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
