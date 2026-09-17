package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RecordAndGetPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if base.Ui32(l3) < base.Ui32(int32(_a_F_RecordAndGetPageWithFreeSpace_0)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(4069)
	v18 = base.I32_div_u_s(l1, v17)
	v21 = base.I64_extend_i32_u(v18) << (uint(int64(32)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v21
	v25 = v18 * v17
	if base.Ui32(int32(_a_F_RecordAndGetPageWithFreeSpace_1)) < base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L73
	}
L4:
	;
	v32 = int32(-1)
	goto L6
L5:
	;
	v32 = int32(base.Ui32(l2) >> (uint(int32(5)) % 32))
	goto L6
L6:
	;
	v34 = v32 & int32(255)
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v40 = int32(base.Ui32(l3+int32(31)) >> (uint(int32(5)) % 32))
	goto L9
L8:
	;
	v40 = int32(1)
	goto L9
L9:
	;
	v42 = v40 & int32(255)
	v43 = m.G0
	v45 = v43 - int32(16)
	m.G0 = v45
	v48 = v13 + int32(32)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v49
	v54 = F_fsm_readbuf(m, l0, v45+int32(8), int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	F_LockBuffer(m, v54, int32(2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if v54 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v84 = v78 + int32(28)
	v86 = l1 - v25 + int32(4095)
	v87 = v84 + v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 != v34 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_RecordAndGetPageWithFreeSpace[0]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64+(v54^int32(-1))<<(uint(int32(2))%32))))
	v78 = v70
	goto L13
L15:
	;
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_RecordAndGetPageWithFreeSpace[1]))
	v78 = v72 + v54<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L17:
	;
	if v180 != 0 {
		goto L48
	} else {
		goto L49
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v34)
	v97 = v86
	goto L21
L19:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if base.Ui32(v90) < base.Ui32(v34) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v180 = int32(0)
	goto L17
L21:
	;
	v101 = int32(1)
	v102 = v97 - v101
	v103 = int32(2)
	v104 = base.I32_div_s(v102, v103)
	v106 = v104 << (uint(v101) % 32)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v106)+1)))
	v110 = v106 + v103
	if base.Ui32(v110) <= base.Ui32(int32(_a_F_RecordAndGetPageWithFreeSpace_2)) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if base.Ui32(v129) < base.Ui32(v34) {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v114 = v108 & int32(255)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v110))))
	if base.Ui32(v116) < base.Ui32(v114) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v119 = v108
	goto L25
L25:
	;
	v121 = v104 + v84
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v122 != v119&int32(255) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v118 = v114
	goto L28
L27:
	;
	v118 = v116
	goto L28
L28:
	;
	v119 = v118
	goto L25
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v119)
	if int32(1) < v102 {
		v97 = v104
		goto L21
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L22
L32:
	;
	goto L31
L33:
	;
	v135 = int32(4094)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v180 = int32(1)
	goto L17
L36:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v135) {
		v156 = int32(0)
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v157 = v135 + v84
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v158 != v156&int32(255) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v145 = v135 << (uint(int32(1)) % 32)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+int32(29)+v145))))
	if v135 == int32(4081) {
		v156 = v147
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v84)+2)))
	if base.Ui32(v151) < base.Ui32(v147) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v153 = v147
	goto L43
L42:
	;
	v153 = v151
	goto L43
L43:
	;
	v156 = v153
	goto L38
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v156)
	goto L46
L45:
	;
	goto L46
L46:
	;
	if v135 != 0 {
		v135 = v135 - int32(1)
		goto L36
	} else {
		goto L47
	}
L47:
	;
	goto L37
L48:
	;
	F_MarkBufferDirtyHint(m, v54, int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v42 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v188 = F_fsm_search_avail(m, v54, v42, base.B2i32(v184 == int32(0)), int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	v191 = int32(-1)
	goto L54
L54:
	;
	F_UnlockReleaseBuffer(m, v54)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	v191 = v188
	goto L54
L56:
	;
	m.G0 = v45 + int32(16)
	if v191 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v13 + int32(48)
	return v243
L58:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v201 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v240 = F_fsm_search(m, l0, v42)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L72
	}
L61:
	;
	v227 = v201
	goto L63
L62:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v203
	v205 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v205
	v209 = F_smgropen(m, v13+int32(16), v202)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L64
	}
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+20))
	v231 = v191&int32(_a_F_RecordAndGetPageWithFreeSpace_3) + v25
	if base.B2i32(v228 != int32(-1))&base.B2i32(base.Ui32(v231) < base.Ui32(v228)) != 0 {
		v243 = v231
		goto L57
	} else {
		goto L69
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v209
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+72))
	if v213 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v227 = v225
	goto L63
L66:
	;
	v221 = v213
	goto L68
L67:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v209)+72))
	v221 = v219
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+72)) = v221 + int32(1)
	goto L65
L69:
	;
	v235 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	if base.Ui32(v231) < base.Ui32(v235) {
		v243 = v231
		goto L57
	} else {
		goto L71
	}
L71:
	;
	goto L60
L72:
	;
	v243 = v240
	goto L57
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
	F_errmsg_internal(m, int32(_a_F_RecordAndGetPageWithFreeSpace_4), v13)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_RecordAndGetPageWithFreeSpace_5), int32(438), int32(_a_F_RecordAndGetPageWithFreeSpace_6))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
