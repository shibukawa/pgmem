package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildWaitSnapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L55
	} else {
		goto L66
	}
L2:
	;
	v13 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v170 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v13<<(uint(int32(2))%32))))
	if base.Ui32(v19) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if v139 != 0 {
		goto L1
	} else {
		goto L47
	}
L8:
	;
	v139 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	if v30 == v19 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v139 = int32(1)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v34 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v139 = v131
	goto L7
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	if v38 == int32(0) {
		v131 = int32(0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v102 = int32(0)
	v104 = v34 - int32(1)
	goto L37
L18:
	;
	v43 = v38
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v48 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v131 = int32(0)
	goto L14
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v95 != 0 {
		v43 = v95
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v51 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v54 = int32(1)
	if v19 == v51 {
		v131 = v54
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v58 = v56 - int32(1)
	if v58 < int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v63 = int32(0)
	v65 = v58
	goto L26
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v71 = int32(2)
	v72 = base.I32_div_s(v65-v63, v71)
	v73 = v72 + v63
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69+v73<<(uint(v71)%32))))
	if v77 == v19 {
		v131 = v54
		goto L14
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	v81 = F_TransactionIdPrecedes(m, v77, v19)
	mBase = m.M
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = v73 + int32(1)
	goto L31
L30:
	;
	v82 = v63
	goto L31
L31:
	;
	if v81 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = v65
	goto L34
L33:
	;
	v85 = v73 - int32(1)
	goto L34
L34:
	;
	if v82 <= v85 {
		v63 = v82
		v65 = v85
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	goto L20
L37:
	;
	v109 = int32(2)
	v110 = base.I32_div_s(v104-v102, v109)
	v111 = v110 + v102
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v100+v111<<(uint(v109)%32))))
	v116 = base.B2i32(v115 == v19)
	if v115 == v19 {
		v131 = v116
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v131 = v116
	goto L14
L39:
	;
	v119 = base.B2i32(base.Ui32(v115) < base.Ui32(v19))
	if base.Ui32(v115) < base.Ui32(v19) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = v111 + int32(1)
	goto L42
L41:
	;
	v120 = v102
	goto L42
L42:
	;
	if base.Ui32(v115) < base.Ui32(v19) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v123 = v104
	goto L45
L44:
	;
	v123 = v111 - int32(1)
	goto L45
L45:
	;
	if v120 <= v123 {
		v102 = v120
		v104 = v123
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v19)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v151 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v151 = base.B2i32(base.Ui32(l1) < base.Ui32(v19))
	goto L48
L50:
	;
	goto L51
L51:
	;
	v151 = base.B2i32(int32(0) < v19-l1)
	goto L48
L52:
	;
	v154 = int32(0)
	F_XactLockTableWait(m, v19, v154, v154, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v160 = v13 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v160 < v161 {
		v13 = v160
		goto L5
	} else {
		goto L57
	}
L55:
	;
	return
L56:
	;
	goto L54
L57:
	;
	goto L6
L58:
	;
	if v180 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+316))
	v178 = base.B2i32(v176 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v178)
	v180 = v178
	goto L61
L60:
	;
	v180 = int32(0)
	goto L61
L61:
	;
	goto L58
L62:
	;
	v183 = F_LogStandbySnapshot(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L55
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	return
L65:
	;
	goto L64
L66:
	;
	F_errmsg_internal(m, int32(155432), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(491980), int32(1449), int32(86035))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
