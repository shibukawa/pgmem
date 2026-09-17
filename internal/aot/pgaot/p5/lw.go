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
			v6 = int32(_a_F_LWLockRelease_0)
			v8 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockRelease[0])) = v8 - int32(1)
			return
		}
	}
}
func F_LWLockShmemSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	v1 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemSize[0]))
	if v11 <= v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v98 = int32(214)
	goto L3
L2:
	;
	v16 = v11 & int32(3)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemSize[1]))
	if base.Ui32(int32(4)) <= base.Ui32(v11) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v100 = F_mul_size(m, v98, int32(128))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v98 = v80 + int32(214)
	goto L3
L5:
	;
	v23 = v1
	v25 = v1
	v30 = v1
	goto L8
L6:
	;
	v50 = v1
	v52 = v1
	goto L7
L7:
	;
	v59 = v50
	v61 = v52
	v67 = v1
	goto L12
L8:
	;
	v34 = v18 + v23*int32(68)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+268))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+200))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+132))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	v42 = v35 + (v36 + (v37 + (v38 + v25)))
	v43 = int32(4)
	v44 = v23 + v43
	v46 = v30 + v43
	if v46 != v11&int32(2147483644) {
		v23 = v44
		v25 = v42
		v30 = v46
		goto L8
	} else {
		goto L10
	}
L9:
	;
	if v16 == int32(0) {
		v80 = v42
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v50 = v44
	v52 = v42
	goto L7
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v18+v59*int32(68))+64))
	v72 = v71 + v61
	v73 = int32(1)
	v76 = v67 + v73
	if v76 != v16 {
		v59 = v59 + v73
		v61 = v72
		v67 = v76
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v80 = v72
	goto L4
L14:
	;
	goto L13
L15:
	;
	return int32(0)
L16:
	;
	v105 = F_add_size(m, v100, int32(132))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemSize[0]))
	v110 = F_mul_size(m, v108, int32(8))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v112 = F_add_size(m, v105, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemSize[0]))
	if int32(0) < v115 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v118 = v112
	v123 = v1
	goto L23
L21:
	;
	v142 = v112
	goto L22
L22:
	;
	return v142
L23:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemSize[1]))
	v132 = F_strlen(m, v128+v123*int32(68))
	mBase = m.M
	v135 = F_add_size(m, v118, v132+int32(1))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L15
	} else {
		goto L25
	}
L24:
	;
	v142 = v135
	goto L22
L25:
	;
	v138 = v123 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockShmemSize[0]))
	if v138 < v140 {
		v118 = v135
		v123 = v138
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
}
