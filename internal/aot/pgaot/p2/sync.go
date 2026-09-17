package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SyncRepInitConfig(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[0])))
	if v10 != 0 {
		v101 = v1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[1]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+72))
	if v104 == v101 {
		goto L32
	} else {
		goto L33
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[2]))
	if v12 == int32(0) {
		v101 = v1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v15 == int32(0) {
		v101 = v1
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[3]))
	if v19 == int32(0) {
		v101 = v1
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v22 <= int32(0) {
		v101 = v1
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = v19 + int32(16)
	v29 = int32(1)
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[4]))
	v36 = v28
	v37 = v33
	goto L11
L8:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[3]))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)))
	if v96 != 0 {
		goto L29
	} else {
		goto L30
	}
L9:
	;
	goto L8
L10:
	;
	if v74 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v40 == v41 {
		v63 = v40
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v74 = int32(0)
	goto L10
L13:
	;
	v65 = int32(1)
	if v63 != 0 {
		v36 = v36 + v65
		v37 = v37 + v65
		goto L11
	} else {
		goto L22
	}
L14:
	;
	if base.Ui32((v40-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v51 = v40 | int32(32)
	goto L17
L16:
	;
	v51 = v40
	goto L17
L17:
	;
	if base.Ui32((v41-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v41 | int32(32)
	goto L20
L19:
	;
	v60 = v41
	goto L20
L20:
	;
	if v51 == v60 {
		v63 = v51
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v74 = v51 - v60
	goto L10
L22:
	;
	goto L12
L23:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v77 == int32(42) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v80 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v83 = F_strlen(m, v28)
	mBase = m.M
	v85 = int32(1)
	v88 = v29 + v85
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[3]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	if v88 <= v91 {
		v28 = v83 + v28 + v85
		v29 = v88
		goto L7
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v101 = v1
	goto L1
L29:
	;
	v97 = int32(1)
	goto L31
L30:
	;
	v97 = v29
	goto L31
L31:
	;
	v101 = v97
	goto L1
L32:
	;
	m.G0 = v7 + int32(16)
	return
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+76)) = int32(1)
	if v106 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[1]))
	F_s_lock(m, v110+int32(76), int32(_a_F_SyncRepInitConfig_0), int32(456), int32(_a_F_SyncRepInitConfig_1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[1]))
	v120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+76)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v119)+72)) = v101
	v125 = F_errstart(m, int32(14), v120)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L37
	} else {
		goto L39
	}
L37:
	;
	return
L38:
	;
	goto L36
L39:
	;
	if v125 == int32(0) {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v101
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_SyncRepInitConfig[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v131
	F_errmsg_internal(m, int32(_a_F_SyncRepInitConfig_2), v7)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_SyncRepInitConfig_0), int32(462), int32(_a_F_SyncRepInitConfig_1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L32
}
