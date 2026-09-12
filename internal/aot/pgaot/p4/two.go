package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessTwoPhaseBuffer(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[180]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = F_TransactionIdDidCommit(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L74
	}
L2:
	;
	m.G0 = v12 + int32(96)
	return v173
L3:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v16))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L4:
	;
	return int32(0)
L5:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = F_TransactionIdDidAbort(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v29 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	if v23 == int32(0) {
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
	if v29 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v29 != 0 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(45169), v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(499009), int32(2200), int32(226791))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v173 = v6
	goto L2
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg(m, int32(44890), v12+int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_PrepareRedoRemove(m, l0, int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	F_errfinish(m, int32(499009), int32(2207), int32(226791))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v173 = v6
	goto L2
L27:
	;
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v68 = base.B2i32(base.Ui32(v16) <= base.Ui32(l0))
	goto L27
L29:
	;
	goto L30
L30:
	;
	v68 = base.B2i32(int32(0) <= l0-v16)
	goto L27
L31:
	;
	v71 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l2 != 0 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	if l2 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v71 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if v71 != 0 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg(m, int32(45113), v12+int32(32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_RemoveTwoPhaseFile(m, l0, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(499009), int32(2220), int32(226791))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v173 = v6
	goto L2
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l0
	F_errmsg(m, int32(44827), v12+int32(48))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_PrepareRedoRemove(m, l0, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	F_errfinish(m, int32(499009), int32(2227), int32(226791))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v173 = v6
	goto L2
L50:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+28))
	if v133 <= int32(0) {
		v173 = v132
		goto L2
	} else {
		goto L62
	}
L51:
	;
	v102 = F_ReadTwoPhaseFile(m, l0, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_XlogReadTwoPhaseData(m, l1, v12+int32(92), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L60
	}
L54:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if v104 == l0 {
		v132 = v102
		goto L50
	} else {
		goto L55
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	F_errmsg(m, int32(45224), v12-int32(-64))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(499009), int32(2252), int32(226791))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v130 != l0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v132 = v129
	goto L50
L62:
	;
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+54)))
	v147 = int32(0)
	goto L63
L63:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v132+(v136+int32(7))&int32(131064)+int32(72)+v147<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v173 = v132
	goto L2
L65:
	;
	F_AdvanceNextFullTransactionIdPastXid(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if l3 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	F_SubTransSetParent(m, v157, l0)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v163 = v147 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v132)+28))
	if v163 < v164 {
		v147 = v163
		goto L63
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	goto L64
L74:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = l0
	F_errmsg(m, int32(44772), v12+int32(80))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(499009), int32(2257), int32(226791))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
