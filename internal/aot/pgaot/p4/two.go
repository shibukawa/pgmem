package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessTwoPhaseBuffer(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int64
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessTwoPhaseBuffer[0]))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v19 = base.I32_wrap_i64(l0)
	v20 = F_TransactionIdDidCommit(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L70
	}
L2:
	;
	m.G0 = v14 + int32(96)
	return v182
L3:
	;
	if base.Ui64(v18) <= base.Ui64(l0) {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	return int32(0)
L5:
	;
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = F_TransactionIdDidAbort(m, v19)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v32 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	if v26 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v32 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v32 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v19
	v36 = int64(base.Ui64(l0) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v36)
	F_errmsg(m, int32(_a_F_ProcessTwoPhaseBuffer_0), v14)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_RemoveTwoPhaseFile(m, l0, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(_a_F_ProcessTwoPhaseBuffer_1), int32(2218), int32(_a_F_ProcessTwoPhaseBuffer_2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v182 = v6
	goto L2
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v19
	v51 = int64(base.Ui64(l0) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+20)) = uint32(v51)
	F_errmsg(m, int32(_a_F_ProcessTwoPhaseBuffer_3), v14+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_PrepareRedoRemoveFull(m, l0, int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ProcessTwoPhaseBuffer_1), int32(2226), int32(_a_F_ProcessTwoPhaseBuffer_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v182 = v6
	goto L2
L27:
	;
	v69 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l2 != 0 {
		goto L47
	} else {
		goto L48
	}
L30:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v69 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v69 != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v19
	v73 = int64(base.Ui64(l0) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+68)) = uint32(v73)
	F_errmsg(m, int32(_a_F_ProcessTwoPhaseBuffer_4), v14-int32(-64))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_RemoveTwoPhaseFile(m, l0, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	F_errfinish(m, int32(_a_F_ProcessTwoPhaseBuffer_1), int32(2240), int32(_a_F_ProcessTwoPhaseBuffer_2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v182 = v6
	goto L2
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v19
	v90 = int64(base.Ui64(l0) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+84)) = uint32(v90)
	F_errmsg(m, int32(_a_F_ProcessTwoPhaseBuffer_5), v14+int32(80))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_PrepareRedoRemoveFull(m, l0, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	F_errfinish(m, int32(_a_F_ProcessTwoPhaseBuffer_1), int32(2248), int32(_a_F_ProcessTwoPhaseBuffer_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v182 = v6
	goto L2
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	if v140 <= int32(0) {
		v182 = v139
		goto L2
	} else {
		goto L58
	}
L47:
	;
	v106 = F_ReadTwoPhaseFile(m, l0, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_XlogReadTwoPhaseData(m, l1, v14+int32(92), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L56
	}
L50:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v108 == v19 {
		v139 = v106
		goto L46
	} else {
		goto L51
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v118 = int64(base.Ui64(l0) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+36)) = uint32(v118)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v19
	F_errmsg(m, int32(_a_F_ProcessTwoPhaseBuffer_6), v14+int32(32))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ProcessTwoPhaseBuffer_1), int32(2274), int32(_a_F_ProcessTwoPhaseBuffer_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v137 != v19 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v139 = v136
	goto L46
L58:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+54)))
	v154 = int32(0)
	goto L59
L59:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v139+(v143+int32(7))&int32(_a_F_ProcessTwoPhaseBuffer_7)+int32(72)+v154<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v182 = v139
	goto L2
L61:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if l3 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	F_SubTransSetParent(m, v166, v19)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v172 = v154 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	if v172 < v173 {
		v154 = v172
		goto L59
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	goto L60
L70:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v198 = int64(base.Ui64(l0) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+52)) = uint32(v198)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v19
	F_errmsg(m, int32(_a_F_ProcessTwoPhaseBuffer_8), v14+int32(48))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ProcessTwoPhaseBuffer_1), int32(2280), int32(_a_F_ProcessTwoPhaseBuffer_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
