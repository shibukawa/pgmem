package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetSequence(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v3 = l2
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	F_init_sequence(m, l0, v11+int32(108), v11+int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[0]))
	v24 = F_pg_class_aclcheck(m, v20, v22, int64(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L57
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = F_SearchSysCache1(m, int32(61), base.I64_extend_i32_u(l0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L50
	}
L9:
	;
	if v30 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = v34 + v35
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+32))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v36)+24))
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_PreventCommandIfReadOnly(m, int32(_a_F_SetSequence_3))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_PreventCommandIfParallelMode(m, int32(_a_F_SetSequence_3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v55 = F_read_seq_tuple(m, v41, v11+int32(100), v11+int32(80))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if base.B2i32(l1 < v37)|base.B2i32(v38 < l1) != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+118)))
	if v69 != int32(112) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v66 = v62
	goto L19
L21:
	;
	goto L22
L22:
	;
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v63)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = l1
	v66 = l1
	goto L19
L23:
	;
	v80 = int32(_a_F_SetSequence_5)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_SetSequence[1])) = v82 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+16)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v55)+8)) = int64(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	F_MarkBufferDirty(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[2]))
	if v73 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	if v76 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v78 = F_GetTopTransactionId(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	if v77 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L23
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+118)))
	if v94 != int32(112) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v147 = int32(_a_F_SetSequence_5)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_SetSequence[1])) = v149 - int32(1)
	F_UnlockReleaseBuffer(m, v90)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L48
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[2]))
	if v98 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	if v101 != 0 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v90 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	if v102 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[3]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106+(v90^int32(-1))<<(uint(int32(2))%32))))
	v120 = v112
	goto L39
L41:
	;
	goto L42
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_SetSequence[4]))
	v120 = v114 + v90<<(uint(int32(13))%32) + int32(-8192)
	goto L39
L43:
	;
	F_XLogRegisterBuffer(m, int32(0), v90, int32(6))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v127
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v129
	F_XLogRegisterData(m, v11-int32(-64), int32(12))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	F_XLogRegisterData(m, v136, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v142 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_rotl(v142, int64(32))
	goto L32
L48:
	;
	F_relation_close(m, v41, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	m.G0 = v11 + int32(112)
	return
L50:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v169 + int32(4)
	F_errmsg(m, int32(_a_F_SetSequence_6), v11+int32(48))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_SetSequence_1), int32(965), int32(_a_F_SetSequence_2))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_SetSequence_0), v11)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_SetSequence_1), int32(969), int32(_a_F_SetSequence_2))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v203 + int32(4)
	F_errmsg(m, int32(_a_F_SetSequence_4), v11+int32(16))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_SetSequence_1), int32(994), int32(_a_F_SetSequence_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
