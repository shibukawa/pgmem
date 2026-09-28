package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cfb_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L42
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = l1
	v27 = l2
	v28 = l3
	v31 = v24
	goto L4
L4:
	;
	if int32(0) < v31 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v64 = l0 + int32(24)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = v26
	v68 = v27
	v69 = v28
	v73 = v65
	goto L21
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = v40 - v31
	if v27 < v41 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	v43 = v27
	goto L11
L10:
	;
	v43 = v41
	goto L11
L11:
	;
	v44 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v26, v43, v28)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v46 = v27 - v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 == v48 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v47 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v54 = v47
	goto L16
L16:
	;
	if int32(0) < v46 {
		v26 = v26 + v44
		v27 = v46
		v28 = v28 + v44
		v31 = v54
		goto L4
	} else {
		goto L20
	}
L17:
	;
	base.MemoryCopy(m, l0+int32(24), l0+int32(88), v47)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v51
	v54 = v51
	goto L16
L20:
	;
	goto L2
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v85 = m.T0[v84].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v80, int32(0), v64, v73, l0+int32(56), v16+int32(12))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L2
L23:
	;
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v87 = l5
	goto L26
L25:
	;
	v87 = int32(1)
	goto L26
L26:
	;
	if v87 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v90 <= int32(4) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v90 + int32(1)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 < v96 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v98 = v68
	goto L33
L32:
	;
	v98 = v96
	goto L33
L33:
	;
	v99 = m.T0[l4].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v67, v98, v69)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v101 = v68 - v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 == v103 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v102 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if int32(0) < v101 {
		v67 = v67 + v99
		v68 = v101
		v69 = v69 + v99
		v73 = v103
		goto L21
	} else {
		goto L41
	}
L38:
	;
	base.MemoryCopy(m, v64, l0+int32(88), v102)
	goto L40
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L37
L41:
	;
	goto L22
L42:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	if v85 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v164
	F_errmsg(m, int32(_a_F_cfb_process_0), v16)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L58
	}
L45:
	;
	v164 = int32(_a_F_cfb_process_1)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v143 = int32(_a_F_cfb_process_2)
	goto L49
L48:
	;
	v164 = v158
	goto L44
L49:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v85 != v146 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v158 = v156
	goto L48
L51:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	if v148 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v164 = int32(_a_F_cfb_process_3)
	goto L44
L55:
	;
	goto L56
L56:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	if v85 != v152 {
		v143 = v143 + int32(16)
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v158 = v148
	goto L48
L58:
	;
	F_errfinish(m, int32(_a_F_cfb_process_4), int32(239), int32(_a_F_cfb_process_5))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
