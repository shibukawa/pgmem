package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_disable_timeout(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var __phi76 int32
	_ = __phi76
	var v77 int32
	_ = v77
	var __phi77 int32
	_ = __phi77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v137 int32
	_ = v137
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[0])) = v2
	v14 = l0 * int32(40)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_disable_timeout[1]))))
	if v15 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	if v20 <= int32(0) {
		v47 = int32(-1)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_disable_timeout[3]))) = uint8(v112)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	if v112 < v115 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_disable_timeout[4])))
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)) = uint8(v69)
	v72 = v24 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	if v72 < v74 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v24 = v2
	goto L8
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	if v24 < v40 {
		goto L4
	} else {
		goto L12
	}
L8:
	;
	v29 = v24 << (uint(int32(2)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_disable_timeout[4])))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 == l0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v47 = int32(-1)
	goto L5
L10:
	;
	v34 = v24 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	if v34 < v36 {
		v24 = v34
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v47 = v24
	goto L5
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v47
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v54 - int32(1)
	F_errmsg_internal(m, int32(_a_F_disable_timeout_0), v8)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_disable_timeout_1), int32(143), int32(_a_F_disable_timeout_2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	__phi76 = v72
	__phi77 = v24
	v76 = __phi76
	v77 = __phi77
	goto L20
L18:
	;
	goto L19
L19:
	;
	v99 = int32(_a_F_disable_timeout_3)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2])) = v101 - int32(1)
	goto L3
L20:
	;
	v81 = int32(2)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(v81)%32))+uint32(_c_F_disable_timeout[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(v81)%32))+uint32(_c_F_disable_timeout[4]))) = v87
	v90 = v76 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_disable_timeout[2]))
	if v90 < v92 {
		__phi76 = v90
		__phi77 = v76
		v76 = __phi76
		v77 = __phi77
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	goto L21
L23:
	;
	v121 = m.G0
	v122 = int32(16)
	v123 = v121 - v122
	m.G0 = v123
	F_gettimeofday(m, v123)
	mBase = m.M
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(v123)+8)))
	m.G0 = v123 + v122
	goto L26
L24:
	;
	goto L25
L25:
	;
	m.G0 = v8 + int32(16)
	return
L26:
	;
	F_schedule_alarm(m, v127+v126*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
