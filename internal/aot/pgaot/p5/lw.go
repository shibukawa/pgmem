package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v2 = F_LWLockDisownInternal(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_LWLockReleaseInternal(m, l0, v2)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = int32(4543676)
			v8 = *(*int32)(unsafe.Add(mBase, _consts[412]))
			*(*int32)(unsafe.Add(mBase, _consts[412])) = v8 - int32(1)
			return
		}
	}
}
func F_LWLockShmemSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v1 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if v10 <= v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v116 = int32(214)
	goto L3
L2:
	;
	v15 = v10 & int32(3)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	if base.Ui32(int32(4)) <= base.Ui32(v10) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v118 = F_mul_size(m, v116, int32(128))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v23 = v1
	v25 = int32(0)
	v26 = v1
	goto L7
L5:
	;
	v70 = v1
	v73 = v1
	goto L6
L6:
	;
	if v15 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v33 = int32(68)
	v36 = int32(-64)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v23|int32(3))*v33-v36)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v23|int32(2))*v33-v36)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v23|int32(1))*v33-v36)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17+v23*v33-v36)))
	v64 = v38 + (v46 + (v54 + (v60 + v26)))
	v65 = int32(4)
	v66 = v23 + v65
	v68 = v25 + v65
	if v68 != v10&int32(2147483644) {
		v23 = v66
		v25 = v68
		v26 = v64
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v70 = v66
	v73 = v64
	goto L6
L9:
	;
	goto L8
L10:
	;
	v78 = v70
	v81 = v73
	v84 = v1
	goto L13
L11:
	;
	v101 = v73
	goto L12
L12:
	;
	v116 = v101 + int32(214)
	goto L3
L13:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v17+v78*int32(68)-int32(-64))))
	v92 = v91 + v81
	v93 = int32(1)
	v96 = v84 + v93
	if v96 != v15 {
		v78 = v78 + v93
		v81 = v92
		v84 = v96
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v101 = v92
	goto L12
L15:
	;
	goto L14
L16:
	;
	return int32(0)
L17:
	;
	v123 = F_add_size(m, v118, int32(132))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	v128 = F_mul_size(m, v126, int32(8))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v130 = F_add_size(m, v123, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if int32(0) < v133 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v136 = v130
	v140 = v1
	goto L24
L22:
	;
	v159 = v130
	goto L23
L23:
	;
	return v159
L24:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	v149 = F_strlen(m, v145+v140*int32(68))
	mBase = m.M
	v152 = F_add_size(m, v136, v149+int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L16
	} else {
		goto L26
	}
L25:
	;
	v159 = v152
	goto L23
L26:
	;
	v155 = v140 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if v155 < v157 {
		v136 = v152
		v140 = v155
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
