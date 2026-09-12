package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SyncRepGetCandidateStandbys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	v12 = F_palloc(m, v9*int32(48))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v12
	v18 = *(*int32)(unsafe.Add(mBase, _consts[714]))
	if v18 == int32(0) {
		v115 = v2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v115
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)))
	if v101 != 0 {
		v115 = v97
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v95 = v18
	v97 = v2
	goto L5
L7:
	;
	goto L8
L8:
	;
	v28 = v2
	v29 = v2
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v36 = v33 + v29*int32(96)
	v38 = v36 + int32(164)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(1)
	v45 = v40 + v28*int32(48)
	v47 = v36 + int32(88)
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[714]))
	v95 = v93
	v97 = v86
	goto L5
L11:
	;
	F_s_lock(m, v38, int32(495748), int32(779), int32(113078))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v47)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v47)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v47)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+24)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = v62
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+76)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v66 == v64 {
		v86 = v28
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v88 = v29 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	if v88 < v90 {
		v28 = v86
		v29 = v88
		goto L9
	} else {
		goto L20
	}
L16:
	;
	if base.Ui32(v55-int32(5)) < base.Ui32(int32(-2)) {
		v86 = v28
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v45)+32))
	if v73 == int32(0) {
		v86 = v28
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v45)+16))
	if v76 == int64(0) {
		v86 = v28
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+36)) = v29
	v81 = *(*int32)(unsafe.Add(mBase, _consts[715]))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+40)) = uint8(base.B2i32(v47 == v81))
	v86 = v28 + int32(1)
	goto L15
L20:
	;
	goto L10
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v97 <= v102 {
		v115 = v97
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pg_qsort(m, v104, v97, int32(48), int32(1028))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[714]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v115 = v111
	goto L3
}
