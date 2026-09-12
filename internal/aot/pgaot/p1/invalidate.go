package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v16 & int32(-4194305)
	v20 = F_BufTableHashCode(m, v8)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v30 = v23 + v20&int32(127)<<(uint(int32(7))%32) + int32(6912)
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L75
	}
L4:
	;
	m.G0 = v8 + int32(48)
	return
L5:
	;
	v37 = F_LWLockAcquire(m, v30, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	if v71&int32(33554432) != 0 {
		goto L69
	} else {
		goto L70
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(219672)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(476033)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v50 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v49 | v50
	if v49&v50 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v71 = v49
	goto L10
L10:
	;
	v78 = int32(4069244)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(24))+8))
	if v81 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	F_perform_spin_delay(m, v8+int32(24))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v71 = v64
	goto L10
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v65 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v64 | v65
	if v64&v65 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v98 != v99 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v96
	goto L16
L18:
	;
	if int32(999) < v79 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v79 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v86 = int32(900)
	if v86 <= v79 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v89 = v86
	goto L24
L23:
	;
	v89 = v79
	goto L24
L24:
	;
	v96 = v89 + int32(100)
	goto L17
L25:
	;
	v96 = v79 - int32(1)
	goto L17
L26:
	;
	if v71&int32(262143) != 0 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v71 & int32(-4194305)
	F_LWLockRelease(m, v30)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v101 != v102 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v104 != v105 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v107 != v108 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v110 == v111 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	goto L4
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v71 & int32(-4194305)
	F_LWLockRelease(m, v30)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L6
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v127 = v125 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v127
	v130 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	if v127 == v130 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L68
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if int32(0) < v176 {
		goto L3
	} else {
		goto L67
	}
L40:
	;
	v175 = int32(4365360)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[267]))
	if v127 == v134 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = int32(4365368)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	if v127 == v138 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = int32(4365376)
	goto L39
L47:
	;
	goto L48
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[269]))
	if v127 == v142 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v175 = int32(4365384)
	goto L39
L50:
	;
	goto L51
L51:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[270]))
	if v127 == v146 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = int32(4365392)
	goto L39
L53:
	;
	goto L54
L54:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if v127 == v150 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v175 = int32(4365400)
	goto L39
L56:
	;
	goto L57
L57:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	if v127 == v154 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = int32(4365408)
	goto L39
L59:
	;
	goto L60
L60:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if v127 == v158 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v175 = int32(4365416)
	goto L39
L62:
	;
	goto L63
L63:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[274]))
	if v162 == int32(0) {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v169 = int32(0)
	v171 = F_hash_search(m, v166, v8+int32(24), v169, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v171 == int32(0) {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	v175 = v171
	goto L39
L67:
	;
	goto L38
L68:
	;
	goto L5
L69:
	;
	F_BufTableDelete(m, v8, v20)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_LWLockRelease(m, v30)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_StrategyFreeBuffer(m, l0)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	F_errmsg_internal(m, int32(217229), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(476033), int32(2236), int32(217249))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InvalidateTSCacheCallBack(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	F_hash_seq_init(m, v6+int32(12), l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1170]))
	if l0 == v29 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v19)
	v23 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v23 != 0 {
		v18 = v23
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1171])) = int32(0)
	goto L13
L12:
	;
	goto L13
L13:
	;
	m.G0 = v6 + int32(32)
	return
}
