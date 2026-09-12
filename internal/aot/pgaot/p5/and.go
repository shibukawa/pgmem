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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if base.Ui32(l3) < base.Ui32(int32(8161)) {
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
	if base.Ui32(int32(8159)) < base.Ui32(l2) {
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
	v248 = m.ExcPending
	if v248 != 0 {
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
	v83 = v78 + int32(28)
	v85 = l1 - v25 + int32(4095)
	v86 = v83 + v85
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v87 != v34 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64+(v54^int32(-1))<<(uint(int32(2))%32))))
	v78 = v70
	goto L13
L15:
	;
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v78 = v72 + v54<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L17:
	;
	if v177 != 0 {
		goto L48
	} else {
		goto L49
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v34)
	v97 = v85
	goto L21
L19:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if base.Ui32(v89) < base.Ui32(v34) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v177 = int32(0)
	goto L17
L21:
	;
	v99 = int32(1)
	v100 = v97 - v99
	v101 = int32(2)
	v102 = base.I32_div_s(v100, v101)
	v104 = v102 << (uint(v99) % 32)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v104)+1)))
	v108 = v104 + v101
	if base.Ui32(v108) <= base.Ui32(int32(8163)) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if base.Ui32(v127) < base.Ui32(v34) {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v112 = v106 & int32(255)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v108))))
	if base.Ui32(v114) < base.Ui32(v112) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v117 = v106
	goto L25
L25:
	;
	v119 = v83 + v102
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v120 != v117&int32(255) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v116 = v112
	goto L28
L27:
	;
	v116 = v114
	goto L28
L28:
	;
	v117 = v116
	goto L25
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v117)
	if int32(1) < v100 {
		v97 = v102
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
	v133 = int32(4094)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v177 = int32(1)
	goto L17
L36:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v133) {
		v155 = int32(0)
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v156 = v83 + v133
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v157 != v155&int32(255) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v140 = v133 << (uint(int32(1)) % 32)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v140)+1)))
	if v133 == int32(4081) {
		v155 = v142
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v146 = v142 & int32(255)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+(v140+int32(2))))))
	if base.Ui32(v150) < base.Ui32(v146) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v152 = v146
	goto L43
L42:
	;
	v152 = v150
	goto L43
L43:
	;
	v155 = v152
	goto L38
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v155)
	goto L46
L45:
	;
	goto L46
L46:
	;
	if v133 != 0 {
		v133 = v133 - int32(1)
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
	v180 = m.ExcPending
	if v180 != 0 {
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
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v185 = F_fsm_search_avail(m, v54, v42, base.B2i32(v181 == int32(0)), int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	v188 = int32(-1)
	goto L54
L54:
	;
	F_UnlockReleaseBuffer(m, v54)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	v188 = v185
	goto L54
L56:
	;
	m.G0 = v45 + int32(16)
	if v188 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v13 + int32(48)
	return v240
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v198 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v237 = F_fsm_search(m, l0, v42)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L72
	}
L61:
	;
	v224 = v198
	goto L63
L62:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v200
	v202 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v202
	v206 = F_smgropen(m, v13+int32(16), v199)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L64
	}
L63:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	v228 = v188&int32(65535) + v25
	if base.B2i32(v225 != int32(-1))&base.B2i32(base.Ui32(v228) < base.Ui32(v225)) != 0 {
		v240 = v228
		goto L57
	} else {
		goto L69
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v206
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206)+72))
	if v210 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v224 = v222
	goto L63
L66:
	;
	v218 = v210
	goto L68
L67:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)+76))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v206)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v206)+72))
	v218 = v216
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+72)) = v218 + int32(1)
	goto L65
L69:
	;
	v232 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	if base.Ui32(v228) < base.Ui32(v232) {
		v240 = v228
		goto L57
	} else {
		goto L71
	}
L71:
	;
	goto L60
L72:
	;
	v240 = v237
	goto L57
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
	F_errmsg_internal(m, int32(37144), v13)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(494061), int32(438), int32(111410))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
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
