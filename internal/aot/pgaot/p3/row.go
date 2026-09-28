package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SendRowDescriptionMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v280 int32
	_ = v280
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v24 = v22
	goto L3
L2:
	;
	v24 = int32(0)
	goto L3
L3:
	;
	F_resetStringInfo(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(84)
	goto L4
L4:
	;
	F_enlargeStringInfo(m, l0, int32(2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = int32(8)
	v40 = v21<<(uint(v34)%32) | int32(base.Ui32(v21&int32(_a_F_SendRowDescriptionMessage_0))>>(uint(v34)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v31+v32))) = uint16(v40)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31 + int32(2)
	F_enlargeStringInfo(m, l0, v21*int32(274))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if int32(0) < v21 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v24
	v64 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pq_endmessage_reuse(m, l0)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L42
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = l1 + v67<<(uint(int32(3))%32) + v64*int32(100)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+96))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v75
	v79 = F_getBaseTypeAndTypmod(m, v74, v19+int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if l3 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v153 = int32(0)
	v164 = v153
	v168 = v153
	v172 = v153
	goto L14
L16:
	;
	v89 = v57
	goto L17
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+26)))
	if v100 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+24)))
	v112 = int32(8)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v118 = int32(16711935)
	v128 = v89 + int32(4)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v128) < base.Ui32(v130+v131<<(uint(int32(2))%32)) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v104 = v89 + int32(4)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v104) < base.Ui32(v105+v106<<(uint(int32(2))%32)) {
		v89 = v104
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L15
L23:
	;
	v136 = v128
	goto L25
L24:
	;
	v136 = int32(0)
	goto L25
L25:
	;
	v164 = base.I32_rotr(v117&v118, v112) | base.I32_rotr(v117, int32(24))&v118
	v168 = v111<<(uint(v112)%32) | int32(base.Ui32(v111)>>(uint(v112)%32))
	v172 = v136
	goto L14
L26:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v64<<(uint(int32(1))%32)))))
	v179 = int32(8)
	v186 = v178<<(uint(v179)%32) | int32(base.Ui32(v178)>>(uint(v179)%32))
	goto L28
L27:
	;
	v186 = int32(0)
	goto L28
L28:
	;
	v188 = v73 + int32(32)
	v189 = F_strlen(m, v188)
	mBase = m.M
	v190 = F_pg_server_to_client(m, v188, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	v210 = v208 + v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v210+v211))) = v164
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint16)(unsafe.Add(mBase, uint32(v214+v210)+4)) = uint16(v168)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v219 = int32(24)
	v221 = int32(16711935)
	v225 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v217+v210)+6)) = base.I32_rotr(v79, v219)&v221 | base.I32_rotr(v79&v221, v225)
	v230 = v210 + int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(28))+72)))
	v239 = v234<<(uint(v225)%32) | int32(base.Ui32(v234)>>(uint(v225)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v232+v230))) = uint16(v239)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v210)+12)) = base.I32_rotr(v243&v221, v225) | base.I32_rotr(v243, v219)&v221
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint16)(unsafe.Add(mBase, uint32(v254+v210)+16)) = uint16(v186)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v210 + int32(18)
	v261 = v64 + int32(1)
	if v261 != v21 {
		v57 = v172
		v64 = v261
		goto L11
	} else {
		goto L41
	}
L30:
	;
	if v190 != v188 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v194 = F_strlen(m, v190)
	mBase = m.M
	v196 = v194 + int32(1)
	if v196 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = v189 + int32(1)
	if v204 != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v197+v193, v190, v196)
	goto L36
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, v190)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v208 = v193
	v209 = v196
	goto L29
L38:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v205+v202, v190, v204)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v208 = v202
	v209 = v204
	goto L29
L41:
	;
	goto L12
L42:
	;
	m.G0 = v19 + int32(16)
	return
}
func F_row_is_in_frame(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v27 int64
	_ = v27
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v110 int64
	_ = v110
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v136 int32
	_ = v136
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+228))
	F_update_frameheadpos(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v15)+184))
	if l1 < v21 {
		v136 = v5
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v136
L4:
	;
	if v16&int32(1024) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v16&int32(_a_F_row_is_in_frame_0) != 0 {
		goto L39
	} else {
		goto L40
	}
L6:
	;
	if v16&int32(4) != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v16&int32(_a_F_row_is_in_frame_4) == int32(0) {
		goto L5
	} else {
		goto L27
	}
L9:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	if l1 <= v27 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v16&int32(10) == int32(0) {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v136 = int32(-1)
	goto L3
L13:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	if l1 <= v34 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if l3 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	if v42 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L16:
	;
	v38 = F_window_gettupleslot(m, l0, l1, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v38 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v136 = int32(-1)
	goto L3
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+380))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = l2
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	if v49 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	F_MemoryContextReset(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v55 = int32(_a_F_row_is_in_frame_3)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_row_is_in_frame[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_row_is_in_frame[0])) = v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	v63 = m.T0[v62].(func(*base.Module, int32, int32, int32) int64)(m, v49, v45, v13+int32(15))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_row_is_in_frame[0])) = v56
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	F_MemoryContextReset(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v63 != int64(0) {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v136 = int32(-1)
	goto L3
L27:
	;
	if v16&int32(4) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+248))
	if v16&int32(_a_F_row_is_in_frame_5) != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if v16&int32(10) == int32(0) {
		goto L5
	} else {
		goto L35
	}
L31:
	;
	v84 = int64(0) - v80
	goto L33
L32:
	;
	v84 = v80
	goto L33
L33:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	v88 = v84 + v87
	if base.B2i32(v84 < int64(0))^base.B2i32(v88 < v87)|base.B2i32(l1 <= v88) != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v136 = int32(-1)
	goto L3
L35:
	;
	F_update_frametailpos(m, v15)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v15)+192))
	if l1 < v100 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(-1)
	goto L3
L38:
	;
	v136 = int32(1)
	goto L3
L39:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	if l1 != v110 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v16&int32(_a_F_row_is_in_frame_1) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v136 = v5
	goto L3
L43:
	;
	if v16&int32(_a_F_row_is_in_frame_2) == int32(0) {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+96))
	if v123 == int32(0) {
		v136 = v5
		goto L3
	} else {
		goto L48
	}
L46:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	if l1 == v120 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v15)+352))
	if l1 < v126 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	F_update_grouptailpos(m, v15)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v15)+360))
	if l1 < v130 {
		v136 = v5
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L38
}
func F_row_security_policy_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v4 != int32(0))
L2:
	;
	goto L3
L3:
	;
	if v4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v39 - v40
L8:
	;
	goto L7
L9:
	;
	v24 = v6
	v25 = v4
	goto L10
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v39 = v29
	v40 = v28
	goto L8
L12:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
