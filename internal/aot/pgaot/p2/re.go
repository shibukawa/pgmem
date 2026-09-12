package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecScanReScan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	m.T0[v13].(func(*base.Module, int32))(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+156))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = v20 - int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v22))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))) = uint8(v27)
	goto L3
L6:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	switch v30 - int32(354) {
	case 0:
		v49 = int32(116)
		goto L8
	case 1:
		goto L9
	default:
		goto L10
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19+v49)))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v49 = int32(100)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v38
	F_errmsg_internal(m, int32(510877), v8)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(523101), int32(146), int32(299756))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	if v108 < int32(0) {
		goto L3
	} else {
		goto L25
	}
L15:
	;
	v108 = base.I32_ctz(v94) | v95<<(uint(int32(5))%32)
	goto L14
L16:
	;
	v108 = int32(-2)
	goto L14
L17:
	;
	v61 = base.I32_div_s(int32(0), int32(32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v62 <= v61 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v65 = v51 + int32(8)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v61<<(uint(int32(2))%32))))
	v72 = v69 & int32(-1)
	if v72 != 0 {
		v94 = v72
		v95 = v61
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v74 = v61 + int32(1)
	if v74 == v62 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v77 = v74
	goto L21
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65+v77<<(uint(int32(2))%32))))
	if v84 != 0 {
		v94 = v84
		v95 = v77
		goto L15
	} else {
		goto L23
	}
L22:
	;
	goto L16
L23:
	;
	v86 = v77 + int32(1)
	if v86 != v62 {
		v77 = v86
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v111 = v108
	goto L26
L26:
	;
	v117 = v111 - int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v117))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117+v118))) = uint8(v122)
	if v51 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L3
L28:
	;
	if int32(0) <= v179 {
		v111 = v179
		goto L26
	} else {
		goto L39
	}
L29:
	;
	v179 = base.I32_ctz(v165) | v166<<(uint(int32(5))%32)
	goto L28
L30:
	;
	v179 = int32(-2)
	goto L28
L31:
	;
	v130 = v111 + int32(1)
	v132 = base.I32_div_s(v130, int32(32))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v133 <= v132 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v136 = v51 + int32(8)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v132<<(uint(int32(2))%32))))
	v143 = v140 & (int32(-1) << (uint(v130) % 32))
	if v143 != 0 {
		v165 = v143
		v166 = v132
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v145 = v132 + int32(1)
	if v145 == v133 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v148 = v145
	goto L35
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v136+v148<<(uint(int32(2))%32))))
	if v155 != 0 {
		v165 = v155
		v166 = v148
		goto L29
	} else {
		goto L37
	}
L36:
	;
	goto L30
L37:
	;
	v157 = v148 + int32(1)
	if v157 != v133 {
		v148 = v157
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L27
}
