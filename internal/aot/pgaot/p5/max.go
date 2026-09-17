package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_set_max_safe_fds(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v1 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[0]))
	v16 = F_palloc(m, int32(_a_F_set_max_safe_fds_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = v9 + int32(-16)
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v59 = int32(1)
	if v14 <= v59 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	if v41 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L5:
	;
	switch int32(4) {
	case 0:
		goto L10
	default:
		goto L9
	case 4:
		goto L11
	}
L6:
	;
	goto L7
L7:
	;
	v41 = F___syscall_ret(m, int32(0))
	mBase = m.M
	goto L4
L8:
	;
	goto L7
L9:
	;
	v34 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v34
	goto L8
L10:
	;
	v28 = m.G2
	v29 = m.G1
	v31 = base.I64_extend_i32_u(v28 - v29)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v31
	goto L8
L11:
	;
	v24 = int64(4096)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v24
	goto L8
L12:
	;
	v46 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_set_max_safe_fds_1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_set_max_safe_fds_2), int32(988), int32(_a_F_set_max_safe_fds_3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	v62 = v59
	goto L19
L18:
	;
	v62 = v14
	goto L19
L19:
	;
	v66 = v1
	v67 = v16
	v68 = int32(1024)
	v69 = v1
	goto L22
L20:
	;
	F_pfree(m, v156)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L52
	}
L21:
	;
	v137 = int32(0)
	goto L49
L22:
	;
	if v41 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v131 = v62
	v132 = v120
	v134 = v127
	goto L21
L24:
	;
	if v68 <= v66 {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	if v66 != 0 {
		v131 = v66
		v132 = v67
		v134 = v69
		goto L21
	} else {
		goto L40
	}
L26:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	if base.Ui64(v75-int64(1)) <= base.Ui64(base.I64_extend_i32_u(v69)) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v80 = m.Env.X__syscall_dup(m, int32(2))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v80) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L28
L30:
	;
	if int32(0) <= v88 {
		goto L24
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[1])) = int32(0) - v80
	v88 = int32(-1)
	goto L33
L32:
	;
	v88 = v80
	goto L33
L33:
	;
	goto L30
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[1]))
	switch v92 - int32(33) {
	case 0, 8:
		goto L25
	default:
		goto L35
	}
L35:
	;
	v97 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v97 == int32(0) {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v66
	F_errmsg_internal(m, int32(_a_F_set_max_safe_fds_4), v9+int32(-32))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_set_max_safe_fds_2), int32(1011), int32(_a_F_set_max_safe_fds_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L25
L40:
	;
	v155 = v66
	v156 = v67
	v158 = v69
	goto L20
L41:
	;
	v116 = F_repalloc(m, v67, v68<<(uint(int32(3))%32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v120 = v67
	v121 = v68
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120+v66<<(uint(int32(2))%32)))) = v88
	if v88 < v69 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v120 = v116
	v121 = v68 << (uint(int32(1)) % 32)
	goto L43
L45:
	;
	v127 = v69
	goto L47
L46:
	;
	v127 = v88
	goto L47
L47:
	;
	v129 = v66 + int32(1)
	if v129 != v62 {
		v66 = v129
		v67 = v120
		v68 = v121
		v69 = v127
		goto L22
	} else {
		goto L48
	}
L48:
	;
	goto L23
L49:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v132+v137<<(uint(int32(2))%32))))
	v149 = F_close(m, v148)
	mBase = m.M
	v151 = v137 + int32(1)
	if v151 != v131 {
		v137 = v151
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v155 = v131
	v156 = v132
	v158 = v134
	goto L20
L51:
	;
	goto L50
L52:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[0]))
	if v155 < v165 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v167 = v155
	goto L55
L54:
	;
	v167 = v165
	goto L55
L55:
	;
	v169 = v167 - int32(10)
	*(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[2])) = v169
	v173 = v158 - v155 + int32(1)
	if int32(47) < v169 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v178 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L65
	}
L59:
	;
	if v178 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v173
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v183
	F_errmsg_internal(m, int32(_a_F_set_max_safe_fds_5), v9+int32(-48))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	m.G0 = v11 - int32(-64)
	return
L63:
	;
	F_errfinish(m, int32(_a_F_set_max_safe_fds_2), int32(1086), int32(_a_F_set_max_safe_fds_6))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_set_max_safe_fds_7), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(58)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v173
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_set_max_safe_fds[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v213 + int32(10)
	F_errdetail(m, int32(_a_F_set_max_safe_fds_8), v11)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_set_max_safe_fds_2), int32(1083), int32(_a_F_set_max_safe_fds_6))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
